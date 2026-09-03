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

package webhook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type actuationModeAnnotationValidator struct {
	client client.Client
}

// NewActuationModeAnnotationValidatorHandler creates an instance of
// actuationModeAnnotationValidator to handle actuation-mode annotation
// validation and warnings.
func NewActuationModeAnnotationValidatorHandler() HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &actuationModeAnnotationValidator{client: mgr.GetClient()}
	}
}

func (a *actuationModeAnnotationValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}

	var raw []byte
	if req.Operation == admissionv1.Delete {
		raw = req.AdmissionRequest.OldObject.Raw
	} else {
		raw = req.AdmissionRequest.Object.Raw
	}

	if len(raw) == 0 {
		return allowedResponse
	}

	if _, _, err := deserializer.Decode(raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding object: %w", err))
	}

	value, ok := k8s.GetAnnotation(k8s.ActuationModeAnnotation, obj)
	if ok && value == "Paused" {
		name := obj.GetName()
		if name == "" {
			name = req.Name
		}
		warning := fmt.Sprintf("Resource '%s' has %s: \"Paused\". All actuation against GCP (including Create, Update, and Delete) is halted until unpaused.", name, k8s.ActuationModeAnnotation)
		return allowedResponse.WithWarnings(warning)
	}

	return allowedResponse
}
