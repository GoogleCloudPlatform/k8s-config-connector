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
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	admissionregistration "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var (
	admissionReviewVersions = []string{"v1", "v1beta1"}
)

func GenerateWebhookManifests(validatingWebhookConfigurationName, mutatingWebhookConfigurationName,
	serviceName string, whCfgs []Config) (*admissionregistration.ValidatingWebhookConfiguration, *admissionregistration.MutatingWebhookConfiguration) {
	validating := validatingWebhookConfig(validatingWebhookConfigurationName, serviceName, whCfgs)
	mutating := mutatingWebhookConfig(mutatingWebhookConfigurationName, serviceName, whCfgs)
	return validating, mutating
}

func generateService(name string, selector map[string]string) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: k8s.SystemNamespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: selector,
			Ports: []corev1.ServicePort{
				{
					Port:       ServicePort,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: ServicePort},
				},
			},
		},
	}
}

func mutatingWebhooksForWebhookConfigs(whCfgs []Config, svcName string, whType webhookType) []admissionregistration.MutatingWebhook {
	whs := make([]admissionregistration.MutatingWebhook, 0)
	for _, whCfg := range whCfgs {
		if whCfg.Type != whType {
			continue
		}
		wh := &admissionregistration.MutatingWebhook{
			Name:                    whCfg.Name,
			Rules:                   whCfg.Rules,
			ObjectSelector:          whCfg.ObjectSelector,
			FailurePolicy:           &whCfg.FailurePolicy,
			SideEffects:             &whCfg.SideEffects,
			AdmissionReviewVersions: admissionReviewVersions,
			// This field is removed due to go/kcccl/65384.
			// It is currently not necessary to set this field because 10s is the default in Kubernetes.
			// TimeoutSeconds:          &k8s.WebhookTimeoutSeconds,
		}
		cc := getClientConfig(svcName, whCfg.Path)
		wh.ClientConfig = *cc
		whs = append(whs, *wh)
	}
	sort.Slice(whs, func(i, j int) bool {
		return whs[i].Name < whs[j].Name
	})
	return whs
}

func validatingWebhooksForWebhookConfigs(whCfgs []Config, svcName string, whType webhookType) []admissionregistration.ValidatingWebhook {
	whs := make([]admissionregistration.ValidatingWebhook, 0)
	for _, whCfg := range whCfgs {
		if whCfg.Type != whType {
			continue
		}
		wh := &admissionregistration.ValidatingWebhook{
			Name:                    whCfg.Name,
			Rules:                   whCfg.Rules,
			ObjectSelector:          whCfg.ObjectSelector,
			FailurePolicy:           &whCfg.FailurePolicy,
			SideEffects:             &whCfg.SideEffects,
			AdmissionReviewVersions: admissionReviewVersions,
			// This field is removed due to go/kcccl/65384.
			// It is currently not necessary to set this field because 10s is the default in Kubernetes.
			// TimeoutSeconds:          &k8s.WebhookTimeoutSeconds,
		}
		cc := getClientConfig(svcName, whCfg.Path)
		wh.ClientConfig = *cc
		whs = append(whs, *wh)
	}
	sort.Slice(whs, func(i, j int) bool {
		return whs[i].Name < whs[j].Name
	})
	return whs
}

func mutatingWebhookConfig(name, svcName string, whCfgs []Config) *admissionregistration.MutatingWebhookConfiguration {
	whs := mutatingWebhooksForWebhookConfigs(whCfgs, svcName, Mutating)
	if len(whs) > 0 {
		return &admissionregistration.MutatingWebhookConfiguration{
			TypeMeta: metav1.TypeMeta{
				APIVersion: admissionregistration.SchemeGroupVersion.String(),
				Kind:       "MutatingWebhookConfiguration",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
				Labels: map[string]string{
					k8s.KCCSystemLabel: "true",
				},
			},
			Webhooks: whs,
		}
	}
	return nil
}

func validatingWebhookConfig(name, svcName string, whCfgs []Config) *admissionregistration.ValidatingWebhookConfiguration {
	whs := validatingWebhooksForWebhookConfigs(whCfgs, svcName, Validating)
	if len(whs) > 0 {
		return &admissionregistration.ValidatingWebhookConfiguration{
			TypeMeta: metav1.TypeMeta{
				APIVersion: admissionregistration.SchemeGroupVersion.String(),
				Kind:       "ValidatingWebhookConfiguration",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
				Labels: map[string]string{
					k8s.KCCSystemLabel: "true",
				},
			},
			Webhooks: whs,
		}
	}
	return nil
}

func getClientConfig(serviceName, path string) *admissionregistration.WebhookClientConfig {
	return &admissionregistration.WebhookClientConfig{
		CABundle: []byte{},
		Service: &admissionregistration.ServiceReference{
			Name:      serviceName,
			Namespace: k8s.SystemNamespace,
			Path:      &path,
		},
	}
}
