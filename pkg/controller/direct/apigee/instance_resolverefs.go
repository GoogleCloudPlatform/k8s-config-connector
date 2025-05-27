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

package apigee

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveApigeeInstanceRefs(ctx context.Context, kube client.Reader, obj *krm.ApigeeInstance) error {
	if ref := obj.Spec.DiskEncryptionKMSCryptoKeyRef; ref != nil {
		_, err := ref.NormalizedExternal(ctx, kube, obj.Namespace)
		if err != nil {
			return err
		}
	}
	return nil
}
