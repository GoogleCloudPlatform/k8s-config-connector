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

package kccfeatureflags

import (
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// UseDirectReconciler is true if we should use the direct reconciler to actuate the specified resource.
func UseDirectReconciler(gk schema.GroupKind) bool {
	directReconcilers := os.Getenv("KCC_USE_DIRECT_RECONCILERS")
	if directReconcilers == "" {
		return false
	}

	for _, directReconciler := range strings.Split(directReconcilers, ",") {
		if directReconciler == gk.Kind {
			return true
		}
	}

	return false
}
