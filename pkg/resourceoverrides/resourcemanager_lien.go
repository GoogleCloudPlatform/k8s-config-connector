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
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetResourceManagerLienResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ResourceManagerLien",
	}
	ro.Overrides = append(ro.Overrides, buildResourceManagerLien())
	return ro
}

func buildResourceManagerLien() ResourceOverride {
	h := &ResourceManagerLien{}
	o := ResourceOverride{
		PreActuationTransform: h.PreActuationTransform,
	}
	return o
}

type ResourceManagerLien struct {
}

func (h *ResourceManagerLien) PreActuationTransform(r *k8s.Resource) error {
	parentRef, exists, err := unstructured.NestedString(r.Spec, "parent", "projectRef", "external")
	if err != nil {
		return fmt.Errorf("error reading spec.parent.projectRef.external: %w", err)
	}
	if exists && parentRef != "" {
		normalized := normalizeLienParent(parentRef)
		if normalized != parentRef {
			if err := unstructured.SetNestedField(r.Spec, normalized, "parent", "projectRef", "external"); err != nil {
				return fmt.Errorf("error updating spec.parent.projectRef.external: %w", err)
			}
		}
	}
	return nil
}

func normalizeLienParent(parent string) string {
	projectStr := parent
	if strings.HasPrefix(parent, "projects/") {
		projectStr = strings.TrimPrefix(parent, "projects/")
	}
	// Strip sovereign partition prefix if present (e.g. 'partition:project-id' -> 'project-id')
	if idx := strings.Index(projectStr, ":"); idx != -1 {
		projectStr = projectStr[idx+1:]
	}
	return fmt.Sprintf("projects/%s", projectStr)
}
