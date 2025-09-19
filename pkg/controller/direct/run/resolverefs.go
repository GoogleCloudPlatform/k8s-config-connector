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

package run

import (
	"context"
	"errors"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveRunJobRefs(ctx context.Context, kube client.Reader, desired *krm.RunJob) error {
	var totalErr, err error
	if desired.Spec.Template != nil {
		template := desired.Spec.Template.Template
		if template != nil {
			if template.EncryptionKeyRef != nil {
				template.EncryptionKeyRef, err = refs.ResolveKMSCryptoKeyRef(ctx, kube, desired, template.EncryptionKeyRef)
				if err != nil {
					errors.Join(totalErr, err)
				}
			}
			if template.ServiceAccountRef != nil {
				err = template.ServiceAccountRef.Resolve(ctx, kube, desired)
				if err != nil {
					errors.Join(totalErr, err)
				}
			}
			if template.VPCAccess != nil && template.VPCAccess.ConnectorRef != nil {
				template.VPCAccess.ConnectorRef.External, err = template.VPCAccess.ConnectorRef.NormalizedExternal(ctx, kube, desired.GetNamespace())
				if err != nil {
					errors.Join(totalErr, err)
				}
			}
			for _, v := range template.Volumes {
				for _, sqlInstance := range v.CloudSQLInstance.InstanceRefs {
					instanceRef, err := refs.ResolveSQLInstanceRef(ctx, kube, desired, sqlInstance)
					if err != nil {
						errors.Join(totalErr, err)
					}
					sqlInstance.External = instanceRef.ConnectionName()
				}
			}
		}
	}
	return totalErr
}
