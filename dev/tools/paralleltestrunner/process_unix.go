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

//go:build !windows

package main

import (
	"os/exec"
	"syscall"
)

// setProcessGroup configures the Cmd to start in a new process group on Unix/Linux systems,
// and sets a custom Cancel function so that when the Context is cancelled (due to timeout),
// SIGKILL is sent to the entire process group (-PID) rather than just the parent process PID.
// This ensures any nested subprocesses spawned during the test are terminated, preventing
// them from running indefinitely and keeping stdout/stderr pipes open.
func setProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
}
