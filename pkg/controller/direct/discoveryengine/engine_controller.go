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
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineEngineGVK, NewDiscoveryEngineEngineModel)
}

func NewDiscoveryEngineEngineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDiscoveryEngineEngine{config: *config}, nil
}

var _ directbase.Model = &modelDiscoveryEngineEngine{}

type modelDiscoveryEngineEngine struct {
	config config.ControllerConfig
}

func (m *modelDiscoveryEngineEngine) client(ctx context.Context, projectID string) (*gcp.EngineClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	config.UserProjectOverride = true
	config.BillingProject = projectID

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewEngineRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building engine client: %w", err)
	}

	return gcpClient, err
}

func (m *modelDiscoveryEngineEngine) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DiscoveryEngineEngine{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewDiscoveryEngineEngineIDFromObject(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	var projectRef *refs.Project
	if err := common.NormalizeReferences(ctx, reader, obj, projectRef); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineEngineSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &DiscoveryEngineEngineAdapter{
		model:   m,
		id:      id,
		desired: desired,
	}, nil
}

func (m *modelDiscoveryEngineEngine) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DiscoveryEngineEngineAdapter struct {
	model   *modelDiscoveryEngineEngine
	id      *krm.DiscoveryEngineEngineID
	desired *pb.Engine
	actual  *pb.Engine
}

var _ directbase.Adapter = &DiscoveryEngineEngineAdapter{}

func (a *DiscoveryEngineEngineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine engine", "name", a.id)

	gcpClient, err := a.model.client(ctx, a.id.ProjectID)
	if err != nil {
		return false, err
	}
	req := &pb.GetEngineRequest{Name: a.id.String()}
	enginepb, err := gcpClient.GetEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine engine %q from gcp: %w", a.id.String(), err)
	}

	a.actual = enginepb
	return true, nil
}

func (a *DiscoveryEngineEngineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Engine", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	gcpClient, err := a.model.client(ctx, a.id.ProjectID)
	if err != nil {
		return err
	}

	req := &pb.CreateEngineRequest{
		Parent:   a.id.CollectionLink.String(),
		Engine:   desired,
		EngineId: a.id.Engine,
	}
	op, err := gcpClient.CreateEngine(ctx, req)
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

func (a *DiscoveryEngineEngineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Engine", "name", a.id)

	gcpClient, err := a.model.client(ctx, a.id.ProjectID)
	if err != nil {
		return err
	}

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	req := &pb.UpdateEngineRequest{
		UpdateMask: updateMask,
		Engine:     desired,
	}
	updated, err := gcpClient.UpdateEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Engine %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Engine", "name", a.id)

	status := &krm.DiscoveryEngineEngineStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineEngineObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *DiscoveryEngineEngineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

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

	u.SetName(a.id.Engine)
	u.SetGroupVersionKind(krm.DiscoveryEngineEngineGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *DiscoveryEngineEngineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Engine", "name", a.id)

	gcpClient, err := a.model.client(ctx, a.id.ProjectID)
	if err != nil {
		return false, err
	}

	req := &pb.DeleteEngineRequest{Name: a.id.String()}
	op, err := gcpClient.DeleteEngine(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting discoveryengine engine %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine engine", "name", a.id)

	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return false, fmt.Errorf("waiting for deletion of discoveryengine engine %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}
