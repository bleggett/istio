//go:build integ
// +build integ

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

package helm

import (
	"fmt"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/types"

	"istio.io/istio/pkg/test/framework"
	kubecluster "istio.io/istio/pkg/test/framework/components/cluster/kube"
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/pkg/test/helm"
	"istio.io/istio/tests/util/sanitycheck"
)

// TestDefaultInstall tests Istio installation using Helm with default options
func TestDefaultInstall(t *testing.T) {
	overrideValuesStr := `
global:
  hub: %s
  %s
  variant: %q
`
	framework.
		NewTest(t).
		Run(setupInstallation(false, DefaultNamespaceConfig, ""))
}

// TestAmbientInstall tests Istio ambient profile installation using Helm
func TestAmbientInstall(t *testing.T) {
	framework.
		NewTest(t).
		Run(setupInstallation(ambientProfileOverride, true, DefaultNamespaceConfig, ""))
}

func TestAmbientInstallMultiNamespace(t *testing.T) {
	tests := []struct {
		name     string
		nsConfig NamespaceConfig
	}{{
		name: "isolated-istio-cni",
		nsConfig: NewNamespaceConfig(types.NamespacedName{
			Name: CniReleaseName, Namespace: "istio-cni",
		}),
	}, {
		name: "isolated-istio-cni-and-ztunnel",
		nsConfig: NewNamespaceConfig(types.NamespacedName{
			Name: CniReleaseName, Namespace: "istio-cni",
		}, types.NamespacedName{
			Name: ZtunnelReleaseName, Namespace: "kube-system",
		}),
	}, {
		name: "isolated-istio-cni-ztunnel-and-gateway",
		nsConfig: NewNamespaceConfig(types.NamespacedName{
			Name: CniReleaseName, Namespace: "istio-cni",
		}, types.NamespacedName{
			Name: ZtunnelReleaseName, Namespace: "ztunnel",
		}, types.NamespacedName{
			Name: IngressReleaseName, Namespace: "ingress-release",
		}),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			framework.
				NewTest(t).
				Run(setupInstallation(ambientProfileOverride, true, tt.nsConfig, ""))
		})
	}
}

// TestReleaseChannels tests that non-stable CRDs and fields get blocked
// by the default ValidatingAdmissionPolicy
func TestReleaseChannels(t *testing.T) {
	overrideValuesStr := `
global:
  hub: %s
  %s
  variant: %q
profile: stable
`

	framework.
		NewTest(t).
		RequireKubernetesMinorVersion(30).
		Run(setupInstallationWithCustomCheck(overrideValuesStr, false, DefaultNamespaceConfig, func(t framework.TestContext) {
			// Try to apply an EnvoyFilter (it should be rejected)
			expectedErrorPrefix := `%s "sample" is forbidden: ValidatingAdmissionPolicy 'stable-channel-default-policy.istio.io' ` +
				`with binding 'stable-channel-default-policy-binding.istio.io' denied request`
			err := t.ConfigIstio().Eval("default", nil, sampleEnvoyFilter).Apply()
			if err == nil {
				t.Errorf("Did not receive an error while applying sample EnvoyFilter with stable admission policy")
			} else {
				msg := fmt.Sprintf(expectedErrorPrefix, "envoyfilters.networking.istio.io")
				if !strings.Contains(err.Error(), msg) {
					t.Errorf("Expected error %q to contain %q", err.Error(), msg)
				}
			}

			// Now test field-level blocks with Telemetry
			err = t.ConfigIstio().Eval("default", nil, extendedTelemetry).Apply()
			if err == nil {
				t.Error("Did not receive an error while applying extended Telemetry resource with stable admission policy")
			} else {
				msg := fmt.Sprintf(expectedErrorPrefix, "telemetries.telemetry.istio.io")
				if !strings.Contains(err.Error(), msg) {
					t.Errorf("Expected error %q to contain %q", err.Error(), msg)
				}
			}
		}, ""))
}

// TestRevisionedReleaseChannels tests that non-stable CRDs and fields get blocked
// by the revisioned ValidatingAdmissionPolicy
func TestRevisionedReleaseChannels(t *testing.T) {
	overrideValuesStr := `
global:
  hub: %s
  %s
  variant: %q
profile: stable
revision: 1-x
defaultRevision: ""
`
	revision := "1-x"
	framework.
		NewTest(t).
		RequireKubernetesMinorVersion(30).
		Run(setupInstallationWithCustomCheck(overrideValuesStr, false, DefaultNamespaceConfig, func(t framework.TestContext) {
			// Try to apply an EnvoyFilter (it should be rejected)
			expectedErrorPrefix := `%s "sample" is forbidden: ValidatingAdmissionPolicy 'stable-channel-policy-1-x-istio-system.istio.io' ` +
				`with binding 'stable-channel-policy-binding-1-x-istio-system.istio.io' denied request`
			err := t.ConfigIstio().Eval("default", nil, fmt.Sprintf(revisionedSampleEnvoyFilter, revision)).Apply()
			if err == nil {
				t.Errorf("Did not receive an error while applying sample EnvoyFilter with stable admission policy")
			} else {
				msg := fmt.Sprintf(expectedErrorPrefix, "envoyfilters.networking.istio.io")
				if !strings.Contains(err.Error(), msg) {
					t.Errorf("Expected error %q to contain %q", err.Error(), msg)
				}
			}

			// Now test field-level blocks with Telemetry
			err = t.ConfigIstio().Eval("default", nil, fmt.Sprintf(revisionedExtendedTelemetry, revision)).Apply()
			if err == nil {
				t.Error("Did not receive an error while applying extended Telemetry resource with stable admission policy")
			} else {
				msg := fmt.Sprintf(expectedErrorPrefix, "telemetries.telemetry.istio.io")
				if !strings.Contains(err.Error(), msg) {
					t.Errorf("Expected error %q to contain %q", err.Error(), msg)
				}
			}
		}, revision))
}

func setupInstallation(isAmbient bool, config NamespaceConfig, revision string) func(t framework.TestContext) {
	return baseSetup(overrideValuesStr, isAmbient, config, func(t framework.TestContext) {
		sanitycheck.RunTrafficTest(t, t)
	}, revision)
}

func setupInstallationWithCustomCheck(isAmbient bool, config NamespaceConfig,
	check func(t framework.TestContext), revision string,
) func(t framework.TestContext) {
	return baseSetup(isAmbient, config, check, revision)
}

func baseSetup(isAmbient bool, config NamespaceConfig,
	check func(t framework.TestContext), revision string,
) func(t framework.TestContext) {
	return func(t framework.TestContext) {
		cs := t.Clusters().Default().(*kubecluster.Cluster)
		h := helm.New(cs.Filename())
		s := t.Settings()

		overrideValuesFile := GetValuesOverrides(t, s.Image.Hub, s.Image.Tag, s.Image.Variant, revision, isAmbient)
		t.Cleanup(func() {
			if !t.Failed() {
				return
			}
			if t.Settings().CIMode {
				for _, ns := range config.AllNamespaces() {
					namespace.Dump(t, ns)
				}
			}
		})

		InstallIstio(t, cs, h, overrideValuesFile, "", true, isAmbient, config)

		VerifyInstallation(t, cs, config, true, isAmbient, revision)
		verifyValidation(t, revision)

		check(t)
		t.Cleanup(func() {
			DeleteIstio(t, h, cs, config, isAmbient)
		})
	}
}
