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
// proto.message: google.cloud.kms.v1.KeyRing
// crd.type: KMSKeyRing
// crd.version: v1beta1

package kms

import (
	"context"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.KMSKeyRingGVK, NewKeyRingModel)
}

func NewKeyRingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &keyRingModel{config: *config}, nil
}

var _ directbase.Model = &keyRingModel{}

type keyRingModel struct {
	config config.ControllerConfig
}

func (m *keyRingModel) client(ctx context.Context, projectID string) (*kms.KeyManagementClient, error) {
	var opts []option.ClientOption

	config := m.config

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := kms.NewKeyManagementRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building kms keyring client: %w", err)
	}

	return gcpClient, err
}

func (m *keyRingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.KMSKeyRing{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	resolvedID, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := resolvedID.(*krm.KMSKeyRingIdentity)

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &keyRingAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *keyRingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type keyRingAdapter struct {
	gcpClient *kms.KeyManagementClient
	id        *krm.KMSKeyRingIdentity
	desired   *krm.KMSKeyRing
	actual    *kmspb.KeyRing
	reader    client.Reader
}

var _ directbase.Adapter = &keyRingAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *keyRingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting kms keyring", "name", a.id)

	req := &kmspb.GetKeyRingRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetKeyRing(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting kms keyring %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *keyRingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating kms keyring", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := KMSKeyRingSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &kmspb.CreateKeyRingRequest{
		Parent:    a.id.ParentString(),
		KeyRingId: a.id.Keyring,
		KeyRing:   desired,
	}
	created, err := a.gcpClient.CreateKeyRing(ctx, req)
	if err != nil {
		return fmt.Errorf("creating kms keyring %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created kms keyring in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *keyRingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KMSKeyRing", "name", a.id)

	// KeyRing is immutable in GCP, no fields can be updated.
	return a.updateStatus(ctx, updateOp, a.actual)
}

// Export implements the Adapter interface.
func (a *keyRingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.KMSKeyRing{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(KMSKeyRingSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = direct.PtrTo(a.id.Location)
	obj.Spec.ResourceID = direct.PtrTo(a.id.Keyring)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Keyring)
	u.SetGroupVersionKind(krm.KMSKeyRingGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *keyRingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("No-op Delete for key ring", "name", a.id.String())
	// KMS API does not support deleting key rings.
	// Return success to remove the finalizer, so the resource can be deleted in k8s.
	return true, nil
}

func (a *keyRingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *kmspb.KeyRing) error {
	mapCtx := &direct.MapContext{}
	status := KMSKeyRingStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
