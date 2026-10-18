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
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
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

type tokSpan struct {
	pos int
	end int
	tok token.Token
	lit string
}

type packageDeclarations struct {
	types     map[string]bool
	protoTags map[string]bool
	funcs     map[string]bool
}

// mergeWithExisting merges declarations, imports, and FileAnnotation attributes
// from an existing generated Go file on disk into f so sequential
// generate-types or generate-mapper invocations (for example across multiple
// --service packages) accumulate additively.
func (f *generatedFile) mergeWithExisting(existingBytes []byte, dir string) {
	matchingExistingAnn, compatible := findCompatibleFileAnnotation(existingBytes, f.fileAnnotation)
	if !compatible || !f.canMergeTargetFile(dir) {
		return
	}

	existingBlocks, existingImports, currentBlocks, ok := f.parseBlocksForMerge(existingBytes)
	if !ok {
		return
	}

	filteredExisting, hasActiveExternal := f.filterExistingBlocks(existingBlocks, currentBlocks, existingImports, dir)
	if !hasActiveExternal {
		return
	}

	if matchingExistingAnn != nil && f.fileAnnotation != nil {
		f.fileAnnotation = mergeFileAnnotations(matchingExistingAnn, f.fileAnnotation)
	}

	aliasRenames := f.mergeExistingImports(existingImports)
	rewriteBlockQualifiers(filteredExisting, aliasRenames)

	merged := mergeBlocksForFile(f.key.FileName, filteredExisting, currentBlocks)
	f.replaceBodyWithBlocks(merged)
}

func findCompatibleFileAnnotation(existingBytes []byte, current *annotations.FileAnnotation) (*annotations.FileAnnotation, bool) {
	if current == nil {
		return nil, true
	}
	existingAnns, err := annotations.FindFileAnnotations(existingBytes, []string{"+generated:"})
	if err != nil || len(existingAnns) == 0 {
		return nil, true
	}
	for i := range existingAnns {
		if existingAnns[i].Key == current.Key {
			if !hasCompatibleKRMGroup(&existingAnns[i], current) {
				return nil, false
			}
			annCopy := existingAnns[i]
			return &annCopy, true
		}
	}
	return nil, true
}

func hasCompatibleKRMGroup(existing, current *annotations.FileAnnotation) bool {
	eg := existing.Attributes["krm.group"]
	cg := current.Attributes["krm.group"]
	return len(eg) == 0 || len(cg) == 0 || eg[0] == cg[0]
}

func (f *generatedFile) canMergeTargetFile(dir string) bool {
	if f.key.FileName == "types.generated.go" && isTargetOfSeparateGeneratedFile(dir, f.key.FileName, f.fileAnnotation) {
		return false
	}
	return true
}

func (f *generatedFile) parseBlocksForMerge(existingBytes []byte) (existingBlocks []generatedBlock, existingImports map[string]string, currentBlocks []generatedBlock, ok bool) {
	pkgName := f.goPackage
	if pkgName == "" {
		pkgName = "p"
	}
	existingBlocks, existingImports, err := parseGeneratedGoBlocks(existingBytes, pkgName, false)
	if err != nil || len(existingBlocks) == 0 {
		return nil, nil, nil, false
	}
	currentBlocks, _, err = parseGeneratedGoBlocks(f.body.Bytes(), pkgName, true)
	if err != nil {
		return nil, nil, nil, false
	}
	if f.key.FileName == "mapper.generated.go" && len(currentBlocks) > 0 {
		if hasVersionedMapperFuncs(existingBlocks) != hasVersionedMapperFuncs(currentBlocks) {
			return nil, nil, nil, false
		}
	}
	return existingBlocks, existingImports, currentBlocks, true
}

func (f *generatedFile) filterExistingBlocks(existingBlocks, currentBlocks []generatedBlock, existingImports map[string]string, dir string) (filtered []generatedBlock, hasActiveExternal bool) {
	currentByName := indexBlocksByName(currentBlocks)
	pkgDecls := scanHandwrittenDirDeclarations(dir, f.key.FileName)
	validKRMTypes := f.resolveValidKRMTypes(existingImports)

	for _, eb := range existingBlocks {
		if eb.name == "" {
			continue
		}
		if cb, inCurrent := currentByName[eb.name]; inCurrent && shouldCurrentReplaceExisting(eb, cb) {
			filtered = append(filtered, eb)
			continue
		}
		if !f.shouldKeepExistingBlock(eb, pkgDecls, validKRMTypes) {
			continue
		}
		if !eb.isCommented {
			hasActiveExternal = true
		}
		filtered = append(filtered, eb)
	}
	return filtered, hasActiveExternal
}

