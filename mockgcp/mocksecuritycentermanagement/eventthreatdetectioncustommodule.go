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

package mocksecuritycentermanagement

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
)

type SecurityCenterManagementServer struct {
	*MockService
	pb.UnimplementedSecurityCenterManagementServer
}

func (s *SecurityCenterManagementServer) GetEventThreatDetectionCustomModule(ctx context.Context, req *pb.GetEventThreatDetectionCustomModuleRequest) (*pb.EventThreatDetectionCustomModule, error) {
	name, err := s.parseEventThreatDetectionCustomModuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EventThreatDetectionCustomModule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterManagementServer) CreateEventThreatDetectionCustomModule(ctx context.Context, req *pb.CreateEventThreatDetectionCustomModuleRequest) (*pb.EventThreatDetectionCustomModule, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Parent is required")
	}

	module := req.GetEventThreatDetectionCustomModule()
	if module == nil {
		return nil, status.Errorf(codes.InvalidArgument, "EventThreatDetectionCustomModule is required")
	}

	// Generate a deterministic ID based on parent and display name
	idSuffix := generateDeterministicID(parent, module.GetDisplayName())
	reqName := fmt.Sprintf("%s/eventThreatDetectionCustomModules/%s", parent, idSuffix)

	name, err := s.parseEventThreatDetectionCustomModuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(module).(*pb.EventThreatDetectionCustomModule)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.UpdateTime = now
	obj.LastEditor = "mock-editor@google.com"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterManagementServer) UpdateEventThreatDetectionCustomModule(ctx context.Context, req *pb.UpdateEventThreatDetectionCustomModuleRequest) (*pb.EventThreatDetectionCustomModule, error) {
	module := req.GetEventThreatDetectionCustomModule()
	if module == nil {
		return nil, status.Errorf(codes.InvalidArgument, "EventThreatDetectionCustomModule is required")
	}

	name, err := s.parseEventThreatDetectionCustomModuleName(module.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EventThreatDetectionCustomModule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	// Apply field-mask or update specific fields
	updateMask := req.GetUpdateMask()
	if updateMask == nil || len(updateMask.GetPaths()) == 0 {
		// If no update mask, update everything mutable
		obj.Config = module.Config
		obj.EnablementState = module.EnablementState
		obj.DisplayName = module.DisplayName
		obj.Description = module.Description
	} else {
		for _, path := range updateMask.GetPaths() {
			switch path {
			case "config":
				obj.Config = module.Config
			case "enablement_state", "enablementState":
				obj.EnablementState = module.EnablementState
			case "display_name", "displayName":
				obj.DisplayName = module.DisplayName
			case "description":
				obj.Description = module.Description
			default:
				return nil, status.Errorf(codes.InvalidArgument, "updating field %q is not supported", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()
	obj.LastEditor = "mock-editor@google.com"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterManagementServer) DeleteEventThreatDetectionCustomModule(ctx context.Context, req *pb.DeleteEventThreatDetectionCustomModuleRequest) (*emptypb.Empty, error) {
	name, err := s.parseEventThreatDetectionCustomModuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EventThreatDetectionCustomModule{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func generateDeterministicID(parent, displayName string) string {
	h := sha256.New()
	h.Write([]byte(parent))
	h.Write([]byte(displayName))
	sum := h.Sum(nil)
	val := binary.BigEndian.Uint64(sum[:8])
	// Return a 16-digit string representing a large number, similar to GCP's IDs
	return fmt.Sprintf("%d", val%10000000000000000)
}
