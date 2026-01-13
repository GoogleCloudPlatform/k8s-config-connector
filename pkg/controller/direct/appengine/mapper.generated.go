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

// +generated:mapper
// krm.group: appengine.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.appengine.v1

package appengine

import (
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppEngineDomainMappingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DomainMapping) *krm.AppEngineDomainMappingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AppEngineDomainMappingObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ID
	// MISSING: SSLSettings
	out.ResourceRecords = direct.Slice_FromProto(mapCtx, in.ResourceRecords, ResourceRecord_FromProto)
	return out
}
func AppEngineDomainMappingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AppEngineDomainMappingObservedState) *pb.DomainMapping {
	if in == nil {
		return nil
	}
	out := &pb.DomainMapping{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ID
	// MISSING: SSLSettings
	out.ResourceRecords = direct.Slice_ToProto(mapCtx, in.ResourceRecords, ResourceRecord_ToProto)
	return out
}
func AppEngineDomainMappingSpec_FromProto(mapCtx *direct.MapContext, in *pb.DomainMapping) *krm.AppEngineDomainMappingSpec {
	if in == nil {
		return nil
	}
	out := &krm.AppEngineDomainMappingSpec{}
	// MISSING: ID
	// MISSING: SSLSettings
	// (near miss): "SSLSettings" vs "SslSettings"
	return out
}
func AppEngineDomainMappingSpec_ToProto(mapCtx *direct.MapContext, in *krm.AppEngineDomainMappingSpec) *pb.DomainMapping {
	if in == nil {
		return nil
	}
	out := &pb.DomainMapping{}
	// MISSING: ID
	// MISSING: SSLSettings
	// (near miss): "SSLSettings" vs "SslSettings"
	return out
}
func ResourceRecord_FromProto(mapCtx *direct.MapContext, in *pb.ResourceRecord) *krm.ResourceRecord {
	if in == nil {
		return nil
	}
	out := &krm.ResourceRecord{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Rrdata = direct.LazyPtr(in.GetRrdata())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func ResourceRecord_ToProto(mapCtx *direct.MapContext, in *krm.ResourceRecord) *pb.ResourceRecord {
	if in == nil {
		return nil
	}
	out := &pb.ResourceRecord{}
	out.Name = direct.ValueOf(in.Name)
	out.Rrdata = direct.ValueOf(in.Rrdata)
	out.Type = direct.Enum_ToProto[pb.ResourceRecord_RecordType](mapCtx, in.Type)
	return out
}
func SSLSettings_FromProto(mapCtx *direct.MapContext, in *pb.SslSettings) *krm.SSLSettings {
	if in == nil {
		return nil
	}
	out := &krm.SSLSettings{}
	out.CertificateID = direct.LazyPtr(in.GetCertificateId())
	out.SSLManagementType = direct.Enum_FromProto(mapCtx, in.GetSslManagementType())
	out.PendingManagedCertificateID = direct.LazyPtr(in.GetPendingManagedCertificateId())
	return out
}
func SSLSettings_ToProto(mapCtx *direct.MapContext, in *krm.SSLSettings) *pb.SslSettings {
	if in == nil {
		return nil
	}
	out := &pb.SslSettings{}
	out.CertificateId = direct.ValueOf(in.CertificateID)
	out.SslManagementType = direct.Enum_ToProto[pb.SslSettings_SslManagementType](mapCtx, in.SSLManagementType)
	out.PendingManagedCertificateId = direct.ValueOf(in.PendingManagedCertificateID)
	return out
}
