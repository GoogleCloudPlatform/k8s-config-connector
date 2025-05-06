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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type stateIntoSpecAnnotationValidator struct {
	client client.Client
}

// NewStateIntoSpecAnnotationValidatorHandler creates an instance of
// stateIntoSpecAnnotationValidator to handle state-into-spec annotation
// validation.
func NewStateIntoSpecAnnotationValidatorHandler() HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &stateIntoSpecAnnotationValidator{client: mgr.GetClient()}
	}
}

func (a *stateIntoSpecAnnotationValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding object: %w", err))
	}

	value, ok := k8s.GetAnnotation(k8s.StateIntoSpecAnnotation, obj)
	if ok && value == apis.StateMergeIntoSpec {
		return allowedResponse.WithWarnings(
			fmt.Sprintf("'%v: %v' is unsupported for CRDs added in "+
				"1.114.0 and later. Use '%v' instead. More details can be "+
				"found at https://cloud.google.com/config-connector/docs/concepts/ignore-unspecified-fields.",
				k8s.StateIntoSpecAnnotation, apis.StateMergeIntoSpec, apis.StateAbsentInSpec))
	}
	// TODO: Verify if the state-into-spec annotation will be defaulted to 'merge'.
	return allowedResponse
}
