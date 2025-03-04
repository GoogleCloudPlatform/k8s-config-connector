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

package mockdocumentai

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/documentai/v1"
)

func (s *DocumentProcessorV1) GetProcessor(ctx context.Context, req *pb.GetProcessorRequest) (*pb.Processor, error) {
	name, err := s.ParseProcessorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Processor{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DocumentProcessorV1) CreateProcessor(ctx context.Context, req *pb.CreateProcessorRequest) (*pb.Processor, error) {
	name, err := s.ParseProcessorName(req.GetProcessor().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	processorVersion := proto.Clone(req.GetProcessor()).(*pb.ProcessorVersion)
	processorVersion.Name = fqn
	now := time.Now()
	req.Processor.CreateTime = timestamppb.New(now)
	req.Processor.State = pb.Processor_ENABLED

	if err := s.storage.Create(ctx, fqn, req.Processor); err != nil {
		return nil, err
	}

	return req.Processor, nil
}

func (s *DocumentProcessorV1) DeleteProcessor(ctx context.Context, req *pb.DeleteProcessorRequest) (*longrunningpb.Operation, error) {
	name, err := s.ParseProcessorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObject := &pb.Processor{}
	if err := s.storage.Delete(ctx, fqn, deletedObject); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

// ProcessorName format: `projects/{project}/locations/{location}/processors/{processor}`
type ProcessorName struct {
	Project       *projects.ProjectData
	Location      string
	ProcessorName string
}

func (n *ProcessorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/processors/%s", n.Project.ID, n.Location, n.ProcessorName)
}

// parseProcessorName parses a string into a processorName.
// The expected form is projects/{project}/locations/{location}/processors/{processor}
func (s *DocumentProcessorV1) ParseProcessorName(name string) (*ProcessorName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "processors" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &ProcessorName{
			Project:       project,
			Location:      tokens[3],
			ProcessorName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}

}
