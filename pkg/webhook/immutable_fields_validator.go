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

package webhook

import (
	"container/list"
	"context"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclcontainer "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension/container"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/typeutil"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	scheme              = runtime.NewScheme()
	codecs              = serializer.NewCodecFactory(scheme)
	ErrTFSchemaNotFound = fmt.Errorf("schema does not exist")
)

type immutableFieldsValidatorHandler struct {
	smLoader              *servicemappingloader.ServiceMappingLoader
	tfResourceMap         map[string]*tfschema.Resource
	dclSchemaLoader       dclschemaloader.DCLSchemaLoader
	serviceMetadataLoader dclmetadata.ServiceMetadataLoader
}

var (
	allowedResponse = admission.ValidationResponse(true, "admission controller passed")
)

func NewImmutableFieldsValidatorHandler(smLoader *servicemappingloader.ServiceMappingLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &immutableFieldsValidatorHandler{
			smLoader:              smLoader,
			tfResourceMap:         provider.ResourceMap(),
			dclSchemaLoader:       dclSchemaLoader,
			serviceMetadataLoader: serviceMetadataLoader,
		}
	}
}

func (a *immutableFieldsValidatorHandler) Handle(_ context.Context, req admission.Request) admission.Response {
	if regexp.MustCompile(ControllerManagerServiceAccountRegex).MatchString(req.AdmissionRequest.UserInfo.Username) {
		return admission.ValidationResponse(true, "ignore non-user requests")
	}

	// decode the existing object and the updated object
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding object: %w", err))
	}
	oldObj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.OldObject.Raw, nil, oldObj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding old object: %w", err))
	}

	spec, ok := obj.Object["spec"].(map[string]interface{})
	if obj.Object["spec"] != nil && !ok {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("the type of spec field is not map[string]interface{}"))
	}
	oldSpec, ok := oldObj.Object["spec"].(map[string]interface{})
	if oldObj.Object["spec"] != nil && !ok {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("the type of spec field is not map[string]interface{}"))
	}

	if isIAMResource(oldObj) {
		return validateImmutableFieldsForIAMResource(oldObj, oldSpec, spec)
	}

	if err := validateImmutableStateIntoSpecAnnotation(obj, oldObj); err != nil {
		return admission.Errored(http.StatusForbidden, err)
	}

	gk := oldObj.GroupVersionKind().GroupKind()
	switch gk {
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
		return validateImmutableFieldsForLoggingLogMetricResource(oldSpec, spec)
	case schema.GroupKind{Group: "gkehub.cnrm.cloud.google.com", Kind: "GKEHubFeatureMembership"}:
		return validateImmutableFieldsForGKEHubFeatureMembershipResource(oldSpec, spec)
	}

	isDirect := supportedgvks.IsDirectByGVK(obj.GroupVersionKind())
	if isDirect {
		return allowedResponse
	}

	if dclmetadata.IsDCLBasedResourceKind(obj.GroupVersionKind(), a.serviceMetadataLoader) {
		return validateImmutableFieldsForDCLBasedResource(obj, oldObj, spec, oldSpec, a.dclSchemaLoader, a.serviceMetadataLoader)
	}
	return validateImmutableFieldsForTFBasedResource(obj, oldObj, spec, oldSpec, a.smLoader, a.tfResourceMap)
}

func validateImmutableStateIntoSpecAnnotation(obj, oldObj *unstructured.Unstructured) error {
	val, found := k8s.GetAnnotation(k8s.StateIntoSpecAnnotation, obj)
	prevVal, prevFound := k8s.GetAnnotation(k8s.StateIntoSpecAnnotation, oldObj)
	if found != prevFound || val != prevVal {
		return fmt.Errorf("annotation %v is immutable", k8s.StateIntoSpecAnnotation)
	}
	return nil
}

func validateImmutableFieldsForDCLBasedResource(obj, oldObj *unstructured.Unstructured, spec, oldSpec map[string]interface{}, dclSchemaLoader dclschemaloader.DCLSchemaLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) admission.Response {
	gvk := obj.GroupVersionKind()
	schema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting the DCL Schema for GroupVersionKind %v: %w", gvk, err))
	}
	containers, err := dclcontainer.GetContainersForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting containers supported by GroupVersionKind %v: %w", gvk, err))
	}
	hierarchicalRefs, err := dcl.GetHierarchicalReferencesForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting hierarchical references supported by GroupVersionKind %v: %w", gvk, err))
	}
	if err := validateContainerAnnotationsForResource(gvk.Kind, obj.GetAnnotations(), oldObj.GetAnnotations(), containers, hierarchicalRefs); err != nil {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error validating container annotations: %w", err))
	}
	if isResourceIDModified(spec, oldSpec) {
		return admission.Errored(http.StatusForbidden,
			k8s.NewImmutableFieldsMutationError([]string{k8s.ResourceIDFieldPath}))
	}
	res, err := getChangesOnImmutableFields(spec, oldSpec, []string{"spec"}, []string{}, schema, hierarchicalRefs)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("unexpected error: %w", err))
	}
	if len(res) != 0 {
		return admission.Errored(http.StatusForbidden,
			k8s.NewImmutableFieldsMutationError(res))
	}
	return allowedResponse
}

