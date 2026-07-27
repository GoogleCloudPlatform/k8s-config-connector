// Copyright 2026 Google LLC
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
// See the License for the_internal:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockcontentwarehouse

import (
	"context"

	pb "google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SynonymSetServer struct {
	*MockService
	pb.UnimplementedSynonymSetServiceServer
}

func (s *SynonymSetServer) GetSynonymSet(ctx context.Context, req *pb.GetSynonymSetRequest) (*pb.SynonymSet, error) {
	name, err := s.parseSynonymSetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SynonymSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SynonymSetServer) CreateSynonymSet(ctx context.Context, req *pb.CreateSynonymSetRequest) (*pb.SynonymSet, error) {
	reqName := req.Parent + "/synonymSets/" + req.SynonymSet.Context
	name, err := s.parseSynonymSetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.SynonymSet).(*pb.SynonymSet)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SynonymSetServer) UpdateSynonymSet(ctx context.Context, req *pb.UpdateSynonymSetRequest) (*pb.SynonymSet, error) {
	name, err := s.parseSynonymSetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SynonymSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Context = req.SynonymSet.Context
	obj.Synonyms = req.SynonymSet.Synonyms

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SynonymSetServer) DeleteSynonymSet(ctx context.Context, req *pb.DeleteSynonymSetRequest) (*emptypb.Empty, error) {
	name, err := s.parseSynonymSetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SynonymSet{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
