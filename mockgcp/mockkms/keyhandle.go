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

	lro "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/kms/v1"
)

type autokeyServer struct {
	*MockService
	pb.UnimplementedAutokeyServer
}

func (r *autokeyServer) GetKeyHandle(ctx context.Context, req *pb.GetKeyHandleRequest) (*pb.KeyHandle, error) {
	parent, resourceID, err := r.parseKeyHandleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := parent.String() + "/keyHandles/" + resourceID

	obj := &pb.KeyHandle{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *autokeyServer) CreateKeyHandle(ctx context.Context, req *pb.CreateKeyHandleRequest) (*lro.Operation, error) {
	var reqName string
	if req.KeyHandleId != "" {
		reqName = req.Parent + "/keyHandles/" + req.KeyHandleId
	} else if req.KeyHandle.Name != "" {
		reqName = req.KeyHandle.Name
	} else {
		reqName = req.Parent + "/keyHandles/" + "5fe9854c-4a75-4ec9-8c27-c235754b981d"

	}
	parent, resourceID, err := r.parseKeyHandleName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := parent.String() + "/keyHandles/" + resourceID

	obj := proto.Clone(req.GetKeyHandle()).(*pb.KeyHandle)
	obj.Name = fqn

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return r.operations.StartLRO(ctx, parent.String(), nil, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.KeyHandle)
		return result, nil
	})
}

type parentName struct {
	projectID string
	location  string
}

func (a *parentName) String() string {
	return "projects/" + a.projectID + "/locations/" + a.location
}

func (r *autokeyServer) parseParentName(name string) (*parentName, error) {
	name = strings.TrimPrefix(name, "/")
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, fmt.Errorf("format of KMSKeyHandle external=%q was not known (use projects/<projectId>/locations/<location>/keyHandles/<keyhandleID>)", name)
	}
	return &parentName{
		projectID: tokens[1],
		location:  tokens[3],
	}, nil
}

// parseKeyHandleName parses a string into an KeyHandle name.
// The expected form is `projects/{projectId}/locations/<location>/keyHandles/<resourceId>`.
func (r *autokeyServer) parseKeyHandleName(name string) (*parentName, string, error) {
	name = strings.TrimPrefix(name, "/")
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "keyHandles" {
		return nil, "", fmt.Errorf("format of KMSKeyHandle external=%q was not known (use projects/<projectId>/locations/<location>/keyHandles/<keyhandleID>)", name)
	}
	parent := &parentName{
		projectID: tokens[1],
		location:  tokens[3],
	}
	return parent, tokens[5], nil
}
