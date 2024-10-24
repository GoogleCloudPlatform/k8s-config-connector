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

package v1beta1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetLocation(u *unstructured.Unstructured) (string, error) {
	location, _, err := unstructured.NestedString(u.Object, "spec", "location")
	if err != nil {
		return "", fmt.Errorf("reading spec.location from %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}
	if location == "" {
		return "", fmt.Errorf("spec.location not set in %v %v/%v", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName())
	}
	return location, nil
}
