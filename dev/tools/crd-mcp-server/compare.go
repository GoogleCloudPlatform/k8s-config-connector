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

package main

import (
	"fmt"
	"maps"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/yaml"
)

// parseCRD parses a CRD from YAML bytes.
// Note: yaml.Unmarshal only reads the first document; files with multiple CRDs
// will have only the first one checked.
func parseCRD(data []byte) (*apiextensionsv1.CustomResourceDefinition, error) {
	var crd apiextensionsv1.CustomResourceDefinition
	if err := yaml.Unmarshal(data, &crd); err != nil {
		return nil, fmt.Errorf("unmarshaling CRD: %w", err)
	}
	return &crd, nil
}

// gitShow retrieves the content of a file at the given git ref.
// Returns (content, isNew, error) where isNew is true if the file did not exist at the ref.
func gitShow(ref, file string) ([]byte, bool, error) {
	// Prevent argument injection: git flag-like refs would be misinterpreted.
	if strings.HasPrefix(ref, "-") {
		return nil, false, fmt.Errorf("invalid ref %q: must not start with '-'", ref)
	}

	rootCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	rootOut, err := rootCmd.Output()
	if err != nil {
		return nil, false, fmt.Errorf("getting git root: %w", err)
	}
	root := strings.TrimSpace(string(rootOut))

	absFile, err := filepath.Abs(file)
	if err != nil {
		return nil, false, fmt.Errorf("getting absolute path of %q: %w", file, err)
	}

	relPath, err := filepath.Rel(root, absFile)
	if err != nil {
		return nil, false, fmt.Errorf("computing relative path of %q from repo root %q: %w", absFile, root, err)
	}

	// Reject paths outside the repository root (filepath.Rel can produce "../..." paths).
	if strings.HasPrefix(relPath, "..") {
		return nil, false, fmt.Errorf("file %q is outside the repository root %q", absFile, root)
	}

	showCmd := exec.Command("git", "show", fmt.Sprintf("%s:%s", ref, relPath))
	output, err := showCmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr := string(exitErr.Stderr)
			// git show reports these messages when the path doesn't exist at the ref.
			if strings.Contains(stderr, "exists on disk, but not in") ||
				strings.Contains(stderr, "does not exist in") {
				return nil, true, nil
			}
		}
		return nil, false, fmt.Errorf("git show %s:%s: %w", ref, relPath, err)
	}
	return output, false, nil
}

const maxWalkDepth = 50

// walk converts a JSONSchemaProps into a simplified nested structure for comparison.
// Descriptions are intentionally discarded.
func walk(s *apiextensionsv1.JSONSchemaProps) any {
	return walkDepth(s, 0)
}

