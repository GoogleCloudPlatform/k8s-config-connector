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

package crdgeneration

import (
	"fmt"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration/crdboilerplate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	APIDomain         = "cnrm.cloud.google.com"
	ManagedByKCCLabel = "cnrm.cloud.google.com/managed-by-kcc"
	GCPCategory       = "gcp"
)

// FileNameForCRD determines the file name for the given CRD.
// File names take the form of "$group_$version_$kind.yaml"
// Example: "pubsub_v1alpha1_pubsubtopic.yaml"
func FileNameForCRD(crd *apiextensions.CustomResourceDefinition) (string, error) {
	group, err := getGroup(crd)
	if err != nil {
		return "", err
	}
	version := k8s.GetVersionFromCRD(crd)
	kind := getKind(crd)
	fileName := strings.Join([]string{group, version, kind}, "_") + ".yaml"
	return fileName, nil
}

func getGroup(crd *apiextensions.CustomResourceDefinition) (string, error) {
	groupSplit := strings.SplitN(crd.Spec.Group, ".", 2)
	if len(groupSplit) != 2 {
		return "", fmt.Errorf("unable to parse group %v", crd.Spec.Group)
	}
	return groupSplit[0], nil
}

func getKind(crd *apiextensions.CustomResourceDefinition) string {
	return strings.ToLower(crd.Spec.Names.Kind)
}

func GenerateShortNames(kind string) []string {
	return []string{
		formatGCPShortName(kind),
		formatGCPShortName(text.Pluralize(kind)),
	}
}

func formatGCPShortName(kind string) string {
	return fmt.Sprintf("gcp%v", strings.ToLower(kind))
}

func GetCustomResourceDefinition(kind, group string, versions []string, storageVersion string, openAPIV3Schema *apiextensions.JSONSchemaProps, engineLabel string) *apiextensions.CustomResourceDefinition {
	// `storageVersion` must be unset if there is only one version.
	if len(versions) == 1 && storageVersion != "" {
		panic(fmt.Sprintf("invalid storage version %v: must be empty "+
			"when there is only one version", storageVersion))
	}
	// `storageVersion` is required when there are more than one version, and it
	// needs to be either `v1alpha1` or `v1beta1`.
	if len(versions) > 1 && !IsValidStorageVersion(storageVersion) {
		panic(fmt.Sprintf("invalid storage version %v: must be %v or "+
			"%v when there are more than one version", storageVersion,
			k8s.KCCAPIVersionV1Alpha1, k8s.KCCAPIVersionV1Beta1))
	}
	singular := strings.ToLower(kind)
	plural := text.Pluralize(singular)
	fullName := plural + "." + group
	crdVersions := make([]apiextensions.CustomResourceDefinitionVersion, len(versions))
	for i, version := range versions {
		// There should be only one storage version. When there is only one
		// version in the CRD, or if the current version is the same as the
		// `storageVersion` value, then the current version is the storage
		// version.
		storage := false
		if len(versions) == 1 || version == storageVersion {
			storage = true
		}
		// When v1alpha1 is supported and the storage version is a different
		// version, v1alpha1 is considered deprecated.
		deprecated := false
		if len(versions) > 1 && version == k8s.KCCAPIVersionV1Alpha1 &&
			version != storageVersion {
			deprecated = true
		}
		crdVersions[i] = apiextensions.CustomResourceDefinitionVersion{
			Schema: &apiextensions.CustomResourceValidation{
				OpenAPIV3Schema: openAPIV3Schema,
			},
			AdditionalPrinterColumns: crdboilerplate.GetAdditionalPrinterColumns(),
			Subresources: &apiextensions.CustomResourceSubresources{
				Status: &apiextensions.CustomResourceSubresourceStatus{},
			},
			Name:       version,
			Served:     true,
			Storage:    storage,
			Deprecated: deprecated,
		}
		if deprecated {
			deprecationWarning := fmt.Sprintf("%s/%s %s is deprecated, use %s/%s %s instead", group, version, kind, group, storageVersion, kind)
			crdVersions[i].DeprecationWarning = &deprecationWarning
		}
	}

	result := &apiextensions.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apiextensions.k8s.io/v1",
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fullName,
			Annotations: map[string]string{
				k8s.KCCVersionLabel: "0.0.0-dev",
			},
			Labels: map[string]string{
				ManagedByKCCLabel:  "true",
				engineLabel:        "true",
				k8s.KCCSystemLabel: "true",
			},
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:    group,
			Versions: []apiextensions.CustomResourceDefinitionVersion{},
			Scope:    apiextensions.NamespaceScoped,
			Names: apiextensions.CustomResourceDefinitionNames{
				Singular:   singular,
				Plural:     plural,
				Kind:       kind,
				Categories: []string{GCPCategory},
				ShortNames: GenerateShortNames(kind),
			},
		},
		Status: apiextensions.CustomResourceDefinitionStatus{
			Conditions:     []apiextensions.CustomResourceDefinitionCondition{},
			StoredVersions: []string{},
		},
	}

	for _, crdVersion := range crdVersions {
		result.Spec.Versions = append(result.Spec.Versions, crdVersion)
	}
	return result
}

