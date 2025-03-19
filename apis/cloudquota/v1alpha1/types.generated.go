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

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaConfig
type QuotaConfig struct {
	// Required. The preferred value. Must be greater than or equal to -1. If set
	//  to -1, it means the value is "unlimited".
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.preferred_value
	//+required
	PreferredValue *int64 `json:"preferredValue,omitempty"`

	// Optional. The annotations map for clients to store small amounts of
	//  arbitrary data. Do not put PII or other sensitive information here. See
	//  https://google.aip.dev/128#annotations
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.api.cloudquotas.v1beta.QuotaConfig
type QuotaConfigObservedState struct {
	// Output only. Optional details about the state of this quota preference.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.state_detail
	StateDetail *string `json:"stateDetail,omitempty"`

	// Output only. Granted quota value.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.granted_value
	GrantedValue *int64 `json:"grantedValue,omitempty"`

	// Output only. The trace id that the Google Cloud uses to provision the
	//  requested quota. This trace id may be used by the client to contact Cloud
	//  support to track the state of a quota preference request. The trace id is
	//  only produced for increase requests and is unique for each request. The
	//  quota decrease requests do not have a trace id.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.trace_id
	TraceID *string `json:"traceID,omitempty"`

	// Output only. The origin of the quota preference request.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaConfig.request_origin
	RequestOrigin *string `json:"requestOrigin,omitempty"`
}
