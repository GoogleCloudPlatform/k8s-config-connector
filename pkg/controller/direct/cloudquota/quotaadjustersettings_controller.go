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
// proto.service: google.api.cloudquotas.v1beta.QuotaAdjusterSettingsManager
// proto.message: google.api.cloudquotas.v1beta.QuotaAdjusterSettings
// crd.type: APIQuotaAdjusterSettings
// crd.version: v1alpha1

package cloudquota

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/cloudquotas/apiv1beta"
	pb "cloud.google.com/go/cloudquotas/apiv1beta/cloudquotaspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudquota/v1alpha1"
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
	registry.RegisterModel(krm.APIQuotaAdjusterSettingsGVK, NewQuotaAdjusterSettingsModel)
}

func NewQuotaAdjusterSettingsModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &apiQuotaAdjusterSettingsModel{config: config}, nil
}

var _ directbase.Model = &apiQuotaAdjusterSettingsModel{}

type apiQuotaAdjusterSettingsModel struct {
	config *config.ControllerConfig
}

func (m *apiQuotaAdjusterSettingsModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.APIQuotaAdjusterSettings{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewQuotaAdjusterSettingsIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get cloudquotas GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newQuotaAdjusterSettingsManagerClient(ctx)
	if err != nil {
		return nil, err
	}
	return &apiQuotaAdjusterSettingsAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *apiQuotaAdjusterSettingsModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type apiQuotaAdjusterSettingsAdapter struct {
	gcpClient *gcp.QuotaAdjusterSettingsManagerClient
	id        *krm.QuotaAdjusterSettingsIdentity
	desired   *krm.APIQuotaAdjusterSettings
	actual    *pb.QuotaAdjusterSettings
}

var _ directbase.Adapter = &apiQuotaAdjusterSettingsAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
func (a *apiQuotaAdjusterSettingsAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting cloudquotas quotaadjustersettings", "name", a.id)

	req := &pb.GetQuotaAdjusterSettingsRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetQuotaAdjusterSettings(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting cloudquotas quotaadjustersettings %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create is not supported for QuotaAdjusterSettings. This resource is managed by GCP implicitly.
func (a *apiQuotaAdjusterSettingsAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("create operation is not supported for resource %s", krm.APIQuotaAdjusterSettingsGVK)
	return fmt.Errorf("create operation is not supported for resource %s", krm.APIQuotaAdjusterSettingsGVK)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *apiQuotaAdjusterSettingsAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating cloudquotas quotaadjustersettings", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := APIQuotaAdjusterSettingsSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Enablement != nil && !reflect.DeepEqual(resource.Enablement, a.actual.Enablement) {
		paths = append(paths, "enablement")
	}

	var updated *pb.QuotaAdjusterSettings
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// Even if there are no updates, we still need to update the status
		// for potential acquisition scenarios.
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		// Set etag for optimistic concurrency control if provided
		if a.desired.Status.ObservedState.Etag != nil {
			resource.Etag = *a.desired.Status.ObservedState.Etag
		} else {
			// Use the etag from the last read if not specified in spec
			resource.Etag = a.actual.Etag
		}
		req := &pb.UpdateQuotaAdjusterSettingsRequest{
			QuotaAdjusterSettings: resource,
			UpdateMask:            &fieldmaskpb.FieldMask{Paths: paths},
		}
		var err error
		updated, err = a.gcpClient.UpdateQuotaAdjusterSettings(ctx, req)
		if err != nil {
			return fmt.Errorf("updating cloudquotas quotaadjustersettings %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated cloudquotas quotaadjustersettings", "name", a.id)
	}

	status := &krm.APIQuotaAdjusterSettingsStatus{}
	status.ObservedState = APIQuotaAdjusterSettingsObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *apiQuotaAdjusterSettingsAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.APIQuotaAdjusterSettings{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(APIQuotaAdjusterSettingsSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID()) // Name is fixed as 'quotaAdjusterSettings'
	u.SetGroupVersionKind(krm.APIQuotaAdjusterSettingsGVK)
	u.Object = uObj
	return u, nil
}

// Delete is not supported for QuotaAdjusterSettings.
func (a *apiQuotaAdjusterSettingsAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("delete operation is not supported for QuotaAdjusterSettings, treating as no-op", "name", a.id.String())
	return false, nil // Indicate that deletion is not supported and should not be retried.
}
