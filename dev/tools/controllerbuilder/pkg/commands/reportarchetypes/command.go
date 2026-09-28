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

// Package reportarchetypes implements the report-archetypes command, which
// classifies every kind mapped in apis/*/generate.sh by the shape of its proto
// API, using the model in pkg/protoapi.
package reportarchetypes

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/spf13/cobra"
)

// ReportArchetypesOptions configures the report-archetypes command.
type ReportArchetypesOptions struct {
	*options.GenerateOptions

	// APIsDir holds the <service>/generate.sh scripts to scan.
	APIsDir string
	// Output is the file to write the report to; stdout when empty.
	Output string
}

// InitDefaults points APIsDir at the repository's apis directory.
func (o *ReportArchetypesOptions) InitDefaults() error {
	root, err := options.RepoRoot()
	if err != nil {
		return err
	}
	o.APIsDir = filepath.Join(root, "apis")
	return nil
}

// BindFlags registers the command's own flags. The descriptor set comes from
// the shared --proto-source-path flag.
func (o *ReportArchetypesOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.APIsDir, "apis-dir", o.APIsDir, "directory holding the <service>/generate.sh scripts to scan")
	cmd.Flags().StringVarP(&o.Output, "output", "o", o.Output, "file to write the TSV report to; stdout when empty")
}

// BuildCommand returns the report-archetypes command.
func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &ReportArchetypesOptions{
		GenerateOptions: baseOptions,
	}

	if err := opt.InitDefaults(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing defaults: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "report-archetypes",
		Short: "classify every kind mapped in apis/*/generate.sh by the shape of its proto API",
		Long: `report-archetypes reads the generate-types invocations in apis/*/generate.sh,
looks up each kind's proto message in the descriptor set, and writes one TSV row
per kind: its archetype, the service hosting its standard methods and that
service's default host, its resource pattern and the shape of its standard
methods. Per-archetype counts go to stderr.

The result depends on the descriptor set, so pass the one for the googleapis
version pinned in apis/git.versions, e.g.
  --proto-source-path .build/googleapis-<sha>.pb
The protoSource column names the descriptor set generate.sh builds a kind
against when that is not the default one.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunReportArchetypes(opt)
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

// RunReportArchetypes writes the report and prints per-archetype counts to
// stderr.
func RunReportArchetypes(o *ReportArchetypesOptions) error {
	if o.ProtoSourcePath == "" {
		return fmt.Errorf("`--proto-source-path` is required")
	}
	api, err := protoapi.LoadProto(o.ProtoSourcePath, o.ProtoOverlayPath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}
	rows, err := Report(api, o.APIsDir)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err := WriteTSV(&out, rows); err != nil {
		return err
	}
	if o.Output == "" {
		if _, err := os.Stdout.Write(out.Bytes()); err != nil {
			return err
		}
	} else if err := os.WriteFile(o.Output, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("writing %s: %w", o.Output, err)
	}

	fmt.Fprintf(os.Stderr, "descriptor set: %s\n", o.ProtoSourcePath)
	return writeCounts(os.Stderr, rows)
}

// Report scans apisDir/*/generate.sh and classifies every kind it maps.
func Report(api *protoapi.Proto, apisDir string) ([]Row, error) {
	mappings, err := ScanGenerateScripts(apisDir)
	if err != nil {
		return nil, fmt.Errorf("scanning generate scripts: %w", err)
	}
	return BuildReport(api, mappings), nil
}

func writeCounts(w io.Writer, rows []Row) error {
	counts := CountByArchetype(rows)
	for _, a := range protoapi.AllArchetypes() {
		if _, err := fmt.Fprintf(w, "%s %-20s %4d\n", a, a.Name(), counts[a]); err != nil {
			return err
		}
	}
	_, err := fmt.Fprintf(w, "  %-20s %4d\n", "total", len(rows))
	return err
}
