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

package iptables

import (
	"net/netip"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"

	// Create a new mount namespace.
	"github.com/howardjohn/unshare-go/mountns"

	// Create a new network namespace. This will have the 'lo' interface ready but nothing else.
	_ "github.com/howardjohn/unshare-go/netns"
	"github.com/howardjohn/unshare-go/userns"

	"istio.io/istio/cni/pkg/ipset"
	"istio.io/istio/cni/pkg/scopes"
	istiolog "istio.io/istio/pkg/log"
	"istio.io/istio/pkg/test/util/assert"
	iptablescapture "istio.io/istio/tools/istio-iptables/pkg/capture"
	dep "istio.io/istio/tools/istio-iptables/pkg/dependencies"
)

func createHostsideProbeIpset(isV6 bool) (ipset.IPSet, error) {
	linDeps := ipset.RealNlDeps()
	probeSet, err := ipset.NewIPSet(ProbeIPSet, isV6, linDeps)
	if err != nil {
		return probeSet, err
	}
	probeSet.Flush()
	return probeSet, nil
}

func TestIdempotentEquivalentInPodRerun(t *testing.T) {
	setup(t)

	tests := []struct {
		name        string
		config      func(cfg *Config)
		ingressMode bool
	}{
		{
			name: "default",
			config: func(cfg *Config) {
				cfg.RedirectDNS = true
			},
			ingressMode: false,
		},
		{
			name: "tproxy",
			config: func(cfg *Config) {
				cfg.TPROXYRedirection = true
				cfg.RedirectDNS = true
			},
			ingressMode: false,
		},
		{
			name: "ingress",
			config: func(cfg *Config) {
			},
			ingressMode: true,
		},
	}

	probeSNATipv4 := netip.MustParseAddr("169.254.7.127")
	probeSNATipv6 := netip.MustParseAddr("e9ac:1e77:90ca:399f:4d6d:ece2:2f9b:3164")
	ext := &dep.RealDependencies{
		HostFilesystemPodNetwork: false,
		NetworkNamespace:         "",
	}
	iptVer, err := ext.DetectIptablesVersion(false)
	if err != nil {
		t.Fatalf("Can't detect iptables version: %v", err)
	}

	ipt6Ver, err := ext.DetectIptablesVersion(true)
	if err != nil {
		t.Fatalf("Can't detect ip6tables version")
	}
	scope := istiolog.FindScope(istiolog.DefaultScopeName)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{}
			tt.config(cfg)

			deps := &dep.RealDependencies{}
			iptConfigurator, _, err := NewIptablesConfigurator(cfg, cfg, deps, deps, EmptyNlDeps())
			builder := iptConfigurator.AppendInpodRules(probeSNATipv4, probeSNATipv6, tt.ingressMode)
			if err != nil {
				t.Fatalf("failed to setup iptables configurator: %v", err)
			}
			defer func() {
				assert.NoError(t, iptConfigurator.DeleteInpodRules())
				residueExists, deltaExists := iptablescapture.VerifyIptablesState(scope, iptConfigurator.ext, builder, &iptVer, &ipt6Ver)
				assert.Equal(t, residueExists, false)
				assert.Equal(t, deltaExists, true)
			}()
			assert.NoError(t, iptConfigurator.CreateInpodRules(scopes.CNIAgent, probeSNATipv4, probeSNATipv6, tt.ingressMode))
			residueExists, deltaExists := iptablescapture.VerifyIptablesState(scope, iptConfigurator.ext, builder, &iptVer, &ipt6Ver)
			assert.Equal(t, residueExists, true)
			assert.Equal(t, deltaExists, false)

			t.Log("Starting cleanup")
			// Cleanup, should work
			assert.NoError(t, iptConfigurator.DeleteInpodRules())
			residueExists, deltaExists = iptablescapture.VerifyIptablesState(scope, iptConfigurator.ext, builder, &iptVer, &ipt6Ver)
			assert.Equal(t, residueExists, false)
			assert.Equal(t, deltaExists, true)

			t.Log("Second run")
			// Add again, should still work
			assert.NoError(t, iptConfigurator.CreateInpodRules(scopes.CNIAgent, probeSNATipv4, probeSNATipv6, false))
			residueExists, deltaExists = iptablescapture.VerifyIptablesState(scope, iptConfigurator.ext, builder, &iptVer, &ipt6Ver)
			assert.Equal(t, residueExists, true)
			assert.Equal(t, deltaExists, false)

			t.Log("Third run")
			assert.NoError(t, iptConfigurator.CreateInpodRules(scopes.CNIAgent, probeSNATipv4, probeSNATipv6, false))
		})
	}
}

