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

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/managementconflict"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type managementConflictAnnotationDefaulter struct {
	client                client.Client
	dclSchemaLoader       dclschemaloader.DCLSchemaLoader
	serviceMetadataLoader dclmetadata.ServiceMetadataLoader
	smLoader              *servicemappingloader.ServiceMappingLoader
	tfResourceMap         map[string]*tfschema.Resource
}

func NewManagementConflictAnnotationDefaulter(smLoader *servicemappingloader.ServiceMappingLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &managementConflictAnnotationDefaulter{
			client:                mgr.GetClient(),
			smLoader:              smLoader,
			serviceMetadataLoader: serviceMetadataLoader,
			dclSchemaLoader:       dclSchemaLoader,
			tfResourceMap:         provider.ResourceMap(),
		}
	}
}

func (a *managementConflictAnnotationDefaulter) Handle(ctx context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding object: %w", err))
	}
	ns := &corev1.Namespace{}
	if err := a.client.Get(ctx, apimachinerytypes.NamespacedName{Name: obj.GetNamespace()}, ns); err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting Namespace %v: %w", obj.GetNamespace(), err))
	}
	if supportedgvks.IsDirectByGVK(obj.GroupVersionKind()) {
		return admission.Allowed("")
	}
	if dclmetadata.IsDCLBasedResourceKind(obj.GroupVersionKind(), a.serviceMetadataLoader) {
		return defaultManagementConflictAnnotationForDCLBasedResources(obj, ns, a.dclSchemaLoader, a.serviceMetadataLoader)
	}
	return defaultManagementConflictAnnotationForTFBasedResources(obj, ns, a.smLoader, a.tfResourceMap)
}

func defaultManagementConflictAnnotationForDCLBasedResources(obj *unstructured.Unstructured, ns *corev1.Namespace, dclSchemaLoader dclschemaloader.DCLSchemaLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) admission.Response {
	gvk := obj.GroupVersionKind()
	stv, err := dclmetadata.ToServiceTypeVersion(gvk, serviceMetadataLoader)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting DCL ServiceTypeVersion for GroupVersionKind %v: %w", gvk, err))
	}
	schema, err := dclSchemaLoader.GetDCLSchema(stv)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting the DCL Schema for GroupVersionKind %v: %w", gvk, err))
	}
	newObj := obj.DeepCopy()
	if err := managementconflict.ValidateOrDefaultManagementConflictPreventionAnnotationForDCLBasedResource(newObj, ns, schema); err != nil {
		return admission.Errored(http.StatusBadRequest, fmt.Errorf("error validating or defaulting management conflict policy annotation: %w", err))
	}
	return constructPatchResponse(obj, newObj)
}

func defaultManagementConflictAnnotationForTFBasedResources(obj *unstructured.Unstructured, ns *corev1.Namespace, smLoader *servicemappingloader.ServiceMappingLoader, tfResourceMap map[string]*tfschema.Resource) admission.Response {
	rc, err := smLoader.GetResourceConfig(obj)
	if err != nil {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error getting ResourceConfig for kind %v: %w", obj.GetKind(), err))
	}
	newObj := obj.DeepCopy()
	if err := managementconflict.ValidateOrDefaultManagementConflictPreventionAnnotationForTFBasedResource(newObj, ns, rc, tfResourceMap); err != nil {
		return admission.Errored(http.StatusBadRequest, fmt.Errorf("error validating or defaulting management conflict policy annotation: %w", err))
	}
	return constructPatchResponse(obj, newObj)
}
