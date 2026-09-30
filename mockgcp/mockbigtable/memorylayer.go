package mockbigtable

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
)

func (s *instanceAdminServer) GetMemoryLayer(ctx context.Context, req *pb.GetMemoryLayerRequest) (*pb.MemoryLayer, error) {
	name, err := s.parseMemoryLayerName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MemoryLayer{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "MemoryLayer %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *instanceAdminServer) UpdateMemoryLayer(ctx context.Context, req *pb.UpdateMemoryLayerRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseMemoryLayerName(req.MemoryLayer.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MemoryLayer{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// If not found, we create it.
			obj = proto.CloneOf(req.MemoryLayer).(*pb.MemoryLayer)
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// Update existing
		paths := req.GetUpdateMask().GetPaths()
		for _, path := range paths {
			if path == "memory_config" || path == "*" {
				obj.MemoryConfig = req.MemoryLayer.GetMemoryConfig()
			}
		}

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	now := time.Now()
	// Pick some zone string, this might just be the cluster's zone, but name only has project, instance, cluster
	// prefix format ops
	zone := "us-central1-a"
	prefix := fmt.Sprintf("operations/projects/%s/instances/%s/clusters/%s/locations/%s", name.Project, name.Instance, name.Cluster, zone)
	metadata := &pb.UpdateMemoryLayerMetadata{
		RequestTime:     timestamppb.New(now),
		OriginalRequest: req,
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.FinishTime = timestamppb.Now()
		return obj, nil
	})
}

// memoryLayerName represents a parsed MemoryLayer FQN
type memoryLayerName struct {
	Project     string
	Instance    string
	Cluster     string
	MemoryLayer string
}

func (n *memoryLayerName) String() string {
	return "projects/" + n.Project + "/instances/" + n.Instance + "/clusters/" + n.Cluster + "/memoryLayer"
}

// parseMemoryLayerName parses a string into a memoryLayerName.
// The expected form is projects/<project>/instances/<instance>/clusters/<cluster>/memoryLayer
func (s *MockService) parseMemoryLayerName(name string) (*memoryLayerName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "clusters" && tokens[6] == "memoryLayer" {
		return &memoryLayerName{
			Project:     tokens[1],
			Instance:    tokens[3],
			Cluster:     tokens[5],
			MemoryLayer: tokens[7], // usually empty or we might not need it, the url is simply .../memoryLayer, wait, wait...
		}, nil
	} else if len(tokens) == 7 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "clusters" && tokens[6] == "memoryLayer" {
		return &memoryLayerName{
			Project:  tokens[1],
			Instance: tokens[3],
			Cluster:  tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid memoryLayer name", name)
}
