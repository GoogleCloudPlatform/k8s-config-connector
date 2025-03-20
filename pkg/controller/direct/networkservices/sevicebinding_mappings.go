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

// +generated:mapper
// krm.group: networkservices.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networkservices.v1

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesServiceBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceBinding) *krm.NetworkServicesServiceBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesServiceBindingSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.Service != "" {
		out.Service = &v1alpha1.ServiceDirectoryServiceRef{
			External: in.Service,
		}
	}
	out.Labels = in.Labels
	return out
}
func NetworkServicesServiceBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesServiceBindingSpec) *pb.ServiceBinding {
	if in == nil {
		return nil
	}
	out := &pb.ServiceBinding{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	if in.Service != nil {
		out.Service = in.Service.External
	}
	out.Labels = in.Labels
	return out
}
