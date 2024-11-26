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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ImagesV1 struct {
	*MockService
	pb.UnimplementedImagesServer
}

func (s *ImagesV1) GetFromFamily(ctx context.Context, req *pb.GetFromFamilyImageRequest) (*pb.Image, error) {
	obj := &pb.Image{}

	// Details from gcloud compute images describe-from-family debian-11 --project debian-cloud --log-http
	obj.Kind = PtrTo("compute#image")
	obj.Name = PtrTo("debian-11-bullseye-v20231010")
	obj.Description = PtrTo("Debian, Debian GNU/Linux, 11 (bullseye), amd64 built on 20231010")
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, "projects/debian-cloud/global/images/debian-11-bullseye-v20231010"))
	obj.Family = PtrTo("debian-11")
	obj.Status = PtrTo("UP")

	return obj, nil
}

func (s *ImagesV1) Get(ctx context.Context, req *pb.GetImageRequest) (*pb.Image, error) {
	// Get from family
	if req.GetProject() == "debian-cloud" && req.GetImage() == "debian-11" {
		return nil, status.Errorf(codes.NotFound, "image not found")
	}

	name, err := s.parseImageName(req.GetProject(), req.GetImage())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Image{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ImagesV1) Insert(ctx context.Context, req *pb.InsertImageRequest) (*pb.Operation, error) {
	name, err := s.parseImageName(req.GetProject(), req.GetImageResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetImageResource()).(*pb.Image)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#image")

	if obj.DiskSizeGb == nil {
		obj.DiskSizeGb = PtrTo(int64(500))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ImagesV1) Patch(ctx context.Context, req *pb.PatchImageRequest) (*pb.Operation, error) {
	name, err := s.parseImageName(req.GetProject(), req.GetImage())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Image{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if obj.DiskSizeGb == nil {
		obj.DiskSizeGb = PtrTo(int64(500))
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ImagesV1) SetLabels(ctx context.Context, req *pb.SetLabelsImageRequest) (*pb.Operation, error) {
	name, err := s.parseImageName(req.GetProject(), req.GetResource())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Image{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetGlobalSetLabelsRequestResource().GetLabels()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *ImagesV1) Delete(ctx context.Context, req *pb.DeleteImageRequest) (*pb.Operation, error) {
	name, err := s.parseImageName(req.GetProject(), req.GetImage())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Image{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type ImageName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *ImageName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/images/" + n.Name
}

// parseImageName parses a string into an imageName.
// The expected form is `projects/*/global/images/*`.
func (s *MockService) parseImageName(projectName, name string) (*ImageName, error) {
	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &ImageName{
		Project: project,
		Name:    name,
	}, nil
}
