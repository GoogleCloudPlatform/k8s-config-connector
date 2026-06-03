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

package binaryauthorization

import (
	pb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/binaryauthorization/v1beta1"
	containeranalysisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containeranalysis/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func UserOwnedGrafeasNote_FromProto(mapCtx *direct.MapContext, in *pb.UserOwnedGrafeasNote) *krmv1beta1.UserOwnedGrafeasNote {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.UserOwnedGrafeasNote{}
	if in.GetNoteReference() != "" {
		out.NoteRef = &containeranalysisv1beta1.ContainerAnalysisNoteRef{External: in.GetNoteReference()}
	}
	out.PublicKeys = direct.Slice_FromProto(mapCtx, in.PublicKeys, AttestorPublicKey_FromProto)
	return out
}

func UserOwnedGrafeasNote_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.UserOwnedGrafeasNote) *pb.UserOwnedGrafeasNote {
	if in == nil {
		return nil
	}
	out := &pb.UserOwnedGrafeasNote{}
	if in.NoteRef != nil {
		out.NoteReference = in.NoteRef.External
	}
	out.PublicKeys = direct.Slice_ToProto(mapCtx, in.PublicKeys, AttestorPublicKey_ToProto)
	return out
}
