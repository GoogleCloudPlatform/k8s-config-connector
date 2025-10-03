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
// proto.message: google.cloud.speech.v2.PhraseSet

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

func (s *SpeechV2) GetPhraseSet(ctx context.Context, req *pb.GetPhraseSetRequest) (*pb.PhraseSet, error) {
	name, err := s.parsePhraseSetName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.PhraseSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Adjusted error message based on typical Google API responses
			return nil, status.Errorf(codes.NotFound, "Unable to find PhraseSet %s from project %d.", name.PhraseSetID, name.Project.Number)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SpeechV2) CreatePhraseSet(ctx context.Context, req *pb.CreatePhraseSetRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/phraseSets/%s", req.GetParent(), req.GetPhraseSetId())
	name, err := s.parsePhraseSetName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetPhraseSet()).(*pb.PhraseSet)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = uuid.New().String()
	obj.State = pb.PhraseSet_ACTIVE // Assume immediate activation for mock
	obj.Etag = fields.ComputeWeakEtag(obj)

	// Validate boost values
	if obj.Boost < 0 || obj.Boost > 20 {
		return nil, status.Errorf(codes.InvalidArgument, "phrase set boost %v must be between 0 and 20", obj.Boost)
	}
	for _, phrase := range obj.Phrases {
		if phrase.Boost < 0 || phrase.Boost > 20 {
			return nil, status.Errorf(codes.InvalidArgument, "phrase boost %v for phrase '%s' must be between 0 and 20", phrase.Boost, phrase.Value)
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Method:          "google.cloud.speech.v2.Speech.CreatePhraseSet",
		ProgressPercent: 100,
	}

	// change project ID to project number
	metadata.Resource = strings.Replace(obj.GetName(), "projects/"+name.Project.ID, "projects/"+strconv.FormatInt(name.Project.Number, 10), 1)

	// change project ID to project number in request details
	req.Parent = strings.Replace(req.GetParent(), "projects/"+name.Project.ID, "projects/"+strconv.FormatInt(name.Project.Number, 10), 1)
	metadata.Request = &pb.OperationMetadata_CreatePhraseSetRequest{
		CreatePhraseSetRequest: req,
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	return s.operations.DoneLRO(ctx, prefix, metadata, obj)
}

func (s *SpeechV2) UpdatePhraseSet(ctx context.Context, req *pb.UpdatePhraseSetRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePhraseSetName(req.GetPhraseSet().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.PhraseSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for i, path := range paths {
		switch path {
		case "displayName": // proto field name is display_name
			obj.DisplayName = req.GetPhraseSet().GetDisplayName()
			// HACK: to make the field mask valid when returning
			req.UpdateMask.Paths[i] = "display_name"
		case "phrases":
			obj.Phrases = req.GetPhraseSet().GetPhrases()
		case "boost":
			obj.Boost = req.GetPhraseSet().GetBoost()
		case "annotations":
			obj.Annotations = req.GetPhraseSet().GetAnnotations()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid for PhraseSet update", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Method:          "google.cloud.speech.v2.Speech.UpdatePhraseSet",
		ProgressPercent: 100,
	}

	// change project ID to project number
	metadata.Resource = strings.Replace(obj.GetName(), "projects/"+name.Project.ID, "projects/"+strconv.FormatInt(name.Project.Number, 10), 1)

	// change project ID to project number in request details
	req.PhraseSet.Name = strings.Replace(req.PhraseSet.GetName(), "projects/"+name.Project.ID, "projects/"+strconv.FormatInt(name.Project.Number, 10), 1)
	metadata.Request = &pb.OperationMetadata_UpdatePhraseSetRequest{
		UpdatePhraseSetRequest: req,
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	return s.operations.DoneLRO(ctx, prefix, metadata, obj)
}

func (s *SpeechV2) DeletePhraseSet(ctx context.Context, req *pb.DeletePhraseSetRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePhraseSetName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	// change project ID to project number in request details
	req.Name = strings.Replace(req.GetName(), "projects/"+name.Project.ID, "projects/"+strconv.FormatInt(name.Project.Number, 10), 1)

	obj := &pb.PhraseSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			if req.GetAllowMissing() {
				// Return a completed LRO indicating success (no-op)
				metadata := &pb.OperationMetadata{
					CreateTime:      timestamppb.New(now),
					UpdateTime:      timestamppb.New(now),
					Resource:        strings.Replace(fqn, "projects/"+name.Project.ID, "projects/"+strconv.FormatInt(name.Project.Number, 10), 1), // Use project number in metadata
					Method:          "google.cloud.speech.v2.Speech.DeletePhraseSet",
					ProgressPercent: 100,
					Request: &pb.OperationMetadata_DeletePhraseSetRequest{
						DeletePhraseSetRequest: req,
					},
				}
				// Return a placeholder object matching the LRO response type
				deletedPlaceholder := &pb.PhraseSet{
					Name:  strings.Replace(fqn, "projects/"+name.Project.ID, "projects/"+strconv.FormatInt(name.Project.Number, 10), 1),
					State: pb.PhraseSet_DELETED,
				}
				return s.operations.DoneLRO(ctx, prefix, metadata, deletedPlaceholder)
			}
			return nil, status.Errorf(codes.NotFound, "PhraseSet %q was not found.", name.PhraseSetID)
		}
		return nil, err
	}

	// Validate Etag if provided
	if req.GetEtag() != "" && req.GetEtag() != obj.Etag {
		return nil, status.Errorf(codes.Aborted, "etag mismatch for PhraseSet %q", name.PhraseSetID)
	}

	// Mark as deleted conceptually (although we delete immediately)
	obj.State = pb.PhraseSet_DELETED
	obj.DeleteTime = timestamppb.New(now)
	// Set expire time, e.g., 30 days from now, though it won't be used if we delete immediately
	obj.ExpireTime = timestamppb.New(now.Add(30 * 24 * time.Hour))

	// Delete from storage
	if err := s.storage.Delete(ctx, fqn, &pb.PhraseSet{}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete PhraseSet %q: %v", fqn, err)
	}

	metadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Method:          "google.cloud.speech.v2.Speech.DeletePhraseSet",
		ProgressPercent: 100,
		Request: &pb.OperationMetadata_DeletePhraseSetRequest{
			DeletePhraseSetRequest: req,
		},
		Resource: obj.GetName(),
	}

	return s.operations.DoneLRO(ctx, prefix, metadata, obj)
}

type phraseSetName struct {
	Project     *projects.ProjectData
	Location    string
	PhraseSetID string
}

func (n *phraseSetName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/phraseSets/%s", n.Project.Number, n.Location, n.PhraseSetID)
}

// parsePhraseSetName parses a string into a phraseSetName.
// The expected form is `projects/*/locations/*/phraseSets/*`.
func (s *MockService) parsePhraseSetName(name string) (*phraseSetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "phraseSets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "project %q not found: %v", tokens[1], err)
		}

		nameObj := &phraseSetName{
			Project:     project,
			Location:    tokens[3],
			PhraseSetID: tokens[5],
		}

		// Basic validation for IDs - should not be empty
		if nameObj.Location == "" || nameObj.PhraseSetID == "" {
			return nil, status.Errorf(codes.InvalidArgument, "name %q has empty location or phrase set ID", name)
		}

		return nameObj, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in the expected format projects/*/locations/*/phraseSets/*", name)
}
