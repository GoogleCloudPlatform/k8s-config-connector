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
// krm.kind: ChronicleDataAccessLabel

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
	registry.RegisterModel(krm.ChronicleDataAccessLabelGVK, NewDataAccessLabelModel)
}

func NewDataAccessLabelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDataAccessLabel{config: *config}, nil
}

var _ directbase.Model = &modelDataAccessLabel{}

type modelDataAccessLabel struct {
	config config.ControllerConfig
}

func (m *modelDataAccessLabel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ChronicleDataAccessLabel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ChronicleDataAccessLabelIdentity)

	// Get chronicle GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DataAccessLabelAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDataAccessLabel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//chronicle.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ChronicleDataAccessLabelIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DataAccessLabelAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDataAccessLabel) client(ctx context.Context) (*gcp.DataAccessControlClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataAccessControlRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataAccessLabel client: %w", err)
	}
	return gcpClient, err
}

type DataAccessLabelAdapter struct {
	id        *krm.ChronicleDataAccessLabelIdentity
	gcpClient *gcp.DataAccessControlClient
	desired   *krm.ChronicleDataAccessLabel
	actual    *chroniclepb.DataAccessLabel
}

var _ directbase.Adapter = &DataAccessLabelAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DataAccessLabelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataAccessLabel", "name", a.id)

	req := &chroniclepb.GetDataAccessLabelRequest{
		Name: a.id.String(),
	}
	dataaccesslabelpb, err := a.gcpClient.GetDataAccessLabel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataAccessLabel %q: %w", a.id, err)
	}

	a.actual = dataaccesslabelpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAccessLabelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataAccessLabel", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ChronicleDataAccessLabelSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &chroniclepb.CreateDataAccessLabelRequest{
		Parent:            a.id.ParentString(),
		DataAccessLabelId: a.id.DataAccessLabel,
		DataAccessLabel:   resource,
	}
	created, err := a.gcpClient.CreateDataAccessLabel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataAccessLabel %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DataAccessLabel", "name", a.id)

	status := &krm.ChronicleDataAccessLabelStatus{}
	status.ObservedState = ChronicleDataAccessLabelObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAccessLabelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataAccessLabel", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ChronicleDataAccessLabelSpec_FromProto,
		ChronicleDataAccessLabelSpec_ToProto,
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
		req := &chroniclepb.UpdateDataAccessLabelRequest{
			UpdateMask:      updateMask,
			DataAccessLabel: resource,
		}
		updated, err = a.gcpClient.UpdateDataAccessLabel(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DataAccessLabel %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DataAccessLabel", "name", a.id)
	}

	status := &krm.ChronicleDataAccessLabelStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ChronicleDataAccessLabelObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DataAccessLabelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ChronicleDataAccessLabel{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ChronicleDataAccessLabelSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ChronicleDataAccessLabelGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DataAccessLabelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataAccessLabel", "name", a.id)

	req := &chroniclepb.DeleteDataAccessLabelRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteDataAccessLabel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DataAccessLabel, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DataAccessLabel %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DataAccessLabel", "name", a.id)
	return true, nil
}
