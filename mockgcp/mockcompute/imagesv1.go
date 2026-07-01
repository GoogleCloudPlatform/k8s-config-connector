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

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ImagesV1 struct {
	*MockService
	pb.UnimplementedImagesServer

	pendingLabels map[string]*pb.Image
}

func (s *ImagesV1) GetFromFamily(ctx context.Context, req *pb.GetFromFamilyImageRequest) (*pb.Image, error) {
	obj := &pb.Image{}

	// Details from gcloud compute images describe-from-family debian-11 --project debian-cloud --log-http
	obj.Kind = PtrTo("compute#image")
	obj.Name = PtrTo("debian-11-bullseye-v20231010")
	obj.Description = PtrTo("Debian, Debian GNU/Linux, 11 (bullseye), amd64 built on 20231010")
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, "projects/debian-cloud/global/images/debian-11-bullseye-v20231010"))
	obj.Family = PtrTo("debian-11")
	obj.Status = PtrTo("UP")

	return obj, nil
}

func (s *ImagesV1) Get(ctx context.Context, req *pb.GetImageRequest) (*pb.Image, error) {
	// Get from family
	if req.GetProject() == "debian-cloud" && req.GetImage() == "debian-11" {
		fqn := "projects/debian-cloud/global/images/debian-11"
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	name, err := s.parseImageName(req.GetProject(), req.GetImage())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Image{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Eventually consistent labels support
	if s.pendingLabels != nil {
		if oldObj, exists := s.pendingLabels[fqn]; exists {
			delete(s.pendingLabels, fqn)
			return oldObj, nil
		}
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

	obj := proto.CloneOf(req.GetImageResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#image")
	obj.EnableConfidentialCompute = PtrTo(false)
	obj.SourceType = PtrTo("RAW")
	obj.Status = PtrTo("READY")
	obj.StorageLocations = []string{"us"}
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.GetLabels()))

	sourceDisk := obj.GetSourceDisk()
	if sourceDisk != "" {
		// Convert relative or absolute URL to relative path starting with "projects/"
		u := sourceDisk
		if idx := strings.Index(u, "projects/"); idx != -1 {
			u = u[idx:]
		}
		if diskName, err := s.parseZonalDiskName(u); err == nil {
			diskFQN := diskName.String()
			diskObj := &pb.Disk{}
			if err := s.storage.Get(ctx, diskFQN, diskObj); err == nil {
				// Disk exists! Let's get its ID and size
				obj.SourceDiskId = PtrTo(fmt.Sprintf("%d", diskObj.GetId()))
				if diskObj.SizeGb != nil {
					obj.DiskSizeGb = diskObj.SizeGb
				}
			} else {
				obj.SourceDiskId = PtrTo("1234567890123456789")
			}
		}
		if !strings.HasPrefix(sourceDisk, "http") {
			obj.SourceDisk = PtrTo(BuildComputeSelfLink(ctx, strings.TrimPrefix(sourceDisk, "/")))
		}
	}

	if obj.DiskSizeGb == nil {
		obj.DiskSizeGb = PtrTo(int64(500))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if obj.DiskSizeGb == nil {
		obj.DiskSizeGb = PtrTo(int64(500))
	}

	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.GetLabels()))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	if s.pendingLabels == nil {
		s.pendingLabels = make(map[string]*pb.Image)
	}
	s.pendingLabels[fqn] = proto.CloneOf(obj)

	obj.Labels = req.GetGlobalSetLabelsRequestResource().GetLabels()
	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.GetLabels()))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("setLabels"),
		User:          PtrTo("user@example.com"),
		Status:        PtrTo(pb.Operation_DONE),
		Progress:      PtrTo(int32(100)),
		EndTime:       PtrTo(s.nowString()),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
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
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
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
