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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AdvancedSettings_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedSettings) *krm.AdvancedSettings {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedSettings{}
	out.AudioExportGcsDestination = GcsDestination_FromProto(mapCtx, in.GetAudioExportGcsDestination())
	out.SpeechSettings = AdvancedSettings_SpeechSettings_FromProto(mapCtx, in.GetSpeechSettings())
	out.DtmfSettings = AdvancedSettings_DtmfSettings_FromProto(mapCtx, in.GetDtmfSettings())
	out.LoggingSettings = AdvancedSettings_LoggingSettings_FromProto(mapCtx, in.GetLoggingSettings())
	return out
}
func AdvancedSettings_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedSettings) *pb.AdvancedSettings {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedSettings{}
	out.AudioExportGcsDestination = GcsDestination_ToProto(mapCtx, in.AudioExportGcsDestination)
	out.SpeechSettings = AdvancedSettings_SpeechSettings_ToProto(mapCtx, in.SpeechSettings)
	out.DtmfSettings = AdvancedSettings_DtmfSettings_ToProto(mapCtx, in.DtmfSettings)
	out.LoggingSettings = AdvancedSettings_LoggingSettings_ToProto(mapCtx, in.LoggingSettings)
	return out
}
func AdvancedSettings_DtmfSettings_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedSettings_DtmfSettings) *krm.AdvancedSettings_DtmfSettings {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedSettings_DtmfSettings{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.MaxDigits = direct.LazyPtr(in.GetMaxDigits())
	out.FinishDigit = direct.LazyPtr(in.GetFinishDigit())
	out.InterdigitTimeoutDuration = direct.StringDuration_FromProto(mapCtx, in.GetInterdigitTimeoutDuration())
	out.EndpointingTimeoutDuration = direct.StringDuration_FromProto(mapCtx, in.GetEndpointingTimeoutDuration())
	return out
}
func AdvancedSettings_DtmfSettings_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedSettings_DtmfSettings) *pb.AdvancedSettings_DtmfSettings {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedSettings_DtmfSettings{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.MaxDigits = direct.ValueOf(in.MaxDigits)
	out.FinishDigit = direct.ValueOf(in.FinishDigit)
	out.InterdigitTimeoutDuration = direct.StringDuration_ToProto(mapCtx, in.InterdigitTimeoutDuration)
	out.EndpointingTimeoutDuration = direct.StringDuration_ToProto(mapCtx, in.EndpointingTimeoutDuration)
	return out
}
func AdvancedSettings_LoggingSettings_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedSettings_LoggingSettings) *krm.AdvancedSettings_LoggingSettings {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedSettings_LoggingSettings{}
	out.EnableStackdriverLogging = direct.LazyPtr(in.GetEnableStackdriverLogging())
	out.EnableInteractionLogging = direct.LazyPtr(in.GetEnableInteractionLogging())
	out.EnableConsentBasedRedaction = direct.LazyPtr(in.GetEnableConsentBasedRedaction())
	return out
}
func AdvancedSettings_LoggingSettings_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedSettings_LoggingSettings) *pb.AdvancedSettings_LoggingSettings {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedSettings_LoggingSettings{}
	out.EnableStackdriverLogging = direct.ValueOf(in.EnableStackdriverLogging)
	out.EnableInteractionLogging = direct.ValueOf(in.EnableInteractionLogging)
	out.EnableConsentBasedRedaction = direct.ValueOf(in.EnableConsentBasedRedaction)
	return out
}
func AdvancedSettings_SpeechSettings_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedSettings_SpeechSettings) *krm.AdvancedSettings_SpeechSettings {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedSettings_SpeechSettings{}
	out.EndpointerSensitivity = direct.LazyPtr(in.GetEndpointerSensitivity())
	out.NoSpeechTimeout = direct.StringDuration_FromProto(mapCtx, in.GetNoSpeechTimeout())
	out.UseTimeoutBasedEndpointing = direct.LazyPtr(in.GetUseTimeoutBasedEndpointing())
	out.Models = in.Models
	return out
}
func AdvancedSettings_SpeechSettings_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedSettings_SpeechSettings) *pb.AdvancedSettings_SpeechSettings {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedSettings_SpeechSettings{}
	out.EndpointerSensitivity = direct.ValueOf(in.EndpointerSensitivity)
	out.NoSpeechTimeout = direct.StringDuration_ToProto(mapCtx, in.NoSpeechTimeout)
	out.UseTimeoutBasedEndpointing = direct.ValueOf(in.UseTimeoutBasedEndpointing)
	out.Models = in.Models
	return out
}
func Agent_FromProto(mapCtx *direct.MapContext, in *pb.Agent) *krm.Agent {
	if in == nil {
		return nil
	}
	out := &krm.Agent{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DefaultLanguageCode = direct.LazyPtr(in.GetDefaultLanguageCode())
	out.SupportedLanguageCodes = in.SupportedLanguageCodes
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.AvatarURI = direct.LazyPtr(in.GetAvatarUri())
	out.SpeechToTextSettings = SpeechToTextSettings_FromProto(mapCtx, in.GetSpeechToTextSettings())
	out.StartFlow = direct.LazyPtr(in.GetStartFlow())
	out.StartPlaybook = direct.LazyPtr(in.GetStartPlaybook())
	out.SecuritySettings = direct.LazyPtr(in.GetSecuritySettings())
	out.EnableStackdriverLogging = direct.LazyPtr(in.GetEnableStackdriverLogging())
	out.EnableSpellCorrection = direct.LazyPtr(in.GetEnableSpellCorrection())
	out.EnableMultiLanguageTraining = direct.LazyPtr(in.GetEnableMultiLanguageTraining())
	out.Locked = direct.LazyPtr(in.GetLocked())
	out.AdvancedSettings = AdvancedSettings_FromProto(mapCtx, in.GetAdvancedSettings())
	out.GitIntegrationSettings = Agent_GitIntegrationSettings_FromProto(mapCtx, in.GetGitIntegrationSettings())
	out.BigqueryExportSettings = BigQueryExportSettings_FromProto(mapCtx, in.GetBigqueryExportSettings())
	out.TextToSpeechSettings = TextToSpeechSettings_FromProto(mapCtx, in.GetTextToSpeechSettings())
	out.GenAppBuilderSettings = Agent_GenAppBuilderSettings_FromProto(mapCtx, in.GetGenAppBuilderSettings())
	out.AnswerFeedbackSettings = Agent_AnswerFeedbackSettings_FromProto(mapCtx, in.GetAnswerFeedbackSettings())
	out.PersonalizationSettings = Agent_PersonalizationSettings_FromProto(mapCtx, in.GetPersonalizationSettings())
	out.ClientCertificateSettings = Agent_ClientCertificateSettings_FromProto(mapCtx, in.GetClientCertificateSettings())
	return out
}
func Agent_ToProto(mapCtx *direct.MapContext, in *krm.Agent) *pb.Agent {
	if in == nil {
		return nil
	}
	out := &pb.Agent{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DefaultLanguageCode = direct.ValueOf(in.DefaultLanguageCode)
	out.SupportedLanguageCodes = in.SupportedLanguageCodes
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.Description = direct.ValueOf(in.Description)
	out.AvatarUri = direct.ValueOf(in.AvatarURI)
	out.SpeechToTextSettings = SpeechToTextSettings_ToProto(mapCtx, in.SpeechToTextSettings)
	if oneof := Agent_StartFlow_ToProto(mapCtx, in.StartFlow); oneof != nil {
		out.SessionEntryResource = oneof
	}
	if oneof := Agent_StartPlaybook_ToProto(mapCtx, in.StartPlaybook); oneof != nil {
		out.SessionEntryResource = oneof
	}
	out.SecuritySettings = direct.ValueOf(in.SecuritySettings)
	out.EnableStackdriverLogging = direct.ValueOf(in.EnableStackdriverLogging)
	out.EnableSpellCorrection = direct.ValueOf(in.EnableSpellCorrection)
	out.EnableMultiLanguageTraining = direct.ValueOf(in.EnableMultiLanguageTraining)
	out.Locked = direct.ValueOf(in.Locked)
	out.AdvancedSettings = AdvancedSettings_ToProto(mapCtx, in.AdvancedSettings)
	out.GitIntegrationSettings = Agent_GitIntegrationSettings_ToProto(mapCtx, in.GitIntegrationSettings)
	out.BigqueryExportSettings = BigQueryExportSettings_ToProto(mapCtx, in.BigqueryExportSettings)
	out.TextToSpeechSettings = TextToSpeechSettings_ToProto(mapCtx, in.TextToSpeechSettings)
	if oneof := Agent_GenAppBuilderSettings_ToProto(mapCtx, in.GenAppBuilderSettings); oneof != nil {
		out.GenAppBuilderSettings = &pb.Agent_GenAppBuilderSettings_{GenAppBuilderSettings: oneof}
	}
	out.AnswerFeedbackSettings = Agent_AnswerFeedbackSettings_ToProto(mapCtx, in.AnswerFeedbackSettings)
	out.PersonalizationSettings = Agent_PersonalizationSettings_ToProto(mapCtx, in.PersonalizationSettings)
	out.ClientCertificateSettings = Agent_ClientCertificateSettings_ToProto(mapCtx, in.ClientCertificateSettings)
	return out
}
func Agent_AnswerFeedbackSettings_FromProto(mapCtx *direct.MapContext, in *pb.Agent_AnswerFeedbackSettings) *krm.Agent_AnswerFeedbackSettings {
	if in == nil {
		return nil
	}
	out := &krm.Agent_AnswerFeedbackSettings{}
	out.EnableAnswerFeedback = direct.LazyPtr(in.GetEnableAnswerFeedback())
	return out
}
func Agent_AnswerFeedbackSettings_ToProto(mapCtx *direct.MapContext, in *krm.Agent_AnswerFeedbackSettings) *pb.Agent_AnswerFeedbackSettings {
	if in == nil {
		return nil
	}
	out := &pb.Agent_AnswerFeedbackSettings{}
	out.EnableAnswerFeedback = direct.ValueOf(in.EnableAnswerFeedback)
	return out
}
func Agent_ClientCertificateSettings_FromProto(mapCtx *direct.MapContext, in *pb.Agent_ClientCertificateSettings) *krm.Agent_ClientCertificateSettings {
	if in == nil {
		return nil
	}
	out := &krm.Agent_ClientCertificateSettings{}
	out.SslCertificate = direct.LazyPtr(in.GetSslCertificate())
	out.PrivateKey = direct.LazyPtr(in.GetPrivateKey())
	out.Passphrase = direct.LazyPtr(in.GetPassphrase())
	return out
}
func Agent_ClientCertificateSettings_ToProto(mapCtx *direct.MapContext, in *krm.Agent_ClientCertificateSettings) *pb.Agent_ClientCertificateSettings {
	if in == nil {
		return nil
	}
	out := &pb.Agent_ClientCertificateSettings{}
	out.SslCertificate = direct.ValueOf(in.SslCertificate)
	out.PrivateKey = direct.ValueOf(in.PrivateKey)
	out.Passphrase = direct.ValueOf(in.Passphrase)
	return out
}
func Agent_GenAppBuilderSettings_FromProto(mapCtx *direct.MapContext, in *pb.Agent_GenAppBuilderSettings) *krm.Agent_GenAppBuilderSettings {
	if in == nil {
		return nil
	}
	out := &krm.Agent_GenAppBuilderSettings{}
	out.Engine = direct.LazyPtr(in.GetEngine())
	return out
}
func Agent_GenAppBuilderSettings_ToProto(mapCtx *direct.MapContext, in *krm.Agent_GenAppBuilderSettings) *pb.Agent_GenAppBuilderSettings {
	if in == nil {
		return nil
	}
	out := &pb.Agent_GenAppBuilderSettings{}
	out.Engine = direct.ValueOf(in.Engine)
	return out
}
func Agent_GitIntegrationSettings_FromProto(mapCtx *direct.MapContext, in *pb.Agent_GitIntegrationSettings) *krm.Agent_GitIntegrationSettings {
	if in == nil {
		return nil
	}
	out := &krm.Agent_GitIntegrationSettings{}
	out.GithubSettings = Agent_GitIntegrationSettings_GithubSettings_FromProto(mapCtx, in.GetGithubSettings())
	return out
}
func Agent_GitIntegrationSettings_ToProto(mapCtx *direct.MapContext, in *krm.Agent_GitIntegrationSettings) *pb.Agent_GitIntegrationSettings {
	if in == nil {
		return nil
	}
	out := &pb.Agent_GitIntegrationSettings{}
	if oneof := Agent_GitIntegrationSettings_GithubSettings_ToProto(mapCtx, in.GithubSettings); oneof != nil {
		out.GitSettings = &pb.Agent_GitIntegrationSettings_GithubSettings_{GithubSettings: oneof}
	}
	return out
}
func Agent_GitIntegrationSettings_GithubSettings_FromProto(mapCtx *direct.MapContext, in *pb.Agent_GitIntegrationSettings_GithubSettings) *krm.Agent_GitIntegrationSettings_GithubSettings {
	if in == nil {
		return nil
	}
	out := &krm.Agent_GitIntegrationSettings_GithubSettings{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.RepositoryURI = direct.LazyPtr(in.GetRepositoryUri())
	out.TrackingBranch = direct.LazyPtr(in.GetTrackingBranch())
	out.AccessToken = direct.LazyPtr(in.GetAccessToken())
	out.Branches = in.Branches
	return out
}
func Agent_GitIntegrationSettings_GithubSettings_ToProto(mapCtx *direct.MapContext, in *krm.Agent_GitIntegrationSettings_GithubSettings) *pb.Agent_GitIntegrationSettings_GithubSettings {
	if in == nil {
		return nil
	}
	out := &pb.Agent_GitIntegrationSettings_GithubSettings{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.RepositoryUri = direct.ValueOf(in.RepositoryURI)
	out.TrackingBranch = direct.ValueOf(in.TrackingBranch)
	out.AccessToken = direct.ValueOf(in.AccessToken)
	out.Branches = in.Branches
	return out
}
func Agent_PersonalizationSettings_FromProto(mapCtx *direct.MapContext, in *pb.Agent_PersonalizationSettings) *krm.Agent_PersonalizationSettings {
	if in == nil {
		return nil
	}
	out := &krm.Agent_PersonalizationSettings{}
	out.DefaultEndUserMetadata = DefaultEndUserMetadata_FromProto(mapCtx, in.GetDefaultEndUserMetadata())
	return out
}
func Agent_PersonalizationSettings_ToProto(mapCtx *direct.MapContext, in *krm.Agent_PersonalizationSettings) *pb.Agent_PersonalizationSettings {
	if in == nil {
		return nil
	}
	out := &pb.Agent_PersonalizationSettings{}
	out.DefaultEndUserMetadata = DefaultEndUserMetadata_ToProto(mapCtx, in.DefaultEndUserMetadata)
	return out
}
func BigQueryExportSettings_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryExportSettings) *krm.BigQueryExportSettings {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryExportSettings{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.BigqueryTable = direct.LazyPtr(in.GetBigqueryTable())
	return out
}
func BigQueryExportSettings_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryExportSettings) *pb.BigQueryExportSettings {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryExportSettings{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.BigqueryTable = direct.ValueOf(in.BigqueryTable)
	return out
}
func GcsDestination_FromProto(mapCtx *direct.MapContext, in *pb.GcsDestination) *krm.GcsDestination {
	if in == nil {
		return nil
	}
	out := &krm.GcsDestination{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func GcsDestination_ToProto(mapCtx *direct.MapContext, in *krm.GcsDestination) *pb.GcsDestination {
	if in == nil {
		return nil
	}
	out := &pb.GcsDestination{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func SpeechToTextSettings_FromProto(mapCtx *direct.MapContext, in *pb.SpeechToTextSettings) *krm.SpeechToTextSettings {
	if in == nil {
		return nil
	}
	out := &krm.SpeechToTextSettings{}
	out.EnableSpeechAdaptation = direct.LazyPtr(in.GetEnableSpeechAdaptation())
	return out
}
func SpeechToTextSettings_ToProto(mapCtx *direct.MapContext, in *krm.SpeechToTextSettings) *pb.SpeechToTextSettings {
	if in == nil {
		return nil
	}
	out := &pb.SpeechToTextSettings{}
	out.EnableSpeechAdaptation = direct.ValueOf(in.EnableSpeechAdaptation)
	return out
}
func SynthesizeSpeechConfig_FromProto(mapCtx *direct.MapContext, in *pb.SynthesizeSpeechConfig) *krm.SynthesizeSpeechConfig {
	if in == nil {
		return nil
	}
	out := &krm.SynthesizeSpeechConfig{}
	out.SpeakingRate = direct.LazyPtr(in.GetSpeakingRate())
	out.Pitch = direct.LazyPtr(in.GetPitch())
	out.VolumeGainDb = direct.LazyPtr(in.GetVolumeGainDb())
	out.EffectsProfileID = in.EffectsProfileId
	out.Voice = VoiceSelectionParams_FromProto(mapCtx, in.GetVoice())
	return out
}
func SynthesizeSpeechConfig_ToProto(mapCtx *direct.MapContext, in *krm.SynthesizeSpeechConfig) *pb.SynthesizeSpeechConfig {
	if in == nil {
		return nil
	}
	out := &pb.SynthesizeSpeechConfig{}
	out.SpeakingRate = direct.ValueOf(in.SpeakingRate)
	out.Pitch = direct.ValueOf(in.Pitch)
	out.VolumeGainDb = direct.ValueOf(in.VolumeGainDb)
	out.EffectsProfileId = in.EffectsProfileID
	out.Voice = VoiceSelectionParams_ToProto(mapCtx, in.Voice)
	return out
}
func TextToSpeechSettings_FromProto(mapCtx *direct.MapContext, in *pb.TextToSpeechSettings) *krm.TextToSpeechSettings {
	if in == nil {
		return nil
	}
	out := &krm.TextToSpeechSettings{}
	// MISSING: SynthesizeSpeechConfigs
	return out
}
func TextToSpeechSettings_ToProto(mapCtx *direct.MapContext, in *krm.TextToSpeechSettings) *pb.TextToSpeechSettings {
	if in == nil {
		return nil
	}
	out := &pb.TextToSpeechSettings{}
	// MISSING: SynthesizeSpeechConfigs
	return out
}
func VoiceSelectionParams_FromProto(mapCtx *direct.MapContext, in *pb.VoiceSelectionParams) *krm.VoiceSelectionParams {
	if in == nil {
		return nil
	}
	out := &krm.VoiceSelectionParams{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SsmlGender = direct.Enum_FromProto(mapCtx, in.GetSsmlGender())
	return out
}
func VoiceSelectionParams_ToProto(mapCtx *direct.MapContext, in *krm.VoiceSelectionParams) *pb.VoiceSelectionParams {
	if in == nil {
		return nil
	}
	out := &pb.VoiceSelectionParams{}
	out.Name = direct.ValueOf(in.Name)
	out.SsmlGender = direct.Enum_ToProto[pb.SsmlVoiceGender](mapCtx, in.SsmlGender)
	return out
}
