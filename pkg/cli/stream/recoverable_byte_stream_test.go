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

package stream

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestRecoverableByteStream(t *testing.T) {
	ctx := context.TODO()

	s := &panicStream{}
	recoverableStream := NewRecoverableByteStream(s)
	_, _, err := recoverableStream.Next(ctx)
	if err == nil {
		t.Fatalf("got nil, but expect to have an error")
	}
}

type panicStream struct{}

func (s *panicStream) Next(_ context.Context) (bytes []byte, u *unstructured.Unstructured, err error) {
	panic("intentionally panic")
}
