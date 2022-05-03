// Copyright 2022 Google LLC
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

package asserts

import (
	"strings"
	"testing"
)

func AssertErrorIsExpected(t *testing.T, actual error, expect error) {
	if expect == nil {
		if actual != nil {
			t.Fatalf("expect to get nil err, but got %v", actual)
		}
	} else {
		if actual == nil {
			t.Fatalf("expect to get %v err, but got nil", expect)
		}
		if !strings.Contains(actual.Error(), expect.Error()) {
			t.Fatalf("expect to have error about %v, but got %v", expect, actual)
		}
	}
}
