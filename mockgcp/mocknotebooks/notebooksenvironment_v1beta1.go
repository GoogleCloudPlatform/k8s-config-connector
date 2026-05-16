package mocknotebooks

import (
        "context"
        "fmt"
        "time"

        "google.golang.org/grpc/codes"
        "google.golang.org/grpc/status"
        "google.golang.org/protobuf/proto"
        "google.golang.org/protobuf/types/known/emptypb"
        "google.golang.org/protobuf/types/known/timestamppb"

        pb_v1beta1 "cloud.google.com/go/notebooks/apiv1beta1/notebookspb"
        longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NotebookServiceV1beta1) GetEnvironment(ctx context.Context, req *pb_v1beta1.GetEnvironmentRequest) (*pb_v1beta1.Environment, error) {
        name, err := s.parseEnvironmentName(req.GetName())
        if err != nil {
                return nil, err
        }
        fqn := name.String()

        obj := &pb_v1beta1.Environment{}
        if err := s.storage.Get(ctx, fqn, obj); err != nil {
                if status.Code(err) == codes.NotFound {
                        return nil, status.Errorf(codes.NotFound, "environment %q not found", fqn)
                }
                return nil, err
        }

        return obj, nil
}

func (s *NotebookServiceV1beta1) CreateEnvironment(ctx context.Context, req *pb_v1beta1.CreateEnvironmentRequest) (*longrunningpb.Operation, error) {
        reqName := fmt.Sprintf("%s/environments/%s", req.GetParent(), req.GetEnvironmentId())
        name, err := s.parseEnvironmentName(reqName)
        if err != nil {
                return nil, err
        }

        fqn := name.String()
        obj := proto.Clone(req.GetEnvironment()).(*pb_v1beta1.Environment)
        obj.Name = fqn
        obj.CreateTime = timestamppb.New(time.Now())

        if err := s.storage.Create(ctx, fqn, obj); err != nil {
                return nil, err
        }

        prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
        metadata := &pb_v1beta1.OperationMetadata{
                ApiVersion:            "v1beta1",
                CreateTime:            timestamppb.New(time.Now()),
                RequestedCancellation: false,
                Target:                name.String(),
                Verb:                  "create",
                Endpoint:              "CreateEnvironment",
        }
        return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
                metadata.EndTime = timestamppb.New(time.Now())
                return obj, nil
        })
}

func (s *NotebookServiceV1beta1) DeleteEnvironment(ctx context.Context, req *pb_v1beta1.DeleteEnvironmentRequest) (*longrunningpb.Operation, error) {
        name, err := s.parseEnvironmentName(req.GetName())
        if err != nil {
                return nil, err
        }
        fqn := name.String()

        deletedObj := &pb_v1beta1.Environment{}
        if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
                return nil, err
        }

        prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
        metadata := &pb_v1beta1.OperationMetadata{
                ApiVersion:            "v1beta1",
                CreateTime:            timestamppb.Now(),
                RequestedCancellation: false,
                Target:                name.String(),
                Verb:                  "delete",
                Endpoint:              "DeleteEnvironment",
        }
        return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
                metadata.EndTime = timestamppb.New(time.Now())
                return &emptypb.Empty{}, nil
        })
}
