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
// proto.service: google.cloud.networkservices.v1.NetworkServices
// proto.message: google.cloud.networkservices.v1.WasmPlugin

package mocknetworkservices

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NetworkServicesServer) GetWasmPlugin(ctx context.Context, req *pb.GetWasmPluginRequest) (*pb.WasmPlugin, error) {
	name, err := s.parseWasmPluginName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WasmPlugin{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) ListWasmPlugins(ctx context.Context, req *pb.ListWasmPluginsRequest) (*pb.ListWasmPluginsResponse, error) {
	response := &pb.ListWasmPluginsResponse{}

	parent, err := s.parseWasmPluginParent(req.Parent)
	if err != nil {
		return nil, err
	}
	prefix := parent.String() + "/wasmPlugins/"

	findKind := (&pb.WasmPlugin{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb.WasmPlugin)
		response.WasmPlugins = append(response.WasmPlugins, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *NetworkServicesServer) CreateWasmPlugin(ctx context.Context, req *pb.CreateWasmPluginRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/wasmPlugins/" + req.WasmPluginId
	name, err := s.parseWasmPluginName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.WasmPlugin)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	// In CreateWasmPlugin, versions can be provided in the map
	for versionID, details := range obj.Versions {
		versionName := fqn + "/versions/" + versionID
		version := &pb.WasmPluginVersion{
			Name:        versionName,
			Description: details.Description,
			Labels:      details.Labels,
			ImageUri:    details.ImageUri,
			ImageDigest: "sha256:abcdef1234567890", // Mock digest
			CreateTime:  timestamppb.New(now),
			UpdateTime:  timestamppb.New(now),
		}
		if data := details.GetPluginConfigData(); data != nil {
			version.PluginConfigSource = &pb.WasmPluginVersion_PluginConfigData{PluginConfigData: data}
			version.PluginConfigDigest = "sha256:config1234567890"
		} else if uri := details.GetPluginConfigUri(); uri != "" {
			version.PluginConfigSource = &pb.WasmPluginVersion_PluginConfigUri{PluginConfigUri: uri}
			version.PluginConfigDigest = "sha256:config1234567890"
		}

		if err := s.storage.Create(ctx, versionName, version); err != nil {
			return nil, err
		}
	}
	// Clear versions from the main object to match real GCP behavior (versions are separate resources)
	obj.Versions = nil

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NetworkServicesServer) UpdateWasmPlugin(ctx context.Context, req *pb.UpdateWasmPluginRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetWasmPlugin().GetName()

	name, err := s.parseWasmPluginName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.WasmPlugin{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		req.WasmPlugin.CreateTime = obj.CreateTime
		req.WasmPlugin.UpdateTime = timestamppb.New(now)
		req.WasmPlugin.Name = obj.Name
		obj = req.WasmPlugin
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetWasmPlugin().GetLabels()
			case "description":
				obj.Description = req.GetWasmPlugin().GetDescription()
			case "logConfig", "log_config":
				obj.LogConfig = req.GetWasmPlugin().GetLogConfig()
			case "mainVersionId", "main_version_id":
				obj.MainVersionId = req.GetWasmPlugin().GetMainVersionId()
			case "versions":
				// Replace all versions
				// 1. Collect existing versions
				var versionNames []string
				prefix := fqn + "/versions/"
				findKind := (&pb.WasmPluginVersion{}).ProtoReflect().Descriptor()
				if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
					versionNames = append(versionNames, msg.(*pb.WasmPluginVersion).Name)
					return nil
				}); err != nil {
					return nil, err
				}
				// 2. Delete existing
				for _, versionName := range versionNames {
					if err := s.storage.Delete(ctx, versionName, &pb.WasmPluginVersion{}); err != nil {
						return nil, err
					}
				}

				// 3. Create new
				for versionID, details := range req.GetWasmPlugin().GetVersions() {
					versionName := fqn + "/versions/" + versionID
					version := &pb.WasmPluginVersion{
						Name:        versionName,
						Description: details.Description,
						Labels:      details.Labels,
						ImageUri:    details.ImageUri,
						ImageDigest: "sha256:abcdef1234567890",
						CreateTime:  timestamppb.New(now),
						UpdateTime:  timestamppb.New(now),
					}
					if data := details.GetPluginConfigData(); data != nil {
						version.PluginConfigSource = &pb.WasmPluginVersion_PluginConfigData{PluginConfigData: data}
						version.PluginConfigDigest = "sha256:config1234567890"
					} else if uri := details.GetPluginConfigUri(); uri != "" {
						version.PluginConfigSource = &pb.WasmPluginVersion_PluginConfigUri{PluginConfigUri: uri}
						version.PluginConfigDigest = "sha256:config1234567890"
					}

					if err := s.storage.Create(ctx, versionName, version); err != nil {
						return nil, err
					}
				}
				// obj.Versions = req.GetWasmPlugin().GetVersions() // Versions not stored in main object
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NetworkServicesServer) DeleteWasmPlugin(ctx context.Context, req *pb.DeleteWasmPluginRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseWasmPluginName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.WasmPlugin{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// Delete versions
	// 1. Collect versions
	var versionNames []string
	prefix := fqn + "/versions/"
	findKind := (&pb.WasmPluginVersion{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		versionNames = append(versionNames, msg.(*pb.WasmPluginVersion).Name)
		return nil
	}); err != nil {
		return nil, err
	}
	// 2. Delete versions
	for _, versionName := range versionNames {
		if err := s.storage.Delete(ctx, versionName, &pb.WasmPluginVersion{}); err != nil {
			return nil, err
		}
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := &emptypb.Empty{}
		return result, nil
	})
}

func (s *NetworkServicesServer) GetWasmPluginVersion(ctx context.Context, req *pb.GetWasmPluginVersionRequest) (*pb.WasmPluginVersion, error) {
	name, err := s.parseWasmPluginVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WasmPluginVersion{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) ListWasmPluginVersions(ctx context.Context, req *pb.ListWasmPluginVersionsRequest) (*pb.ListWasmPluginVersionsResponse, error) {
	response := &pb.ListWasmPluginVersionsResponse{}

	parent, err := s.parseWasmPluginName(req.Parent)
	if err != nil {
		return nil, err
	}
	prefix := parent.String() + "/versions/"

	findKind := (&pb.WasmPluginVersion{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb.WasmPluginVersion)
		response.WasmPluginVersions = append(response.WasmPluginVersions, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *NetworkServicesServer) CreateWasmPluginVersion(ctx context.Context, req *pb.CreateWasmPluginVersionRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/versions/" + req.WasmPluginVersionId
	name, err := s.parseWasmPluginVersionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.WasmPluginVersion)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.ImageDigest = "sha256:abcdef1234567890"
	if obj.PluginConfigSource != nil {
		obj.PluginConfigDigest = "sha256:config1234567890"
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NetworkServicesServer) DeleteWasmPluginVersion(ctx context.Context, req *pb.DeleteWasmPluginVersionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseWasmPluginVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.WasmPluginVersion{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := &emptypb.Empty{}
		return result, nil
	})
}

type wasmPluginParent struct {
	Project  *projects.ProjectData
	Location string
}

func (p *wasmPluginParent) String() string {
	return "projects/" + p.Project.ID + "/locations/" + p.Location
}

func (s *NetworkServicesServer) parseWasmPluginParent(parent string) (*wasmPluginParent, error) {
	tokens := strings.Split(parent, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &wasmPluginParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}
}

type wasmPluginName struct {
	Project        *projects.ProjectData
	Location       string
	WasmPluginName string
}

func (n *wasmPluginName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/wasmPlugins/" + n.WasmPluginName
}

func (s *NetworkServicesServer) parseWasmPluginName(name string) (*wasmPluginName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "wasmPlugins" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &wasmPluginName{
			Project:        project,
			Location:       tokens[3],
			WasmPluginName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

type wasmPluginVersionName struct {
	Project               *projects.ProjectData
	Location              string
	WasmPluginName        string
	WasmPluginVersionName string
}

func (n *wasmPluginVersionName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/wasmPlugins/" + n.WasmPluginName + "/versions/" + n.WasmPluginVersionName
}

func (s *NetworkServicesServer) parseWasmPluginVersionName(name string) (*wasmPluginVersionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "wasmPlugins" && tokens[6] == "versions" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &wasmPluginVersionName{
			Project:               project,
			Location:              tokens[3],
			WasmPluginName:        tokens[5],
			WasmPluginVersionName: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
