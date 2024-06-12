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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containeranalysis/v1beta1"

	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func Note_KRMToApi(in *krm.ContainerAnalysisNote) *grafeaspb.Note {
	if in == nil {
		return nil
	}

	out := &grafeaspb.Note{
		ShortDescription: ValueOf(in.Spec.ShortDescription),
		LongDescription:  ValueOf(in.Spec.LongDescription),
		RelatedUrl:       RelatedUrl_KRMToApi(in.Spec.RelatedUrl),
		ExpirationTime:   Time_KRMToApi(ValueOf(in.Spec.ExpirationTime)),
		//RelatedNoteNames: ResourceNames_KRMToApi(in.Spec.RelatedNoteNames),
		Type: &grafeaspb.Note_Attestation{Attestation: Attestation_KRMToApi(in.Spec.Attestation)},
	}
	return out
}

func RelatedUrl_KRMToApi(in []krm.NoteRelatedUrl) []*grafeaspb.RelatedUrl {
	out := make([]*grafeaspb.RelatedUrl, len(in))
	for i, v := range in {
		out[i] = &grafeaspb.RelatedUrl{
			Url:   ValueOf(v.Url),
			Label: ValueOf(v.Label),
		}
	}
	return out

}

func Time_KRMToApi(in string) *timestamppb.Timestamp {
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

//func ResourceNames_KRMToApi(in []v1alpha1.ResourceRef) []string {
//	if in == nil {
//		return nil
//	}
//	var out []string
//	for _, ref := range in {
//		out = append(out, ref.External)
//	}
//	return out
//}

func Attestation_KRMToApi(in *krm.NoteAttestation) *grafeaspb.AttestationNote {
	if in == nil {
		return nil
	}
	return &grafeaspb.AttestationNote{
		Hint: &grafeaspb.AttestationNote_Hint{
			HumanReadableName: in.Hint.HumanReadableName,
		},
	}
}
