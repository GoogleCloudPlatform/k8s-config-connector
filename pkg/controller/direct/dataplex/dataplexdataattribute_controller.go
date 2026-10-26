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
// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DataplexDataAttribute

package dataplex

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dataplex/apiv1"
	dataplexpb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DataplexDataAttributeGVK, NewDataAttributeModel)
}

func NewDataAttributeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDataAttribute{config: *config}, nil
}

var _ directbase.Model = &modelDataAttribute{}

type modelDataAttribute struct {
	config config.ControllerConfig
}

func (m *modelDataAttribute) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexDataAttribute{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DataplexDataAttributeIdentity)

	// Get dataplex GCP client
	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &DataAttributeAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDataAttribute) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}
func (m *modelDataAttribute) client(ctx context.Context, location string) (*gcp.DataTaxonomyClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("https://dataplex.%s.rep.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewDataTaxonomyRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataAttribute client: %w", err)
	}
	return gcpClient, err
}

type DataAttributeAdapter struct {
	id        *krm.DataplexDataAttributeIdentity
	gcpClient *gcp.DataTaxonomyClient
	desired   *krm.DataplexDataAttribute
	actual    *dataplexpb.DataAttribute
}

var _ directbase.Adapter = &DataAttributeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DataAttributeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataAttribute", "name", a.id)

	req := &dataplexpb.GetDataAttributeRequest{
		Name: a.id.String(),
	}
	dataattributepb, err := a.gcpClient.GetDataAttribute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataAttribute %q: %w", a.id, err)
	}

	a.actual = dataattributepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAttributeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataAttribute", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataplexDataAttributeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dataplexpb.CreateDataAttributeRequest{
		Parent:          a.id.ParentString(),
		DataAttributeId: a.id.DataAttributeId,
		DataAttribute:   resource,
	}
	op, err := a.gcpClient.CreateDataAttribute(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataAttribute %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DataAttribute %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DataAttribute", "name", a.id)

	status := &krm.DataplexDataAttributeStatus{}
	status.ObservedState = DataplexDataAttributeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAttributeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataAttribute", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DataplexDataAttributeSpec_FromProto,
		DataplexDataAttributeSpec_ToProto,
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
		req := &dataplexpb.UpdateDataAttributeRequest{
			UpdateMask:    updateMask,
			DataAttribute: resource,
		}
		op, err := a.gcpClient.UpdateDataAttribute(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DataAttribute %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DataAttribute %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DataAttribute", "name", a.id)
	}

	status := &krm.DataplexDataAttributeStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DataplexDataAttributeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DataAttributeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataplexDataAttribute{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexDataAttributeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DataplexDataAttributeGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DataAttributeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataAttribute", "name", a.id)

	req := &dataplexpb.DeleteDataAttributeRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteDataAttribute(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DataAttribute, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DataAttribute %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DataAttribute", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete DataAttribute %s: %w", a.id, err)
	}
	return true, nil
}
