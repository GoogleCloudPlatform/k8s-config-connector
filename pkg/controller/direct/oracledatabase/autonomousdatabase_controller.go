// Copyright 2025 Google LLC
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

package oracledatabase

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/oracledatabase/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	oracledatabasepb "cloud.google.com/go/oracledatabase/v1/oracledatabasepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.OracledatabaseAutonomousDatabaseGVK, NewAutonomousDatabaseModel)
}

func NewAutonomousDatabaseModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAutonomousDatabase{config: *config}, nil
}

var _ directbase.Model = &modelAutonomousDatabase{}

type modelAutonomousDatabase struct {
	config config.ControllerConfig
}

func (m *modelAutonomousDatabase) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AutonomousDatabase client: %w", err)
	}
	return gcpClient, err
}

func (m *modelAutonomousDatabase) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.OracledatabaseAutonomousDatabase{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAutonomousDatabaseIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get oracledatabase GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AutonomousDatabaseAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAutonomousDatabase) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AutonomousDatabaseAdapter struct {
	id        *krm.AutonomousDatabaseIdentity
	gcpClient *gcp.Client
	desired   *krm.OracledatabaseAutonomousDatabase
	actual    *oracledatabasepb.AutonomousDatabase
}

var _ directbase.Adapter = &AutonomousDatabaseAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AutonomousDatabaseAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AutonomousDatabase", "name", a.id)

	req := &oracledatabasepb.GetAutonomousDatabaseRequest{Name: a.id}
	autonomousdatabasepb, err := a.gcpClient.GetAutonomousDatabase(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AutonomousDatabase %q: %w", a.id, err)
	}

	a.actual = autonomousdatabasepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AutonomousDatabaseAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AutonomousDatabase", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := OracledatabaseAutonomousDatabaseSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &oracledatabasepb.CreateAutonomousDatabaseRequest{
		Parent:             a.id.Parent().String(),
		AutonomousDatabase: resource,
	}
	op, err := a.gcpClient.CreateAutonomousDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AutonomousDatabase %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AutonomousDatabase %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created AutonomousDatabase", "name", a.id)

	status := &krm.OracledatabaseAutonomousDatabaseStatus{}
	status.ObservedState = OracledatabaseAutonomousDatabaseObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AutonomousDatabaseAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AutonomousDatabase", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := OracledatabaseAutonomousDatabaseSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	// Option 1: This option is good for proto that has `field_mask` for output-only, immutable, required/optional.
	// TODO(contributor): If choosing this option, remove the "Option 2" code.
	{
		var err error
		paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
		if err != nil {
			return err
		}
	}

	// Option 2: manually add all mutable fields.
	// TODO(contributor): If choosing this option, remove the "Option 1" code.
	{
		if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
			paths = append(paths, "display_name")
		}
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		status := &krm.OracledatabaseAutonomousDatabaseStatus{}
		status.ObservedState = OracledatabaseAutonomousDatabaseObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &oracledatabasepb.UpdateAutonomousDatabaseRequest{
		Name:               a.id.External,
		UpdateMask:         updateMask,
		AutonomousDatabase: desiredPb,
	}
	op, err := a.gcpClient.UpdateAutonomousDatabase(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AutonomousDatabase %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AutonomousDatabase %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated AutonomousDatabase", "name", a.id.External)

	status := &krm.OracledatabaseAutonomousDatabaseStatus{}
	status.ObservedState = OracledatabaseAutonomousDatabaseObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AutonomousDatabaseAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.OracledatabaseAutonomousDatabase{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(OracledatabaseAutonomousDatabaseSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Id)
	u.SetGroupVersionKind(krm.OracledatabaseAutonomousDatabaseGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AutonomousDatabaseAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AutonomousDatabase", "name", a.id)

	req := &oracledatabasepb.DeleteAutonomousDatabaseRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAutonomousDatabase(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting AutonomousDatabase %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AutonomousDatabase", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete AutonomousDatabase %s: %w", a.id, err)
	}
	return true, nil
}
