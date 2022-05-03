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

// This program generates the admin and the viewer cluster roles for cnrm api groups.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	"github.com/ghodss/yaml"
	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const outputFileMode = 0600

func main() {
	clusterRolesPath := repo.GetClusterRolesPath()
	if _, err := os.Stat(clusterRolesPath); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("path %v does not exist probably because the path to cluster roles has changed", clusterRolesPath)
		}
		log.Fatal(err)
	}
	smLoader, err := servicemappingloader.New()
	if err != nil {
		log.Fatalf("error getting new service mapping loader: %v", err)
	}
	serviceMetadataLoader := dclmetadata.New()
	gvks := supportedgvks.All(smLoader, serviceMetadataLoader)

	apis := make(map[string]bool)
	for _, gvk := range gvks {
		apis[gvk.Group] = true
	}
	apiGroupList := make([]string, 0)
	for api, _ := range apis {
		apiGroupList = slice.IncludeString(apiGroupList, api)
	}

	viewerRole := &rbacv1.ClusterRole{
		TypeMeta: v1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRole",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "viewer",
			Labels: map[string]string{
				"rbac.authorization.k8s.io/aggregate-to-view": "true",
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: apiGroupList,
				Verbs:     []string{"get", "list", "watch"},
				Resources: []string{"*"},
			},
		},
	}
	adminRole := &rbacv1.ClusterRole{
		TypeMeta: v1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRole",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "admin",
			Labels: map[string]string{
				"rbac.authorization.k8s.io/aggregate-to-admin": "true",
				"rbac.authorization.k8s.io/aggregate-to-edit":  "true",
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: apiGroupList,
				Verbs:     []string{"get", "list", "watch", "create", "update", "patch", "delete"},
				Resources: []string{"*"},
			},
		},
	}

	viewerRoleFileName := "cnrm_viewer.yaml"
	if err := outputClusterRoleToFile(clusterRolesPath, viewerRoleFileName, viewerRole); err != nil {
		log.Fatalf("error generating %v in %v: %v", viewerRoleFileName, clusterRolesPath, err)
	}

	adminRoleFileName := "cnrm_admin.yaml"
	if err := outputClusterRoleToFile(clusterRolesPath, adminRoleFileName, adminRole); err != nil {
		log.Fatalf("error generating %v in %v: %v", adminRoleFileName, clusterRolesPath, err)
	}
}

func outputClusterRoleToFile(outputDirPath, outputFileName string, r *rbacv1.ClusterRole) error {
	outputPath := path.Join(outputDirPath, outputFileName)
	b, err := yaml.Marshal(r)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outputPath, b, outputFileMode)
	if err != nil {
		return err
	}
	return nil
}
