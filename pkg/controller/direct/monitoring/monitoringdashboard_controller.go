// Copyright 2024 Google LLC
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
	"encoding/json"
	"fmt"
	"strings"

	api "google.golang.org/api/monitoring/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	monitoringprojects "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/monitoring/dashboard/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.MonitoringDashboardGVK, newDashboardModel)
}

func newDashboardModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dashboardModel{config: config}, nil
}

type dashboardModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &dashboardModel{}

type dashboardAdapter struct {
	id *krm.MonitoringDashboardIdentity

	desired *pb.Dashboard
	actual  *pb.Dashboard

	dashboardsService *api.ProjectsDashboardsService
	projectMapper     *monitoringprojects.ProjectMapper
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &dashboardAdapter{}

// AdapterForObject implements the Model interface.
func (m *dashboardModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	dashboardsService, err := gcpClient.newDashboardsService(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.MonitoringDashboard{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	projectRef, err := refs.ResolveProject(ctx, kube, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, kube, obj, projectRef); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := MonitoringDashboardSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	dashboardProject := projectRef.ProjectID
	if err := normalizeDashboardProto(ctx, m.config.ProjectMapper, desiredProto, dashboardProject); err != nil {
		return nil, err
	}

	return &dashboardAdapter{
		id:                id.(*krm.MonitoringDashboardIdentity),
		desired:           desiredProto,
		dashboardsService: dashboardsService,
		projectMapper:     m.config.ProjectMapper,
	}, nil
}

func (m *dashboardModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //monitoring.googleapis.com/projects/PROJECT_NUMBER/dashboards/DASHBOARD_ID
	if !strings.HasPrefix(url, "//monitoring.googleapis.com/") {
		return nil, nil
	}

	id := &krm.MonitoringDashboardIdentity{}
	if err := id.FromExternal(strings.TrimPrefix(url, "//monitoring.googleapis.com/")); err != nil {
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	dashboardsService, err := gcpClient.newDashboardsService(ctx)
	if err != nil {
		return nil, err
	}

	return &dashboardAdapter{
		id:                id,
		dashboardsService: dashboardsService,
		projectMapper:     m.config.ProjectMapper,
	}, nil
}

// Find implements the Adapter interface.
func (a *dashboardAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Dashboard == "" {
		return false, nil
	}

	dashboard, err := a.dashboardsService.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	dashboardJSON, err := dashboard.MarshalJSON()
	if err != nil {
		return false, fmt.Errorf("marshalling dashboard to json: %w", err)
	}

	pbDashboard := &pb.Dashboard{}
	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(dashboardJSON, pbDashboard); err != nil {
		return false, fmt.Errorf("unmarshalling dashboard to proto: %w", err)
	}

	if err := normalizeDashboardProto(ctx, a.projectMapper, pbDashboard, a.id.Project); err != nil {
		return false, err
	}

	a.actual = pbDashboard

	return true, nil
}

// Delete implements the Adapter interface.
func (a *dashboardAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Check if exists / already deleted
	// Technically we can just delete, but this is a little cleaner in logs etc.
	exists, err := a.Find(ctx)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	if _, err := a.dashboardsService.Delete(a.fullyQualifiedName()).Context(ctx).Do(); err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting dashboard %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *dashboardAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	parent := "projects/" + a.id.Project

	a.desired.Name = a.fullyQualifiedName()

	desiredJSON, err := protojson.Marshal(a.desired)
	if err != nil {
		return fmt.Errorf("marshalling desired dashboard: %w", err)
	}

	apiDashboard := &api.Dashboard{}
	if err := json.Unmarshal(desiredJSON, apiDashboard); err != nil {
		return fmt.Errorf("converting desired dashboard to api: %w", err)
	}

	log.V(2).Info("creating dashboard", "apiDashboard", apiDashboard)
	created, err := a.dashboardsService.Create(parent, apiDashboard).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating dashboard: %w", err)
	}
	log.V(2).Info("created dashboard", "dashboard", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	createdJSON, err := created.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marshalling created dashboard: %w", err)
	}

	createdPB := &pb.Dashboard{}
	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(createdJSON, createdPB); err != nil {
		return fmt.Errorf("unmarshalling created dashboard: %w", err)
	}

	mapCtx := &direct.MapContext{}
	status := MonitoringDashboardStatus_FromProto(mapCtx, createdPB)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

// Update implements the Adapter interface.
func (a *dashboardAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating object", "u", u)

	// TODO: Where/how do we want to enforce immutability?

	if ShouldReconcileBasedOnEtag(ctx, u, a.actual.Etag) {
		a.desired.Name = a.fullyQualifiedName()
		a.desired.Etag = a.actual.Etag

		desiredJSON, err := protojson.Marshal(a.desired)
		if err != nil {
			return fmt.Errorf("marshalling desired dashboard: %w", err)
		}

		apiDashboard := &api.Dashboard{}
		if err := json.Unmarshal(desiredJSON, apiDashboard); err != nil {
			return fmt.Errorf("converting desired dashboard to api: %w", err)
		}

		log.V(2).Info("updating dashboard", "apiDashboard", apiDashboard)
		updated, err := a.dashboardsService.Patch(a.fullyQualifiedName(), apiDashboard).Context(ctx).Do()
		if err != nil {
			return err
		}
		log.V(2).Info("updated dashboard", "dashboard", updated)

		updatedJSON, err := updated.MarshalJSON()
		if err != nil {
			return fmt.Errorf("marshalling updated dashboard: %w", err)
		}

		updatedPB := &pb.Dashboard{}
		if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(updatedJSON, updatedPB); err != nil {
			return fmt.Errorf("unmarshalling updated dashboard: %w", err)
		}

		a.actual = updatedPB
	}

	mapCtx := &direct.MapContext{}
	status := MonitoringDashboardStatus_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *dashboardAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("dashboard %q not found", a.fullyQualifiedName())
	}

	mc := &direct.MapContext{}
	spec := MonitoringDashboardSpec_FromProto(mc, a.actual)
	if err := mc.Err(); err != nil {
		return nil, fmt.Errorf("error converting dashboard from API %w", err)
	}

	spec.ProjectRef.External = a.id.Project

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting dashboard spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.Dashboard)
	u.SetGroupVersionKind(krm.MonitoringDashboardGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *dashboardAdapter) fullyQualifiedName() string {
	return a.id.String()
}
