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
// proto.service: google.appengine.v1.Versions

package mockappengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type versionsServer struct {
	*MockService
	pb.UnimplementedVersionsServer
}

type versionParent struct {
	App     string
	Service string
}

func (s *MockService) parseVersionParent(parent string) (*versionParent, error) {
	parts := strings.Split(parent, "/")
	if len(parts) == 4 && parts[0] == "apps" && parts[2] == "services" {
		return &versionParent{App: parts[1], Service: parts[3]}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid parent %q", parent)
}

type versionName struct {
	App     string
	Service string
	Version string
}

func (s *MockService) parseVersionName(name string) (*versionName, error) {
	parts := strings.Split(name, "/")
	if len(parts) == 6 && parts[0] == "apps" && parts[2] == "services" && parts[4] == "versions" {
		return &versionName{App: parts[1], Service: parts[3], Version: parts[5]}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid version name %q", name)
}

func (n *versionName) String() string {
	return fmt.Sprintf("apps/%s/services/%s/versions/%s", n.App, n.Service, n.Version)
}

func (s *versionsServer) GetVersion(ctx context.Context, req *pb.GetVersionRequest) (*pb.Version, error) {
	name, err := s.parseVersionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Version{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource 'projects/%s/locations/us-central1/applications/%s/services/%s/versions/%s' was not found", name.App, name.App, name.Service, name.Version)
		}
		return nil, err
	}

	result := proto.Clone(obj).(*pb.Version)
	result.Entrypoint = nil
	if req.GetView() != pb.VersionView_FULL {
		result.Handlers = nil
		result.ErrorHandlers = nil
		result.Libraries = nil
		result.ApiConfig = nil
		result.EnvVariables = nil
		result.BuildEnvVariables = nil
		result.DefaultExpiration = nil
		result.HealthCheck = nil
		result.ReadinessCheck = nil
		result.LivenessCheck = nil
		result.NobuildFilesRegex = ""
		result.Deployment = nil
	}

	return result, nil
}

func (s *versionsServer) CreateVersion(ctx context.Context, req *pb.CreateVersionRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseVersionParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	obj := proto.Clone(req.GetVersion()).(*pb.Version)
	versionID := obj.GetId()
	if versionID == "" {
		versionID = fmt.Sprintf("v%d", time.Now().UnixNano())
		obj.Id = versionID
	}

	name := &versionName{
		App:     parent.App,
		Service: parent.Service,
		Version: versionID,
	}
	fqn := name.String()

	now := time.Now()
	s.populateDefaultsForVersion(parent, name, obj, now)

	existing := &pb.Version{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		if len(req.GetVersion().GetHandlers()) > 0 && len(existing.GetHandlers()) > 0 {
			obj.Handlers = append(existing.GetHandlers(), obj.GetHandlers()...)
		}
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	opMetadata := &pb.OperationMetadataV1{
		Method:     "google.appengine.v1.Versions.CreateVersion",
		InsertTime: timestamppb.New(now),
		User:       fmt.Sprintf("overseer-kcc-tester@%s.iam.gserviceaccount.com", parent.App),
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("apps/%s", parent.App)
	return s.operations.StartLROWithDone(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(now)
		return proto.Clone(obj), nil
	})
}

func (s *versionsServer) UpdateVersion(ctx context.Context, req *pb.UpdateVersionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVersionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.Version{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Version)
	updateObj := req.GetVersion()

	if updateObj.GetInstanceClass() != "" {
		updated.InstanceClass = updateObj.GetInstanceClass()
	}
	if updateObj.GetAutomaticScaling() != nil {
		if updated.GetAutomaticScaling() == nil {
			updated.Scaling = &pb.Version_AutomaticScaling{AutomaticScaling: &pb.AutomaticScaling{}}
		}
		as := updated.GetAutomaticScaling()
		inAs := updateObj.GetAutomaticScaling()
		if inAs.GetMaxConcurrentRequests() != 0 {
			as.MaxConcurrentRequests = inAs.GetMaxConcurrentRequests()
		}
		if inAs.GetMaxIdleInstances() != 0 {
			as.MaxIdleInstances = inAs.GetMaxIdleInstances()
		}
		if inAs.GetMinIdleInstances() != 0 {
			as.MinIdleInstances = inAs.GetMinIdleInstances()
		}
		if inAs.GetMaxPendingLatency() != nil {
			as.MaxPendingLatency = inAs.GetMaxPendingLatency()
		}
		if inAs.GetMinPendingLatency() != nil {
			as.MinPendingLatency = inAs.GetMinPendingLatency()
		}
		if inAs.GetStandardSchedulerSettings() != nil {
			if as.StandardSchedulerSettings == nil {
				as.StandardSchedulerSettings = &pb.StandardSchedulerSettings{}
			}
			inSched := inAs.GetStandardSchedulerSettings()
			if inSched.GetMaxInstances() != 0 {
				as.StandardSchedulerSettings.MaxInstances = inSched.GetMaxInstances()
			}
			if inSched.GetMinInstances() != 0 {
				as.StandardSchedulerSettings.MinInstances = inSched.GetMinInstances()
			}
			if inSched.GetTargetCpuUtilization() != 0 {
				as.StandardSchedulerSettings.TargetCpuUtilization = inSched.GetTargetCpuUtilization()
			}
			if inSched.GetTargetThroughputUtilization() != 0 {
				as.StandardSchedulerSettings.TargetThroughputUtilization = inSched.GetTargetThroughputUtilization()
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	now := time.Now()
	opMetadata := &pb.OperationMetadataV1{
		Method:     "google.appengine.v1.Versions.UpdateVersion",
		InsertTime: timestamppb.New(now),
		User:       fmt.Sprintf("overseer-kcc-tester@%s.iam.gserviceaccount.com", name.App),
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("apps/%s", name.App)
	return s.operations.StartLROWithDone(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(now)
		return proto.Clone(updated), nil
	})
}

func (s *versionsServer) DeleteVersion(ctx context.Context, req *pb.DeleteVersionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVersionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Version{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	opMetadata := &pb.OperationMetadataV1{
		Method:     "google.appengine.v1.Versions.DeleteVersion",
		InsertTime: timestamppb.New(now),
		User:       fmt.Sprintf("overseer-kcc-tester@%s.iam.gserviceaccount.com", name.App),
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("apps/%s", name.App)
	return s.operations.StartLROWithDone(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(now)
		return &emptypb.Empty{}, nil
	})
}

func (s *versionsServer) ListVersions(ctx context.Context, req *pb.ListVersionsRequest) (*pb.ListVersionsResponse, error) {
	parent, err := s.parseVersionParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListVersionsResponse{}
	kind := (&pb.Version{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		version := obj.(*pb.Version)
		if strings.HasPrefix(version.GetName(), fmt.Sprintf("apps/%s/services/%s/versions/", parent.App, parent.Service)) {
			v := proto.Clone(version).(*pb.Version)
			v.Entrypoint = nil
			if req.GetView() != pb.VersionView_FULL {
				v.Handlers = nil
				v.ErrorHandlers = nil
				v.Libraries = nil
				v.ApiConfig = nil
				v.EnvVariables = nil
				v.BuildEnvVariables = nil
				v.DefaultExpiration = nil
				v.HealthCheck = nil
				v.ReadinessCheck = nil
				v.LivenessCheck = nil
				v.NobuildFilesRegex = ""
				v.Deployment = nil
			}
			response.Versions = append(response.Versions, v)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *MockService) populateDefaultsForVersion(parent *versionParent, name *versionName, obj *pb.Version, now time.Time) {
	obj.Name = name.String()
	obj.Id = name.Version

	if obj.Env == "" {
		obj.Env = "standard"
	}
	if obj.Env == "flex" {
		obj.Env = "flexible"
	}
	if obj.ServingStatus == pb.ServingStatus_SERVING_STATUS_UNSPECIFIED {
		obj.ServingStatus = pb.ServingStatus_SERVING
	}
	if obj.Env == "standard" {
		if obj.InstanceClass == "" {
			obj.InstanceClass = "F1"
		}
		if obj.RuntimeChannel == "" {
			obj.RuntimeChannel = "default"
		}
		if obj.DiskUsageBytes == 0 {
			obj.DiskUsageBytes = 2978869
		}
	}
	obj.Threadsafe = true

	if obj.ServiceAccount == "" {
		obj.ServiceAccount = fmt.Sprintf("%s@appspot.gserviceaccount.com", parent.App)
	}
	if obj.CreatedBy == "" {
		obj.CreatedBy = fmt.Sprintf("overseer-kcc-tester@%s.iam.gserviceaccount.com", parent.App)
	}
	if obj.CreateTime == nil {
		obj.CreateTime = timestamppb.New(now)
	}
	if obj.VersionUrl == "" {
		obj.VersionUrl = fmt.Sprintf("https://%s-dot-%s.uc.r.appspot.com", name.Version, parent.App)
	}
	if obj.Network == nil {
		obj.Network = &pb.Network{}
	}
	if obj.Deployment != nil {
		if obj.Deployment.Files != nil {
			for _, f := range obj.Deployment.Files {
				if f.Sha1Sum == "" && f.SourceUrl != "" {
					parts := strings.Split(f.SourceUrl, "/")
					hash := parts[len(parts)-1]
					if len(hash) == 40 {
						f.Sha1Sum = hash
					}
				}
			}
		}
		if obj.Deployment.Container != nil {
			if obj.Deployment.Container.Image == "gcr.io/google-samples/hello-app:1.0" {
				obj.Deployment.Container.Image = "gcr.io/google-samples/hello-app@sha256:649188051906c9df7f87fb2ef04a21acbdb009c0877fa67b3132b83230411f6e"
			}
		}
	}
	if obj.Env == "flexible" {
		if obj.BetaSettings == nil {
			obj.BetaSettings = make(map[string]string)
		}
		if obj.BetaSettings["no_appserver_affinity"] == "" {
			obj.BetaSettings["no_appserver_affinity"] = "true"
		}
		if obj.BetaSettings["use_deployment_manager"] == "" {
			obj.BetaSettings["use_deployment_manager"] = "true"
		}
		if obj.BetaSettings["minor_version_id_ccfe"] == "" {
			obj.BetaSettings["minor_version_id_ccfe"] = "1234567890123456789"
		}
	}
	if obj.Env == "standard" {
		if len(obj.Handlers) == 0 {
			obj.Handlers = []*pb.UrlMap{
				{
					Login:         pb.LoginRequirement_LOGIN_OPTIONAL,
					SecurityLevel: pb.SecurityLevel_SECURE_OPTIONAL,
					UrlRegex:      ".*",
					HandlerType: &pb.UrlMap_Script{
						Script: &pb.ScriptHandler{
							ScriptPath: "auto",
						},
					},
				},
			}
		} else {
			for _, h := range obj.Handlers {
				if h.AuthFailAction == pb.AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED {
					h.AuthFailAction = pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT
				}
			}
		}
	} else {
		for _, h := range obj.Handlers {
			if h.AuthFailAction == pb.AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED {
				h.AuthFailAction = pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT
			}
		}
	}
}
