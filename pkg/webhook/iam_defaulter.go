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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	iamResourceRefPath  = []string{"spec", "resourceRef"}
	iamResourceRefField = strings.Join(iamResourceRefPath, ".")
)

type iamDefaulter struct {
	smLoader              *servicemappingloader.ServiceMappingLoader
	serviceMetadataLoader metadata.ServiceMetadataLoader
}

func NewIAMDefaulter(smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetadataLoader metadata.ServiceMetadataLoader) HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &iamDefaulter{
			smLoader:              smLoader,
			serviceMetadataLoader: serviceMetadataLoader,
		}
	}
}

func (a *iamDefaulter) Handle(_ context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding object: %w", err))
	}

	if !isIAMResource(obj) {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("object of GroupVersionKind %v is not a supported IAM resource", obj.GroupVersionKind()))
	}
	return defaultAPIVersionForIAMResourceRef(obj, a.smLoader, a.serviceMetadataLoader)
}

func defaultAPIVersionForIAMResourceRef(obj *unstructured.Unstructured,
	smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetadataLoader metadata.ServiceMetadataLoader) admission.Response {
	resourceRef, found, err := unstructured.NestedMap(obj.Object, iamResourceRefPath...)
	if err != nil {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error getting '%v': %w", iamResourceRefField, err))
	}
	if !found {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("couldn't find '%v'", iamResourceRefField))
	}

	// No need to default spec.resourceRef.apiVersion if already set.
	if _, ok := resourceRef["apiVersion"]; ok {
		return constructPatchResponse(obj, obj)
	}

	if _, ok := resourceRef["kind"]; !ok {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("couldn't find %v.kind'", iamResourceRefField))
	}
	kind, ok := resourceRef["kind"].(string)
	if !ok || kind == "" {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("'%v.kind' must be set to a non-empty string", iamResourceRefField))
	}

	apiVersion, err := apiVersionForKind(kind, smLoader, serviceMetadataLoader)
	if err != nil {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error determining apiVersion for kind '%v': %w", kind, err))
	}

	resourceRef["apiVersion"] = apiVersion
	newObj := obj.DeepCopy()
	if err := unstructured.SetNestedMap(newObj.Object, resourceRef, iamResourceRefPath...); err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error setting '%v': %w", iamResourceRefField, err))
	}
	return constructPatchResponse(obj, newObj)

}

func apiVersionForKind(kind string,
	smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetadataLoader metadata.ServiceMetadataLoader) (string, error) {
	gvk, ok, err := gvks.GVKForKind(kind, smLoader, serviceMetadataLoader)
	if err != nil {
		return "", fmt.Errorf("error finding a GroupVersionKind for kind '%v': %w", kind, err)
	}
	if !ok {
		return "", fmt.Errorf("couldn't find a GroupVersionKind for kind '%v'", kind)
	}
	return gvk.GroupVersion().String(), nil
}
