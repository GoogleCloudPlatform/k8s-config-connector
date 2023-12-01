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

package mockcompute

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type GlobalSSLCertificatesV1 struct {
	*MockService
	pb.UnimplementedSslCertificatesServer
}

func (s *GlobalSSLCertificatesV1) Get(ctx context.Context, req *pb.GetSslCertificateRequest) (*pb.SslCertificate, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/sslCertificates/" + req.GetSslCertificate()
	name, err := s.parseGlobalSslCertificateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SslCertificate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "sslCertificate %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading sslCertificate: %v", err)
		}
	}

	return obj, nil
}

func (s *GlobalSSLCertificatesV1) Insert(ctx context.Context, req *pb.InsertSslCertificateRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/sslCertificates/" + req.GetSslCertificateResource().GetName()
	name, err := s.parseGlobalSslCertificateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetSslCertificateResource()).(*pb.SslCertificate)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#sslCertificate")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating sslCertificate: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *GlobalSSLCertificatesV1) Delete(ctx context.Context, req *pb.DeleteSslCertificateRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/sslCertificates/" + req.GetSslCertificate()
	name, err := s.parseGlobalSslCertificateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.SslCertificate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "sslCertificate %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting sslCertificate: %v", err)
		}
	}

	return s.newLRO(ctx, name.Project.ID)
}

type globalSSLCertificateName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalSSLCertificateName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/sslCertificates/" + n.Name
}

// parseGlobalSslCertificateName parses a string into a globalSslCertificateName.
// The expected form is `projects/*/regions/*/sslcertificate/*`.
func (s *MockService) parseGlobalSslCertificateName(name string) (*globalSSLCertificateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "sslCertificates" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalSSLCertificateName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
