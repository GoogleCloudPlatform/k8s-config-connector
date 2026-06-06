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

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func UptimeCheckConfig_ResourceGroup_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_ResourceGroup) *krm.UptimeCheckConfig_ResourceGroup {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_ResourceGroup{}
	if in.GetGroupId() != "" {
		out.GroupRef = &krm.MonitoringGroupRef{
			External: in.GetGroupId(),
		}
	}
	out.ResourceType = direct.Enum_FromProto(mapCtx, in.GetResourceType())
	return out
}

func UptimeCheckConfig_ResourceGroup_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_ResourceGroup) *pb.UptimeCheckConfig_ResourceGroup {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_ResourceGroup{}
	if in.GroupRef != nil {
		out.GroupId = in.GroupRef.External
	}
	out.ResourceType = direct.Enum_ToProto[pb.GroupResourceType](mapCtx, in.ResourceType)
	return out
}

func UptimeCheckConfig_HTTPCheck_BasicAuthentication_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_HttpCheck_BasicAuthentication) *krm.UptimeCheckConfig_HTTPCheck_BasicAuthentication {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_HTTPCheck_BasicAuthentication{}
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}

func UptimeCheckConfig_HTTPCheck_BasicAuthentication_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_HTTPCheck_BasicAuthentication) *pb.UptimeCheckConfig_HttpCheck_BasicAuthentication {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_HttpCheck_BasicAuthentication{}
	out.Username = direct.ValueOf(in.Username)
	if in.Password != nil {
		out.Password = direct.ValueOf(in.Password.Value)
	}
	return out
}

func UptimeCheckConfig_HTTPCheck_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_HttpCheck) *krm.UptimeCheckConfig_HTTPCheck {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_HTTPCheck{}
	out.AuthInfo = UptimeCheckConfig_HTTPCheck_BasicAuthentication_FromProto(mapCtx, in.GetAuthInfo())
	if len(in.GetBody()) > 0 {
		out.Body = direct.LazyPtr(string(in.GetBody()))
	}
	out.ContentType = direct.Enum_FromProto(mapCtx, in.GetContentType())
	out.Headers = in.GetHeaders()
	out.MaskHeaders = direct.LazyPtr(in.GetMaskHeaders())
	out.Path = direct.LazyPtr(in.GetPath())
	if in.GetPort() != 0 {
		out.Port = direct.LazyPtr(int64(in.GetPort()))
	}
	out.RequestMethod = direct.Enum_FromProto(mapCtx, in.GetRequestMethod())
	out.UseSsl = direct.LazyPtr(in.GetUseSsl())
	out.ValidateSsl = direct.LazyPtr(in.GetValidateSsl())
	return out
}

func UptimeCheckConfig_HTTPCheck_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_HTTPCheck) *pb.UptimeCheckConfig_HttpCheck {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_HttpCheck{}
	out.AuthInfo = UptimeCheckConfig_HTTPCheck_BasicAuthentication_ToProto(mapCtx, in.AuthInfo)
	if in.Body != nil {
		out.Body = []byte(*in.Body)
	}
	out.ContentType = direct.Enum_ToProto[pb.UptimeCheckConfig_HttpCheck_ContentType](mapCtx, in.ContentType)
	out.Headers = in.Headers
	out.MaskHeaders = direct.ValueOf(in.MaskHeaders)
	out.Path = direct.ValueOf(in.Path)
	out.Port = int32(direct.ValueOf(in.Port))
	out.RequestMethod = direct.Enum_ToProto[pb.UptimeCheckConfig_HttpCheck_RequestMethod](mapCtx, in.RequestMethod)
	out.UseSsl = direct.ValueOf(in.UseSsl)
	out.ValidateSsl = direct.ValueOf(in.ValidateSsl)
	return out
}

func UptimeCheckConfig_TCPCheck_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_TcpCheck) *krm.UptimeCheckConfig_TCPCheck {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_TCPCheck{}
	if in.GetPort() != 0 {
		out.Port = direct.LazyPtr(int64(in.GetPort()))
	}
	return out
}

func UptimeCheckConfig_TCPCheck_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_TCPCheck) *pb.UptimeCheckConfig_TcpCheck {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_TcpCheck{}
	out.Port = int32(direct.ValueOf(in.Port))
	return out
}

func UptimeCheckConfig_ContentMatcher_FromProto(mapCtx *direct.MapContext, in *pb.UptimeCheckConfig_ContentMatcher) *krm.UptimeCheckConfig_ContentMatcher {
	if in == nil {
		return nil
	}
	out := &krm.UptimeCheckConfig_ContentMatcher{}
	out.Content = direct.LazyPtr(in.GetContent())
	out.Matcher = direct.Enum_FromProto(mapCtx, in.GetMatcher())
	return out
}

func UptimeCheckConfig_ContentMatcher_ToProto(mapCtx *direct.MapContext, in *krm.UptimeCheckConfig_ContentMatcher) *pb.UptimeCheckConfig_ContentMatcher {
	if in == nil {
		return nil
	}
	out := &pb.UptimeCheckConfig_ContentMatcher{}
	out.Content = direct.ValueOf(in.Content)
	out.Matcher = direct.Enum_ToProto[pb.UptimeCheckConfig_ContentMatcher_ContentMatcherOption](mapCtx, in.Matcher)
	return out
}
