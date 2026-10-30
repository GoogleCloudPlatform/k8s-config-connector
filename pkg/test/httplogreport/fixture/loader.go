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

package fixture

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

// KRMResource holds KRM metadata extracted from create.yaml, update.yaml, or dependencies.yaml.
type KRMResource struct {
	Kind        string
	APIVersion  string
	Group       string
	Name        string
	ResourceID  string
	ExternalRef string
	Obj         *unstructured.Unstructured
}

// Data contains all parsed KRM and golden artifacts for a single E2E resource fixture directory.
type Data struct {
	FixtureDir     string
	FixtureName    string
	PrimaryCreate  *KRMResource
	PrimaryUpdate  *KRMResource
	Dependencies   []*KRMResource
	GoldenObj      *unstructured.Unstructured
	GoldenFileName string
	ExportedObj    *unstructured.Unstructured
	RawLogContent  string
}

// Load reads _http.log and all KRM YAML manifests from fixtureDir.
func Load(fixtureDir string) (*Data, error) {
	logPath := filepath.Join(fixtureDir, "_http.log")
	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", logPath, err)
	}
	return LoadWithLogContent(fixtureDir, string(logBytes))
}

// LoadWithLogContent loads create.yaml, optional update.yaml, optional dependencies.yaml,
// and optional golden/exported YAML files from fixtureDir using the provided rawLogContent.
func LoadWithLogContent(fixtureDir, rawLogContent string) (*Data, error) {
	fd := &Data{
		FixtureDir:    fixtureDir,
		FixtureName:   filepath.Base(fixtureDir),
		RawLogContent: rawLogContent,
	}

	createPath := filepath.Join(fixtureDir, "create.yaml")
	createDocs, err := parseMultiDocYAMLFile(createPath)
	if err != nil {
		return nil, fmt.Errorf("loading create.yaml: %w", err)
	}
	if len(createDocs) == 0 {
		return nil, fmt.Errorf("no YAML documents found in %s", createPath)
	}

	primaryIdx := selectPrimaryDocIndex(createDocs, fixtureDir)
	fd.PrimaryCreate = extractKRMResource(createDocs[primaryIdx])
	for i, doc := range createDocs {
		if i != primaryIdx {
			fd.Dependencies = append(fd.Dependencies, extractKRMResource(doc))
		}
	}

	updatePath := filepath.Join(fixtureDir, "update.yaml")
	if FileExists(updatePath) {
		updateDocs, err := parseMultiDocYAMLFile(updatePath)
		if err == nil && len(updateDocs) > 0 {
			uIdx := selectMatchingDocIndex(updateDocs, fd.PrimaryCreate)
			fd.PrimaryUpdate = extractKRMResource(updateDocs[uIdx])
		}
	}

	depPath := filepath.Join(fixtureDir, "dependencies.yaml")
	if FileExists(depPath) {
		depDocs, err := parseMultiDocYAMLFile(depPath)
		if err == nil {
			for _, doc := range depDocs {
				meta := extractKRMResource(doc)
				if meta.Kind != "" {
					fd.Dependencies = append(fd.Dependencies, meta)
				}
			}
		}
	}

	goldenMatches, _ := filepath.Glob(filepath.Join(fixtureDir, "_generated_object_*.golden.yaml"))
	if len(goldenMatches) > 0 {
		fd.GoldenFileName = filepath.Base(goldenMatches[0])
		if docs, err := parseMultiDocYAMLFile(goldenMatches[0]); err == nil && len(docs) > 0 {
			fd.GoldenObj = docs[len(docs)-1]
			if fd.PrimaryCreate != nil {
				if extRef, _, _ := unstructured.NestedString(fd.GoldenObj.Object, "status", "externalRef"); extRef != "" {
					fd.PrimaryCreate.ExternalRef = extRef
				} else if selfLink, _, _ := unstructured.NestedString(fd.GoldenObj.Object, "status", "selfLink"); selfLink != "" {
					fd.PrimaryCreate.ExternalRef = selfLink
				}
			}
		}
	}

	exportedPath := filepath.Join(fixtureDir, "_exported.yaml")
	if FileExists(exportedPath) {
		if docs, err := parseMultiDocYAMLFile(exportedPath); err == nil && len(docs) > 0 {
			fd.ExportedObj = docs[0]
		}
	}

	return fd, nil
}

