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

package ha

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildClusterHandles(t *testing.T) {
	useTestKubeconfig(t, "cluster-a")
	contexts := []string{"missing-1", "missing-2"}

	_, err := buildClusterHandles(context.Background(), contexts)

	if err == nil {
		t.Fatalf("expected error for missing contexts, got nil")
	}
	if !strings.Contains(err.Error(), "missing-1") || !strings.Contains(err.Error(), "missing-2") {
		t.Fatalf("expected accumulated errors for both missing-1 and missing-2, got %v", err)
	}
}

func useTestKubeconfig(t *testing.T, contextNames ...string) {
	t.Helper()
	var b strings.Builder
	b.WriteString("apiVersion: v1\nkind: Config\nclusters:\n")
	for _, name := range contextNames {
		b.WriteString("- name: " + name + "\n  cluster:\n    server: https://127.0.0.1:6443\n")
	}
	b.WriteString("contexts:\n")
	for _, name := range contextNames {
		b.WriteString("- name: " + name + "\n  context:\n    cluster: " + name + "\n")
	}
	path := filepath.Join(t.TempDir(), "kubeconfig")
	if err := os.WriteFile(path, []byte(b.String()), 0600); err != nil {
		t.Fatalf("writing test kubeconfig: %v", err)
	}
	t.Setenv("KUBECONFIG", path)
	t.Setenv("KUBECACHEDIR", filepath.Join(t.TempDir(), "cache"))
}
