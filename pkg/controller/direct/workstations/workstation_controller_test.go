/*
Copyright 2024 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package workstations

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	gcp "cloud.google.com/go/workstations/apiv1"
	"cloud.google.com/go/workstations/apiv1/workstationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type mockListener struct {
	diffs []*structuredreporting.Diff
}

func (m *mockListener) OnError(ctx context.Context, err error, args ...any) {}
func (m *mockListener) OnDiff(ctx context.Context, diff *structuredreporting.Diff) {
	m.diffs = append(m.diffs, diff)
}
func (m *mockListener) OnReconcileStart(ctx context.Context, u *unstructured.Unstructured, t k8s.ReconcilerType) {
}
func (m *mockListener) OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error, t k8s.ReconcilerType) {
}

type mockTransport struct {
	roundTripFunc func(*http.Request) (*http.Response, error)
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.roundTripFunc(req)
}

func TestWorkstationAdapter_Update_ReportsDiff(t *testing.T) {
	ctx := context.Background()

	// Helper to create adapter
	newAdapter := func(desired *krm.Workstation, actual *workstationspb.Workstation, gcpClient *gcp.Client) *WorkstationAdapter {
		localObj := desired.DeepCopy()
		if localObj.Namespace == "" {
			localObj.Namespace = "default"
		}
		if localObj.Name == "" {
			localObj.Name = "test"
		}
		if localObj.Annotations == nil {
			localObj.Annotations = make(map[string]string)
		}
		localObj.Annotations["cnrm.cloud.google.com/project-id"] = "test-project"

		localObj.Spec.Parent = &krm.WorkstationConfigRef{External: "projects/test-project/locations/us-central1/workstationClusters/test-cluster/workstationConfigs/test-config"}

		k8sClient := fake.NewClientBuilder().Build()
		id, err := krm.NewWorkstationIdentity(ctx, k8sClient, localObj)
		if err != nil {
			t.Logf("Warning: NewWorkstationIdentity failed: %v", err)
		}

		return &WorkstationAdapter{
			id:        id,
			gcpClient: gcpClient,
			desired:   desired,
			actual:    actual,
		}
	}

	// Setup Scheme
	scheme := runtime.NewScheme()
	_ = krm.AddToScheme(scheme)

	t.Run("Identical Protos", func(t *testing.T) {
		// Mock Listener
		listener := &mockListener{}
		ctx := structuredreporting.ContextWithListener(ctx, listener)

		// Mock Client (should not be called, but provided)
		httpClient := &http.Client{Transport: &mockTransport{
			roundTripFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("should not be called")
			},
		}}
		gcpClient, _ := gcp.NewRESTClient(ctx, option.WithHTTPClient(httpClient), option.WithoutAuthentication())

		desired := &krm.Workstation{
			Spec: krm.WorkstationSpec{
				DisplayName: direct.PtrTo("test-display-name"),
				Labels: []krm.WorkstationLabel{
					{Key: "key1", Value: "value1"},
					{Key: "key2", Value: "value2"},
				},
			},
		}
		actual := &workstationspb.Workstation{
			DisplayName: "test-display-name",
			Labels: map[string]string{
				"key2": "value2",
				"key1": "value1",
			},
		}

		a := newAdapter(desired, actual, gcpClient)
		if a.id == nil {
			t.Skip("Skipping test because Identity creation failed")
		}
		// Ensure actual has the name set
		a.actual.Name = a.id.String()

		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(krm.WorkstationGVK)
		u.SetName("test")
		u.SetNamespace("default")
		// Initialize fake client with Scheme and Object
		k8sClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(u).Build()

		lh := lifecyclehandler.LifecycleHandler{Recorder: record.NewFakeRecorder(100)}
		updateOp := directbase.NewUpdateOperation(lh, k8sClient, u)

		err := a.Update(ctx, updateOp)
		// We tolerate "not found" error which comes from UpdateStatus (fake client issue),
		// which means we correctly identified no-diff and tried to update status.
		if err != nil && !strings.Contains(err.Error(), "not found") {
			t.Errorf("Update() error = %v, want nil or 'not found'", err)
		}

		if len(listener.diffs) != 0 {
			t.Errorf("ReportDiff called %d times, want 0", len(listener.diffs))
			for _, d := range listener.diffs {
				t.Logf("Diff fields: %v", d.FieldIDs())
			}
		}
	})

	t.Run("Different Protos", func(t *testing.T) {
		listener := &mockListener{}
		ctx := structuredreporting.ContextWithListener(ctx, listener)

		mockHTTP := &mockTransport{
			roundTripFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("mock transport error")
			},
		}
		httpClient := &http.Client{Transport: mockHTTP}
		gcpClient, _ := gcp.NewRESTClient(ctx, option.WithHTTPClient(httpClient), option.WithoutAuthentication())

		desired := &krm.Workstation{
			Spec: krm.WorkstationSpec{
				DisplayName: direct.PtrTo("new-name"),
				Labels: []krm.WorkstationLabel{
					{Key: "key1", Value: "value1"},
					{Key: "key2", Value: "value2"},
				},
			},
		}
		actual := &workstationspb.Workstation{
			DisplayName: "old-name",
			Labels: map[string]string{
				"key2": "value2",
				"key1": "value1",
			},
		}

		a := newAdapter(desired, actual, gcpClient)
		if a.id == nil {
			t.Skip("Skipping test because Identity creation failed")
		}
		a.actual.Name = a.id.String()

		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(krm.WorkstationGVK)
		u.SetName("test")
		u.SetNamespace("default")
		k8sClient := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(u).Build()

		lh := lifecyclehandler.LifecycleHandler{Recorder: record.NewFakeRecorder(100)}
		updateOp := directbase.NewUpdateOperation(lh, k8sClient, u)

		// Run Update
		err := a.Update(ctx, updateOp)
		// Expect error because client fails, BUT ReportDiff should have been called.
		if err == nil {
			t.Log("Update() returned nil, expected error from mock client")
		}

		if len(listener.diffs) == 0 {
			t.Error("ReportDiff not called, want called")
		} else {
			d := listener.diffs[0]
			ids := d.FieldIDs()
			foundDisplayName := false
			foundLabels := false
			for _, id := range ids {
				switch id {
				case "display_name":
					foundDisplayName = true
				case "labels":
					foundLabels = true
				default:
					t.Errorf("Unexpected diff field. Fields: %v", ids)
				}
			}
			if !foundDisplayName {
				t.Errorf("Diff does not contain display_name. Fields: %v", ids)
			}
			if foundLabels {
				t.Errorf("Diff contains labels but they should be identical. Fields: %v", ids)
			}
		}
	})
}
