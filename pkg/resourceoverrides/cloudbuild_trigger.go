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

func GetCloudBuildTriggerResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "CloudBuildTrigger",
	}
	// Default the optional and immutable 'location' field to 'global' because
	// implicitly defaulted value will be removed in the live state.
	ro.Overrides = append(ro.Overrides, defaultLocationFieldToGlobal())
	return ro
}

// defaultLocationFieldToGlobal defaults the optional top-level 'location' field
// to 'global'.
func defaultLocationFieldToGlobal() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return KeepTopLevelFieldOptionalWithDefault(crd, "global", "location")
	}
	return o
}
