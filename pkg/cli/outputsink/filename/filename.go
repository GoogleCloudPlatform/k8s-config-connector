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

package filename

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	iamapi "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/gcpclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceskeleton"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/ghodss/yaml" //nolint:depguard
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var refFields = []string{"projectRef", "folderRef", "organizationRef"}

func Get(ctx context.Context, u *unstructured.Unstructured, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *tfschema.Provider) (string, error) {
	if iamapi.IsHandwrittenIAM(u.GroupVersionKind()) {
		return getIAM(ctx, u, smLoader, tfProvider)
	}
	parentPrefix, err := getParentPrefix(u, smLoader)
	if err != nil {
		return "", fmt.Errorf("error getting parent for %v '%v': %w", u.GetKind(), u.GetName(), err)
	}
	location, err := getLocation(u)
	if err != nil {
		return "", fmt.Errorf("error getting location field for %v '%v': %w", u.GetKind(), u.GetName(), err)
	}
	return path.Join(parentPrefix, u.GetKind(), location, u.GetName()), nil
}

func getIAM(ctx context.Context, u *unstructured.Unstructured, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *tfschema.Provider) (string, error) {
	resourceRefField := "resourceRef"
	resourceRef, found, err := getFieldAsIAMResourceRef(u, resourceRefField)
	if err != nil {
		return "", err
	}
	if !found {
		return "", fmt.Errorf("%v '%v' does not contain required field '%v'", u.GetKind(), u.GetName(), resourceRefField)
	}
	// turn the import id into a resource and then get the path for it
	sm, err := smLoader.GetServiceMapping(resourceRef.GroupVersionKind().Group)
	if err != nil {
		return "", fmt.Errorf("error getting service mapping for %v: %w", resourceRef.GroupVersionKind().Group, err)
	}
	uri := fmt.Sprintf("//%v/%v", sm.Spec.ServiceHostName, resourceRef.External)
	parentSkel, err := resourceskeleton.NewFromURI(uri, smLoader, tfProvider)
	if err != nil {
		// not all resources support NewFromURI, example: ServiceAccount, for those, fall back to the external value
		// converted to a legal filename -- this could benefit from the same refactor that would enable not fetching
		// hierarchal resources again. Example for ServiceAccount: IAMPolicy/ServiceAccount/full-external-path-resource-name.yaml
		name := MakeSafeFilename(fmt.Sprintf("%v-%v", resourceRef.External, u.GetName()))
		return path.Join(u.GetKind(), resourceRef.Kind, name), nil
	}
	if isHierarchalKind(resourceRef.Kind) {
		// we do not have enough information in the external id to properly reconstruct a hierarchal resource so we need to fetch it again
		// this fetch could be avoided if we calculated the for a resource and its associated IAMPolicy at the same time,
		// but that will require more refactoring and the benefit would be low (it would mean less GCP api requests)
		parent, err := getResource(ctx, parentSkel, smLoader, tfProvider)
		if err != nil {
			return "", fmt.Errorf("error getting parent resource for ref '%v': %w", *resourceRef, err)
		}
		parentSkel = parent
	}
	parentPath, err := Get(ctx, parentSkel, smLoader, tfProvider)
	if err != nil {
		return "", fmt.Errorf("error getting parent path for resource ref '%v': %w", *resourceRef, err)
	}
	//replace the last element of the path with the name of this iam resource
	lastSlashIdx := strings.LastIndex(parentPath, "/")
	if lastSlashIdx < 0 {
		return "", fmt.Errorf("invalid parent path '%v': does not contain a '/' character", parentPath)
	}
	return path.Join(parentPath[0:lastSlashIdx], u.GetName()), nil
}

func isHierarchalKind(kind string) bool {
	return kind == "Folder" || kind == "Project" || kind == "Organization"
}

func getResource(ctx context.Context, u *unstructured.Unstructured, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *tfschema.Provider) (*unstructured.Unstructured, error) {
	client := gcpclient.New(tfProvider, smLoader)
	return client.Get(ctx, u)
}

