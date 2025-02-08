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

package recaptchaenterprise

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AccountDefenderAssessment_FromProto(mapCtx *direct.MapContext, in *pb.AccountDefenderAssessment) *krm.AccountDefenderAssessment {
	if in == nil {
		return nil
	}
	out := &krm.AccountDefenderAssessment{}
	// MISSING: Labels
	return out
}
func AccountDefenderAssessment_ToProto(mapCtx *direct.MapContext, in *krm.AccountDefenderAssessment) *pb.AccountDefenderAssessment {
	if in == nil {
		return nil
	}
	out := &pb.AccountDefenderAssessment{}
	// MISSING: Labels
	return out
}
func AccountDefenderAssessmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountDefenderAssessment) *krm.AccountDefenderAssessmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountDefenderAssessmentObservedState{}
	out.Labels = direct.EnumSlice_FromProto(mapCtx, in.Labels)
	return out
}
func AccountDefenderAssessmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountDefenderAssessmentObservedState) *pb.AccountDefenderAssessment {
	if in == nil {
		return nil
	}
	out := &pb.AccountDefenderAssessment{}
	out.Labels = direct.EnumSlice_ToProto[pb.AccountDefenderAssessment_AccountDefenderLabel](mapCtx, in.Labels)
	return out
}
func AccountVerificationInfo_FromProto(mapCtx *direct.MapContext, in *pb.AccountVerificationInfo) *krm.AccountVerificationInfo {
	if in == nil {
		return nil
	}
	out := &krm.AccountVerificationInfo{}
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, EndpointVerificationInfo_FromProto)
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	// MISSING: LatestVerificationResult
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}
func AccountVerificationInfo_ToProto(mapCtx *direct.MapContext, in *krm.AccountVerificationInfo) *pb.AccountVerificationInfo {
	if in == nil {
		return nil
	}
	out := &pb.AccountVerificationInfo{}
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, EndpointVerificationInfo_ToProto)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	// MISSING: LatestVerificationResult
	out.Username = direct.ValueOf(in.Username)
	return out
}
func AccountVerificationInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountVerificationInfo) *krm.AccountVerificationInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountVerificationInfoObservedState{}
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, EndpointVerificationInfoObservedState_FromProto)
	// MISSING: LanguageCode
	out.LatestVerificationResult = direct.Enum_FromProto(mapCtx, in.GetLatestVerificationResult())
	// MISSING: Username
	return out
}
func AccountVerificationInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountVerificationInfoObservedState) *pb.AccountVerificationInfo {
	if in == nil {
		return nil
	}
	out := &pb.AccountVerificationInfo{}
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, EndpointVerificationInfoObservedState_ToProto)
	// MISSING: LanguageCode
	out.LatestVerificationResult = direct.Enum_ToProto[pb.AccountVerificationInfo_Result](mapCtx, in.LatestVerificationResult)
	// MISSING: Username
	return out
}
func Assessment_FromProto(mapCtx *direct.MapContext, in *pb.Assessment) *krm.Assessment {
	if in == nil {
		return nil
	}
	out := &krm.Assessment{}
	// MISSING: Name
	out.Event = Event_FromProto(mapCtx, in.GetEvent())
	// MISSING: RiskAnalysis
	// MISSING: TokenProperties
	out.AccountVerification = AccountVerificationInfo_FromProto(mapCtx, in.GetAccountVerification())
	// MISSING: AccountDefenderAssessment
	out.PrivatePasswordLeakVerification = PrivatePasswordLeakVerification_FromProto(mapCtx, in.GetPrivatePasswordLeakVerification())
	// MISSING: FirewallPolicyAssessment
	// MISSING: FraudPreventionAssessment
	// MISSING: FraudSignals
	// MISSING: PhoneFraudAssessment
	out.AssessmentEnvironment = AssessmentEnvironment_FromProto(mapCtx, in.GetAssessmentEnvironment())
	return out
}
func Assessment_ToProto(mapCtx *direct.MapContext, in *krm.Assessment) *pb.Assessment {
	if in == nil {
		return nil
	}
	out := &pb.Assessment{}
	// MISSING: Name
	out.Event = Event_ToProto(mapCtx, in.Event)
	// MISSING: RiskAnalysis
	// MISSING: TokenProperties
	out.AccountVerification = AccountVerificationInfo_ToProto(mapCtx, in.AccountVerification)
	// MISSING: AccountDefenderAssessment
	out.PrivatePasswordLeakVerification = PrivatePasswordLeakVerification_ToProto(mapCtx, in.PrivatePasswordLeakVerification)
	// MISSING: FirewallPolicyAssessment
	// MISSING: FraudPreventionAssessment
	// MISSING: FraudSignals
	// MISSING: PhoneFraudAssessment
	out.AssessmentEnvironment = AssessmentEnvironment_ToProto(mapCtx, in.AssessmentEnvironment)
	return out
}
func AssessmentEnvironment_FromProto(mapCtx *direct.MapContext, in *pb.AssessmentEnvironment) *krm.AssessmentEnvironment {
	if in == nil {
		return nil
	}
	out := &krm.AssessmentEnvironment{}
	out.Client = direct.LazyPtr(in.GetClient())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func AssessmentEnvironment_ToProto(mapCtx *direct.MapContext, in *krm.AssessmentEnvironment) *pb.AssessmentEnvironment {
	if in == nil {
		return nil
	}
	out := &pb.AssessmentEnvironment{}
	out.Client = direct.ValueOf(in.Client)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func AssessmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Assessment) *krm.AssessmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssessmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Event
	out.RiskAnalysis = RiskAnalysis_FromProto(mapCtx, in.GetRiskAnalysis())
	out.TokenProperties = TokenProperties_FromProto(mapCtx, in.GetTokenProperties())
	out.AccountVerification = AccountVerificationInfoObservedState_FromProto(mapCtx, in.GetAccountVerification())
	out.AccountDefenderAssessment = AccountDefenderAssessment_FromProto(mapCtx, in.GetAccountDefenderAssessment())
	out.PrivatePasswordLeakVerification = PrivatePasswordLeakVerificationObservedState_FromProto(mapCtx, in.GetPrivatePasswordLeakVerification())
	out.FirewallPolicyAssessment = FirewallPolicyAssessment_FromProto(mapCtx, in.GetFirewallPolicyAssessment())
	out.FraudPreventionAssessment = FraudPreventionAssessment_FromProto(mapCtx, in.GetFraudPreventionAssessment())
	out.FraudSignals = FraudSignals_FromProto(mapCtx, in.GetFraudSignals())
	out.PhoneFraudAssessment = PhoneFraudAssessment_FromProto(mapCtx, in.GetPhoneFraudAssessment())
	// MISSING: AssessmentEnvironment
	return out
}
func AssessmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssessmentObservedState) *pb.Assessment {
	if in == nil {
		return nil
	}
	out := &pb.Assessment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Event
	out.RiskAnalysis = RiskAnalysis_ToProto(mapCtx, in.RiskAnalysis)
	out.TokenProperties = TokenProperties_ToProto(mapCtx, in.TokenProperties)
	out.AccountVerification = AccountVerificationInfoObservedState_ToProto(mapCtx, in.AccountVerification)
	out.AccountDefenderAssessment = AccountDefenderAssessment_ToProto(mapCtx, in.AccountDefenderAssessment)
	out.PrivatePasswordLeakVerification = PrivatePasswordLeakVerificationObservedState_ToProto(mapCtx, in.PrivatePasswordLeakVerification)
	out.FirewallPolicyAssessment = FirewallPolicyAssessment_ToProto(mapCtx, in.FirewallPolicyAssessment)
	out.FraudPreventionAssessment = FraudPreventionAssessment_ToProto(mapCtx, in.FraudPreventionAssessment)
	out.FraudSignals = FraudSignals_ToProto(mapCtx, in.FraudSignals)
	out.PhoneFraudAssessment = PhoneFraudAssessment_ToProto(mapCtx, in.PhoneFraudAssessment)
	// MISSING: AssessmentEnvironment
	return out
}
func EndpointVerificationInfo_FromProto(mapCtx *direct.MapContext, in *pb.EndpointVerificationInfo) *krm.EndpointVerificationInfo {
	if in == nil {
		return nil
	}
	out := &krm.EndpointVerificationInfo{}
	out.EmailAddress = direct.LazyPtr(in.GetEmailAddress())
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	// MISSING: RequestToken
	// MISSING: LastVerificationTime
	return out
}
func EndpointVerificationInfo_ToProto(mapCtx *direct.MapContext, in *krm.EndpointVerificationInfo) *pb.EndpointVerificationInfo {
	if in == nil {
		return nil
	}
	out := &pb.EndpointVerificationInfo{}
	if oneof := EndpointVerificationInfo_EmailAddress_ToProto(mapCtx, in.EmailAddress); oneof != nil {
		out.Endpoint = oneof
	}
	if oneof := EndpointVerificationInfo_PhoneNumber_ToProto(mapCtx, in.PhoneNumber); oneof != nil {
		out.Endpoint = oneof
	}
	// MISSING: RequestToken
	// MISSING: LastVerificationTime
	return out
}
func EndpointVerificationInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EndpointVerificationInfo) *krm.EndpointVerificationInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointVerificationInfoObservedState{}
	// MISSING: EmailAddress
	// MISSING: PhoneNumber
	out.RequestToken = direct.LazyPtr(in.GetRequestToken())
	out.LastVerificationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastVerificationTime())
	return out
}
func EndpointVerificationInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointVerificationInfoObservedState) *pb.EndpointVerificationInfo {
	if in == nil {
		return nil
	}
	out := &pb.EndpointVerificationInfo{}
	// MISSING: EmailAddress
	// MISSING: PhoneNumber
	out.RequestToken = direct.ValueOf(in.RequestToken)
	out.LastVerificationTime = direct.StringTimestamp_ToProto(mapCtx, in.LastVerificationTime)
	return out
}
func Event_FromProto(mapCtx *direct.MapContext, in *pb.Event) *krm.Event {
	if in == nil {
		return nil
	}
	out := &krm.Event{}
	out.Token = direct.LazyPtr(in.GetToken())
	out.SiteKey = direct.LazyPtr(in.GetSiteKey())
	out.UserAgent = direct.LazyPtr(in.GetUserAgent())
	out.UserIPAddress = direct.LazyPtr(in.GetUserIpAddress())
	out.ExpectedAction = direct.LazyPtr(in.GetExpectedAction())
	out.HashedAccountID = in.GetHashedAccountId()
	out.Express = direct.LazyPtr(in.GetExpress())
	out.RequestedURI = direct.LazyPtr(in.GetRequestedUri())
	out.WafTokenAssessment = direct.LazyPtr(in.GetWafTokenAssessment())
	out.Ja3 = direct.LazyPtr(in.GetJa3())
	out.Headers = in.Headers
	out.FirewallPolicyEvaluation = direct.LazyPtr(in.GetFirewallPolicyEvaluation())
	out.TransactionData = TransactionData_FromProto(mapCtx, in.GetTransactionData())
	out.UserInfo = UserInfo_FromProto(mapCtx, in.GetUserInfo())
	out.FraudPrevention = direct.Enum_FromProto(mapCtx, in.GetFraudPrevention())
	return out
}
func Event_ToProto(mapCtx *direct.MapContext, in *krm.Event) *pb.Event {
	if in == nil {
		return nil
	}
	out := &pb.Event{}
	out.Token = direct.ValueOf(in.Token)
	out.SiteKey = direct.ValueOf(in.SiteKey)
	out.UserAgent = direct.ValueOf(in.UserAgent)
	out.UserIpAddress = direct.ValueOf(in.UserIPAddress)
	out.ExpectedAction = direct.ValueOf(in.ExpectedAction)
	out.HashedAccountId = in.HashedAccountID
	out.Express = direct.ValueOf(in.Express)
	out.RequestedUri = direct.ValueOf(in.RequestedURI)
	out.WafTokenAssessment = direct.ValueOf(in.WafTokenAssessment)
	out.Ja3 = direct.ValueOf(in.Ja3)
	out.Headers = in.Headers
	out.FirewallPolicyEvaluation = direct.ValueOf(in.FirewallPolicyEvaluation)
	out.TransactionData = TransactionData_ToProto(mapCtx, in.TransactionData)
	out.UserInfo = UserInfo_ToProto(mapCtx, in.UserInfo)
	out.FraudPrevention = direct.Enum_ToProto[pb.Event_FraudPrevention](mapCtx, in.FraudPrevention)
	return out
}
func FirewallAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction) *krm.FirewallAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction{}
	out.Allow = FirewallAction_AllowAction_FromProto(mapCtx, in.GetAllow())
	out.Block = FirewallAction_BlockAction_FromProto(mapCtx, in.GetBlock())
	out.IncludeRecaptchaScript = FirewallAction_IncludeRecaptchaScriptAction_FromProto(mapCtx, in.GetIncludeRecaptchaScript())
	out.Redirect = FirewallAction_RedirectAction_FromProto(mapCtx, in.GetRedirect())
	out.Substitute = FirewallAction_SubstituteAction_FromProto(mapCtx, in.GetSubstitute())
	out.SetHeader = FirewallAction_SetHeaderAction_FromProto(mapCtx, in.GetSetHeader())
	return out
}
func FirewallAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction) *pb.FirewallAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction{}
	if oneof := FirewallAction_AllowAction_ToProto(mapCtx, in.Allow); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_Allow{Allow: oneof}
	}
	if oneof := FirewallAction_BlockAction_ToProto(mapCtx, in.Block); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_Block{Block: oneof}
	}
	if oneof := FirewallAction_IncludeRecaptchaScriptAction_ToProto(mapCtx, in.IncludeRecaptchaScript); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_IncludeRecaptchaScript{IncludeRecaptchaScript: oneof}
	}
	if oneof := FirewallAction_RedirectAction_ToProto(mapCtx, in.Redirect); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_Redirect{Redirect: oneof}
	}
	if oneof := FirewallAction_SubstituteAction_ToProto(mapCtx, in.Substitute); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_Substitute{Substitute: oneof}
	}
	if oneof := FirewallAction_SetHeaderAction_ToProto(mapCtx, in.SetHeader); oneof != nil {
		out.FirewallActionOneof = &pb.FirewallAction_SetHeader{SetHeader: oneof}
	}
	return out
}
func FirewallAction_AllowAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_AllowAction) *krm.FirewallAction_AllowAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_AllowAction{}
	return out
}
func FirewallAction_AllowAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_AllowAction) *pb.FirewallAction_AllowAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_AllowAction{}
	return out
}
func FirewallAction_BlockAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_BlockAction) *krm.FirewallAction_BlockAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_BlockAction{}
	return out
}
func FirewallAction_BlockAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_BlockAction) *pb.FirewallAction_BlockAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_BlockAction{}
	return out
}
func FirewallAction_IncludeRecaptchaScriptAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_IncludeRecaptchaScriptAction) *krm.FirewallAction_IncludeRecaptchaScriptAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_IncludeRecaptchaScriptAction{}
	return out
}
func FirewallAction_IncludeRecaptchaScriptAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_IncludeRecaptchaScriptAction) *pb.FirewallAction_IncludeRecaptchaScriptAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_IncludeRecaptchaScriptAction{}
	return out
}
func FirewallAction_RedirectAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_RedirectAction) *krm.FirewallAction_RedirectAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_RedirectAction{}
	return out
}
func FirewallAction_RedirectAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_RedirectAction) *pb.FirewallAction_RedirectAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_RedirectAction{}
	return out
}
func FirewallAction_SetHeaderAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_SetHeaderAction) *krm.FirewallAction_SetHeaderAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_SetHeaderAction{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func FirewallAction_SetHeaderAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_SetHeaderAction) *pb.FirewallAction_SetHeaderAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_SetHeaderAction{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func FirewallAction_SubstituteAction_FromProto(mapCtx *direct.MapContext, in *pb.FirewallAction_SubstituteAction) *krm.FirewallAction_SubstituteAction {
	if in == nil {
		return nil
	}
	out := &krm.FirewallAction_SubstituteAction{}
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func FirewallAction_SubstituteAction_ToProto(mapCtx *direct.MapContext, in *krm.FirewallAction_SubstituteAction) *pb.FirewallAction_SubstituteAction {
	if in == nil {
		return nil
	}
	out := &pb.FirewallAction_SubstituteAction{}
	out.Path = direct.ValueOf(in.Path)
	return out
}
func FirewallPolicy_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Condition = direct.LazyPtr(in.GetCondition())
	out.Actions = direct.Slice_FromProto(mapCtx, in.Actions, FirewallAction_FromProto)
	return out
}
func FirewallPolicy_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicy) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Path = direct.ValueOf(in.Path)
	out.Condition = direct.ValueOf(in.Condition)
	out.Actions = direct.Slice_ToProto(mapCtx, in.Actions, FirewallAction_ToProto)
	return out
}
func FirewallPolicyAssessment_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyAssessment) *krm.FirewallPolicyAssessment {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyAssessment{}
	// MISSING: Error
	// MISSING: FirewallPolicy
	return out
}
func FirewallPolicyAssessment_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyAssessment) *pb.FirewallPolicyAssessment {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyAssessment{}
	// MISSING: Error
	// MISSING: FirewallPolicy
	return out
}
func FirewallPolicyAssessmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicyAssessment) *krm.FirewallPolicyAssessmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FirewallPolicyAssessmentObservedState{}
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.FirewallPolicy = FirewallPolicy_FromProto(mapCtx, in.GetFirewallPolicy())
	return out
}
func FirewallPolicyAssessmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FirewallPolicyAssessmentObservedState) *pb.FirewallPolicyAssessment {
	if in == nil {
		return nil
	}
	out := &pb.FirewallPolicyAssessment{}
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.FirewallPolicy = FirewallPolicy_ToProto(mapCtx, in.FirewallPolicy)
	return out
}
func FraudPreventionAssessment_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment) *krm.FraudPreventionAssessment {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessment{}
	// MISSING: TransactionRisk
	// MISSING: StolenInstrumentVerdict
	// MISSING: CardTestingVerdict
	// MISSING: BehavioralTrustVerdict
	return out
}
func FraudPreventionAssessment_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment) *pb.FraudPreventionAssessment {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment{}
	// MISSING: TransactionRisk
	// MISSING: StolenInstrumentVerdict
	// MISSING: CardTestingVerdict
	// MISSING: BehavioralTrustVerdict
	return out
}
func FraudPreventionAssessmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment) *krm.FraudPreventionAssessmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessmentObservedState{}
	out.TransactionRisk = direct.LazyPtr(in.GetTransactionRisk())
	out.StolenInstrumentVerdict = FraudPreventionAssessment_StolenInstrumentVerdict_FromProto(mapCtx, in.GetStolenInstrumentVerdict())
	out.CardTestingVerdict = FraudPreventionAssessment_CardTestingVerdict_FromProto(mapCtx, in.GetCardTestingVerdict())
	out.BehavioralTrustVerdict = FraudPreventionAssessment_BehavioralTrustVerdict_FromProto(mapCtx, in.GetBehavioralTrustVerdict())
	return out
}
func FraudPreventionAssessmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessmentObservedState) *pb.FraudPreventionAssessment {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment{}
	out.TransactionRisk = direct.ValueOf(in.TransactionRisk)
	out.StolenInstrumentVerdict = FraudPreventionAssessment_StolenInstrumentVerdict_ToProto(mapCtx, in.StolenInstrumentVerdict)
	out.CardTestingVerdict = FraudPreventionAssessment_CardTestingVerdict_ToProto(mapCtx, in.CardTestingVerdict)
	out.BehavioralTrustVerdict = FraudPreventionAssessment_BehavioralTrustVerdict_ToProto(mapCtx, in.BehavioralTrustVerdict)
	return out
}
func FraudPreventionAssessment_BehavioralTrustVerdict_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment_BehavioralTrustVerdict) *krm.FraudPreventionAssessment_BehavioralTrustVerdict {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessment_BehavioralTrustVerdict{}
	// MISSING: Trust
	return out
}
func FraudPreventionAssessment_BehavioralTrustVerdict_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_BehavioralTrustVerdict) *pb.FraudPreventionAssessment_BehavioralTrustVerdict {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment_BehavioralTrustVerdict{}
	// MISSING: Trust
	return out
}
func FraudPreventionAssessment_BehavioralTrustVerdictObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment_BehavioralTrustVerdict) *krm.FraudPreventionAssessment_BehavioralTrustVerdictObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessment_BehavioralTrustVerdictObservedState{}
	out.Trust = direct.LazyPtr(in.GetTrust())
	return out
}
func FraudPreventionAssessment_BehavioralTrustVerdictObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_BehavioralTrustVerdictObservedState) *pb.FraudPreventionAssessment_BehavioralTrustVerdict {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment_BehavioralTrustVerdict{}
	out.Trust = direct.ValueOf(in.Trust)
	return out
}
func FraudPreventionAssessment_CardTestingVerdict_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment_CardTestingVerdict) *krm.FraudPreventionAssessment_CardTestingVerdict {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessment_CardTestingVerdict{}
	// MISSING: Risk
	return out
}
func FraudPreventionAssessment_CardTestingVerdict_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_CardTestingVerdict) *pb.FraudPreventionAssessment_CardTestingVerdict {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment_CardTestingVerdict{}
	// MISSING: Risk
	return out
}
func FraudPreventionAssessment_CardTestingVerdictObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment_CardTestingVerdict) *krm.FraudPreventionAssessment_CardTestingVerdictObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessment_CardTestingVerdictObservedState{}
	out.Risk = direct.LazyPtr(in.GetRisk())
	return out
}
func FraudPreventionAssessment_CardTestingVerdictObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_CardTestingVerdictObservedState) *pb.FraudPreventionAssessment_CardTestingVerdict {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment_CardTestingVerdict{}
	out.Risk = direct.ValueOf(in.Risk)
	return out
}
func FraudPreventionAssessment_StolenInstrumentVerdict_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment_StolenInstrumentVerdict) *krm.FraudPreventionAssessment_StolenInstrumentVerdict {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessment_StolenInstrumentVerdict{}
	// MISSING: Risk
	return out
}
func FraudPreventionAssessment_StolenInstrumentVerdict_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_StolenInstrumentVerdict) *pb.FraudPreventionAssessment_StolenInstrumentVerdict {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment_StolenInstrumentVerdict{}
	// MISSING: Risk
	return out
}
func FraudPreventionAssessment_StolenInstrumentVerdictObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment_StolenInstrumentVerdict) *krm.FraudPreventionAssessment_StolenInstrumentVerdictObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessment_StolenInstrumentVerdictObservedState{}
	out.Risk = direct.LazyPtr(in.GetRisk())
	return out
}
func FraudPreventionAssessment_StolenInstrumentVerdictObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_StolenInstrumentVerdictObservedState) *pb.FraudPreventionAssessment_StolenInstrumentVerdict {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment_StolenInstrumentVerdict{}
	out.Risk = direct.ValueOf(in.Risk)
	return out
}
func FraudSignals_FromProto(mapCtx *direct.MapContext, in *pb.FraudSignals) *krm.FraudSignals {
	if in == nil {
		return nil
	}
	out := &krm.FraudSignals{}
	// MISSING: UserSignals
	// MISSING: CardSignals
	return out
}
func FraudSignals_ToProto(mapCtx *direct.MapContext, in *krm.FraudSignals) *pb.FraudSignals {
	if in == nil {
		return nil
	}
	out := &pb.FraudSignals{}
	// MISSING: UserSignals
	// MISSING: CardSignals
	return out
}
func FraudSignalsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FraudSignals) *krm.FraudSignalsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FraudSignalsObservedState{}
	out.UserSignals = FraudSignals_UserSignals_FromProto(mapCtx, in.GetUserSignals())
	out.CardSignals = FraudSignals_CardSignals_FromProto(mapCtx, in.GetCardSignals())
	return out
}
func FraudSignalsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FraudSignalsObservedState) *pb.FraudSignals {
	if in == nil {
		return nil
	}
	out := &pb.FraudSignals{}
	out.UserSignals = FraudSignals_UserSignals_ToProto(mapCtx, in.UserSignals)
	out.CardSignals = FraudSignals_CardSignals_ToProto(mapCtx, in.CardSignals)
	return out
}
func FraudSignals_CardSignals_FromProto(mapCtx *direct.MapContext, in *pb.FraudSignals_CardSignals) *krm.FraudSignals_CardSignals {
	if in == nil {
		return nil
	}
	out := &krm.FraudSignals_CardSignals{}
	// MISSING: CardLabels
	return out
}
func FraudSignals_CardSignals_ToProto(mapCtx *direct.MapContext, in *krm.FraudSignals_CardSignals) *pb.FraudSignals_CardSignals {
	if in == nil {
		return nil
	}
	out := &pb.FraudSignals_CardSignals{}
	// MISSING: CardLabels
	return out
}
func FraudSignals_CardSignalsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FraudSignals_CardSignals) *krm.FraudSignals_CardSignalsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FraudSignals_CardSignalsObservedState{}
	out.CardLabels = direct.EnumSlice_FromProto(mapCtx, in.CardLabels)
	return out
}
func FraudSignals_CardSignalsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FraudSignals_CardSignalsObservedState) *pb.FraudSignals_CardSignals {
	if in == nil {
		return nil
	}
	out := &pb.FraudSignals_CardSignals{}
	out.CardLabels = direct.EnumSlice_ToProto[pb.FraudSignals_CardSignals_CardLabel](mapCtx, in.CardLabels)
	return out
}
func FraudSignals_UserSignals_FromProto(mapCtx *direct.MapContext, in *pb.FraudSignals_UserSignals) *krm.FraudSignals_UserSignals {
	if in == nil {
		return nil
	}
	out := &krm.FraudSignals_UserSignals{}
	// MISSING: ActiveDaysLowerBound
	// MISSING: SyntheticRisk
	return out
}
func FraudSignals_UserSignals_ToProto(mapCtx *direct.MapContext, in *krm.FraudSignals_UserSignals) *pb.FraudSignals_UserSignals {
	if in == nil {
		return nil
	}
	out := &pb.FraudSignals_UserSignals{}
	// MISSING: ActiveDaysLowerBound
	// MISSING: SyntheticRisk
	return out
}
func FraudSignals_UserSignalsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FraudSignals_UserSignals) *krm.FraudSignals_UserSignalsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.FraudSignals_UserSignalsObservedState{}
	out.ActiveDaysLowerBound = direct.LazyPtr(in.GetActiveDaysLowerBound())
	out.SyntheticRisk = direct.LazyPtr(in.GetSyntheticRisk())
	return out
}
func FraudSignals_UserSignalsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.FraudSignals_UserSignalsObservedState) *pb.FraudSignals_UserSignals {
	if in == nil {
		return nil
	}
	out := &pb.FraudSignals_UserSignals{}
	out.ActiveDaysLowerBound = direct.ValueOf(in.ActiveDaysLowerBound)
	out.SyntheticRisk = direct.ValueOf(in.SyntheticRisk)
	return out
}
func PhoneFraudAssessment_FromProto(mapCtx *direct.MapContext, in *pb.PhoneFraudAssessment) *krm.PhoneFraudAssessment {
	if in == nil {
		return nil
	}
	out := &krm.PhoneFraudAssessment{}
	// MISSING: SmsTollFraudVerdict
	return out
}
func PhoneFraudAssessment_ToProto(mapCtx *direct.MapContext, in *krm.PhoneFraudAssessment) *pb.PhoneFraudAssessment {
	if in == nil {
		return nil
	}
	out := &pb.PhoneFraudAssessment{}
	// MISSING: SmsTollFraudVerdict
	return out
}
func PhoneFraudAssessmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PhoneFraudAssessment) *krm.PhoneFraudAssessmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PhoneFraudAssessmentObservedState{}
	out.SmsTollFraudVerdict = SmsTollFraudVerdict_FromProto(mapCtx, in.GetSmsTollFraudVerdict())
	return out
}
func PhoneFraudAssessmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PhoneFraudAssessmentObservedState) *pb.PhoneFraudAssessment {
	if in == nil {
		return nil
	}
	out := &pb.PhoneFraudAssessment{}
	out.SmsTollFraudVerdict = SmsTollFraudVerdict_ToProto(mapCtx, in.SmsTollFraudVerdict)
	return out
}
func PrivatePasswordLeakVerification_FromProto(mapCtx *direct.MapContext, in *pb.PrivatePasswordLeakVerification) *krm.PrivatePasswordLeakVerification {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePasswordLeakVerification{}
	out.LookupHashPrefix = in.GetLookupHashPrefix()
	out.EncryptedUserCredentialsHash = in.GetEncryptedUserCredentialsHash()
	// MISSING: EncryptedLeakMatchPrefixes
	// MISSING: ReencryptedUserCredentialsHash
	return out
}
func PrivatePasswordLeakVerification_ToProto(mapCtx *direct.MapContext, in *krm.PrivatePasswordLeakVerification) *pb.PrivatePasswordLeakVerification {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePasswordLeakVerification{}
	out.LookupHashPrefix = in.LookupHashPrefix
	out.EncryptedUserCredentialsHash = in.EncryptedUserCredentialsHash
	// MISSING: EncryptedLeakMatchPrefixes
	// MISSING: ReencryptedUserCredentialsHash
	return out
}
func PrivatePasswordLeakVerificationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivatePasswordLeakVerification) *krm.PrivatePasswordLeakVerificationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePasswordLeakVerificationObservedState{}
	// MISSING: LookupHashPrefix
	// MISSING: EncryptedUserCredentialsHash
	out.EncryptedLeakMatchPrefixes = in.EncryptedLeakMatchPrefixes
	out.ReencryptedUserCredentialsHash = in.GetReencryptedUserCredentialsHash()
	return out
}
func PrivatePasswordLeakVerificationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivatePasswordLeakVerificationObservedState) *pb.PrivatePasswordLeakVerification {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePasswordLeakVerification{}
	// MISSING: LookupHashPrefix
	// MISSING: EncryptedUserCredentialsHash
	out.EncryptedLeakMatchPrefixes = in.EncryptedLeakMatchPrefixes
	out.ReencryptedUserCredentialsHash = in.ReencryptedUserCredentialsHash
	return out
}
func RecaptchaenterpriseAssessmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Assessment) *krm.RecaptchaenterpriseAssessmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseAssessmentObservedState{}
	// MISSING: Name
	// MISSING: Event
	// MISSING: RiskAnalysis
	// MISSING: TokenProperties
	// MISSING: AccountVerification
	// MISSING: AccountDefenderAssessment
	// MISSING: PrivatePasswordLeakVerification
	// MISSING: FirewallPolicyAssessment
	// MISSING: FraudPreventionAssessment
	// MISSING: FraudSignals
	// MISSING: PhoneFraudAssessment
	// MISSING: AssessmentEnvironment
	return out
}
func RecaptchaenterpriseAssessmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseAssessmentObservedState) *pb.Assessment {
	if in == nil {
		return nil
	}
	out := &pb.Assessment{}
	// MISSING: Name
	// MISSING: Event
	// MISSING: RiskAnalysis
	// MISSING: TokenProperties
	// MISSING: AccountVerification
	// MISSING: AccountDefenderAssessment
	// MISSING: PrivatePasswordLeakVerification
	// MISSING: FirewallPolicyAssessment
	// MISSING: FraudPreventionAssessment
	// MISSING: FraudSignals
	// MISSING: PhoneFraudAssessment
	// MISSING: AssessmentEnvironment
	return out
}
func RecaptchaenterpriseAssessmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Assessment) *krm.RecaptchaenterpriseAssessmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseAssessmentSpec{}
	// MISSING: Name
	// MISSING: Event
	// MISSING: RiskAnalysis
	// MISSING: TokenProperties
	// MISSING: AccountVerification
	// MISSING: AccountDefenderAssessment
	// MISSING: PrivatePasswordLeakVerification
	// MISSING: FirewallPolicyAssessment
	// MISSING: FraudPreventionAssessment
	// MISSING: FraudSignals
	// MISSING: PhoneFraudAssessment
	// MISSING: AssessmentEnvironment
	return out
}
func RecaptchaenterpriseAssessmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseAssessmentSpec) *pb.Assessment {
	if in == nil {
		return nil
	}
	out := &pb.Assessment{}
	// MISSING: Name
	// MISSING: Event
	// MISSING: RiskAnalysis
	// MISSING: TokenProperties
	// MISSING: AccountVerification
	// MISSING: AccountDefenderAssessment
	// MISSING: PrivatePasswordLeakVerification
	// MISSING: FirewallPolicyAssessment
	// MISSING: FraudPreventionAssessment
	// MISSING: FraudSignals
	// MISSING: PhoneFraudAssessment
	// MISSING: AssessmentEnvironment
	return out
}
func RiskAnalysis_FromProto(mapCtx *direct.MapContext, in *pb.RiskAnalysis) *krm.RiskAnalysis {
	if in == nil {
		return nil
	}
	out := &krm.RiskAnalysis{}
	// MISSING: Score
	// MISSING: Reasons
	// MISSING: ExtendedVerdictReasons
	// MISSING: Challenge
	return out
}
func RiskAnalysis_ToProto(mapCtx *direct.MapContext, in *krm.RiskAnalysis) *pb.RiskAnalysis {
	if in == nil {
		return nil
	}
	out := &pb.RiskAnalysis{}
	// MISSING: Score
	// MISSING: Reasons
	// MISSING: ExtendedVerdictReasons
	// MISSING: Challenge
	return out
}
func RiskAnalysisObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RiskAnalysis) *krm.RiskAnalysisObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RiskAnalysisObservedState{}
	out.Score = direct.LazyPtr(in.GetScore())
	out.Reasons = direct.EnumSlice_FromProto(mapCtx, in.Reasons)
	out.ExtendedVerdictReasons = in.ExtendedVerdictReasons
	out.Challenge = direct.Enum_FromProto(mapCtx, in.GetChallenge())
	return out
}
func RiskAnalysisObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RiskAnalysisObservedState) *pb.RiskAnalysis {
	if in == nil {
		return nil
	}
	out := &pb.RiskAnalysis{}
	out.Score = direct.ValueOf(in.Score)
	out.Reasons = direct.EnumSlice_ToProto[pb.RiskAnalysis_ClassificationReason](mapCtx, in.Reasons)
	out.ExtendedVerdictReasons = in.ExtendedVerdictReasons
	out.Challenge = direct.Enum_ToProto[pb.RiskAnalysis_Challenge](mapCtx, in.Challenge)
	return out
}
func SmsTollFraudVerdict_FromProto(mapCtx *direct.MapContext, in *pb.SmsTollFraudVerdict) *krm.SmsTollFraudVerdict {
	if in == nil {
		return nil
	}
	out := &krm.SmsTollFraudVerdict{}
	// MISSING: Risk
	// MISSING: Reasons
	return out
}
func SmsTollFraudVerdict_ToProto(mapCtx *direct.MapContext, in *krm.SmsTollFraudVerdict) *pb.SmsTollFraudVerdict {
	if in == nil {
		return nil
	}
	out := &pb.SmsTollFraudVerdict{}
	// MISSING: Risk
	// MISSING: Reasons
	return out
}
func SmsTollFraudVerdictObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SmsTollFraudVerdict) *krm.SmsTollFraudVerdictObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SmsTollFraudVerdictObservedState{}
	out.Risk = direct.LazyPtr(in.GetRisk())
	out.Reasons = direct.EnumSlice_FromProto(mapCtx, in.Reasons)
	return out
}
func SmsTollFraudVerdictObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SmsTollFraudVerdictObservedState) *pb.SmsTollFraudVerdict {
	if in == nil {
		return nil
	}
	out := &pb.SmsTollFraudVerdict{}
	out.Risk = direct.ValueOf(in.Risk)
	out.Reasons = direct.EnumSlice_ToProto[pb.SmsTollFraudVerdict_SmsTollFraudReason](mapCtx, in.Reasons)
	return out
}
func TokenProperties_FromProto(mapCtx *direct.MapContext, in *pb.TokenProperties) *krm.TokenProperties {
	if in == nil {
		return nil
	}
	out := &krm.TokenProperties{}
	// MISSING: Valid
	// MISSING: InvalidReason
	// MISSING: CreateTime
	// MISSING: Hostname
	// MISSING: AndroidPackageName
	// MISSING: IosBundleID
	// MISSING: Action
	return out
}
func TokenProperties_ToProto(mapCtx *direct.MapContext, in *krm.TokenProperties) *pb.TokenProperties {
	if in == nil {
		return nil
	}
	out := &pb.TokenProperties{}
	// MISSING: Valid
	// MISSING: InvalidReason
	// MISSING: CreateTime
	// MISSING: Hostname
	// MISSING: AndroidPackageName
	// MISSING: IosBundleID
	// MISSING: Action
	return out
}
func TokenPropertiesObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TokenProperties) *krm.TokenPropertiesObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TokenPropertiesObservedState{}
	out.Valid = direct.LazyPtr(in.GetValid())
	out.InvalidReason = direct.Enum_FromProto(mapCtx, in.GetInvalidReason())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.AndroidPackageName = direct.LazyPtr(in.GetAndroidPackageName())
	out.IosBundleID = direct.LazyPtr(in.GetIosBundleId())
	out.Action = direct.LazyPtr(in.GetAction())
	return out
}
func TokenPropertiesObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TokenPropertiesObservedState) *pb.TokenProperties {
	if in == nil {
		return nil
	}
	out := &pb.TokenProperties{}
	out.Valid = direct.ValueOf(in.Valid)
	out.InvalidReason = direct.Enum_ToProto[pb.TokenProperties_InvalidReason](mapCtx, in.InvalidReason)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Hostname = direct.ValueOf(in.Hostname)
	out.AndroidPackageName = direct.ValueOf(in.AndroidPackageName)
	out.IosBundleId = direct.ValueOf(in.IosBundleID)
	out.Action = direct.ValueOf(in.Action)
	return out
}
func TransactionData_FromProto(mapCtx *direct.MapContext, in *pb.TransactionData) *krm.TransactionData {
	if in == nil {
		return nil
	}
	out := &krm.TransactionData{}
	out.TransactionID = in.TransactionId
	out.PaymentMethod = direct.LazyPtr(in.GetPaymentMethod())
	out.CardBin = direct.LazyPtr(in.GetCardBin())
	out.CardLastFour = direct.LazyPtr(in.GetCardLastFour())
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Value = direct.LazyPtr(in.GetValue())
	out.ShippingValue = direct.LazyPtr(in.GetShippingValue())
	out.ShippingAddress = TransactionData_Address_FromProto(mapCtx, in.GetShippingAddress())
	out.BillingAddress = TransactionData_Address_FromProto(mapCtx, in.GetBillingAddress())
	out.User = TransactionData_User_FromProto(mapCtx, in.GetUser())
	out.Merchants = direct.Slice_FromProto(mapCtx, in.Merchants, TransactionData_User_FromProto)
	out.Items = direct.Slice_FromProto(mapCtx, in.Items, TransactionData_Item_FromProto)
	out.GatewayInfo = TransactionData_GatewayInfo_FromProto(mapCtx, in.GetGatewayInfo())
	return out
}
func TransactionData_ToProto(mapCtx *direct.MapContext, in *krm.TransactionData) *pb.TransactionData {
	if in == nil {
		return nil
	}
	out := &pb.TransactionData{}
	out.TransactionId = in.TransactionID
	out.PaymentMethod = direct.ValueOf(in.PaymentMethod)
	out.CardBin = direct.ValueOf(in.CardBin)
	out.CardLastFour = direct.ValueOf(in.CardLastFour)
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Value = direct.ValueOf(in.Value)
	out.ShippingValue = direct.ValueOf(in.ShippingValue)
	out.ShippingAddress = TransactionData_Address_ToProto(mapCtx, in.ShippingAddress)
	out.BillingAddress = TransactionData_Address_ToProto(mapCtx, in.BillingAddress)
	out.User = TransactionData_User_ToProto(mapCtx, in.User)
	out.Merchants = direct.Slice_ToProto(mapCtx, in.Merchants, TransactionData_User_ToProto)
	out.Items = direct.Slice_ToProto(mapCtx, in.Items, TransactionData_Item_ToProto)
	out.GatewayInfo = TransactionData_GatewayInfo_ToProto(mapCtx, in.GatewayInfo)
	return out
}
func TransactionData_Address_FromProto(mapCtx *direct.MapContext, in *pb.TransactionData_Address) *krm.TransactionData_Address {
	if in == nil {
		return nil
	}
	out := &krm.TransactionData_Address{}
	out.Recipient = direct.LazyPtr(in.GetRecipient())
	out.Address = in.Address
	out.Locality = direct.LazyPtr(in.GetLocality())
	out.AdministrativeArea = direct.LazyPtr(in.GetAdministrativeArea())
	out.RegionCode = direct.LazyPtr(in.GetRegionCode())
	out.PostalCode = direct.LazyPtr(in.GetPostalCode())
	return out
}
func TransactionData_Address_ToProto(mapCtx *direct.MapContext, in *krm.TransactionData_Address) *pb.TransactionData_Address {
	if in == nil {
		return nil
	}
	out := &pb.TransactionData_Address{}
	out.Recipient = direct.ValueOf(in.Recipient)
	out.Address = in.Address
	out.Locality = direct.ValueOf(in.Locality)
	out.AdministrativeArea = direct.ValueOf(in.AdministrativeArea)
	out.RegionCode = direct.ValueOf(in.RegionCode)
	out.PostalCode = direct.ValueOf(in.PostalCode)
	return out
}
func TransactionData_GatewayInfo_FromProto(mapCtx *direct.MapContext, in *pb.TransactionData_GatewayInfo) *krm.TransactionData_GatewayInfo {
	if in == nil {
		return nil
	}
	out := &krm.TransactionData_GatewayInfo{}
	out.Name = direct.LazyPtr(in.GetName())
	out.GatewayResponseCode = direct.LazyPtr(in.GetGatewayResponseCode())
	out.AvsResponseCode = direct.LazyPtr(in.GetAvsResponseCode())
	out.CvvResponseCode = direct.LazyPtr(in.GetCvvResponseCode())
	return out
}
func TransactionData_GatewayInfo_ToProto(mapCtx *direct.MapContext, in *krm.TransactionData_GatewayInfo) *pb.TransactionData_GatewayInfo {
	if in == nil {
		return nil
	}
	out := &pb.TransactionData_GatewayInfo{}
	out.Name = direct.ValueOf(in.Name)
	out.GatewayResponseCode = direct.ValueOf(in.GatewayResponseCode)
	out.AvsResponseCode = direct.ValueOf(in.AvsResponseCode)
	out.CvvResponseCode = direct.ValueOf(in.CvvResponseCode)
	return out
}
func TransactionData_Item_FromProto(mapCtx *direct.MapContext, in *pb.TransactionData_Item) *krm.TransactionData_Item {
	if in == nil {
		return nil
	}
	out := &krm.TransactionData_Item{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	out.Quantity = direct.LazyPtr(in.GetQuantity())
	out.MerchantAccountID = direct.LazyPtr(in.GetMerchantAccountId())
	return out
}
func TransactionData_Item_ToProto(mapCtx *direct.MapContext, in *krm.TransactionData_Item) *pb.TransactionData_Item {
	if in == nil {
		return nil
	}
	out := &pb.TransactionData_Item{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	out.Quantity = direct.ValueOf(in.Quantity)
	out.MerchantAccountId = direct.ValueOf(in.MerchantAccountID)
	return out
}
func TransactionData_User_FromProto(mapCtx *direct.MapContext, in *pb.TransactionData_User) *krm.TransactionData_User {
	if in == nil {
		return nil
	}
	out := &krm.TransactionData_User{}
	out.AccountID = direct.LazyPtr(in.GetAccountId())
	out.CreationMs = direct.LazyPtr(in.GetCreationMs())
	out.Email = direct.LazyPtr(in.GetEmail())
	out.EmailVerified = direct.LazyPtr(in.GetEmailVerified())
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	out.PhoneVerified = direct.LazyPtr(in.GetPhoneVerified())
	return out
}
func TransactionData_User_ToProto(mapCtx *direct.MapContext, in *krm.TransactionData_User) *pb.TransactionData_User {
	if in == nil {
		return nil
	}
	out := &pb.TransactionData_User{}
	out.AccountId = direct.ValueOf(in.AccountID)
	out.CreationMs = direct.ValueOf(in.CreationMs)
	out.Email = direct.ValueOf(in.Email)
	out.EmailVerified = direct.ValueOf(in.EmailVerified)
	out.PhoneNumber = direct.ValueOf(in.PhoneNumber)
	out.PhoneVerified = direct.ValueOf(in.PhoneVerified)
	return out
}
func UserId_FromProto(mapCtx *direct.MapContext, in *pb.UserId) *krm.UserId {
	if in == nil {
		return nil
	}
	out := &krm.UserId{}
	out.Email = direct.LazyPtr(in.GetEmail())
	out.PhoneNumber = direct.LazyPtr(in.GetPhoneNumber())
	out.Username = direct.LazyPtr(in.GetUsername())
	return out
}
func UserId_ToProto(mapCtx *direct.MapContext, in *krm.UserId) *pb.UserId {
	if in == nil {
		return nil
	}
	out := &pb.UserId{}
	if oneof := UserId_Email_ToProto(mapCtx, in.Email); oneof != nil {
		out.IdOneof = oneof
	}
	if oneof := UserId_PhoneNumber_ToProto(mapCtx, in.PhoneNumber); oneof != nil {
		out.IdOneof = oneof
	}
	if oneof := UserId_Username_ToProto(mapCtx, in.Username); oneof != nil {
		out.IdOneof = oneof
	}
	return out
}
func UserInfo_FromProto(mapCtx *direct.MapContext, in *pb.UserInfo) *krm.UserInfo {
	if in == nil {
		return nil
	}
	out := &krm.UserInfo{}
	out.CreateAccountTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateAccountTime())
	out.AccountID = direct.LazyPtr(in.GetAccountId())
	out.UserIds = direct.Slice_FromProto(mapCtx, in.UserIds, UserId_FromProto)
	return out
}
func UserInfo_ToProto(mapCtx *direct.MapContext, in *krm.UserInfo) *pb.UserInfo {
	if in == nil {
		return nil
	}
	out := &pb.UserInfo{}
	out.CreateAccountTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateAccountTime)
	out.AccountId = direct.ValueOf(in.AccountID)
	out.UserIds = direct.Slice_ToProto(mapCtx, in.UserIds, UserId_ToProto)
	return out
}
