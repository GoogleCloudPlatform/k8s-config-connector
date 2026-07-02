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

package compute

import (
	computepb "cloud.google.com/go/compute/apiv1/computepb"
)

// SignedURLKey_ToProto builds a computepb.SignedUrlKey from (keyName, keyValue).
func SignedURLKey_ToProto(keyName, keyValue string) *computepb.SignedUrlKey {
	return &computepb.SignedUrlKey{
		KeyName:  &keyName,
		KeyValue: &keyValue,
	}
}

// SignedURLKey_KeyNameFromProto returns the key name from a SignedUrlKey proto.
// keyValue is intentionally not returned — it is write-only in the GCP API.
func SignedURLKey_KeyNameFromProto(key *computepb.SignedUrlKey) string {
	if key == nil {
		return ""
	}
	return key.GetKeyName()
}
