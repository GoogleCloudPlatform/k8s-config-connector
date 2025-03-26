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
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.CustomTargetType
// crd.type: DeployCustomTargetType
// crd.version: v1alpha1

package deploy

import (
	"context"
	"fmt"
	"reflect"

	deploy "cloud.google.com/go/deploy/apiv1"
	deploypb "cloud.google.com/go/deploy/apiv1/deploypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/deploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.DeployCustomTargetTypeGVK, NewCustomTargetTypeModel)
}

func NewCustomTargetTypeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &customTargetTypeModel{config: *config}, nil
}

var _ directbase.Model = &customTargetTypeModel{}

type customTargetTypeModel struct {
	config config.ControllerConfig
}

func (m *customTargetTypeModel) client(ctx context.Context, projectID string) (*deploy.Client, error) {
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

	gcpClient, err := deploy.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building deploy custom target type client: %w", err)
	}

	return gcpClient, err
}

func (m *customTargetTypeModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DeployCustomTargetType{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewCustomTargetTypeIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &customTargetTypeAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *customTargetTypeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type customTargetTypeAdapter struct {
	gcpClient *deploy.Client
	id        *krm.CustomTargetTypeIdentity
	desired   *krm.DeployCustomTargetType
	actual    *deploypb.CustomTargetType
}

var _ directbase.Adapter = &customTargetTypeAdapter{}

func (a *customTargetTypeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting deploy custom target type", "name", a.id)

	req := &deploypb.GetCustomTargetTypeRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCustomTargetType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting deploy custom target type %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *customTargetTypeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating deploy custom target type", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DeployCustomTargetTypeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &deploypb.CreateCustomTargetTypeRequest{
		Parent:             a.id.Parent().String(),
		CustomTargetTypeId: a.id.ID(),
		CustomTargetType:   resource,
	}
	op, err := a.gcpClient.CreateCustomTargetType(ctx, req)
	if err != nil {
		return fmt.Errorf("creating deploy custom target type %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("deploy custom target type %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created deploy custom target type in gcp", "name", a.id)

	status := &krm.DeployCustomTargetTypeStatus{}
	status.ObservedState = DeployCustomTargetTypeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *customTargetTypeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating deploy custom target type", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DeployCustomTargetTypeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.GetDescription(), a.actual.GetDescription()) {
		paths = append(paths, "description")
	}
	if desired.Spec.Annotations != nil && !reflect.DeepEqual(resource.GetAnnotations(), a.actual.GetAnnotations()) {
		paths = append(paths, "annotations")
	}
	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.GetLabels(), a.actual.GetLabels()) {
		paths = append(paths, "labels")
	}
	if desired.Spec.CustomActions != nil && !reflect.DeepEqual(resource.GetCustomActions(), a.actual.GetCustomActions()) {
		paths = append(paths, "custom_actions")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
	req := &deploypb.UpdateCustomTargetTypeRequest{
		CustomTargetType: resource,
		UpdateMask:       &fieldmaskpb.FieldMask{Paths: paths},
	}
	op, err := a.gcpClient.UpdateCustomTargetType(ctx, req)
	if err != nil {
		return fmt.Errorf("updating deploy custom target type %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("deploy custom target type %s waiting for update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated deploy custom target type", "name", a.id)

	status := &krm.DeployCustomTargetTypeStatus{}
	status.ObservedState = DeployCustomTargetTypeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *customTargetTypeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DeployCustomTargetType{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DeployCustomTargetTypeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.DeployCustomTargetTypeGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *customTargetTypeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting deploy custom target type", "name", a.id)

	req := &deploypb.DeleteCustomTargetTypeRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCustomTargetType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent deploy custom target type, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting deploy custom target type %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted deploy custom target type", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete deploy custom target type %s: %w", a.id, err)
	}
	return true, nil
}
```
</out>


