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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/paymentgateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/paymentgateway/issuerswitch/apiv1/issuerswitchpb"
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
func MetadataTransaction_FromProto(mapCtx *direct.MapContext, in *pb.MetadataTransaction) *krm.MetadataTransaction {
	if in == nil {
		return nil
	}
	out := &krm.MetadataTransaction{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Info = TransactionInfo_FromProto(mapCtx, in.GetInfo())
	// MISSING: Initiator
	return out
}
func MetadataTransaction_ToProto(mapCtx *direct.MapContext, in *krm.MetadataTransaction) *pb.MetadataTransaction {
	if in == nil {
		return nil
	}
	out := &pb.MetadataTransaction{}
	out.Name = direct.ValueOf(in.Name)
	out.Info = TransactionInfo_ToProto(mapCtx, in.Info)
	// MISSING: Initiator
	return out
}
func MetadataTransactionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataTransaction) *krm.MetadataTransactionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataTransactionObservedState{}
	// MISSING: Name
	out.Info = TransactionInfoObservedState_FromProto(mapCtx, in.GetInfo())
	out.Initiator = Participant_FromProto(mapCtx, in.GetInitiator())
	return out
}
func MetadataTransactionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MetadataTransactionObservedState) *pb.MetadataTransaction {
	if in == nil {
		return nil
	}
	out := &pb.MetadataTransaction{}
	// MISSING: Name
	out.Info = TransactionInfoObservedState_ToProto(mapCtx, in.Info)
	out.Initiator = Participant_ToProto(mapCtx, in.Initiator)
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
func PaymentgatewayMetadataTransactionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MetadataTransaction) *krm.PaymentgatewayMetadataTransactionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayMetadataTransactionObservedState{}
	// MISSING: Name
	// MISSING: Info
	// MISSING: Initiator
	return out
}
func PaymentgatewayMetadataTransactionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayMetadataTransactionObservedState) *pb.MetadataTransaction {
	if in == nil {
		return nil
	}
	out := &pb.MetadataTransaction{}
	// MISSING: Name
	// MISSING: Info
	// MISSING: Initiator
	return out
}
func PaymentgatewayMetadataTransactionSpec_FromProto(mapCtx *direct.MapContext, in *pb.MetadataTransaction) *krm.PaymentgatewayMetadataTransactionSpec {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayMetadataTransactionSpec{}
	// MISSING: Name
	// MISSING: Info
	// MISSING: Initiator
	return out
}
func PaymentgatewayMetadataTransactionSpec_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayMetadataTransactionSpec) *pb.MetadataTransaction {
	if in == nil {
		return nil
	}
	out := &pb.MetadataTransaction{}
	// MISSING: Name
	// MISSING: Info
	// MISSING: Initiator
	return out
}
func TransactionInfo_FromProto(mapCtx *direct.MapContext, in *pb.TransactionInfo) *krm.TransactionInfo {
	if in == nil {
		return nil
	}
	out := &krm.TransactionInfo{}
	// MISSING: ID
	// MISSING: ApiType
	// MISSING: TransactionType
	// MISSING: TransactionSubType
	// MISSING: State
	out.Metadata = TransactionInfo_TransactionMetadata_FromProto(mapCtx, in.GetMetadata())
	// MISSING: ErrorDetails
	// MISSING: AdapterInfo
	out.RiskInfo = direct.Slice_FromProto(mapCtx, in.RiskInfo, TransactionInfo_TransactionRiskInfo_FromProto)
	return out
}
func TransactionInfo_ToProto(mapCtx *direct.MapContext, in *krm.TransactionInfo) *pb.TransactionInfo {
	if in == nil {
		return nil
	}
	out := &pb.TransactionInfo{}
	// MISSING: ID
	// MISSING: ApiType
	// MISSING: TransactionType
	// MISSING: TransactionSubType
	// MISSING: State
	out.Metadata = TransactionInfo_TransactionMetadata_ToProto(mapCtx, in.Metadata)
	// MISSING: ErrorDetails
	// MISSING: AdapterInfo
	out.RiskInfo = direct.Slice_ToProto(mapCtx, in.RiskInfo, TransactionInfo_TransactionRiskInfo_ToProto)
	return out
}
func TransactionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransactionInfo) *krm.TransactionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TransactionInfoObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	out.ApiType = direct.Enum_FromProto(mapCtx, in.GetApiType())
	out.TransactionType = direct.Enum_FromProto(mapCtx, in.GetTransactionType())
	out.TransactionSubType = direct.Enum_FromProto(mapCtx, in.GetTransactionSubType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Metadata = TransactionInfo_TransactionMetadataObservedState_FromProto(mapCtx, in.GetMetadata())
	out.ErrorDetails = TransactionInfo_TransactionErrorDetails_FromProto(mapCtx, in.GetErrorDetails())
	out.AdapterInfo = TransactionInfo_AdapterInfo_FromProto(mapCtx, in.GetAdapterInfo())
	// MISSING: RiskInfo
	return out
}
func TransactionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TransactionInfoObservedState) *pb.TransactionInfo {
	if in == nil {
		return nil
	}
	out := &pb.TransactionInfo{}
	out.Id = direct.ValueOf(in.ID)
	out.ApiType = direct.Enum_ToProto[pb.ApiType](mapCtx, in.ApiType)
	out.TransactionType = direct.Enum_ToProto[pb.TransactionType](mapCtx, in.TransactionType)
	out.TransactionSubType = direct.Enum_ToProto[pb.TransactionInfo_TransactionSubType](mapCtx, in.TransactionSubType)
	out.State = direct.Enum_ToProto[pb.TransactionInfo_State](mapCtx, in.State)
	out.Metadata = TransactionInfo_TransactionMetadataObservedState_ToProto(mapCtx, in.Metadata)
	out.ErrorDetails = TransactionInfo_TransactionErrorDetails_ToProto(mapCtx, in.ErrorDetails)
	out.AdapterInfo = TransactionInfo_AdapterInfo_ToProto(mapCtx, in.AdapterInfo)
	// MISSING: RiskInfo
	return out
}
func TransactionInfo_AdapterInfo_FromProto(mapCtx *direct.MapContext, in *pb.TransactionInfo_AdapterInfo) *krm.TransactionInfo_AdapterInfo {
	if in == nil {
		return nil
	}
	out := &krm.TransactionInfo_AdapterInfo{}
	// MISSING: RequestIds
	// MISSING: ResponseMetadata
	return out
}
func TransactionInfo_AdapterInfo_ToProto(mapCtx *direct.MapContext, in *krm.TransactionInfo_AdapterInfo) *pb.TransactionInfo_AdapterInfo {
	if in == nil {
		return nil
	}
	out := &pb.TransactionInfo_AdapterInfo{}
	// MISSING: RequestIds
	// MISSING: ResponseMetadata
	return out
}
func TransactionInfo_AdapterInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransactionInfo_AdapterInfo) *krm.TransactionInfo_AdapterInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TransactionInfo_AdapterInfoObservedState{}
	out.RequestIds = direct.LazyPtr(in.GetRequestIds())
	out.ResponseMetadata = TransactionInfo_AdapterInfo_ResponseMetadata_FromProto(mapCtx, in.GetResponseMetadata())
	return out
}
func TransactionInfo_AdapterInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TransactionInfo_AdapterInfoObservedState) *pb.TransactionInfo_AdapterInfo {
	if in == nil {
		return nil
	}
	out := &pb.TransactionInfo_AdapterInfo{}
	out.RequestIds = direct.ValueOf(in.RequestIds)
	out.ResponseMetadata = TransactionInfo_AdapterInfo_ResponseMetadata_ToProto(mapCtx, in.ResponseMetadata)
	return out
}
func TransactionInfo_TransactionErrorDetails_FromProto(mapCtx *direct.MapContext, in *pb.TransactionInfo_TransactionErrorDetails) *krm.TransactionInfo_TransactionErrorDetails {
	if in == nil {
		return nil
	}
	out := &krm.TransactionInfo_TransactionErrorDetails{}
	// MISSING: ErrorCode
	// MISSING: ErrorMessage
	// MISSING: UpiErrorCode
	return out
}
func TransactionInfo_TransactionErrorDetails_ToProto(mapCtx *direct.MapContext, in *krm.TransactionInfo_TransactionErrorDetails) *pb.TransactionInfo_TransactionErrorDetails {
	if in == nil {
		return nil
	}
	out := &pb.TransactionInfo_TransactionErrorDetails{}
	// MISSING: ErrorCode
	// MISSING: ErrorMessage
	// MISSING: UpiErrorCode
	return out
}
func TransactionInfo_TransactionErrorDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransactionInfo_TransactionErrorDetails) *krm.TransactionInfo_TransactionErrorDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TransactionInfo_TransactionErrorDetailsObservedState{}
	out.ErrorCode = direct.LazyPtr(in.GetErrorCode())
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	out.UpiErrorCode = direct.LazyPtr(in.GetUpiErrorCode())
	return out
}
func TransactionInfo_TransactionErrorDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TransactionInfo_TransactionErrorDetailsObservedState) *pb.TransactionInfo_TransactionErrorDetails {
	if in == nil {
		return nil
	}
	out := &pb.TransactionInfo_TransactionErrorDetails{}
	out.ErrorCode = direct.ValueOf(in.ErrorCode)
	out.ErrorMessage = direct.ValueOf(in.ErrorMessage)
	out.UpiErrorCode = direct.ValueOf(in.UpiErrorCode)
	return out
}
func TransactionInfo_TransactionRiskInfo_FromProto(mapCtx *direct.MapContext, in *pb.TransactionInfo_TransactionRiskInfo) *krm.TransactionInfo_TransactionRiskInfo {
	if in == nil {
		return nil
	}
	out := &krm.TransactionInfo_TransactionRiskInfo{}
	out.Provider = direct.LazyPtr(in.GetProvider())
	out.Type = direct.LazyPtr(in.GetType())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func TransactionInfo_TransactionRiskInfo_ToProto(mapCtx *direct.MapContext, in *krm.TransactionInfo_TransactionRiskInfo) *pb.TransactionInfo_TransactionRiskInfo {
	if in == nil {
		return nil
	}
	out := &pb.TransactionInfo_TransactionRiskInfo{}
	out.Provider = direct.ValueOf(in.Provider)
	out.Type = direct.ValueOf(in.Type)
	out.Value = direct.ValueOf(in.Value)
	return out
}
