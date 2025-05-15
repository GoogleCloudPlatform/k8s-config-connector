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
// proto.message: google.cloud.speech.v2.Recognizer

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

// recognizerName represents the fully qualified name of a Recognizer.
type recognizerName struct {
	Project      *projects.ProjectData
	Location     string
	RecognizerID string
}

func (n *recognizerName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/recognizers/%s", n.Project.Number, n.Location, n.RecognizerID)
}

// parseRecognizerName parses a string into a recognizerName.
// The expected form is `projects/{project}/locations/{location}/recognizers/{recognizer}`.
func (s *SpeechV2) parseRecognizerName(name string) (*recognizerName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "recognizers" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "project %q not found: %v", tokens[1], err)
		}

		nameObj := &recognizerName{
			Project:      project,
			Location:     tokens[3],
			RecognizerID: tokens[5],
		}
		if nameObj.Location == "" || nameObj.RecognizerID == "" {
			return nil, status.Errorf(codes.InvalidArgument, "name %q has empty location or recognizer id", name)
		}

		return nameObj, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not in the expected format projects/{project}/locations/{location}/recognizers/{recognizer}", name)
}

// GetRecognizer returns the requested Recognizer.
func (s *SpeechV2) GetRecognizer(ctx context.Context, req *pb.GetRecognizerRequest) (*pb.Recognizer, error) {
	name, err := s.parseRecognizerName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Recognizer{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Unable to find Recognizer %s from project %d.", name.RecognizerID, name.Project.Number)
		}
		return nil, err
	}

	return obj, nil
}

// CreateRecognizer creates a new Recognizer.
func (s *SpeechV2) CreateRecognizer(ctx context.Context, req *pb.CreateRecognizerRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/recognizers/%s", req.GetParent(), req.GetRecognizerId())
	name, err := s.parseRecognizerName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetRecognizer()).(*pb.Recognizer)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = uuid.New().String()
	obj.State = pb.Recognizer_ACTIVE

	if obj.DefaultRecognitionConfig == nil {
		obj.DefaultRecognitionConfig = &pb.RecognitionConfig{
			Model:         obj.Model,
			LanguageCodes: obj.LanguageCodes,
		}
	}

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Resource:        obj.GetName(),
		Method:          "google.cloud.speech.v2.Speech.CreateRecognizer",
		ProgressPercent: 100,
	}

	// Ensure parent in metadata request uses project number
	metadataReq := proto.Clone(req).(*pb.CreateRecognizerRequest)
	parentProjectID := strings.Split(metadataReq.GetParent(), "/")[1]
	if _, err := strconv.ParseInt(parentProjectID, 10, 64); err != nil { // if not a number, it's an ID
		metadataReq.Parent = strings.Replace(metadataReq.GetParent(), "projects/"+parentProjectID, "projects/"+strconv.FormatInt(name.Project.Number, 10), 1)
	}
	metadata.Request = &pb.OperationMetadata_CreateRecognizerRequest{
		CreateRecognizerRequest: metadataReq,
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, prefix, metadata, obj)
}

// UpdateRecognizer updates an existing Recognizer.
func (s *SpeechV2) UpdateRecognizer(ctx context.Context, req *pb.UpdateRecognizerRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRecognizerName(req.GetRecognizer().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.Recognizer{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	updatedRecognizerFromRequest := req.GetRecognizer()

	for i, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = updatedRecognizerFromRequest.GetDisplayName()

			// HACK: to make the field mask valid when returning
			// original error:
			// proto: google.protobuf.FieldMask.paths contains irreversible value "displayName"
			req.UpdateMask.Paths[i] = "display_name"
		case "model":
			obj.Model = updatedRecognizerFromRequest.GetModel()
		case "language_codes":
			obj.LanguageCodes = updatedRecognizerFromRequest.GetLanguageCodes()
		case "default_recognition_config":
			obj.DefaultRecognitionConfig = updatedRecognizerFromRequest.GetDefaultRecognitionConfig()
		case "annotations":
			obj.Annotations = updatedRecognizerFromRequest.GetAnnotations()
		default:
			if strings.HasPrefix(path, "default_recognition_config.") {
				return nil, status.Errorf(codes.Unimplemented, "granular updates to sub-fields of default_recognition_config (e.g., %q) are not supported by this mock. Provide 'default_recognition_config' in update_mask to replace the entire config.", path)
			}
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid for Recognizer update or not supported by mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now), // LRO create time
		UpdateTime:      timestamppb.New(now), // LRO update time
		Resource:        obj.GetName(),
		Method:          "google.cloud.speech.v2.Speech.UpdateRecognizer",
		ProgressPercent: 100,
	}

	metadataReq := proto.Clone(req).(*pb.UpdateRecognizerRequest)
	parsedReqRecognizerNameForMeta, _ := s.parseRecognizerName(metadataReq.GetRecognizer().GetName())
	if parsedReqRecognizerNameForMeta != nil { // Use parsed name if successful
		metadataReq.GetRecognizer().Name = parsedReqRecognizerNameForMeta.String()
	}
	metadata.Request = &pb.OperationMetadata_UpdateRecognizerRequest{
		UpdateRecognizerRequest: metadataReq,
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, prefix, metadata, obj)
}