func (f *generatedFile) resolveValidKRMTypes(existingImports map[string]string) map[string]bool {
	if f.validKRMTypes != nil {
		return f.validKRMTypes
	}
	if f.key.FileName == "mapper.generated.go" && f.baseDir != "" {
		return discoverValidKRMTypes(f.baseDir, f.imports, existingImports)
	}
	return nil
}

func (f *generatedFile) shouldKeepExistingBlock(eb generatedBlock, pkgDecls packageDeclarations, validKRMTypes map[string]bool) bool {
	switch eb.kind {
	case "type":
		if eb.isCommented {
			return true
		}
		if pkgDecls.types[eb.name] || f.reservedTypeNames[eb.name] {
			return false
		}
		if eb.protoFQN != "" && pkgDecls.protoTags[eb.protoFQN] {
			baseName := strings.TrimSuffix(eb.name, "ObservedState")
			if isBaseProtoStructName(eb.name, eb.protoFQN) || isBaseProtoStructName(baseName, eb.protoFQN) {
				return false
			}
		}
	case "func":
		if !eb.isCommented && pkgDecls.funcs[eb.name] {
			return false
		}
		if validKRMTypes != nil && eb.krmTypeName != "" && !validKRMTypes[eb.krmTypeName] {
			return false
		}
	}
	return true
}

func indexBlocksByName(blocks []generatedBlock) map[string]generatedBlock {
	byName := make(map[string]generatedBlock, len(blocks))
	for _, b := range blocks {
		if b.name != "" {
			byName[b.name] = b
		}
	}
	return byName
}

func shouldCurrentReplaceExisting(eb, cb generatedBlock) bool {
	if !eb.isCommented && cb.isCommented && eb.protoFQN != "" && cb.protoFQN != "" && eb.protoFQN != cb.protoFQN {
		return false
	}
	return true
}

// mergeExistingImports merges imports from the existing file into f.imports.
// When an existing import alias collides with a different package in f.imports
// (or the same package uses a different alias), it returns the old->new
// qualifier map so callers can rewrite selector prefixes in preserved blocks.
func (f *generatedFile) mergeExistingImports(existingImports map[string]string) map[string]string {
	usedAliases := f.currentImportQualifiers()
	aliasRenames := make(map[string]string)

	existingPkgs := make([]string, 0, len(existingImports))
	for pkg := range existingImports {
		existingPkgs = append(existingPkgs, pkg)
	}
	sort.Strings(existingPkgs)

	for _, pkg := range existingPkgs {
		alias := existingImports[pkg]
		oldQual := importQualifier(pkg, alias)
		if currentAlias, already := f.imports[pkg]; already {
			newQual := importQualifier(pkg, currentAlias)
			if oldQual != "" && newQual != "" && oldQual != newQual {
				aliasRenames[oldQual] = newQual
			}
			continue
		}
		chosenAlias := alias
		if oldQual != "" {
			if otherPkg, taken := usedAliases[oldQual]; taken && otherPkg != pkg {
				chosenAlias = nextAvailableAlias(oldQual, usedAliases)
				aliasRenames[oldQual] = chosenAlias
			}
			usedAliases[importQualifier(pkg, chosenAlias)] = pkg
		}
		f.addImport(chosenAlias, pkg)
	}
	return aliasRenames
}

func (f *generatedFile) currentImportQualifiers() map[string]string {
	usedAliases := make(map[string]string, len(f.imports))
	for pkg, alias := range f.imports {
		if qual := importQualifier(pkg, alias); qual != "" {
			usedAliases[qual] = pkg
		}
	}
	return usedAliases
}

func importQualifier(pkg, alias string) string {
	if alias != "" {
		return alias
	}
	return lastGoComponent(pkg)
}

func nextAvailableAlias(baseQual string, usedAliases map[string]string) string {
	for idx := 2; ; idx++ {
		candidate := fmt.Sprintf("%s_%d", baseQual, idx)
		if _, taken := usedAliases[candidate]; !taken {
			return candidate
		}
	}
}

