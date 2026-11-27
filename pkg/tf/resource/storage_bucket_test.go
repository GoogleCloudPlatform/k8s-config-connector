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

package resource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	tpgstorage "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/storage"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/storage/v1"
)

func TestStorageBucketEncryptionEnforcement(t *testing.T) {
	const key = "projects/test-project/locations/us/keyRings/test/cryptoKeys/test"
	for _, tc := range []struct {
		name       string
		initialKey string
		encryption map[string]interface{}
		wantKey    string
		wantMode   string
	}{
		{
			name:     "acquire with encryption omitted",
			wantMode: "FullyRestricted",
		},
		{
			name:       "acquire with key and encryption omitted",
			initialKey: key,
			wantKey:    key,
			wantMode:   "FullyRestricted",
		},
		{
			name:       "update key with enforcement omitted",
			initialKey: key,
			encryption: map[string]interface{}{"default_kms_key_name": key + "-updated"},
			wantKey:    key + "-updated",
			wantMode:   "FullyRestricted",
		},
		{
			name:       "update enforcement with key omitted",
			initialKey: key,
			encryption: map[string]interface{}{
				"customer_supplied_encryption_enforcement_config": []interface{}{
					map[string]interface{}{"restriction_mode": "NotRestricted"},
				},
			},
			wantKey:  key,
			wantMode: "NotRestricted",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bucket := storage.Bucket{
				Name: "test-bucket", Id: "test-bucket", Location: "US", StorageClass: "STANDARD",
				Encryption: &storage.BucketEncryption{
					DefaultKmsKeyName: tc.initialKey,
					CustomerSuppliedEncryptionEnforcementConfig: &storage.BucketEncryptionCustomerSuppliedEncryptionEnforcementConfig{
						RestrictionMode: "FullyRestricted",
					},
					CustomerManagedEncryptionEnforcementConfig: &storage.BucketEncryptionCustomerManagedEncryptionEnforcementConfig{
						RestrictionMode: "NotRestricted",
					},
					GoogleManagedEncryptionEnforcementConfig: &storage.BucketEncryptionGoogleManagedEncryptionEnforcementConfig{
						RestrictionMode: "NotRestricted",
					},
				},
			}
			var patches int
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch r.Method {
				case http.MethodGet:
				case http.MethodPatch:
					patches++
					var patch map[string]json.RawMessage
					if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
						t.Errorf("decoding patch: %v", err)
					}
					if raw, ok := patch["encryption"]; ok {
						bucket.Encryption = nil
						if err := json.Unmarshal(raw, &bucket.Encryption); err != nil {
							t.Errorf("decoding encryption: %v", err)
						}
					}
					if raw, ok := patch["labels"]; ok {
						if err := json.Unmarshal(raw, &bucket.Labels); err != nil {
							t.Errorf("decoding labels: %v", err)
						}
					}
				default:
					t.Errorf("unexpected request: %s %s", r.Method, r.URL)
				}
				if err := json.NewEncoder(w).Encode(bucket); err != nil {
					t.Errorf("encoding bucket: %v", err)
				}
			}))
			defer server.Close()
			config := &transport_tpg.Config{
				Context: context.Background(), Client: server.Client(), StorageBasePath: server.URL + "/",
				Project: "test-project",
			}
			resource := tpgstorage.ResourceStorageBucket()
			state, diagnostics := resource.RefreshWithoutUpgrade(context.Background(), &terraform.InstanceState{
				ID: "test-bucket", Attributes: map[string]string{"name": "test-bucket"},
			}, config)
			if diagnostics.HasError() {
				t.Fatalf("reading existing bucket: %v", diagnostics)
			}
			raw := map[string]interface{}{
				"name": "test-bucket", "location": "US",
				"labels": map[string]interface{}{"managed-by-cnrm": "true"},
			}
			if tc.encryption != nil {
				raw["encryption"] = []interface{}{tc.encryption}
			}
			diff, err := resource.Diff(context.Background(), state, terraform.NewResourceConfigRaw(raw), config)
			if err != nil {
				t.Fatalf("planning update: %v", err)
			}
			state, diagnostics = resource.Apply(context.Background(), state, diff, config)
			if diagnostics.HasError() {
				t.Fatalf("updating bucket: %v", diagnostics)
			}
			if patches != 1 {
				t.Fatalf("got %d PATCH requests, want 1", patches)
			}
			if bucket.Encryption == nil {
				t.Fatal("update cleared encryption")
			}
			if got := bucket.Encryption.DefaultKmsKeyName; got != tc.wantKey {
				t.Errorf("default KMS key = %q, want %q", got, tc.wantKey)
			}
			if got := bucket.Encryption.CustomerSuppliedEncryptionEnforcementConfig; got == nil || got.RestrictionMode != tc.wantMode {
				t.Errorf("CSEK enforcement = %+v, want %s", got, tc.wantMode)
			}
			if got := bucket.Encryption.CustomerManagedEncryptionEnforcementConfig; got == nil || got.RestrictionMode != "NotRestricted" {
				t.Errorf("CMEK enforcement was changed: %+v", got)
			}
			if got := bucket.Encryption.GoogleManagedEncryptionEnforcementConfig; got == nil || got.RestrictionMode != "NotRestricted" {
				t.Errorf("GMEK enforcement was changed: %+v", got)
			}
			state, diagnostics = resource.RefreshWithoutUpgrade(context.Background(), state, config)
			if diagnostics.HasError() {
				t.Fatalf("refreshing updated bucket: %v", diagnostics)
			}
			diff, err = resource.Diff(context.Background(), state, terraform.NewResourceConfigRaw(raw), config)
			if err != nil {
				t.Fatalf("planning second update: %v", err)
			}
			if !diff.Empty() {
				t.Errorf("unexpected diff after update: %v", diff)
			}
		})
	}
}

