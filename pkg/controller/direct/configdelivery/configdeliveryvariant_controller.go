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

// +generated:controller
// krm.group: configdelivery.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ConfigdeliveryVariant

package configdelivery

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdelivery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/configdelivery/apiv1"
	configdeliverypb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ConfigdeliveryVariantGVK, NewVariantModel)
}

func NewVariantModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelVariant{config: *config}, nil
}

var _ directbase.Model = &modelVariant{}

type modelVariant struct {
	config config.ControllerConfig
}

func (m *modelVariant) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConfigdeliveryVariant{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ConfigdeliveryVariantIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get configdelivery GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &VariantAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelVariant) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//configdelivery.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ConfigdeliveryVariantIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &VariantAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelVariant) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Variant client: %w", err)
	}
	return gcpClient, err
}

type VariantAdapter struct {
	id        *krm.ConfigdeliveryVariantIdentity
	gcpClient *gcp.Client
	desired   *krm.ConfigdeliveryVariant
	actual    *configdeliverypb.Variant
}

var _ directbase.Adapter = &VariantAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *VariantAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Variant", "name", a.id)

	req := &configdeliverypb.GetVariantRequest{
		Name: a.id.String(),
	}
	variantpb, err := a.gcpClient.GetVariant(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Variant %q: %w", a.id, err)
	}

	a.actual = variantpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *VariantAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Variant", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ConfigdeliveryVariantSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	req := &configdeliverypb.CreateVariantRequest{
		Parent:    a.id.ParentString(),
		VariantId: a.id.Variant,
		Variant:   resource,
	}
	op, err := a.gcpClient.CreateVariant(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Variant %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Variant %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Variant", "name", a.id)

	status := &krm.ConfigdeliveryVariantStatus{}
	status.ObservedState = ConfigdeliveryVariantObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *VariantAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Variant", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ConfigdeliveryVariantSpec_FromProto,
		ConfigdeliveryVariantSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &configdeliverypb.UpdateVariantRequest{
			UpdateMask: updateMask,
			Variant:    resource,
		}
		op, err := a.gcpClient.UpdateVariant(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Variant %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Variant %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Variant", "name", a.id)
	}

	status := &krm.ConfigdeliveryVariantStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ConfigdeliveryVariantObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *VariantAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConfigdeliveryVariant{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConfigdeliveryVariantSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ConfigdeliveryVariantGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *VariantAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Variant", "name", a.id)

	req := &configdeliverypb.DeleteVariantRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteVariant(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Variant, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Variant %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Variant", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Variant %s: %w", a.id, err)
	}
	return true, nil
}
