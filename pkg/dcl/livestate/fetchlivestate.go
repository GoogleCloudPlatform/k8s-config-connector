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

package livestate

import (
	"context"
	"fmt"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/kcclite"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FetchLiveState returns the live state of the underlying resource.

// It invokes DCL Get() to read on the underlying resource and then convert response to KCC lite format.
func FetchLiveState(ctx context.Context, resource *dcl.Resource, dclConfig *mmdcl.Config, converter *conversion.Converter, serviceMappingLoader *servicemappingloader.ServiceMappingLoader, kubeClient client.Client) (*unstructured.Unstructured, error) {
	serverGeneratedIDNotConfigured, err := resource.HasServerGeneratedIDButNotConfigured()
	if err != nil {
		return nil, err
	}
	if serverGeneratedIDNotConfigured {
		// If the resource ID cannot be determined because it requires a server-
		// generated ID that has not been set, this means the resource has not
		// yet been created. Return as if the read returned a nonexistent
		// resource.
		return nil, nil
	}

	var lite *unstructured.Unstructured
	secretVersions := make(map[string]string)
	if k8s.IsDeleted(&resource.ObjectMeta) {
		lite, err = kcclite.ToKCCLiteBestEffort(resource, converter.MetadataLoader, converter.SchemaLoader, serviceMappingLoader, kubeClient)
	} else {
		lite, secretVersions, err = kcclite.ToKCCLiteAndSecretVersions(resource, converter.MetadataLoader, converter.SchemaLoader, serviceMappingLoader, kubeClient)
	}
	if err != nil {
		return nil, fmt.Errorf("error converting to lite state: %w", err)
	}

	dclResource, err := converter.KRMObjectToDCLObject(lite)
	if err != nil {
		return nil, fmt.Errorf("error converting to DCL resource object: %w", err)
	}

	liveState, err := dclunstruct.Get(ctx, dclConfig, dclResource)
	if err != nil {
		if gcp.IsNotFoundError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("error reading underlying resource: %w", err)
	}
	liveLite, err := converter.DCLObjectToKRMObject(liveState)
	if err != nil {
		return nil, fmt.Errorf("error converting DCL object to KRM resource obj: %w", err)
	}

	liveLite, err = withMutableButUnreadableFields(liveLite, resource, secretVersions, kubeClient)
	if err != nil {
		return nil, fmt.Errorf("error setting mutable-but-unreadable fields for live state: %w", err)
	}

	return liveLite, nil
}

func SetMutableButUnreadableFields(kccLite *unstructured.Unstructured, mutableButUnreadableSpec map[string]interface{}, path []string, schema *openapi.Schema, secretVersions map[string]string, namespace string, kubeClient client.Client) (*unstructured.Unstructured, error) {
	if len(mutableButUnreadableSpec) == 0 {
		return kccLite, nil
	}

	if schema.Type != "object" {
		return nil, fmt.Errorf("wrong type for provided schema: %s, expect to have object", schema.Type)
	}

	for k, v := range mutableButUnreadableSpec {
		path := append(path, k)

		subSchema, ok := schema.Properties[k]
		if !ok {
			return nil, fmt.Errorf("unknown mutable-but-unreadable path '%v'", pathslice.ToString(path))
		}
		switch subSchema.Type {
		case "integer":
			v, err := dcl.CanonicalizeIntegerValue(v)
			if err != nil {
				return nil, fmt.Errorf("error canonicalizing the integer value for path '%v': %w", pathslice.ToString(path), err)
			}

			if err := unstructured.SetNestedField(kccLite.Object, v, path...); err != nil {
				return nil, fmt.Errorf("error setting mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
			}
		case "number":
			v, err := dcl.CanonicalizeNumberValue(v)
			if err != nil {
				return nil, fmt.Errorf("error canonicalizing the number value for path '%v': %w", pathslice.ToString(path), err)
			}

			if err := unstructured.SetNestedField(kccLite.Object, v, path...); err != nil {
				return nil, fmt.Errorf("error setting mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
			}

		case "string":
			isSensitive, err := extension.IsSensitiveField(subSchema)
			if err != nil {
				return nil, fmt.Errorf("error checking sensitivity for mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
			}

			if !isSensitive {
				if err := unstructured.SetNestedField(kccLite.Object, v, path...); err != nil {
					return nil, fmt.Errorf("error setting mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
				}
				continue
			}

			val, err := resolveSensitiveValueForLiveState(v, secretVersions, namespace, kubeClient)
			if err != nil {
				return nil, fmt.Errorf("error parsing the secret for mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
			}
			valMap := make(map[string]interface{})
			valMap["value"] = val
			if err := unstructured.SetNestedField(kccLite.Object, valMap, path...); err != nil {
				return nil, fmt.Errorf("error setting mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
			}
		case "boolean":
			if err := unstructured.SetNestedField(kccLite.Object, v, path...); err != nil {
				return nil, fmt.Errorf("error setting mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
			}
		case "array":
			v, ok := v.([]interface{})
			if !ok {
				return nil, fmt.Errorf("wrong type for path '%v': %T, expect to have []interface{}", pathslice.ToString(path), v)
			}

			itemSchema := subSchema.Items
			switch itemSchema.Type {
			case "string", "boolean", "number", "integer":
				// List/set of primitives
				path := strings.Split(pathslice.ToString(path), ".")
				if err := unstructured.SetNestedField(kccLite.Object, deepcopy.DeepCopy(v), path...); err != nil {
					return nil, fmt.Errorf("error setting mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
				}
			case "array", "object":
				// List/set of non-primitives
				return nil, fmt.Errorf("error setting mutable-but-unreadable path '%v': mutable-but-unreadable list/set of non-primitive types is not yet supported", pathslice.ToString(path))
			default:
				// List of unknown types
				return nil, fmt.Errorf("unknown list/set type %T for mutable-but-unreadable path '%v'", itemSchema, pathslice.ToString(path))
			}
		case "object":
			v, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("wrong type for path '%v': %T, expect to have map[string]interface{}", pathslice.ToString(path), v)
			}

			if subSchema.AdditionalProperties != nil {
				// Map
				// Currently KCC doesn't support the nested path in the value
				// of the map to be mutable-but-unreadable, so the map is copied
				// as a whole.
				if err := unstructured.SetNestedField(kccLite.Object, deepcopy.DeepCopy(v), path...); err != nil {
					return nil, fmt.Errorf("error setting mutable-but-unreadable path '%v': %w", pathslice.ToString(path), err)
				}
			} else {
				// Object
				var err error
				kccLite, err = SetMutableButUnreadableFields(kccLite, v, path, subSchema, secretVersions, namespace, kubeClient)
				if err != nil {
					return nil, fmt.Errorf("error setting nested mutable-but-unreadable path: %w", err)
				}
			}
		default:
			// Unknown types
			return nil, fmt.Errorf("unknown type %T for mutable-but-unreadable path '%v'", subSchema, pathslice.ToString(path))
		}
	}

	return kccLite, nil
}

func resolveSensitiveValueForLiveState(value interface{}, secretVersions map[string]string, namespace string, kubeClient client.Client) (string, error) {
	sensitiveField := corekccv1alpha1.SensitiveField{}
	if err := util.Marshal(value, &sensitiveField); err != nil {
		return "", fmt.Errorf("error parsing %v onto a SensitiveField struct: %w", value, err)
	}

	if sensitiveField.Value != nil {
		return *sensitiveField.Value, nil
	}

	secretKeyRef := sensitiveField.ValueFrom.SecretKeyRef
	secretVal, secretVer, err := k8s.GetSecretVal(secretKeyRef, namespace, kubeClient)
	if err != nil {
		// If a previously referenced Secret cannot be resolved, it is possible
		// it simply no longer exists. Don't error out in this case; just skip
		// the presetting of the path in the live state so that a diff is
		// generated if the path had been updated in the spec. If the path
		// still points to the same Secret in the spec, then the DCL KCClite
		// conversion will appropriately error out due to the nonexistent
		// Secret.
		return "", nil
	}
	// Preset sensitive path only if we can be sure that the
	// referenced Secret had not been changed.
	prevSecretVer, ok := secretVersions[secretKeyRef.Name]
	if ok && secretVer == prevSecretVer {
		return secretVal, nil
	}

	// If the referenced Secret couldn't be found in the secretVersions map, or
	// the referenced Secret has a different Secret version, then we need to
	// explicitly set the secret value to an empty string ("") to trigger the diff.
	return "", nil
}

func withMutableButUnreadableFields(kccLite *unstructured.Unstructured, resource *dcl.Resource, currSecretVersions map[string]string, kubeClient client.Client) (*unstructured.Unstructured, error) {
	hasMutableButUnreadableFields, err := resource.HasMutableButUnreadableFields()
	if err != nil {
		return nil, fmt.Errorf("error checking if resource has mutable-but-unreadable fields: %w", err)
	}
	if !hasMutableButUnreadableFields {
		return kccLite, nil
	}

	lastSeenValues, err := dcl.GetMutableButUnreadableFieldsFromAnnotations(resource)
	if err != nil {
		return nil, fmt.Errorf("error getting last-seen values of mutable-but-unreadable fields from annotations: %w", err)
	}

	updatedLite := kccLite.DeepCopy()
	if len(lastSeenValues) == 0 { // if len(lastSeenValues) == 0, the mutable-but-unreadable field is probably optional and unset
		return updatedLite, nil
	}

	secretVersions, err := k8s.GetSecretVersionsFromAnnotations(&resource.Resource)
	if err != nil {
		return nil, fmt.Errorf("error getting secret versions from annotations: %w", err)
	}
	// When secretVersions are not found in the annotations, there is a possibility
	// that this is either (1) a resource acquisition or (2) a resource created
	// before it supported the annotation. To avoid unnecessarily updating the
	// resource in both cases, use the current Secret versions.
	if secretVersions == nil {
		secretVersions = currSecretVersions
	}

	updatedLite, err = SetMutableButUnreadableFields(updatedLite, lastSeenValues["spec"].(map[string]interface{}), []string{"spec"}, resource.Schema, secretVersions, resource.GetNamespace(), kubeClient)
	if err != nil {
		return nil, err
	}
	return updatedLite, nil
}
