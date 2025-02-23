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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"

	gcp "cloud.google.com/go/workstations/apiv1"
	"cloud.google.com/go/workstations/apiv1/workstationspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.WorkstationGVK, NewWorkstationModel)
	fuzztesting.RegisterKRMFuzzer(workstationFuzzer())
}

func workstationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&workstationspb.Workstation{},
		WorkstationSpec_FromProto, WorkstationSpec_ToProto,
		WorkstationObservedState_FromProto, WorkstationObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".reconciling")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".start_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".host")

	return f
}

func NewWorkstationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkstation{config: *config}, nil
}

var _ directbase.Model = &modelWorkstation{}

type modelWorkstation struct {
	config config.ControllerConfig
}

func (m *modelWorkstation) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Workstation client: %w", err)
	}
	return gcpClient, err
}

func (m *modelWorkstation) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.Workstation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewWorkstationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get workstations GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WorkstationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWorkstation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type WorkstationAdapter struct {
	id        *krm.WorkstationIdentity
	gcpClient *gcp.Client
	desired   *krm.Workstation
	actual    *workstationspb.Workstation
}

var _ directbase.Adapter = &WorkstationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *WorkstationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Workstation", "name", a.id)

	req := &workstationspb.GetWorkstationRequest{Name: a.id.String()}
	workstationpb, err := a.gcpClient.GetWorkstation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Workstation %q: %w", a.id, err)
	}

	a.actual = workstationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WorkstationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Workstation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkstationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set name manually, because it is not filled-in by WorkstationSpec_ToProto.
	resource.Name = a.id.String()

	req := &workstationspb.CreateWorkstationRequest{
		Parent:        a.id.Parent().String(),
		WorkstationId: a.id.ID(),
		Workstation:   resource,
	}
	op, err := a.gcpClient.CreateWorkstation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Workstation %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting Workstation %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Workstation", "name", a.id)

	status := &krm.WorkstationStatus{}
	status.ObservedState = WorkstationObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WorkstationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Workstation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := WorkstationSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set name and etag manually, because they are not filled-in by WorkstationConfigSpec_ToProto.
	desiredPb.Name = a.id.String()
	desiredPb.Etag = a.actual.Etag

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.WorkstationStatus{}
		status.ObservedState = WorkstationObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{Paths: sets.List(paths)}

	req := &workstationspb.UpdateWorkstationRequest{
		Workstation: desiredPb,
		UpdateMask:  updateMask,
	}
	op, err := a.gcpClient.UpdateWorkstation(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Workstation %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting Workstation %s update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Workstation", "name", a.id.String())

	status := &krm.WorkstationStatus{}
	status.ObservedState = WorkstationObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *WorkstationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.Workstation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(WorkstationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Parent = &krm.WorkstationConfigRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.WorkstationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *WorkstationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Workstation", "name", a.id)

	req := &workstationspb.DeleteWorkstationRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteWorkstation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Workstation, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Workstation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Workstation", "name", a.id)

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Workstation %s: %w", a.id, err)
	}
	return true, nil
}
