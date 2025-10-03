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

package mockcloudidentity

import (
	"fmt"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func PtrTo[T any](t T) *T {
	return &t
}

func ValueOf[T any](p *T) T {
	var v T
	if p != nil {
		v = *p
	}
	return v
}

func buildLRO(obj proto.Message) (*longrunning.Operation, error) {
	responseAny, err := anypb.New(obj)
	if err != nil {
		return nil, fmt.Errorf("error building anypb for response: %w", err)
	}
	lro := &longrunning.Operation{}
	lro.Done = true
	lro.Result = &longrunning.Operation_Response{
		Response: responseAny,
	}
	return lro, nil
}
