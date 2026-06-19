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

package vertexaimodeldeploymentmonitoringjob

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	directcommon "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.VertexAIModelDeploymentMonitoringJobGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.JobClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewJobClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building JobRESTClient client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	reader := op.Reader
	u := op.GetUnstructured()
	obj := &krm.VertexAIModelDeploymentMonitoringJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := directcommon.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id.(*krm.VertexAIModelDeploymentMonitoringJobIdentity),
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.VertexAIModelDeploymentMonitoringJobIdentity
	gcpClient *gcp.JobClient
	desired   *krm.VertexAIModelDeploymentMonitoringJob
	reader    client.Reader
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAIModelDeploymentMonitoringJob", "name", a.id.String())

	req := &pb.GetModelDeploymentMonitoringJobRequest{
		Name: a.id.String(),
	}
	job, err := a.gcpClient.GetModelDeploymentMonitoringJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAIModelDeploymentMonitoringJob %q: %w", a.id.String(), err)
	}

	mapCtx := &direct.MapContext{}
	observedState := VertexAIModelDeploymentMonitoringJobObservedState_FromProto(mapCtx, job)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("decoding state into %T: %w", observedState, mapCtx.Err())
	}
	a.desired.Status.ObservedState = observedState
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAIModelDeploymentMonitoringJob", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredObj := VertexAIModelDeploymentMonitoringJobSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateModelDeploymentMonitoringJobRequest{
		Parent:                       "projects/" + a.id.Project + "/locations/" + a.id.Location,
		ModelDeploymentMonitoringJob: desiredObj,
	}

	op, err := a.gcpClient.CreateModelDeploymentMonitoringJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAIModelDeploymentMonitoringJob %s: %w", a.id.String(), err)
	}
	job := op
	_ = ctx
	if err != nil {
		return fmt.Errorf("waiting for VertexAIModelDeploymentMonitoringJob %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created VertexAIModelDeploymentMonitoringJob", "name", a.id.String())

	status := VertexAIModelDeploymentMonitoringJobObservedState_FromProto(mapCtx, job)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	a.desired.Status.ObservedState = status
	a.desired.Status.ExternalRef = direct.LazyPtr(job.Name)
	return nil
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAIModelDeploymentMonitoringJob", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredObj := VertexAIModelDeploymentMonitoringJobSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredObj.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{
		Paths: []string{
			"display_name", "model_deployment_monitoring_objective_configs",
			"model_deployment_monitoring_schedule_config", "logging_sampling_strategy",
			"model_monitoring_alert_config", "predict_instance_schema_uri",
			"sample_predict_instance", "analysis_instance_schema_uri",
			"log_ttl", "labels", "stats_anomalies_base_directory",
			"encryption_spec", "enable_monitoring_pipeline_logs",
		},
	}

	req := &pb.UpdateModelDeploymentMonitoringJobRequest{
		ModelDeploymentMonitoringJob: desiredObj,
		UpdateMask:                   updateMask,
	}

	op, err := a.gcpClient.UpdateModelDeploymentMonitoringJob(ctx, req)
	if err != nil {
		return fmt.Errorf("updating VertexAIModelDeploymentMonitoringJob %s: %w", a.id.String(), err)
	}
	job, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update: %w", err)
	}
	_ = ctx
	if err != nil {
		return fmt.Errorf("waiting for VertexAIModelDeploymentMonitoringJob %s update: %w", a.id.String(), err)
	}

	status := VertexAIModelDeploymentMonitoringJobObservedState_FromProto(mapCtx, job)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	a.desired.Status.ObservedState = status
	return nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.desired == nil {
		return nil, fmt.Errorf("adapter has no desired state")
	}
	var err error
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(a.desired)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: obj}, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAIModelDeploymentMonitoringJob", "name", a.id.String())

	req := &pb.DeleteModelDeploymentMonitoringJobRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteModelDeploymentMonitoringJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting VertexAIModelDeploymentMonitoringJob %s: %w", a.id.String(), err)
	}
	// op.Wait(ctx) is not needed since Delete returns LRO and we don't block unless we want to wait.
	// Actually, DeleteModelDeploymentMonitoringJob might return an LRO.
	if err = op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for VertexAIModelDeploymentMonitoringJob %s deletion: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted VertexAIModelDeploymentMonitoringJob", "name", a.id.String())
	return true, nil
}
