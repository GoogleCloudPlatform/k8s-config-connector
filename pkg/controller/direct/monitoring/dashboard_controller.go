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
	"reflect"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/durationpb"
	"k8s.io/klog/v2"

	pb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/golang/protobuf/ptypes/empty"

	dashboard "cloud.google.com/go/monitoring/dashboard/apiv1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

// Add creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddDashboardController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.MonitoringDashboardGVK

	return directbase.Add(mgr, gvk, &dashboardModel{config: *config})
}

type dashboardModel struct {
	config controller.Config
}

var dashboardMapping = NewMapping(&pb.Dashboard{}, &krm.MonitoringDashboard{},
	TODO("name"), // TODO: Should be ResourceID?
	Spec("displayName"),
	Ignore("etag"),
	Spec("gridLayout"),
	Spec("mosaicLayout"),
	Spec("rowLayout"),
	Spec("columnLayout"),
	TODO("dashboardFilters"),
	TODO("labels"),
).
	MapNested(&pb.GridLayout{}, &krm.DashboardGridLayout{}, "columns", "widgets").
	MapNested(&pb.MosaicLayout{}, &krm.DashboardMosaicLayout{}, "columns", "tiles").
	MapNested(&pb.MosaicLayout_Tile{}, &krm.DashboardTiles{},
		"xPos", "yPos", "width", "height", "widget").
	MapNested(&pb.RowLayout{}, &krm.DashboardRowLayout{}, "rows").
	MapNested(&pb.RowLayout_Row{}, &krm.DashboardRows{}, "weight", "widgets").
	MapNested(&pb.ColumnLayout{}, &krm.DashboardColumnLayout{}, "columns").
	MapNested(&pb.ColumnLayout_Column{}, &krm.DashboardColumns{}, "weight", "widgets").
	// TODO: Why DashboardWidget & DashboardWidgets ... looks like a typo?
	MapNested(&pb.Widget{}, &krm.DashboardWidget{}, "title", "xyChart", "scorecard", "text", "blank",
		TODO("timeSeriesTable"), TODO("alertChart"), TODO("timeSeriesTable"), TODO("collapsibleGroup"), "logsPanel").
	MapNested(&pb.Widget{}, &krm.DashboardWidgets{}, "title", "xyChart", "scorecard", "text", "blank",
		TODO("timeSeriesTable"), TODO("alertChart"), TODO("timeSeriesTable"), TODO("collapsibleGroup"), "logsPanel").
	MapNested(&pb.LogsPanel{}, &krm.DashboardLogsPanel{},
		"filter",
		ResourceRef("resourceNames", &refMapProjects{})).
	MapNested(&pb.XyChart{}, &krm.DashboardXyChart{},
		"dataSets",
		Transformed("timeshiftDuration", &durationTransform{}),
		"thresholds",
		"xAxis",
		"yAxis",
		TODO("y2Axis"),
		"chartOptions").
	MapNested(&pb.XyChart_Axis{}, &krm.DashboardXAxis{},
		"label",
		"scale").
	MapNested(&pb.XyChart_Axis{}, &krm.DashboardYAxis{},
		"label",
		"scale").
	MapNested(&pb.XyChart_DataSet{}, &krm.DashboardDataSets{},
		"plotType",
		"timeSeriesQuery",
		"minAlignmentPeriod",
		TODO("targetAxis"),
		"legendTemplate").
	MapNested(&pb.Scorecard{}, &krm.DashboardScorecard{},
		"timeSeriesQuery",
		"gaugeView",
		"sparkChartView",
		"thresholds").
	MapNested(&pb.Text{}, &krm.DashboardText{}, "content", "format").
	MapNested(&empty.Empty{}, &krm.DashboardBlank{}).
	MapNested(&pb.Threshold{}, &krm.DashboardThresholds{},
		"label",
		"value",
		"color",
		"direction",
		TODO("targetAxis")).
	MapNested(&pb.TimeSeriesQuery{}, &krm.DashboardTimeSeriesQuery{},
		"timeSeriesFilter",
		"timeSeriesFilterRatio",
		"timeSeriesQueryLanguage",
		TODO("prometheusQuery"),
		"unitOverride").
	MapNested(&pb.TimeSeriesFilter{}, &krm.DashboardTimeSeriesFilter{},
		"filter",
		"aggregation",
		"secondaryAggregation",
		"pickTimeSeriesFilter",
		TODO("statisticalTimeSeriesFilter")).
	MapNested(&pb.PickTimeSeriesFilter{}, &krm.DashboardPickTimeSeriesFilter{},
		"direction",
		"numTimeSeries",
		"rankingMethod").
	// TODO: We need to map statisticalTimeSeriesFilter
	// MapNested(&pb.StatisticalTimeSeriesFilter{}, &krm.DashboardBlank{},
	// 	TODO("numTimeSeries"),
	// 	TODO("rankingMethod")).
	MapNested(&pb.Aggregation{}, &krm.DashboardAggregation{},
		"alignmentPeriod",
		"perSeriesAligner",
		"crossSeriesReducer",
		"groupByFields").
	MustBuild()

