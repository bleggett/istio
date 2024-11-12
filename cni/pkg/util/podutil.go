// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"context"
	"fmt"
	"net/netip"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"

	"istio.io/api/annotation"
	"istio.io/api/label"
	"istio.io/istio/pkg/config/constants"
	"istio.io/istio/pkg/log"
)

var annotationPatch = []byte(fmt.Sprintf(
	`{"metadata":{"annotations":{"%s":"%s"}}}`,
	annotation.AmbientRedirection.Name,
	constants.AmbientRedirectionEnabled,
))

var annotationRemovePatch = []byte(fmt.Sprintf(
	`{"metadata":{"annotations":{"%s":null}}}`,
	annotation.AmbientRedirection.Name,
))

// PodRedirectionEnabled determines if a pod should or should not be configured
// to have traffic redirected thru the node proxy.
//
// NOTE that this is NOT checking if a pod WAS enabled previously, for that see PodRedirectionActive.
// This is just checking if it SHOULD be enabled under current state.
func PodRedirectionEnabled(namespace *corev1.Namespace, pod *corev1.Pod, autoEnrollEnabled bool, excludeNamespaces []string) bool {
	// Check for user opt-out - do the namespace and/or pod have the explicit opt-out labels?
	hasExplicitOptOut := (namespace.GetLabels()[label.IoIstioDataplaneMode.Name] == constants.DataplaneModeNone ||
		pod.GetLabels()[label.IoIstioDataplaneMode.Name] == constants.DataplaneModeNone)

	// Check for user opt-in - do the namespace and/or pod have the explicit opt-in labels?
	hasExplicitOptIn := (namespace.GetLabels()[label.IoIstioDataplaneMode.Name] == constants.DataplaneModeAmbient ||
		pod.GetLabels()[label.IoIstioDataplaneMode.Name] == constants.DataplaneModeAmbient)

	// Always unconditionally respect explicit opt-out (pod or namespace) if present.
	if hasExplicitOptOut {
		// Pod or namespace explicitly asked to not have ambient redirection enabled
		return false
	}

	// If no autoenroll and no explicit opt-in, pod should be ignored.
	if !autoEnrollEnabled && !hasExplicitOptIn {
		return false
	}

	// If autoenroll enabled and pod does NOT have an explicit opt-in, see
	// if it's in our list of excluded namespaces.
	//
	// This is designed to allow people to always explicitly opt-in pods with labels,
	// even if they are in a namespace we have been told not to autoenroll.
	// Basically, explicit labels always win, even for autoenroll-ignored namespaces.
	if autoEnrollEnabled && !hasExplicitOptIn {
		// If pod in our autoenroll excluded namespace list (shared with the CNI plugin), it can never be enabled.
		for _, ns := range excludeNamespaces {
			if ns == namespace.Name {
				log.Debugf("excluding %s/%s from autoenroll due to excluded namespace: %s", namespace.Name, pod.Name, ns)
				return false
			}
		}
	}

	// If we get here, pod is either auto-enrollable OR has an explicit opt-in.

	// Ztunnel and sidecar for a single pod is currently not supported; opt out.
	if podHasSidecar(pod) {
		return false
	}

	// Host network pods cannot be captured, as we require inserting rules into the pod network namespace.
	// If we were to allow them, we would be writing these rules into the host network namespace, effectively breaking the host.
	if pod.Spec.HostNetwork {
		return false
	}

	return true
}

// PodRedirectionActive reports on whether the pod _has_ actually been configured for traffic redirection.
//
// That is, have we annotated it after successfully sending it to the node proxy and set up iptables rules.
//
// If you just want to know if the pod _should be_ configured for traffic redirection, see PodRedirectionEnabled
func PodRedirectionActive(pod *corev1.Pod) bool {
	return pod.GetAnnotations()[annotation.AmbientRedirection.Name] == constants.AmbientRedirectionEnabled
}

func podHasSidecar(pod *corev1.Pod) bool {
	if _, f := pod.GetAnnotations()[annotation.SidecarStatus.Name]; f {
		return true
	}
	return false
}

func IsZtunnelPod(systemNs string, pod *corev1.Pod) bool {
	return pod.Namespace == systemNs && pod.GetLabels()["app"] == "ztunnel"
}

func AnnotateEnrolledPod(client kubernetes.Interface, pod *metav1.ObjectMeta) error {
	_, err := client.CoreV1().
		Pods(pod.Namespace).
		Patch(
			context.Background(),
			pod.Name,
			types.MergePatchType,
			annotationPatch,
			metav1.PatchOptions{},
			// Both "pods" and "pods/status" can mutate the metadata. However, pods/status is lower privilege, so we use that instead.
			"status",
		)
	return err
}

func AnnotateUnenrollPod(client kubernetes.Interface, pod *metav1.ObjectMeta) error {
	if pod.Annotations[annotation.AmbientRedirection.Name] != constants.AmbientRedirectionEnabled {
		return nil
	}
	// TODO: do not overwrite if already none
	_, err := client.CoreV1().
		Pods(pod.Namespace).
		Patch(
			context.Background(),
			pod.Name,
			types.MergePatchType,
			annotationRemovePatch,
			metav1.PatchOptions{},
			// Both "pods" and "pods/status" can mutate the metadata. However, pods/status is lower privilege, so we use that instead.
			"status",
		)
	if errors.IsNotFound(err) {
		return nil
	}
	return err
}

// Get any IPs currently assigned to the Pod.
//
// If 'PodIPs' exists, it is preferred (and should be guaranteed to contain the address in 'PodIP'),
// otherwise fallback to 'PodIP'.
//
// Note that very early in the pod's lifecycle (before all the node CNI plugin invocations finish)
// K8S may not have received the pod IPs yet, and may not report the pod as having any.
func GetPodIPsIfPresent(pod *corev1.Pod) []netip.Addr {
	var podIPs []netip.Addr
	if len(pod.Status.PodIPs) != 0 {
		for _, pip := range pod.Status.PodIPs {
			ip := netip.MustParseAddr(pip.IP)
			podIPs = append(podIPs, ip)
		}
	} else if len(pod.Status.PodIP) != 0 {
		ip := netip.MustParseAddr(pod.Status.PodIP)
		podIPs = append(podIPs, ip)
	}
	return podIPs
}
