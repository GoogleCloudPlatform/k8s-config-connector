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
// proto.message: google.cloud.kms.v1.ImportJob
// crd.type: KMSKeyRingImportJob
// crd.version: v1alpha1

package kms

import (
	"context"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.KMSKeyRingImportJobGVK, NewKeyRingImportJobModel, registry.CannotBeDeleted())
}

func NewKeyRingImportJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &keyRingImportJobModel{config: *config}, nil
}

var _ directbase.Model = &keyRingImportJobModel{}

type keyRingImportJobModel struct {
	config config.ControllerConfig
}

func (m *keyRingImportJobModel) client(ctx context.Context) (*kms.KeyManagementClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := kms.NewKeyManagementRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building kms keyringimportjob client: %w", err)
	}

	return gcpClient, nil
}

func (m *keyRingImportJobModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.KMSKeyRingImportJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idObj, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idObj.(*krm.KMSKeyRingImportJobIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idObj)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &keyRingImportJobAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *keyRingImportJobModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.KMSKeyRingImportJobIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &keyRingImportJobAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type keyRingImportJobAdapter struct {
	gcpClient *kms.KeyManagementClient
	id        *krm.KMSKeyRingImportJobIdentity
	desired   *krm.KMSKeyRingImportJob
	actual    *kmspb.ImportJob
	reader    client.Reader
}

var _ directbase.Adapter = &keyRingImportJobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *keyRingImportJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting KMSKeyRingImportJob", "name", a.id.String())

	req := &kmspb.GetImportJobRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetImportJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting kms keyringimportjob %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *keyRingImportJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating KMSKeyRingImportJob", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := KMSKeyRingImportJobSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &kmspb.CreateImportJobRequest{
		Parent:      a.id.ParentString(),
		ImportJobId: a.id.ImportJob,
		ImportJob:   desired,
	}
	_, err := a.gcpClient.CreateImportJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating kms keyringimportjob %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created KMSKeyRingImportJob in GCP", "name", a.id.String())

	// Cloud KMS CreateImportJob returns state=PENDING_GENERATION without public key/attestation.
	// Fetch the created ImportJob from GCP (matching TF's Read() after Create()) so that
	// status.publicKey, status.attestation, and state=ACTIVE are populated.
	actual, err := a.gcpClient.GetImportJob(ctx, &kmspb.GetImportJobRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("reading kms keyringimportjob %s after creation: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, actual)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *keyRingImportJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KMSKeyRingImportJob", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := KMSKeyRingImportJobSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, err := compareKeyRingImportJob(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("KMSKeyRingImportJob is immutable and cannot be updated. Detected changes: %v", diffs.Fields)
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

// Export implements the Adapter interface.
func (a *keyRingImportJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("exporting KMSKeyRingImportJob", "name", a.id.String())
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.KMSKeyRingImportJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(KMSKeyRingImportJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.KeyRing = a.id.ParentString()
	obj.Spec.ImportJobId = a.id.ImportJob

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ImportJob)
	u.SetGroupVersionKind(krm.KMSKeyRingImportJobGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *keyRingImportJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("no-op deleting KMSKeyRingImportJob", "name", a.id.String())
	// KMS API does not support deleting an import job. An import job expires after three days.
	// Return success to remove the finalizer, so the resource can be deleted in k8s.
	return true, nil
}

func (a *keyRingImportJobAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *kmspb.ImportJob) error {
	mapCtx := &direct.MapContext{}
	status := KMSKeyRingImportJobStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareKeyRingImportJob(ctx context.Context, actual, desired *kmspb.ImportJob) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, KMSKeyRingImportJobSpec_FromProto, KMSKeyRingImportJobSpec_ToProto)
	if err != nil {
		return nil, err
	}
	diffs, _, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
