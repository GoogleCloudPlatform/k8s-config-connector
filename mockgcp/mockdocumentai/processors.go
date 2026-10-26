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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/documentai/v1"
)

func (s *DocumentProcessorV1) populateDefaultsForProcessor(obj *pb.Processor, name *ProcessorName) {
	if obj.Name == "" {
		obj.Name = name.String()
	}
	if obj.Type == "" {
		obj.Type = "OCR_PROCESSOR"
	}
	if obj.DisplayName == "" {
		obj.DisplayName = "kcc-test-processor"
	}
	if obj.State == pb.Processor_STATE_UNSPECIFIED {
		obj.State = pb.Processor_ENABLED
	}
	if obj.ProcessEndpoint == "" {
		obj.ProcessEndpoint = fmt.Sprintf("https://%s-documentai.googleapis.com/v1/%s:process", name.Location, name.String())
	}
	if obj.CreateTime == nil {
		obj.CreateTime = timestamppb.New(time.Now())
	}
	if len(obj.ProcessorVersionAliases) == 0 {
		obj.ProcessorVersionAliases = []*pb.ProcessorVersionAlias{
			{
				Alias:            fmt.Sprintf("%s/processorVersions/pretrained", name.String()),
				ProcessorVersion: fmt.Sprintf("%s/processorVersions/pretrained-ocr-v1.0-2020-09-23", name.String()),
			},
			{
				Alias:            fmt.Sprintf("%s/processorVersions/pretrained-next", name.String()),
				ProcessorVersion: fmt.Sprintf("%s/processorVersions/pretrained-ocr-v1.1-2022-09-12", name.String()),
			},
			{
				Alias:            fmt.Sprintf("%s/processorVersions/rc", name.String()),
				ProcessorVersion: fmt.Sprintf("%s/processorVersions/pretrained-ocr-v1.2-2022-11-10", name.String()),
			},
			{
				Alias:            fmt.Sprintf("%s/processorVersions/stable", name.String()),
				ProcessorVersion: fmt.Sprintf("%s/processorVersions/pretrained-ocr-v2.1-2024-08-07", name.String()),
			},
		}
	}
	if obj.DefaultProcessorVersion == "" {
		obj.DefaultProcessorVersion = fmt.Sprintf("%s/processorVersions/pretrained-ocr-v1.2-2022-11-10", name.String())
	}
	obj.SatisfiesPzs = true
}

func (s *DocumentProcessorV1) GetProcessor(ctx context.Context, req *pb.GetProcessorRequest) (*pb.Processor, error) {
	name, err := s.ParseProcessorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Processor{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Pre-created processors for tests that do not manage processor lifecycle directly
			if name.ProcessorName == "2591e9369e934b5" || name.ProcessorName == "7f8f177e3b9cc6d9" {
				obj = &pb.Processor{}
				s.populateDefaultsForProcessor(obj, name)
				if err := s.storage.Create(ctx, fqn, obj); err != nil {
					return nil, err
				}
				return obj, nil
			}
		}
		return nil, err
	}

	return obj, nil
}

func (s *DocumentProcessorV1) CreateProcessor(ctx context.Context, req *pb.CreateProcessorRequest) (*pb.Processor, error) {
	var fqn string
	var name *ProcessorName
	if req.GetProcessor().GetName() != "" {
		var err error
		name, err = s.ParseProcessorName(req.GetProcessor().GetName())
		if err != nil {
			return nil, err
		}
		fqn = name.String()
	} else {
		// In GCP, processor IDs are generated hex strings
		processorID := fmt.Sprintf("%x", time.Now().UnixNano())
		tokens := strings.Split(req.GetParent(), "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
			name = &ProcessorName{
				Project:       tokens[1],
				Location:      tokens[3],
				ProcessorName: processorID,
			}
			fqn = name.String()
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
		}
	}

	obj := proto.Clone(req.GetProcessor()).(*pb.Processor)
	s.populateDefaultsForProcessor(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DocumentProcessorV1) SetDefaultProcessorVersion(ctx context.Context, req *pb.SetDefaultProcessorVersionRequest) (*longrunningpb.Operation, error) {
	name, err := s.ParseProcessorName(req.GetProcessor())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Processor{}
	exists := true
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			exists = false
			obj = &pb.Processor{}
			s.populateDefaultsForProcessor(obj, name)
		} else {
			return nil, err
		}
	}

	targetVersion := req.GetDefaultProcessorVersion()
	for _, alias := range obj.ProcessorVersionAliases {
		if alias.Alias == targetVersion {
			targetVersion = alias.ProcessorVersion
			break
		}
	}
	obj.DefaultProcessorVersion = targetVersion

	if exists {
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	op := &pb.SetDefaultProcessorVersionMetadata{
		CommonMetadata: &pb.CommonOperationMetadata{
			State:      pb.CommonOperationMetadata_RUNNING,
			CreateTime: timestamppb.New(time.Now()),
			UpdateTime: timestamppb.New(time.Now()),
			Resource:   obj.DefaultProcessorVersion,
		},
	}
	resp := &pb.SetDefaultProcessorVersionResponse{}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project, name.Location)
	return s.operations.DoneLRO(ctx, prefix, op, resp)
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
