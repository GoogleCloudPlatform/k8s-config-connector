// Copyright 2023 Google LLC
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

package main

import (
	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	addonsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/api/v1alpha1"
	commonoperator "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/pkg/commonoperator"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/pkg/controllers"
)

func main() {
	op := commonoperator.Operator{
		LeaderElectionID: "experiments.cnrm.cloud.google.com",
	}
	op.RegisterSchema(addonsv1alpha1.AddToScheme)
	op.RegisterReconciler(&controllers.CompositeDefinitionReconciler{})
	op.RunMain()
}
