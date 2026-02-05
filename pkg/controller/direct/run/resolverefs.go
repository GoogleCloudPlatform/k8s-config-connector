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
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	sqlv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveRunJobRefs(ctx context.Context, kube client.Reader, desired *krm.RunJob) error {
	var err error
	if desired.Spec.Template == nil {
		return nil
	}
	template := desired.Spec.Template.Template
	if template != nil {
		if template.EncryptionKeyRef != nil {
			template.EncryptionKeyRef, err = refs.ResolveKMSCryptoKeyRef(ctx, kube, desired, template.EncryptionKeyRef)
			if err != nil {
				return err
			}
		}
		if template.ServiceAccountRef != nil {
			err = template.ServiceAccountRef.Resolve(ctx, kube, desired)
			if err != nil {
				return err
			}
		}
		if template.VPCAccess != nil && template.VPCAccess.ConnectorRef != nil {
			template.VPCAccess.ConnectorRef.External, err = template.VPCAccess.ConnectorRef.NormalizedExternal(ctx, kube, desired.GetNamespace())
			if err != nil {
				return err
			}
		}
		for _, c := range template.Containers {
			for _, env := range c.Env {
				if env.ValueSource != nil && env.ValueSource.SecretKeyRef != nil {
					sm := env.ValueSource.SecretKeyRef
					// Different from the TF-based RunJob, which requires both secret and secret version to be specified,
					// the KRM-based RunJob allows users to only specify the secret version, and the secret will be inferred.
					if sm.SecretRef != nil {
						sm.SecretRef.External, err = sm.SecretRef.NormalizedExternal(ctx, kube, desired.GetNamespace())
						if err != nil {
							return fmt.Errorf("missing required field `.spec.template.template.containers[].env[].valueSource.secretKeyRef.secretRef` when using the short version secretVersionRef: %w", err)
						}
					}
					if sm.VersionRef != nil && sm.VersionRef.External == "" {
						sm.VersionRef.External, err = sm.VersionRef.NormalizedExternal(ctx, kube, desired.GetNamespace())
						if err != nil {
							return err
						}
						fullSecretVersionExternal := sm.VersionRef.External

						// GCP server has special requirements for SecretManager.
						// 1. SecretManager must be specified
						// 2. SecertVersion must be in the short version.
						if sm.SecretRef == nil {
							sm.SecretRef = &secretmanagerv1beta1.SecretRef{
								External: strings.Split(fullSecretVersionExternal, "/versions/")[0],
							}
						}
						sm.VersionRef.External = strings.Split(sm.VersionRef.External, "/versions/")[1]
					}
				}
			}
		}
		for _, v := range template.Volumes {
			if v.CloudSQLInstance != nil {
				for _, sqlInstance := range v.CloudSQLInstance.InstanceRefs {
					if err := sqlInstance.Normalize(ctx, kube, desired.GetNamespace()); err != nil {
						return err
					}
					id := &sqlv1beta1.SQLInstanceIdentity{}
					if err := id.FromExternal(sqlInstance.External); err != nil {
						return err
					}
					sqlInstance.External = id.ConnectionName()
				}
			}
			if v.GCS != nil && v.GCS.BucketRef != nil {
				err = v.GCS.BucketRef.Normalize(ctx, kube, desired.GetNamespace())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
