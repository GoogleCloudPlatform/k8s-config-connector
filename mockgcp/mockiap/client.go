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

package mockiap

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/iap/apiv1/iappb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type clientName struct {
	Brand  *brandName
	Client string
}

func (n *clientName) String() string {
	return fmt.Sprintf("%s/identityAwareProxyClients/%s", n.Brand.String(), n.Client)
}

func (s *MockService) parseClientName(name string) (*clientName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "brands" && tokens[4] == "identityAwareProxyClients" {
		brand, err := s.parseBrandName(strings.Join(tokens[0:4], "/"))
		if err != nil {
			return nil, err
		}
		clientID := tokens[5]
		return &clientName{Brand: brand, Client: clientID}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid client name format: %q", name)
}

// CreateIdentityAwareProxyClient creates an Identity Aware Proxy owned client.
func (s *IdentityAwareProxyOAuthService) CreateIdentityAwareProxyClient(ctx context.Context, req *pb.CreateIdentityAwareProxyClientRequest) (*pb.IdentityAwareProxyClient, error) {
	brand, err := s.parseBrandName(req.GetParent())
	if err != nil {
		return nil, err
	}

	// Verify the brand exists
	brandObj := &pb.Brand{}
	if err := s.storage.Get(ctx, brand.String(), brandObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Brand %q not found.", brand.String())
		}
		return nil, status.Errorf(codes.Internal, "failed to get brand: %v", err)
	}

	clientID := fmt.Sprintf("%s-mock-client.apps.googleusercontent.com", brand.Brand)

	name := &clientName{
		Brand:  brand,
		Client: clientID,
	}

	fqn := name.String()

	obj := proto.Clone(req.GetIdentityAwareProxyClient()).(*pb.IdentityAwareProxyClient)
	obj.Name = fqn
	obj.Secret = "mock-client-secret-xyz123abc456"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create client: %v", err)
	}

	return obj, nil
}

// GetIdentityAwareProxyClient retrieves the Identity Aware Proxy owned client.
func (s *IdentityAwareProxyOAuthService) GetIdentityAwareProxyClient(ctx context.Context, req *pb.GetIdentityAwareProxyClientRequest) (*pb.IdentityAwareProxyClient, error) {
	name, err := s.parseClientName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.IdentityAwareProxyClient{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Client %q not found.", name.String())
		}
		return nil, status.Errorf(codes.Internal, "failed to get client: %v", err)
	}

	return obj, nil
}

// ListIdentityAwareProxyClients lists the Identity Aware Proxy owned clients.
func (s *IdentityAwareProxyOAuthService) ListIdentityAwareProxyClients(ctx context.Context, req *pb.ListIdentityAwareProxyClientsRequest) (*pb.ListIdentityAwareProxyClientsResponse, error) {
	brand, err := s.parseBrandName(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := brand.String() + "/identityAwareProxyClients/"

	response := &pb.ListIdentityAwareProxyClientsResponse{}

	clientKind := (&pb.IdentityAwareProxyClient{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, clientKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		client := obj.(*pb.IdentityAwareProxyClient)
		response.IdentityAwareProxyClients = append(response.IdentityAwareProxyClients, client)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteIdentityAwareProxyClient deletes the Identity Aware Proxy owned client.
func (s *IdentityAwareProxyOAuthService) DeleteIdentityAwareProxyClient(ctx context.Context, req *pb.DeleteIdentityAwareProxyClientRequest) (*emptypb.Empty, error) {
	name, err := s.parseClientName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.IdentityAwareProxyClient{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Client %q not found.", name.String())
		}
		return nil, status.Errorf(codes.Internal, "failed to delete client: %v", err)
	}

	return &emptypb.Empty{}, nil
}

// ResetIdentityAwareProxyClientSecret resets the client secret.
func (s *IdentityAwareProxyOAuthService) ResetIdentityAwareProxyClientSecret(ctx context.Context, req *pb.ResetIdentityAwareProxyClientSecretRequest) (*pb.IdentityAwareProxyClient, error) {
	name, err := s.parseClientName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.IdentityAwareProxyClient{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Client %q not found.", name.String())
		}
		return nil, status.Errorf(codes.Internal, "failed to get client: %v", err)
	}

	obj.Secret = "mock-client-secret-reset-789"
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update client secret: %v", err)
	}

	return obj, nil
}
