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

package mockcompute

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func getPresentFields(ctx context.Context) map[string]bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}
	values := md["x-gcp-present-fields"]
	if len(values) == 0 {
		return nil
	}
	presentFields := make(map[string]bool)
	for _, val := range strings.Split(values[0], ",") {
		presentFields[val] = true
	}
	return presentFields
}

type RoutersV1 struct {
	*MockService
	pb.UnimplementedRoutersServer
}

func (s *RoutersV1) Get(ctx context.Context, req *pb.GetRouterRequest) (*pb.Router, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouter()
	name, err := s.parseRouterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Router{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RoutersV1) Insert(ctx context.Context, req *pb.InsertRouterRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouterResource().GetName()
	name, err := s.parseRouterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.CloneOf(req.GetRouterResource())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#router")

	if obj.Description == nil {
		obj.Description = PtrTo("")
	}

	if obj.EncryptedInterconnectRouter == nil {
		obj.EncryptedInterconnectRouter = PtrTo(false)
	}

	if obj.Network != nil {
		networkName, err := s.parseNetworkSelfLink(obj.GetNetwork())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "network %q is not valid", obj.GetNetwork())
		}
		obj.Network = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/networks/%s", networkName.Project.ID, networkName.Name)))
	}

	// output only fields
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))

	s.populateRouter(obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RoutersV1) Patch(ctx context.Context, req *pb.PatchRouterRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouter()
	name, err := s.parseRouterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Router{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	reqResource := req.GetRouterResource()

	hasNats := false
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get(httpmux.MetadataKeyHttpRequestQuery)
		if len(values) > 0 {
			q, _ := url.ParseQuery(values[0])
			if q.Get("natsSpecified") == "true" {
				hasNats = true
			}
		}
	}
	if !hasNats {
		hasNats = req.GetRouterResource().Nats != nil
	}

	presentFields := getPresentFields(ctx)

	// Save repeated fields and clear them from reqResource so proto.Merge doesn't append them
	reqNats := reqResource.Nats
	reqResource.Nats = nil

	reqInterfaces := reqResource.Interfaces
	reqResource.Interfaces = nil

	reqBgpPeers := reqResource.BgpPeers
	reqResource.BgpPeers = nil

	reqMd5 := reqResource.Md5AuthenticationKeys
	reqResource.Md5AuthenticationKeys = nil

	proto.Merge(obj, reqResource)

	// Restore them in reqResource
	reqResource.Nats = reqNats
	reqResource.Interfaces = reqInterfaces
	reqResource.BgpPeers = reqBgpPeers
	reqResource.Md5AuthenticationKeys = reqMd5

	// Manually replace repeated fields if present
	hasNatsField := false
	if presentFields != nil {
		hasNatsField = presentFields["nats"]
	} else {
		hasNatsField = hasNats
	}
	if hasNatsField {
		obj.Nats = reqNats
	}
	hasInterfacesField := false
	if presentFields != nil {
		hasInterfacesField = presentFields["interfaces"]
	} else {
		hasInterfacesField = reqInterfaces != nil
	}
	if hasInterfacesField {
		obj.Interfaces = reqInterfaces
	}

	hasBgpPeersField := false
	if presentFields != nil {
		hasBgpPeersField = presentFields["bgp_peers"]
	} else {
		hasBgpPeersField = reqBgpPeers != nil
	}
	if hasBgpPeersField {
		obj.BgpPeers = reqBgpPeers
	}

	hasMd5Field := false
	if presentFields != nil {
		hasMd5Field = presentFields["md5_authentication_keys"]
	} else {
		hasMd5Field = reqMd5 != nil
	}
	if hasMd5Field {
		obj.Md5AuthenticationKeys = reqMd5
	}

	s.populateRouter(obj)
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RoutersV1) Update(ctx context.Context, req *pb.UpdateRouterRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouter()
	name, err := s.parseRouterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Router{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(req.GetRouterResource()).(*pb.Router)
	updated.SelfLink = existing.SelfLink
	updated.CreationTimestamp = existing.CreationTimestamp
	updated.Id = existing.Id
	updated.Kind = existing.Kind
	updated.Region = existing.Region

	s.populateRouter(updated)
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      updated.Id,
		TargetLink:    updated.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return updated, nil
	})
}

func (s *RoutersV1) Delete(ctx context.Context, req *pb.DeleteRouterRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/routers/" + req.GetRouter()
	name, err := s.parseRouterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Router{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type routerName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *routerName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/routers/" + n.Name
}

func (s *MockService) parseRouterName(name string) (*routerName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "routers" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &routerName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
func (s *RoutersV1) populateRouter(obj *pb.Router) {
	for _, iface := range obj.Interfaces {
		if iface.IpVersion == nil {
			iface.IpVersion = PtrTo("IPV4")
		}
	}
	for _, nat := range obj.Nats {
		if nat.Type == nil {
			nat.Type = PtrTo("PUBLIC")
		}
	}
}
