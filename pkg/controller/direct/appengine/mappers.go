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

// AppEngineFlexibleAppVersionSpec_v1alpha1_FromProto converts the protobuf Version to the KRM AppEngineFlexibleAppVersionSpec.
// This is handcoded because the baseline CRD schema uses customized field structures, integer/float differences,
// and string-encoded durations or request counts.
func AppEngineFlexibleAppVersionSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krmappenginev1alpha1.AppEngineFlexibleAppVersionSpec {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.AppEngineFlexibleAppVersionSpec{}
	out.ResourceID = direct.LazyPtr(in.GetId())
	out.Runtime = in.GetRuntime()
	out.RuntimeChannel = direct.LazyPtr(in.GetRuntimeChannel())
	out.RuntimeApiVersion = direct.LazyPtr(in.GetRuntimeApiVersion())
	out.RuntimeMainExecutablePath = direct.LazyPtr(in.GetRuntimeMainExecutablePath())
	out.InstanceClass = direct.LazyPtr(in.GetInstanceClass())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ServingStatus = direct.Enum_FromProto(mapCtx, in.GetServingStatus())
	out.InboundServices = direct.EnumSlice_FromProto(mapCtx, in.InboundServices)
	out.BetaSettings = in.BetaSettings
	out.EnvVariables = in.EnvVariables
	out.DefaultExpiration = direct.StringDuration_FromProto(mapCtx, in.GetDefaultExpiration())
	out.NobuildFilesRegex = direct.LazyPtr(in.GetNobuildFilesRegex())

	if in.GetAutomaticScaling() != nil {
		as := in.GetAutomaticScaling()
		outAS := &krmappenginev1alpha1.AppEngineFlexibleAppVersionAutomaticScaling{
			CoolDownPeriod:        direct.StringDuration_FromProto(mapCtx, as.GetCoolDownPeriod()),
			MaxConcurrentRequests: direct.LazyPtr(int(as.GetMaxConcurrentRequests())),
			MaxIdleInstances:      direct.LazyPtr(int(as.GetMaxIdleInstances())),
			MaxTotalInstances:     direct.LazyPtr(int(as.GetMaxTotalInstances())),
			MaxPendingLatency:     direct.StringDuration_FromProto(mapCtx, as.GetMaxPendingLatency()),
			MinIdleInstances:      direct.LazyPtr(int(as.GetMinIdleInstances())),
			MinTotalInstances:     direct.LazyPtr(int(as.GetMinTotalInstances())),
			MinPendingLatency:     direct.StringDuration_FromProto(mapCtx, as.GetMinPendingLatency()),
		}
		if as.GetCpuUtilization() != nil {
			outAS.CpuUtilization = &krmappenginev1alpha1.AppEngineFlexibleAppVersionCpuUtilization{
				AggregationWindowLength: direct.StringDuration_FromProto(mapCtx, as.GetCpuUtilization().GetAggregationWindowLength()),
				TargetUtilization:       as.GetCpuUtilization().GetTargetUtilization(),
			}
		}
		if as.GetDiskUtilization() != nil {
			du := as.GetDiskUtilization()
			outAS.DiskUtilization = &krmappenginev1alpha1.AppEngineFlexibleAppVersionDiskUtilization{
				TargetReadBytesPerSecond:  direct.LazyPtr(int(du.GetTargetReadBytesPerSecond())),
				TargetReadOpsPerSecond:    direct.LazyPtr(int(du.GetTargetReadOpsPerSecond())),
				TargetWriteBytesPerSecond: direct.LazyPtr(int(du.GetTargetWriteBytesPerSecond())),
				TargetWriteOpsPerSecond:   direct.LazyPtr(int(du.GetTargetWriteOpsPerSecond())),
			}
		}
		if as.GetNetworkUtilization() != nil {
			nu := as.GetNetworkUtilization()
			outAS.NetworkUtilization = &krmappenginev1alpha1.AppEngineFlexibleAppVersionNetworkUtilization{
				TargetReceivedBytesPerSecond:   direct.LazyPtr(int(nu.GetTargetReceivedBytesPerSecond())),
				TargetReceivedPacketsPerSecond: direct.LazyPtr(int(nu.GetTargetReceivedPacketsPerSecond())),
				TargetSentBytesPerSecond:       direct.LazyPtr(int(nu.GetTargetSentBytesPerSecond())),
				TargetSentPacketsPerSecond:     direct.LazyPtr(int(nu.GetTargetSentPacketsPerSecond())),
			}
		}
		if as.GetRequestUtilization() != nil {
			ru := as.GetRequestUtilization()
			outAS.RequestUtilization = &krmappenginev1alpha1.AppEngineFlexibleAppVersionRequestUtilization{}
			if ru.GetTargetRequestCountPerSecond() != 0 {
				val := strconv.FormatInt(int64(ru.GetTargetRequestCountPerSecond()), 10)
				outAS.RequestUtilization.TargetRequestCountPerSecond = &val
			}
			if ru.GetTargetConcurrentRequests() != 0 {
				val := float64(ru.GetTargetConcurrentRequests())
				outAS.RequestUtilization.TargetConcurrentRequests = &val
			}
		}
		out.AutomaticScaling = outAS
	}

	if in.GetManualScaling() != nil {
		out.ManualScaling = &krmappenginev1alpha1.AppEngineFlexibleAppVersionManualScaling{
			Instances: int(in.GetManualScaling().GetInstances()),
		}
	}

	if in.GetNetwork() != nil {
		net := in.GetNetwork()
		out.Network = &krmappenginev1alpha1.AppEngineFlexibleAppVersionNetwork{
			ForwardedPorts:  net.GetForwardedPorts(),
			InstanceTag:     direct.LazyPtr(net.GetInstanceTag()),
			Name:            net.GetName(),
			SessionAffinity: direct.LazyPtr(net.GetSessionAffinity()),
			Subnetwork:      direct.LazyPtr(net.GetSubnetworkName()),
		}
	}

	if in.GetResources() != nil {
		res := in.GetResources()
		outRes := &krmappenginev1alpha1.AppEngineFlexibleAppVersionResources{
			MemoryGb: direct.LazyPtr(res.GetMemoryGb()),
		}
		if res.GetCpu() != 0 {
			cpu := int(res.GetCpu())
			outRes.Cpu = &cpu
		}
		if res.GetDiskGb() != 0 {
			diskGb := int(res.GetDiskGb())
			outRes.DiskGb = &diskGb
		}
		for _, v := range res.GetVolumes() {
			outRes.Volumes = append(outRes.Volumes, krmappenginev1alpha1.AppEngineFlexibleAppVersionVolumes{
				Name:       v.GetName(),
				VolumeType: v.GetVolumeType(),
				SizeGb:     int(v.GetSizeGb()),
			})
		}
		out.Resources = outRes
	}

	if in.GetReadinessCheck() != nil {
		rc := in.GetReadinessCheck()
		out.ReadinessCheck = &krmappenginev1alpha1.AppEngineFlexibleAppVersionReadinessCheck{
			Path:            rc.GetPath(),
			Host:            direct.LazyPtr(rc.GetHost()),
			CheckInterval:   direct.StringDuration_FromProto(mapCtx, rc.GetCheckInterval()),
			Timeout:         direct.StringDuration_FromProto(mapCtx, rc.GetTimeout()),
			AppStartTimeout: direct.StringDuration_FromProto(mapCtx, rc.GetAppStartTimeout()),
		}
		if rc.GetFailureThreshold() != 0 {
			val := float64(rc.GetFailureThreshold())
			out.ReadinessCheck.FailureThreshold = &val
		}
		if rc.GetSuccessThreshold() != 0 {
			val := float64(rc.GetSuccessThreshold())
			out.ReadinessCheck.SuccessThreshold = &val
		}
	}

	if in.GetLivenessCheck() != nil {
		lc := in.GetLivenessCheck()
		out.LivenessCheck = &krmappenginev1alpha1.AppEngineFlexibleAppVersionLivenessCheck{
			Path:          lc.GetPath(),
			Host:          direct.LazyPtr(lc.GetHost()),
			CheckInterval: direct.StringDuration_FromProto(mapCtx, lc.GetCheckInterval()),
			Timeout:       direct.StringDuration_FromProto(mapCtx, lc.GetTimeout()),
			InitialDelay:  direct.StringDuration_FromProto(mapCtx, lc.GetInitialDelay()),
		}
		if lc.GetFailureThreshold() != 0 {
			val := float64(lc.GetFailureThreshold())
			out.LivenessCheck.FailureThreshold = &val
		}
		if lc.GetSuccessThreshold() != 0 {
			val := float64(lc.GetSuccessThreshold())
			out.LivenessCheck.SuccessThreshold = &val
		}
	}

	if in.GetDeployment() != nil {
		dep := in.GetDeployment()
		outDep := &krmappenginev1alpha1.AppEngineFlexibleAppVersionDeployment{}
		if dep.GetCloudBuildOptions() != nil {
			cbo := dep.GetCloudBuildOptions()
			outDep.CloudBuildOptions = &krmappenginev1alpha1.AppEngineFlexibleAppVersionCloudBuildOptions{
				AppYamlPath:       cbo.GetAppYamlPath(),
				CloudBuildTimeout: direct.StringDuration_FromProto(mapCtx, cbo.GetCloudBuildTimeout()),
			}
		}
		if dep.GetContainer() != nil {
			outDep.Container = &krmappenginev1alpha1.AppEngineFlexibleAppVersionContainer{
				Image: dep.GetContainer().GetImage(),
			}
		}
		if dep.GetZip() != nil {
			z := dep.GetZip()
			outDep.Zip = &krmappenginev1alpha1.AppEngineFlexibleAppVersionZip{
				SourceUrl:  z.GetSourceUrl(),
				FilesCount: direct.LazyPtr(int(z.GetFilesCount())),
			}
		}
		if len(dep.GetFiles()) > 0 {
			names := make([]string, 0, len(dep.GetFiles()))
			for k := range dep.GetFiles() {
				names = append(names, k)
			}
			sort.Strings(names)
			for _, k := range names {
				f := dep.GetFiles()[k]
				outDep.Files = append(outDep.Files, krmappenginev1alpha1.AppEngineFlexibleAppVersionFiles{
					Name:      k,
					Sha1Sum:   direct.LazyPtr(f.GetSha1Sum()),
					SourceUrl: f.GetSourceUrl(),
				})
			}
		}
		out.Deployment = outDep
	}

	if in.GetEntrypoint() != nil {
		out.Entrypoint = &krmappenginev1alpha1.AppEngineFlexibleAppVersionEntrypoint{
			Shell: in.GetEntrypoint().GetShell(),
		}
	}

	if in.GetEndpointsApiService() != nil {
		eas := in.GetEndpointsApiService()
		out.EndpointsApiService = &krmappenginev1alpha1.AppEngineFlexibleAppVersionEndpointsApiService{
			Name:                 eas.GetName(),
			ConfigId:             direct.LazyPtr(eas.GetConfigId()),
			RolloutStrategy:      direct.Enum_FromProto(mapCtx, eas.GetRolloutStrategy()),
			DisableTraceSampling: direct.LazyPtr(eas.GetDisableTraceSampling()),
		}
	}

	if in.GetVpcAccessConnector() != nil {
		out.VpcAccessConnector = &krmappenginev1alpha1.AppEngineFlexibleAppVersionVpcAccessConnector{
			Name: in.GetVpcAccessConnector().GetName(),
		}
	}

	if in.GetApiConfig() != nil {
		ac := in.GetApiConfig()
		out.ApiConfig = &krmappenginev1alpha1.AppEngineFlexibleAppVersionApiConfig{
			AuthFailAction: direct.Enum_FromProto(mapCtx, ac.GetAuthFailAction()),
			Login:          direct.Enum_FromProto(mapCtx, ac.GetLogin()),
			Script:         ac.GetScript(),
			SecurityLevel:  direct.Enum_FromProto(mapCtx, ac.GetSecurityLevel()),
			Url:            direct.LazyPtr(ac.GetUrl()),
		}
	}

	for _, h := range in.GetHandlers() {
		outH := krmappenginev1alpha1.AppEngineFlexibleAppVersionHandlers{
			AuthFailAction:           direct.Enum_FromProto(mapCtx, h.GetAuthFailAction()),
			Login:                    direct.Enum_FromProto(mapCtx, h.GetLogin()),
			RedirectHttpResponseCode: direct.Enum_FromProto(mapCtx, h.GetRedirectHttpResponseCode()),
			SecurityLevel:            direct.Enum_FromProto(mapCtx, h.GetSecurityLevel()),
			UrlRegex:                 direct.LazyPtr(h.GetUrlRegex()),
		}
		if h.GetScript() != nil {
			outH.Script = &krmappenginev1alpha1.AppEngineFlexibleAppVersionScript{
				ScriptPath: h.GetScript().GetScriptPath(),
			}
		}
		if h.GetStaticFiles() != nil {
			sf := h.GetStaticFiles()
			outH.StaticFiles = &krmappenginev1alpha1.AppEngineFlexibleAppVersionStaticFiles{
				ApplicationReadable: direct.LazyPtr(sf.GetApplicationReadable()),
				Expiration:          direct.StringDuration_FromProto(mapCtx, sf.GetExpiration()),
				HttpHeaders:         sf.GetHttpHeaders(),
				MimeType:            direct.LazyPtr(sf.GetMimeType()),
				Path:                direct.LazyPtr(sf.GetPath()),
				RequireMatchingFile: direct.LazyPtr(sf.GetRequireMatchingFile()),
				UploadPathRegex:     direct.LazyPtr(sf.GetUploadPathRegex()),
			}
		}
		out.Handlers = append(out.Handlers, outH)
	}

	return out
}

// AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto converts the KRM AppEngineFlexibleAppVersionSpec to the protobuf Version.
// This is handcoded because the baseline CRD schema uses customized field structures, integer/float differences,
// and string-encoded durations or request counts.
func AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.AppEngineFlexibleAppVersionSpec) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	out.Id = direct.ValueOf(in.ResourceID)
	out.Runtime = in.Runtime
	out.RuntimeChannel = direct.ValueOf(in.RuntimeChannel)
	out.RuntimeApiVersion = direct.ValueOf(in.RuntimeApiVersion)
	out.RuntimeMainExecutablePath = direct.ValueOf(in.RuntimeMainExecutablePath)
	out.InstanceClass = direct.ValueOf(in.InstanceClass)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ServingStatus = direct.Enum_ToProto[pb.ServingStatus](mapCtx, in.ServingStatus)
	out.InboundServices = direct.EnumSlice_ToProto[pb.InboundServiceType](mapCtx, in.InboundServices)
	out.BetaSettings = in.BetaSettings
	out.EnvVariables = in.EnvVariables
	out.DefaultExpiration = direct.StringDuration_ToProto(mapCtx, in.DefaultExpiration)
	out.NobuildFilesRegex = direct.ValueOf(in.NobuildFilesRegex)

	if in.AutomaticScaling != nil {
		as := in.AutomaticScaling
		outAS := &pb.AutomaticScaling{
			CoolDownPeriod:        direct.StringDuration_ToProto(mapCtx, as.CoolDownPeriod),
			MaxConcurrentRequests: int32(direct.ValueOf(as.MaxConcurrentRequests)),
			MaxIdleInstances:      int32(direct.ValueOf(as.MaxIdleInstances)),
			MaxTotalInstances:     int32(direct.ValueOf(as.MaxTotalInstances)),
			MaxPendingLatency:     direct.StringDuration_ToProto(mapCtx, as.MaxPendingLatency),
			MinIdleInstances:      int32(direct.ValueOf(as.MinIdleInstances)),
			MinTotalInstances:     int32(direct.ValueOf(as.MinTotalInstances)),
			MinPendingLatency:     direct.StringDuration_ToProto(mapCtx, as.MinPendingLatency),
		}
		if as.CpuUtilization != nil {
			outAS.CpuUtilization = &pb.CpuUtilization{
				AggregationWindowLength: direct.StringDuration_ToProto(mapCtx, as.CpuUtilization.AggregationWindowLength),
				TargetUtilization:       as.CpuUtilization.TargetUtilization,
			}
		}
		if as.DiskUtilization != nil {
			outAS.DiskUtilization = &pb.DiskUtilization{
				TargetReadBytesPerSecond:  int32(direct.ValueOf(as.DiskUtilization.TargetReadBytesPerSecond)),
				TargetReadOpsPerSecond:    int32(direct.ValueOf(as.DiskUtilization.TargetReadOpsPerSecond)),
				TargetWriteBytesPerSecond: int32(direct.ValueOf(as.DiskUtilization.TargetWriteBytesPerSecond)),
				TargetWriteOpsPerSecond:   int32(direct.ValueOf(as.DiskUtilization.TargetWriteOpsPerSecond)),
			}
		}
		if as.NetworkUtilization != nil {
			outAS.NetworkUtilization = &pb.NetworkUtilization{
				TargetReceivedBytesPerSecond:   int32(direct.ValueOf(as.NetworkUtilization.TargetReceivedBytesPerSecond)),
				TargetReceivedPacketsPerSecond: int32(direct.ValueOf(as.NetworkUtilization.TargetReceivedPacketsPerSecond)),
				TargetSentBytesPerSecond:       int32(direct.ValueOf(as.NetworkUtilization.TargetSentBytesPerSecond)),
				TargetSentPacketsPerSecond:     int32(direct.ValueOf(as.NetworkUtilization.TargetSentPacketsPerSecond)),
			}
		}
		if as.RequestUtilization != nil {
			outAS.RequestUtilization = &pb.RequestUtilization{}
			if as.RequestUtilization.TargetRequestCountPerSecond != nil {
				val, err := strconv.ParseInt(*as.RequestUtilization.TargetRequestCountPerSecond, 10, 32)
				if err != nil {
					mapCtx.Errorf("cannot parse targetRequestCountPerSecond %q: %v", *as.RequestUtilization.TargetRequestCountPerSecond, err)
				} else {
					outAS.RequestUtilization.TargetRequestCountPerSecond = int32(val)
				}
			}
			if as.RequestUtilization.TargetConcurrentRequests != nil {
				outAS.RequestUtilization.TargetConcurrentRequests = int32(*as.RequestUtilization.TargetConcurrentRequests)
			}
		}
		out.Scaling = &pb.Version_AutomaticScaling{AutomaticScaling: outAS}
	}

	if in.ManualScaling != nil {
		out.Scaling = &pb.Version_ManualScaling{
			ManualScaling: &pb.ManualScaling{
				Instances: int32(in.ManualScaling.Instances),
			},
		}
	}

	if in.Network != nil {
		out.Network = &pb.Network{
			ForwardedPorts:  in.Network.ForwardedPorts,
			InstanceTag:     direct.ValueOf(in.Network.InstanceTag),
			Name:            in.Network.Name,
			SessionAffinity: direct.ValueOf(in.Network.SessionAffinity),
			SubnetworkName:  direct.ValueOf(in.Network.Subnetwork),
		}
	}

	if in.Resources != nil {
		outRes := &pb.Resources{
			MemoryGb: direct.ValueOf(in.Resources.MemoryGb),
		}
		if in.Resources.Cpu != nil {
			outRes.Cpu = float64(*in.Resources.Cpu)
		}
		if in.Resources.DiskGb != nil {
			outRes.DiskGb = float64(*in.Resources.DiskGb)
		}
		for _, v := range in.Resources.Volumes {
			outRes.Volumes = append(outRes.Volumes, &pb.Volume{
				Name:       v.Name,
				VolumeType: v.VolumeType,
				SizeGb:     float64(v.SizeGb),
			})
		}
		out.Resources = outRes
	}

	if in.ReadinessCheck != nil {
		rc := in.ReadinessCheck
		out.ReadinessCheck = &pb.ReadinessCheck{
			Path:            rc.Path,
			Host:            direct.ValueOf(rc.Host),
			CheckInterval:   direct.StringDuration_ToProto(mapCtx, rc.CheckInterval),
			Timeout:         direct.StringDuration_ToProto(mapCtx, rc.Timeout),
			AppStartTimeout: direct.StringDuration_ToProto(mapCtx, rc.AppStartTimeout),
		}
		if rc.FailureThreshold != nil {
			out.ReadinessCheck.FailureThreshold = uint32(*rc.FailureThreshold)
		}
		if rc.SuccessThreshold != nil {
			out.ReadinessCheck.SuccessThreshold = uint32(*rc.SuccessThreshold)
		}
	}

	if in.LivenessCheck != nil {
		lc := in.LivenessCheck
		out.LivenessCheck = &pb.LivenessCheck{
			Path:          lc.Path,
			Host:          direct.ValueOf(lc.Host),
			CheckInterval: direct.StringDuration_ToProto(mapCtx, lc.CheckInterval),
			Timeout:       direct.StringDuration_ToProto(mapCtx, lc.Timeout),
			InitialDelay:  direct.StringDuration_ToProto(mapCtx, lc.InitialDelay),
		}
		if lc.FailureThreshold != nil {
			out.LivenessCheck.FailureThreshold = uint32(*lc.FailureThreshold)
		}
		if lc.SuccessThreshold != nil {
			out.LivenessCheck.SuccessThreshold = uint32(*lc.SuccessThreshold)
		}
	}

	if in.Deployment != nil {
		dep := in.Deployment
		outDep := &pb.Deployment{}
		if dep.CloudBuildOptions != nil {
			outDep.CloudBuildOptions = &pb.CloudBuildOptions{
				AppYamlPath:       dep.CloudBuildOptions.AppYamlPath,
				CloudBuildTimeout: direct.StringDuration_ToProto(mapCtx, dep.CloudBuildOptions.CloudBuildTimeout),
			}
		}
		if dep.Container != nil {
			outDep.Container = &pb.ContainerInfo{
				Image: dep.Container.Image,
			}
		}
		if dep.Zip != nil {
			outDep.Zip = &pb.ZipInfo{
				SourceUrl:  dep.Zip.SourceUrl,
				FilesCount: int32(direct.ValueOf(dep.Zip.FilesCount)),
			}
		}
		if len(dep.Files) > 0 {
			outDep.Files = make(map[string]*pb.FileInfo, len(dep.Files))
			for _, f := range dep.Files {
				outDep.Files[f.Name] = &pb.FileInfo{
					Sha1Sum:   direct.ValueOf(f.Sha1Sum),
					SourceUrl: f.SourceUrl,
				}
			}
		}
		out.Deployment = outDep
	}

	if in.Entrypoint != nil {
		out.Entrypoint = &pb.Entrypoint{
			Command: &pb.Entrypoint_Shell{
				Shell: in.Entrypoint.Shell,
			},
		}
	}

	if in.EndpointsApiService != nil {
		eas := in.EndpointsApiService
		out.EndpointsApiService = &pb.EndpointsApiService{
			Name:                 eas.Name,
			ConfigId:             direct.ValueOf(eas.ConfigId),
			RolloutStrategy:      direct.Enum_ToProto[pb.EndpointsApiService_RolloutStrategy](mapCtx, eas.RolloutStrategy),
			DisableTraceSampling: direct.ValueOf(eas.DisableTraceSampling),
		}
	}

	if in.VpcAccessConnector != nil {
		out.VpcAccessConnector = &pb.VpcAccessConnector{
			Name: in.VpcAccessConnector.Name,
		}
	}

	if in.ApiConfig != nil {
		ac := in.ApiConfig
		out.ApiConfig = &pb.ApiConfigHandler{
			AuthFailAction: direct.Enum_ToProto[pb.AuthFailAction](mapCtx, ac.AuthFailAction),
			Login:          direct.Enum_ToProto[pb.LoginRequirement](mapCtx, ac.Login),
			Script:         ac.Script,
			SecurityLevel:  direct.Enum_ToProto[pb.SecurityLevel](mapCtx, ac.SecurityLevel),
			Url:            direct.ValueOf(ac.Url),
		}
	}

	for _, h := range in.Handlers {
		outH := &pb.UrlMap{
			AuthFailAction:           direct.Enum_ToProto[pb.AuthFailAction](mapCtx, h.AuthFailAction),
			Login:                    direct.Enum_ToProto[pb.LoginRequirement](mapCtx, h.Login),
			RedirectHttpResponseCode: direct.Enum_ToProto[pb.UrlMap_RedirectHttpResponseCode](mapCtx, h.RedirectHttpResponseCode),
			SecurityLevel:            direct.Enum_ToProto[pb.SecurityLevel](mapCtx, h.SecurityLevel),
			UrlRegex:                 direct.ValueOf(h.UrlRegex),
		}
		if h.Script != nil {
			outH.HandlerType = &pb.UrlMap_Script{
				Script: &pb.ScriptHandler{
					ScriptPath: h.Script.ScriptPath,
				},
			}
		}
		if h.StaticFiles != nil {
			sf := h.StaticFiles
			outH.HandlerType = &pb.UrlMap_StaticFiles{
				StaticFiles: &pb.StaticFilesHandler{
					ApplicationReadable: direct.ValueOf(sf.ApplicationReadable),
					Expiration:          direct.StringDuration_ToProto(mapCtx, sf.Expiration),
					HttpHeaders:         sf.HttpHeaders,
					MimeType:            direct.ValueOf(sf.MimeType),
					Path:                direct.ValueOf(sf.Path),
					RequireMatchingFile: direct.ValueOf(sf.RequireMatchingFile),
					UploadPathRegex:     direct.ValueOf(sf.UploadPathRegex),
				},
			}
		}
		out.Handlers = append(out.Handlers, outH)
	}

	return out
}

// AppEngineFlexibleAppVersionStatus_v1alpha1_FromProto converts the protobuf Version to the KRM AppEngineFlexibleAppVersionStatus.
func AppEngineFlexibleAppVersionStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krmappenginev1alpha1.AppEngineFlexibleAppVersionStatus {
	if in == nil {
		return nil
	}
	out := &krmappenginev1alpha1.AppEngineFlexibleAppVersionStatus{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}

// AppEngineFlexibleAppVersionStatus_v1alpha1_ToProto converts the KRM AppEngineFlexibleAppVersionStatus to the protobuf Version.
func AppEngineFlexibleAppVersionStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmappenginev1alpha1.AppEngineFlexibleAppVersionStatus) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
