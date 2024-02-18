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

package gcpclient_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/gcpclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/serviceclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceskeleton"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	"github.com/ghodss/yaml"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// This test demonstrates how to use the gcp client, it does an Apply, Get, and then Delete for a ComputeDisk resource,
// note this file has a build rule of 'integration'
func TestQuickStart(t *testing.T) {
	ctx := context.TODO()

	client, smLoader, tfProvider, serviceClient, err := newDependencies(t)
	if err != nil {
		t.Fatalf("error creating dependencies: %v", err)
	}
	projectID := testgcp.GetDefaultProjectID(t)
	// define a basic disk resource with minimal fields
	diskSkeleton, err := newDiskSkeleton(projectID, "us-central1-a", "my-compute-disk")
	if err != nil {
		t.Fatalf("error creating disk skeleton: %v", err)
	}
	// apply the resource, this will create the resource if it does not exist or update the resource if it does exist
	disk, err := client.Apply(diskSkeleton)
	if err != nil {
		t.Fatalf("error creating disk: %v", err)
	}
	fmt.Println("Disk after applying:")
	printUnstructuredIndented(disk)
	selfLinkPath := "status.selfLink"
	selfLink, ok, err := unstructured.NestedString(disk.Object, strings.Split(selfLinkPath, ".")...)
	if !ok {
		t.Fatalf("expected disk's yaml to contain field '%v'", selfLinkPath)
	}
	if err != nil {
		t.Fatalf("error retrieving the value at '%v': %v", selfLinkPath, err)
	}
	// fetch the disk given only the self link
	diskAsset := asset.Asset{
		Name:      selfLink,
		AssetType: "compute.googleapis.com/Disk",
	}
	disk, err = getForAsset(ctx, client, smLoader, tfProvider, serviceClient, diskAsset)
	if err != nil {
		t.Fatalf("error getting yaml from asset '%v': %v", diskAsset, err)
	}
	fmt.Println(fmt.Sprintf("Disk from selflink '%v':", selfLink))
	printUnstructuredIndented(disk)
	// delete the disk
	if err := client.Delete(disk); err != nil {
		t.Fatalf("error deleting disk: %v", err)
	}
}

func getForAsset(ctx context.Context, client gcpclient.Client, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider, serviceClient serviceclient.ServiceClient, a asset.Asset) (*unstructured.Unstructured, error) {
	skeleton, err := resourceskeleton.NewFromAsset(&a, smLoader, tfProvider, serviceClient)
	if err != nil {
		return nil, fmt.Errorf("error converting asset to skeleton: %v", err)
	}
	return client.Get(ctx, skeleton)
}

func newDependencies(t *testing.T) (gcpclient.Client, *servicemappingloader.ServiceMappingLoader, *schema.Provider, serviceclient.ServiceClient, error) {
	ctx := context.TODO()
	smLoader, err := servicemappingloader.New()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error loading service mappings: %v", err)
	}
	config := tfprovider.Config{
		GCPAccessToken: "", // <- insert a valid oauth2 token here
	}
	tfProvider, err := tfprovider.New(ctx, config)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error loading terraform provider: %v", err)
	}
	serviceClient := serviceclient.NewMockServiceClient(t)
	return gcpclient.New(tfProvider, smLoader), smLoader, tfProvider, &serviceClient, nil
}

const (
	diskTemplate = `apiVersion: compute.cnrm.cloud.google.com/v1beta1
kind: ComputeDisk
metadata:
  name: ${name}
spec:
  projectRef:
    external: "${projectId}"
  description: an example disk without reference
  location: ${location}`
)

func printUnstructuredIndented(u *unstructured.Unstructured) {
	bytes, err := unstructToYAML(u)
	if err != nil {
		panic(fmt.Sprintf("error converting unstructured to YAML: %v", err))
	}
	stringVal := string(bytes)
	stringVal = strings.ReplaceAll(stringVal, "\n", "\n\t")
	stringVal = fmt.Sprintf("\t%v", stringVal)
	fmt.Println(stringVal)
}

func newDiskSkeleton(projectID, location, name string) (*unstructured.Unstructured, error) {
	diskYAML := strings.Replace(diskTemplate, "${projectId}", projectID, 1)
	diskYAML = strings.Replace(diskYAML, "${location}", location, 1)
	diskYAML = strings.Replace(diskYAML, "${name}", name, 1)
	var u *unstructured.Unstructured
	if err := yaml.Unmarshal([]byte(diskYAML), &u); err != nil {
		return nil, fmt.Errorf("error marshalling yaml to unstructured: %v", err)
	}
	return u, nil
}

func unstructToYAML(u *unstructured.Unstructured) ([]byte, error) {
	bytes, err := yaml.Marshal(u.Object)
	if err != nil {
		return nil, fmt.Errorf("error marshalling unstructured to yaml: %v", err)
	}
	return bytes, nil
}
