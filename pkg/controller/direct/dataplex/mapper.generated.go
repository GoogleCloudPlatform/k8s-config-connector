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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Content_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krm.Content {
	if in == nil {
		return nil
	}
	out := &krm.Content{}
	// MISSING: Name
	// MISSING: Uid
	out.Path = direct.LazyPtr(in.GetPath())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DataText = direct.LazyPtr(in.GetDataText())
	out.SQLScript = Content_SQLScript_FromProto(mapCtx, in.GetSqlScript())
	out.Notebook = Content_Notebook_FromProto(mapCtx, in.GetNotebook())
	return out
}
func Content_ToProto(mapCtx *direct.MapContext, in *krm.Content) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	// MISSING: Name
	// MISSING: Uid
	out.Path = direct.ValueOf(in.Path)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	if oneof := Content_DataText_ToProto(mapCtx, in.DataText); oneof != nil {
		out.Data = oneof
	}
	if oneof := Content_SQLScript_ToProto(mapCtx, in.SQLScript); oneof != nil {
		out.Content = &pb.Content_SqlScript_{SqlScript: oneof}
	}
	if oneof := Content_Notebook_ToProto(mapCtx, in.Notebook); oneof != nil {
		out.Content = &pb.Content_Notebook_{Notebook: oneof}
	}
	return out
}
func ContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krm.ContentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ContentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Path
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
func ContentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ContentObservedState) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Path
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
func Content_Notebook_FromProto(mapCtx *direct.MapContext, in *pb.Content_Notebook) *krm.Content_Notebook {
	if in == nil {
		return nil
	}
	out := &krm.Content_Notebook{}
	out.KernelType = direct.Enum_FromProto(mapCtx, in.GetKernelType())
	return out
}
func Content_Notebook_ToProto(mapCtx *direct.MapContext, in *krm.Content_Notebook) *pb.Content_Notebook {
	if in == nil {
		return nil
	}
	out := &pb.Content_Notebook{}
	out.KernelType = direct.Enum_ToProto[pb.Content_Notebook_KernelType](mapCtx, in.KernelType)
	return out
}
func Content_SQLScript_FromProto(mapCtx *direct.MapContext, in *pb.Content_SqlScript) *krm.Content_SQLScript {
	if in == nil {
		return nil
	}
	out := &krm.Content_SQLScript{}
	out.Engine = direct.Enum_FromProto(mapCtx, in.GetEngine())
	return out
}
func Content_SQLScript_ToProto(mapCtx *direct.MapContext, in *krm.Content_SQLScript) *pb.Content_SqlScript {
	if in == nil {
		return nil
	}
	out := &pb.Content_SqlScript{}
	out.Engine = direct.Enum_ToProto[pb.Content_SqlScript_QueryEngine](mapCtx, in.Engine)
	return out
}
func DataplexContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krm.DataplexContentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexContentObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Path
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
func DataplexContentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexContentObservedState) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Path
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
func DataplexContentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krm.DataplexContentSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexContentSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Path
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
func DataplexContentSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexContentSpec) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Path
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
