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
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ControllerSelector struct {
	Config *ResourcesControllerMap
}

func NewControllerSelector(config *ResourcesControllerMap) *ControllerSelector {
	return &ControllerSelector{
		Config: config,
	}
}

func (s *ControllerSelector) SelectController(gvk schema.GroupVersionKind, namespace string, ccc *v1beta1.ConfigConnectorContextSpec) k8s.ReconcilerType {
	// Check if there is an override in ConfigConnectorContext for the specific GVK.
	if ccc != nil && ccc.Experiments.ControllerOverrides != nil {
		if controllerType, ok := ccc.Experiments.ControllerOverrides[gvk.GroupKind()]; ok {
			return controllerType
		}
	}

	// Fall back to the default from the central configuration.
	for k, v := range *s.Config {
		if k.Group == gvk.Group && k.Kind == gvk.Kind {
			return v.DefaultController
		}
	}

	// Should not happen if the config is complete.
	return ""
}
