// Copyright 2022 Google LLC
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

// This program creates two directories in the root directory, THIRD_PARTY_NOTICES and MIRRORED_LIBRARY_SOURCE,
// that contain the licenses of our third-party code and MPL-mandated mirrored library source code.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	inputDir                 = "temp-vendor"
	thirdPartyNoticeDir      = "THIRD_PARTY_NOTICES"
	mirroredLibrarySourceDir = "MIRRORED_LIBRARY_SOURCE"
	dirMode                  = 0700
	fileMode                 = 0600
)

func main() {
	var files []string
	os.RemoveAll(thirdPartyNoticeDir)
	os.RemoveAll(mirroredLibrarySourceDir)

	// find all the LICENSE files in the vendor directory
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if !strings.Contains(path, "LICENSE") && !strings.Contains(path, "LICENCE") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking vendor directory: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		licensePath := strings.TrimPrefix(file, "temp-vendor/")
		repo, licenseFilename := splitLicensePath(licensePath)
		licenseURL := repoToLicenseURL(repo, licenseFilename)
		fmt.Println(licenseURL)

		outputFilename := thirdPartyNoticeDir + "/" + licensePath
		outputFileDir := thirdPartyNoticeDir + "/" + repo
		input, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := os.MkdirAll(outputFileDir, dirMode); err != nil {
			fmt.Printf("error creating output directory '%v': %v\n", outputFileDir, err)
			os.Exit(1)
		}

		// copy the license
		if err := ioutil.WriteFile(outputFilename, input, fileMode); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if licenseRequiresSourceCodeMirroring(input) {
			fmt.Printf("REQUIRES SOURCE MIRRORING: %v\n", file)
			outputSourceDir := mirroredLibrarySourceDir + "/" + repo

			if err := os.MkdirAll(outputSourceDir, dirMode); err != nil {
				fmt.Printf("error creating output directory '%v': %v\n", outputFileDir, err)
				os.Exit(1)
			}
			// need to remove the actual dir so 'cp' works
			os.Remove(outputSourceDir)

			sourceDir := "temp-vendor/" + repo
			cmd := exec.Command("cp", "-r", sourceDir, outputSourceDir)
			if output, err := cmd.CombinedOutput(); err != nil {
				fmt.Printf("error copying source code for '%v': %v", sourceDir, string(output))
				os.Exit(1)
			}
		}
	}
}

func licenseRequiresSourceCodeMirroring(licenseText []byte) bool {
	normalizedText := strings.ToLower(string(licenseText))
	licenseRequiresSourceCodeMirroring := map[string]bool{
		"mozilla public license":                      true,
		"common development and distribution license": true,
		"eclipse public license":                      true,
		"gnu general public license":                  true,
		"lesser general public license":               true,
	}

	for licenseType := range licenseRequiresSourceCodeMirroring {
		if strings.Contains(normalizedText, licenseType) {
			return true
		}
	}
	return false
}

func splitLicensePath(path string) (repo string, licenseFilename string) {
	splitPath := strings.Split(path, "/")
	repo = strings.Join(splitPath[:len(splitPath)-1], "/")
	licenseFilename = splitPath[len(splitPath)-1]
	return repo, licenseFilename
}

func repoToLicenseURL(repo string, licenseFilename string) string {
	if manualLicenseURLMapping[repo] != "" {
		return manualLicenseURLMapping[repo]
	}
	domain, repoRoot, subrepoPath := splitRepo(repo)

	licensePathInRepo := licenseFilename
	if subrepoPath != "" {
		licensePathInRepo = strings.Join([]string{subrepoPath, licenseFilename}, "/")
	}

	// TODO: instead of assuming "blob/master", link to the specific SHA we use
	switch domain {
	case "cloud.google.com":
		splitRepoRoot := strings.Split(repoRoot, "/")
		if len(splitRepoRoot) == 2 && splitRepoRoot[0] == "go" {
			return fmt.Sprintf("https://github.com/googleapis/google-cloud-go/blob/master/LICENSE")
		} else {
			panic(fmt.Sprintf("unrecognized repo under cloud.google.com: %v", repoRoot))
		}
	case "sigs.k8s.io":
		return fmt.Sprintf("https://github.com/kubernetes-sigs/%v/blob/master/%v", repoRoot, licensePathInRepo)
	case "github.com":
		return fmt.Sprintf("https://github.com/%v/blob/master/%v", repoRoot, licensePathInRepo)
	case "golang.org":
		if !strings.HasPrefix(repoRoot, "x/") {
			panic(fmt.Sprintf("unhandled domain for repo %v", repo))
		}
		newRepoRoot := strings.TrimLeft(repoRoot, "x/")
		if newRepoRoot == "tools" && strings.Contains(licensePathInRepo, "third_party") {
			// This SHA still contains the licenses for the third_party dir
			return fmt.Sprintf("https://github.com/golang/tools/blob/7414d4c1f71cec71978b1aec0539171a2e42d230/%v", licensePathInRepo)
		} else {
			return fmt.Sprintf("https://github.com/golang/%v/blob/master/%v", newRepoRoot, licensePathInRepo)
		}
	case "k8s.io":
		return fmt.Sprintf("https://github.com/kubernetes/%v/blob/master/%v", repoRoot, licensePathInRepo)
	case "go.uber.org":
		return fmt.Sprintf("https://github.com/uber-go/%v/blob/master/%v", repoRoot, licensePathInRepo)
	case "gopkg.in":
		switch repoRoot {
		case "yaml.v2":
			return fmt.Sprintf("https://github.com/go-yaml/yaml/blob/v2.2.2/%v", licensePathInRepo)
		case "yaml.v3":
			return fmt.Sprintf("https://github.com/go-yaml/yaml/blob/v3/%v", licensePathInRepo)
		case "dnaeon/go-vcr.v3":
			return fmt.Sprintf("https://github.com/dnaeon/go-vcr/blob/v3/%v", licensePathInRepo)
		case "evanphx/json-patch.v4":
			return fmt.Sprintf("https://github.com/evanphx/json-patch/blob/master/%v", licensePathInRepo)
		default:
			panic(fmt.Sprintf("unhandled domain for repo domain %v, root %v", domain, repoRoot))
		}
	case "go.opencensus.io":
		return fmt.Sprintf("https://github.com/census-instrumentation/opencensus-go/blob/master/%v", licensePathInRepo)
	case "honnef.co":
		return fmt.Sprintf("https://github.com/dominikh/go-tools/blob/master/%v", licensePathInRepo)
	case "go.opentelemetry.io":
		return fmt.Sprintf("https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/%v", licensePathInRepo)
	default:
		panic(fmt.Sprintf("unhandled domain %q for repo %q", domain, repo))
	}
}

