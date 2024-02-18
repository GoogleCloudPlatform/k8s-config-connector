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

package parameters

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/valutil"
)

const (
	OutputFormatParam = "output-format"
	TableOutputFormat = "table"
	JSONOutputFormat  = "json"
	YAMLOutputFormat  = "yaml"

	DefaultOutputFormat = TableOutputFormat
)

type Parameters struct {
	OutputFormat string
}

func Validate(params *Parameters) error {
	return validateResourceFormat(params.OutputFormat)
}

func validateResourceFormat(value string) error {
	resourceFormatOptions := []string{TableOutputFormat, JSONOutputFormat, YAMLOutputFormat}
	if valutil.IsDefaultValue(value) {
		return fmt.Errorf("invalid empty value for %v: must be one of {%v}", OutputFormatParam, strings.Join(resourceFormatOptions, ", "))
	}
	for _, o := range resourceFormatOptions {
		if value == o {
			return nil
		}
	}
	return fmt.Errorf("invalid %v value of '%v': must be one of {%v}", OutputFormatParam, value, strings.Join(resourceFormatOptions, ", "))
}
