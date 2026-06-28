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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/privatecarefs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig) *krm.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &krm.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig{}
	if in.GetCaPool() != "" {
		out.CAPoolRef = &privatecarefs.PrivateCACAPoolRef{External: in.GetCaPool()}
	}
	return out
}

func CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig) *pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateIssuanceConfig_CertificateAuthorityConfig_CertificateAuthorityServiceConfig{}
	if in.CAPoolRef != nil {
		out.CaPool = privatecarefs.StripCAPoolPrefix(in.CAPoolRef.External)
	}
	return out
}