func getChangesOnImmutableFields(spec, oldSpec map[string]interface{}, krmPath, dclPath []string, schema *openapi.Schema, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) ([]string, error) {
	if schema.Type != "object" {
		return nil, fmt.Errorf("expect the schame type to be 'object', but got %v", schema.Type)
	}
	ret := make([]string, 0)
	for f, s := range schema.Properties {
		if s.ReadOnly {
			continue
		}
		isImmutable, err := dclextension.IsImmutableField(s)
		if err != nil {
			return nil, fmt.Errorf("error determining if field %v is immutable", f)
		}
		if dclextension.IsReferenceField(s) {
			if !isImmutable {
				continue
			}
			// If parent reference field is immutable and resource supports
			// multiple parent types, changes to either the hierarchical
			// reference key or value should be rejected.
			if dcl.IsMultiTypeParentReferenceField(append(dclPath, f)) {
				for _, h := range hierarchicalRefs {
					if !reflect.DeepEqual(oldSpec[h.Key], spec[h.Key]) {
						krmPathToField := pathslice.ToString(append(krmPath, h.Key))
						ret = append(ret, krmPathToField)
					}
				}
				continue
			}
			refField, err := dclextension.GetReferenceFieldName(append(dclPath, f), s)
			if err != nil {
				return nil, err
			}
			if !reflect.DeepEqual(oldSpec[refField], spec[refField]) {
				krmPathToField := pathslice.ToString(append(krmPath, refField))
				ret = append(ret, krmPathToField)
			}
			continue
		}
		krmPathToField := pathslice.ToString(append(krmPath, f))
		oldVal := oldSpec[f]
		newVal := spec[f]
		if oldVal == nil && newVal == nil {
			continue
		}
		if isImmutable && (oldVal == nil || newVal == nil) {
			ret = append(ret, krmPathToField)
			continue
		}
		switch s.Type {
		case "object":
			var v1 map[string]interface{}
			var v2 map[string]interface{}
			if oldVal != nil {
				v1 = oldVal.(map[string]interface{})
			}
			if newVal != nil {
				v2 = newVal.(map[string]interface{})
			}
			// Field is a map of key-value pairs
			if s.AdditionalProperties != nil {
				// If map field is immutable, reject any mutations.
				if isImmutable {
					if !reflect.DeepEqual(v1, v2) {
						ret = append(ret, krmPathToField)
					}
					continue
				}
				if typeutil.IsPrimitiveType(s.AdditionalProperties.Type) {
					continue
				}
				// If map field is mutable, but its key-value pairs have
				// non-primitive values (e.g. objects, arrays of objects), the
				// values themselves may have immutable fields. For now, let
				// such cases pass through the webhook, and let DCL or the GCP
				// API handle them instead (see b/216381382).
				continue
			}
			// Field is an object
			nestedFields, err := getChangesOnImmutableFields(v1, v2, append(krmPath, f), append(dclPath, f), s, hierarchicalRefs)
			if err != nil {
				return nil, err
			}
			ret = append(ret, nestedFields...)
		case "array":
			if typeutil.IsPrimitiveType(s.Items.Type) {
				if isImmutable && !reflect.DeepEqual(oldVal, newVal) {
					ret = append(ret, krmPathToField)
				}
				continue
			}
			// Kubernetes considers all lists of objects to be atomic, and so all subsequent
			// applies will currently wipe out defaulted immutable fields. Temporarily delegate validation to DCL.
		case "string", "boolean", "number", "integer":
			if isImmutable && !reflect.DeepEqual(oldSpec[f], spec[f]) {
				ret = append(ret, krmPathToField)
			}
		default:
			return nil, fmt.Errorf("unknown schema type %v", s.Type)
		}
	}
	return ret, nil
}

func getQualifiedFieldName(prefix string, fieldName string) string {
	qualifiedName := fieldName
	if prefix != "" {
		qualifiedName = prefix + "." + fieldName
	}
	return qualifiedName
}

