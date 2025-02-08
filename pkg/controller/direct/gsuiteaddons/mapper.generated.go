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

package gsuiteaddons

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/gsuiteaddons/apiv1/gsuiteaddonspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gsuiteaddons/v1alpha1"
)
func AddOns_FromProto(mapCtx *direct.MapContext, in *pb.AddOns) *krm.AddOns {
	if in == nil {
		return nil
	}
	out := &krm.AddOns{}
	out.Common = CommonAddOnManifest_FromProto(mapCtx, in.GetCommon())
	out.Gmail = GmailAddOnManifest_FromProto(mapCtx, in.GetGmail())
	out.Drive = DriveAddOnManifest_FromProto(mapCtx, in.GetDrive())
	out.Calendar = CalendarAddOnManifest_FromProto(mapCtx, in.GetCalendar())
	out.Docs = DocsAddOnManifest_FromProto(mapCtx, in.GetDocs())
	out.Sheets = SheetsAddOnManifest_FromProto(mapCtx, in.GetSheets())
	out.Slides = SlidesAddOnManifest_FromProto(mapCtx, in.GetSlides())
	out.HTTPOptions = HttpOptions_FromProto(mapCtx, in.GetHttpOptions())
	return out
}
func AddOns_ToProto(mapCtx *direct.MapContext, in *krm.AddOns) *pb.AddOns {
	if in == nil {
		return nil
	}
	out := &pb.AddOns{}
	out.Common = CommonAddOnManifest_ToProto(mapCtx, in.Common)
	out.Gmail = GmailAddOnManifest_ToProto(mapCtx, in.Gmail)
	out.Drive = DriveAddOnManifest_ToProto(mapCtx, in.Drive)
	out.Calendar = CalendarAddOnManifest_ToProto(mapCtx, in.Calendar)
	out.Docs = DocsAddOnManifest_ToProto(mapCtx, in.Docs)
	out.Sheets = SheetsAddOnManifest_ToProto(mapCtx, in.Sheets)
	out.Slides = SlidesAddOnManifest_ToProto(mapCtx, in.Slides)
	out.HttpOptions = HttpOptions_ToProto(mapCtx, in.HTTPOptions)
	return out
}
func Deployment_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.Deployment {
	if in == nil {
		return nil
	}
	out := &krm.Deployment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.OauthScopes = in.OauthScopes
	out.AddOns = AddOns_FromProto(mapCtx, in.GetAddOns())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Deployment_ToProto(mapCtx *direct.MapContext, in *krm.Deployment) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	out.Name = direct.ValueOf(in.Name)
	out.OauthScopes = in.OauthScopes
	out.AddOns = AddOns_ToProto(mapCtx, in.AddOns)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func GsuiteaddonsDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.GsuiteaddonsDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GsuiteaddonsDeploymentObservedState{}
	// MISSING: Name
	// MISSING: OauthScopes
	// MISSING: AddOns
	// MISSING: Etag
	return out
}
func GsuiteaddonsDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GsuiteaddonsDeploymentObservedState) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	// MISSING: OauthScopes
	// MISSING: AddOns
	// MISSING: Etag
	return out
}
func GsuiteaddonsDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.GsuiteaddonsDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.GsuiteaddonsDeploymentSpec{}
	// MISSING: Name
	// MISSING: OauthScopes
	// MISSING: AddOns
	// MISSING: Etag
	return out
}
func GsuiteaddonsDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.GsuiteaddonsDeploymentSpec) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	// MISSING: Name
	// MISSING: OauthScopes
	// MISSING: AddOns
	// MISSING: Etag
	return out
}
