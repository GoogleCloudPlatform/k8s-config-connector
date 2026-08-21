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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.MetadataJob
// crd.type: DataplexMetadataJob
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexMetadataJobGVK, NewMetadataJobModel, registry.CannotBeDeleted())
}

func NewMetadataJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &metadataJobModel{config: config}, nil
}

var _ directbase.Model = &metadataJobModel{}

type metadataJobModel struct {
	config *config.ControllerConfig
}

func (m *metadataJobModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexMetadataJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	copied := obj.DeepCopy()
	mapCtx := &direct.MapContext{}
	desired := DataplexMetadataJobSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	idI, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idI.(*krm.MetadataJobIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idI)
	}

	adapter := &metadataJobAdapter{
		id:      id,
		desired: desired,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	client, err := gcpClient.catalogClient(ctx)
	if err != nil {
		return nil, err
	}
	adapter.gcpClient = client

	return adapter, nil
}

func (m *metadataJobModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type metadataJobAdapter struct {
	gcpClient *gcp.CatalogClient
	id        *krm.MetadataJobIdentity
	desired   *pb.MetadataJob
	actual    *pb.MetadataJob
	reader    client.Reader
}

var _ directbase.Adapter = &metadataJobAdapter{}

func (a *metadataJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex metadatajob", "name", a.id)

	req := &pb.GetMetadataJobRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetMetadataJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex metadatajob %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *metadataJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex metadatajob", "name", a.id)

	req := &pb.CreateMetadataJobRequest{
		Parent:        a.id.ParentString(),
		MetadataJob:   a.desired,
		MetadataJobId: a.id.MetadataJob,
	}
	op, err := a.gcpClient.CreateMetadataJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex metadatajob %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex metadatajob %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex metadatajob in gcp", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.DataplexMetadataJobStatus{}
	status.ObservedState = DataplexMetadataJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *metadataJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex metadatajob", "name", a.id)

	paths, err := common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing spec: %w", err)
	}
	if len(paths) > 0 {
		return fmt.Errorf("DataplexMetadataJob is immutable and cannot be updated. Detected changes: %v", paths.UnsortedList())
	}

	// Even though there is no update, we still want to update KRM status with the latest state from Find().
	mapCtx := &direct.MapContext{}
	status := &krm.DataplexMetadataJobStatus{}
	status.ObservedState = DataplexMetadataJobObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *metadataJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexMetadataJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexMetadataJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.MetadataJob)
	u.SetGroupVersionKind(krm.DataplexMetadataJobGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *metadataJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex metadatajob (no-op)", "name", a.id)
	// MetadataJob does not support deletion. Returning true allows KCC to remove the finalizer.
	return true, nil
}
