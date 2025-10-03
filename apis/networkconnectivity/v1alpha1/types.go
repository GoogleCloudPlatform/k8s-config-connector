// Copyright 2024 Google LLC
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

// +kcc:status:proto=mockgcp.cloud.networkconnectivity.v1.GoogleRpcStatus
type GoogleRpcStatus struct {
	// The status code, which should be an enum value of google.rpc.Code.
	Code *int32 `json:"code,omitempty"`

	/*NOTYET
	// A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Details []google_protobuf_Any `json:"details,omitempty"`
	*/

	// A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
	Message *string `json:"message,omitempty"`
}
