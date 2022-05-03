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

package repo

import (
	"os"
	"path/filepath"

	"github.com/golang/glog"
)

const serviceMappingDirEnvVar = "SERVICE_MAPPING_DIR"
const serviceMappingDir = "/configconnector/servicemappings/"

func GetServiceMappingsPathOrLogFatal() string {
	if dir, ok := os.LookupEnv(serviceMappingDirEnvVar); ok {
		val, err := filepath.Abs(dir)
		if err != nil {
			glog.Fatalf("error getting absolute path for '%v': %v", dir, err)
		}
		return val
	}
	repoRoot, err := GetRoot()
	if err == nil {
		return filepath.Join(repoRoot, "config", "servicemappings")
	}
	dir := serviceMappingDir
	absPath, err := filepath.Abs(dir)
	if err != nil {
		glog.Fatalf("error getting absolute path for '%v': %v", dir, err)
	}
	return absPath
}

func GetDCLSchemasPathOrLogFatal() string {
	return filepath.Join(GetRootOrLogFatal(), "vendor/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google")
}

func GetServiceAccountCRDPath() string {
	return filepath.Join(GetCRDsPath(), "apiextensions.k8s.io_v1_customresourcedefinition_iamserviceaccounts.iam.cnrm.cloud.google.com.yaml")
}

func GetServiceAccountKeyCRDPath() string {
	return filepath.Join(GetCRDsPath(), "apiextensions.k8s.io_v1_customresourcedefinition_iamserviceaccountkeys.iam.cnrm.cloud.google.com.yaml")
}

func GetResourcesSamplesPath() string {
	return filepath.Join(GetRootOrLogFatal(), "config", "samples", "resources")
}

func GetCRDsPath() string {
	return filepath.Join(GetRootOrLogFatal(), "config", "crds", "resources")
}

func GetResourcesSnippetsPath() string {
	return filepath.Join(GetRootOrLogFatal(), "config", "cloudcodesnippets")
}

func GetClusterRolesPath() string {
	return filepath.Join(GetRootOrLogFatal(), "config", "installbundle", "components", "clusterroles")
}

func GetG3ResourceReferenceTemplatesPath() string {
	return filepath.Join(GetRootOrLogFatal(), "scripts", "generate-google3-docs", "resource-reference", "templates/")
}

func GetG3ResourceReferenceGeneratedPath() string {
	return filepath.Join(GetRootOrLogFatal(), "scripts", "generate-google3-docs", "resource-reference", "generated", "resource-docs")
}

func GetG3ResourceListsTemplatePath() string {
	return filepath.Join(GetRootOrLogFatal(), "scripts", "generate-google3-docs", "resource-lists", "template.tmpl")
}

func GetG3ResourceListsGeneratedPath() string {
	return filepath.Join(GetRootOrLogFatal(), "scripts", "generate-google3-docs", "resource-lists", "generated")
}

func GetClientGenerationPath() string {
	return filepath.Join(GetRootOrLogFatal(), "scripts", "generate-go-crd-clients")
}

func GetTypesGeneratedApisPath() string {
	return filepath.Join(GetRootOrLogFatal(), "pkg", "clients", "generated", "apis")
}

func GetTypesTemplatePath() string {
	return filepath.Join(GetRootOrLogFatal(), "scripts", "generate-go-crd-clients")
}
