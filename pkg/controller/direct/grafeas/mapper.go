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
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
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

func Any_FromProto(mapCtx *direct.MapContext, in *anypb.Any) *krm.Any {
	if in == nil {
		return nil
	}
	return &krm.Any{
		TypeURL: direct.LazyPtr(in.GetTypeUrl()),
		Value:   in.GetValue(),
	}
}

func Any_ToProto(mapCtx *direct.MapContext, in *krm.Any) *anypb.Any {
	if in == nil {
		return nil
	}
	return &anypb.Any{
		TypeUrl: direct.ValueOf(in.TypeURL),
		Value:   in.Value,
	}
}

func Annotations_FromProto(mapCtx *direct.MapContext, in map[string]*structpb.Value) map[string]krm.Value {
	return nil
}

func Annotations_ToProto(mapCtx *direct.MapContext, in map[string]krm.Value) map[string]*structpb.Value {
	return nil
}

func Metadata_FromProto(mapCtx *direct.MapContext, in *pb.Metadata) *krm.Metadata {
	if in == nil {
		return nil
	}
	return &krm.Metadata{
		BuildInvocationID: direct.LazyPtr(in.GetBuildInvocationId()),
		BuildStartedOn:    direct.StringTimestamp_FromProto(mapCtx, in.GetBuildStartedOn()),
		BuildFinishedOn:   direct.StringTimestamp_FromProto(mapCtx, in.GetBuildFinishedOn()),
		Completeness:      Completeness_FromProto(mapCtx, in.GetCompleteness()),
		Reproducible:      direct.LazyPtr(in.GetReproducible()),
	}
}

func Metadata_ToProto(mapCtx *direct.MapContext, in *krm.Metadata) *pb.Metadata {
	if in == nil {
		return nil
	}
	return &pb.Metadata{
		BuildInvocationId: direct.ValueOf(in.BuildInvocationID),
		BuildStartedOn:    direct.StringTimestamp_ToProto(mapCtx, in.BuildStartedOn),
		BuildFinishedOn:   direct.StringTimestamp_ToProto(mapCtx, in.BuildFinishedOn),
		Completeness:      Completeness_ToProto(mapCtx, in.Completeness),
		Reproducible:      direct.ValueOf(in.Reproducible),
	}
}

func InTotoSlsaProvenanceV1_BuildMetadata_FromProto(mapCtx *direct.MapContext, in *pb.InTotoSlsaProvenanceV1_BuildMetadata) *krm.InTotoSlsaProvenanceV1_BuildMetadata {
	if in == nil {
		return nil
	}
	return &krm.InTotoSlsaProvenanceV1_BuildMetadata{
		InvocationID: direct.LazyPtr(in.GetInvocationId()),
		StartedOn:    direct.StringTimestamp_FromProto(mapCtx, in.GetStartedOn()),
		FinishedOn:   direct.StringTimestamp_FromProto(mapCtx, in.GetFinishedOn()),
	}
}

func InTotoSlsaProvenanceV1_BuildMetadata_ToProto(mapCtx *direct.MapContext, in *krm.InTotoSlsaProvenanceV1_BuildMetadata) *pb.InTotoSlsaProvenanceV1_BuildMetadata {
	if in == nil {
		return nil
	}
	return &pb.InTotoSlsaProvenanceV1_BuildMetadata{
		InvocationId: direct.ValueOf(in.InvocationID),
		StartedOn:    direct.StringTimestamp_ToProto(mapCtx, in.StartedOn),
		FinishedOn:   direct.StringTimestamp_ToProto(mapCtx, in.FinishedOn),
	}
}

func SlsaProvenance_SlsaMetadata_FromProto(mapCtx *direct.MapContext, in *pb.SlsaProvenance_SlsaMetadata) *krm.SlsaProvenance_SlsaMetadata {
	if in == nil {
		return nil
	}
	return &krm.SlsaProvenance_SlsaMetadata{
		BuildInvocationID: direct.LazyPtr(in.GetBuildInvocationId()),
		BuildStartedOn:    direct.StringTimestamp_FromProto(mapCtx, in.GetBuildStartedOn()),
		BuildFinishedOn:   direct.StringTimestamp_FromProto(mapCtx, in.GetBuildFinishedOn()),
		Completeness:      SlsaProvenance_SlsaCompleteness_FromProto(mapCtx, in.GetCompleteness()),
		Reproducible:      direct.LazyPtr(in.GetReproducible()),
	}
}

