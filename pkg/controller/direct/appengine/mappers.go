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

package appengine

import (
	"sort"
	"strconv"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	krmappenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// AppEngineFirewallRuleSpec_v1alpha1_FromProto converts the protobuf FirewallRule to the KRM AppEngineFirewallRuleSpec.
// This is handcoded because Action is represented as an enum in protobuf but as a required non-pointer string in the existing CRD schema,
// SourceRange is a required non-pointer string in KRM, and Priority is an int32 in protobuf while represented by ResourceID (*string) in KRM.
func AppEngineFirewallRuleSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.FirewallRule) *krmappenginev1alpha1.AppEngineFirewallRuleSpec {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.AppEngineFirewallRuleSpec{}
	out.Action = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetAction()))
	out.SourceRange = in.GetSourceRange()
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetPriority() != 0 {
		priorityStr := strconv.FormatInt(int64(in.GetPriority()), 10)
		out.ResourceID = &priorityStr
	}
	return out
}

// AppEngineFirewallRuleSpec_v1alpha1_ToProto converts the KRM AppEngineFirewallRuleSpec to the protobuf FirewallRule.
// This is handcoded because Action is represented as an enum in protobuf but as a required non-pointer string in the existing CRD schema,
// SourceRange is a required non-pointer string in KRM, and Priority is an int32 in protobuf while represented by ResourceID (*string) in KRM.
func AppEngineFirewallRuleSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.AppEngineFirewallRuleSpec) *pb.FirewallRule {
	if in == nil {
		return nil
	}
	out := &pb.FirewallRule{}
	out.Action = direct.Enum_ToProto[pb.FirewallRule_Action](mapCtx, &in.Action)
	out.SourceRange = in.SourceRange
	out.Description = direct.ValueOf(in.Description)
	if in.ResourceID != nil {
		priorityVal, err := strconv.ParseInt(*in.ResourceID, 10, 32)
		if err != nil {
			mapCtx.Errorf("cannot parse resourceID %q as int32 priority: %v", *in.ResourceID, err)
		} else {
			out.Priority = int32(priorityVal)
		}
	}
	return out
}

// AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto converts protobuf Service to KRM AppEngineServiceSplitTrafficSpec.
// This is handcoded because Service.id maps to KRM ResourceID, and migrateTraffic / project are KRM-specific fields.
func AppEngineServiceSplitTrafficSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec{}
	out.Split = AppEngineServiceSplitTrafficSplit_v1alpha1_FromProto(mapCtx, in.GetSplit())
	if in.GetId() != "" {
		out.ResourceID = direct.LazyPtr(in.GetId())
	}
	return out
}

// AppEngineServiceSplitTrafficSpec_v1alpha1_ToProto converts KRM AppEngineServiceSplitTrafficSpec to protobuf Service.
// This is handcoded because ResourceID maps to Service.id, and migrateTraffic / project are KRM-specific fields.
func AppEngineServiceSplitTrafficSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.AppEngineServiceSplitTrafficSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	out.Split = AppEngineServiceSplitTrafficSplit_v1alpha1_ToProto(mapCtx, in.Split)
	if in.ResourceID != nil {
		out.Id = *in.ResourceID
	}
	return out
}

// Allocations_FromProto converts protobuf float64 traffic allocations to KRM string representations.
// This is handcoded because in protobuf allocations is map[string]double whereas in the existing CRD it is map[string]string.
func Allocations_FromProto(mapCtx *direct.MapContext, in map[string]float64) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return out
}

// Allocations_ToProto converts KRM string traffic allocations to protobuf float64 representations.
// This is handcoded because in protobuf allocations is map[string]double whereas in the existing CRD it is map[string]string.
func Allocations_ToProto(mapCtx *direct.MapContext, in map[string]string) map[string]float64 {
	if in == nil {
		return nil
	}
	out := make(map[string]float64, len(in))
	for k, v := range in {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			mapCtx.Errorf("cannot parse allocation %q for version %q as float64: %v", v, k, err)
			continue
		}
		out[k] = val
	}
	return out
}

