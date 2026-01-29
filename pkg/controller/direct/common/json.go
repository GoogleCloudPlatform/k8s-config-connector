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

package common

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtoToAPI maps a proto Message to a typed-go struct, assuming they have the same json schema
func ProtoToAPI[T any](in proto.Message, out *T) error {
	j, err := protojson.Marshal(in)
	if err != nil {
		return fmt.Errorf("converting proto %v to json: %w", in.ProtoReflect().Descriptor().FullName(), err)
	}
	if err := json.Unmarshal(j, out); err != nil {
		return fmt.Errorf("converting json to %T: %w", out, err)
	}
	return nil
}

// APIToProto maps a typed-go struct to a proto Message, assuming they have the same json schema
func APIToProto(in any, out proto.Message) error {
	j, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("converting api %T to json: %w", in, err)
	}
	if err := protojson.Unmarshal(j, out); err != nil {
		return fmt.Errorf("converting json to proto %v: %w", out.ProtoReflect().Descriptor().FullName(), err)
	}
	return nil
}
