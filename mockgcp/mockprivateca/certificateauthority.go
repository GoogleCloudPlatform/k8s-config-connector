// Copyright 2022 Google LLC
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

package mockprivateca

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/security/privateca/v1"
)

type PrivateCAV1 struct {
	*MockService
	pb.UnimplementedCertificateAuthorityServiceServer
}

func (s *PrivateCAV1) GetCaPool(ctx context.Context, req *pb.GetCaPoolRequest) (*pb.CaPool, error) {
	name, err := s.parseCAPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CaPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *PrivateCAV1) CreateCaPool(ctx context.Context, req *pb.CreateCaPoolRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/caPools/" + req.CaPoolId
	name, err := s.parseCAPoolName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()

	obj := proto.Clone(req.CaPool).(*pb.CaPool)
	obj.Name = fqn

	// service seems to remove "zero" values
	baseKeyUsage := obj.GetIssuancePolicy().GetBaselineValues().GetKeyUsage().GetBaseKeyUsage()
	if baseKeyUsage != nil {
		if proto.Equal(baseKeyUsage, &pb.KeyUsage_KeyUsageOptions{}) {
			obj.IssuancePolicy.BaselineValues.KeyUsage.BaseKeyUsage = nil
		}
	}

	extendedKeyUsage := obj.GetIssuancePolicy().GetBaselineValues().GetKeyUsage().GetExtendedKeyUsage()
	if extendedKeyUsage != nil {
		if proto.Equal(extendedKeyUsage, &pb.KeyUsage_ExtendedKeyUsageOptions{}) {
			obj.IssuancePolicy.BaselineValues.KeyUsage.ExtendedKeyUsage = nil
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "create",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *PrivateCAV1) UpdateCaPool(ctx context.Context, req *pb.UpdateCaPoolRequest) (*longrunning.Operation, error) {
	reqName := req.GetCaPool().GetName()

	name, err := s.parseCAPoolName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()
	obj := &pb.CaPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "issuancePolicy":
			obj.IssuancePolicy = req.GetCaPool().GetIssuancePolicy()
		case "publishingOptions":
			obj.PublishingOptions = req.GetCaPool().GetPublishingOptions()
		case "labels":
			obj.Labels = req.GetCaPool().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "update",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *PrivateCAV1) DeleteCaPool(ctx context.Context, req *pb.DeleteCaPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseCAPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	oldObj := &pb.CaPool{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "delete",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type caPoolName struct {
	Project    *projects.ProjectData
	Location   string
	CAPoolName string
}

func (n *caPoolName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/caPools/" + n.CAPoolName
}

// parseCAPoolName parses a string into a caPoolName.
// The expected form is projects/<projectID>/locations/<region>/caPools/<capoolName>
func (s *MockService) parseCAPoolName(name string) (*caPoolName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "caPools" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &caPoolName{
			Project:    project,
			Location:   tokens[3],
			CAPoolName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
