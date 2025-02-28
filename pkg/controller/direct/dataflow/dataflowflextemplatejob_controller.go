// Copyright 2024 Google LLC
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
	"encoding/json"
	"fmt"
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
	"google.golang.org/protobuf/proto"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataflowFlexTemplateJobGVK, newDataflowFlexTemplateJobModel)
}

func newDataflowFlexTemplateJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataFlowFlexTemplateJobModel{config: config}, nil
}

type dataFlowFlexTemplateJobModel struct {
	config *config.ControllerConfig
}

// dataFlowFlexTemplateJobModel implements the Model interface.
var _ directbase.Model = &dataFlowFlexTemplateJobModel{}

type dataflowFlexTemplateJobAdapter struct {
	projectID  string
	location   string
	resourceID string
	jobID      string

	desiredObj *krm.DataflowFlexTemplateJob
	desired    *pb.LaunchFlexTemplateParameter
	actual     *pb.Job

	flexTemplatesClient *api.FlexTemplatesClient
	jobsClient          *api.JobsV1Beta3Client
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &dataflowFlexTemplateJobAdapter{}

// AdapterForObject implements the Model interface.
func (m *dataFlowFlexTemplateJobModel) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	flexTemplatesClient, err := gcpClient.newFlexTemplatesClient(ctx)
	if err != nil {
		return nil, err
	}
	jobsClient, err := gcpClient.newJobsClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.DataflowFlexTemplateJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// TODO: Why don't we support resourceID?
	resourceID := "" // direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := direct.ValueOf(obj.Spec.Region)
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

	jobID := obj.Status.JobID

	// TODO: Move from monitoring package into shared package (and make refs implement an interface)
	if err := common.VisitFields(obj, &refNormalizer{ctx: ctx, src: obj, project: *projectRef, kube: kube}); err != nil {
		return nil, err
	}

	desired, err := toLaunchParameter(ctx, resourceID, obj)
	if err != nil {
		return nil, err
	}

	return &dataflowFlexTemplateJobAdapter{
		projectID:           projectID,
		location:            location,
		resourceID:          resourceID,
		jobID:               jobID,
		desired:             desired,
		desiredObj:          obj,
		flexTemplatesClient: flexTemplatesClient,
		jobsClient:          jobsClient,
	}, nil
}

// To preserve backwards compatibility, as we previously supported map[string]any, rather than map[string]string.
// For example: `foo: true` as well as `foo: "true"`
func toMapStringString(fieldPath string, ext *runtime.RawExtension) (map[string]string, error) {
	if ext == nil || ext.Raw == nil {
		return nil, nil
	}
	m := make(map[string]any)
	if err := json.Unmarshal(ext.Raw, &m); err != nil {
		return nil, fmt.Errorf("error parsing %v field: %w", fieldPath, err)
	}

	out := make(map[string]string)
	for k, v := range m {
		switch v := v.(type) {
		case string:
			out[k] = v
		default:
			vString := fmt.Sprintf("%v", v)
			out[k] = vString
		}
	}

	return out, nil
}

func toLaunchParameter(ctx context.Context, resourceID string, obj *krm.DataflowFlexTemplateJob) (*pb.LaunchFlexTemplateParameter, error) {
	mapCtx := &direct.MapContext{}
	environment := DataflowFlexTemplateJobSpec_ToProto(mapCtx, &obj.Spec)

	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	launchParameter := &pb.LaunchFlexTemplateParameter{
		JobName:     resourceID,
		Environment: environment,
	}

	var err error
	launchParameter.Parameters, err = toMapStringString("spec.parameters", obj.Spec.Parameters)
	if err != nil {
		return nil, err
	}

	// Only applies to update, but it's easier to build this now
	launchParameter.TransformNameMappings, err = toMapStringString("spec.transformNameMapping", obj.Spec.TransformNameMapping)
	if err != nil {
		return nil, err
	}

	if obj.Spec.ContainerSpecGcsPath != nil {
		launchParameter.Template = &pb.LaunchFlexTemplateParameter_ContainerSpecGcsPath{
			ContainerSpecGcsPath: *obj.Spec.ContainerSpecGcsPath,
		}
	}

	return launchParameter, nil
}

func (m *dataFlowFlexTemplateJobModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *dataflowFlexTemplateJobAdapter) Find(ctx context.Context) (bool, error) {
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
func (a *dataflowFlexTemplateJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	jobID := a.actual.Id

	// To terminate a dataflow job, we update the job with a requested
	// terminal state.
	updateJob := &pb.Job{
		RequestedState: pb.JobState_JOB_STATE_CANCELLED,
	}
	// TODO: Delete via status selfLink?
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
			// TODO: not right!
			return false, fmt.Errorf("getting state of job")
		}
		switch latest.CurrentState {
		case pb.JobState_JOB_STATE_CANCELLED:
			deleted = true
		default:
			log.Info("unexpected job state waiting for job cancellation", "state", latest.CurrentState)
		}
	}

	// TODO: Poll for status
	log.Info("updated job", "updated", updatedJob)
	return true, nil
}

