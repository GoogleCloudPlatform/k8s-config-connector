// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestComputeInstance_networkIPCustomizedDiff(t *testing.T) {
	t.Parallel()

	d := &tpgresource.ResourceDiffMock{
		Before: map[string]interface{}{
			"network_interface.#": 0,
		},
		After: map[string]interface{}{
			"network_interface.#": 1,
		},
	}

	err := forceNewIfNetworkIPNotUpdatableFunc(d, failIfCalled(t))
	if err != nil {
		t.Error(err)
	}

	if d.IsForceNew {
		t.Errorf("Expected not force new if network_interface array size changes")
	}

	type NetworkInterface struct {
		Network           string
		Subnetwork        string
		SubnetworkProject string
		NetworkIP         string
	}
	NIBefore := NetworkInterface{
		Network:           "a",
		Subnetwork:        "a",
		SubnetworkProject: "a",
		NetworkIP:         "a",
	}

	cases := map[string]struct {
		ExpectedForceNew bool
		Before           NetworkInterface
		After            NetworkInterface
	}{
		"NetworkIP only change": {
			ExpectedForceNew: true,
			Before:           NIBefore,
			After: NetworkInterface{
				Network:           "a",
				Subnetwork:        "a",
				SubnetworkProject: "a",
				NetworkIP:         "b",
			},
		},
		"NetworkIP and Network change": {
			ExpectedForceNew: false,
			Before:           NIBefore,
			After: NetworkInterface{
				Network:           "b",
				Subnetwork:        "a",
				SubnetworkProject: "a",
				NetworkIP:         "b",
			},
		},
		"NetworkIP and Subnetwork change": {
			ExpectedForceNew: false,
			Before:           NIBefore,
			After: NetworkInterface{
				Network:           "a",
				Subnetwork:        "b",
				SubnetworkProject: "a",
				NetworkIP:         "b",
			},
		},
		"NetworkIP and SubnetworkProject change": {
			ExpectedForceNew: false,
			Before:           NIBefore,
			After: NetworkInterface{
				Network:           "a",
				Subnetwork:        "a",
				SubnetworkProject: "b",
				NetworkIP:         "b",
			},
		},
		"All change": {
			ExpectedForceNew: false,
			Before:           NIBefore,
			After: NetworkInterface{
				Network:           "b",
				Subnetwork:        "b",
				SubnetworkProject: "b",
				NetworkIP:         "b",
			},
		},
		"No change": {
			ExpectedForceNew: false,
			Before:           NIBefore,
			After: NetworkInterface{
				Network:           "a",
				Subnetwork:        "a",
				SubnetworkProject: "a",
				NetworkIP:         "a",
			},
		},
	}

	for tn, tc := range cases {
		d := &tpgresource.ResourceDiffMock{
			Before: map[string]interface{}{
				"network_interface.#":                    1,
				"network_interface.0.network":            tc.Before.Network,
				"network_interface.0.subnetwork":         tc.Before.Subnetwork,
				"network_interface.0.subnetwork_project": tc.Before.SubnetworkProject,
				"network_interface.0.network_ip":         tc.Before.NetworkIP,
			},
			After: map[string]interface{}{
				"network_interface.#":                    1,
				"network_interface.0.network":            tc.After.Network,
				"network_interface.0.subnetwork":         tc.After.Subnetwork,
				"network_interface.0.subnetwork_project": tc.After.SubnetworkProject,
				"network_interface.0.network_ip":         tc.After.NetworkIP,
			},
		}
		err := forceNewIfNetworkIPNotUpdatableFunc(d, failIfCalled(t))
		if err != nil {
			t.Error(err)
		}
		if tc.ExpectedForceNew != d.IsForceNew {
			t.Errorf("%v: expected d.IsForceNew to be %v, but was %v", tn, tc.ExpectedForceNew, d.IsForceNew)
		}
	}
}

