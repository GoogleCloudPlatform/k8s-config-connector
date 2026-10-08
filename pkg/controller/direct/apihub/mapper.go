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

package apihub

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func APIHubInstance_Config_FromProto(mapCtx *direct.MapContext, in *pb.ApiHubInstance_Config) *krm.APIHubInstance_Config {
	if in == nil {
		return nil
	}
	out := &krm.APIHubInstance_Config{}

	if in.GetCmekKeyName() != "" {
		out.CmekKeyRef = &refs.KMSCryptoKeyRef{External: in.GetCmekKeyName()}
	}

	return out
}

func APIHubInstance_Config_ToProto(mapCtx *direct.MapContext, in *krm.APIHubInstance_Config) *pb.ApiHubInstance_Config {
	if in == nil {
		return nil
	}
	out := &pb.ApiHubInstance_Config{}

	if in.CmekKeyRef != nil {
		out.CmekKeyName = in.CmekKeyRef.External
	}

	return out
}

func APIHubAttributeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Attribute) *krm.APIHubAttributeSpec {
	if in == nil {
		return nil
	}
	out := &krm.APIHubAttributeSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Scope = direct.Enum_FromProto(mapCtx, in.GetScope())
	out.DataType = direct.Enum_FromProto(mapCtx, in.GetDataType())
	out.AllowedValues = direct.Slice_FromProto(mapCtx, in.AllowedValues, Attribute_AllowedValue_FromProto)

	cardinality := int32(in.GetCardinality())
	if cardinality == 0 {
		cardinality = 1
	}
	out.Cardinality = &cardinality
	return out
}

func APIHubAttributeSpec_ToProto(mapCtx *direct.MapContext, in *krm.APIHubAttributeSpec) *pb.Attribute {
	if in == nil {
		return nil
	}
	out := &pb.Attribute{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Scope = direct.Enum_ToProto[pb.Attribute_Scope](mapCtx, in.Scope)
	out.DataType = direct.Enum_ToProto[pb.Attribute_DataType](mapCtx, in.DataType)
	out.AllowedValues = direct.Slice_ToProto(mapCtx, in.AllowedValues, Attribute_AllowedValue_ToProto)
	if in.Cardinality != nil {
		out.Cardinality = *in.Cardinality
	}
	return out
}
