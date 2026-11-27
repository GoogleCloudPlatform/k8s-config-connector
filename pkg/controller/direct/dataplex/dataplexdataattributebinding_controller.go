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
// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DataplexDataAttributeBinding

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

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DataplexDataAttributeBindingGVK, NewDataAttributeBindingModel)
}

func NewDataAttributeBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDataAttributeBinding{config: *config}, nil
}

var _ directbase.Model = &modelDataAttributeBinding{}

type modelDataAttributeBinding struct {
	config config.ControllerConfig
}

func (m *modelDataAttributeBinding) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexDataAttributeBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DataplexDataAttributeBindingIdentity)

	// Get dataplex GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	dtClient, err := gcpClient.dataTaxonomyClient(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &DataAttributeBindingAdapter{
		id:        id,
		gcpClient: dtClient,
		desired:   obj,
	}, nil
}

func (m *modelDataAttributeBinding) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DataAttributeBindingAdapter struct {
	id        *krm.DataplexDataAttributeBindingIdentity
	gcpClient *gcp.DataTaxonomyClient
	desired   *krm.DataplexDataAttributeBinding
	actual    *dataplexpb.DataAttributeBinding
}

var _ directbase.Adapter = &DataAttributeBindingAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DataAttributeBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataAttributeBinding", "name", a.id)

	req := &dataplexpb.GetDataAttributeBindingRequest{
		Name: a.id.String(),
	}
	dataattributebindingpb, err := a.gcpClient.GetDataAttributeBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataAttributeBinding %q: %w", a.id, err)
	}

	a.actual = dataattributebindingpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAttributeBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataAttributeBinding", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataplexDataAttributeBindingSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dataplexpb.CreateDataAttributeBindingRequest{
		Parent:                 fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		DataAttributeBindingId: a.id.Dataattributebinding,
		DataAttributeBinding:   resource,
	}
	op, err := a.gcpClient.CreateDataAttributeBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataAttributeBinding %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DataAttributeBinding %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DataAttributeBinding", "name", a.id)

	status := &krm.DataplexDataAttributeBindingStatus{}
	status.ObservedState = DataplexDataAttributeBindingObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAttributeBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataAttributeBinding", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DataplexDataAttributeBindingSpec_FromProto,
		DataplexDataAttributeBindingSpec_ToProto,
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
		req := &dataplexpb.UpdateDataAttributeBindingRequest{
			UpdateMask:           updateMask,
			DataAttributeBinding: resource,
		}
		op, err := a.gcpClient.UpdateDataAttributeBinding(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DataAttributeBinding %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DataAttributeBinding %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DataAttributeBinding", "name", a.id)
	}

	status := &krm.DataplexDataAttributeBindingStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DataplexDataAttributeBindingObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DataAttributeBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataplexDataAttributeBinding{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexDataAttributeBindingSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DataplexDataAttributeBindingGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DataAttributeBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataAttributeBinding", "name", a.id)

	req := &dataplexpb.DeleteDataAttributeBindingRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteDataAttributeBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DataAttributeBinding, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DataAttributeBinding %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DataAttributeBinding", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete DataAttributeBinding %s: %w", a.id, err)
	}
	return true, nil
}
