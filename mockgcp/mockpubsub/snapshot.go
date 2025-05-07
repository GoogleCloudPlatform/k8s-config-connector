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
// proto.service: google.pubsub.v1.Subscriber
// proto.message: google.pubsub.v1.Snapshot

package mockpubsub

import (
	"context"
	"strings"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/pubsub/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *subscriberService) GetSnapshot(ctx context.Context, req *pb.GetSnapshotRequest) (*pb.Snapshot, error) {
	name, err := parseSnapshotName(req.Snapshot)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Snapshot{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource not found (resource=%s).", name.Snapshot)
		}
		return nil, err
	}
	return obj, nil
}

func (s *subscriberService) CreateSnapshot(ctx context.Context, req *pb.CreateSnapshotRequest) (*pb.Snapshot, error) {
	name, err := parseSnapshotName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	topicName := strings.Replace(strings.Replace(req.GetSubscription(), "subscriptions", "topics", 1), "subscription", "topic", 1)

	obj := &pb.Snapshot{
		ExpireTime: timestamppb.New(time.Now()),
		Name:       fqn,
		Topic:      topicName,
		Labels:     req.Labels,
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *subscriberService) DeleteSnapshot(ctx context.Context, req *pb.DeleteSnapshotRequest) (*emptypb.Empty, error) {
	name, err := parseSnapshotName(req.Snapshot)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Snapshot{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *subscriberService) ListSnapshots(ctx context.Context, req *pb.ListSnapshotsRequest) (*pb.ListSnapshotsResponse, error) {
	prefix := req.GetProject() + "/snapshots/"
	list := make([]*pb.Snapshot, 0)
	snapshotKind := (&pb.Snapshot{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, snapshotKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj protoreflect.ProtoMessage) error {
		snapshot, ok := obj.(*pb.Snapshot)
		if !ok {
			return status.Errorf(codes.Internal, "unexpected type %T in ListSnapshots", obj)
		}
		if strings.HasPrefix(snapshot.Name, prefix) {
			list = append(list, snapshot)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.ListSnapshotsResponse{Snapshots: list}, nil
}

type snapshotName struct {
	Project  string
	Snapshot string
}

func (n *snapshotName) String() string {
	return "projects/" + n.Project + "/snapshots/" + n.Snapshot
}

// parseSnapshotName parses a string into a snapshotName.
// The expected form is `projects/*/snapshots/*`.
func parseSnapshotName(name string) (*snapshotName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "snapshots" {
		name := &snapshotName{
			Project:  tokens[1],
			Snapshot: tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
