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

package certificatemanager

import (
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CertificateManagerDNSAuthorizationSpec_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization) *krm.CertificateManagerDNSAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerDNSAuthorizationSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Domain = in.GetDomain()
	return out
}
func CertificateManagerDNSAuthorizationSpec_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerDNSAuthorizationSpec) *pb.DnsAuthorization {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization{}
	out.Description = direct.ValueOf(in.Description)
	out.Domain = in.Domain
	return out
}
func DnsAuthorization_DnsResourceRecord_FromProto(mapCtx *direct.MapContext, in *pb.DnsAuthorization_DnsResourceRecord) *krm.DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := &krm.DnsAuthorization_DnsResourceRecord{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Data = direct.LazyPtr(in.GetData())
	return out
}
func DnsAuthorization_DnsResourceRecord_ToProto(mapCtx *direct.MapContext, in *krm.DnsAuthorization_DnsResourceRecord) *pb.DnsAuthorization_DnsResourceRecord {
	if in == nil {
		return nil
	}
	out := &pb.DnsAuthorization_DnsResourceRecord{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.Data = direct.ValueOf(in.Data)
	return out
}
