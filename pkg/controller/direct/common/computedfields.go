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

package common

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// ParentPair tracks matching actual and desired proto messages for a given parent path.
type ParentPair struct {
	Actual  protoreflect.Message
	Desired protoreflect.Message
}

// BuildParentMap collects all parent sub-messages required by paths from actualPb,
// and ensures matching container messages are initialized on desiredPb.
func BuildParentMap(desiredPb, actualPb protoreflect.ProtoMessage, paths []string) map[string]ParentPair {
	parents := make(map[string]ParentPair)
	if actualPb == nil || desiredPb == nil {
		return parents
	}

	parents[""] = ParentPair{Actual: actualPb.ProtoReflect(), Desired: desiredPb.ProtoReflect()}

	for _, path := range paths {
		segments := strings.Split(path, ".")
		if len(segments) <= 1 {
			continue
		}

		parentSegments := segments[:len(segments)-1]
		currActual := actualPb.ProtoReflect()
		currDesired := desiredPb.ProtoReflect()
		currPath := ""

		for _, seg := range parentSegments {
			if currPath == "" {
				currPath = seg
			} else {
				currPath = currPath + "." + seg
			}

			if pair, exists := parents[currPath]; exists {
				currActual = pair.Actual
				currDesired = pair.Desired
				continue
			}

			fd := FindProtoField(currActual.Descriptor(), seg)
			if fd == nil || fd.Kind() != protoreflect.MessageKind || !currActual.Has(fd) {
				break
			}

			currActual = currActual.Get(fd).Message()
			if !currDesired.Has(fd) {
				currDesired.Set(fd, currDesired.NewField(fd))
			}
			currDesired = currDesired.Get(fd).Message()

			parents[currPath] = ParentPair{Actual: currActual, Desired: currDesired}
		}
	}

	return parents
}

// FindProtoField matches a KRM leaf field name to its corresponding proto field descriptor
// by comparing the field's JSONName against the KRM field name (ignoring case).
func FindProtoField(desc protoreflect.MessageDescriptor, krmLeaf string) protoreflect.FieldDescriptor {
	krmName := strings.TrimSuffix(krmLeaf, "Ref")
	for i := 0; i < desc.Fields().Len(); i++ {
		fd := desc.Fields().Get(i)
		if strings.EqualFold(fd.JSONName(), krmName) {
			return fd
		}
	}
	return nil
}
