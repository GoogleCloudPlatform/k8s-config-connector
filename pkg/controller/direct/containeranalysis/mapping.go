// Copyright 2024 Google LLC
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

package containeranalysis

import (
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containeranalysis/v1beta1"

	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func Note_ToProto(in *krm.ContainerAnalysisNote) *grafeaspb.Note {
	if in == nil {
		return nil
	}

	spec := in.Spec
	out := &grafeaspb.Note{
		ShortDescription: ValueOf(spec.ShortDescription),
		LongDescription:  ValueOf(spec.LongDescription),
		RelatedUrl:       RelatedUrl_ToProto(spec.RelatedUrl),
		ExpirationTime:   Time_ToProto(ValueOf(spec.ExpirationTime)),
		//RelatedNoteNames: ResourceNames_ToProto(spec.RelatedNoteNames),
	}

	if spec.Attestation != nil {
		out.Type = &grafeaspb.Note_Attestation{Attestation: Attestation_ToProto(spec.Attestation)}
	}
	if spec.Build != nil {
		out.Type = &grafeaspb.Note_Build{Build: Build_ToProto(spec.Build)}
	}
	if spec.Deployment != nil {
		out.Type = &grafeaspb.Note_Deployment{Deployment: Deployment_ToProto(spec.Deployment)}
	}
	if spec.Discovery != nil {
		out.Type = &grafeaspb.Note_Discovery{Discovery: Discovery_ToProto(spec.Discovery)}
	}
	if spec.Image != nil {
		out.Type = &grafeaspb.Note_Image{Image: Image_ToProto(spec.Image)}
	}
	if spec.Package != nil {
		out.Type = &grafeaspb.Note_Package{Package: Package_ToProto(spec.Package)}
	}
	if spec.Vulnerability != nil {
		out.Type = &grafeaspb.Note_Vulnerability{Vulnerability: Vulnerability_ToProto(spec.Vulnerability)}
	}

	return out
}

func RelatedUrl_ToProto(in []krm.NoteRelatedUrl) []*grafeaspb.RelatedUrl {
	out := make([]*grafeaspb.RelatedUrl, len(in))
	for i, v := range in {
		out[i] = &grafeaspb.RelatedUrl{
			Url:   ValueOf(v.Url),
			Label: ValueOf(v.Label),
		}
	}
	return out

}

//func ResourceNames_ToProto(in []v1alpha1.ResourceRef) []string {
//	if in == nil {
//		return nil
//	}
//	var out []string
//	for _, ref := range in {
//		out = append(out, ref.External)
//	}
//	return out
//}

func Attestation_ToProto(in *krm.NoteAttestation) *grafeaspb.AttestationNote {
	if in == nil {
		return nil
	}
	return &grafeaspb.AttestationNote{
		Hint: &grafeaspb.AttestationNote_Hint{
			HumanReadableName: in.Hint.HumanReadableName,
		},
	}
}

func Build_ToProto(in *krm.NoteBuild) *grafeaspb.BuildNote {
	if in == nil {
		return nil
	}
	return &grafeaspb.BuildNote{
		BuilderVersion: in.BuilderVersion,
	}
}

func Deployment_ToProto(in *krm.NoteDeployment) *grafeaspb.DeploymentNote {
	if in == nil {
		return nil
	}

	out := make([]string, len(in.ResourceUri))
	for i, v := range in.ResourceUri {
		out[i] = v
	}

	return &grafeaspb.DeploymentNote{
		ResourceUri: out,
	}
}

func Discovery_ToProto(in *krm.NoteDiscovery) *grafeaspb.DiscoveryNote {
	if in == nil {
		return nil
	}
	return &grafeaspb.DiscoveryNote{
		AnalysisKind: Enum_ToProto[grafeaspb.NoteKind](in.AnalysisKind),
	}
}

func Image_ToProto(in *krm.NoteImage) *grafeaspb.ImageNote {
	if in == nil {
		return nil
	}
	return &grafeaspb.ImageNote{
		ResourceUrl: in.ResourceUrl,
		Fingerprint: &grafeaspb.Fingerprint{V1Name: in.Fingerprint.V1Name, V2Blob: in.Fingerprint.V2Blob},
	}
}

func Package_ToProto(in *krm.NotePackage) *grafeaspb.PackageNote {
	if in == nil {
		return nil
	}
	return &grafeaspb.PackageNote{
		Name: in.Name,
		// Distribution is deprecated.
	}
}

func Vulnerability_ToProto(in *krm.NoteVulnerability) *grafeaspb.VulnerabilityNote {
	if in == nil {
		return nil
	}
	return &grafeaspb.VulnerabilityNote{
		CvssScore: float32(*in.CvssScore),
		Severity:  Enum_ToProto[grafeaspb.Severity](*in.Severity),
		Details:   Details_ToProto(in.Details),
	}
}

func Details_ToProto(in []krm.NoteDetails) []*grafeaspb.VulnerabilityNote_Detail {
	if in == nil {
		return nil
	}
	out := make([]*grafeaspb.VulnerabilityNote_Detail, len(in))
	for i, v := range in {
		out[i] = &grafeaspb.VulnerabilityNote_Detail{
			SeverityName:    ValueOf(v.SeverityName),
			Description:     ValueOf(v.Description),
			PackageType:     ValueOf(v.PackageType),
			AffectedCpeUri:  v.AffectedCpeUri,
			AffectedPackage: v.AffectedPackage,
			//AffectedVersionStart: v.AffectedVersionStart,
			//AffectedVersionEnd:   v.AffectedVersionEnd,
			FixedCpeUri:  ValueOf(v.FixedCpeUri),
			FixedPackage: ValueOf(v.FixedPackage),
			//FixedVersion:         v.FixedVersion,
			IsObsolete:       ValueOf(v.IsObsolete),
			SourceUpdateTime: Time_ToProto(ValueOf(v.SourceUpdateTime)),
			//Source: v.Source,
			//Vendor: v.Vendor,
		}
	}

	return out
}

func Time_ToProto(in string) *timestamppb.Timestamp {
	if in == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, in)
	if err != nil {
		return nil
	}
	out := timestamppb.New(t)
	return out
}

type ProtoEnum interface {
	~int32
	Descriptor() protoreflect.EnumDescriptor
}

func Enum_ToProto[U ProtoEnum](in string) U {
	var defaultU U
	descriptor := defaultU.Descriptor()

	inValue := in
	if inValue == "" {
		unspecifiedValue := U(0)
		return unspecifiedValue
	}

	n := descriptor.Values().Len()
	for i := 0; i < n; i++ {
		value := descriptor.Values().Get(i)
		if string(value.Name()) == inValue {
			v := U(value.Number())
			return v
		}
	}

	var validValues []string
	for i := 0; i < n; i++ {
		value := descriptor.Values().Get(i)
		validValues = append(validValues, string(value.Name()))
	}

	return 0
}
