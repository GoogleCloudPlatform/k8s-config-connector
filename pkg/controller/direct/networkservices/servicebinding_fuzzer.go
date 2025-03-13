// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.networkservices.v1.ServiceBinding
// api.group: networkservices.cnrm.cloud.google.com

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(networkServicesServiceBindingFuzzer())
}

func NetworkServicesServiceBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.ServiceBinding) *krm.NetworkServicesServiceBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesServiceBindingSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetService() != "" {
		out.Service = direct.LazyPtr(in.GetService())
	}
	out.Labels = in.Labels
	return out
}
func NetworkServicesServiceBindingSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesServiceBindingSpec) *pb.ServiceBinding {
	if in == nil {
		return nil
	}
	out := &pb.ServiceBinding{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	if in.Service != nil {
		out.Service = direct.ValueOf(in.Service)
	}
	out.Labels = in.Labels
	return out
}

func NetworkServicesServiceBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceBinding) *krm.NetworkServicesServiceBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesServiceBindingObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func NetworkServicesServiceBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesServiceBindingObservedState) *pb.ServiceBinding {
	if in == nil {
		return nil
	}
	out := &pb.ServiceBinding{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func networkServicesServiceBindingFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServiceBinding{},
		NetworkServicesServiceBindingSpec_FromProto, NetworkServicesServiceBindingSpec_ToProto,
		NetworkServicesServiceBindingObservedState_FromProto, NetworkServicesServiceBindingObservedState_ToProto,
	)

	f.SpecFields.Insert(".name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".service")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	return f
}
