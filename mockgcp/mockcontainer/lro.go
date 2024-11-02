// Copyright 2023 Google LLC
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

package mockcontainer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/container/v1beta1"
	pbstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"
)

func (s *ClusterManagerV1) GetOperation(ctx context.Context, req *pb.GetOperationRequest) (*pb.Operation, error) {
	name, err := s.parseOperationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Operation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ClusterManagerV1) startLRO(ctx context.Context, project *projects.ProjectData, op *pb.Operation, callback func() (proto.Message, error)) (*pb.Operation, error) {
	now := time.Now()
	op.StartTime = now.UTC().Format(time.RFC3339Nano)

	operationID := fmt.Sprintf("operation-%d-%s", now.UnixMilli(), uuid.NewUUID())
	op.Name = operationID

	name := operationName{
		Project:   project,
		Operation: operationID,
	}

	if op.Location != "" {
		name.Location = op.Location
	} else if op.Zone != "" {
		name.Location = op.Zone
	}

	fqn := name.String()
	op.SelfLink = "https://container.googleapis.com/v1beta1/" + AsZonalLink(fqn)

	op.Status = pb.Operation_RUNNING

	if err := s.storage.Create(ctx, fqn, op); err != nil {
		return nil, err
	}

	go func() {
		_, err := callback()
		finished := &pb.Operation{}
		if err2 := s.storage.Get(ctx, fqn, finished); err2 != nil {
			klog.Warningf("error getting LRO: %v", err2)
			return
		}

		// Progress might have been updated by callback
		finished.Progress = op.Progress

		finished.Status = pb.Operation_DONE
		if err != nil {
			finished.Error = &pbstatus.Status{
				Message: err.Error(),
			}
		}
		finished.EndTime = time.Now().Format(time.RFC3339)

		if err := s.storage.Update(ctx, fqn, finished); err != nil {
			klog.Warningf("error updating LRO: %v", err)
			return
		}
	}()

	return op, nil
}

type operationName struct {
	Project   *projects.ProjectData
	Location  string
	Operation string
}

func (n *operationName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/operations/%s", n.Project.Number, n.Location, n.Operation)
}

// parseOperationName parses a string into a operationName.
// The expected form is `projects/*/locations/*/operations/*`.
func (s *MockService) parseOperationName(name string) (*operationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "operations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &operationName{
			Project:   project,
			Location:  tokens[3],
			Operation: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
