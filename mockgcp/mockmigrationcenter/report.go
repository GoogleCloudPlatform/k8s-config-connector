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

package mockmigrationcenter

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/migrationcenter/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type reportName struct {
	Project        *projects.ProjectData
	Location       string
	ReportConfigID string
	ReportID       string
}

func (n *reportName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/reportConfigs/%s/reports/%s", n.Project.ID, n.Location, n.ReportConfigID, n.ReportID)
}

func (s *MockService) parseReportName(name string) (*reportName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "reportConfigs" && tokens[6] == "reports" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &reportName{
			Project:        project,
			Location:       tokens[3],
			ReportConfigID: tokens[5],
			ReportID:       tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in expected format", name)
}

func (s *MigrationCenterV1) GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.Report, error) {
	name, err := s.parseReportName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Report{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *MigrationCenterV1) CreateReport(ctx context.Context, req *pb.CreateReportRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/reports/%s", req.GetParent(), req.GetReportId())
	name, err := s.parseReportName(reqName)
	if err != nil {
		return nil, err
	}

	parentConfigFQN := fmt.Sprintf("projects/%s/locations/%s/reportConfigs/%s", name.Project.ID, name.Location, name.ReportConfigID)
	parentConfig := &pb.ReportConfig{}
	if err := s.storage.Get(ctx, parentConfigFQN, parentConfig); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parentConfigFQN)
		}
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetReport()).(*pb.Report)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Report_SUCCEEDED

	obj.Summary = s.buildReportSummary(ctx, name, parentConfig)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return proto.Clone(obj).(*pb.Report), nil
	})
}

func (s *MigrationCenterV1) DeleteReport(ctx context.Context, req *pb.DeleteReportRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseReportName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Report{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *MigrationCenterV1) ListReports(ctx context.Context, req *pb.ListReportsRequest) (*pb.ListReportsResponse, error) {
	parentName, err := s.parseReportConfigName(req.GetParent())
	if err != nil {
		return nil, err
	}

	parentConfigFQN := parentName.String()
	parentConfig := &pb.ReportConfig{}
	if err := s.storage.Get(ctx, parentConfigFQN, parentConfig); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parentConfigFQN)
		}
		return nil, err
	}

	var reports []*pb.Report
	listKind := (&pb.Report{}).ProtoReflect().Descriptor()
	err = s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		r := obj.(*pb.Report)
		name, err := s.parseReportName(r.GetName())
		if err == nil && name.Project.ID == parentName.Project.ID && name.Location == parentName.Location && name.ReportConfigID == parentName.ReportConfigID {
			reports = append(reports, r)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListReportsResponse{
		Reports: reports,
	}, nil
}

func (s *MockService) buildReportSummary(ctx context.Context, name *reportName, reportConfig *pb.ReportConfig) *pb.ReportSummary {
	summary := &pb.ReportSummary{
		AllAssetsStats: defaultAssetAggregateStats(),
	}

	for _, assignment := range reportConfig.GetGroupPreferencesetAssignments() {
		groupFinding := &pb.ReportSummary_GroupFinding{
			AssetAggregateStats: defaultAssetAggregateStats(),
		}

		if parsedGroup, err := s.parseGroupName(assignment.GetGroup()); err == nil {
			group := &pb.Group{}
			if err := s.storage.Get(ctx, parsedGroup.String(), group); err == nil {
				groupFinding.DisplayName = group.GetDisplayName()
				groupFinding.Description = group.GetDescription()
			}
		}

		prefFinding := &pb.ReportSummary_GroupPreferenceSetFinding{
			ComputeEngineFinding: &pb.ReportSummary_ComputeEngineFinding{
				AllocatedRegions: []string{name.Location},
			},
			MachinePreferences:       defaultMachinePreferences(name.Location),
			MonthlyCostCompute:       &money.Money{CurrencyCode: "USD"},
			MonthlyCostNetworkEgress: &money.Money{CurrencyCode: "USD"},
			MonthlyCostOsLicense:     &money.Money{CurrencyCode: "USD"},
			MonthlyCostOther:         &money.Money{CurrencyCode: "USD"},
			MonthlyCostStorage:       &money.Money{CurrencyCode: "USD"},
			MonthlyCostTotal:         &money.Money{CurrencyCode: "USD"},
		}

		if parsedPref, err := s.parsePreferenceSetName(assignment.GetPreferenceSet()); err == nil {
			pref := &pb.PreferenceSet{}
			if err := s.storage.Get(ctx, parsedPref.String(), pref); err == nil {
				prefFinding.DisplayName = pref.GetDisplayName()
				prefFinding.Description = pref.GetDescription()
			}
		}

		groupFinding.PreferenceSetFindings = []*pb.ReportSummary_GroupPreferenceSetFinding{prefFinding}
		summary.GroupFindings = append(summary.GroupFindings, groupFinding)
	}

	return summary
}

