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
	registry.RegisterModel(krm.SaaSServiceMgmtUnitGVK, NewUnitModel)
}

func NewUnitModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &unitModel{config: config}, nil
}

var _ directbase.Model = &unitModel{}

type unitModel struct {
	config *config.ControllerConfig
}

func (m *unitModel) client(ctx context.Context) (*gcp.SaasDeploymentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewSaasDeploymentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SaasServiceMgmt deployments client: %w", err)
	}
	return gcpClient, nil
}

func (m *unitModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SaaSServiceMgmtUnit{}
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
	id, ok := idVal.(*krm.SaaSServiceMgmtUnitIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := SaaSServiceMgmtUnitSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(obj)
	desired.Annotations = filterSystemAnnotations(obj.GetAnnotations())
	desired.Name = id.String()

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &SaaSServiceMgmtUnitAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *unitModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type SaaSServiceMgmtUnitAdapter struct {
	id        *krm.SaaSServiceMgmtUnitIdentity
	gcpClient *gcp.SaasDeploymentsClient
	desired   *pb.Unit
	actual    *pb.Unit
}

var _ directbase.Adapter = &SaaSServiceMgmtUnitAdapter{}

func (a *SaaSServiceMgmtUnitAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting SaaSServiceMgmtUnit", "name", fqn)

	req := &pb.GetUnitRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetUnit(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SaaSServiceMgmtUnit %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *SaaSServiceMgmtUnitAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating SaaSServiceMgmtUnit", "name", fqn)

	req := &pb.CreateUnitRequest{
		Parent: parent,
		UnitId: a.id.Unit,
		Unit:   a.desired,
	}
	created, err := a.gcpClient.CreateUnit(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SaaSServiceMgmtUnit %s: %w", a.id.Unit, err)
	}
	log.V(2).Info("successfully created SaaSServiceMgmtUnit", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *SaaSServiceMgmtUnitAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating SaaSServiceMgmtUnit", "name", fqn)

	diffs, updateMask, err := compareUnit(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = fqn

		req := &pb.UpdateUnitRequest{
			Unit:       desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateUnit(ctx, req)
		if err != nil {
			return fmt.Errorf("updating SaaSServiceMgmtUnit %s: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *SaaSServiceMgmtUnitAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Unit) error {
	mapCtx := &direct.MapContext{}
	status := SaaSServiceMgmtUnitStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *SaaSServiceMgmtUnitAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SaaSServiceMgmtUnit{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SaaSServiceMgmtUnitSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.SaaSServiceMgmtUnitGVK)
	return u, nil
}

func (a *SaaSServiceMgmtUnitAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting SaaSServiceMgmtUnit", "name", fqn)

	req := &pb.DeleteUnitRequest{Name: fqn}
	err := a.gcpClient.DeleteUnit(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent SaaSServiceMgmtUnit, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting SaaSServiceMgmtUnit %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted SaaSServiceMgmtUnit", "name", fqn)
	return true, nil
}

func compareUnit(ctx context.Context, actual, desired *pb.Unit) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, SaaSServiceMgmtUnitSpec_FromProto, SaaSServiceMgmtUnitSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels
	maskedActual.Annotations = filterSystemAnnotations(actual.Annotations)
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func SaaSServiceMgmtUnitStatus_FromProto(mapCtx *direct.MapContext, in *pb.Unit) *krm.SaaSServiceMgmtUnitStatus {
	if in == nil {
		return nil
	}
	out := &krm.SaaSServiceMgmtUnitStatus{}
	if in.Name != "" {
		out.ExternalRef = &in.Name
	}
	out.ObservedState = SaaSServiceMgmtUnitObservedState_FromProto(mapCtx, in)
	return out
}
