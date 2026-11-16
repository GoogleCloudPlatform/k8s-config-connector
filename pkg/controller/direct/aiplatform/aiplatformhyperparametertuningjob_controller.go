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
// krm.kind: AIPlatformHyperparameterTuningJob

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
	registry.RegisterModel(krm.AIPlatformHyperparameterTuningJobGVK, NewHyperparameterTuningJobModel)
}

func NewHyperparameterTuningJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelHyperparameterTuningJob{config: *config}, nil
}

var _ directbase.Model = &modelHyperparameterTuningJob{}

type modelHyperparameterTuningJob struct {
	config config.ControllerConfig
}

func (m *modelHyperparameterTuningJob) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformHyperparameterTuningJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AIPlatformHyperparameterTuningJobIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get aiplatform GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &HyperparameterTuningJobAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelHyperparameterTuningJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AIPlatformHyperparameterTuningJobIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &HyperparameterTuningJobAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelHyperparameterTuningJob) client(ctx context.Context, location string) (*gcp.JobClient, error) {
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
		return nil, fmt.Errorf("building HyperparameterTuningJob client: %w", err)
	}
	return gcpClient, err
}

type HyperparameterTuningJobAdapter struct {
	id        *krm.AIPlatformHyperparameterTuningJobIdentity
	gcpClient *gcp.JobClient
	desired   *krm.AIPlatformHyperparameterTuningJob
	actual    *aiplatformpb.HyperparameterTuningJob
}

var _ directbase.Adapter = &HyperparameterTuningJobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *HyperparameterTuningJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting HyperparameterTuningJob", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &aiplatformpb.GetHyperparameterTuningJobRequest{
		Name: a.id.String(),
	}
	hyperparametertuningjobpb, err := a.gcpClient.GetHyperparameterTuningJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting HyperparameterTuningJob %q: %w", a.id, err)
	}

	a.actual = hyperparametertuningjobpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *HyperparameterTuningJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating HyperparameterTuningJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformHyperparameterTuningJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &aiplatformpb.CreateHyperparameterTuningJobRequest{
		Parent:                  a.id.ParentString(),
		HyperparameterTuningJob: resource,
	}
	created, err := a.gcpClient.CreateHyperparameterTuningJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating HyperparameterTuningJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created HyperparameterTuningJob", "name", a.id)

	status := &krm.AIPlatformHyperparameterTuningJobStatus{}
	status.ObservedState = AIPlatformHyperparameterTuningJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.AIPlatformHyperparameterTuningJobIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.HyperparameterTuningJob = parsedID.HyperparameterTuningJob
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *HyperparameterTuningJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating HyperparameterTuningJob", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformHyperparameterTuningJobSpec_FromProto,
		AIPlatformHyperparameterTuningJobSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating AIPlatformHyperparameterTuningJob is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.AIPlatformHyperparameterTuningJobStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIPlatformHyperparameterTuningJobObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *HyperparameterTuningJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformHyperparameterTuningJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformHyperparameterTuningJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AIPlatformHyperparameterTuningJobGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *HyperparameterTuningJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting HyperparameterTuningJob", "name", a.id)

	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	// 1. Get the latest HyperparameterTuningJob state before deletion
	getReq := &aiplatformpb.GetHyperparameterTuningJobRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetHyperparameterTuningJob(ctx, getReq)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("getting HyperparameterTuningJob before deletion %q: %w", a.id.String(), err)
	}

	// 2. If the state is not terminal, initiate cancellation and wait
	isTerminal := func(state aiplatformpb.JobState) bool {
		return state == aiplatformpb.JobState_JOB_STATE_SUCCEEDED ||
			state == aiplatformpb.JobState_JOB_STATE_FAILED ||
			state == aiplatformpb.JobState_JOB_STATE_CANCELLED
	}

	if !isTerminal(actual.State) {
		if actual.State != aiplatformpb.JobState_JOB_STATE_CANCELLING {
			log.V(2).Info("cancelling AIPlatformHyperparameterTuningJob before deletion", "name", a.id.String(), "currentState", actual.State)
			cancelReq := &aiplatformpb.CancelHyperparameterTuningJobRequest{
				Name: a.id.String(),
			}
			if err := a.gcpClient.CancelHyperparameterTuningJob(ctx, cancelReq); err != nil {
				log.Error(err, "error cancelling AIPlatformHyperparameterTuningJob", "name", a.id.String())
				return false, fmt.Errorf("cancelling AIPlatformHyperparameterTuningJob %q: %w", a.id.String(), err)
			}
		}
		return false, fmt.Errorf("AIPlatformHyperparameterTuningJob %q is in state %s; waiting for terminal state before deletion", a.id.String(), actual.State)
	}

	// 3. Delete the HyperparameterTuningJob once in a terminal state
	req := &aiplatformpb.DeleteHyperparameterTuningJobRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteHyperparameterTuningJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent HyperparameterTuningJob, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting HyperparameterTuningJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted HyperparameterTuningJob", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete HyperparameterTuningJob %s: %w", a.id, err)
	}
	return true, nil
}
