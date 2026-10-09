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

package scaffold

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// isDeprecated reports whether the proto marks field [deprecated = true].
func isDeprecated(field protoreflect.FieldDescriptor) bool {
	opts, ok := field.Options().(*descriptorpb.FieldOptions)
	return ok && opts.GetDeprecated()
}

// NestedDeprecatedFields returns a queue item for each field of a nested Spec
// struct that the proto marks deprecated.
//
// PrepopulateSpec leaves deprecated top-level fields out of the Spec, and
// walkSpecFields skips them, so every item here is for a nested field. A nested
// field stays in its struct. The struct is written to types.generated.go, where
// existing Kinds can use it too, and dropping the field there would change
// their CRDs.
//
// The items use the same reason as the top-level ones from PrepopulateSpec.
// Their paths differ, so their queue keys do not collide.
func NestedDeprecatedFields(msg protoreflect.MessageDescriptor, opts codegen.WriteOptions) []JudgementItem {
	var out []JudgementItem
	walkSpecFields(msg, ".spec", opts, true, map[protoreflect.FullName]bool{}, func(path, _ string, field protoreflect.FieldDescriptor) {
		if !isDeprecated(field) {
			return
		}
		out = append(out, JudgementItem{
			FieldPath: path,
			Reason:    "deprecated-field",
			Detail: "the proto marks this field deprecated. It is kept because generated structs " +
				"can be shared with other Kinds. Remove it by hand if no existing Kind uses it",
		})
	})
	return out
}
