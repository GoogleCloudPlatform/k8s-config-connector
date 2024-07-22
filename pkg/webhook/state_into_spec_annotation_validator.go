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
	klog.Infof("maqiuyu...kind %q, obj %q, annotation value %v, ok? %v", obj.GroupVersionKind().Kind, obj.GetName(), value, ok)
	if ok && value == k8s.StateMergeIntoSpec {
		return allowedResponse.WithWarnings(
			fmt.Sprintf("'%v: %v' is unsupported for CRDs added in "+
				"1.114.0 and later. Use '%v' instead. More details can be "+
				"found at https://cloud.google.com/config-connector/docs/concepts/ignore-unspecified-fields.",
				k8s.StateIntoSpecAnnotation, k8s.StateMergeIntoSpec, k8s.StateAbsentInSpec))
	} else if !ok {
		stateIntoSpecDefaultVal, err := k8s.GetStateIntoSpecDefaultValue(ctx, a.client, obj)
		klog.Infof("maqiuyu...kind %q, obj %q, default value %v, err? %v", obj.GroupVersionKind().Kind, obj.GetName(), stateIntoSpecDefaultVal, err)
		if err != nil {
			// Something wrong happened while trying to fetch the default value
			// 'state-into-spec' annotation. This should not block the request.
			return allowedResponse
		}
		klog.Infof("maqiuyu...kind %q, obj %q, will start comparing the default value", obj.GroupVersionKind().Kind, obj.GetName())
		if stateIntoSpecDefaultVal == k8s.StateMergeIntoSpec {
			return allowedResponse.WithWarnings(
				fmt.Sprintf("'%v: %v' is unsupported for CRDs added in "+
					"1.114.0 and later. Use '%v' instead. More details can be "+
					"found at https://cloud.google.com/config-connector/docs/concepts/ignore-unspecified-fields.",
					k8s.StateIntoSpecAnnotation, k8s.StateMergeIntoSpec, k8s.StateAbsentInSpec))
		}
	}
	return allowedResponse
}
