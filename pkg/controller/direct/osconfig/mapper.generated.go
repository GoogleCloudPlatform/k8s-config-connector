// Copyright 2025 Google LLC
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

package osconfig

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
)
func OSPolicyAssignmentReport_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignmentReport) *krm.OSPolicyAssignmentReport {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignmentReport{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Instance = direct.LazyPtr(in.GetInstance())
	out.OsPolicyAssignment = direct.LazyPtr(in.GetOsPolicyAssignment())
	out.OsPolicyCompliances = direct.Slice_FromProto(mapCtx, in.OsPolicyCompliances, OSPolicyAssignmentReport_OSPolicyCompliance_FromProto)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LastRunID = direct.LazyPtr(in.GetLastRunId())
	return out
}
func OSPolicyAssignmentReport_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignmentReport) *pb.OSPolicyAssignmentReport {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignmentReport{}
	out.Name = direct.ValueOf(in.Name)
	out.Instance = direct.ValueOf(in.Instance)
	out.OsPolicyAssignment = direct.ValueOf(in.OsPolicyAssignment)
	out.OsPolicyCompliances = direct.Slice_ToProto(mapCtx, in.OsPolicyCompliances, OSPolicyAssignmentReport_OSPolicyCompliance_ToProto)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LastRunId = direct.ValueOf(in.LastRunID)
	return out
}
func OSPolicyAssignmentReport_OSPolicyCompliance_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignmentReport_OSPolicyCompliance) *krm.OSPolicyAssignmentReport_OSPolicyCompliance {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignmentReport_OSPolicyCompliance{}
	out.OsPolicyID = direct.LazyPtr(in.GetOsPolicyId())
	out.ComplianceState = direct.Enum_FromProto(mapCtx, in.GetComplianceState())
	out.ComplianceStateReason = direct.LazyPtr(in.GetComplianceStateReason())
	out.OsPolicyResourceCompliances = direct.Slice_FromProto(mapCtx, in.OsPolicyResourceCompliances, OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_FromProto)
	return out
}
func OSPolicyAssignmentReport_OSPolicyCompliance_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignmentReport_OSPolicyCompliance) *pb.OSPolicyAssignmentReport_OSPolicyCompliance {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignmentReport_OSPolicyCompliance{}
	out.OsPolicyId = direct.ValueOf(in.OsPolicyID)
	out.ComplianceState = direct.Enum_ToProto[pb.OSPolicyAssignmentReport_OSPolicyCompliance_ComplianceState](mapCtx, in.ComplianceState)
	out.ComplianceStateReason = direct.ValueOf(in.ComplianceStateReason)
	out.OsPolicyResourceCompliances = direct.Slice_ToProto(mapCtx, in.OsPolicyResourceCompliances, OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ToProto)
	return out
}
func OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance) *krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance{}
	out.OsPolicyResourceID = direct.LazyPtr(in.GetOsPolicyResourceId())
	out.ConfigSteps = direct.Slice_FromProto(mapCtx, in.ConfigSteps, OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep_FromProto)
	out.ComplianceState = direct.Enum_FromProto(mapCtx, in.GetComplianceState())
	out.ComplianceStateReason = direct.LazyPtr(in.GetComplianceStateReason())
	out.ExecResourceOutput = OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput_FromProto(mapCtx, in.GetExecResourceOutput())
	return out
}
func OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance) *pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance{}
	out.OsPolicyResourceId = direct.ValueOf(in.OsPolicyResourceID)
	out.ConfigSteps = direct.Slice_ToProto(mapCtx, in.ConfigSteps, OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep_ToProto)
	out.ComplianceState = direct.Enum_ToProto[pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ComplianceState](mapCtx, in.ComplianceState)
	out.ComplianceStateReason = direct.ValueOf(in.ComplianceStateReason)
	if oneof := OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput_ToProto(mapCtx, in.ExecResourceOutput); oneof != nil {
		out.Output = &pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput_{ExecResourceOutput: oneof}
	}
	return out
}
func OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput) *krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput{}
	out.EnforcementOutput = in.GetEnforcementOutput()
	return out
}
func OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput) *pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_ExecResourceOutput{}
	out.EnforcementOutput = in.EnforcementOutput
	return out
}
func OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep) *krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	return out
}
func OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep) *pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep{}
	out.Type = direct.Enum_ToProto[pb.OSPolicyAssignmentReport_OSPolicyCompliance_OSPolicyResourceCompliance_OSPolicyResourceConfigStep_Type](mapCtx, in.Type)
	out.ErrorMessage = direct.ValueOf(in.ErrorMessage)
	return out
}
func OsconfigOSPolicyAssignmentReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignmentReport) *krm.OsconfigOSPolicyAssignmentReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigOSPolicyAssignmentReportObservedState{}
	// MISSING: Name
	// MISSING: Instance
	// MISSING: OsPolicyAssignment
	// MISSING: OsPolicyCompliances
	// MISSING: UpdateTime
	// MISSING: LastRunID
	return out
}
func OsconfigOSPolicyAssignmentReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigOSPolicyAssignmentReportObservedState) *pb.OSPolicyAssignmentReport {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignmentReport{}
	// MISSING: Name
	// MISSING: Instance
	// MISSING: OsPolicyAssignment
	// MISSING: OsPolicyCompliances
	// MISSING: UpdateTime
	// MISSING: LastRunID
	return out
}
func OsconfigOSPolicyAssignmentReportSpec_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignmentReport) *krm.OsconfigOSPolicyAssignmentReportSpec {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigOSPolicyAssignmentReportSpec{}
	// MISSING: Name
	// MISSING: Instance
	// MISSING: OsPolicyAssignment
	// MISSING: OsPolicyCompliances
	// MISSING: UpdateTime
	// MISSING: LastRunID
	return out
}
func OsconfigOSPolicyAssignmentReportSpec_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigOSPolicyAssignmentReportSpec) *pb.OSPolicyAssignmentReport {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignmentReport{}
	// MISSING: Name
	// MISSING: Instance
	// MISSING: OsPolicyAssignment
	// MISSING: OsPolicyCompliances
	// MISSING: UpdateTime
	// MISSING: LastRunID
	return out
}
