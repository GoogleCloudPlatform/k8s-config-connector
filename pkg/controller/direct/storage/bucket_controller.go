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

package storage

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gcp "google.golang.org/api/storage/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.StorageBucketGVK, NewBucketModel)
}

func NewBucketModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBucket{config: *config}, nil
}

var _ directbase.Model = &modelBucket{}

type modelBucket struct {
	config config.ControllerConfig
}

func (m *modelBucket) client(ctx context.Context) (*gcp.Service, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Bucket client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBucket) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.StorageBucket{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBucketIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := resolveReferences(ctx, reader, obj); err != nil {
		return nil, err
	}

	// Get storage GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &BucketAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func resolveReferences(ctx context.Context, reader client.Reader, obj *krm.StorageBucket) error {
	if obj.Spec.IPFilter != nil {
		for i := range obj.Spec.IPFilter.VpcNetworkSources {
			if err := obj.Spec.IPFilter.VpcNetworkSources[i].NetworkRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
	}
	if obj.Spec.Encryption != nil {
		if err := (&obj.Spec.Encryption.KmsKeyRef).Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	return nil
}

func (m *modelBucket) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BucketAdapter struct {
	id        *krm.StorageBucketIdentity
	gcpClient *gcp.Service
	desired   *krm.StorageBucket
	actual    *gcp.Bucket
}

var _ directbase.Adapter = &BucketAdapter{}

// Find retrieves the GCP resource.
func (a *BucketAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Bucket", "name", a.id.Bucket)

	bucket, err := a.gcpClient.Buckets.Get(a.id.Bucket).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Bucket %q: %w", a.id.Bucket, err)
	}

	a.actual = bucket
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BucketAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Bucket", "name", a.id.Bucket)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := StorageBucketSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.Bucket

	created, err := a.gcpClient.Buckets.Insert(a.id.Project, resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating Bucket %s: %w", a.id.Bucket, err)
	}
	log.V(2).Info("successfully created Bucket", "name", a.id.Bucket)

	status := StorageBucketStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BucketAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Bucket", "name", a.id.Bucket)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := StorageBucketSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.Bucket

	// JSON API Update uses PUT (replaces everything), Patch uses PATCH (merges).
	// Most direct controllers use Patch for better compatibility.
	updated, err := a.gcpClient.Buckets.Patch(a.id.Bucket, resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating Bucket %s: %w", a.id.Bucket, err)
	}
	log.V(2).Info("successfully updated Bucket", "name", a.id.Bucket)

	status := StorageBucketStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BucketAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.StorageBucket{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(StorageBucketSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Bucket)
	u.SetGroupVersionKind(krm.StorageBucketGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BucketAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Bucket", "name", a.id.Bucket)

	err := a.gcpClient.Buckets.Delete(a.id.Bucket).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Bucket, assuming it was already deleted", "name", a.id.Bucket)
			return true, nil
		}
		return false, fmt.Errorf("deleting Bucket %s: %w", a.id.Bucket, err)
	}
	log.V(2).Info("successfully deleted Bucket", "name", a.id.Bucket)

	return true, nil
}
