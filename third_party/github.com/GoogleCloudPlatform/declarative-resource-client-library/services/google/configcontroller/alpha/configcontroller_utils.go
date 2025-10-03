// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package configcontroller provides functions to manage Config Controller
// resources.
package alpha

// EncodeInstanceCreateRequest enables the Config Controller option (adds the
// `bundlesConfig` field) for all the create requests for a ConfigController
// Instance.
func EncodeInstanceCreateRequest(m map[string]interface{}) map[string]interface{} {
	m["bundlesConfig"] = map[string]interface{}{
		"configControllerConfig": map[string]interface{}{
			"enabled": true,
		},
	}
	return m
}
