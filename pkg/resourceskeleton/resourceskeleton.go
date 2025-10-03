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

package resourceskeleton

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/serviceclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	uri2 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceskeleton/uri"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const ProjectKind = "Project"

var ResourceManagerAPIGroupName = fmt.Sprintf("resourcemanager.%v", crdgeneration.APIDomain)

func NewProject(projectID string, smLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	sm, err := smLoader.GetServiceMapping(ResourceManagerAPIGroupName)
	if err != nil {
		return nil, fmt.Errorf("error getting service mapping for '%v': %w", ResourceManagerAPIGroupName, err)
	}
	u := &unstructured.Unstructured{}
	gvk := schema.GroupVersionKind{
		Group:   sm.Name,
		Version: sm.Spec.Version,
		Kind:    ProjectKind,
	}
	u.SetGroupVersionKind(gvk)
	u.SetName(projectID)
	annotations := make(map[string]string, 1)
	annotations[k8s.FolderIDAnnotation] = "skeleton-folder"
	u.SetAnnotations(annotations)
	return u, nil
}

func NewFromURI(uri string, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *tfschema.Provider) (*unstructured.Unstructured, error) {
	parsedURL, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("error parsing '%v' as url: %w", uri, err)
	}
	canonicalHost := trimRegionPrefix(parsedURL.Host) // e.g. "us-central1-aiplatform.googleapis.com" -> "aiplatform.googleapis.com"
	sm, rc, err := uri2.GetServiceMappingAndResourceConfig(smLoader, canonicalHost, parsedURL.Path)
	if err != nil {
		return nil, fmt.Errorf("error getting service mapping and resource config for url '%v': %w", uri, err)
	}
	tfInfo := terraform.InstanceInfo{
		Type: rc.Name,
	}
	state, err := krmtotf.ImportState(context.Background(), strings.TrimPrefix(parsedURL.Path, "/"), &tfInfo, tfProvider)
	if err != nil {
		return nil, fmt.Errorf("error importing resource name to TF state: %w", err)
	}
	resource, err := tfStateToResource(state, sm, rc, tfProvider)
	if err != nil {
		return nil, fmt.Errorf("error creating new resource: %w", err)
	}
	return resource.MarshalAsUnstructured()
}

func NewFromAsset(a *asset.Asset, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *tfschema.Provider, serviceClient serviceclient.ServiceClient) (*unstructured.Unstructured, error) {
	sm, rc, err := asset.GetServiceMappingAndResourceConfig(smLoader, a)
	if err != nil {
		return nil, err
	}
	tfInfo := terraform.InstanceInfo{
		Type: rc.Name,
	}
	name := trimServiceHostName(a, sm)
	importID, err := convertAssetNameToImportID(rc, name)
	if err != nil {
		return nil, fmt.Errorf("error converting cloud asset inventory name '%v' to resource id: %w", name, err)
	}
	state, err := krmtotf.ImportState(context.Background(), importID, &tfInfo, tfProvider)
	if err != nil {
		return nil, fmt.Errorf("error importing resource name to TF state: %w", err)
	}
	resource, err := tfStateToResource(state, sm, rc, tfProvider)
	if err != nil {
		return nil, fmt.Errorf("error creating new resource: %w", err)
	}
	err = applyAssetKRMResourceHacks(resource, a, serviceClient, state)
	if err != nil {
		return nil, fmt.Errorf("unable to apply asset KRM hacks on asset %v: %w", a, err)
	}
	return resource.MarshalAsUnstructured()
}

func tfStateToResource(state *terraform.InstanceState, sm *v1alpha1.ServiceMapping, rc *v1alpha1.ResourceConfig, tfProvider *tfschema.Provider) (*krmtotf.Resource, error) {
	resource, err := krmtotf.NewResourceFromResourceConfig(rc, tfProvider)
	if err != nil {
		return nil, fmt.Errorf("error creating new resource: %w", err)
	}
	gvk := schema.GroupVersionKind{
		Group:   sm.Name,
		Version: sm.GetVersionFor(rc),
		Kind:    rc.Kind,
	}

	resource.SetGroupVersionKind(gvk)
	resource.Spec, resource.Status = krmtotf.GetSpecAndStatusFromState(resource, state)
	resource.Labels = krmtotf.GetLabelsFromState(resource, state)
	resource.Annotations = krmtotf.GetAnnotationsFromState(resource, state)
	resource.Name = krmtotf.GetNameFromState(resource, state)
	return resource, nil
}