func rewriteBlockQualifiers(blocks []generatedBlock, aliasRenames map[string]string) {
	if len(aliasRenames) == 0 {
		return
	}
	for i := range blocks {
		if !blocks[i].isCommented {
			blocks[i].rawText = rewritePackageQualifiers(blocks[i].rawText, aliasRenames)
		}
	}
}

func rewritePackageQualifiers(src string, renames map[string]string) string {
	if len(renames) == 0 {
		return src
	}
	toks := scanTokens(src)
	var out strings.Builder
	last := 0
	for i, t := range toks {
		if t.tok != token.IDENT {
			continue
		}
		newAlias, ok := renames[t.lit]
		if !ok {
			continue
		}
		prevIsPeriod := i > 0 && toks[i-1].tok == token.PERIOD
		nextIsPeriod := i+1 < len(toks) && toks[i+1].tok == token.PERIOD
		if !prevIsPeriod && nextIsPeriod {
			out.WriteString(src[last:t.pos])
			out.WriteString(newAlias)
			last = t.end
		}
	}
	out.WriteString(src[last:])
	return out.String()
}

func scanTokens(src string) []tokSpan {
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	var s scanner.Scanner
	s.Init(file, []byte(src), nil, scanner.ScanComments)

	var toks []tokSpan
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		off := file.Offset(pos)
		length := len(lit)
		if length == 0 {
			length = len(tok.String())
		}
		toks = append(toks, tokSpan{pos: off, end: off + length, tok: tok, lit: lit})
	}
	return toks
}

func mergeBlocksForFile(fileName string, existing, current []generatedBlock) []generatedBlock {
	if fileName == "types.generated.go" && allBlocksHaveProtoFQN(existing, current) {
		return mergeAndSortTypeBlocks(existing, current)
	}
	return mergeOrderedBlocks(existing, current)
}

func (f *generatedFile) replaceBodyWithBlocks(blocks []generatedBlock) {
	f.body.Reset()
	for _, b := range blocks {
		f.body.WriteString("\n")
		f.body.WriteString(b.rawText)
		f.body.WriteString("\n")
	}
}

func discoverValidKRMTypes(baseDir string, currentImports, existingImports map[string]string) map[string]bool {
	apiDirs := importedAPIDirs(baseDir, currentImports, existingImports)
	if len(apiDirs) == 0 {
		return nil
	}
	valid := make(map[string]bool)
	fset := token.NewFileSet()
	for _, apiDir := range apiDirs {
		collectExportedTypesInDir(apiDir, fset, valid)
	}
	return valid
}

func importedAPIDirs(baseDir string, importMaps ...map[string]string) []string {
	const krmPrefix = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/"
	seenDirs := make(map[string]bool)
	var apiDirs []string
	for _, impMap := range importMaps {
		for pkg := range impMap {
			_, rel, ok := strings.Cut(pkg, krmPrefix)
			if !ok || rel == "" || strings.HasPrefix(rel, "refs/") || strings.HasPrefix(rel, "common/") {
				continue
			}
			apiDir := filepath.Join(baseDir, "apis", filepath.FromSlash(rel))
			if seenDirs[apiDir] {
				continue
			}
			if info, err := os.Stat(apiDir); err == nil && info.IsDir() {
				seenDirs[apiDir] = true
				apiDirs = append(apiDirs, apiDir)
			}
		}
	}
	return apiDirs
}

func collectExportedTypesInDir(apiDir string, fset *token.FileSet, valid map[string]bool) {
	entries, err := os.ReadDir(apiDir)
	if err != nil {
		return
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		src, err := os.ReadFile(filepath.Join(apiDir, e.Name()))
		if err != nil {
			continue
		}
		astFile, err := parser.ParseFile(fset, e.Name(), src, 0)
		if err != nil {
			continue
		}
		for _, decl := range astFile.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}
			for _, spec := range genDecl.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok && ts.Name != nil {
					valid[ts.Name.Name] = true
				}
			}
		}
	}
}

