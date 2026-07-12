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
	"strings"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	rootrefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.MonitoringUptimeCheckConfigGVK, newUptimeCheckConfigModel)
}

func newUptimeCheckConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &uptimeCheckConfigModel{config: config}, nil
}

type uptimeCheckConfigModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &uptimeCheckConfigModel{}

type uptimeCheckConfigAdapter struct {
	id *krm.MonitoringUptimeCheckConfigIdentity

	desired *pb.UptimeCheckConfig
	actual  *pb.UptimeCheckConfig

	uptimeCheckClient *monitoring.UptimeCheckClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &uptimeCheckConfigAdapter{}

// AdapterForObject implements the Model interface.
func (m *uptimeCheckConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	uptimeCheckClient, err := gcpClient.newUptimeCheckClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.MonitoringUptimeCheckConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := MonitoringUptimeCheckConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &uptimeCheckConfigAdapter{
		id:                id.(*krm.MonitoringUptimeCheckConfigIdentity),
		desired:           desiredProto,
		uptimeCheckClient: uptimeCheckClient,
	}, nil
}

func (m *uptimeCheckConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //monitoring.googleapis.com/projects/PROJECT_NUMBER/uptimeCheckConfigs/UPTIME_CHECK_ID
	if !strings.HasPrefix(url, "//monitoring.googleapis.com/") {
		return nil, nil
	}

	id := &krm.MonitoringUptimeCheckConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	uptimeCheckClient, err := gcpClient.newUptimeCheckClient(ctx)
	if err != nil {
		return nil, err
	}

	return &uptimeCheckConfigAdapter{
		id:                id,
		uptimeCheckClient: uptimeCheckClient,
	}, nil
}

// Find implements the Adapter interface.
func (a *uptimeCheckConfigAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.UptimeCheckConfig == "" {
		return false, nil
	}

	req := &pb.GetUptimeCheckConfigRequest{
		Name: a.fullyQualifiedName(),
	}
	uptimeCheckConfig, err := a.uptimeCheckClient.GetUptimeCheckConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = uptimeCheckConfig

	return true, nil
}

// Delete implements the Adapter interface.
func (a *uptimeCheckConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	exists, err := a.Find(ctx)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	req := &pb.DeleteUptimeCheckConfigRequest{
		Name: a.fullyQualifiedName(),
	}

	if err := a.uptimeCheckClient.DeleteUptimeCheckConfig(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting uptimeCheckConfig %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *uptimeCheckConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating uptimeCheckConfig", "id", fqn)

	parent := "projects/" + a.id.Project

	req := &pb.CreateUptimeCheckConfigRequest{
		Parent:            parent,
		UptimeCheckConfig: a.desired,
	}

	created, err := a.uptimeCheckClient.CreateUptimeCheckConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating uptimeCheckConfig: %w", err)
	}
	log.V(2).Info("created uptimeCheckConfig", "id", fqn)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func compareMonitoringUptimeCheckConfig(ctx context.Context, actual, desired *pb.UptimeCheckConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, MonitoringUptimeCheckConfigSpec_FromProto, MonitoringUptimeCheckConfigSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}

	// Copy write-only HTTP basic auth password from desired to maskedActual to avoid triggering continuous updates
	// because the GCP GET/Find API returns the password as unreadable/empty.
	if desired.GetHttpCheck() != nil && desired.GetHttpCheck().GetAuthInfo() != nil {
		if maskedActual.GetHttpCheck() != nil && maskedActual.GetHttpCheck().GetAuthInfo() != nil {
			maskedActual.GetHttpCheck().GetAuthInfo().Password = desired.GetHttpCheck().GetAuthInfo().GetPassword()
		}
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

// Update implements the Adapter interface.
func (a *uptimeCheckConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating object", "id", fqn)

	diffs, updateMask, err := compareMonitoringUptimeCheckConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateUptimeCheckConfigRequest{
			UptimeCheckConfig: a.desired,
			UpdateMask:        updateMask,
		}
		req.UptimeCheckConfig.Name = a.fullyQualifiedName()

		updated, err := a.uptimeCheckClient.UpdateUptimeCheckConfig(ctx, req)
		if err != nil {
			return err
		}
		log.V(2).Info("updated uptimeCheckConfig", "id", fqn)
		a.actual = updated
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *uptimeCheckConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("uptimeCheckConfig %q not found", a.fullyQualifiedName())
	}

	mc := &direct.MapContext{}
	spec := MonitoringUptimeCheckConfigSpec_FromProto(mc, a.actual)
	if err := mc.Err(); err != nil {
		return nil, fmt.Errorf("error converting uptimeCheckConfig from API %w", err)
	}

	spec.ProjectRef = &rootrefs.ProjectRef{
		External: a.id.Project,
	}

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting uptimeCheckConfig spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.UptimeCheckConfig)
	u.SetGroupVersionKind(krm.MonitoringUptimeCheckConfigGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *uptimeCheckConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.UptimeCheckConfig) error {
	mapCtx := &direct.MapContext{}
	status := MonitoringUptimeCheckConfigStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *uptimeCheckConfigAdapter) fullyQualifiedName() string {
	return a.id.String()
}
