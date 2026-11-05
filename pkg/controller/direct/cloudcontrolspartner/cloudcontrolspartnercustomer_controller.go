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

// +generated:controller
// krm.group: cloudcontrolspartner.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CloudcontrolspartnerCustomer

package cloudcontrolspartner

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/cloudcontrolspartner/apiv1"
	cloudcontrolspartnerpb "cloud.google.com/go/cloudcontrolspartner/apiv1/cloudcontrolspartnerpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudcontrolspartnerCustomerGVK, NewCustomerModel)
}

func NewCustomerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCustomer{config: *config}, nil
}

var _ directbase.Model = &modelCustomer{}

type modelCustomer struct {
	config config.ControllerConfig
}

func (m *modelCustomer) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudcontrolspartnerCustomer{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CloudcontrolspartnerCustomerIdentity)

	// Get cloudcontrolspartner GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CustomerAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCustomer) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//cloudcontrolspartner.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CloudcontrolspartnerCustomerIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &CustomerAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelCustomer) client(ctx context.Context) (*gcp.CloudControlsPartnerCoreClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudControlsPartnerCoreRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Customer client: %w", err)
	}
	return gcpClient, err
}

type CustomerAdapter struct {
	id        *krm.CloudcontrolspartnerCustomerIdentity
	gcpClient *gcp.CloudControlsPartnerCoreClient
	desired   *krm.CloudcontrolspartnerCustomer
	actual    *cloudcontrolspartnerpb.Customer
}

var _ directbase.Adapter = &CustomerAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *CustomerAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Customer", "name", a.id)

	req := &cloudcontrolspartnerpb.GetCustomerRequest{
		Name: a.id.String(),
	}
	customerpb, err := a.gcpClient.GetCustomer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Customer %q: %w", a.id, err)
	}

	a.actual = customerpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CustomerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Customer", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudcontrolspartnerCustomerSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cloudcontrolspartnerpb.CreateCustomerRequest{
		Parent:     a.id.ParentString(),
		CustomerId: a.id.Customer,
		Customer:   resource,
	}
	created, err := a.gcpClient.CreateCustomer(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Customer %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Customer", "name", a.id)

	status := &krm.CloudcontrolspartnerCustomerStatus{}
	status.ObservedState = CloudcontrolspartnerCustomerObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *CustomerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Customer", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CloudcontrolspartnerCustomerSpec_FromProto,
		CloudcontrolspartnerCustomerSpec_ToProto,
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
		req := &cloudcontrolspartnerpb.UpdateCustomerRequest{
			UpdateMask: updateMask,
			Customer:   resource,
		}
		updated, err = a.gcpClient.UpdateCustomer(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Customer %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Customer", "name", a.id)
	}

	status := &krm.CloudcontrolspartnerCustomerStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudcontrolspartnerCustomerObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *CustomerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudcontrolspartnerCustomer{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudcontrolspartnerCustomerSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CloudcontrolspartnerCustomerGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *CustomerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Customer", "name", a.id)

	req := &cloudcontrolspartnerpb.DeleteCustomerRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteCustomer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Customer, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Customer %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Customer", "name", a.id)
	return true, nil
}
