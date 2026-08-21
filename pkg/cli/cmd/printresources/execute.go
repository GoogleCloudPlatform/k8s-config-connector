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

package printresources

import (
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/printresources/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/printresources/printer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/printresources/resourcedescription"
)

func Execute(params *parameters.Parameters, output io.Writer) error {
	resourceDescs, err := resourcedescription.GetAll()
	if err != nil {
		return err
	}
	switch params.OutputFormat {
	case parameters.JSONOutputFormat:
		return printer.PrintJSON(resourceDescs, output)
	case parameters.TableOutputFormat:
		return printer.PrintTable(resourceDescs, output)
	case parameters.YAMLOutputFormat:
		return printer.PrintYAML(resourceDescs, output)
	default:
		return fmt.Errorf("unknown output-format '%v'", params.OutputFormat)
	}
}
