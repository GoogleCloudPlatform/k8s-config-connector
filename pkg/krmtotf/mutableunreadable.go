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

package krmtotf

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
)

func MutableButUnreadableFieldsAnnotationFor(r *Resource) (string, error) {
	paths := getMutableButUnreadablePaths(r)
	return k8s.GenerateMutableButUnreadableFieldsAnnotation(&r.Resource, paths)
}

// getMutableButUnreadablePaths returns the list of fields supported by the
// resource that are mutable but unreadable. Each field is broken down into its
// path elements (i.e. ["spec", "fooBar"] instead of "spec.fooBar").
func getMutableButUnreadablePaths(r *Resource) [][]string {
	tfFields := r.ResourceConfig.MutableButUnreadableFields
	lowerCamelCasePaths := make([][]string, 0)
	for _, tfField := range tfFields {
		tfPath := strings.Split(tfField, ".")
		lowerCamelCasePaths = append(lowerCamelCasePaths, text.SnakeCaseStrsToLowerCamelCaseStrs(tfPath))
	}
	return lowerCamelCasePaths
}

func getMutableButUnreadableFieldsFromAnnotations(r *Resource) (map[string]interface{}, error) {
	paths := getMutableButUnreadablePaths(r)
	return k8s.GetMutableButUnreadableFieldsFromAnnotations(&r.Resource, paths)
}
