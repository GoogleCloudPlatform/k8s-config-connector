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
	krmrunv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	vpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vpcaccess/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveRunWorkerPoolRefs(ctx context.Context, kube client.Reader, desired *krmrunv1alpha1.RunWorkerPool) error {
	if desired.Spec.Template == nil {
		return nil
	}
	template := desired.Spec.Template
	if err := resolveCommonTemplateRefs(ctx, kube, desired, template.EncryptionKeyRef, template.ServiceAccountRef, template.VPCAccess); err != nil {
		return err
	}

	for i := range template.Containers {
		c := &template.Containers[i]
		for j := range c.Env {
			env := &c.Env[j]
			if env.ValueSource != nil && env.ValueSource.SecretKeyRef != nil {
				sm := env.ValueSource.SecretKeyRef
				if err := resolveSecretKeySelector(ctx, kube, desired.GetNamespace(), &sm.SecretRef, sm.VersionRef); err != nil {
					return fmt.Errorf("resolving .spec.template.containers[%d].env[%d].valueSource.secretKeyRef: %w", i, j, err)
				}
			}
		}
	}
	for i := range template.Volumes {
		v := &template.Volumes[i]
		if err := resolveCommonVolumeRefs(ctx, kube, desired, v.Secret, v.CloudSQLInstance, v.GCS); err != nil {
			return fmt.Errorf("resolving .spec.template.volumes[%d]: %w", i, err)
		}
	}
	return nil
}

func ResolveRunJobRefs(ctx context.Context, kube client.Reader, desired *krm.RunJob) error {
	if desired.Spec.Template == nil || desired.Spec.Template.Template == nil {
		return nil
	}
	template := desired.Spec.Template.Template
	if err := resolveCommonTemplateRefs(ctx, kube, desired, template.EncryptionKeyRef, template.ServiceAccountRef, template.VPCAccess); err != nil {
		return err
	}

	for i := range template.Containers {
		c := &template.Containers[i]
		for j := range c.Env {
			env := &c.Env[j]
			if env.ValueSource != nil && env.ValueSource.SecretKeyRef != nil {
				sm := env.ValueSource.SecretKeyRef
				if err := resolveSecretKeySelector(ctx, kube, desired.GetNamespace(), &sm.SecretRef, sm.VersionRef); err != nil {
					return fmt.Errorf("resolving .spec.template.template.containers[%d].env[%d].valueSource.secretKeyRef: %w", i, j, err)
				}
			}
		}
	}
	for i := range template.Volumes {
		v := &template.Volumes[i]
		if err := resolveCommonVolumeRefs(ctx, kube, desired, v.Secret, v.CloudSQLInstance, v.GCS); err != nil {
			return fmt.Errorf("resolving .spec.template.template.volumes[%d]: %w", i, err)
		}
	}
	return nil
}

// SecretVolumeSource is structurally identical in v1alpha1 and v1beta1,
// but they are different types. We can use a small interface or just take the fields.
// Since we have items with VersionRef, it's better to take the common fields.

type genericSecretVolumeSource interface {
	GetSecretRef() *secretmanagerv1beta1.SecretRef
	GetVersionRefs() []*secretmanagerv1beta1.SecretVersionRef
}

