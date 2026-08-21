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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataSource_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.DataSource {
	if in == nil {
		return nil
	}
	out := &krm.DataSource{}
	if in.GetEntity() != "" {
		out.EntityRef = &krm.DataplexEntityRef{External: in.GetEntity()}
	}
	out.Resource = direct.LazyPtr(in.GetResource())
	return out
}

func DataSource_ToProto(mapCtx *direct.MapContext, in *krm.DataSource) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	if in.EntityRef != nil {
		out.Source = &pb.DataSource_Entity{
			Entity: in.EntityRef.External,
		}
	}
	if in.Resource != nil {
		out.Source = &pb.DataSource_Resource{
			Resource: *in.Resource,
		}
	}
	return out
}
