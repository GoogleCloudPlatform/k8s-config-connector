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
	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/genproto/googleapis/longrunning"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
)

func (s *DocumentProcessorV1) GetProcessor(ctx context.Context, req *pb.GetProcessorRequest) (*pb.Processor, error) {
	name, err := s.ParseProcessorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Processor{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "processor %q not found", req.GetName())
		}
		return nil, err
	}

	return obj, nil
}

func (s *DocumentProcessorV1) CreateProcessor(ctx context.Context, req *pb.CreateProcessorRequest) (*pb.Processor, error) {
	processorID := fmt.Sprintf("%x", time.Now().UnixNano())

	// documentai uses project number in the name
	parentTokens := strings.Split(req.GetParent(), "/")
	if len(parentTokens) != 4 || parentTokens[0] != "projects" || parentTokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q", req.GetParent())
	}
	projectID := parentTokens[1]
	location := parentTokens[3]

	project, err := s.Projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	projectNumber := fmt.Sprintf("%d", project.Number)

	reqName := fmt.Sprintf("projects/%s/locations/%s/processors/%s", projectNumber, location, processorID)
	name, err := s.ParseProcessorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.GetProcessor())
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.Processor_ENABLED

	processorVersionID := "stable"
	obj.DefaultProcessorVersion = fmt.Sprintf("%s/processorVersions/%s", fqn, processorVersionID)
	obj.ProcessEndpoint = fmt.Sprintf("https://%s-documentai.googleapis.com/v1/%s:process", location, fqn)

	for _, alias := range []string{"pretrained", "pretrained-next", "rc", "stable"} {
		obj.ProcessorVersionAliases = append(obj.ProcessorVersionAliases, &pb.ProcessorVersionAlias{
			Alias:            fmt.Sprintf("%s/processorVersions/%s", fqn, alias),
			ProcessorVersion: fmt.Sprintf("%s/processorVersions/%s", fqn, processorVersionID),
		})
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DocumentProcessorV1) DeleteProcessor(ctx context.Context, req *pb.DeleteProcessorRequest) (*longrunning.Operation, error) {
	name, err := s.ParseProcessorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObject := &pb.Processor{}
	if err := s.storage.Delete(ctx, fqn, deletedObject); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "processor %q not found", req.GetName())
		}
		return nil, err
	}

	return s.operations.DoneLRO(ctx, "", nil, &emptypb.Empty{})
}

// ProcessorName format: `projects/{project}/locations/{location}/processors/{processor}`
type ProcessorName struct {
	Project       string
	Location      string
	ProcessorName string
}

func (n *ProcessorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/processors/%s", n.Project, n.Location, n.ProcessorName)
}

// parseProcessorName parses a string into a processorName.
// The expected form is projects/{project}/locations/{location}/processors/{processor}
func (s *DocumentProcessorV1) ParseProcessorName(name string) (*ProcessorName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "processors" {
		name := &ProcessorName{
			Project:       tokens[1],
			Location:      tokens[3],
			ProcessorName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}

}
