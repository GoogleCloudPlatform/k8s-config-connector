package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	rootDir := "."
	tmpDir := "/tmp"

	// Walk the files in rootDir, copy them to tmpDir
	if err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return fmt.Errorf("getting relative path for %q: %w", path, err)
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", path, err)
		}

		tmpPath := filepath.Join(tmpDir, relativePath)
		if err := os.MkdirAll(filepath.Dir(tmpPath), 0755); err != nil {
			return fmt.Errorf("creating directory %q: %w", filepath.Dir(tmpPath), err)
		}
		if err := os.WriteFile(tmpPath, b, 0644); err != nil {
			return fmt.Errorf("writing file %q: %w", tmpPath, err)
		}
	}); err != nil {
		return fmt.Errorf("walking rootDir dir %q: %w", rootDir, err)
	}

	return nil
}
