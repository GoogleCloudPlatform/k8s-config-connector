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

package monitoring

import (
	"context"
	"fmt"

	gcpmonitoring "cloud.google.com/go/monitoring/apiv3/v2"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func init() {
	registry.RegisterModel(krm.MonitoringUptimeCheckConfigGVK, newUptimeCheckModel)
}

func newUptimeCheckModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &uptimeCheckModel{config: config}, nil
}

type uptimeCheckModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &uptimeCheckModel{}

type UptimeCheckAdapter struct {
	id        *krm.MonitoringUptimeCheckConfigIdentity
	gcpClient *gcpmonitoring.UptimeCheckClient
	desired   *pb.UptimeCheckConfig
	actual    *pb.UptimeCheckConfig
}

var _ directbase.Adapter = &UptimeCheckAdapter{}

func (m *uptimeCheckModel) client(ctx context.Context) (*gcpmonitoring.UptimeCheckClient, error) {
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	uptimeCheckClient, err := gcpClient.newUptimeCheckClient(ctx)
	if err != nil {
		return nil, err
	}
	return uptimeCheckClient, nil
}

func (m *uptimeCheckModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MonitoringUptimeCheckConfig{}
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
	desiredPb := MonitoringUptimeCheckConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &UptimeCheckAdapter{
		id:        id.(*krm.MonitoringUptimeCheckConfigIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *uptimeCheckModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *UptimeCheckAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.UptimeCheckConfig == "" {
		return false, nil
	}

	req := &pb.GetUptimeCheckConfigRequest{
		Name: a.fullyQualifiedName(),
	}
	uptimeCheck, err := a.gcpClient.GetUptimeCheckConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = uptimeCheck

	return true, nil
}

func (a *UptimeCheckAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MonitoringUptimeCheckConfig", "name", a.id)

	req := &pb.DeleteUptimeCheckConfigRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteUptimeCheckConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent MonitoringUptimeCheckConfig, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting MonitoringUptimeCheckConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MonitoringUptimeCheckConfig", "name", a.id)
	return true, nil
}

func (a *UptimeCheckAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating MonitoringUptimeCheckConfig", "id", fqn)

	parent := "projects/" + a.id.Project

	req := &pb.CreateUptimeCheckConfigRequest{
		Parent:            parent,
		UptimeCheckConfig: a.desired,
	}

	log.V(2).Info("creating MonitoringUptimeCheckConfig", "req", req)
	created, err := a.gcpClient.CreateUptimeCheckConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MonitoringUptimeCheckConfig: %w", err)
	}
	log.V(2).Info("created MonitoringUptimeCheckConfig", "fqn", fqn)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *UptimeCheckAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MonitoringUptimeCheckConfig", "name", a.id.String())

	diffs, updateMask, err := compareUptimeCheckConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateUptimeCheckConfigRequest{
			UptimeCheckConfig: a.desired,
			UpdateMask:        updateMask,
		}
		req.UptimeCheckConfig.Name = a.fullyQualifiedName()

		updated, err := a.gcpClient.UpdateUptimeCheckConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating MonitoringUptimeCheckConfig %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *UptimeCheckAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.UptimeCheckConfig) error {
	mapCtx := &direct.MapContext{}
	status := MonitoringUptimeCheckConfigStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *UptimeCheckAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MonitoringUptimeCheckConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MonitoringUptimeCheckConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.MonitoringUptimeCheckConfigGVK)
	return u, nil
}

func (a *UptimeCheckAdapter) fullyQualifiedName() string {
	return a.id.String()
}

func compareUptimeCheckConfig(ctx context.Context, actual, desired *pb.UptimeCheckConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *pb.UptimeCheckConfig
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := MonitoringUptimeCheckConfigSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = MonitoringUptimeCheckConfigSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
