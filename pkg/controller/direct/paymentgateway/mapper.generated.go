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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/paymentgateway/issuerswitch/accountmanager/apiv1/accountmanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/paymentgateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AccountManagerMerchantInfo_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerMerchantInfo) *krm.AccountManagerMerchantInfo {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerMerchantInfo{}
	out.CategoryCode = direct.LazyPtr(in.GetCategoryCode())
	out.ID = direct.LazyPtr(in.GetId())
	return out
}
func AccountManagerMerchantInfo_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerMerchantInfo) *pb.AccountManagerMerchantInfo {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerMerchantInfo{}
	out.CategoryCode = direct.ValueOf(in.CategoryCode)
	out.Id = direct.ValueOf(in.ID)
	return out
}
func AccountManagerParticipant_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerParticipant) *krm.AccountManagerParticipant {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerParticipant{}
	out.PaymentAddress = direct.LazyPtr(in.GetPaymentAddress())
	out.Persona = direct.Enum_FromProto(mapCtx, in.GetPersona())
	out.Account = AccountReference_FromProto(mapCtx, in.GetAccount())
	return out
}
func AccountManagerParticipant_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerParticipant) *pb.AccountManagerParticipant {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerParticipant{}
	out.PaymentAddress = direct.ValueOf(in.PaymentAddress)
	out.Persona = direct.Enum_ToProto[pb.AccountManagerParticipant_Persona](mapCtx, in.Persona)
	out.Account = AccountReference_ToProto(mapCtx, in.Account)
	return out
}
func AccountManagerParticipantObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerParticipant) *krm.AccountManagerParticipantObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerParticipantObservedState{}
	// MISSING: PaymentAddress
	// MISSING: Persona
	out.Account = AccountReferenceObservedState_FromProto(mapCtx, in.GetAccount())
	return out
}
func AccountManagerParticipantObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerParticipantObservedState) *pb.AccountManagerParticipant {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerParticipant{}
	// MISSING: PaymentAddress
	// MISSING: Persona
	out.Account = AccountReferenceObservedState_ToProto(mapCtx, in.Account)
	return out
}
func AccountManagerSettlementParticipant_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerSettlementParticipant) *krm.AccountManagerSettlementParticipant {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerSettlementParticipant{}
	out.Participant = AccountManagerParticipant_FromProto(mapCtx, in.GetParticipant())
	out.MerchantInfo = AccountManagerMerchantInfo_FromProto(mapCtx, in.GetMerchantInfo())
	return out
}
func AccountManagerSettlementParticipant_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerSettlementParticipant) *pb.AccountManagerSettlementParticipant {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerSettlementParticipant{}
	out.Participant = AccountManagerParticipant_ToProto(mapCtx, in.Participant)
	out.MerchantInfo = AccountManagerMerchantInfo_ToProto(mapCtx, in.MerchantInfo)
	return out
}
func AccountManagerSettlementParticipantObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerSettlementParticipant) *krm.AccountManagerSettlementParticipantObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerSettlementParticipantObservedState{}
	out.Participant = AccountManagerParticipantObservedState_FromProto(mapCtx, in.GetParticipant())
	// MISSING: MerchantInfo
	return out
}
func AccountManagerSettlementParticipantObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerSettlementParticipantObservedState) *pb.AccountManagerSettlementParticipant {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerSettlementParticipant{}
	out.Participant = AccountManagerParticipantObservedState_ToProto(mapCtx, in.Participant)
	// MISSING: MerchantInfo
	return out
}
func AccountManagerTransaction_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransaction) *krm.AccountManagerTransaction {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerTransaction{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AccountID = direct.LazyPtr(in.GetAccountId())
	out.Info = AccountManagerTransactionInfo_FromProto(mapCtx, in.GetInfo())
	out.Payer = AccountManagerSettlementParticipant_FromProto(mapCtx, in.GetPayer())
	out.Payee = AccountManagerSettlementParticipant_FromProto(mapCtx, in.GetPayee())
	out.ReconciliationInfo = AccountManagerTransactionReconciliationInfo_FromProto(mapCtx, in.GetReconciliationInfo())
	out.Amount = Money_FromProto(mapCtx, in.GetAmount())
	return out
}
func AccountManagerTransaction_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerTransaction) *pb.AccountManagerTransaction {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransaction{}
	out.Name = direct.ValueOf(in.Name)
	out.AccountId = direct.ValueOf(in.AccountID)
	out.Info = AccountManagerTransactionInfo_ToProto(mapCtx, in.Info)
	out.Payer = AccountManagerSettlementParticipant_ToProto(mapCtx, in.Payer)
	out.Payee = AccountManagerSettlementParticipant_ToProto(mapCtx, in.Payee)
	out.ReconciliationInfo = AccountManagerTransactionReconciliationInfo_ToProto(mapCtx, in.ReconciliationInfo)
	out.Amount = Money_ToProto(mapCtx, in.Amount)
	return out
}
func AccountManagerTransactionInfo_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransactionInfo) *krm.AccountManagerTransactionInfo {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerTransactionInfo{}
	out.ID = direct.LazyPtr(in.GetId())
	out.TransactionType = direct.Enum_FromProto(mapCtx, in.GetTransactionType())
	// MISSING: State
	out.Metadata = AccountManagerTransactionInfo_AccountManagerTransactionMetadata_FromProto(mapCtx, in.GetMetadata())
	// MISSING: ErrorDetails
	return out
}
func AccountManagerTransactionInfo_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerTransactionInfo) *pb.AccountManagerTransactionInfo {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransactionInfo{}
	out.Id = direct.ValueOf(in.ID)
	out.TransactionType = direct.Enum_ToProto[pb.AccountManagerTransactionType](mapCtx, in.TransactionType)
	// MISSING: State
	out.Metadata = AccountManagerTransactionInfo_AccountManagerTransactionMetadata_ToProto(mapCtx, in.Metadata)
	// MISSING: ErrorDetails
	return out
}
func AccountManagerTransactionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransactionInfo) *krm.AccountManagerTransactionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerTransactionInfoObservedState{}
	// MISSING: ID
	// MISSING: TransactionType
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Metadata = AccountManagerTransactionInfo_AccountManagerTransactionMetadataObservedState_FromProto(mapCtx, in.GetMetadata())
	out.ErrorDetails = AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails_FromProto(mapCtx, in.GetErrorDetails())
	return out
}
func AccountManagerTransactionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerTransactionInfoObservedState) *pb.AccountManagerTransactionInfo {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransactionInfo{}
	// MISSING: ID
	// MISSING: TransactionType
	out.State = direct.Enum_ToProto[pb.AccountManagerTransactionInfo_State](mapCtx, in.State)
	out.Metadata = AccountManagerTransactionInfo_AccountManagerTransactionMetadataObservedState_ToProto(mapCtx, in.Metadata)
	out.ErrorDetails = AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails_ToProto(mapCtx, in.ErrorDetails)
	return out
}
func AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails) *krm.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails{}
	// MISSING: ErrorCode
	// MISSING: ErrorMessage
	return out
}
func AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails) *pb.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails{}
	// MISSING: ErrorCode
	// MISSING: ErrorMessage
	return out
}
func AccountManagerTransactionInfo_AccountManagerTransactionErrorDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails) *krm.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetailsObservedState{}
	out.ErrorCode = direct.LazyPtr(in.GetErrorCode())
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	return out
}
func AccountManagerTransactionInfo_AccountManagerTransactionErrorDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetailsObservedState) *pb.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails{}
	out.ErrorCode = direct.ValueOf(in.ErrorCode)
	out.ErrorMessage = direct.ValueOf(in.ErrorMessage)
	return out
}
func AccountManagerTransactionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransaction) *krm.AccountManagerTransactionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerTransactionObservedState{}
	// MISSING: Name
	// MISSING: AccountID
	out.Info = AccountManagerTransactionInfoObservedState_FromProto(mapCtx, in.GetInfo())
	out.Payer = AccountManagerSettlementParticipantObservedState_FromProto(mapCtx, in.GetPayer())
	// MISSING: Payee
	out.ReconciliationInfo = AccountManagerTransactionReconciliationInfoObservedState_FromProto(mapCtx, in.GetReconciliationInfo())
	// MISSING: Amount
	return out
}
func AccountManagerTransactionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerTransactionObservedState) *pb.AccountManagerTransaction {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransaction{}
	// MISSING: Name
	// MISSING: AccountID
	out.Info = AccountManagerTransactionInfoObservedState_ToProto(mapCtx, in.Info)
	out.Payer = AccountManagerSettlementParticipantObservedState_ToProto(mapCtx, in.Payer)
	// MISSING: Payee
	out.ReconciliationInfo = AccountManagerTransactionReconciliationInfoObservedState_ToProto(mapCtx, in.ReconciliationInfo)
	// MISSING: Amount
	return out
}
func AccountManagerTransactionReconciliationInfo_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransactionReconciliationInfo) *krm.AccountManagerTransactionReconciliationInfo {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerTransactionReconciliationInfo{}
	// MISSING: State
	out.ReconciliationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetReconciliationTime())
	return out
}
func AccountManagerTransactionReconciliationInfo_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerTransactionReconciliationInfo) *pb.AccountManagerTransactionReconciliationInfo {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransactionReconciliationInfo{}
	// MISSING: State
	out.ReconciliationTime = direct.StringTimestamp_ToProto(mapCtx, in.ReconciliationTime)
	return out
}
func AccountManagerTransactionReconciliationInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransactionReconciliationInfo) *krm.AccountManagerTransactionReconciliationInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccountManagerTransactionReconciliationInfoObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: ReconciliationTime
	return out
}
func AccountManagerTransactionReconciliationInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccountManagerTransactionReconciliationInfoObservedState) *pb.AccountManagerTransactionReconciliationInfo {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransactionReconciliationInfo{}
	out.State = direct.Enum_ToProto[pb.AccountManagerTransactionReconciliationInfo_ReconciliationState](mapCtx, in.State)
	// MISSING: ReconciliationTime
	return out
}
func PaymentgatewayAccountManagerTransactionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransaction) *krm.PaymentgatewayAccountManagerTransactionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayAccountManagerTransactionObservedState{}
	// MISSING: Name
	// MISSING: AccountID
	// MISSING: Info
	// MISSING: Payer
	// MISSING: Payee
	// MISSING: ReconciliationInfo
	// MISSING: Amount
	return out
}
func PaymentgatewayAccountManagerTransactionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayAccountManagerTransactionObservedState) *pb.AccountManagerTransaction {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransaction{}
	// MISSING: Name
	// MISSING: AccountID
	// MISSING: Info
	// MISSING: Payer
	// MISSING: Payee
	// MISSING: ReconciliationInfo
	// MISSING: Amount
	return out
}
func PaymentgatewayAccountManagerTransactionSpec_FromProto(mapCtx *direct.MapContext, in *pb.AccountManagerTransaction) *krm.PaymentgatewayAccountManagerTransactionSpec {
	if in == nil {
		return nil
	}
	out := &krm.PaymentgatewayAccountManagerTransactionSpec{}
	// MISSING: Name
	// MISSING: AccountID
	// MISSING: Info
	// MISSING: Payer
	// MISSING: Payee
	// MISSING: ReconciliationInfo
	// MISSING: Amount
	return out
}
func PaymentgatewayAccountManagerTransactionSpec_ToProto(mapCtx *direct.MapContext, in *krm.PaymentgatewayAccountManagerTransactionSpec) *pb.AccountManagerTransaction {
	if in == nil {
		return nil
	}
	out := &pb.AccountManagerTransaction{}
	// MISSING: Name
	// MISSING: AccountID
	// MISSING: Info
	// MISSING: Payer
	// MISSING: Payee
	// MISSING: ReconciliationInfo
	// MISSING: Amount
	return out
}
