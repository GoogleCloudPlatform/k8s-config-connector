// Copyright 2026 Google LLC
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

package tf

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// Regression tests for https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/293:
// once a ComputeInstance exists, reconciling an unchanged (or label-only changed) spec must not
// report an immutable-field change or a perpetual diff, however networkIpRef / sourceImageRef are spelled.
const (
	issue293Project         = "test-project"
	issue293NetworkURL      = "https://www.googleapis.com/compute/v1/projects/test-project/global/networks/default"
	issue293AddressURL      = "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/addresses/addr"
	issue293AddressPath     = "projects/test-project/regions/us-central1/addresses/addr"
	issue293OtherAddressURL = "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/addresses/other"
	issue293AddressIP       = "10.128.0.42"
	issue293OtherAddressIP  = "10.128.0.99"
	issue293LiveCOSImage    = "https://www.googleapis.com/compute/v1/projects/cos-cloud/global/images/cos-81-12871-1196-0"
	issue293LiveDebian      = "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-11-bullseye-v20231010"
	// The boot disk image field, as named by tfresource.ImmutableFieldsFromDiff.
	issue293ImageField = "bootDisk.0.InitializeParams.0.Image"
	// The network IP field, as named by tfresource.ImmutableFieldsFromDiff.
	issue293NetworkIPField = "networkInterface.0.NetworkIp"
)

// fakeComputeAddresses answers the regional address lookups the provider makes to resolve a
// networkIpRef.external address reference. Addresses "addr" and "other" are known; everything else is a 404.
// The path match hard-codes project test-project and region us-central1, the only project and region the
// ComputeInstance fixtures in this test use.
type fakeComputeAddresses struct{}

