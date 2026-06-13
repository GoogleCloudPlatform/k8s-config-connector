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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// CertificateManagerCertificateMapEntrySpec_v1beta1_FromProto converts the CertificateMapEntry protobuf to the KRM Spec.
// Handcoded because of custom reference types (CertificatesRefs) and parent references (MapRef).
func CertificateManagerCertificateMapEntrySpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMapEntry) *krm.CertificateManagerCertificateMapEntrySpec {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerCertificateMapEntrySpec{}
	out.Description = direct.LazyPtr(in.GetDescription())

	// Map oneof match
	if in.Match != nil {
		switch m := in.Match.(type) {
		case *pb.CertificateMapEntry_Hostname:
			out.Hostname = direct.LazyPtr(m.Hostname)
		case *pb.CertificateMapEntry_Matcher_:
			if m.Matcher == pb.CertificateMapEntry_PRIMARY {
				out.Matcher = direct.LazyPtr("PRIMARY")
			} else {
				out.Matcher = direct.LazyPtr("MATCHER_UNSPECIFIED")
			}
		}
	}

	if len(in.Certificates) > 0 {
		out.CertificatesRefs = make([]krm.CertificateManagerCertificateRef, len(in.Certificates))
		for i, v := range in.Certificates {
			out.CertificatesRefs[i] = krm.CertificateManagerCertificateRef{External: v}
		}
	}

	return out
}

// CertificateManagerCertificateMapEntrySpec_v1beta1_ToProto converts the KRM Spec to the CertificateMapEntry protobuf.
// Handcoded because of custom reference types (CertificatesRefs) and parent references (MapRef).
func CertificateManagerCertificateMapEntrySpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerCertificateMapEntrySpec) *pb.CertificateMapEntry {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMapEntry{}
	out.Description = direct.ValueOf(in.Description)

	// Map oneof match
	if in.Hostname != nil {
		out.Match = &pb.CertificateMapEntry_Hostname{
			Hostname: *in.Hostname,
		}
	} else if in.Matcher != nil {
		var matcherVal pb.CertificateMapEntry_Matcher
		if *in.Matcher == "PRIMARY" {
			matcherVal = pb.CertificateMapEntry_PRIMARY
		} else {
			matcherVal = pb.CertificateMapEntry_MATCHER_UNSPECIFIED
		}
		out.Match = &pb.CertificateMapEntry_Matcher_{
			Matcher: matcherVal,
		}
	}

	if len(in.CertificatesRefs) > 0 {
		out.Certificates = make([]string, len(in.CertificatesRefs))
		for i, v := range in.CertificatesRefs {
			out.Certificates[i] = v.External
		}
	}

	return out
}

// CertificateManagerCertificateMapEntryStatus_v1beta1_FromProto converts the CertificateMapEntry protobuf to the KRM Status.
// Handcoded because of status-level field mappings.
func CertificateManagerCertificateMapEntryStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMapEntry) *krm.CertificateManagerCertificateMapEntryStatus {
	if in == nil {
		return nil
	}
	out := &krm.CertificateManagerCertificateMapEntryStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

// CertificateManagerCertificateMapEntryStatus_v1beta1_ToProto converts the KRM Status to the CertificateMapEntry protobuf.
// Handcoded because of status-level field mappings.
func CertificateManagerCertificateMapEntryStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.CertificateManagerCertificateMapEntryStatus) *pb.CertificateMapEntry {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMapEntry{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.ServingState](mapCtx, in.State)
	return out
}
