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

package mockapikeys

import (
	"context"
	"fmt"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/api/apikeys/v2"
)

type APIKeysV2 struct {
	*MockService
	pb.UnimplementedApiKeysServer
}

func (s *APIKeysV2) GetKey(ctx context.Context, req *pb.GetKeyRequest) (*pb.Key, error) {
	name, err := s.parseAPIKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Key{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *APIKeysV2) GetKeyString(ctx context.Context, req *pb.GetKeyStringRequest) (*pb.GetKeyStringResponse, error) {
	name, err := s.parseAPIKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Key{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	keyString := "dummy-encrypted-value"
	return &pb.GetKeyStringResponse{
		KeyString: keyString,
	}, nil
}

func (s *APIKeysV2) CreateKey(ctx context.Context, req *pb.CreateKeyRequest) (*longrunning.Operation, error) {
	keyID := req.KeyId
	if keyID == "" {
		keyID = fmt.Sprintf("key-%d", time.Now().UnixNano())
	}
	reqName := req.Parent + "/keys/" + keyID
	name, err := s.parseAPIKeyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Key).(*pb.Key)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *APIKeysV2) UpdateKey(ctx context.Context, req *pb.UpdateKeyRequest) (*longrunning.Operation, error) {
	keyName := req.GetKey().GetName()

	name, err := s.parseAPIKeyName(keyName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Key{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// From the proto:
	// You can update only the `display_name`, `restrictions`, and `annotations` fields.

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()

	// If an update mask is not provided, the service treats it as an implied mask
	// equivalent to all allowed fields that are set on the wire
	if updateMask == nil {
		updateMask = &fieldmaskpb.FieldMask{}
		if req.GetKey().GetDisplayName() != "" {
			updateMask.Paths = append(updateMask.Paths, "display_name")
		}
		if req.GetKey().GetRestrictions() != nil {
			updateMask.Paths = append(updateMask.Paths, "restrictions")
		}
		if req.GetKey().GetAnnotations() != nil {
			updateMask.Paths = append(updateMask.Paths, "annotations")
		}
	}

	isSpecified := func(s string) bool {
		for _, p := range updateMask.Paths {
			if p == s {
				return true
			}
			if p == "*" {
				return true
			}
		}
		return false
	}

	if isSpecified("display_name") || isSpecified("displayName") {
		obj.DisplayName = req.GetKey().GetDisplayName()
	}
	if isSpecified("restrictions") {
		obj.Restrictions = req.GetKey().GetRestrictions()
	}
	if isSpecified("annotations") {
		obj.Annotations = req.GetKey().GetAnnotations()
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *APIKeysV2) DeleteKey(ctx context.Context, req *pb.DeleteKeyRequest) (*longrunning.Operation, error) {
	name, err := s.parseAPIKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Key{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, "", nil, func() (proto.Message, error) {
		return deleted, nil
	})
}
