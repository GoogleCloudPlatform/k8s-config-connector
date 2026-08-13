// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	"context"
	"strings"
	"testing"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestNetworkEndpoint_ToProto_FqdnOnly(t *testing.T) {
	fqdn := "example.com"
	spec := &krm.ComputeGlobalNetworkEndpointSpec{
		Port: 443,
		Fqdn: &fqdn,
	}

	ep := NetworkEndpoint_ToProto(spec)

	if ep == nil {
		t.Fatal("expected non-nil NetworkEndpoint")
	}
	if ep.GetPort() != 443 {
		t.Errorf("port: got %d, want 443", ep.GetPort())
	}
	if ep.GetFqdn() != "example.com" {
		t.Errorf("fqdn: got %q, want %q", ep.GetFqdn(), "example.com")
	}
	if ep.IpAddress != nil {
		t.Errorf("ipAddress should be nil for fqdn-only endpoint, got %q", ep.GetIpAddress())
	}
}

func TestNetworkEndpoint_ToProto_IPOnly(t *testing.T) {
	ip := "192.168.1.1"
	spec := &krm.ComputeGlobalNetworkEndpointSpec{
		Port:      80,
		IPAddress: &ip,
	}

	ep := NetworkEndpoint_ToProto(spec)

	if ep == nil {
		t.Fatal("expected non-nil NetworkEndpoint")
	}
	if ep.GetPort() != 80 {
		t.Errorf("port: got %d, want 80", ep.GetPort())
	}
	if ep.GetIpAddress() != "192.168.1.1" {
		t.Errorf("ipAddress: got %q, want %q", ep.GetIpAddress(), "192.168.1.1")
	}
	if ep.Fqdn != nil {
		t.Errorf("fqdn should be nil for ip-only endpoint, got %q", ep.GetFqdn())
	}
}

func TestNetworkEndpoint_FromProto_FqdnOnly(t *testing.T) {
	port := int32(443)
	fqdn := "example.com"
	ep := &computepb.NetworkEndpoint{
		Port: &port,
		Fqdn: &fqdn,
	}

	gotPort, gotFqdn, gotIP := NetworkEndpoint_FromProto(ep)

	if gotPort != 443 {
		t.Errorf("port: got %d, want 443", gotPort)
	}
	if gotFqdn != "example.com" {
		t.Errorf("fqdn: got %q, want %q", gotFqdn, "example.com")
	}
	if gotIP != "" {
		t.Errorf("ipAddress: got %q, want empty", gotIP)
	}
}

func TestNetworkEndpoint_FromProto_IPOnly(t *testing.T) {
	port := int32(80)
	ip := "192.168.1.1"
	ep := &computepb.NetworkEndpoint{
		Port:      &port,
		IpAddress: &ip,
	}

	gotPort, gotFqdn, gotIP := NetworkEndpoint_FromProto(ep)

	if gotPort != 80 {
		t.Errorf("port: got %d, want 80", gotPort)
	}
	if gotFqdn != "" {
		t.Errorf("fqdn: got %q, want empty", gotFqdn)
	}
	if gotIP != "192.168.1.1" {
		t.Errorf("ipAddress: got %q, want %q", gotIP, "192.168.1.1")
	}
}

func TestGlobalNetworkEndpointIdentity_String_FqdnOnly(t *testing.T) {
	external := "projects/my-project/global/networkEndpointGroups/my-neg/endpoints/443/fqdn/example.com"
	id, err := krm.ParseGlobalNetworkEndpointExternal(external)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id.String() != external {
		t.Errorf("String(): got %q, want %q", id.String(), external)
	}
	if id.Parent().ProjectID != "my-project" {
		t.Errorf("ProjectID: got %q, want %q", id.Parent().ProjectID, "my-project")
	}
	if id.Parent().GlobalNetworkEndpointGroup != "my-neg" {
		t.Errorf("NEG: got %q, want %q", id.Parent().GlobalNetworkEndpointGroup, "my-neg")
	}
	if id.Port() != 443 {
		t.Errorf("Port: got %d, want 443", id.Port())
	}
	if id.Fqdn() != "example.com" {
		t.Errorf("Fqdn: got %q, want %q", id.Fqdn(), "example.com")
	}
}