func getParentPrefix(u *unstructured.Unstructured, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	parentRefs, err := getParentResourceReferences(u, smLoader)
	if err != nil {
		return "", err
	}
	if len(parentRefs) == 0 {
		return getHierarchalParentPath(u)
	}
	// the assumption there can only be a single parent filled in at any time is carried over from tf/controller.go
	// so the loop terminates if a parent is found
	for _, refConfig := range parentRefs {
		resourceRef, found, err := getFieldAsKCCResourceRef(u, refConfig.Key)
		if err != nil {
			return "", err
		}
		if !found {
			continue
		}
		// it is assumed only external references are being used as that is all that is supported by the CLI
		external := trimExternalRef(resourceRef.External)
		if external == "" {
			return "", fmt.Errorf("field '%v.external' contains an unexpected empty value in spec.%v '%v'",
				refConfig.Key, u.GetKind(), u.GetName())
		}
		if !isHierarchalQualifier(strings.Split(external, "/")[0]) {
			hierarchalParent, err := getHierarchalParentPath(u)
			if err != nil {
				return "", err
			}
			external = fmt.Sprintf("%v/%v", hierarchalParent, external)
		}
		// insert the kind of the parent to the path
		splits := strings.Split(external, "/")
		if len(splits) < 2 {
			return "", fmt.Errorf("external name %q was not valid (gvk=%q, name=%q)", external, u.GroupVersionKind(), u.GetName())
		}
		newSplits := make([]string, 0, len(splits)+1)
		newSplits = append(newSplits, splits[0:2]...)
		newSplits = append(newSplits, refConfig.GVK.Kind)
		newSplits = append(newSplits, splits[2:]...)
		return strings.Join(newSplits, "/"), nil
	}
	return "", nil
}

func getFieldAsKCCResourceRef(u *unstructured.Unstructured, fieldName string) (*v1alpha1.ResourceReference, bool, error) {
	key := fmt.Sprintf("spec.%v", fieldName)
	value, ok, err := unstructured.NestedMap(u.Object, strings.Split(key, ".")...)
	if err != nil {
		return nil, false, fmt.Errorf("error getting field '%v' from %v: %w", key, u.GetKind(), err)
	}
	if !ok {
		return nil, false, nil
	}
	resourceRef, err := toGeneralResourceReference(value)
	if err != nil {
		return nil, false, fmt.Errorf("error converting field '%v' in %v to resource reference: %w", key, u.GetKind(), err)
	}
	return &resourceRef, true, nil
}

func getFieldAsIAMResourceRef(u *unstructured.Unstructured, fieldName string) (*iamapi.ResourceReference, bool, error) {
	key := fmt.Sprintf("spec.%v", fieldName)
	value, ok, err := unstructured.NestedMap(u.Object, strings.Split(key, ".")...)
	if err != nil {
		return nil, false, fmt.Errorf("error getting field '%v' from %v: %w", key, u.GetKind(), err)
	}
	if !ok {
		return nil, false, nil
	}
	resourceRef, err := toIAMResourceReference(value)
	if err != nil {
		return nil, false, fmt.Errorf("error converting field '%v' in %v to resource reference: %w", key, u.GetKind(), err)
	}
	return &resourceRef, true, nil
}

// trimExternalRef will take an external ref and remove any portion of the path before the hierarchal portion
// for example,
//   - https://www.googleapis.com/compute/v1/projects/kcc-test/global/networks/default
//
// will turn into,
//   - projects/kcc-test/global/networks/default
func trimExternalRef(externalRef string) string {
	splits := strings.Split(externalRef, "/")
	for i, subStr := range splits {
		if isHierarchalQualifier(subStr) {
			return strings.Join(splits[i:], "/")
		}
	}
	return externalRef
}

func isHierarchalQualifier(s string) bool {
	return s == "projects" || s == "folders" || s == "organizations"
}

func getParentResourceReferences(u *unstructured.Unstructured, smLoader *servicemappingloader.ServiceMappingLoader) ([]v1alpha1.ReferenceConfig, error) {
	rc, err := smLoader.GetResourceConfig(u)
	if err != nil {
		return nil, fmt.Errorf("error getting resource config for %v: %w", u.GetKind(), err)
	}
	refs := make([]v1alpha1.ReferenceConfig, 0)
	for _, resourceRef := range rc.ResourceReferences {
		if resourceRef.Parent {
			refs = append(refs, resourceRef)
		}
	}
	return refs, nil
}

