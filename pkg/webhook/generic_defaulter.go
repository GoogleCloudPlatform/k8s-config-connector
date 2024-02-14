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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// A generic webhook defaulter that is meant to handle various miscellaneous defaulting.
// The reasons to have a generic defaulter are 1) to avoid the heavyweight boilerplate code
// of registering new webhook handlers and 2) to avoid adding a new entry in the MutatingWebhookConfiguration object with almost all KCC resources listed
// every time we introduce a new feature across the board.
//
// This defaulter should only contain lightweight defaulting logic, for complicated defaulting
// functions, it's recommended to create dedicated webhook handlers.
type genericDefaulter struct {
}

func NewGenericDefaulter() HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &genericDefaulter{}
	}
}

func (a *genericDefaulter) Handle(ctx context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding object: %w", err))
	}
	newObj := obj.DeepCopy()
	if err := k8s.ValidateOrDefaultStateIntoSpecAnnotation(newObj); err != nil {
		return admission.Errored(http.StatusBadRequest, fmt.Errorf("error validating or defaulting '%v' annotation: %w", k8s.StateIntoSpecAnnotation, err))
	}
	return constructPatchResponse(obj, newObj)
}
