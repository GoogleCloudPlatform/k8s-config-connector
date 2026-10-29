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

package mockartifactregistry

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *ArtifactRegistryV1) GetRule(ctx context.Context, req *pb.GetRuleRequest) (*pb.Rule, error) {
	name, err := s.parseRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Rule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	if err := s.populateDefaultsForRule(ctx, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ArtifactRegistryV1) CreateRule(ctx context.Context, req *pb.CreateRuleRequest) (*pb.Rule, error) {
	parentName, err := s.parseArtifactRegistryName(req.Parent)
	if err != nil {
		return nil, err
	}

	parentRepo := &pb.Repository{}
	if err := s.storage.Get(ctx, parentName.String(), parentRepo); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	ruleId := req.RuleId
	if ruleId == "" && req.GetRule().GetName() != "" {
		parts := strings.Split(req.GetRule().GetName(), "/")
		ruleId = parts[len(parts)-1]
	}
	if ruleId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "rule_id is required")
	}

	reqName := parentName.String() + "/rules/" + ruleId
	name, err := s.parseRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.Rule)
	obj.Name = fqn

	if err := s.populateDefaultsForRule(ctx, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ArtifactRegistryV1) UpdateRule(ctx context.Context, req *pb.UpdateRuleRequest) (*pb.Rule, error) {
	reqName := req.GetRule().GetName()

	name, err := s.parseRuleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Rule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		name := obj.Name
		obj = proto.CloneOf(req.GetRule())
		obj.Name = name
	} else {
		for _, path := range paths {
			switch path {
			case "action":
				obj.Action = req.GetRule().GetAction()
			case "operation":
				obj.Operation = req.GetRule().GetOperation()
			case "condition":
				obj.Condition = req.GetRule().GetCondition()
			case "package_id", "packageId":
				obj.PackageId = req.GetRule().GetPackageId()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
			}
		}
	}

	if err := s.populateDefaultsForRule(ctx, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ArtifactRegistryV1) DeleteRule(ctx context.Context, req *pb.DeleteRuleRequest) (*emptypb.Empty, error) {
	name, err := s.parseRuleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Rule{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ArtifactRegistryV1) ListRules(ctx context.Context, req *pb.ListRulesRequest) (*pb.ListRulesResponse, error) {
	parentName, err := s.parseArtifactRegistryName(req.Parent)
	if err != nil {
		return nil, err
	}

	parentRepo := &pb.Repository{}
	if err := s.storage.Get(ctx, parentName.String(), parentRepo); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	findPrefix := parentName.String() + "/rules/"

	response := &pb.ListRulesResponse{}
	ruleKind := (&pb.Rule{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, ruleKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		rule := obj.(*pb.Rule)
		ruleCopy := proto.CloneOf(rule)
		if err := s.populateDefaultsForRule(ctx, ruleCopy); err != nil {
			return err
		}
		response.Rules = append(response.Rules, ruleCopy)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ArtifactRegistryV1) populateDefaultsForRule(ctx context.Context, obj *pb.Rule) error {
	return nil
}
