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

package iam

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/iam/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	iampb "cloud.google.com/go/iam/v1beta/iampb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.IamWorkloadIdentityPoolProviderGVK, NewWorkloadIdentityPoolProviderModel)
}

func NewWorkloadIdentityPoolProviderModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkloadIdentityPoolProvider{config: *config}, nil
}

var _ directbase.Model = &modelWorkloadIdentityPoolProvider{}

type modelWorkloadIdentityPoolProvider struct {
	config config.ControllerConfig
}

func (m *modelWorkloadIdentityPoolProvider) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadIdentityPoolProvider client: %w", err)
	}
	return gcpClient, err
}

func (m *modelWorkloadIdentityPoolProvider) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.IamWorkloadIdentityPoolProvider{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewWorkloadIdentityPoolProviderIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get iam GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WorkloadIdentityPoolProviderAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWorkloadIdentityPoolProvider) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type WorkloadIdentityPoolProviderAdapter struct {
	id        *krm.WorkloadIdentityPoolProviderIdentity
	gcpClient *gcp.Client
	desired   *krm.IamWorkloadIdentityPoolProvider
	actual    *iampb.WorkloadIdentityPoolProvider
}

var _ directbase.Adapter = &WorkloadIdentityPoolProviderAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *WorkloadIdentityPoolProviderAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting WorkloadIdentityPoolProvider", "name", a.id)

	req := &iampb.GetWorkloadIdentityPoolProviderRequest{Name: a.id}
	workloadidentitypoolproviderpb, err := a.gcpClient.GetWorkloadIdentityPoolProvider(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WorkloadIdentityPoolProvider %q: %w", a.id, err)
	}

	a.actual = workloadidentitypoolproviderpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WorkloadIdentityPoolProviderAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating WorkloadIdentityPoolProvider", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := IamWorkloadIdentityPoolProviderSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &iampb.CreateWorkloadIdentityPoolProviderRequest{
		Parent:                       a.id.Parent().String(),
		WorkloadIdentityPoolProvider: resource,
	}
	op, err := a.gcpClient.CreateWorkloadIdentityPoolProvider(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WorkloadIdentityPoolProvider %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkloadIdentityPoolProvider %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created WorkloadIdentityPoolProvider", "name", a.id)

	status := &krm.IamWorkloadIdentityPoolProviderStatus{}
	status.ObservedState = IamWorkloadIdentityPoolProviderObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WorkloadIdentityPoolProviderAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating WorkloadIdentityPoolProvider", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := IamWorkloadIdentityPoolProviderSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
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
		status := &krm.IamWorkloadIdentityPoolProviderStatus{}
		status.ObservedState = IamWorkloadIdentityPoolProviderObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &iampb.UpdateWorkloadIdentityPoolProviderRequest{
		Name:                         a.id.External,
		UpdateMask:                   updateMask,
		WorkloadIdentityPoolProvider: desiredPb,
	}
	op, err := a.gcpClient.UpdateWorkloadIdentityPoolProvider(ctx, req)
	if err != nil {
		return fmt.Errorf("updating WorkloadIdentityPoolProvider %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkloadIdentityPoolProvider %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated WorkloadIdentityPoolProvider", "name", a.id.External)

	status := &krm.IamWorkloadIdentityPoolProviderStatus{}
	status.ObservedState = IamWorkloadIdentityPoolProviderObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *WorkloadIdentityPoolProviderAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.IamWorkloadIdentityPoolProvider{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(IamWorkloadIdentityPoolProviderSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.IamWorkloadIdentityPoolProviderGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *WorkloadIdentityPoolProviderAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting WorkloadIdentityPoolProvider", "name", a.id)

	req := &iampb.DeleteWorkloadIdentityPoolProviderRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteWorkloadIdentityPoolProvider(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting WorkloadIdentityPoolProvider %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted WorkloadIdentityPoolProvider", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete WorkloadIdentityPoolProvider %s: %w", a.id, err)
	}
	return true, nil
}
