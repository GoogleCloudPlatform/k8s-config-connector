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

// GetLatestGVKFromCRD returns the GVK the CRD is for.
// One CRD can describe multiple versions; we choose the latest.
func GetLatestGVKFromCRD(crd *apiextensions.CustomResourceDefinition) schema.GroupVersionKind {
	panicIfNoVersionPresent(crd)
	latestVersion := getLatestVersion(crd)
	return schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: latestVersion,
		Kind:    crd.Spec.Names.Kind,
	}
}

var scoreForVersion = map[string]int{
	"v1alpha1": 111,
	"v1beta1":  121,
	"v1":       131,
}

func getLatestVersion(crd *apiextensions.CustomResourceDefinition) string {
	versions := crd.Spec.Versions
	bestScore := -1
	bestVersion := ""
	for _, version := range versions {
		score, found := scoreForVersion[version.Name]
		if !found {
			panic(fmt.Sprintf("version %q is not known in getLatestVersion", version.Name))
		}
		if score > bestScore {
			bestVersion = version.Name
			bestScore = score
		}
	}
	return bestVersion
}

func GetAllVersionsFromCRD(crd *apiextensions.CustomResourceDefinition) []string {
	panicIfNoVersionPresent(crd)
	var versions []string
	for _, version := range crd.Spec.Versions {
		versions = append(versions, version.Name)
	}
	return versions
}

func GetOpenAPIV3SchemaFromCRD(crd *apiextensions.CustomResourceDefinition, version string) *apiextensions.JSONSchemaProps {
	// Currently KCC CRDs only support one version.
	for i := range crd.Spec.Versions {
		v := &crd.Spec.Versions[i]
		if v.Name == version {
			return v.Schema.OpenAPIV3Schema
		}
	}
	panic(fmt.Sprintf("version %q is not known in GetOpenAPIV3SchemaFromCRD", version))
}

func panicIfNoVersionPresent(crd *apiextensions.CustomResourceDefinition) {
	if len(crd.Spec.Versions) == 0 {
		panic(fmt.Sprintf("no versions present in CRD %v\n", crd.GetName()))
	}
}
