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
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1beta1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AccountDefenderAssessment_FromProto(mapCtx *direct.MapContext, in *pb.AccountDefenderAssessment) *krm.AccountDefenderAssessment {
	if in == nil {
		return nil
	}
	out := &krm.AccountDefenderAssessment{}
	out.Labels = direct.EnumSlice_FromProto(mapCtx, in.Labels)
	return out
}
func AccountDefenderAssessment_ToProto(mapCtx *direct.MapContext, in *krm.AccountDefenderAssessment) *pb.AccountDefenderAssessment {
	if in == nil {
		return nil
	}
	out := &pb.AccountDefenderAssessment{}
	out.Labels = direct.EnumSlice_ToProto[pb.AccountDefenderAssessment_AccountDefenderLabel](mapCtx, in.Labels)
	return out
}
func Assessment_FromProto(mapCtx *direct.MapContext, in *pb.Assessment) *krm.Assessment {
	if in == nil {
		return nil
	}
	out := &krm.Assessment{}
	// MISSING: Name
	out.Event = Event_FromProto(mapCtx, in.GetEvent())
	// MISSING: Score
	// MISSING: TokenProperties
	// MISSING: Reasons
	out.PasswordLeakVerification = PasswordLeakVerification_FromProto(mapCtx, in.GetPasswordLeakVerification())
	out.AccountDefenderAssessment = AccountDefenderAssessment_FromProto(mapCtx, in.GetAccountDefenderAssessment())
	out.FraudPreventionAssessment = FraudPreventionAssessment_FromProto(mapCtx, in.GetFraudPreventionAssessment())
	return out
}
func Assessment_ToProto(mapCtx *direct.MapContext, in *krm.Assessment) *pb.Assessment {
	if in == nil {
		return nil
	}
	out := &pb.Assessment{}
	// MISSING: Name
	out.Event = Event_ToProto(mapCtx, in.Event)
	// MISSING: Score
	// MISSING: TokenProperties
	// MISSING: Reasons
	out.PasswordLeakVerification = PasswordLeakVerification_ToProto(mapCtx, in.PasswordLeakVerification)
	out.AccountDefenderAssessment = AccountDefenderAssessment_ToProto(mapCtx, in.AccountDefenderAssessment)
	out.FraudPreventionAssessment = FraudPreventionAssessment_ToProto(mapCtx, in.FraudPreventionAssessment)
	return out
}
func AssessmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Assessment) *krm.AssessmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssessmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Event
	out.Score = direct.LazyPtr(in.GetScore())
	out.TokenProperties = TokenProperties_FromProto(mapCtx, in.GetTokenProperties())
	out.Reasons = direct.EnumSlice_FromProto(mapCtx, in.Reasons)
	out.PasswordLeakVerification = PasswordLeakVerificationObservedState_FromProto(mapCtx, in.GetPasswordLeakVerification())
	// MISSING: AccountDefenderAssessment
	// MISSING: FraudPreventionAssessment
	return out
}
func AssessmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssessmentObservedState) *pb.Assessment {
	if in == nil {
		return nil
	}
	out := &pb.Assessment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Event
	out.Score = direct.ValueOf(in.Score)
	out.TokenProperties = TokenProperties_ToProto(mapCtx, in.TokenProperties)
	out.Reasons = direct.EnumSlice_ToProto[pb.Assessment_ClassificationReason](mapCtx, in.Reasons)
	out.PasswordLeakVerification = PasswordLeakVerificationObservedState_ToProto(mapCtx, in.PasswordLeakVerification)
	// MISSING: AccountDefenderAssessment
	// MISSING: FraudPreventionAssessment
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
	out.TransactionData = TransactionData_FromProto(mapCtx, in.GetTransactionData())
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
	out.TransactionData = TransactionData_ToProto(mapCtx, in.TransactionData)
	out.FraudPrevention = direct.Enum_ToProto[pb.Event_FraudPrevention](mapCtx, in.FraudPrevention)
	return out
}
func FraudPreventionAssessment_FromProto(mapCtx *direct.MapContext, in *pb.FraudPreventionAssessment) *krm.FraudPreventionAssessment {
	if in == nil {
		return nil
	}
	out := &krm.FraudPreventionAssessment{}
	out.TransactionRisk = direct.LazyPtr(in.GetTransactionRisk())
	out.StolenInstrumentVerdict = FraudPreventionAssessment_StolenInstrumentVerdict_FromProto(mapCtx, in.GetStolenInstrumentVerdict())
	out.CardTestingVerdict = FraudPreventionAssessment_CardTestingVerdict_FromProto(mapCtx, in.GetCardTestingVerdict())
	out.BehavioralTrustVerdict = FraudPreventionAssessment_BehavioralTrustVerdict_FromProto(mapCtx, in.GetBehavioralTrustVerdict())
	return out
}
func FraudPreventionAssessment_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment) *pb.FraudPreventionAssessment {
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
	out.Trust = direct.LazyPtr(in.GetTrust())
	return out
}
func FraudPreventionAssessment_BehavioralTrustVerdict_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_BehavioralTrustVerdict) *pb.FraudPreventionAssessment_BehavioralTrustVerdict {
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
	out.Risk = direct.LazyPtr(in.GetRisk())
	return out
}
func FraudPreventionAssessment_CardTestingVerdict_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_CardTestingVerdict) *pb.FraudPreventionAssessment_CardTestingVerdict {
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
	out.Risk = direct.LazyPtr(in.GetRisk())
	return out
}
func FraudPreventionAssessment_StolenInstrumentVerdict_ToProto(mapCtx *direct.MapContext, in *krm.FraudPreventionAssessment_StolenInstrumentVerdict) *pb.FraudPreventionAssessment_StolenInstrumentVerdict {
	if in == nil {
		return nil
	}
	out := &pb.FraudPreventionAssessment_StolenInstrumentVerdict{}
	out.Risk = direct.ValueOf(in.Risk)
	return out
}
func PasswordLeakVerification_FromProto(mapCtx *direct.MapContext, in *pb.PasswordLeakVerification) *krm.PasswordLeakVerification {
	if in == nil {
		return nil
	}
	out := &krm.PasswordLeakVerification{}
	out.HashedUserCredentials = in.GetHashedUserCredentials()
	// MISSING: CredentialsLeaked
	out.CanonicalizedUsername = direct.LazyPtr(in.GetCanonicalizedUsername())
	return out
}
func PasswordLeakVerification_ToProto(mapCtx *direct.MapContext, in *krm.PasswordLeakVerification) *pb.PasswordLeakVerification {
	if in == nil {
		return nil
	}
	out := &pb.PasswordLeakVerification{}
	out.HashedUserCredentials = in.HashedUserCredentials
	// MISSING: CredentialsLeaked
	out.CanonicalizedUsername = direct.ValueOf(in.CanonicalizedUsername)
	return out
}
func PasswordLeakVerificationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PasswordLeakVerification) *krm.PasswordLeakVerificationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PasswordLeakVerificationObservedState{}
	// MISSING: HashedUserCredentials
	out.CredentialsLeaked = direct.LazyPtr(in.GetCredentialsLeaked())
	// MISSING: CanonicalizedUsername
	return out
}
func PasswordLeakVerificationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PasswordLeakVerificationObservedState) *pb.PasswordLeakVerification {
	if in == nil {
		return nil
	}
	out := &pb.PasswordLeakVerification{}
	// MISSING: HashedUserCredentials
	out.CredentialsLeaked = direct.ValueOf(in.CredentialsLeaked)
	// MISSING: CanonicalizedUsername
	return out
}
func TokenProperties_FromProto(mapCtx *direct.MapContext, in *pb.TokenProperties) *krm.TokenProperties {
	if in == nil {
		return nil
	}
	out := &krm.TokenProperties{}
	out.Valid = direct.LazyPtr(in.GetValid())
	out.InvalidReason = direct.Enum_FromProto(mapCtx, in.GetInvalidReason())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.Action = direct.LazyPtr(in.GetAction())
	return out
}
func TokenProperties_ToProto(mapCtx *direct.MapContext, in *krm.TokenProperties) *pb.TokenProperties {
	if in == nil {
		return nil
	}
	out := &pb.TokenProperties{}
	out.Valid = direct.ValueOf(in.Valid)
	out.InvalidReason = direct.Enum_ToProto[pb.TokenProperties_InvalidReason](mapCtx, in.InvalidReason)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Hostname = direct.ValueOf(in.Hostname)
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
