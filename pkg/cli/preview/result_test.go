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
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
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
	var logBuf bytes.Buffer
	klog.LogToStderr(false)
	klog.SetOutput(&logBuf)
	defer func() {
		klog.LogToStderr(true)
		klog.SetOutput(os.Stderr)
	}()

	r := &RecorderReconciledResults{
		results: map[GKNN]*GKNNReconciledResult{
			{Group: "g1", Kind: "K1", Namespace: "n1", Name: "name1"}: {
				GKNN:            GKNN{Group: "g1", Kind: "K1", Namespace: "n1", Name: "name1"},
				CurrentStatus:   "UpToDate",
				ControllerType:  k8s.ReconcilerType("tf"),
				ReconcileStatus: ReconcileStatusHealthy,
			},
			{Group: "g3", Kind: "K3", Namespace: "n3", Name: "name3"}: {
				GKNN:            GKNN{Group: "g3", Kind: "K3", Namespace: "n3", Name: "name3"},
				CurrentStatus:   "UpdateFailed",
				ControllerType:  k8s.ReconcilerType("tf"),
				ReconcileStatus: ReconcileStatusUnhealthy,
			},
			{Group: "g4", Kind: "K4", Namespace: "n4", Name: "name4"}: {
				GKNN:            GKNN{Group: "g4", Kind: "K4", Namespace: "n4", Name: "name4"},
				CurrentStatus:   "UpdateFailed",
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusUnhealthy,
			},
			{Group: "g5", Kind: "K5", Namespace: "n5", Name: "name5"}: {
				GKNN:            GKNN{Group: "g5", Kind: "K5", Namespace: "n5", Name: "name5"},
				CurrentStatus:   "UpToDate",
				ControllerType:  k8s.ReconcilerType("tf"),
				ReconcileStatus: ReconcileStatusHealthy,
			},
			{Group: "g6", Kind: "K6", Namespace: "n6", Name: "name6"}: {
				GKNN:            GKNN{Group: "g6", Kind: "K6", Namespace: "n6", Name: "name6"},
				CurrentStatus:   "UpToDate",
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusHealthy,
			},
		},
	}
	alt := &RecorderReconciledResults{
		results: map[GKNN]*GKNNReconciledResult{
			{Group: "g1", Kind: "K1", Namespace: "n1", Name: "name1"}: {
				GKNN:            GKNN{Group: "g1", Kind: "K1", Namespace: "n1", Name: "name1"},
				CurrentStatus:   "UpToDate",
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusUnhealthy,
			},
			{Group: "g2", Kind: "K2", Namespace: "n2", Name: "name2"}: {
				GKNN:            GKNN{Group: "g2", Kind: "K2", Namespace: "n2", Name: "name2"},
				CurrentStatus:   "UpToDate",
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusHealthy,
			},
			{Group: "g3", Kind: "K3", Namespace: "n3", Name: "name3"}: {
				GKNN:            GKNN{Group: "g3", Kind: "K3", Namespace: "n3", Name: "name3"},
				CurrentStatus:   "UpdateFailed",
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusUnhealthy,
			},
			{Group: "g4", Kind: "K4", Namespace: "n4", Name: "name4"}: {
				GKNN:            GKNN{Group: "g4", Kind: "K4", Namespace: "n4", Name: "name4"},
				CurrentStatus:   "UpdateFailed",
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusUnhealthy,
			},
			{Group: "g5", Kind: "K5", Namespace: "n5", Name: "name5"}: {
				GKNN:            GKNN{Group: "g5", Kind: "K5", Namespace: "n5", Name: "name5"},
				CurrentStatus:   "UpToDate",
				ControllerType:  k8s.ReconcilerType("direct"),
				ReconcileStatus: ReconcileStatusHealthy,
			},
			{Group: "g6", Kind: "K6", Namespace: "n6", Name: "name6"}: {
				GKNN:            GKNN{Group: "g6", Kind: "K6", Namespace: "n6", Name: "name6"},
				CurrentStatus:   "UpToDate",
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
	defer os.Remove(tmpFile.Name() + "-detail")
	tmpFile.Close()

	altExpectedMap := map[schema.GroupKind]k8s.ReconcilerType{
		{Group: "g1", Kind: "K1"}: k8s.ReconcilerType("direct"),
		{Group: "g2", Kind: "K2"}: k8s.ReconcilerType("direct"),
		{Group: "g3", Kind: "K3"}: k8s.ReconcilerType("direct"),
		{Group: "g4", Kind: "K4"}: k8s.ReconcilerType("direct"),
		{Group: "g5", Kind: "K5"}: k8s.ReconcilerType("direct"),
		{Group: "g6", Kind: "K6"}: k8s.ReconcilerType("direct"),
	}

	if err := r.CombinedSummaryReport(tmpFile.Name(), alt, altExpectedMap); err != nil {
		t.Fatalf("CombinedSummaryReport failed: %v", err)
	}
	klog.Flush()

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read summary file: %v", err)
	}

	// Verify the content contains the expected headers and rows
	expectedRows := []string{
		"GROUP   KIND   NAMESPACE   NAME    CURRENT-STATUS   DEFAULT-CONTROLLER   DEFAULT-RESULT   ALTERNATIVE-CONTROLLER   ALTERNATIVE-RESULT   DEFAULT-DIFFS   ALTERNATIVE-DIFFS",
		"g1      K1     n1          name1   UpToDate         tf                   HEALTHY          direct                   UNHEALTHY            N/A             N/A",
		"g2      K2     n2          name2   UpToDate         N/A                  N/A              direct                   HEALTHY              N/A             N/A",
		"g3      K3     n3          name3   UpdateFailed     tf                   UNHEALTHY        direct                   UNHEALTHY            N/A             N/A",
		"g4      K4     n4          name4   UpdateFailed     direct               UNHEALTHY        direct                   UNHEALTHY            N/A             N/A",
		"g5      K5     n5          name5   UpToDate         tf                   HEALTHY          direct                   HEALTHY              N/A             N/A",
		"g6      K6     n6          name6   UpToDate         direct               HEALTHY          direct                   HEALTHY              N/A             N/A",
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

	// Verify klog output prints alternative run results when controller type differs even if healthy/unhealthy status is the same
	logs := logBuf.String()
	expectedLogSubstrings := []string{
		`name="name3" group="g3" kind="K3" current_status="UpdateFailed" controller_type="tf"`,
		`name="name3" group="g3" kind="K3" current_status="UpdateFailed" controller_type="direct"`,
		`name="name5" group="g5" kind="K5" current_status="UpToDate" controller_type="tf"`,
		`name="name5" group="g5" kind="K5" current_status="UpToDate" controller_type="direct"`,
	}
	for _, sub := range expectedLogSubstrings {
		if !strings.Contains(logs, sub) {
			t.Errorf("expected klog output to contain %q, got:\n%s", sub, logs)
		}
	}

	// Verify klog output does not duplicate when both controller type and status/diffs are the same (name4 and name6)
	if count := strings.Count(logs, `name="name4"`); count != 1 {
		t.Errorf("expected name4 to be logged once, got %d times in:\n%s", count, logs)
	}
	if count := strings.Count(logs, `name="name6"`); count != 1 {
		t.Errorf("expected name6 to be logged once, got %d times in:\n%s", count, logs)
	}

	// Verify the detail report contains the expected bad results and is deduplicated/sorted correctly
	detailContent, err := os.ReadFile(tmpFile.Name() + "-detail")
	if err != nil {
		t.Fatalf("failed to read detail file: %v", err)
	}

	var badResults []*GKNNReconciledResult
	if err := json.Unmarshal(detailContent, &badResults); err != nil {
		t.Fatalf("failed to unmarshal detail file: %v", err)
	}

	// Expected bad results:
	// 1. name1 (alt: direct)
	// 2. name3 (def: tf)
	// 3. name3 (alt: direct) - same status/diffs as def, but different controller type -> included
	// 4. name4 (def: direct) - alt has same controller type AND same status/diffs -> deduplicated (only 1 entry)
	expectedBadCount := 4
	if len(badResults) != expectedBadCount {
		t.Fatalf("expected %d bad results in detail report, got %d: %s", expectedBadCount, len(badResults), string(detailContent))
	}

	expectedEntries := []struct {
		name           string
		controllerType k8s.ReconcilerType
	}{
		{name: "name1", controllerType: k8s.ReconcilerType("direct")},
		{name: "name3", controllerType: k8s.ReconcilerType("tf")},
		{name: "name3", controllerType: k8s.ReconcilerType("direct")},
		{name: "name4", controllerType: k8s.ReconcilerType("direct")},
	}
	for i, exp := range expectedEntries {
		if badResults[i].GKNN.Name != exp.name || badResults[i].ControllerType != exp.controllerType {
			t.Errorf("entry %d: expected name=%q controller=%q, got name=%q controller=%q", i, exp.name, exp.controllerType, badResults[i].GKNN.Name, badResults[i].ControllerType)
		}
	}
}
