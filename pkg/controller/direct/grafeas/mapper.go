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

package grafeas

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/grafeas/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/grafeas/v1"
)

func VulnerabilityAssessmentNote_Assessment_FromProto(mapCtx *direct.MapContext, in *pb.VulnerabilityAssessmentNote_Assessment) *krm.VulnerabilityAssessmentNote_Assessment {
	if in == nil {
		return nil
	}
	out := &krm.VulnerabilityAssessmentNote_Assessment{}
	out.Cve = direct.LazyPtr(in.GetCve())
	out.VulnerabilityID = direct.LazyPtr(in.GetVulnerabilityId())
	out.ShortDescription = direct.LazyPtr(in.GetShortDescription())
	out.LongDescription = direct.LazyPtr(in.GetLongDescription())
	out.RelatedURIs = direct.Slice_FromProto(mapCtx, in.RelatedUris, RelatedURL_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Impacts = in.Impacts
	out.Justification = VulnerabilityAssessmentNote_Assessment_Justification_FromProto(mapCtx, in.GetJustification())
	out.Remediations = direct.Slice_FromProto(mapCtx, in.Remediations, VulnerabilityAssessmentNote_Assessment_Remediation_FromProto)
	return out
}

func VulnerabilityAssessmentNote_Assessment_ToProto(mapCtx *direct.MapContext, in *krm.VulnerabilityAssessmentNote_Assessment) *pb.VulnerabilityAssessmentNote_Assessment {
	if in == nil {
		return nil
	}
	out := &pb.VulnerabilityAssessmentNote_Assessment{}
	out.Cve = direct.ValueOf(in.Cve)
	out.VulnerabilityId = direct.ValueOf(in.VulnerabilityID)
	out.ShortDescription = direct.ValueOf(in.ShortDescription)
	out.LongDescription = direct.ValueOf(in.LongDescription)
	out.RelatedUris = direct.Slice_ToProto(mapCtx, in.RelatedURIs, RelatedURL_ToProto)
	out.State = direct.Enum_ToProto[pb.VulnerabilityAssessmentNote_Assessment_State](mapCtx, in.State)
	out.Impacts = in.Impacts
	out.Justification = VulnerabilityAssessmentNote_Assessment_Justification_ToProto(mapCtx, in.Justification)
	out.Remediations = direct.Slice_ToProto(mapCtx, in.Remediations, VulnerabilityAssessmentNote_Assessment_Remediation_ToProto)
	return out
}

func WindowsUpdate_FromProto(mapCtx *direct.MapContext, in *pb.WindowsUpdate) *krm.WindowsUpdate {
	if in == nil {
		return nil
	}
	out := &krm.WindowsUpdate{}
	out.Identity = WindowsUpdate_Identity_FromProto(mapCtx, in.GetIdentity())
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Categories = direct.Slice_FromProto(mapCtx, in.Categories, WindowsUpdate_Category_FromProto)
	out.KbArticleIDs = in.KbArticleIds
	out.SupportURL = direct.LazyPtr(in.GetSupportUrl())
	out.LastPublishedTimestamp = direct.StringTimestamp_FromProto(mapCtx, in.GetLastPublishedTimestamp())
	return out
}

func WindowsUpdate_ToProto(mapCtx *direct.MapContext, in *krm.WindowsUpdate) *pb.WindowsUpdate {
	if in == nil {
		return nil
	}
	out := &pb.WindowsUpdate{}
	out.Identity = WindowsUpdate_Identity_ToProto(mapCtx, in.Identity)
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Categories = direct.Slice_ToProto(mapCtx, in.Categories, WindowsUpdate_Category_ToProto)
	out.KbArticleIds = in.KbArticleIDs
	out.SupportUrl = direct.ValueOf(in.SupportURL)
	out.LastPublishedTimestamp = direct.StringTimestamp_ToProto(mapCtx, in.LastPublishedTimestamp)
	return out
}
