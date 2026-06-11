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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// CertificateMap_GclbTarget_IPConfigObservedState_v1beta1_FromProto converts the IpConfig from Proto to KRM status.
// It is handcoded to map between uint32 ports array in the Protobuf and int32 array in KRM types.
func CertificateMap_GclbTarget_IPConfigObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CertificateMap_GclbTarget_IpConfig) *v1beta1.CertificateMap_GclbTarget_IPConfigObservedState {
	if in == nil {
		return nil
	}
	out := &v1beta1.CertificateMap_GclbTarget_IPConfigObservedState{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	if in.Ports != nil {
		out.Ports = make([]int32, len(in.Ports))
		for i, p := range in.Ports {
			out.Ports[i] = int32(p)
		}
	}
	return out
}

// CertificateMap_GclbTarget_IPConfigObservedState_v1beta1_ToProto converts the IpConfig from KRM status to Proto.
// It is handcoded to map between int32 array in KRM types and uint32 ports array in the Protobuf.
func CertificateMap_GclbTarget_IPConfigObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *v1beta1.CertificateMap_GclbTarget_IPConfigObservedState) *pb.CertificateMap_GclbTarget_IpConfig {
	if in == nil {
		return nil
	}
	out := &pb.CertificateMap_GclbTarget_IpConfig{}
	out.IpAddress = direct.ValueOf(in.IPAddress)
	if in.Ports != nil {
		out.Ports = make([]uint32, len(in.Ports))
		for i, p := range in.Ports {
			out.Ports[i] = uint32(p)
		}
	}
	return out
}
