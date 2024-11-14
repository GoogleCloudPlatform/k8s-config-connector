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

package workstations

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/workstations/apiv1"
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.WorkstationConfigGVK, NewWorkstationConfigModel)
}

func NewWorkstationConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkstationConfig{config: *config}, nil
}

var _ directbase.Model = &modelWorkstationConfig{}

type modelWorkstationConfig struct {
	config config.ControllerConfig
}

func (m *modelWorkstationConfig) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building WorkstationConfig client: %w", err)
	}
	return gcpClient, err
}

func (m *modelWorkstationConfig) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.WorkstationConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewWorkstationConfigRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// TODO: Resolve Refs

	// Get workstations GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WorkstationConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWorkstationConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type WorkstationConfigAdapter struct {
	id        *krm.WorkstationConfigRef
	gcpClient *gcp.Client
	desired   *krm.WorkstationConfig
	actual    *pb.WorkstationConfig
}

var _ directbase.Adapter = &WorkstationConfigAdapter{}

func (a *WorkstationConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting WorkstationConfig", "name", a.id.External)

	req := &pb.GetWorkstationConfigRequest{Name: a.id.External}
	workstationconfigpb, err := a.gcpClient.GetWorkstationConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WorkstationConfig %q: %w", a.id.External, err)
	}

	a.actual = workstationconfigpb
	return true, nil
}

func (a *WorkstationConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating WorkstationConfig", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkstationConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.id.Parent()
	if err != nil {
		return err
	}
	req := &pb.CreateWorkstationConfigRequest{
		Parent:              parent.String(),
		WorkstationConfigId: a.id.Name,
		WorkstationConfig:   resource,
	}
	op, err := a.gcpClient.CreateWorkstationConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WorkstationConfig %s: %w", a.id.External, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkstationConfig %s waiting creation: %w", a.id.External, err)
	}
	log.V(2).Info("successfully created WorkstationConfig", "name", a.id.External)

	status := &krm.WorkstationConfigStatus{}
	status.ObservedState = WorkstationConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *WorkstationConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating WorkstationConfig", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkstationConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set name and etag manually, because they are not filled-in by WorkstationConfigSpec_ToProto.
	resource.Name = a.id.External
	resource.Etag = a.actual.Etag

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		return nil
	}
	req := &pb.UpdateWorkstationConfigRequest{
		WorkstationConfig: resource,
		UpdateMask:        &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}

	op, err := a.gcpClient.UpdateWorkstationConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating WorkstationConfig %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkstationConfig %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated WorkstationConfig", "name", a.id.External)

	status := &krm.WorkstationConfigStatus{}
	status.ObservedState = WorkstationConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *WorkstationConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	/*
		if a.actual == nil {
			return nil, fmt.Errorf("Find() not called")
		}
		u := &unstructured.Unstructured{}

		obj := &krm.WorkstationConfig{}
		mapCtx := &direct.MapContext{}
		obj.Spec = direct.ValueOf(WorkstationConfigSpec_FromProto(mapCtx, a.actual))
		if mapCtx.Err() != nil {
			return nil, mapCtx.Err()
		}
		// TODO(user): Update other resource references
		parent, err := a.id.Parent()
		if err != nil {
			return nil, err
		}
		obj.Spec.ProjectRef = &refs.ProjectRef{External: parent.String()}
		obj.Spec.Location = parent.Location
		uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return nil, err
		}

		u.SetName(a.actual.Id)
		u.SetGroupVersionKind(krm.WorkstationConfigGVK)

		u.Object = uObj
		return u, nil
	*/
	return nil, nil
}

// Delete implements the Adapter interface.
func (a *WorkstationConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting WorkstationConfig", "name", a.id.External)

	req := &pb.DeleteWorkstationConfigRequest{Name: a.id.External}
	op, err := a.gcpClient.DeleteWorkstationConfig(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting WorkstationConfig %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully deleted WorkstationConfig", "name", a.id.External)

	_, err = op.Wait(ctx)
	if err != nil {
		// todo (b/368419476): Workstation service does not provide a valid response on success.
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting delete WorkstationConfig %s: %w", a.id.External, err)
		}
	}

	return true, nil
}
