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

package tfidentity

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/identity"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/services/resourcemanager"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Folder IAM reads go through a hand-written client that sends requests
// without a context, and carry no project. The resolver must still decide the
// identity, through the provider copy.
func TestFolderIAMReadUsesResolvedIdentity(t *testing.T) {
	var gotPath, gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(`{"etag":"BwY="}`)); err != nil {
			t.Errorf("writing IAM response: %v", err)
		}
	}))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	provider, base := newTestProvider(ctx)
	base.ResourceManagerV3BasePath = server.URL + "/v3/"

	member := &unstructured.Unstructured{Object: map[string]interface{}{
		"spec": map[string]interface{}{
			"resourceRef": map[string]interface{}{"kind": "Folder", "external": "folders/130065489"},
		},
	}}
	member.SetNamespace("team-a")
	member.SetName("folder-viewer")
	ctx, err := identity.ForObject(identity.WithResolver(ctx, byNamespace), identity.Target{
		GroupKind: schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicyMember"},
		Object:    member,
	})
	if err != nil {
		t.Fatalf("ForObject: %v", err)
	}
	p, err := Provider(ctx, provider)
	if err != nil {
		t.Fatalf("Provider: %v", err)
	}

	if _, err := resourcemanager.GetFolderIamPolicyByFolderName("folders/130065489", "kcc/test", p.Meta().(*transport_tpg.Config)); err != nil {
		t.Fatalf("GetFolderIamPolicyByFolderName: %v", err)
	}
	if want := "/folders/130065489:getIamPolicy"; !strings.HasSuffix(gotPath, want) {
		t.Errorf("path = %q, want suffix %q", gotPath, want)
	}
	if want := "Bearer team-a"; gotAuth != want {
		t.Errorf("Authorization = %q, want %q", gotAuth, want)
	}
}
