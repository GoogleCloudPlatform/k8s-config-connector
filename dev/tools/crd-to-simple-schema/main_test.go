package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

func TestRun(t *testing.T) {
	dirs, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatalf("reading testdata: %v", err)
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		t.Run(d.Name(), func(t *testing.T) {
			dir := filepath.Join("testdata", d.Name())

			crd1 := filepath.Join(dir, "crd1.yaml")
			crd2 := filepath.Join(dir, "crd2.yaml")
			outputFile := filepath.Join(dir, "_output.txt")

			var opt ConvertOptions
			opt.CRDFile = crd1
			opt.DiffCRDFile = crd2
			opt.Flatten = true

			var buf bytes.Buffer
			if err := Run(t.Context(), opt, &buf); err != nil {
				t.Fatalf("Run failed: %v", err)
			}

			test.CompareGoldenFile(t, outputFile, buf.String())
		})
	}
}
