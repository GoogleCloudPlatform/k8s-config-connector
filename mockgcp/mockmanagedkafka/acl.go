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

package mockmanagedkafka

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"

	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type aclName struct {
	Project  *projects.ProjectData
	Location string
	Cluster  string
	AclID    string
}

func (n *aclName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/acls/%s", n.Project.ID, n.Location, n.Cluster, n.AclID)
}

func (s *MockService) parseAclName(name string) (*aclName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) >= 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" && tokens[6] == "acls" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &aclName{
			Project:  project,
			Location: tokens[3],
			Cluster:  tokens[5],
			AclID:    strings.Join(tokens[7:], "/"),
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid acl name %q", name)
}

func deriveAclFields(aclID string) (resourceType, resourceName, patternType string, err error) {
	switch {
	case aclID == "cluster":
		return "CLUSTER", "kafka-cluster", "LITERAL", nil
	case aclID == "allTopics":
		return "TOPIC", "*", "LITERAL", nil
	case aclID == "allConsumerGroups":
		return "GROUP", "*", "LITERAL", nil
	case aclID == "allTransactionalIds":
		return "TRANSACTIONAL_ID", "*", "LITERAL", nil
	case strings.HasPrefix(aclID, "topic/"):
		return "TOPIC", strings.TrimPrefix(aclID, "topic/"), "LITERAL", nil
	case strings.HasPrefix(aclID, "consumerGroup/"):
		return "GROUP", strings.TrimPrefix(aclID, "consumerGroup/"), "LITERAL", nil
	case strings.HasPrefix(aclID, "transactionalId/"):
		return "TRANSACTIONAL_ID", strings.TrimPrefix(aclID, "transactionalId/"), "LITERAL", nil
	case strings.HasPrefix(aclID, "topicPrefixed/"):
		return "TOPIC", strings.TrimPrefix(aclID, "topicPrefixed/"), "PREFIXED", nil
	case strings.HasPrefix(aclID, "consumerGroupPrefixed/"):
		return "GROUP", strings.TrimPrefix(aclID, "consumerGroupPrefixed/"), "PREFIXED", nil
	case strings.HasPrefix(aclID, "transactionalIdPrefixed/"):
		return "TRANSACTIONAL_ID", strings.TrimPrefix(aclID, "transactionalIdPrefixed/"), "PREFIXED", nil
	default:
		return "", "", "", status.Errorf(codes.InvalidArgument, "invalid aclId %q", aclID)
	}
}

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func (s *managedKafka) ListAcls(ctx context.Context, req *pb.ListAclsRequest) (*pb.ListAclsResponse, error) {
	name, err := s.parseClusterName(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListAclsResponse{}
	findPrefix := name.String() + "/acls/"

	listKind := (&pb.Acl{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		acl := obj.(*pb.Acl)
		if strings.HasPrefix(acl.GetName(), findPrefix) {
			response.Acls = append(response.Acls, acl)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *managedKafka) GetAcl(ctx context.Context, req *pb.GetAclRequest) (*pb.Acl, error) {
	name, err := s.parseAclName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Acl{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Acl 'projects/%d/locations/%s/clusters/%s/acls/%s' not found.", name.Project.Number, name.Location, name.Cluster, name.AclID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) CreateAcl(ctx context.Context, req *pb.CreateAclRequest) (*pb.Acl, error) {
	reqName := fmt.Sprintf("%s/acls/%s", req.GetParent(), req.GetAclId())
	name, err := s.parseAclName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetAcl()).(*pb.Acl)
	obj.Name = fqn

	rType, rName, pType, err := deriveAclFields(name.AclID)
	if err != nil {
		return nil, err
	}
	obj.ResourceType = rType
	obj.ResourceName = rName
	obj.PatternType = pType
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) UpdateAcl(ctx context.Context, req *pb.UpdateAclRequest) (*pb.Acl, error) {
	name, err := s.parseAclName(req.GetAcl().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Acl{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().GetPaths()) > 0 {
		for _, path := range req.GetUpdateMask().GetPaths() {
			switch path {
			case "acl_entries", "aclEntries":
				obj.AclEntries = req.GetAcl().GetAclEntries()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
			}
		}
	} else {
		obj.AclEntries = req.GetAcl().GetAclEntries()
	}

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) DeleteAcl(ctx context.Context, req *pb.DeleteAclRequest) (*emptypb.Empty, error) {
	name, err := s.parseAclName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Acl{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Acl 'projects/%d/locations/%s/clusters/%s/acls/%s' not found.", name.Project.Number, name.Location, name.Cluster, name.AclID)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
