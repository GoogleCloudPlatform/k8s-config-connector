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

package workstations

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveWorkstationConfigRefs(ctx context.Context, kube client.Reader, obj *krm.WorkstationConfig) error {
	var err error
	if obj.Spec.Host != nil && obj.Spec.Host.GceInstance != nil && obj.Spec.Host.GceInstance.ServiceAccountRef != nil {
		err = obj.Spec.Host.GceInstance.ServiceAccountRef.Resolve(ctx, kube, obj)
		if err != nil {
			return err
		}
	}
	if obj.Spec.EncryptionKey != nil {
		if obj.Spec.EncryptionKey.KMSCryptoKeyRef != nil {
			ref := obj.Spec.EncryptionKey.KMSCryptoKeyRef
			_, err := ref.NormalizedExternal(ctx, kube, obj.Namespace)
			if err != nil {
				return err
			}
		}
		if obj.Spec.EncryptionKey.ServiceAccountRef != nil {
			err = obj.Spec.EncryptionKey.ServiceAccountRef.Resolve(ctx, kube, obj)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
