// Copyright 2024 Google LLC
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

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1beta"
)

type GKEHubFeature struct {
	*MockService
	pb.UnimplementedGkeHubServer
}

func (s *GKEHubFeature) GetFeature(ctx context.Context, req *pb.GetFeatureRequest) (*pb.Feature, error) {
	name, err := s.parseFeatureName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Feature{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *GKEHubFeature) CreateFeature(ctx context.Context, req *pb.CreateFeatureRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/features/" + req.FeatureId
	name, err := s.parseFeatureName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Resource).(*pb.Feature)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *GKEHubFeature) UpdateFeature(ctx context.Context, req *pb.UpdateFeatureRequest) (*longrunning.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseFeatureName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Feature{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.Resource.GetLabels()
		case "multiclusteringress":
			obj.Spec.FeatureSpec = req.GetResource().Spec.GetFeatureSpec().(*pb.CommonFeatureSpec_Multiclusteringress)
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *GKEHubFeature) DeleteFeature(ctx context.Context, req *pb.DeleteFeatureRequest) (*longrunning.Operation, error) {
	name, err := s.parseFeatureName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Feature{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
