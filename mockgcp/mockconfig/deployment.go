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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/config/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type deploymentName struct {
	Project        *projects.ProjectData
	Location       string
	DeploymentName string
}

func (n *deploymentName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/deployments/" + n.DeploymentName
}

func (s *MockService) parseDeploymentName(name string) (*deploymentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deployments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &deploymentName{
			Project:        project,
			Location:       tokens[3],
			DeploymentName: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ConfigV1) calculateStateDetail(obj *pb.Deployment) string {
	if tb := obj.GetTerraformBlueprint(); tb != nil {
		for _, v := range tb.GetExternalValues() {
			if dep := v.GetDeploymentSource().GetDeployment(); dep != "" {
				return fmt.Sprintf("Revision failed: Failed to import values from an external source: Invalid resource state for %q: the provided deployment doesn't exist", dep)
			}
		}
	}
	sa := obj.GetServiceAccount()
	if idx := strings.LastIndex(sa, "/"); idx != -1 {
		sa = sa[idx+1:]
	}
	bucket := strings.TrimPrefix(obj.GetArtifactsGcsBucket(), "gs://")
	return fmt.Sprintf("Revision failed: The revision failed due to an internal error: unable to create actuation build for revision %s. invalid bucket %q; service account %s does not have access to the bucket", obj.GetLatestRevision(), bucket, sa)
}

func (s *ConfigV1) GetDeployment(ctx context.Context, req *pb.GetDeploymentRequest) (*pb.Deployment, error) {
	name, err := s.parseDeploymentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ConfigV1) ListDeployments(ctx context.Context, req *pb.ListDeploymentsRequest) (*pb.ListDeploymentsResponse, error) {
	response := &pb.ListDeploymentsResponse{}
	findPrefix := req.Parent + "/deployments/"

	deploymentKind := (&pb.Deployment{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, deploymentKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		deployment := obj.(*pb.Deployment)
		response.Deployments = append(response.Deployments, deployment)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ConfigV1) CreateDeployment(ctx context.Context, req *pb.CreateDeploymentRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/deployments/" + req.DeploymentId
	name, err := s.parseDeploymentName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()

	obj := proto.Clone(req.Deployment).(*pb.Deployment)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.ArtifactsGcsBucket == nil || *obj.ArtifactsGcsBucket == "" {
		defaultBucket := fmt.Sprintf("gs://%d-%s-blueprint-config", name.Project.Number, name.Location)
		obj.ArtifactsGcsBucket = &defaultBucket
	}

	obj.ImportExistingResources = nil
	obj.LockState = pb.Deployment_UNLOCKED
	obj.State = pb.Deployment_FAILED
	obj.ErrorCode = pb.Deployment_REVISION_FAILED
	obj.LatestRevision = fqn + "/revisions/r-0"

	if obj.TfVersionConstraint != nil && *obj.TfVersionConstraint != "" {
		obj.TfVersion = strings.TrimPrefix(*obj.TfVersionConstraint, "=")
	} else {
		obj.TfVersion = "1.5.7"
	}

	obj.StateDetail = s.calculateStateDetail(obj)

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
		opMetadata.ResourceMetadata = &pb.OperationMetadata_DeploymentMetadata{
			DeploymentMetadata: &pb.DeploymentOperationMetadata{
				Step: pb.DeploymentOperationMetadata_SUCCEEDED,
			},
		}
		lroResponse := proto.Clone(obj).(*pb.Deployment)
		lroResponse.Annotations = nil
		return lroResponse, nil
	})
}

func (s *ConfigV1) UpdateDeployment(ctx context.Context, req *pb.UpdateDeploymentRequest) (*longrunning.Operation, error) {
	reqName := req.GetDeployment().GetName()

	name, err := s.parseDeploymentName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()
	obj := &pb.Deployment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetDeployment().GetLabels()
		case "annotations":
			obj.Annotations = req.GetDeployment().GetAnnotations()
		case "terraformBlueprint":
			obj.Blueprint = req.GetDeployment().GetBlueprint()
			obj.LatestRevision = fqn + "/revisions/r-1"
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.StateDetail = s.calculateStateDetail(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "update",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		opMetadata.ResourceMetadata = &pb.OperationMetadata_DeploymentMetadata{
			DeploymentMetadata: &pb.DeploymentOperationMetadata{
				Step: pb.DeploymentOperationMetadata_SUCCEEDED,
			},
		}
		lroResponse := proto.Clone(obj).(*pb.Deployment)
		lroResponse.Annotations = nil
		return lroResponse, nil
	})
}

func (s *ConfigV1) DeleteDeployment(ctx context.Context, req *pb.DeleteDeploymentRequest) (*longrunning.Operation, error) {
	name, err := s.parseDeploymentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	oldObj := &pb.Deployment{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	deleteBucket := oldObj.GetArtifactsGcsBucket()
	responseObj := &pb.Deployment{
		Name:               oldObj.Name,
		ArtifactsGcsBucket: oldObj.ArtifactsGcsBucket,
		DeleteLogs:         deleteBucket + "/" + name.DeploymentName + "/delete_results/logs",
		DeleteResults: &pb.ApplyResults{
			Artifacts: deleteBucket + "/" + name.DeploymentName + "/delete_results/artifacts",
			Content:   deleteBucket + "/" + name.DeploymentName + "/delete_results/content",
		},
		LatestRevision: oldObj.LatestRevision,
		LockState:      pb.Deployment_UNLOCKED,
		ServiceAccount: oldObj.ServiceAccount,
		State:          pb.Deployment_DELETED,
		Blueprint:      oldObj.Blueprint,
		TfVersion:      oldObj.TfVersion,
	}

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
		opMetadata.ResourceMetadata = &pb.OperationMetadata_DeploymentMetadata{
			DeploymentMetadata: &pb.DeploymentOperationMetadata{
				Step: pb.DeploymentOperationMetadata_SUCCEEDED,
			},
		}
		return responseObj, nil
	})
}
