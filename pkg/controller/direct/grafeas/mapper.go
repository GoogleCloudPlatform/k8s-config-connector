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

func GrafeasNoteSpec_ToProto(mapCtx *direct.MapContext, in *krm.GrafeasNoteSpec) *pb.Note {
	if in == nil {
		return nil
	}
	out := &pb.Note{}
	out.ShortDescription = direct.ValueOf(in.ShortDescription)
	out.LongDescription = direct.ValueOf(in.LongDescription)
	out.RelatedUrl = direct.Slice_ToProto(mapCtx, in.RelatedURL, RelatedURL_ToProto)
	out.ExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpirationTime)
	out.RelatedNoteNames = in.RelatedNoteNames

	// Map oneof types
	if in.Vulnerability != nil {
		out.Type = &pb.Note_Vulnerability{Vulnerability: VulnerabilityNote_ToProto(mapCtx, in.Vulnerability)}
	} else if in.Build != nil {
		out.Type = &pb.Note_Build{Build: BuildNote_ToProto(mapCtx, in.Build)}
	} else if in.Image != nil {
		out.Type = &pb.Note_Image{Image: ImageNote_ToProto(mapCtx, in.Image)}
	} else if in.Package != nil {
		out.Type = &pb.Note_Package{Package: PackageNote_ToProto(mapCtx, in.Package)}
	} else if in.Deployment != nil {
		out.Type = &pb.Note_Deployment{Deployment: DeploymentNote_ToProto(mapCtx, in.Deployment)}
	} else if in.Discovery != nil {
		out.Type = &pb.Note_Discovery{Discovery: DiscoveryNote_ToProto(mapCtx, in.Discovery)}
	} else if in.Attestation != nil {
		out.Type = &pb.Note_Attestation{Attestation: AttestationNote_ToProto(mapCtx, in.Attestation)}
	} else if in.Upgrade != nil {
		out.Type = &pb.Note_Upgrade{Upgrade: UpgradeNote_ToProto(mapCtx, in.Upgrade)}
	} else if in.Compliance != nil {
		out.Type = &pb.Note_Compliance{Compliance: ComplianceNote_ToProto(mapCtx, in.Compliance)}
	} else if in.DsseAttestation != nil {
		out.Type = &pb.Note_DsseAttestation{DsseAttestation: DsseAttestationNote_ToProto(mapCtx, in.DsseAttestation)}
	} else if in.VulnerabilityAssessment != nil {
		out.Type = &pb.Note_VulnerabilityAssessment{VulnerabilityAssessment: VulnerabilityAssessmentNote_ToProto(mapCtx, in.VulnerabilityAssessment)}
	} else if in.SbomReference != nil {
		out.Type = &pb.Note_SbomReference{SbomReference: SbomReferenceNote_ToProto(mapCtx, in.SbomReference)}
	}

	// Patch fields missed due to casing differences in code generator
	if in.VulnerabilityAssessment != nil && out.GetVulnerabilityAssessment() != nil {
		if in.VulnerabilityAssessment.Assessment != nil && out.GetVulnerabilityAssessment().GetAssessment() != nil {
			out.GetVulnerabilityAssessment().GetAssessment().RelatedUris = direct.Slice_ToProto(mapCtx, in.VulnerabilityAssessment.Assessment.RelatedURIs, RelatedURL_ToProto)
		}
	}
	if in.Upgrade != nil && out.GetUpgrade() != nil {
		if in.Upgrade.WindowsUpdate != nil && out.GetUpgrade().GetWindowsUpdate() != nil {
			out.GetUpgrade().GetWindowsUpdate().KbArticleIds = in.Upgrade.WindowsUpdate.KbArticleIDs
		}
	}

	return out
}

