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

package preview

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestGenerateAlternativeReconcilerOverride(t *testing.T) {
	tests := []struct {
		name          string
		config        resourceconfig.ResourcesControllerMap
		expectedMap   map[string]string
		expectedError bool
	}{
		{
			name: "single supported controller, same as default",
			config: resourceconfig.ResourcesControllerMap{
				schema.GroupKind{Group: "example.cnrm.cloud.google.com", Kind: "Single"}: {
					DefaultController:    k8s.ReconcilerTypeTerraform,
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerTypeTerraform},
				},
			},
			expectedMap: map[string]string{},
		},
		{
			name: "multiple supported controllers, different default",
			config: resourceconfig.ResourcesControllerMap{
				schema.GroupKind{Group: "example.cnrm.cloud.google.com", Kind: "Multi"}: {
					DefaultController:    k8s.ReconcilerTypeTerraform,
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerTypeDirect, k8s.ReconcilerTypeTerraform},
				},
			},
			expectedMap: map[string]string{
				"Multi.example.cnrm.cloud.google.com": "direct",
			},
		},
		{
			name: "multiple supported controllers, default is direct",
			config: resourceconfig.ResourcesControllerMap{
				schema.GroupKind{Group: "example.cnrm.cloud.google.com", Kind: "MultiDirect"}: {
					DefaultController:    k8s.ReconcilerTypeDirect,
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerTypeDirect, k8s.ReconcilerTypeTerraform},
				},
			},
			expectedMap: map[string]string{
				"MultiDirect.example.cnrm.cloud.google.com": "tf",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := generateAlternativeReconcilerOverride(tc.config)
			if (err != nil) != tc.expectedError {
				t.Fatalf("expected error: %v, got: %v", tc.expectedError, err)
			}
			if len(result) != len(tc.expectedMap) {
				t.Errorf("expected %d overrides, got %d", len(tc.expectedMap), len(result))
			}
			for k, v := range tc.expectedMap {
				if gotVal, ok := result[k]; !ok || gotVal != v {
					t.Errorf("expected override for %s to be %s, got %s", k, v, gotVal)
				}
			}
		})
	}
}

func TestGenerateFilename(t *testing.T) {
	prefix := "test-prefix"
	filename := generateFilename(prefix)

	// Check prefix
	if !strings.HasPrefix(filename, prefix+"-") {
		t.Errorf("expected filename to start with %q, but got %q", prefix+"-", filename)
	}

	// Check format "YYYYMMDD-HHMMSS.milliseconds" which is 25 chars.
	// 20060102-150405.000
	// 4+2+2 + 1 + 2+2+2 + 1 + 3 = 19 ? No
	// 2006 (4) 01 (2) 02 (2) - (1) 15 (2) 04 (2) 05 (2) . (1) 000 (3)
	// Total length of timestamp part: 4+2+2 + 1 + 2+2+2 + 1 + 3 = 19
	// 17 characters is what I see in `20060102-150405.000`? No.
	// 20060102 is 8.
	// - is 1.
	// 150405 is 6.
	// . is 1.
	// 000 is 3.
	// 8 + 1 + 6 + 1 + 3 = 19.

	timePart := strings.TrimPrefix(filename, prefix+"-")
	if len(timePart) != 19 {
		t.Errorf("expected timestamp part length to be 19, got %d (part: %q)", len(timePart), timePart)
	}

	// Regex check for format
	match, err := regexp.MatchString(`^\d{8}-\d{6}\.\d{3}$`, timePart)
	if err != nil {
		t.Fatalf("regex error: %v", err)
	}
	if !match {
		t.Errorf("timestamp part %q does not match expected format YYYYMMDD-HHMMSS.mmm", timePart)
	}

	// Optional: Check if time is recent (within a second)
	// This helps catch if we somehow generate a stale time or something weird, but might be flaky if test runner is super slow.
	// Given we are testing the format mostly, the above is enough, but parsing it back is a good sanity check.

	_, err = time.Parse("20060102-150405.000", timePart)
	if err != nil {
		t.Errorf("failed to parse timestamp %q: %v", timePart, err)
	}
}
