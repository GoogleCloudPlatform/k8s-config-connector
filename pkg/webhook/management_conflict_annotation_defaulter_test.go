// Copyright 2022 Google LLC
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

package webhook

import (
	"context"
	"errors"
	"maps"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/managementconflict"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testutil "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
	"github.com/nasa9084/go-openapi"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func TestDefaultManagementConflictAnnotationForDCLBasedResources(t *testing.T) {
	tests := []struct {
		name   string
		obj    *unstructured.Unstructured
		newObj *unstructured.Unstructured
		ns     *corev1.Namespace
		schema *openapi.Schema
		denied bool
	}{
		{
			name: "default annotation to 'none' if not set on either namespace or resource",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "none",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
		},
		{
			name: "default annotation from namespace",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "none",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
					Annotations: map[string]string{
						managementconflict.FullyQualifiedAnnotation: "none",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
		},
		{
			name: "no defaulting needed since resource specifies annotation",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "resource",
						},
					},
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "resource",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
					Annotations: map[string]string{
						managementconflict.FullyQualifiedAnnotation: "none",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": &openapi.Schema{
						Type: "string",
					},
				},
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
			},
		},
		{
			name: "default annotation to 'none' if resource doesn't support management conflict prevention",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "none",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
				},
			},
			schema: &openapi.Schema{
				Type: "object",
			},
		},
		{
			name: "default annotation to 'none' if resource doesn't support management conflict prevention even if namespace has annotation",
			obj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
				},
			},
			newObj: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "test1.cnrm.cloud.google.com/v1alpha1",
					"kind":       "Test1Bar",
					"metadata": map[string]interface{}{
						"annotations": map[string]interface{}{
							managementconflict.FullyQualifiedAnnotation: "none",
						},
					},
				},
			},
			ns: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ns",
					Annotations: map[string]string{
						managementconflict.FullyQualifiedAnnotation: "resource",
					},
				},
			},
			schema: &openapi.Schema{
				Type: "object",
			},
		},
	}
	smLoader := testservicemetadataloader.NewForUnitTest()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dclSchemaMap := make(map[string]*openapi.Schema)
			dclSchemaMap["test1_beta_bar"] = tc.schema
			dclSchemaLoader := testdclschemaloader.New(dclSchemaMap)
			response := defaultManagementConflictAnnotationForDCLBasedResources(tc.obj, tc.ns, dclSchemaLoader, smLoader)
			expectedResponse := constructPatchResponse(tc.obj, tc.newObj)
			if !testutil.Equals(t, response, expectedResponse) {
				t.Fatalf("expect to get response %v, but got %v", expectedResponse, response)
			}
		})
	}
}

func TestRequestsResourcePolicy(t *testing.T) {
	tests := []struct {
		name          string
		objAnnotation string
		nsAnnotation  string
		want          bool
	}{
		{name: "object asks for resource", objAnnotation: "resource", want: true},
		{name: "object asks for none in a resource namespace", objAnnotation: "none", nsAnnotation: "resource", want: false},
		{name: "object without the annotation in a resource namespace", nsAnnotation: "resource", want: true},
		{name: "neither asks for resource", want: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obj := &unstructured.Unstructured{Object: map[string]any{}}
			if tc.objAnnotation != "" {
				obj.SetAnnotations(map[string]string{managementconflict.FullyQualifiedAnnotation: tc.objAnnotation})
			}
			ns := &corev1.Namespace{}
			if tc.nsAnnotation != "" {
				ns.SetAnnotations(map[string]string{managementconflict.FullyQualifiedAnnotation: tc.nsAnnotation})
			}

			got := requestsResourcePolicy(obj, ns)

			if got != tc.want {
				t.Errorf("requestsResourcePolicy(object policy %q, namespace policy %q) = %v, want %v", tc.objAnnotation, tc.nsAnnotation, got, tc.want)
			}
		})
	}
}

