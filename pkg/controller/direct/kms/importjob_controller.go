// Copyright 2024 Google LLC
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
// crd.type: KMSImportJob
// crd.version: v1alpha1

package kms

import (
	"context"
	"fmt"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.KMSImportJobGVK, NewImportJobModel)
}

func NewImportJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &importJobModel{config: *config}, nil
}

var _ directbase.Model = &importJobModel{}

type importJobModel struct {
	config config.ControllerConfig
}

func (m *importJobModel) client(ctx context.Context, projectID string) (*kms.KeyManagementClient, error) {
	var opts []option.ClientOption

	config := m.config

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := kms.NewKeyManagementRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building kms importjob client: %w", err)
	}

	return gcpClient, err
}

func (m *importJobModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.KMSImportJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewImportJobIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &importJobAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *importJobModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type importJobAdapter struct {
	gcpClient *kms.KeyManagementClient
	id        *krm.ImportJobIdentity
	desired   *krm.KMSImportJob
	actual    *kmspb.ImportJob
	reader    client.Reader
}

var _ directbase.Adapter = &importJobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *importJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting kms importjob", "name", a.id)

	req := &kmspb.GetImportJobRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetImportJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting kms importjob %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *importJobAdapter) normalizeReferences(ctx context.Context) error {
	if a.desired.Spec.KMSKeyRingRef != nil {
		if _, err := refs.ResolveKMSKeyRingRef(ctx, a.reader, a.desired, a.desired.Spec.KMSKeyRingRef); err != nil {
			return err
		}
	}
	return nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *importJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating kms importjob", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := KMSImportJobSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &kmspb.CreateImportJobRequest{
		Parent:      a.id.Parent().String(),
		ImportJobId: a.id.ID(),
		ImportJob:   desired,
	}
	created, err := a.gcpClient.CreateImportJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating kms importjob %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created kms importjob in gcp", "name", a.id)

	status := &krm.KMSImportJobStatus{}
	status.ObservedState = KMSImportJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *importJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating KMSImportJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desiredPb := KMSImportJobSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	// Match output-only fields that are not marked as output-only in the proto.
	desiredPb.PublicKey = a.actual.PublicKey
	desiredPb.Attestation = a.actual.Attestation

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	log.V(2).Info("KMSImportJob doesn't support update", "name", a.id)
	return fmt.Errorf("updating KMSImportJob %s: update is not supported", a.id)
}

// Export implements the Adapter interface.
func (a *importJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.KMSImportJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(KMSImportJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.KMSKeyRingRef = &refs.KMSKeyRingRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.KMSImportJobGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *importJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("No-op Delete for import job", "name", a.id.String())
	// KMS API does not support deleting an import job. An import job expires after three days.
	// Return success to remove the finalizer, so the resource can be deleted in k8s.
	return true, nil
}
