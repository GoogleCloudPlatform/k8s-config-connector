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

package resourceoverrides_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	secretmanager "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/secretmanager"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSecretManagerSecretCreateTimeStatus(t *testing.T) {
	loader, err := servicemappingloader.New()
	if err != nil {
		t.Fatal(err)
	}
	mapping, err := loader.GetServiceMapping("secretmanager.cnrm.cloud.google.com")
	if err != nil {
		t.Fatal(err)
	}
	object := &unstructured.Unstructured{Object: map[string]interface{}{
		"apiVersion": "secretmanager.cnrm.cloud.google.com/v1beta1",
		"kind":       "SecretManagerSecret",
		"metadata":   map[string]interface{}{"name": "test-secret", "namespace": "test-project"},
		"spec":       map[string]interface{}{"replication": map[string]interface{}{"automatic": true}},
	}}
	providerResource := secretmanager.ResourceSecretManagerSecret()
	resource, err := krmtotf.NewResource(object, mapping, &tfschema.Provider{ResourcesMap: map[string]*tfschema.Resource{
		"google_secret_manager_secret": providerResource,
	}})
	if err != nil {
		t.Fatal(err)
	}
	const createTime = "2026-10-04T20:17:50Z"
	const name = "projects/test-project/secrets/test-secret"
	state := krmtotf.MapToInstanceState(providerResource, map[string]interface{}{
		"project": "test-project", "secret_id": "test-secret", "name": name,
		"create_time": createTime,
		"replication": []interface{}{map[string]interface{}{"automatic": true}},
	})
	state.ID = name
	resource.Spec, resource.Status = krmtotf.GetSpecAndStatusFromState(resource, state)
	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource, state, nil); err != nil {
		t.Fatal(err)
	}
	if _, found := resource.Status["createTime"]; found {
		t.Fatal("Terraform status still contains the undeclared top-level createTime")
	}
	got, found, err := unstructured.NestedString(resource.Status, "observedState", "createTime")
	if err != nil || !found || got != createTime {
		t.Fatalf("observedState.createTime = %q, found=%v, err=%v; want %q", got, found, err, createTime)
	}
	if resource.Status["name"] != name {
		t.Fatalf("legacy status.name changed: got %v, want %q", resource.Status["name"], name)
	}
}