func TestStorageBucketCreateEncryptionEnforcement(t *testing.T) {
	for tfField, apiField := range map[string]string{
		"customer_managed_encryption_enforcement_config":  "customerManagedEncryptionEnforcementConfig",
		"customer_supplied_encryption_enforcement_config": "customerSuppliedEncryptionEnforcementConfig",
		"google_managed_encryption_enforcement_config":    "googleManagedEncryptionEnforcementConfig",
	} {
		for _, mode := range []string{"NotRestricted", "FullyRestricted"} {
			t.Run(tfField+"/"+mode, func(t *testing.T) {
				var bucket storage.Bucket
				var inserts int
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					switch r.Method {
					case http.MethodPost:
						inserts++
						var body map[string]interface{}
						if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
							t.Errorf("decoding insert: %v", err)
						}
						enc, ok := body["encryption"].(map[string]interface{})
						if !ok {
							t.Error("insert omitted encryption")
						} else {
							cfg, ok := enc[apiField].(map[string]interface{})
							if !ok || cfg["restrictionMode"] != mode {
								t.Errorf("insert encryption = %v, want %s = %s", enc, apiField, mode)
							}
							if _, ok := enc["defaultKmsKeyName"]; ok {
								t.Error("insert included an undeclared default KMS key")
							}
						}
						encoded, err := json.Marshal(body)
						if err != nil {
							t.Errorf("encoding insert: %v", err)
						}
						if err := json.Unmarshal(encoded, &bucket); err != nil {
							t.Errorf("decoding bucket: %v", err)
						}
						bucket.Id = bucket.Name
					case http.MethodGet:
					default:
						t.Errorf("unexpected request: %s %s", r.Method, r.URL)
					}
					if err := json.NewEncoder(w).Encode(bucket); err != nil {
						t.Errorf("encoding bucket: %v", err)
					}
				}))
				defer server.Close()
				config := &transport_tpg.Config{
					Context: context.Background(), Client: server.Client(), StorageBasePath: server.URL + "/",
					Project: "test-project",
				}
				raw := map[string]interface{}{
					"name": "test-bucket", "location": "US",
					"encryption": []interface{}{map[string]interface{}{
						tfField: []interface{}{map[string]interface{}{"restriction_mode": mode}},
					}},
				}
				resource := tpgstorage.ResourceStorageBucket()
				if diagnostics := resource.Validate(terraform.NewResourceConfigRaw(raw)); diagnostics.HasError() {
					t.Fatalf("validating configuration: %v", diagnostics)
				}
				data := schema.TestResourceDataRaw(t, resource.Schema, raw)
				if err := resource.Create(data, config); err != nil {
					t.Fatalf("creating bucket: %v", err)
				}
				if inserts != 1 {
					t.Errorf("got %d insert requests, want 1", inserts)
				}
				if got := data.Get("encryption.0." + tfField + ".0.restriction_mode"); got != mode {
					t.Errorf("read restriction mode = %v, want %s", got, mode)
				}
			})
		}
	}
}

func TestStorageBucketEncryptionEnforcementValidation(t *testing.T) {
	resource := tpgstorage.ResourceStorageBucket()
	if err := resource.InternalValidate(nil, true); err != nil {
		t.Fatalf("invalid resource schema: %v", err)
	}
	for _, field := range []string{
		"customer_managed_encryption_enforcement_config",
		"customer_supplied_encryption_enforcement_config",
		"google_managed_encryption_enforcement_config",
	} {
		for _, mode := range []string{"", "fullyrestricted", "invalid"} {
			t.Run(field+"/"+mode, func(t *testing.T) {
				raw := map[string]interface{}{
					"name": "test-bucket", "location": "US",
					"encryption": []interface{}{map[string]interface{}{
						field: []interface{}{map[string]interface{}{"restriction_mode": mode}},
					}},
				}
				if diagnostics := resource.Validate(terraform.NewResourceConfigRaw(raw)); !diagnostics.HasError() {
					t.Errorf("accepted invalid restriction mode %q", mode)
				}
			})
		}
	}
}
