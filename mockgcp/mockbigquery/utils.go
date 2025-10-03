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

package mockbigquery

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func PtrTo[T any](t T) *T {
	return &t
}

func ValueOf[T any](t *T) T {
	var zeroVal T
	if t == nil {
		return zeroVal
	}
	return *t
}

func CloneProto[T protoreflect.ProtoMessage](t T) T {
	o := proto.Clone(t)
	return o.(T)
}
