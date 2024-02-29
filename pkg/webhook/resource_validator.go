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
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type resourceValidatorHandler struct {
}

func NewResourceValidatorHandler() HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &resourceValidatorHandler{}
	}
}

func (a *resourceValidatorHandler) Handle(_ context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err := resourceoverrides.Handler.ConfigValidate(obj); err != nil {
		return admission.Errored(http.StatusForbidden, err)
	}
	return admission.ValidationResponse(true, "admission controller passed")
}
