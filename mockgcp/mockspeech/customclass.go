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

// +tool:mockgcp-support
// proto.service: google.cloud.speech.v2.Speech
// proto.message: google.cloud.speech.v2.CustomClass

package mockspeech

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/speech/v2"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *SpeechV2) GetCustomClass(ctx context.Context, req *pb.GetCustomClassRequest) (*pb.CustomClass, error) {
	name, err := s.parseCustomClassName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.CustomClass{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Unable to find CustomClass %s from project %d.", name.CustomClassID, name.Project.Number)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SpeechV2) CreateCustomClass(ctx context.Context, req *pb.CreateCustomClassRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/customClasses/%s", req.GetParent(), req.GetCustomClassId())
	name, err := s.parseCustomClassName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetCustomClass()).(*pb.CustomClass)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = uuid.New().String()
	obj.State = pb.CustomClass_ACTIVE // Assume immediate activation for mock
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Method:          "google.cloud.speech.v2.Speech.CreateCustomClass",
		ProgressPercent: 100,
	}

	// change project ID to project number
	metadata.Resource = strings.Replace(obj.GetName(), "projects/"+name.Project.ID+"/locations/"+name.Location+"/customClasses/"+name.CustomClassID, "projects/"+strconv.FormatInt(name.Project.Number, 10)+"/locations/"+name.Location+"/customClasses/"+name.CustomClassID, 1)

	req.Parent = strings.Replace(req.GetParent(), "projects/"+name.Project.ID+"/locations/"+name.Location, "projects/"+strconv.FormatInt(name.Project.Number, 10)+"/locations/"+name.Location, 1)
	metadata.Request = &pb.OperationMetadata_CreateCustomClassRequest{
		CreateCustomClassRequest: req,
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	return s.operations.DoneLRO(ctx, prefix, metadata, obj)
}

func (s *SpeechV2) UpdateCustomClass(ctx context.Context, req *pb.UpdateCustomClassRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCustomClassName(req.GetCustomClass().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.CustomClass{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for i, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetCustomClass().GetDisplayName()

			// HACK: to make the field mask valid when returning
			// original error:
			// proto: google.protobuf.FieldMask.paths contains irreversible value "displayName"
			req.UpdateMask.Paths[i] = "display_name"
		case "items":
			obj.Items = req.GetCustomClass().GetItems()
		case "annotations":
			obj.Annotations = req.GetCustomClass().GetAnnotations()
		// Add other updatable fields if necessary
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid for CustomClass update", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Resource:        obj.GetName(),
		Method:          "google.cloud.speech.v2.Speech.UpdateCustomClass",
		ProgressPercent: 100,
	}

	// change project ID to project number
	req.CustomClass.Name = strings.Replace(req.CustomClass.GetName(), "projects/"+name.Project.ID+"/locations/"+name.Location, "projects/"+strconv.FormatInt(name.Project.Number, 10)+"/locations/"+name.Location, 1)
	metadata.Request = &pb.OperationMetadata_UpdateCustomClassRequest{
		UpdateCustomClassRequest: req,
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	return s.operations.DoneLRO(ctx, prefix, metadata, obj)
}

func (s *SpeechV2) DeleteCustomClass(ctx context.Context, req *pb.DeleteCustomClassRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCustomClassName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	// change project ID to project number
	req.Name = strings.Replace(req.GetName(), "projects/"+name.Project.ID+"/locations/"+name.Location, "projects/"+strconv.FormatInt(name.Project.Number, 10)+"/locations/"+name.Location, 1)

	obj := &pb.CustomClass{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Allow deleting a non-existent resource if allow_missing is true
			if req.GetAllowMissing() {
				// Return a completed LRO indicating success (no-op)
				metadata := &pb.OperationMetadata{
					CreateTime:      timestamppb.New(now),
					UpdateTime:      timestamppb.New(now),
					Resource:        fqn,
					Method:          "google.cloud.speech.v2.Speech.DeleteCustomClass",
					ProgressPercent: 100,
					Request: &pb.OperationMetadata_DeleteCustomClassRequest{
						DeleteCustomClassRequest: req,
					},
				}
				// The response type for delete is CustomClass according to the LRO info,
				// but typically delete returns the deleted resource or empty.
				// Let's return the (non-existent) resource name in metadata and Empty in response.
				// Or maybe return the (non-existent) obj itself in the response as per LRO info? Let's try obj.
				// Based on testing, delete LRO response is the resource itself before deletion (or with deleted state).
				// Since it doesn't exist, returning a placeholder might be best or erroring?
				// Let's return a placeholder indicating deletion.
				deletedPlaceholder := &pb.CustomClass{Name: fqn, State: pb.CustomClass_DELETED}
				return s.operations.DoneLRO(ctx, prefix, metadata, deletedPlaceholder)
			}
			return nil, status.Errorf(codes.NotFound, "customClass %q not found", fqn)
		}
		return nil, err
	}

	// Validate Etag if provided
	if req.GetEtag() != "" && req.GetEtag() != obj.Etag {
		return nil, status.Errorf(codes.Aborted, "etag mismatch")
	}

	obj.State = pb.CustomClass_DELETED
	obj.DeleteTime = timestamppb.New(now)
	obj.ExpireTime = timestamppb.New(now.Add(30 * 24 * time.Hour))

	if err := s.storage.Delete(ctx, fqn, &pb.CustomClass{}); err != nil {
		// Log or handle the error if removing the original fails after marking deleted
		// This indicates an inconsistent state.
		fmt.Printf("Warning: failed to remove original customClass %q after marking as deleted: %v\n", fqn, err)
	}

	metadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Resource:        obj.GetName(),
		Method:          "google.cloud.speech.v2.Speech.DeleteCustomClass",
		ProgressPercent: 100,
		Request: &pb.OperationMetadata_DeleteCustomClassRequest{
			DeleteCustomClassRequest: req,
		},
	}

	return s.operations.DoneLRO(ctx, prefix, metadata, obj)
}

