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

package mockbase

import (
	"encoding/json"
	"fmt"
)

func MarshalAny(data interface{}, typeName string) ([]byte, error) {
	// "response": {
	//     "@type": "type.googleapis.com/google.cloud.resourcemanager.v2.Folder",
	//     "name": "folders/######",
	//     "parent": "organizations/######",
	//     "displayName": "folder-test-1",
	//     "lifecycleState": "ACTIVE",
	//     "createTime": "2022-02-06T23:53:55.450Z"
	//   }
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error converting object to JSON: %w", err)
	}

	// We need to introduce @type; this is a little hacky
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("error deserializing JSON: %w", err)
	}
	m["@type"] = typeName

	b, err = json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error converting response data to JSON: %w", err)
	}

	return b, nil
}
