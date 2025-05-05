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
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcptests"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/yaml"
)

// TestLocalModePreview creates an object using KCC, and then runs the preview command to look for additional changes
func TestLocalModePreview(t *testing.T) {
	log.SetLogger(klogr.New())

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	harness := mockgcptests.NewHarness(ctx, t, mockgcptests.WithGCPTarget("mock"))
	harness.Init()

	localKube := NewLocalModeKube()

	{
		crd := &unstructured.Unstructured{}
		crd.SetGroupVersionKind(schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "CustomResourceDefinition"})
		crd.SetName("pubsubtopics.pubsub.cnrm.cloud.google.com")
		unstructured.SetNestedField(crd.Object, "pubsub.cnrm.cloud.google.com", "spec", "group")
		unstructured.SetNestedField(crd.Object, "PubSubTopic", "spec", "names", "kind")
		unstructured.SetNestedField(crd.Object, "pubsubtopics", "spec", "names", "plural")
		unstructured.SetNestedField(crd.Object, false, "spec", "preserveUnknownFields")
		unstructured.SetNestedField(crd.Object, "Namespaced", "spec", "scope")
		versions := []any{
			map[string]any{
				"name": "v1beta1",
				"schema": map[string]any{
					"openAPIV3Schema": map[string]any{
						"type": "object",
					},
				},
			},
		}
		unstructured.SetNestedField(crd.Object, versions, "spec", "versions")

		crd.SetLabels(map[string]string{
			"cnrm.cloud.google.com/managed-by-kcc":  "true",
			"cnrm.cloud.google.com/stability-level": "stable",
			"cnrm.cloud.google.com/system":          "true",
			"cnrm.cloud.google.com/tf2crd":          "true",
		})

		localKube.AddObject(crd)
	}

	{
		group := "core.cnrm.cloud.google.com"
		kind := "ConfigConnectorContext"
		resource := "configconnectorcontexts"
		scope := "Namespaced"
		version := "v1beta1"

		crd := &unstructured.Unstructured{}
		crd.SetGroupVersionKind(schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "CustomResourceDefinition"})
		crd.SetName(resource + "." + group)
		unstructured.SetNestedField(crd.Object, group, "spec", "group")
		unstructured.SetNestedField(crd.Object, kind, "spec", "names", "kind")
		unstructured.SetNestedField(crd.Object, resource, "spec", "names", "plural")
		unstructured.SetNestedField(crd.Object, false, "spec", "preserveUnknownFields")
		unstructured.SetNestedField(crd.Object, scope, "spec", "scope")
		versions := []any{
			map[string]any{
				"name": version,
				"schema": map[string]any{
					"openAPIV3Schema": map[string]any{
						"type": "object",
					},
				},
			},
		}
		unstructured.SetNestedField(crd.Object, versions, "spec", "versions")

		crd.SetLabels(map[string]string{})

		localKube.AddObject(crd)
	}

	{
		group := "core.cnrm.cloud.google.com"
		kind := "ConfigConnector"
		resource := "configconnectors"
		scope := "Cluster"
		version := "v1beta1"

		crd := &unstructured.Unstructured{}
		crd.SetGroupVersionKind(schema.GroupVersionKind{Group: "apiextensions.k8s.io", Version: "v1", Kind: "CustomResourceDefinition"})
		crd.SetName(resource + "." + group)
		unstructured.SetNestedField(crd.Object, group, "spec", "group")
		unstructured.SetNestedField(crd.Object, kind, "spec", "names", "kind")
		unstructured.SetNestedField(crd.Object, resource, "spec", "names", "plural")
		unstructured.SetNestedField(crd.Object, false, "spec", "preserveUnknownFields")
		unstructured.SetNestedField(crd.Object, scope, "spec", "scope")
		versions := []any{
			map[string]any{
				"name": version,
				"schema": map[string]any{
					"openAPIV3Schema": map[string]any{
						"type": "object",
					},
				},
			},
		}
		unstructured.SetNestedField(crd.Object, versions, "spec", "versions")

		crd.SetLabels(map[string]string{})

		localKube.AddObject(crd)
	}
	// Create KCC objects
	ns := &unstructured.Unstructured{}
	ns.SetGroupVersionKind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"})
	ns.SetName(harness.Project.ProjectID)
	localKube.AddObject(ns)

	cc := &unstructured.Unstructured{}
	cc.SetGroupVersionKind(schema.GroupVersionKind{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnector"})
	cc.SetName("configconnector.core.cnrm.cloud.google.com")
	MustSetNestedField(t, cc, "spec.mode", "namespaced")
	localKube.AddObject(cc)

	ccc := &unstructured.Unstructured{}
	ccc.SetGroupVersionKind(schema.GroupVersionKind{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnectorContext"})
	ccc.SetName("configconnectorcontext.core.cnrm.cloud.google.com")
	ccc.SetNamespace(ns.GetName())
	MustSetNestedField(t, ccc, "spec.googleServiceAccount", "fake@fake.iam.gserviceaccount.com")
	localKube.AddObject(ccc)

	// Create a pubsub topic (should be created in mock gcp)
	testObj := &unstructured.Unstructured{}

	{
		y := `
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: pubsubtopic-example
spec:
  messageRetentionDuration: "3600s"
`

		if err := yaml.Unmarshal([]byte(y), testObj); err != nil {
			t.Fatalf("unmarshaling yaml: %v", err)
		}
		testObj.SetNamespace(ns.GetName())
		annotations := testObj.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}
		// TODO: we should not rely on webhooks
		annotations["cnrm.cloud.google.com/project-id"] = harness.Project.ProjectID
		testObj.SetAnnotations(annotations)

		// TODO: precreate finalizers?
		finalizers := testObj.GetFinalizers()
		finalizers = append(finalizers, "cnrm.cloud.google.com/finalizer", "cnrm.cloud.google.com/deletion-defender")
		testObj.SetFinalizers(finalizers)

		localKube.AddObject(testObj)

		// Wait for object to be ready
		// create.WaitForReadySingleResource(ctx, localKube.BuildControllerRuntimeClient(), testObj, time.Minute)
	}

	recorder := NewRecorder()

	authorization := harness.GCPAuthorization()

	previewInstanceOptions := PreviewInstanceOptions{
		UpstreamKubeClient:       localKube.BuildKubeClient(),
		UpstreamKubeRESTMapper:   localKube.BuildRESTMapper(),
		UpstreamGCPAuthorization: authorization,
		UpstreamGCPHTTPClient:    harness.GCPHTTPClient(),
	}

	preview, err := NewPreviewInstance(recorder, previewInstanceOptions)
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
			if len(hasReconciled) > 0 {
				break
			}
		}
		if time.Now().After(timeoutAt) {
			t.Fatalf("did not see captured object in recorder")
		}
		time.Sleep(time.Second)
	}

	t.Logf("Printing captured changes")
	if len(recorder.objects) != 1 {
		t.Errorf("expected exactly one object to be reconciled; got %v", len(recorder.objects))
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

		expectedGKNN := GKNN{
			Group:     testObj.GroupVersionKind().Group,
			Kind:      testObj.GroupVersionKind().Kind,
			Name:      testObj.GetName(),
			Namespace: testObj.GetNamespace(),
		}
		if gknn != expectedGKNN {
			t.Errorf("unexpected object in changelist; got %v; want %v", gknn, expectedGKNN)
		}

		for _, event := range obj.events {
			switch event.eventType {
			case EventTypeDiff:
				// We aren't expected changes
				t.Errorf("unexpected diff in changelist: %+v", event.diff)

			case EventTypeKubeAction:
				// We aren't expected changes
				t.Errorf("unexpected kubeAction in changelist: %+v", event.kubeAction)

			case EventTypeGCPAction:
				// We aren't expected changes
				t.Errorf("unexpected gcpAction in changelist: %+v", event.gcpAction)

			case EventTypeReconcileStart, EventTypeReconcileEnd:
				// We do expect this!

			default:
				t.Errorf("unexpected event type in changelist: %v", event.eventType)
			}
		}
	}
}
