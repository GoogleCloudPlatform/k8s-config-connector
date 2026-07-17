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

package changecookies

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func ComputeChangeCookie(desired, actual proto.Message) (string, error) {
	desiredHash, err := ComputeHash(desired)
	if err != nil {
		return "", err
	}
	actualHash, err := ComputeHash(actual)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("c/%s/%s", desiredHash, actualHash), nil
}

func ComputeHash(obj proto.Message) (string, error) {
	if obj == nil {
		return "", nil
	}

	// We use a deterministic proto marshaler.
	j, err := (proto.MarshalOptions{Deterministic: true}).Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("cannot marshal proto: %w", err)
	}

	h := sha1.Sum(j)
	return base64.URLEncoding.EncodeToString(h[:]), nil
}
