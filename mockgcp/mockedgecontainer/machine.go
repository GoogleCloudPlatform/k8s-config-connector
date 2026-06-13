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

package mockedgecontainer

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

var machineKind = (&pb.Machine{}).ProtoReflect().Descriptor()

func (s *EdgeContainerV1) GetMachine(ctx context.Context, req *pb.GetMachineRequest) (*pb.Machine, error) {
	name, err := s.parseMachineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Machine{}
	err = s.storage.Get(ctx, fqn, obj)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			now := time.Now()
			obj = &pb.Machine{
				Name:       fqn,
				CreateTime: timestamppb.New(now.Add(-24 * time.Hour)),
				UpdateTime: timestamppb.New(now),
				Version:    "1.0.0",
				Zone:       "us-central1-a",
				Disabled:   false,
				Labels: map[string]string{
					"mock-label": "mock-value",
				},
			}
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}

	return obj, nil
}

func (s *EdgeContainerV1) ListMachines(ctx context.Context, req *pb.ListMachinesRequest) (*pb.ListMachinesResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}

	prefix := req.Parent + "/machines/"

	var machines []*pb.Machine
	err := s.storage.List(ctx, machineKind, storage.ListOptions{}, func(obj proto.Message) error {
		m := obj.(*pb.Machine)
		if strings.HasPrefix(m.Name, prefix) {
			machines = append(machines, m)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(machines) == 0 {
		defaultMachineName := prefix + "default-machine"
		now := time.Now()
		m := &pb.Machine{
			Name:       defaultMachineName,
			CreateTime: timestamppb.New(now.Add(-24 * time.Hour)),
			UpdateTime: timestamppb.New(now),
			Version:    "1.0.0",
			Zone:       "us-central1-a",
			Disabled:   false,
			Labels: map[string]string{
				"mock-label": "mock-value",
			},
		}
		if err := s.storage.Create(ctx, defaultMachineName, m); err != nil {
			return nil, err
		}
		machines = append(machines, m)
	}

	return &pb.ListMachinesResponse{
		Machines: machines,
	}, nil
}
