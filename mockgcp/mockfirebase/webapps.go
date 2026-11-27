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

package mockfirebase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/firebase/v1beta1"
	"github.com/golang/protobuf/ptypes/timestamp"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type webAppsServer struct {
	*MockService
	pb.UnimplementedProjectsWebAppsServerServer
}

type webAppName struct {
	Project *projects.ProjectData
	AppID   string
}

func (n *webAppName) String() string {
	return fmt.Sprintf("projects/%s/webApps/%s", n.Project.ID, n.AppID)
}

func (s *MockService) parseWebAppName(name string) (*webAppName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "webApps" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &webAppName{
			Project: project,
			AppID:   tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid web app name %q", name)
}

func (s *webAppsServer) GetProjectsWebApp(ctx context.Context, req *pb.GetProjectsWebAppRequest) (*pb.WebApp, error) {
	name, err := s.parseWebAppName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WebApp{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *webAppsServer) CreateProjectsWebApp(ctx context.Context, req *pb.CreateProjectsWebAppRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseProjectName(req.GetParent())
	if err != nil {
		return nil, err
	}

	app := req.GetProjectsWebApp()
	if app == nil {
		return nil, status.Errorf(codes.InvalidArgument, "projects_web_app is required")
	}

	h := sha256.Sum256([]byte(fmt.Sprintf("%s/%s", parent.Project.ID, app.GetDisplayName())))
	hashStr := hex.EncodeToString(h[:])[:22]
	appID := fmt.Sprintf("1:%d:web:%s", parent.Project.Number, hashStr)

	name := &webAppName{
		Project: parent.Project,
		AppID:   appID,
	}

	fqn := name.String()

	obj := proto.Clone(app).(*pb.WebApp)
	obj.Name = PtrTo(fqn)
	obj.AppId = PtrTo(appID)
	obj.ProjectId = PtrTo(parent.Project.ID)
	obj.State = PtrTo("ACTIVE")
	obj.Etag = PtrTo("abcdef0123A=")
	if obj.ApiKeyId == nil {
		obj.ApiKeyId = PtrTo("647489ca-09d1-4295-b1c5-7d156b5804e5")
	}
	obj.WebId = PtrTo("MGEzYWJkY2EtZDQ3OS00Y2VjLTgxNjktNzA1YTdhOTZjY2E4")
	obj.ExpireTime = &timestamp.Timestamp{
		Seconds: time.Now().Add(30 * 24 * time.Hour).Unix(),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.StartLRO(ctx, "operations/workflows/{{operationID}}", nil, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *webAppsServer) PatchProjectsWebApp(ctx context.Context, req *pb.PatchProjectsWebAppRequest) (*pb.WebApp, error) {
	name, err := s.parseWebAppName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WebApp{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	patch := req.GetProjectsWebApp()
	if patch != nil {
		if patch.DisplayName != nil {
			obj.DisplayName = patch.DisplayName
		}
		if patch.AppUrls != nil {
			obj.AppUrls = patch.AppUrls
		}
		if patch.ApiKeyId != nil {
			obj.ApiKeyId = patch.ApiKeyId
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *webAppsServer) DeleteProjectsWebApp(ctx context.Context, req *pb.DeleteProjectsWebAppRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseWebAppName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.WebApp{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadataAny, err := anypb.New(&pb.OperationMetadata{})
	if err != nil {
		return nil, err
	}
	metadataAny.TypeUrl = "type.googleapis.com/google.firebase.service.v1beta1.OperationMetadata"
	return &longrunningpb.Operation{
		Done:     true,
		Metadata: metadataAny,
	}, nil
}
