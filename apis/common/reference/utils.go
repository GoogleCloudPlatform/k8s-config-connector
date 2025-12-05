// Copyright 2025 Google LLC
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

package common

import (
	"fmt"
	"strings"
)

// FixStaleComputeExternalFormat converts the "External" reference field to the right format if a SelfLink value is used.
// This guarantees the backward compatibility for Compute Beta resources.
func FixStaleComputeExternalFormat(external string) string {
	// Extra handling for compatibility with the deprecated externalRef format.
	// See comment https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/5234#discussion_r2470673611
	external = strings.TrimPrefix(external, "/compute.googleapis.com")

	prefix := "https://www.googleapis.com/compute/"
	if strings.HasPrefix(external, prefix) {
		external = strings.TrimPrefix(external, prefix)
		tokens := strings.Split(external, "/")
		if tokens[0] == "v1" || tokens[0] == "v1beta1" || tokens[0] == "beta" {
			external = strings.TrimPrefix(external, tokens[0])
		} else {
			// The external with unknown version will fail the validation later in ParseComputeNetworkExternal,
			// todo: shall we log an error here?
			fmt.Printf("WARNING: received Compute selfLink with unknown version %s, accepted versions are v1, v1beta1 and beta.", tokens[0])
			return external
		}
	}
	return strings.TrimPrefix(external, "/")
}