func TestManagementConflictAnnotationDefaulterForDirectReconciledObjects(t *testing.T) {
	const namespace = "test-ns"
	const policyPatchPath = "/metadata/annotations/cnrm.cloud.google.com~1management-conflict-prevention-policy"
	const arObjectPolicyWarning = "cnrm.cloud.google.com/management-conflict-prevention-policy=resource is not enforced: " +
		"ArtifactRegistryRepository objects reconciled by the Direct controller do not support resource leasing"
	const arNamespacePolicyWarning = "the namespace default cnrm.cloud.google.com/management-conflict-prevention-policy=resource is not applied: " +
		"ArtifactRegistryRepository objects reconciled by the Direct controller do not support resource leasing"
	const documentAINamespacePolicyWarning = "the namespace default cnrm.cloud.google.com/management-conflict-prevention-policy=resource is not applied: " +
		"DocumentAIProcessorDefaultVersion objects reconciled by the Direct controller do not support resource leasing"
	const arAPIVersion, arKind = "artifactregistry.cnrm.cloud.google.com/v1beta1", "ArtifactRegistryRepository"
	const documentAIAPIVersion, documentAIKind = "documentai.cnrm.cloud.google.com/v1alpha1", "DocumentAIProcessorDefaultVersion"
	const policy = managementconflict.FullyQualifiedAnnotation
	resourceNamespace := map[string]string{policy: "resource"}

	scheme := runtime.NewScheme()
	if err := operatorv1beta1.AddToScheme(scheme); err != nil {
		t.Fatalf("adding operator types to scheme: %v", err)
	}
	routedToDirect := fake.NewClientBuilder().WithScheme(scheme).WithObjects(
		&operatorv1beta1.ConfigConnector{
			ObjectMeta: metav1.ObjectMeta{Name: operatorv1beta1.ConfigConnectorAllowedName},
			Spec:       operatorv1beta1.ConfigConnectorSpec{Mode: "namespaced"},
		},
		&operatorv1beta1.ConfigConnectorContext{
			ObjectMeta: metav1.ObjectMeta{Namespace: namespace, Name: operatorv1beta1.ConfigConnectorContextAllowedName},
			Spec: operatorv1beta1.ConfigConnectorContextSpec{
				GoogleServiceAccount: "kcc@test-project.iam.gserviceaccount.com",
				Experiments: &operatorv1beta1.Experiments{ControllerOverrides: map[string]k8s.ReconcilerType{
					"ArtifactRegistryRepository.artifactregistry.cnrm.cloud.google.com": k8s.ReconcilerTypeDirect,
				}},
			},
		},
	).Build()
	unreadableRouting := fake.NewClientBuilder().WithInterceptorFuncs(interceptor.Funcs{
		Get: func(context.Context, client.WithWatch, client.ObjectKey, client.Object, ...client.GetOption) error {
			return apierrors.NewForbidden(schema.GroupResource{Group: "core.cnrm.cloud.google.com", Resource: "configconnectors"}, "", errors.New("RBAC: access denied"))
		},
	}).Build()

	tests := []struct {
		name       string
		reader     client.Reader
		apiVersion string
		kind       string
		objAnnots  map[string]string
		nsAnnots   map[string]string
		// wantPolicy is the policy that the response patch sets. "" means no patch.
		wantPolicy   string
		wantWarnings []string
	}{
		{
			name:         "Direct object that asks for resource",
			apiVersion:   arAPIVersion,
			kind:         arKind,
			objAnnots:    map[string]string{k8s.AlphaReconcilerAnnotation: "direct", policy: "resource"},
			wantWarnings: []string{arObjectPolicyWarning},
		},
		{
			name:         "Direct object in a resource namespace",
			apiVersion:   arAPIVersion,
			kind:         arKind,
			objAnnots:    map[string]string{k8s.AlphaReconcilerAnnotation: "direct"},
			nsAnnots:     resourceNamespace,
			wantWarnings: []string{arNamespacePolicyWarning},
		},
		{
			name:         "object routed to Direct by the namespace",
			reader:       routedToDirect,
			apiVersion:   arAPIVersion,
			kind:         arKind,
			nsAnnots:     resourceNamespace,
			wantWarnings: []string{arNamespacePolicyWarning},
		},
		{
			name:         "Direct-by-default object in a resource namespace",
			apiVersion:   documentAIAPIVersion,
			kind:         documentAIKind,
			nsAnnots:     resourceNamespace,
			wantWarnings: []string{documentAINamespacePolicyWarning},
		},
		{
			name:       "Terraform object in a resource namespace",
			apiVersion: arAPIVersion,
			kind:       arKind,
			nsAnnots:   resourceNamespace,
			wantPolicy: "resource",
		},
		{
			name:       "object in a resource namespace when the routing can't be read",
			reader:     unreadableRouting,
			apiVersion: arAPIVersion,
			kind:       arKind,
			nsAnnots:   resourceNamespace,
			wantPolicy: "resource",
		},
		{
			name:       "Direct object without a policy",
			apiVersion: arAPIVersion,
			kind:       arKind,
			objAnnots:  map[string]string{k8s.AlphaReconcilerAnnotation: "direct"},
			wantPolicy: "none",
		},
		{
			name:       "Direct object that asks for none in a resource namespace",
			apiVersion: arAPIVersion,
			kind:       arKind,
			objAnnots:  map[string]string{k8s.AlphaReconcilerAnnotation: "direct", policy: "none"},
			nsAnnots:   resourceNamespace,
		},
	}
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatalf("creating service mapping loader: %v", err)
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace, Annotations: tc.nsAnnots}}
			a := &managementConflictAnnotationDefaulter{
				client:                fake.NewClientBuilder().WithObjects(ns).Build(),
				reader:                tc.reader,
				smLoader:              smLoader,
				serviceMetadataLoader: testservicemetadataloader.NewForUnitTest(),
				tfResourceMap:         provider.ResourceMap(),
			}
			obj := &unstructured.Unstructured{Object: map[string]any{"spec": map[string]any{}}}
			obj.SetAPIVersion(tc.apiVersion)
			obj.SetKind(tc.kind)
			obj.SetNamespace(namespace)
			obj.SetName("test-object")
			// Start from a non-empty annotations map, so that a defaulted policy is a single "add" patch.
			annotations := map[string]string{k8s.ProjectIDAnnotation: "test-project"}
			maps.Copy(annotations, tc.objAnnots)
			obj.SetAnnotations(annotations)
			raw, err := obj.MarshalJSON()
			if err != nil {
				t.Fatalf("marshaling object: %v", err)
			}
			req := admission.Request{AdmissionRequest: admissionv1.AdmissionRequest{
				Operation: admissionv1.Create,
				Object:    runtime.RawExtension{Raw: raw},
			}}

			resp := a.Handle(t.Context(), req)

			if !resp.Allowed {
				t.Fatalf("Handle(create of %v) allowed = false, want true (result: %v)", tc.kind, resp.Result)
			}
			wantPatches := 0
			if tc.wantPolicy != "" {
				wantPatches = 1
			}
			if len(resp.Patches) != wantPatches {
				t.Fatalf("Handle(create of %v) patches = %+v, want %d patch(es)", tc.kind, resp.Patches, wantPatches)
			}
			if tc.wantPolicy != "" {
				p := resp.Patches[0]
				if p.Operation != "add" || p.Path != policyPatchPath || p.Value != tc.wantPolicy {
					t.Errorf("Handle(create of %v) patch = %+v, want add %v with value %q", tc.kind, p, policyPatchPath, tc.wantPolicy)
				}
			}
			if diff := cmp.Diff(tc.wantWarnings, resp.Warnings, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("Handle(create of %v) warnings mismatch (-want +got):\n%s", tc.kind, diff)
			}
		})
	}
}
