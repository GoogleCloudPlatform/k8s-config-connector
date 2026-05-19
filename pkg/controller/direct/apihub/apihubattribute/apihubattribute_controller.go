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

package apihubattribute

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/apihub/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	apihubpb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.APIHubAttributeGVK, NewAttributeModel)
}

func NewAttributeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAttribute{config: *config}, nil
}

var _ directbase.Model = &modelAttribute{}

type modelAttribute struct {
	config config.ControllerConfig
}

func (m *modelAttribute) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Attribute client: %w", err)
	}
	return gcpClient, err
}

func (m *modelAttribute) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.APIHubAttribute{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAPIHubAttributeIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get apihub GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AttributeAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAttribute) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AttributeAdapter struct {
	id        *krm.APIHubAttributeIdentity
	gcpClient *gcp.Client
	desired   *krm.APIHubAttribute
	actual    *apihubpb.Attribute
}

var _ directbase.Adapter = &AttributeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AttributeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Attribute", "name", a.id)

	req := &apihubpb.GetAttributeRequest{Name: a.id.String()}
	attributepb, err := a.gcpClient.GetAttribute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Attribute %q: %w", a.id, err)
	}

	a.actual = attributepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AttributeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Attribute", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := APIHubAttributeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &apihubpb.CreateAttributeRequest{
		Parent:      "projects/" + a.id.Project + "/locations/" + a.id.Location,
		AttributeId: a.id.Attribute,
		Attribute:   resource,
	}
	created, err := a.gcpClient.CreateAttribute(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Attribute %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Attribute", "name", a.id)

	status := &krm.APIHubAttributeStatus{}
	status.ObservedState = APIHubAttributeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AttributeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Attribute", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := APIHubAttributeSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])
	// Option 1: This option is good for proto that has `field_mask` for output-only, immutable, required/optional.
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
			paths = paths.Insert("display_name")
		}
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		desiredPb.Name = a.id.String()
		// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
		req := &apihubpb.UpdateAttributeRequest{
			UpdateMask: updateMask,
			Attribute:  desiredPb,
		}
		var err error
		updated, err = a.gcpClient.UpdateAttribute(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Attribute %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Attribute", "name", a.id)
	}

	status := &krm.APIHubAttributeStatus{}
	status.ObservedState = APIHubAttributeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AttributeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.APIHubAttribute{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(APIHubAttributeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Attribute)
	u.SetGroupVersionKind(krm.APIHubAttributeGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AttributeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Attribute", "name", a.id)

	req := &apihubpb.DeleteAttributeRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAttribute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Attribute, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Attribute %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Attribute", "name", a.id)

	return true, nil
}
