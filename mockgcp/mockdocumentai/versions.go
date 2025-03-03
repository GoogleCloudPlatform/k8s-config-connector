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
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/documentai/v1"
)

func (s *DocumentProcessorV1) GetProcessorVersion(ctx context.Context, req *pb.GetProcessorVersionRequest) (*pb.ProcessorVersion, error) {
	name, err := s.parseProcessorVersionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ProcessorVersion{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DocumentProcessorV1) DeleteProcessorVersion(ctx context.Context, req *pb.DeleteProcessorVersionRequest) (*longrunning.Operation, error) {
	name, err := s.parseProcessorVersionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.ProcessorVersion{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *DocumentProcessorV1) TrainProcessorVersion(ctx context.Context, req *pb.TrainProcessorVersionRequest) (*longrunning.Operation, error) {
	name, err := s.parseProcessorVersionName(req.GetProcessorVersion().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	processorVersion := proto.Clone(req.GetProcessorVersion()).(*pb.ProcessorVersion)
	processorVersion.Name = fqn

	if err := s.storage.Create(ctx, fqn, processorVersion); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

type processorVersionName struct {
	processorName        *ProcessorName
	ProcessorVersionName string
}

func (n *processorVersionName) String() string {
	return n.processorName.String() + "/processorVersions/" + n.ProcessorVersionName
}

// parseProcessorVersionName parses a string into a processorVersionName.
// The expected form is projects/{project}/locations/{location}/processors/{processor}/processorVersions/{processorVersion}
func (s *DocumentProcessorV1) parseProcessorVersionName(name string) (*processorVersionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "processors" && tokens[6] == "processorVersions" {
		processorName, err := s.ParseProcessorName(strings.Join(tokens[0:6], "/"))
		if err != nil {
			return nil, err
		}
		name := &processorVersionName{
			processorName:        processorName,
			ProcessorVersionName: tokens[7],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
