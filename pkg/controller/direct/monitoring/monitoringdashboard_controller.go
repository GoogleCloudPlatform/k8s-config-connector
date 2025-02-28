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
	"fmt"
	"strings"

	api "cloud.google.com/go/monitoring/dashboard/apiv1"
	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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
	projectID  string
	resourceID string

	desired *pb.Dashboard
	actual  *pb.Dashboard

	dashboardsClient *api.DashboardsClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &dashboardAdapter{}

// AdapterForObject implements the Model interface.
func (m *dashboardModel) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	dashboardsClient, err := gcpClient.newDashboardsClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.MonitoringDashboard{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectRef, err := refs.ResolveProject(ctx, kube, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if err := common.NormalizeReferences(ctx, kube, obj, projectRef); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := MonitoringDashboardSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &dashboardAdapter{
		projectID:        projectID,
		resourceID:       resourceID,
		desired:          desiredProto,
		dashboardsClient: dashboardsClient,
	}, nil
}

func (m *dashboardModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //monitoring.googleapis.com/projects/PROJECT_NUMBER/dashboards/DASHBOARD_ID
	if !strings.HasPrefix(url, "//monitoring.googleapis.com/") {
		return nil, nil
	}

	tokens := strings.Split(strings.TrimPrefix(url, "//monitoring.googleapis.com/"), "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "dashboards" {
		gcpClient, err := newGCPClient(m.config)
		if err != nil {
			return nil, fmt.Errorf("building gcp client: %w", err)
		}

		dashboardsClient, err := gcpClient.newDashboardsClient(ctx)
		if err != nil {
			return nil, err
		}

		return &dashboardAdapter{
			projectID:        tokens[1],
			resourceID:       tokens[3],
			dashboardsClient: dashboardsClient,
		}, nil
	}

	return nil, nil

}

// Find implements the Adapter interface.
func (a *dashboardAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &pb.GetDashboardRequest{
		Name: a.fullyQualifiedName(),
	}
	dashboard, err := a.dashboardsClient.GetDashboard(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = dashboard

	return true, nil
}

// Delete implements the Adapter interface.
func (a *dashboardAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteDashboardRequest{
		Name: a.fullyQualifiedName(),
	}

	if err := a.dashboardsClient.DeleteDashboard(ctx, req); err != nil {
		return false, fmt.Errorf("deleting dashboard %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *dashboardAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	parent := "projects/" + a.projectID

	req := &pb.CreateDashboardRequest{
		Parent:    parent,
		Dashboard: a.desired,
	}
	req.Dashboard.Name = a.fullyQualifiedName()

	log.V(2).Info("creating dashboard", "req", req)
	created, err := a.dashboardsClient.CreateDashboard(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dashboard: %w", err)
	}
	log.V(2).Info("created dashboard", "dashboard", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	mapCtx := &direct.MapContext{}
	status := MonitoringDashboardStatus_FromProto(mapCtx, created)
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
		req := &pb.UpdateDashboardRequest{
			Dashboard: a.desired,
		}
		req.Dashboard.Name = a.fullyQualifiedName()
		req.Dashboard.Etag = a.actual.Etag

		log.V(2).Info("updating dashboard", "request", req)
		updated, err := a.dashboardsClient.UpdateDashboard(ctx, req)
		if err != nil {
			return err
		}
		log.V(2).Info("updated dashboard", "dashboard", updated)
		a.actual = updated
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

	spec.ProjectRef.External = a.projectID

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting dashboard spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.resourceID)
	u.SetGroupVersionKind(krm.MonitoringDashboardGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *dashboardAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/dashboards/%s", a.projectID, a.resourceID)
}
