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

// +tool:controller
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.message: google.cloud.datacatalog.v1.TagTemplate
// crd.type: DataCatalogTagTemplate
// crd.version: v1alpha1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogTagTemplateSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.TagTemplate) *krm.DataCatalogTagTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogTagTemplateSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.IsPubliclyReadable = direct.LazyPtr(in.GetIsPubliclyReadable())
	out.Fields = make(map[string]krm.TagTemplateField)
	for k, v := range in.GetFields() {
		out.Fields[k] = *TagTemplateField_v1alpha1_FromProto(mapCtx, v)
	}
	return out
}
func DataCatalogTagTemplateSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogTagTemplateSpec) *pb.TagTemplate {
	if in == nil {
		return nil
	}
	out := &pb.TagTemplate{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.IsPubliclyReadable = direct.ValueOf(in.IsPubliclyReadable)

	out.Fields = make(map[string]*pb.TagTemplateField)
	for k, v := range in.Fields {
		out.Fields[k] = TagTemplateField_v1alpha1_ToProto(mapCtx, &v)
	}
	return out
}
