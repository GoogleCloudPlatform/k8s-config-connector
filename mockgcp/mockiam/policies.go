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
// proto.service: google.iam.v2.Policies
// proto.message: google.iam.v2.Policy

package mockiam

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	pb "cloud.google.com/go/iam/apiv2/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

// IAMV2PoliciesServer is a simple mock for the IAM V2 Policies service.
type IAMV2PoliciesServer struct {
	*MockService
	pb.UnimplementedPoliciesServer
}

type attachmentPointName struct {
	Project *projects.ProjectData
}

func (n *attachmentPointName) String() string {
	s := fmt.Sprintf("cloudresourcemanager.googleapis.com/projects/%d", n.Project.Number)
	return url.PathEscape(s)
}

// PolicyName represents the parsed name of an IAM V2 Policy.
type policyName struct {
	// AttachmentPoint specifies the entity this policy is attached to.
	AttachmentPoint attachmentPointName
	// PolicyID is the user-defined ID for the deny policy.
	PolicyID string
}

func (n *policyName) String() string {
	return fmt.Sprintf("policies/%s/denypolicies/%s", n.AttachmentPoint.String(), n.PolicyID)
}

// parseAttachmentPointName parses a string into an attachmentPointName.
// The expected form is the url-encoding of `cloudresourcemanager/projects/<project-id-or-number>`.
func (s *MockService) parseAttachmentPointName(escapedName string) (*attachmentPointName, error) {
	name, err := url.PathUnescape(escapedName)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid url-encoded attachment point name", escapedName)
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 3 && tokens[0] == "cloudresourcemanager.googleapis.com" && tokens[1] == "projects" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[2])
		if err != nil {
			return nil, err
		}

		return &attachmentPointName{
			Project: project,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid IAM V2 attachment point", name)
}

// parsePolicyName parses a string into an policyName.
// The expected form is `policies/{attachment_point}/denypolicies/{policy_id}`.
func (s *MockService) parsePolicyName(name string) (*policyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "policies" && tokens[2] == "denypolicies" {
		attachmentPointName, err := s.parseAttachmentPointName(tokens[1])
		if err != nil {
			return nil, err
		}
		return &policyName{
			AttachmentPoint: *attachmentPointName,
			PolicyID:        tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid IAM V2 policy name", name)
}

// parseIAMV2PolicyParent parses a parent string into the (escaped) attachment point part of a policy name.
// The expected form is `policies/{attachment_point}/denypolicies`.
func (s *MockService) parseIAMV2PolicyParent(parent string) (*attachmentPointName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 3 && tokens[0] == "policies" && tokens[2] == "denypolicies" {
		return s.parseAttachmentPointName(tokens[1])
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not a valid IAM V2 policy parent", parent)
}

// GetPolicy retrieves a policy.
func (s *IAMV2PoliciesServer) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.Policy, error) {
	name, err := s.parsePolicyName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Policy %q not found.", fqn)
		}
		return nil, status.Errorf(codes.Internal, "Error getting policy: %v", err)
	}

	return obj, nil
}

// ListPolicies lists policies attached to a resource.
func (s *IAMV2PoliciesServer) ListPolicies(ctx context.Context, req *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	attachmentPoint, err := s.parseIAMV2PolicyParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefixToList := fmt.Sprintf("policies/%s/denypolicies/", attachmentPoint.String())

	response := &pb.ListPoliciesResponse{}
	policyKind := (&pb.Policy{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, policyKind, storage.ListOptions{Prefix: prefixToList}, func(objproto proto.Message) error {
		obj := objproto.(*pb.Policy)
		// Omit rules in list response as per documentation.
		policyToList := proto.Clone(obj).(*pb.Policy)
		policyToList.Rules = nil
		response.Policies = append(response.Policies, policyToList)
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Error listing policies: %v", err)
	}

	// TODO: Implement pagination.
	return response, nil
}

// CreatePolicy creates a new policy.
func (s *IAMV2PoliciesServer) CreatePolicy(ctx context.Context, req *pb.CreatePolicyRequest) (*longrunningpb.Operation, error) {
	attachmentPoint, err := s.parseIAMV2PolicyParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	policyID := req.GetPolicyId()
	if policyID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "policy_id is required")
	}

	name := &policyName{
		AttachmentPoint: *attachmentPoint,
		PolicyID:        policyID,
	}
	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.GetPolicy()).(*pb.Policy)

	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.Kind = "DenyPolicy"
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.DeleteTime = nil // Not deleted on create
	obj.Etag = computeIAMV2PolicyEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "Policy %q already exists.", fqn)
		}
		return nil, status.Errorf(codes.Internal, "Error creating policy: %v", err)
	}

	lroMetadata := &pb.PolicyOperationMetadata{}
	lroMetadata.CreateTime = timestamppb.New(now)

	return s.operations.DoneLRO(ctx, fqn, lroMetadata, obj)
}

