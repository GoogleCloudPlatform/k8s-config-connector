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

	opv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type actuationModeAnnotationValidator struct{}

// NewActuationModeAnnotationValidatorHandler creates an instance of
// actuationModeAnnotationValidator to handle actuation-mode annotation
// validation and warnings.
func NewActuationModeAnnotationValidatorHandler() HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &actuationModeAnnotationValidator{}
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
	if ok {
		mode := opv1beta1.ActuationMode(value)
		if mode != opv1beta1.Paused && mode != opv1beta1.Reconciling {
			return admission.Denied(fmt.Sprintf("invalid value %q for annotation %s; allowed values are %q and %q",
				value, k8s.ActuationModeAnnotation, opv1beta1.Paused, opv1beta1.Reconciling))
		}
		if mode == opv1beta1.Paused {
			name := obj.GetName()
			if name == "" {
				name = req.Name
			}
			warning := fmt.Sprintf("Resource '%s' has %s: %q. Actuation against GCP (Create and Update) is halted until unpaused. Deletion will reconcile normally.",
				name, k8s.ActuationModeAnnotation, opv1beta1.Paused)
			return allowedResponse.WithWarnings(warning)
		}
	}

	return allowedResponse
}