// AutomaticScaling_v1alpha1_FromProto converts protobuf AutomaticScaling to KRM AutomaticScaling.
// This is handcoded because integer fields are int in KRM (to match openapi integer schema without format int32)
// but int32 in protobuf.
func AutomaticScaling_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AutomaticScaling) *krmappenginev1alpha1.AutomaticScaling {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.AutomaticScaling{}
	if in.GetMaxConcurrentRequests() != 0 {
		v := int(in.GetMaxConcurrentRequests())
		out.MaxConcurrentRequests = &v
	}
	if in.GetMaxIdleInstances() != 0 {
		v := int(in.GetMaxIdleInstances())
		out.MaxIdleInstances = &v
	}
	if in.GetMinIdleInstances() != 0 {
		v := int(in.GetMinIdleInstances())
		out.MinIdleInstances = &v
	}
	out.MaxPendingLatency = direct.StringDuration_FromProto(mapCtx, in.GetMaxPendingLatency())
	out.MinPendingLatency = direct.StringDuration_FromProto(mapCtx, in.GetMinPendingLatency())
	out.StandardSchedulerSettings = StandardSchedulerSettings_v1alpha1_FromProto(mapCtx, in.GetStandardSchedulerSettings())
	return out
}

// AutomaticScaling_v1alpha1_ToProto converts KRM AutomaticScaling to protobuf AutomaticScaling.
// This is handcoded because integer fields are int in KRM but int32 in protobuf.
func AutomaticScaling_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.AutomaticScaling) *pb.AutomaticScaling {
	if in == nil {
		return nil
	}
	out := &pb.AutomaticScaling{}
	if in.MaxConcurrentRequests != nil {
		out.MaxConcurrentRequests = int32(*in.MaxConcurrentRequests)
	}
	if in.MaxIdleInstances != nil {
		out.MaxIdleInstances = int32(*in.MaxIdleInstances)
	}
	if in.MinIdleInstances != nil {
		out.MinIdleInstances = int32(*in.MinIdleInstances)
	}
	out.MaxPendingLatency = direct.StringDuration_ToProto(mapCtx, in.MaxPendingLatency)
	out.MinPendingLatency = direct.StringDuration_ToProto(mapCtx, in.MinPendingLatency)
	out.StandardSchedulerSettings = StandardSchedulerSettings_v1alpha1_ToProto(mapCtx, in.StandardSchedulerSettings)
	return out
}

// StandardSchedulerSettings_v1alpha1_FromProto converts protobuf StandardSchedulerSettings to KRM StandardSchedulerSettings.
// This is handcoded because MaxInstances and MinInstances are int in KRM (to match openapi integer schema without format int32)
// but int32 in protobuf.
func StandardSchedulerSettings_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.StandardSchedulerSettings) *krmappenginev1alpha1.StandardSchedulerSettings {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.StandardSchedulerSettings{}
	if in.GetMaxInstances() != 0 {
		v := int(in.GetMaxInstances())
		out.MaxInstances = &v
	}
	if in.GetMinInstances() != 0 {
		v := int(in.GetMinInstances())
		out.MinInstances = &v
	}
	out.TargetCPUUtilization = direct.LazyPtr(in.GetTargetCpuUtilization())
	out.TargetThroughputUtilization = direct.LazyPtr(in.GetTargetThroughputUtilization())
	return out
}

// StandardSchedulerSettings_v1alpha1_ToProto converts KRM StandardSchedulerSettings to protobuf StandardSchedulerSettings.
// This is handcoded because MaxInstances and MinInstances are int in KRM but int32 in protobuf.
func StandardSchedulerSettings_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.StandardSchedulerSettings) *pb.StandardSchedulerSettings {
	if in == nil {
		return nil
	}
	out := &pb.StandardSchedulerSettings{}
	if in.MaxInstances != nil {
		out.MaxInstances = int32(*in.MaxInstances)
	}
	if in.MinInstances != nil {
		out.MinInstances = int32(*in.MinInstances)
	}
	out.TargetCpuUtilization = direct.ValueOf(in.TargetCPUUtilization)
	out.TargetThroughputUtilization = direct.ValueOf(in.TargetThroughputUtilization)
	return out
}

// Deployment_v1alpha1_FromProto converts protobuf Deployment to KRM Deployment.
// This is handcoded because in protobuf files are modeled as a map[string]*FileInfo,
// whereas in the existing KRM CRD files are modeled as a slice of DeploymentFile structs,
// and FilesCount is int in KRM but int32 in protobuf.
func Deployment_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krmappenginev1alpha1.Deployment {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.Deployment{}
	if in.Files != nil {
		names := make([]string, 0, len(in.Files))
		for name := range in.Files {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fileInfo := in.Files[name]
			if fileInfo == nil {
				continue
			}
			f := krmappenginev1alpha1.DeploymentFile{
				Name:      name,
				SourceUrl: fileInfo.GetSourceUrl(),
			}
			if fileInfo.GetSha1Sum() != "" {
				f.Sha1Sum = direct.LazyPtr(fileInfo.GetSha1Sum())
			}
			out.Files = append(out.Files, f)
		}
	}
	if in.Zip != nil {
		out.Zip = &krmappenginev1alpha1.DeploymentZip{
			SourceUrl: in.Zip.GetSourceUrl(),
		}
		if in.Zip.GetFilesCount() != 0 {
			count := int(in.Zip.GetFilesCount())
			out.Zip.FilesCount = &count
		}
	}
	return out
}

