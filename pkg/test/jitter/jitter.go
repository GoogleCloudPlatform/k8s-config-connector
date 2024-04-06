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

package testjitter

import (
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// static jitter generator used for testing
type TestJitterGenerator struct {
}

// JitteredReenqueue implements jitter.Generator.
func (*TestJitterGenerator) JitteredReenqueue(gvk schema.GroupVersionKind, obj v1.Object) (time.Duration, error) {
	return time.Second, nil
}

// WatchJitteredTimeout implements jitter.Generator.
func (*TestJitterGenerator) WatchJitteredTimeout() time.Duration {
	return time.Second
}

var _ jitter.Generator = &TestJitterGenerator{}