func GenerateResourceIDFieldDescription(targetField string, isServerGeneratedResourceID bool) string {
	if isServerGeneratedResourceID {
		return fmt.Sprintf("Immutable. Optional. The service-generated "+
			"%s of the resource. Used for acquisition only. Leave unset to "+
			"create a new resource.", targetField)
	}

	return fmt.Sprintf("Immutable. Optional. The %s of the resource. "+
		"Used for creation and acquisition. When unset, the value of "+
		"`metadata.name` is used as the default.", targetField)
}

// MarkReferencedKindsNotSupported changes the description of the direct reference field 'name' to reflect that some of
// the referenced resource types are not yet supported in KCC.
func MarkReferencedKindsNotSupported(refSchema *apiextensions.JSONSchemaProps, kinds []string) {
	prop := refSchema.Properties["name"]
	prop.Description = fmt.Sprintf("[WARNING] %v not yet supported in Config Connector, use 'external' field to reference existing resources.\n%v", strings.Join(kinds, ","), prop.Description)
	refSchema.Properties["name"] = prop
}

// MarkHierarchicalReferencesOptionalButMutuallyExclusive returns a modified
// copy of the given JSON schema so that keys for hierarchical references are
// marked optional but mutually exclusive (i.e. at most one may be specified).
func MarkHierarchicalReferencesOptionalButMutuallyExclusive(spec *apiextensions.JSONSchemaProps, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) *apiextensions.JSONSchemaProps {
	specCopy := spec.DeepCopy()
	if !resourceSupportsHierarchicalRefs(hierarchicalRefs) {
		return specCopy
	}

	// Remove hierarchical references from the list of required fields in case
	// they're included.
	for _, h := range hierarchicalRefs {
		specCopy.Required = slice.RemoveStringFromStringSlice(specCopy.Required, h.Key)
	}

	// If only one hierarchical reference is supported, nothing else needs to
	// be done.
	if len(hierarchicalRefs) == 1 {
		return specCopy
	}

	// Add rule so that _at most one_ hierarchical reference can be specified.
	for _, h := range hierarchicalRefs {
		specCopy.OneOf = append(specCopy.OneOf, apiextensions.JSONSchemaProps{
			Required: []string{h.Key},
		})
	}
	canSpecifyNoRefRule := apiextensions.JSONSchemaProps{
		Not: &apiextensions.JSONSchemaProps{},
	}
	for _, h := range hierarchicalRefs {
		canSpecifyNoRefRule.Not.AnyOf = append(canSpecifyNoRefRule.Not.AnyOf, apiextensions.JSONSchemaProps{
			Required: []string{h.Key},
		})
	}
	specCopy.OneOf = append(specCopy.OneOf, canSpecifyNoRefRule)
	return specCopy
}

// MarkHierarchicalReferencesRequiredButMutuallyExclusive returns a modified
// copy of the given JSON schema so that keys for hierarchical references are
// marked required but mutually exclusive (i.e. one and only one must be
// specified).
func MarkHierarchicalReferencesRequiredButMutuallyExclusive(spec *apiextensions.JSONSchemaProps, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) *apiextensions.JSONSchemaProps {
	specCopy := spec.DeepCopy()
	if !resourceSupportsHierarchicalRefs(hierarchicalRefs) {
		return specCopy
	}

	// If only one hierarchical reference is supported, add it to the list of
	// required fields in case it's not already included.
	if len(hierarchicalRefs) == 1 {
		specCopy.Required = slice.IncludeString(specCopy.Required, hierarchicalRefs[0].Key)
		return specCopy
	}

	for _, h := range hierarchicalRefs {
		specCopy.OneOf = append(specCopy.OneOf, apiextensions.JSONSchemaProps{
			Required: []string{h.Key},
		})

		// Remove reference from the list of required fields in case it's
		// included since this would conflict with the OneOf rule.
		specCopy.Required = slice.RemoveStringFromStringSlice(specCopy.Required, h.Key)
	}
	return specCopy
}

func resourceSupportsHierarchicalRefs(hierarchicalRefs []corekccv1alpha1.HierarchicalReference) bool {
	return len(hierarchicalRefs) > 0
}

func IsValidStorageVersion(version string) bool {
	if version == k8s.KCCAPIVersionV1Beta1 || version == k8s.KCCAPIVersionV1Alpha1 {
		return true
	}
	return false
}
