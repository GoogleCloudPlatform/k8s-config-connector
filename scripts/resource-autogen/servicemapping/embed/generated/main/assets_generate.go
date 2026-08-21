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

	cnrmvfsgen "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/vfsgen"
	generatedembed "github.com/GoogleCloudPlatform/k8s-config-connector/scripts/resource-autogen/servicemapping/embed/generated"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var inputFS http.FileSystem = cnrmvfsgen.ConsistentModTimeFileSystem{
		HTTPFS: generatedembed.AssetsDir,
	}
	err := vfsgen.Generate(inputFS, generatedembed.VfsgenOptions)
	if err != nil {
		log.Fatalln(fmt.Sprintf("error generating embedded service mappings: %w", err))
	}
}
