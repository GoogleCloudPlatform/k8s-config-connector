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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	registry.RegisterModel(krm.SaaSServiceMgmtRolloutKindGVK, NewRolloutKindModel)
}

func NewRolloutKindModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &rolloutKindModel{config: config}, nil
}

var _ directbase.Model = &rolloutKindModel{}

type rolloutKindModel struct {
	config *config.ControllerConfig
}

func (m *rolloutKindModel) client(ctx context.Context) (*gcp.SaasRolloutsClient, error) {
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

func (m *rolloutKindModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SaaSServiceMgmtRolloutKind{}
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
	id, ok := idVal.(*krm.SaaSServiceMgmtRolloutKindIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := SaaSServiceMgmtRolloutKindSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(obj)
	desired.Annotations = filterSystemAnnotations(obj.GetAnnotations())

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &SaaSServiceMgmtRolloutKindAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *rolloutKindModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type SaaSServiceMgmtRolloutKindAdapter struct {
	id        *krm.SaaSServiceMgmtRolloutKindIdentity
	gcpClient *gcp.SaasRolloutsClient
	desired   *pb.RolloutKind
	actual    *pb.RolloutKind
}

var _ directbase.Adapter = &SaaSServiceMgmtRolloutKindAdapter{}

func (a *SaaSServiceMgmtRolloutKindAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting SaaSServiceMgmtRolloutKind", "name", fqn)

	req := &pb.GetRolloutKindRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetRolloutKind(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SaaSServiceMgmtRolloutKind %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *SaaSServiceMgmtRolloutKindAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating SaaSServiceMgmtRolloutKind", "name", fqn)

	req := &pb.CreateRolloutKindRequest{
		Parent:        parent,
		RolloutKindId: a.id.RolloutKind,
		RolloutKind:   a.desired,
	}
	created, err := a.gcpClient.CreateRolloutKind(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SaaSServiceMgmtRolloutKind %s: %w", a.id.RolloutKind, err)
	}
	log.V(2).Info("successfully created SaaSServiceMgmtRolloutKind", "name", created.Name)

	getReq := &pb.GetRolloutKindRequest{Name: fqn}
	latest, err := a.gcpClient.GetRolloutKind(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting SaaSServiceMgmtRolloutKind %s after create: %w", fqn, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *SaaSServiceMgmtRolloutKindAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating SaaSServiceMgmtRolloutKind", "name", fqn)

	diffs, updateMask, err := compareRolloutKind(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = fqn

		req := &pb.UpdateRolloutKindRequest{
			RolloutKind: desired,
			UpdateMask:  updateMask,
		}

		updated, err := a.gcpClient.UpdateRolloutKind(ctx, req)
		if err != nil {
			return fmt.Errorf("updating SaaSServiceMgmtRolloutKind %s: %w", fqn, err)
		}
		latest = updated

		getReq := &pb.GetRolloutKindRequest{Name: fqn}
		latest, err = a.gcpClient.GetRolloutKind(ctx, getReq)
		if err != nil {
			return fmt.Errorf("getting SaaSServiceMgmtRolloutKind %s after update: %w", fqn, err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *SaaSServiceMgmtRolloutKindAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.RolloutKind) error {
	mapCtx := &direct.MapContext{}
	status := SaaSServiceMgmtRolloutKindStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *SaaSServiceMgmtRolloutKindAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SaaSServiceMgmtRolloutKind{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SaaSServiceMgmtRolloutKindSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.SaaSServiceMgmtRolloutKindGVK)
	return u, nil
}

func (a *SaaSServiceMgmtRolloutKindAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting SaaSServiceMgmtRolloutKind", "name", fqn)

	req := &pb.DeleteRolloutKindRequest{Name: fqn}
	err := a.gcpClient.DeleteRolloutKind(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent SaaSServiceMgmtRolloutKind, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting SaaSServiceMgmtRolloutKind %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted SaaSServiceMgmtRolloutKind", "name", fqn)
	return true, nil
}

func compareRolloutKind(ctx context.Context, actual, desired *pb.RolloutKind) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, SaaSServiceMgmtRolloutKindSpec_FromProto, SaaSServiceMgmtRolloutKindSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels
	maskedActual.Annotations = filterSystemAnnotations(actual.Annotations)

	clonedDesired := proto.CloneOf(desired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
