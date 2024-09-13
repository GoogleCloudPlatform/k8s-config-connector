// Copyright 2022 Google LLC
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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type GlobalTargetHTTPSProxiesV1 struct {
	*MockService
	pb.UnimplementedTargetHttpsProxiesServer
}

func (s *GlobalTargetHTTPSProxiesV1) Get(ctx context.Context, req *pb.GetTargetHttpsProxyRequest) (*pb.TargetHttpsProxy, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
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

func (s *GlobalTargetHTTPSProxiesV1) Insert(ctx context.Context, req *pb.InsertTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxyResource().GetName()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetHttpsProxyResource()).(*pb.TargetHttpsProxy)
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + name.String())
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
				certs = append(certs, fmt.Sprintf("//certificatemanager.googleapis.com/projects/%s/locations/global/certificates/%s", tokens[1], tokens[5]))
			} else {
				sslCertName, err := s.parseGlobalSslCertificateName(cert)
				if err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "sslCertName %q is not valid", sslCertName)
				}
				certs = append(certs, fmt.Sprintf("https://www.googleapis.com/compute/beta/projects/%s/global/sslCertificates/%s", sslCertName.Project.ID, sslCertName.Name))
			}
			obj.SslCertificates = certs
		}
	}
	if obj.UrlMap != nil {
		mapName, err := s.parseGlobalUrlMapName(obj.GetUrlMap())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "mapName %q is not valid", mapName)
		}
		obj.UrlMap = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/beta/projects/%s/global/urlMaps/%s", mapName.Project.ID, mapName.Name))
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

// Updates a TargetHttpsProxy resource in the specified project using the data included in the request.
// This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (s *GlobalTargetHTTPSProxiesV1) Patch(ctx context.Context, req *pb.PatchTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxyResource().GetName()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetHttpsProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetTargetHttpsProxyResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalTargetHTTPSProxiesV1) SetUrlMap(ctx context.Context, req *pb.SetUrlMapTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetHttpsProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if obj.UrlMap != nil {
		mapName, err := s.parseGlobalUrlMapName(req.GetUrlMapReferenceResource().GetUrlMap())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "mapName %q is not valid", mapName)
		}
		obj.UrlMap = PtrTo(fmt.Sprintf("https://www.googleapis.com/compute/beta/projects/%s/global/urlMaps/%s", mapName.Project.ID, mapName.Name))
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
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalTargetHTTPSProxiesV1) SetQuicOverride(ctx context.Context, req *pb.SetQuicOverrideTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.TargetHttpsProxy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.QuicOverride = req.GetTargetHttpsProxiesSetQuicOverrideRequestResource().QuicOverride

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.targetHttpsProxies.setQuicOverride"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalTargetHTTPSProxiesV1) Delete(ctx context.Context, req *pb.DeleteTargetHttpsProxyRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/global" + "/targetHttpsProxies/" + req.GetTargetHttpsProxy()
	name, err := s.parseGlobalTargetHttpsProxyName(reqName)
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
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

type globalTargetHttpsProxyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalTargetHttpsProxyName) String() string {
	return "projects/" + n.Project.ID + "/global" + "/targetHttpsProxies/" + n.Name
}

// parseGlobalTargetHttpsProxyName parses a string into a globalTargetHttpsProxyName.
// The expected form is `projects/*/regions/*/targetHttpsproxy/*`.
func (s *MockService) parseGlobalTargetHttpsProxyName(name string) (*globalTargetHttpsProxyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "targetHttpsProxies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &globalTargetHttpsProxyName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