func defaultAssetAggregateStats() *pb.ReportSummary_AssetAggregateStats {
	return &pb.ReportSummary_AssetAggregateStats{
		CoreCountHistogram: &pb.ReportSummary_HistogramChartData{
			Buckets: []*pb.ReportSummary_HistogramChartData_Bucket{
				{UpperBound: 2},
				{LowerBound: 2, UpperBound: 4},
				{LowerBound: 4, UpperBound: 16},
				{LowerBound: 16, UpperBound: 32},
				{LowerBound: 32, UpperBound: 9223372036854775807},
			},
		},
		MemoryBytesHistogram: &pb.ReportSummary_HistogramChartData{
			Buckets: []*pb.ReportSummary_HistogramChartData_Bucket{
				{UpperBound: 4294967296},
				{LowerBound: 4294967296, UpperBound: 17179869184},
				{LowerBound: 17179869184, UpperBound: 68719476736},
				{LowerBound: 68719476736, UpperBound: 137438953472},
				{LowerBound: 137438953472, UpperBound: 9223372036854775807},
			},
		},
		StorageBytesHistogram: &pb.ReportSummary_HistogramChartData{
			Buckets: []*pb.ReportSummary_HistogramChartData_Bucket{
				{UpperBound: 107374182400},
				{LowerBound: 107374182400, UpperBound: 268435456000},
				{LowerBound: 268435456000, UpperBound: 536870912000},
				{LowerBound: 536870912000, UpperBound: 1073741824000},
				{LowerBound: 1073741824000, UpperBound: 9223372036854775807},
			},
		},
	}
}

func defaultMachinePreferences(location string) *pb.VirtualMachinePreferences {
	return &pb.VirtualMachinePreferences{
		CommitmentPlan: pb.CommitmentPlan_COMMITMENT_PLAN_THREE_YEARS,
		ComputeEnginePreferences: &pb.ComputeEnginePreferences{
			LicenseType: pb.LicenseType_LICENSE_TYPE_DEFAULT,
			MachinePreferences: &pb.MachinePreferences{
				AllowedMachineSeries: []*pb.MachineSeries{
					{Code: "C2"},
					{Code: "C2D"},
					{Code: "E2"},
					{Code: "M1"},
					{Code: "M2"},
					{Code: "N1"},
					{Code: "N2"},
					{Code: "N2D"},
					{Code: "T2D"},
					{Code: "T2A"},
					{Code: "C3"},
					{Code: "M3"},
					{Code: "N4"},
					{Code: "C3D"},
					{Code: "C4"},
					{Code: "C4A"},
					{Code: "C4D"},
					{Code: "M4"},
					{Code: "N4D"},
					{Code: "N4A"},
				},
			},
		},
		RegionPreferences: &pb.RegionPreferences{
			PreferredRegions: []string{location},
		},
		SizingOptimizationStrategy: pb.SizingOptimizationStrategy_SIZING_OPTIMIZATION_STRATEGY_MODERATE,
		SoleTenancyPreferences: &pb.SoleTenancyPreferences{
			CommitmentPlan:        pb.SoleTenancyPreferences_COMMITMENT_3_YEAR,
			CpuOvercommitRatio:    2,
			HostMaintenancePolicy: pb.SoleTenancyPreferences_HOST_MAINTENANCE_POLICY_MIGRATE_WITHIN_NODE_GROUP,
			NodeTypes: []*pb.SoleTenantNodeType{
				{NodeName: "c2-node-60-240"},
				{NodeName: "m1-node-160-3844"},
				{NodeName: "m1-node-96-1433"},
				{NodeName: "m2-node-416-11776"},
				{NodeName: "m3-node-128-1952"},
				{NodeName: "m3-node-128-3904"},
				{NodeName: "n1-node-96-624"},
				{NodeName: "n2-node-128-864"},
				{NodeName: "n2-node-80-640"},
				{NodeName: "n2d-node-224-896"},
				{NodeName: "c3-node-176-1408"},
				{NodeName: "c3-node-176-352"},
				{NodeName: "c3-node-176-704"},
				{NodeName: "c3d-node-360-1440"},
				{NodeName: "c3d-node-360-2880"},
				{NodeName: "c3d-node-360-708"},
				{NodeName: "c4-node-192-1488"},
				{NodeName: "c4-node-192-384"},
				{NodeName: "c4-node-192-720"},
				{NodeName: "c4a-node-72-144"},
				{NodeName: "c4a-node-72-288"},
				{NodeName: "c4a-node-72-576"},
				{NodeName: "m2-node-416-8832"},
				{NodeName: "m4-node-224-2976"},
				{NodeName: "m4-node-224-5952"},
				{NodeName: "n2d-node-224-1792"},
				{NodeName: "n4-node-224-1372"},
				{NodeName: "n4a-node-96-624"},
			},
		},
		TargetProduct: pb.ComputeMigrationTargetProduct_COMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE,
		VmwareEnginePreferences: &pb.VmwareEnginePreferences{
			CommitmentPlan:                       pb.VmwareEnginePreferences_COMMITMENT_3_YEAR_MONTHLY_PAYMENTS,
			CpuOvercommitRatio:                   4,
			MemoryOvercommitRatio:                1,
			StorageDeduplicationCompressionRatio: 1.7,
		},
	}
}
