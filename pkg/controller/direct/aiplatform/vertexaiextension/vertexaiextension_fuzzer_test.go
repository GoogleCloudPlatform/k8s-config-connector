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

package vertexaiextension

import (
	"testing"
	"time"
)

func TestVertexAIExtensionFuzzer(t *testing.T) {
	fuzzer := vertexAIExtensionFuzzer()
	seed := time.Now().UnixNano()

	t.Logf("Running fuzzer with seed %d", seed)
	for i := 0; i < 1000; i++ {
		fuzzer.FuzzSpec(t, seed+int64(i))
		fuzzer.FuzzStatus(t, seed+int64(i))
	}
}
