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
// proto.service: google.cloud.assuredworkloads.v1.AssuredWorkloadsService
// proto.message: google.cloud.assuredworkloads.v1.Workload
// crd.type: AssuredWorkloadsWorkload
// crd.version: v1alpha1

package assuredworkloads

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/assuredworkloads/apiv1"
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.AssuredWorkloadsWorkloadGVK, NewAssuredWorkloadsWorkloadModel)
}

func NewAssuredWorkloadsWorkloadModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAssuredWorkloadsWorkload{config: *config}, nil
}

var _ directbase.Model = &modelAssuredWorkloadsWorkload{}

type modelAssuredWorkloadsWorkload struct {
	config config.ControllerConfig
}

func (m *modelAssuredWorkloadsWorkload) client(ctx context.Context) (*gcp.Client, error) {
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	return gcpClient.newAssuredWorkloadsClient(ctx)
}

func (m *modelAssuredWorkloadsWorkload) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AssuredWorkloadsWorkload{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &AssuredWorkloadsWorkloadAdapter{
		id:        id.(*krm.AssuredWorkloadsWorkloadIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelAssuredWorkloadsWorkload) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AssuredWorkloadsWorkloadIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &AssuredWorkloadsWorkloadAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type AssuredWorkloadsWorkloadAdapter struct {
	id        *krm.AssuredWorkloadsWorkloadIdentity
	gcpClient *gcp.Client
	desired   *pb.Workload
	actual    *pb.Workload
}

var _ directbase.Adapter = &AssuredWorkloadsWorkloadAdapter{}

func (a *AssuredWorkloadsWorkloadAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AssuredWorkloadsWorkload", "name", a.id)

	if a.id.Workload == "" {
		log.V(2).Info("no workload ID indicates creation intention", "name", a.id)
		return false, nil
	}

	req := &pb.GetWorkloadRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AssuredWorkloadsWorkload %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *AssuredWorkloadsWorkloadAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating AssuredWorkloadsWorkload", "id", fqn)

	parent := a.id.ParentString()

	req := &pb.CreateWorkloadRequest{
		Parent:   parent,
		Workload: a.desired,
	}
	op, err := a.gcpClient.CreateWorkload(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AssuredWorkloadsWorkload %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AssuredWorkloadsWorkload creation %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created AssuredWorkloadsWorkload", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *AssuredWorkloadsWorkloadAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AssuredWorkloadsWorkload", "name", a.id.String())

	diffs, updateMask, err := compareWorkload(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		a.desired.Name = a.id.String()
		req := &pb.UpdateWorkloadRequest{
			Workload:   a.desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateWorkload(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AssuredWorkloadsWorkload %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *AssuredWorkloadsWorkloadAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Workload) error {
	mapCtx := &direct.MapContext{}
	status := AssuredWorkloadsWorkloadObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	krmStatus := &krm.AssuredWorkloadsWorkloadStatus{
		ObservedState: status,
		ExternalRef:   direct.LazyPtr(latest.Name),
	}
	return op.UpdateStatus(ctx, krmStatus, nil)
}

func (a *AssuredWorkloadsWorkloadAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AssuredWorkloadsWorkload{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AssuredWorkloadsWorkloadSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.Workload)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Workload)
	u.SetGroupVersionKind(krm.AssuredWorkloadsWorkloadGVK)
	return u, nil
}

func (a *AssuredWorkloadsWorkloadAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AssuredWorkloadsWorkload", "name", a.id)

	req := &pb.DeleteWorkloadRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent AssuredWorkloadsWorkload, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting AssuredWorkloadsWorkload %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AssuredWorkloadsWorkload", "name", a.id)
	return true, nil
}

func compareWorkload(ctx context.Context, actual, desired *pb.Workload) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, AssuredWorkloadsWorkloadSpec_FromProto, AssuredWorkloadsWorkloadSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
