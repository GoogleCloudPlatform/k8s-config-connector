// Copyright 2026 Google LLC
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

package cleanupemptyobjects

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/crd-tools/pkg/objectvisitor"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func AddCommand(parent *cobra.Command) {
	var opt Options
	cmd := &cobra.Command{
		Use:   "cleanup-empty-objects",
		Short: "Cleanup empty object types in CRDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd.Context(), opt)
		},
	}
	cmd.Flags().StringVar(&opt.Dir, "dir", "", "Directory to process")
	parent.AddCommand(cmd)
}

type Options struct {
	Dir string
}

func (o *Options) Validate() error {
	if o.Dir == "" {
		return fmt.Errorf("--dir is required")
	}
	return nil
}

type visitor struct{}

func (v *visitor) VisitObject(obj *unstructured.Unstructured) error {
	if obj.GetKind() != "CustomResourceDefinition" {
		return nil
	}

	spec, found, err := unstructured.NestedMap(obj.Object, "spec")
	if err != nil || !found {
		return nil
	}

	versions, found, err := unstructured.NestedSlice(spec, "versions")
	if err != nil || !found {
		return nil
	}

	for i, vRaw := range versions {
		version, ok := vRaw.(map[string]interface{})
		if !ok {
			continue
		}

		schema, found, err := unstructured.NestedMap(version, "schema", "openAPIV3Schema")
		if err != nil || !found {
			continue
		}

		v.cleanupSchema(schema)

		if err := unstructured.SetNestedMap(version, schema, "schema", "openAPIV3Schema"); err != nil {
			return err
		}
		versions[i] = version
	}

	if err := unstructured.SetNestedSlice(obj.Object, versions, "spec", "versions"); err != nil {
		return err
	}

	return nil
}

func (v *visitor) cleanupSchema(schema map[string]interface{}) {
	properties, found, err := unstructured.NestedMap(schema, "properties")
	if err == nil && found {
		for name, propRaw := range properties {
			prop, ok := propRaw.(map[string]interface{})
			if !ok {
				continue
			}
			if name == "observedState" && isTypeObject(prop) && isEmptyObject(prop) {
				delete(properties, name)
				continue
			}
			v.cleanupSchema(prop)
			properties[name] = prop
		}
		if err := unstructured.SetNestedMap(schema, properties, "properties"); err != nil {
			panic(err)
		}
	}

	items, found, err := unstructured.NestedMap(schema, "items")
	if err == nil && found {
		v.cleanupSchema(items)
		if err := unstructured.SetNestedMap(schema, items, "items"); err != nil {
			panic(err)
		}
	}

	for _, key := range []string{"allOf", "anyOf", "oneOf"} {
		list, found, err := unstructured.NestedSlice(schema, key)
		if err == nil && found {
			for i, itemRaw := range list {
				item, ok := itemRaw.(map[string]interface{})
				if ok {
					v.cleanupSchema(item)
					list[i] = item
				}
			}
			if err := unstructured.SetNestedSlice(schema, list, key); err != nil {
				panic(err)
			}
		}
	}

	if isTypeObject(schema) && isEmptyObject(schema) {
		schema["additionalProperties"] = false
	}
}

func isTypeObject(schema map[string]interface{}) bool {
	t, found, _ := unstructured.NestedString(schema, "type")
	return found && t == "object"
}

func isEmptyObject(schema map[string]interface{}) bool {
	properties, foundProps, _ := unstructured.NestedMap(schema, "properties")
	additionalProperties, foundAddProps, _ := unstructured.NestedFieldNoCopy(schema, "additionalProperties")
	preserveUnknownFields, foundPreserve, _ := unstructured.NestedBool(schema, "x-kubernetes-preserve-unknown-fields")

	return (!foundProps || len(properties) == 0) &&
		(!foundAddProps || additionalProperties == nil) &&
		(!foundPreserve || !preserveUnknownFields)
}

func Run(ctx context.Context, options Options) error {
	if err := options.Validate(); err != nil {
		return err
	}

	visitor := &visitor{}
	return objectvisitor.VisitObjectsInDirectory(ctx, options.Dir, visitor)
}
