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

type androidAppsServer struct {
	*MockService
	pb.UnimplementedProjectsAndroidAppsServerServer
}

type androidAppName struct {
	Project *projects.ProjectData
	AppID   string
}

func (n *androidAppName) String() string {
	return fmt.Sprintf("projects/%s/androidApps/%s", n.Project.ID, n.AppID)
}

func (s *MockService) parseAndroidAppName(name string) (*androidAppName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "androidApps" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &androidAppName{
			Project: project,
			AppID:   tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid android app name %q", name)
}

func (s *androidAppsServer) GetProjectsAndroidApp(ctx context.Context, req *pb.GetProjectsAndroidAppRequest) (*pb.AndroidApp, error) {
	name, err := s.parseAndroidAppName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AndroidApp{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *androidAppsServer) CreateProjectsAndroidApp(ctx context.Context, req *pb.CreateProjectsAndroidAppRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseProjectName(req.GetParent())
	if err != nil {
		return nil, err
	}

	app := req.GetProjectsAndroidApp()
	if app == nil {
		return nil, status.Errorf(codes.InvalidArgument, "projects_android_app is required")
	}

	h := sha256.Sum256([]byte(fmt.Sprintf("%s/%s", parent.Project.ID, app.GetDisplayName())))
	hashStr := hex.EncodeToString(h[:])[:22]
	appID := fmt.Sprintf("1:%d:android:%s", parent.Project.Number, hashStr)

	name := &androidAppName{
		Project: parent.Project,
		AppID:   appID,
	}

	fqn := name.String()

	obj := proto.Clone(app).(*pb.AndroidApp)
	obj.Name = PtrTo(fqn)
	obj.AppId = PtrTo(appID)
	obj.ProjectId = PtrTo(parent.Project.ID)
	obj.State = PtrTo("ACTIVE")
	obj.Etag = PtrTo("abcdef0123A=")
	if obj.ApiKeyId == nil {
		obj.ApiKeyId = PtrTo("6d1ee53e-9817-4331-ac63-7fdfba6d5e07")
	}
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

func (s *androidAppsServer) PatchProjectsAndroidApp(ctx context.Context, req *pb.PatchProjectsAndroidAppRequest) (*pb.AndroidApp, error) {
	name, err := s.parseAndroidAppName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AndroidApp{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	patch := req.GetProjectsAndroidApp()
	if patch != nil {
		if patch.DisplayName != nil {
			obj.DisplayName = patch.DisplayName
		}
		if patch.Sha1Hashes != nil {
			obj.Sha1Hashes = patch.Sha1Hashes
		}
		if patch.Sha256Hashes != nil {
			obj.Sha256Hashes = patch.Sha256Hashes
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

func (s *androidAppsServer) DeleteProjectsAndroidApp(ctx context.Context, req *pb.DeleteProjectsAndroidAppRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAndroidAppName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.AndroidApp{}
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
