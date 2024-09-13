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

package resourceoverrides

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func GetFirestoreIndexResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "FirestoreIndex",
	}
	// Default the optional and immutable 'database' field to '(default)'
	// because implicitly defaulted value will be removed in the live state.
	ro.Overrides = append(ro.Overrides, defaultDatabaseFieldToDefault())
	return ro
}

// defaultDatabaseFieldToDefault defaults the optional top-level
// 'database' field to '(default)'.
func defaultDatabaseFieldToDefault() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return KeepTopLevelFieldOptionalWithDefault(crd, "(default)", "database")
	}
	return o
}
