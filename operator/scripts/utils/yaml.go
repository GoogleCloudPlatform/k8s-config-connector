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

package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ghodss/yaml" //nolint:depguard
	goyaml "gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func UnstructToYaml(u *unstructured.Unstructured) ([]byte, error) {
	bytes, err := yaml.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("error marshalling unstruct to yaml: %w", err)
	}
	return bytes, nil
}

func BytesToUnstruct(bytes []byte) (*unstructured.Unstructured, error) {
	u := unstructured.Unstructured{}
	if err := yaml.Unmarshal(bytes, &u); err != nil {
		return nil, fmt.Errorf("error unmarshalling bytes to unstruct: %w", err)
	}
	return &u, nil
}

func ReadFileToUnstructs(filePath string) ([]*unstructured.Unstructured, error) {
	var returnUnstructs []*unstructured.Unstructured
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	yamls, err := SplitYAML(b)
	if err != nil {
		return nil, err
	}
	for _, b = range yamls {
		u, err := BytesToUnstruct(b)
		if err != nil {
			return nil, err
		}
		returnUnstructs = append(returnUnstructs, u)
	}
	return returnUnstructs, nil
}

func SplitYAML(yamlBytes []byte) ([][]byte, error) {
	r := bytes.NewReader(yamlBytes)
	dec := goyaml.NewDecoder(r)
	results := make([][]byte, 0)
	for {
		var value map[string]interface{}
		err := dec.Decode(&value)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("error decoding yaml: %w", err)
		}

		bytes, err := goyaml.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("error marshalling '%v' to YAML: %w", value, err)
		}
		results = append(results, bytes)
	}
	return results, nil
}
