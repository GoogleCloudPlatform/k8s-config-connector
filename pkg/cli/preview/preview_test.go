// Copyright 2025 Google LLC
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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/yaml"
)

// TestPreview creates an object using KCC, and then runs the preview command to look for additional changes
func TestPreview(t *testing.T) {
	log.SetLogger(klogr.New())

	if os.Getenv("KUBEBUILDER_ASSETS") == "" {
		assetDir, err := getKubebuilderAssetDir()
		if err != nil {
			t.Fatalf("getting asset dir: %v", err)
		}
		klog.Warningf("defaulting KUBEBUILDER_ASSETS to %v", assetDir)
		os.Setenv("KUBEBUILDER_ASSETS", assetDir)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	harness := create.NewHarness(ctx, t, create.WithKubeTarget("envtest"), create.WithGCPTarget("mock"))

	// Create KCC objects
	ns := &unstructured.Unstructured{}
	ns.SetGroupVersionKind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"})
	ns.SetName(harness.Project.ProjectID)
	if err := harness.GetClient().Create(ctx, ns); err != nil {
		t.Fatalf("creating object: %v", err)
	}

	cc := &unstructured.Unstructured{}
	cc.SetGroupVersionKind(schema.GroupVersionKind{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnector"})
	cc.SetName("configconnector.core.cnrm.cloud.google.com")
	MustSetNestedField(t, cc, "spec.mode", "namespaced")
	if err := harness.GetClient().Create(ctx, cc); err != nil {
		t.Fatalf("creating object: %v", err)
	}

	ccc := &unstructured.Unstructured{}
	ccc.SetGroupVersionKind(schema.GroupVersionKind{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnectorContext"})
	ccc.SetName("configconnectorcontext.core.cnrm.cloud.google.com")
	ccc.SetNamespace(ns.GetName())
	MustSetNestedField(t, ccc, "spec.googleServiceAccount", "fake@fake.iam.gserviceaccount.com")
	if err := harness.GetClient().Create(ctx, ccc); err != nil {
		t.Fatalf("creating object: %v", err)
	}

	testResources := map[GKNN]struct {
		resourceSpec      string
		expectedEventType []EventType
	}{
		{
			Group:     "pubsub.cnrm.cloud.google.com",
			Kind:      "PubSubTopic",
			Namespace: ns.GetName(),
			Name:      "pubsubtopic-example",
		}: {
			resourceSpec: `apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: pubsubtopic-example
spec:
  messageRetentionDuration: "3600s"
`,
			expectedEventType: []EventType{EventTypeReconcileStart, EventTypeReconcileEnd},
		},
		{
			Group:     "spanner.cnrm.cloud.google.com",
			Kind:      "SpannerInstance",
			Namespace: ns.GetName(),
			Name:      "spannerinstance-sample",
		}: {
			resourceSpec: `apiVersion: spanner.cnrm.cloud.google.com/v1beta1
kind: SpannerInstance
metadata:
  labels:
    label-one: "value-one"
  name: spannerinstance-sample
  annotations:
    alpha.cnrm.cloud.google.com/reconciler: "direct"
spec:
  config: regional-us-west1
  displayName: Spanner Instance Sample
  numNodes: 2
`,
			expectedEventType: []EventType{EventTypeReconcileStart, EventTypeKubeAction, EventTypeReconcileEnd},
		},
	}
	{
		for gknn, resource := range testResources {
			testObj := &unstructured.Unstructured{}
			if err := yaml.Unmarshal([]byte(resource.resourceSpec), testObj); err != nil {
				t.Fatalf("unmarshaling yaml: %v", err)
			}
			testObj.SetNamespace(ns.GetName())
			if err := harness.GetClient().Create(ctx, testObj); err != nil {
				t.Fatalf("creating object `%s`: %v", gknn.Name, err)
			}

			// Wait for object to be ready
			create.WaitForReady(harness, time.Minute, testObj)
		}
	}

	// Now we can run our test ... let's run the preview mode, we expect a read of the GCP object but no write
	upstreamRESTConfig := harness.GetRESTConfig()

	recorder := NewRecorder()

	authorization := harness.GCPAuthorization()

	preview, err := NewPreviewInstance(recorder, PreviewInstanceOptions{
		UpstreamRESTConfig:       upstreamRESTConfig,
		UpstreamGCPAuthorization: authorization,
		UpstreamGCPHTTPClient:    harness.GCPHTTPClient(),
	})
	if err != nil {
		t.Fatalf("building preview instance: %v", err)
	}

	go func() {
		if err := preview.Start(ctx); err != nil {
			t.Errorf("starting preview: %v", err)
		}
	}()

	timeoutAt := time.Now().Add(2 * time.Minute)
	for {
		// Wait for the object to be reconciled
		if len(recorder.objects) > 0 {
			hasReconciled := make(map[GKNN]bool)
			for gknn, obj := range recorder.objects {
				for _, event := range obj.events {
					if event.eventType == EventTypeReconcileEnd {
						hasReconciled[gknn] = true
					}
				}
			}
			if len(hasReconciled) == len(testResources) {
				break
			}
		}
		if time.Now().After(timeoutAt) {
			t.Fatalf("did not see captured object in recorder")
		}
		time.Sleep(time.Second)
	}

	t.Logf("Printing captured changes")
	if len(recorder.objects) != len(testResources) {
		t.Errorf("expected exactly %d object to be reconciled; got %v", len(testResources), len(recorder.objects))
	}

	for gknn, obj := range recorder.objects {
		t.Logf("object %+v", gknn)
		for _, event := range obj.events {
			switch event.eventType {
			case EventTypeDiff:
				t.Logf("  diff %+v", event.diff)

			case EventTypeReconcileStart:
				t.Logf("  reconcileStart %+v", event.object)

			case EventTypeReconcileEnd:
				t.Logf("  reconcileEnd %+v", event.object)

			case EventTypeKubeAction:
				t.Logf("  kubeAction %+v", event.kubeAction)

			case EventTypeGCPAction:
				t.Logf("  gcpAction %+v", event.gcpAction)

			default:
				t.Logf("  unknown event: %+v", event)
			}
		}

		for gknn, resource := range testResources {
			objectInfo, ok := recorder.objects[gknn]
			if !ok {
				t.Logf("expected object not found in changelist; want %v", gknn)
			} else {
				if len(objectInfo.events) != len(resource.expectedEventType) {
					// TODO: enable this error once the reconcile more than once is fixed.
					t.Logf("unexpected number of events in changelist; got %v; want %v", len(objectInfo.events), len(resource.expectedEventType))
				}
				for i, expectedEventType := range resource.expectedEventType {
					if expectedEventType != objectInfo.events[i].eventType {
						t.Errorf("unexpected event type in changelist; got %v; want %v", objectInfo.events[i].eventType, expectedEventType)
					}
				}
			}
		}
	}
}

// getKubebuilderAssetDir returns the path to the kubebuilder assets directory
// which is the latest directory in the ~/.local/share/kubebuilder-envtest/k8s directory
func getKubebuilderAssetDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting home dir: %w", err)
	}
	dir := filepath.Join(homeDir, ".local", "share", "kubebuilder-envtest", "k8s")
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("reading directory %v: %w", dir, err)
	}
	var candidates []string
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		candidates = append(candidates, file.Name())
	}
	if len(candidates) == 0 {
		return "", fmt.Errorf("found no kubebuilder assets in %v", dir)
	}
	bestCandidate := candidates[len(candidates)-1]
	return filepath.Join(dir, bestCandidate), nil
}

// MustSetNestedField sets a nested field on an unstructured object
// and panics if it fails.  This is a helper function for tests.
func MustSetNestedField(t *testing.T, obj *unstructured.Unstructured, path string, value interface{}) {
	fields := strings.Split(path, ".")
	if err := unstructured.SetNestedField(obj.Object, value, fields...); err != nil {
		t.Fatalf("setting nested field %v: %v", path, err)
	}
}
