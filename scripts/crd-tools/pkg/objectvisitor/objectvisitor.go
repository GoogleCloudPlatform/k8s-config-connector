// Copyright 2024 Google LLC
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

package objectvisitor

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	kccyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/yaml"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

type Visitor interface {
	VisitObject(obj *unstructured.Unstructured) error
}

func VisitObjectsInDirectory(ctx context.Context, dir string, visitor Visitor) error {
	if err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(p) != ".yaml" {
			return nil
		}

		if err := processFile(ctx, p, visitor); err != nil {
			return fmt.Errorf("processing file %q: %w", p, err)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func VisitObjectsFromStdin(ctx context.Context, visitor Visitor) error {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("reading from stdin: %w", err)
	}
	out, err := processBytes(ctx, b, visitor)
	if err != nil {
		return fmt.Errorf("processing stdin: %w", err)
	}

	if _, err := os.Stdout.Write(out); err != nil {
		return fmt.Errorf("error writing to stdout: %w", err)
	}

	return nil
}

func processFile(ctx context.Context, p string, visitor Visitor) error {
	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", p, err)
	}

	out, err := processBytes(ctx, b, visitor)
	if err != nil {
		return fmt.Errorf("processing file %q: %w", p, err)
	}

	if err := os.WriteFile(p, out, 0644); err != nil {
		return fmt.Errorf("error writing file %q: %w", p, err)
	}
	return nil
}

func processBytes(ctx context.Context, b []byte, visitor Visitor) ([]byte, error) {
	var out bytes.Buffer

	// We preserve the yaml header (copyright, typically)
	for _, line := range bytes.Split(b, []byte("\n")) {
		if len(line) == 0 || line[0] == '#' {
			out.Write(line)
			out.WriteString("\n")
		} else {
			break
		}
	}

	yamls, err := kccyaml.SplitYAML(b)
	if err != nil {
		return nil, fmt.Errorf("error splitting bytes into YAMLs: %w", err)
	}
	objects := make([]*unstructured.Unstructured, 0)
	for _, y := range yamls {
		crd := &unstructured.Unstructured{}
		if err := yaml.Unmarshal(y, crd); err != nil {
			return nil, fmt.Errorf("error unmarshalling bytes to object: %w", err)
		}
		objects = append(objects, crd)
	}

	for _, obj := range objects {
		if err := visitor.VisitObject(obj); err != nil {
			return nil, err
		}
	}

	for i, obj := range objects {
		b, err := yaml.Marshal(obj)
		if err != nil {
			return nil, fmt.Errorf("error marshalling object to bytes: %w", err)
		}
		out.Write(b)
		if i != 0 {
			out.WriteString("\n---\n")
		}
	}
	return out.Bytes(), nil
}

func PtrTo[T any](t T) *T {
	return &t
}
