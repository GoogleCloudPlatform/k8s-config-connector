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

package mockdeveloperconnect

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type InsightsConfigServer struct {
	*MockService
	pb.UnimplementedInsightsConfigServiceServer
}

func (s *InsightsConfigServer) GetInsightsConfig(ctx context.Context, req *pb.GetInsightsConfigRequest) (*pb.InsightsConfig, error) {
	name, err := s.parseInsightsConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InsightsConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *InsightsConfigServer) CreateInsightsConfig(ctx context.Context, req *pb.CreateInsightsConfigRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/insightsConfigs/" + req.InsightsConfigId
	name, err := s.parseInsightsConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.InsightsConfig).(*pb.InsightsConfig)
	obj.Name = fqn
	obj.Labels = nil // Real GCP ignores labels during creation!

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = pb.InsightsConfig_PENDING

	s.normalizeInsightsConfigContext(ctx, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "create")
	return s.operations.StartLRO(ctx, req.Parent, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.InsightsConfig)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *InsightsConfigServer) UpdateInsightsConfig(ctx context.Context, req *pb.UpdateInsightsConfigRequest) (*longrunning.Operation, error) {
	reqName := req.GetInsightsConfig().GetName()
	name, err := s.parseInsightsConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InsightsConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update fields
	desired := req.GetInsightsConfig()
	obj.Labels = desired.Labels
	obj.Annotations = desired.Annotations
	if desired.InsightsConfigContext != nil {
		obj.InsightsConfigContext = desired.InsightsConfigContext
		s.normalizeInsightsConfigContext(ctx, obj)
	}
	obj.ArtifactConfigs = desired.ArtifactConfigs
	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "update")
	return s.operations.StartLRO(ctx, "projects/"+name.Project.ID+"/locations/"+name.Location, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.InsightsConfig)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *InsightsConfigServer) DeleteInsightsConfig(ctx context.Context, req *pb.DeleteInsightsConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseInsightsConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InsightsConfig{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "delete")
	return s.operations.StartLRO(ctx, "projects/"+name.Project.ID+"/locations/"+name.Location, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *InsightsConfigServer) ListInsightsConfigs(ctx context.Context, req *pb.ListInsightsConfigsRequest) (*pb.ListInsightsConfigsResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}

	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := "projects/" + project.ID + "/locations/" + tokens[3] + "/insightsConfigs/"

	var response pb.ListInsightsConfigsResponse
	err = s.storage.List(ctx, (&pb.InsightsConfig{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.InsightsConfig)
		response.InsightsConfigs = append(response.InsightsConfigs, item)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *InsightsConfigServer) normalizeProjects(ctx context.Context, projects *pb.Projects) {
	if projects == nil {
		return
	}
	for i, projectID := range projects.ProjectIds {
		idOrNumber := projectID
		if strings.HasPrefix(idOrNumber, "projects/") {
			idOrNumber = strings.TrimPrefix(idOrNumber, "projects/")
		}
		proj, err := s.Projects.GetProjectByIDOrNumber(idOrNumber)
		if err == nil {
			projects.ProjectIds[i] = "projects/" + strconv.FormatInt(proj.Number, 10)
		}
	}
}

func (s *InsightsConfigServer) normalizeAppHubApplication(ctx context.Context, appHubApp string) string {
	if appHubApp == "" {
		return ""
	}
	const servicePrefix = "//apphub.googleapis.com/"
	trimmed := appHubApp
	if strings.HasPrefix(trimmed, servicePrefix) {
		trimmed = strings.TrimPrefix(trimmed, servicePrefix)
	}
	// Format of trimmed is: "projects/<project>/locations/<location>/applications/<app>"
	tokens := strings.Split(trimmed, "/")
	if len(tokens) >= 2 && tokens[0] == "projects" {
		proj, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err == nil {
			tokens[1] = strconv.FormatInt(proj.Number, 10)
		}
	}
	return servicePrefix + strings.Join(tokens, "/")
}

func (s *InsightsConfigServer) normalizeInsightsConfigContext(ctx context.Context, obj *pb.InsightsConfig) {
	if obj == nil {
		return
	}
	if context, ok := obj.InsightsConfigContext.(*pb.InsightsConfig_AppHubApplication); ok {
		context.AppHubApplication = s.normalizeAppHubApplication(ctx, context.AppHubApplication)
	}
	if context, ok := obj.InsightsConfigContext.(*pb.InsightsConfig_Projects); ok {
		s.normalizeProjects(ctx, context.Projects)
	}
}

func constructOperationMetadata(target, verb string) *pb.OperationMetadata {
	return &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		Target:                target,
		Verb:                  verb,
		RequestedCancellation: false,
	}
}
