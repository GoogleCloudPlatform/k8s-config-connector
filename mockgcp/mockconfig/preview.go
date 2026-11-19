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

package mockconfig

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/config/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type previewName struct {
	Project     *projects.ProjectData
	Location    string
	PreviewName string
}

func (n *previewName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/previews/" + n.PreviewName
}

func (s *MockService) parsePreviewName(name string) (*previewName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "previews" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &previewName{
			Project:     project,
			Location:    tokens[3],
			PreviewName: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ConfigV1) calculatePreviewStatus(obj *pb.Preview) {
	obj.State = pb.Preview_FAILED

	if tb := obj.GetTerraformBlueprint(); tb != nil {
		for _, v := range tb.GetExternalValues() {
			if dep := v.GetDeploymentSource().GetDeployment(); dep != "" {
				obj.ErrorCode = pb.Preview_EXTERNAL_VALUE_SOURCE_IMPORT_FAILED
				obj.ErrorStatus = &rpcstatus.Status{
					Code:    400,
					Message: fmt.Sprintf("Failed to import values from an external source: Invalid resource state for %q: the provided deployment doesn't exist", dep),
				}
				return
			}
		}
	}

	bucket := strings.TrimPrefix(obj.GetArtifactsGcsBucket(), "gs://")
	sa := obj.GetServiceAccount()
	if idx := strings.LastIndex(sa, "/"); idx != -1 {
		sa = sa[idx+1:]
	}
	obj.ErrorCode = pb.Preview_PREVIEW_BUILD_API_FAILED
	obj.ErrorStatus = &rpcstatus.Status{
		Code:    500,
		Message: fmt.Sprintf("The preview build failed due to an internal error: invalid bucket %q; service account %s does not have access to the bucket", bucket, sa),
	}
}

func (s *ConfigV1) GetPreview(ctx context.Context, req *pb.GetPreviewRequest) (*pb.Preview, error) {
	name, err := s.parsePreviewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Preview{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ConfigV1) ListPreviews(ctx context.Context, req *pb.ListPreviewsRequest) (*pb.ListPreviewsResponse, error) {
	response := &pb.ListPreviewsResponse{}
	findPrefix := req.Parent + "/previews/"

	previewKind := (&pb.Preview{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, previewKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		preview := obj.(*pb.Preview)
		response.Previews = append(response.Previews, preview)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ConfigV1) CreatePreview(ctx context.Context, req *pb.CreatePreviewRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/previews/" + req.PreviewId
	name, err := s.parsePreviewName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()

	obj := proto.Clone(req.Preview).(*pb.Preview)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	if obj.ArtifactsGcsBucket == nil || *obj.ArtifactsGcsBucket == "" {
		defaultBucket := fmt.Sprintf("gs://%d-%s-blueprint-config", name.Project.Number, name.Location)
		obj.ArtifactsGcsBucket = &defaultBucket
	}

	bucketBase := strings.TrimSuffix(*obj.ArtifactsGcsBucket, "/")
	obj.Logs = bucketBase + "/new-previews/" + name.PreviewName + "/logs"
	obj.PreviewArtifacts = &pb.PreviewArtifacts{
		Artifacts: bucketBase + "/new-previews/" + name.PreviewName + "/preview_results/artifacts",
		Content:   bucketBase + "/new-previews/" + name.PreviewName + "/preview_results/content",
	}

	if obj.TfVersionConstraint != nil && *obj.TfVersionConstraint != "" {
		obj.TfVersion = strings.TrimPrefix(*obj.TfVersionConstraint, "=")
	} else {
		obj.TfVersion = "1.5.7"
	}

	s.calculatePreviewStatus(obj)

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
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		opMetadata.ResourceMetadata = &pb.OperationMetadata_PreviewMetadata{
			PreviewMetadata: &pb.PreviewOperationMetadata{
				Logs:             obj.Logs,
				PreviewArtifacts: obj.PreviewArtifacts,
				Step:             pb.PreviewOperationMetadata_FAILED,
			},
		}
		lroResponse := proto.Clone(obj).(*pb.Preview)
		lroResponse.Annotations = nil
		lroResponse.Labels = nil
		return lroResponse, nil
	})
}

func (s *ConfigV1) DeletePreview(ctx context.Context, req *pb.DeletePreviewRequest) (*longrunning.Operation, error) {
	name, err := s.parsePreviewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	oldObj := &pb.Preview{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	responseObj := proto.Clone(oldObj).(*pb.Preview)
	responseObj.State = pb.Preview_DELETED
	responseObj.CreateTime = nil
	responseObj.Labels = nil
	responseObj.Annotations = nil
	responseObj.ErrorCode = pb.Preview_ERROR_CODE_UNSPECIFIED

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "delete",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		opMetadata.ResourceMetadata = &pb.OperationMetadata_PreviewMetadata{
			PreviewMetadata: &pb.PreviewOperationMetadata{
				Logs:             oldObj.Logs,
				PreviewArtifacts: oldObj.PreviewArtifacts,
				Step:             pb.PreviewOperationMetadata_PREPARING_STORAGE_BUCKET,
			},
		}
		return responseObj, nil
	})
}