/*
func (s *SpeechV2) UndeleteCustomClass(ctx context.Context, req *pb.UndeleteCustomClassRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCustomClassName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedFQN := fqn + "@deleted"
	now := time.Now()

	obj := &pb.CustomClass{}
	if err := s.storage.Get(ctx, deletedFQN, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Check if it exists in the non-deleted state already
			activeObj := &pb.CustomClass{}
			if activeErr := s.storage.Get(ctx, fqn, activeObj); activeErr == nil {
				// Already active, treat as success? Or precondition failed?
				// Let's treat as success with the active object.
				metadata := &pb.OperationMetadata{
					CreateTime: timestamppb.New(now), // Or use activeObj.UpdateTime?
					UpdateTime: timestamppb.New(now),
					Resource:   activeObj.GetName(),
					Method:     "google.cloud.speech.v2.Speech.UndeleteCustomClass",
				}
				return s.operations.DoneLRO(ctx, name.String(), metadata, activeObj)
			}
			return nil, status.Errorf(codes.NotFound, "customClass %q not found (or not deleted)", fqn)
		}
		return nil, err
	}

	// Validate Etag if provided
	if req.GetEtag() != "" && req.GetEtag() != obj.Etag {
		return nil, status.Errorf(codes.Aborted, "etag mismatch")
	}

	// Check if the expiration time has passed
	if obj.ExpireTime != nil && now.After(obj.ExpireTime.AsTime()) {
		return nil, status.Errorf(codes.FailedPrecondition, "customClass %q has passed its expiration time and cannot be undeleted", fqn)
	}

	// Undelete: Change state back to ACTIVE, clear deletion timestamps
	obj.State = pb.CustomClass_ACTIVE
	obj.DeleteTime = nil
	obj.ExpireTime = nil
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

	// Move back from deleted state to active state
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		// This could happen if another undelete raced or if there's an internal issue
		return nil, status.Errorf(codes.Internal, "failed to restore customClass %q: %v", fqn, err)
	}
	// Remove from the "deleted" state
	if err := s.storage.Delete(ctx, deletedFQN, &pb.CustomClass{}); err != nil {
		// Log or handle the error if removing the deleted marker fails.
		fmt.Printf("Warning: failed to remove deleted marker for customClass %q: %v\n", deletedFQN, err)
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
		Resource:   obj.GetName(),
		Method:     "google.cloud.speech.v2.Speech.UndeleteCustomClass",
	}

	return s.operations.StartLRO(ctx, name.String(), metadata, func() (proto.Message, error) {
		metadata.ProgressPercent = 100
		// Return the object in its restored state
		return obj, nil
	})
}
*/

type customClassName struct {
	Project       *projects.ProjectData
	Location      string
	CustomClassID string
}

func (n *customClassName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/customClasses/%s", n.Project.Number, n.Location, n.CustomClassID)
}

// parseCustomClassName parses a string into a customClassName.
// The expected form is `projects/*/locations/*/customClasses/*`.
func (s *MockService) parseCustomClassName(name string) (*customClassName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "customClasses" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			// Handle project not found potentially, or bubble up the error
			return nil, status.Errorf(codes.NotFound, "project %q not found: %v", tokens[1], err)
		}

		nameObj := &customClassName{
			Project:       project,
			Location:      tokens[3],
			CustomClassID: tokens[5],
		}

		return nameObj, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in the expected format projects/*/locations/*/customClasses/*", name)
}
