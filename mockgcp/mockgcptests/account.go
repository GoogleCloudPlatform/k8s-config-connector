// Copyright 2025 Google LLC
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

package mockgcptests

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func GetDefaultAccount(t *testing.T) string {
	t.Helper()

	account, err := getGCloudDefaultAccount()
	if err != nil {
		t.Fatal(err)
	}

	return account
}

func getGCloudDefaultAccount() (string, error) {
	cmd := exec.Command("gcloud", "config", "get-value", "account")
	bytes, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing command '%v': %w'", cmd, err)
	}
	return strings.TrimSpace(string(bytes)), nil
}
