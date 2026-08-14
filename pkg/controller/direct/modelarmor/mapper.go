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

package modelarmor

import (
	pb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Template_TemplateMetadata_FromProto(mapCtx *direct.MapContext, in *pb.Template_TemplateMetadata) *krm.Template_TemplateMetadata {
	if in == nil {
		return nil
	}
	out := &krm.Template_TemplateMetadata{}
	out.IgnorePartialInvocationFailures = direct.LazyPtr(in.GetIgnorePartialInvocationFailures())
	out.CustomPromptSafetyErrorCode = direct.LazyPtr(in.GetCustomPromptSafetyErrorCode())
	out.CustomPromptSafetyErrorMessage = direct.LazyPtr(in.GetCustomPromptSafetyErrorMessage())
	out.CustomLlmResponseSafetyErrorCode = direct.LazyPtr(in.GetCustomLlmResponseSafetyErrorCode())
	out.CustomLlmResponseSafetyErrorMessage = direct.LazyPtr(in.GetCustomLlmResponseSafetyErrorMessage())
	out.LogTemplateOperations = direct.LazyPtr(in.GetLogTemplateOperations())
	out.LogSanitizeOperations = direct.LazyPtr(in.GetLogSanitizeOperations())
	out.EnforcementType = direct.Enum_FromProto(mapCtx, in.GetEnforcementType())
	out.MultiLanguageDetection = Template_TemplateMetadata_MultiLanguageDetection_FromProto(mapCtx, in.GetMultiLanguageDetection())
	return out
}

func Template_TemplateMetadata_ToProto(mapCtx *direct.MapContext, in *krm.Template_TemplateMetadata) *pb.Template_TemplateMetadata {
	if in == nil {
		return nil
	}
	out := &pb.Template_TemplateMetadata{}
	if in.IgnorePartialInvocationFailures != nil {
		out.IgnorePartialInvocationFailures = *in.IgnorePartialInvocationFailures
	}
	if in.CustomPromptSafetyErrorCode != nil {
		out.CustomPromptSafetyErrorCode = *in.CustomPromptSafetyErrorCode
	}
	if in.CustomPromptSafetyErrorMessage != nil {
		out.CustomPromptSafetyErrorMessage = *in.CustomPromptSafetyErrorMessage
	}
	if in.CustomLlmResponseSafetyErrorCode != nil {
		out.CustomLlmResponseSafetyErrorCode = *in.CustomLlmResponseSafetyErrorCode
	}
	if in.CustomLlmResponseSafetyErrorMessage != nil {
		out.CustomLlmResponseSafetyErrorMessage = *in.CustomLlmResponseSafetyErrorMessage
	}
	if in.LogTemplateOperations != nil {
		out.LogTemplateOperations = *in.LogTemplateOperations
	}
	if in.LogSanitizeOperations != nil {
		out.LogSanitizeOperations = *in.LogSanitizeOperations
	}
	out.EnforcementType = direct.Enum_ToProto[pb.Template_TemplateMetadata_EnforcementType](mapCtx, in.EnforcementType)
	out.MultiLanguageDetection = Template_TemplateMetadata_MultiLanguageDetection_ToProto(mapCtx, in.MultiLanguageDetection)
	return out
}

func Template_TemplateMetadata_MultiLanguageDetection_FromProto(mapCtx *direct.MapContext, in *pb.Template_TemplateMetadata_MultiLanguageDetection) *krm.Template_TemplateMetadata_MultiLanguageDetection {
	if in == nil {
		return nil
	}
	out := &krm.Template_TemplateMetadata_MultiLanguageDetection{}
	out.EnableMultiLanguageDetection = direct.LazyPtr(in.GetEnableMultiLanguageDetection())
	return out
}

func Template_TemplateMetadata_MultiLanguageDetection_ToProto(mapCtx *direct.MapContext, in *krm.Template_TemplateMetadata_MultiLanguageDetection) *pb.Template_TemplateMetadata_MultiLanguageDetection {
	if in == nil {
		return nil
	}
	out := &pb.Template_TemplateMetadata_MultiLanguageDetection{}
	if in.EnableMultiLanguageDetection != nil {
		out.EnableMultiLanguageDetection = *in.EnableMultiLanguageDetection
	}
	return out
}
