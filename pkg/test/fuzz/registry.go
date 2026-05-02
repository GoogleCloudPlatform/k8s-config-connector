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

package fuzz

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	registryCache = make(map[string]*protoregistry.Files)
	registryMu    sync.Mutex
)

// GetProtoRegistry loads a proto descriptor set from the given path and returns a registry.
// It caches the registry for future use.
func GetProtoRegistry(path string) (*protoregistry.Files, error) {
	if !filepath.IsAbs(path) {
		root, err := repoRoot()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(root, path)
	}

	registryMu.Lock()
	defer registryMu.Unlock()

	if r, ok := registryCache[path]; ok {
		return r, nil
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %q: %w", path, err)
	}

	fds := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(b, fds); err != nil {
		return nil, fmt.Errorf("unmarshalling %q: %w", path, err)
	}

	files, err := protodesc.NewFiles(fds)
	if err != nil {
		return nil, fmt.Errorf("building file description from %q: %w", path, err)
	}

	registryCache[path] = files
	return files, nil
}

func repoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("getting repo root: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}