func mergeFileAnnotations(existing, current *annotations.FileAnnotation) *annotations.FileAnnotation {
	if current == nil {
		return existing
	}
	if existing == nil || existing.Key != current.Key || !hasCompatibleKRMGroup(existing, current) {
		return current
	}

	merged := &annotations.FileAnnotation{
		Key:        current.Key,
		Attributes: make(map[string][]string),
	}
	for k := range existing.Attributes {
		merged.Attributes[k] = mergeUniqueStrings(existing.Attributes[k], current.Attributes[k])
	}
	for k := range current.Attributes {
		if _, already := merged.Attributes[k]; !already {
			merged.Attributes[k] = mergeUniqueStrings(nil, current.Attributes[k])
		}
	}
	return merged
}

func mergeUniqueStrings(first, second []string) []string {
	seen := make(map[string]bool, len(first)+len(second))
	var out []string
	for _, list := range [][]string{first, second} {
		for _, v := range list {
			if !seen[v] {
				seen[v] = true
				out = append(out, v)
			}
		}
	}
	return out
}

func isTargetOfSeparateGeneratedFile(dir, currentFileName string, currentAnn *annotations.FileAnnotation) bool {
	if currentAnn == nil {
		return false
	}
	currentResources := currentAnn.Attributes["resource"]
	if len(currentResources) == 0 {
		return false
	}
	resSet := make(map[string]bool, len(currentResources))
	for _, r := range currentResources {
		resSet[r] = true
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		name := e.Name()
		if name == currentFileName || !strings.HasSuffix(name, "_types.generated.go") {
			continue
		}
		if fileClaimsAnyResource(filepath.Join(dir, name), resSet) {
			return true
		}
	}
	return false
}

func fileClaimsAnyResource(filePath string, targetResources map[string]bool) bool {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}
	anns, err := annotations.FindFileAnnotations(b, []string{"+generated:"})
	if err != nil {
		return false
	}
	for _, a := range anns {
		for _, r := range a.Attributes["resource"] {
			if targetResources[r] {
				return true
			}
		}
	}
	return false
}

func hasVersionedMapperFuncs(blocks []generatedBlock) bool {
	for _, b := range blocks {
		if b.kind == "func" && isVersionedMapperFuncName(b.name) {
			return true
		}
	}
	return false
}

func isVersionedMapperFuncName(name string) bool {
	for _, suffix := range []string{"_FromProto", "_ToProto"} {
		base, ok := strings.CutSuffix(name, suffix)
		if !ok {
			continue
		}
		if idx := strings.LastIndex(base, "_"); idx > 0 {
			ver := base[idx+1:]
			if strings.HasPrefix(ver, "v1") || strings.HasPrefix(ver, "v2") {
				return true
			}
		}
	}
	return false
}

func allBlocksHaveProtoFQN(a, b []generatedBlock) bool {
	for _, list := range [][]generatedBlock{a, b} {
		for _, blk := range list {
			if blk.protoFQN == "" {
				return false
			}
		}
	}
	return true
}

func mergeAndSortTypeBlocks(existing, current []generatedBlock) []generatedBlock {
	byName := make(map[string]generatedBlock, len(existing)+len(current))
	for _, b := range existing {
		byName[b.name] = b
	}
	for _, b := range current {
		if eb, exists := byName[b.name]; exists && !shouldCurrentReplaceExisting(eb, b) {
			continue
		}
		byName[b.name] = b
	}
	merged := make([]generatedBlock, 0, len(byName))
	for _, b := range byName {
		merged = append(merged, b)
	}
	sort.Slice(merged, func(i, j int) bool {
		if merged[i].isObservedState != merged[j].isObservedState {
			return !merged[i].isObservedState && merged[j].isObservedState
		}
		if merged[i].protoFQN != merged[j].protoFQN {
			return merged[i].protoFQN < merged[j].protoFQN
		}
		return merged[i].name < merged[j].name
	})
	return merged
}

func mergeOrderedBlocks(existing, current []generatedBlock) []generatedBlock {
	dedupCurrent, currentByName := deduplicateBlocks(current)

	var result []generatedBlock
	placed := make(map[string]bool, len(existing)+len(dedupCurrent))
	for _, eb := range existing {
		if cb, ok := currentByName[eb.name]; ok && shouldCurrentReplaceExisting(eb, cb) {
			result = append(result, cb)
		} else {
			result = append(result, eb)
		}
		placed[eb.name] = true
	}

	for i, cb := range dedupCurrent {
		if placed[cb.name] {
			continue
		}
		result = insertUnplacedBlock(result, dedupCurrent, i, placed)
		placed[cb.name] = true
	}
	return result
}

