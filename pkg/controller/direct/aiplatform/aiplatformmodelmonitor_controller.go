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
// krm.group: aiplatform.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: AIPlatformModelMonitor

package aiplatform

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformModelMonitorGVK, NewModelMonitorModel)
}

func NewModelMonitorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelModelMonitor{config: *config}, nil
}

var _ directbase.Model = &modelModelMonitor{}

type modelModelMonitor struct {
	config config.ControllerConfig
}

func (m *modelModelMonitor) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformModelMonitor{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AIPlatformModelMonitorIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get aiplatform GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &ModelMonitorAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelModelMonitor) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AIPlatformModelMonitorIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &ModelMonitorAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelModelMonitor) client(ctx context.Context, location string) (*gcp.ModelMonitoringClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewModelMonitoringRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ModelMonitor client: %w", err)
	}
	return gcpClient, err
}

type ModelMonitorAdapter struct {
	id        *krm.AIPlatformModelMonitorIdentity
	gcpClient *gcp.ModelMonitoringClient
	desired   *krm.AIPlatformModelMonitor
	actual    *aiplatformpb.ModelMonitor
}

var _ directbase.Adapter = &ModelMonitorAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ModelMonitorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ModelMonitor", "name", a.id)

	req := &aiplatformpb.GetModelMonitorRequest{
		Name: a.id.String(),
	}
	modelmonitorpb, err := a.gcpClient.GetModelMonitor(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ModelMonitor %q: %w", a.id, err)
	}

	a.actual = modelmonitorpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ModelMonitorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ModelMonitor", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformModelMonitorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &aiplatformpb.CreateModelMonitorRequest{
		Parent:         a.id.ParentString(),
		ModelMonitorId: a.id.ModelMonitor,
		ModelMonitor:   resource,
	}
	op, err := a.gcpClient.CreateModelMonitor(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ModelMonitor %s: %w", a.id, err)
	}
	locPath := fmt.Sprintf("/v1beta1/%s/operations/%s", a.id.ParentString(), op.Name()[strings.LastIndex(op.Name(), "/")+1:])
	created, err := op.Wait(ctx, gax.WithPath(locPath))
	if err != nil {
		return fmt.Errorf("ModelMonitor %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ModelMonitor", "name", a.id)

	status := &krm.AIPlatformModelMonitorStatus{}
	status.ObservedState = AIPlatformModelMonitorObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ModelMonitorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ModelMonitor", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformModelMonitorSpec_FromProto,
		AIPlatformModelMonitorSpec_ToProto,
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
		req := &aiplatformpb.UpdateModelMonitorRequest{
			UpdateMask:   updateMask,
			ModelMonitor: resource,
		}
		op, err := a.gcpClient.UpdateModelMonitor(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ModelMonitor %s: %w", a.id, err)
		}
		locPath := fmt.Sprintf("/v1beta1/%s/operations/%s", a.id.ParentString(), op.Name()[strings.LastIndex(op.Name(), "/")+1:])
		updated, err = op.Wait(ctx, gax.WithPath(locPath))
		if err != nil {
			return fmt.Errorf("ModelMonitor %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ModelMonitor", "name", a.id)
	}

	status := &krm.AIPlatformModelMonitorStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIPlatformModelMonitorObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ModelMonitorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformModelMonitor{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformModelMonitorSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AIPlatformModelMonitorGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ModelMonitorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ModelMonitor", "name", a.id)

	req := &aiplatformpb.DeleteModelMonitorRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteModelMonitor(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ModelMonitor, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ModelMonitor %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ModelMonitor", "name", a.id)
	locPath := fmt.Sprintf("/v1beta1/%s/operations/%s", a.id.ParentString(), op.Name()[strings.LastIndex(op.Name(), "/")+1:])
	err = op.Wait(ctx, gax.WithPath(locPath))
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete ModelMonitor %s: %w", a.id, err)
	}
	return true, nil
}
