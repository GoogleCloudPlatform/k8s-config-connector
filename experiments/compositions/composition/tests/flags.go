// Copyright 2024 Google LLC
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

package e2e

import (
	"flag"
)

// Constants used in the file
const (
	Local  = "local"
	None   = "none"
	Remote = "remote"
)

// E2E enables running end-to-end tests.
var E2E = flag.Bool("e2e", true, "if true, run end-to-end tests.")

// KubernetesVersion is the version of Kubernetes to test against. Only has effect
// when testing against test-created Kind clusters.
var KubernetesVersion = flag.String("kubernetes-version", "1.23",
	"The version of Kubernetes to create")

// Operator - Where is the reconciler for the KRM objects running
var Operator = flag.String("operator", "none", "One of [local,none,oncluster].\n"+
	"`local` implies operator is started locally by test code\n"+
	"`none` implies the  user has started the operator already\n"+
	"`remote` implies the test installs the operator manifests using <image>.\n"+
	"In all cases, --kubeconfig is used to connect to admin-cluster.")

// Image - if a custom image is to be used
var Image = flag.String("image", "", "<image> to be used for the operator container")

// Kubeconfig - if a local kubeconfig is to be used
var Kubeconfig = flag.String("kube-config", "", "location of kubeconfig file")

// Debug - for logging
var Debug = flag.Bool("debug", false, "If true, enables verbose logging")