func SlsaProvenance_SlsaMetadata_ToProto(mapCtx *direct.MapContext, in *krm.SlsaProvenance_SlsaMetadata) *pb.SlsaProvenance_SlsaMetadata {
	if in == nil {
		return nil
	}
	return &pb.SlsaProvenance_SlsaMetadata{
		BuildInvocationId: direct.ValueOf(in.BuildInvocationID),
		BuildStartedOn:    direct.StringTimestamp_ToProto(mapCtx, in.BuildStartedOn),
		BuildFinishedOn:   direct.StringTimestamp_ToProto(mapCtx, in.BuildFinishedOn),
		Completeness:      SlsaProvenance_SlsaCompleteness_ToProto(mapCtx, in.Completeness),
		Reproducible:      direct.ValueOf(in.Reproducible),
	}
}

func SlsaProvenanceZeroTwo_SlsaMetadata_FromProto(mapCtx *direct.MapContext, in *pb.SlsaProvenanceZeroTwo_SlsaMetadata) *krm.SlsaProvenanceZeroTwo_SlsaMetadata {
	if in == nil {
		return nil
	}
	return &krm.SlsaProvenanceZeroTwo_SlsaMetadata{
		BuildInvocationID: direct.LazyPtr(in.GetBuildInvocationId()),
		BuildStartedOn:    direct.StringTimestamp_FromProto(mapCtx, in.GetBuildStartedOn()),
		BuildFinishedOn:   direct.StringTimestamp_FromProto(mapCtx, in.GetBuildFinishedOn()),
		Completeness:      SlsaProvenanceZeroTwo_SlsaCompleteness_FromProto(mapCtx, in.GetCompleteness()),
		Reproducible:      direct.LazyPtr(in.GetReproducible()),
	}
}

func SlsaProvenanceZeroTwo_SlsaMetadata_ToProto(mapCtx *direct.MapContext, in *krm.SlsaProvenanceZeroTwo_SlsaMetadata) *pb.SlsaProvenanceZeroTwo_SlsaMetadata {
	if in == nil {
		return nil
	}
	return &pb.SlsaProvenanceZeroTwo_SlsaMetadata{
		BuildInvocationId: direct.ValueOf(in.BuildInvocationID),
		BuildStartedOn:    direct.StringTimestamp_ToProto(mapCtx, in.BuildStartedOn),
		BuildFinishedOn:   direct.StringTimestamp_ToProto(mapCtx, in.BuildFinishedOn),
		Completeness:      SlsaProvenanceZeroTwo_SlsaCompleteness_ToProto(mapCtx, in.Completeness),
		Reproducible:      direct.ValueOf(in.Reproducible),
	}
}

func DiscoveryOccurrence_FromProto(mapCtx *direct.MapContext, in *pb.DiscoveryOccurrence) *krm.DiscoveryOccurrence {
	if in == nil {
		return nil
	}
	out := &krm.DiscoveryOccurrence{}
	out.ContinuousAnalysis = direct.Enum_FromProto(mapCtx, in.GetContinuousAnalysis())
	out.AnalysisStatus = direct.Enum_FromProto(mapCtx, in.GetAnalysisStatus())
	out.AnalysisCompleted = DiscoveryOccurrence_AnalysisCompleted_FromProto(mapCtx, in.GetAnalysisCompleted())
	out.AnalysisError = direct.Slice_FromProto(mapCtx, in.AnalysisError, direct.Status_FromProto)
	out.AnalysisStatusError = direct.Status_FromProto(mapCtx, in.GetAnalysisStatusError())
	out.Cpe = direct.LazyPtr(in.GetCpe())
	out.LastScanTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastScanTime())
	out.SbomStatus = DiscoveryOccurrence_SbomStatus_FromProto(mapCtx, in.GetSbomStatus())
	out.VulnerabilityAttestation = DiscoveryOccurrence_VulnerabilityAttestation_FromProto(mapCtx, in.GetVulnerabilityAttestation())
	return out
}

