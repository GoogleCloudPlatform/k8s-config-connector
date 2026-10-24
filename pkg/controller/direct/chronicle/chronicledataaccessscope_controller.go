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
// krm.group: chronicle.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ChronicleDataAccessScope

package chronicle

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/chronicle/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/chronicle/apiv1"
	chroniclepb "cloud.google.com/go/chronicle/apiv1/chroniclepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ChronicleDataAccessScopeGVK, NewDataAccessScopeModel)
}

func NewDataAccessScopeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDataAccessScope{config: *config}, nil
}

var _ directbase.Model = &modelDataAccessScope{}

type modelDataAccessScope struct {
	config config.ControllerConfig
}

func (m *modelDataAccessScope) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ChronicleDataAccessScope{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ChronicleDataAccessScopeIdentity)

	// Get chronicle GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DataAccessScopeAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDataAccessScope) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//chronicle.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ChronicleDataAccessScopeIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DataAccessScopeAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDataAccessScope) client(ctx context.Context) (*gcp.DataAccessControlClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataAccessControlRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataAccessScope client: %w", err)
	}
	return gcpClient, err
}

type DataAccessScopeAdapter struct {
	id        *krm.ChronicleDataAccessScopeIdentity
	gcpClient *gcp.DataAccessControlClient
	desired   *krm.ChronicleDataAccessScope
	actual    *chroniclepb.DataAccessScope
}

var _ directbase.Adapter = &DataAccessScopeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DataAccessScopeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataAccessScope", "name", a.id)

	req := &chroniclepb.GetDataAccessScopeRequest{
		Name: a.id.String(),
	}
	dataaccessscopepb, err := a.gcpClient.GetDataAccessScope(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataAccessScope %q: %w", a.id, err)
	}

	a.actual = dataaccessscopepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAccessScopeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataAccessScope", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ChronicleDataAccessScopeSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &chroniclepb.CreateDataAccessScopeRequest{
		Parent:            a.id.ParentString(),
		DataAccessScopeId: a.id.DataAccessScope,
		DataAccessScope:   resource,
	}
	created, err := a.gcpClient.CreateDataAccessScope(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataAccessScope %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DataAccessScope", "name", a.id)

	status := &krm.ChronicleDataAccessScopeStatus{}
	status.ObservedState = ChronicleDataAccessScopeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAccessScopeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataAccessScope", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ChronicleDataAccessScopeSpec_FromProto,
		ChronicleDataAccessScopeSpec_ToProto,
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
		req := &chroniclepb.UpdateDataAccessScopeRequest{
			UpdateMask:      updateMask,
			DataAccessScope: resource,
		}
		updated, err = a.gcpClient.UpdateDataAccessScope(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DataAccessScope %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DataAccessScope", "name", a.id)
	}

	status := &krm.ChronicleDataAccessScopeStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ChronicleDataAccessScopeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DataAccessScopeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ChronicleDataAccessScope{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ChronicleDataAccessScopeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ChronicleDataAccessScopeGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DataAccessScopeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataAccessScope", "name", a.id)

	req := &chroniclepb.DeleteDataAccessScopeRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteDataAccessScope(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DataAccessScope, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DataAccessScope %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DataAccessScope", "name", a.id)
	return true, nil
}
