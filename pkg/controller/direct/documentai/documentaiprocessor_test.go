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

package documentai

import (
	"context"
	"net/http"
	"testing"

	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDocumentAIProcessor_KmsKeyRef_Mapping(t *testing.T) {
	mapCtx := &direct.MapContext{}

	// Test ToProto
	spec := &krm.DocumentAIProcessorSpec{
		Type:        direct.LazyPtr("OCR_PROCESSOR"),
		DisplayName: direct.LazyPtr("test-processor"),
		KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
			External: "projects/test-project/locations/us/keyRings/test-ring/cryptoKeys/test-key",
		},
	}
	proto := DocumentAIProcessorSpec_v1alpha1_ToProto(mapCtx, spec)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in ToProto: %v", err)
	}
	if got, want := proto.GetKmsKeyName(), "projects/test-project/locations/us/keyRings/test-ring/cryptoKeys/test-key"; got != want {
		t.Errorf("proto.GetKmsKeyName() = %q, want %q", got, want)
	}
	if got, want := proto.GetType(), "OCR_PROCESSOR"; got != want {
		t.Errorf("proto.GetType() = %q, want %q", got, want)
	}
	if got, want := proto.GetDisplayName(), "test-processor"; got != want {
		t.Errorf("proto.GetDisplayName() = %q, want %q", got, want)
	}

	// Test FromProto
	protoObj := &pb.Processor{
		Type:        "OCR_PROCESSOR",
		DisplayName: "test-processor",
		KmsKeyName:  "projects/test-project/locations/us/keyRings/test-ring/cryptoKeys/test-key",
	}
	specFromProto := DocumentAIProcessorSpec_v1alpha1_FromProto(mapCtx, protoObj)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in FromProto: %v", err)
	}
	if specFromProto.KmsKeyRef == nil {
		t.Fatalf("expected KmsKeyRef to be non-nil")
	}
	if got, want := specFromProto.KmsKeyRef.External, "projects/test-project/locations/us/keyRings/test-ring/cryptoKeys/test-key"; got != want {
		t.Errorf("specFromProto.KmsKeyRef.External = %q, want %q", got, want)
	}
}

func TestDocumentAIProcessor_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	if err := krm.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme failed: %v", err)
	}
	if err := kmsv1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme failed: %v", err)
	}

	kmsKey := &kmsv1beta1.KMSCryptoKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-key",
			Namespace: "test-ns",
		},
		Status: kmsv1beta1.KMSCryptoKeyStatus{
			SelfLink: direct.LazyPtr("projects/test-project/locations/us/keyRings/test-ring/cryptoKeys/test-key"),
		},
	}

	client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(kmsKey).Build()

	processor := &krm.DocumentAIProcessor{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-processor",
			Namespace: "test-ns",
		},
		Spec: krm.DocumentAIProcessorSpec{
			Location: "us",
			KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
				Name: "test-key",
			},
		},
	}

	if err := common.NormalizeReferences(ctx, client, processor, nil); err != nil {
		t.Fatalf("NormalizeReferences failed: %v", err)
	}

	if got, want := processor.Spec.KmsKeyRef.External, "projects/test-project/locations/us/keyRings/test-ring/cryptoKeys/test-key"; got != want {
		t.Errorf("normalized KMSKeyRef.External = %q, want %q", got, want)
	}
}

func TestDocumentAIProcessor_AdapterForObject_NoBlockingOnMissingReferences(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	if err := krm.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme failed: %v", err)
	}
	if err := kmsv1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme failed: %v", err)
	}

	// Fake client without any KMSCryptoKey objects
	client := fake.NewClientBuilder().WithScheme(scheme).Build()

	processor := &krm.DocumentAIProcessor{
		TypeMeta: metav1.TypeMeta{
			APIVersion: krm.DocumentAIProcessorGVK.GroupVersion().String(),
			Kind:       krm.DocumentAIProcessorGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-processor",
			Namespace: "test-ns",
		},
		Spec: krm.DocumentAIProcessorSpec{
			Location: "us",
			KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
				Name: "non-existent-key",
			},
		},
	}
	processor.Spec.ProjectRef = &refs.ProjectRef{
		External: "test-project",
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(processor)
	if err != nil {
		t.Fatalf("ToUnstructured failed: %v", err)
	}
	u := &unstructured.Unstructured{Object: uObj}

	model, err := NewProcessorModel(ctx, &config.ControllerConfig{
		HTTPClient: &http.Client{},
	})
	if err != nil {
		t.Fatalf("NewProcessorModel failed: %v", err)
	}

	op := &directbase.AdapterForObjectOperation{
		Reader: client,
		Object: u,
	}

	// AdapterForObject must succeed even if the referenced KMSCryptoKey does not exist
	// (ensuring Deletion is not blocked by missing referenced resources).
	adapter, err := model.AdapterForObject(ctx, op)
	if err != nil {
		t.Fatalf("AdapterForObject failed unexpectedly: %v", err)
	}
	if adapter == nil {
		t.Fatalf("expected non-nil adapter")
	}

	// However, calling Create on the adapter should fail during reference normalization
	createOp := &directbase.CreateOperation{}
	err = adapter.Create(ctx, createOp)
	if err == nil {
		t.Fatalf("expected Create to fail due to missing KMS key reference normalization, but got nil")
	}

	// When the key exists, reference normalization succeeds in the adapter
	kmsKey := &kmsv1beta1.KMSCryptoKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-key",
			Namespace: "test-ns",
		},
		Status: kmsv1beta1.KMSCryptoKeyStatus{
			SelfLink: direct.LazyPtr("projects/test-project/locations/us/keyRings/test-ring/cryptoKeys/test-key"),
		},
	}
	clientWithKey := fake.NewClientBuilder().WithScheme(scheme).WithObjects(kmsKey).Build()
	processorWithKey := processor.DeepCopy()
	processorWithKey.Spec.KmsKeyRef.Name = "test-key"

	uObjWithKey, err := runtime.DefaultUnstructuredConverter.ToUnstructured(processorWithKey)
	if err != nil {
		t.Fatalf("ToUnstructured failed: %v", err)
	}
	opWithKey := &directbase.AdapterForObjectOperation{
		Reader: clientWithKey,
		Object: &unstructured.Unstructured{Object: uObjWithKey},
	}
	adapterWithKey, err := model.AdapterForObject(ctx, opWithKey)
	if err != nil {
		t.Fatalf("AdapterForObject failed: %v", err)
	}
	procAdapter, ok := adapterWithKey.(*ProcessorAdapter)
	if !ok {
		t.Fatalf("expected *ProcessorAdapter, got %T", adapterWithKey)
	}
	if err := common.NormalizeReferences(ctx, procAdapter.reader, procAdapter.desired, nil); err != nil {
		t.Fatalf("NormalizeReferences failed: %v", err)
	}
	if got, want := procAdapter.desired.Spec.KmsKeyRef.External, "projects/test-project/locations/us/keyRings/test-ring/cryptoKeys/test-key"; got != want {
		t.Errorf("normalized KMSKeyRef.External = %q, want %q", got, want)
	}
}
