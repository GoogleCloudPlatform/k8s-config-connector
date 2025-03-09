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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ImagesV1 struct {
	*MockService
	pb.UnimplementedImagesServer
}

func (s *ImagesV1) GetFromFamily(ctx context.Context, req *pb.GetFromFamilyImageRequest) (*pb.Image, error) {
	ret := &pb.Image{}

	findKind := (&pb.Image{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{}, func(obj proto.Message) error {
		image := obj.(*pb.Image)
		if req.GetFamily() != image.GetFamily() {
			return nil
		}

		name, err := s.parseImageSelfLink(image.GetSelfLink())
		if err != nil {
			return err
		}
		if req.GetProject() != "" && req.GetProject() != name.Project.ID {
			return nil
		}
		ret = image
		return nil
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *ImagesV1) Get(ctx context.Context, req *pb.GetImageRequest) (*pb.Image, error) {
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

	if sourceImage := req.GetImageResource().GetSourceImage(); sourceImage != "" {
		sourceImageName, err := s.parseImageSelfLink(sourceImage)
		if err != nil {
			return nil, fmt.Errorf("invalid source image %q", sourceImage)
		}
		source := &pb.Image{}
		if err := s.storage.Get(ctx, sourceImageName.String(), source); err != nil {
			if status.Code(err) == codes.NotFound {
				return nil, fmt.Errorf("source image %q not found", sourceImage)
			} else {
				return nil, fmt.Errorf("error getting source image %q", sourceImage)
			}
		}
		obj.Architecture = source.Architecture
		obj.ArchiveSizeBytes = source.ArchiveSizeBytes
		obj.DiskSizeGb = source.DiskSizeGb
		obj.EnableConfidentialCompute = source.EnableConfidentialCompute
		obj.GuestOsFeatures = source.GuestOsFeatures
		obj.Licenses = source.Licenses
		obj.LicenseCodes = source.LicenseCodes
		obj.SourceImage = PtrTo(buildComputeSelfLink(ctx, sourceImageName.String()))
		obj.SourceImageId = PtrTo(fmt.Sprintf("%d", source.GetId()))
		obj.SourceType = source.SourceType
	}
	obj.Status = PtrTo("READY")

	obj.StorageLocations = []string{"us"}

	if obj.DiskSizeGb == nil {
		obj.DiskSizeGb = PtrTo(int64(500))
	}

	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.Labels))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	op := &pb.Operation{
		OperationType: PtrTo("patch"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	// Returns an LRO, but the LRO is already done
	op := &pb.Operation{
		OperationType: PtrTo("setLabels"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
		EndTime:       PtrTo(s.nowString()),
		Progress:      PtrTo(int32(100)),
		Status:        PtrTo(pb.Operation_DONE),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
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

	op := &pb.Operation{
		OperationType: PtrTo("delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
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

// parseImageName parses a string into an imageName.
// The expected form is `projects/*/global/images/*`.
func (s *MockService) parseImageSelfLink(selfLink string) (*ImageName, error) {
	fqn := strings.TrimPrefix(selfLink, "https://www.googleapis.com/compute/v1/")
	tokens := strings.Split(fqn, "/")
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "images" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &ImageName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	return nil, fmt.Errorf("invalid image self link %q", selfLink)
}
