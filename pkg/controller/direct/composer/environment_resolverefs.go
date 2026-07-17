// Copyright 2026 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveEnvironmentRefs(ctx context.Context, kube client.Reader, obj *krm.ComposerEnvironment) error {
	var err error
	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return err
	}

	if obj.Spec.Config != nil {
		if obj.Spec.Config.EncryptionConfig != nil && obj.Spec.Config.EncryptionConfig.KMSKeyRef != nil {
			obj.Spec.Config.EncryptionConfig.KMSKeyRef, err = refs.ResolveKMSCryptoKeyRef(ctx, kube, obj, obj.Spec.Config.EncryptionConfig.KMSKeyRef)
			if err != nil {
				return err
			}
		}
		if obj.Spec.Config.NodeConfig != nil {
			nodeConfig := obj.Spec.Config.NodeConfig
			if nodeConfig.SubnetworkRef != nil {
				if err := nodeConfig.SubnetworkRef.Normalize(ctx, kube, obj.GetNamespace()); err != nil {
					return err
				}
			}
			if nodeConfig.ServiceAccountRef != nil {
				err = nodeConfig.ServiceAccountRef.Resolve(ctx, kube, obj)
				if err != nil {
					return err
				}
			}
			if nodeConfig.ComposerNetworkAttachmentRef != nil {
				if err := nodeConfig.ComposerNetworkAttachmentRef.Normalize(ctx, kube, obj.GetNamespace()); err != nil {
					return err
				}
			}
		}
		if obj.Spec.Config.PrivateEnvironmentConfig != nil {
			pec := obj.Spec.Config.PrivateEnvironmentConfig
			if pec.CloudComposerConnectionSubnetworkRef != nil {
				if err := pec.CloudComposerConnectionSubnetworkRef.Normalize(ctx, kube, obj.GetNamespace()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
