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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ResourceControllerConfig struct {
	DefaultController    k8s.ReconcilerType
	SupportedControllers []k8s.ReconcilerType
}

type ResourcesControllerMap map[schema.GroupKind]ResourceControllerConfig

func (c *ResourcesControllerMap) GetControllersForGVK(gvk schema.GroupVersionKind) (*ResourceControllerConfig, error) {
	groupKind := schema.GroupKind{
		Group: gvk.Group,
		Kind:  gvk.Kind,
	}
	if config, ok := (*c)[groupKind]; ok {
		return &config, nil
	}
	return nil, fmt.Errorf("no controller config found for GVK %v", gvk)
}
