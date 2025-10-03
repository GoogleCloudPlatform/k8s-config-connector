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

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	cnrmmanifest "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/scripts/utils"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

const (
	fileMode = 0644
	rbacDir  = "config/rbac"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	channelName := "stable"

	manifests, err := loadManifests(channelName)
	if err != nil {
		return err
	}
	if err := copyDependencies(manifests); err != nil {
		return err
	}
	return nil
}

func copyDependencies(objects []*manifest.Object) error {
	if err := copyViewerRole(objects); err != nil {
		return err
	}
	return nil
}

func copyViewerRole(objects []*manifest.Object) error {
	viewerRoleKind := "ClusterRole"
	viewerRoleName := "cnrm-viewer"
	viewerRole, ok := findObject(objects, "ClusterRole", viewerRoleName)
	if !ok {
		return fmt.Errorf("unable to find %v '%v' in manifests", viewerRoleKind, viewerRoleName)
	}
	outputPath := path.Join(paths.GetOperatorSrcRootOrLogFatal(), rbacDir, "cnrm_viewer_role.yaml")
	if err := writeManifestObjectToFile(viewerRole, outputPath); err != nil {
		return err
	}
	return nil
}

func findObject(objects []*manifest.Object, kind, name string) (*manifest.Object, bool) {
	for _, o := range objects {
		if o.Kind == kind && o.GetName() == name {
			return o, true
		}
	}
	return nil, false
}

func writeManifestObjectToFile(object *manifest.Object, outputPath string) error {
	bytes, err := utils.UnstructToYaml(object.UnstructuredObject())
	if err != nil {
		return fmt.Errorf("error serializing %v '%v' to yaml", object.Kind, object.GetName())
	}
	if err := ioutil.WriteFile(outputPath, bytes, fileMode); err != nil {
		return fmt.Errorf("error writing unstructured %v '%v' to file", object.Kind, object.GetName())
	}
	return nil
}

func loadManifests(channelName string) ([]*manifest.Object, error) {
	ctx := context.Background()
	cc := &corev1beta1.ConfigConnector{
		Spec: corev1beta1.ConfigConnectorSpec{
			Mode: "namespaced",
		},
	}
	operatorSrcRoot := paths.GetOperatorSrcRootOrLogFatal()
	r := cnrmmanifest.NewLocalRepository(path.Join(operatorSrcRoot, "channels"))
	channel, err := r.LoadChannel(ctx, channelName)
	if err != nil {
		return nil, fmt.Errorf("error loading %v channel: %w", channelName, err)
	}
	version, err := channel.Latest(ctx, cc.ComponentName())
	if err != nil {
		return nil, fmt.Errorf("error resolving the version to deploy: %w", err)
	}
	if version == nil {
		return nil, fmt.Errorf("could not find the latest version in channel %v", channelName)
	}
	manifestStrs, err := r.LoadManifest(ctx, cc.ComponentName(), version.Version, cc)
	if err != nil {
		return nil, fmt.Errorf("error loading manifest for package %v of version %v: %w", version.Package, version.Version, err)
	}
	objects := make([]*manifest.Object, 0)
	for _, str := range manifestStrs {
		m, err := manifest.ParseObjects(ctx, str)
		if err != nil {
			return nil, fmt.Errorf("parsing manifest: %w", err)
		}
		objects = append(objects, m.Items...)
	}
	return objects, nil
}
