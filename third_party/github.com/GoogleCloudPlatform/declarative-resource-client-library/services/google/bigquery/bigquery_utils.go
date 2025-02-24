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
// Package bigquery provices methods and types for managing bigquery GCP resources.
package bigquery

import ()

func equalsDatasetAccessRole(m, n *string) bool {
	if m == nil && n == nil {
		return true
	}
	if m == nil || n == nil {
		return false
	}
	if *m == *n {
		return true
	}
	mappings := map[string]string{
		"OWNER":  "roles/bigquery.dataOwner",
		"WRITER": "roles/bigquery.dataEditor",
		"READER": "roles/bigquery.dataViewer",
	}
	if mappings[*m] == *n || mappings[*n] == *m {
		return true
	}
	return false
}

func canonicalizeDatasetAccessRole(m, n interface{}) bool {
	if m == nil && n == nil {
		return true
	}
	mVal, _ := m.(*string)
	nVal, _ := n.(*string)
	return equalsDatasetAccessRole(mVal, nVal)
}
