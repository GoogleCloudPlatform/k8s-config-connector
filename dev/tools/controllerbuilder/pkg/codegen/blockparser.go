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

package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

type generatedBlock struct {
	name            string
	kind            string // "type" or "func"
	isCommented     bool
	protoFQN        string
	isObservedState bool
	krmTypeName     string
	rawText         string
	offset          int
}

type byteSpan struct {
	start int
	end   int
}

// parseGeneratedGoBlocks parses a generated Go source file (or body buffer when
// isBodyOnly is true) into ordered top-level type/func blocks (including
// commented-out blocks) and an import map keyed by package path.
func parseGeneratedGoBlocks(src []byte, defaultPkg string, isBodyOnly bool) ([]generatedBlock, map[string]string, error) {
	goSrc := src
	if isBodyOnly {
		header := fmt.Sprintf("package %s\n\n", defaultPkg)
		goSrc = append([]byte(header), src...)
	}

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", goSrc, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	imports := parseASTImports(astFile)
	blocks, declSpans := extractASTDeclBlocks(fset, astFile, goSrc)
	blocks = append(blocks, extractCommentedBlocks(fset, astFile, declSpans)...)

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].offset < blocks[j].offset
	})

	return blocks, imports, nil
}

func parseASTImports(astFile *ast.File) map[string]string {
	imports := make(map[string]string, len(astFile.Imports))
	for _, imp := range astFile.Imports {
		if imp.Path == nil {
			continue
		}

		pkgPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			continue
		}

		alias := ""
		if imp.Name != nil {
			alias = imp.Name.Name
		}

		imports[pkgPath] = alias
	}

	return imports
}

func extractASTDeclBlocks(fset *token.FileSet, astFile *ast.File, goSrc []byte) ([]generatedBlock, []byteSpan) {
	var blocks []generatedBlock
	var declSpans []byteSpan
	for _, decl := range astFile.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if blk, span, ok := buildTypeDeclBlock(fset, d, goSrc); ok {
				blocks = append(blocks, blk)
				declSpans = append(declSpans, span)
			}
		case *ast.FuncDecl:
			if blk, span, ok := buildFuncDeclBlock(fset, d, goSrc); ok {
				blocks = append(blocks, blk)
				declSpans = append(declSpans, span)
			}
		}
	}

	return blocks, declSpans
}

func buildTypeDeclBlock(fset *token.FileSet, d *ast.GenDecl, goSrc []byte) (generatedBlock, byteSpan, bool) {
	if d.Tok != token.TYPE || len(d.Specs) == 0 {
		return generatedBlock{}, byteSpan{}, false
	}

	typeSpec, ok := d.Specs[0].(*ast.TypeSpec)
	if !ok || typeSpec.Name == nil {
		return generatedBlock{}, byteSpan{}, false
	}

	declStart := fset.Position(d.Pos()).Offset
	docStart := declDocStartOffset(fset, d.Pos(), d.Doc)
	declEnd := fset.Position(d.End()).Offset

	rawText := strings.TrimSpace(string(goSrc[docStart:declEnd]))
	protoFQN, isObservedState := extractProtoInfoFromBlockText(rawText)
	if strings.HasSuffix(typeSpec.Name.Name, "ObservedState") {
		isObservedState = true
	}

	blk := generatedBlock{
		name:            typeSpec.Name.Name,
		kind:            "type",
		protoFQN:        protoFQN,
		isObservedState: isObservedState,
		rawText:         rawText,
		offset:          docStart,
	}

	return blk, byteSpan{start: declStart, end: declEnd}, true
}

func buildFuncDeclBlock(fset *token.FileSet, d *ast.FuncDecl, goSrc []byte) (generatedBlock, byteSpan, bool) {
	if d.Name == nil {
		return generatedBlock{}, byteSpan{}, false
	}

	declStart := fset.Position(d.Pos()).Offset
	docStart := declDocStartOffset(fset, d.Pos(), d.Doc)
	declEnd := fset.Position(d.End()).Offset

	rawText := strings.TrimSpace(string(goSrc[docStart:declEnd]))
	blk := generatedBlock{
		name:        d.Name.Name,
		kind:        "func",
		krmTypeName: extractKRMTypeFromFuncText(rawText),
		rawText:     rawText,
		offset:      docStart,
	}

	return blk, byteSpan{start: declStart, end: declEnd}, true
}

func extractKRMTypeFromFuncText(rawText string) string {
	for _, line := range strings.Split(rawText, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "func ") {
			return extractKRMTypeFromFuncSignature(trimmed)
		}
	}

	return ""
}

func declDocStartOffset(fset *token.FileSet, declPos token.Pos, doc *ast.CommentGroup) int {
	docStart := fset.Position(declPos).Offset
	if doc == nil || len(doc.List) == 0 {
		return docStart
	}

	firstLineCommentIdx := 0
	for idx, c := range doc.List {
		if strings.HasPrefix(c.Text, "/*") {
			firstLineCommentIdx = idx + 1
		}
	}

	if firstLineCommentIdx < len(doc.List) {
		return fset.Position(doc.List[firstLineCommentIdx].Pos()).Offset
	}

	return docStart
}

