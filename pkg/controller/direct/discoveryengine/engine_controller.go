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
// proto.service: google.cloud.discoveryengine.v1.EngineService
// proto.message: google.cloud.discoveryengine.v1.Engine
// crd.type: DiscoveryEngineEngine
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineEngineGVK, NewEngineModel)
}

func NewEngineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &engineModel{config: *config}, nil
}

var _ directbase.Model = &engineModel{}

type engineModel struct {
	config config.ControllerConfig
}

func (m *engineModel) client(ctx context.Context, projectID string) (*gcp.EngineClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewEngineRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine engine client: %w", err)
	}

	return gcpClient, err
}

func (m *engineModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineEngine{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewDiscoveryEngineEngineIDFromObject(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineEngineSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.ProjectID)
	if err != nil {
		return nil, err
	}

	return &engineAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *engineModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Not implemented
	return nil, nil
}

type engineAdapter struct {
	gcpClient *gcp.EngineClient
	id        *krm.DiscoveryEngineEngineID
	desired   *pb.Engine
	actual    *pb.Engine
}

var _ directbase.Adapter = &engineAdapter{}

func (a *engineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine engine", "name", a.id)

	req := &pb.GetEngineRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine engine %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *engineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine engine", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	req := &pb.CreateEngineRequest{
		Parent:   a.id.CollectionLink.String(),
		Engine:   desired,
		EngineId: a.id.Engine,
	}
	op, err := a.gcpClient.CreateEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine engine %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("discoveryengine engine %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created discoveryengine engine in gcp", "name", a.id)

	status := &krm.DiscoveryEngineEngineStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineEngineObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *engineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine engine", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
		report.AddField("display_name", a.actual.DisplayName, a.desired.DisplayName)
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(a.desired.DataStoreIds, a.actual.DataStoreIds) {
		report.AddField("data_store_ids", a.actual.DataStoreIds, a.desired.DataStoreIds)
		updateMask.Paths = append(updateMask.Paths, "data_store_ids")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateEngineRequest{
		UpdateMask: updateMask,
		Engine:     desired,
	}
	updated, err := a.gcpClient.UpdateEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("updating discoveryengine engine %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated discoveryengine engine", "name", a.id)

	status := &krm.DiscoveryEngineEngineStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineEngineObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *engineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineEngine{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineEngineSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.ProjectID}
	obj.Spec.Location = a.id.Location
	obj.Spec.Collection = a.id.Collection
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Engine)
	u.SetGroupVersionKind(krm.DiscoveryEngineEngineGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *engineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine engine", "name", a.id)

	req := &pb.DeleteEngineRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEngine(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting discoveryengine engine %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine engine", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of discoveryengine engine %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
