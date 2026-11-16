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

// +tool:mockgcp-support
// proto.service: google.appengine.v1.DomainMappings
// proto.message: google.appengine.v1.DomainMapping

package mockappengine

import (
	"context"
	"fmt"
	"strings"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type domainMappingsServer struct {
	*MockService
	pb.UnimplementedDomainMappingsServer
}

type domainMappingName struct {
	App           string
	DomainMapping string
}

func (n *domainMappingName) String() string {
	return fmt.Sprintf("apps/%s/domainMappings/%s", n.App, n.DomainMapping)
}

func parseDomainMappingName(name string) (*domainMappingName, error) {
	parts := strings.Split(name, "/")
	if len(parts) != 4 || parts[0] != "apps" || parts[2] != "domainMappings" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid domain mapping name: %q", name)
	}
	return &domainMappingName{
		App:           parts[1],
		DomainMapping: parts[3],
	}, nil
}

func parseDomainMappingParent(parent string) (string, error) {
	parts := strings.Split(parent, "/")
	if len(parts) != 2 || parts[0] != "apps" {
		return "", status.Errorf(codes.InvalidArgument, "invalid parent name: %q", parent)
	}
	return parts[1], nil
}

func (s *domainMappingsServer) GetDomainMapping(ctx context.Context, req *pb.GetDomainMappingRequest) (*pb.DomainMapping, error) {
	name, err := parseDomainMappingName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DomainMapping{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *domainMappingsServer) CreateDomainMapping(ctx context.Context, req *pb.CreateDomainMappingRequest) (*longrunningpb.Operation, error) {
	appID, err := parseDomainMappingParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	reqObj := req.GetDomainMapping()
	if reqObj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "domain_mapping is required")
	}

	domainID := reqObj.GetId()
	if domainID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "domain_mapping.id is required")
	}

	fqn := fmt.Sprintf("apps/%s/domainMappings/%s", appID, domainID)
	obj := proto.Clone(reqObj).(*pb.DomainMapping)
	obj.Name = fqn
	obj.Id = domainID

	if obj.SslSettings == nil {
		obj.SslSettings = &pb.SslSettings{
			SslManagementType: pb.SslSettings_AUTOMATIC,
		}
	}

	if len(obj.ResourceRecords) == 0 {
		obj.ResourceRecords = []*pb.ResourceRecord{
			{
				Name:   domainID,
				Rrdata: "ghs.googlehosted.com.",
				Type:   pb.ResourceRecord_CNAME,
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadataV1{
		Method:     "google.appengine.v1.DomainMappings.CreateDomainMapping",
		InsertTime: timestamppb.Now(),
		Target:     fqn,
		User:       fmt.Sprintf("user@%s.iam.gserviceaccount.com", appID),
	}
	prefix := fmt.Sprintf("apps/%s", appID)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *domainMappingsServer) UpdateDomainMapping(ctx context.Context, req *pb.UpdateDomainMappingRequest) (*longrunningpb.Operation, error) {
	name, err := parseDomainMappingName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DomainMapping{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	reqObj := req.GetDomainMapping()
	if reqObj != nil {
		if reqObj.SslSettings != nil {
			if obj.SslSettings == nil {
				obj.SslSettings = &pb.SslSettings{}
			}
			if reqObj.SslSettings.SslManagementType != pb.SslSettings_SSL_MANAGEMENT_TYPE_UNSPECIFIED {
				obj.SslSettings.SslManagementType = reqObj.SslSettings.SslManagementType
			}
			if reqObj.SslSettings.CertificateId != "" {
				obj.SslSettings.CertificateId = reqObj.SslSettings.CertificateId
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadataV1{
		Method:     "google.appengine.v1.DomainMappings.UpdateDomainMapping",
		InsertTime: timestamppb.Now(),
		Target:     fqn,
		User:       fmt.Sprintf("user@%s.iam.gserviceaccount.com", name.App),
	}
	prefix := fmt.Sprintf("apps/%s", name.App)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *domainMappingsServer) DeleteDomainMapping(ctx context.Context, req *pb.DeleteDomainMappingRequest) (*longrunningpb.Operation, error) {
	name, err := parseDomainMappingName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DomainMapping{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadataV1{
		Method:     "google.appengine.v1.DomainMappings.DeleteDomainMapping",
		InsertTime: timestamppb.Now(),
		Target:     fqn,
		User:       fmt.Sprintf("user@%s.iam.gserviceaccount.com", name.App),
	}
	prefix := fmt.Sprintf("apps/%s", name.App)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *domainMappingsServer) ListDomainMappings(ctx context.Context, req *pb.ListDomainMappingsRequest) (*pb.ListDomainMappingsResponse, error) {
	appID, err := parseDomainMappingParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	var mappings []*pb.DomainMapping
	prefix := fmt.Sprintf("apps/%s/domainMappings/", appID)
	findKind := (&pb.DomainMapping{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		item := msg.(*pb.DomainMapping)
		mappings = append(mappings, proto.Clone(item).(*pb.DomainMapping))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListDomainMappingsResponse{
		DomainMappings: mappings,
	}, nil
}