// walkDepth is the recursive implementation of walk, tracking depth to guard
// against stack overflow on cyclic or deeply nested schemas.
func walkDepth(s *apiextensionsv1.JSONSchemaProps, depth int) any {
	if s == nil {
		return "unknown"
	}
	if depth > maxWalkDepth {
		return "unknown (depth limit exceeded)"
	}

	var res any
	if s.XPreserveUnknownFields != nil && *s.XPreserveUnknownFields {
		res = "json"
	} else if len(s.Properties) > 0 {
		m := make(map[string]any)
		for k, v := range s.Properties {
			val := v
			m[k] = walkDepth(&val, depth+1)
		}
		res = m
	} else if s.Type == "array" {
		if s.Items != nil && s.Items.Schema != nil {
			res = []any{walkDepth(s.Items.Schema, depth+1)}
		} else {
			// TODO: s.Items can also be a slice of schemas (tuple validation); not handled here.
			res = []any{"any"}
		}
	} else if s.AdditionalProperties != nil {
		if s.AdditionalProperties.Schema != nil {
			val := walkDepth(s.AdditionalProperties.Schema, depth+1)
			if str, ok := val.(string); ok {
				res = fmt.Sprintf("map[string]%s", str)
			} else {
				res = map[string]any{"KEY": val}
			}
		} else if !s.AdditionalProperties.Allows {
			// Boolean form: Allows=false means no additional properties are permitted.
			res = "map[string]none"
		} else {
			res = "map[string]any"
		}
	} else {
		t := s.Type
		if t == "" {
			// Walk into anyOf/allOf/oneOf so that changes within those schemas are detected.
			if len(s.AnyOf) > 0 {
				items := make([]any, len(s.AnyOf))
				for i := range s.AnyOf {
					items[i] = walkDepth(&s.AnyOf[i], depth+1)
				}
				res = map[string]any{"anyOf": items}
			} else if len(s.AllOf) > 0 {
				items := make([]any, len(s.AllOf))
				for i := range s.AllOf {
					items[i] = walkDepth(&s.AllOf[i], depth+1)
				}
				res = map[string]any{"allOf": items}
			} else if len(s.OneOf) > 0 {
				items := make([]any, len(s.OneOf))
				for i := range s.OneOf {
					items[i] = walkDepth(&s.OneOf[i], depth+1)
				}
				res = map[string]any{"oneOf": items}
			} else {
				res = "any"
			}
		} else if s.Format != "" {
			switch s.Format {
			case "int32", "int64":
				res = s.Format
			default:
				res = fmt.Sprintf("%s (%s)", t, s.Format)
			}
		} else {
			res = t
		}
	}

	if len(s.XValidations) > 0 {
		m := make(map[string]any)
		if mExisting, ok := res.(map[string]any); ok {
			for k, v := range mExisting {
				m[k] = v
			}
		} else {
			m[":type"] = res
		}
		for i, v := range s.XValidations {
			val := map[string]any{
				"rule": v.Rule,
			}
			if v.Message != "" {
				val["message"] = v.Message
			}
			if v.MessageExpression != "" {
				val["messageExpression"] = v.MessageExpression
			}
			if v.Reason != nil {
				val["reason"] = string(*v.Reason)
			}
			if v.FieldPath != "" {
				val["fieldPath"] = v.FieldPath
			}
			if v.OptionalOldSelf != nil {
				val["optionalOldSelf"] = *v.OptionalOldSelf
			}
			m[fmt.Sprintf(":validation[%d]", i)] = val
		}
		return m
	}

	return res
}

// flatten converts a nested schema structure to a flat path -> type map.
// Note: More complex OpenAPI features (like tuples, anyOf, or allOf) are currently flattened
// into simplified structural representations. Their specific validation logic might require
// future enhancements if strict structural equivalence checking of those complex types is needed.
func flatten(path string, schema any, out map[string]string) {
	switch v := schema.(type) {
	case map[string]any:
		if path != "" {
			out[path] = "object"
		}
		for k, child := range v {
			if k == ":type" {
				flatten(path, child, out)
				continue
			}
			childPath := k
			if path != "" {
				childPath = path + "." + k
			}
			flatten(childPath, child, out)
		}
	case []any:
		if path != "" {
			out[path] = "array"
		}
		for _, item := range v {
			flatten(path+"[]", item, out)
		}
	case string:
		out[path] = v
	}
}

// getVersionSchemas returns a map of version name -> flattened schema (path -> type).
func getVersionSchemas(crd *apiextensionsv1.CustomResourceDefinition) map[string]map[string]string {
	result := make(map[string]map[string]string)
	for _, v := range crd.Spec.Versions {
		if v.Schema == nil || v.Schema.OpenAPIV3Schema == nil {
			continue
		}
		schema := walk(v.Schema.OpenAPIV3Schema)
		paths := make(map[string]string)
		flatten("", schema, paths)
		result[v.Name] = paths
	}
	return result
}

// isUnderStatus returns true if the path is within the status section of the schema.
func isUnderStatus(path string) bool {
	return path == "status" ||
		strings.HasPrefix(path, "status.") ||
		strings.HasPrefix(path, "status[")
}

// isValidationPath returns true if the path is within a validation rule.
func isValidationPath(path string) bool {
	return strings.Contains(path, ":validation[")
}

// CompareResult holds the outcome of a CRD comparison.
type CompareResult struct {
	// Diffs are disqualifying differences (non-empty means not equivalent / not backward compatible).
	Diffs []string
	// Notes are informational messages about allowed changes.
	Notes []string
}

