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

package dlp

import (
	"fmt"
	"time"

	pb "cloud.google.com/go/dlp/apiv2/dlppb"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dlp/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Status_FromProto(mapCtx *direct.MapContext, in *status.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	if in.Details != nil {
		out.Details = make([]krm.Any, len(in.Details))
		for i, d := range in.Details {
			out.Details[i] = krm.Any{
				TypeURL: direct.LazyPtr(d.GetTypeUrl()),
				Value:   d.GetValue(),
			}
		}
	}
	return out
}

func Status_ToProto(mapCtx *direct.MapContext, in *krm.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}

func string_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	if in == nil {
		return nil
	}
	s := in.AsTime().Format(time.RFC3339)
	return &s
}

func string_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	if in == nil || *in == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, *in)
	if err != nil {
		return nil
	}
	return timestamppb.New(t)
}

func InspectJobConfig_FromProto(mapCtx *direct.MapContext, in *pb.InspectJobConfig) *krm.InspectJobConfig {
	if in == nil {
		return nil
	}
	out := &krm.InspectJobConfig{}
	out.StorageConfig = StorageConfig_FromProto(mapCtx, in.GetStorageConfig())
	out.InspectConfig = InspectConfig_FromProto(mapCtx, in.GetInspectConfig())
	if in.GetInspectTemplateName() != "" {
		out.InspectTemplateRef = &krm.DLPInspectTemplateRef{External: in.GetInspectTemplateName()}
	}
	out.Actions = direct.Slice_FromProto(mapCtx, in.Actions, Action_FromProto)
	return out
}

func InspectJobConfig_ToProto(mapCtx *direct.MapContext, in *krm.InspectJobConfig) *pb.InspectJobConfig {
	if in == nil {
		return nil
	}
	out := &pb.InspectJobConfig{}
	out.StorageConfig = StorageConfig_ToProto(mapCtx, in.StorageConfig)
	out.InspectConfig = InspectConfig_ToProto(mapCtx, in.InspectConfig)
	if in.InspectTemplateRef != nil {
		out.InspectTemplateName = in.InspectTemplateRef.External
	}
	out.Actions = direct.Slice_ToProto(mapCtx, in.Actions, Action_ToProto)
	return out
}

func PartitionID_FromProto(mapCtx *direct.MapContext, in *pb.PartitionId) *krm.PartitionID {
	if in == nil {
		return nil
	}
	out := &krm.PartitionID{}
	if in.GetProjectId() != "" {
		out.ProjectRef = &refsv1beta1.ProjectRef{External: in.GetProjectId()}
	}
	out.NamespaceID = direct.LazyPtr(in.GetNamespaceId())
	return out
}

func PartitionID_ToProto(mapCtx *direct.MapContext, in *krm.PartitionID) *pb.PartitionId {
	if in == nil {
		return nil
	}
	out := &pb.PartitionId{}
	if in.ProjectRef != nil {
		out.ProjectId = in.ProjectRef.External
	}
	out.NamespaceId = direct.ValueOf(in.NamespaceID)
	return out
}

func CloudStorageRegexFileSet_FromProto(mapCtx *direct.MapContext, in *pb.CloudStorageRegexFileSet) *krm.CloudStorageRegexFileSet {
	if in == nil {
		return nil
	}
	out := &krm.CloudStorageRegexFileSet{}
	if in.GetBucketName() != "" {
		out.BucketRef = &storagev1beta1.StorageBucketRef{External: in.GetBucketName()}
	}
	out.IncludeRegex = in.IncludeRegex
	out.ExcludeRegex = in.ExcludeRegex
	return out
}

func CloudStorageRegexFileSet_ToProto(mapCtx *direct.MapContext, in *krm.CloudStorageRegexFileSet) *pb.CloudStorageRegexFileSet {
	if in == nil {
		return nil
	}
	out := &pb.CloudStorageRegexFileSet{}
	if in.BucketRef != nil {
		out.BucketName = in.BucketRef.External
	}
	out.IncludeRegex = in.IncludeRegex
	out.ExcludeRegex = in.ExcludeRegex
	return out
}

func Action_PublishToPubSub_FromProto(mapCtx *direct.MapContext, in *pb.Action_PublishToPubSub) *krm.Action_PublishToPubSub {
	if in == nil {
		return nil
	}
	out := &krm.Action_PublishToPubSub{}
	if in.GetTopic() != "" {
		out.TopicRef = &pubsubv1beta1.PubSubTopicRef{External: in.GetTopic()}
	}
	return out
}

