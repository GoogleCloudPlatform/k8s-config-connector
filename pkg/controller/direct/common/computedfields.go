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
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
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

// PopulateComputedFields populates dynamic/computed server-generated values in O(N) linear time
// by checking omitted KRM fields against non-nil parents in actualPb and directly assigning their values to desiredPb.
// If a path points to a composite proto message, omitted child fields are merged recursively.
func PopulateComputedFields(desiredSpec any, desiredPb, actualPb protoreflect.ProtoMessage, paths []string) {
	if desiredPb == nil || actualPb == nil {
		return
	}

	// 1. Build map of non-nil parents from actualPb and initialize them on desiredPb
	parentMap := BuildParentMap(desiredPb, actualPb, paths)

	// 2. Collect all paths explicitly set in desiredSpec in a single O(N) pass
	var presentFields sets.Set[string]
	if desiredSpec != nil {
		presentFields = CollectPresentFields(desiredSpec)
	} else {
		presentFields = sets.New[string]()
	}
	normPresent := normalizePresentFields(presentFields)

	// 3. For any computed field omitted in desiredSpec, copy from actualPb if its parent exists
	for _, path := range paths {
		lastDot := strings.LastIndex(path, ".")
		var parentPath, leafName string
		if lastDot == -1 {
			parentPath = ""
			leafName = path
		} else {
			parentPath = path[:lastDot]
			leafName = path[lastDot+1:]
		}

		pair, ok := parentMap[parentPath]
		if !ok {
			continue
		}

		fd := FindProtoField(pair.Actual.Descriptor(), leafName)
		if fd == nil {
			klog.V(0).Infof("internal error: field %q not found on proto message %s", leafName, pair.Actual.Descriptor().FullName())
			continue
		}
		copyComputedField(pair.Desired, pair.Actual, fd, normalizeFieldPath(parentPath), normPresent)
	}
}

// normalizeFieldPath converts a dot-separated KRM or proto JSON path into a canonical
// lowercase path with any trailing "Ref" stripped from each segment.
func normalizeFieldPath(path string) string {
	if path == "" {
		return ""
	}
	segs := strings.Split(path, ".")
	for i, seg := range segs {
		segs[i] = strings.ToLower(strings.TrimSuffix(seg, "Ref"))
	}
	return strings.Join(segs, ".")
}

// normalizePresentFields converts a set of KRM field paths (from CollectPresentFields)
// into a set of canonical normalized paths for O(1) lookup against protobuf descriptors.
func normalizePresentFields(present sets.Set[string]) sets.Set[string] {
	norm := sets.New[string]()
	for p := range present {
		norm.Insert(normalizeFieldPath(p))
	}
	return norm
}

// copyComputedField copies a computed field (or recursively merges a computed composite message)
// from actualParent to desiredParent if omitted from normPresent.
func copyComputedField(
	desiredParent, actualParent protoreflect.Message,
	fd protoreflect.FieldDescriptor,
	normParentPath string,
	normPresent sets.Set[string],
) {
	if !actualParent.Has(fd) {
		return
	}

	leafNorm := strings.ToLower(fd.JSONName())
	fieldNormPath := leafNorm
	if normParentPath != "" {
		fieldNormPath = normParentPath + "." + leafNorm
	}

	// If the field is a non-map, non-list composite message:
	if fd.Kind() == protoreflect.MessageKind && !fd.IsMap() && !fd.IsList() {
		if !normPresent.Has(fieldNormPath) {
			// Entire sub-message was omitted in desired spec; copy whole message from actual.
			desiredParent.Set(fd, actualParent.Get(fd))
			return
		}
		// Sub-message was partially specified in desired spec; recurse into its fields.
		if !desiredParent.Has(fd) {
			desiredParent.Set(fd, desiredParent.NewField(fd))
		}
		desiredSub := desiredParent.Get(fd).Message()
		actualSub := actualParent.Get(fd).Message()
		fields := actualSub.Descriptor().Fields()
		for i := 0; i < fields.Len(); i++ {
			copyComputedField(desiredSub, actualSub, fields.Get(i), fieldNormPath, normPresent)
		}
		return
	}

	// Scalar, enum, list, or map field: copy only if omitted in desired spec.
	if !normPresent.Has(fieldNormPath) {
		desiredParent.Set(fd, actualParent.Get(fd))
	}
}
