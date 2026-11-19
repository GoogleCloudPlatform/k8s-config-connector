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
// krm.kind: AIPlatformBatchPredictionJob

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
	registry.RegisterModel(krm.AIPlatformBatchPredictionJobGVK, NewBatchPredictionJobModel)
}

func NewBatchPredictionJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBatchPredictionJob{config: *config}, nil
}

var _ directbase.Model = &modelBatchPredictionJob{}

type modelBatchPredictionJob struct {
	config config.ControllerConfig
}

func (m *modelBatchPredictionJob) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformBatchPredictionJob{}
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
	id := i.(*krm.AIPlatformBatchPredictionJobIdentity)

	// Get aiplatform GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &BatchPredictionJobAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBatchPredictionJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AIPlatformBatchPredictionJobIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &BatchPredictionJobAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelBatchPredictionJob) client(ctx context.Context, location string) (*gcp.JobClient, error) {
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
		return nil, fmt.Errorf("building BatchPredictionJob client: %w", err)
	}
	return gcpClient, err
}

type BatchPredictionJobAdapter struct {
	id        *krm.AIPlatformBatchPredictionJobIdentity
	gcpClient *gcp.JobClient
	desired   *krm.AIPlatformBatchPredictionJob
	actual    *aiplatformpb.BatchPredictionJob
}

var _ directbase.Adapter = &BatchPredictionJobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BatchPredictionJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BatchPredictionJob", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &aiplatformpb.GetBatchPredictionJobRequest{
		Name: a.id.String(),
	}
	batchpredictionjobpb, err := a.gcpClient.GetBatchPredictionJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BatchPredictionJob %q: %w", a.id, err)
	}

	a.actual = batchpredictionjobpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BatchPredictionJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BatchPredictionJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformBatchPredictionJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &aiplatformpb.CreateBatchPredictionJobRequest{
		Parent:             a.id.ParentString(),
		BatchPredictionJob: resource,
	}
	created, err := a.gcpClient.CreateBatchPredictionJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BatchPredictionJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BatchPredictionJob", "name", a.id)

	status := &krm.AIPlatformBatchPredictionJobStatus{}
	status.ObservedState = AIPlatformBatchPredictionJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.AIPlatformBatchPredictionJobIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.BatchPredictionJob = parsedID.BatchPredictionJob
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BatchPredictionJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BatchPredictionJob", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformBatchPredictionJobSpec_FromProto,
		AIPlatformBatchPredictionJobSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating AIPlatformBatchPredictionJob is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.AIPlatformBatchPredictionJobStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIPlatformBatchPredictionJobObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BatchPredictionJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformBatchPredictionJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformBatchPredictionJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AIPlatformBatchPredictionJobGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BatchPredictionJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BatchPredictionJob", "name", a.id)

	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	// 1. Get the latest BatchPredictionJob state before deletion
	getReq := &aiplatformpb.GetBatchPredictionJobRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetBatchPredictionJob(ctx, getReq)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("getting BatchPredictionJob before deletion %q: %w", a.id.String(), err)
	}

	// 2. If the state is not terminal, initiate cancellation and wait
	isTerminal := func(state aiplatformpb.JobState) bool {
		return state == aiplatformpb.JobState_JOB_STATE_SUCCEEDED ||
			state == aiplatformpb.JobState_JOB_STATE_FAILED ||
			state == aiplatformpb.JobState_JOB_STATE_CANCELLED
	}

	if !isTerminal(actual.State) {
		if actual.State != aiplatformpb.JobState_JOB_STATE_CANCELLING {
			log.V(2).Info("cancelling BatchPredictionJob before deletion", "name", a.id.String(), "currentState", actual.State)
			cancelReq := &aiplatformpb.CancelBatchPredictionJobRequest{
				Name: a.id.String(),
			}
			if err := a.gcpClient.CancelBatchPredictionJob(ctx, cancelReq); err != nil {
				log.Error(err, "error cancelling BatchPredictionJob", "name", a.id.String())
				return false, fmt.Errorf("cancelling BatchPredictionJob %q: %w", a.id.String(), err)
			}
		}
		return false, fmt.Errorf("BatchPredictionJob %q is in state %s; waiting for terminal state before deletion", a.id.String(), actual.State)
	}

	req := &aiplatformpb.DeleteBatchPredictionJobRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteBatchPredictionJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BatchPredictionJob, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BatchPredictionJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BatchPredictionJob", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete BatchPredictionJob %s: %w", a.id, err)
	}
	return true, nil
}