func validateImmutableFieldsForTFBasedResource(obj, oldObj *unstructured.Unstructured, spec, oldSpec map[string]interface{}, smLoader *servicemappingloader.ServiceMappingLoader, tfResourceMap map[string]*tfschema.Resource) admission.Response {
	rc, err := smLoader.GetResourceConfig(obj)
	if err != nil {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("couldn't get ResourceConfig for kind %v: %w", obj.GetKind(), err))
	}

	if err := validateContainerAnnotationsForResource(obj.GetKind(), obj.GetAnnotations(), oldObj.GetAnnotations(), rc.Containers, rc.HierarchicalReferences); err != nil {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error validating container annotations: %w", err))
	}

	r, ok := tfResourceMap[rc.Name]
	if !ok {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("unknown resource %v", rc.Name))
	}

	if findChangesOnImmutableResourceIDField(spec, oldSpec, rc) {
		return admission.Errored(http.StatusForbidden,
			k8s.NewImmutableFieldsMutationError([]string{k8s.ResourceIDFieldPath}))
	}

	if findChangesOnImmutableLocationField(spec, oldSpec, rc) {
		return admission.Errored(http.StatusForbidden,
			k8s.NewImmutableFieldsMutationError([]string{"spec.location"}))
	}

	fields := list.New()
	compareAndFindChangesOnImmutableFields(spec, oldSpec, r.Schema, "", rc, getIgnoredFields(rc), fields)
	if fields.Len() != 0 {
		res := make([]string, 0)
		for e := fields.Front(); e != nil; e = e.Next() {
			res = append(res, constructCamelCasePath(e.Value.(string)))
		}
		return admission.Errored(http.StatusBadRequest,
			k8s.NewImmutableFieldsMutationError(res))
	}

	return allowedResponse
}

func isImmutableFieldModified(oldSpec, newSpec map[string]interface{}, field string) bool {
	tokens := strings.Split(field, ".")
	oldVal, ok1, err1 := unstructured.NestedFieldCopy(oldSpec, tokens...)
	newVal, ok2, err2 := unstructured.NestedFieldCopy(newSpec, tokens...)
	if oldVal == nil && newVal == nil {
		return false
	}
	if !ok1 || err1 != nil {
		return true
	}
	if !ok2 || err2 != nil {
		return true
	}
	return !reflect.DeepEqual(oldVal, newVal)
}

func validateImmutableFieldsForGKEHubFeatureMembershipResource(oldSpec, spec map[string]interface{}) admission.Response {
	ImmutableFields := []string{"featureRef", "location", "projectRef", "membershipLocation", "membershipRef"}
	var res []string
	for _, field := range ImmutableFields {
		if isImmutableFieldModified(oldSpec, spec, field) {
			res = append(res, field)
		}
	}
	if len(res) != 0 {
		return admission.Errored(http.StatusForbidden,
			k8s.NewImmutableFieldsMutationError(res))
	}
	return allowedResponse
}

func validateImmutableFieldsForLoggingLogMetricResource(oldSpec, spec map[string]interface{}) admission.Response {
	if isResourceIDModified(oldSpec, spec) {
		return admission.Errored(http.StatusForbidden,
			k8s.NewImmutableFieldsMutationError([]string{k8s.ResourceIDFieldPath}))
	}
	ImmutableFields := []string{"metricDescriptor.metricKind", "metricDescriptor.valueType", "projectRef"}
	var res []string
	for _, field := range ImmutableFields {
		if isImmutableFieldModified(oldSpec, spec, field) {
			res = append(res, field)
		}
	}
	if len(res) != 0 {
		return admission.Errored(http.StatusForbidden,
			k8s.NewImmutableFieldsMutationError(res))
	}
	return allowedResponse
}

func validateContainerAnnotationsForResource(kind string, annotations, oldAnnotations map[string]string, containers []corekccv1alpha1.Container, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) error {
	// TODO(b/193177782): Delete this if-block once all resources support
	// hierarchical references.
	if len(hierarchicalRefs) == 0 {
		return validateContainerAnnotations(kind, annotations, oldAnnotations, containers)
	}
	return validateDeprecatedContainerAnnotations(annotations, oldAnnotations, containers, hierarchicalRefs)
}

