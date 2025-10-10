// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you under the Apache License, Version 2.0 (the "License");
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
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// HashProto calculates a hash of a proto message.
// We use this to detect changes to the GCP resource.
func HashProto(obj proto.Message) (string, error) {
	// We normalize the proto by clearing output-only fields etc
	// We do this on a copy
	obj = proto.Clone(obj)
	NormalizeProto(obj)

	// We use a deterministic proto marshaler.
	j, err := (proto.MarshalOptions{Deterministic: true}).Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("cannot marshal proto: %w", err)
	}

	h := sha256.Sum256(j)
	return hex.EncodeToString(h[:]), nil
}

// NormalizeProto clears fields that are not significant for comparison.
// It modifies the passed-in proto.
func NormalizeProto(pb proto.Message) {
	// TODO: Should we also clear fields like `uid` and `name`?
	// The problem is that they are not consistently marked as output-only.

	// We clear fields that are known to be volatile
	volatileFieldNames := []string{"etag", "update_time", "updated_time"}

	// We also clear fields that are marked as output-only
	// We do this by walking the fields and checking for `field_behavior: OUTPUT_ONLY`
	// We build a field mask and then use that to clear the fields.
	var paths []string
	pb.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		options := fd.Options()
		if options != nil {
			if proto.HasExtension(options, annotations.E_FieldBehavior) {
				fieldBehavior := proto.GetExtension(options, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
				for _, b := range fieldBehavior {
					if b == annotations.FieldBehavior_OUTPUT_ONLY {
						paths = append(paths, string(fd.Name()))
					}
				}
			}
		}
		for _, fieldName := range volatileFieldNames {
			if string(fd.Name()) == fieldName {
				paths = append(paths, string(fd.Name()))
			}
		}
		return true
	})

	if len(paths) > 0 {
		fm, err := fieldmaskpb.New(pb, paths...)
		if err != nil {
			// This should not happen
			panic(fmt.Sprintf("error creating fieldmask: %v", err))
		}
		fm.Normalize()
		if !fm.IsValid(pb) {
			// This should not happen
			panic(fmt.Sprintf("invalid fieldmask %v for %T", fm, pb))
		}
		clearFields(pb.ProtoReflect(), fm.GetPaths())
	}
}

// clearFields clears the given fields from the message.
func clearFields(m protoreflect.Message, paths []string) {
	for _, path := range paths {
		// Note: We don't support nested fields yet
		fd := m.Descriptor().Fields().ByName(protoreflect.Name(path))
		if fd == nil {
			// This should not happen
			panic(fmt.Sprintf("field %q not found in %v", path, m.Descriptor().FullName()))
		}
		m.Clear(fd)
	}
}

// Cookie is used for stateful reconciliation.
// It is stored in the status of the KCC resource.
type Cookie struct {
	SpecHash string `json:"specHash"`
	GCPHash  string `json:"gcpHash"`
}

// ComposeCookie creates a cookie string from the spec and gcp hashes.
func ComposeCookie(specHash, gcpHash string) (string, error) {
	cookie := &Cookie{
		SpecHash: specHash,
		GCPHash:  gcpHash,
	}
	b, err := json.Marshal(cookie)
	if err != nil {
		return "", fmt.Errorf("error marshalling cookie: %w", err)
	}
	return string(b), nil
}

// ParseCookie parses a cookie string.
func ParseCookie(s string) (*Cookie, error) {
	cookie := &Cookie{}
	if err := json.Unmarshal([]byte(s), cookie); err != nil {
		return nil, fmt.Errorf("error unmarshalling cookie: %w", err)
	}
	return cookie, nil
}