const (
	networkIPTestAddrIP   = "10.128.0.42"
	networkIPTestOtherIP  = "10.128.0.99"
	networkIPTestAddrURL  = "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/addresses/addr"
	networkIPTestOtherURL = "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/addresses/other"
)

// failIfCalled returns a resolver that fails the test if an address lookup is attempted.
func failIfCalled(t *testing.T) networkIPAddressResolver {
	t.Helper()
	return func(project, region, name string) (string, error) {
		t.Errorf("unexpected address lookup for projects/%s/regions/%s/addresses/%s", project, region, name)
		return "", nil
	}
}

// fakeAddressResolver resolves addresses from a fixed name-to-IP map and records every lookup.
type fakeAddressResolver struct {
	err   error
	calls []string
}

func (f *fakeAddressResolver) resolve(project, region, name string) (string, error) {
	f.calls = append(f.calls, project+"/"+region+"/"+name)
	if f.err != nil {
		return "", f.err
	}
	return map[string]string{"addr": networkIPTestAddrIP, "other": networkIPTestOtherIP, "v6": "2600:1900:0:0:0:0:0:1"}[name], nil
}

// The diff cases below use ResourceDiffMock, whose Clear only records the key and does not enforce the SDK's
// rules for clearing a diff. The real Clear path is covered by TestComputeInstanceIssue293 in pkg/controller/tf.
func TestComputeInstance_networkIPAddressRefCustomizedDiff(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		OldNetworkIP           string
		NewNetworkIP           string
		NetworkChanges         bool
		ResolverErr            error
		ExpectedResolverCalls  int
		ExpectedCleared        bool
		ExpectedForceNew       bool
		ExpectedErrContainsAll []string
		NetworkBefore          string
		NetworkAfter           string
	}{
		"same IP, network spelled differently": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          networkIPTestAddrURL,
			NetworkBefore:         "https://www.googleapis.com/compute/v1/projects/test-project/global/networks/default",
			NetworkAfter:          "projects/test-project/global/networks/default",
			ExpectedResolverCalls: 1,
			ExpectedCleared:       true,
		},
		"same IP, full v1 selfLink": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          networkIPTestAddrURL,
			ExpectedResolverCalls: 1,
			ExpectedCleared:       true,
		},
		"same IP, relative path": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          "projects/test-project/regions/us-central1/addresses/addr",
			ExpectedResolverCalls: 1,
			ExpectedCleared:       true,
		},
		"same IP, compute.googleapis.com beta URL": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          "https://compute.googleapis.com/compute/beta/projects/test-project/regions/us-central1/addresses/addr",
			ExpectedResolverCalls: 1,
			ExpectedCleared:       true,
		},
		"IPv6 address, different spelling of the same IP": {
			OldNetworkIP:          "2600:1900::1",
			NewNetworkIP:          "projects/test-project/regions/us-central1/addresses/v6",
			ExpectedResolverCalls: 1,
			ExpectedCleared:       true,
		},
		"different IP": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          networkIPTestOtherURL,
			ExpectedResolverCalls: 1,
			ExpectedForceNew:      true,
		},
		// A network change moves the instance in place and the update patches NetworkIP with the address reference,
		// so the lookup is skipped and the existing logic (no ForceNew) applies.
		"different IP and network also changes": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          networkIPTestOtherURL,
			NetworkChanges:        true,
			ExpectedResolverCalls: 0,
		},
		"same IP and network also changes": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          networkIPTestAddrURL,
			NetworkChanges:        true,
			ExpectedResolverCalls: 0,
		},
		"create, network and network_ip both unset before": {
			OldNetworkIP:          "",
			NewNetworkIP:          networkIPTestAddrURL,
			NetworkChanges:        true,
			ExpectedResolverCalls: 0,
		},
		"bare IP to bare IP": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          "10.128.0.43",
			ExpectedResolverCalls: 0,
			ExpectedForceNew:      true,
		},
		"lookup error names the reference and key": {
			OldNetworkIP:           networkIPTestAddrIP,
			NewNetworkIP:           networkIPTestAddrURL,
			ResolverErr:            errors.New("boom"),
			ExpectedResolverCalls:  1,
			ExpectedErrContainsAll: []string{networkIPTestAddrURL, "network_interface.0.network_ip", "boom"},
		},
		"global address path is not a regional reference": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          "projects/test-project/global/addresses/addr",
			ExpectedResolverCalls: 0,
			ExpectedForceNew:      true,
		},
		"regional path without projects prefix is not a reference": {
			OldNetworkIP:          networkIPTestAddrIP,
			NewNetworkIP:          "regions/us-central1/addresses/addr",
			ExpectedResolverCalls: 0,
			ExpectedForceNew:      true,
		},
	}

	for tn, tc := range cases {
		networkBefore := "a"
		if tc.NetworkBefore != "" {
			networkBefore = tc.NetworkBefore
		}
		networkAfter := networkBefore
		if tc.NetworkChanges {
			networkAfter = "b"
		}
		if tc.NetworkAfter != "" {
			networkAfter = tc.NetworkAfter
		}
		d := &tpgresource.ResourceDiffMock{
			Before: map[string]interface{}{
				"network_interface.#":                    1,
				"network_interface.0.network":            networkBefore,
				"network_interface.0.subnetwork":         "a",
				"network_interface.0.subnetwork_project": "a",
				"network_interface.0.network_ip":         tc.OldNetworkIP,
			},
			After: map[string]interface{}{
				"network_interface.#":                    1,
				"network_interface.0.network":            networkAfter,
				"network_interface.0.subnetwork":         "a",
				"network_interface.0.subnetwork_project": "a",
				"network_interface.0.network_ip":         tc.NewNetworkIP,
			},
		}
		resolver := &fakeAddressResolver{err: tc.ResolverErr}
		err := forceNewIfNetworkIPNotUpdatableFunc(d, resolver.resolve)
		if got := len(resolver.calls); got != tc.ExpectedResolverCalls {
			t.Errorf("%v: expected %d address lookups, got %d (%v)", tn, tc.ExpectedResolverCalls, got, resolver.calls)
		}
		_, cleared := d.Cleared["network_interface.0.network_ip"]
		if cleared != tc.ExpectedCleared {
			t.Errorf("%v: expected network_ip cleared to be %v, but was %v", tn, tc.ExpectedCleared, cleared)
		}
		if len(tc.ExpectedErrContainsAll) > 0 {
			if err == nil {
				t.Fatalf("%v: expected an error", tn)
			}
			for _, want := range tc.ExpectedErrContainsAll {
				if !strings.Contains(err.Error(), want) {
					t.Errorf("%v: expected error %q to contain %q", tn, err.Error(), want)
				}
			}
			continue
		}
		if err != nil {
			t.Fatalf("%v: unexpected error: %v", tn, err)
		}
		if tc.ExpectedForceNew != d.IsForceNew {
			t.Errorf("%v: expected d.IsForceNew to be %v, but was %v", tn, tc.ExpectedForceNew, d.IsForceNew)
		}
	}
}