func validateContainerAnnotations(kind string, annotations, oldAnnotations map[string]string, containers []corekccv1alpha1.Container) error {
	for _, c := range containers {
		a := k8s.GetAnnotationForContainerType(c.Type)

		// No changes to the container annotation.
		if oldAnnotations[a] == annotations[a] {
			continue
		}

		// Reject changes to container annotations except for Projects and
		// Folders which rely on container annotation updates to allow for
		// migrations across different parent Folders and Organizations.
		if kind != "Project" && kind != "Folder" {
			return fmt.Errorf("cannot make changes to container annotation %v", a)
		}

		// Reject changes from one container annotation type to another.
		for _, otherC := range containers {
			if c == otherC {
				continue
			}
			otherA := k8s.GetAnnotationForContainerType(otherC.Type)
			_, ok := oldAnnotations[a]
			_, otherOk := annotations[otherA]
			if ok && otherOk {
				return fmt.Errorf("cannot change from container annotation %v to container annotation %v", a, otherA)
			}
		}
	}
	return nil
}

func validateDeprecatedContainerAnnotations(annotations, oldAnnotations map[string]string, containers []corekccv1alpha1.Container, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) error {
	for _, c := range containers {
		a := k8s.GetAnnotationForContainerType(c.Type)

		// No changes to the container annotation.
		if oldAnnotations[a] == annotations[a] {
			continue
		}

		// Container annotation was removed. This is a change that we allow to
		// give users the ability to "clean" their configurations now that
		// container annotations have been deprecated for this resource.
		if annotations[a] == "" {
			continue
		}

		// Container annotation was either added or changed.
		possibleFields := k8s.HierarchicalReferencesToFields(hierarchicalRefs)
		return fmt.Errorf("cannot add/change container annotation %v as it is no longer supported by the resource; set one of [%v] instead", a, strings.Join(possibleFields, ", "))
	}
	return nil
}

func validateImmutableFieldsForIAMResource(oldObj *unstructured.Unstructured, oldSpec, newSpec map[string]interface{}) admission.Response {
	if isIAMPolicy(oldObj) {
		return handleIAMPolicy(oldSpec, newSpec)
	}
	if isIAMPartialPolicy(oldObj) {
		return handleIAMPartialPolicy(oldSpec, newSpec)
	}
	if isIAMPolicyMember(oldObj) {
		return handleIAMPolicyMember(oldSpec, newSpec)
	}
	if isIAMAuditConfig(oldObj) {
		return handleIAMAuditConfig(oldSpec, newSpec)
	}
	return admission.ValidationResponse(false, fmt.Sprintf("unknown IAM resource type: %v", oldObj.GroupVersionKind()))
}

func handleIAMPolicy(oldSpec, newSpec map[string]interface{}) admission.Response {
	if isIAMResourceReferenceModified(oldSpec, newSpec) {
		msg := fmt.Sprintf("the IAMPolicy's spec.resourceRef is immutable")
		return admission.ValidationResponse(false, msg)
	}
	return allowedResponse
}

func handleIAMPartialPolicy(oldSpec, newSpec map[string]interface{}) admission.Response {
	if isIAMResourceReferenceModified(oldSpec, newSpec) {
		msg := fmt.Sprintf("the IAMPartialPolicy's spec.resourceRef is immutable")
		return admission.ValidationResponse(false, msg)
	}
	return allowedResponse
}

func handleIAMPolicyMember(oldSpec, newSpec map[string]interface{}) admission.Response {
	if isIAMSpecModified(oldSpec, newSpec) {
		msg := fmt.Sprintf("the IAMPolicyMember's spec is immutable")
		return admission.ValidationResponse(false, msg)
	}
	return allowedResponse
}

func handleIAMAuditConfig(oldSpec, newSpec map[string]interface{}) admission.Response {
	if isIAMResourceReferenceModified(oldSpec, newSpec) {
		msg := fmt.Sprintf("the IAMAuditConfig's spec.resourceRef is immutable")
		return admission.ValidationResponse(false, msg)
	}
	if isIAMAuditConfigServiceModified(oldSpec, newSpec) {
		msg := fmt.Sprintf("the IAMAuditConfig's spec.service is immutable")
		return admission.ValidationResponse(false, msg)
	}
	return allowedResponse
}

func findChangesOnImmutableResourceIDField(spec, oldSpec map[string]interface{}, rc *corekccv1alpha1.ResourceConfig) bool {
	if rc.ResourceID.TargetField == "" {
		return false
	}

	return isResourceIDModified(spec, oldSpec)
}

func isResourceIDModified(spec, oldSpec map[string]interface{}) bool {
	return !reflect.DeepEqual(spec[k8s.ResourceIDFieldName], oldSpec[k8s.ResourceIDFieldName])
}