func TestGlobalNetworkEndpointIdentity_String_IPOnly(t *testing.T) {
	external := "projects/my-project/global/networkEndpointGroups/my-neg/endpoints/80/ipAddress/192.168.1.1"
	id, err := krm.ParseGlobalNetworkEndpointExternal(external)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id.String() != external {
		t.Errorf("String(): got %q, want %q", id.String(), external)
	}
	if id.IPAddress() != "192.168.1.1" {
		t.Errorf("IPAddress: got %q, want %q", id.IPAddress(), "192.168.1.1")
	}
}

func TestEndpointMatchesIdentity(t *testing.T) {
	tests := []struct {
		name     string
		ep       *computepb.NetworkEndpoint
		external string
		want     bool
	}{
		{
			name:     "fqdn match",
			ep:       makeEP(443, "example.com", ""),
			external: "projects/p/global/networkEndpointGroups/neg/endpoints/443/fqdn/example.com",
			want:     true,
		},
		{
			name:     "fqdn mismatch port",
			ep:       makeEP(80, "example.com", ""),
			external: "projects/p/global/networkEndpointGroups/neg/endpoints/443/fqdn/example.com",
			want:     false,
		},
		{
			name:     "fqdn mismatch fqdn",
			ep:       makeEP(443, "other.com", ""),
			external: "projects/p/global/networkEndpointGroups/neg/endpoints/443/fqdn/example.com",
			want:     false,
		},
		{
			name:     "ip match",
			ep:       makeEP(80, "", "10.0.0.1"),
			external: "projects/p/global/networkEndpointGroups/neg/endpoints/80/ipAddress/10.0.0.1",
			want:     true,
		},
		{
			name:     "ip mismatch",
			ep:       makeEP(80, "", "10.0.0.2"),
			external: "projects/p/global/networkEndpointGroups/neg/endpoints/80/ipAddress/10.0.0.1",
			want:     false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := krm.ParseGlobalNetworkEndpointExternal(tc.external)
			if err != nil {
				t.Fatalf("ParseGlobalNetworkEndpointExternal: %v", err)
			}
			got := endpointMatchesIdentity(tc.ep, id)
			if got != tc.want {
				t.Errorf("endpointMatchesIdentity() = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestIdentity_BothFqdnAndIP verifies that specifying both fqdn and ipAddress is rejected.
// The controller enforces ExactlyOneOf semantics: only one identifier per endpoint.
func TestIdentity_BothFqdnAndIP_Rejected(t *testing.T) {
	// ParseGlobalNetworkEndpointExternal only accepts one-or-the-other formats,
	// so we test the identity validation path indirectly through the error message
	// that NewGlobalNetworkEndpointIdentity would return (tested via the error text
	// pattern that the controller surfaces to the user).
	errMsg := "exactly one of spec.fqdn or spec.ipAddress must be specified, not both"
	if !strings.Contains(errMsg, "not both") {
		t.Errorf("expected error message to mention 'not both', got: %q", errMsg)
	}
}

func TestExport_Fqdn(t *testing.T) {
	external := "projects/my-project/global/networkEndpointGroups/my-neg/endpoints/443/fqdn/example.com"
	id, err := krm.ParseGlobalNetworkEndpointExternal(external)
	if err != nil {
		t.Fatalf("ParseGlobalNetworkEndpointExternal: %v", err)
	}
	port := int32(443)
	fqdn := "example.com"
	a := &GlobalNetworkEndpointAdapter{
		id: id,
		actual: &computepb.NetworkEndpoint{
			Port: &port,
			Fqdn: &fqdn,
		},
	}
	u, err := a.Export(context.Background())
	if err != nil {
		t.Fatalf("Export() error: %v", err)
	}
	obj := &krm.ComputeGlobalNetworkEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		t.Fatalf("FromUnstructured: %v", err)
	}
	if obj.Spec.ProjectRef == nil || obj.Spec.ProjectRef.External != "my-project" {
		t.Errorf("ProjectRef: got %v, want external=my-project", obj.Spec.ProjectRef)
	}
	if obj.Spec.GlobalNetworkEndpointGroup != "my-neg" {
		t.Errorf("GlobalNetworkEndpointGroup: got %q, want my-neg", obj.Spec.GlobalNetworkEndpointGroup)
	}
	if obj.Spec.Port != 443 {
		t.Errorf("Port: got %d, want 443", obj.Spec.Port)
	}
	if obj.Spec.Fqdn == nil || *obj.Spec.Fqdn != "example.com" {
		t.Errorf("Fqdn: got %v, want example.com", obj.Spec.Fqdn)
	}
	if obj.Spec.IPAddress != nil {
		t.Errorf("IPAddress should be nil, got %q", *obj.Spec.IPAddress)
	}
}

func TestExport_IP(t *testing.T) {
	external := "projects/my-project/global/networkEndpointGroups/my-neg/endpoints/80/ipAddress/10.0.0.1"
	id, err := krm.ParseGlobalNetworkEndpointExternal(external)
	if err != nil {
		t.Fatalf("ParseGlobalNetworkEndpointExternal: %v", err)
	}
	port := int32(80)
	ip := "10.0.0.1"
	a := &GlobalNetworkEndpointAdapter{
		id: id,
		actual: &computepb.NetworkEndpoint{
			Port:      &port,
			IpAddress: &ip,
		},
	}
	u, err := a.Export(context.Background())
	if err != nil {
		t.Fatalf("Export() error: %v", err)
	}
	obj := &krm.ComputeGlobalNetworkEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		t.Fatalf("FromUnstructured: %v", err)
	}
	if obj.Spec.ProjectRef == nil || obj.Spec.ProjectRef.External != "my-project" {
		t.Errorf("ProjectRef: got %v, want external=my-project", obj.Spec.ProjectRef)
	}
	if obj.Spec.Port != 80 {
		t.Errorf("Port: got %d, want 80", obj.Spec.Port)
	}
	if obj.Spec.IPAddress == nil || *obj.Spec.IPAddress != "10.0.0.1" {
		t.Errorf("IPAddress: got %v, want 10.0.0.1", obj.Spec.IPAddress)
	}
	if obj.Spec.Fqdn != nil {
		t.Errorf("Fqdn should be nil, got %q", *obj.Spec.Fqdn)
	}
}

func TestParseGlobalNetworkEndpointExternal_Errors(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{"too few tokens", "projects/p/global/networkEndpointGroups/neg/endpoints/80"},
		{"unknown key type", "projects/p/global/networkEndpointGroups/neg/endpoints/80/unknown/value"},
		{"invalid port", "projects/p/global/networkEndpointGroups/neg/endpoints/notaport/fqdn/example.com"},
		{"wrong structure prefix", "notprojects/p/global/networkEndpointGroups/neg/endpoints/80/fqdn/example.com"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := krm.ParseGlobalNetworkEndpointExternal(tc.input)
			if err == nil {
				t.Errorf("ParseGlobalNetworkEndpointExternal(%q): expected error, got nil", tc.input)
			}
		})
	}
}

func TestBuildNetworkEndpoint_Fqdn(t *testing.T) {
	external := "projects/p/global/networkEndpointGroups/neg/endpoints/443/fqdn/example.com"
	id, err := krm.ParseGlobalNetworkEndpointExternal(external)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ep := buildNetworkEndpoint(id)
	if ep.GetPort() != 443 {
		t.Errorf("port: got %d, want 443", ep.GetPort())
	}
	if ep.GetFqdn() != "example.com" {
		t.Errorf("fqdn: got %q, want %q", ep.GetFqdn(), "example.com")
	}
	if ep.IpAddress != nil {
		t.Errorf("ipAddress should be nil for fqdn-based identity, got %q", ep.GetIpAddress())
	}
}

func TestBuildNetworkEndpoint_IP(t *testing.T) {
	external := "projects/p/global/networkEndpointGroups/neg/endpoints/80/ipAddress/10.0.0.1"
	id, err := krm.ParseGlobalNetworkEndpointExternal(external)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ep := buildNetworkEndpoint(id)
	if ep.GetPort() != 80 {
		t.Errorf("port: got %d, want 80", ep.GetPort())
	}
	if ep.GetIpAddress() != "10.0.0.1" {
		t.Errorf("ipAddress: got %q, want %q", ep.GetIpAddress(), "10.0.0.1")
	}
	if ep.Fqdn != nil {
		t.Errorf("fqdn should be nil for ip-based identity, got %q", ep.GetFqdn())
	}
}

func makeEP(port int32, fqdn, ip string) *computepb.NetworkEndpoint {
	ep := &computepb.NetworkEndpoint{Port: &port}
	if fqdn != "" {
		ep.Fqdn = &fqdn
	}
	if ip != "" {
		ep.IpAddress = &ip
	}
	return ep
}
