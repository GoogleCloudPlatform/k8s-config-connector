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

const (
	ControllerManagerServiceAccountRegex = "system:serviceaccount:[a-z0-9.-]+:cnrm-controller-manager"
	certDir                              = "/tmp/cert"
	certSecretName                       = "cnrm-webhook-cert"
)

// ServicePort is the port that the webhook binary will bind to, as well as use as the service port.
//
// must be 443 as private GKE clusters have opened up 443 specifically
// as a port that GKE masters can send requests to nodes to, and the requests are sent
// directly to the targetPort of the pod rather than the service port.
// see b/180354275
//
// Since the TargetPort is effectively being used as a public port,
// standardizing public ports removes ambiguity.
//
// This is a mutable variable for tests.
var ServicePort = 443
