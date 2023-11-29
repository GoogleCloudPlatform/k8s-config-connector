// Copyright 2023 Google LLC
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
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned"
)

func main() {
	if err := Run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Run(ctx context.Context) error {
	// Get a rest.Config
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	restConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return fmt.Errorf("loading kubernetes client configuration: %w", err)
	}
	httpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		return fmt.Errorf("loading kubernetes http client: %w", err)
	}

	// Build a typed client
	client, err := versioned.NewForConfigAndClient(restConfig, httpClient)
	if err != nil {
		return fmt.Errorf("building dynamic client: %w", err)
	}

	// Populate the StorageBucket object we want to create
	resourceID := fmt.Sprintf("example-bucket-%d", time.Now().UnixNano())
	obj := &v1beta1.StorageBucket{}
	obj.SetNamespace("config-control")
	obj.SetName("example-bucket")
	obj.Spec.ResourceID = &resourceID

	// Create the object the object
	applied, err := client.StorageV1beta1().StorageBuckets(obj.GetNamespace()).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("applying object: %w", err)
	}
	b, err := yaml.Marshal(applied)
	if err != nil {
		return fmt.Errorf("converting to yaml: %w", err)
	}
	klog.Infof("applied object, is now %v", string(b))

	// Wait for the object to be ready
	klog.Infof("waiting for object to be ready")
	for {
		latest, err := client.StorageV1beta1().StorageBuckets(obj.GetNamespace()).Get(ctx, obj.GetName(), metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("checking if object is ready: %w", err)
		}

		ready := false
		for _, condition := range latest.Status.Conditions {
			if condition.Type == "Ready" && condition.Status == "True" {
				ready = true
			}
		}

		b, err := yaml.Marshal(latest)
		if err != nil {
			return fmt.Errorf("converting to yaml: %w", err)
		}
		klog.Infof("object is: %v", string(b))

		if ready {
			klog.Infof("object is ready")
			break
		}
		klog.Infof("object is not ready; will wait and try again")
		time.Sleep(2 * time.Second)
	}

	// Give users a chance to verify the object bucket exists etc
	klog.Infof("sleeping for 30 seconds before deleting")
	time.Sleep(30 * time.Second)

	// Delete the object
	if err := client.StorageV1beta1().StorageBuckets(obj.GetNamespace()).Delete(ctx, obj.GetName(), metav1.DeleteOptions{}); err != nil {
		return fmt.Errorf("deleting object: %w", err)
	}

	// Wait for the object to be deleted
	klog.Infof("waiting for object deletion")
	for {
		time.Sleep(2 * time.Second)

		latest, err := client.StorageV1beta1().StorageBuckets(obj.GetNamespace()).Get(ctx, obj.GetName(), metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			break
		}
		if err != nil {
			return fmt.Errorf("checking if object still exists: %w", err)
		}

		b, err := yaml.Marshal(latest)
		if err != nil {
			return fmt.Errorf("converting to yaml: %w", err)
		}
		klog.Infof("object still exists, is now %v", string(b))
	}
	klog.Infof("object deleted")

	return nil
}
