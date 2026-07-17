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

func TestConfigDeliveryResourceBundleRef(t *testing.T) {
	r := &v1alpha1.ConfigDeliveryResourceBundleRef{
		Name:      "my-bundle",
		Namespace: "default",
	}

	gvk := r.GetGVK()
	if gvk.Kind != "ConfigDeliveryResourceBundle" {
		t.Errorf("expected Kind to be ConfigDeliveryResourceBundle, got %s", gvk.Kind)
	}

	nsName := r.GetNamespacedName()
	if nsName.Name != "my-bundle" || nsName.Namespace != "default" {
		t.Errorf("expected NamespacedName to be default/my-bundle, got %s", nsName.String())
	}

	r.SetExternal("projects/my-project/locations/global/resourceBundles/my-bundle")
	if r.GetExternal() != "projects/my-project/locations/global/resourceBundles/my-bundle" {
		t.Errorf("expected external to be projects/my-project/locations/global/resourceBundles/my-bundle, got %s", r.GetExternal())
	}
}
