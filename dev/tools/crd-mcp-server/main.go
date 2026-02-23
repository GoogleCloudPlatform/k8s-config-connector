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
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"crd-compatibility-checker",
		"1.0.0",
	)

	s.AddTool(equivalenceTool(), handleCheckEquivalence)
	s.AddTool(backwardCompatTool(), handleCheckBackwardCompat)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func equivalenceTool() mcp.Tool {
	return mcp.NewTool("check_crd_equivalence",
		mcp.WithDescription(`Check whether a CRD file change is equivalent to its previous git-committed version.

Equivalent means:
  - No fields are added or deleted (new fields MAY be added under 'status')
  - Field names and types do not change
  - Adding spec.names.listKind is fine
  - Descriptions may change freely

The file is compared against the version stored at the given git ref (default: HEAD).`),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Path to the CRD YAML file to check."),
		),
		mcp.WithString("ref",
			mcp.Description("Git ref for the old version (default: HEAD)."),
		),
	)
}

func backwardCompatTool() mcp.Tool {
	return mcp.NewTool("check_crd_backward_compatibility",
		mcp.WithDescription(`Check whether a CRD file change is backward compatible with its previous git-committed version.

Backward compatible means:
  - No fields are removed or renamed
  - Field types do not change
  - New fields may be added anywhere
  - Descriptions may change freely

The file is compared against the version stored at the given git ref (default: HEAD).`),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Path to the CRD YAML file to check."),
		),
		mcp.WithString("ref",
			mcp.Description("Git ref for the old version (default: HEAD)."),
		),
	)
}

func handleCheckEquivalence(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	file, ref, err := parseArgs(req)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %v", err)), nil
	}

	result, err := runEquivalenceCheck(file, ref)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %v", err)), nil
	}
	return mcp.NewToolResultText(result), nil
}

func handleCheckBackwardCompat(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	file, ref, err := parseArgs(req)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %v", err)), nil
	}

	result, err := runBackwardCompatCheck(file, ref)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error: %v", err)), nil
	}
	return mcp.NewToolResultText(result), nil
}

func parseArgs(req mcp.CallToolRequest) (file, ref string, err error) {
	f, err := req.RequireString("file")
	if err != nil {
		return "", "", fmt.Errorf("'file' parameter is required: %w", err)
	}
	ref = req.GetString("ref", "HEAD")
	return f, ref, nil
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
