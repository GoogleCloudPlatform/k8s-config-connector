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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/execution"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// RecoverableByteStream stream wraps a ByteStream and can recover from panics on Next().
// It should be used to wrap the top level stream so that it can catch panics from any nested stream.
type RecoverableByteStream struct {
	byteStream ByteStream
}

func NewRecoverableByteStream(inputStream ByteStream) *RecoverableByteStream {
	r := RecoverableByteStream{}
	r.byteStream = inputStream
	return &r
}

func (s *RecoverableByteStream) Next(ctx context.Context) (bytes []byte, unstructured *unstructured.Unstructured, err error) {
	defer execution.RecoverWithGenericError(&err)
	bytes, unstructured, err = s.byteStream.Next(ctx)
	return bytes, unstructured, err
}
