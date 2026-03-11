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

package mockgkehub

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1"
)

type GKEHubV1 struct {
	*MockService
	pb.UnimplementedProjectsLocationsScopesServerServer
	pb.UnimplementedProjectsLocationsScopesNamespacesServerServer
}

func (s *GKEHubV1) GetProjectsLocationsScope(ctx context.Context, req *pb.GetProjectsLocationsScopeRequest) (*pb.Scope, error) {
	fqn := req.Name
	obj := &pb.Scope{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *GKEHubV1) CreateProjectsLocationsScope(ctx context.Context, req *pb.CreateProjectsLocationsScopeRequest) (*longrunning.Operation, error) {
	fqn := fmt.Sprintf("%s/scopes/%s", req.Parent, req.ScopeId)
	now := timestamppb.Now()

	obj := proto.Clone(req.ProjectsLocationsScope).(*pb.Scope)
	obj.Name = fqn
	obj.Uid = "google-generated-uuid"
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = &pb.ScopeLifecycleState{Code: "READY"}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &longrunning.Operation{} // Minimal metadata
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubV1) PatchProjectsLocationsScope(ctx context.Context, req *pb.PatchProjectsLocationsScopeRequest) (*longrunning.Operation, error) {
	fqn := req.Name
	obj := &pb.Scope{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Simple patch implementation
	paths := strings.Split(req.UpdateMask, ",")
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.ProjectsLocationsScope.Labels
		case "namespaceLabels":
			obj.NamespaceLabels = req.ProjectsLocationsScope.NamespaceLabels
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &longrunning.Operation{}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubV1) DeleteProjectsLocationsScope(ctx context.Context, req *pb.DeleteProjectsLocationsScopeRequest) (*longrunning.Operation, error) {
	fqn := req.Name
	oldObj := &pb.Scope{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}
	metadata := &longrunning.Operation{}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return &pb.Scope{}, nil
	})
}

func (s *GKEHubV1) GetProjectsLocationsScopesNamespace(ctx context.Context, req *pb.GetProjectsLocationsScopesNamespaceRequest) (*pb.Namespace, error) {
	fqn := req.Name
	obj := &pb.Namespace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *GKEHubV1) CreateProjectsLocationsScopesNamespace(ctx context.Context, req *pb.CreateProjectsLocationsScopesNamespaceRequest) (*longrunning.Operation, error) {
	fqn := fmt.Sprintf("%s/namespaces/%s", req.Parent, req.ScopeNamespaceId)
	now := timestamppb.Now()

	obj := proto.Clone(req.ProjectsLocationsScopesNamespace).(*pb.Namespace)
	obj.Name = fqn
	obj.Uid = "google-generated-uuid"
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = &pb.NamespaceLifecycleState{Code: "READY"}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &longrunning.Operation{}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubV1) PatchProjectsLocationsScopesNamespace(ctx context.Context, req *pb.PatchProjectsLocationsScopesNamespaceRequest) (*longrunning.Operation, error) {
	fqn := req.Name
	obj := &pb.Namespace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := strings.Split(req.UpdateMask, ",")
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.ProjectsLocationsScopesNamespace.Labels
		case "namespaceLabels":
			obj.NamespaceLabels = req.ProjectsLocationsScopesNamespace.NamespaceLabels
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &longrunning.Operation{}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GKEHubV1) DeleteProjectsLocationsScopesNamespace(ctx context.Context, req *pb.DeleteProjectsLocationsScopesNamespaceRequest) (*longrunning.Operation, error) {
	fqn := req.Name
	oldObj := &pb.Namespace{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}
	metadata := &longrunning.Operation{}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return &pb.Namespace{}, nil
	})
}
