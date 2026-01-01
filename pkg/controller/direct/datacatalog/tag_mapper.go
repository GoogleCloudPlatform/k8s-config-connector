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
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.datacatalog.v1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogTagSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Tag) *krm.DataCatalogTagSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogTagSpec{}
	// MISSING: Name
	if in.GetTemplate() != "" {
		out.TemplateRef = &krm.TagTemplateRef{External: in.GetTemplate()}
	}
	out.Column = direct.LazyPtr(in.GetColumn())
	out.Fields = make(map[string]krm.TagField)
	for k, v := range in.GetFields() {
		out.Fields[k] = *TagField_v1alpha1_FromProto(mapCtx, v)
	}
	return out
}

func DataCatalogTagSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogTagSpec) *pb.Tag {
	if in == nil {
		return nil
	}
	out := &pb.Tag{}
	// MISSING: Name
	if in.TemplateRef != nil {
		out.Template = in.TemplateRef.External
	}
	if oneof := DataCatalogTagSpec_Column_ToProto(mapCtx, in.Column); oneof != nil {
		out.Scope = oneof
	}
	out.Fields = make(map[string]*pb.TagField)
	for k, v := range in.Fields {
		out.Fields[k] = TagField_v1alpha1_ToProto(mapCtx, &v)
	}
	return out
}