func (fakeComputeAddresses) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == http.MethodGet && strings.Contains(req.URL.Path, "/projects/test-project/regions/us-central1/addresses/") {
		name := req.URL.Path[strings.LastIndex(req.URL.Path, "/")+1:]
		if ip := map[string]string{"addr": issue293AddressIP, "other": issue293OtherAddressIP}[name]; ip != "" {
			return jsonResponse(http.StatusOK, fmt.Sprintf(`{"name":%q,"address":%q,"addressType":"INTERNAL"}`, name, ip)), nil
		}
	}
	return jsonResponse(http.StatusNotFound, `{"error":{"code":404,"message":"not found","status":"NOT_FOUND"}}`), nil
}

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestComputeInstanceIssue293(t *testing.T) {
	ctx := context.Background()
	savedTransformer := transport_tpg.DefaultHTTPClientTransformer
	transport_tpg.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		return &http.Client{Transport: fakeComputeAddresses{}}
	}
	t.Cleanup(func() { transport_tpg.DefaultHTTPClientTransformer = savedTransformer })
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatalf("error loading service mappings: %v", err)
	}
	sm, err := smLoader.GetServiceMapping("compute.cnrm.cloud.google.com")
	if err != nil {
		t.Fatalf("error getting compute service mapping: %v", err)
	}
	crd, err := crdloader.GetCRD("compute.cnrm.cloud.google.com", "v1beta1", "ComputeInstance")
	if err != nil {
		t.Fatalf("error loading ComputeInstance CRD: %v", err)
	}
	jsonSchema := k8s.GetOpenAPIV3SchemaFromCRD(crd)

	ready := []interface{}{map[string]interface{}{"type": "Ready", "status": "True", "reason": "UpToDate"}}
	kube := fake.NewClientBuilder().WithObjects(
		&unstructured.Unstructured{Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
			"kind":       "ComputeNetwork",
			"metadata":   map[string]interface{}{"name": "default", "namespace": "ns"},
			"status":     map[string]interface{}{"selfLink": issue293NetworkURL, "conditions": ready},
		}},
		&unstructured.Unstructured{Object: map[string]interface{}{
			"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
			"kind":       "ComputeAddress",
			"metadata":   map[string]interface{}{"name": "addr", "namespace": "ns"},
			"spec":       map[string]interface{}{"location": "us-central1", "address": issue293AddressIP, "addressType": "INTERNAL"},
			"status":     map[string]interface{}{"selfLink": issue293AddressURL, "conditions": ready},
		}},
	).Build()

	const (
		debianSource = "debian-cloud/debian-11"
		cosSource    = "cos-cloud/cos-81-lts"
		cosFamilySrc = "projects/cos-cloud/global/images/family/cos-81-lts"
	)
	networkByName := map[string]interface{}{"name": "default"}
	addressIP := map[string]interface{}{"external": issue293AddressIP}
	tests := []struct {
		name         string
		networkRef   map[string]interface{}
		networkIPRef map[string]interface{}
		sourceImage  string
		liveImage    string
		// wantImmutableField is set for controls whose spec really differs from the live resource:
		// the instance must require replacement, and the immutable field must be the boot disk image.
		wantImmutableField string
	}{
		{name: "networkIpRef by name", networkRef: networkByName, networkIPRef: map[string]interface{}{"kind": "ComputeAddress", "name": "addr"},
			sourceImage: debianSource, liveImage: issue293LiveDebian},
		{name: "networkIpRef.external IP", networkRef: networkByName, networkIPRef: addressIP,
			sourceImage: debianSource, liveImage: issue293LiveDebian},
		// Real GCP returns the network as a full URL, so the live state holds a full URL here. The address
		// selfLink resolves to the live IP, so the provider treats network_ip as unchanged and the spec is a no-op.
		{name: "networkIpRef.external address selfLink", networkRef: networkByName, networkIPRef: map[string]interface{}{"external": issue293AddressURL},
			sourceImage: debianSource, liveImage: issue293LiveDebian},
		{name: "networkIpRef.external address path", networkRef: networkByName, networkIPRef: map[string]interface{}{"external": issue293AddressPath},
			sourceImage: debianSource, liveImage: issue293LiveDebian},
		// Same address as above, but networkRef.external is a relative path, as the mock stores it. The address
		// reference is still resolved to the live IP, so the network_ip diff is cleared rather than persisting.
		{name: "networkIpRef.external address selfLink, networkRef.external path", networkRef: map[string]interface{}{"external": "projects/test-project/global/networks/default"}, networkIPRef: map[string]interface{}{"external": issue293AddressURL},
			sourceImage: debianSource, liveImage: issue293LiveDebian},
		// The address selfLink resolves to a different IP than the live one, so the change is real and the
		// instance must be replaced (the network_ip stays immutable).
		{name: "control: networkIpRef.external different address selfLink", networkRef: networkByName, networkIPRef: map[string]interface{}{"external": issue293OtherAddressURL},
			sourceImage: debianSource, liveImage: issue293LiveDebian, wantImmutableField: issue293NetworkIPField},
		{name: "sourceImageRef COS family shorthand", networkRef: networkByName, networkIPRef: addressIP,
			sourceImage: cosSource, liveImage: issue293LiveCOSImage},
		{name: "sourceImageRef COS family path", networkRef: networkByName, networkIPRef: addressIP,
			sourceImage: cosFamilySrc, liveImage: issue293LiveCOSImage},
		{name: "control: sourceImageRef different COS family", networkRef: networkByName, networkIPRef: addressIP,
			sourceImage: "cos-cloud/cos-85-lts", liveImage: issue293LiveCOSImage, wantImmutableField: issue293ImageField},
	}
	for _, tc := range tests {
		newResource := func(t *testing.T, labelValue string) *krmtotf.Resource {
			u := &unstructured.Unstructured{Object: map[string]interface{}{
				"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
				"kind":       "ComputeInstance",
				"metadata": map[string]interface{}{
					"name": "inst", "namespace": "ns",
					"annotations": map[string]interface{}{"cnrm.cloud.google.com/project-id": issue293Project},
					"labels":      map[string]interface{}{"env": labelValue},
				},
				"spec": map[string]interface{}{
					"zone":        "us-central1-a",
					"machineType": "n1-standard-1",
					"bootDisk": map[string]interface{}{"initializeParams": map[string]interface{}{
						"sourceImageRef": map[string]interface{}{"external": tc.sourceImage}}},
					"networkInterface": []interface{}{map[string]interface{}{
						"networkRef":   tc.networkRef,
						"networkIpRef": tc.networkIPRef,
					}},
				},
			}}
			r, err := krmtotf.NewResource(u, sm, provider)
			if err != nil {
				t.Fatalf("error building resource: %v", err)
			}
			if err := resourceoverrides.Handler.PreActuationTransform(&r.Resource); err != nil {
				t.Fatalf("error applying pre-actuation transform: %v", err)
			}
			return r
		}
		// liveState is what Terraform reads back after GCE created the instance:
		// networkIP is always an IP, and the boot image is the concrete image the family resolved to.
		liveState := func(r *krmtotf.Resource) *terraform.InstanceState {
			state := krmtotf.MapToInstanceState(r.TFResource, map[string]interface{}{
				"name": "inst", "zone": "us-central1-a", "project": issue293Project, "machine_type": "n1-standard-1",
				"labels":              map[string]interface{}{"env": "a", "managed-by-cnrm": "true"},
				"can_ip_forward":      false,
				"deletion_protection": false,
				"guest_accelerator":   []interface{}{},
				"scheduling": []interface{}{map[string]interface{}{
					"automatic_restart": true, "on_host_maintenance": "MIGRATE", "preemptible": false, "provisioning_model": "STANDARD",
				}},
				"reservation_affinity":         []interface{}{},
				"confidential_instance_config": []interface{}{},
				"network_interface": []interface{}{map[string]interface{}{
					"network":    issue293NetworkURL,
					"network_ip": issue293AddressIP,
				}},
				"boot_disk": []interface{}{map[string]interface{}{
					"auto_delete": true, "mode": "READ_WRITE", "device_name": "persistent-disk-0",
					"source": "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/disks/inst",
					"initialize_params": []interface{}{map[string]interface{}{
						"image": tc.liveImage, "labels": map[string]interface{}{}, "size": 10, "type": "pd-standard",
					}},
				}},
			})
			state.ID = "projects/test-project/zones/us-central1-a/instances/inst"
			return state
		}
		diffAgainstLive := func(t *testing.T, labelValue string) *terraform.InstanceDiff {
			r := newResource(t, labelValue)
			live := liveState(r)
			config, _, err := krmtotf.KRMResourceToTFResourceConfigFull(r, kube, smLoader, live, jsonSchema, true)
			if err != nil {
				t.Fatalf("error converting to TF config: %v", err)
			}
			diff, err := r.TFResource.Diff(ctx, live, config, provider.Meta())
			if err != nil {
				t.Fatalf("error calculating diff: %v", err)
			}
			return diff
		}

		t.Run(tc.name, func(t *testing.T) {
			t.Run("unchanged spec", func(t *testing.T) {
				diff := diffAgainstLive(t, "a")
				if tc.wantImmutableField != "" {
					if !diff.RequiresNew() {
						t.Fatalf("expected an immutable-field change, got diff %v", changedAttributes(diff))
					}
					immutable := tfresource.ImmutableFieldsFromDiff(diff)
					if !containsField(immutable, tc.wantImmutableField) {
						t.Fatalf("expected immutable field %q, got %v", tc.wantImmutableField, immutable)
					}
					return
				}
				if diff.RequiresNew() {
					t.Fatalf("re-reconciling an unchanged spec fails with: %v", k8s.NewImmutableFieldsMutationError(tfresource.ImmutableFieldsFromDiff(diff)))
				}
				if got := changedAttributes(diff); len(got) != 0 {
					t.Fatalf("re-reconciling an unchanged spec is not a no-op (perpetual diff): %v", got)
				}
			})
			if tc.wantImmutableField != "" {
				return
			}
			t.Run("label update", func(t *testing.T) {
				diff := diffAgainstLive(t, "b")
				if diff.RequiresNew() {
					t.Fatalf("label update fails with: %v", k8s.NewImmutableFieldsMutationError(tfresource.ImmutableFieldsFromDiff(diff)))
				}
				if got, want := changedAttributes(diff), []string{"labels.env: a -> b"}; !equalStrings(got, want) {
					t.Fatalf("label update diff = %v, want %v", got, want)
				}
			})
		})
	}
}

func changedAttributes(diff *terraform.InstanceDiff) []string {
	var changed []string
	if diff == nil {
		return changed
	}
	for k, a := range diff.Attributes {
		if a.Old != a.New || a.RequiresNew || a.NewRemoved {
			changed = append(changed, k+": "+a.Old+" -> "+a.New)
		}
	}
	sort.Strings(changed)
	return changed
}

// containsField reports whether one of the ImmutableFieldsFromDiff entries names the given field.
func containsField(entries []string, field string) bool {
	for _, e := range entries {
		if strings.HasPrefix(e, "Field Name: "+field+",") {
			return true
		}
	}
	return false
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
