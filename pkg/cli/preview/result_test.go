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
	"fmt"
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

func TestFormatReconciledStatus(t *testing.T) {
	tests := []struct {
		name     string
		result   *GKNNReconciledResult
		expected string
	}{
		{
			name:     "nil result",
			result:   nil,
			expected: "N/A",
		},
		{
			name: "healthy result",
			result: &GKNNReconciledResult{
				ReconcileStatus: ReconcileStatusHealthy,
			},
			expected: "HEALTHY",
		},
		{
			name: "unhealthy result with no reasons",
			result: &GKNNReconciledResult{
				ReconcileStatus: ReconcileStatusUnhealthy,
			},
			expected: "UNHEALTHY",
		},
		{
			name: "unhealthy result with RECONCILE_ERROR",
			result: &GKNNReconciledResult{
				ReconcileStatus:  ReconcileStatusUnhealthy,
				UnhealthyReasons: []UnhealthyReason{UnhealthyReasonError},
			},
			expected: "UNHEALTHY (RECONCILE_ERROR)",
		},
		{
			name: "unhealthy result with GCP_WRITE",
			result: &GKNNReconciledResult{
				ReconcileStatus:  ReconcileStatusUnhealthy,
				UnhealthyReasons: []UnhealthyReason{UnhealthyReasonGCPWrite},
			},
			expected: "UNHEALTHY (GCP_WRITE)",
		},
		{
			name: "unhealthy result with multiple reasons",
			result: &GKNNReconciledResult{
				ReconcileStatus:  ReconcileStatusUnhealthy,
				UnhealthyReasons: []UnhealthyReason{UnhealthyReasonGCPWrite, UnhealthyReasonError},
			},
			expected: "UNHEALTHY (GCP_WRITE, RECONCILE_ERROR)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := formatReconciledStatus(tc.result)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestGenerateRecorderReconciledResults(t *testing.T) {
	recorder := NewRecorder()

	gknnHealthy := GKNN{Group: "storage.cnrm.cloud.google.com", Kind: "StorageBucket", Namespace: "default", Name: "healthy-bucket"}
	gknnReconcileEndErr := GKNN{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerInstance", Namespace: "default", Name: "err-instance"}
	gknnReportErr := GKNN{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInstance", Namespace: "default", Name: "report-err-instance"}
	gknnGCPWrite := GKNN{Group: "pubsub.cnrm.cloud.google.com", Kind: "PubSubTopic", Namespace: "default", Name: "gcp-write-topic"}
	gknnMultiErr := GKNN{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicy", Namespace: "default", Name: "multi-err-policy"}

	// 1. Healthy object
	info1 := recorder.getObjectInfo(gknnHealthy)
	info1.currentStatus = "UpToDate"
	info1.events = []event{
		{eventType: EventTypeReconcileStart, reconcilerType: k8s.ReconcilerTypeDirect},
		{eventType: EventTypeReconcileEnd, reconcilerType: k8s.ReconcilerTypeDirect},
	}

	// 2. Early reconcile failure via ReconcileEnd error
	info2 := recorder.getObjectInfo(gknnReconcileEndErr)
	info2.currentStatus = "UpdateFailed"
	info2.events = []event{
		{eventType: EventTypeReconcileStart, reconcilerType: k8s.ReconcilerTypeDirect},
		{eventType: EventTypeReconcileEnd, reconcilerType: k8s.ReconcilerTypeDirect, err: fmt.Errorf("mapping error: proto conversion failed")},
	}

	// 3. Error reported via OnError (EventTypeError)
	info3 := recorder.getObjectInfo(gknnReportErr)
	info3.currentStatus = "UpdateFailed"
	info3.events = []event{
		{eventType: EventTypeReconcileStart, reconcilerType: k8s.ReconcilerTypeTerraform},
		{eventType: EventTypeError, err: fmt.Errorf("reference resolution failed")},
		{eventType: EventTypeReconcileEnd, reconcilerType: k8s.ReconcilerTypeTerraform},
	}

	// 4. Blocked GCP action
	info4 := recorder.getObjectInfo(gknnGCPWrite)
	info4.currentStatus = "UpToDate"
	info4.events = []event{
		{eventType: EventTypeReconcileStart, reconcilerType: k8s.ReconcilerTypeDirect},
		{eventType: EventTypeGCPAction, gcpAction: &gcpAction{Method: "POST", URL: "https://pubsub.googleapis.com/v1/projects/p/topics/t"}},
		{eventType: EventTypeReconcileEnd, reconcilerType: k8s.ReconcilerTypeDirect},
	}

	// 5. Both GCP Action and multiple errors (including duplicate error)
	info5 := recorder.getObjectInfo(gknnMultiErr)
	info5.currentStatus = "UpdateFailed"
	info5.events = []event{
		{eventType: EventTypeReconcileStart, reconcilerType: k8s.ReconcilerTypeIAMPolicy},
		{eventType: EventTypeGCPAction, gcpAction: &gcpAction{Method: "PUT", URL: "https://cloudresourcemanager.googleapis.com/v1/projects/p:setIamPolicy"}},
		{eventType: EventTypeError, err: fmt.Errorf("iam member invalid")},
		{eventType: EventTypeReconcileEnd, reconcilerType: k8s.ReconcilerTypeIAMPolicy, err: fmt.Errorf("iam member invalid")},
	}

	results := recorder.GenerateRecorderReconciledResults()

	if results.goodCount != 1 {
		t.Errorf("expected goodCount=1, got %d", results.goodCount)
	}
	if results.badCount != 4 {
		t.Errorf("expected badCount=4, got %d", results.badCount)
	}

	// Check healthy
	res1 := results.results[gknnHealthy]
	if res1.ReconcileStatus != ReconcileStatusHealthy {
		t.Errorf("expected healthy status, got %v", res1.ReconcileStatus)
	}
	if len(res1.UnhealthyReasons) != 0 {
		t.Errorf("expected no unhealthy reasons for healthy object, got %v", res1.UnhealthyReasons)
	}
	if len(res1.Errors) != 0 {
		t.Errorf("expected no errors for healthy object, got %v", res1.Errors)
	}

	// Check ReconcileEnd error
	res2 := results.results[gknnReconcileEndErr]
	if res2.ReconcileStatus != ReconcileStatusUnhealthy {
		t.Errorf("expected unhealthy status, got %v", res2.ReconcileStatus)
	}
	if !reflect.DeepEqual(res2.UnhealthyReasons, []UnhealthyReason{UnhealthyReasonError}) {
		t.Errorf("expected UnhealthyReasons [%s], got %v", UnhealthyReasonError, res2.UnhealthyReasons)
	}
	if len(res2.Errors) != 1 || res2.Errors[0] != "mapping error: proto conversion failed" {
		t.Errorf("expected error 'mapping error: proto conversion failed', got %v", res2.Errors)
	}

	// Check EventTypeError
	res3 := results.results[gknnReportErr]
	if res3.ReconcileStatus != ReconcileStatusUnhealthy {
		t.Errorf("expected unhealthy status, got %v", res3.ReconcileStatus)
	}
	if !reflect.DeepEqual(res3.UnhealthyReasons, []UnhealthyReason{UnhealthyReasonError}) {
		t.Errorf("expected UnhealthyReasons [%s], got %v", UnhealthyReasonError, res3.UnhealthyReasons)
	}
	if len(res3.Errors) != 1 || res3.Errors[0] != "reference resolution failed" {
		t.Errorf("expected error 'reference resolution failed', got %v", res3.Errors)
	}

	// Check GCPAction
	res4 := results.results[gknnGCPWrite]
	if res4.ReconcileStatus != ReconcileStatusUnhealthy {
		t.Errorf("expected unhealthy status, got %v", res4.ReconcileStatus)
	}
	if !reflect.DeepEqual(res4.UnhealthyReasons, []UnhealthyReason{UnhealthyReasonGCPWrite}) {
		t.Errorf("expected UnhealthyReasons [%s], got %v", UnhealthyReasonGCPWrite, res4.UnhealthyReasons)
	}
	if len(res4.Errors) != 0 {
		t.Errorf("expected no errors for gcp action object, got %v", res4.Errors)
	}

	// Check MultiErr (GCP Write + Error deduplicated)
	res5 := results.results[gknnMultiErr]
	if res5.ReconcileStatus != ReconcileStatusUnhealthy {
		t.Errorf("expected unhealthy status, got %v", res5.ReconcileStatus)
	}
	expectedReasons := []UnhealthyReason{UnhealthyReasonGCPWrite, UnhealthyReasonError}
	if !reflect.DeepEqual(res5.UnhealthyReasons, expectedReasons) {
		t.Errorf("expected UnhealthyReasons %v, got %v", expectedReasons, res5.UnhealthyReasons)
	}
	if len(res5.Errors) != 1 || res5.Errors[0] != "iam member invalid" {
		t.Errorf("expected deduplicated single error 'iam member invalid', got %v", res5.Errors)
	}
}

func TestBadResultReport_DetailJSON(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "bad-detail-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	badResults := []*GKNNReconciledResult{
		{
			GKNN:             GKNN{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerInstance", Namespace: "default", Name: "instance-1"},
			CurrentStatus:    "UpdateFailed",
			ControllerType:   k8s.ReconcilerTypeDirect,
			ReconcileStatus:  ReconcileStatusUnhealthy,
			UnhealthyReasons: []UnhealthyReason{UnhealthyReasonError},
			Errors:           []string{"mapping error occurred"},
		},
	}

	r := &RecorderReconciledResults{}
	if err := r.BadResultReport(tmpFile.Name(), badResults); err != nil {
		t.Fatalf("BadResultReport failed: %v", err)
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to read detail file: %v", err)
	}

	var parsed []*GKNNReconciledResult
	if err := json.Unmarshal(content, &parsed); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if len(parsed) != 1 {
		t.Fatalf("expected 1 result, got %d", len(parsed))
	}
	if !reflect.DeepEqual(parsed[0].UnhealthyReasons, []UnhealthyReason{UnhealthyReasonError}) {
		t.Errorf("expected UnhealthyReasons [%s], got %v", UnhealthyReasonError, parsed[0].UnhealthyReasons)
	}
	if !reflect.DeepEqual(parsed[0].Errors, []string{"mapping error occurred"}) {
		t.Errorf("expected Errors ['mapping error occurred'], got %v", parsed[0].Errors)
	}
}
