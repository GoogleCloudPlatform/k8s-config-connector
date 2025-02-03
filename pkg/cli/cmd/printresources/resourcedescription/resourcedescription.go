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

package resourcedescription

import (
	"context"
	"fmt"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ResourceDescription struct {
	GVK                schema.GroupVersionKind
	ResourceNameFormat string
	SupportsExport     bool
	SupportsBulkExport bool
	SupportsIAM        bool // true if the Resource supports IAM binding export alongside the resource
}

func GetAll() ([]ResourceDescription, error) {
	ctx := context.TODO()

	smLoader, err := servicemappingloader.New()
	if err != nil {
		return nil, fmt.Errorf("error creating service mapping loader: %w", err)
	}
	// tfprovider.New(...) configures the provider which requires valid access credentials. This is unnecessary
	// for the purposes of getting the schema information. To get around this but not create a new codepath for creating
	// a google provider, pass in an invalid, placeholder oauth2 token.
	tfProvider, err := tf.NewProvider(ctx, "invalid token")
	if err != nil {
		return nil, fmt.Errorf("error creating tf provider: %w", err)
	}
	resourceDescs := make([]ResourceDescription, 0)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			rDesc := getSMResourceDescription(tfProvider, sm, rc)
			resourceDescs = append(resourceDescs, rDesc)
		}
	}
	sortResourceDescriptions(resourceDescs)
	return resourceDescs, nil
}

func getSMResourceDescription(tfProvider *tfschema.Provider, sm v1alpha1.ServiceMapping, rc v1alpha1.ResourceConfig) ResourceDescription {
	supportsExport := doesResourceSupportExport(tfProvider, sm, rc)
	resourceNameFormat := ""
	if supportsExport {
		resourceNameFormat = getResourceNameFormat(sm, rc)
	}
	rDesc := ResourceDescription{
		GVK: schema.GroupVersionKind{
			Group:   sm.Name,
			Version: sm.GetVersionFor(&rc),
			Kind:    rc.Kind,
		},
		ResourceNameFormat: resourceNameFormat,
		SupportsBulkExport: doesResourceSupportBulkExport(tfProvider, sm, rc),
		SupportsExport:     supportsExport,
		SupportsIAM:        krmtotf.SupportsIAM(&rc),
	}
	return rDesc
}

func getResourceNameFormat(sm v1alpha1.ServiceMapping, rc v1alpha1.ResourceConfig) string {
	idTemplate := rc.IDTemplate
	// folder doesn't yet have a valid ID template due to some issues with the way the terraform provider is representing the ID
	if rc.Kind == "Folder" {
		idTemplate = "folders/{{folder_id}}"
	}
	return fmt.Sprintf("//%v/%v", sm.Spec.ServiceHostName, idTemplate)
}

func doesResourceSupportBulkExport(tfProvider *tfschema.Provider, _ v1alpha1.ServiceMapping, rc v1alpha1.ResourceConfig) bool {
	return rc.ResourceAvailableInAssetInventory && resourceHasTFImporter(rc, tfProvider)
}

func doesResourceSupportExport(tfProvider *tfschema.Provider, sm v1alpha1.ServiceMapping, rc v1alpha1.ResourceConfig) bool {
	if isHierarchicalResource(sm, rc) {
		return true
	}
	return resourceHasTFImporter(rc, tfProvider) && *rc.IDTemplateCanBeUsedToMatchResourceName
}

func resourceHasTFImporter(rc v1alpha1.ResourceConfig, tfProvider *tfschema.Provider) bool {
	// Every value for rc.Name should be in the ResourcesMap.
	//
	// TODO: remove 'Direct' field from ResourceConfig and remove the if statement.
	// The 'Direct' indicator won't be needed after we finish all the migrations.
	// The 'Direct' indicator is necessary during the migration so
	// that Config Connector uses direct approach to generate CRDs
	// but still allow TF-based controller to reconcile the resource.
	// Once a resource is migrated to direct, we need to figure out a different
	// way to support export for backwards compatibility if needed.
	if rc.Direct {
		return false
	}
	resource := tfProvider.ResourcesMap[rc.Name]
	return resource.Importer != nil
}

func isHierarchicalResource(sm v1alpha1.ServiceMapping, rc v1alpha1.ResourceConfig) bool {
	if sm.Spec.Name != "ResourceManager" {
		return false
	}
	return rc.Kind == "Project" || rc.Kind == "Folder"
}

func sortResourceDescriptions(resourceDescs []ResourceDescription) {
	sort.SliceStable(resourceDescs, func(i, j int) bool {
		if resourceDescs[i].GVK.Group == resourceDescs[j].GVK.Group {
			return resourceDescs[i].GVK.Kind < resourceDescs[j].GVK.Kind
		}
		return resourceDescs[i].GVK.Group < resourceDescs[j].GVK.Group
	})
}

// toMap converts a resourceDescription to a map keyed by GVK.
func toMap(rd []ResourceDescription) map[schema.GroupVersionKind]ResourceDescription {
	output := make(map[schema.GroupVersionKind]ResourceDescription, len(rd))
	for _, resourceDescription := range rd {
		output[resourceDescription.GVK] = resourceDescription
	}
	return output
}
