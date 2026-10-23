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

package mockhealthcare

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/healthcare/v1"
)

type consentStoresServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsDatasetsConsentStoresServerServer
}

func (s *consentStoresServer) GetProjectsLocationsDatasetsConsentStore(ctx context.Context, req *pb.GetProjectsLocationsDatasetsConsentStoreRequest) (*pb.ConsentStore, error) {
	name, err := s.parseConsentStoreName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConsentStore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *consentStoresServer) CreateProjectsLocationsDatasetsConsentStore(ctx context.Context, req *pb.CreateProjectsLocationsDatasetsConsentStoreRequest) (*pb.ConsentStore, error) {
	reqName := req.GetParent() + "/consentStores/" + req.GetConsentStoreId()
	name, err := s.parseConsentStoreName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetProjectsLocationsDatasetsConsentStore()).(*pb.ConsentStore)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *consentStoresServer) DeleteProjectsLocationsDatasetsConsentStore(ctx context.Context, req *pb.DeleteProjectsLocationsDatasetsConsentStoreRequest) (*pb.Empty, error) {
	name, err := s.parseConsentStoreName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ConsentStore{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *consentStoresServer) PatchProjectsLocationsDatasetsConsentStore(ctx context.Context, req *pb.PatchProjectsLocationsDatasetsConsentStoreRequest) (*pb.ConsentStore, error) {
	name, err := s.parseConsentStoreName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConsentStore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	if updateMask == "" {
		proto.Merge(obj, req.GetProjectsLocationsDatasetsConsentStore())
	} else {
		for _, path := range strings.Split(updateMask, ",") {
			switch path {
			case "defaultConsentTtl":
				obj.DefaultConsentTtl = req.GetProjectsLocationsDatasetsConsentStore().GetDefaultConsentTtl()
			case "enableConsentCreateOnUpdate":
				obj.EnableConsentCreateOnUpdate = req.GetProjectsLocationsDatasetsConsentStore().GetEnableConsentCreateOnUpdate()
			case "labels":
				obj.Labels = req.GetProjectsLocationsDatasetsConsentStore().GetLabels()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "unsupported update_mask path %q", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
