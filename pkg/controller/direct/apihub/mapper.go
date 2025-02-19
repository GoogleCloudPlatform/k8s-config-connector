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

package apihub

import (
	"cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Attribute_AllowedValue_FromProto(mapCtx *direct.MapContext, in *apihubpb.Attribute_AllowedValue) *krm.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	out := &krm.Attribute_AllowedValue{}
	out.ID = &in.Id
	out.DisplayName = &in.DisplayName
	out.Description = &in.Description
	out.Immutable = &in.Immutable
	return out
}

func Attribute_AllowedValue_ToProto(mapCtx *direct.MapContext, in *krm.Attribute_AllowedValue) *apihubpb.Attribute_AllowedValue {
	if in == nil {
		return nil
	}
	id := ""
	if in.ID != nil {
		id = *in.ID
	}
	displayName := ""
	if in.DisplayName != nil {
		displayName = *in.DisplayName
	}
	description := ""
	if in.Description != nil {
		description = *in.Description
	}
	immutable := false
	if in.Immutable != nil {
		immutable = *in.Immutable
	}

	out := &apihubpb.Attribute_AllowedValue{
		Id:          id,
		DisplayName: displayName,
		Description: description,
		Immutable:   immutable,
	}
	return out
}
