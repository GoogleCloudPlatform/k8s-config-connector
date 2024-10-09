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

package mockcompute

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type RegionalSSLCertificatesV1 struct {
	*MockService
	pb.UnimplementedRegionSslCertificatesServer
}

func (s *RegionalSSLCertificatesV1) Get(ctx context.Context, req *pb.GetRegionSslCertificateRequest) (*pb.SslCertificate, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.Region + "/sslCertificates/" + req.GetSslCertificate()
	name, err := s.parseRegionalSslCertificateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SslCertificate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RegionalSSLCertificatesV1) Insert(ctx context.Context, req *pb.InsertRegionSslCertificateRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/sslCertificates/" + req.GetSslCertificateResource().GetName()
	name, err := s.parseRegionalSslCertificateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetSslCertificateResource()).(*pb.SslCertificate)
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#sslCertificate")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionalSSLCertificatesV1) Delete(ctx context.Context, req *pb.DeleteRegionSslCertificateRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/sslCertificates/" + req.GetSslCertificate()
	name, err := s.parseRegionalSslCertificateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.SslCertificate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type regionalSSLCertificateName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalSSLCertificateName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/sslCertificates/" + n.Name
}

// parseRegionalSslCertificateName parses a string into a regionalSSLCertificateName.
// The expected form is `projects/*/regions/*/sslcertificate/*`.
func (s *MockService) parseRegionalSslCertificateName(name string) (*regionalSSLCertificateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "sslCertificates" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalSSLCertificateName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
