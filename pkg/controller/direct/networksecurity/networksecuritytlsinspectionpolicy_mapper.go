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

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	krmcertificatemanagerv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	krmnetworksecurityv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/privatecarefs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkSecurityTLSInspectionPolicySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.TlsInspectionPolicy) *krmnetworksecurityv1alpha1.NetworkSecurityTLSInspectionPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmnetworksecurityv1alpha1.NetworkSecurityTLSInspectionPolicySpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetCaPool() != "" {
		out.CaPoolRef = &privatecarefs.PrivateCACAPoolRef{External: in.GetCaPool()}
	}
	if in.GetTrustConfig() != "" {
		out.TrustConfigRef = &krmcertificatemanagerv1alpha1.CertificateManagerTrustConfigRef{External: in.GetTrustConfig()}
	}
	out.ExcludePublicCASet = in.ExcludePublicCaSet
	out.MinTLSVersion = direct.Enum_FromProto(mapCtx, in.GetMinTlsVersion())
	out.TLSFeatureProfile = direct.Enum_FromProto(mapCtx, in.GetTlsFeatureProfile())
	out.CustomTLSFeatures = in.CustomTlsFeatures
	return out
}
func NetworkSecurityTLSInspectionPolicySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworksecurityv1alpha1.NetworkSecurityTLSInspectionPolicySpec) *pb.TlsInspectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.TlsInspectionPolicy{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	if in.CaPoolRef != nil {
		out.CaPool = privatecarefs.StripCAPoolPrefix(in.CaPoolRef.External)
	}
	if in.TrustConfigRef != nil {
		out.TrustConfig = in.TrustConfigRef.External
	}
	out.ExcludePublicCaSet = in.ExcludePublicCASet
	out.MinTlsVersion = direct.Enum_ToProto[pb.TlsInspectionPolicy_TlsVersion](mapCtx, in.MinTLSVersion)
	out.TlsFeatureProfile = direct.Enum_ToProto[pb.TlsInspectionPolicy_Profile](mapCtx, in.TLSFeatureProfile)
	out.CustomTlsFeatures = in.CustomTLSFeatures
	return out
}
