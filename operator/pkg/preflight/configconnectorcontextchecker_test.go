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
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/asserts"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConfigConnectorContextChecker(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		ccc  *corev1beta1.ConfigConnectorContext
		err  error
	}{
		{
			name: "CCC has spec.billingProject set and spec.requestProjectPolicy set to BILLING_PROJECT",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "BILLING_PROJECT",
					BillingProject:       "BILL_ME",
				},
			},
			err: nil,
		},

		{
			name: "CCC has spec.billingProject omitted and spec.requestProjectPolicy set to RESOURCE_PROJECT",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "RESOURCE_PROJECT",
				},
			},
			err: nil,
		},

		{
			name: "CCC has spec.billingProject set to empty and spec.requestProjectPolicy set to RESOURCE_PROJECT",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "RESOURCE_PROJECT",
					BillingProject:       "",
				},
			},
			err: nil,
		},

		{
			name: "CCC has spec.billingProject omitted and spec.requestProjectPolicy set to SERVICE_ACCOUNT_PROJECT",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "SERVICE_ACCOUNT_PROJECT",
				},
			},
			err: nil,
		},

		{
			name: "CCC has spec.billingProject unset and requestProjectPolicy set to BILLING_PROJECT",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "BILLING_PROJECT",
				},
			},
			err: fmt.Errorf("spec.billingProject must be set if spec.requestProjectPolicy is set to %v", k8s.BillingProjectPolicy),
		},

		{
			name: "CCC has spec.billingProject set to empty and requestProjectPolicy set to BILLING_PROJECT",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "BILLING_PROJECT",
					BillingProject:       "",
				},
			},
			err: fmt.Errorf("spec.billingProject must be set if spec.requestProjectPolicy is set to %v", k8s.BillingProjectPolicy),
		},

		{
			name: "CCC has spec.billingProject unset and requestProjectPolicy set to RESOURCE_PROJECT",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "RESOURCE_PROJECT",
					BillingProject:       "BILL_ME",
				},
			},
			err: fmt.Errorf("spec.billingProject cannot be set if spec.requestProjectPolicy is not set to %v", k8s.BillingProjectPolicy),
		},

		{
			name: "CCC has spec.billingProject unset and requestProjectPolicy set to SERVICE_ACCOUNT_PROJECT",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "SERVICE_ACCOUNT_PROJECT",
					BillingProject:       "BILL_ME",
				},
			},
			err: fmt.Errorf("spec.billingProject cannot be set if spec.requestProjectPolicy is not set to %v", k8s.BillingProjectPolicy),
		},
		{
			name: "CCC with WIF only is valid",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					WorkloadIdentityFederation: &corev1beta1.WorkloadIdentityFederationSpec{
						CredentialSecretName: "my-wif-secret",
						Audience:             "//iam.googleapis.com/projects/12345/locations/global/workloadIdentityPools/pool/providers/provider",
					},
				},
			},
			err: nil,
		},
		{
			name: "CCC with both GSA and WIF is invalid",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					WorkloadIdentityFederation: &corev1beta1.WorkloadIdentityFederationSpec{
						CredentialSecretName: "my-wif-secret",
						Audience:             "//iam.googleapis.com/projects/12345/locations/global/workloadIdentityPools/pool/providers/provider",
					},
				},
			},
			err: fmt.Errorf("spec.googleServiceAccount and spec.workloadIdentityFederation are mutually exclusive"),
		},
		{
			name: "CCC with WIF missing credentialSecretName is invalid",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					WorkloadIdentityFederation: &corev1beta1.WorkloadIdentityFederationSpec{
						Audience: "//iam.googleapis.com/projects/12345/locations/global/workloadIdentityPools/pool/providers/provider",
					},
				},
			},
			err: fmt.Errorf("spec.workloadIdentityFederation.credentialSecretName is required"),
		},
		{
			name: "CCC with WIF missing audience is invalid",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					WorkloadIdentityFederation: &corev1beta1.WorkloadIdentityFederationSpec{
						CredentialSecretName: "my-wif-secret",
					},
				},
			},
			err: fmt.Errorf("spec.workloadIdentityFederation.audience is required"),
		},
		{
			name: "CCC with neither GSA nor WIF is invalid",
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{},
			},
			err: fmt.Errorf("one of spec.googleServiceAccount or spec.workloadIdentityFederation must be set"),
		},
	}

	checker := NewConfigConnectorContextChecker()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := checker.Preflight(context.TODO(), tc.ccc)
			asserts.AssertErrorIsExpected(t, err, tc.err)
		})
	}
}

func TestValidateGSAFormat(t *testing.T) {
	tests := []struct {
		name string
		gsa  string
		err  error
	}{
		{
			name: "empty",
			gsa:  "",
			err:  nil,
		},
		{
			name: "valid GSA format",
			gsa:  "foo@abc.gserviceaccount.com",
			err:  nil,
		},
		{
			name: "valid GSA format",
			gsa:  "foo@abc.def.gserviceaccount.com",
			err:  nil,
		},
		{
			name: "valid GSA format",
			gsa:  "foo@abc.def.ghi.gserviceaccount.com",
			err:  nil,
		},
		{
			name: "invalid GSA format",
			gsa:  "abc",
			err:  fmt.Errorf("invalid GoogleServiceAccount format for %q", "abc"),
		},
		{
			name: "invalid GSA format",
			gsa:  "foo@bar.com",
			err:  fmt.Errorf("invalid GoogleServiceAccount format for %q", "foo@bar.com"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateGSAFormat(tc.gsa)
			asserts.AssertErrorIsExpected(t, err, tc.err)
		})
	}
}
