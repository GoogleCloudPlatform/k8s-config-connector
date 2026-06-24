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

	api "cloud.google.com/go/dataflow/apiv1beta3"
	pb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataflow/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func init() {
	registry.RegisterModel(krm.DataflowJobGVK, NewDataflowJobModel)
}

func NewDataflowJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataflowJobModel{config: *config}, nil
}

type dataflowJobModel struct {
	config config.ControllerConfig
}

var _ directbase.Model = &dataflowJobModel{}

func (m *dataflowJobModel) client(ctx context.Context) (*api.JobsV1Beta3Client, error) {
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	return gcpClient.newJobsClient(ctx)
}

func (m *dataflowJobModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataflowJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &dataflowJobAdapter{
		id:        identity.(*krm.DataflowJobIdentity),
		gcpClient: gcpClient,
	}, nil
}

func (m *dataflowJobModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.DataflowJobIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &dataflowJobAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type dataflowJobAdapter struct {
	id        *krm.DataflowJobIdentity
	gcpClient *api.JobsV1Beta3Client
	actual    *pb.Job
}

var _ directbase.Adapter = &dataflowJobAdapter{}

func (a *dataflowJobAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Job == "" {
		return false, nil
	}

	jobFQN := fmt.Sprintf("projects/%s/locations/%s/jobs/%s", a.id.Project, a.id.Location, a.id.Job)

	req := &pb.GetJobRequest{
		JobId:     a.id.Job,
		ProjectId: a.id.Project,
		Location:  a.id.Location,
		View:      pb.JobView_JOB_VIEW_SUMMARY,
	}
	job, err := a.gcpClient.GetJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataflow job %q: %w", jobFQN, err)
	}

	a.actual = job
	return true, nil
}

func (a *dataflowJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return fmt.Errorf("direct controller creation not implemented for DataflowJob")
}

func (a *dataflowJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return fmt.Errorf("direct controller update not implemented for DataflowJob")
}

func (a *dataflowJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	return false, fmt.Errorf("direct controller deletion not implemented for DataflowJob")
}

func (a *dataflowJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataflowJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataflowJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set identity fields manually
	obj.Spec.Region = direct.LazyPtr(a.id.Location)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Job)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Job)
	u.SetGroupVersionKind(krm.DataflowJobGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}