// UpdatePolicy updates an existing policy.
func (s *IAMV2PoliciesServer) UpdatePolicy(ctx context.Context, req *pb.UpdatePolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePolicyName(req.GetPolicy().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	existing := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Policy %q not found.", fqn)
		}
		return nil, status.Errorf(codes.Internal, "Error getting policy for update: %v", err)
	}

	// Check etag for optimistic concurrency control.
	if req.GetPolicy().GetEtag() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "etag is required for UpdatePolicy")
	}
	if req.GetPolicy().GetEtag() != existing.GetEtag() {
		return nil, status.Errorf(codes.Aborted, "etag mismatch for policy %q", fqn)
	}

	updatedPolicy := proto.Clone(existing).(*pb.Policy)

	// Only 'display_name' and 'rules' can be updated.
	updateRequestPolicy := req.GetPolicy()
	updatedPolicy.DisplayName = updateRequestPolicy.GetDisplayName()
	updatedPolicy.Rules = updateRequestPolicy.GetRules()
	updatedPolicy.Annotations = updateRequestPolicy.GetAnnotations() // Annotations are also updatable

	updatedPolicy.UpdateTime = timestamppb.New(now)
	updatedPolicy.Etag = computeIAMV2PolicyEtag(updatedPolicy)

	if err := s.storage.Update(ctx, fqn, updatedPolicy); err != nil {
		return nil, status.Errorf(codes.Internal, "Error updating policy: %v", err)
	}

	lroMetadata := &pb.PolicyOperationMetadata{}
	lroMetadata.CreateTime = timestamppb.New(now)

	return s.operations.DoneLRO(ctx, fqn, lroMetadata, updatedPolicy)
}

// DeletePolicy deletes a policy.
func (s *IAMV2PoliciesServer) DeletePolicy(ctx context.Context, req *pb.DeletePolicyRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePolicyName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.Policy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			// According to API behavior for some GCP services, deleting a non-existent
			// resource can be a no-op or return a specific LRO.
			// Let's assume it should "succeed" by returning an LRO that indicates deletion.
			// However, the test expects a NotFound error.
			return nil, status.Errorf(codes.NotFound, "Policy %q not found.", fqn)
		}
		return nil, status.Errorf(codes.Internal, "Error getting policy for delete: %v", err)
	}

	now := time.Now()

	// Check etag for optimistic concurrency control if provided.
	if req.GetEtag() != "" && req.GetEtag() != existing.GetEtag() {
		return nil, status.Errorf(codes.Aborted, "etag mismatch for policy %q", fqn)
	}

	if err := s.storage.Delete(ctx, fqn, &pb.Policy{}); err != nil {
		return nil, status.Errorf(codes.Internal, "Error deleting policy: %v", err)
	}

	// The LRO response type is Policy. Return the policy as it was, but mark it as deleted.
	deletedPolicy := proto.Clone(existing).(*pb.Policy)
	deletedPolicy.DeleteTime = timestamppb.New(now)
	deletedPolicy.UpdateTime = deletedPolicy.GetDeleteTime()   // Update time is also set to delete time
	deletedPolicy.Etag = computeIAMV2PolicyEtag(deletedPolicy) // Etag changes upon deletion

	lroMetadata := &pb.PolicyOperationMetadata{}
	lroMetadata.CreateTime = timestamppb.New(now)

	return s.operations.DoneLRO(ctx, fqn, lroMetadata, deletedPolicy)
}

// computeIAMV2PolicyEtag computes a simple etag for a Policy object.
func computeIAMV2PolicyEtag(obj *pb.Policy) string {
	// Create a copy and clear output-only or server-set fields that shouldn't affect etag
	temp := proto.Clone(obj).(*pb.Policy)
	temp.Name = ""
	temp.Uid = ""
	temp.CreateTime = nil
	temp.UpdateTime = nil
	temp.DeleteTime = nil
	// temp.Etag = "" // Don't include the etag itself in etag calculation

	b, err := proto.Marshal(temp)
	if err != nil {
		klog.Fatalf("failed to marshal proto object for etag: %v", err)
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
