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

package networksecuritybackendauthenticationconfig

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networksecurity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkSecurityBackendAuthenticationConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackendAuthenticationConfig) *krm.NetworkSecurityBackendAuthenticationConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityBackendAuthenticationConfigObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}

func NetworkSecurityBackendAuthenticationConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityBackendAuthenticationConfigObservedState) *pb.BackendAuthenticationConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackendAuthenticationConfig{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}

func NetworkSecurityBackendAuthenticationConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackendAuthenticationConfig) *krm.NetworkSecurityBackendAuthenticationConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSecurityBackendAuthenticationConfigSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	if in.GetClientCertificate() != "" {
		out.ClientCertificateRef = &refsv1beta1.CertificateManagerCertificateRef{External: in.GetClientCertificate()}
	}
	if in.GetTrustConfig() != "" {
		out.TrustConfigRef = &refsv1beta1.CertificateManagerTrustConfigRef{External: in.GetTrustConfig()}
	}
	out.WellKnownRoots = direct.Enum_FromProto(mapCtx, in.GetWellKnownRoots())
	return out
}

func NetworkSecurityBackendAuthenticationConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSecurityBackendAuthenticationConfigSpec) *pb.BackendAuthenticationConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackendAuthenticationConfig{}
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	if in.ClientCertificateRef != nil {
		out.ClientCertificate = in.ClientCertificateRef.External
	}
	if in.TrustConfigRef != nil {
		out.TrustConfig = in.TrustConfigRef.External
	}
	out.WellKnownRoots = direct.Enum_ToProto[pb.BackendAuthenticationConfig_WellKnownRoots](mapCtx, in.WellKnownRoots)
	return out
}
