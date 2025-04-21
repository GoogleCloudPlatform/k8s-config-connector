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
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	apitypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type noUnknownFieldsValidatorHandler struct {
	client   client.Client
	smLoader *servicemappingloader.ServiceMappingLoader
}

func NewNoUnknownFieldsValidatorHandler(smLoader *servicemappingloader.ServiceMappingLoader) HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &noUnknownFieldsValidatorHandler{
			client:   mgr.GetClient(),
			smLoader: smLoader,
		}
	}
}

func (a *noUnknownFieldsValidatorHandler) Handle(_ context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest, err)
	}
	crd := &apiextensions.CustomResourceDefinition{}
	nn := apitypes.NamespacedName{
		Name: text.Pluralize(strings.ToLower(obj.GetKind())) + "." + obj.GroupVersionKind().Group,
	}
	if err := a.client.Get(context.Background(), nn, crd); err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	if err := validateNoUnknownFields(schema, obj.Object, ""); err != nil {
		return admission.Errored(http.StatusForbidden, err)
	}
	return admission.ValidationResponse(true, "admission controller passed")
}

func validateNoUnknownFields(schema *apiextensions.JSONSchemaProps, field interface{}, key string) error {
	switch schema.Type {
	case "object":
		m, ok := field.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unrecognized type for field %v; expected object", key)
		}
		if len(schema.Properties) == 0 {
			// If the schema is of type object but it has no properties, it should allow
			// any keys.
			return nil
		}
		for k, v := range m {
			subfieldKey := key + "." + k
			if key == "" {
				subfieldKey = k
			}
			subfieldSchema, ok := schema.Properties[k]
			if !ok {
				return fmt.Errorf("unknown field \"%v\"", subfieldKey)
			}
			if err := validateNoUnknownFields(&subfieldSchema, v, subfieldKey); err != nil {
				return err
			}
		}
	case "array":
		a, ok := field.([]interface{})
		if !ok {
			return fmt.Errorf("unrecognized type for field %v; expected array", key)
		}
		for i, v := range a {
			itemKey := fmt.Sprintf("%v[%v]", key, i)
			if err := validateNoUnknownFields(schema.Items.Schema, v, itemKey); err != nil {
				return err
			}
		}
	case "string", "number", "integer", "boolean", "null":
		return nil
	default:
		// If we are dealing with a schemaless field, we don't want to validate the type
		if schema.XPreserveUnknownFields != nil && *schema.XPreserveUnknownFields {
			return nil
		}
		return fmt.Errorf("unrecognized schema type: %v", schema.Type)
	}
	return nil
}
