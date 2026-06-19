// Copyright 2026 Google LLC
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/security/privateca/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type certificateTemplateName struct {
	Project                 *projects.ProjectData
	Location                string
	CertificateTemplateName string
}

func (n *certificateTemplateName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/certificateTemplates/" + n.CertificateTemplateName
}

func (s *MockService) parseCertificateTemplateName(name string) (*certificateTemplateName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "certificateTemplates" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &certificateTemplateName{
			Project:                 project,
			Location:                tokens[3],
			CertificateTemplateName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *PrivateCAV1) GetCertificateTemplate(ctx context.Context, req *pb.GetCertificateTemplateRequest) (*pb.CertificateTemplate, error) {
	name, err := s.parseCertificateTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CertificateTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *PrivateCAV1) CreateCertificateTemplate(ctx context.Context, req *pb.CreateCertificateTemplateRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificateTemplates/" + req.CertificateTemplateId
	name, err := s.parseCertificateTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()

	obj := proto.Clone(req.CertificateTemplate).(*pb.CertificateTemplate)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

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
	return s.operations.StartLROWithDone(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

func (s *PrivateCAV1) UpdateCertificateTemplate(ctx context.Context, req *pb.UpdateCertificateTemplateRequest) (*longrunning.Operation, error) {
	reqName := req.GetCertificateTemplate().GetName()

	name, err := s.parseCertificateTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.CertificateTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(obj).(*pb.CertificateTemplate)

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(updated, req.CertificateTemplate, req.UpdateMask.Paths); err != nil {
			return nil, err
		}
	} else {
		updated = proto.Clone(req.CertificateTemplate).(*pb.CertificateTemplate)
	}

	updated.Name = fqn
	updated.CreateTime = obj.CreateTime
	updated.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
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
	return s.operations.StartLROWithDone(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(now)
		return updated, nil
	})
}

func (s *PrivateCAV1) DeleteCertificateTemplate(ctx context.Context, req *pb.DeleteCertificateTemplateRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.CertificateTemplate{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
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
	return s.operations.StartLROWithDone(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(now)
		return &emptypb.Empty{}, nil
	})
}

func (s *PrivateCAV1) ListCertificateTemplates(ctx context.Context, req *pb.ListCertificateTemplatesRequest) (*pb.ListCertificateTemplatesResponse, error) {
	objs := []*pb.CertificateTemplate{}
	kind := (&pb.CertificateTemplate{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		objs = append(objs, obj.(*pb.CertificateTemplate))
		return nil
	}); err != nil {
		return nil, err
	}

	// Filter by parent (projects/*/locations/*)
	// Parents: projects/<projectID>/locations/<region>
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectID := tokens[1]
		location := tokens[3]
		filtered := []*pb.CertificateTemplate{}
		for _, obj := range objs {
			name, err := s.parseCertificateTemplateName(obj.Name)
			if err == nil && name.Project.ID == projectID && name.Location == location {
				filtered = append(filtered, obj)
			}
		}
		objs = filtered
	}

	return &pb.ListCertificateTemplatesResponse{
		CertificateTemplates: objs,
	}, nil
}
