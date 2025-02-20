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

package logging

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

        loggingapiv2 "cloud.google.com/go/logging/apiv2"
	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.LoggingLogBucketGVK, NewLogBucketModel)
}

func NewLogBucketModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLogBucket{config: *config}, nil
}

var _ directbase.Model = &modelLogBucket{}

type modelLogBucket struct {
	config config.ControllerConfig
}

func (m *modelLogBucket) client(ctx context.Context) (*loggingapiv2.ConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := loggingapiv2.NewConfigClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building logbucket client: %w", err)
	}
	return client, nil
}

func (m *modelLogBucket) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.LoggingLogBucket{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewLogBucketIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get logging GCP client
	client, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &LogBucketAdapter{
		id:        id,
		client: client,
		desired:   obj,
	}, nil
}

func (m *modelLogBucket) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LogBucketAdapter struct {
        id        *krm.LogBucketIdentity
        client *loggingapiv2.ConfigClient
        desired   *krm.LoggingLogBucket
        actual    *loggingpb.LogBucket
}

var _ directbase.Adapter = &LogBucketAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *LogBucketAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LogBucket", "name", a.id)

	req := &loggingpb.GetLogBucketRequest{Name: a.id.String()}
	logbucket, err := a.client.GetBucket(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LogBucket %q: %w", a.id, err)
	}

	a.actual = logbucket
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LogBucketAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LogBucket", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := LoggingLogBucketSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &loggingpb.CreateLogBucketRequest{
		Parent:    a.id.Parent().String(),
		LogBucket: resource,
		BucketId:  a.id.ID,
	}
	created, err := a.client.CreateBucket(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LogBucket %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created LogBucket", "name", a.id)

	status := &krm.LoggingLogBucketStatus{}
	status.ObservedState = LoggingLogBucketObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)

}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LogBucketAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LogBucket", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := LoggingLogBucketSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
        if err != nil {
               return err
        }
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.LoggingLogBucketStatus{}
		status.ObservedState = LoggingLogBucketObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	req := &loggingpb.UpdateLogBucketRequest{
		Name:       a.id.String(),
		UpdateMask: updateMask,
		LogBucket:  desiredPb,
	}
	updated, err := a.client.UpdateBucket(ctx, req)
	if err != nil {
		return fmt.Errorf("updating LogBucket %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated LogBucket", "name", a.id)

	status := &krm.LoggingLogBucketStatus{}
	status.ObservedState = LoggingLogBucketObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *LogBucketAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLogBucket{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLogBucketSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.LoggingLogBucketGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *LogBucketAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LogBucket", "name", a.id)

	req := &loggingpb.DeleteLogBucketRequest{Name: a.id.String()}
	err := a.client.DeleteBucket(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent LogBucket, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LogBucket %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LogBucket", "name", a.id)
	return true, nil
}