func findChangesOnImmutableLocationField(obj map[string]interface{}, oldObj map[string]interface{}, rc *corekccv1alpha1.ResourceConfig) bool {
	if rc.Locationality == "" {
		return false
	}
	// Location is immutable by default as it's part of the URL in underlying api.
	return !reflect.DeepEqual(obj["location"], oldObj["location"])
}

// TODO: get rid of list.List by changing the function to return a []string recursively
func compareAndFindChangesOnImmutableFields(obj map[string]interface{}, oldObj map[string]interface{}, schemaMap map[string]*tfschema.Schema, prefix string, resourceConfig *corekccv1alpha1.ResourceConfig, ignoredFields map[string]bool, fields *list.List) {
	for k, s := range schemaMap {
		qualifiedName := getQualifiedFieldName(prefix, k)
		if ignoredFields[qualifiedName] {
			continue
		}

		if ok, refConfig := krmtotf.IsReferenceField(qualifiedName, resourceConfig); ok {
			if !s.ForceNew {
				continue
			}
			modified, refKey := isReferenceValRawModified(obj, oldObj, refConfig)
			if modified {
				refKey = getQualifiedFieldName(prefix, refKey)
				fields.PushBack(refKey)
			}
			continue
		}

		camelCaseKey := text.SnakeCaseToLowerCamelCase(k)
		v1 := obj[camelCaseKey]
		v2 := oldObj[camelCaseKey]
		if v1 == nil && v2 == nil {
			continue
		}
		if (v1 == nil || v2 == nil) && s.ForceNew {
			fields.PushBack(qualifiedName)
			continue
		}

		switch s.Type {
		// TODO: terraform schema doc says that TypeMap only support Elem to be a *Schema with a Type that is one of the primitives
		// Is there any edge cases to handle?
		case tfschema.TypeBool, tfschema.TypeFloat, tfschema.TypeString, tfschema.TypeInt, tfschema.TypeMap:
			if s.ForceNew && !reflect.DeepEqual(v1, v2) {
				fields.PushBack(qualifiedName)
			}
		case tfschema.TypeList, tfschema.TypeSet:
			switch s.Elem.(type) {
			case *tfschema.Schema:
				// it's a list of primitives
				if s.ForceNew && !reflect.DeepEqual(v1, v2) {
					fields.PushBack(qualifiedName)
				}
			case *tfschema.Resource:
				if s.MaxItems == 1 {
					// A list with MaxItems == 1 is actually a nested object due to limitations with TF schemas.
					tfObjSchemaMap := s.Elem.(*tfschema.Resource).Schema
					var o1 map[string]interface{}
					var o2 map[string]interface{}
					if v1 != nil {
						o1 = v1.(map[string]interface{})
					}
					if v2 != nil {
						o2 = v2.(map[string]interface{})
					}
					compareAndFindChangesOnImmutableFields(o1, o2, tfObjSchemaMap, qualifiedName, resourceConfig, ignoredFields, fields)
				} else { //nolint:revive
					// TODO(kcc-eng): Kubernetes considers all lists of objects to be atomic, and so all subsequent
					//  applies will currently wipe out defaulted immutable fields. Temporarily delegate validation
					//  to the controller, which will determine via comparing the config with calculated fields in
					//  the live state to detect if a diff in immutable fields is present.
				}
			}
		}
	}
}

func isReferenceValRawModified(obj map[string]interface{}, oldObj map[string]interface{}, refConfig *corekccv1alpha1.ReferenceConfig) (bool, string) {
	// currently we choose to treat switching between different reference approaches as modification, e.g. change from referencing to ComputeAddress to directly specifying ip address
	referenceFieldKey := krmtotf.GetKeyForReferenceField(refConfig)
	return !reflect.DeepEqual(obj[referenceFieldKey], oldObj[referenceFieldKey]), referenceFieldKey
}

func constructCamelCasePath(path string) string {
	segs := make([]string, 0)
	for _, f := range strings.Split(path, ".") {
		segs = append(segs, text.SnakeCaseToLowerCamelCase(f))
	}
	return strings.Join(segs, ".")
}

func getIgnoredFields(rc *corekccv1alpha1.ResourceConfig) map[string]bool {
	ignoredFields := make(map[string]bool)
	for _, f := range rc.IgnoredFields {
		ignoredFields[f] = true
	}

	// metadata mapping can be ignored because there are only two k8s metadata fields mapping to TF fields: name and label
	// k8s object names are unique identifiers and labels are totally mutable by default
	ignoredFields[rc.MetadataMapping.Name] = true
	ignoredFields[rc.MetadataMapping.Labels] = true
	return ignoredFields
}
