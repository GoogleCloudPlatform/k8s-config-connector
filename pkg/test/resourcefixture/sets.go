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

package resourcefixture

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
)

// returns an id that is unique for each resource config
func GetUniqueResourceConfigID(rc v1alpha1.ResourceConfig) string {
	if rc.Locationality != "" {
		return fmt.Sprintf("%v:%v", rc.Kind, rc.Locationality)
	}
	if rc.Name == "google_compute_instance" || rc.Name == "google_compute_instance_from_template" {
		return fmt.Sprintf("%v:%v", rc.Kind, rc.Name)
	}

	return rc.Kind
}