// compareEquivalence checks whether the change from oldCRD to newCRD is equivalent.
//
// Equivalent means:
//   - No fields are added or deleted (only status.externalRef, status.observedState,
//     and status.observedState's sub-fields, may be added under 'status')
//   - Field names and types do not change (descriptions may change freely)
//   - Adding spec.names.listKind is fine
func compareEquivalence(oldCRD, newCRD *apiextensionsv1.CustomResourceDefinition) CompareResult {
	var r CompareResult

	// spec.names.listKind: adding is fine, changing is not.
	oldListKind := oldCRD.Spec.Names.ListKind
	newListKind := newCRD.Spec.Names.ListKind
	switch {
	case oldListKind == "" && newListKind != "":
		r.Notes = append(r.Notes, fmt.Sprintf("spec.names.listKind added: %q (allowed)", newListKind))
	case oldListKind != newListKind:
		r.Diffs = append(r.Diffs, fmt.Sprintf("spec.names.listKind changed: %q -> %q", oldListKind, newListKind))
	}

	oldVersions := getVersionSchemas(oldCRD)
	newVersions := getVersionSchemas(newCRD)

	for _, vName := range slices.Sorted(maps.Keys(oldVersions)) {
		if _, ok := newVersions[vName]; !ok {
			r.Diffs = append(r.Diffs, fmt.Sprintf("version %q removed", vName))
		}
	}
	for _, vName := range slices.Sorted(maps.Keys(newVersions)) {
		if _, ok := oldVersions[vName]; !ok {
			r.Diffs = append(r.Diffs, fmt.Sprintf("version %q added", vName))
		}
	}

	for _, vName := range slices.Sorted(maps.Keys(oldVersions)) {
		newPaths, ok := newVersions[vName]
		if !ok {
			continue
		}
		d, n := schemaEquivalenceDiff(vName, oldVersions[vName], newPaths)
		r.Diffs = append(r.Diffs, d...)
		r.Notes = append(r.Notes, n...)
	}

	return r
}

// schemaEquivalenceDiff compares two flattened schemas and returns equivalence diffs and notes.
func schemaEquivalenceDiff(version string, oldPaths, newPaths map[string]string) (diffs, notes []string) {
	prefix := ""
	if version != "" {
		prefix = fmt.Sprintf("[%s] ", version)
	}

	lastRemovedPrefix := ""
	for _, path := range slices.Sorted(maps.Keys(oldPaths)) {
		oldType := oldPaths[path]
		newType, ok := newPaths[path]
		if !ok {
			if lastRemovedPrefix != "" && (strings.HasPrefix(path, lastRemovedPrefix+".") || strings.HasPrefix(path, lastRemovedPrefix+"[")) {
				continue
			}
			diffs = append(diffs, fmt.Sprintf("%sfield removed: %s (was %s)", prefix, path, oldType))
			lastRemovedPrefix = path
			continue
		}
		lastRemovedPrefix = ""
		if oldType != newType {
			if isAllowedTypeChange(path, oldType, newType) {
				notes = append(notes, fmt.Sprintf("%sfield type changed: %s (%s -> %s) (allowed)", prefix, path, oldType, newType))
			} else {
				diffs = append(diffs, fmt.Sprintf("%sfield type changed: %s (%s -> %s)", prefix, path, oldType, newType))
			}
		}
	}

	lastDiffPrefix := ""
	for _, path := range slices.Sorted(maps.Keys(newPaths)) {
		if _, ok := oldPaths[path]; ok {
			continue
		}
		if lastDiffPrefix != "" && (strings.HasPrefix(path, lastDiffPrefix+".") || strings.HasPrefix(path, lastDiffPrefix+"[")) {
			continue
		}
		if isUnderStatus(path) {
			if isAllowedNewStatusField(path) {
				if path == "status" {
					notes = append(notes, fmt.Sprintf("%sstatus block added (allowed)", prefix))
				} else {
					notes = append(notes, fmt.Sprintf("%sfield added under status: %s (type: %s, allowed)", prefix, path, newPaths[path]))
				}
			} else {
				diffs = append(diffs, fmt.Sprintf("%sfield added under status: %s (type: %s)", prefix, path, newPaths[path]))
				lastDiffPrefix = path
			}
		} else {
			diffs = append(diffs, fmt.Sprintf("%sfield added: %s (type: %s)", prefix, path, newPaths[path]))
			lastDiffPrefix = path
		}
	}

	return diffs, notes
}

