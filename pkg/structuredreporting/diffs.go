// Copyright 2025 Google LLC
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

package structuredreporting

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Diff allows reporting of a detected difference
type Diff struct {
	Object *unstructured.Unstructured

	// Fields contains field-level diffs
	// Likely empty if NewObject is true
	Fields []DiffField

	// IsNewObject is true if the object does not exist in GCP
	IsNewObject bool
}

type DiffField struct {
	// ID is the identity of the field.  Note that this might be the proto or terraform name.
	ID  string
	Old any
	New any
}

// ReportDiff should be called by a controller when it detects diffs
func ReportDiff(ctx context.Context, diff *Diff) {
	if listener, ok := GetListenerFromContext(ctx); ok {
		listener.OnDiff(ctx, diff)
	}
}
