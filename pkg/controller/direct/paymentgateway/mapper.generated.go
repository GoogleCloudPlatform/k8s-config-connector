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

package paymentgateway

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/paymentgateway/issuerswitch/apiv1/issuerswitchpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/paymentgateway/v1alpha1"
)
func AccountReference_FromProto(mapCtx *direct.MapContext, in *pb.AccountReference) *krm.AccountReference {
	if in == nil {
		return nil
	}
	out := &krm.AccountReference{}
	out.Ifsc = direct.LazyPtr(in.GetIfsc())
	// MISSING: AccountType
	out.AccountNumber = direct.LazyPtr(in.GetAccountNumber())
	return out
}
func AccountReference_ToProto(mapCtx *direct.MapContext, in *krm.AccountReference) *pb.AccountReference {
	if in == nil {
		return nil
	}
	out := &pb.AccountReference{}
	out.Ifsc = direct.ValueOf(in.Ifsc)
	// MISSING: AccountType
	out.AccountNumber = direct.ValueOf(in.AccountNumber)
	return out
}
func AccountReferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountReference) *krm.AccountReferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountReferenceObservedState{}
	// MISSING: Ifsc
	out.AccountType = direct.LazyPtr(in.GetAccountType())
	// MISSING: AccountNumber
	return out
}
func AccountReferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountReferenceObservedState) *pb.AccountReference {
	if in == nil {
		return nil
	}
	out := &pb.AccountReference{}
	// MISSING: Ifsc
	out.AccountType = direct.ValueOf(in.AccountType)
	// MISSING: AccountNumber
	return out
}
func CaseDetails_FromProto(mapCtx *direct.MapContext, in *pb.CaseDetails) *krm.CaseDetails {
	if in == nil {
		return nil
	}
	out := &krm.CaseDetails{}
	out.OriginalTransaction = OriginalTransaction_FromProto(mapCtx, in.GetOriginalTransaction())
	out.TransactionSubType = direct.Enum_FromProto(mapCtx, in.GetTransactionSubType())
	out.Amount = Money_FromProto(mapCtx, in.GetAmount())
	out.OriginalSettlementResponseCode = direct.LazyPtr(in.GetOriginalSettlementResponseCode())
	out.CurrentCycle = direct.LazyPtr(in.GetCurrentCycle())
	return out
}
func CaseDetails_ToProto(mapCtx *direct.MapContext, in *krm.CaseDetails) *pb.CaseDetails {
	if in == nil {
		return nil
	}
	out := &pb.CaseDetails{}
	out.OriginalTransaction = OriginalTransaction_ToProto(mapCtx, in.OriginalTransaction)
	out.TransactionSubType = direct.Enum_ToProto[pb.TransactionSubType](mapCtx, in.TransactionSubType)
	out.Amount = Money_ToProto(mapCtx, in.Amount)
	out.OriginalSettlementResponseCode = direct.ValueOf(in.OriginalSettlementResponseCode)
	out.CurrentCycle = direct.ValueOf(in.CurrentCycle)
	return out
}
func DeviceDetails_FromProto(mapCtx *direct.MapContext, in *pb.DeviceDetails) *krm.DeviceDetails {
	if in == nil {
		return nil
	}
	out := &krm.DeviceDetails{}
	out.PaymentApp = direct.LazyPtr(in.GetPaymentApp())
	out.Capability = direct.LazyPtr(in.GetCapability())
	out.GeoCode = LatLng_FromProto(mapCtx, in.GetGeoCode())
	out.ID = direct.LazyPtr(in.GetId())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.OperatingSystem = direct.LazyPtr(in.GetOperatingSystem())
	out.TelecomProvider = direct.LazyPtr(in.GetTelecomProvider())
	out.Type = direct.LazyPtr(in.GetType())
	return out
}
func DeviceDetails_ToProto(mapCtx *direct.MapContext, in *krm.DeviceDetails) *pb.DeviceDetails {
	if in == nil {
		return nil
	}
	out := &pb.DeviceDetails{}
	out.PaymentApp = direct.ValueOf(in.PaymentApp)
	out.Capability = direct.ValueOf(in.Capability)
	out.GeoCode = LatLng_ToProto(mapCtx, in.GeoCode)
	out.Id = direct.ValueOf(in.ID)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Location = direct.ValueOf(in.Location)
	out.OperatingSystem = direct.ValueOf(in.OperatingSystem)
	out.TelecomProvider = direct.ValueOf(in.TelecomProvider)
	out.Type = direct.ValueOf(in.Type)
	return out
}
func Dispute_FromProto(mapCtx *direct.MapContext, in *pb.Dispute) *krm.Dispute {
	if in == nil {
		return nil
	}
	out := &krm.Dispute{}
	out.Name = direct.LazyPtr(in.GetName())
	out.RaiseDisputeAdjustment = RaiseDisputeAdjustment_FromProto(mapCtx, in.GetRaiseDisputeAdjustment())
	out.Details = CaseDetails_FromProto(mapCtx, in.GetDetails())
	// MISSING: Response
	out.ResolveDisputeAdjustment = ResolveDisputeAdjustment_FromProto(mapCtx, in.GetResolveDisputeAdjustment())
	return out
}
func Dispute_ToProto(mapCtx *direct.MapContext, in *krm.Dispute) *pb.Dispute {
	if in == nil {
		return nil
	}
	out := &pb.Dispute{}
	out.Name = direct.ValueOf(in.Name)
	out.RaiseDisputeAdjustment = RaiseDisputeAdjustment_ToProto(mapCtx, in.RaiseDisputeAdjustment)
	out.Details = CaseDetails_ToProto(mapCtx, in.Details)
	// MISSING: Response
	out.ResolveDisputeAdjustment = ResolveDisputeAdjustment_ToProto(mapCtx, in.ResolveDisputeAdjustment)
	return out
}
func DisputeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dispute) *krm.DisputeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DisputeObservedState{}
	// MISSING: Name
	// MISSING: RaiseDisputeAdjustment
	// MISSING: Details
	out.Response = CaseResponse_FromProto(mapCtx, in.GetResponse())
	// MISSING: ResolveDisputeAdjustment
	return out
}
func DisputeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DisputeObservedState) *pb.Dispute {
	if in == nil {
		return nil
	}
	out := &pb.Dispute{}
	// MISSING: Name
	// MISSING: RaiseDisputeAdjustment
	// MISSING: Details
	out.Response = CaseResponse_ToProto(mapCtx, in.Response)
	// MISSING: ResolveDisputeAdjustment
	return out
}
func OriginalTransaction_FromProto(mapCtx *direct.MapContext, in *pb.OriginalTransaction) *krm.OriginalTransaction {
	if in == nil {
		return nil
	}
	out := &krm.OriginalTransaction{}
	out.TransactionID = direct.LazyPtr(in.GetTransactionId())
	out.RetrievalReferenceNumber = direct.LazyPtr(in.GetRetrievalReferenceNumber())
	out.RequestTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRequestTime())
	return out
}
func OriginalTransaction_ToProto(mapCtx *direct.MapContext, in *krm.OriginalTransaction) *pb.OriginalTransaction {
	if in == nil {
		return nil
	}
	out := &pb.OriginalTransaction{}
	out.TransactionId = direct.ValueOf(in.TransactionID)
	out.RetrievalReferenceNumber = direct.ValueOf(in.RetrievalReferenceNumber)
	out.RequestTime = direct.StringTimestamp_ToProto(mapCtx, in.RequestTime)
	return out
}
func Participant_FromProto(mapCtx *direct.MapContext, in *pb.Participant) *krm.Participant {
	if in == nil {
		return nil
	}
	out := &krm.Participant{}
	out.PaymentAddress = direct.LazyPtr(in.GetPaymentAddress())
	out.Persona = direct.Enum_FromProto(mapCtx, in.GetPersona())
	out.User = direct.LazyPtr(in.GetUser())
	// MISSING: Account
	// MISSING: DeviceDetails
	// MISSING: MobileNumber
	return out
}
func Participant_ToProto(mapCtx *direct.MapContext, in *krm.Participant) *pb.Participant {
	if in == nil {
		return nil
	}
	out := &pb.Participant{}
	out.PaymentAddress = direct.ValueOf(in.PaymentAddress)
	out.Persona = direct.Enum_ToProto[pb.Participant_Persona](mapCtx, in.Persona)
	out.User = direct.ValueOf(in.User)
	// MISSING: Account
	// MISSING: DeviceDetails
	// MISSING: MobileNumber
	return out
}
func ParticipantObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Participant) *krm.ParticipantObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ParticipantObservedState{}
	// MISSING: PaymentAddress
	// MISSING: Persona
	// MISSING: User
	out.Account = AccountReference_FromProto(mapCtx, in.GetAccount())
	out.DeviceDetails = DeviceDetails_FromProto(mapCtx, in.GetDeviceDetails())
	out.MobileNumber = direct.LazyPtr(in.GetMobileNumber())
	return out
}
func ParticipantObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ParticipantObservedState) *pb.Participant {
	if in == nil {
		return nil
	}
	out := &pb.Participant{}
	// MISSING: PaymentAddress
	// MISSING: Persona
	// MISSING: User
	out.Account = AccountReference_ToProto(mapCtx, in.Account)
	out.DeviceDetails = DeviceDetails_ToProto(mapCtx, in.DeviceDetails)
	out.MobileNumber = direct.ValueOf(in.MobileNumber)
	return out
}
func PaymentgatewayDisputeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dispute) *krm.PaymentgatewayDisputeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayDisputeObservedState{}
	// MISSING: Name
	// MISSING: RaiseDisputeAdjustment
	// MISSING: Details
	// MISSING: Response
	// MISSING: ResolveDisputeAdjustment
	return out
}
func PaymentgatewayDisputeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayDisputeObservedState) *pb.Dispute {
	if in == nil {
		return nil
	}
	out := &pb.Dispute{}
	// MISSING: Name
	// MISSING: RaiseDisputeAdjustment
	// MISSING: Details
	// MISSING: Response
	// MISSING: ResolveDisputeAdjustment
	return out
}
func PaymentgatewayDisputeSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dispute) *krm.PaymentgatewayDisputeSpec {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayDisputeSpec{}
	// MISSING: Name
	// MISSING: RaiseDisputeAdjustment
	// MISSING: Details
	// MISSING: Response
	// MISSING: ResolveDisputeAdjustment
	return out
}
func PaymentgatewayDisputeSpec_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayDisputeSpec) *pb.Dispute {
	if in == nil {
		return nil
	}
	out := &pb.Dispute{}
	// MISSING: Name
	// MISSING: RaiseDisputeAdjustment
	// MISSING: Details
	// MISSING: Response
	// MISSING: ResolveDisputeAdjustment
	return out
}
func RaiseDisputeAdjustment_FromProto(mapCtx *direct.MapContext, in *pb.RaiseDisputeAdjustment) *krm.RaiseDisputeAdjustment {
	if in == nil {
		return nil
	}
	out := &krm.RaiseDisputeAdjustment{}
	out.AdjustmentFlag = direct.Enum_FromProto(mapCtx, in.GetAdjustmentFlag())
	out.AdjustmentCode = direct.Enum_FromProto(mapCtx, in.GetAdjustmentCode())
	return out
}
func RaiseDisputeAdjustment_ToProto(mapCtx *direct.MapContext, in *krm.RaiseDisputeAdjustment) *pb.RaiseDisputeAdjustment {
	if in == nil {
		return nil
	}
	out := &pb.RaiseDisputeAdjustment{}
	out.AdjustmentFlag = direct.Enum_ToProto[pb.RaiseDisputeAdjustment_AdjustmentFlag](mapCtx, in.AdjustmentFlag)
	out.AdjustmentCode = direct.Enum_ToProto[pb.RaiseDisputeAdjustment_ReasonCode](mapCtx, in.AdjustmentCode)
	return out
}
func ResolveDisputeAdjustment_FromProto(mapCtx *direct.MapContext, in *pb.ResolveDisputeAdjustment) *krm.ResolveDisputeAdjustment {
	if in == nil {
		return nil
	}
	out := &krm.ResolveDisputeAdjustment{}
	out.AdjustmentFlag = direct.Enum_FromProto(mapCtx, in.GetAdjustmentFlag())
	out.AdjustmentCode = direct.Enum_FromProto(mapCtx, in.GetAdjustmentCode())
	return out
}
func ResolveDisputeAdjustment_ToProto(mapCtx *direct.MapContext, in *krm.ResolveDisputeAdjustment) *pb.ResolveDisputeAdjustment {
	if in == nil {
		return nil
	}
	out := &pb.ResolveDisputeAdjustment{}
	out.AdjustmentFlag = direct.Enum_ToProto[pb.ResolveDisputeAdjustment_AdjustmentFlag](mapCtx, in.AdjustmentFlag)
	out.AdjustmentCode = direct.Enum_ToProto[pb.ResolveDisputeAdjustment_ReasonCode](mapCtx, in.AdjustmentCode)
	return out
}
