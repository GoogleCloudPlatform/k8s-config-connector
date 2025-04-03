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

// +tool:controller
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Environment
// crd.type: DataplexEnvironment
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexEnvironmentGVK, NewEnvironmentModel)
}

func NewEnvironmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &environmentModel{config: config}, nil
}

var _ directbase.Model = &environmentModel{}

type environmentModel struct {
	config *config.ControllerConfig
}

func (m *environmentModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataplexEnvironment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEnvironmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	environmentAdapter := &environmentAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	environmentClient, err := gcpClient.client(ctx)
	if err != nil {
		return nil, err
	}
	environmentAdapter.gcpClient = environmentClient

	return environmentAdapter, nil
}

func (m *environmentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type environmentAdapter struct {
	gcpClient *gcp.Client
	id        *krm.EnvironmentIdentity
	desired   *krm.DataplexEnvironment
	actual    *pb.Environment
	reader    client.Reader
}

var _ directbase.Adapter = &environmentAdapter{}

func (a *environmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex environment", "name", a.id)

	req := &pb.GetEnvironmentRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetEnvironment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex environment %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *environmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex environment", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()

	environment := DataplexEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateEnvironmentRequest{
		Parent:        a.id.Parent().String(),
		Environment:   environment,
		EnvironmentId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateEnvironment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex environment %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex environment %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex environment in gcp", "name", a.id)

	status := &krm.DataplexEnvironmentStatus{}
	status.ObservedState = DataplexEnvironmentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *environmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex environment", "name", a.id)

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	environment := DataplexEnvironmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	environment.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	if desired.Spec.DisplayName != nil && !reflect.DeepEqual(environment.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if desired.Spec.Description != nil && !reflect.DeepEqual(environment.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(environment.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if desired.Spec.InfrastructureSpec != nil && !reflect.DeepEqual(environment.InfrastructureSpec, a.actual.InfrastructureSpec) {
		updateMask.Paths = append(updateMask.Paths, "infrastructure_spec")
	}
	if desired.Spec.SessionSpec != nil && !reflect.DeepEqual(environment.SessionSpec, a.actual.SessionSpec) {
		updateMask.Paths = append(updateMask.Paths, "session_spec")
	}

	var updated *pb.Environment
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		req := &pb.UpdateEnvironmentRequest{
			UpdateMask:  updateMask,
			Environment: environment,
		}
		op, err := a.gcpClient.UpdateEnvironment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataplex environment %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of dataplex environment %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataplex environment", "name", a.id)
	}

	status := &krm.DataplexEnvironmentStatus{}
	status.ObservedState = DataplexEnvironmentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *environmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexEnvironment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexEnvironmentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Parent().Location
	obj.Spec.LakeRef = &refs.DataplexLakeRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.DataplexEnvironmentGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *environmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex environment", "name", a.id)

	req := &pb.DeleteEnvironmentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEnvironment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent dataplex environment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex environment %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex environment", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of dataplex environment %s: %w", a.id.String(), err)
	}
	return true, nil
}
