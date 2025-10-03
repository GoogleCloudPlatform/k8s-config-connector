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

	"google.golang.org/protobuf/proto"
)

// ComputeWeakEtag computes the etag of the proto object with weak indicator.
func ComputeWeakEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}

	h := sha256.Sum256(b)
	str := base64.StdEncoding.EncodeToString(h[:])
	return fmt.Sprintf(`W/"%s"`, str) // ETag must be quoted.
}
