// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type RegionalTargetHTTPSProxiesV1 struct {
	*MockService
	pb.UnimplementedRegionTargetHttpsProxiesServer
}

func (s *RegionalTargetHTTPSProxiesV1) Get(ctx context.Context, req *pb.GetRegionTargetHttpsProxyRequest) (*pb.TargetHttpsProxy, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseRegionalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetHttpsProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RegionalTargetHTTPSProxiesV1) Insert(ctx context.Context, req *pb.InsertRegionTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpsProxies/" + req.GetTargetHttpsProxyResource().GetName()
	name, err := s.parseRegionalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetHttpsProxyResource()).(*pb.TargetHttpsProxy)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#targetHttpsProxy")

	if obj.Fingerprint == nil {
		obj.Fingerprint = PtrTo(computeFingerprint(obj))
	}

	if obj.SslCertificates != nil {
		var certs []string
		for _, cert := range obj.GetSslCertificates() {
			// todo(yuhou): this is a strange design of TF/GCP API.
			// GCP field `sslCertificates` refers to SSL Certificate resource or Certificate Manager Certificate resource,
			// Mixing Classic Certificates and Certificate Manager Certificates is not allowed.
			// TF handled it by adding a new field `certificateManagerCertificates` and using `conflictWith` to avoid the mixed values.
			// ref: https://github.com/hashicorp/terraform-provider-google/blob/31e35e8baaee132be5e25cd5d4740b9ac920dd57/google/services/compute/resource_compute_target_https_proxy.go#L1073s
			if strings.Contains(cert, "certificates") {
				cert := strings.TrimPrefix(cert, "https://certificatemanager.googleapis.com/v1/")
				tokens := strings.Split(cert, "/")
				if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "certificates" {
				} else {
					return nil, status.Errorf(codes.InvalidArgument, "certificateManagerCertificate %q is not valid", cert)
				}
				certs = append(certs, fmt.Sprintf("//certificatemanager.googleapis.com/projects/%s/locations/%s/certificates/%s", tokens[1], tokens[3], tokens[5]))

			} else {
				sslCertName, err := s.parseRegionalSslCertificateName(cert)
				if err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "sslCertName %q is not valid", sslCertName)
				}
				certs = append(certs, buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/sslCertificates/%s", sslCertName.Project.ID, sslCertName.Region, sslCertName.Name)))
			}
			obj.SslCertificates = certs
		}
	}

	if obj.UrlMap != nil {
		mapName, err := s.parseRegionalUrlMapName(obj.GetUrlMap())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "mapName %q is not valid", mapName)
		}
		obj.UrlMap = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/urlMaps/%s", mapName.Project.ID, mapName.Region, mapName.Name)))
	}
	obj.Region = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))

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

func (s *RegionalTargetHTTPSProxiesV1) SetUrlMap(ctx context.Context, req *pb.SetUrlMapRegionTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseRegionalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetHttpsProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.UrlMap = req.GetUrlMapReferenceResource().UrlMap

	if obj.UrlMap != nil {
		mapName, err := s.parseRegionalUrlMapName(obj.GetUrlMap())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "mapName %q is not valid", mapName)
		}
		obj.UrlMap = PtrTo(buildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s/urlMaps/%s", mapName.Project.ID, mapName.Region, mapName.Name)))
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("SetUrlMap"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionalTargetHTTPSProxiesV1) Delete(ctx context.Context, req *pb.DeleteRegionTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseRegionalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TargetHttpsProxy{}
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

type regionalTargetHttpsProxyName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalTargetHttpsProxyName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/targetHttpsProxies/" + n.Name
}

// parseRegionalTargetHttpsProxyName parses a string into a targethttpsproxyName.
// The expected form is `projects/*/regions/*/targethttpsproxy/*`.
func (s *MockService) parseRegionalTargetHttpsProxyName(name string) (*regionalTargetHttpsProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "targetHttpsProxies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalTargetHttpsProxyName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
