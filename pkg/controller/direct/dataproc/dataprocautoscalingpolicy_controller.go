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
// proto.service: google.cloud.dataproc.v1.AutoscalingPolicyService
// proto.message: google.cloud.dataproc.v1.AutoscalingPolicy
// crd.type: DataprocAutoscalingPolicy
// crd.version: v1beta1

package dataproc

import (
	"context"
	"fmt"

	dataproc "cloud.google.com/go/dataproc/v2/apiv1"
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DataprocAutoscalingPolicyGVK, NewDataprocAutoscalingPolicyModel)
}

func NewDataprocAutoscalingPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataprocAutoscalingPolicyModel{config: *config}, nil
}

var _ directbase.Model = &dataprocAutoscalingPolicyModel{}

type dataprocAutoscalingPolicyModel struct {
	config config.ControllerConfig
}

func (m *dataprocAutoscalingPolicyModel) client(ctx context.Context) (*dataproc.AutoscalingPolicyClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	return dataproc.NewAutoscalingPolicyRESTClient(ctx, opts...)
}

func (m *dataprocAutoscalingPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataprocAutoscalingPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	policyID, ok := id.(*krm.DataprocAutoscalingPolicyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", id)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &dataprocAutoscalingPolicyAdapter{
		gcpClient: gcpClient,
		id:        policyID,
		desired:   obj,
	}, nil
}

func (m *dataprocAutoscalingPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

var _ directbase.Adapter = &dataprocAutoscalingPolicyAdapter{}

type dataprocAutoscalingPolicyAdapter struct {
	gcpClient *dataproc.AutoscalingPolicyClient
	id        *krm.DataprocAutoscalingPolicyIdentity
	desired   *krm.DataprocAutoscalingPolicy
	actual    *pb.AutoscalingPolicy
}

func (a *dataprocAutoscalingPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataproc autoscalingpolicy", "name", a.id)

	req := &pb.GetAutoscalingPolicyRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetAutoscalingPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataproc autoscalingpolicy %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *dataprocAutoscalingPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataproc autoscalingpolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	policy := DataprocAutoscalingPolicySpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	policy.Id = a.id.AutoscalingPolicy

	req := &pb.CreateAutoscalingPolicyRequest{
		Parent: a.id.ParentString(),
		Policy: policy,
	}
	created, err := a.gcpClient.CreateAutoscalingPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataproc autoscalingpolicy %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created dataproc autoscalingpolicy in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *dataprocAutoscalingPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataproc autoscalingpolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	policy := DataprocAutoscalingPolicySpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	policy.Id = a.id.AutoscalingPolicy
	policy.Name = a.id.String()

	diffs, err := compareDataprocAutoscalingPolicy(ctx, a.actual, policy)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateAutoscalingPolicyRequest{
			Policy: policy,
		}
		updated, err := a.gcpClient.UpdateAutoscalingPolicy(ctx, req)
		if err != nil {
			return fmt.Errorf("updating dataproc autoscalingpolicy %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated dataproc autoscalingpolicy", "name", a.id)
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *dataprocAutoscalingPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AutoscalingPolicy) error {
	status := &krm.DataprocAutoscalingPolicyStatus{}
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}

func compareDataprocAutoscalingPolicy(ctx context.Context, actual, desired *pb.AutoscalingPolicy) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DataprocAutoscalingPolicySpec_v1beta1_FromProto, DataprocAutoscalingPolicySpec_v1beta1_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Id = desired.Id
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.AutoscalingPolicy)

	diffs, _, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func (a *dataprocAutoscalingPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataproc autoscalingpolicy", "name", a.id)

	req := &pb.DeleteAutoscalingPolicyRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAutoscalingPolicy(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting dataproc autoscalingpolicy %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataproc autoscalingpolicy", "name", a.id)

	return true, nil
}

// Export implements the Adapter interface.
func (a *dataprocAutoscalingPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataprocAutoscalingPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataprocAutoscalingPolicySpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &parent.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Region
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.AutoscalingPolicy)
	u.SetGroupVersionKind(krm.DataprocAutoscalingPolicyGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}
