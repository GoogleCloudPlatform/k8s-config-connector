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
// See the License for the specific language governing permissions and
// limitations under the License.

package mockbigtable

import (
	"context"
	"fmt"
	"strings"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	// MemoryLayer is only defined in the newer admin proto.
	"cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

// mockMemoryLayerStorageSizeGibPerNode is the size per cluster node in GiB that
// the mock reports for an enabled memory layer, matching the real service's 8 GiB/node.
const mockMemoryLayerStorageSizeGibPerNode = 8

// mockMemoryLayerEtag is the etag the mock hands out for an enabled memory layer,
// matching the opaque value the other mockbigtable resources use.
const mockMemoryLayerEtag = "abcdef0123A="

type memoryLayerName struct {
	cluster *clusterName
}

func (n *memoryLayerName) String() string {
	return n.cluster.String() + "/memoryLayer"
}

// parseMemoryLayerName parses a string into a memoryLayerName.
// The expected form is <clusterName>/memoryLayer.
func (s *MockService) parseMemoryLayerName(name string) (*memoryLayerName, error) {
	clusterPart, ok := strings.CutSuffix(name, "/memoryLayer")
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
	cluster, err := s.parseClusterName(clusterPart)
	if err != nil {
		return nil, err
	}
	return &memoryLayerName{cluster: cluster}, nil
}

func (s *instanceAdminServer) GetMemoryLayer(ctx context.Context, req *adminpb.GetMemoryLayerRequest) (*adminpb.MemoryLayer, error) {
	name, err := s.parseMemoryLayerName(req.GetName())
	if err != nil {
		return nil, err
	}

	// A memory layer only exists as part of its cluster.
	cluster := &pb.Cluster{}
	if err := s.storage.Get(ctx, name.cluster.String(), cluster); err != nil {
		return nil, err
	}

	return s.getOrDefaultMemoryLayer(ctx, name)
}

// defaultMemoryLayer builds the disabled memory layer that the service reports
// for every cluster that has never turned one on. Like the real service, a
// disabled memory layer does not carry an etag.
func defaultMemoryLayer(fqn string) *adminpb.MemoryLayer {
	return &adminpb.MemoryLayer{
		Name:  fqn,
		State: adminpb.MemoryLayer_DISABLED,
	}
}

// getOrDefaultMemoryLayer reads the stored memory layer, falling back to the
// disabled memory layer that the service reports for every cluster that has
// never turned one on.
func (s *MockService) getOrDefaultMemoryLayer(ctx context.Context, name *memoryLayerName) (*adminpb.MemoryLayer, error) {
	obj := &adminpb.MemoryLayer{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}
		return defaultMemoryLayer(name.String()), nil
	}
	return obj, nil
}

func (s *instanceAdminServer) UpdateMemoryLayer(ctx context.Context, req *adminpb.UpdateMemoryLayerRequest) (*longrunning.Operation, error) {
	name, err := s.parseMemoryLayerName(req.GetMemoryLayer().GetName())
	if err != nil {
		return nil, err
	}

	cluster := &pb.Cluster{}
	if err := s.storage.Get(ctx, name.cluster.String(), cluster); err != nil {
		return nil, err
	}

	fqn := name.String()

	// Distinguish "never enabled" from "stored", so we know whether to Create
	// or Update the storage entry below.
	stored := &adminpb.MemoryLayer{}
	exists := true
	if err := s.storage.Get(ctx, fqn, stored); err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}
		exists = false
		stored = defaultMemoryLayer(fqn)
	}

	// The service requires an up-to-date etag when updating or disabling an
	// already-enabled memory layer, but not on the initial enable from DISABLED.
	reqEtag := req.GetMemoryLayer().GetEtag()
	if stored.GetState() != adminpb.MemoryLayer_DISABLED {
		if reqEtag == "" {
			return nil, status.Errorf(codes.InvalidArgument, "etag is required when updating an enabled memory layer")
		}
		if reqEtag != stored.GetEtag() {
			return nil, status.Errorf(codes.Aborted, "etag %q does not match current etag %q", reqEtag, stored.GetEtag())
		}
	}

	now := timestamppb.Now()

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "memory_config":
			if req.GetMemoryLayer().GetMemoryConfig() == nil {
				stored.MemoryConfig = nil
				stored.State = adminpb.MemoryLayer_DISABLED
				stored.Etag = ""
			} else {
				// storage_size_gib is output only; the service decides it based on cluster nodes.
				nodes := cluster.GetServeNodes()
				if nodes <= 0 {
					nodes = 1
				}
				stored.MemoryConfig = &adminpb.MemoryLayer_MemoryConfig{
					StorageSizeGib: mockMemoryLayerStorageSizeGibPerNode * nodes,
				}
				stored.State = adminpb.MemoryLayer_READY
				stored.Etag = mockMemoryLayerEtag
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "mock does not implement update of MemoryLayer field %q", path)
		}
	}

	if exists {
		if err := s.storage.Update(ctx, fqn, stored); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, fqn, stored); err != nil {
			return nil, err
		}
	}

	zone := lastComponent(cluster.GetLocation())
	lroPrefix := fmt.Sprintf("operations/%s/locations/%s", fqn, zone)
	metadata := &adminpb.UpdateMemoryLayerMetadata{
		RequestTime:     now,
		OriginalRequest: req,
	}
	return s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		metadata.FinishTime = timestamppb.Now()
		return stored, nil
	})
}
