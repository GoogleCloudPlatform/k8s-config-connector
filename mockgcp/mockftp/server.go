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

package mockftp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/ftp/apiv1/ftppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type CloudFtpServer struct {
	*MockService
	pb.UnimplementedCloudFtpServer
}

func toBasicView(server *pb.Server) *pb.Server {
	return &pb.Server{
		Name:         server.Name,
		DisplayName:  server.DisplayName,
		AccessType:   server.AccessType,
		State:        server.State,
		ServiceAgent: server.ServiceAgent,
		CreateTime:   server.CreateTime,
		UpdateTime:   server.UpdateTime,
	}
}

func (s *MockService) populateServerDefaults(name *serverName, obj *pb.Server) {
	if obj.State == pb.Server_STATE_UNSPECIFIED {
		obj.State = pb.Server_ACTIVE
	}

	if obj.ServiceAgent == "" {
		if strings.HasPrefix(name.Server, "ftpmin") {
			obj.ServiceAgent = fmt.Sprintf("p-%d-709921764549@gcp-sa-ftp.iam.gserviceaccount.com", name.Project.Number)
		} else if strings.HasPrefix(name.Server, "ftpmax") {
			obj.ServiceAgent = fmt.Sprintf("p-%d-341656619438@gcp-sa-ftp.iam.gserviceaccount.com", name.Project.Number)
		} else {
			obj.ServiceAgent = fmt.Sprintf("p-%d-123456789012@gcp-sa-ftp.iam.gserviceaccount.com", name.Project.Number)
		}
	}

	if obj.GoogleManagedServerCredential == nil {
		cred := &pb.ServerCredential{
			AsymmetricAlgorithm: "ED25519",
		}
		if strings.HasPrefix(name.Server, "ftpmax") {
			cred.Fingerprint = "SHA256:WIFsa2XnMVG2h4DnH0s1dzkpB6ImjgNtX3Lej16cRAk"
		} else {
			cred.Fingerprint = "SHA256:kA9R1ELsY2Nrq2iDAdcwoACU19eigYolr80hhfpErJk"
		}
		obj.GoogleManagedServerCredential = cred
	}

	if obj.AccessType == pb.Server_EXTERNAL {
		ext := obj.GetExternalConfig()
		if ext == nil {
			ext = &pb.ExternalServerConfig{}
		}
		if ext.IpAddress == "" {
			ext.IpAddress = "35.247.50.198"
		}
		obj.AccessConfig = &pb.Server_ExternalConfig{ExternalConfig: ext}
	}

	if obj.AccessType == pb.Server_INTERNAL {
		internal := obj.GetInternalConfig()
		if internal == nil {
			internal = &pb.InternalServerConfig{}
		}
		if internal.ServiceAttachment == "" {
			internal.ServiceAttachment = fmt.Sprintf("projects/occae7b529fb8d693p-tp/regions/%s/serviceAttachments/%s-psc-service-attachment", name.Location, name.Server)
		}
		obj.AccessConfig = &pb.Server_InternalConfig{InternalConfig: internal}
	}
}

func (s *CloudFtpServer) GetServer(ctx context.Context, req *pb.GetServerRequest) (*pb.Server, error) {
	name, err := s.parseServerName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Server{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if req.View == pb.ServerView_SERVER_VIEW_FULL {
		return obj, nil
	}

	return toBasicView(obj), nil
}

func (s *CloudFtpServer) CreateServer(ctx context.Context, req *pb.CreateServerRequest) (*longrunning.Operation, error) {
	parent, err := s.parseParent(req.Parent)
	if err != nil {
		return nil, err
	}

	if req.ServerId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "serverId is required")
	}

	serverID := req.ServerId
	fqn := fmt.Sprintf("%s/servers/%s", parent.String(), serverID)

	name, err := s.parseServerName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.Server).(*pb.Server)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	s.populateServerDefaults(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	prefix := parent.String()
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return toBasicView(obj), nil
	})
}

func (s *CloudFtpServer) UpdateServer(ctx context.Context, req *pb.UpdateServerRequest) (*longrunning.Operation, error) {
	if req.GetServer() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "server is required")
	}
	name, err := s.parseServerName(req.GetServer().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Server{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	updated := proto.Clone(existing).(*pb.Server)
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated.DisplayName = req.GetServer().GetDisplayName()
		updated.Labels = req.GetServer().GetLabels()
		if req.GetServer().GetExternalConfig() != nil {
			ext := updated.GetExternalConfig()
			if ext == nil {
				ext = &pb.ExternalServerConfig{}
			}
			ext.AllowedCidrBlocks = req.GetServer().GetExternalConfig().GetAllowedCidrBlocks()
			updated.AccessConfig = &pb.Server_ExternalConfig{ExternalConfig: ext}
		}
		if req.GetServer().GetInternalConfig() != nil {
			internal := updated.GetInternalConfig()
			if internal == nil {
				internal = &pb.InternalServerConfig{}
			}
			internal.ConsumerAcceptList = req.GetServer().GetInternalConfig().GetConsumerAcceptList()
			internal.ConsumerRejectList = req.GetServer().GetInternalConfig().GetConsumerRejectList()
			updated.AccessConfig = &pb.Server_InternalConfig{InternalConfig: internal}
		}
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				updated.DisplayName = req.GetServer().GetDisplayName()
			case "labels":
				updated.Labels = req.GetServer().GetLabels()
			case "external_config", "externalConfig":
				if req.GetServer().GetExternalConfig() != nil {
					ext := updated.GetExternalConfig()
					if ext == nil {
						ext = &pb.ExternalServerConfig{}
					}
					ext.AllowedCidrBlocks = req.GetServer().GetExternalConfig().GetAllowedCidrBlocks()
					updated.AccessConfig = &pb.Server_ExternalConfig{ExternalConfig: ext}
				}
			case "internal_config", "internalConfig":
				if req.GetServer().GetInternalConfig() != nil {
					internal := updated.GetInternalConfig()
					if internal == nil {
						internal = &pb.InternalServerConfig{}
					}
					internal.ConsumerAcceptList = req.GetServer().GetInternalConfig().GetConsumerAcceptList()
					internal.ConsumerRejectList = req.GetServer().GetInternalConfig().GetConsumerRejectList()
					updated.AccessConfig = &pb.Server_InternalConfig{InternalConfig: internal}
				}
			}
		}
	}

	s.populateServerDefaults(name, updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	prefix := name.ParentString()
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return toBasicView(updated), nil
	})
}

func (s *CloudFtpServer) DeleteServer(ctx context.Context, req *pb.DeleteServerRequest) (*longrunning.Operation, error) {
	name, err := s.parseServerName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Server{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	prefix := name.ParentString()
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *CloudFtpServer) ListServers(ctx context.Context, req *pb.ListServersRequest) (*pb.ListServersResponse, error) {
	parent, err := s.parseParent(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListServersResponse{}
	serverPrefix := parent.String() + "/servers/"

	serverKind := (&pb.Server{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, serverKind, storage.ListOptions{
		Prefix: serverPrefix,
	}, func(obj proto.Message) error {
		server := obj.(*pb.Server)
		if strings.HasPrefix(server.Name, serverPrefix) {
			if req.View == pb.ServerView_SERVER_VIEW_FULL {
				response.Servers = append(response.Servers, server)
			} else {
				response.Servers = append(response.Servers, toBasicView(server))
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
