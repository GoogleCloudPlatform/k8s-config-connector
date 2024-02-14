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

package yaml

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	goyaml "gopkg.in/yaml.v2"
)

func SplitYAML(yamlBytes []byte) ([][]byte, error) {
	r := bytes.NewReader(yamlBytes)
	dec := goyaml.NewDecoder(r)
	results := make([][]byte, 0)
	var value map[string]interface{}
	for eof := dec.Decode(&value); !errors.Is(eof, io.EOF); eof = dec.Decode(&value) {
		if eof != nil {
			return nil, eof
		}
		bytes, err := goyaml.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("error marshalling '%v' to YAML: %w", value, err)
		}
		results = append(results, bytes)
		value = make(map[string]interface{})
	}
	return results, nil
}
