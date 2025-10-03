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

package mockmonitoring

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/dashboard/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type DashboardsService struct {
	*MockService
	pb.UnimplementedDashboardsServiceServer
}

func (s *DashboardsService) GetDashboard(ctx context.Context, req *pb.GetDashboardRequest) (*pb.Dashboard, error) {
	name, err := s.parseDashboardName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dashboard{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *DashboardsService) CreateDashboard(ctx context.Context, req *pb.CreateDashboardRequest) (*pb.Dashboard, error) {
	reqName := req.GetDashboard().GetName() // req.GetParent() + "/dashboards/" + req.GetDashboard().GetName()
	name, err := s.parseDashboardName(reqName)
	if err != nil {
		return nil, err
	}

	if req.ValidateOnly {
		return nil, fmt.Errorf("ValidateOnly not implemented")
	}

	fqn := name.String()

	obj := proto.Clone(req.Dashboard).(*pb.Dashboard)

	defaulter := &dashboardDefaulter{}
	defaulter.visitDashboard(obj)

	validator := &dashboardValidator{}
	validator.visitDashboard(obj)
	if len(validator.errors) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", validator.errors[0])
	}

	obj.Name = fqn
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type dashboardDefaulter struct {
}

func (d *dashboardDefaulter) visitDashboard(obj *pb.Dashboard) {
	switch layout := obj.Layout.(type) {
	case *pb.Dashboard_ColumnLayout:
		d.visitColumnLayout(layout.ColumnLayout)
	case *pb.Dashboard_MosaicLayout:
		d.visitMosaicLayout(layout.MosaicLayout)
	}
}

func (d *dashboardDefaulter) visitColumnLayout(obj *pb.ColumnLayout) {
	for _, column := range obj.Columns {
		for _, widget := range column.Widgets {
			d.visitWidget(widget)
		}
	}
}

func (d *dashboardDefaulter) visitMosaicLayout(obj *pb.MosaicLayout) {
	for _, tile := range obj.Tiles {
		d.visitWidget(tile.Widget)
	}
}

func (d *dashboardDefaulter) visitWidget(obj *pb.Widget) {
	switch content := obj.Content.(type) {
	case *pb.Widget_XyChart:
		d.visitXYChartWidget(content)

	case *pb.Widget_Scorecard:
		d.visitScorecardWidget(content)

	case *pb.Widget_Text:
		d.visitTextWidget(content)

	case *pb.Widget_PieChart:
		d.visitPieChart(content.PieChart)

	case *pb.Widget_TimeSeriesTable:
		d.visitTimeSeriesTable(content.TimeSeriesTable)
	}
}

func (d *dashboardDefaulter) visitXYChartWidget(obj *pb.Widget_XyChart) {
	for _, dataSet := range obj.XyChart.DataSets {
		// TargetAxis defaults to Y1
		if dataSet.TargetAxis == pb.XyChart_DataSet_TARGET_AXIS_UNSPECIFIED {
			dataSet.TargetAxis = pb.XyChart_DataSet_Y1
		}
	}
	if yaxis := obj.XyChart.YAxis; yaxis != nil {
		// Defaults to linear
		if yaxis.Scale == pb.XyChart_Axis_SCALE_UNSPECIFIED {
			yaxis.Scale = pb.XyChart_Axis_LINEAR
		}
	}
}

func (d *dashboardDefaulter) visitScorecardWidget(obj *pb.Widget_Scorecard) {
}

func (d *dashboardDefaulter) visitTextWidget(obj *pb.Widget_Text) {
	if obj.Text.Style == nil {
		obj.Text.Style = &pb.Text_TextStyle{}
	}
	if obj.Text.Format == pb.Text_FORMAT_UNSPECIFIED {
		obj.Text.Format = pb.Text_MARKDOWN
	}
}

func (d *dashboardDefaulter) visitPieChart(obj *pb.PieChart) {
}

func (d *dashboardDefaulter) visitTimeSeriesTable(obj *pb.TimeSeriesTable) {
	for _, dataSet := range obj.DataSets {
		if dataSet.MinAlignmentPeriod == nil {
			dataSet.MinAlignmentPeriod = durationpb.New(time.Duration(0))
		}
	}
}

type dashboardValidator struct {
	errors []error
}

func (d *dashboardValidator) errorf(format string, args ...interface{}) {
	d.errors = append(d.errors, fmt.Errorf(format, args...))
}

func (d *dashboardValidator) invalidArgumentf(format string, args ...interface{}) {
	d.errors = append(d.errors, status.Errorf(codes.InvalidArgument, format, args...))
}

func (d *dashboardValidator) visitDashboard(obj *pb.Dashboard) {
	switch layout := obj.Layout.(type) {
	case *pb.Dashboard_ColumnLayout:
		d.visitColumnLayout(layout.ColumnLayout)
	case *pb.Dashboard_MosaicLayout:
		d.visitMosaicLayout(layout.MosaicLayout)
	}
}

func (d *dashboardValidator) visitColumnLayout(obj *pb.ColumnLayout) {
	for _, column := range obj.Columns {
		for _, widget := range column.Widgets {
			d.visitWidget(widget, obj)
		}
	}
}

func (d *dashboardValidator) visitMosaicLayout(obj *pb.MosaicLayout) {
	for _, tile := range obj.Tiles {
		d.visitWidget(tile.Widget, obj)
	}
}

func (d *dashboardValidator) visitWidget(obj *pb.Widget, layout proto.Message) {
	if obj.Content == nil {
		// Actually should be more like Dashboard is missing required field mosaicLayout.tiles[0].widget.content
		d.invalidArgumentf("Dashboard is missing required field ...widget.content.")
	}

	switch content := obj.Content.(type) {
	case *pb.Widget_XyChart:
		d.visitXYChartWidget(content.XyChart)

	case *pb.Widget_Scorecard:
		d.visitScorecardWidget(content)
	case *pb.Widget_Text:
		d.visitTextWidget(content)

	case *pb.Widget_SingleViewGroup:
		switch layout.(type) {
		case *pb.MosaicLayout:
			// OK
		default:
			// This is the error we get from GCP (should probably be singleViewGroup though)
			d.invalidArgumentf("dropdownGroup is only allowed in MosaicLayout.")
		}

	case *pb.Widget_SectionHeader:
		switch layout.(type) {
		case *pb.MosaicLayout:
			// OK
		default:
			d.invalidArgumentf("sectionHeader is only allowed in MosaicLayout.")
		}

	case *pb.Widget_CollapsibleGroup:
		switch layout.(type) {
		case *pb.MosaicLayout:
			// OK
		default:
			d.invalidArgumentf("collapsibleGroup is only allowed in MosaicLayout.")
		}
	}
}

func formatDuration(d *durationpb.Duration) string {
	return fmt.Sprintf("%ds", d.Seconds)
}

func (d *dashboardValidator) visitXYChartWidget(obj *pb.XyChart) {
	timeshiftDuration := obj.TimeshiftDuration
	if timeshiftDuration != nil && timeshiftDuration.AsDuration() != 0 {
		if timeshiftDuration.Seconds < 60 {
			// Should be columnLayout.columns[0].widgets[0].xyChart.timeshiftDuration ...
			d.errorf("Field columnLayout.columns[].widgets[].xyChart.timeshiftDuration has an invalid value of %q: must be greater than or equal to one minute.", formatDuration(timeshiftDuration))
			return
		}

		for _, dataSet := range obj.DataSets {
			switch dataSet.GetPlotType() {
			case pb.XyChart_DataSet_STACKED_BAR:
				// TODO: Should be Field columnLayout.columns[0].widgets[2].xyChart.dataSets[0].plotType ...
				d.errorf("Field columnLayout.columns[].widgets[].xyChart.dataSets[].plotType has an invalid value of %q: plot type is incompatible with XyChart's timeshiftDuration.", dataSet.GetPlotType())
			}
		}
	}
}

func (d *dashboardValidator) visitScorecardWidget(obj *pb.Widget_Scorecard) {
}

func (d *dashboardValidator) visitTextWidget(obj *pb.Widget_Text) {
}

func (s *DashboardsService) UpdateDashboard(ctx context.Context, req *pb.UpdateDashboardRequest) (*pb.Dashboard, error) {
	name, err := s.parseDashboardName(req.GetDashboard().GetName())
	if err != nil {
		return nil, err
	}

	if req.GetDashboard().GetEtag() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Update Dashboard should specify a non empty etag.")
	}

	if req.ValidateOnly {
		return nil, fmt.Errorf("ValidateOnly not implemented")
	}

	fqn := name.String()

	existing := &pb.Dashboard{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.Dashboard).(*pb.Dashboard)

	defaulter := &dashboardDefaulter{}
	defaulter.visitDashboard(updated)

	validator := &dashboardValidator{}
	validator.visitDashboard(updated)
	if len(validator.errors) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", validator.errors[0])
	}

	updated.Name = fqn
	updated.Etag = computeEtag(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *DashboardsService) DeleteDashboard(ctx context.Context, req *pb.DeleteDashboardRequest) (*empty.Empty, error) {
	name, err := s.parseDashboardName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Dashboard{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type dashboardName struct {
	Project       *projects.ProjectData
	DashboardName string
}

func (n *dashboardName) String() string {
	return fmt.Sprintf("projects/%d/dashboards/%s", n.Project.Number, n.DashboardName)
}

// parseDashboardName parses a string into a dashboardName.
// The expected form is projects/[PROJECT_ID_OR_NUMBER]/dashboards/[DASHBOARD_ID]
func (s *MockService) parseDashboardName(name string) (*dashboardName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "dashboards" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &dashboardName{
			Project:       project,
			DashboardName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
