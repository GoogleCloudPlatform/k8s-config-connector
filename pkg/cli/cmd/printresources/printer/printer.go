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

package printer

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/printresources/resourcedescription"

	"github.com/ghodss/yaml" //nolint:depguard
	"github.com/olekukonko/tablewriter"
)

func PrintTable(resourceDescs []resourcedescription.ResourceDescription, output io.Writer) error {
	tblWriter := tablewriter.NewWriter(output)
	tblWriter.SetHeader([]string{"Kind", "Bulk Export?", "Export?", "Supports IAM Export?", "Resource Name Format"})
	tblWriter.SetBorder(true)
	for _, rd := range resourceDescs {
		row := []string{
			rd.GVK.Kind,
			formatTableBool(rd.SupportsBulkExport),
			formatTableBool(rd.SupportsExport),
			formatTableBool(rd.SupportsIAM),
			rd.ResourceNameFormat,
		}
		tblWriter.Append(row)
	}
	tblWriter.Render()
	return nil
}

func formatTableBool(value bool) string {
	if value {
		return "x"
	}
	return ""
}

func PrintJSON(resourceDescs []resourcedescription.ResourceDescription, output io.Writer) error {
	bytes, err := json.MarshalIndent(resourceDescs, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling value to json: %w", err)
	}
	_, err = output.Write(bytes)
	return err
}

func PrintYAML(resourceDescs []resourcedescription.ResourceDescription, output io.Writer) error {
	bytes, err := yaml.Marshal(resourceDescs)
	if err != nil {
		return fmt.Errorf("error marshalling value to yaml: %w", err)
	}
	_, err = output.Write(bytes)
	return err
}
