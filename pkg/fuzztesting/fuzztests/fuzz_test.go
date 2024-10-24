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

package fuzztesting

import (
	"math/rand"
	"testing"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func FuzzAllMappers(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		fuzzer := fuzztesting.ChooseFuzzer(randStream.Int63())
		nextSeed := randStream.Int63()
		fuzzer(t, nextSeed)
	})
}
