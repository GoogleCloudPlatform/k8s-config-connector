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

package cloudidentity

import (
	groupspb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/apps/cloudidentity/groups/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	cloudidentity "google.golang.org/api/cloudidentity/v1beta1"
)

/*
Convert between gcp

	type Group struct {
	    Description              string                `json:"description,omitempty"`
	    DisplayName              string                `json:"displayName,omitempty"`
	    GroupKey                 *EntityKey            `json:"groupKey,omitempty"`
	    Labels                   map[string]string     `json:"labels,omitempty"`
	    Parent                   string                `json:"parent,omitempty"`
	}

and proto

	type Group struct {
	    Description          *string               `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	    DisplayName          *string               `protobuf:"bytes,4,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	    GroupKey             *EntityKey            `protobuf:"bytes,6,opt,name=group_key,json=groupKey" json:"group_key,omitempty"`
	    Labels               map[string]string     `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	    Parent               *string               `protobuf:"bytes,9,opt,name=parent" json:"parent,omitempty"`
	}

Focus only on configurable fields and eliminate the output-only fields.

The direct type conversion isn't possible since the Go structs that have the same fields but different field types,
we need to do a manual field-by-field conversion here.
*/
func convertProtoToAPI(in *groupspb.Group) *cloudidentity.Group {
	if in == nil {
		return nil
	}

	out := &cloudidentity.Group{
		Labels: in.Labels,
	}

	if in.Description != nil {
		out.Description = direct.ValueOf(in.Description)
	}

	if in.DisplayName != nil {
		out.DisplayName = direct.ValueOf(in.DisplayName)
	}

	if in.GroupKey != nil {
		out.GroupKey = &cloudidentity.EntityKey{
			Id:        direct.ValueOf(in.GroupKey.Id),
			Namespace: direct.ValueOf(in.GroupKey.Namespace),
		}
	}

	if in.Parent != nil {
		out.Parent = direct.ValueOf(in.Parent)
	}

	if in.AdditionalGroupKeys != nil {
		for _, key := range in.AdditionalGroupKeys {
			keypb := &cloudidentity.EntityKey{
				Id:        direct.ValueOf(key.Id),
				Namespace: direct.ValueOf(key.Namespace),
			}
			out.AdditionalGroupKeys = append(out.AdditionalGroupKeys, keypb)
		}
	}
	return out
}

func convertAPIToProto(in *cloudidentity.Group) *groupspb.Group {
	if in == nil {
		return nil
	}

	out := &groupspb.Group{
		Labels: in.Labels,
	}

	if in.Description != "" {
		out.Description = direct.LazyPtr(in.Description)
	}
	if in.DisplayName != "" {
		out.DisplayName = direct.LazyPtr(in.DisplayName)
	}
	if in.GroupKey != nil {
		out.GroupKey = &groupspb.EntityKey{
			Id:        direct.LazyPtr(in.GroupKey.Id),
			Namespace: direct.LazyPtr(in.GroupKey.Namespace),
		}
	}
	if in.Parent != "" {
		out.Parent = direct.LazyPtr(in.Parent)
	}
	if in.AdditionalGroupKeys != nil {
		for _, key := range in.AdditionalGroupKeys {
			keypb := &groupspb.EntityKey{
				Id:        direct.LazyPtr(key.Id),
				Namespace: direct.LazyPtr(key.Namespace),
			}
			out.AdditionalGroupKeys = append(out.AdditionalGroupKeys, keypb)
		}
	}
	return out
}
