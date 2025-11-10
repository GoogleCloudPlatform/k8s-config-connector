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

package composer

import (
	"context"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

func ResolveComposerEnvironmentRefs(ctx context.Context, reader client.Reader, obj *krm.ComposerEnvironment) error {
	config := obj.Spec.Config
	if config != nil {
		nodeConfig := config.NodeConfig
		if nodeConfig != nil {
			if nodeConfig.NetworkRef != nil {
				if err := nodeConfig.NetworkRef.Normalize(ctx, reader, obj); err != nil {
					return err
				}
			}
			if nodeConfig.SubnetworkRef != nil {
				subnetwork, err := refs.ResolveComputeSubnetwork(ctx, reader, obj, nodeConfig.SubnetworkRef)
				if err != nil {
					return err
				}
				nodeConfig.SubnetworkRef = subnetwork
			}
			if nodeConfig.ServiceAccountRef != nil {
				if err := nodeConfig.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
					return err
				}
			}
			if nodeConfig.ComposerNetworkAttachmentRef != nil {
				_, err := nodeConfig.ComposerNetworkAttachmentRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
				if err != nil {
					return err
				}
			}
		}
		if config.PrivateEnvironmentConfig != nil && config.PrivateEnvironmentConfig.CloudComposerConnectionSubnetworkRef != nil {
			ref := config.PrivateEnvironmentConfig.CloudComposerConnectionSubnetworkRef
			subnetwork, err := refs.ResolveComputeSubnetwork(ctx, reader, obj, ref)
			if err != nil {
				return err
			}
			ref = subnetwork
		}
		if config.EncryptionConfig != nil && config.EncryptionConfig.KMSKeyRef != nil {
			ref := config.EncryptionConfig.KMSKeyRef
			key, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, obj, ref)
			if err != nil {
				return err
			}
			ref = key
		}
	}
	if obj.Spec.StorageConfig != nil && obj.Spec.StorageConfig.BucketRef != nil {
		ref := obj.Spec.StorageConfig.BucketRef
		bucket, err := ref.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return err
		}
		// refine external to the API required format
		// The name of the Cloud Storage bucket used by the environment. No
		// `gs://` prefix.
		tokens := strings.Split(bucket, "/")
		storageBucketName := tokens[len(tokens)-1]
		ref.External = storageBucketName
		return nil
	}
	return nil
}
