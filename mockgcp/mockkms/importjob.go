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
// proto.service: google.cloud.kms.v1.KeyManagementService
// proto.message: google.cloud.kms.v1.KeyRing

package mockkms

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/kms/v1"
)

func (r *kmsServer) GetKeyRing(ctx context.Context, req *pb.GetKeyRingRequest) (*pb.KeyRing, error) {
	name, err := r.parseKeyRingName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.KeyRing{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "KeyRing %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *kmsServer) CreateKeyRing(ctx context.Context, req *pb.CreateKeyRingRequest) (*pb.KeyRing, error) {
	reqName := fmt.Sprintf("%s/keyRings/%s", req.GetParent(), req.GetKeyRingId())
	name, err := r.parseKeyRingName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetKeyRing()).(*pb.KeyRing)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	r.populateDefaultsForKeyRing(name, obj)

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *kmsServer) populateDefaultsForKeyRing(name *KeyRingName, obj *pb.KeyRing) {

}

type KeyRingName struct {
	Project   *projects.ProjectData
	Location  string
	KeyRingID string
}

func (n *KeyRingName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/keyRings/" + n.KeyRingID
}

// parseKeyRingName parses a string into an KeyRingName.
// The expected form is `projects/*/locations/*/keyRings/*`.
func (r *kmsServer) parseKeyRingName(name string) (*KeyRingName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "keyRings" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &KeyRingName{
			Project:   project,
			Location:  tokens[3],
			KeyRingID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

