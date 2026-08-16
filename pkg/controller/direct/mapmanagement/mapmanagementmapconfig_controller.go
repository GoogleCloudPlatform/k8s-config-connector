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
// proto.service: google.maps.mapmanagement.v2beta.MapManagementService
// proto.message: google.maps.mapmanagement.v2beta.MapConfig
// crd.type: MapManagementMapConfig
// crd.version: v1alpha1

package mapmanagement

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/maps/mapmanagement/apiv2beta"
	pb "cloud.google.com/go/maps/mapmanagement/apiv2beta/mapmanagementpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/mapmanagement/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.MapManagementMapConfigGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context, projectID string) (*gcp.Client, error) {
	var opts []option.ClientOption

	cfg := m.config
	if !cfg.UserProjectOverride || cfg.BillingProject == "" {
		cfg.UserProjectOverride = true
		cfg.BillingProject = projectID
	}

	opts, err := cfg.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building maps mapmanagement client: %w", err)
	}

	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MapManagementMapConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.MapManagementMapConfigIdentity)

	mapCtx := &direct.MapContext{}
	desired := MapManagementMapConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//mapmanagement.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//mapmanagement.googleapis.com/")
		id := &krm.MapManagementMapConfigIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			log.V(2).Error(err, "url did not match MapManagementMapConfig format", "url", url)
			return nil, nil
		}
		gcpClient, err := m.client(ctx, id.Project)
		if err != nil {
			return nil, err
		}
		return &adapter{
			gcpClient: gcpClient,
			id:        id,
		}, nil
	}
	return nil, nil
}

type adapter struct {
	gcpClient *gcp.Client
	id        *krm.MapManagementMapConfigIdentity
	desired   *pb.MapConfig
	actual    *pb.MapConfig
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting mapmanagement map config", "name", a.id)

	req := &pb.GetMapConfigRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetMapConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting mapmanagement map config %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating mapmanagement map config", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = "" // Name is server-assigned and read-only, must be unset for create payload

	req := &pb.CreateMapConfigRequest{
		Parent:    a.id.ParentString(),
		MapConfig: desired,
	}
	created, err := a.gcpClient.CreateMapConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating mapconfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created mapconfig in gcp", "name", created.GetName())

	// Fetch latest after creation as per guideline (Rule: Always perform a GET operation after a Create LRO/Operation)
	getReq := &pb.GetMapConfigRequest{Name: created.GetName()}
	latest, err := a.gcpClient.GetMapConfig(ctx, getReq)
	if err != nil {
		latest = created
	}

	// Update the adapter id to use the server-assigned MapConfig ID so that subsequent updates use the correct ID.
	if err := a.id.FromExternal(created.GetName()); err != nil {
		return fmt.Errorf("parsing generated mapconfig name: %w", err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating mapmanagement map config", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	diffs, updateMask, err := a.compare(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateMapConfigRequest{
		MapConfig:  desired,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateMapConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating mapconfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated mapconfig", "name", a.id)

	// Fetch latest after update as per guideline
	getReq := &pb.GetMapConfigRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetMapConfig(ctx, getReq)
	if err != nil {
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *adapter) compare(ctx context.Context, actual, desired *pb.MapConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, MapManagementMapConfigSpec_FromProto, MapManagementMapConfigSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MapConfig) error {
	status := &krm.MapManagementMapConfigStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = MapManagementMapConfigObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting mapmanagement map config", "name", a.id)

	req := &pb.DeleteMapConfigRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteMapConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting mapconfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted mapconfig", "name", a.id)
	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.MapManagementMapConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MapManagementMapConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	if a.id.MapConfig != "" {
		obj.Spec.ResourceID = direct.LazyPtr(a.id.MapConfig)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}
