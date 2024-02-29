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

package testcontroller

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Name returns a suitable name for an object based on the test name.
func Name(t *testing.T) string {
	return strings.ToLower(strings.TrimPrefix(t.Name(), "TestReconcile"))
}

// UniqueName takes a name and returns a unique version.
func UniqueName(_ *testing.T, name string) string {
	return fmt.Sprintf("%v-%v", name, rand.Intn(1000000))
}
