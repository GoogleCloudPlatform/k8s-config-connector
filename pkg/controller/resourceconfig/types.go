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

type ResourcesControllerMap map[schema.GroupKind]ResourceConfig

// ResourceConfig defines the controller configuration for a specific resource Kind.
type ResourceConfig struct {
	DefaultController    k8s.ReconcilerType
	SupportedControllers []k8s.ReconcilerType
}

func (c *ResourcesControllerMap) GetControllerForGVK(gvk schema.GroupVersionKind) (*ResourceConfig, error) {
	groupKind := schema.GroupKind{
		Group: gvk.Group,
		Kind:  gvk.Kind,
	}
	if config, ok := (*c)[groupKind]; ok {
		return &config, nil
	}
	return nil, nil
}
