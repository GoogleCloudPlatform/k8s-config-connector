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

package datalabelingjob

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAIDataLabelingJobGVK, NewDataLabelingJobModel)
}

func NewDataLabelingJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDataLabelingJob{config: *config}, nil
}

var _ directbase.Model = &modelDataLabelingJob{}

type modelDataLabelingJob struct {
	config config.ControllerConfig
}

func (m *modelDataLabelingJob) client(ctx context.Context, location string) (*gcp.JobClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewJobClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataLabelingJob client: %w", err)
	}
	return gcpClient, err
}

func (m *modelDataLabelingJob) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAIDataLabelingJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewVertexAIDataLabelingJobIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &DataLabelingJobAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDataLabelingJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DataLabelingJobAdapter struct {
	id        *krm.VertexAIDataLabelingJobIdentity
	gcpClient *gcp.JobClient
	desired   *krm.VertexAIDataLabelingJob
	actual    *pb.DataLabelingJob
}

var _ directbase.Adapter = &DataLabelingJobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DataLabelingJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataLabelingJob", "name", a.id)

	req := &pb.GetDataLabelingJobRequest{Name: a.id.String()}
	datalabelingjobpb, err := a.gcpClient.GetDataLabelingJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataLabelingJob %q: %w", a.id, err)
	}

	a.actual = datalabelingjobpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataLabelingJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataLabelingJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VertexAIDataLabelingJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Parent format: projects/{project}/locations/{location}
	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateDataLabelingJobRequest{
		Parent:          parent,
		DataLabelingJob: resource,
	}
	op, err := a.gcpClient.CreateDataLabelingJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataLabelingJob %s: %w", a.id, err)
	}

	created := op
	a.actual = created

	status := &krm.VertexAIDataLabelingJobStatus{}
	status.ObservedState = VertexAIDataLabelingJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := created.GetName()
	status.ExternalRef = &externalRef

	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataLabelingJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataLabelingJob", "name", a.id)

	desired := a.desired.DeepCopy()
	mapCtx := &direct.MapContext{}
	resource := VertexAIDataLabelingJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Compare proto representations to check for drift / unsupported updates
	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) != 0 {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)
		log.V(2).Info("VertexAIDataLabelingJob does not support updates", "name", a.id.String())
		return nil
	}

	status := &krm.VertexAIDataLabelingJobStatus{}
	status.ObservedState = VertexAIDataLabelingJobObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return updateOp.UpdateStatus(ctx, status, nil)
}

// Delete deletes the resource from GCP.
func (a *DataLabelingJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataLabelingJob", "name", a.id)

	req := &pb.DeleteDataLabelingJobRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteDataLabelingJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting DataLabelingJob %s: %w", a.id.String(), err)
	}

	// DeleteDataLabelingJob returns an Operation, wait for it to complete.
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for DataLabelingJob deletion %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted DataLabelingJob", "name", a.id)

	return true, nil
}

// Export fetches the cloud provider's representation of the object as an unstructured.Unstructured.
func (a *DataLabelingJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.VertexAIDataLabelingJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAIDataLabelingJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.DataLabelingJob)
	return u, nil
}
