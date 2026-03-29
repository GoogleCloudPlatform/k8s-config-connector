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
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestGetAlternativeControllerExpectedMap(t *testing.T) {
	testCases := []struct {
		name      string
		configMap resourceconfig.ResourcesControllerMap
		expected  map[schema.GroupKind]k8s.ReconcilerType
	}{
		{
			name: "single config with alternative controller",
			configMap: resourceconfig.ResourcesControllerMap{
				{Group: "foo", Kind: "Bar"}: {
					DefaultController:    k8s.ReconcilerType("tf"),
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerType("tf"), k8s.ReconcilerType("direct")},
				},
			},
			expected: map[schema.GroupKind]k8s.ReconcilerType{
				{Group: "foo", Kind: "Bar"}: k8s.ReconcilerType("direct"),
			},
		},
		{
			name: "single config without alternative controller",
			configMap: resourceconfig.ResourcesControllerMap{
				{Group: "foo", Kind: "Bar"}: {
					DefaultController:    k8s.ReconcilerType("direct"),
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerType("direct")},
				},
			},
			expected: map[schema.GroupKind]k8s.ReconcilerType{},
		},
		{
			name: "multiple configs mixed",
			configMap: resourceconfig.ResourcesControllerMap{
				{Group: "foo", Kind: "Bar"}: {
					DefaultController:    k8s.ReconcilerType("tf"),
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerType("tf"), k8s.ReconcilerType("direct")},
				},
				{Group: "storage", Kind: "StorageBucket"}: {
					DefaultController:    k8s.ReconcilerType("direct"),
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerType("direct")},
				},
				{Group: "compute", Kind: "Instance"}: {
					DefaultController:    k8s.ReconcilerType("tf"),
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerType("tf"), k8s.ReconcilerType("dcl")},
				},
			},
			expected: map[schema.GroupKind]k8s.ReconcilerType{
				{Group: "foo", Kind: "Bar"}:          k8s.ReconcilerType("direct"),
				{Group: "compute", Kind: "Instance"}: k8s.ReconcilerType("dcl"),
			},
		},
		{
			name: "multiple alternatives picks first non-default",
			configMap: resourceconfig.ResourcesControllerMap{
				{Group: "foo", Kind: "Quad"}: {
					DefaultController:    k8s.ReconcilerType("tf"),
					SupportedControllers: []k8s.ReconcilerType{k8s.ReconcilerType("tf"), k8s.ReconcilerType("direct"), k8s.ReconcilerType("dcl")},
				},
			},
			expected: map[schema.GroupKind]k8s.ReconcilerType{
				{Group: "foo", Kind: "Quad"}: k8s.ReconcilerType("direct"),
			},
		},
		{
			name:      "empty config",
			configMap: resourceconfig.ResourcesControllerMap{},
			expected:  map[schema.GroupKind]k8s.ReconcilerType{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetAlternativeControllerExpectedMap(tc.configMap)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("expected \n%v\n, got \n%v", tc.expected, actual)
			}
		})
	}
}

func TestCombinedSummaryReport(t *testing.T) {
	r := &RecorderReconciledResults{
		results: map[GKNN]*GKNNReconciledResult{
			{Group: "g1", Kind: "K1", Namespace: "n1", Name: "name1"}: {
				GKNN:            GKNN{Group: "g1", Kind: "K1", Namespace: "n1", Name: "name1"},
				ControllerType:  k8s.ReconcilerType("tf"),
				ReconcileStatus: ReconcileStatusHealthy,
			},
		},
	}
	alt := &RecorderReconciledResults{
		results: map[GKNN]*GKNNReconciledResult{
			{Group: "g1", Kind: "K1", Namespace: "n1", Name: "name1"}: {
				GKNN:            GKNN{Group: "g1", Kind: "K1", Namespace: "n1", Name: "name1"},
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusUnhealthy,
			},
			{Group: "g2", Kind: "K2", Namespace: "n2", Name: "name2"}: {
				GKNN:            GKNN{Group: "g2", Kind: "K2", Namespace: "n2", Name: "name2"},
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusHealthy,
			},
		},
	}

	tmpFile, err := os.CreateTemp("", "summary-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	altExpectedMap := map[schema.GroupKind]k8s.ReconcilerType{
		{Group: "g1", Kind: "K1"}: k8s.ReconcilerType("direct"),
		{Group: "g2", Kind: "K2"}: k8s.ReconcilerType("direct"),
	}

	if err := r.CombinedSummaryReport(tmpFile.Name(), alt, altExpectedMap); err != nil {
		t.Fatalf("CombinedSummaryReport failed: %v", err)
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read summary file: %v", err)
	}

	// Verify the content contains the expected headers and rows
	expectedRows := []string{
		"GROUP   KIND   NAME    DEFAULT-CONTROLLER   DEFAULT-RESULT   DEFAULT-DIFFS   ALTERNATIVE-CONTROLLER   ALTERNATIVE-RESULT   ALTERNATIVE-DIFFS",
		"g1      K1     name1   tf                   HEALTHY          N/A             direct                   UNHEALTHY            N/A",
		"g2      K2     name2   N/A                  N/A              N/A             direct                   HEALTHY              N/A",
	}

	for _, row := range expectedRows {
		if !strings.Contains(string(content), row) {
			// tabwriter might use different spacing, let's just check for the key parts
			parts := strings.Fields(row)
			allFound := true
			for _, part := range parts {
				if !strings.Contains(string(content), part) {
					allFound = false
					break
				}
			}
			if !allFound {
				t.Errorf("expected row %q not found in content:\n%s", row, string(content))
			}
		}
	}
}
