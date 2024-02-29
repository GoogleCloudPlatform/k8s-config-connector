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

package test

import (
	"context"
	"testing"
)

type testKeyType int

// testKey is the key value for testing.T in a context.Context
var testKey testKeyType

func FromContext(ctx context.Context) *testing.T {
	v := ctx.Value(testKey)
	if v == nil {
		return nil
	}
	return v.(*testing.T)
}

func WithContext(ctx context.Context, t *testing.T) context.Context {
	return context.WithValue(ctx, testKey, t)
}
