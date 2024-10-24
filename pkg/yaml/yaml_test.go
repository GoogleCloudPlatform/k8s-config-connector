package yaml

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/google/go-cmp/cmp"
)

func TestSplitYAML(t *testing.T) {
	baseDir := "testdata/splityaml"

	var testDirs []string
	if err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			relPath, err := filepath.Rel(baseDir, path)
			if err != nil {
				return err
			}
			testDirs = append(testDirs, relPath)
		}
		return nil
	}); err != nil {
		t.Fatalf("error walking directory %q: %v", baseDir, err)
	}

	for _, testDir := range testDirs {
		t.Run(testDir, func(t *testing.T) {
			inPath := filepath.Join(baseDir, testDir, "in.yaml")
			outPath := filepath.Join(baseDir, testDir, "_expected.yaml")

			in := test.MustReadFile(t, inPath)

			out, err := SplitYAML(in)
			if err != nil {
				t.Fatalf("error from SplitYAML: %v", err)
			}

			gotBytes := bytes.Join(out, []byte("\n---\n\n"))
			got := strings.TrimSpace(string(gotBytes))
			CompareGoldenFile(t, outPath, got)
		})
	}
}

func CompareGoldenFile(h *testing.T, p string, got string, normalizers ...func(s string) string) {
	if os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		// Short-circuit when the output is correct
		b, err := os.ReadFile(p)
		if err == nil {
			want := string(b)
			for _, normalizer := range normalizers {
				got = normalizer(got)
				want = normalizer(want)
			}
			if want == got {
				return
			}
		}

		if err := os.WriteFile(p, []byte(got), 0644); err != nil {
			h.Fatalf("failed to write golden output %s: %v", p, err)
		}
		h.Errorf("wrote output to %s", p)
	} else {
		want := string(test.MustReadFile(h, p))

		for _, normalizer := range normalizers {
			got = normalizer(got)
			want = normalizer(want)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			h.Errorf("unexpected diff in %s: %s", p, diff)
		}
	}
}
