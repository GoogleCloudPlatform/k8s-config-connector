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

package saasservicemgmt

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1"
	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/saasservicemgmt/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SaaSServiceMgmtRolloutGVK, NewRolloutModel)
}

func NewRolloutModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &rolloutModel{config: config}, nil
}

var _ directbase.Model = &rolloutModel{}

type rolloutModel struct {
	config *config.ControllerConfig
}

func (m *rolloutModel) client(ctx context.Context) (*gcp.SaasRolloutsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewSaasRolloutsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SaasServiceMgmt rollouts client: %w", err)
	}
	return gcpClient, nil
}

func (m *rolloutModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SaaSServiceMgmtRollout{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.SaaSServiceMgmtRolloutIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := SaaSServiceMgmtRolloutSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())
	desired.Annotations = filterSystemAnnotations(obj.GetAnnotations())

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &SaaSServiceMgmtRolloutAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *rolloutModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.SaaSServiceMgmtRolloutIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SaaSServiceMgmtRolloutAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type SaaSServiceMgmtRolloutAdapter struct {
	id        *krm.SaaSServiceMgmtRolloutIdentity
	gcpClient *gcp.SaasRolloutsClient
	desired   *pb.Rollout
	actual    *pb.Rollout
}

var _ directbase.Adapter = &SaaSServiceMgmtRolloutAdapter{}

func (a *SaaSServiceMgmtRolloutAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting SaaSServiceMgmtRollout", "name", fqn)

	req := &pb.GetRolloutRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetRollout(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SaaSServiceMgmtRollout %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *SaaSServiceMgmtRolloutAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating SaaSServiceMgmtRollout", "name", fqn)

	req := &pb.CreateRolloutRequest{
		Parent:    parent,
		RolloutId: a.id.Rollout,
		Rollout:   a.desired,
	}
	created, err := a.gcpClient.CreateRollout(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SaaSServiceMgmtRollout %s: %w", a.id.Rollout, err)
	}
	log.V(2).Info("successfully created SaaSServiceMgmtRollout", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *SaaSServiceMgmtRolloutAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating SaaSServiceMgmtRollout", "name", fqn)

	diffs, updateMask, err := compareRollout(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = fqn

		req := &pb.UpdateRolloutRequest{
			Rollout:    desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateRollout(ctx, req)
		if err != nil {
			return fmt.Errorf("updating SaaSServiceMgmtRollout %s: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *SaaSServiceMgmtRolloutAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Rollout) error {
	mapCtx := &direct.MapContext{}
	status := SaaSServiceMgmtRolloutStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *SaaSServiceMgmtRolloutAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SaaSServiceMgmtRollout{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SaaSServiceMgmtRolloutSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Rollout)
	u.SetGroupVersionKind(krm.SaaSServiceMgmtRolloutGVK)
	return u, nil
}

func (a *SaaSServiceMgmtRolloutAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting SaaSServiceMgmtRollout", "name", fqn)

	req := &pb.DeleteRolloutRequest{Name: fqn}
	err := a.gcpClient.DeleteRollout(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent SaaSServiceMgmtRollout, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting SaaSServiceMgmtRollout %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted SaaSServiceMgmtRollout", "name", fqn)
	return true, nil
}

func compareRollout(ctx context.Context, actual, desired *pb.Rollout) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, SaaSServiceMgmtRolloutSpec_FromProto, SaaSServiceMgmtRolloutSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Labels = actual.Labels
	maskedActual.Annotations = filterSystemAnnotations(actual.Annotations)

	clonedDesired := proto.Clone(desired).(*pb.Rollout)

	if clonedDesired.Control == nil {
		clonedDesired.Control = maskedActual.Control
	} else if clonedDesired.Control.Action == pb.RolloutAction_ROLLOUT_ACTION_UNSPECIFIED {
		if maskedActual.Control != nil && maskedActual.Control.Action == pb.RolloutAction_ROLLOUT_ACTION_RUN {
			clonedDesired.Control.Action = pb.RolloutAction_ROLLOUT_ACTION_RUN
		}
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func SaaSServiceMgmtRolloutStatus_FromProto(mapCtx *direct.MapContext, in *pb.Rollout) *krm.SaaSServiceMgmtRolloutStatus {
	if in == nil {
		return nil
	}
	out := &krm.SaaSServiceMgmtRolloutStatus{}
	if in.Name != "" {
		out.ExternalRef = &in.Name
	}
	out.ObservedState = SaaSServiceMgmtRolloutObservedState_FromProto(mapCtx, in)
	return out
}
