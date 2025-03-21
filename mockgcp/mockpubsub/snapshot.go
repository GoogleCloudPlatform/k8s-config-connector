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
        "encoding/json"
        "net/http"
        "strings"

	"google.golang.org/grpc/codes"
        "google.golang.org/grpc/status"
        "google.golang.org/protobuf/reflect/protoreflect"
        "google.golang.org/protobuf/types/known/emptypb"
        pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/pubsub/v1"
)



import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/pubsub/v1"
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
			return nil, status.Errorf(codes.NotFound, "snapshot %q not found", req.GetSnapshot())
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
	_, err = parseSubscriptionName(req.GetSubscription())
        if err != nil {
            return nil, err
        }

	fqn := name.String()

        topicName := strings.Replace(req.GetSubscription(), "subscriptions", "topics", 1)

	obj := &pb.Snapshot{
		Name:  fqn,
		Topic: topicName,
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type subscriptionNameForSnapshot struct {
	Project      string
	Subscription string
}

func (n *subscriptionNameForSnapshot) String() string {
	return "projects/" + n.Project + "/subscriptions/" + n.Subscription
}

// parseSubscriptionName parses a string into a subscriptionNameForSnapshot.
// The expected form is `projects/*/subscriptions/*`.
func parseSubscriptionName(name string) (*subscriptionNameForSnapshot, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "subscriptions" {
		name := &subscriptionNameForSnapshot{
			Project:      tokens[1],
			Subscription: tokens[3],
		}
		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "subscription name %q is not valid", name)
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


func (s *subscriberService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/v1/projects/") {
		suffix := strings.TrimPrefix(r.URL.Path, "/v1/")
		if strings.HasSuffix(suffix, ":get") && r.Method == http.MethodPost{
			// get snapshot
			var req pb.GetSnapshotRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			req.Snapshot = strings.TrimSuffix(suffix,":get")

			resp, err := s.GetSnapshot(r.Context(), &req)
			if err != nil {
				status, ok := status.FromError(err)
				if !ok {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					http.Error(w, err.Error(), int(status.Code()))
				}
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if strings.HasSuffix(suffix, ":create") && r.Method == http.MethodPost{
			var req pb.CreateSnapshotRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			req.Name = strings.TrimSuffix(suffix,":create")

			resp, err := s.CreateSnapshot(r.Context(), &req)
            if err != nil {
				status, ok := status.FromError(err)
				if !ok {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					http.Error(w, err.Error(), int(status.Code()))
				}
				return
			}
			w.Header().Set("Content-Type", "application/json")
            if err := json.NewEncoder(w).Encode(resp); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

		} else if strings.HasSuffix(suffix, ":delete") && r.Method == http.MethodPost{
			// delete snapshot
			var req pb.DeleteSnapshotRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			req.Snapshot = strings.TrimSuffix(suffix,":delete")
			resp, err := s.DeleteSnapshot(r.Context(), &req)
			if err != nil {
				status, ok := status.FromError(err)
				if !ok {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					http.Error(w, err.Error(), int(status.Code()))
				}
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if r.Method == http.MethodGet {
			// list snapshots
			var req pb.ListSnapshotsRequest
			req.Project = strings.TrimSuffix(suffix, "/snapshots")

			resp, err := s.ListSnapshots(r.Context(), &req)
			if err != nil {
				status, ok := status.FromError(err)
				if !ok {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					http.Error(w, err.Error(), int(status.Code()))
				}
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.NotFound(w, r)
		}
	} else {
		http.NotFound(w, r)
	}
}

func (s *subscriberService) ListSnapshots(ctx context.Context, req *pb.ListSnapshotsRequest) (*pb.ListSnapshotsResponse, error) {

	projectName, err := parseProjectName(req.GetProject())
	if err != nil {
		return nil, err
	}
	prefix := "projects/" + projectName.Project + "/snapshots/"
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

type projectName struct {
	Project string
}

func (n *projectName) String() string {
	return "projects/" + n.Project
}

func parseProjectName(name string) (*projectName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		name := &projectName{Project: tokens[1]}
		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "project name %q is not valid", name)
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
