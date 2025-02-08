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

package security

import (
	"context"
	"fmt"
	"reflect"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/security/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/security/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	securitypb "cloud.google.com/go/security/v1beta1/securitypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.SecurityReusableConfigGVK, NewReusableConfigModel)
}

func NewReusableConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelReusableConfig{config: *config}, nil
}

var _ directbase.Model = &modelReusableConfig{}

type modelReusableConfig struct {
	config config.ControllerConfig
}

func (m *modelReusableConfig) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ReusableConfig client: %w", err)
	}
	return gcpClient, err
}

func (m *modelReusableConfig) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SecurityReusableConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewReusableConfigIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get security GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReusableConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelReusableConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ReusableConfigAdapter struct {
	id        *krm.ReusableConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.SecurityReusableConfig
	actual    *securitypb.ReusableConfig
}

var _ directbase.Adapter = &ReusableConfigAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ReusableConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ReusableConfig", "name", a.id)

	req := &securitypb.GetReusableConfigRequest{Name: a.id}
	reusableconfigpb, err := a.gcpClient.GetReusableConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ReusableConfig %q: %w", a.id, err)
	}

	a.actual = reusableconfigpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReusableConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ReusableConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecurityReusableConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &securitypb.CreateReusableConfigRequest{
		Parent:         a.id.Parent().String(),
		ReusableConfig: resource,
	}
	op, err := a.gcpClient.CreateReusableConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ReusableConfig %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ReusableConfig %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ReusableConfig", "name", a.id)

	status := &krm.SecurityReusableConfigStatus{}
	status.ObservedState = SecurityReusableConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReusableConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ReusableConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := SecurityReusableConfigSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
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
		status := &krm.SecurityReusableConfigStatus{}
		status.ObservedState = SecurityReusableConfigObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &securitypb.UpdateReusableConfigRequest{
		Name:           a.id.External,
		UpdateMask:     updateMask,
		ReusableConfig: desiredPb,
	}
	op, err := a.gcpClient.UpdateReusableConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ReusableConfig %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ReusableConfig %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated ReusableConfig", "name", a.id.External)

	status := &krm.SecurityReusableConfigStatus{}
	status.ObservedState = SecurityReusableConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReusableConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecurityReusableConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecurityReusableConfigSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.SecurityReusableConfigGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReusableConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ReusableConfig", "name", a.id)

	req := &securitypb.DeleteReusableConfigRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteReusableConfig(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting ReusableConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ReusableConfig", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ReusableConfig %s: %w", a.id, err)
	}
	return true, nil
}
