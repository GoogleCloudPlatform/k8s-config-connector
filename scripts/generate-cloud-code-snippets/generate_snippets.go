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

// This program generates snippets for Cloud Code using our samples.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/snippet/snippetgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"gopkg.in/yaml.v2"
)

const outputFileMode = 0600
const dirMode = 0700

func main() {
	// Clear all snippets in directory and ensure generated snippets are up-to-date
	snippetPath := repo.GetResourcesSnippetsPath()
	if err := os.RemoveAll(snippetPath); err != nil {
		log.Fatalf("error deleting dir %v: %v", snippetPath, err)
	}
	if err := os.Mkdir(snippetPath, dirMode); err != nil {
		log.Fatalf("error recreating dir %v: %v", snippetPath, err)
	}

	samplesPath := repo.GetResourcesSamplesPath()
	resources, err := fileutil.SubdirsIn(samplesPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range resources {
		sampleFilePath, err := snippetgeneration.PathToSampleFileUsedForSnippets(resource)
		if err != nil {
			log.Fatal(err)
		}

		content, err := ioutil.ReadFile(sampleFilePath)
		if err != nil {
			log.Fatalf("error reading file: %v", err)
		}

		snippet, err := snippetgeneration.SnippifyResourceConfig(content)
		if err != nil {
			log.Fatal(err)
		}

		sampleFileName := path.Base(sampleFilePath)
		err = outputSnippetToFile(sampleFileName, snippet)
		if err != nil {
			log.Fatalf("error writing snippet to file: %v", err)
		}
	}
}

func outputSnippetToFile(outputFileName string, s snippetgeneration.Snippet) error {
	outputPath := path.Join(repo.GetResourcesSnippetsPath(), outputFileName)
	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outputPath, b, outputFileMode)
	if err != nil {
		return err
	}
	return nil
}
