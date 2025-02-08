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

package websecurityscanner

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/websecurityscanner/apiv1beta/websecurityscannerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/websecurityscanner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Finding_FromProto(mapCtx *direct.MapContext, in *pb.Finding) *krm.Finding {
	if in == nil {
		return nil
	}
	out := &krm.Finding{}
	out.Name = direct.LazyPtr(in.GetName())
	out.FindingType = direct.LazyPtr(in.GetFindingType())
	out.HTTPMethod = direct.LazyPtr(in.GetHttpMethod())
	out.FuzzedURL = direct.LazyPtr(in.GetFuzzedUrl())
	out.Body = direct.LazyPtr(in.GetBody())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ReproductionURL = direct.LazyPtr(in.GetReproductionUrl())
	out.FrameURL = direct.LazyPtr(in.GetFrameUrl())
	out.FinalURL = direct.LazyPtr(in.GetFinalUrl())
	out.TrackingID = direct.LazyPtr(in.GetTrackingId())
	out.Form = Form_FromProto(mapCtx, in.GetForm())
	out.OutdatedLibrary = OutdatedLibrary_FromProto(mapCtx, in.GetOutdatedLibrary())
	out.ViolatingResource = ViolatingResource_FromProto(mapCtx, in.GetViolatingResource())
	out.VulnerableHeaders = VulnerableHeaders_FromProto(mapCtx, in.GetVulnerableHeaders())
	out.VulnerableParameters = VulnerableParameters_FromProto(mapCtx, in.GetVulnerableParameters())
	out.Xss = Xss_FromProto(mapCtx, in.GetXss())
	return out
}
func Finding_ToProto(mapCtx *direct.MapContext, in *krm.Finding) *pb.Finding {
	if in == nil {
		return nil
	}
	out := &pb.Finding{}
	out.Name = direct.ValueOf(in.Name)
	out.FindingType = direct.ValueOf(in.FindingType)
	out.HttpMethod = direct.ValueOf(in.HTTPMethod)
	out.FuzzedUrl = direct.ValueOf(in.FuzzedURL)
	out.Body = direct.ValueOf(in.Body)
	out.Description = direct.ValueOf(in.Description)
	out.ReproductionUrl = direct.ValueOf(in.ReproductionURL)
	out.FrameUrl = direct.ValueOf(in.FrameURL)
	out.FinalUrl = direct.ValueOf(in.FinalURL)
	out.TrackingId = direct.ValueOf(in.TrackingID)
	out.Form = Form_ToProto(mapCtx, in.Form)
	out.OutdatedLibrary = OutdatedLibrary_ToProto(mapCtx, in.OutdatedLibrary)
	out.ViolatingResource = ViolatingResource_ToProto(mapCtx, in.ViolatingResource)
	out.VulnerableHeaders = VulnerableHeaders_ToProto(mapCtx, in.VulnerableHeaders)
	out.VulnerableParameters = VulnerableParameters_ToProto(mapCtx, in.VulnerableParameters)
	out.Xss = Xss_ToProto(mapCtx, in.Xss)
	return out
}
func Form_FromProto(mapCtx *direct.MapContext, in *pb.Form) *krm.Form {
	if in == nil {
		return nil
	}
	out := &krm.Form{}
	out.ActionURI = direct.LazyPtr(in.GetActionUri())
	out.Fields = in.Fields
	return out
}
func Form_ToProto(mapCtx *direct.MapContext, in *krm.Form) *pb.Form {
	if in == nil {
		return nil
	}
	out := &pb.Form{}
	out.ActionUri = direct.ValueOf(in.ActionURI)
	out.Fields = in.Fields
	return out
}
func OutdatedLibrary_FromProto(mapCtx *direct.MapContext, in *pb.OutdatedLibrary) *krm.OutdatedLibrary {
	if in == nil {
		return nil
	}
	out := &krm.OutdatedLibrary{}
	out.LibraryName = direct.LazyPtr(in.GetLibraryName())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.LearnMoreUrls = in.LearnMoreUrls
	return out
}
func OutdatedLibrary_ToProto(mapCtx *direct.MapContext, in *krm.OutdatedLibrary) *pb.OutdatedLibrary {
	if in == nil {
		return nil
	}
	out := &pb.OutdatedLibrary{}
	out.LibraryName = direct.ValueOf(in.LibraryName)
	out.Version = direct.ValueOf(in.Version)
	out.LearnMoreUrls = in.LearnMoreUrls
	return out
}
func ViolatingResource_FromProto(mapCtx *direct.MapContext, in *pb.ViolatingResource) *krm.ViolatingResource {
	if in == nil {
		return nil
	}
	out := &krm.ViolatingResource{}
	out.ContentType = direct.LazyPtr(in.GetContentType())
	out.ResourceURL = direct.LazyPtr(in.GetResourceUrl())
	return out
}
func ViolatingResource_ToProto(mapCtx *direct.MapContext, in *krm.ViolatingResource) *pb.ViolatingResource {
	if in == nil {
		return nil
	}
	out := &pb.ViolatingResource{}
	out.ContentType = direct.ValueOf(in.ContentType)
	out.ResourceUrl = direct.ValueOf(in.ResourceURL)
	return out
}
func VulnerableHeaders_FromProto(mapCtx *direct.MapContext, in *pb.VulnerableHeaders) *krm.VulnerableHeaders {
	if in == nil {
		return nil
	}
	out := &krm.VulnerableHeaders{}
	out.Headers = direct.Slice_FromProto(mapCtx, in.Headers, VulnerableHeaders_Header_FromProto)
	out.MissingHeaders = direct.Slice_FromProto(mapCtx, in.MissingHeaders, VulnerableHeaders_Header_FromProto)
	return out
}
func VulnerableHeaders_ToProto(mapCtx *direct.MapContext, in *krm.VulnerableHeaders) *pb.VulnerableHeaders {
	if in == nil {
		return nil
	}
	out := &pb.VulnerableHeaders{}
	out.Headers = direct.Slice_ToProto(mapCtx, in.Headers, VulnerableHeaders_Header_ToProto)
	out.MissingHeaders = direct.Slice_ToProto(mapCtx, in.MissingHeaders, VulnerableHeaders_Header_ToProto)
	return out
}
func VulnerableHeaders_Header_FromProto(mapCtx *direct.MapContext, in *pb.VulnerableHeaders_Header) *krm.VulnerableHeaders_Header {
	if in == nil {
		return nil
	}
	out := &krm.VulnerableHeaders_Header{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func VulnerableHeaders_Header_ToProto(mapCtx *direct.MapContext, in *krm.VulnerableHeaders_Header) *pb.VulnerableHeaders_Header {
	if in == nil {
		return nil
	}
	out := &pb.VulnerableHeaders_Header{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func VulnerableParameters_FromProto(mapCtx *direct.MapContext, in *pb.VulnerableParameters) *krm.VulnerableParameters {
	if in == nil {
		return nil
	}
	out := &krm.VulnerableParameters{}
	out.ParameterNames = in.ParameterNames
	return out
}
func VulnerableParameters_ToProto(mapCtx *direct.MapContext, in *krm.VulnerableParameters) *pb.VulnerableParameters {
	if in == nil {
		return nil
	}
	out := &pb.VulnerableParameters{}
	out.ParameterNames = in.ParameterNames
	return out
}
func Xss_FromProto(mapCtx *direct.MapContext, in *pb.Xss) *krm.Xss {
	if in == nil {
		return nil
	}
	out := &krm.Xss{}
	out.StackTraces = in.StackTraces
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	return out
}
func Xss_ToProto(mapCtx *direct.MapContext, in *krm.Xss) *pb.Xss {
	if in == nil {
		return nil
	}
	out := &pb.Xss{}
	out.StackTraces = in.StackTraces
	out.ErrorMessage = direct.ValueOf(in.ErrorMessage)
	return out
}
