// Copyright 2024 Google LLC
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

package mockmonitoring

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type UptimeCheckService struct {
	*MockService
	pb.UnimplementedUptimeCheckServiceServer
}

func (s *UptimeCheckService) GetUptimeCheckConfig(ctx context.Context, req *pb.GetUptimeCheckConfigRequest) (*pb.UptimeCheckConfig, error) {
	name, err := s.parseUptimeCheckConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UptimeCheckConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Config not found for check %s in project %s", name.Name, name.Project.ID)
		}
		return nil, err
	}

	return redactUptimeCheckConfig(obj), nil
}

func redactUptimeCheckConfig(obj *pb.UptimeCheckConfig) *pb.UptimeCheckConfig {
	// Fields containing sensitive information like authentication tokens or contact info are only partially populated on retrieval.
	redacted := proto.Clone(obj).(*pb.UptimeCheckConfig)
	if authInfo := redacted.GetHttpCheck().GetAuthInfo(); authInfo != nil {
		authInfo.Password = strings.Repeat("*", 6)
	}
	if headers := redacted.GetHttpCheck().GetHeaders(); headers != nil {
		for k := range headers {
			headers[k] = "******"
		}
	}
	return redacted
}

func populateDefaultsForUptimeCheckConfig(obj *pb.UptimeCheckConfig) {
	if obj.CheckerType == pb.UptimeCheckConfig_CHECKER_TYPE_UNSPECIFIED {
		obj.CheckerType = pb.UptimeCheckConfig_STATIC_IP_CHECKERS
	}

	if httpCheck := obj.GetHttpCheck(); httpCheck != nil {
		if httpCheck.Body != nil {
			// Users can provide a `Content-Length` header via the `headers` field or the API will do so.
			if httpCheck.Headers == nil {
				httpCheck.Headers = make(map[string]string)
			}
			foundContentLength := false
			for k := range httpCheck.Headers {
				if strings.ToLower(k) == "content-length" {
					foundContentLength = true
				}
			}
			if !foundContentLength {
				httpCheck.Headers["Content-Length"] = strconv.Itoa(len(httpCheck.Body))
			}
		}
	}
}
func (s *UptimeCheckService) CreateUptimeCheckConfig(ctx context.Context, req *pb.CreateUptimeCheckConfigRequest) (*pb.UptimeCheckConfig, error) {
	now := time.Now()

	uptimeCheckConfigID := fmt.Sprintf("%d", now.UnixNano())

	reqName := req.GetParent() + "/uptimeCheckConfigs/" + uptimeCheckConfigID
	name, err := s.parseUptimeCheckConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.UptimeCheckConfig).(*pb.UptimeCheckConfig)
	obj.Name = fqn

	populateDefaultsForUptimeCheckConfig(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating uptimeCheckConfig: %v", err)
	}

	return redactUptimeCheckConfig(obj), nil
}

func (s *UptimeCheckService) UpdateUptimeCheckConfig(ctx context.Context, req *pb.UpdateUptimeCheckConfigRequest) (*pb.UptimeCheckConfig, error) {
	name, err := s.parseUptimeCheckConfigName(req.GetUptimeCheckConfig().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.UptimeCheckConfig{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.UptimeCheckConfig)

	for _, path := range req.GetUpdateMask().GetPaths() {
		// TODO: Validate path?
		if err := setField(updated, req.GetUptimeCheckConfig(), path); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "error setting field %q: %v", path, err)
		}
	}

	populateDefaultsForUptimeCheckConfig(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating uptimeCheckConfig: %v", err)
	}

	return redactUptimeCheckConfig(updated), nil
}

func (s *UptimeCheckService) DeleteUptimeCheckConfig(ctx context.Context, req *pb.DeleteUptimeCheckConfigRequest) (*empty.Empty, error) {
	name, err := s.parseUptimeCheckConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.UptimeCheckConfig{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

type uptimeCheckConfigName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *uptimeCheckConfigName) String() string {
	return "projects/" + n.Project.ID + "/uptimeCheckConfigs/" + n.Name
}

// parseUptimeCheckConfigName parses a string into a uptimeCheckConfigName.
// The expected form is projects/[PROJECT_ID_OR_NUMBER]/uptimeCheckConfigs/[UPTIME_CHECK_ID]
func (s *MockService) parseUptimeCheckConfigName(name string) (*uptimeCheckConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "uptimeCheckConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &uptimeCheckConfigName{
			Project: project,
			Name:    tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
