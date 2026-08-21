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

package certificatemanager

import (
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CertificateSelfManaged_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_SelfManagedCertificate) *krm.CertificateSelfManaged {
	if in == nil {
		return nil
	}
	out := &krm.CertificateSelfManaged{}
	out.PemCertificate = direct.LazyPtr(in.GetPemCertificate())
	if in.PemCertificate != "" {
		out.CertificatePem = &krm.CertificateCertificatePem{Value: direct.LazyPtr(in.GetPemCertificate())}
	}
	return out
}

func CertificateSelfManaged_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.CertificateSelfManaged) *pb.Certificate_SelfManagedCertificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_SelfManagedCertificate{}
	if in.PemCertificate != nil {
		out.PemCertificate = *in.PemCertificate
	} else if in.CertificatePem != nil && in.CertificatePem.Value != nil {
		out.PemCertificate = *in.CertificatePem.Value
	}
	if in.PemPrivateKey != nil && in.PemPrivateKey.Value != nil {
		out.PemPrivateKey = *in.PemPrivateKey.Value
	} else if in.PrivateKeyPem != nil && in.PrivateKeyPem.Value != nil {
		out.PemPrivateKey = *in.PrivateKeyPem.Value
	}
	return out
}

func CertificateManaged_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_ManagedCertificate) *krm.CertificateManaged {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManaged{}
	out.Domains = in.Domains
	if len(in.DnsAuthorizations) > 0 {
		out.DnsAuthorizationsRefs = make([]krm.CertificateManagerDNSAuthorizationRef, len(in.DnsAuthorizations))
		for i, v := range in.DnsAuthorizations {
			out.DnsAuthorizationsRefs[i] = krm.CertificateManagerDNSAuthorizationRef{External: v}
		}
	}
	if in.IssuanceConfig != "" {
		out.IssuanceConfigRef = &krmv1alpha1.CertificateManagerCertificateIssuanceConfigRef{External: in.IssuanceConfig}
	}
	return out
}

func CertificateManaged_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManaged) *pb.Certificate_ManagedCertificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_ManagedCertificate{}
	out.Domains = in.Domains
	if len(in.DnsAuthorizationsRefs) > 0 {
		out.DnsAuthorizations = make([]string, len(in.DnsAuthorizationsRefs))
		for i, v := range in.DnsAuthorizationsRefs {
			out.DnsAuthorizations[i] = v.External
		}
	}
	if in.IssuanceConfigRef != nil {
		out.IssuanceConfig = in.IssuanceConfigRef.External
	}
	return out
}

func CertificateManagedStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Certificate_ManagedCertificate) *krm.CertificateManagedStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagedStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	if v := in.GetProvisioningIssue(); v != nil {
		if issue := CertificateProvisioningIssueStatus_v1beta1_FromProto(mapCtx, v); issue != nil {
			out.ProvisioningIssue = []krm.CertificateProvisioningIssueStatus{*issue}
		}
	}
	out.AuthorizationAttemptInfo = direct.Slice_FromProto(mapCtx, in.AuthorizationAttemptInfo, CertificateAuthorizationAttemptInfoStatus_v1beta1_FromProto)
	return out
}

func CertificateManagedStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagedStatus) *pb.Certificate_ManagedCertificate {
	if in == nil {
		return nil
	}
	out := &pb.Certificate_ManagedCertificate{}
	out.State = direct.Enum_ToProto[pb.Certificate_ManagedCertificate_State](mapCtx, in.State)
	if len(in.ProvisioningIssue) > 0 {
		out.ProvisioningIssue = CertificateProvisioningIssueStatus_v1beta1_ToProto(mapCtx, &in.ProvisioningIssue[0])
	}
	out.AuthorizationAttemptInfo = direct.Slice_ToProto(mapCtx, in.AuthorizationAttemptInfo, CertificateAuthorizationAttemptInfoStatus_v1beta1_ToProto)
	return out
}
