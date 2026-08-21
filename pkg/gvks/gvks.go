// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gvks

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/externalonlygvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// All returns GroupVersionKinds corresponding to GCP resources known to KCC,
// including those unsupported by KCC but commonly referenced by KCC resources.
func All(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader) ([]schema.GroupVersionKind, error) {
	gvks, err := supportedgvks.All(smLoader, serviceMetaLoader)
	if err != nil {
		return gvks, fmt.Errorf("error loading all supported GVKs: %w", err)
	}
	gvks = append(gvks, externalonlygvks.All()...)
	return gvks, nil
}

func GVKForKind(kind string, smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetaLoader metadata.ServiceMetadataLoader) (gvk schema.GroupVersionKind, found bool, err error) {
	allGVKs, err := All(smLoader, serviceMetaLoader)
	if err != nil {
		return schema.GroupVersionKind{}, false, fmt.Errorf("error loading all supported GVKs: %w", err)
	}
	for _, v := range allGVKs {
		if v.Kind == kind {
			return v, true, nil
		}
	}
	return schema.GroupVersionKind{}, false, nil
}
