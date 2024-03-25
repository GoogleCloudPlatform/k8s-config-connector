// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockaccesscontextmanager

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/identity/accesscontextmanager/v1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"
)

func (s *AccessContextManagerV1) GetAccessPolicy(ctx context.Context, req *pb.GetAccessPolicyRequest) (*pb.AccessPolicy, error) {
	name, err := s.parseAccessPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AccessPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AccessContextManagerV1) CreateAccessPolicy(ctx context.Context, req *pb.AccessPolicy) (*longrunning.Operation, error) {
	accessPolicy := req.Name
	if accessPolicy == "" {
		accessPolicy = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	reqName := "accessPolicies/" + accessPolicy
	name, err := s.parseAccessPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req).(*pb.AccessPolicy)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AccessContextManagerV1) UpdateAccessPolicy(ctx context.Context, req *pb.UpdateAccessPolicyRequest) (*longrunning.Operation, error) {
	reqName := req.GetPolicy().GetName()

	name, err := s.parseAccessPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AccessPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AccessContextManagerV1) DeleteAccessPolicy(ctx context.Context, req *pb.DeleteAccessPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseAccessPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.AccessPolicy{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func computeEtag(obj proto.Message) []byte {
	// TODO: Do we risk exposing internal fields?  Doesn't matter on a mock, I guess
	b, err := proto.Marshal(obj)
	if err != nil {
		klog.Fatalf("failed to marshal proto object: %v", err)
	}
	hash := md5.Sum(b)
	return hash[:]
}
