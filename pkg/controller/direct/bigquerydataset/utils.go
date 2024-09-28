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

package bigquerydataset

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func convertProtoToAPI(u protoreflect.ProtoMessage, v any) error {
	if u == nil {
		return nil
	}

	j, err := protojson.Marshal(u)
	if err != nil {
		return fmt.Errorf("converting proto to json: %w", err)
	}

	if err := json.Unmarshal(j, v); err != nil {
		return fmt.Errorf("converting json to cloud API type: %w", err)
	}
	return nil
}

func convertAPIToProto[V protoreflect.ProtoMessage](u any, pV *V) error {
	if u == nil {
		return nil
	}

	j, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("converting proto to json: %w", err)
	}

	var v V
	if err := json.Unmarshal(j, &v); err != nil {
		return fmt.Errorf("converting json to proto type: %w", err)
	}
	*pV = v
	return nil
}
