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

package sql

import (
	"context"
	"fmt"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/kms/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveSQLInstanceRefs(ctx context.Context, kube client.Reader, obj *krm.SQLInstance) error {
	if err := resolveCryptoKeyRef(ctx, kube, obj); err != nil {
		return err
	}
	if err := resolveMasterInstanceRef(ctx, kube, obj); err != nil {
		return err
	}
	if err := resolveReplicaPasswordRef(ctx, kube, obj); err != nil {
		return err
	}
	if err := resolveRootPasswordRef(ctx, kube, obj); err != nil {
		return err
	}
	if err := resolvePrivateNetworkRef(ctx, kube, obj); err != nil {
		return err
	}
	if err := resolveAuditLogBucketRef(ctx, kube, obj); err != nil {
		return err
	}
	if err := resolveSourceSQLInstanceRef(ctx, kube, obj); err != nil {
		return err
	}
	return nil
}

func resolveCryptoKeyRef(ctx context.Context, kube client.Reader, obj *krm.SQLInstance) error {
	if obj.Spec.EncryptionKMSCryptoKeyRef == nil {
		return nil
	}

	keyRef := obj.Spec.EncryptionKMSCryptoKeyRef

	if keyRef.External != "" && keyRef.Name != "" {
		return fmt.Errorf("cannot specify both spec.encryptionKMSCryptoKeyRef.external and spec.encryptionKMSCryptoKeyRef.name")
	}

	if keyRef.External != "" {
		return nil
	} else if keyRef.Name != "" {
		if keyRef.Namespace == "" {
			keyRef.Namespace = obj.Namespace
		}

		key := types.NamespacedName{
			Namespace: keyRef.Namespace,
			Name:      keyRef.Name,
		}

		cryptoKey := &unstructured.Unstructured{}
		cryptoKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
		if err := kube.Get(ctx, key, cryptoKey); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(kmsv1beta1.KMSCryptoKeyGVK, key)
			}
			return fmt.Errorf("error reading referenced KMSCryptoKey %v: %w", cryptoKey, err)
		}

		keyLink, _, err := unstructured.NestedString(cryptoKey.Object, "status", "selfLink")
		if err != nil || keyLink == "" {
			return fmt.Errorf("reading status.selfLink from %v %v/%v: %w", cryptoKey.GroupVersionKind().Kind, cryptoKey.GetNamespace(), cryptoKey.GetName(), err)
		}

		obj.Spec.EncryptionKMSCryptoKeyRef.External = keyLink

		return nil
	} else {
		return fmt.Errorf("must specify either spec.encryptionKMSCryptoKeyRef.external or spec.encryptionKMSCryptoKeyRef.name")
	}
}

func resolveMasterInstanceRef(ctx context.Context, kube client.Reader, obj *krm.SQLInstance) error {
	if obj.Spec.MasterInstanceRef == nil {
		return nil
	}

	if obj.Spec.MasterInstanceRef.External != "" && obj.Spec.MasterInstanceRef.Name != "" {
		return fmt.Errorf("cannot specify both spec.masterInstanceRef.external and spec.masterInstanceRef.name")
	}

	if obj.Spec.MasterInstanceRef.External != "" {
		return nil
	} else if obj.Spec.MasterInstanceRef.Name != "" {
		if obj.Spec.MasterInstanceRef.Namespace == "" {
			obj.Spec.MasterInstanceRef.Namespace = obj.Namespace
		}

		key := types.NamespacedName{
			Namespace: obj.Spec.MasterInstanceRef.Namespace,
			Name:      obj.Spec.MasterInstanceRef.Name,
		}

		masterInstance := &unstructured.Unstructured{}
		masterInstance.SetGroupVersionKind(krm.SQLInstanceGVK)
		if err := kube.Get(ctx, key, masterInstance); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(krm.SQLInstanceGVK, key)
			}
			return fmt.Errorf("error reading referenced master instance %v: %w", key, err)
		}

		masterInstanceName, err := refs.GetResourceID(masterInstance)
		if err != nil {
			return err
		}

		masterInstanceProject, ok := masterInstance.GetAnnotations()[k8s.ProjectIDAnnotation]
		if !ok {
			masterInstanceProject = masterInstance.GetNamespace()
		}

		obj.Spec.MasterInstanceRef.External = fmt.Sprintf("%s:%s", masterInstanceProject, masterInstanceName)

		return nil
	} else {
		return fmt.Errorf("must specify either spec.masterInstanceRef.external or spec.masterInstanceRef.name")
	}
}

func resolveReplicaPasswordRef(ctx context.Context, kube client.Reader, obj *krm.SQLInstance) error {
	if obj.Spec.ReplicaConfiguration == nil || obj.Spec.ReplicaConfiguration.Password == nil {
		return nil
	}

	if obj.Spec.ReplicaConfiguration.Password.Value != nil && obj.Spec.ReplicaConfiguration.Password.ValueFrom != nil {
		return fmt.Errorf("cannot specify both spec.replicaConfiguration.password.value and spec.replicaConfiguration.password.valueFrom")
	}

	if obj.Spec.ReplicaConfiguration.Password.Value != nil {
		return nil
	} else if obj.Spec.ReplicaConfiguration.Password.ValueFrom != nil {
		key := types.NamespacedName{
			Namespace: obj.Namespace,
			Name:      obj.Spec.ReplicaConfiguration.Password.ValueFrom.SecretKeyRef.Name,
		}

		secret := &corev1.Secret{}
		if err := kube.Get(ctx, key, secret); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewSecretNotFoundError(key)
			}
			return fmt.Errorf("error reading referenced Secret %v: %w", key, err)
		}

		password := string(secret.Data[obj.Spec.ReplicaConfiguration.Password.ValueFrom.SecretKeyRef.Key])

		obj.Spec.ReplicaConfiguration.Password.Value = direct.PtrTo(password)

		return nil
	}
	return nil
}

