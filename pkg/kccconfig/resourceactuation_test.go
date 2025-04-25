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

package kccconfig

import (
	"testing"

	opv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
)

func TestDecideActuationMode(t *testing.T) {
	tests := []struct {
		name                  string
		cc                    *opv1beta1.ConfigConnector
		ccc                   *opv1beta1.ConfigConnectorContext
		expectedActuationMode opv1beta1.ActuationMode
	}{
		{
			name: "both CC and CCC specify actuationMode in namespaced mode: defer to CCC",
			cc: &opv1beta1.ConfigConnector{
				Spec: opv1beta1.ConfigConnectorSpec{
					Mode:      opv1beta1.NamespacedMode,
					Actuation: opv1beta1.Reconciling,
				},
			},
			ccc: &opv1beta1.ConfigConnectorContext{
				Spec: opv1beta1.ConfigConnectorContextSpec{
					Actuation: opv1beta1.Paused,
				},
			},
			expectedActuationMode: opv1beta1.Paused,
		},
		{
			name: "only CC specifies in namespaced mode: Use CC",
			cc: &opv1beta1.ConfigConnector{
				Spec: opv1beta1.ConfigConnectorSpec{
					Mode:      opv1beta1.NamespacedMode,
					Actuation: opv1beta1.Paused,
				},
			},
			ccc: &opv1beta1.ConfigConnectorContext{
				Spec: opv1beta1.ConfigConnectorContextSpec{},
			},
			expectedActuationMode: opv1beta1.Paused,
		},
		{
			name: "both CC and CCC specify an actuationMode in cluster mode: ignore CCC",
			cc: &opv1beta1.ConfigConnector{
				Spec: opv1beta1.ConfigConnectorSpec{
					Mode:      opv1beta1.ClusterMode,
					Actuation: opv1beta1.Reconciling,
				},
			},
			ccc: &opv1beta1.ConfigConnectorContext{
				Spec: opv1beta1.ConfigConnectorContextSpec{
					Actuation: opv1beta1.Paused,
				},
			},
			expectedActuationMode: opv1beta1.Reconciling,
		},
		{
			name:                  "neither CC nor CCC specify an actuationMode: Use default",
			cc:                    &opv1beta1.ConfigConnector{},
			ccc:                   &opv1beta1.ConfigConnectorContext{},
			expectedActuationMode: opv1beta1.Reconciling,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := &Configuration{cc: test.cc, ccc: test.ccc}
			actualMode := config.ActuationMode()
			if test.expectedActuationMode != actualMode {
				t.Errorf("DecideActuationMode failed; got %v, want %v", actualMode, test.expectedActuationMode)
			}
		})
	}
}
