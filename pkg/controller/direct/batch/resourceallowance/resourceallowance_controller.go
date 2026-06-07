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

// +tool:controller
// proto.service: google.cloud.batch.v1alpha.BatchService
// proto.message: google.cloud.batch.v1alpha.ResourceAllowance
// crd.type: CloudBatchResourceAllowance
// crd.version: v1alpha1

package resourceallowance

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/batch/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/batch/resourceallowance/pb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func init() {
	registry.RegisterModel(krm.CloudBatchResourceAllowanceGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (pb.BatchServiceClient, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("dialing batch service: %w", err)
	}

	return pb.NewBatchServiceClient(conn), nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudBatchResourceAllowance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}
	mapCtx := &direct.MapContext{}
	desired := CloudBatchResourceAllowanceSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id.(*v1alpha1.ResourceAllowanceIdentity),
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type adapter struct {
	gcpClient pb.BatchServiceClient
	id        *v1alpha1.ResourceAllowanceIdentity
	desired   *pb.ResourceAllowance
	actual    *pb.ResourceAllowance
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting batch resource allowance", "name", a.id)

	req := &pb.GetResourceAllowanceRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetResourceAllowance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting batch resource allowance %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating batch resource allowance", "name", a.id)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateResourceAllowanceRequest{
		Parent:              parent,
		ResourceAllowance:   a.desired,
		ResourceAllowanceId: a.id.ResourceAllowance,
	}
	created, err := a.gcpClient.CreateResourceAllowance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating batch resource allowance %s: %w", a.id.String(), err)
	}
	log.Info("successfully created batch resource allowance in gcp", "name", a.id)

	status := &krm.CloudBatchResourceAllowanceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudBatchResourceAllowanceObservedState_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating batch resource allowance", "name", a.id)

	desiredpb := proto.Clone(a.desired).(*pb.ResourceAllowance)
	desiredpb.Name = a.id.String()

	paths, report, err := common.CompareProtoMessageStructuredDiff(desiredpb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		status := &krm.CloudBatchResourceAllowanceStatus{}
		mapCtx := &direct.MapContext{}
		status.ObservedState = CloudBatchResourceAllowanceObservedState_FromProto(mapCtx, a.actual)
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateResourceAllowanceRequest{
		ResourceAllowance: desiredpb,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		},
	}

	updated, err := a.gcpClient.UpdateResourceAllowance(ctx, req)
	if err != nil {
		return fmt.Errorf("updating batch resource allowance %s: %w", a.id.String(), err)
	}
	log.Info("successfully updated batch resource allowance", "name", a.id)

	status := &krm.CloudBatchResourceAllowanceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudBatchResourceAllowanceObservedState_FromProto(mapCtx, updated)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting batch resource allowance", "name", a.id)

	req := &pb.DeleteResourceAllowanceRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteResourceAllowance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting batch resource allowance %s: %w", a.id.String(), err)
	}
	log.Info("successfully initiated deletion of batch resource allowance", "name", a.id)

	doneFunc := func() (bool, error) {
		exists, err := a.Find(ctx)
		if err != nil {
			return false, err
		}
		return !exists, nil
	}
	err = common.WaitForDoneOrTimeout(ctx, 2*time.Second, doneFunc)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudBatchResourceAllowance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudBatchResourceAllowanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ResourceAllowance)
	u.SetGroupVersionKind(krm.CloudBatchResourceAllowanceGVK)

	u.Object = uObj
	return u, nil
}