func parseMultiDocYAMLFile(path string) ([]*unstructured.Unstructured, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	normalized := strings.ReplaceAll(string(b), "\r\n", "\n")
	rawDocs := strings.Split(normalized, "\n---")
	var out []*unstructured.Unstructured
	for _, raw := range rawDocs {
		trimmed := strings.TrimSpace(raw)
		if trimmed == "" {
			continue
		}
		hasNonComment := false
		for _, line := range strings.Split(trimmed, "\n") {
			lt := strings.TrimSpace(line)
			if lt != "" && !strings.HasPrefix(lt, "#") {
				hasNonComment = true
				break
			}
		}
		if !hasNonComment {
			continue
		}
		u := &unstructured.Unstructured{}
		if err := yaml.Unmarshal([]byte(trimmed), u); err != nil {
			continue
		}
		if u.GetKind() != "" {
			out = append(out, u)
		}
	}
	return out, nil
}

func selectPrimaryDocIndex(docs []*unstructured.Unstructured, fixtureDir string) int {
	if len(docs) == 1 {
		return 0
	}
	lowerDir := strings.ToLower(fixtureDir)
	for i, doc := range docs {
		if strings.Contains(lowerDir, "/"+strings.ToLower(doc.GetKind())+"/") {
			return i
		}
	}
	return len(docs) - 1
}

func selectMatchingDocIndex(docs []*unstructured.Unstructured, primary *KRMResource) int {
	if primary == nil || len(docs) <= 1 {
		return 0
	}
	for i, doc := range docs {
		if doc.GetKind() == primary.Kind && doc.GetName() == primary.Name {
			return i
		}
	}
	for i, doc := range docs {
		if doc.GetKind() == primary.Kind {
			return i
		}
	}
	return len(docs) - 1
}

func extractKRMResource(u *unstructured.Unstructured) *KRMResource {
	if u == nil {
		return nil
	}
	gvk := u.GroupVersionKind()
	resID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
	return &KRMResource{
		Kind:       gvk.Kind,
		APIVersion: u.GetAPIVersion(),
		Group:      gvk.Group,
		Name:       u.GetName(),
		ResourceID: resID,
		Obj:        u,
	}
}

// DiscoverDirs finds all fixture directories under rootPath containing both create.yaml and _http.log.
func DiscoverDirs(rootPath string) ([]string, error) {
	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		dir := filepath.Dir(rootPath)
		if FileExists(filepath.Join(dir, "create.yaml")) && FileExists(filepath.Join(dir, "_http.log")) {
			return []string{dir}, nil
		}
		return nil, fmt.Errorf("%s is not a valid fixture directory", rootPath)
	}

	if FileExists(filepath.Join(rootPath, "create.yaml")) && FileExists(filepath.Join(rootPath, "_http.log")) {
		return []string{rootPath}, nil
	}

	var dirs []string
	err = filepath.WalkDir(rootPath, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			if FileExists(filepath.Join(p, "create.yaml")) && FileExists(filepath.Join(p, "_http.log")) {
				dirs = append(dirs, p)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(dirs)
	return dirs, nil
}

// FileExists reports whether path exists on the local filesystem.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsNonDeletable reports whether the primary resource in the fixture does not issue a DELETE RPC on teardown.
func (d *Data) IsNonDeletable() bool {
	if d == nil || d.PrimaryCreate == nil {
		return false
	}
	for _, meta := range []*KRMResource{d.PrimaryCreate, d.PrimaryUpdate} {
		if meta == nil || meta.Obj == nil {
			continue
		}
		ann := meta.Obj.GetAnnotations()
		if strings.EqualFold(ann["cnrm.cloud.google.com/deletion-policy"], "abandon") {
			return true
		}
		if dp, _, _ := unstructured.NestedString(meta.Obj.Object, "spec", "deletionPolicy"); strings.Contains(strings.ToUpper(dp), "ABANDON") {
			return true
		}
	}
	switch d.PrimaryCreate.Kind {
	case "IAMPolicy", "IAMPartialPolicy", "IAMPolicyMember", "IAMAuditConfig",
		"ComputeProjectMetadata", "ComputeSharedVPCHostProject", "Service", "ServiceIdentity",
		"KMSAutokeyConfig", "KMSKeyRing", "KMSCryptoKey", "KMSCryptoKeyVersion", "KMSKeyRingImportJob",
		"FirebaseProject", "FirebaseWebApp", "FirebaseAndroidApp",
		"LoggingLogBucket", "Project", "Folder":
		return true
	}
	return false
}
