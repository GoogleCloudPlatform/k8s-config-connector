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

package designcenter

import (
	"fmt"
	"strings"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	krmstoragev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DesignCenterSpaceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Space) *krm.DesignCenterSpaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.DesignCenterSpaceSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.EnableGCPSharedTemplates = in.EnableGcpSharedTemplates
	if in.GetGcsBucket() != "" {
		bucket := in.GetGcsBucket()
		if !strings.Contains(bucket, "/") && in.GetName() != "" {
			if tokens := strings.Split(in.GetName(), "/"); len(tokens) >= 2 && tokens[0] == "projects" {
				bucket = fmt.Sprintf("projects/%s/buckets/%s", tokens[1], bucket)
			}
		}
		out.GCSBucketRef = &krmstoragev1beta1.StorageBucketRef{External: bucket}
	}
	out.Tags = in.Tags
	return out
}

func DesignCenterSpaceSpec_ToProto(mapCtx *direct.MapContext, in *krm.DesignCenterSpaceSpec) *pb.Space {
	if in == nil {
		return nil
	}
	out := &pb.Space{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.EnableGcpSharedTemplates = in.EnableGCPSharedTemplates
	if in.GCSBucketRef != nil {
		out.GcsBucket = bucketNameFromExternal(in.GCSBucketRef.External)
	}
	out.Tags = in.Tags
	return out
}

func bucketNameFromExternal(external string) string {
	if strings.Contains(external, "/buckets/") {
		parts := strings.Split(external, "/buckets/")
		return parts[len(parts)-1]
	}
	return strings.TrimPrefix(external, "gs://")
}