func TestComputeInstance_networkIPAddressRefMatchesIP(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		OldIP                 string
		NewValue              string
		ExpectedResolverCalls int
		ExpectedMatch         bool
	}{
		"full v1 URL, same IP": {
			OldIP:                 networkIPTestAddrIP,
			NewValue:              networkIPTestAddrURL,
			ExpectedResolverCalls: 1,
			ExpectedMatch:         true,
		},
		"relative path, same IP": {
			OldIP:                 networkIPTestAddrIP,
			NewValue:              "projects/test-project/regions/us-central1/addresses/addr",
			ExpectedResolverCalls: 1,
			ExpectedMatch:         true,
		},
		"full v1 URL, different IP": {
			OldIP:                 networkIPTestAddrIP,
			NewValue:              networkIPTestOtherURL,
			ExpectedResolverCalls: 1,
			ExpectedMatch:         false,
		},
		"IPv6 spelled differently, same IP": {
			OldIP:                 "2600:1900::1",
			NewValue:              "projects/test-project/regions/us-central1/addresses/v6",
			ExpectedResolverCalls: 1,
			ExpectedMatch:         true,
		},
		"old value is not an IP": {
			OldIP:                 "addr",
			NewValue:              networkIPTestAddrURL,
			ExpectedResolverCalls: 0,
		},
		"old value is empty": {
			OldIP:                 "",
			NewValue:              networkIPTestAddrURL,
			ExpectedResolverCalls: 0,
		},
		"global address path": {
			OldIP:                 networkIPTestAddrIP,
			NewValue:              "projects/test-project/global/addresses/addr",
			ExpectedResolverCalls: 0,
		},
		"prefix before projects is not a path segment": {
			OldIP:                 networkIPTestAddrIP,
			NewValue:              "xprojects/test-project/regions/us-central1/addresses/addr",
			ExpectedResolverCalls: 0,
		},
		"trailing slash is not a reference": {
			OldIP:                 networkIPTestAddrIP,
			NewValue:              "projects/test-project/regions/us-central1/addresses/addr/",
			ExpectedResolverCalls: 0,
		},
	}

	for tn, tc := range cases {
		resolver := &fakeAddressResolver{}
		match, err := networkIPAddressRefMatchesIP(tc.OldIP, tc.NewValue, resolver.resolve)
		if err != nil {
			t.Errorf("%v: unexpected error: %v", tn, err)
		}
		if got := len(resolver.calls); got != tc.ExpectedResolverCalls {
			t.Errorf("%v: expected %d address lookups, got %d", tn, tc.ExpectedResolverCalls, got)
		}
		if match != tc.ExpectedMatch {
			t.Errorf("%v: expected match to be %v, but was %v", tn, tc.ExpectedMatch, match)
		}
	}
}

