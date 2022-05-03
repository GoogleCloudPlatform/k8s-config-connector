/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loaders

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	semver "github.com/blang/semver/v4"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/yaml"
)

type Repository interface {
	LoadChannel(ctx context.Context, name string) (*Channel, error)
	LoadManifest(ctx context.Context, packageName string, id string) (map[string]string, error)
}

// FSRepository is a Repository backed by a filesystem
type FSRepository struct {
	basedir string
}

var _ Repository = &FSRepository{}

// NewFSRepository is the constructor for an FSRepository
func NewFSRepository(basedir string) *FSRepository {
	return &FSRepository{
		basedir: basedir,
	}
}

var safelistChannelName = "abcdefghijklmnopqrstuvwxyz"

// We validate the channel name - keeping it to a small subset helps with path traversal,
// and also ensures that we can back easily this by other stores (e.g. https)
func allowedChannelName(name string) bool {
	if !matchesSafelist(name, safelistChannelName) {
		return false
	}

	// Double check!
	if strings.HasPrefix(name, ".") {
		return false
	}

	return true
}

var safelistVersion = "abcdefghijklmnopqrstuvwxyz0123456789-."

func allowedManifestId(name string) bool {
	if !matchesSafelist(name, safelistVersion) {
		return false
	}

	// Double check!
	if strings.HasPrefix(name, ".") {
		return false
	}

	return true
}

func matchesSafelist(s string, safelist string) bool {
	for _, c := range s {
		if strings.IndexRune(safelist, c) == -1 {
			return false
		}
	}
	return true
}

func (r *FSRepository) LoadChannel(ctx context.Context, name string) (*Channel, error) {
	if !allowedChannelName(name) {
		return nil, fmt.Errorf("invalid channel name: %q", name)
	}

	log := log.Log
	log.WithValues("channel", name).WithValues("base", r.basedir).Info("loading channel")

	p := filepath.Join(r.basedir, name)
	b, err := os.ReadFile(p)
	if err != nil {
		log.WithValues("path", p).Error(err, "error reading channel")
		return nil, fmt.Errorf("error reading channel %s: %v", p, err)
	}

	channel := &Channel{}
	if err := yaml.Unmarshal(b, channel); err != nil {
		return nil, fmt.Errorf("error parsing channel %s: %v", p, err)
	}

	return channel, nil
}

func (r *FSRepository) LoadManifest(ctx context.Context, packageName string, id string) (map[string]string, error) {
	if !allowedManifestId(packageName) {
		return nil, fmt.Errorf("invalid package name: %q", id)
	}

	if !allowedManifestId(id) {
		return nil, fmt.Errorf("invalid manifest id: %q", id)
	}

	log := log.Log
	log.WithValues("package", packageName).Info("loading package")

	dirPath := filepath.Join(r.basedir, "packages", packageName, id)
	filesPath, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %v", dirPath, err)
	}
	result := make(map[string]string)
	for _, p := range filesPath {
		if p.IsDir() {
			log.V(2).Info("skipping directory", "directory", p.Name())
			continue
		}

		filePath := filepath.Join(dirPath, p.Name())
		b, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %v", filePath, err)
		}
		result[filePath] = string(b)
	}

	return result, nil
}

type Channel struct {
	Manifests []Version `json:"manifests,omitempty"`
}

type Version struct {
	Package string `json:"name"`
	Version string `json:"version"`
}

func (c *Channel) Latest(packageName string) (*Version, error) {
	var latest *Version
	for i := range c.Manifests {
		v := &c.Manifests[i]
		if v.Package != "" && v.Package != packageName {
			continue
		}
		if latest == nil {
			latest = v
		} else if latest.Compare(v) < 0 {
			// Tie-break by taking the later version
			latest = v
		}
	}

	return latest, nil
}

// Compare compares two Versions, returning >0 for l>r, =0 if l=r, <0 if l<r
func (l *Version) Compare(r *Version) int {
	// If the package name is specified, it "wins"
	if l.Package != r.Package {
		if l.Package == "" {
			return -1
		}
		if r.Package == "" {
			return 1
		}
	}

	lSemver, lErr := semver.ParseTolerant(l.Version)
	rSemver, rErr := semver.ParseTolerant(r.Version)
	if lErr != nil {
		log.Log.Info("invalid semver in version", "version", l)
		if rErr != nil {
			log.Log.Info("invalid semver in version", "version", r)
			return 0
		}
		return -1
	}
	if rErr != nil {
		log.Log.Info("invalid semver in version", "version", r)
		return 1
	}

	return lSemver.Compare(rSemver)
}
