//go:build !without_kustomize
// +build !without_kustomize

package kustomize

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

// Run calls the kustomize/api library to run `kustomize build`. This method is differentiated by go build
// tag `without_kustomize`
func Run(ctx context.Context, fs filesys.FileSystem, manifestPath string) ([]byte, error) {
	log := log.FromContext(ctx)
	log.WithValues("manifestPath", manifestPath).Info("running kustomize")
	// run kustomize to create final manifest
	opts := krusty.MakeDefaultOptions()
	kustomizer := krusty.MakeKustomizer(opts)
	m, err := kustomizer.Run(fs, manifestPath)
	if err != nil {
		return nil, fmt.Errorf("error running kustomize: %v", err)
	}

	manifestYaml, err := m.AsYaml()
	if err != nil {
		return nil, fmt.Errorf("error converting kustomize output to yaml: %v", err)
	}
	return manifestYaml, nil
}
