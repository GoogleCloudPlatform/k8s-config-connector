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
// krm.version: v1alpha1
// proto.service: google.cloud.datacatalog.v1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FieldType_PrimitiveType_ToProto(mapCtx *direct.MapContext, in *string) *pb.FieldType_PrimitiveType_ {
	if in == nil {
		return nil
	}
	primitiveType := pb.FieldType_PrimitiveType(pb.FieldType_PrimitiveType_value[*in])
	return &pb.FieldType_PrimitiveType_{
		PrimitiveType: primitiveType,
	}
}
