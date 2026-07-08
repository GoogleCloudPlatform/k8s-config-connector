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

package datalabelingannotationspecset

import (
	pb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datalabeling/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func InstructionCsvInstruction_FromProto(mapCtx *direct.MapContext, in *pb.CsvInstruction) *krm.InstructionCsvInstruction {
	if in == nil {
		return nil
	}
	out := &krm.InstructionCsvInstruction{}
	out.GcsFileURI = direct.LazyPtr(in.GetGcsFileUri())
	return out
}

func InstructionCsvInstruction_ToProto(mapCtx *direct.MapContext, in *krm.InstructionCsvInstruction) *pb.CsvInstruction {
	if in == nil {
		return nil
	}
	out := &pb.CsvInstruction{}
	out.GcsFileUri = direct.ValueOf(in.GcsFileURI)
	return out
}

func InstructionPdfInstruction_FromProto(mapCtx *direct.MapContext, in *pb.PdfInstruction) *krm.InstructionPdfInstruction {
	if in == nil {
		return nil
	}
	out := &krm.InstructionPdfInstruction{}
	out.GcsFileURI = direct.LazyPtr(in.GetGcsFileUri())
	return out
}

func InstructionPdfInstruction_ToProto(mapCtx *direct.MapContext, in *krm.InstructionPdfInstruction) *pb.PdfInstruction {
	if in == nil {
		return nil
	}
	out := &pb.PdfInstruction{}
	out.GcsFileUri = direct.ValueOf(in.GcsFileURI)
	return out
}

func TextMetadata_FromProto(mapCtx *direct.MapContext, in *pb.TextMetadata) *krm.TextMetadata {
	if in == nil {
		return nil
	}
	out := &krm.TextMetadata{}
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}

func TextMetadata_ToProto(mapCtx *direct.MapContext, in *krm.TextMetadata) *pb.TextMetadata {
	if in == nil {
		return nil
	}
	out := &pb.TextMetadata{}
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}

func ClassificationMetadata_FromProto(mapCtx *direct.MapContext, in *pb.ClassificationMetadata) *krm.ClassificationMetadata {
	if in == nil {
		return nil
	}
	out := &krm.ClassificationMetadata{}
	out.IsMultiLabel = direct.LazyPtr(in.GetIsMultiLabel())
	return out
}

func ClassificationMetadata_ToProto(mapCtx *direct.MapContext, in *krm.ClassificationMetadata) *pb.ClassificationMetadata {
	if in == nil {
		return nil
	}
	out := &pb.ClassificationMetadata{}
	out.IsMultiLabel = direct.ValueOf(in.IsMultiLabel)
	return out
}
