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

package mapmanagement

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/maps/mapmanagement/apiv2beta"
	pb "cloud.google.com/go/maps/mapmanagement/apiv2beta/mapmanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/mapmanagement/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
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
	registry.RegisterModel(krm.MapManagementMapConfigGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
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
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building MapManagement REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MapManagementMapConfig{}
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
	id, ok := idVal.(*krm.MapManagementMapConfigIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := MapManagementMapConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &MapManagementMapConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.MapManagementMapConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &MapManagementMapConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type MapManagementMapConfigAdapter struct {
	id        *krm.MapManagementMapConfigIdentity
	gcpClient *gcp.Client
	desired   *pb.MapConfig
	actual    *pb.MapConfig
}

var _ directbase.Adapter = &MapManagementMapConfigAdapter{}

func (a *MapManagementMapConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting MapManagementMapConfig", "name", fqn)

	req := &pb.GetMapConfigRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetMapConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MapManagementMapConfig %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *MapManagementMapConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating MapManagementMapConfig", "name", fqn)

	req := &pb.CreateMapConfigRequest{
		Parent:    parent,
		MapConfig: a.desired,
	}
	created, err := a.gcpClient.CreateMapConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MapManagementMapConfig %s: %w", a.id.MapConfig, err)
	}
	log.V(2).Info("successfully created MapManagementMapConfig", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *MapManagementMapConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating MapManagementMapConfig", "name", fqn)

	diffs, updateMask, err := compareMapConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.MapConfig)
		desired.Name = fqn

		req := &pb.UpdateMapConfigRequest{
			MapConfig:  desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateMapConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating MapManagementMapConfig %s: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *MapManagementMapConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MapConfig) error {
	mapCtx := &direct.MapContext{}
	status := MapManagementMapConfigStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *MapManagementMapConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MapManagementMapConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MapManagementMapConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.MapConfig)
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{
		External: a.id.Project,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.MapConfig)
	u.SetGroupVersionKind(krm.MapManagementMapConfigGVK)

	return u, nil
}

func (a *MapManagementMapConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting MapManagementMapConfig", "name", fqn)

	req := &pb.DeleteMapConfigRequest{
		Name:  fqn,
		Force: true,
	}
	err := a.gcpClient.DeleteMapConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent MapManagementMapConfig, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting MapManagementMapConfig %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted MapManagementMapConfig", "name", fqn)
	return true, nil
}

func compareMapConfig(ctx context.Context, actual, desired *pb.MapConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, MapManagementMapConfigSpec_FromProto, MapManagementMapConfigSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func MapManagementMapConfigStatus_FromProto(mapCtx *direct.MapContext, latest *pb.MapConfig) *krm.MapManagementMapConfigStatus {
	if latest == nil {
		return nil
	}
	status := &krm.MapManagementMapConfigStatus{}
	status.ObservedState = MapManagementMapConfigObservedState_FromProto(mapCtx, latest)
	return status
}
