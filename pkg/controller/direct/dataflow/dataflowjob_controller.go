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

package dataflow

import (
	"context"
	"fmt"
	"strings"
	"time"

	api "cloud.google.com/go/dataflow/apiv1beta3"
	pb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataflow/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DataflowJobGVK, newDataflowJobModel)
}

func newDataflowJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataFlowJobModel{config: config}, nil
}

type dataFlowJobModel struct {
	config *config.ControllerConfig
}

// dataFlowJobModel implements the Model interface.
var _ directbase.Model = &dataFlowJobModel{}

type dataflowJobAdapter struct {
	projectID  string
	location   string
	resourceID string
	jobID      string

	desiredObj *krm.DataflowJob
	actual     *pb.Job

	templatesClient *api.TemplatesClient
	jobsClient      *api.JobsV1Beta3Client
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &dataflowJobAdapter{}

// AdapterForObject implements the Model interface.
func (m *dataFlowJobModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	templatesClient, err := gcpClient.newTemplatesClient(ctx)
	if err != nil {
		return nil, err
	}
	jobsClient, err := gcpClient.newJobsClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.DataflowJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := direct.ValueOf(obj.Spec.Region)
	if location == "" {
		zone := direct.ValueOf(obj.Spec.Zone)
		if zone != "" {
			lastDash := strings.LastIndex(zone, "-")
			if lastDash != -1 {
				location = zone[:lastDash]
			} else {
				location = zone
			}
		}
	}
	if location == "" {
		return nil, fmt.Errorf("cannot resolve region")
	}

	projectRef, err := refs.ResolveProjectFromAnnotation(ctx, kube, obj)
	if err != nil {
		return nil, err
	}
	if projectRef == nil || projectRef.ProjectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	projectID := projectRef.ProjectID

	jobID := direct.ValueOf(obj.Status.JobId)

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if err := common.VisitFields(obj, &refNormalizer{ctx: ctx, src: obj, project: *projectRef, kube: kube}); err != nil {
		return nil, err
	}

	return &dataflowJobAdapter{
		projectID:       projectID,
		location:        location,
		resourceID:      resourceID,
		jobID:           jobID,
		desiredObj:      obj,
		templatesClient: templatesClient,
		jobsClient:      jobsClient,
	}, nil
}

func (m *dataFlowJobModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *dataflowJobAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	jobID := a.jobID
	if jobID == "" {
		return false, nil
	}

	jobFQN := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", a.projectID, a.location, jobID)

	req := &pb.GetJobRequest{
		JobId:     jobID,
		ProjectId: a.projectID,
		Location:  a.location,
		View:      pb.JobView_JOB_VIEW_SUMMARY,
	}
	job, err := a.jobsClient.GetJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataflow job %q: %w", jobFQN, err)
	}

	a.actual = job
	return true, nil
}

// Delete implements the Adapter interface.
func (a *dataflowJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	// Already deleted
	if a.actual == nil {
		return false, nil
	}

	jobID := a.actual.Id

	if isTerminalState(a.actual.CurrentState) {
		return true, nil
	}

	// To terminate a dataflow job, we update the job with a requested terminal state.
	updateJob := &pb.Job{
		RequestedState: pb.JobState_JOB_STATE_CANCELLED,
	}
	req := &pb.UpdateJobRequest{
		ProjectId: a.projectID,
		JobId:     jobID,
		Job:       updateJob,
		Location:  a.location,
	}

	jobFQN := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", a.projectID, a.location, jobID)

	updatedJob, err := a.jobsClient.UpdateJob(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting dataflow job %s: %w", jobFQN, err)
	}

	deleted := false
	for !deleted {
		time.Sleep(2 * time.Second)
		latest, err := a.getJob(ctx, jobID)
		if err != nil {
			return false, fmt.Errorf("getting state of job: %w", err)
		}
		if latest == nil {
			deleted = true
			break
		}
		if isTerminalState(latest.CurrentState) {
			deleted = true
		} else {
			log.Info("unexpected job state waiting for job cancellation", "state", latest.CurrentState)
		}
	}

	log.Info("updated job", "updated", updatedJob)
	return true, nil
}

func (a *dataflowJobAdapter) getJob(ctx context.Context, jobID string) (*pb.Job, error) {
	req := &pb.GetJobRequest{
		ProjectId: a.projectID,
		JobId:     jobID,
		View:      pb.JobView_JOB_VIEW_SUMMARY,
		Location:  a.location,
	}

	jobFQN := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", a.projectID, a.location, jobID)

	job, err := a.jobsClient.GetJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("getting dataflow job %s: %w", jobFQN, err)
	}

	return job, nil
}