func resolveCommonTemplateRefs(ctx context.Context, kube client.Reader, owner client.Object, encryptionKeyRef *refs.KMSCryptoKeyRef, serviceAccountRef *refs.IAMServiceAccountRef, vpcAccess any) error {
	var err error
	if encryptionKeyRef != nil {
		if _, err := refs.ResolveKMSCryptoKeyRef(ctx, kube, owner, encryptionKeyRef); err != nil {
			return err
		}
	}
	if serviceAccountRef != nil {
		if err := serviceAccountRef.Resolve(ctx, kube, owner); err != nil {
			return err
		}
	}
	if vpcAccess != nil {
		// Use reflection or just check the connectorRef manually if we can.
		// Actually, let's just pass the connectorRef directly if it exists.
		var connectorRef *vpcaccessv1beta1.VPCAccessConnectorRef
		switch v := vpcAccess.(type) {
		case *krmrunv1alpha1.VPCAccess:
			if v != nil {
				connectorRef = v.ConnectorRef
			}
		case *krm.VPCAccess:
			if v != nil {
				connectorRef = v.ConnectorRef
			}
		}
		if connectorRef != nil {
			connectorRef.External, err = connectorRef.NormalizedExternal(ctx, kube, owner.GetNamespace())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func resolveSecretKeySelector(ctx context.Context, kube client.Reader, namespace string, secretRef **secretmanagerv1beta1.SecretRef, versionRef *secretmanagerv1beta1.SecretVersionRef) error {
	var err error
	if *secretRef != nil {
		(*secretRef).External, err = (*secretRef).NormalizedExternal(ctx, kube, namespace)
		if err != nil {
			return err
		}
	}
	if versionRef != nil && versionRef.External == "" {
		versionRef.External, err = versionRef.NormalizedExternal(ctx, kube, namespace)
		if err != nil {
			return err
		}
		fullSecretVersionExternal := versionRef.External

		if *secretRef == nil {
			*secretRef = &secretmanagerv1beta1.SecretRef{
				External: strings.Split(fullSecretVersionExternal, "/versions/")[0],
			}
		}
		versionRef.External = strings.Split(versionRef.External, "/versions/")[1]
	}
	return nil
}

func resolveCommonVolumeRefs(ctx context.Context, kube client.Reader, owner client.Object, secret any, cloudSQLInstance any, gcs any) error {
	var err error

	// Resolve Secret Volume
	if secret != nil {
		var secretRef *secretmanagerv1beta1.SecretRef
		var versionRefs []*secretmanagerv1beta1.SecretVersionRef

		switch s := secret.(type) {
		case *krmrunv1alpha1.SecretVolumeSource:
			if s != nil {
				secretRef = s.SecretRef
				for i := range s.Items {
					versionRefs = append(versionRefs, s.Items[i].VersionRef)
				}
			}
		case *krm.SecretVolumeSource:
			if s != nil {
				secretRef = s.SecretRef
				for i := range s.Items {
					versionRefs = append(versionRefs, s.Items[i].VersionRef)
				}
			}
		}

		if secretRef != nil {
			secretRef.External, err = secretRef.NormalizedExternal(ctx, kube, owner.GetNamespace())
			if err != nil {
				return err
			}
		}
		for _, vRef := range versionRefs {
			if vRef != nil && vRef.External == "" {
				vRef.External, err = vRef.NormalizedExternal(ctx, kube, owner.GetNamespace())
				if err != nil {
					return err
				}
				if strings.Contains(vRef.External, "/versions/") {
					vRef.External = strings.Split(vRef.External, "/versions/")[1]
				}
			}
		}
	}

	// Resolve Cloud SQL Volume
	var instanceRefs []*refs.SQLInstanceRef
	switch c := cloudSQLInstance.(type) {
	case *krmrunv1alpha1.CloudSQLInstance:
		if c != nil {
			instanceRefs = c.InstanceRefs
		}
	case *krm.CloudSQLInstance:
		if c != nil {
			instanceRefs = c.InstanceRefs
		}
	}
	for _, sqlInstance := range instanceRefs {
		instanceRef, err := refs.ResolveSQLInstanceRef(ctx, kube, owner, sqlInstance)
		if err != nil {
			return err
		}
		sqlInstance.External = instanceRef.ConnectionName()
	}

	// Resolve GCS Volume
	var bucketRef *storagev1beta1.StorageBucketRef
	switch g := gcs.(type) {
	case *krmrunv1alpha1.GCSVolumeSource:
		if g != nil {
			bucketRef = g.BucketRef
		}
	case *krm.GCSVolumeSource:
		if g != nil {
			bucketRef = g.BucketRef
		}
	}
	if bucketRef != nil {
		err = bucketRef.Normalize(ctx, kube, owner.GetNamespace())
		if err != nil {
			return err
		}
	}

	return nil
}
