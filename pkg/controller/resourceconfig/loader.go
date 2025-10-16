// Copyright 2025 Google LLC
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

package resourceconfig

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func LoadConfig() *ResourcesControllerMap {
	return ControllerConfigStatic
}

func IsControllerSupported(gvk schema.GroupVersionKind, controllerType k8s.ReconcilerType) bool {
	config := LoadConfig()
	if config == nil {
		return false
	}
	controllerConfig, exists := (*config)[gvk.GroupKind()]
	if !exists {
		return false
	}
	for _, supportedController := range controllerConfig.SupportedControllers {
		if supportedController == controllerType {
			return true
		}
	}
	return false
}
