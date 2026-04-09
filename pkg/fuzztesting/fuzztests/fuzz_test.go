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
	"os"
	"strconv"
	"testing"
	"time"

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

func TestSomeMappers(t *testing.T) {
	seed := time.Now().UnixNano()
	if s := os.Getenv("FUZZ_SEED"); s != "" {
		var err error
		seed, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			t.Fatalf("error parsing FUZZ_SEED %q: %v", s, err)
		}
	}
	t.Logf("Random seed: %d", seed)
	t.Logf("Number of registered fuzzers: %d", fuzztesting.GetFuzzerCount())

	iterations := 100000
	if os.Getenv("CI") != "" {
		iterations = 200000
	}

	randStream := rand.New(rand.NewSource(seed))

	for i := 0; i < iterations; i++ {
		fuzzer := fuzztesting.ChooseFuzzer(randStream.Int63())
		nextSeed := randStream.Int63()
		fuzzer(t, nextSeed)
	}
}