// Deployment_v1alpha1_ToProto converts KRM Deployment to protobuf Deployment.
// This is handcoded because in protobuf files are modeled as a map[string]*FileInfo,
// whereas in the existing KRM CRD files are modeled as a slice of DeploymentFile structs.
func Deployment_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.Deployment) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	if len(in.Files) > 0 {
		out.Files = make(map[string]*pb.FileInfo)
		for _, f := range in.Files {
			fileInfo := &pb.FileInfo{
				SourceUrl: f.SourceUrl,
			}
			if f.Sha1Sum != nil {
				fileInfo.Sha1Sum = *f.Sha1Sum
			}
			out.Files[f.Name] = fileInfo
		}
	}
	if in.Zip != nil {
		out.Zip = &pb.ZipInfo{
			SourceUrl: in.Zip.SourceUrl,
		}
		if in.Zip.FilesCount != nil {
			out.Zip.FilesCount = int32(*in.Zip.FilesCount)
		}
	}
	return out
}

// AppEngineStandardAppVersionSpec_Runtime_ToProto maps a required non-pointer string to string.
func AppEngineStandardAppVersionSpec_Runtime_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

// BasicScaling_v1alpha1_FromProto converts protobuf BasicScaling to KRM BasicScaling.
// This is handcoded because MaxInstances is a required scalar int in KRM (non-pointer) but int32 in protobuf.
func BasicScaling_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.BasicScaling) *krmappenginev1alpha1.BasicScaling {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.BasicScaling{}
	out.IdleTimeout = direct.StringDuration_FromProto(mapCtx, in.GetIdleTimeout())
	out.MaxInstances = int(in.GetMaxInstances())
	return out
}

// BasicScaling_v1alpha1_ToProto converts KRM BasicScaling to protobuf BasicScaling.
// This is handcoded because MaxInstances is a required scalar int in KRM (non-pointer).
func BasicScaling_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.BasicScaling) *pb.BasicScaling {
	if in == nil {
		return nil
	}
	out := &pb.BasicScaling{}
	out.IdleTimeout = direct.StringDuration_ToProto(mapCtx, in.IdleTimeout)
	out.MaxInstances = int32(in.MaxInstances)
	return out
}

// ManualScaling_v1alpha1_FromProto converts protobuf ManualScaling to KRM ManualScaling.
// This is handcoded because Instances is a required scalar int in KRM (non-pointer) but int32 in protobuf.
func ManualScaling_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ManualScaling) *krmappenginev1alpha1.ManualScaling {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.ManualScaling{}
	out.Instances = int(in.GetInstances())
	return out
}

// ManualScaling_v1alpha1_ToProto converts KRM ManualScaling to protobuf ManualScaling.
// This is handcoded because Instances is a required scalar int in KRM (non-pointer).
func ManualScaling_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.ManualScaling) *pb.ManualScaling {
	if in == nil {
		return nil
	}
	out := &pb.ManualScaling{}
	out.Instances = int32(in.Instances)
	return out
}

// Entrypoint_v1alpha1_ToProto converts KRM Entrypoint to protobuf Entrypoint.
// This is handcoded because Shell is a required scalar string in KRM (non-pointer) mapped to a oneof in protobuf.
func Entrypoint_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.Entrypoint) *pb.Entrypoint {
	if in == nil {
		return nil
	}
	out := &pb.Entrypoint{}
	if in.Shell != "" {
		out.Command = &pb.Entrypoint_Shell{Shell: in.Shell}
	}
	return out
}

// Entrypoint_Shell_ToProto maps a required non-pointer string to *pb.Entrypoint_Shell.
func Entrypoint_Shell_ToProto(mapCtx *direct.MapContext, in string) *pb.Entrypoint_Shell {
	return &pb.Entrypoint_Shell{Shell: in}
}

// ScriptHandler_ScriptPath_ToProto maps a required non-pointer string to string.
func ScriptHandler_ScriptPath_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}

// VPCAccessConnector_Name_ToProto maps a required non-pointer string to string.
func VPCAccessConnector_Name_ToProto(mapCtx *direct.MapContext, in string) string {
	return in
}
