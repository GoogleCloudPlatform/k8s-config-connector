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

package contentwarehouse

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/contentwarehouse/v1"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ContentWarehouseSynonymSetGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ContentWarehouse service: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ContentWarehouseSynonymSet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.ContentWarehouseSynonymSetIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	mapCtx := &direct.MapContext{}
	desired := ContentWarehouseSynonymSetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ContentWarehouseSynonymSetAdapter{
		id:            id,
		gcpClient:     gcpClient,
		desired:       desired,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ContentWarehouseSynonymSetAdapter struct {
	id            *krm.ContentWarehouseSynonymSetIdentity
	gcpClient     *api.Service
	desired       *pb.SynonymSet
	actual        *pb.SynonymSet
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &ContentWarehouseSynonymSetAdapter{}

func (a *ContentWarehouseSynonymSetAdapter) fullyQualifiedName(ctx context.Context) (string, error) {
	projectNumber, err := a.projectMapper.LookupProjectNumber(ctx, a.id.Project)
	if err != nil {
		return "", fmt.Errorf("error converting project ID %s to project number: %w", a.id.Project, err)
	}
	return fmt.Sprintf("projects/%d/locations/%s/synonymSets/%s", projectNumber, a.id.Location, a.id.Context), nil
}

func (a *ContentWarehouseSynonymSetAdapter) parentName(ctx context.Context) (string, error) {
	projectNumber, err := a.projectMapper.LookupProjectNumber(ctx, a.id.Project)
	if err != nil {
		return "", fmt.Errorf("error converting project ID %s to project number: %w", a.id.Project, err)
	}
	return fmt.Sprintf("projects/%d/locations/%s", projectNumber, a.id.Location), nil
}

func (a *ContentWarehouseSynonymSetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ContentWarehouseSynonymSet", "name", a.id)

	name, err := a.fullyQualifiedName(ctx)
	if err != nil {
		return false, err
	}

	restObj, err := a.gcpClient.Projects.Locations.SynonymSets.Get(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ContentWarehouseSynonymSet %q: %w", a.id, err)
	}

	actual, err := RESTToProto(restObj)
	if err != nil {
		return false, err
	}

	a.actual = actual
	return true, nil
}

func (a *ContentWarehouseSynonymSetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ContentWarehouseSynonymSet", "name", a.id)

	parent, err := a.parentName(ctx)
	if err != nil {
		return err
	}

	name, err := a.fullyQualifiedName(ctx)
	if err != nil {
		return err
	}

	a.desired.Name = name
	a.desired.Context = a.id.Context

	restObj, err := ProtoToREST(a.desired)
	if err != nil {
		return err
	}

	createdREST, err := a.gcpClient.Projects.Locations.SynonymSets.Create(parent, restObj).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ContentWarehouseSynonymSet %q: %w", a.id, err)
	}

	created, err := RESTToProto(createdREST)
	if err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ContentWarehouseSynonymSetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ContentWarehouseSynonymSet", "name", a.id)

	name, err := a.fullyQualifiedName(ctx)
	if err != nil {
		return err
	}
	a.desired.Name = name
	a.desired.Context = a.id.Context

	diffResults, err := compareSynonymSet(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffResults.HasDiff() {
		diffResults.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffResults)

		restObj, err := ProtoToREST(a.desired)
		if err != nil {
			return err
		}

		updatedREST, err := a.gcpClient.Projects.Locations.SynonymSets.Patch(name, restObj).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating ContentWarehouseSynonymSet %q: %w", a.id, err)
		}

		updated, err := RESTToProto(updatedREST)
		if err != nil {
			return err
		}
		latest = updated
	} else {
		log.V(2).Info("no diff detected for ContentWarehouseSynonymSet", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *ContentWarehouseSynonymSetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ContentWarehouseSynonymSet", "name", a.id)

	name, err := a.fullyQualifiedName(ctx)
	if err != nil {
		return false, err
	}

	_, err = a.gcpClient.Projects.Locations.SynonymSets.Delete(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ContentWarehouseSynonymSet %q: %w", a.id, err)
	}

	return true, nil
}

func (a *ContentWarehouseSynonymSetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.ContentWarehouseSynonymSet{}
	mapCtx := &direct.MapContext{}
	spec := ContentWarehouseSynonymSetSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	spec.ResourceID = direct.LazyPtr(a.id.Context)
	spec.Location = a.id.Location
	spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec = *spec
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Context)
	u.SetGroupVersionKind(krm.ContentWarehouseSynonymSetGVK)

	return u, nil
}

func compareSynonymSet(ctx context.Context, actual, desired *pb.SynonymSet) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ContentWarehouseSynonymSetSpec_FromProto, ContentWarehouseSynonymSetSpec_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name // Restore identifier field

	clonedDesired := proto.Clone(desired).(*pb.SynonymSet)

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	return diffs, err
}

func (a *ContentWarehouseSynonymSetAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SynonymSet) error {
	mapCtx := &direct.MapContext{}
	status := ContentWarehouseSynonymSetStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func ContentWarehouseSynonymSetStatus_FromProto(mapCtx *direct.MapContext, in *pb.SynonymSet) *krm.ContentWarehouseSynonymSetStatus {
	if in == nil {
		return nil
	}
	out := &krm.ContentWarehouseSynonymSetStatus{}
	out.ExternalRef = &in.Name
	out.ObservedState = ContentWarehouseSynonymSetObservedState_FromProto(mapCtx, in)
	return out
}

func ContentWarehouseSynonymSetStatus_ToProto(mapCtx *direct.MapContext, in *krm.ContentWarehouseSynonymSetStatus) *pb.SynonymSet {
	if in == nil {
		return nil
	}
	out := &pb.SynonymSet{}
	if in.ExternalRef != nil {
		out.Name = *in.ExternalRef
	}
	return out
}

func RESTToProto(in *api.GoogleCloudContentwarehouseV1SynonymSet) (*pb.SynonymSet, error) {
	if in == nil {
		return nil, nil
	}
	bytes, err := in.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling REST struct: %w", err)
	}
	out := &pb.SynonymSet{}
	unmarshalOpts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
	if err := unmarshalOpts.Unmarshal(bytes, out); err != nil {
		return nil, fmt.Errorf("error unmarshaling proto struct: %w", err)
	}
	return out, nil
}

func ProtoToREST(in *pb.SynonymSet) (*api.GoogleCloudContentwarehouseV1SynonymSet, error) {
	if in == nil {
		return nil, nil
	}
	marshalOpts := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: false,
	}
	bytes, err := marshalOpts.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("error marshaling proto struct: %w", err)
	}
	out := &api.GoogleCloudContentwarehouseV1SynonymSet{}
	if err := json.Unmarshal(bytes, out); err != nil {
		return nil, fmt.Errorf("error unmarshaling REST struct: %w", err)
	}
	return out, nil
}
