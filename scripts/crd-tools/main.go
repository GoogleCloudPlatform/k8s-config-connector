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

// This program parses CRDs found in a given YAML file and outputs them onto
// individual CRD files.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/cmd/deleteannotation"
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/cmd/deletefield"
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/cmd/reflowdescriptions"
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/cmd/removedescriptions"
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/cmd/setannotation"
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/cmd/setfield"
	"github.com/spf13/cobra"
)

func main() {
	err := run(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	rootCmd := &cobra.Command{
		Use: "crd-tools",
	}

	deleteannotation.AddCommand(rootCmd)
	deletefield.AddCommand(rootCmd)
	reflowdescriptions.AddCommand(rootCmd)
	removedescriptions.AddCommand(rootCmd)
	setannotation.AddCommand(rootCmd)
	setfield.AddCommand(rootCmd)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		return err
	}
	return nil
}