func splitRepo(repo string) (domain string, repoRoot string, subrepoPath string) {
	splitRepo := strings.Split(repo, "/")
	domain = splitRepo[0]
	if len(splitRepo) == 2 {
		repoRoot = splitRepo[1]
	} else if len(splitRepo) > 2 {
		repoRoot = strings.Join(splitRepo[1:3], "/")
	}
	subrepoPath = ""
	if len(splitRepo) > 3 {
		subrepoPath = strings.Join(splitRepo[3:], "/")
	}
	return domain, repoRoot, subrepoPath
}

var manualLicenseURLMapping = map[string]string{
	"bitbucket.org/creachadair/stringset":                     "https://bitbucket.org/creachadair/stringset/src/master/LICENSE",
	"cel.dev/expr":                                            "https://github.com/google/cel-spec/blob/master/LICENSE",
	"cloud.google.com/go":                                     "https://github.com/googleapis/google-cloud-go/blob/master/LICENSE",
	"contrib.go.opencensus.io/exporter/prometheus":            "https://github.com/census-ecosystem/opencensus-go-exporter-prometheus/blob/master/LICENSE",
	"dario.cat/mergo":                                         "https://github.com/darccio/mergo/blob/master/LICENSE",
	"go.starlark.net":                                         "https://github.com/google/starlark-go/blob/master/LICENSE",
	"gomodules.xyz/jsonpatch/v2":                              "https://https://github.com/gomodules/jsonpatch/blob/master/LICENSE",
	"google.golang.org/api":                                   "https://github.com/googleapis/google-api-go-client/blob/master/LICENSE",
	"google.golang.org/api/googleapi/internal/uritemplates":   "https://github.com/googleapis/google-api-go-client/blob/master/googleapi/internal/uritemplates/LICENSE",
	"google.golang.org/api/internal/third_party/uritemplates": "https://github.com/googleapis/google-api-go-client/blob/master/internal/third_party/uritemplates/LICENSE",
	"google.golang.org/appengine":                             "https://github.com/golang/appengine/blob/master/LICENSE",
	"google.golang.org/genproto":                              "https://github.com/google/go-genproto/blob/master/LICENSE",
	"google.golang.org/genproto/googleapis/api":               "https://github.com/google/go-genproto/blob/master/LICENSE",
	"google.golang.org/genproto/googleapis/api/serviceusage":  "https://github.com/google/go-genproto/blob/master/LICENSE",
	"google.golang.org/genproto/googleapis/rpc":               "https://github.com/google/go-genproto/blob/master/LICENSE",
	"google.golang.org/grpc":                                  "https://github.com/grpc/grpc-go/blob/master/LICENSE",
	"google.golang.org/protobuf":                              "https://github.com/protocolbuffers/protobuf-go/blob/master/LICENSE",
	"gopkg.in/fsnotify.v1":                                    "https://github.com/fsnotify/fsnotify/blob/master/LICENSE",
	"gopkg.in/inf.v0":                                         "https://github.com/go-inf/inf/blob/master/LICENSE",
	"gopkg.in/tomb.v1":                                        "https://github.com/go-tomb/tomb/blob/v1/LICENSE",
	"gopkg.in/vmihailenco/msgpack.v4":                         "https://github.com/vmihailenco/msgpack/blob/master/LICENSE",
	"gopkg.in/warnings.v0":                                    "https://github.com/go-warnings/warnings/blob/master/LICENSE",
	"honnef.co/go/tools":                                      "https://github.com/dominikh/go-tools/blob/master/LICENSE",
	"honnef.co/go/tools/lint":                                 "https://github.com/dominikh/go-tools/blob/master/lint/LICENSE",
	"honnef.co/go/tools/ssa":                                  "https://github.com/dominikh/go-tools/blob/master/ssa/LICENSE",
}
