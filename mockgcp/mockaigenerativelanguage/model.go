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

// +tool:mockgcp-support
// proto.service: google.ai.generativelanguage.v1.Models
// proto.message: google.ai.generativelanguage.v1.Model

package mockaigenerativelanguage

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/ai/generativelanguage/v1"
)

type modelService struct {
	*MockService
	pb.UnimplementedModelServiceServer
}

func (s *modelService) GetModel(ctx context.Context, req *pb.GetModelRequest) (*pb.Model, error) {
	name, err := s.parseModelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Model{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Model %v not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

// TODO acpana there is not Delete or Create in the pb

type modelName struct {
	Model string
}

func (n *modelName) String() string {
	return fmt.Sprintf("models/%s", n.Model)
}

func (s *MockService) parseModelName(name string) (*modelName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "models" {
		return &modelName{
			Model: tokens[1],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}


