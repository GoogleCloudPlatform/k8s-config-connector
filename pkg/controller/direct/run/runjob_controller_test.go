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

package run

import (
	"context"
	"testing"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestRunJob_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-key")
	readyKey.SetNamespace("test-ns")
	selfLink := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-key"
	if err := unstructured.SetNestedField(readyKey.Object, selfLink, "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}
	if err := unstructured.SetNestedField(readyKey.Object, selfLink, "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey).Build()

	t.Run("resolves EncryptionKeyRef by name", func(t *testing.T) {
		obj := &krm.RunJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-ns",
			},
			Spec: krm.RunJobSpec{
				Template: &krm.ExecutionTemplate{
					Template: &krm.TaskTemplate{
						EncryptionKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							Name: "test-key",
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Template.Template.EncryptionKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Template.Template.EncryptionKeyRef.External, selfLink)
		}

		if err := ResolveRunJobRefs(ctx, reader, obj); err != nil {
			t.Fatalf("ResolveRunJobRefs() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := TaskTemplate_v1beta1_ToProto(mapCtx, obj.Spec.Template.Template)
		if mapCtx.Err() != nil {
			t.Fatalf("TaskTemplate_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.EncryptionKey != selfLink {
			t.Errorf("protoObj EncryptionKey = %q, want %q", protoObj.EncryptionKey, selfLink)
		}

		fromProto := TaskTemplate_v1beta1_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("TaskTemplate_v1beta1_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.EncryptionKeyRef == nil || fromProto.EncryptionKeyRef.External != selfLink {
			t.Errorf("fromProto EncryptionKeyRef.External = %v, want %q", fromProto.EncryptionKeyRef, selfLink)
		}
	})

	t.Run("handles external EncryptionKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.RunJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-ns",
			},
			Spec: krm.RunJobSpec{
				Template: &krm.ExecutionTemplate{
					Template: &krm.TaskTemplate{
						EncryptionKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
							External: externalKey,
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.Template.Template.EncryptionKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.Template.Template.EncryptionKeyRef.External, externalKey)
		}

		if err := ResolveRunJobRefs(ctx, reader, obj); err != nil {
			t.Fatalf("ResolveRunJobRefs() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := TaskTemplate_v1beta1_ToProto(mapCtx, obj.Spec.Template.Template)
		if mapCtx.Err() != nil {
			t.Fatalf("TaskTemplate_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.EncryptionKey != externalKey {
			t.Errorf("protoObj EncryptionKey = %q, want %q", protoObj.EncryptionKey, externalKey)
		}
	})
}

func TestRunJob_NormalizeReferences_SecretRef(t *testing.T) {
	ctx := context.Background()

	readySecret := &unstructured.Unstructured{}
	readySecret.SetGroupVersionKind(secretmanagerv1beta1.SecretManagerSecretGVK)
	readySecret.SetName("test-secret")
	readySecret.SetNamespace("test-ns")
	secretExternal := "projects/test-project/secrets/test-secret"
	if err := unstructured.SetNestedField(readySecret.Object, secretExternal, "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}

	readyVersion := &unstructured.Unstructured{}
	readyVersion.SetGroupVersionKind(secretmanagerv1beta1.SecretManagerSecretVersionGVK)
	readyVersion.SetName("test-version")
	readyVersion.SetNamespace("test-ns")
	versionExternal := "projects/test-project/secrets/test-secret/versions/1"
	if err := unstructured.SetNestedField(readyVersion.Object, versionExternal, "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readySecret, readyVersion).Build()

	t.Run("resolves SecretKeySelector secretRef and versionRef by name", func(t *testing.T) {
		obj := &krm.RunJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-ns",
			},
			Spec: krm.RunJobSpec{
				Template: &krm.ExecutionTemplate{
					Template: &krm.TaskTemplate{
						Containers: []krm.Container{
							{
								Env: []krm.EnvVar{
									{
										Name: direct.LazyPtr("MY_SECRET"),
										ValueSource: &krm.EnvVarSource{
											SecretKeyRef: &krm.SecretKeySelector{
												SecretRef: &secretmanagerv1beta1.SecretRef{
													Name: "test-secret",
												},
												VersionRef: &secretmanagerv1beta1.SecretVersionRef{
													Name: "test-version",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if err := ResolveRunJobRefs(ctx, reader, obj); err != nil {
			t.Fatalf("ResolveRunJobRefs() unexpected error: %v", err)
		}

		skr := obj.Spec.Template.Template.Containers[0].Env[0].ValueSource.SecretKeyRef
		if skr.SecretRef.External != secretExternal {
			t.Errorf("SecretRef.External = %q, want %q", skr.SecretRef.External, secretExternal)
		}
		if skr.VersionRef.External != "1" {
			t.Errorf("VersionRef.External = %q, want \"1\"", skr.VersionRef.External)
		}

		mapCtx := &direct.MapContext{}
		protoObj := TaskTemplate_v1beta1_ToProto(mapCtx, obj.Spec.Template.Template)
		if mapCtx.Err() != nil {
			t.Fatalf("TaskTemplate_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		protoSKR := protoObj.Containers[0].Env[0].GetValueSource().GetSecretKeyRef()
		if protoSKR.GetSecret() != secretExternal {
			t.Errorf("protoSKR.Secret = %q, want %q", protoSKR.GetSecret(), secretExternal)
		}
		if protoSKR.GetVersion() != "1" {
			t.Errorf("protoSKR.Version = %q, want \"1\"", protoSKR.GetVersion())
		}
	})

	t.Run("resolves SecretVolumeSource secretRef and versionRef by name", func(t *testing.T) {
		obj := &krm.RunJob{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-job",
				Namespace: "test-ns",
			},
			Spec: krm.RunJobSpec{
				Template: &krm.ExecutionTemplate{
					Template: &krm.TaskTemplate{
						Volumes: []krm.Volume{
							{
								Name: direct.LazyPtr("secret-vol"),
								Secret: &krm.SecretVolumeSource{
									SecretRef: &secretmanagerv1beta1.SecretRef{
										Name: "test-secret",
									},
									Items: []krm.VersionToPath{
										{
											Path: direct.LazyPtr("token"),
											VersionRef: &secretmanagerv1beta1.SecretVersionRef{
												Name: "test-version",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if err := ResolveRunJobRefs(ctx, reader, obj); err != nil {
			t.Fatalf("ResolveRunJobRefs() unexpected error: %v", err)
		}

		sv := obj.Spec.Template.Template.Volumes[0].Secret
		if sv.SecretRef.External != secretExternal {
			t.Errorf("SecretRef.External = %q, want %q", sv.SecretRef.External, secretExternal)
		}
		if sv.Items[0].VersionRef.External != "1" {
			t.Errorf("Items[0].VersionRef.External = %q, want \"1\"", sv.Items[0].VersionRef.External)
		}

		mapCtx := &direct.MapContext{}
		protoObj := TaskTemplate_v1beta1_ToProto(mapCtx, obj.Spec.Template.Template)
		if mapCtx.Err() != nil {
			t.Fatalf("TaskTemplate_v1beta1_ToProto unexpected error: %v", mapCtx.Err())
		}
		protoSecretVol := protoObj.Volumes[0].GetSecret()
		if protoSecretVol.GetSecret() != secretExternal {
			t.Errorf("protoSecretVol.Secret = %q, want %q", protoSecretVol.GetSecret(), secretExternal)
		}
		if protoSecretVol.GetItems()[0].GetVersion() != "1" {
			t.Errorf("protoSecretVol.Items[0].Version = %q, want \"1\"", protoSecretVol.GetItems()[0].GetVersion())
		}
	})
}

