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
		external, err := obj.Spec.UrlMapRef.NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.UrlMapRef.External = external
	}

	if obj.Spec.SslPolicyRef != nil {
		external, err := obj.Spec.SslPolicyRef.NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.SslPolicyRef.External = external
	}

	if obj.Spec.CertificateMapRef != nil {
		external, err := obj.Spec.CertificateMapRef.NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.CertificateMapRef.External = external
	}

	if obj.Spec.ServerTlsPolicyRef != nil {
		external, err := obj.Spec.ServerTlsPolicyRef.NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.ServerTlsPolicyRef.External = external
	}

	for i := range obj.Spec.SslCertificates {
		external, err := obj.Spec.SslCertificates[i].NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.SslCertificates[i].External = external
	}

	for i := range obj.Spec.CertificateManagerCertificates {
		external, err := obj.Spec.CertificateManagerCertificates[i].NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.CertificateManagerCertificates[i].External = external
	}

	return nil
}
