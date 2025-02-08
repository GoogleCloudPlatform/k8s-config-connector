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

package securitycenter

import (
	"context"
	"fmt"
	"reflect"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/securitycenter/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	securitycenterpb "cloud.google.com/go/securitycenter/v1/securitycenterpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleGVK, NewEffectiveSecurityHealthAnalyticsCustomModuleModel)
}

func NewEffectiveSecurityHealthAnalyticsCustomModuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEffectiveSecurityHealthAnalyticsCustomModule{config: *config}, nil
}

var _ directbase.Model = &modelEffectiveSecurityHealthAnalyticsCustomModule{}

type modelEffectiveSecurityHealthAnalyticsCustomModule struct {
	config config.ControllerConfig
}

func (m *modelEffectiveSecurityHealthAnalyticsCustomModule) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building EffectiveSecurityHealthAnalyticsCustomModule client: %w", err)
	}
	return gcpClient, err
}

func (m *modelEffectiveSecurityHealthAnalyticsCustomModule) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SecuritycenterEffectiveSecurityHealthAnalyticsCustomModule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEffectiveSecurityHealthAnalyticsCustomModuleIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get securitycenter GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EffectiveSecurityHealthAnalyticsCustomModuleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEffectiveSecurityHealthAnalyticsCustomModule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EffectiveSecurityHealthAnalyticsCustomModuleAdapter struct {
	id        *krm.EffectiveSecurityHealthAnalyticsCustomModuleIdentity
	gcpClient *gcp.Client
	desired   *krm.SecuritycenterEffectiveSecurityHealthAnalyticsCustomModule
	actual    *securitycenterpb.EffectiveSecurityHealthAnalyticsCustomModule
}

var _ directbase.Adapter = &EffectiveSecurityHealthAnalyticsCustomModuleAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EffectiveSecurityHealthAnalyticsCustomModuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EffectiveSecurityHealthAnalyticsCustomModule", "name", a.id)

	req := &securitycenterpb.GetEffectiveSecurityHealthAnalyticsCustomModuleRequest{Name: a.id}
	effectivesecurityhealthanalyticscustommodulepb, err := a.gcpClient.GetEffectiveSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EffectiveSecurityHealthAnalyticsCustomModule %q: %w", a.id, err)
	}

	a.actual = effectivesecurityhealthanalyticscustommodulepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EffectiveSecurityHealthAnalyticsCustomModuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EffectiveSecurityHealthAnalyticsCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &securitycenterpb.CreateEffectiveSecurityHealthAnalyticsCustomModuleRequest{
		Parent: a.id.Parent().String(),
		EffectiveSecurityHealthAnalyticsCustomModule: resource,
	}
	op, err := a.gcpClient.CreateEffectiveSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating EffectiveSecurityHealthAnalyticsCustomModule %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("EffectiveSecurityHealthAnalyticsCustomModule %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created EffectiveSecurityHealthAnalyticsCustomModule", "name", a.id)

	status := &krm.SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleStatus{}
	status.ObservedState = SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EffectiveSecurityHealthAnalyticsCustomModuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EffectiveSecurityHealthAnalyticsCustomModule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
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
		status := &krm.SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleStatus{}
		status.ObservedState = SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &securitycenterpb.UpdateEffectiveSecurityHealthAnalyticsCustomModuleRequest{
		Name:       a.id.External,
		UpdateMask: updateMask,
		EffectiveSecurityHealthAnalyticsCustomModule: desiredPb,
	}
	op, err := a.gcpClient.UpdateEffectiveSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		return fmt.Errorf("updating EffectiveSecurityHealthAnalyticsCustomModule %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("EffectiveSecurityHealthAnalyticsCustomModule %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated EffectiveSecurityHealthAnalyticsCustomModule", "name", a.id.External)

	status := &krm.SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleStatus{}
	status.ObservedState = SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EffectiveSecurityHealthAnalyticsCustomModuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecuritycenterEffectiveSecurityHealthAnalyticsCustomModule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.SecuritycenterEffectiveSecurityHealthAnalyticsCustomModuleGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EffectiveSecurityHealthAnalyticsCustomModuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EffectiveSecurityHealthAnalyticsCustomModule", "name", a.id)

	req := &securitycenterpb.DeleteEffectiveSecurityHealthAnalyticsCustomModuleRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEffectiveSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting EffectiveSecurityHealthAnalyticsCustomModule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted EffectiveSecurityHealthAnalyticsCustomModule", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete EffectiveSecurityHealthAnalyticsCustomModule %s: %w", a.id, err)
	}
	return true, nil
}