func (a *dataflowJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

// Create implements the Adapter interface.
func (a *dataflowJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	mapCtx := &direct.MapContext{}
	env := toRuntimeEnvironment(mapCtx, &a.desiredObj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parameters, err := toMapStringString("spec.parameters", a.desiredObj.Spec.Parameters)
	if err != nil {
		return err
	}

	req := &pb.CreateJobFromTemplateRequest{
		ProjectId: a.projectID,
		Location:  a.location,
		JobName:   a.resourceID,
		Template: &pb.CreateJobFromTemplateRequest_GcsPath{
			GcsPath: a.desiredObj.Spec.TemplateGcsPath,
		},
		Parameters:  parameters,
		Environment: env,
	}
	log.V(0).Info("making dataflow CreateJobFromTemplate call", "request", req)

	job, err := a.templatesClient.CreateJobFromTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataflow job from template: %w", err)
	}

	if err := a.updateStatus(ctx, createOp, job); err != nil {
		return err
	}

	return nil
}

func isTerminalState(state pb.JobState) bool {
	switch state {
	case pb.JobState_JOB_STATE_DONE,
		pb.JobState_JOB_STATE_FAILED,
		pb.JobState_JOB_STATE_CANCELLED,
		pb.JobState_JOB_STATE_DRAINED,
		pb.JobState_JOB_STATE_UPDATED:
		return true
	default:
		return false
	}
}

func (a *dataflowJobAdapter) updateStatus(ctx context.Context, op directbase.Operation, job *pb.Job) error {
	status := &krm.DataflowJobStatus{
		JobId: direct.PtrTo(job.GetId()),
		State: direct.PtrTo(job.CurrentState.String()),
		Type:  direct.PtrTo(job.Type.String()),
	}

	var readyCondition *v1alpha1.Condition

	switch job.CurrentState {
	case pb.JobState_JOB_STATE_RUNNING:
		readyCondition = &v1alpha1.Condition{
			Type:    v1alpha1.ReadyConditionType,
			Status:  v1.ConditionTrue,
			Reason:  k8s.UpToDate,
			Message: "The resource is up to date",
		}

	case pb.JobState_JOB_STATE_FAILED, pb.JobState_JOB_STATE_CANCELLED, pb.JobState_JOB_STATE_DONE, pb.JobState_JOB_STATE_DRAINED:
		readyCondition = &v1alpha1.Condition{
			Type:    v1alpha1.ReadyConditionType,
			Status:  v1.ConditionFalse,
			Reason:  k8s.UpToDate, // Because we are stopping reconciliation
			Message: fmt.Sprintf("Dataflow job has reached terminal state '%v'", job.CurrentState),
		}

	default:
		readyCondition = &v1alpha1.Condition{
			Type:    v1alpha1.ReadyConditionType,
			Status:  v1.ConditionFalse,
			Reason:  k8s.Updating,
			Message: fmt.Sprintf("Waiting for Dataflow job to be running (state is %v)", job.CurrentState),
		}
		op.RequestRequeue()
	}

	if err := op.UpdateStatus(ctx, status, readyCondition); err != nil {
		return fmt.Errorf("updating status: %w", err)
	}

	return nil
}

// Update implements the Adapter interface.
func (a *dataflowJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	observedGeneration, _, err := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
	if err != nil {
		return fmt.Errorf("reading status.observedGeneration: %w", err)
	}
	metadataGeneration := u.GetGeneration()

	if observedGeneration == metadataGeneration {
		log.V(0).Info("object status.observedGeneration matches metadata.generation, skipping reconcile", "status.observedGeneration", observedGeneration, "metadata.generation", metadataGeneration)

		if a.actual != nil {
			if err := a.updateStatus(ctx, updateOp, a.actual); err != nil {
				return err
			}
		}
		return nil
	}

	if a.actual != nil && a.actual.Type == pb.JobType_JOB_TYPE_BATCH {
		return fmt.Errorf("Batch jobs cannot be updated.")
	}

	mapCtx := &direct.MapContext{}
	env := toRuntimeEnvironment(mapCtx, &a.desiredObj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parameters, err := toMapStringString("spec.parameters", a.desiredObj.Spec.Parameters)
	if err != nil {
		return err
	}

	transformNameMapping, err := toMapStringString("spec.transformNameMapping", a.desiredObj.Spec.TransformNameMapping)
	if err != nil {
		return err
	}

	launchParameters := &pb.LaunchTemplateParameters{
		JobName:              a.resourceID,
		Parameters:           parameters,
		TransformNameMapping: transformNameMapping,
		Environment:          env,
		Update:               true,
	}

	req := &pb.LaunchTemplateRequest{
		ProjectId: a.projectID,
		Location:  a.location,
		Template: &pb.LaunchTemplateRequest_GcsPath{
			GcsPath: a.desiredObj.Spec.TemplateGcsPath,
		},
		LaunchParameters: launchParameters,
	}
	log.V(0).Info("making dataflow LaunchTemplate call", "request", req)

	response, err := a.templatesClient.LaunchTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("updating dataflow template: %w", err)
	}

	job := response.GetJob()

	if err := a.updateStatus(ctx, updateOp, job); err != nil {
		return err
	}

	return nil
}

func toRuntimeEnvironment(mapCtx *direct.MapContext, in *krm.DataflowJobSpec) *pb.RuntimeEnvironment {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeEnvironment{}

	if in.MaxWorkers != nil {
		out.MaxWorkers = int32(*in.MaxWorkers)
	}
	if in.Zone != nil {
		out.Zone = *in.Zone
	}
	if in.ServiceAccountRef != nil {
		out.ServiceAccountEmail = in.ServiceAccountRef.External
	}
	out.TempLocation = in.TempGcsLocation
	if in.MachineType != nil {
		out.MachineType = *in.MachineType
	}
	if len(in.AdditionalExperiments) > 0 {
		out.AdditionalExperiments = in.AdditionalExperiments
	}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	if in.IpConfiguration != nil {
		out.IpConfiguration = direct.Enum_ToProto[pb.WorkerIPAddressConfiguration](mapCtx, in.IpConfiguration)
	}
	if in.EnableStreamingEngine != nil {
		out.EnableStreamingEngine = *in.EnableStreamingEngine
	}

	return out
}
