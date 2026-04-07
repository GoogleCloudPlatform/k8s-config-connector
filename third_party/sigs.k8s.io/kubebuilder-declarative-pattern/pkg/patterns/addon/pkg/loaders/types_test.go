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
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"sigs.k8s.io/kustomize/kyaml/filesys"
)

func Test_allowManifestChannelName(t *testing.T) {
	var testcast = []struct {
		description string
		name        string
		expected    bool
	}{
		{
			description: "success pattern",
			name:        "stable",
			expected:    true,
		},
		{
			description: "fail pattern which has hyphen",
			name:        "stable-alpha",
			expected:    false,
		},
		{
			description: "fail pattern which has prefix dot",
			name:        ".stable",
			expected:    false,
		},
	}
	for _, test := range testcast {
		t.Run(test.description, func(t *testing.T) {
			actual := allowedChannelName(test.name)
			expected := test.expected
			if actual != expected {
				t.Fatalf("expected %+v but got %+v", expected, actual)
			}
		})
	}
}

func Test_allowManifestId(t *testing.T) {
	var testcast = []struct {
		description string
		name        string
		expected    bool
	}{
		{
			description: "success pattern",
			name:        "v1.0",
			expected:    true,
		},
		{
			description: "fail pattern which has @",
			name:        "v1.0@stable",
			expected:    false,
		},
		{
			description: "fail pattern which has prefix dot",
			name:        ".v1.0",
			expected:    false,
		},
	}
	for _, test := range testcast {
		t.Run(test.description, func(t *testing.T) {
			actual := allowedManifestId(test.name)
			expected := test.expected
			if actual != expected {
				t.Fatalf("expected %+v but got %+v", expected, actual)
			}
		})
	}
}

func TestFSRepository_LoadChannel(t *testing.T) {
	var testcases = []struct {
		description   string
		name          string
		errorString   string
		channelString string
	}{
		{
			description: "success pattern",
			name:        "stable",
			errorString: "",
			channelString: `# Versions for the stable channel
manifests:
- name: nginx
  version: 0.1.0`,
		},
		{
			description: "fail pattern which has @ in name",
			name:        "stable@-stable",
			errorString: "invalid channel name: ",
			channelString: `# Versions for the stable channel
manifests:
- name: nginx
  version: 0.1.0`,
		},
		{
			description: "fail pattern which has invalid channel",
			name:        "xxxxx",
			errorString: "error reading channel ",
			channelString: `# Versions for the stable channel
manifests:
- name: nginx
  version: 0.1.0`,
		},
	}

	for _, test := range testcases {
		t.Run(test.description, func(t *testing.T) {
			fSys := filesys.MakeFsOnDisk()
			baseDir := "/tmp/packages/"
			err := fSys.MkdirAll(baseDir)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			defer fSys.RemoveAll(baseDir)

			filePath := filepath.Join(baseDir, "stable")
			err = fSys.WriteFile(filePath, []byte(test.channelString))
			if err != nil {
				t.Fatalf("writing channel file: %v", err)
			}

			expected := Channel{
				Manifests: []Version{
					{
						Package: "nginx",
						Version: "0.1.0",
					},
				},
			}

			ctx := context.Background()
			var fs = NewFSRepository("/tmp/packages/")

			actual, err := fs.LoadChannel(ctx, test.name)

			if err != nil {
				if strings.HasPrefix(err.Error(), test.errorString) == false {
					t.Fatalf("expected start with: \"%s\" but got: \"%s\"", test.errorString, err.Error())
				}
			} else {
				if !reflect.DeepEqual(*actual, expected) {
					t.Fatalf("expected %+v but got %+v", expected, actual)
				}
			}
		})
	}
}

func TestFSRepository_Latest(t *testing.T) {
	var testcases = []struct {
		description   string
		name          string
		expected      string
		channelString string
	}{
		{
			description: "success pattern",
			name:        "stable",
			expected:    "0.2.0",
			channelString: `# Versions for the stable channel
manifests:
- name: nginx
  version: 0.1.0
- name: nginx
  version: 0.2.0
- name: apache
  version: 1.1.0`,
		},
	}

	for _, test := range testcases {
		t.Run(test.description, func(t *testing.T) {
			ctx := context.TODO()

			fSys := filesys.MakeFsOnDisk()
			baseDir := "/tmp/packages/"
			err := fSys.MkdirAll(baseDir)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			defer fSys.RemoveAll(baseDir)

			filePath := filepath.Join(baseDir, "stable")
			err = fSys.WriteFile(filePath, []byte(test.channelString))
			if err != nil {
				t.Fatalf("writing channel file: %v", err)
			}

			var fs = NewFSRepository("/tmp/packages/")

			channel, err := fs.LoadChannel(ctx, test.name)
			if err != nil {
				t.Fatalf("error from fs.LoadChannel: %v", err)
			}
			actual, err := channel.Latest(ctx, "nginx")
			if err != nil {
				t.Fatalf("error from channel.Latest: %v", err)
			}

			if actual.Version != test.expected {
				t.Fatalf("expected %+v but got %+v", test.expected, actual.Version)
			}
		})
	}
}

func TestFSRepository_LoadManifest(t *testing.T) {

	fSys := filesys.MakeFsOnDisk()
	baseDir := "/tmp/packages/nginx/1.2.3/"
	err := fSys.MkdirAll(baseDir)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	defer fSys.RemoveAll(baseDir)

	filePath := filepath.Join(baseDir, "manifest.yaml")
	var manifestStr = `
	kind: Deployment
	metadata:
	  labels:
		app: nginx2
	  name: foo
	  annotations:
		app: nginx2
	spec:
	  replicas: 1
	---
	kind: Service
	metadata:
	  name: foo
	  annotations:
		app: nginx
	spec:
	  selector:
		app: nginx`

	err = fSys.WriteFile(filePath, []byte(manifestStr))
	if err != nil {
		t.Fatalf("writing manifest file: %v", err)
	}

	expected := map[string]string{
		filePath: manifestStr,
	}

	ctx := context.Background()
	var fs = NewFSRepository("/tmp")

	actual, err := fs.LoadManifest(ctx, "nginx", "1.2.3")

	if err != nil {
		t.Fatalf("loading manifest: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %+v but got %+v", expected, actual)
	}
}