func DiscoveryOccurrence_ToProto(mapCtx *direct.MapContext, in *krm.DiscoveryOccurrence) *pb.DiscoveryOccurrence {
	if in == nil {
		return nil
	}
	out := &pb.DiscoveryOccurrence{}
	out.ContinuousAnalysis = direct.Enum_ToProto[pb.DiscoveryOccurrence_ContinuousAnalysis](mapCtx, in.ContinuousAnalysis)
	out.AnalysisStatus = direct.Enum_ToProto[pb.DiscoveryOccurrence_AnalysisStatus](mapCtx, in.AnalysisStatus)
	out.AnalysisCompleted = DiscoveryOccurrence_AnalysisCompleted_ToProto(mapCtx, in.AnalysisCompleted)
	out.AnalysisError = direct.Slice_ToProto(mapCtx, in.AnalysisError, direct.Status_ToProto)
	out.AnalysisStatusError = direct.Status_ToProto(mapCtx, in.AnalysisStatusError)
	out.Cpe = direct.ValueOf(in.Cpe)
	out.LastScanTime = direct.StringTimestamp_ToProto(mapCtx, in.LastScanTime)
	out.SbomStatus = DiscoveryOccurrence_SbomStatus_ToProto(mapCtx, in.SbomStatus)
	out.VulnerabilityAttestation = DiscoveryOccurrence_VulnerabilityAttestation_ToProto(mapCtx, in.VulnerabilityAttestation)
	return out
}

func InTotoSlsaProvenanceV1_FromProto(mapCtx *direct.MapContext, in *pb.InTotoSlsaProvenanceV1) *krm.InTotoSlsaProvenanceV1 {
	if in == nil {
		return nil
	}
	out := &krm.InTotoSlsaProvenanceV1{}
	out.Type = direct.LazyPtr(in.GetType())
	out.Subject = direct.Slice_FromProto(mapCtx, in.Subject, Subject_FromProto)
	out.PredicateType = direct.LazyPtr(in.GetPredicateType())
	out.Predicate = InTotoSlsaProvenanceV1_SlsaProvenanceV1_FromProto(mapCtx, in.GetPredicate())
	return out
}

func InTotoSlsaProvenanceV1_ToProto(mapCtx *direct.MapContext, in *krm.InTotoSlsaProvenanceV1) *pb.InTotoSlsaProvenanceV1 {
	if in == nil {
		return nil
	}
	out := &pb.InTotoSlsaProvenanceV1{}
	out.Type = direct.ValueOf(in.Type)
	out.Subject = direct.Slice_ToProto(mapCtx, in.Subject, Subject_ToProto)
	out.PredicateType = direct.ValueOf(in.PredicateType)
	out.Predicate = InTotoSlsaProvenanceV1_SlsaProvenanceV1_ToProto(mapCtx, in.Predicate)
	return out
}

func InTotoStatement_FromProto(mapCtx *direct.MapContext, in *pb.InTotoStatement) *krm.InTotoStatement {
	if in == nil {
		return nil
	}
	out := &krm.InTotoStatement{}
	out.Type = direct.LazyPtr(in.GetType())
	out.Subject = direct.Slice_FromProto(mapCtx, in.Subject, Subject_FromProto)
	out.PredicateType = direct.LazyPtr(in.GetPredicateType())
	out.Provenance = InTotoProvenance_FromProto(mapCtx, in.GetProvenance())
	out.SlsaProvenance = SlsaProvenance_FromProto(mapCtx, in.GetSlsaProvenance())
	out.SlsaProvenanceZeroTwo = SlsaProvenanceZeroTwo_FromProto(mapCtx, in.GetSlsaProvenanceZeroTwo())
	return out
}

