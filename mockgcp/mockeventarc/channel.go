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
// proto.service: google.cloud.eventarc.v1.Eventarc
// proto.message: google.cloud.eventarc.v1.Channel

package mockeventarc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/eventarc/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *EventarcV1) GetChannel(ctx context.Context, req *pb.GetChannelRequest) (*pb.Channel, error) {
	name, err := s.parseChannelName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Channel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *EventarcV1) CreateChannel(ctx context.Context, req *pb.CreateChannelRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/channels/%s", req.GetParent(), req.GetChannelId())
	name, err := s.parseChannelName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetChannel()).(*pb.Channel)
	obj.Name = fqn
	obj.Uid = name.Channel
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Channel_ACTIVE // By default, new resources are considered ACTIVE.
	parsedName, err := s.parseChannelName(obj.Name)
	if err != nil {
		return nil, err
	}
	obj.State = pb.Channel_ACTIVE
	obj.Transport = &pb.Channel_PubsubTopic{PubsubTopic: fmt.Sprintf("projects/%s/topics/eventarc-channel-%s-%s-368", parsedName.Project.ID, parsedName.Location, parsedName.Channel)}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Return an LRO that doesnt finish immediately
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	lro.Done = false
	lro.Metadata, err = anypb.New(lroMetadata)
	if err != nil {
		return nil, err
	}
	// Use the fully qualified type name to ensure compatibility with the expected output.
	lro.Metadata.TypeUrl = "type.googleapis.com/google.cloud.eventarc.v1.OperationMetadata"
	lro.Name = fmt.Sprintf("projects/%s/locations/%s/operations/%s", name.Project.ID, name.Location, strings.Split(lro.Name, "/")[len(strings.Split(lro.Name, "/"))-1])

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		updated, err := s.updateChannel(ctx, fqn, func(obj *pb.Channel) {
			obj.State = pb.Channel_ACTIVE
		})
		if err != nil {
			return nil, err
		}
		return updated, err
	})
}

// updateService will read-modify-write the object with optimistic locking
func (s *EventarcV1) updateChannel(ctx context.Context, fqn string, update func(obj *pb.Channel)) (*pb.Channel, error) {
	obj := &pb.Channel{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	update(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *EventarcV1) UpdateChannel(ctx context.Context, req *pb.UpdateChannelRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetChannel().GetName()
	name, err := s.parseChannelName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Channel{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "crypto_key_name", "cryptoKeyName":
			obj.CryptoKeyName = req.GetChannel().GetCryptoKeyName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not supported for update", path)
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroRet := proto.Clone(obj).(*pb.Channel)
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return lroRet, nil
	})
}

func (s *EventarcV1) DeleteChannel(ctx context.Context, req *pb.DeleteChannelRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseChannelName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.Channel{}

	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	deletedObj.State = pb.Channel_INACTIVE
	deletedObj.Transport = &pb.Channel_PubsubTopic{PubsubTopic: ""}
	deletedObj.UpdateTime = nil
	deletedObj.CreateTime = nil
	deletedObj.Uid = ""

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return deletedObj, nil
	})
}

func (s *EventarcV1) populateDefaultsForChannel(obj *pb.Channel) {

}

type channelName struct {
	Project  *projects.ProjectData
	Location string
	Channel  string
}

func (n *channelName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/channels/%s", n.Project.ID, n.Location, n.Channel)
}

// parseChannelName parses a string into an channelName.
// The expected form is `projects/*/locations/*/channels/*`.
func (s *MockService) parseChannelName(name string) (*channelName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "channels" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &channelName{
			Project:  project,
			Location: tokens[3],
			Channel:  tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