func Action_PublishToPubSub_ToProto(mapCtx *direct.MapContext, in *krm.Action_PublishToPubSub) *pb.Action_PublishToPubSub {
	if in == nil {
		return nil
	}
	out := &pb.Action_PublishToPubSub{}
	if in.TopicRef != nil {
		out.Topic = in.TopicRef.External
	}
	return out
}

func BigQueryTable_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryTable) *krm.BigQueryTable {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryTable{}
	projectID := in.GetProjectId()
	datasetID := in.GetDatasetId()
	tableID := in.GetTableId()
	if projectID != "" {
		out.ProjectRef = &refsv1beta1.ProjectRef{External: projectID}
	}
	if datasetID != "" && projectID != "" {
		out.DatasetRef = &bigqueryv1beta1.DatasetRef{External: fmt.Sprintf("projects/%s/datasets/%s", projectID, datasetID)}
	}
	if tableID != "" && datasetID != "" && projectID != "" {
		out.TableRef = &bigqueryv1beta1.BigQueryTableRef{External: fmt.Sprintf("projects/%s/datasets/%s/tables/%s", projectID, datasetID, tableID)}
	}
	return out
}

func BigQueryTable_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryTable) *pb.BigQueryTable {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryTable{}
	if in.ProjectRef != nil {
		out.ProjectId = in.ProjectRef.External
	}
	if in.DatasetRef != nil {
		id := &bigqueryv1beta1.DatasetIdentity{}
		if err := id.FromExternal(in.DatasetRef.External); err == nil {
			out.DatasetId = id.Dataset
			if out.ProjectId == "" {
				out.ProjectId = id.Project
			}
		} else {
			out.DatasetId = in.DatasetRef.External
		}
	}
	if in.TableRef != nil {
		id := &bigqueryv1beta1.BigQueryTableIdentity{}
		if err := id.FromExternal(in.TableRef.External); err == nil {
			out.TableId = id.Table
			if out.DatasetId == "" {
				out.DatasetId = id.Dataset
			}
			if out.ProjectId == "" {
				out.ProjectId = id.Project
			}
		} else {
			out.TableId = in.TableRef.External
		}
	}
	return out
}

func StoredType_FromProto(mapCtx *direct.MapContext, in *pb.StoredType) *krm.StoredType {
	if in == nil {
		return nil
	}
	out := &krm.StoredType{}
	if in.GetName() != "" {
		out.NameRef = &krm.DLPStoredInfoTypeRef{External: in.GetName()}
	}
	out.CreateTime = string_FromProto(mapCtx, in.GetCreateTime())
	return out
}

func StoredType_ToProto(mapCtx *direct.MapContext, in *krm.StoredType) *pb.StoredType {
	if in == nil {
		return nil
	}
	out := &pb.StoredType{}
	if in.NameRef != nil {
		out.Name = in.NameRef.External
	}
	out.CreateTime = string_ToProto(mapCtx, in.CreateTime)
	return out
}

func TransformationConfig_FromProto(mapCtx *direct.MapContext, in *pb.TransformationConfig) *krm.TransformationConfig {
	if in == nil {
		return nil
	}
	out := &krm.TransformationConfig{}
	if in.GetDeidentifyTemplate() != "" {
		out.DeidentifyTemplateRef = &krm.DLPDeidentifyTemplateRef{External: in.GetDeidentifyTemplate()}
	}
	if in.GetStructuredDeidentifyTemplate() != "" {
		out.StructuredDeidentifyTemplateRef = &krm.DLPDeidentifyTemplateRef{External: in.GetStructuredDeidentifyTemplate()}
	}
	if in.GetImageRedactTemplate() != "" {
		out.ImageRedactTemplateRef = &krm.DLPDeidentifyTemplateRef{External: in.GetImageRedactTemplate()}
	}
	return out
}

func TransformationConfig_ToProto(mapCtx *direct.MapContext, in *krm.TransformationConfig) *pb.TransformationConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransformationConfig{}
	if in.DeidentifyTemplateRef != nil {
		out.DeidentifyTemplate = in.DeidentifyTemplateRef.External
	}
	if in.StructuredDeidentifyTemplateRef != nil {
		out.StructuredDeidentifyTemplate = in.StructuredDeidentifyTemplateRef.External
	}
	if in.ImageRedactTemplateRef != nil {
		out.ImageRedactTemplate = in.ImageRedactTemplateRef.External
	}
	return out
}
