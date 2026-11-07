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
	"strings"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Asset_ResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Asset_ResourceSpec) *krm.Asset_ResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.Asset_ResourceSpec{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ReadAccessMode = direct.Enum_FromProto(mapCtx, in.GetReadAccessMode())

	resName := in.GetName()
	if resName != "" {
		if strings.Contains(resName, "/buckets/") || in.GetType() == pb.Asset_ResourceSpec_STORAGE_BUCKET {
			out.StorageBucketRef = &storagev1beta1.StorageBucketRef{External: resName}
		} else if strings.Contains(resName, "/datasets/") || in.GetType() == pb.Asset_ResourceSpec_BIGQUERY_DATASET {
			out.BigQueryDatasetRef = &bigqueryv1beta1.DatasetRef{External: resName}
		}
	}
	return out
}

func Asset_ResourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.Asset_ResourceSpec) *pb.Asset_ResourceSpec {
	if in == nil {
		return nil
	}
	out := &pb.Asset_ResourceSpec{}
	out.Type = direct.Enum_ToProto[pb.Asset_ResourceSpec_Type](mapCtx, in.Type)
	out.ReadAccessMode = direct.Enum_ToProto[pb.Asset_ResourceSpec_AccessMode](mapCtx, in.ReadAccessMode)

	if in.StorageBucketRef != nil {
		out.Name = in.StorageBucketRef.External
	} else if in.BigQueryDatasetRef != nil {
		out.Name = in.BigQueryDatasetRef.External
	}
	return out
}
