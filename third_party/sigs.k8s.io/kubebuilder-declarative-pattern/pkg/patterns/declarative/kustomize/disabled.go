//go:build without_kustomize
// +build without_kustomize

package kustomize

import (
	"context"
	"fmt"

	"sigs.k8s.io/kustomize/kyaml/filesys"
)

// Run ignore the `kustomization.yaml` file and won't run `kustomize build`.
func Run(_ context.Context, _ filesys.FileSystem, _ string) ([]byte, error) {
	return nil, fmt.Errorf("kustomize support is not compiled in (built with tag `without_kustomize`)")
}
