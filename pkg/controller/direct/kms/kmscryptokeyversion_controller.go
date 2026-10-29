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
// proto.message: google.cloud.kms.v1.CryptoKeyVersion
// crd.type: KMSCryptoKeyVersion
// crd.version: v1alpha1

package kms

import (
	"context"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.KMSCryptoKeyVersionGVK, NewKMSCryptoKeyVersionModel)
}

func NewKMSCryptoKeyVersionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &kmsCryptoKeyVersionModel{config: *config}, nil
}

var _ directbase.Model = &kmsCryptoKeyVersionModel{}

type kmsCryptoKeyVersionModel struct {
	config config.ControllerConfig
}

func (m *kmsCryptoKeyVersionModel) client(ctx context.Context) (*kms.KeyManagementClient, error) {
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	return gcpClient.newKeyManagementClient(ctx)
}

func (m *kmsCryptoKeyVersionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.KMSCryptoKeyVersion{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.KMSCryptoKeyVersionIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &kmsCryptoKeyVersionAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *kmsCryptoKeyVersionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.KMSCryptoKeyVersionIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	if !id.HasIdentitySpecified() {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &kmsCryptoKeyVersionAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type kmsCryptoKeyVersionAdapter struct {
	gcpClient *kms.KeyManagementClient
	id        *krm.KMSCryptoKeyVersionIdentity
	desired   *krm.KMSCryptoKeyVersion
	actual    *kmspb.CryptoKeyVersion
	reader    client.Reader
}

var _ directbase.Adapter = &kmsCryptoKeyVersionAdapter{}

func (a *kmsCryptoKeyVersionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting kms cryptokeyversion", "name", a.id)

	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &kmspb.GetCryptoKeyVersionRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCryptoKeyVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting kms cryptokeyversion %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *kmsCryptoKeyVersionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating kms cryptokeyversion", "parent", a.id.ParentString())

	mapCtx := &direct.MapContext{}
	desiredPb := KMSCryptoKeyVersionSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Rule 3 for service-generated ID: leave desiredPb.Name empty.
	desiredPb.Name = ""

	req := &kmspb.CreateCryptoKeyVersionRequest{
		Parent:           a.id.ParentString(),
		CryptoKeyVersion: desiredPb,
	}
	created, err := a.gcpClient.CreateCryptoKeyVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating kms cryptokeyversion in parent %s: %w", a.id.ParentString(), err)
	}
	log.V(2).Info("successfully created kms cryptokeyversion in gcp", "name", created.GetName())

	return a.updateStatus(ctx, createOp, created)
}

func (a *kmsCryptoKeyVersionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KMSCryptoKeyVersion", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredPb := KMSCryptoKeyVersionSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	diffs, updateMask, err := compareCryptoKeyVersion(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for KMSCryptoKeyVersion", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &kmspb.UpdateCryptoKeyVersionRequest{
		CryptoKeyVersion: desiredPb,
		UpdateMask:       updateMask,
	}

	updated, err := a.gcpClient.UpdateCryptoKeyVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("updating KMSCryptoKeyVersion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated KMSCryptoKeyVersion", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *kmsCryptoKeyVersionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.KMSCryptoKeyVersion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(KMSCryptoKeyVersionSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.CryptoKey = a.id.ParentString()
	obj.Spec.ResourceID = direct.PtrTo(a.id.CryptoKeyVersion)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.CryptoKeyVersion)
	u.SetGroupVersionKind(krm.KMSCryptoKeyVersionGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *kmsCryptoKeyVersionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("Deleting KMSCryptoKeyVersion", "name", a.id)

	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	if a.actual != nil && (a.actual.State == kmspb.CryptoKeyVersion_DESTROYED || a.actual.State == kmspb.CryptoKeyVersion_DESTROY_SCHEDULED) {
		return true, nil
	}

	destroyReq := &kmspb.DestroyCryptoKeyVersionRequest{Name: a.id.String()}
	if _, err := a.gcpClient.DestroyCryptoKeyVersion(ctx, destroyReq); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("destroying kms cryptokeyversion %q: %w", a.id.String(), err)
	}

	return true, nil
}

func compareCryptoKeyVersion(ctx context.Context, actual, desired *kmspb.CryptoKeyVersion) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	diff := &structuredreporting.Diff{}
	var paths []string

	if desired.State != kmspb.CryptoKeyVersion_CRYPTO_KEY_VERSION_STATE_UNSPECIFIED && desired.State != actual.State {
		diff.AddField("state", actual.State.String(), desired.State.String())
		paths = append(paths, "state")
	}

	updateMask := &fieldmaskpb.FieldMask{Paths: paths}
	return diff, updateMask, nil
}

func (a *kmsCryptoKeyVersionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *kmspb.CryptoKeyVersion) error {
	mapCtx := &direct.MapContext{}
	status := KMSCryptoKeyVersionStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
