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

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.MetadataTransaction
type MetadataTransaction struct {
	// The name of the metadata transaction. This uniquely identifies the
	//  transaction. Format of name is
	//  projects/{project_id}/metadataTransaction/{metadata_transaction_id}.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.MetadataTransaction.name
	Name *string `json:"name,omitempty"`

	// Information about the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.MetadataTransaction.info
	Info *TransactionInfo `json:"info,omitempty"`
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

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo
type TransactionInfo struct {

	// Metadata about the API transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.metadata
	Metadata *TransactionInfo_TransactionMetadata `json:"metadata,omitempty"`

	// Risk information as provided by the payments orchestrator.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.risk_info
	RiskInfo []TransactionInfo_TransactionRiskInfo `json:"riskInfo,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.AdapterInfo
type TransactionInfo_AdapterInfo struct {
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.AdapterInfo.ResponseMetadata
type TransactionInfo_AdapterInfo_ResponseMetadata struct {
	// A map of name-value pairs.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.AdapterInfo.ResponseMetadata.values
	Values map[string]string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionErrorDetails
type TransactionInfo_TransactionErrorDetails struct {
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata
type TransactionInfo_TransactionMetadata struct {
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionRiskInfo
type TransactionInfo_TransactionRiskInfo struct {
	// Entity providing the risk score. This could either be the payment service
	//  provider or the payment orchestrator (UPI, etc).
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionRiskInfo.provider
	Provider *string `json:"provider,omitempty"`

	// Type of risk. Examples include `TXNRISK`.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionRiskInfo.type
	Type *string `json:"type,omitempty"`

	// Numeric value of risk evaluation ranging from 0 (No Risk) to 100 (Maximum
	//  Risk).
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionRiskInfo.value
	Value *string `json:"value,omitempty"`
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

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.AccountReference
type AccountReferenceObservedState struct {
	// Output only. Type of account. Examples include SAVINGS, CURRENT, etc.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.AccountReference.account_type
	AccountType *string `json:"accountType,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.MetadataTransaction
type MetadataTransactionObservedState struct {
	// Information about the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.MetadataTransaction.info
	Info *TransactionInfoObservedState `json:"info,omitempty"`

	// Output only. The initiator of the metadata transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.MetadataTransaction.initiator
	Initiator *Participant `json:"initiator,omitempty"`
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

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo
type TransactionInfoObservedState struct {
	// Output only. An identifier that is mandatorily present in every transaction
	//  processed via UPI. This maps to UPI's transaction ID.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.id
	ID *string `json:"id,omitempty"`

	// Output only. The API type of the transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.api_type
	ApiType *string `json:"apiType,omitempty"`

	// Output only. The transaction type.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.transaction_type
	TransactionType *string `json:"transactionType,omitempty"`

	// Output only. The transaction subtype.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.transaction_sub_type
	TransactionSubType *string `json:"transactionSubType,omitempty"`

	// Output only. The transaction's state.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.state
	State *string `json:"state,omitempty"`

	// Metadata about the API transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.metadata
	Metadata *TransactionInfo_TransactionMetadataObservedState `json:"metadata,omitempty"`

	// Output only. Any error details for the current API transaction, if the
	//  state is `FAILED`.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.error_details
	ErrorDetails *TransactionInfo_TransactionErrorDetails `json:"errorDetails,omitempty"`

	// Output only. Information about the adapter invocation from the issuer
	//  switch for processing this API transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.adapter_info
	AdapterInfo *TransactionInfo_AdapterInfo `json:"adapterInfo,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.AdapterInfo
type TransactionInfo_AdapterInfoObservedState struct {
	// Output only. List of adapter request IDs (colon separated) used when
	//  invoking the Adapter APIs for fulfilling a transaction request.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.AdapterInfo.request_ids
	RequestIds *string `json:"requestIds,omitempty"`

	// Output only. Response metadata included by the adapter in its response to
	//  an API invocation from the issuer switch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.AdapterInfo.response_metadata
	ResponseMetadata *TransactionInfo_AdapterInfo_ResponseMetadata `json:"responseMetadata,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionErrorDetails
type TransactionInfo_TransactionErrorDetailsObservedState struct {
	// Output only. Error code of the failed transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionErrorDetails.error_code
	ErrorCode *string `json:"errorCode,omitempty"`

	// Output only. Error description for the failed transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionErrorDetails.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Output only. Error code as per the UPI specification. The issuer switch
	//  maps the ErrorCode to an appropriate error code that complies with the
	//  UPI specification.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionErrorDetails.upi_error_code
	UpiErrorCode *string `json:"upiErrorCode,omitempty"`
}

// +kcc:proto=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata
type TransactionInfo_TransactionMetadataObservedState struct {
	// Output only. The time at which the transaction resource was created by
	//  the issuer switch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the transaction resource was last updated
	//  by the issuer switch.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A reference id for the API transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata.reference_id
	ReferenceID *string `json:"referenceID,omitempty"`

	// Output only. A reference URI to this API transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata.reference_uri
	ReferenceURI *string `json:"referenceURI,omitempty"`

	// Output only. A descriptive note about this API transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata.description
	Description *string `json:"description,omitempty"`

	// Output only. The initiation mode of this API transaction. In UPI, the
	//  values are as defined by the UPI API specification.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata.initiation_mode
	InitiationMode *string `json:"initiationMode,omitempty"`

	// Output only. The purpose code of this API transaction. In UPI, the values
	//  are as defined by the UPI API specification.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata.purpose_code
	PurposeCode *string `json:"purposeCode,omitempty"`

	// Output only. The reference category of this API transaction.
	// +kcc:proto:field=google.cloud.paymentgateway.issuerswitch.v1.TransactionInfo.TransactionMetadata.reference_category
	ReferenceCategory *string `json:"referenceCategory,omitempty"`
}
