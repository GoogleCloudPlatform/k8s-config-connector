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

import "testing"

func TestHACmdNotRewrittenByBulkExportFallback(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "ha parent command",
			args: []string{"config-connector", "ha"},
		},
		{
			name: "ha setup subcommand",
			args: []string{"config-connector", "ha", "setup", "--contexts=cluster-a,cluster-b", "--bucket=my-bucket"},
		},
		{
			name: "ha status subcommand",
			args: []string{"config-connector", "ha", "status", "--contexts=cluster-a,cluster-b"},
		},
		{
			name: "ha failover subcommand",
			args: []string{"config-connector", "ha", "failover", "--contexts=cluster-a,cluster-b", "--target-cluster=cluster-b"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			args := tc.args

			got := isLegacyArgs(args)

			if got {
				t.Fatalf("isLegacyArgs(%v) = true, want false so ha command is not rewritten to bulk-export", args)
			}
		})
	}
}