func isAllowedNewStatusField(path string) bool {
	switch {
	case path == "status":
		return true
	case path == "status.externalRef":
		return true
	case path == "status.observedState" || strings.HasPrefix(path, "status.observedState."):
		return true
	default:
		return false
	}
}

func isAllowedTypeChange(path, oldType, newType string) bool {
	// If it is changed from integer to int32, then it should always be allowed.
	if oldType == "integer" && newType == "int32" {
		return true
	}
	// If it is changed from integer to int64, it should be an allowed change if this happens to a status field.
	if oldType == "integer" && newType == "int64" && isUnderStatus(path) {
		return true
	}
	return false
}

// compareBackwardCompatibility checks whether the change from oldCRD to newCRD is backward compatible.
//
// Backward compatible means:
//   - No fields are removed or renamed
//   - Field types do not change (descriptions may change freely)
//   - New fields may be added anywhere
func compareBackwardCompatibility(oldCRD, newCRD *apiextensionsv1.CustomResourceDefinition) CompareResult {
	var r CompareResult

	oldVersions := getVersionSchemas(oldCRD)
	newVersions := getVersionSchemas(newCRD)

	for _, vName := range slices.Sorted(maps.Keys(oldVersions)) {
		if _, ok := newVersions[vName]; !ok {
			r.Diffs = append(r.Diffs, fmt.Sprintf("version %q removed", vName))
		}
	}
	for _, vName := range slices.Sorted(maps.Keys(newVersions)) {
		if _, ok := oldVersions[vName]; !ok {
			r.Notes = append(r.Notes, fmt.Sprintf("version %q added (allowed)", vName))
		}
	}

	for _, vName := range slices.Sorted(maps.Keys(oldVersions)) {
		newPaths, ok := newVersions[vName]
		if !ok {
			continue
		}
		d, n := schemaBackwardCompatDiff(vName, oldVersions[vName], newPaths)
		r.Diffs = append(r.Diffs, d...)
		r.Notes = append(r.Notes, n...)
	}

	return r
}

// schemaBackwardCompatDiff compares two flattened schemas for backward compatibility.
func schemaBackwardCompatDiff(version string, oldPaths, newPaths map[string]string) (diffs, notes []string) {
	prefix := ""
	if version != "" {
		prefix = fmt.Sprintf("[%s] ", version)
	}

	lastRemovedPrefix := ""
	for _, path := range slices.Sorted(maps.Keys(oldPaths)) {
		// x-kubernetes-validations are not checked for backward compatibility for now.
		// These are determined on a case-by-case basis.
		if isValidationPath(path) {
			continue
		}
		oldType := oldPaths[path]
		newType, ok := newPaths[path]
		if !ok {
			if lastRemovedPrefix != "" && (strings.HasPrefix(path, lastRemovedPrefix+".") || strings.HasPrefix(path, lastRemovedPrefix+"[")) {
				continue
			}
			diffs = append(diffs, fmt.Sprintf("%sfield removed: %s (was %s)", prefix, path, oldType))
			lastRemovedPrefix = path
			continue
		}
		lastRemovedPrefix = ""
		if oldType != newType {
			if isAllowedTypeChange(path, oldType, newType) {
				notes = append(notes, fmt.Sprintf("%sfield type changed: %s (%s -> %s) (allowed)", prefix, path, oldType, newType))
			} else {
				diffs = append(diffs, fmt.Sprintf("%sfield type changed: %s (%s -> %s)", prefix, path, oldType, newType))
			}
		}
	}

	for _, path := range slices.Sorted(maps.Keys(newPaths)) {
		// x-kubernetes-validations are not checked for backward compatibility for now.
		// These are determined on a case-by-case basis.
		if isValidationPath(path) {
			continue
		}
		if _, ok := oldPaths[path]; !ok {
			notes = append(notes, fmt.Sprintf("%sfield added: %s (type: %s, allowed)", prefix, path, newPaths[path]))
		}
	}

	return diffs, notes
}