func GrafeasNoteSpec_FromProto(mapCtx *direct.MapContext, in *pb.Note) *krm.GrafeasNoteSpec {
	if in == nil {
		return nil
	}
	out := &krm.GrafeasNoteSpec{}
	out.ShortDescription = direct.LazyPtr(in.GetShortDescription())
	out.LongDescription = direct.LazyPtr(in.GetLongDescription())
	out.RelatedURL = direct.Slice_FromProto(mapCtx, in.GetRelatedUrl(), RelatedURL_FromProto)
	out.ExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpirationTime())
	out.RelatedNoteNames = in.GetRelatedNoteNames()

	// Map oneof types
	switch t := in.Type.(type) {
	case *pb.Note_Vulnerability:
		out.Vulnerability = VulnerabilityNote_FromProto(mapCtx, t.Vulnerability)
	case *pb.Note_Build:
		out.Build = BuildNote_FromProto(mapCtx, t.Build)
	case *pb.Note_Image:
		out.Image = ImageNote_FromProto(mapCtx, t.Image)
	case *pb.Note_Package:
		out.Package = PackageNote_FromProto(mapCtx, t.Package)
	case *pb.Note_Deployment:
		out.Deployment = DeploymentNote_FromProto(mapCtx, t.Deployment)
	case *pb.Note_Discovery:
		out.Discovery = DiscoveryNote_FromProto(mapCtx, t.Discovery)
	case *pb.Note_Attestation:
		out.Attestation = AttestationNote_FromProto(mapCtx, t.Attestation)
	case *pb.Note_Upgrade:
		out.Upgrade = UpgradeNote_FromProto(mapCtx, t.Upgrade)
	case *pb.Note_Compliance:
		out.Compliance = ComplianceNote_FromProto(mapCtx, t.Compliance)
	case *pb.Note_DsseAttestation:
		out.DsseAttestation = DsseAttestationNote_FromProto(mapCtx, t.DsseAttestation)
	case *pb.Note_VulnerabilityAssessment:
		out.VulnerabilityAssessment = VulnerabilityAssessmentNote_FromProto(mapCtx, t.VulnerabilityAssessment)
	case *pb.Note_SbomReference:
		out.SbomReference = SbomReferenceNote_FromProto(mapCtx, t.SbomReference)
	}

	// Patch fields missed due to casing differences in code generator
	if in.GetVulnerabilityAssessment() != nil && out.VulnerabilityAssessment != nil {
		if in.GetVulnerabilityAssessment().GetAssessment() != nil && out.VulnerabilityAssessment.Assessment != nil {
			out.VulnerabilityAssessment.Assessment.RelatedURIs = direct.Slice_FromProto(mapCtx, in.GetVulnerabilityAssessment().GetAssessment().GetRelatedUris(), RelatedURL_FromProto)
		}
	}
	if in.GetUpgrade() != nil && out.Upgrade != nil {
		if in.GetUpgrade().GetWindowsUpdate() != nil && out.Upgrade.WindowsUpdate != nil {
			out.Upgrade.WindowsUpdate.KbArticleIDs = in.GetUpgrade().GetWindowsUpdate().GetKbArticleIds()
		}
	}

	return out
}

func GrafeasNoteObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Note) *krm.GrafeasNoteObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GrafeasNoteObservedState{}
	kindStr := in.GetKind().String()
	out.Kind = &kindStr
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func GrafeasNoteObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GrafeasNoteObservedState) *pb.Note {
	if in == nil {
		return nil
	}
	out := &pb.Note{}
	if in.Kind != nil {
		if val, ok := pb.NoteKind_value[*in.Kind]; ok {
			out.Kind = pb.NoteKind(val)
		}
	}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func GrafeasNoteStatus_FromProto(mapCtx *direct.MapContext, in *pb.Note) *krm.GrafeasNoteStatus {
	if in == nil {
		return nil
	}
	out := &krm.GrafeasNoteStatus{}
	out.ObservedState = GrafeasNoteObservedState_FromProto(mapCtx, in)
	return out
}
