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
	"crypto/sha256"
	"fmt"
	"slices"
	"sort"
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

func versionForReadView(v *pb.Version, view pb.VersionView) *pb.Version {
	result := proto.Clone(v).(*pb.Version)
	result.Entrypoint = nil
	if result.Deployment != nil {
		result.Deployment.Zip = nil
	}
	if view != pb.VersionView_FULL {
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
	return result
}

func versionForOperationResponse(v *pb.Version) *pb.Version {
	result := versionForReadView(v, pb.VersionView_BASIC)
	result.CreatedBy = ""
	return result
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

	return versionForReadView(obj, req.GetView()), nil
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
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(now)
		return versionForOperationResponse(obj), nil
	})
}

// toProtoFieldPath converts a camelCase or snake_case field path ("envVariables")
// to proto snake_case ("env_variables").
func toProtoFieldPath(p string) string {
	var b strings.Builder
	for _, r := range p {
		if r >= 'A' && r <= 'Z' {
			b.WriteByte('_')
			b.WriteRune(r - 'A' + 'a')
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// updateVersionAllowedPaths are the update_mask paths App Engine's UpdateVersion accepts
// (Standard and Flexible, per the public App Engine Admin API reference), sorted as in
// App Engine's error message.
var updateVersionAllowedPaths = []string{
	"automatic_scaling.cool_down_period_sec",
	"automatic_scaling.cpu_utilization.target_utilization",
	"automatic_scaling.max_idle_instances",
	"automatic_scaling.max_total_instances",
	"automatic_scaling.min_idle_instances",
	"automatic_scaling.min_total_instances",
	"automatic_scaling.standard_scheduler_settings.max_instances",
	"automatic_scaling.standard_scheduler_settings.min_instances",
	"automatic_scaling.standard_scheduler_settings.target_cpu_utilization",
	"automatic_scaling.standard_scheduler_settings.target_throughput_utilization",
	"instance_class",
	"manual_scaling.instances",
	"serving_status",
}

// validateUpdateVersionMask mirrors App Engine's update_mask validation for UpdateVersion.
func validateUpdateVersionMask(rawPaths []string) ([]string, error) {
	if len(rawPaths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "At least one field must be specified for this operation.")
	}
	allowed := make(map[string]bool, len(updateVersionAllowedPaths))
	for _, p := range updateVersionAllowedPaths {
		allowed[p] = true
	}
	snakePaths := make([]string, len(rawPaths))
	hasDisallowed := false
	for i, p := range rawPaths {
		sp := toProtoFieldPath(p)
		snakePaths[i] = sp
		if !allowed[sp] {
			hasDisallowed = true
		}
	}
	if hasDisallowed {
		got := slices.Clone(snakePaths)
		sort.Strings(got)
		return nil, status.Errorf(codes.InvalidArgument,
			"This operation is only supported on the following field(s): [%s], but got field(s): [%s] from the update request.",
			strings.Join(updateVersionAllowedPaths, ", "), strings.Join(got, ", "))
	}
	return snakePaths, nil
}

// validateUpdateVersionRequest mirrors App Engine's per-field UpdateVersion checks for
// Standard and Flexible versions.
func validateUpdateVersionRequest(existing, patch *pb.Version, paths []string) error {
	isFlex := existing.GetEnv() == "flexible" || existing.GetEnv() == "flex" || existing.GetVm()
	isManual := existing.GetManualScaling() != nil
	isBasic := !isManual && existing.GetBasicScaling() != nil
	isAutomatic := !isManual && !isBasic

	for _, p := range paths {
		switch p {
		case "automatic_scaling.cool_down_period_sec",
			"automatic_scaling.cpu_utilization.target_utilization",
			"automatic_scaling.max_total_instances",
			"automatic_scaling.min_total_instances":
			if !isFlex {
				return status.Errorf(codes.FailedPrecondition, "This request contained fields not supported on this version")
			}
		case "automatic_scaling.min_idle_instances",
			"automatic_scaling.max_idle_instances":
			if !isAutomatic {
				return status.Errorf(codes.FailedPrecondition, "You cannot change the scaling type of a version after deployment.")
			}
		case "automatic_scaling.standard_scheduler_settings.max_instances",
			"automatic_scaling.standard_scheduler_settings.min_instances",
			"automatic_scaling.standard_scheduler_settings.target_cpu_utilization",
			"automatic_scaling.standard_scheduler_settings.target_throughput_utilization":
			if isFlex {
				return status.Errorf(codes.FailedPrecondition, "This request contained fields not supported on this version")
			}
		case "manual_scaling.instances":
			if !isManual {
				return status.Errorf(codes.FailedPrecondition, "You cannot change the scaling type of a version after deployment.")
			}
		case "instance_class":
			if cls := patch.GetInstanceClass(); cls != "" {
				if isAutomatic && !strings.HasPrefix(cls, "F") {
					return status.Errorf(codes.InvalidArgument, "Instance class for automatic scaling versions must be set to one of the F_ instance classes")
				}
				if !isAutomatic && !strings.HasPrefix(cls, "B") {
					return status.Errorf(codes.InvalidArgument, "Instance class for basic or manual scaled versions must be set to one of the B_ instance classes")
				}
			}
		case "serving_status":
			if isAutomatic && !isFlex {
				return status.Errorf(codes.InvalidArgument, "serving status cannot be changed for Automatic Scaling versions")
			}
		}
	}
	return nil
}

func ensureAutomaticScaling(v *pb.Version) *pb.AutomaticScaling {
	if v.GetAutomaticScaling() == nil {
		v.Scaling = &pb.Version_AutomaticScaling{AutomaticScaling: &pb.AutomaticScaling{}}
	}
	return v.GetAutomaticScaling()
}

func ensureStandardSchedulerSettings(v *pb.Version) *pb.StandardSchedulerSettings {
	as := ensureAutomaticScaling(v)
	if as.StandardSchedulerSettings == nil {
		as.StandardSchedulerSettings = &pb.StandardSchedulerSettings{}
	}
	return as.StandardSchedulerSettings
}

func applyUpdateVersionMask(updated, patch *pb.Version, paths []string) error {
	for _, p := range paths {
		switch p {
		case "instance_class":
			updated.InstanceClass = patch.GetInstanceClass()
		case "serving_status":
			updated.ServingStatus = patch.GetServingStatus()
		case "manual_scaling.instances":
			if updated.GetManualScaling() == nil {
				updated.Scaling = &pb.Version_ManualScaling{ManualScaling: &pb.ManualScaling{}}
			}
			updated.GetManualScaling().Instances = patch.GetManualScaling().GetInstances()
		case "automatic_scaling.min_idle_instances":
			ensureAutomaticScaling(updated).MinIdleInstances = patch.GetAutomaticScaling().GetMinIdleInstances()
		case "automatic_scaling.max_idle_instances":
			ensureAutomaticScaling(updated).MaxIdleInstances = patch.GetAutomaticScaling().GetMaxIdleInstances()
		case "automatic_scaling.standard_scheduler_settings.max_instances":
			ensureStandardSchedulerSettings(updated).MaxInstances = patch.GetAutomaticScaling().GetStandardSchedulerSettings().GetMaxInstances()
		case "automatic_scaling.standard_scheduler_settings.min_instances":
			ensureStandardSchedulerSettings(updated).MinInstances = patch.GetAutomaticScaling().GetStandardSchedulerSettings().GetMinInstances()
		case "automatic_scaling.standard_scheduler_settings.target_cpu_utilization":
			ensureStandardSchedulerSettings(updated).TargetCpuUtilization = patch.GetAutomaticScaling().GetStandardSchedulerSettings().GetTargetCpuUtilization()
		case "automatic_scaling.standard_scheduler_settings.target_throughput_utilization":
			ensureStandardSchedulerSettings(updated).TargetThroughputUtilization = patch.GetAutomaticScaling().GetStandardSchedulerSettings().GetTargetThroughputUtilization()
		default:
			return status.Errorf(codes.Unimplemented, "update_mask path %q is not implemented in mockappengine", p)
		}
	}
	return nil
}

func (s *versionsServer) UpdateVersion(ctx context.Context, req *pb.UpdateVersionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVersionName(req.GetName())
	if err != nil {
		return nil, err
	}

	paths, err := validateUpdateVersionMask(req.GetUpdateMask().GetPaths())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.Version{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	if err := validateUpdateVersionRequest(existing, req.GetVersion(), paths); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Version)
	if err := applyUpdateVersionMask(updated, req.GetVersion(), paths); err != nil {
		return nil, err
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
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.New(now)
		return versionForOperationResponse(updated), nil
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
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
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
			response.Versions = append(response.Versions, versionForReadView(version, req.GetView()))
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
		if obj.Deployment.Container != nil && obj.Deployment.Container.Image != "" {
			image := obj.Deployment.Container.Image
			if !strings.Contains(image, "@sha256:") {
				// App Engine Flexible resolves container image tags to a sha256 digest.
				base := image
				if idx := strings.LastIndex(image, ":"); idx != -1 {
					base = image[:idx]
				}
				h := sha256.Sum256([]byte(image))
				obj.Deployment.Container.Image = fmt.Sprintf("%s@sha256:%x", base, h)
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
		// Real App Engine appends the default URL map only when the runtime defines one
		// (2nd-gen runtimes). The mock appends for all Standard runtimes; fixtures use python311.
		defaultHandler := &pb.UrlMap{
			Login:         pb.LoginRequirement_LOGIN_OPTIONAL,
			SecurityLevel: pb.SecurityLevel_SECURE_OPTIONAL,
			UrlRegex:      ".*",
			HandlerType: &pb.UrlMap_Script{
				Script: &pb.ScriptHandler{
					ScriptPath: "auto",
				},
			},
		}
		if len(obj.Handlers) == 0 {
			obj.Handlers = []*pb.UrlMap{defaultHandler}
		} else {
			for _, h := range obj.Handlers {
				if h.Login == pb.LoginRequirement_LOGIN_UNSPECIFIED {
					h.Login = pb.LoginRequirement_LOGIN_OPTIONAL
				}
				if h.SecurityLevel == pb.SecurityLevel_SECURE_UNSPECIFIED {
					h.SecurityLevel = pb.SecurityLevel_SECURE_OPTIONAL
				}
				if h.AuthFailAction == pb.AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED {
					h.AuthFailAction = pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT
				}
			}
			obj.Handlers = append(obj.Handlers, defaultHandler)
		}
	} else {
		for _, h := range obj.Handlers {
			if h.AuthFailAction == pb.AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED {
				h.AuthFailAction = pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT
			}
		}
	}
}