func TestIptablesHostCleanRoundTrip(t *testing.T) {
	setup(t)

	tests := []struct {
		name   string
		config func(cfg *Config)
	}{
		{
			name: "default",
			config: func(cfg *Config) {
				cfg.RedirectDNS = true
			},
		},
	}

	probeSNATipv4 := netip.MustParseAddr("169.254.7.127")
	probeSNATipv6 := netip.MustParseAddr("e9ac:1e77:90ca:399f:4d6d:ece2:2f9b:3164")
	ext := &dep.RealDependencies{
		HostFilesystemPodNetwork: false,
		NetworkNamespace:         "",
	}
	iptVer, err := ext.DetectIptablesVersion(false)
	if err != nil {
		t.Fatalf("Can't detect iptables version: %v", err)
	}

	ipt6Ver, err := ext.DetectIptablesVersion(true)
	if err != nil {
		t.Fatalf("Can't detect ip6tables version")
	}
	scope := istiolog.FindScope(istiolog.DefaultScopeName)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{}
			tt.config(cfg)

			deps := &dep.RealDependencies{}
			set, err := createHostsideProbeIpset(true)
			if err != nil {
				t.Fatalf("failed to create hostside probe ipset: %v", err)
			}
			defer func() {
				assert.NoError(t, set.DestroySet())
			}()

			iptConfigurator, _, err := NewIptablesConfigurator(cfg, cfg, deps, deps, RealNlDeps())
			builder := iptConfigurator.AppendHostRules(&probeSNATipv4, &probeSNATipv6)
			if err != nil {
				t.Fatalf("failed to setup iptables configurator: %v", err)
			}
			defer func() {
				iptConfigurator.DeleteHostRules(probeSNATipv4, probeSNATipv6)
				residueExists, deltaExists := iptablescapture.VerifyIptablesState(scope, iptConfigurator.ext, builder, &iptVer, &ipt6Ver)
				assert.Equal(t, residueExists, false)
				assert.Equal(t, deltaExists, true)
			}()

			assert.NoError(t, iptConfigurator.CreateHostRulesForHealthChecks(&probeSNATipv4, &probeSNATipv6))
			residueExists, deltaExists := iptablescapture.VerifyIptablesState(scope, iptConfigurator.ext, builder, &iptVer, &ipt6Ver)
			assert.Equal(t, residueExists, true)
			assert.Equal(t, deltaExists, false)

			// Round-trip deletion and recreation to test clean-up and re-setup
			t.Log("Starting cleanup")
			iptConfigurator.DeleteHostRules(probeSNATipv4, probeSNATipv6)
			residueExists, deltaExists = iptablescapture.VerifyIptablesState(scope, iptConfigurator.ext, builder, &iptVer, &ipt6Ver)
			assert.Equal(t, residueExists, false)
			assert.Equal(t, deltaExists, true)

			t.Log("Second run")
			assert.NoError(t, iptConfigurator.CreateHostRulesForHealthChecks(&probeSNATipv4, &probeSNATipv6))
			residueExists, deltaExists = iptablescapture.VerifyIptablesState(scope, iptConfigurator.ext, builder, &iptVer, &ipt6Ver)
			assert.Equal(t, residueExists, true)
			assert.Equal(t, deltaExists, false)

			t.Log("Third run")
			assert.NoError(t, iptConfigurator.CreateHostRulesForHealthChecks(&probeSNATipv4, &probeSNATipv6))
		})
	}
}

var initialized = &sync.Once{}

func setup(t *testing.T) {
	initialized.Do(func() {
		// Setup group namespace so iptables --gid-owner will work
		assert.NoError(t, userns.WriteGroupMap(map[uint32]uint32{userns.OriginalGID(): 0}))
		// Istio iptables expects to find a non-localhost IP in some interface
		assert.NoError(t, exec.Command("ip", "addr", "add", "240.240.240.240/32", "dev", "lo").Run())
		// Put a new file we have permission to access over xtables.lock
		xtables := filepath.Join(t.TempDir(), "xtables.lock")
		_, err := os.Create(xtables)
		assert.NoError(t, err)
		_ = os.Mkdir("/run", 0o777)
		_ = mountns.BindMount(xtables, "/run/xtables.lock")
	})
}

func iptablesSave(t *testing.T) string {
	res, err := exec.Command("iptables-save").CombinedOutput()
	assert.NoError(t, err)
	return string(res)
}
