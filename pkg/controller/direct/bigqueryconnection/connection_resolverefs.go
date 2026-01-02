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

package bigqueryconnection

import (
	"context"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveBigQueryConnectionConnectionRefs(ctx context.Context, reader client.Reader, obj *krm.BigQueryConnectionConnection) error {
	// Resolve SQLInstanceRef and SQLDatabaseRef
	if obj.Spec.CloudSQLSpec != nil {
		sql := obj.Spec.CloudSQLSpec
		if sql.InstanceRef != nil {
			instance, err := refs.ResolveSQLInstanceRef(ctx, reader, obj, sql.InstanceRef)
			if err != nil {
				return err
			}
			sql.InstanceRef.External = instance.ConnectionName()
		}
		if sql.DatabaseRef != nil {
			database, err := refs.ResolveSQLDatabaseRef(ctx, reader, obj, sql.DatabaseRef)
			if err != nil {
				return err
			}
			sql.DatabaseRef.External = database.Name()
		}
		if sql.Credential != nil {
			if err := refsv1beta1secret.NormalizedSecret(ctx, sql.Credential.SecretRef, reader, obj.Namespace); err != nil {
				return err
			}
		}
	}

	// Resolve SpannerDatabaseRef
	if obj.Spec.CloudSpannerSpec != nil {
		if ref := obj.Spec.CloudSpannerSpec.DatabaseRef; ref != nil {
			_, err := ref.NormalizedExternal(ctx, reader, obj.Namespace)
			if err != nil {
				return err
			}
		}
	}

	// Resolve Spark.DataprocClusterRef and Spark.MetastoreServiceRef
	if obj.Spec.SparkSpec != nil {
		if obj.Spec.SparkSpec.SparkHistoryServer != nil {
			if ref := obj.Spec.SparkSpec.SparkHistoryServer.DataprocClusterRef; ref != nil {
				_, err := ref.NormalizedExternal(ctx, reader, obj.GetNamespace())
				if err != nil {
					return err
				}
			}
		}

		if obj.Spec.SparkSpec.MetastoreService != nil {
			if ref := obj.Spec.SparkSpec.MetastoreService.MetastoreServiceRef; ref != nil {
				_, err := ref.NormalizedExternal(ctx, reader, obj.GetNamespace())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
