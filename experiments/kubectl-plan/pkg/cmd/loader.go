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

package cmd

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kubectl-plan/pkg/plan"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// readObjects reads objects from the arguments supplied on the command line
func readObjects(ctx context.Context, cmd *cobra.Command, args []string) ([]*unstructured.Unstructured, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("must pass file or - (for stdin) with objects to apply")
	}
	if len(args) != 1 {
		return nil, fmt.Errorf("expected exactly 1 arg, got %d", len(args))
	}

	if args[0] == "-" {
		objs, err := plan.ParseObjects(cmd.InOrStdin())
		if err != nil {
			return nil, fmt.Errorf("error reading objects: %w", err)
		}
		return objs, nil
	} else {
		return plan.LoadObjectsFromFilesystem(args[0])
	}
}
