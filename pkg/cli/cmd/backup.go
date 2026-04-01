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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/backup"
	"github.com/spf13/cobra"
)

func NewBackupCmd() *cobra.Command {
	backupCmd := &cobra.Command{
		Use:   "backup",
		Short: "Backup and restore Config Connector resources",
		Long:  `Backup and restore Config Connector resources`,
		Args:  cobra.NoArgs,
	}

	backupCmd.AddCommand(backup.NewConfigureCmd())
	backupCmd.AddCommand(backup.NewCreateCmd())
	backupCmd.AddCommand(backup.NewStatusCmd())
	backupCmd.AddCommand(backup.NewRestoreCmd())

	return backupCmd
}
