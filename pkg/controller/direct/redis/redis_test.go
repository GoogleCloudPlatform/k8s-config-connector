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

package redis

import (
	"strconv"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func TestRedisClusterFuzzer(t *testing.T) {
	var fuzzer fuzztesting.KRMFuzzer = redisClusterFuzzer()
	for i := int64(0); i < 500; i++ {
		t.Run(strconv.FormatInt(i, 10), func(t *testing.T) {
			fuzzer.FuzzSpec(t, i)
		})
	}
}

func TestRedisInstanceFuzzer(t *testing.T) {
	var fuzzer fuzztesting.KRMFuzzer = redisInstanceFuzzer()
	for i := int64(0); i < 500; i++ {
		t.Run(strconv.FormatInt(i, 10), func(t *testing.T) {
			fuzzer.FuzzSpec(t, i)
		})
	}
}

func TestRedisClusterEndpointFuzzer(t *testing.T) {
	var fuzzer fuzztesting.KRMFuzzer = redisClusterEndpointFuzzer()
	for i := int64(0); i < 500; i++ {
		t.Run(strconv.FormatInt(i, 10), func(t *testing.T) {
			fuzzer.FuzzSpec(t, i)
		})
	}
}