type durationTransform struct {
}

var _ Mapper = &durationTransform{}

func (*durationTransform) CloudToKRM(in reflect.Value) (reflect.Value, error) {
	var d *durationpb.Duration

	switch in.Kind() {
	case reflect.Ptr:
		if in.IsNil() {
			return reflect.ValueOf(d), nil
		}
		in = in.Elem()
	}

	switch in.Kind() {
	case reflect.Struct:
		switch v := in.Interface().(type) {
		case durationpb.Duration:
			s := v.AsDuration().String()
			return reflect.ValueOf(s), nil
		default:
			return reflect.Value{}, fmt.Errorf("unhandled kind in durationTransform::KRMToCloud: %T", v)
		}
	default:
		return reflect.Value{}, fmt.Errorf("unhandled kind in durationTransform::KRMToCloud: %v", in.Kind())
	}

}

func (*durationTransform) KRMToCloud(in reflect.Value) (reflect.Value, error) {
	var d *durationpb.Duration

	s := ""
	if in.Kind() == reflect.Ptr {
		if in.IsNil() {
			return reflect.ValueOf(d), nil
		}
		in = in.Elem()
	}
	switch in.Kind() {
	case reflect.String:
		s = in.String()
	default:
		return reflect.Value{}, fmt.Errorf("unhandled kind in durationTransform::CloudToKRM: %v", in.Kind())
	}
	if s == "" {
		return reflect.ValueOf(d), nil
	}
	duration, err := time.ParseDuration(s)
	if err != nil {
		return reflect.Value{}, fmt.Errorf("invalid duration %q", s)
	}
	d = durationpb.New(duration)
	return reflect.ValueOf(d), nil
}

type dashboardAdapter struct {
	projectID   string
	dashboardID string

	desired *krm.MonitoringDashboard
	actual  *krm.MonitoringDashboard

	gcp *dashboard.DashboardsClient
}

func (m *dashboardModel) client(ctx context.Context) (*dashboard.DashboardsClient, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	gcpClient, err := dashboard.NewDashboardsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dashboard client: %w", err)
	}
	return gcpClient, err
}

func (m *dashboardModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.MonitoringDashboard{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	// TODO: Use name or request resourceID to be set on create?
	dashboardID := ValueOf(obj.Spec.ResourceID)
	if dashboardID == "" {
		dashboardID = obj.GetName()
	}
	if dashboardID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	return &dashboardAdapter{
		projectID:   projectID,
		dashboardID: dashboardID,
		desired:     obj,
		gcp:         gcp,
	}, nil
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

func (a *dashboardAdapter) Find(ctx context.Context) (bool, error) {
	if a.dashboardID == "" {
		return false, nil
	}

	req := &pb.GetDashboardRequest{
		Name: a.fullyQualifiedName(),
	}
	gcpObject, err := a.gcp.GetDashboard(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("dashboard was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.MonitoringDashboard{}
	if err := dashboardMapping.Map(gcpObject, u); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *dashboardAdapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteDashboardRequest{
		Name: a.fullyQualifiedName(),
	}
	if err := a.gcp.DeleteDashboard(ctx, req); err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting dashboard: %w", err)
	}

	return true, nil
}

func (a *dashboardAdapter) Create(ctx context.Context, obj *unstructured.Unstructured) error {
	desired := &pb.Dashboard{}
	if err := dashboardMapping.Map(a.desired, desired); err != nil {
		return err
	}

	desired.Name = a.fullyQualifiedName()

	req := &pb.CreateDashboardRequest{
		Parent:    fmt.Sprintf("projects/%s", a.projectID),
		Dashboard: desired,
	}

	created, err := a.gcp.CreateDashboard(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dashboard: %w", err)
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("created dashboard", "dashboard", created)
	// TODO: Return created object
	return nil
}

func (a *dashboardAdapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {
	desired := &pb.Dashboard{}
	if err := dashboardMapping.Map(a.desired, desired); err != nil {
		return nil, err
	}

	desired.Name = a.fullyQualifiedName()

	req := &pb.UpdateDashboardRequest{
		Dashboard: desired,
	}

	updated, err := a.gcp.UpdateDashboard(ctx, req)
	if err != nil {
		return nil, err
	}
	log := klog.FromContext(ctx)
	log.V(2).Info("updated dashboard", "dashboard", updated)
	// TODO: Return updated object
	return nil, nil
}

func (a *dashboardAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/dashboards/%s", a.projectID, a.dashboardID)
}
