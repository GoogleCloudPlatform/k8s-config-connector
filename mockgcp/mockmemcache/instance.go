package mockmemcache

import (
	"context"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/memcache/apiv1beta2/memcachepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *cloudMemcacheServer) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Instance %s not found", name.String())
		}
		return nil, err
	}

	return obj, nil
}

func (s *cloudMemcacheServer) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Parent + "/instances/" + req.InstanceId)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Resource).(*pb.Instance)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.Instance_CREATING

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fqn
	metadata := &pb.OperationMetadata{}
	op, err := s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		obj.State = pb.Instance_READY
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (s *cloudMemcacheServer) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Resource.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// For simple mock, just update the whole object
	updatedObj := proto.Clone(req.Resource).(*pb.Instance)
	updatedObj.CreateTime = obj.CreateTime
	updatedObj.UpdateTime = timestamppb.Now()
	updatedObj.State = pb.Instance_READY

	lroPrefix := fqn
	metadata := &pb.OperationMetadata{}
	op, err := s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		if err := s.storage.Update(ctx, fqn, updatedObj); err != nil {
			return nil, err
		}
		return updatedObj, nil
	})
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (s *cloudMemcacheServer) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fqn
	metadata := &pb.OperationMetadata{}
	op, err := s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		if err := s.storage.Delete(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return &anypb.Any{}, nil
	})
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (s *cloudMemcacheServer) UpdateParameters(ctx context.Context, req *pb.UpdateParametersRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParameters not implemented")
}

func (s *cloudMemcacheServer) ApplyParameters(ctx context.Context, req *pb.ApplyParametersRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyParameters not implemented")
}

func (s *cloudMemcacheServer) RescheduleMaintenance(ctx context.Context, req *pb.RescheduleMaintenanceRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RescheduleMaintenance not implemented")
}

