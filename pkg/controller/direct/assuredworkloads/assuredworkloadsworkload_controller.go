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

package assuredworkloads

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/assuredworkloads/apiv1"
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AssuredWorkloadsWorkloadGVK, NewWorkloadModel)
}

func NewWorkloadModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkload{config: *config}, nil
}

var _ directbase.Model = &modelWorkload{}

type modelWorkload struct {
	config config.ControllerConfig
}

func (m *modelWorkload) client(ctx context.Context) (*gcp.Client, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AssuredWorkloads client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelWorkload) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AssuredWorkloadsWorkload{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &WorkloadAdapter{
		id:        id.(*krm.AssuredWorkloadsWorkloadIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWorkload) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id, err := krm.ParseWorkloadExternal(url)
	if err != nil {
		return nil, nil // Not a workload URL
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &WorkloadAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type WorkloadAdapter struct {
	id        *krm.AssuredWorkloadsWorkloadIdentity
	gcpClient *gcp.Client
	desired   *krm.AssuredWorkloadsWorkload
	actual    *pb.Workload
}

var _ directbase.Adapter = &WorkloadAdapter{}

func (a *WorkloadAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Workload == "" {
		return false, nil
	}

	req := &pb.GetWorkloadRequest{
		Name: a.id.String(),
	}
	workload, err := a.gcpClient.GetWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AssuredWorkloadsWorkload %q: %w", a.id.String(), err)
	}

	a.actual = workload
	return true, nil
}

func (a *WorkloadAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Workload", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredPb := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Assured Workloads requires labels in the Create request if we want them from the start.
	// We use the helper to get KCC labels.
	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(a.desired.GetLabels())

	parent := a.id.ParentString()
	req := &pb.CreateWorkloadRequest{
		Parent:   parent,
		Workload: desiredPb,
	}

	op, err := a.gcpClient.CreateWorkload(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Workload %q: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for Workload creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created Workload", "name", a.id.String())

	a.actual = created
	return a.updateStatus(ctx, createOp, created)
}

func (a *WorkloadAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Workload", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredPb := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(a.desired.GetLabels())
	desiredPb.Name = a.actual.Name

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	latest := a.actual
	if len(paths) != 0 {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)

		req := &pb.UpdateWorkloadRequest{
			Workload:   desiredPb,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
		}

		updated, err := a.gcpClient.UpdateWorkload(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Workload %q: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *WorkloadAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Workload) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AssuredWorkloadsWorkloadStatus{}
	status.ObservedState = AssuredWorkloadsWorkloadObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := latest.GetName()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *WorkloadAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.AssuredWorkloadsWorkload{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *AssuredWorkloadsWorkloadSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set identity fields
	obj.Spec.OrganizationRef = &refs.OrganizationRef{External: "organizations/" + a.id.Organization}
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = &a.id.Workload

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Workload)
	u.SetGroupVersionKind(krm.AssuredWorkloadsWorkloadGVK)

	return u, nil
}

func (a *WorkloadAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Workload", "name", a.id.String())

	req := &pb.DeleteWorkloadRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting Workload %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Workload", "name", a.id.String())
	return true, nil
}
