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

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
)

type DeveloperConnectServer struct {
	*MockService
	pb.UnimplementedDeveloperConnectServer
}

func (s *DeveloperConnectServer) GetAccountConnector(ctx context.Context, req *pb.GetAccountConnectorRequest) (*pb.AccountConnector, error) {
	name, err := s.parseAccountConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AccountConnector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DeveloperConnectServer) CreateAccountConnector(ctx context.Context, req *pb.CreateAccountConnectorRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/accountConnectors/" + req.AccountConnectorId
	name, err := s.parseAccountConnectorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.AccountConnector).(*pb.AccountConnector)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	// Generate oauthStartUri
	obj.OauthStartUri = "https://console.cloud.google.com/developer-connect/start-oauth/" + name.Location + "/" + name.AccountConnector + "?project=" + name.Project.ID

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "create")
	return s.operations.StartLRO(ctx, req.Parent, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.AccountConnector)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *DeveloperConnectServer) UpdateAccountConnector(ctx context.Context, req *pb.UpdateAccountConnectorRequest) (*longrunning.Operation, error) {
	reqName := req.GetAccountConnector().GetName()
	name, err := s.parseAccountConnectorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AccountConnector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) > 0 {
		for _, path := range paths {
			switch path {
			case "annotations":
				obj.Annotations = req.GetAccountConnector().GetAnnotations()
			case "labels":
				obj.Labels = req.GetAccountConnector().GetLabels()
			case "providerOauthConfig":
				obj.AccountConnectorConfig = &pb.AccountConnector_ProviderOauthConfig{
					ProviderOauthConfig: req.GetAccountConnector().GetProviderOauthConfig(),
				}
			}
		}
	} else {
		// If no update mask, replace fields
		obj.Labels = req.GetAccountConnector().GetLabels()
		obj.Annotations = req.GetAccountConnector().GetAnnotations()
		if req.GetAccountConnector().GetProviderOauthConfig() != nil {
			obj.AccountConnectorConfig = &pb.AccountConnector_ProviderOauthConfig{
				ProviderOauthConfig: req.GetAccountConnector().GetProviderOauthConfig(),
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "update")
	return s.operations.StartLRO(ctx, "projects/"+name.Project.ID+"/locations/"+name.Location, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.AccountConnector)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *DeveloperConnectServer) DeleteAccountConnector(ctx context.Context, req *pb.DeleteAccountConnectorRequest) (*longrunning.Operation, error) {
	name, err := s.parseAccountConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AccountConnector{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "delete")
	return s.operations.StartLRO(ctx, "projects/"+name.Project.ID+"/locations/"+name.Location, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
