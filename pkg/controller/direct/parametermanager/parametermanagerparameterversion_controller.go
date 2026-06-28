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

package parametermanager

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/parametermanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/parametermanager/apiv1"
	parametermanagerpb "cloud.google.com/go/parametermanager/apiv1/parametermanagerpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ParameterManagerParameterVersionGVK, NewParameterVersionModel)
}

func NewParameterVersionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelParameterVersion{config: *config}, nil
}

var _ directbase.Model = &modelParameterVersion{}

type modelParameterVersion struct {
	config config.ControllerConfig
}

func (m *modelParameterVersion) client(ctx context.Context, location string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	// Add regional endpoint if region is specified
	if location != "" && location != "global" {
		endpoint := fmt.Sprintf("parametermanager.%s.rep.googleapis.com:443", location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Parameter Version client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelParameterVersion) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ParameterManagerParameterVersion{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	parameter := id.(*krm.ParameterVersionIdentity).Parent()

	location := parameter.Parent().Location

	// Get parmetermanager GCP client
	gcpClient, err := m.client(ctx, location)
	if err != nil {
		return nil, err
	}
	return &ParameterVersionAdapter{
		id:        id.(*krm.ParameterVersionIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelParameterVersion) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ParameterVersionAdapter struct {
	id        *krm.ParameterVersionIdentity
	gcpClient *gcp.Client
	desired   *krm.ParameterManagerParameterVersion
	actual    *parametermanagerpb.ParameterVersion
}

var _ directbase.Adapter = &ParameterVersionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ParameterVersionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Parameter Version", "name", a.id.String())

	req := &parametermanagerpb.GetParameterVersionRequest{Name: a.id.String()}
	parameterVersionPb, err := a.gcpClient.GetParameterVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Parameter Version %q: %w", a.id.String(), err)
	}

	a.actual = parameterVersionPb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *ParameterVersionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Parameter Version", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ParameterManagerParameterVersionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &parametermanagerpb.CreateParameterVersionRequest{
		Parent:             a.id.Parent().String(),
		ParameterVersion:   resource,
		ParameterVersionId: a.id.ID(),
	}
	created, err := a.gcpClient.CreateParameterVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Parameter Version %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Parameter Version", "name", a.id.String())

	status := &krm.ParameterManagerParameterVersionStatus{}
	status.ObservedState = ParameterManagerParameterVersionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and updates the Config Connector object `status` based on the GCP response.
func (a *ParameterVersionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Parameter Version", "name", a.id.String())
	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := ParameterManagerParameterVersionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.String()

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	// If parameterversion is disabled, payload is not available for retrieval leading to mismatch
	// For disabled parameterversion, ignore difference in payload
	if a.actual.GetDisabled() {
		paths = paths.Delete("payload.data")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		if a.desired.Status.ExternalRef == nil {
			status := &krm.ParameterManagerParameterVersionStatus{}
			status.ObservedState = ParameterManagerParameterVersionObservedState_FromProto(mapCtx, a.actual)
			if mapCtx.Err() != nil {
				return mapCtx.Err()
			}
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	req := &parametermanagerpb.UpdateParameterVersionRequest{
		UpdateMask:       updateMask,
		ParameterVersion: resource,
	}
	updated, err := a.gcpClient.UpdateParameterVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Parameter Version %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated Parameter Version", "name", a.id.String())

	status := &krm.ParameterManagerParameterVersionStatus{}
	status.ObservedState = ParameterManagerParameterVersionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.desired.Status.ExternalRef == nil {
		status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ParameterVersionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ParameterManagerParameterVersion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ParameterManagerParameterVersionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	externalRef := a.actual.GetName()
	var id *krm.ParameterVersionIdentity
	if err := id.FromExternal(externalRef); err != nil {
		return nil, fmt.Errorf("parsing external ref %q: %w", externalRef, err)
	}

	obj.Spec.ParameterRef = &krm.ParameterRef{External: id.Parent().String()}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.ParameterManagerParameterVersionGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ParameterVersionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Parameter Version", "name", a.id.String())

	req := &parametermanagerpb.DeleteParameterVersionRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteParameterVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Parameter Version, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Parameter Version %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Parameter Version", "name", a.id.String())

	return true, nil
}