func deduplicateBlocks(blocks []generatedBlock) ([]generatedBlock, map[string]generatedBlock) {
	byName := make(map[string]generatedBlock, len(blocks))
	seenIdx := make(map[string]int, len(blocks))
	var dedup []generatedBlock
	for _, b := range blocks {
		byName[b.name] = b
		if idx, ok := seenIdx[b.name]; ok {
			dedup[idx] = b
		} else {
			seenIdx[b.name] = len(dedup)
			dedup = append(dedup, b)
		}
	}
	return dedup, byName
}

func insertUnplacedBlock(result, dedupCurrent []generatedBlock, idx int, placed map[string]bool) []generatedBlock {
	cb := dedupCurrent[idx]
	if idx > 0 && placed[dedupCurrent[idx-1].name] {
		if updated, ok := insertBlockAtNeighbor(result, dedupCurrent[idx-1].name, cb, true); ok {
			return updated
		}
	}
	for j := idx + 1; j < len(dedupCurrent); j++ {
		if placed[dedupCurrent[j].name] {
			if updated, ok := insertBlockAtNeighbor(result, dedupCurrent[j].name, cb, false); ok {
				return updated
			}
			break
		}
	}
	return append(result, cb)
}

func insertBlockAtNeighbor(blocks []generatedBlock, neighborName string, cb generatedBlock, after bool) ([]generatedBlock, bool) {
	for idx := range blocks {
		if blocks[idx].name == neighborName {
			insertPos := idx
			if after {
				insertPos = idx + 1
			}
			updated := append(blocks[:insertPos], append([]generatedBlock{cb}, blocks[insertPos:]...)...)
			return updated, true
		}
	}
	return blocks, false
}

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

func extractProtoInfoFromBlockText(rawText string) (protoFQN string, isObservedState bool) {
	for _, line := range strings.Split(rawText, "\n") {
		trimmed := strings.TrimSpace(line)
		if msg, kind, ok := GetProtoMessageAndKindFromAnnotation(trimmed); ok {
			return msg, kind == KCCProtoMessageAnnotationObservedState
		}
	}
	return "", false
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

func extractKRMTypeFromFuncText(rawText string) string {
	for _, line := range strings.Split(rawText, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "func ") {
			return extractKRMTypeFromFuncSignature(trimmed)
		}
	}
	return ""
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

func scanHandwrittenDirDeclarations(srcDir, currentFileName string) packageDeclarations {
	decls := packageDeclarations{
		types:     make(map[string]bool),
		protoTags: make(map[string]bool),
		funcs:     make(map[string]bool),
	}
	files, err := os.ReadDir(srcDir)
	if err != nil {
		return decls
	}
	for _, f := range files {
		name := f.Name()
		if !strings.HasSuffix(name, ".go") || name == currentFileName || name == "zz_generated.deepcopy.go" || strings.HasSuffix(name, "_test.go") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(srcDir, name))
		if err != nil {
			continue
		}
		if strings.HasSuffix(name, "generated.go") {
			recordGeneratedFileDecls(b, &decls)
		} else {
			recordHandwrittenFileDecls(string(b), &decls)
		}
	}
	return decls
}

func recordGeneratedFileDecls(content []byte, decls *packageDeclarations) {
	blocks, _, err := parseGeneratedGoBlocks(content, "p", false)
	if err != nil {
		return
	}
	for _, blk := range blocks {
		if blk.isCommented || blk.name == "" {
			continue
		}
		switch blk.kind {
		case "type":
			decls.types[blk.name] = true
			if blk.protoFQN != "" {
				decls.protoTags[blk.protoFQN] = true
			}
		case "func":
			decls.funcs[blk.name] = true
		}
	}
}

func recordHandwrittenFileDecls(content string, decls *packageDeclarations) {
	protoLines := make(map[string]string)
	scanFileDeclarations(content, decls.types, protoLines)
	for protoMsg := range protoLines {
		decls.protoTags[protoMsg] = true
	}
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if rest, ok := strings.CutPrefix(trimmed, "func "); ok {
			if fnName, _, found := strings.Cut(rest, "("); found && fnName != "" && !strings.Contains(fnName, " ") {
				decls.funcs[fnName] = true
			}
		}
	}
}
