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
// proto.message: google.cloud.kms.v1.CryptoKey
// crd.type: KMSCryptoKey
// crd.version: v1beta1

package kms

import (
	"context"
	"fmt"
	"strings"

	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
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
	registry.RegisterModel(krm.KMSCryptoKeyGVK, NewCryptoKeyModel)
}

func NewCryptoKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &cryptoKeyModel{config: *config}, nil
}

var _ directbase.Model = &cryptoKeyModel{}

type cryptoKeyModel struct {
	config config.ControllerConfig
}

func (m *cryptoKeyModel) client(ctx context.Context, projectID string) (*kms.KeyManagementClient, error) {
	var opts []option.ClientOption

	config := m.config

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := kms.NewKeyManagementRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building kms cryptokey client: %w", err)
	}

	return gcpClient, err
}

func (m *cryptoKeyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.KMSCryptoKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewKMSCryptoKeyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().Parent.ProjectID)
	if err != nil {
		return nil, err
	}

	return &cryptoKeyAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *cryptoKeyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support AdpaterForURL
	return nil, nil
}

type cryptoKeyAdapter struct {
	gcpClient *kms.KeyManagementClient
	id        *krm.KMSCryptoKeyIdentity
	desired   *krm.KMSCryptoKey
	actual    *kmspb.CryptoKey
	reader    client.Reader
}

var _ directbase.Adapter = &cryptoKeyAdapter{}

func (a *cryptoKeyAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.ID() == "" {
		return false, nil
	}

	req := &kmspb.GetCryptoKeyRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetCryptoKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting kms cryptokey %s: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *cryptoKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating KMSCryptoKey", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := KMSCryptoKeySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &kmspb.CreateCryptoKeyRequest{
		Parent:      a.id.Parent().String(),
		CryptoKeyId: a.id.ID(),
		CryptoKey:   desired,
	}
	if a.desired.Spec.SkipInitialVersionCreation != nil {
		req.SkipInitialVersionCreation = *a.desired.Spec.SkipInitialVersionCreation
	}

	created, err := a.gcpClient.CreateCryptoKey(ctx, req)
	if err != nil {
		return fmt.Errorf("creating kms cryptokey %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created kms cryptokey in gcp", "name", a.id)

	status := &krm.KMSCryptoKeyStatus{}
	status.ObservedState = KMSCryptoKeyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.PtrTo(a.id.String())
	status.SelfLink = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *cryptoKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KMSCryptoKey", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := KMSCryptoKeySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	if a.desired.Spec.Purpose == nil {
		desiredPb.Purpose = a.actual.Purpose
	}
	if a.desired.Spec.ImportOnly == nil {
		desiredPb.ImportOnly = a.actual.ImportOnly
	}
	if a.desired.Spec.DestroyScheduledDuration == nil {
		desiredPb.DestroyScheduledDuration = a.actual.DestroyScheduledDuration
	}
	if a.desired.Spec.VersionTemplate == nil {
		desiredPb.VersionTemplate = a.actual.VersionTemplate
	} else {
		if a.desired.Spec.VersionTemplate.ProtectionLevel == nil && a.actual.VersionTemplate != nil && desiredPb.VersionTemplate != nil {
			desiredPb.VersionTemplate.ProtectionLevel = a.actual.VersionTemplate.ProtectionLevel
		}
		if a.desired.Spec.VersionTemplate.Algorithm == nil && a.actual.VersionTemplate != nil && desiredPb.VersionTemplate != nil {
			desiredPb.VersionTemplate.Algorithm = a.actual.VersionTemplate.Algorithm
		}
	}

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	// Purpose is immutable
	paths.Delete("purpose")
	// ImportOnly is immutable
	paths.Delete("import_only")
	// DestroyScheduledDuration is immutable
	paths.Delete("destroy_scheduled_duration")
	// NextRotationTime is not supported for update through UpdateCryptoKey
	paths.Delete("next_rotation_time")

	// Primary is output only
	for path := range paths {
		if path == "primary" || strings.HasPrefix(path, "primary.") {
			paths.Delete(path)
		}
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	updateMask := &fieldmaskpb.FieldMask{Paths: paths.UnsortedList()}
	req := &kmspb.UpdateCryptoKeyRequest{
		CryptoKey:  desiredPb,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateCryptoKey(ctx, req)
	if err != nil {
		return fmt.Errorf("updating kms cryptokey %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated kms cryptokey in gcp", "name", a.id)

	status := &krm.KMSCryptoKeyStatus{}
	status.ObservedState = KMSCryptoKeyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.PtrTo(a.id.String())
	status.SelfLink = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *cryptoKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.KMSCryptoKey{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(KMSCryptoKeySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.KeyRingRef = &kmsv1beta1.KMSKeyRingRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.KMSCryptoKeyGVK)

	return u, nil
}

func (a *cryptoKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("No-op Delete for crypto key", "name", a.id.String())
	// KMS API does not support deleting a crypto key.
	// Return success to remove the finalizer, so the resource can be deleted in k8s.
	return true, nil
}
