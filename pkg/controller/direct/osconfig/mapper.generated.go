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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CVSSv3_FromProto(mapCtx *direct.MapContext, in *pb.CVSSv3) *krm.CVSSv3 {
	if in == nil {
		return nil
	}
	out := &krm.CVSSv3{}
	out.BaseScore = direct.LazyPtr(in.GetBaseScore())
	out.ExploitabilityScore = direct.LazyPtr(in.GetExploitabilityScore())
	out.ImpactScore = direct.LazyPtr(in.GetImpactScore())
	out.AttackVector = direct.Enum_FromProto(mapCtx, in.GetAttackVector())
	out.AttackComplexity = direct.Enum_FromProto(mapCtx, in.GetAttackComplexity())
	out.PrivilegesRequired = direct.Enum_FromProto(mapCtx, in.GetPrivilegesRequired())
	out.UserInteraction = direct.Enum_FromProto(mapCtx, in.GetUserInteraction())
	out.Scope = direct.Enum_FromProto(mapCtx, in.GetScope())
	out.ConfidentialityImpact = direct.Enum_FromProto(mapCtx, in.GetConfidentialityImpact())
	out.IntegrityImpact = direct.Enum_FromProto(mapCtx, in.GetIntegrityImpact())
	out.AvailabilityImpact = direct.Enum_FromProto(mapCtx, in.GetAvailabilityImpact())
	return out
}
func CVSSv3_ToProto(mapCtx *direct.MapContext, in *krm.CVSSv3) *pb.CVSSv3 {
	if in == nil {
		return nil
	}
	out := &pb.CVSSv3{}
	out.BaseScore = direct.ValueOf(in.BaseScore)
	out.ExploitabilityScore = direct.ValueOf(in.ExploitabilityScore)
	out.ImpactScore = direct.ValueOf(in.ImpactScore)
	out.AttackVector = direct.Enum_ToProto[pb.CVSSv3_AttackVector](mapCtx, in.AttackVector)
	out.AttackComplexity = direct.Enum_ToProto[pb.CVSSv3_AttackComplexity](mapCtx, in.AttackComplexity)
	out.PrivilegesRequired = direct.Enum_ToProto[pb.CVSSv3_PrivilegesRequired](mapCtx, in.PrivilegesRequired)
	out.UserInteraction = direct.Enum_ToProto[pb.CVSSv3_UserInteraction](mapCtx, in.UserInteraction)
	out.Scope = direct.Enum_ToProto[pb.CVSSv3_Scope](mapCtx, in.Scope)
	out.ConfidentialityImpact = direct.Enum_ToProto[pb.CVSSv3_Impact](mapCtx, in.ConfidentialityImpact)
	out.IntegrityImpact = direct.Enum_ToProto[pb.CVSSv3_Impact](mapCtx, in.IntegrityImpact)
	out.AvailabilityImpact = direct.Enum_ToProto[pb.CVSSv3_Impact](mapCtx, in.AvailabilityImpact)
	return out
}
func OsconfigVulnerabilityReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityReport) *krm.OsconfigVulnerabilityReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigVulnerabilityReportObservedState{}
	// MISSING: Name
	// MISSING: Vulnerabilities
	// MISSING: UpdateTime
	return out
}
func OsconfigVulnerabilityReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigVulnerabilityReportObservedState) *pb.VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityReport{}
	// MISSING: Name
	// MISSING: Vulnerabilities
	// MISSING: UpdateTime
	return out
}
func OsconfigVulnerabilityReportSpec_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityReport) *krm.OsconfigVulnerabilityReportSpec {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigVulnerabilityReportSpec{}
	// MISSING: Name
	// MISSING: Vulnerabilities
	// MISSING: UpdateTime
	return out
}
func OsconfigVulnerabilityReportSpec_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigVulnerabilityReportSpec) *pb.VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityReport{}
	// MISSING: Name
	// MISSING: Vulnerabilities
	// MISSING: UpdateTime
	return out
}
func VulnerabilityReport_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityReport) *krm.VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := &krm.VulnerabilityReport{}
	// MISSING: Name
	// MISSING: Vulnerabilities
	// MISSING: UpdateTime
	return out
}
func VulnerabilityReport_ToProto(mapCtx *direct.MapContext, in *krm.VulnerabilityReport) *pb.VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityReport{}
	// MISSING: Name
	// MISSING: Vulnerabilities
	// MISSING: UpdateTime
	return out
}
func VulnerabilityReportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityReport) *krm.VulnerabilityReportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VulnerabilityReportObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Vulnerabilities = direct.Slice_FromProto(mapCtx, in.Vulnerabilities, VulnerabilityReport_Vulnerability_FromProto)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func VulnerabilityReportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VulnerabilityReportObservedState) *pb.VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityReport{}
	out.Name = direct.ValueOf(in.Name)
	out.Vulnerabilities = direct.Slice_ToProto(mapCtx, in.Vulnerabilities, VulnerabilityReport_Vulnerability_ToProto)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func VulnerabilityReport_Vulnerability_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityReport_Vulnerability) *krm.VulnerabilityReport_Vulnerability {
	if in == nil {
		return nil
	}
	out := &krm.VulnerabilityReport_Vulnerability{}
	out.Details = VulnerabilityReport_Vulnerability_Details_FromProto(mapCtx, in.GetDetails())
	out.InstalledInventoryItemIds = in.InstalledInventoryItemIds
	out.AvailableInventoryItemIds = in.AvailableInventoryItemIds
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, VulnerabilityReport_Vulnerability_Item_FromProto)
	return out
}
func VulnerabilityReport_Vulnerability_ToProto(mapCtx *direct.MapContext, in *krm.VulnerabilityReport_Vulnerability) *pb.VulnerabilityReport_Vulnerability {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityReport_Vulnerability{}
	out.Details = VulnerabilityReport_Vulnerability_Details_ToProto(mapCtx, in.Details)
	out.InstalledInventoryItemIds = in.InstalledInventoryItemIds
	out.AvailableInventoryItemIds = in.AvailableInventoryItemIds
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, VulnerabilityReport_Vulnerability_Item_ToProto)
	return out
}
func VulnerabilityReport_Vulnerability_Details_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityReport_Vulnerability_Details) *krm.VulnerabilityReport_Vulnerability_Details {
	if in == nil {
		return nil
	}
	out := &krm.VulnerabilityReport_Vulnerability_Details{}
	out.Cve = direct.LazyPtr(in.GetCve())
	out.CvssV2Score = direct.LazyPtr(in.GetCvssV2Score())
	out.CvssV3 = CVSSv3_FromProto(mapCtx, in.GetCvssV3())
	out.Severity = direct.LazyPtr(in.GetSeverity())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.References = direct.Slice_FromProto(mapCtx, in.References, VulnerabilityReport_Vulnerability_Details_Reference_FromProto)
	return out
}
func VulnerabilityReport_Vulnerability_Details_ToProto(mapCtx *direct.MapContext, in *krm.VulnerabilityReport_Vulnerability_Details) *pb.VulnerabilityReport_Vulnerability_Details {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityReport_Vulnerability_Details{}
	out.Cve = direct.ValueOf(in.Cve)
	out.CvssV2Score = direct.ValueOf(in.CvssV2Score)
	out.CvssV3 = CVSSv3_ToProto(mapCtx, in.CvssV3)
	out.Severity = direct.ValueOf(in.Severity)
	out.Description = direct.ValueOf(in.Description)
	out.References = direct.Slice_ToProto(mapCtx, in.References, VulnerabilityReport_Vulnerability_Details_Reference_ToProto)
	return out
}
func VulnerabilityReport_Vulnerability_Details_Reference_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityReport_Vulnerability_Details_Reference) *krm.VulnerabilityReport_Vulnerability_Details_Reference {
	if in == nil {
		return nil
	}
	out := &krm.VulnerabilityReport_Vulnerability_Details_Reference{}
	out.URL = direct.LazyPtr(in.GetUrl())
	out.Source = direct.LazyPtr(in.GetSource())
	return out
}
func VulnerabilityReport_Vulnerability_Details_Reference_ToProto(mapCtx *direct.MapContext, in *krm.VulnerabilityReport_Vulnerability_Details_Reference) *pb.VulnerabilityReport_Vulnerability_Details_Reference {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityReport_Vulnerability_Details_Reference{}
	out.Url = direct.ValueOf(in.URL)
	out.Source = direct.ValueOf(in.Source)
	return out
}
func VulnerabilityReport_Vulnerability_Item_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityReport_Vulnerability_Item) *krm.VulnerabilityReport_Vulnerability_Item {
	if in == nil {
		return nil
	}
	out := &krm.VulnerabilityReport_Vulnerability_Item{}
	out.InstalledInventoryItemID = direct.LazyPtr(in.GetInstalledInventoryItemId())
	out.AvailableInventoryItemID = direct.LazyPtr(in.GetAvailableInventoryItemId())
	out.FixedCpeURI = direct.LazyPtr(in.GetFixedCpeUri())
	out.UpstreamFix = direct.LazyPtr(in.GetUpstreamFix())
	return out
}
func VulnerabilityReport_Vulnerability_Item_ToProto(mapCtx *direct.MapContext, in *krm.VulnerabilityReport_Vulnerability_Item) *pb.VulnerabilityReport_Vulnerability_Item {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityReport_Vulnerability_Item{}
	out.InstalledInventoryItemId = direct.ValueOf(in.InstalledInventoryItemID)
	out.AvailableInventoryItemId = direct.ValueOf(in.AvailableInventoryItemID)
	out.FixedCpeUri = direct.ValueOf(in.FixedCpeURI)
	out.UpstreamFix = direct.ValueOf(in.UpstreamFix)
	return out
}
