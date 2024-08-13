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

package resource

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/cel-go/common/types"
	"gopkg.in/yaml.v2"
)

type Resource struct {
	Name      string
	Data      map[string]interface{}
	Raw       []byte
	Variables []*Variable
}

func NewResourceFromRaw(name string, raw []byte) (*Resource, error) {
	var data map[string]interface{}
	err := yaml.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	variables, err := extractVariables(raw)
	if err != nil {
		return nil, err
	}

	resource := &Resource{
		Name:      name,
		Data:      data,
		Raw:       raw,
		Variables: variables,
	}
	return resource, nil
}

func (r *Resource) replaceVariables(vars map[string]string) []byte {
	copy := bytes.Clone(r.Raw)
	for expr, elem := range vars {
		trapRegex := regexp.MustCompile(regexExpression(expr))
		copy = trapRegex.ReplaceAll(copy, []byte(elem))
	}
	return copy
}

func (r *Resource) ApplyResolvedVariables() error {
	vars := make(map[string]string)
	for _, variable := range r.Variables {
		if variable.ResolvedValue != nil {
			switch v := variable.ResolvedValue.(type) {
			case string:
				vars[variable.Expression] = v
			case types.String:
				vars[variable.Expression] = v.Value().(string)
			case int, int64, int32, int16, int8:
				vars[variable.Expression] = fmt.Sprintf("%d", v)
			case types.Int:
				vars[variable.Expression] = fmt.Sprintf("%d", v.Value().(int))
			case float64, float32:
				vars[variable.Expression] = fmt.Sprintf("%f", v)
			case types.Double:
				vars[variable.Expression] = fmt.Sprintf("%f", v.Value().(float64))
			case bool:
				vars[variable.Expression] = fmt.Sprintf("%t", v)
			case types.Bool:
				vars[variable.Expression] = fmt.Sprintf("%t", v.Value().(bool))
			case types.Null:
				vars[variable.Expression] = ""
			default:
				return fmt.Errorf("unknown variable type: %v: type=%v", variable.ResolvedValue, v)
			}
			vv := vars[variable.Expression]
			if strings.Contains(vv, "\"") && strings.Contains(vv, "{") {
				// This is a json string, so we need to wrap it with single quotes and
				// escape the double quotes.
				vv = strings.ReplaceAll(vv, "\"", "\\\"")
				if !strings.HasPrefix(vv, "'") && !strings.HasSuffix(vv, "'") {
					// vv = "'" + vv + "'"
				}
				vars[variable.Expression] = vv
			}
		}
	}

	r.Raw = r.replaceVariables(vars)
	var newData map[string]interface{}
	err := yaml.Unmarshal(r.Raw, &newData)
	if err != nil {
		return err
	}
	r.Data = newData

	return nil
}
