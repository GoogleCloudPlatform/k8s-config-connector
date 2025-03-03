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
// Package identitytoolkit provides types and functions for representing identitytoolkit GCP resources.
package identitytoolkit

import (
	"errors"
	"regexp"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func fetchName(o map[string]interface{}) (*string, error) {
	name, ok := o["name"].(string)
	if !ok || name == "" {
		return nil, errors.New("unable to fetch name from output")
	}
	return &name, nil
}

// Returns a copy of the given map normalizing format of phone numbers in it.
func normalizeNumbers(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	re := regexp.MustCompile("[^0-9+]")
	for phoneNumber, code := range m {
		normalized := re.ReplaceAllString(phoneNumber, "")
		n[normalized] = code
	}
	return n
}

// canonicalizeConfigTestPhoneNumbers compares two maps with phone number keys with the phone numbers
// normalized.
func canonicalizeConfigTestPhoneNumbers(m, n interface{}) bool {
	if m == nil && n == nil {
		return true
	}
	if m == nil || n == nil {
		return false
	}
	mMap, _ := m.(map[string]string)
	nMap, _ := n.(map[string]string)
	mNormalized := normalizeNumbers(mMap)
	nNormalized := normalizeNumbers(nMap)
	ds, err := dcl.Diff(mNormalized, nNormalized, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, dcl.FieldName{})
	return len(ds) == 0 && err == nil
}
