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

// +tool:controller
// crd.type: ComputeBackendBucketSignedURLKey
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"
	"slices"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.ComputeBackendBucketSignedURLKeyGVK, NewBackendBucketSignedURLKeyModel)
}

func NewBackendBucketSignedURLKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &backendBucketSignedURLKeyModel{config: config}, nil
}

var _ directbase.Model = &backendBucketSignedURLKeyModel{}

type backendBucketSignedURLKeyModel struct {
	config *config.ControllerConfig
}

func (m *backendBucketSignedURLKeyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeBackendBucketSignedURLKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBackendBucketSignedURLKeyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Resolve the key value from the spec (inline value or secretKeyRef).
	// Key value is only needed for Create; skip resolution during deletion to
	// avoid blocking on a missing secret when the parent resources are gone.
	var resolvedKeyValue string
	if u.GetDeletionTimestamp() == nil {
		keyValue, err := obj.Spec.KeyValue.ReadSecretValue(ctx, "spec.keyValue", obj.GetNamespace(), reader)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.keyValue: %w", err)
		}
		if keyValue == nil || *keyValue == "" {
			return nil, fmt.Errorf("spec.keyValue resolved to an empty string")
		}
		resolvedKeyValue = *keyValue
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	bbClient, err := gcpClient.newBackendBucketsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &BackendBucketSignedURLKeyAdapter{
		gcpClient: bbClient,
		id:        id,
		desired:   obj,
		keyValue:  resolvedKeyValue,
	}, nil
}

func (m *backendBucketSignedURLKeyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BackendBucketSignedURLKeyAdapter struct {
	gcpClient *compute.BackendBucketsClient
	id        *krm.BackendBucketSignedURLKeyIdentity
	desired   *krm.ComputeBackendBucketSignedURLKey
	// keyValue is the resolved (plaintext) key value, used only during Create.
	keyValue string
	// found indicates whether the key was found in GCP during Find.
	found bool
}

var _ directbase.Adapter = &BackendBucketSignedURLKeyAdapter{}

// Find retrieves the GCP resource.
// Returns true if found (triggers Update), false if not found (triggers Create).
// Since GCP never returns the key value, we look for the key name in the bucket's CDN policy.
func (a *BackendBucketSignedURLKeyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding BackendBucketSignedURLKey", "id", a.id)

	req := &computepb.GetBackendBucketRequest{
		Project:       a.id.Parent().ProjectID,
		BackendBucket: a.id.Parent().BackendBucket,
	}
	bucket, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackendBucket %s: %w", a.id.Parent().BackendBucket, err)
	}

	if bucket.GetCdnPolicy() == nil {
		return false, nil
	}
	keyNames := bucket.GetCdnPolicy().GetSignedUrlKeyNames()
	a.found = slices.Contains(keyNames, a.id.KeyName())
	return a.found, nil
}

// Create adds the signed URL key to the backend bucket.
func (a *BackendBucketSignedURLKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BackendBucketSignedURLKey", "id", a.id)

	keyName := a.id.KeyName()
	keyValue := a.keyValue
	req := &computepb.AddSignedUrlKeyBackendBucketRequest{
		Project:       a.id.Parent().ProjectID,
		BackendBucket: a.id.Parent().BackendBucket,
		SignedUrlKeyResource: &computepb.SignedUrlKey{
			KeyName:  &keyName,
			KeyValue: &keyValue,
		},
	}

	op, err := a.gcpClient.AddSignedUrlKey(ctx, req)
	if err != nil {
		return fmt.Errorf("adding BackendBucketSignedURLKey %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for BackendBucketSignedURLKey %s create: %w", a.id, err)
	}

	log.Info("successfully added BackendBucketSignedURLKey in GCP", "id", a.id)

	status := &krm.ComputeBackendBucketSignedURLKeyStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update is a no-op: the key value is immutable and write-only in GCP.
// A key cannot be updated in-place; to rotate a key, delete and recreate it.
func (a *BackendBucketSignedURLKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("BackendBucketSignedURLKey is immutable; no update needed", "id", a.id)

	status := &krm.ComputeBackendBucketSignedURLKeyStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource spec.
// Note: keyValue is write-only in GCP and cannot be exported.
func (a *BackendBucketSignedURLKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if !a.found {
		return nil, fmt.Errorf("Find() not called or key not found")
	}

	obj := &krm.ComputeBackendBucketSignedURLKey{}
	obj.Spec.ProjectRef = &refv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.BackendBucketRef = krm.BackendBucketRef{External: a.id.Parent().BackendBucket}
	keyName := a.id.KeyName()
	obj.Spec.ResourceID = &keyName
	// keyValue is intentionally left empty — it cannot be retrieved from GCP.

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{}
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeBackendBucketSignedURLKeyGVK)
	u.Object = uObj
	return u, nil
}

// Delete removes the signed URL key from the backend bucket.
func (a *BackendBucketSignedURLKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BackendBucketSignedURLKey", "id", a.id)

	// Confirm the key still exists before attempting deletion.
	found, err := a.Find(ctx)
	if err != nil {
		return false, fmt.Errorf("finding BackendBucketSignedURLKey %s before delete: %w", a.id, err)
	}
	if !found {
		log.V(2).Info("BackendBucketSignedURLKey not found, presuming already deleted", "id", a.id)
		return false, nil
	}

	keyName := a.id.KeyName()
	req := &computepb.DeleteSignedUrlKeyBackendBucketRequest{
		Project:       a.id.Parent().ProjectID,
		BackendBucket: a.id.Parent().BackendBucket,
		KeyName:       keyName,
	}

	op, err := a.gcpClient.DeleteSignedUrlKey(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting BackendBucketSignedURLKey %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for BackendBucketSignedURLKey %s delete: %w", a.id, err)
	}

	log.Info("successfully deleted BackendBucketSignedURLKey from GCP", "id", a.id)
	return true, nil
}
