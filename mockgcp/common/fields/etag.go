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

package fields

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

var mustFields = []string{
	"displayName",
	"state",
}

// ComputeEtag computes the etag of the proto object with weak indicator.
func ComputeEtag(obj proto.Message) string {
	pb := proto.Clone(obj)

	// ignore dynamic fields like timestampe or uniqueId.
	descriptor := pb.ProtoReflect().Descriptor()
	fieldDescs := descriptor.Fields()
	for i := 0; i < fieldDescs.Len(); i++ {
		fieldDesc := fieldDescs.Get(i)
		must := false
		for _, mustField := range mustFields {
			if fieldDesc.JSONName() == mustField {
				must = true
			}
		}
		if !must {
			pb.ProtoReflect().Clear(fieldDesc)
		}
	}

	m, err := prototext.Marshal(pb)
	if err != nil {
		panic(fmt.Sprintf("converting to prototext: %v", err))
	}
	h := sha256.Sum256([]byte(m))
	str := base64.StdEncoding.EncodeToString(h[:])
	strong := fmt.Sprintf(`"%s"`, str) // ETag must be quoted.
	return "W/" + strong
}

func ComputeEtagBytes(obj proto.Message) []byte {
	return []byte(ComputeEtag(obj))
}

func ComputeEtagPtr(obj proto.Message) *string {
	return ptrTo(ComputeEtag(obj))
}

func ptrTo[T any](t T) *T {
	return &t
}
