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

package k8s

import (
	"fmt"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var scoreForVersion = map[string]int{
	"v1alpha1": 111,
	"v1beta1":  121,
	"v1":       131,
}

func PreferredVersion(crd *apiextensions.CustomResourceDefinition) *apiextensions.CustomResourceDefinitionVersion {
	bestScore := -1
	var preferredVersion *apiextensions.CustomResourceDefinitionVersion
	for _, version := range crd.Spec.Versions {
		score, found := scoreForVersion[version.Name]
		if !found {
			panic(fmt.Sprintf("version %q is not known in getLatestVersion", version.Name))
		}

		if score > bestScore {
			preferredVersion = &version
			bestScore = score
		}

	}
	return preferredVersion
}

// GetGroupKindFromCRD returns the GroupKind that the CRD defines.
func GetGroupKindFromCRD(crd *apiextensions.CustomResourceDefinition) schema.GroupKind {
	return schema.GroupKind{
		Group: crd.Spec.Group,
		Kind:  crd.Spec.Names.Kind,
	}
}

func GetAPIVersionFromCRD(crd *apiextensions.CustomResourceDefinition) string {
	panicIfNoVersionPresent(crd)
	// Currently KCC CRDs only support one version.
	return fmt.Sprintf("%v/%v", crd.Spec.Group, PreferredVersion(crd).Name)
}

func GetVersionFromCRD(crd *apiextensions.CustomResourceDefinition) string {
	panicIfNoVersionPresent(crd)
	// Currently KCC CRDs only support one version.
	return PreferredVersion(crd).Name
}

func GetOpenAPIV3SchemaFromCRD(crd *apiextensions.CustomResourceDefinition) *apiextensions.JSONSchemaProps {
	panicIfNoVersionPresent(crd)
	// Currently KCC CRDs only support one version.
	return PreferredVersion(crd).Schema.OpenAPIV3Schema
}

func panicIfNoVersionPresent(crd *apiextensions.CustomResourceDefinition) {
	if len(crd.Spec.Versions) == 0 {
		panic(fmt.Sprintf("no versions present in CRD %v\n", crd.GetName()))
	}
}
