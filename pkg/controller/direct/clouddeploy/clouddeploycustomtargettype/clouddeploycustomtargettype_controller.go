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

package clouddeploycustomtargettype

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/clouddeploy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/deploy/apiv1"
	clouddeploypb "cloud.google.com/go/deploy/apiv1/deploypb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CustomTargetTypeGVK, NewCustomTargetTypeModel)
}

func NewCustomTargetTypeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCustomTargetType{config: *config}, nil
}

var _ directbase.Model = &modelCustomTargetType{}

type modelCustomTargetType struct {
	config config.ControllerConfig
}

func (m *modelCustomTargetType) client(ctx context.Context) (*gcp.CloudDeployClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudDeployRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CustomTargetType client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCustomTargetType) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudDeployCustomTargetType{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewCustomTargetTypeIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get clouddeploy GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CustomTargetTypeAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCustomTargetType) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type CustomTargetTypeAdapter struct {
	id        *krm.CustomTargetTypeIdentity
	gcpClient *gcp.CloudDeployClient
	desired   *krm.CloudDeployCustomTargetType
	actual    *clouddeploypb.CustomTargetType
}

var _ directbase.Adapter = &CustomTargetTypeAdapter{}

// Find retrieves the GCP resource.
func (a *CustomTargetTypeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CustomTargetType", "name", a.id)

	req := &clouddeploypb.GetCustomTargetTypeRequest{Name: a.id.String()}
	customtargettypepb, err := a.gcpClient.GetCustomTargetType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CustomTargetType %q: %w", a.id, err)
	}

	a.actual = customtargettypepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CustomTargetTypeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CustomTargetType", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := clouddeploy.CustomTargetTypeSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &clouddeploypb.CreateCustomTargetTypeRequest{
		Parent:             a.id.Parent().String(),
		CustomTargetType:   resource,
		CustomTargetTypeId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateCustomTargetType(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CustomTargetType %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("CustomTargetType %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created CustomTargetType", "name", a.id)

	status := &krm.CustomTargetTypeStatus{}
	status.ObservedState = clouddeploy.CustomTargetTypeObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CustomTargetTypeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CustomTargetType", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := clouddeploy.CustomTargetTypeSpec_v1alpha1_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}

	desiredPb.Name = a.id.String()
	req := &clouddeploypb.UpdateCustomTargetTypeRequest{
		UpdateMask:       updateMask,
		CustomTargetType: desiredPb,
	}
	op, err := a.gcpClient.UpdateCustomTargetType(ctx, req)
	if err != nil {
		return fmt.Errorf("updating CustomTargetType %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("CustomTargetType %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated CustomTargetType", "name", a.id)

	status := &krm.CustomTargetTypeStatus{}
	status.ObservedState = clouddeploy.CustomTargetTypeObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CustomTargetTypeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDeployCustomTargetType{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(clouddeploy.CustomTargetTypeSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.CustomTargetTypeGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CustomTargetTypeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CustomTargetType", "name", a.id)

	req := &clouddeploypb.DeleteCustomTargetTypeRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCustomTargetType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent CustomTargetType, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CustomTargetType %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CustomTargetType", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CustomTargetType %s: %w", a.id, err)
	}
	return true, nil
}
