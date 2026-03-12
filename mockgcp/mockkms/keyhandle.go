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
// proto.service: google.cloud.kms.v1.Autokey
// proto.message: google.cloud.kms.v1.KeyHandle

package mockkms

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	lro "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/util/uuid"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/kms/v1"
)

type autokeyServer struct {
	*MockService
	pb.UnimplementedAutokeyServer
}

func (r *autokeyServer) GetKeyHandle(ctx context.Context, req *pb.GetKeyHandleRequest) (*pb.KeyHandle, error) {
	keyHandleName, err := r.parseKeyHandleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := keyHandleName.String()

	obj := &pb.KeyHandle{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "KeyHandle %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *autokeyServer) CreateKeyHandle(ctx context.Context, req *pb.CreateKeyHandleRequest) (*lro.Operation, error) {
	var reqName string
	uuid := string(uuid.NewUUID())
	if req.KeyHandleId != "" {
		reqName = req.Parent + "/keyHandles/" + req.KeyHandleId
	} else {
		reqName = req.Parent + "/keyHandles/" + uuid
	}
	keyHandleName, err := r.parseKeyHandleName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := keyHandleName.String()

	obj := proto.Clone(req.GetKeyHandle()).(*pb.KeyHandle)
	project, err := r.Projects.GetProjectByID(keyHandleName.projectID)
	if err != nil {
		return nil, err
	}
	obj.Name = fqn
	// Autokey full relative name:
	// projects/key-project/locations/us-central1/keyRings/autokey/cryptoKeys/${projectNumber}-compute-disk-ac103725a174885c
	// key-project is a separate project that contains auto keys, see https://cloud.google.com/kms/docs/enable-autokey#set-up-key-project
	// hard-code key-project and the generated key suffix for golden testing
	obj.KmsKey = fmt.Sprintf("projects/${key_project}/locations/%s/keyRings/autokey/cryptoKeys/%d-compute-disk-${generated-id}", keyHandleName.location, project.Number)

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreateKeyHandleMetadata{}
	return r.operations.StartLROWithOptions(ctx, req.Parent, metadata, func() (proto.Message, error) {
		return obj, nil
	}, false)
}

type KeyHandleName struct {
	projectID   string
	location    string
	keyHandleID string
}

func (a *KeyHandleName) String() string {
	return "projects/" + a.projectID + "/locations/" + a.location + "/keyHandles/" + a.keyHandleID
}

// parseKeyHandleName parses a string into an KeyHandle name.
// The expected form is `projects/{projectId}/locations/<location>/keyHandles/<resourceId>`.
func (r *autokeyServer) parseKeyHandleName(name string) (*KeyHandleName, error) {
	name = strings.TrimPrefix(name, "/")
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "keyHandles" {
		return nil, fmt.Errorf("format of KMSKeyHandle external=%q was not known (use projects/<projectId>/locations/<location>/keyHandles/<keyhandleID>)", name)
	}
	return &KeyHandleName{
		projectID:   tokens[1],
		location:    tokens[3],
		keyHandleID: tokens[5],
	}, nil
}
