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

// +tool:mockgcp-support
// proto.service: google.cloud.aiplatform.v1.MetadataService
// proto.message: google.cloud.aiplatform.v1.Artifact

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *metadataStoreService) GetArtifact(ctx context.Context, req *pb.GetArtifactRequest) (*pb.Artifact, error) {
	name, err := s.parseArtifactName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.Artifact{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource not found.; GetArtifact is unable to find artifact resource with name: %s", name.String())
		}
		return nil, err
	}

	result := proto.Clone(obj).(*pb.Artifact)
	result.Name = name.String()
	s.populateArtifactDefaults(result)
	return result, nil
}

func (s *metadataStoreService) CreateArtifact(ctx context.Context, req *pb.CreateArtifactRequest) (*pb.Artifact, error) {
	artifactId := req.GetArtifactId()
	if artifactId == "" {
		artifactId = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	reqName := fmt.Sprintf("%s/artifacts/%s", req.GetParent(), artifactId)
	name, err := s.parseArtifactName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.Clone(req.GetArtifact()).(*pb.Artifact)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateArtifactDefaults(obj)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.Artifact)
	return result, nil
}

func (s *metadataStoreService) UpdateArtifact(ctx context.Context, req *pb.UpdateArtifactRequest) (*pb.Artifact, error) {
	reqObj := req.GetArtifact()
	if reqObj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "artifact is required")
	}

	name, err := s.parseArtifactName(reqObj.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := &pb.Artifact{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource not found.; UpdateArtifact is unable to find artifact resource with name: %s", name.String())
		}
		return nil, err
	}

	var paths []string
	if req.GetUpdateMask() != nil {
		for _, p := range req.GetUpdateMask().GetPaths() {
			paths = append(paths, strings.Split(p, ",")...)
		}
	}

	if len(paths) == 0 {
		obj.DisplayName = reqObj.GetDisplayName()
		obj.Description = reqObj.GetDescription()
		obj.Uri = reqObj.GetUri()
		obj.State = reqObj.GetState()
		obj.SchemaTitle = reqObj.GetSchemaTitle()
		obj.SchemaVersion = reqObj.GetSchemaVersion()
		obj.Metadata = reqObj.GetMetadata()
		obj.Labels = reqObj.GetLabels()
		obj.Etag = reqObj.GetEtag()
	} else {
		for _, path := range paths {
			switch path {
			case "displayName", "display_name":
				obj.DisplayName = reqObj.GetDisplayName()
			case "description":
				obj.Description = reqObj.GetDescription()
			case "uri":
				obj.Uri = reqObj.GetUri()
			case "state":
				obj.State = reqObj.GetState()
			case "schemaTitle", "schema_title":
				obj.SchemaTitle = reqObj.GetSchemaTitle()
			case "schemaVersion", "schema_version":
				obj.SchemaVersion = reqObj.GetSchemaVersion()
			case "metadata":
				obj.Metadata = reqObj.GetMetadata()
			case "labels":
				obj.Labels = reqObj.GetLabels()
			case "etag":
				obj.Etag = reqObj.GetEtag()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
			}
		}
	}

	s.populateArtifactDefaults(obj)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.Artifact)
	return result, nil
}

func (s *metadataStoreService) DeleteArtifact(ctx context.Context, req *pb.DeleteArtifactRequest) (*longrunning.Operation, error) {
	name, err := s.parseArtifactName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	deleted := &pb.Artifact{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource not found.; DeleteArtifact is unable to find artifact resource with name: %s", name.String())
		}
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.DoneLROWithMetadata(ctx, opPrefix, op, &emptypb.Empty{})
}

func (s *metadataStoreService) ListArtifacts(ctx context.Context, req *pb.ListArtifactsRequest) (*pb.ListArtifactsResponse, error) {
	parent, err := s.parseMetadataStoreName(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/metadataStores/%s/artifacts/", parent.Project.Number, parent.Location, parent.MetadataStoreId)

	var artifacts []*pb.Artifact
	artifactKind := (&pb.Artifact{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, artifactKind, storage.ListOptions{}, func(obj proto.Message) error {
		artifact := obj.(*pb.Artifact)
		if strings.HasPrefix(artifact.GetName(), findPrefix) {
			artifacts = append(artifacts, artifact)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListArtifactsResponse{
		Artifacts: artifacts,
	}, nil
}

func (s *MockService) populateArtifactDefaults(obj *pb.Artifact) {
	if obj.Metadata == nil {
		obj.Metadata = &structpb.Struct{}
	}
	if obj.SchemaTitle != "" && obj.SchemaVersion == "" {
		obj.SchemaVersion = "0.0.1"
	}
}

type ArtifactName struct {
	Project         *projects.ProjectData
	Location        string
	MetadataStoreId string
	ArtifactId      string
}

func (n *ArtifactName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/metadataStores/%s/artifacts/%s", n.Project.Number, n.Location, n.MetadataStoreId, n.ArtifactId)
}

func (n *ArtifactName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/metadataStores/%s/artifacts/%s", n.Project.Number, n.Location, n.MetadataStoreId, n.ArtifactId)
}

// parseArtifactName parses a string into an ArtifactName.
// Expected form: projects/<projectID_or_number>/locations/<location>/metadataStores/<metadataStoreID>/artifacts/<artifactID>
func (s *MockService) parseArtifactName(name string) (*ArtifactName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "metadataStores" && tokens[6] == "artifacts" {
		parentName := strings.Join(tokens[0:6], "/")
		parent, err := s.parseMetadataStoreName(parentName)
		if err != nil {
			return nil, err
		}
		return &ArtifactName{
			Project:         parent.Project,
			Location:        parent.Location,
			MetadataStoreId: parent.MetadataStoreId,
			ArtifactId:      tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
