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

package krmtotf

import (
	"fmt"
	"reflect"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Resource is a wrapper around k8s.Resource and adds information regarding its
// corresponding Terraform resource and maintains an original copy of the
// k8s.Resource.
type Resource struct {
	k8s.Resource `json:",inline"`

	Original *k8s.Resource `json:"-"`

	// Fields related to TF provider processing
	TFInfo         *terraform.InstanceInfo        `json:"-"`
	ResourceConfig corekccv1alpha1.ResourceConfig `json:"-"`
	TFResource     *tfschema.Resource             `json:"-"`
}

// NewResource returns a Resource, populating the Resource information from u.Kind,
// using the structs found in sm and p.
func NewResource(u *unstructured.Unstructured, sm *corekccv1alpha1.ServiceMapping, p *tfschema.Provider) (*Resource, error) {
	rc, err := servicemappingloader.GetResourceConfig(sm, u)
	if err != nil {
		return nil, err
	}
	resource, err := NewResourceFromResourceConfig(rc, p)
	if err != nil {
		return nil, err
	}

	r, err := k8s.NewResource(u)
	if err != nil {
		return nil, err
	}
	resource.Resource = *r

	// Intentionally re-create the K8s resource to create a separate copy.
	resource.Original, err = k8s.NewResource(u)
	if err != nil {
		return nil, err
	}

	if err := resource.ValidateResourceIDIfSupported(); err != nil {
		return nil, err
	}

	return resource, nil
}

func NewResourceFromResourceConfig(rc *corekccv1alpha1.ResourceConfig, p *tfschema.Provider) (*Resource, error) {
	tfResource, ok := p.ResourcesMap[rc.Name]
	// Pure Direct Resource does not have ResourceMap.
	//
	// TODO: remove 'Direct' field from ResourceConfig and remove the if statement.
	// The 'Direct' indicator won't be needed after we finish all the migrations.
	// The 'Direct' indicator is necessary during the migration so
	// that Config Connector uses direct approach to generate CRDs
	// but still allow TF-based controller to reconcile the resource.
	if rc.Direct {
		return &Resource{
			TFInfo: &terraform.InstanceInfo{
				Type: rc.Name,
			},
			ResourceConfig: *rc,
		}, nil
	}
	if !ok {
		return nil, fmt.Errorf("error getting TF resource: unknown resource %v", rc.Name)
	}
	resource := &Resource{
		TFInfo: &terraform.InstanceInfo{
			Type: rc.Name,
		},
		ResourceConfig: *rc,
		TFResource:     tfResource,
	}
	return resource, nil
}

func getServerGeneratedIDFromStatus(rc *corekccv1alpha1.ResourceConfig, gvk schema.GroupVersionKind, status map[string]interface{}) (string, bool, error) {
	statusOrObservedState := status
	if stateintospec.OutputOnlyFieldsAreUnderObservedState(gvk) {
		statusOrObservedState = getObservedStateFromStatus(status)
	}
	splitPath := text.SnakeCaseStrsToLowerCamelCaseStrs(
		strings.Split(rc.ServerGeneratedIDField, "."))

	return unstructured.NestedString(statusOrObservedState, splitPath...)
}

// DeepCopyObject is needed to implement the interface of client.Object.
func (r *Resource) DeepCopyObject() runtime.Object {
	panic("unexpected call to resource.DeepCopyObject(...)")
}

func (r *Resource) ValidateResourceIDIfSupported() error {
	if !SupportsResourceIDField(&r.ResourceConfig) {
		return nil
	}

	_, err := r.IsResourceIDConfigured()
	if err != nil {
		return fmt.Errorf("error validating '%s' field: %w", k8s.ResourceIDFieldPath, err)
	}
	return nil
}

func (r *Resource) ConstructServerGeneratedIDInStatusFromResourceID(c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	resourceID, foundInSpec, err := unstructured.NestedString(r.Spec, k8s.ResourceIDFieldName)
	if err != nil {
		return "", fmt.Errorf("error getting '%s': %w",
			k8s.ResourceIDFieldPath, err)
	}

	if !foundInSpec {
		return "", nil
	}

	if foundInSpec && resourceID == "" {
		return "", fmt.Errorf("the value of '%s' is invalid: '' (empty "+
			"string)", k8s.ResourceIDFieldPath)
	}

	resourceID, err = ResolveValueTemplate(
		r.ResourceConfig.ResourceID.ValueTemplate, resourceID, r, c, smLoader)
	if err != nil {
		return "", fmt.Errorf("error expanding resource ID: %w", err)
	}

	return resourceID, nil
}

func (r *Resource) SelfLinkAsID() (string, error) {
	selfLink, found, err := unstructured.NestedString(r.GetStatusOrObservedState(), k8s.SelfLinkFieldName)
	if err != nil {
		return "", fmt.Errorf("error getting '%s': %w",
			k8s.SelfLinkFieldName, err)
	}
	if !found {
		return "", fmt.Errorf("resource %s doesn't have a '%s' field", r.Name, k8s.SelfLinkFieldName)
	}
	return selfLink, nil
}

// GetImportID returns the Terraform import ID for the resource.
// TODO(kcc-eng): Require ID templates for all resources and remove all implicit defaults.
func (r *Resource) GetImportID(c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	template := r.ResourceConfig.IDTemplate
	if r.HasServerGeneratedIDField() {
		// when using a server generated id for import, ensure it is there before importing to get a more specific
		// error of type ServerGeneratedIDNotFoundError
		if template == "" {
			template = r.serverGeneratedIDToTemplate()
			if _, err := r.GetServerGeneratedID(); err != nil {
				return "", err
			}
		} else if r.serverGeneratedIDInIDTemplate() {
			if _, err := r.GetServerGeneratedID(); err != nil {
				return "", err
			}
		}
	} else {
		if template == "" {
			template = fmt.Sprintf("{{project?}}/{{%v}}", r.ResourceConfig.MetadataMapping.Name)
		}
	}
	value, err := expandTemplate(template, r, c, smLoader)
	if err != nil {
		// Some resources, e.g. Project, have (1) a server-generated ID and (2)
		// an ID template that doesn't contain the server-generated ID in it.
		// And they can be imported by either (1) or (2). The following if block
		// is to get import ID via (1) after failing to resolve (2).
		if r.shouldFallBackToServerGeneratedIDIfImportIDFails() {
			template = r.serverGeneratedIDToTemplate()
			return expandTemplate(template, r, c, smLoader)
		}
		return "", err
	}
	return value, nil
}

func (r *Resource) HasIDTemplate() bool {
	return r.ResourceConfig.IDTemplate != ""
}

func (r *Resource) HasServerGeneratedIDField() bool {
	return r.ResourceConfig.ServerGeneratedIDField != ""
}

func (r *Resource) serverGeneratedIDToTemplate() string {
	return ServerGeneratedIDToTemplate(&r.ResourceConfig)
}

func (r *Resource) shouldFallBackToServerGeneratedIDIfImportIDFails() bool {
	return r.HasServerGeneratedIDField() && !r.serverGeneratedIDInIDTemplate()
}

func (r *Resource) serverGeneratedIDInIDTemplate() bool {
	if !r.HasIDTemplate() || !r.HasServerGeneratedIDField() {
		return false
	}
	idTemplateFormOfServerGeneratedID := fmt.Sprintf("{{%v}}", r.ResourceConfig.ServerGeneratedIDField)
	return strings.Contains(r.ResourceConfig.IDTemplate, idTemplateFormOfServerGeneratedID)
}

// GetServerGeneratedID gets the value of the resource's server-generated ID.
// There are two cases:
// (1) If the resource supports a server-generated `spec.resourceID`, return
//
//	its value if specified.  If unspecified, continue to case (2) but
//	extract out the resource ID segment from the server-generated ID field
//	using the value template of the resource ID field.
//
// (2) If the resource doesn't support a server-generated `spec.resourceID`
//
//	field, then look up the field defined in ResourceConfig.ServerGeneratedIDField
//	in `status` and return its value. Note: this value is not a resource ID,
//	but a raw value in the status field.
func (r *Resource) GetServerGeneratedID() (string, error) {
	if SupportsResourceIDField(&r.ResourceConfig) && IsResourceIDFieldServerGenerated(&r.ResourceConfig) {
		id, exists, err := unstructured.NestedString(r.Spec, k8s.ResourceIDFieldName)
		if err != nil {
			return "", fmt.Errorf("error getting server-generated resource ID: %w", err)
		}
		if exists {
			if id == "" {
				return "", fmt.Errorf("invalid empty value for \"spec.%s\"",
					k8s.ResourceIDFieldName)
			}
			return id, nil
		}
	}

	// If the resource doesn't support a server-generated `spec.resourceID` or
	// if the field is not specified, fallback to resolve it from status.
	idInStatus, exists, err := getServerGeneratedIDFromStatus(&r.ResourceConfig, r.GroupVersionKind(), r.Status)
	if err != nil {
		return "", fmt.Errorf("error getting server-generated ID: %w", err)
	}
	if !exists {
		return "", k8s.NewServerGeneratedIDNotFoundError(r.GroupVersionKind(),
			k8s.GetNamespacedName(r))
	}

	if idInStatus == "" {
		return "", fmt.Errorf("invalid empty value for \"status.%s\"",
			text.SnakeCaseToLowerCamelCase(r.ResourceConfig.ServerGeneratedIDField))
	}

	if SupportsResourceIDField(&r.ResourceConfig) && IsResourceIDFieldServerGenerated(&r.ResourceConfig) {
		id, err := extractValueSegmentFromIDInStatus(idInStatus,
			r.ResourceConfig.ResourceID.ValueTemplate)
		if err != nil {
			return "", fmt.Errorf("error getting server-generated "+
				"resource ID from the value of '%s': %w", fmt.Sprintf("status.%s",
				text.SnakeCaseToLowerCamelCase(r.ResourceConfig.ServerGeneratedIDField)), err)
		}

		if id == "" {
			return "", fmt.Errorf("invalid empty value for server-generated resource ID")
		}
		return id, nil
	}
	return idInStatus, nil
}

// GetResourceID gets the resource's resource ID. The assumption is
// that the resource supports the `spec.resourceID` field.
// There are two cases:
// (1) If `spec.resourceID` is specified, return its value.
// (2) Otherwise, (happens during KCC upgrade or resource creation), fall back to:
//   - Value of `metadata.name` if the resource ID is user-specified.
//   - Value of the server generated ID field in status if the resource ID is
//     server-generated.
func (r *Resource) GetResourceID() (string, error) {
	resourceID, exists, err := unstructured.NestedString(r.Spec, k8s.ResourceIDFieldName)
	if err != nil {
		return "", fmt.Errorf("error getting the value of "+
			"\"spec.%s\": %w", k8s.ResourceIDFieldName, err)
	}

	if !exists {
		if !IsResourceIDFieldServerGenerated(&r.ResourceConfig) {
			resourceID = r.GetName()
		} else {
			resourceID, err = r.GetServerGeneratedID()
			if err != nil {
				return "", err
			}
		}
	}

	if resourceID == "" {
		return "", fmt.Errorf("invalid empty value for resource ID")
	}
	return resourceID, nil
}

func (r *Resource) Unreadable() bool {
	return r.ResourceConfig.Unreadable != nil && *r.ResourceConfig.Unreadable
}

// AllTopLevelFieldsAreImmutableOrComputed returns true if the resource schema only
// contains top level fields that are immutable and/or computed.
func (r *Resource) AllTopLevelFieldsAreImmutableOrComputed() bool {
	for _, schema := range r.TFResource.Schema {
		if !schema.Computed && !schema.ForceNew {
			return false
		}
	}
	return true
}

func getObservedStateFromStatus(status map[string]interface{}) map[string]interface{} {
	observedState, _, _ := unstructured.NestedMap(status, k8s.ObservedStateFieldName)
	return observedState
}

func (r *Resource) GetStatusOrObservedState() map[string]interface{} {
	if stateintospec.OutputOnlyFieldsAreUnderObservedState(r.GroupVersionKind()) {
		return getObservedStateFromStatus(r.Status)
	}
	return r.Status
}

func SupportsResourceIDField(rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.ResourceID.TargetField != ""
}

func IsResourceIDFieldServerGenerated(rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.ResourceID.TargetField == rc.ServerGeneratedIDField
}

func SupportsServerGeneratedIDField(rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.ServerGeneratedIDField != ""
}

func SupportsHierarchicalReferences(rc *corekccv1alpha1.ResourceConfig) bool {
	return len(rc.HierarchicalReferences) > 0
}

func SupportsIAM(rc *corekccv1alpha1.ResourceConfig) bool {
	emptyIAMConfig := corekccv1alpha1.IAMConfig{}
	return !reflect.DeepEqual(rc.IAMConfig, emptyIAMConfig)
}

func GVKForResource(sm *corekccv1alpha1.ServiceMapping, rc *corekccv1alpha1.ResourceConfig) schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   sm.Name,
		Version: sm.GetVersionFor(rc),
		Kind:    rc.Kind,
	}
}

func ServerGeneratedIDToTemplate(rc *corekccv1alpha1.ResourceConfig) string {
	return fmt.Sprintf("{{%v}}", rc.ServerGeneratedIDField)
}

func isMetadataMappingLabelsField(field string, rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.MetadataMapping.Labels != "" && field == rc.MetadataMapping.Labels
}

func isMetadataMappingNameField(field string, rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.MetadataMapping.Name != "" && field == rc.MetadataMapping.Name
}

func isServerGeneratedIDField(field string, rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.ServerGeneratedIDField != "" && field == rc.ServerGeneratedIDField
}
