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
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type RegionalServiceAttachmentV1 struct {
	*MockService
	pb.UnimplementedServiceAttachmentsServer
}

func (s *RegionalServiceAttachmentV1) Get(ctx context.Context, req *pb.GetServiceAttachmentRequest) (*pb.ServiceAttachment, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/serviceAttachments/" + req.GetServiceAttachment()
	name, err := s.parseRegionalServiceAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ServiceAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RegionalServiceAttachmentV1) Insert(ctx context.Context, req *pb.InsertServiceAttachmentRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/serviceAttachments/" + req.GetServiceAttachmentResource().GetName()
	name, err := s.parseRegionalServiceAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetServiceAttachmentResource()).(*pb.ServiceAttachment)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#serviceAttachment")
	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
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
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

// Updates a ServiceAttachment resource in the specified project using the data included in the request.
// This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (s *RegionalServiceAttachmentV1) Patch(ctx context.Context, req *pb.PatchServiceAttachmentRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/serviceAttachments/" + req.GetServiceAttachment()
	name, err := s.parseRegionalServiceAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ServiceAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := mergeProtos(obj.ProtoReflect(), req.GetServiceAttachmentResource().ProtoReflect()); err != nil {
		return nil, err
	}
	obj.Fingerprint = nil
	obj.Fingerprint = PtrTo(computeFingerprint(obj))

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}

	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionalServiceAttachmentV1) Delete(ctx context.Context, req *pb.DeleteServiceAttachmentRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/serviceAttachments/" + req.GetServiceAttachment()
	name, err := s.parseRegionalServiceAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ServiceAttachment{}
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

type regionalServiceAttachmentName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalServiceAttachmentName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/serviceAttachments/" + n.Name
}

// parseRegionalServiceAttachmentName parses a string into a serviceattachmentName.
// The expected form is `projects/*/regions/*/serviceattachment/*`.
func (s *MockService) parseRegionalServiceAttachmentName(name string) (*regionalServiceAttachmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "serviceAttachments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalServiceAttachmentName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

// mergeProtos implements the patch/update semantics of GCP updates.
func mergeProtos(dst protoreflect.Message, src protoreflect.Message) error {
	var errs []error
	src.Range(func(fd protoreflect.FieldDescriptor, srcVal protoreflect.Value) bool {
		if fd.IsList() {
			// errs = append(errs, fmt.Errorf("unhandled list for field %v: %v", fd, fd.Kind()))
			switch fd.Kind() {
			case protoreflect.StringKind:
				// Replace []string
				dstVal := dst.Get(fd)
				klog.Warningf("replacing string list src=%v, dst=%v", srcVal, dstVal)
				dst.Set(fd, srcVal)
			case protoreflect.MessageKind:
				dstVal := dst.Get(fd)
				klog.Warningf("replacing message list src=%v, dst=%v", srcVal, dstVal)
				dst.Set(fd, srcVal)

			default:
				errs = append(errs, fmt.Errorf("unhandled kind for list field %v: %v", fd, fd.Kind()))
			}
		} else if fd.IsMap() {
			errs = append(errs, fmt.Errorf("unhandled list for field %v: %v", fd, fd.Kind()))
		} else {
			switch fd.Kind() {
			case protoreflect.StringKind:
				dst.Set(fd, srcVal)
			default:
				errs = append(errs, fmt.Errorf("unhandled kind for field %v: %v", fd, fd.Kind()))
			}
		}

		return true
	})
	return errors.Join(errs...)
}
