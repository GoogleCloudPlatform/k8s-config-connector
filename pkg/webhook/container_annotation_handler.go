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

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclcontainer "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension/container"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type containerAnnotationHandler struct {
	client                client.Client
	dclSchemaLoader       dclschemaloader.DCLSchemaLoader
	serviceMetadataLoader dclmetadata.ServiceMetadataLoader
	smLoader              *servicemappingloader.ServiceMappingLoader
}

func NewContainerAnnotationHandler(smLoader *servicemappingloader.ServiceMappingLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &containerAnnotationHandler{
			client:                mgr.GetClient(),
			smLoader:              smLoader,
			serviceMetadataLoader: serviceMetadataLoader,
			dclSchemaLoader:       dclSchemaLoader,
		}
	}
}

func (a *containerAnnotationHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
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
	if dclmetadata.IsDCLBasedResourceKind(obj.GroupVersionKind(), a.serviceMetadataLoader) {
		return handleContainerAnnotationsForDCLBasedResources(obj, ns, a.dclSchemaLoader, a.serviceMetadataLoader)
	}
	return handleContainerAnnotationsForTFBasedResources(obj, ns, a.smLoader)
}

func handleContainerAnnotationsForDCLBasedResources(obj *unstructured.Unstructured, ns *corev1.Namespace, dclSchemaLoader dclschemaloader.DCLSchemaLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) admission.Response {
	gvk := obj.GroupVersionKind()
	r, found := serviceMetadataLoader.GetResourceWithGVK(gvk)
	if !found {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("ServiceMetadata for resource with GroupVersionKind %v not found", gvk))
	}
	containers, err := dclcontainer.GetContainersForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting containers supported by GroupVersionKind %v: %w", gvk, err))
	}

	// TODO(b/186159460): Delete this if-block once all resources support
	// hierarchical references.
	if !r.SupportsHierarchicalReferences {
		return setDefaultContainerAnnotation(obj, ns, containers)
	}

	hierarchicalRefs, err := dcl.GetHierarchicalReferencesForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting hierarchical references supported by GroupVersionKind %v: %w", gvk, err))
	}
	return setDefaultHierarchicalReference(obj, ns, hierarchicalRefs, containers)
}

func handleContainerAnnotationsForTFBasedResources(obj *unstructured.Unstructured, ns *corev1.Namespace, smLoader *servicemappingloader.ServiceMappingLoader) admission.Response {
	rc, err := smLoader.GetResourceConfig(obj)
	if err != nil {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error getting ResourceConfig for kind %v: %w", obj.GetKind(), err))
	}

	// TODO(b/193177782): Delete this if-block once all resources support
	// hierarchical references.
	if !krmtotf.SupportsHierarchicalReferences(rc) {
		return setDefaultContainerAnnotation(obj, ns, rc.Containers)
	}
	return setDefaultHierarchicalReference(obj, ns, rc.HierarchicalReferences, rc.Containers)
}

func setDefaultContainerAnnotation(obj *unstructured.Unstructured, ns *corev1.Namespace, containers []corekccv1alpha1.Container) admission.Response {
	newObj := obj.DeepCopy()
	if err := k8s.SetDefaultContainerAnnotation(newObj, ns, containers); err != nil {
		return admission.Errored(http.StatusBadRequest, fmt.Errorf("error setting container annotation: %w", err))
	}
	return constructPatchResponse(obj, newObj)
}

func setDefaultHierarchicalReference(obj *unstructured.Unstructured, ns *corev1.Namespace, hierarchicalRefs []corekccv1alpha1.HierarchicalReference, containers []corekccv1alpha1.Container) admission.Response {
	resource, err := k8s.NewResource(obj)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, fmt.Errorf("error converting object to k8s resource: %w", err))
	}
	if err := k8s.SetDefaultHierarchicalReference(resource, ns, hierarchicalRefs, containers); err != nil {
		return admission.Errored(http.StatusInternalServerError, fmt.Errorf("error setting hierarchical reference: %w", err))
	}
	newObj, err := resource.MarshalAsUnstructured()
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, fmt.Errorf("error marshalling k8s resource to unstructured: %w", err))
	}
	return constructPatchResponse(obj, newObj)
}

func constructPatchResponse(obj, newObj *unstructured.Unstructured) admission.Response {
	objRaw, err := obj.MarshalJSON()
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error marshaling object as JSON: %w", err))
	}
	newObjRaw, err := newObj.MarshalJSON()
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error marshaling new object as JSON: %w", err))
	}
	return admission.PatchResponseFromRaw(objRaw, newObjRaw)
}
