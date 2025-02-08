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

package v1alpha1


// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerMerchantInfo
type AccountManagerMerchantInfo struct {
	// Merchant Category Code (MCC) as specified by UPI. This is a four-digit
	//  number listed in ISO 18245 for retail financial services.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerMerchantInfo.category_code
	CategoryCode *string `json:"categoryCode,omitempty"`

	// ID of the merchant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerMerchantInfo.id
	ID *string `json:"id,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerParticipant
type AccountManagerParticipant struct {
	// The payment address of the participant. In the UPI system, this will be the
	//  virtual payment address (VPA) of the participant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerParticipant.payment_address
	PaymentAddress *string `json:"paymentAddress,omitempty"`

	// The persona of the participant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerParticipant.persona
	Persona *string `json:"persona,omitempty"`

	// Unique identification of an account.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerParticipant.account
	Account *AccountReference `json:"account,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerSettlementParticipant
type AccountManagerSettlementParticipant struct {
	// The participant information.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerSettlementParticipant.participant
	Participant *AccountManagerParticipant `json:"participant,omitempty"`

	// Information about a merchant who is a participant in the payment. This
	//  field will be specified only if the participant is a merchant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerSettlementParticipant.merchant_info
	MerchantInfo *AccountManagerMerchantInfo `json:"merchantInfo,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction
type AccountManagerTransaction struct {
	// The name of the transaction. This uniquely identifies the
	//  transaction. Format of name is
	//  projects/{project}/accountManagers/{account_manager}/transactions/{transaction}.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.name
	Name *string `json:"name,omitempty"`

	// The account ID for which the transaction was processed.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.account_id
	AccountID *string `json:"accountID,omitempty"`

	// Information about the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.info
	Info *AccountManagerTransactionInfo `json:"info,omitempty"`

	// The payer in the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.payer
	Payer *AccountManagerSettlementParticipant `json:"payer,omitempty"`

	// The payee in the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.payee
	Payee *AccountManagerSettlementParticipant `json:"payee,omitempty"`

	// Reconciliation information for the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.reconciliation_info
	ReconciliationInfo *AccountManagerTransactionReconciliationInfo `json:"reconciliationInfo,omitempty"`

	// The amount for payment settlement in the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.amount
	Amount *Money `json:"amount,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo
type AccountManagerTransactionInfo struct {
	// An identifier that is mandatorily present in every transaction processed
	//  via account manager.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.id
	ID *string `json:"id,omitempty"`

	// The transaction type.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.transaction_type
	TransactionType *string `json:"transactionType,omitempty"`

	// Metadata about the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.metadata
	Metadata *AccountManagerTransactionInfo_AccountManagerTransactionMetadata `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionErrorDetails
type AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails struct {
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionMetadata
type AccountManagerTransactionInfo_AccountManagerTransactionMetadata struct {
	// The time at which the transaction took place.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionMetadata.transaction_time
	TransactionTime *string `json:"transactionTime,omitempty"`

	// Retrieval reference number (RRN) for the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionMetadata.retrieval_reference_number
	RetrievalReferenceNumber *string `json:"retrievalReferenceNumber,omitempty"`

	// The initiation mode of this transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionMetadata.initiation_mode
	InitiationMode *string `json:"initiationMode,omitempty"`

	// The purpose code of this transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionMetadata.purpose_code
	PurposeCode *string `json:"purposeCode,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionReconciliationInfo
type AccountManagerTransactionReconciliationInfo struct {

	// Time at which reconciliation was performed for the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionReconciliationInfo.reconciliation_time
	ReconciliationTime *string `json:"reconciliationTime,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReference struct {
	// IFSC of the account's bank branch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.ifsc
	Ifsc *string `json:"ifsc,omitempty"`

	// Unique number for an account in a bank and branch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_number
	AccountNumber *string `json:"accountNumber,omitempty"`
}

// +kcc:proto=google.type.Money
type Money struct {
	// The three-letter currency code defined in ISO 4217.
	// +kcc:proto:field=google.type.Money.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// The whole units of the amount.
	//  For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	// +kcc:proto:field=google.type.Money.units
	Units *int64 `json:"units,omitempty"`

	// Number of nano (10^-9) units of the amount.
	//  The value must be between -999,999,999 and +999,999,999 inclusive.
	//  If `units` is positive, `nanos` must be positive or zero.
	//  If `units` is zero, `nanos` can be positive, zero, or negative.
	//  If `units` is negative, `nanos` must be negative or zero.
	//  For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	// +kcc:proto:field=google.type.Money.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerParticipant
type AccountManagerParticipantObservedState struct {
	// Unique identification of an account.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerParticipant.account
	Account *AccountReferenceObservedState `json:"account,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerSettlementParticipant
type AccountManagerSettlementParticipantObservedState struct {
	// The participant information.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerSettlementParticipant.participant
	Participant *AccountManagerParticipantObservedState `json:"participant,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction
type AccountManagerTransactionObservedState struct {
	// Information about the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.info
	Info *AccountManagerTransactionInfoObservedState `json:"info,omitempty"`

	// The payer in the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.payer
	Payer *AccountManagerSettlementParticipantObservedState `json:"payer,omitempty"`

	// Reconciliation information for the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransaction.reconciliation_info
	ReconciliationInfo *AccountManagerTransactionReconciliationInfoObservedState `json:"reconciliationInfo,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo
type AccountManagerTransactionInfoObservedState struct {
	// Output only. The transaction's state.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.state
	State *string `json:"state,omitempty"`

	// Metadata about the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.metadata
	Metadata *AccountManagerTransactionInfo_AccountManagerTransactionMetadataObservedState `json:"metadata,omitempty"`

	// Output only. Any error details for the current transaction, if the state is
	//  `FAILED`.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.error_details
	ErrorDetails *AccountManagerTransactionInfo_AccountManagerTransactionErrorDetails `json:"errorDetails,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionErrorDetails
type AccountManagerTransactionInfo_AccountManagerTransactionErrorDetailsObservedState struct {
	// Output only. Error code of the failed transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionErrorDetails.error_code
	ErrorCode *string `json:"errorCode,omitempty"`

	// Output only. Error description for the failed transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionErrorDetails.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionMetadata
type AccountManagerTransactionInfo_AccountManagerTransactionMetadataObservedState struct {
	// Output only. The time at which the transaction resource was created by
	//  the account manager.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionMetadata.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the transaction resource was last updated
	//  by the account manager.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionInfo.AccountManagerTransactionMetadata.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionReconciliationInfo
type AccountManagerTransactionReconciliationInfoObservedState struct {
	// Output only. State of reconciliation.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.accountmanager.v1.AccountManagerTransactionReconciliationInfo.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReferenceObservedState struct {
	// Output only. Type of account. Examples include SAVINGS, CURRENT, etc.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_type
	AccountType *string `json:"accountType,omitempty"`
}
