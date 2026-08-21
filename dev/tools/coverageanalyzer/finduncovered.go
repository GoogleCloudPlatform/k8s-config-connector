// Copyright 2025 Google LLC
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

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/cover"
)

// UncoveredBranch represents a specific branch that wasn't tested.
type UncoveredBranch struct {
	FileName    string
	BranchType  string
	StartLine   int
	EndLine     int
	Description string
	Content     string // The actual source code of the branch
}

// findUncoveredBlocks parses the coverage file and returns a map of
// fileName -> []cover.ProfileBlock for blocks with a count of 0.
func findUncoveredBlocks(file string) (map[string][]cover.ProfileBlock, error) {
	profiles, err := cover.ParseProfiles(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse coverage profile: %w", err)
	}

	uncovered := make(map[string][]cover.ProfileBlock)
	for _, p := range profiles {
		// We need to resolve the file path relative to the module root.
		// coverage.out uses package paths, not system paths.
		// This is a simplification; a real tool might need to find the module root.
		fileName := p.FileName
		if !filepath.IsAbs(fileName) {
			// It assumes the tool is run from the module root.
			fileName = resolvePath(p.FileName)
		}

		for _, b := range p.Blocks {
			if b.Count == 0 {
				uncovered[fileName] = append(uncovered[fileName], b)
			}
		}
	}
	return uncovered, nil
}

func findBranchDetails(uncoveredBlocks map[string][]cover.ProfileBlock) []UncoveredBranch {
	var branches []UncoveredBranch
	fset := token.NewFileSet()

	for fileName, blocks := range uncoveredBlocks {
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Printf("Could not read source file %s: %v", fileName, err)
			continue
		}

		node, err := parser.ParseFile(fset, fileName, content, 0)
		if err != nil {
			log.Printf("Could not parse source file %s: %v", fileName, err)
			continue
		}

		ast.Inspect(node, func(n ast.Node) bool {
			if n == nil {
				return true
			}

			for i, block := range blocks {
				// Avoid reprocessing a block
				if block.StartLine == 0 {
					continue
				}

				startPos := fset.File(n.Pos()).Line(n.Pos())
				endPos := fset.File(n.End()).Line(n.End())

				if isInside(startPos, endPos, block.StartLine, block.EndLine) {
					// Pass file content to the analyzer
					branch := analyzeAndFilterNode(n, block, fileName, fset, content)
					if branch != nil {
						branches = append(branches, *branch)
					}
					blocks[i].StartLine = 0
					return false
				}
			}
			return true
		})
	}
	return branches
}

func analyzeAndFilterNode(n ast.Node, block cover.ProfileBlock, fileName string, fset *token.FileSet, content []byte) *UncoveredBranch {
	var stmts []ast.Stmt
	var bodyNode ast.Node
	var branchType, description string
	isFilterable := true

	switch node := n.(type) {
	case *ast.IfStmt:
		elseLine := 0
		if node.Else != nil {
			elseLine = fset.File(node.Else.Pos()).Line(node.Else.Pos())
		}
		if elseLine > 0 && block.StartLine >= elseLine {
			stmts, bodyNode = getBodyFromElse(node.Else)
			branchType = "Else Body"
			description = "The 'else' part of an if-statement was not executed."
		} else {
			stmts = node.Body.List
			bodyNode = node.Body
			branchType = "If Body"
			description = "The body of an if-statement was not executed."
		}
	case *ast.CaseClause:
		stmts = node.Body
		bodyNode = node
		branchType = "Switch Case"
		description = "A case in a switch statement was not executed."
	case *ast.FuncDecl:
		isFilterable = false
		bodyNode = node.Body
		branchType = "Function Body"
		description = fmt.Sprintf("The entire function '%s' was never called.", node.Name.Name)
	default:
		return nil
	}

	if isFilterable && isSkippableErrorHandlingBlock(stmts) {
		fmt.Printf("Skipping uncovered branch at %s:%d (reason: ignorable error handling)\n", fileName, block.StartLine)
		return nil // Filtered out!
	}

	var branchContent string
	if bodyNode != nil {
		branchContent = extractContent(bodyNode, fset, content)
	}

	return &UncoveredBranch{
		FileName:    fileName,
		BranchType:  branchType,
		StartLine:   fset.File(n.Pos()).Line(n.Pos()),
		EndLine:     fset.File(n.End()).Line(n.End()),
		Description: description,
		Content:     branchContent,
	}
}

