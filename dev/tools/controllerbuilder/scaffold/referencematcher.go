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

package scaffold

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
)

// sharedRefsPackage holds the reference types services use in common, relative
// to the repo root.
const sharedRefsPackage = "apis/refs/v1beta1"

// candidatesForTarget finds reference types matching a known proto resource target type.
func (a *APIScaffolder) candidatesForTarget(targetType, serviceDir string) map[string]string {
	targetService, targetKind := protoapi.ParseResourceTarget(targetType)
	if targetKind == "" {
		return nil
	}
	currentService := a.serviceName()
	want := strings.ToLower(targetKind) + "ref"
	re := regexp.MustCompile(`(?m)^type ([A-Z]\w*Ref) struct`)

	// 1. Scan serviceDir if service matches
	if serviceMatchesPrefix(currentService, targetService) || serviceMatchesPrefix(targetService, currentService) {
		found := map[string]string{}
		entries, err := os.ReadDir(serviceDir)
		if err == nil {
			for _, e := range entries {
				if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
					continue
				}
				body, err := os.ReadFile(filepath.Join(serviceDir, e.Name()))
				if err != nil {
					continue
				}
				for _, m := range re.FindAllStringSubmatch(string(body), -1) {
					if strings.HasSuffix(strings.ToLower(m[1]), want) {
						found[m[1]] = ""
					}
				}
			}
		}
		if len(found) > 0 {
			return found
		}
	}

	// 2. Scan shared refs package
	sharedDir := filepath.Join(a.repoRoot(), sharedRefsPackage)
	found := map[string]string{}
	entries, err := os.ReadDir(sharedDir)
	if err == nil {
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
				continue
			}
			body, err := os.ReadFile(filepath.Join(sharedDir, e.Name()))
			if err != nil {
				continue
			}
			for _, m := range re.FindAllStringSubmatch(string(body), -1) {
				typeName := m[1]
				lowerType := strings.ToLower(typeName)
				if !strings.HasSuffix(lowerType, want) {
					continue
				}
				// Exact match (e.g. ProjectRef, FolderRef, OrganizationRef)
				if strings.EqualFold(typeName, want) {
					found[typeName] = "refsv1beta1"
					continue
				}
				// Service prefix match
				prefix := strings.TrimSuffix(lowerType, want)
				if serviceMatchesPrefix(targetService, prefix) || serviceMatchesPrefix(currentService, prefix) {
					found[typeName] = "refsv1beta1"
				}
			}
		}
	}
	return found
}

// parentRefTypes returns the reference types that could name a parent whose
// collection singularises to name, as a map of type name to import qualifier.
//
// Services spell a reference type either as a bare noun, ClusterRef in alloydb,
// or with the Kind's own prefix, as in DNSManagedZoneRef, so the match is on
// the suffix. The service's own package wins over the shared one, and
// unexported types are skipped: apis/kms/v1beta1 declares kmsCryptoKeyRef
// beside the real KMSCryptoKeyRef, which would make a single match look
// ambiguous.
//
// Shared reference types in apis/refs/v1beta1 are only matched if they are
// un-prefixed generic references or their prefix matches the current service,
// preventing foreign services (e.g. Chronicle) from accidentally matching
// service-specific types (e.g. SQLInstanceRef).
func parentRefTypes(repoRoot, serviceDir, currentService, name string) map[string]string {
	want := strings.ToLower(name) + "ref"
	re := regexp.MustCompile(`(?m)^type ([A-Z]\w*Ref) struct`)

	scan := func(dir, qualifier string, isShared bool) map[string]string {
		found := map[string]string{}
		entries, err := os.ReadDir(dir)
		if err != nil {
			return found
		}
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
				continue
			}
			body, err := os.ReadFile(filepath.Join(dir, e.Name()))
			if err != nil {
				continue
			}
			for _, m := range re.FindAllStringSubmatch(string(body), -1) {
				typeName := m[1]
				lowerType := strings.ToLower(typeName)
				if !strings.HasSuffix(lowerType, want) {
					continue
				}
				if isShared {
					if !isAllowedSharedRef(typeName, currentService, want) {
						continue
					}
				}
				found[typeName] = qualifier
			}
		}
		return found
	}

	if found := scan(serviceDir, "", false); len(found) > 0 {
		return found
	}
	return scan(filepath.Join(repoRoot, sharedRefsPackage), "refsv1beta1", true)
}

// isAllowedSharedRef ensures that types in apis/refs/v1beta1 only match if
// they are generic un-prefixed references or their prefix matches the current service.
func isAllowedSharedRef(typeName, currentService, want string) bool {
	if strings.EqualFold(typeName, want) {
		return true
	}
	if currentService == "" {
		return false
	}
	prefix := strings.TrimSuffix(strings.ToLower(typeName), want)
	return serviceMatchesPrefix(currentService, prefix)
}

// serviceMatchesPrefix checks if a service name matches a prefix on a reference type.
func serviceMatchesPrefix(service, prefix string) bool {
	service = strings.ToLower(service)
	prefix = strings.ToLower(prefix)
	if service == "" || prefix == "" {
		return false
	}
	if service == prefix {
		return true
	}
	switch service {
	case "sql", "sqladmin":
		return prefix == "sql" || prefix == "sqladmin"
	case "iam":
		return prefix == "iam" || prefix == "gcpserviceaccount"
	case "secretmanager":
		return prefix == "secretmanager" || prefix == "secret"
	case "bigquery":
		return prefix == "bigquery"
	case "compute":
		return prefix == "compute"
	case "alloydb":
		return prefix == "alloydb"
	case "appengine":
		return prefix == "appengine"
	case "kms":
		return prefix == "kms"
	}
	return false
}

// serviceName extracts the service name from APIScaffolder's GoPackage or Group.
func (a *APIScaffolder) serviceName() string {
	if a.GoPackage != "" {
		clean := filepath.Clean(a.GoPackage)
		parts := strings.Split(clean, string(filepath.Separator))
		if len(parts) > 0 && parts[0] != "" && parts[0] != "." {
			return parts[0]
		}
	}
	if a.Group != "" {
		parts := strings.Split(a.Group, ".")
		if len(parts) > 0 {
			return parts[0]
		}
	}
	return ""
}

// repoRoot returns the directory above BaseDir when BaseDir is an apis
// directory, and BaseDir otherwise, so the shared refs package can be found.
//
// generate-types defaults BaseDir to "<repo>/apis/", with a trailing slash,
// and filepath.Dir of that is "<repo>/apis", so the path is cleaned first.
func (a *APIScaffolder) repoRoot() string {
	dir := filepath.Clean(a.BaseDir)
	if filepath.Base(dir) == "apis" {
		return filepath.Dir(dir)
	}
	return dir
}
