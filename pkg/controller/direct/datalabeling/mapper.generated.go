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

package datalabeling

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CsvInstruction_FromProto(mapCtx *direct.MapContext, in *pb.CsvInstruction) *krm.CsvInstruction {
	if in == nil {
		return nil
	}
	out := &krm.CsvInstruction{}
	out.GcsFileURI = direct.LazyPtr(in.GetGcsFileUri())
	return out
}
func CsvInstruction_ToProto(mapCtx *direct.MapContext, in *krm.CsvInstruction) *pb.CsvInstruction {
	if in == nil {
		return nil
	}
	out := &pb.CsvInstruction{}
	out.GcsFileUri = direct.ValueOf(in.GcsFileURI)
	return out
}
func DatalabelingInstructionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instruction) *krm.DatalabelingInstructionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingInstructionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataType
	// MISSING: CsvInstruction
	// MISSING: PdfInstruction
	// MISSING: BlockingResources
	return out
}
func DatalabelingInstructionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingInstructionObservedState) *pb.Instruction {
	if in == nil {
		return nil
	}
	out := &pb.Instruction{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataType
	// MISSING: CsvInstruction
	// MISSING: PdfInstruction
	// MISSING: BlockingResources
	return out
}
func DatalabelingInstructionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instruction) *krm.DatalabelingInstructionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatalabelingInstructionSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataType
	// MISSING: CsvInstruction
	// MISSING: PdfInstruction
	// MISSING: BlockingResources
	return out
}
func DatalabelingInstructionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatalabelingInstructionSpec) *pb.Instruction {
	if in == nil {
		return nil
	}
	out := &pb.Instruction{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DataType
	// MISSING: CsvInstruction
	// MISSING: PdfInstruction
	// MISSING: BlockingResources
	return out
}
func Instruction_FromProto(mapCtx *direct.MapContext, in *pb.Instruction) *krm.Instruction {
	if in == nil {
		return nil
	}
	out := &krm.Instruction{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DataType = direct.Enum_FromProto(mapCtx, in.GetDataType())
	out.CsvInstruction = CsvInstruction_FromProto(mapCtx, in.GetCsvInstruction())
	out.PdfInstruction = PdfInstruction_FromProto(mapCtx, in.GetPdfInstruction())
	out.BlockingResources = in.BlockingResources
	return out
}
func Instruction_ToProto(mapCtx *direct.MapContext, in *krm.Instruction) *pb.Instruction {
	if in == nil {
		return nil
	}
	out := &pb.Instruction{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DataType = direct.Enum_ToProto[pb.DataType](mapCtx, in.DataType)
	out.CsvInstruction = CsvInstruction_ToProto(mapCtx, in.CsvInstruction)
	out.PdfInstruction = PdfInstruction_ToProto(mapCtx, in.PdfInstruction)
	out.BlockingResources = in.BlockingResources
	return out
}
func PdfInstruction_FromProto(mapCtx *direct.MapContext, in *pb.PdfInstruction) *krm.PdfInstruction {
	if in == nil {
		return nil
	}
	out := &krm.PdfInstruction{}
	out.GcsFileURI = direct.LazyPtr(in.GetGcsFileUri())
	return out
}
func PdfInstruction_ToProto(mapCtx *direct.MapContext, in *krm.PdfInstruction) *pb.PdfInstruction {
	if in == nil {
		return nil
	}
	out := &pb.PdfInstruction{}
	out.GcsFileUri = direct.ValueOf(in.GcsFileURI)
	return out
}