// DeleteRecognizer marks a Recognizer as DELETED.
func (s *SpeechV2) DeleteRecognizer(ctx context.Context, req *pb.DeleteRecognizerRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRecognizerName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	metadataReq := proto.Clone(req).(*pb.DeleteRecognizerRequest)
	metadataReq.Name = name.String() // Use name with project number for metadata

	obj := &pb.Recognizer{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			if req.GetAllowMissing() {
				opMetadata := &pb.OperationMetadata{CreateTime: timestamppb.New(now), UpdateTime: timestamppb.New(now), Resource: metadataReq.GetName(), Method: "google.cloud.speech.v2.Speech.DeleteRecognizer", ProgressPercent: 100, Request: &pb.OperationMetadata_DeleteRecognizerRequest{DeleteRecognizerRequest: metadataReq}}
				deletedPlaceholder := &pb.Recognizer{Name: metadataReq.GetName(), State: pb.Recognizer_DELETED}
				return s.operations.DoneLRO(ctx, prefix, opMetadata, deletedPlaceholder)
			}
			return nil, status.Errorf(codes.NotFound, "Recognizer %q not found.", fqn)
		}
		return nil, err
	}

	if req.GetEtag() != "" && req.GetEtag() != obj.Etag {
		return nil, status.Errorf(codes.Aborted, "etag mismatch for recognizer %q", fqn)
	}

	if obj.State != pb.Recognizer_DELETED {
		obj.State = pb.Recognizer_DELETED
		obj.DeleteTime = timestamppb.New(now)
		obj.ExpireTime = timestamppb.New(now.Add(30 * 24 * time.Hour))
		obj.UpdateTime = timestamppb.New(now)
		obj.Etag = fields.ComputeWeakEtag(obj)

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to mark recognizer %q as deleted: %v", fqn, err)
		}
	}

	opMetadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Resource:        obj.GetName(),
		Method:          "google.cloud.speech.v2.Speech.DeleteRecognizer",
		ProgressPercent: 100,
		Request:         &pb.OperationMetadata_DeleteRecognizerRequest{DeleteRecognizerRequest: metadataReq},
	}

	return s.operations.DoneLRO(ctx, prefix, opMetadata, obj)
}

/*
// UndeleteRecognizer undeletes a Recognizer.
func (s *SpeechV2) UndeleteRecognizer(ctx context.Context, req *pb.UndeleteRecognizerRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRecognizerName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()
	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)

	metadataReq := proto.Clone(req).(*pb.UndeleteRecognizerRequest)
	metadataReq.Name = name.String()

	obj := &pb.Recognizer{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.NotFound, "Recognizer %q not found, cannot undelete.", fqn)
	}

	if obj.State != pb.Recognizer_DELETED {
		return nil, status.Errorf(codes.FailedPrecondition, "Recognizer %q is not in a DELETED state.", fqn)
	}

	if req.GetEtag() != "" && req.GetEtag() != obj.Etag {
		return nil, status.Errorf(codes.Aborted, "etag mismatch for recognizer %q", fqn)
	}

	if obj.ExpireTime != nil && now.After(obj.ExpireTime.AsTime()) {
		return nil, status.Errorf(codes.FailedPrecondition, "Recognizer %q has passed its expiration time and cannot be undeleted.", fqn)
	}

	obj.State = pb.Recognizer_ACTIVE
	obj.DeleteTime = nil
	obj.ExpireTime = nil
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to undelete recognizer %q: %v", fqn, err)
	}

	opMetadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		UpdateTime:      timestamppb.New(now),
		Resource:        obj.GetName(),
		Method:          "google.cloud.speech.v2.Speech.UndeleteRecognizer",
		ProgressPercent: 100,
		Request:         &pb.OperationMetadata_UndeleteRecognizerRequest{UndeleteRecognizerRequest: metadataReq},
	}

	return s.operations.DoneLRO(ctx, prefix, opMetadata, obj)
}
*/
