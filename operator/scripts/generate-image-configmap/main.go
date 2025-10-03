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

// This script will extract images of KCC components from stable `manifest.yaml` and generate a ConfigMap `image_configmap.yaml` under config/release directory.
// The ConfigMap will be used for Component Release Pipeline to pre-load and validate images deployed by the operator.
//
// run `make manifests' to invoke the script, or directly go run scripts/generate-image-configmap/main.go"
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/ghodss/yaml" //nolint:depguard
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	cnrmmanifest "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"
)

const (
	fileMode       = 0600
	outputFilename = "image_configmap.yaml"
	outputDir      = "config/gke-addon"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	log := klog.FromContext(ctx)

	channelName := "stable"

	cc := &corev1beta1.ConfigConnector{
		Spec: corev1beta1.ConfigConnectorSpec{
			Mode:                 k8s.NamespacedMode,
			GoogleServiceAccount: "someGSA",
		},
	}
	operatorSrcRoot := paths.GetOperatorSrcRootOrLogFatal()
	r := cnrmmanifest.NewLocalRepository(path.Join(operatorSrcRoot, "channels"))
	channel, err := r.LoadChannel(ctx, channelName)
	if err != nil {
		return fmt.Errorf("error loading %v channel: %w", channelName, err)
	}
	version, err := channel.Latest(ctx, cc.ComponentName())
	if err != nil {
		return fmt.Errorf("error resolving the version to deploy: %w", err)
	}
	if version == nil {
		return fmt.Errorf("could not find the latest version in channel %v", channelName)
	}

	log.Info("got latest version from channel", "version", version.Version)
	manifestStrs, err := r.LoadManifest(ctx, cc.ComponentName(), version.Version, cc)
	if err != nil {
		return fmt.Errorf("error loading manifest for package %v of version %v: %w", version.Package, version.Version, err)
	}
	objects := make([]*manifest.Object, 0)
	for _, str := range manifestStrs {
		m, err := manifest.ParseObjects(ctx, str)
		if err != nil {
			return fmt.Errorf("parsing manifest: %w", err)
		}
		objects = append(objects, m.Items...)
	}

	namespacedStrs, err := r.LoadNamespacedComponents(ctx, cc.ComponentName(), version.Version)
	if err != nil {
		return fmt.Errorf("error loading namespaced components for package %v of version %v: %w", version.Package, version.Version, err)
	}
	for _, str := range namespacedStrs {
		m, err := manifest.ParseObjects(ctx, str)
		if err != nil {
			return fmt.Errorf("parsing manifest: %w", err)
		}
		objects = append(objects, m.Items...)
	}

	cm := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "config-images",
			Namespace: k8s.OperatorSystemNamespace,
			Labels: map[string]string{
				"addonmanager.kubernetes.io/mode":       "Reconcile",
				"cnrm.cloud.google.com/operator-system": "true",
			},
			Annotations: map[string]string{
				"components.gke.io/image-map": "Images deployed by operator",
			},
		},
		Data: make(map[string]string, 0),
	}

	for _, obj := range objects {
		log.Info("found object", "kind", obj.Kind, "name", obj.GetName())

		// controller image
		if obj.Kind == "StatefulSet" && strings.Contains(obj.GetName(), "cnrm-controller-manager") {
			image, err := extractImageFromStatefulSet(obj.UnstructuredObject(), "manager")
			if err != nil {
				return fmt.Errorf("error resolving manager image: %w", err)
			}
			cm.Data["cnrm.controller"] = image
		}
		// deletion defender image
		if obj.Kind == "StatefulSet" && obj.GetName() == "cnrm-deletiondefender" {
			image, err := extractImageFromStatefulSet(obj.UnstructuredObject(), "deletiondefender")
			if err != nil {
				return fmt.Errorf("error resolving manager image: %w", err)
			}
			cm.Data["cnrm.deletiondefender"] = image
		}
		// unmanaged detector image
		if obj.Kind == "StatefulSet" && obj.GetName() == "cnrm-unmanaged-detector" {
			image, err := extractImageFromStatefulSet(obj.UnstructuredObject(), "unmanageddetector")
			if err != nil {
				return fmt.Errorf("error resolving manager image: %w", err)
			}
			cm.Data["cnrm.unmanageddetector"] = image
		}
		// webhook image
		if obj.Kind == "Deployment" && obj.GetName() == "cnrm-webhook-manager" {
			image, err := extractImageFromDeployment(obj.UnstructuredObject(), "webhook")
			if err != nil {
				return fmt.Errorf("error resolving webhook image: %w", err)
			}
			cm.Data["cnrm.webhook"] = image
		}
		// recorder image
		if obj.Kind == "Deployment" && obj.GetName() == "cnrm-resource-stats-recorder" {
			image, err := extractImageFromDeployment(obj.UnstructuredObject(), "recorder")
			if err != nil {
				return fmt.Errorf("error resolving recorder image: %w", err)
			}
			cm.Data["cnrm.recorder"] = image
		}
		// prom-to-sd sidecar image
		if obj.Kind == "Deployment" && obj.GetName() == "cnrm-resource-stats-recorder" {
			image, err := extractImageFromDeployment(obj.UnstructuredObject(), "prom-to-sd")
			if err != nil {
				return fmt.Errorf("error resolving prom-to-sd sidecar image: %w", err)
			}
			cm.Data["prom-to-sd"] = image
		}
	}
	outputFilepath := path.Join(operatorSrcRoot, outputDir, outputFilename)
	if err := outputConfigMapToFile(&cm, outputFilepath); err != nil {
		return fmt.Errorf("error writing ConfigMap %v to file: %w", cm, err)
	}
	log.Info("successfully generated the image configmap", "path", outputFilepath)
	return nil
}

func extractImageFromStatefulSet(obj *unstructured.Unstructured, containerName string) (string, error) {
	b, err := obj.MarshalJSON()
	if err != nil {
		return "", err
	}
	ss := &appsv1.StatefulSet{}
	if err := json.Unmarshal(b, ss); err != nil {
		return "", err
	}

	for _, container := range ss.Spec.Template.Spec.Containers {
		if container.Name == containerName {
			return container.Image, nil
		}
	}
	return "", fmt.Errorf("could not find container with name %v in StatefulSet %v", containerName, obj.GetName())
}

func extractImageFromDeployment(obj *unstructured.Unstructured, containerName string) (string, error) {
	b, err := obj.MarshalJSON()
	if err != nil {
		return "", err
	}
	ss := &appsv1.Deployment{}
	if err := json.Unmarshal(b, ss); err != nil {
		return "", err
	}

	for _, container := range ss.Spec.Template.Spec.Containers {
		if container.Name == containerName {
			return container.Image, nil
		}
	}
	return "", fmt.Errorf("could not find container with name %v in Deployment %v", containerName, obj.GetName())
}

func outputConfigMapToFile(crd *corev1.ConfigMap, outputFilepath string) error {
	crdBytes, err := yaml.Marshal(crd)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(outputFilepath, crdBytes, fileMode); err != nil {
		return err
	}
	return nil
}
