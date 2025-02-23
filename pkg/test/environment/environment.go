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

package testenvironment

import (
	"fmt"
	"log"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"

	admissionregistration "k8s.io/api/admissionregistration/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func init() {
	s := scheme.Scheme
	if err := apiextensions.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering apiextensions v1beta1 scheme: %v", err)
	}
	if err := apis.AddToScheme(s); err != nil {
		log.Fatalf("error registering schemes: %v", err)
	}
	if err := operatorv1beta1.SchemeBuilder.AddToScheme(s); err != nil {
		log.Fatalf("error registering operator v1beta1 schemes: %v", err)
	}
}

func StartTestEnvironmentOrLogFatal(testType test.Type, crds []*apiextensions.CustomResourceDefinition, whCfgs []webhook.Config) *envtest.Environment {
	env, err := startTestEnvironment(testType, crds, whCfgs)
	if err != nil {
		log.Fatal(err)
	}
	return env
}

func startTestEnvironment(testType test.Type, crds []*apiextensions.CustomResourceDefinition, whCfgs []webhook.Config) (*envtest.Environment, error) {
	env := &envtest.Environment{
		ControlPlaneStartTimeout: time.Minute,
		ControlPlaneStopTimeout:  time.Minute,
	}
	switch {
	case len(crds) > 0:
		env.CRDs = crds
	case testType == test.UnitTestType:
		env.CRDs = test.FakeCRDs()
	case testType == test.IntegrationTestType:
		env.CRDDirectoryPaths = []string{repo.GetCRDsPath()}
		env.CRDDirectoryPaths = append(env.CRDDirectoryPaths, paths.GetOperatorCRDsPaths()...)
	}
	if testType == test.IntegrationTestType {
		ConfigureWebhookInstallOptions(env, whCfgs)
	}
	_, err := env.Start()
	if err != nil {
		return nil, fmt.Errorf("error starting test environment: %w", err)
	}
	return env, nil
}

func ConfigureWebhookInstallOptions(env *envtest.Environment, whCfgs []webhook.Config) {
	validatingWebhookCfg, mutatingWebhookCfg := webhook.GenerateWebhookManifests(
		webhook.ValidatingWebhookConfigurationName,
		webhook.MutatingWebhookConfigurationName,
		webhook.CommonWebhookServiceName,
		whCfgs,
	)
	env.WebhookInstallOptions = envtest.WebhookInstallOptions{}
	if validatingWebhookCfg != nil {
		env.WebhookInstallOptions.ValidatingWebhooks = []*admissionregistration.ValidatingWebhookConfiguration{validatingWebhookCfg}
	}
	if mutatingWebhookCfg != nil {
		env.WebhookInstallOptions.MutatingWebhooks = []*admissionregistration.MutatingWebhookConfiguration{mutatingWebhookCfg}
	}
}
