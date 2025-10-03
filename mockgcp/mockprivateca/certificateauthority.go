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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/security/privateca/v1"
)

func (s *PrivateCAV1) GetCertificateAuthority(ctx context.Context, req *pb.GetCertificateAuthorityRequest) (*pb.CertificateAuthority, error) {
	name, err := s.parseCertificateAuthorityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CertificateAuthority{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *PrivateCAV1) CreateCertificateAuthority(ctx context.Context, req *pb.CreateCertificateAuthorityRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificateAuthorities/" + req.CertificateAuthorityId
	name, err := s.parseCertificateAuthorityName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()

	obj := proto.Clone(req.CertificateAuthority).(*pb.CertificateAuthority)
	obj.Name = fqn

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

func (s *PrivateCAV1) DeleteCertificateAuthority(ctx context.Context, req *pb.DeleteCertificateAuthorityRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateAuthorityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	oldObj := &pb.CertificateAuthority{}
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

type certificateAuthorityName struct {
	caPoolName
	CertificateAuthorityID string
}

func (n *certificateAuthorityName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/caPools/" + n.CAPoolName + "/certificateAuthorities" + n.CertificateAuthorityID
}

// parseCertificateAuthorityName parses a string into a certificateAuthorityName.
// The expected form is projects/<projectID>/locations/<region>/caPools/<capoolName>/certificateAuthorities/<caName>
func (s *MockService) parseCertificateAuthorityName(name string) (*certificateAuthorityName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "caPools" && tokens[6] == "certificateAuthorities" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &certificateAuthorityName{
			caPoolName: caPoolName{
				Project:    project,
				Location:   tokens[3],
				CAPoolName: tokens[5],
			},
			CertificateAuthorityID: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
