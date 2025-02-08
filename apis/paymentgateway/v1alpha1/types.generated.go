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


// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReference struct {
	// IFSC of the account's bank branch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.ifsc
	Ifsc *string `json:"ifsc,omitempty"`

	// Unique number for an account in a bank and branch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_number
	AccountNumber *string `json:"accountNumber,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.CaseDetails
type CaseDetails struct {
	// Required. Details of original transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseDetails.original_transaction
	OriginalTransaction *OriginalTransaction `json:"originalTransaction,omitempty"`

	// Required. Initiator of the complaint / dispute.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseDetails.transaction_sub_type
	TransactionSubType *string `json:"transactionSubType,omitempty"`

	// Required. The adjustment amount in URCS for the complaint / dispute. This
	//  maps to `reqAdjAmount` in complaint request.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseDetails.amount
	Amount *Money `json:"amount,omitempty"`

	// The original response code which has been updated in the complaint
	//  Response. This should map to settlement response code currently available
	//  in URCS system.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseDetails.original_settlement_response_code
	OriginalSettlementResponseCode *string `json:"originalSettlementResponseCode,omitempty"`

	// Required. Set to true if the complaint / dispute belongs to current
	//  settlement cycle, false otherwise.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseDetails.current_cycle
	CurrentCycle *bool `json:"currentCycle,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse
