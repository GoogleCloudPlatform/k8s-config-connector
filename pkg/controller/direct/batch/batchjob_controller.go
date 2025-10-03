// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:controller
// proto.service: google.cloud.batch.v1.BatchService
// proto.message: google.cloud.batch.v1.Job
// crd.type: BatchJob
// crd.version: v1alpha1

package batch

import (
	"context"
	"fmt"

	batch "cloud.google.com/go/batch/apiv1"
	batchpb "cloud.google.com/go/batch/apiv1/batchpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/golang/protobuf/proto"
)

func init() {
	registry.RegisterModel(krm.BatchJobGVK, NewJobModel)
}

func NewJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &jobModel{config: *config}, nil
}

var _ directbase.Model = &jobModel{}

type jobModel struct {
	config config.ControllerConfig
}

func (m *jobModel) Client(ctx context.Context, projectID string) (*batch.Client, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := batch.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building batch job client: %w", err)
	}

	return gcpClient, err
}

func (m *jobModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BatchJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewJobIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := BatchJobSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, err := m.Client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &jobAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *jobModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type jobAdapter struct {
	gcpClient *batch.Client
	id        *v1alpha1.JobIdentity
	desired   *batchpb.Job
	actual    *batchpb.Job
}

var _ directbase.Adapter = &jobAdapter{}

func (a *jobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting batch job", "name", a.id)

	req := &batchpb.GetJobRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting batch job %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *jobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating batch job", "name", a.id)

	req := &batchpb.CreateJobRequest{
		Parent: a.id.Parent().String(),
		Job:    a.desired,
		JobId:  a.id.ID(),
	}
	created, err := a.gcpClient.CreateJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating batch job %s: %w", a.id.String(), err)
	}
	log.Info("successfully created batch job in gcp", "name", a.id)

	status := &krm.BatchJobStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = BatchJobObservedState_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// BatchJob does not support update.
func (a *jobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating batch job", "name", a.id)

	desiredpb := proto.Clone(a.desired).(*batchpb.Job)
	paths, err := common.CompareProtoMessage(desiredpb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) != 0 {
		log.V(2).Info("This resource does not support update", "name", a.id.String())
		return nil
	}

	status := &krm.BatchJobStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *jobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting batch job", "name", a.id)

	req := &batchpb.DeleteJobRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteJob(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting batch job %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted batch job", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of batch job %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *jobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BatchJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BatchJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = direct.LazyPtr(a.id.Parent().Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.BatchJobGVK)

	u.Object = uObj
	return u, nil
}
