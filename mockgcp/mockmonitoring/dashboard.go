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
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/dashboard/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *DashboardsService) GetDashboard(ctx context.Context, req *pb.GetDashboardRequest) (*pb.Dashboard, error) {
	name, err := s.parseDashboardName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dashboard{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
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
	obj.Name = fqn

	defaulter := &dashboardDefaulter{}
	defaulter.visitDashboard(obj)

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
	}
}

func (d *dashboardDefaulter) visitColumnLayout(obj *pb.ColumnLayout) {
	for _, column := range obj.Columns {
		for _, widget := range column.Widgets {
			d.visitWidget(widget)
		}
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
	// TODO: Set style to empty when protos are updated
}

func (s *DashboardsService) UpdateDashboard(ctx context.Context, req *pb.UpdateDashboardRequest) (*pb.Dashboard, error) {
	name, err := s.parseDashboardName(req.GetDashboard().GetName())
	if err != nil {
		return nil, err
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
	return "projects/" + n.Project.ID + "/dashboards/" + n.DashboardName
}

// parseDashboardName parses a string into a dashboardName.
// The expected form is projects/[PROJECT_ID_OR_NUMBER]/dashboards/[DASHBOARD_ID]
func (s *MockService) parseDashboardName(name string) (*dashboardName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "dashboards" {
		project, err := s.projects.GetProjectByID(tokens[1])
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

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
