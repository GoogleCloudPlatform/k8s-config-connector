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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vectorsearch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/vectorsearch/apiv1"
	pb "cloud.google.com/go/vectorsearch/apiv1/vectorsearchpb"
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

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelVectorSearchCollection{config: *config}, nil
}

var _ directbase.Model = &modelVectorSearchCollection{}

type modelVectorSearchCollection struct {
	config config.ControllerConfig
}

func (m *modelVectorSearchCollection) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building VectorSearch client: %w", err)
	}
	return gcpClient, err
}

func (m *modelVectorSearchCollection) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VectorSearchCollection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := VectorSearchCollectionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &VectorSearchCollectionAdapter{
		id:        id.(*krm.VectorSearchCollectionIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelVectorSearchCollection) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type VectorSearchCollectionAdapter struct {
	id        *krm.VectorSearchCollectionIdentity
	gcpClient *gcp.Client
	desired   *pb.Collection
	actual    *pb.Collection
}

var _ directbase.Adapter = &VectorSearchCollectionAdapter{}

func (a *VectorSearchCollectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VectorSearchCollection", "name", a.id)

	req := &pb.GetCollectionRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCollection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VectorSearchCollection %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *VectorSearchCollectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating VectorSearchCollection", "id", fqn)

	parent := a.id.ParentString()

	a.desired.Name = a.id.String()

	req := &pb.CreateCollectionRequest{
		Parent:       parent,
		CollectionId: a.id.Collection,
		Collection:   a.desired,
	}
	op, err := a.gcpClient.CreateCollection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VectorSearchCollection %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for VectorSearchCollection %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created VectorSearchCollection", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *VectorSearchCollectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VectorSearchCollection", "name", a.id.String())

	diffs, updateMask, err := compareCollection(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		a.desired.Name = a.id.String()

		req := &pb.UpdateCollectionRequest{
			Collection: a.desired,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateCollection(ctx, req)
		if err != nil {
			return fmt.Errorf("updating VectorSearchCollection %s: %w", a.id.String(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for VectorSearchCollection %s update: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *VectorSearchCollectionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Collection) error {
	mapCtx := &direct.MapContext{}
	status := VectorSearchCollectionStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.LazyPtr(a.id.String())

	return op.UpdateStatus(ctx, status, nil)
}

func (a *VectorSearchCollectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VectorSearchCollectionGVK)

	u.Object = uObj
	return u, nil
}

func (a *VectorSearchCollectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VectorSearchCollection", "name", a.id.String())

	req := &pb.DeleteCollectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCollection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting VectorSearchCollection %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for VectorSearchCollection %s deletion: %w", a.id.String(), err)
	}
	return true, nil
}

func compareCollection(ctx context.Context, actual, desired *pb.Collection) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	clonedActual := proto.Clone(actual).(*pb.Collection)
	clonedActual.Name = desired.Name
	clonedActual.CreateTime = desired.CreateTime
	clonedActual.UpdateTime = desired.UpdateTime

	return tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), clonedActual.ProtoReflect())
}

func VectorSearchCollectionStatus_FromProto(mapCtx *direct.MapContext, in *pb.Collection) *krm.VectorSearchCollectionStatus {
	if in == nil {
		return nil
	}
	out := &krm.VectorSearchCollectionStatus{}
	out.ObservedState = VectorSearchCollectionObservedState_FromProto(mapCtx, in)
	return out
}
