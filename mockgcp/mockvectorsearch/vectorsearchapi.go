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

package mockvectorsearch

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/vectorsearch/apiv1/vectorsearchpb"
)

type VectorSearchServer struct {
	*MockService
	pb.UnimplementedVectorSearchServiceServer
}

func (s *VectorSearchServer) GetCollection(ctx context.Context, req *pb.GetCollectionRequest) (*pb.Collection, error) {
	name, err := s.parseCollectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Collection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VectorSearchServer) CreateCollection(ctx context.Context, req *pb.CreateCollectionRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/collections/" + req.CollectionId
	name, err := s.parseCollectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Collection).(*pb.Collection)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := s.constructOperationMetadata(fqn, "create")
	return s.operations.StartLRO(ctx, req.Parent, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Collection)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *VectorSearchServer) UpdateCollection(ctx context.Context, req *pb.UpdateCollectionRequest) (*longrunning.Operation, error) {
	reqName := req.GetCollection().GetName()
	name, err := s.parseCollectionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Collection{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) > 0 {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetCollection().GetDisplayName()
			case "description":
				obj.Description = req.GetCollection().GetDescription()
			case "labels":
				obj.Labels = req.GetCollection().GetLabels()
			case "vector_schema", "vectorSchema":
				obj.VectorSchema = req.GetCollection().GetVectorSchema()
			case "data_schema", "dataSchema":
				obj.DataSchema = req.GetCollection().GetDataSchema()
			case "encryption_spec", "encryptionSpec":
				obj.EncryptionSpec = req.GetCollection().GetEncryptionSpec()
			}
		}
	} else {
		// If no update mask, replace fields
		obj.DisplayName = req.GetCollection().GetDisplayName()
		obj.Description = req.GetCollection().GetDescription()
		obj.Labels = req.GetCollection().GetLabels()
		obj.VectorSchema = req.GetCollection().GetVectorSchema()
		obj.DataSchema = req.GetCollection().GetDataSchema()
		obj.EncryptionSpec = req.GetCollection().GetEncryptionSpec()
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := s.constructOperationMetadata(fqn, "update")
	return s.operations.StartLRO(ctx, "projects/"+name.Project.ID+"/locations/"+name.Location, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Collection)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *VectorSearchServer) DeleteCollection(ctx context.Context, req *pb.DeleteCollectionRequest) (*longrunning.Operation, error) {
	name, err := s.parseCollectionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Collection{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	metadata := s.constructOperationMetadata(fqn, "delete")
	return s.operations.StartLRO(ctx, "projects/"+name.Project.ID+"/locations/"+name.Location, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *VectorSearchServer) GetIndex(ctx context.Context, req *pb.GetIndexRequest) (*pb.Index, error) {
	name, err := s.parseIndexName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Index{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *VectorSearchServer) CreateIndex(ctx context.Context, req *pb.CreateIndexRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/indexes/" + req.IndexId
	name, err := s.parseIndexName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Index).(*pb.Index)
	obj.Name = fqn

	s.populateIndexDefaults(obj)

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := s.constructOperationMetadata(fqn, "create")
	return s.operations.StartLRO(ctx, name.ParentString(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Index)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *VectorSearchServer) UpdateIndex(ctx context.Context, req *pb.UpdateIndexRequest) (*longrunning.Operation, error) {
	reqName := req.GetIndex().GetName()
	name, err := s.parseIndexName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Index{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) > 0 {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetIndex().GetDisplayName()
			case "description":
				obj.Description = req.GetIndex().GetDescription()
			case "dedicated_infrastructure", "dedicatedInfrastructure":
				if req.GetIndex().GetDedicatedInfrastructure() != nil {
					obj.InfraType = &pb.Index_DedicatedInfrastructure{
						DedicatedInfrastructure: req.GetIndex().GetDedicatedInfrastructure(),
					}
				}
			case "dense_scann", "denseScann":
				if req.GetIndex().GetDenseScann() != nil {
					obj.IndexType = &pb.Index_DenseScann{
						DenseScann: req.GetIndex().GetDenseScann(),
					}
				}
			case "distance_metric", "distanceMetric":
				obj.DistanceMetric = req.GetIndex().GetDistanceMetric()
			case "index_field", "indexField":
				obj.IndexField = req.GetIndex().GetIndexField()
			case "filter_fields", "filterFields":
				obj.FilterFields = req.GetIndex().GetFilterFields()
			case "store_fields", "storeFields":
				obj.StoreFields = req.GetIndex().GetStoreFields()
			}
		}
	} else {
		obj.DisplayName = req.GetIndex().GetDisplayName()
		obj.Description = req.GetIndex().GetDescription()
		if req.GetIndex().GetDedicatedInfrastructure() != nil {
			obj.InfraType = &pb.Index_DedicatedInfrastructure{
				DedicatedInfrastructure: req.GetIndex().GetDedicatedInfrastructure(),
			}
		}
		if req.GetIndex().GetDenseScann() != nil {
			obj.IndexType = &pb.Index_DenseScann{
				DenseScann: req.GetIndex().GetDenseScann(),
			}
		}
		obj.DistanceMetric = req.GetIndex().GetDistanceMetric()
		obj.IndexField = req.GetIndex().GetIndexField()
		obj.FilterFields = req.GetIndex().GetFilterFields()
		obj.StoreFields = req.GetIndex().GetStoreFields()
	}

	s.populateIndexDefaults(obj)
	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := s.constructOperationMetadata(fqn, "update")
	return s.operations.StartLRO(ctx, name.ParentString(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Index)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *VectorSearchServer) DeleteIndex(ctx context.Context, req *pb.DeleteIndexRequest) (*longrunning.Operation, error) {
	name, err := s.parseIndexName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Index{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	metadata := s.constructOperationMetadata(fqn, "delete")
	return s.operations.StartLRO(ctx, name.ParentString(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *VectorSearchServer) populateIndexDefaults(obj *pb.Index) {
	if obj.GetDedicatedInfrastructure() == nil {
		mode := pb.DedicatedInfrastructure_PERFORMANCE_OPTIMIZED
		obj.InfraType = &pb.Index_DedicatedInfrastructure{
			DedicatedInfrastructure: &pb.DedicatedInfrastructure{
				Mode: &mode,
				AutoscalingSpec: &pb.DedicatedInfrastructure_AutoscalingSpec{
					MinReplicaCount: 2,
					MaxReplicaCount: 2,
				},
			},
		}
	} else {
		infra := obj.GetDedicatedInfrastructure()
		if infra.Mode == nil {
			mode := pb.DedicatedInfrastructure_PERFORMANCE_OPTIMIZED
			infra.Mode = &mode
		}
		if infra.AutoscalingSpec == nil {
			infra.AutoscalingSpec = &pb.DedicatedInfrastructure_AutoscalingSpec{
				MinReplicaCount: 2,
				MaxReplicaCount: 2,
			}
		} else {
			if infra.AutoscalingSpec.MinReplicaCount == 0 {
				infra.AutoscalingSpec.MinReplicaCount = 2
			}
			if infra.AutoscalingSpec.MaxReplicaCount == 0 {
				infra.AutoscalingSpec.MaxReplicaCount = 2
			}
		}
	}

	if obj.DistanceMetric == pb.DistanceMetric_DISTANCE_METRIC_UNSPECIFIED {
		obj.DistanceMetric = pb.DistanceMetric_DOT_PRODUCT
	}

	if obj.IndexType == nil {
		obj.IndexType = &pb.Index_DenseScann{
			DenseScann: &pb.DenseScannIndex{},
		}
	}
}

func (s *MockService) constructOperationMetadata(target string, verb string) *pb.OperationMetadata {
	return &pb.OperationMetadata{
		CreateTime: timestamppb.Now(),
		Target:     target,
		Verb:       verb,
		ApiVersion: "v1",
	}
}