func (a *dataflowFlexTemplateJobAdapter) getJob(ctx context.Context, jobID string) (*pb.Job, error) {
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
		return nil, fmt.Errorf("getting dataFlow job %s: %w", jobFQN, err)
	}

	return job, nil
}

func (a *dataflowFlexTemplateJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataflowFlexTemplateJob{}
	mapCtx := &direct.MapContext{}
	actualFlexTemplateParameter, err := toLaunchParameter(ctx, a.resourceID, a.desiredObj)
	if err != nil {
		return nil, err
	}

	obj.Spec = direct.ValueOf(DataflowFlexTemplateJobSpec_FromProto(mapCtx, actualFlexTemplateParameter.Environment))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Region = direct.PtrTo(a.location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.SetName(a.actual.Id)
	u.SetGroupVersionKind(krm.DataflowFlexTemplateJobGVK)
	u.SetLabels(a.actual.Labels)

	u.Object = uObj
	return u, nil
}

// Create implements the Adapter interface.
func (a *dataflowFlexTemplateJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	launchParameter := proto.Clone(a.desired).(*pb.LaunchFlexTemplateParameter)
	launchParameter.Update = false
	launchParameter.TransformNameMappings = nil

	req := &pb.LaunchFlexTemplateRequest{
		ProjectId:       a.projectID,
		Location:        a.location,
		LaunchParameter: launchParameter,
	}
	log.V(0).Info("making dataflow LaunchFlexTemplate call", "request", req)

	response, err := a.flexTemplatesClient.LaunchFlexTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataflow flexTemplate: %w", err)
	}

	job := response.GetJob()

	if err := a.updateStatus(ctx, createOp, job); err != nil {
		return err
	}

	return nil
}

func (a *dataflowFlexTemplateJobAdapter) updateStatus(ctx context.Context, op directbase.Operation, job *pb.Job) error {
	status := &krm.DataflowFlexTemplateJobStatus{
		JobID:        job.GetId(),
		CurrentState: direct.PtrTo(job.CurrentState.String()),
		Type:         direct.PtrTo(job.Type.String()),
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
		// We are healthy - no requeue needed

	case pb.JobState_JOB_STATE_FAILED, pb.JobState_JOB_STATE_CANCELLED, pb.JobState_JOB_STATE_DONE, pb.JobState_JOB_STATE_DRAINED:
		readyCondition = &v1alpha1.Condition{
			Type:    v1alpha1.ReadyConditionType,
			Status:  v1.ConditionFalse,
			Reason:  k8s.UpToDate, // Because we are stopping reconciliation
			Message: fmt.Sprintf("Dataflow job has reached terminal state '%v'", job.CurrentState),
		}
		// State is terminal, no need to requeue

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
func (a *dataflowFlexTemplateJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	observedGeneration, _, err := unstructured.NestedInt64(u.Object, "status", "observedGeneration")
	if err != nil {
		return fmt.Errorf("reading status.observedGeneration: %w", err)
	}
	metadataGeneration := u.GetGeneration()

	// We rely on the FlexJob being immutable, so drift at the GCP level should not be possible.
	// Instead, we reconcile whenever we see a different metadata.generation
	if observedGeneration == metadataGeneration {
		log.V(0).Info("object status.observedGeneration matches metadata.generation, skipping reconcile", "status.observedGeneration", observedGeneration, "metadata.generation", metadataGeneration)

		j, _ := u.MarshalJSON()
		log.V(0).Info("object is", "json", string(j))

		// If we are waiting on the existing job, update the status
		if a.actual != nil {
			if err := a.updateStatus(ctx, updateOp, a.actual); err != nil {
				return err
			}

			// TODO: If job fails to update, we probably need to "put back" the old job id
			// But we also want to avoid repeatedly launching a failing job...
			// The current behaviour will only try to launch once, which seems reasonable.
		}

		return nil
	}

	launchParameter := proto.Clone(a.desired).(*pb.LaunchFlexTemplateParameter)
	launchParameter.Update = true

	req := &pb.LaunchFlexTemplateRequest{
		ProjectId:       a.projectID,
		Location:        a.location,
		LaunchParameter: launchParameter,
	}
	log.V(0).Info("making dataflow LaunchFlexTemplate call", "request", req)

	response, err := a.flexTemplatesClient.LaunchFlexTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("updating dataflow flexTemplate: %w", err)
	}

	job := response.GetJob()

	if err := a.updateStatus(ctx, updateOp, job); err != nil {
		return err
	}

	return nil
}