func TestComputeInstance_networkIPAddressResolver(t *testing.T) {
	t.Parallel()

	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"address":"10.0.0.5"}`))
	}))
	defer srv.Close()

	newConfig := func(ctx context.Context) *transport_tpg.Config {
		return &transport_tpg.Config{
			Client:          srv.Client(),
			Context:         ctx,
			ComputeBasePath: srv.URL + "/compute/beta/",
			UserAgent:       "test",
		}
	}

	// A live context resolves the address with a GET to addresses.get.
	resolve := newNetworkIPAddressResolver(context.Background(), newConfig(context.Background()))
	ip, err := resolve("p", "r", "a")
	if err != nil {
		t.Fatalf("live context: unexpected error: %v", err)
	}
	if ip != "10.0.0.5" {
		t.Errorf("live context: expected IP 10.0.0.5, got %q", ip)
	}
	if want := "/compute/beta/projects/p/regions/r/addresses/a"; gotPath != want {
		t.Errorf("live context: expected request path %q, got %q", want, gotPath)
	}

	// The lookup must use the context passed to the resolver, so a cancelled context aborts it.
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	resolve = newNetworkIPAddressResolver(cancelledCtx, newConfig(cancelledCtx))
	if _, err := resolve("p", "r", "a"); !errors.Is(err, context.Canceled) {
		t.Errorf("cancelled context: expected error wrapping context.Canceled, got %v", err)
	}

	// Provider meta of the wrong type is reported without attempting a lookup.
	resolve = newNetworkIPAddressResolver(context.Background(), "not-a-config")
	if _, err := resolve("p", "r", "a"); err == nil || !strings.Contains(err.Error(), "provider meta has type string") {
		t.Errorf("wrong meta type: expected error containing %q, got %v", "provider meta has type string", err)
	}
}
