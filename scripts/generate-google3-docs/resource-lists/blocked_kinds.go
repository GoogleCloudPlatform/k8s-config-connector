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
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// releaseBundleScript is the script that defines BLOCKED_KINDS, the kinds
// whose CRDs are excluded from the release bundle until they are verified.
const releaseBundleScript = "dev/tasks/build-release-bundle"

var (
	blockedKindsArrayRegex = regexp.MustCompile(`(?s)\bBLOCKED_KINDS=\((.*?)\)`)
	blockedKindRegex       = regexp.MustCompile(`"([^"]+)"`)
)

// blockedKinds returns the (lowercased) kinds listed in BLOCKED_KINDS in the
// release bundle script. These kinds are not released, so they should not be
// listed in the docs.
func blockedKinds() (map[string]bool, error) {
	scriptPath := filepath.Join(repo.GetRootOrLogFatal(), releaseBundleScript)
	b, err := os.ReadFile(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("error reading %v: %w", scriptPath, err)
	}

	// Drop comments so that commented-out kinds are not treated as blocked.
	var lines []string
	for _, line := range strings.Split(string(b), "\n") {
		if i := strings.Index(line, "#"); i >= 0 {
			line = line[:i]
		}
		lines = append(lines, line)
	}

	match := blockedKindsArrayRegex.FindStringSubmatch(strings.Join(lines, "\n"))
	if match == nil {
		return nil, fmt.Errorf("could not find BLOCKED_KINDS in %v", scriptPath)
	}
	kinds := make(map[string]bool)
	for _, m := range blockedKindRegex.FindAllStringSubmatch(match[1], -1) {
		// The release bundle script matches kinds case-insensitively.
		kinds[strings.ToLower(m[1])] = true
	}
	return kinds, nil
}

// excludeBlockedKinds returns the GVKs in resources whose kind is not listed
// in BLOCKED_KINDS.
func excludeBlockedKinds(resources []schema.GroupVersionKind, blocked map[string]bool) []schema.GroupVersionKind {
	filtered := make([]schema.GroupVersionKind, 0, len(resources))
	for _, gvk := range resources {
		if blocked[strings.ToLower(gvk.Kind)] {
			continue
		}
		filtered = append(filtered, gvk)
	}
	return filtered
}
