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

package securesourcemanager

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func setObservedState(u *unstructured.Unstructured, observedState any) error {
	observedStateUnstructured, err := runtime.DefaultUnstructuredConverter.ToUnstructured(observedState)
	if err != nil {
		return fmt.Errorf("error converting observedState to unstructured: %w", err)
	}

	if err := unstructured.SetNestedMap(u.Object, observedStateUnstructured, "status", "observedState"); err != nil {
		return fmt.Errorf("setting observedState: %w", err)
	}

	return nil
}

func lastComponent(s string) string {
	return s[strings.LastIndex(s, "/")+1:]
}
