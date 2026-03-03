// Copyright 2026 Google LLC
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

package migrate

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestMigrateGKEAddonCmd(t *testing.T) {
	// Verify command and subcommands are registered correctly
	root := &cobra.Command{Use: "config-connector"}
	AddCommands(root)

	migrateCmd, _, err := root.Find([]string{"gke-addon"})
	if err != nil {
		t.Fatalf("could not find gke-addon command: %v", err)
	}

	prepareCmd, _, err := migrateCmd.Find([]string{"prepare"})
	if err != nil {
		t.Errorf("could not find prepare subcommand: %v", err)
	}
	if prepareCmd.Name() != "prepare" {
		t.Errorf("expected prepare subcommand name to be 'prepare', got %q", prepareCmd.Name())
	}

	finishCmd, _, err := migrateCmd.Find([]string{"finish"})
	if err != nil {
		t.Errorf("could not find finish subcommand: %v", err)
	}
	if finishCmd.Name() != "finish" {
		t.Errorf("expected finish subcommand name to be 'finish', got %q", finishCmd.Name())
	}
}
