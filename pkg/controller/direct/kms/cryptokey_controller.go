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
	"errors"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.KMSCryptoKeyGVK, NewKMSCryptoKeyModel)
}

func NewKMSCryptoKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &kmsCryptoKeyModel{config: *config}, nil
}

var _ directbase.Model = &kmsCryptoKeyModel{}

type kmsCryptoKeyModel struct {
	config config.ControllerConfig
}

func (m *kmsCryptoKeyModel) client(ctx context.Context) (*kms.KeyManagementClient, error) {
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	return gcpClient.newKeyManagementClient(ctx)
}

func (m *kmsCryptoKeyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.KMSCryptoKey{}
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
	id := idRaw.(*krm.KMSCryptoKeyIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := KMSCryptoKeySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	skipInitialVersionCreation := false
	if obj.Spec.SkipInitialVersionCreation != nil {
		skipInitialVersionCreation = *obj.Spec.SkipInitialVersionCreation
	}

	return &kmsCryptoKeyAdapter{
		gcpClient:                         gcpClient,
		id:                                id,
		desired:                           desired,
		desiredSkipInitialVersionCreation: skipInitialVersionCreation,
		reader:                            reader,
	}, nil
}

func (m *kmsCryptoKeyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type kmsCryptoKeyAdapter struct {
	gcpClient                         *kms.KeyManagementClient
	id                                *krm.KMSCryptoKeyIdentity
	desired                           *kmspb.CryptoKey
	desiredSkipInitialVersionCreation bool
	actual                            *kmspb.CryptoKey
	reader                            client.Reader
}

var _ directbase.Adapter = &kmsCryptoKeyAdapter{}

func (a *kmsCryptoKeyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting kms cryptokey", "name", a.id)

	req := &kmspb.GetCryptoKeyRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCryptoKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting kms cryptokey %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *kmsCryptoKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating kms cryptokey", "name", a.id)

	req := &kmspb.CreateCryptoKeyRequest{
		Parent:                     a.id.ParentString(),
		CryptoKeyId:                a.id.CryptoKey,
		CryptoKey:                  a.desired,
		SkipInitialVersionCreation: a.desiredSkipInitialVersionCreation,
	}
	created, err := a.gcpClient.CreateCryptoKey(ctx, req)
	if err != nil {
		return fmt.Errorf("creating kms cryptokey %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created kms cryptokey in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *kmsCryptoKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KMSCryptoKey", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareCryptoKey(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for KMSCryptoKey", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &kmspb.UpdateCryptoKeyRequest{
		CryptoKey:  a.desired,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateCryptoKey(ctx, req)
	if err != nil {
		return fmt.Errorf("updating KMSCryptoKey %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated KMSCryptoKey", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *kmsCryptoKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.KMSCryptoKey{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(KMSCryptoKeySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	keyRingExternal := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s", a.id.Project, a.id.Location, a.id.KeyRing)
	obj.Spec.KeyRingRef = krm.KMSKeyRingRef{External: keyRingExternal}

	if len(a.actual.Labels) > 0 {
		obj.SetLabels(a.actual.Labels)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.CryptoKey)
	u.SetGroupVersionKind(krm.KMSCryptoKeyGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *kmsCryptoKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("Deleting KMSCryptoKey versions", "name", a.id)

	it := a.gcpClient.ListCryptoKeyVersions(ctx, &kmspb.ListCryptoKeyVersionsRequest{Parent: a.id.String()})
	for {
		version, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return false, fmt.Errorf("listing crypto key versions: %w", err)
		}
		if version.State != kmspb.CryptoKeyVersion_DESTROYED && version.State != kmspb.CryptoKeyVersion_DESTROY_SCHEDULED {
			destroyReq := &kmspb.DestroyCryptoKeyVersionRequest{Name: version.Name}
			if _, err := a.gcpClient.DestroyCryptoKeyVersion(ctx, destroyReq); err != nil {
				return false, fmt.Errorf("destroying crypto key version %q: %w", version.Name, err)
			}
		}
	}

	if a.actual != nil && a.actual.GetRotationPeriod() != nil {
		desiredPb := &kmspb.CryptoKey{
			Name:             a.id.String(),
			RotationSchedule: &kmspb.CryptoKey_RotationPeriod{RotationPeriod: nil},
		}
		updateMask, err := fieldmaskpb.New(desiredPb, "rotation_period")
		if err != nil {
			return false, fmt.Errorf("creating update mask for rotation_period: %w", err)
		}
		req := &kmspb.UpdateCryptoKeyRequest{
			CryptoKey:  desiredPb,
			UpdateMask: updateMask,
		}
		if _, err := a.gcpClient.UpdateCryptoKey(ctx, req); err != nil {
			return false, fmt.Errorf("disabling rotation period on delete: %w", err)
		}
	}

	return true, nil
}

func compareCryptoKey(ctx context.Context, actual, desired *kmspb.CryptoKey) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, KMSCryptoKeySpec_FromProto, KMSCryptoKeySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *kmsCryptoKeyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *kmspb.CryptoKey) error {
	mapCtx := &direct.MapContext{}
	status := KMSCryptoKeyStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.SelfLink = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
