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
// krm.kind: DataplexEntry

package dataplex

import (
	"context"
	"fmt"

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
	registry.RegisterModel(krm.DataplexEntryGVK, NewEntryModel)
}

func NewEntryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEntry{config: *config}, nil
}

var _ directbase.Model = &modelEntry{}

type modelEntry struct {
	config config.ControllerConfig
}

func (m *modelEntry) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexEntry{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DataplexEntryIdentity)

	// Get dataplex GCP client
	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &EntryAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEntry) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}
func (m *modelEntry) client(ctx context.Context, location string) (*gcp.CatalogClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("https://dataplex.%s.rep.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewCatalogRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Entry client: %w", err)
	}
	return gcpClient, err
}

type EntryAdapter struct {
	id        *krm.DataplexEntryIdentity
	gcpClient *gcp.CatalogClient
	desired   *krm.DataplexEntry
	actual    *dataplexpb.Entry
}

var _ directbase.Adapter = &EntryAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EntryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Entry", "name", a.id)

	req := &dataplexpb.GetEntryRequest{
		Name: a.id.String(),
	}
	entrypb, err := a.gcpClient.GetEntry(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Entry %q: %w", a.id, err)
	}

	a.actual = entrypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EntryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Entry", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataplexEntrySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dataplexpb.CreateEntryRequest{
		Parent:  a.id.ParentString(),
		EntryId: a.id.Entry,
		Entry:   resource,
	}
	created, err := a.gcpClient.CreateEntry(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Entry %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Entry", "name", a.id)

	status := &krm.DataplexEntryStatus{}
	status.ObservedState = DataplexEntryObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EntryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Entry", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DataplexEntrySpec_FromProto,
		DataplexEntrySpec_ToProto,
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
		req := &dataplexpb.UpdateEntryRequest{
			UpdateMask: updateMask,
			Entry:      resource,
		}
		updated, err = a.gcpClient.UpdateEntry(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Entry %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Entry", "name", a.id)
	}

	status := &krm.DataplexEntryStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DataplexEntryObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EntryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataplexEntry{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexEntrySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DataplexEntryGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EntryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Entry", "name", a.id)

	req := &dataplexpb.DeleteEntryRequest{
		Name: a.id.String(),
	}
	_, err := a.gcpClient.DeleteEntry(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Entry, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Entry %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Entry", "name", a.id)
	return true, nil
}
