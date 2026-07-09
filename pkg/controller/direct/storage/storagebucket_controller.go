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

// +tool:controller
// proto.service: google.storage.v1.Storage
// proto.message: google.storage.v1.Bucket
// crd.type: StorageBucket
// crd.version: v1beta1

package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/api/option"
	gcs "google.golang.org/api/storage/v1"
	pb "google.golang.org/genproto/googleapis/storage/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.StorageBucketGVK, NewStorageBucketModel)
}

func NewStorageBucketModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &storageBucketModel{config: *config}, nil
}

var _ directbase.Model = &storageBucketModel{}

type storageBucketModel struct {
	config config.ControllerConfig
}

func (m *storageBucketModel) client(ctx context.Context) (*gcs.Service, error) {
	var opts []option.ClientOption
	var err error
	opts, err = m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcs.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building GCS REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *storageBucketModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.StorageBucket{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Resolve identity (bucket name and project)
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, err
	}
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	id := &krm.StorageBucketIdentity{
		Project: projectID,
		Bucket:  resourceID,
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := StorageBucketSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set name on desiredPb
	desiredPb.Name = id.Bucket
	// Support labels
	desiredPb.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &storageBucketAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desiredPb,
		reader:    reader,
	}, nil
}

func (m *storageBucketModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type storageBucketAdapter struct {
	gcpClient *gcs.Service
	id        *krm.StorageBucketIdentity
	desired   *pb.Bucket
	actual    *pb.Bucket
	reader    client.Reader
}

var _ directbase.Adapter = &storageBucketAdapter{}

func protoToREST(in *pb.Bucket) (*gcs.Bucket, error) {
	if in == nil {
		return nil, nil
	}
	data, err := protojson.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := &gcs.Bucket{}
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func restToProto(in *gcs.Bucket) (*pb.Bucket, error) {
	if in == nil {
		return nil, nil
	}
	data, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := &pb.Bucket{}
	unmarshalOpts := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := unmarshalOpts.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (a *storageBucketAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting StorageBucket", "name", a.id.Bucket)

	actualRest, err := a.gcpClient.Buckets.Get(a.id.Bucket).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting StorageBucket %s: %w", a.id.Bucket, err)
	}

	actual, err := restToProto(actualRest)
	if err != nil {
		return false, err
	}

	a.actual = actual
	return true, nil
}

func (a *storageBucketAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating StorageBucket", "name", a.id.Bucket)

	reqRest, err := protoToREST(a.desired)
	if err != nil {
		return err
	}

	createdRest, err := a.gcpClient.Buckets.Insert(a.id.Project, reqRest).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating StorageBucket %s: %w", a.id.Bucket, err)
	}

	created, err := restToProto(createdRest)
	if err != nil {
		return err
	}

	log.V(2).Info("successfully created StorageBucket", "name", a.id.Bucket)
	return a.updateStatus(ctx, createOp, created)
}

func (a *storageBucketAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating StorageBucket", "name", a.id.Bucket)

	diffs, _, err := compareBucket(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		reqRest, err := protoToREST(a.desired)
		if err != nil {
			return err
		}

		updatedRest, err := a.gcpClient.Buckets.Patch(a.id.Bucket, reqRest).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating StorageBucket %s: %w", a.id.Bucket, err)
		}

		updated, err := restToProto(updatedRest)
		if err != nil {
			return err
		}
		latest = updated
		log.V(2).Info("successfully updated StorageBucket", "name", a.id.Bucket)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *storageBucketAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Bucket)
	u.SetGroupVersionKind(krm.StorageBucketGVK)
	u.Object = uObj
	return u, nil
}

func (a *storageBucketAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting StorageBucket", "name", a.id.Bucket)

	err := a.gcpClient.Buckets.Delete(a.id.Bucket).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent StorageBucket, assuming it was already deleted", "name", a.id.Bucket)
			return true, nil
		}
		return false, fmt.Errorf("deleting StorageBucket %s: %w", a.id.Bucket, err)
	}

	log.V(2).Info("successfully deleted StorageBucket", "name", a.id.Bucket)
	return true, nil
}

func (a *storageBucketAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Bucket) error {
	mapCtx := &direct.MapContext{}
	status := &krm.StorageBucketStatus{}

	if latest.Name != "" {
		status.Url = direct.LazyPtr("gs://" + latest.Name)
		status.SelfLink = direct.LazyPtr("https://www.googleapis.com/storage/v1/b/" + latest.Name)
	}

	status.ObservedState = &krm.StorageBucketObservedState{}

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareBucket(ctx context.Context, actual, desired *pb.Bucket) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *pb.Bucket
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := StorageBucketSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = StorageBucketSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
