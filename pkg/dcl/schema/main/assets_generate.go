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

//go:build integration
// +build integration

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	pathpkg "path"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	cnrmvfsgen "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/vfsgen"

	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/vfsgen"
)

var (
	AssetsDir     = http.Dir(GetDCLSchemasPathOrLogFatal())
	outputPath    = path.Join(repo.GetCallerPackagePathOrLogFatal(), "../embed/dcl_assets_vfsdata.go")
	VfsgenOptions = vfsgen.Options{
		PackageName:  "embed",
		VariableName: "Assets",
		Filename:     outputPath,
	}
)

func GetDCLSchemasPathOrLogFatal() string {
	return filepath.Join(repo.GetRootOrLogFatal(), "temp-vendor/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google")
}

func main() {
	loader := metadata.New()
	var fs = filter.Keep(
		AssetsDir,
		func(path string, fi os.FileInfo) bool {
			if fi.IsDir() {
				if path == "/" {
					// keep the root dir
					return true
				}
				path := strings.TrimPrefix(path, "/")
				components := strings.Split(path, "/")
				s, found := loader.GetServiceMetadata(components[0])
				if len(components) == 1 && found {
					// include the service dir if the service is supported
					return true
				}
				if len(components) == 2 {
					for _, r := range s.Resources {
						// include the version dir if there is a resource with corresponding version
						if r.DCLVersion == components[1] {
							return true
						}
					}
					return false
				}
				return false
			}
			// only save OpenAPI schema .yaml files as assets
			if pathpkg.Ext(path) != ".yaml" {
				return false
			}
			service, version, resource := getServiceVersionResourceFromPath(path)
			sm, found := loader.GetServiceMetadata(service)
			if !found {
				return false
			}
			for _, r := range sm.Resources {
				gvk := metadata.GVKForResource(sm, r)
				kindWithoutService := k8s.KindWithoutServicePrefix(gvk)
				if strings.ToLower(resource) == strings.ToLower(kindWithoutService) && r.DCLVersion == version {
					return true
				}
			}
			return false
		},
	)
	var inputFS http.FileSystem = cnrmvfsgen.ConsistentModTimeFileSystem{
		HTTPFS: fs,
	}
	err := vfsgen.Generate(inputFS, VfsgenOptions)
	if err != nil {
		log.Fatalln(fmt.Sprintf("error generating embedded dcl schemas: %v", err))
	}
}

func getServiceVersionResourceFromPath(path string) (service, version, resource string) {
	p := strings.TrimPrefix(path, "/")
	components := strings.Split(strings.TrimSuffix(p, ".yaml"), "/")
	switch len(components) {
	case 2:
		service = components[0]
		resource = strings.ReplaceAll(components[1], "_", "")
		version = "ga"
	case 3:
		service = components[0]
		version = components[1]
		resource = strings.ReplaceAll(components[2], "_", "")
	default:
		panic(fmt.Sprintf("path to the dcl schema yaml file has invalid format: %v", path))
	}
	return service, version, resource
}
