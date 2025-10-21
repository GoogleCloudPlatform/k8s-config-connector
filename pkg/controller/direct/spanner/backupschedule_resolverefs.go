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

package spanner

import (
	"context"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func resolveBackupScheduleRefs(ctx context.Context, reader client.Reader, obj *krm.SpannerBackupSchedule) error {
	if obj.Spec.EncryptionConfig != nil && obj.Spec.EncryptionConfig.KMSKeyRef != nil {
		key, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, obj, obj.Spec.EncryptionConfig.KMSKeyRef)
		if err != nil {
			return err
		}
		obj.Spec.EncryptionConfig.KMSKeyRef = key
	}

	if obj.Spec.EncryptionConfig != nil && obj.Spec.EncryptionConfig.KMSKeyRefs != nil {
		var keys []*refs.KMSCryptoKeyRef
		for _, kmsKey := range obj.Spec.EncryptionConfig.KMSKeyRefs {
			key, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, obj, kmsKey)
			if err != nil {
				return err
			}
			keys = append(keys, key)
		}
		obj.Spec.EncryptionConfig.KMSKeyRefs = keys

	}
	return nil
}