// Helper to get body statements from an `else` clause, which could be `else if` or `else`.
func getBodyFromElse(elseStmt ast.Stmt) ([]ast.Stmt, ast.Node) {
	switch e := elseStmt.(type) {
	case *ast.BlockStmt: // This is a final `else { ... }`
		return e.List, e
	case *ast.IfStmt: // This is an `else if ...`
		return e.Body.List, e.Body
	}
	return nil, nil
}

// Helper function to extract source code from a node.
func extractContent(node ast.Node, fset *token.FileSet, content []byte) string {
	start := fset.File(node.Pos()).Offset(node.Pos())
	end := fset.File(node.End()).Offset(node.End())

	if start < 0 || end < 0 || start > end || end > len(content) {
		return "[could not extract content]"
	}

	return string(content[start:end])
}

// Helper to indent multi-line strings for pretty printing.
func indent(text, prefix string) string {
	return prefix + strings.ReplaceAll(text, "\n", "\n"+prefix)
}

// isSkippableErrorHandlingBlock returns true if all statements are either log calls or a `return ..., err`.
func isSkippableErrorHandlingBlock(stmts []ast.Stmt) bool {
	if len(stmts) == 0 {
		return false
	}

	for _, stmt := range stmts {
		// A statement is skippable if it is a log call, OR a simple `return err`,
		// OR a `return fmt.Errorf(...)`. If it's none of these, the block is not skippable.
		if !isLogStatement(stmt) && !isReturnErrStatement(stmt) && !isReturnFmtErrorfStatement(stmt) {
			return false
		}
	}

	// All statements in the block were of a skippable type.
	return true
}

// Helper to check if a statement is a `log.*` call.
func isLogStatement(stmt ast.Stmt) bool {
	// A log call is an expression statement.
	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return false
	}
	// The expression must be a function call.
	callExpr, ok := exprStmt.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	// The function call must be a selector (e.g., `pkg.Func`).
	selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	// The receiver of the selector must be an identifier.
	ident, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}
	// The identifier's name must be "log".
	return ident.Name == "log"
}

// Helper to check for a `return ..., err` statement.
func isReturnErrStatement(stmt ast.Stmt) bool {
	// The statement must be a return statement.
	ret, ok := stmt.(*ast.ReturnStmt)
	if !ok {
		return false
	}
	// It must return at least one value.
	if len(ret.Results) == 0 {
		return false
	}
	// The last returned value must be an identifier...
	lastResult := ret.Results[len(ret.Results)-1]
	ident, ok := lastResult.(*ast.Ident)
	if !ok {
		return false
	}
	// ...and that identifier's name must be "err".
	return ident.Name == "err"
}

// Helper to check for a `return ..., fmt.Errorf(...)` statement.
func isReturnFmtErrorfStatement(stmt ast.Stmt) bool {
	// The statement must be a return statement.
	ret, ok := stmt.(*ast.ReturnStmt)
	if !ok {
		return false
	}

	// We check if any of the returned expressions is a call to fmt.Errorf.
	for _, res := range ret.Results {
		// The result must be a function call.
		call, ok := res.(*ast.CallExpr)
		if !ok {
			continue // Not a call expression, check the next result
		}

		// The function being called must be a selector expression (like `pkg.Func`).
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		// The package identifier must be "fmt".
		if pkgIdent, ok := sel.X.(*ast.Ident); !ok || pkgIdent.Name != "fmt" {
			continue
		}

		// The function name identifier must be "Errorf".
		if sel.Sel.Name == "Errorf" {
			return true
		}
	}

	return false
}

// isInside checks if nodeStart/end is within blockStart/end.
func isInside(nodeStart, nodeEnd, blockStart, blockEnd int) bool {
	return nodeStart >= blockStart && nodeEnd <= blockEnd
}

// resolvePath is a helper to find the full system path from a Go package path.
// NOTE: This is a simplification. A robust tool would use `go list`.
func resolvePath(pkgPath string) string {
	for _, srcDir := range strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator)) {
		if srcDir == "" {
			continue
		}
		path := filepath.Join(srcDir, "src", pkgPath)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	// Fallback for Go Modules structure (often works when run from module root).
	return pkgPath
}