func resolveRootPasswordRef(ctx context.Context, kube client.Reader, obj *krm.SQLInstance) error {
	if obj.Spec.RootPassword == nil {
		return nil
	}

	if obj.Spec.RootPassword.Value != nil && obj.Spec.RootPassword.ValueFrom != nil {
		return fmt.Errorf("cannot specify both spec.rootPassword.value and spec.rootPassword.valueFrom")
	}

	if obj.Spec.RootPassword.Value != nil {
		return nil
	} else if obj.Spec.RootPassword.ValueFrom != nil {
		key := types.NamespacedName{
			Namespace: obj.Namespace,
			Name:      obj.Spec.RootPassword.ValueFrom.SecretKeyRef.Name,
		}

		secret := &corev1.Secret{}
		if err := kube.Get(ctx, key, secret); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewSecretNotFoundError(key)
			}
			return fmt.Errorf("error reading referenced Secret %v: %w", key, err)
		}

		password := string(secret.Data[obj.Spec.RootPassword.ValueFrom.SecretKeyRef.Key])

		obj.Spec.RootPassword.Value = direct.PtrTo(password)

		return nil
	}
	return nil
}

func resolvePrivateNetworkRef(ctx context.Context, kube client.Reader, obj *krm.SQLInstance) error {
	if obj.Spec.Settings.IpConfiguration == nil || obj.Spec.Settings.IpConfiguration.PrivateNetworkRef == nil {
		return nil
	}

	resRef := obj.Spec.Settings.IpConfiguration.PrivateNetworkRef
	netRef := &computev1beta1.ComputeNetworkRef{
		External:  resRef.External,
		Name:      resRef.Name,
		Namespace: resRef.Namespace,
	}
	external, err := netRef.NormalizedExternal(ctx, kube, obj.GetNamespace())
	if err != nil {
		return err
	}

	obj.Spec.Settings.IpConfiguration.PrivateNetworkRef.External = external

	return nil
}

func resolveAuditLogBucketRef(ctx context.Context, kube client.Reader, obj *krm.SQLInstance) error {
	if obj.Spec.Settings.SqlServerAuditConfig == nil {
		return nil
	}

	if obj.Spec.Settings.SqlServerAuditConfig.BucketRef == nil {
		return fmt.Errorf("must specify bucket for audit config")
	}

	bucketRef := obj.Spec.Settings.SqlServerAuditConfig.BucketRef

	if bucketRef.External != "" && bucketRef.Name != "" {
		return fmt.Errorf("cannot specify both spec.settings.sqlServerAuditConfig.bucketRef.external and spec.settings.sqlServerAuditConfig.bucketRef.name")
	}

	if bucketRef.External != "" {
		return nil
	} else if bucketRef.Name != "" {
		if bucketRef.Namespace == "" {
			bucketRef.Namespace = obj.Namespace
		}

		key := types.NamespacedName{
			Namespace: bucketRef.Namespace,
			Name:      bucketRef.Name,
		}

		bucket := &unstructured.Unstructured{}
		bucket.SetGroupVersionKind(storagev1beta1.StorageBucketGVK)
		if err := kube.Get(ctx, key, bucket); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(storagev1beta1.StorageBucketGVK, key)
			}
			return fmt.Errorf("error reading referenced StorageBucket %v: %w", key, err)
		}

		storageBucketName, err := refs.GetResourceID(bucket)
		if err != nil {
			return err
		}
		obj.Spec.Settings.SqlServerAuditConfig.BucketRef.External = "gs://" + storageBucketName

		return nil
	} else {
		return fmt.Errorf("must specify either spec.settings.sqlServerAuditConfig.bucketRef.external or spec.settings.sqlServerAuditConfig.bucketRef.name")
	}
}

func resolveSourceSQLInstanceRef(ctx context.Context, kube client.Reader, obj *krm.SQLInstance) error {
	if obj.Spec.CloneSource == nil {
		return nil
	}

	sqlInstanceRef := obj.Spec.CloneSource.SQLInstanceRef

	if sqlInstanceRef.External != "" && sqlInstanceRef.Name != "" {
		return fmt.Errorf("cannot specify both spec.settings.cloneSource.sqlInstanceRef.external and spec.settings.cloneSource.sqlInstanceRef.name")
	}

	if sqlInstanceRef.External != "" {
		return nil
	} else if sqlInstanceRef.Name != "" {
		if sqlInstanceRef.Namespace == "" {
			sqlInstanceRef.Namespace = obj.Namespace
		}

		key := types.NamespacedName{
			Namespace: sqlInstanceRef.Namespace,
			Name:      sqlInstanceRef.Name,
		}

		sqlInstance := &unstructured.Unstructured{}
		sqlInstance.SetGroupVersionKind(krm.SQLInstanceGVK)
		if err := kube.Get(ctx, key, sqlInstance); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(krm.SQLInstanceGVK, key)
			}
			return fmt.Errorf("error reading referenced SQLInstance %v: %w", key, err)
		}

		sqlInstanceName, err := refs.GetResourceID(sqlInstance)
		if err != nil {
			return err
		}

		obj.Spec.CloneSource.SQLInstanceRef.External = sqlInstanceName

		return nil
	} else {
		return fmt.Errorf("must specify either spec.settings.cloneSource.sqlInstanceRef.external or spec.settings.cloneSource.sqlInstanceRef.name")
	}
}
