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

package mockagentidentity

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/agentidentity/apiv1/agentidentitypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type authProviderServer struct {
	*MockService
	pb.UnimplementedAuthProviderServiceServer
}

func populateDefaultsForAuthProvider(obj *pb.AuthProvider) {
	// State defaults to ENABLED in GCP Agent Identity service.
	if obj.State == pb.AuthProvider_STATE_UNSPECIFIED {
		obj.State = pb.AuthProvider_ENABLED
	}

	// redirectUrl is a deterministic output-only field generated for threeLeggedOauth.
	if params := obj.GetAuthProviderTypeParams(); params != nil {
		if tlo := params.GetThreeLeggedOauth(); tlo != nil {
			tlo.RedirectUrl = fmt.Sprintf("https://agentidentitycredentials.googleapis.com/v1/%s/oauthcallback", obj.GetName())
		}
	}

	// TODO: AuthProvider proto3 zero-value fields:
	// - ThreeLeggedOAuth.enable_pkce (bool): proto3 omits false values on wire; real GCP also omits false.
	// - State (enum): AuthProvider_STATE_UNSPECIFIED (0) is defaulted to AuthProvider_ENABLED (1).
}

func stripInputOnly(obj *pb.AuthProvider) {
	obj.WorkloadIds = nil
	if params := obj.GetAuthProviderTypeParams(); params != nil {
		if apiKey := params.GetApiKey(); apiKey != nil {
			apiKey.ApiKey = ""
		}
		if tlo := params.GetThreeLeggedOauth(); tlo != nil {
			tlo.ClientSecret = ""
		}
		if two := params.GetTwoLeggedOauth(); two != nil {
			two.ClientSecret = ""
		}
	}
}

func (s *authProviderServer) GetAuthProvider(ctx context.Context, req *pb.GetAuthProviderRequest) (*pb.AuthProvider, error) {
	name, err := s.parseAuthProviderName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AuthProvider{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The requested resource was not found.")
		}
		return nil, err
	}

	res := proto.Clone(obj).(*pb.AuthProvider)
	populateDefaultsForAuthProvider(res)
	stripInputOnly(res)
	return res, nil
}

func (s *authProviderServer) CreateAuthProvider(ctx context.Context, req *pb.CreateAuthProviderRequest) (*pb.AuthProvider, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	authProviderID := req.GetAuthProviderId()
	if authProviderID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "authProviderId is required")
	}

	if req.GetAuthProvider() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "authProvider is required")
	}

	name := &authProviderName{
		Project:      parent.Project,
		Location:     parent.Location,
		AuthProvider: authProviderID,
	}
	fqn := name.String()

	now := timestamppb.Now()

	obj := proto.Clone(req.GetAuthProvider()).(*pb.AuthProvider)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now

	populateDefaultsForAuthProvider(obj)
	stripInputOnly(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return proto.Clone(obj).(*pb.AuthProvider), nil
}

func (s *authProviderServer) UpdateAuthProvider(ctx context.Context, req *pb.UpdateAuthProviderRequest) (*pb.AuthProvider, error) {
	if req.GetAuthProvider() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "authProvider is required")
	}

	name, err := s.parseAuthProviderName(req.GetAuthProvider().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.AuthProvider{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The requested resource was not found.")
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.AuthProvider)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated = proto.Clone(req.GetAuthProvider()).(*pb.AuthProvider)
		updated.Name = fqn
		updated.CreateTime = existing.CreateTime
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				updated.Description = req.GetAuthProvider().GetDescription()
			case "allowed_scopes", "allowedScopes":
				updated.AllowedScopes = req.GetAuthProvider().GetAllowedScopes()
			case "blocked_scopes", "blockedScopes":
				updated.BlockedScopes = req.GetAuthProvider().GetBlockedScopes()
			case "auth_provider_type_params", "authProviderTypeParams":
				updated.AuthProviderTypeParams = req.GetAuthProvider().GetAuthProviderTypeParams()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q not supported for update", path)
			}
		}
	}

	updated.UpdateTime = timestamppb.Now()
	populateDefaultsForAuthProvider(updated)
	stripInputOnly(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return proto.Clone(updated).(*pb.AuthProvider), nil
}

func (s *authProviderServer) DeleteAuthProvider(ctx context.Context, req *pb.DeleteAuthProviderRequest) (*emptypb.Empty, error) {
	name, err := s.parseAuthProviderName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.AuthProvider{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The requested resource was not found.")
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *authProviderServer) ListAuthProviders(ctx context.Context, req *pb.ListAuthProvidersRequest) (*pb.ListAuthProvidersResponse, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/authProviders/"
	var items []*pb.AuthProvider

	findKind := (&pb.AuthProvider{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := proto.Clone(obj).(*pb.AuthProvider)
		populateDefaultsForAuthProvider(item)
		stripInputOnly(item)
		items = append(items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListAuthProvidersResponse{
		AuthProviders: items,
	}, nil
}