type CaseResponse struct {
	// Complaint Reference Number(CRN) sent by UPI as a reference against the
	//  generated complaint / dispute.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.complaint_reference_number
	ComplaintReferenceNumber *string `json:"complaintReferenceNumber,omitempty"`

	// The adjustment amount of the response. This maps to `adjAmt` in
	//  complaint response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.amount
	Amount *Money `json:"amount,omitempty"`

	// The adjustment flag in response to the complaint. This maps adjustment flag
	//  in URCS for the complaint transaction to `Resp.Ref.adjFlag` in complaint
	//  response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.adjustment_flag
	AdjustmentFlag *string `json:"adjustmentFlag,omitempty"`

	// The adjustment code in response to the complaint. This maps reason code in
	//  URCS for the complaint transaction to `Resp.Ref.adjCode` in complaint
	//  response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.adjustment_code
	AdjustmentCode *string `json:"adjustmentCode,omitempty"`

	// It defines the Adjustment Reference ID which has been updated in the
	//  complaint response. This maps to `adjRefID` in complaint response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.adjustment_reference_id
	AdjustmentReferenceID *string `json:"adjustmentReferenceID,omitempty"`

	// Adjustment Remarks. This maps to `adjRemarks` in complaint response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.adjustment_remarks
	AdjustmentRemarks *string `json:"adjustmentRemarks,omitempty"`

	// The Approval Reference Number. This maps to `approvalNum` in complaint
	//  response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.approval_number
	ApprovalNumber *string `json:"approvalNumber,omitempty"`

	// Process Status of the transaction. This maps to `procStatus` in complaint
	//  response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.process_status
	ProcessStatus *string `json:"processStatus,omitempty"`

	// The adjustment timestamp when bank performs the adjustment for the received
	//  complaint request. This maps to `adjTs` in complaint response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.adjustment_time
	AdjustmentTime *string `json:"adjustmentTime,omitempty"`

	// The payer in the original financial transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.payer
	Payer *Participant `json:"payer,omitempty"`

	// The payee in the original financial transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.payee
	Payee *Participant `json:"payee,omitempty"`

	// The result of the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.result
	Result *string `json:"result,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.Complaint
type Complaint struct {
	// The name of the complaint. This uniquely identifies the complaint.
	//  Format of name is
	//  projects/{project_id}/complaints/{complaint_id}.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Complaint.name
	Name *string `json:"name,omitempty"`

	// The reason for raising the complaint. This maps adjustment flag
	//  and reason code for the complaint to `reqAdjFlag` and `reqAdjCode` in
	//  complaint request respectively while raising a complaint.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Complaint.raise_complaint_adjustment
	RaiseComplaintAdjustment *RaiseComplaintAdjustment `json:"raiseComplaintAdjustment,omitempty"`

	// Required. Details required for raising / resolving a complaint.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Complaint.details
	Details *CaseDetails `json:"details,omitempty"`

	// The reason for resolving the complaint. It provides adjustment values while
	//  resolving and for already resolved complaints. This maps adjustment flag
	//  and reason code for the complaint to `reqAdjFlag` and `reqAdjCode` in
	//  complaint request respectively when a complete resolution is done via
	//  Resolve Complaint API otherwise maps to `respAdjFlag` and `respAdjCode` in
	//  complaint response respectively when a complaint request from UPI is
	//  directly resolved by issuer switch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Complaint.resolve_complaint_adjustment
	ResolveComplaintAdjustment *ResolveComplaintAdjustment `json:"resolveComplaintAdjustment,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails
type DeviceDetails struct {
	// The payment application on the device.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.payment_app
	PaymentApp *string `json:"paymentApp,omitempty"`

	// The capability of the device.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.capability
	Capability *string `json:"capability,omitempty"`

	// The geo-code of the device.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.geo_code
	GeoCode *LatLng `json:"geoCode,omitempty"`

	// The device's ID.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.id
	ID *string `json:"id,omitempty"`

	// The device's IP address.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The coarse location of the device.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.location
	Location *string `json:"location,omitempty"`

	// The operating system on the device.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.operating_system
	OperatingSystem *string `json:"operatingSystem,omitempty"`

	// The device's telecom provider.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.telecom_provider
	TelecomProvider *string `json:"telecomProvider,omitempty"`

	// The type of device.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.DeviceDetails.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.OriginalTransaction
type OriginalTransaction struct {
	// Required. Uniquely identifies the original transaction. This maps to the
	//  `Txn.Id` value of the original transaction in India's UPI system.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.OriginalTransaction.transaction_id
	TransactionID *string `json:"transactionID,omitempty"`

	// Required. Retrieval Reference Number (RRN) of the original transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.OriginalTransaction.retrieval_reference_number
	RetrievalReferenceNumber *string `json:"retrievalReferenceNumber,omitempty"`

	// Timestamp of the original transaction request.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.OriginalTransaction.request_time
	RequestTime *string `json:"requestTime,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.Participant
type Participant struct {
	// The payment address of the participant. In the UPI system, this will be the
	//  virtual payment address (VPA) of the participant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Participant.payment_address
	PaymentAddress *string `json:"paymentAddress,omitempty"`

	// The persona of the participant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Participant.persona
	Persona *string `json:"persona,omitempty"`

	// The name of the participant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Participant.user
	User *string `json:"user,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.RaiseComplaintAdjustment
type RaiseComplaintAdjustment struct {
	// Required. The adjustment flag in URCS for the complaint transaction. This
	//  maps to `reqAdjFlag` in complaint request and `respAdjFlag` in complaint
	//  response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.RaiseComplaintAdjustment.adjustment_flag
	AdjustmentFlag *string `json:"adjustmentFlag,omitempty"`

	// Required. The adjustment code in URCS for the complaint transaction. This
	//  maps to `reqAdjCode` in complaint request.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.RaiseComplaintAdjustment.adjustment_code
	AdjustmentCode *string `json:"adjustmentCode,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.ResolveComplaintAdjustment
type ResolveComplaintAdjustment struct {
	// Required. The adjustment flag in URCS for the complaint transaction. This
	//  maps to `reqAdjFlag` in complaint request and `respAdjFlag` in complaint
	//  response.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.ResolveComplaintAdjustment.adjustment_flag
	AdjustmentFlag *string `json:"adjustmentFlag,omitempty"`

	// Required. The adjustment code in URCS for the complaint transaction. This
	//  maps to `reqAdjCode` in complaint request.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.ResolveComplaintAdjustment.adjustment_code
	AdjustmentCode *string `json:"adjustmentCode,omitempty"`
}

// +kcc:proto=google.type.LatLng
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	// +kcc:proto:field=google.type.LatLng.latitude
	Latitude *float64 `json:"latitude,omitempty"`

	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	// +kcc:proto:field=google.type.LatLng.longitude
	Longitude *float64 `json:"longitude,omitempty"`
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

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReferenceObservedState struct {
	// Output only. Type of account. Examples include SAVINGS, CURRENT, etc.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_type
	AccountType *string `json:"accountType,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse
type CaseResponseObservedState struct {
	// The payer in the original financial transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.CaseResponse.payer
	Payer *ParticipantObservedState `json:"payer,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.Complaint
type ComplaintObservedState struct {
	// Output only. Response to the raised / resolved complaint.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Complaint.response
	Response *CaseResponse `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.Participant
type ParticipantObservedState struct {
	// Output only. Unique identification of an account according to India's UPI
	//  standards.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Participant.account
	Account *AccountReference `json:"account,omitempty"`

	// Output only. The device info of the participant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Participant.device_details
	DeviceDetails *DeviceDetails `json:"deviceDetails,omitempty"`

	// Output only. The mobile number of the participant.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.Participant.mobile_number
	MobileNumber *string `json:"mobileNumber,omitempty"`
}
