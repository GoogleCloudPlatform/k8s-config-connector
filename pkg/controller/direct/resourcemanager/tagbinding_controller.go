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

package resourcemanager

import (
	"context"
	"fmt"
	"reflect"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/resourcemanager/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	resourcemanagerpb "cloud.google.com/go/resourcemanager/v3/resourcemanagerpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ResourcemanagerTagBindingGVK, NewTagBindingModel)
}

func NewTagBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTagBinding{config: *config}, nil
}

var _ directbase.Model = &modelTagBinding{}

type modelTagBinding struct {
	config config.ControllerConfig
}

func (m *modelTagBinding) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TagBinding client: %w", err)
	}
	return gcpClient, err
}

func (m *modelTagBinding) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ResourcemanagerTagBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTagBindingIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get resourcemanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TagBindingAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTagBinding) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type TagBindingAdapter struct {
	id        *krm.TagBindingIdentity
	gcpClient *gcp.Client
	desired   *krm.ResourcemanagerTagBinding
	actual    *resourcemanagerpb.TagBinding
}

var _ directbase.Adapter = &TagBindingAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TagBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TagBinding", "name", a.id)

	req := &resourcemanagerpb.GetTagBindingRequest{Name: a.id}
	tagbindingpb, err := a.gcpClient.GetTagBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TagBinding %q: %w", a.id, err)
	}

	a.actual = tagbindingpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TagBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TagBinding", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ResourcemanagerTagBindingSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &resourcemanagerpb.CreateTagBindingRequest{
		Parent:     a.id.Parent().String(),
		TagBinding: resource,
	}
	op, err := a.gcpClient.CreateTagBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TagBinding %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("TagBinding %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created TagBinding", "name", a.id)

	status := &krm.ResourcemanagerTagBindingStatus{}
	status.ObservedState = ResourcemanagerTagBindingObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TagBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating TagBinding", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := ResourcemanagerTagBindingSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
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
		status := &krm.ResourcemanagerTagBindingStatus{}
		status.ObservedState = ResourcemanagerTagBindingObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &resourcemanagerpb.UpdateTagBindingRequest{
		Name:       a.id.External,
		UpdateMask: updateMask,
		TagBinding: desiredPb,
	}
	op, err := a.gcpClient.UpdateTagBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("updating TagBinding %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("TagBinding %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated TagBinding", "name", a.id.External)

	status := &krm.ResourcemanagerTagBindingStatus{}
	status.ObservedState = ResourcemanagerTagBindingObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TagBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ResourcemanagerTagBinding{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ResourcemanagerTagBindingSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.ResourcemanagerTagBindingGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TagBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TagBinding", "name", a.id)

	req := &resourcemanagerpb.DeleteTagBindingRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTagBinding(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting TagBinding %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted TagBinding", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete TagBinding %s: %w", a.id, err)
	}
	return true, nil
}
