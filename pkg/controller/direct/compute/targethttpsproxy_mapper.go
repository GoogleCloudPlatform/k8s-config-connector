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

package compute

import (
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeTargetHTTPSProxySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetHTTPSProxySpec) *pb.TargetHttpsProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetHttpsProxy{}
	if in.HttpKeepAliveTimeoutSec != nil {
		timeout := int32(*in.HttpKeepAliveTimeoutSec)
		out.HttpKeepAliveTimeoutSec = &timeout
	}
	out.ProxyBind = in.ProxyBind
	out.QuicOverride = in.QuicOverride
	out.Description = in.Description

	if in.CertificateMapRef != nil {
		out.CertificateMap = &in.CertificateMapRef.External
	}
	if in.UrlMapRef != nil {
		out.UrlMap = &in.UrlMapRef.External
	}
	if in.SslPolicyRef != nil {
		out.SslPolicy = &in.SslPolicyRef.External
	}
	if in.ServerTlsPolicyRef != nil {
		out.ServerTlsPolicy = &in.ServerTlsPolicyRef.External
	}

	for _, ref := range in.SslCertificates {
		out.SslCertificates = append(out.SslCertificates, ref.External)
	}

	for _, ref := range in.CertificateManagerCertificates {
		// CertificateManagerCertificates are stored in the same SslCertificates field in the proto
		// but they have a specific format.
		out.SslCertificates = append(out.SslCertificates, ref.External)
	}

	return out
}

func ComputeTargetHTTPSProxySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetHttpsProxy) *krm.ComputeTargetHTTPSProxySpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetHTTPSProxySpec{}
	if in.HttpKeepAliveTimeoutSec != nil {
		timeout := int(*in.HttpKeepAliveTimeoutSec)
		out.HttpKeepAliveTimeoutSec = &timeout
	}
	out.ProxyBind = in.ProxyBind
	out.QuicOverride = in.QuicOverride
	out.Description = in.Description

	if in.CertificateMap != nil {
		out.CertificateMapRef = &krm.CertificateManagerCertificateMapRef{External: *in.CertificateMap}
	}
	if in.UrlMap != nil {
		out.UrlMapRef = &krm.ComputeURLMapRef{External: *in.UrlMap}
	}
	if in.SslPolicy != nil {
		out.SslPolicyRef = &krm.ComputeSSLPolicyRef{External: *in.SslPolicy}
	}
	if in.ServerTlsPolicy != nil {
		out.ServerTlsPolicyRef = &krm.NetworkSecurityServerTLSPolicyRef{External: *in.ServerTlsPolicy}
	}

	for _, cert := range in.SslCertificates {
		// Distinguish between Classic SSL Certificates and Certificate Manager Certificates
		// based on the format.
		// Classic: projects/{{project}}/global/sslCertificates/{{name}}
		// CertManager: //certificatemanager.googleapis.com/projects/{{project}}/locations/{{location}}/certificates/{{name}}
		if cert != "" && strings.Contains(cert, "certificatemanager.googleapis.com") {
			out.CertificateManagerCertificates = append(out.CertificateManagerCertificates, krm.CertificateManagerCertificateRef{External: cert})
		} else {
			out.SslCertificates = append(out.SslCertificates, krm.ComputeSSLCertificateRef{External: cert})
		}
	}

	return out
}

func ComputeTargetHTTPSProxyStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetHttpsProxy) *krm.ComputeTargetHTTPSProxyStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetHTTPSProxyStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	if in.Id != nil {
		id := int(*in.Id)
		out.ProxyId = &id
	}
	out.SelfLink = in.SelfLink
	return out
}

func ComputeTargetHTTPSProxyObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetHttpsProxy) *krm.ComputeTargetHTTPSProxyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetHTTPSProxyObservedState{}
	out.Fingerprint = in.Fingerprint
	return out
}

func ComputeTargetHTTPSProxyObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetHTTPSProxyObservedState) *pb.TargetHttpsProxy {
	if in == nil {
		return nil
	}
	out := &pb.TargetHttpsProxy{}
	out.Fingerprint = in.Fingerprint
	return out
}
