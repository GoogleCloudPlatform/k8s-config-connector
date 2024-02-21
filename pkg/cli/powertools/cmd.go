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

package powertools

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/changestateintospec"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/powertools/forcesetfield"
	"github.com/spf13/cobra"
)

func AddCommands(parent *cobra.Command) {
	powertoolsCmd := &cobra.Command{
		Use:   "powertools",
		Short: "Powertools holds our experimental / dangerous tools",
		Args:  cobra.NoArgs,
	}
	parent.AddCommand(powertoolsCmd)

	forcesetfield.AddCommand(powertoolsCmd)
	changestateintospec.AddCommand(powertoolsCmd)
}
