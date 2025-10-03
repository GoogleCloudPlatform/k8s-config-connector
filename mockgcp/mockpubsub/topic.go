// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockpubsub

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/pubsub/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type publisherService struct {
	*MockService
	pb.UnimplementedPublisherServer
}

func (s *publisherService) CreateTopic(ctx context.Context, req *pb.Topic) (*pb.Topic, error) {
	name, err := s.parseTopicName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req).(*pb.Topic)
	obj.Name = name.String()

	s.populateDefaultsForTopic(obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *publisherService) populateDefaultsForTopic(obj *pb.Topic) {
	// TODO: When _is_ this populated?
	populateMessageStoragePolicy := false
	if populateMessageStoragePolicy {
		obj.MessageStoragePolicy = &pb.MessageStoragePolicy{
			AllowedPersistenceRegions: []string{
				"asia-east1",
				"asia-northeast1",
				"asia-southeast1",
				"australia-southeast1",
				"europe-north1",
				"europe-west1",
				"europe-west2",
				"europe-west3",
				"europe-west4",
				"southamerica-west1",
				"us-central1",
				"us-central2",
				"us-east1",
				"us-east4",
				"us-east5",
				"us-east7",
				"us-south1",
				"us-west1",
				"us-west2",
				"us-west3",
				"us-west4",
				"us-west8",
			},
		}
	}
}

func (s *publisherService) UpdateTopic(ctx context.Context, req *pb.UpdateTopicRequest) (*pb.Topic, error) {
	log := klog.FromContext(ctx)

	reqName := req.Topic.Name
	name, err := s.parseTopicName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.Topic{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)
	updated.Name = name.String()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// https://google.aip.dev/134
		// field mask must be optional, and the service must treat an omitted field mask as an implied field mask equivalent to all fields that are populated (have a non-empty value).

		paths = fields.ComputeImpliedFieldMask(ctx, req.GetTopic())
		log.Info("computed implied field_mask", "paths", paths)
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "name":
			if updated.Name != req.GetTopic().GetName() {
				return nil, status.Errorf(codes.InvalidArgument, "name is immutable")
			}
		case "labels":
			updated.Labels = req.GetTopic().GetLabels()
		case "schema_settings":
			updated.SchemaSettings = req.GetTopic().GetSchemaSettings()
		case "message_retention_duration":
			updated.MessageRetentionDuration = req.GetTopic().GetMessageRetentionDuration()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	// Note that if we try to pass e.g. 10d to message_retention_duration we get this:
	// {
	//   "error": {
	//     "code": 400,
	//     "message": "Invalid value at 'topic.message_retention_duration' (type.googleapis.com/google.protobuf.Duration), Field 'messageRetentionDuration', Illegal duration format; duration must end with 's'",
	//     "status": "INVALID_ARGUMENT",
	//     "details": [
	//       {
	//         "@type": "type.googleapis.com/google.rpc.BadRequest",
	//         "fieldViolations": [
	//           {
	//             "field": "topic.message_retention_duration",
	//             "description": "Invalid value at 'topic.message_retention_duration' (type.googleapis.com/google.protobuf.Duration), Field 'messageRetentionDuration', Illegal duration format; duration must end with 's'"
	//           }
	//         ]
	//       }
	//     ]
	//   }
	// }

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *publisherService) GetTopic(ctx context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {
	name, err := s.parseTopicName(req.Topic)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Topic{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource not found (resource=%s).", name.ID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *publisherService) ListTopics(ctx context.Context, req *pb.ListTopicsRequest) (*pb.ListTopicsResponse, error) {
	project, err := s.Projects.GetProjectByID(req.Project)
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%v/", project.ID)

	var topics []*pb.Topic

	topicKind := (&pb.Topic{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, topicKind, storage.ListOptions{}, func(obj proto.Message) error {
		topic := obj.(*pb.Topic)
		if strings.HasPrefix(topic.Name, findPrefix) {
			topics = append(topics, topic)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListTopicsResponse{
		Topics:        topics,
		NextPageToken: "",
	}, nil
}

func (s *publisherService) DeleteTopic(ctx context.Context, req *pb.DeleteTopicRequest) (*emptypb.Empty, error) {
	name, err := s.parseTopicName(req.Topic)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.Topic{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type topicName struct {
	Project *projects.ProjectData
	ID      string
}

func (n *topicName) String() string {
	return fmt.Sprintf("projects/%s/topics/%s", n.Project.ID, n.ID)
}

// parseTopicName parses a string into a topicName.
// The expected form is `projects/*/topics/*`.
func (s *MockService) parseTopicName(name string) (*topicName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "topics" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &topicName{
			Project: project,
			ID:      tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
