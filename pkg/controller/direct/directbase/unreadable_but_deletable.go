// Copyright 2026 Google LLC
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

package directbase

import (
	"errors"

	"google.golang.org/api/googleapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// IsNotFound is a hook initialized by the parent package (direct)
// to check if a GCP error is a 404 Not Found error.
// This breaks the circular dependency between directbase and direct.
var IsNotFound func(err error) bool

// IsUnreadableButDeletable returns true if the resource's current state allows
// it to be deleted despite being unreadable (Find failed).
func IsUnreadableButDeletable(gvk schema.GroupVersionKind, err error) bool {
	if err == nil {
		return false
	}

	// Return false directly when the error is not found.
	if IsNotFound != nil {
		if IsNotFound(err) {
			return false
		}
	} else {
		if isNotFoundFallback(err) {
			return false
		}
	}

	gk := gvk.GroupKind()
	switch gk {
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryTable"}:
		return true
	}

	return false
}

func isNotFoundFallback(err error) bool {
	var apiErr *googleapi.Error
	if errors.As(err, &apiErr) {
		if apiErr.Code == 404 {
			return true
		}
	}
	return false
}
