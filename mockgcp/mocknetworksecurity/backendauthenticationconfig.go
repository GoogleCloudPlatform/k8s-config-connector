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

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networksecurity/v1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NetworkSecurityV1Server struct {
	*MockService
	pb.UnimplementedNetworkSecurityServer
}

type backendAuthenticationConfigName struct {
	Project                     *projects.ProjectData
	Location                    string
	BackendAuthenticationConfig string
}

func (n *backendAuthenticationConfigName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", n.Project.ID, n.Location, n.BackendAuthenticationConfig)
}

func (s *NetworkSecurityV1Server) parseBackendAuthenticationConfigName(name string) (*backendAuthenticationConfigName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || strings.ToLower(tokens[4]) != "backendauthenticationconfigs" {
		return nil, fmt.Errorf("invalid BackendAuthenticationConfig name %q", name)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	return &backendAuthenticationConfigName{
		Project:                     project,
		Location:                    tokens[3],
		BackendAuthenticationConfig: tokens[5],
	}, nil
}

func (s *NetworkSecurityV1Server) CreateBackendAuthenticationConfig(ctx context.Context, req *pb.CreateBackendAuthenticationConfigRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/backendAuthenticationConfigs/" + req.BackendAuthenticationConfigId
	canonical, err := s.parseBackendAuthenticationConfigName(name)
	if err != nil {
		return nil, err
	}
	fqn := canonical.String()

	obj := proto.Clone(req.BackendAuthenticationConfig).(*pb.BackendAuthenticationConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name,
		Verb:                  "create",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *NetworkSecurityV1Server) GetBackendAuthenticationConfig(ctx context.Context, req *pb.GetBackendAuthenticationConfigRequest) (*pb.BackendAuthenticationConfig, error) {
	name, err := s.parseBackendAuthenticationConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackendAuthenticationConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *NetworkSecurityV1Server) UpdateBackendAuthenticationConfig(ctx context.Context, req *pb.UpdateBackendAuthenticationConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseBackendAuthenticationConfigName(req.GetBackendAuthenticationConfig().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.BackendAuthenticationConfig{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	// Apply update mask safely
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj = proto.Clone(req.GetBackendAuthenticationConfig()).(*pb.BackendAuthenticationConfig)
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetBackendAuthenticationConfig().GetDescription()
			case "labels":
				obj.Labels = req.GetBackendAuthenticationConfig().GetLabels()
			case "client_certificate", "clientCertificate":
				obj.ClientCertificate = req.GetBackendAuthenticationConfig().GetClientCertificate()
			case "trust_config", "trustConfig":
				obj.TrustConfig = req.GetBackendAuthenticationConfig().GetTrustConfig()
			case "well_known_roots", "wellKnownRoots":
				obj.WellKnownRoots = req.GetBackendAuthenticationConfig().GetWellKnownRoots()
			}
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *NetworkSecurityV1Server) DeleteBackendAuthenticationConfig(ctx context.Context, req *pb.DeleteBackendAuthenticationConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseBackendAuthenticationConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.BackendAuthenticationConfig{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
		ApiVersion:            "v1",
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return deleted, nil
	})
}
