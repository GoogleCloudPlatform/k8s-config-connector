// Copyright 2026 Google LLC
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

package v1alpha1_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdelivery/v1alpha1"
)

func TestConfigDeliveryFleetPackageRef(t *testing.T) {
	r := &v1alpha1.ConfigDeliveryFleetPackageRef{
		Name:      "my-fleetpackage",
		Namespace: "default",
	}

	gvk := r.GetGVK()
	if gvk.Kind != "ConfigDeliveryFleetPackage" {
		t.Errorf("expected Kind to be ConfigDeliveryFleetPackage, got %s", gvk.Kind)
	}

	nsName := r.GetNamespacedName()
	if nsName.Name != "my-fleetpackage" || nsName.Namespace != "default" {
		t.Errorf("expected NamespacedName to be default/my-fleetpackage, got %s", nsName.String())
	}

	r.SetExternal("projects/my-project/locations/global/fleetPackages/my-fleetpackage")
	if r.GetExternal() != "projects/my-project/locations/global/fleetPackages/my-fleetpackage" {
		t.Errorf("expected external to be projects/my-project/locations/global/fleetPackages/my-fleetpackage, got %s", r.GetExternal())
	}
}
