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
	admissionregistration "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Config struct {
	Type           webhookType
	Name           string
	Path           string
	HandlerFunc    func(mgr manager.Manager) admission.Handler
	FailurePolicy  admissionregistration.FailurePolicyType
	ObjectSelector *metav1.LabelSelector
	Rules          []admissionregistration.RuleWithOperations
	SideEffects    admissionregistration.SideEffectClass
}

type webhookType string

const (
	Mutating   webhookType = "Mutating"
	Validating webhookType = "Validating"
)

func (c *Config) BuildAdmission(mgr manager.Manager) *webhook.Admission {
	handler := c.HandlerFunc(mgr)
	return &webhook.Admission{Handler: handler}
}
