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

// crd-mcp-server is an MCP server that provides tools for analysing CRD changes.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

const (
	EQUIVALENCE_DESCRIPTION = `Check whether a CRD file change is equivalent to its previous git-committed version.

Equivalent means:
  - No fields are added or deleted (new fields MAY be added under 'status')
  - Field names and types do not change
  - Adding spec.names.listKind is fine
  - Descriptions may change freely

The file is compared against the version stored at the given git ref (default: HEAD).`
    BACKWARD_COMPAT_DESCRIPTION = `Check whether a CRD file change is backward compatible with its previous git-committed version.

Backward compatible means:
  - No fields are removed or renamed
  - Field types do not change
  - New fields may be added anywhere
  - Descriptions may change freely

The file is compared against the version stored at the given git ref (default: HEAD).`
)

type EquivalenceArguments struct {
	File string `json:"file" jsonschema:"required,description=Path to the CRD YAML file to check."`
	Ref  string `json:"ref" jsonschema:"description=Git ref for the old version (default: HEAD)."`
}

type BackwardCompatArguments struct {
	File string `json:"file" jsonschema:"required,description=Path to the CRD YAML file to check."`
	Ref  string `json:"ref" jsonschema:"description=Git ref for the old version (default: HEAD)."`
}

func main() {
	done := make(chan struct{})
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	if err := server.RegisterTool("check_crd_equivalence", EQUIVALENCE_DESCRIPTION, handleCheckEquivalence); err != nil {
		panic(err)
	}
	if err := server.RegisterTool("check_crd_backward_compatibility", BACKWARD_COMPAT_DESCRIPTION, handleCheckBackwardCompat); err != nil {
		panic(err)
	}

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-done
}

func handleCheckEquivalence(arguments EquivalenceArguments) (*mcp_golang.ToolResponse, error) {
	result, err := runEquivalenceCheck(arguments.File, arguments.Ref)
	if err != nil {
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
	}
	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(result)), nil
}

func handleCheckBackwardCompat(arguments BackwardCompatArguments) (*mcp_golang.ToolResponse, error) {
	result, err := runBackwardCompatCheck(arguments.File, arguments.Ref)
	if err != nil {
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
	}
	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(result)), nil
}

func runEquivalenceCheck(file, ref string) (string, error) {
	newData, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("reading %q: %w", file, err)
	}

	oldData, isNew, err := gitShow(ref, file)
	if err != nil {
		return "", fmt.Errorf("retrieving previous version from git: %w", err)
	}
	if isNew {
		return fmt.Sprintf("File is new (not present in %s). No previous version to compare against.", ref), nil
	}

	newCRD, err := parseCRD(newData)
	if err != nil {
		return "", fmt.Errorf("parsing new CRD: %w", err)
	}
	oldCRD, err := parseCRD(oldData)
	if err != nil {
		return "", fmt.Errorf("parsing old CRD from %s: %w", ref, err)
	}

	result := compareEquivalence(oldCRD, newCRD)
	return formatResult("EQUIVALENT", "NOT EQUIVALENT", result), nil
}

func runBackwardCompatCheck(file, ref string) (string, error) {
	newData, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("reading %q: %w", file, err)
	}

	oldData, isNew, err := gitShow(ref, file)
	if err != nil {
		return "", fmt.Errorf("retrieving previous version from git: %w", err)
	}
	if isNew {
		return fmt.Sprintf("File is new (not present in %s). No previous version to compare against.", ref), nil
	}

	newCRD, err := parseCRD(newData)
	if err != nil {
		return "", fmt.Errorf("parsing new CRD: %w", err)
	}
	oldCRD, err := parseCRD(oldData)
	if err != nil {
		return "", fmt.Errorf("parsing old CRD from %s: %w", ref, err)
	}

	result := compareBackwardCompatibility(oldCRD, newCRD)
	return formatResult("BACKWARD COMPATIBLE", "NOT BACKWARD COMPATIBLE", result), nil
}

func formatResult(okLabel, failLabel string, r CompareResult) string {
	var sb strings.Builder

	if len(r.Diffs) == 0 {
		sb.WriteString(okLabel + "\n")
	} else {
		sb.WriteString(failLabel + "\n\nDifferences:\n")
		for _, d := range r.Diffs {
			sb.WriteString("  - " + d + "\n")
		}
	}

	if len(r.Notes) > 0 {
		sb.WriteString("\nAllowed changes (informational):\n")
		for _, n := range r.Notes {
			sb.WriteString("  - " + n + "\n")
		}
	}

	return sb.String()
}
