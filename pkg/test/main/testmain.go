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

package testmain

import (
	//"io"
	"log"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testenvironment "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/environment"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/webhook"
	cnrmwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func ForIntegrationTests(m *testing.M, mgr *manager.Manager) {
	if os.Getenv("E2E_GCP_TARGET") != "" {
		log.Fatalf("dynamic integration tests do not support variable E2E_GCP_TARGET")
	}

	// Since Terraform logging defers to the Go standard logger,
	// here we discard everything logged onto the Go standard logger to
	// disable logging from Terraform Google provider in integration tests.
	//log.SetOutput(io.Discard)
	TestMain(m, test.IntegrationTestType, nil, mgr)
}

func ForUnitTests(m *testing.M, mgr *manager.Manager) {
	TestMain(m, test.UnitTestType, nil, mgr)
}

func ForUnitTestsWithCRDs(m *testing.M, crds []*apiextensions.CustomResourceDefinition, mgr *manager.Manager) {
	TestMain(m, test.UnitTestType, crds, mgr)
}

// TestMain starts a local K8S API server to run tests against. These tests do
// not require an external API server to execute.
func TestMain(m *testing.M, testType test.Type, crds []*apiextensions.CustomResourceDefinition, mgr *manager.Manager) {
	SetupMultipleEnvironments(m, testType, crds, []*manager.Manager{mgr})
}

// SetupMultipleEnvironments starts n API servers to run tests against. The value for 'n' is determined by
// the length of the 'mgrPtrs' argument. This is useful when testing multi-cluster scenarios.
func SetupMultipleEnvironments(m *testing.M, testType test.Type, crds []*apiextensions.CustomResourceDefinition, mgrPtrs []*manager.Manager) {
	logging.SetupLogger()
	var err error

	envs := make([]*envtest.Environment, 0, len(mgrPtrs))
	stops := make([]func(), 0, len(mgrPtrs))
	for _, mp := range mgrPtrs {
		var whCfgs []cnrmwebhook.Config
		if testType == test.IntegrationTestType {
			whCfgs, err = webhook.GetTestCommonWebhookConfigs()
			if err != nil {
				log.Fatalf("error getting common wehbook configs: %v", err)
			}
		}
		env := testenvironment.StartTestEnvironmentOrLogFatal(testType, crds, whCfgs)
		envs = append(envs, env)

		mgr, stop := testcontroller.StartTestManagerInstance(env, testType, whCfgs)
		stops = append(stops, stop)

		if err := testcontroller.EnsureNamespaceExists(mgr.GetClient(), k8s.SystemNamespace); err != nil {
			log.Fatalf("error ensuring namesapce exists: %v", err)
		}
		*mp = mgr
	}

	code := m.Run()

	for _, stop := range stops {
		stop()
	}
	for _, env := range envs {
		if err := env.Stop(); err != nil {
			log.Printf("unable to stop at least one test environment: %v", err)
		}
	}
	os.Exit(code)
}
