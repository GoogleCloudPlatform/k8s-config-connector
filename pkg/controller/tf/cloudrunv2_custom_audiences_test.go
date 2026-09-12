// Copyright 2026 Google LLC
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

package tf_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cloudrunv2 "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/cloudrunv2"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestCloudRunV2ServiceReadPersistsCustomAudiences(t *testing.T) {
	wantAudiences := []string{
		"https://service.example.com",
		"https://service.example.net",
		"https://service.example.org",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
			http.Error(w, "unexpected method", http.StatusBadRequest)
			return
		}
		if r.URL.Path != "/v2/projects/test-project/locations/us-central1/services/test-service" {
			t.Errorf("path = %q", r.URL.Path)
			http.Error(w, "unexpected path", http.StatusBadRequest)
			return
		}
		if err := json.NewEncoder(w).Encode(map[string]any{
			"name":            "projects/test-project/locations/us-central1/services/test-service",
			"customAudiences": wantAudiences,
		}); err != nil {
			t.Fatalf("encode response: %v", err)
		}
	}))
	defer server.Close()

	resource := cloudrunv2.ResourceCloudRunV2Service()
	resourceData := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"name":     "test-service",
		"project":  "test-project",
		"location": "us-central1",
	})
	resourceData.SetId("projects/test-project/locations/us-central1/services/test-service")
	config := &transport_tpg.Config{
		Project:            "test-project",
		CloudRunV2BasePath: server.URL + "/v2/",
		Client:             server.Client(),
	}

	if err := resource.Read(resourceData, config); err != nil {
		t.Fatalf("read Cloud Run service: %v", err)
	}

	gotRaw := resourceData.Get("custom_audiences").([]interface{})
	gotAudiences := make([]string, len(gotRaw))
	for i, audience := range gotRaw {
		gotAudiences[i] = audience.(string)
	}
	if diff := cmp.Diff(wantAudiences, gotAudiences); diff != "" {
		t.Fatalf("custom_audiences mismatch (-want +got):\n%s", diff)
	}
}
