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

package preflight

import (
	"context"
	"fmt"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/asserts"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNameChecker_ConfigConnector(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		cc   *corev1beta1.ConfigConnector
		err  error
	}{
		{
			name: "predefined name",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
			},
			err: nil,
		},
		{
			name: "random name",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-configconnector",
				},
			},
			err: fmt.Errorf("the only allowed name for ConfigConnector object is '%v'", corev1beta1.ConfigConnectorAllowedName),
		},
	}

	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	client := mgr.GetClient()
	checker := NewNameChecker(client, corev1beta1.ConfigConnectorAllowedName)
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := checker.Preflight(context.TODO(), tc.cc)
			asserts.AssertErrorIsExpected(t, err, tc.err)
		})
	}
}

func TestNameChecker_ConfigConnectorContext(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		ccc  *corev1beta1.ConfigConnectorContext
		err  error
	}{
		{
			name: "valid ConfigConnectorContext name",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorContextAllowedName,
				},
			},
			err: nil,
		},
		{
			name: "random name",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-ccc",
				},
			},
			err: fmt.Errorf("the only allowed name for ConfigConnectorContext object is '%v'", corev1beta1.ConfigConnectorContextAllowedName),
		},
	}

	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	client := mgr.GetClient()
	checker := NewNameChecker(client, corev1beta1.ConfigConnectorContextAllowedName)
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := checker.Preflight(context.TODO(), tc.ccc)
			asserts.AssertErrorIsExpected(t, err, tc.err)
		})
	}
}