func InTotoStatement_ToProto(mapCtx *direct.MapContext, in *krm.InTotoStatement) *pb.InTotoStatement {
	if in == nil {
		return nil
	}
	out := &pb.InTotoStatement{}
	out.Type = direct.ValueOf(in.Type)
	out.Subject = direct.Slice_ToProto(mapCtx, in.Subject, Subject_ToProto)
	out.PredicateType = direct.ValueOf(in.PredicateType)
	if oneof := InTotoProvenance_ToProto(mapCtx, in.Provenance); oneof != nil {
		out.Predicate = &pb.InTotoStatement_Provenance{Provenance: oneof}
	}
	if oneof := SlsaProvenance_ToProto(mapCtx, in.SlsaProvenance); oneof != nil {
		out.Predicate = &pb.InTotoStatement_SlsaProvenance{SlsaProvenance: oneof}
	}
	if oneof := SlsaProvenanceZeroTwo_ToProto(mapCtx, in.SlsaProvenanceZeroTwo); oneof != nil {
		out.Predicate = &pb.InTotoStatement_SlsaProvenanceZeroTwo{SlsaProvenanceZeroTwo: oneof}
	}
	return out
}

func SbomReferenceIntotoPayload_FromProto(mapCtx *direct.MapContext, in *pb.SbomReferenceIntotoPayload) *krm.SbomReferenceIntotoPayload {
	if in == nil {
		return nil
	}
	out := &krm.SbomReferenceIntotoPayload{}
	out.Type = direct.LazyPtr(in.GetType())
	out.PredicateType = direct.LazyPtr(in.GetPredicateType())
	out.Subject = direct.Slice_FromProto(mapCtx, in.Subject, Subject_FromProto)
	out.Predicate = SbomReferenceIntotoPredicate_FromProto(mapCtx, in.GetPredicate())
	return out
}

func SbomReferenceIntotoPayload_ToProto(mapCtx *direct.MapContext, in *krm.SbomReferenceIntotoPayload) *pb.SbomReferenceIntotoPayload {
	if in == nil {
		return nil
	}
	out := &pb.SbomReferenceIntotoPayload{}
	out.Type = direct.ValueOf(in.Type)
	out.PredicateType = direct.ValueOf(in.PredicateType)
	out.Subject = direct.Slice_ToProto(mapCtx, in.Subject, Subject_ToProto)
	out.Predicate = SbomReferenceIntotoPredicate_ToProto(mapCtx, in.Predicate)
	return out
}

func Source_FromProto(mapCtx *direct.MapContext, in *pb.Source) *krm.Source {
	if in == nil {
		return nil
	}
	out := &krm.Source{}
	out.ArtifactStorageSourceURI = direct.LazyPtr(in.GetArtifactStorageSourceUri())
	if in.FileHashes != nil {
		out.FileHashes = make(map[string]krm.FileHashes, len(in.FileHashes))
		for k, v := range in.FileHashes {
			if mapped := FileHashes_FromProto(mapCtx, v); mapped != nil {
				out.FileHashes[k] = *mapped
			}
		}
	}
	out.Context = SourceContext_FromProto(mapCtx, in.GetContext())
	out.AdditionalContexts = direct.Slice_FromProto(mapCtx, in.AdditionalContexts, SourceContext_FromProto)
	return out
}

func Source_ToProto(mapCtx *direct.MapContext, in *krm.Source) *pb.Source {
	if in == nil {
		return nil
	}
	out := &pb.Source{}
	out.ArtifactStorageSourceUri = direct.ValueOf(in.ArtifactStorageSourceURI)
	if in.FileHashes != nil {
		out.FileHashes = make(map[string]*pb.FileHashes, len(in.FileHashes))
		for k, v := range in.FileHashes {
			v := v
			out.FileHashes[k] = FileHashes_ToProto(mapCtx, &v)
		}
	}
	out.Context = SourceContext_ToProto(mapCtx, in.Context)
	out.AdditionalContexts = direct.Slice_ToProto(mapCtx, in.AdditionalContexts, SourceContext_ToProto)
	return out
}
