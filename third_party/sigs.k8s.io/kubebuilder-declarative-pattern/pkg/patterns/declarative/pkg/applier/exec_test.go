//go:build !without_exec_applier
// +build !without_exec_applier

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package applier

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

// collector is a commandSite implementation that stubs cmd.Run() calls for tests
type collector struct {
	Error error
	Cmds  []*exec.Cmd
}

func (s *collector) Run(c *exec.Cmd) error {
	s.Cmds = append(s.Cmds, c)
	return s.Error
}

func TestKubectlApply(t *testing.T) {
	configMapYAML := `
apiVersion: v1
kind: ConfigMap
metadata:
  name: foo
`
	configMapJSON := `{"apiVersion":"v1","kind":"ConfigMap","metadata":{"name":"foo"}}`

	tests := []struct {
		name        string
		namespace   string
		manifest    string
		validate    bool
		args        []string
		err         error
		expectStdin string
		expectArgs  []string
	}{
		{
			name:        "manifest",
			namespace:   "",
			manifest:    configMapYAML,
			expectStdin: configMapJSON,
			expectArgs:  []string{"kubectl", "apply", "--validate=false", "-f", "-"},
		},
		{
			name:        "manifest with apply",
			namespace:   "kube-system",
			manifest:    configMapYAML,
			expectStdin: configMapJSON,
			expectArgs:  []string{"kubectl", "apply", "-n", "kube-system", "--validate=false", "-f", "-"},
		},
		{
			name:        "manifest with validate",
			namespace:   "",
			manifest:    configMapYAML,
			expectStdin: configMapJSON,
			validate:    true,
			expectArgs:  []string{"kubectl", "apply", "--validate=true", "-f", "-"},
		},
		{
			name:       "error propagation",
			expectArgs: []string{"kubectl", "apply", "--validate=false", "-f", "-"},
			err:        errors.New("error"),
		},
		{
			name:        "manifest with prune",
			namespace:   "kube-system",
			manifest:    configMapYAML,
			expectStdin: configMapJSON,
			args:        []string{"--prune=true", "--prune-whitelist=hello-world"},
			expectArgs:  []string{"kubectl", "apply", "-n", "kube-system", "--validate=false", "--prune=true", "--prune-whitelist=hello-world", "-f", "-"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()

			cs := collector{Error: test.err}
			kubectl := &ExecKubectl{cmdSite: &cs}

			objects, err := manifest.ParseObjects(ctx, test.manifest)
			if err != nil {
				t.Fatalf("error parsing manifest: %v", err)
			}

			opts := ApplierOptions{
				Namespace: test.namespace,
				Objects:   objects.GetItems(),
				Validate:  test.validate,
				ExtraArgs: test.args,
			}
			err = kubectl.Apply(ctx, opts)

			if test.err != nil && err == nil {
				t.Error("expected error to occur")
			} else if test.err == nil && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(cs.Cmds) != 1 {
				t.Errorf("expected 1 command to be invoked, got: %d", len(cs.Cmds))
			}

			cmd := cs.Cmds[0]
			if !reflect.DeepEqual(cmd.Args, test.expectArgs) {
				t.Errorf("argument mistmatch, expected: %v, got: %v", test.expectArgs, cmd.Args)
			}

			stdinBytes, err := io.ReadAll(cmd.Stdin)
			if err != nil {
				t.Fatalf("error reading manifest from stdin: %v", err)
			}
			if got, want := strings.TrimSpace(string(stdinBytes)), test.expectStdin; got != want {
				t.Errorf("manifest mismatch: got: %v, want: %v", got, want)
				t.Errorf("diff=%v", cmp.Diff(got, want))
			}
		})
	}
}

func TestKubectlApplier(t *testing.T) {
	log.SetLogger(klogr.New())

	kubectlPath, err := exec.LookPath("kubectl")
	if err != nil {
		t.Fatalf("failed to find kubectl on path: %v", err)
	}
	t.Logf("kubectl found at %q", kubectlPath)

	kubectlVersion, err := exec.Command("kubectl", "version", "--client").CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run kubectl version: %v", err)
	}
	t.Logf("kubectl version is %q", kubectlVersion)

	applier := NewExec()
	runApplierGoldenTests(t, "testdata/kubectl", true, applier)
}
