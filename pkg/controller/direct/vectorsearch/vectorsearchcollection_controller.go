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

package vectorsearch

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/vectorsearch/apiv1"
	pb "cloud.google.com/go/vectorsearch/apiv1/vectorsearchpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vectorsearch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VectorSearchCollectionGVK, NewModel)
}

func NewModel(_ context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	opts = append(opts, option.WithEndpoint("https://vectorsearch.googleapis.com"))
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building vectorsearch client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VectorSearchCollection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.VectorSearchCollectionIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := VectorSearchCollectionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = id.String()

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &vectorSearchCollectionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(_ context.Context, _ string) (directbase.Adapter, error) {
	return nil, nil
}

type vectorSearchCollectionAdapter struct {
	id        *krm.VectorSearchCollectionIdentity
	gcpClient *gcp.Client
	desired   *pb.Collection
	actual    *pb.Collection
}

var _ directbase.Adapter = &vectorSearchCollectionAdapter{}

func (a *vectorSearchCollectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting VectorSearchCollection", "name", fqn)

	req := &pb.GetCollectionRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetCollection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VectorSearchCollection %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *vectorSearchCollectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating VectorSearchCollection", "name", fqn)

	req := &pb.CreateCollectionRequest{
		Parent:       parent,
		CollectionId: a.id.Collection,
		Collection:   a.desired,
	}
	op, err := a.gcpClient.CreateCollection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VectorSearchCollection %s: %w", a.id.Collection, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for VectorSearchCollection %s creation: %w", a.id.Collection, err)
	}
	log.V(2).Info("successfully created VectorSearchCollection", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *vectorSearchCollectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating VectorSearchCollection", "name", fqn)

	diffs, updateMask, err := compareCollection(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = fqn

		req := &pb.UpdateCollectionRequest{
			Collection: desired,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateCollection(ctx, req)
		if err != nil {
			return fmt.Errorf("updating VectorSearchCollection %s: %w", fqn, err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for VectorSearchCollection %s update: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *vectorSearchCollectionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Collection) error {
	mapCtx := &direct.MapContext{}
	status := vectorSearchCollectionStatusFromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *vectorSearchCollectionAdapter) Export(_ context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VectorSearchCollection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VectorSearchCollectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Collection

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Collection)
	u.SetGroupVersionKind(krm.VectorSearchCollectionGVK)
	return u, nil
}

func (a *vectorSearchCollectionAdapter) Delete(ctx context.Context, _ *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting VectorSearchCollection", "name", fqn)

	req := &pb.DeleteCollectionRequest{
		Name:  fqn,
		Force: true,
	}
	op, err := a.gcpClient.DeleteCollection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent VectorSearchCollection, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting VectorSearchCollection %s: %w", fqn, err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for VectorSearchCollection %s deletion: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted VectorSearchCollection", "name", fqn)
	return true, nil
}

func compareCollection(ctx context.Context, actual, desired *pb.Collection) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VectorSearchCollectionSpec_FromProto, VectorSearchCollectionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	// Restore Name if needed
	maskedActual.Name = desired.Name

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func vectorSearchCollectionStatusFromProto(mapCtx *direct.MapContext, in *pb.Collection) *krm.VectorSearchCollectionStatus {
	if in == nil {
		return nil
	}
	out := &krm.VectorSearchCollectionStatus{}
	if in.Name != "" {
		out.ExternalRef = &in.Name
	}
	out.ObservedState = VectorSearchCollectionObservedState_FromProto(mapCtx, in)
	return out
}
