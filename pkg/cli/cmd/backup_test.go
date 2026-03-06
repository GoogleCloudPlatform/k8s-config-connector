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

package cmd

import (
	"testing"
)

func TestBackupCommandRegistration(t *testing.T) {
	root := rootCmd
	found := false
	for _, cmd := range root.Commands() {
		if cmd.Name() == "backup" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("backup command not found in root command")
	}
}

func TestBackupSubcommandsRegistration(t *testing.T) {
	backupCmd := NewBackupCmd()
	subcommands := []string{"configure", "create", "status"}
	for _, sub := range subcommands {
		found := false
		for _, cmd := range backupCmd.Commands() {
			if cmd.Name() == sub {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("subcommand %s not found in backup command", sub)
		}
	}
}