func getHierarchalParentPath(u *unstructured.Unstructured) (string, error) {
	val, ok, err := getHierarchalParentPathFromSpec(u)
	if err != nil {
		return "", err
	}
	if ok {
		return val, nil
	}
	return getHierarchalParentPathFromAnnotations(u)
}

func getHierarchalParentPathFromAnnotations(u *unstructured.Unstructured) (string, error) {
	annotations := u.GetAnnotations()
	if annotations == nil {
		return "", fmt.Errorf("%v '%v' unexpectedly contains no annotations", u.GetKind(), u.GetName())
	}
	if val, ok := annotations[k8s.ProjectIDAnnotation]; ok {
		if strings.HasPrefix(val, "projects/") {
			return val, nil
		}
		return path.Join("projects", val), nil
	}
	if val, ok := annotations[k8s.FolderIDAnnotation]; ok {
		if strings.HasPrefix(val, "folders/") {
			return val, nil
		}
		return path.Join("folders", val), nil
	}
	if val, ok := annotations[k8s.OrgIDAnnotation]; ok {
		if strings.HasPrefix(val, "organizations/") {
			return val, nil
		}
		return path.Join("organizations", val), nil
	}
	return "", fmt.Errorf("resource contains no hierarchal resource annotation, expected one of {%v}",
		strings.Join(k8s.ContainerAnnotations, ", "))
}

func getHierarchalParentPathFromSpec(u *unstructured.Unstructured) (string, bool, error) {
	for _, refField := range refFields {
		val, ok, err := getHierarchalParentPathFromSpecField(u, refField)
		if err != nil {
			return "", false, err
		}
		if ok {
			return val, true, nil
		}
	}
	return "", false, nil
}

func getHierarchalParentPathFromSpecField(u *unstructured.Unstructured, refFieldName string) (string, bool, error) {
	val, ok, err := unstructured.NestedMap(u.Object, "spec", refFieldName)
	if err != nil {
		return "", false, fmt.Errorf("error retrieving 'spec.%v' from unstructured with kind '%v': '%w'", refFieldName, u.GetKind(), err)
	}
	if !ok {
		return "", false, nil
	}
	resourceRef, err := toGeneralResourceReference(val)
	if err != nil {
		return "", false, fmt.Errorf("error converting field '%v' in '%v' to a resource reference: %w",
			refFieldName, u.GetKind(), err)
	}
	return resourceRef.External, true, nil
}

func toGeneralResourceReference(value map[string]interface{}) (v1alpha1.ResourceReference, error) {
	bytes, err := yaml.Marshal(value)
	if err != nil {
		return v1alpha1.ResourceReference{}, fmt.Errorf("error marshalling '%v' to YAML: %w", value, err)
	}
	var resourceReference v1alpha1.ResourceReference
	if err := yaml.Unmarshal(bytes, &resourceReference); err != nil {
		return v1alpha1.ResourceReference{}, fmt.Errorf("error unmarshalling to resource reference; %w", err)
	}
	return resourceReference, nil
}

func toIAMResourceReference(value map[string]interface{}) (iamapi.ResourceReference, error) {
	bytes, err := yaml.Marshal(value)
	if err != nil {
		return iamapi.ResourceReference{}, fmt.Errorf("error marshalling '%v' to YAML: %w", value, err)
	}
	var resourceReference iamapi.ResourceReference
	if err := yaml.Unmarshal(bytes, &resourceReference); err != nil {
		return iamapi.ResourceReference{}, fmt.Errorf("error unmarshalling to resource reference; %w", err)
	}
	return resourceReference, nil
}

func getLocation(u *unstructured.Unstructured) (string, error) {
	locationFields := []string{"spec.location", "spec.region", "spec.zone"}
	for _, field := range locationFields {
		val, ok, err := getStringField(u, field)
		if err != nil {
			return "", err
		}
		if ok {
			return val, err
		}
	}
	return "", nil
}

func getStringField(u *unstructured.Unstructured, path string) (string, bool, error) {
	val, ok, err := unstructured.NestedString(u.Object, strings.Split(path, ".")...)
	if err != nil {
		return "", false, fmt.Errorf("error retrieving '%v' from %v: %w", path, u.GetKind(), err)
	}
	if ok {
		return val, true, nil
	}
	return "", false, nil
}