func trimServiceHostName(a *asset.Asset, sm *v1alpha1.ServiceMapping) string {
	return strings.TrimPrefix(a.Name, fmt.Sprintf("//%v/", sm.Spec.ServiceHostName))
}

// convertAssetNameToImportID converts the name of the resource in Asset Inventory into
// the import ID of the resource in KCC.
func convertAssetNameToImportID(rc *v1alpha1.ResourceConfig, name string) (string, error) {
	// IAMCustomRole is a custom resource, and has a bespoke ID format.
	if rc.Kind == "IAMCustomRole" {
		id, err := parseIAMCustomRoleID(name)
		if err != nil {
			return "", fmt.Errorf("unable to parse IAMCustomRole id: %w", err)
		}
		switch id.parentType {
		case Project:
			return fmt.Sprintf("%v##%v", id.parentID, id.roleID), nil
		case Organization:
			return fmt.Sprintf("#%v#%v", id.parentID, id.roleID), nil
		}
	}
	if rc.Kind == "MonitoringAlertPolicy" {
		partitions := strings.Split(name, "/")
		if len(partitions) != 4 {
			return "", fmt.Errorf("expected 4 partitions split by '/' for '%v'", name)
		}
		return fmt.Sprintf("%v projects/%v/alertPolicies/%v", partitions[1], partitions[1], partitions[3]), nil
	}

	return name, nil
}

// Apply any hacks that need to be made because we have not come up with the appropriate abstraction in service mappings (yet)
func applyAssetKRMResourceHacks(resource *krmtotf.Resource, a *asset.Asset, client serviceclient.ServiceClient, state *terraform.InstanceState) error {
	if resource.Kind == "StorageBucket" {
		// the storage bucket properly uses the container annotation of 'project', however, it is only used and
		// verified on creation. The project id is not part of the ResourceName in the asset, but the resource
		// skeleton is not useable without it as krmtotf, etc, need the project-id annotation since StorageBucket requires
		// the project id annotation. For that reason, use the project number in its place as it is mostly correct and
		// will allow the resource to be useable by the rest of the system
		for _, ancestor := range a.Ancestors {
			if strings.HasPrefix(ancestor, "projects/") {
				projectNumber := strings.Replace(ancestor, "projects/", "", 1)
				project, err := client.GetProjectFromProjectIDOrNumber(projectNumber)
				if err != nil {
					return err
				}
				resource.Annotations[k8s.ProjectIDAnnotation] = project.ProjectId
			}
		}
	} else if resource.Kind == "IAMCustomRole" {
		id, err := parseIAMCustomRoleID(state.ID)
		if err != nil {
			return fmt.Errorf("unable to parse IAMCustomRole id: %w", err)
		}
		if resource.Spec == nil {
			resource.Spec = make(map[string]interface{})
		}
		resource.Spec[k8s.ResourceIDFieldName] = id.roleID
		switch id.parentType {
		case Project:
			resource.Annotations[k8s.ProjectIDAnnotation] = id.parentID
		case Organization:
			resource.Annotations[k8s.OrgIDAnnotation] = id.parentID
		}
	}
	return nil
}

type parentType int32

const (
	Project parentType = iota
	Organization
)

type iamCustomRoleID struct {
	parentType parentType
	parentID   string
	roleID     string
}

// parseIAMCustomRoleID parses an asset inventory ID for a Custom Role
// and returns its components.
func parseIAMCustomRoleID(id string) (*iamCustomRoleID, error) {
	partitions := strings.Split(id, "/")
	if len(partitions) != 4 {
		return nil, fmt.Errorf("expected 4 partitions split by '/' for '%v'", id)
	}
	value := iamCustomRoleID{
		parentID: partitions[1],
		roleID:   partitions[3],
	}
	switch partitions[0] {
	case "projects":
		value.parentType = Project
	case "organizations":
		value.parentType = Organization
	default:
		return nil, fmt.Errorf("expected 'projects' or 'organizations' for first partition, got '%v'", partitions[0])
	}
	return &value, nil
}

func trimRegionPrefix(host string) string {
	// e.g. "us-central1-aiplatform.googleapis.com" -> "aiplatform.googleapis.com"
	parts := strings.Split(host, "-")
	return parts[len(parts)-1]
}
