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

package sanitizer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	DeletionPolicyAnnotation   = "cnrm.cloud.google.com/deletion-policy"
	ConflictPolicyAnnotation   = "cnrm.cloud.google.com/management-conflict-prevention-policy"
	StateIntoSpecAnnotation    = "cnrm.cloud.google.com/state-into-spec"
	PolicyAbandonValue         = "abandon"
	PolicyConflictNoneValue    = "none"
	StateIntoSpecAbsentValue   = "absent"
)

// SanitizeOptions configures the sanitization behavior.
type SanitizeOptions struct {
	InjectAbandon       bool
	InjectConflictNone  bool
	InjectStateIntoSpec bool
	TargetNamespace     string
	KeepNamespace       bool
}

// SanitizeFile reads a YAML file containing one or more Kubernetes resources,
// strips runtime metadata and status, injects safety annotations, and writes to target path.
func SanitizeFile(srcPath, dstPath string, opts SanitizeOptions) (int, error) {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return 0, fmt.Errorf("reading source file %s: %w", srcPath, err)
	}

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var sanitizedDocs []map[string]interface{}
	resourceCount := 0

	for {
		var doc map[string]interface{}
		err := decoder.Decode(&doc)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("parsing YAML from %s: %w", srcPath, err)
		}
		if len(doc) == 0 {
			continue
		}

		var rawResources []map[string]interface{}
		if kind, ok := doc["kind"].(string); ok && kind == "List" {
			if items, ok := doc["items"].([]interface{}); ok {
				for _, item := range items {
					if itemMap, ok := item.(map[string]interface{}); ok {
						rawResources = append(rawResources, itemMap)
					}
				}
			}
		} else {
			rawResources = append(rawResources, doc)
		}

		for _, res := range rawResources {
			sanitized, isKCC := SanitizeResource(res, opts)
			if isKCC {
				sanitizedDocs = append(sanitizedDocs, sanitized)
				resourceCount++
			}
		}
	}

	if resourceCount == 0 {
		return 0, nil
	}

	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		return 0, fmt.Errorf("creating destination directory: %w", err)
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return 0, fmt.Errorf("creating destination file %s: %w", dstPath, err)
	}
	defer dstFile.Close()

	encoder := yaml.NewEncoder(dstFile)
	encoder.SetIndent(2)
	defer encoder.Close()

	for _, doc := range sanitizedDocs {
		if err := encoder.Encode(doc); err != nil {
			return 0, fmt.Errorf("encoding YAML to %s: %w", dstPath, err)
		}
	}

	return resourceCount, nil
}

// SanitizeResource cleans a single resource map.
func SanitizeResource(doc map[string]interface{}, opts SanitizeOptions) (map[string]interface{}, bool) {
	apiVersion, _ := doc["apiVersion"].(string)
	kind, _ := doc["kind"].(string)

	if kind == "ConfigConnector" || kind == "ConfigConnectorContext" || kind == "MultiClusterLease" {
		return nil, false
	}

	if !strings.Contains(apiVersion, ".cnrm.cloud.google.com") && !strings.HasPrefix(apiVersion, "core.cnrm.cloud.google.com") {
		return nil, false
	}

	// 1. Remove status block completely
	delete(doc, "status")

	// 2. Clean metadata
	rawMetadata, hasMeta := doc["metadata"]
	if !hasMeta || rawMetadata == nil {
		return doc, true
	}

	metadata, ok := rawMetadata.(map[string]interface{})
	if !ok {
		return doc, true
	}

	// Server-managed fields and existing finalizers to remove
	serverFields := []string{
		"uid",
		"resourceVersion",
		"generation",
		"creationTimestamp",
		"managedFields",
		"selfLink",
		"ownerReferences",
		"finalizers",
	}
	for _, f := range serverFields {
		delete(metadata, f)
	}

	// 3. Clean annotations
	rawAnnotations, hasAnnotations := metadata["annotations"]
	var annotations map[string]interface{}
	if hasAnnotations && rawAnnotations != nil {
		annotations, _ = rawAnnotations.(map[string]interface{})
	}
	if annotations == nil {
		annotations = make(map[string]interface{})
	}

	// Strip Config Sync / ACM / kubectl internal annotations
	cleanedAnnotations := make(map[string]interface{})
	for k, v := range annotations {
		if strings.HasPrefix(k, "configmanagement.gke.io/") ||
			strings.HasPrefix(k, "configsync.gke.io/") ||
			strings.HasPrefix(k, "kubectl.kubernetes.io/") ||
			k == "cnrm.cloud.google.com/observed-secret-versions" {
			continue
		}
		cleanedAnnotations[k] = v
	}

	// Inject safety annotations
	if opts.InjectAbandon {
		cleanedAnnotations[DeletionPolicyAnnotation] = PolicyAbandonValue
	}
	if opts.InjectConflictNone {
		cleanedAnnotations[ConflictPolicyAnnotation] = PolicyConflictNoneValue
	}
	if opts.InjectStateIntoSpec {
		if _, exists := cleanedAnnotations[StateIntoSpecAnnotation]; !exists {
			cleanedAnnotations[StateIntoSpecAnnotation] = StateIntoSpecAbsentValue
		}
	}

	if len(cleanedAnnotations) > 0 {
		metadata["annotations"] = cleanedAnnotations
	} else {
		delete(metadata, "annotations")
	}

	// 4. Handle namespace override
	if opts.TargetNamespace != "" {
		metadata["namespace"] = opts.TargetNamespace
	} else if !opts.KeepNamespace {
		// Keep existing namespace if present
	}

	doc["metadata"] = metadata
	return doc, true
}