func extractCommentedBlocks(fset *token.FileSet, astFile *ast.File, declSpans []byteSpan) []generatedBlock {
	var blocks []generatedBlock
	pkgEndOffset := fset.Position(astFile.Name.End()).Offset
	for _, cg := range astFile.Comments {
		for _, c := range cg.List {
			if !strings.HasPrefix(c.Text, "/*") {
				continue
			}

			cOffset := fset.Position(c.Pos()).Offset
			if cOffset <= pkgEndOffset || isOffsetInsideSpans(cOffset, declSpans) {
				continue
			}

			if blk, ok := parseCommentedBlock(strings.TrimSpace(c.Text), cOffset); ok {
				blocks = append(blocks, blk)
			}
		}
	}

	return blocks
}

func isOffsetInsideSpans(offset int, spans []byteSpan) bool {
	for _, sp := range spans {
		if offset >= sp.start && offset < sp.end {
			return true
		}
	}

	return false
}

func parseCommentedBlock(rawText string, offset int) (generatedBlock, bool) {
	if name, protoFQN, isObs, ok := parseCommentedTypeBlock(rawText); ok {
		return generatedBlock{
			name:            name,
			kind:            "type",
			isCommented:     true,
			protoFQN:        protoFQN,
			isObservedState: isObs,
			rawText:         rawText,
			offset:          offset,
		}, true
	}

	if name, krmTypeName, ok := parseCommentedFuncBlock(rawText); ok {
		return generatedBlock{
			name:        name,
			kind:        "func",
			isCommented: true,
			krmTypeName: krmTypeName,
			rawText:     rawText,
			offset:      offset,
		}, true
	}

	return generatedBlock{}, false
}

func parseCommentedTypeBlock(rawText string) (typeName, protoFQN string, isObservedState, ok bool) {
	protoFQN, isObservedState = extractProtoInfoFromBlockText(rawText)
	for _, line := range strings.Split(rawText, "\n") {
		trimmed := strings.TrimSpace(line)
		if rest, found := strings.CutPrefix(trimmed, "/* unreachable type "); found {
			if fields := strings.Fields(rest); len(fields) > 0 {
				typeName = fields[0]
				break
			}
		}

		if name, found := goTypeNameFromDeclaration(trimmed); found {
			typeName = name
			break
		}
	}

	if typeName == "" {
		return "", "", false, false
	}

	if strings.HasSuffix(typeName, "ObservedState") {
		isObservedState = true
	}

	return typeName, protoFQN, isObservedState, true
}

func parseCommentedFuncBlock(rawText string) (funcName, krmTypeName string, ok bool) {
	for _, line := range strings.Split(rawText, "\n") {
		trimmed := strings.TrimSpace(line)
		if rest, found := strings.CutPrefix(trimmed, "func "); found {
			if name, _, hasParen := strings.Cut(rest, "("); hasParen && name != "" {
				return strings.TrimSpace(name), extractKRMTypeFromFuncSignature(trimmed), true
			}
		}
	}

	return "", "", false
}

func extractProtoInfoFromBlockText(rawText string) (protoFQN string, isObservedState bool) {
	for _, line := range strings.Split(rawText, "\n") {
		trimmed := strings.TrimSpace(line)
		if msg, kind, ok := GetProtoMessageAndKindFromAnnotation(trimmed); ok {
			return msg, kind == KCCProtoMessageAnnotationObservedState
		}
	}

	return "", false
}

func extractKRMTypeFromFuncSignature(sigLine string) string {
	rest, ok := strings.CutPrefix(sigLine, "func ")
	if !ok {
		return ""
	}

	rest = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(rest), "{"))
	funcName, afterName, ok := strings.Cut(rest, "(")
	if !ok {
		return ""
	}

	params, retType, ok := strings.Cut(afterName, ")")
	if !ok {
		return ""
	}

	if strings.HasSuffix(funcName, "_FromProto") {
		return extractKRMTypeFromFromProtoReturn(strings.TrimSpace(retType))
	}

	if strings.HasSuffix(funcName, "_ToProto") {
		return extractKRMTypeFromToProtoParams(funcName, params)
	}

	return ""
}

func extractKRMTypeFromFromProtoReturn(retType string) string {
	if _, krmType, found := strings.Cut(retType, "."); found {
		return strings.TrimSpace(krmType)
	}

	return ""
}

func extractKRMTypeFromToProtoParams(funcName, params string) string {
	if _, secondParam, found := strings.Cut(params, ","); found {
		if fields := strings.Fields(secondParam); len(fields) > 0 {
			if _, krmType, hasDot := strings.Cut(fields[len(fields)-1], "."); hasDot {
				return strings.TrimSpace(krmType)
			}
		}
	}

	base := strings.TrimSuffix(funcName, "_ToProto")
	if idx := strings.LastIndex(base, "_"); idx > 0 {
		return base[:idx]
	}

	return ""
}
