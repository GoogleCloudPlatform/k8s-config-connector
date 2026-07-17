// Copyright 2025 Google LLC
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

package compute

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func resolveTargetHTTPSProxyRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeTargetHTTPSProxy) error {
	if obj.Spec.UrlMapRef != nil {
		if err := obj.Spec.UrlMapRef.Normalize(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}

	if obj.Spec.SslPolicyRef != nil {
		if err := obj.Spec.SslPolicyRef.Normalize(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}

	if obj.Spec.CertificateMapRef != nil {
		if err := obj.Spec.CertificateMapRef.Normalize(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}

	if obj.Spec.ServerTlsPolicyRef != nil {
		if err := obj.Spec.ServerTlsPolicyRef.Normalize(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}

	for i := range obj.Spec.SslCertificates {
		if err := obj.Spec.SslCertificates[i].Normalize(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}

	for i := range obj.Spec.CertificateManagerCertificates {
		if err := obj.Spec.CertificateManagerCertificates[i].Normalize(ctx, reader, obj.Namespace); err != nil {
			return err
		}
	}

	return nil
}
