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

// +tool:controller
// proto.service: google.cloud.kms.v1.KeyManagementService
// crd.type: KMSSecretCiphertext
// crd.version: v1alpha1

package kms

import (
	"context"
	"encoding/base64"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.KMSSecretCiphertextGVK, NewKMSSecretCiphertextModel, registry.CannotBeDeleted())
}

func NewKMSSecretCiphertextModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &kmsSecretCiphertextModel{config: *config}, nil
}

var _ directbase.Model = &kmsSecretCiphertextModel{}

type kmsSecretCiphertextModel struct {
	config config.ControllerConfig
}

func (m *kmsSecretCiphertextModel) client(ctx context.Context) (*kms.KeyManagementClient, error) {
	var opts []option.ClientOption

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := kms.NewKeyManagementRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building kms keyManagement client: %w", err)
	}

	return gcpClient, nil
}

func (m *kmsSecretCiphertextModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.KMSSecretCiphertext{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if err := obj.Spec.Plaintext.NormalizeSecret(ctx, "spec.plaintext", obj.GetNamespace(), reader); err != nil {
		return nil, fmt.Errorf("normalizing spec.plaintext: %w", err)
	}
	if obj.Spec.AdditionalAuthenticatedData != nil {
		if err := obj.Spec.AdditionalAuthenticatedData.NormalizeSecret(ctx, "spec.additionalAuthenticatedData", obj.GetNamespace(), reader); err != nil {
			return nil, fmt.Errorf("normalizing spec.additionalAuthenticatedData: %w", err)
		}
	}

	resolvedID, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := resolvedID.(*krm.KMSSecretCiphertextIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &kmsSecretCiphertextAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *kmsSecretCiphertextModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type kmsSecretCiphertextAdapter struct {
	gcpClient *kms.KeyManagementClient
	id        *krm.KMSSecretCiphertextIdentity
	desired   *krm.KMSSecretCiphertext
	actual    *kmspb.CryptoKey
	reader    client.Reader
}

var _ directbase.Adapter = &kmsSecretCiphertextAdapter{}

func (a *kmsSecretCiphertextAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding KMSSecretCiphertext", "name", a.id.String())

	// Rule 2: Find() Pre-Check Guard
	if a.id.Ciphertext == "" {
		return false, nil
	}

	req := &kmspb.GetCryptoKeyRequest{
		Name: a.id.ParentString(),
	}
	cryptoKey, err := a.gcpClient.GetCryptoKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CryptoKey %q: %w", a.id.ParentString(), err)
	}
	a.actual = cryptoKey
	return true, nil
}

func (a *kmsSecretCiphertextAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating KMSSecretCiphertext", "name", a.id.String())

	if a.desired.Spec.Plaintext.Value == nil {
		return fmt.Errorf("spec.plaintext is required")
	}
	plaintextBytes := []byte(*a.desired.Spec.Plaintext.Value)

	req := &kmspb.EncryptRequest{
		Name:      a.id.ParentString(),
		Plaintext: plaintextBytes,
	}

	if a.desired.Spec.AdditionalAuthenticatedData != nil && a.desired.Spec.AdditionalAuthenticatedData.Value != nil {
		req.AdditionalAuthenticatedData = []byte(*a.desired.Spec.AdditionalAuthenticatedData.Value)
	}

	resp, err := a.gcpClient.Encrypt(ctx, req)
	if err != nil {
		return fmt.Errorf("encrypting with KMS CryptoKey %s: %w", a.id.ParentString(), err)
	}

	ciphertextStr := base64.StdEncoding.EncodeToString(resp.Ciphertext)

	// Since ciphertext is server-generated, write it back to spec.resourceID if not set.
	if a.desired.Spec.ResourceID == nil || *a.desired.Spec.ResourceID == "" {
		if err := createOp.SetSpecResourceID(ctx, ciphertextStr); err != nil {
			return err
		}
	}

	// Fetch CryptoKey to confirm existence and align with golden logs
	getReq := &kmspb.GetCryptoKeyRequest{
		Name: a.id.ParentString(),
	}
	if _, err := a.gcpClient.GetCryptoKey(ctx, getReq); err != nil {
		return fmt.Errorf("getting CryptoKey %s: %w", a.id.ParentString(), err)
	}

	status := &krm.KMSSecretCiphertextStatus{
		Ciphertext: direct.LazyPtr(ciphertextStr),
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *kmsSecretCiphertextAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KMSSecretCiphertext", "name", a.id.String())

	if a.actual == nil {
		return fmt.Errorf("actual CryptoKey is missing in Update")
	}

	if a.desired.Spec.CryptoKey != a.id.ParentString() {
		return fmt.Errorf("spec.cryptoKey is immutable and cannot be updated")
	}
	if a.desired.Spec.ResourceID != nil && *a.desired.Spec.ResourceID != "" && *a.desired.Spec.ResourceID != a.id.Ciphertext {
		return fmt.Errorf("spec.resourceID is immutable and cannot be updated")
	}

	status := &krm.KMSSecretCiphertextStatus{
		Ciphertext: direct.LazyPtr(a.id.Ciphertext),
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *kmsSecretCiphertextAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *kmsSecretCiphertextAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("No-op Delete for KMSSecretCiphertext", "name", a.id.String())
	return true, nil
}
