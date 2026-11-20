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
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ImagesV1 struct {
	*MockService
	pb.UnimplementedImagesServer

	pendingLabels map[string]*pb.Image
}

func (s *ImagesV1) GetFromFamily(ctx context.Context, req *pb.GetFromFamilyImageRequest) (*pb.Image, error) {
	obj := &pb.Image{}

	family := req.GetFamily()
	if family == "" {
		family = "debian-11"
	}
	imageName := family + "-bullseye-v20231010"
	if family == "debian-12" {
		imageName = "debian-12-bookworm-v20261023"
	}

	obj.Kind = PtrTo("compute#image")
	obj.Name = PtrTo(imageName)
	obj.Description = PtrTo(fmt.Sprintf("Debian, Debian GNU/Linux, %s, amd64", family))
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, "projects/debian-cloud/global/images/"+imageName))
	obj.Family = PtrTo(family)
	obj.Status = PtrTo("UP")

	return obj, nil
}

func (s *ImagesV1) Get(ctx context.Context, req *pb.GetImageRequest) (*pb.Image, error) {
	// Get from family
	if req.GetProject() == "debian-cloud" && (req.GetImage() == "debian-11" || req.GetImage() == "debian-12") {
		fqn := fmt.Sprintf("projects/debian-cloud/global/images/%s", req.GetImage())
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

	obj, err = applyImagePatch(obj, req.GetImageResource())
	if err != nil {
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

// patchableImageFields are the Image fields that images.patch can change.
var patchableImageFields = map[protoreflect.Name]bool{
	"architecture":        true,
	"deletion_protection": true,
	"description":         true,
	"family":              true,
}

// applyImagePatch applies an images.patch body to a copy of stored, the way real GCE does:
//   - A body that changes any field other than patchableImageFields is rejected with
//     "Updating '<field>' is not supported.". Fields sent with their stored value are accepted.
//   - A body that contains labels is rejected even when the labels are unchanged, because labels
//     can only be changed with images.setLabels.
//
// GCE reports a changed field before it reports labels, so the checks run in that order.
func applyImagePatch(stored, patch *pb.Image) (*pb.Image, error) {
	updated := proto.CloneOf(stored)
	src, dst, old := patch.ProtoReflect(), updated.ProtoReflect(), stored.ProtoReflect()
	fields := src.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if !src.Has(fd) || fd.Name() == "labels" {
			continue
		}
		if patchableImageFields[fd.Name()] {
			dst.Set(fd, src.Get(fd))
			continue
		}
		if !src.Get(fd).Equal(old.Get(fd)) {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'resource.%s': '%v'. Updating '%s' is not supported.",
				fd.JSONName(), src.Get(fd).Interface(), fd.JSONName())
		}
	}
	if len(patch.GetLabels()) != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'resource.labels': ''. Updating 'labels' is not supported.")
	}
	return updated, nil
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

	// Like GCE, refuse a stale fingerprint so that a lost update is caught.
	if req.GetGlobalSetLabelsRequestResource().GetLabelFingerprint() != obj.GetLabelFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "Labels fingerprint either invalid or resource labels have changed")
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
