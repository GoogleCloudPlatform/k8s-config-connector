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
// krm.kind: AIPlatformModelDeploymentMonitoringJob

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

	gcp "cloud.google.com/go/aiplatform/apiv1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformModelDeploymentMonitoringJobGVK, NewModelDeploymentMonitoringJobModel)
}

func NewModelDeploymentMonitoringJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelModelDeploymentMonitoringJob{config: *config}, nil
}

var _ directbase.Model = &modelModelDeploymentMonitoringJob{}

type modelModelDeploymentMonitoringJob struct {
	config config.ControllerConfig
}

func (m *modelModelDeploymentMonitoringJob) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformModelDeploymentMonitoringJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AIPlatformModelDeploymentMonitoringJobIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get aiplatform GCP client
	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &ModelDeploymentMonitoringJobAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelModelDeploymentMonitoringJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}
func (m *modelModelDeploymentMonitoringJob) client(ctx context.Context, location string) (*gcp.JobClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewJobClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ModelDeploymentMonitoringJob client: %w", err)
	}
	return gcpClient, err
}

type ModelDeploymentMonitoringJobAdapter struct {
	id        *krm.AIPlatformModelDeploymentMonitoringJobIdentity
	gcpClient *gcp.JobClient
	desired   *krm.AIPlatformModelDeploymentMonitoringJob
	actual    *aiplatformpb.ModelDeploymentMonitoringJob
}

var _ directbase.Adapter = &ModelDeploymentMonitoringJobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ModelDeploymentMonitoringJobAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.ModelDeploymentMonitoringJob == "" {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("getting ModelDeploymentMonitoringJob", "name", a.id)

	req := &aiplatformpb.GetModelDeploymentMonitoringJobRequest{
		Name: a.id.String(),
	}
	modeldeploymentmonitoringjobpb, err := a.gcpClient.GetModelDeploymentMonitoringJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ModelDeploymentMonitoringJob %q: %w", a.id, err)
	}

	a.actual = modeldeploymentmonitoringjobpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ModelDeploymentMonitoringJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ModelDeploymentMonitoringJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformModelDeploymentMonitoringJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if a.id.ModelDeploymentMonitoringJob != "" {
		resource.Name = a.id.String()
	}

	req := &aiplatformpb.CreateModelDeploymentMonitoringJobRequest{
		Parent:                       a.id.ParentString(),
		ModelDeploymentMonitoringJob: resource,
	}
	created, err := a.gcpClient.CreateModelDeploymentMonitoringJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ModelDeploymentMonitoringJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ModelDeploymentMonitoringJob", "name", created.Name)

	if err := a.id.FromExternal(created.Name); err != nil {
		return fmt.Errorf("parsing created ModelDeploymentMonitoringJob name %q: %w", created.Name, err)
	}

	status := &krm.AIPlatformModelDeploymentMonitoringJobStatus{}
	status.ObservedState = AIPlatformModelDeploymentMonitoringJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ModelDeploymentMonitoringJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ModelDeploymentMonitoringJob", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformModelDeploymentMonitoringJobSpec_FromProto,
		AIPlatformModelDeploymentMonitoringJobSpec_ToProto,
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
		req := &aiplatformpb.UpdateModelDeploymentMonitoringJobRequest{
			UpdateMask:                   updateMask,
			ModelDeploymentMonitoringJob: resource,
		}
		op, err := a.gcpClient.UpdateModelDeploymentMonitoringJob(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ModelDeploymentMonitoringJob %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("ModelDeploymentMonitoringJob %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ModelDeploymentMonitoringJob", "name", a.id)
	}

	status := &krm.AIPlatformModelDeploymentMonitoringJobStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIPlatformModelDeploymentMonitoringJobObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ModelDeploymentMonitoringJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformModelDeploymentMonitoringJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformModelDeploymentMonitoringJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AIPlatformModelDeploymentMonitoringJobGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ModelDeploymentMonitoringJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ModelDeploymentMonitoringJob", "name", a.id)

	req := &aiplatformpb.DeleteModelDeploymentMonitoringJobRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteModelDeploymentMonitoringJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ModelDeploymentMonitoringJob, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ModelDeploymentMonitoringJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ModelDeploymentMonitoringJob", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete ModelDeploymentMonitoringJob %s: %w", a.id, err)
	}
	return true, nil
}
