// Copyright 2022 Google LLC
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

// ResourceReference defines a relationship to another resource
type ResourceReference struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	External  string `json:"external,omitempty"`
}

// SensitiveField represents a field that expects sensitive information (e.g. passwords)
type SensitiveField struct {
	Value     *string      `json:"value,omitempty"`
	ValueFrom *ValueSource `json:"valueFrom,omitempty"`
}

// ValueSource represents a source for the value of a SensitiveField
type ValueSource struct {
	SecretKeyRef *SecretKeyReference `json:"secretKeyRef,omitempty"`
}

// SecretKeyReference represents a relationship to a keyed value in a Secret object
type SecretKeyReference struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}
