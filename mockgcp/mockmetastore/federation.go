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
// proto.service: google.cloud.metastore.v1.DataprocMetastoreFederation
// proto.message: google.cloud.metastore.v1.Federation

package mockmetastore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/anypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/metastore/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *DataprocMetastoreFederationV1) GetFederation(ctx context.Context, req *pb.GetFederationRequest) (*pb.Federation, error) {
	name, err := s.parseFederationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Federation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Federation %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataprocMetastoreFederationV1) CreateFederation(ctx context.Context, req *pb.CreateFederationRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/federations/" + req.FederationId
	name, err := s.parseFederationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Federation).(*pb.Federation)

	obj.Name = fqn

	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Federation_CREATING
	obj.EndpointUri = "federation-endpoint-" + name.Name
	obj.StateMessage = "The federation is being created"
	obj.Version = "3.1.2"

	if obj.BackendMetastores == nil {
		obj.BackendMetastores = map[int32]*pb.BackendMetastore{
			1: {
				Name:          "projects/" + name.Project.ID + "/locations/" + name.Location + "/services/" + name.Name,
				MetastoreType: pb.BackendMetastore_DATAPROC_METASTORE,
			},
		}
	}

	if obj.Uid == "" {
		obj.Uid = name.Name
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}

	lro, err := s.operations.NewLRO(ctx)
	lro.Done = false
	if err != nil {
		return nil, err
	}
	lro.Metadata, err = anypb.New(lroMetadata)
	if err != nil {
		return nil, err
	}
	lro.Metadata.TypeUrl = "type.googleapis.com/google.cloud.metastore.v1.OperationMetadata"
	lro.Name = fmt.Sprintf("projects/%s/locations/%s/operations/%s", name.Project.ID, name.Location, strings.Split(lro.Name, "/")[len(strings.Split(lro.Name, "/"))-1])

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		updated, err := s.updateFederation(ctx, fqn, func(obj *pb.Federation) {
			obj.State = pb.Federation_ACTIVE
			obj.StateMessage = "The federation is ready to use"
		})
		if err != nil {
			return nil, err
		}
		return updated, err
	})
}

func (s *DataprocMetastoreFederationV1) UpdateFederation(ctx context.Context, req *pb.UpdateFederationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseFederationName(req.GetFederation().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	update := func(obj *pb.Federation) {
		if req.UpdateMask == nil || len(req.UpdateMask.Paths) == 0 {
			return
		}
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "labels":
				obj.Labels = req.Federation.Labels
			case "version":
				obj.Version = req.Federation.Version
			case "backend_metastores":
				obj.BackendMetastores = req.Federation.BackendMetastores
			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	updated, err := s.updateFederation(ctx, fqn, update)
	if err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return updated, nil
	})
}

func (s *DataprocMetastoreFederationV1) DeleteFederation(ctx context.Context, req *pb.DeleteFederationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseFederationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Federation{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *DataprocMetastoreFederationV1) updateFederation(ctx context.Context, fqn string, update func(obj *pb.Federation)) (*pb.Federation, error) {
	obj := &pb.Federation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	update(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type federationName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *federationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/federations/%s", n.Project.ID, n.Location, n.Name)
}

func (s *DataprocMetastoreFederationV1) parseFederationName(name string) (*federationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "federations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &federationName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) buildFederationName(projectName, region, federation string) (*federationName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &federationName{
		Project:  project,
		Location: region,
		Name:     federation,
	}, nil
}
