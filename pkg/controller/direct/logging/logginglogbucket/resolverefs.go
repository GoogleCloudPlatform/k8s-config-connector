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

package logginglogbucket

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveLoggingLogBucketRefs(ctx context.Context, kube client.Reader, obj *krm.LoggingLogBucket) error {
	if obj.Spec.CmekSettings != nil && obj.Spec.CmekSettings.KmsKeyRef != nil {
		kmsKeyRef, err := refsv1beta1.ResolveKMSCryptoKeyRef(ctx, kube, obj, obj.Spec.CmekSettings.KmsKeyRef)
		if err != nil {
			return err
		}
		obj.Spec.CmekSettings.KmsKeyRef = kmsKeyRef
	}
	return nil
}
