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

package v1beta1

// *** DISCLAIMER ***
// This file is duplicated in the scripts/generate-go-crd-clients/k8s/ directory.
// If you're making modifications to this file, please make sure to modify the
// corresponding file (`types.go`) so the generated go-clients have an accurate
// representation of this package.

import (
	v1 "k8s.io/api/core/v1"
)

const (
	ReadyConditionType = "Ready"
)

type Condition struct {
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last
	// transition.
	Reason string `json:"reason,omitempty"`

	// Status is the status of the condition. Can be True, False, Unknown.
	Status v1.ConditionStatus `json:"status,omitempty"`

	// Type is the type of the condition.
	Type string `json:"type,omitempty"`
}
