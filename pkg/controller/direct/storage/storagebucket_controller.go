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
// proto.service: google.storage.v1
// proto.message: google.storage.v1.Bucket
// crd.type: StorageBucket
// crd.version: v1beta1

package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/api/option"
	gcpstorage "google.golang.org/api/storage/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	pb "google.golang.org/genproto/googleapis/storage/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func (m *modelBucket) client(ctx context.Context) (*gcpstorage.Service, error) {
	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating GCS HTTP client: %w", err)
	}

	gcpClient, err := gcpstorage.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("building GCS REST client: %w", err)
	}

	return gcpClient, nil
}

func (m *modelBucket) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.StorageBucket{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	bucketId := id.(*krm.StorageBucketIdentity)

	mapCtx := &direct.MapContext{}
	desired := StorageBucketSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set Labels from k8s labels if available in the desired state
	desired.Labels = label.GCPLabels(obj)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &BucketAdapter{
		id:         bucketId,
		gcpClient:  gcpClient,
		desiredObj: obj,
		desired:    desired,
		reader:     reader,
	}, nil
}

func (m *modelBucket) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BucketAdapter struct {
	id                     *krm.StorageBucketIdentity
	gcpClient              *gcpstorage.Service
	desiredObj             *krm.StorageBucket
	desired                *pb.Bucket
	actual                 *pb.Bucket
	actualSoftDeletePolicy *gcpstorage.BucketSoftDeletePolicy
	reader                 client.Reader
}

var _ directbase.Adapter = &BucketAdapter{}

func protoToRESTBucket(in *pb.Bucket) (*gcpstorage.Bucket, error) {
	if in == nil {
		return nil, nil
	}
	m := protojson.MarshalOptions{
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}
	data, err := m.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := &gcpstorage.Bucket{}
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}

	if in.IamConfiguration != nil && in.IamConfiguration.UniformBucketLevelAccess != nil {
		out.IamConfiguration.UniformBucketLevelAccess.ForceSendFields = []string{"Enabled"}
	}
	if in.Billing != nil {
		out.Billing.ForceSendFields = []string{"RequesterPays"}
	}
	if in.Versioning != nil {
		out.Versioning.ForceSendFields = []string{"Enabled"}
	}

	return out, nil
}

func restToProtoBucket(in *gcpstorage.Bucket) (*pb.Bucket, error) {
	if in == nil {
		return nil, nil
	}
	data, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	out := &pb.Bucket{}
	u := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := u.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (a *BucketAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting StorageBucket", "name", a.id)

	bucket, err := a.gcpClient.Buckets.Get(a.id.Bucket).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting StorageBucket %q: %w", a.id, err)
	}

	actual, err := restToProtoBucket(bucket)
	if err != nil {
		return false, err
	}

	a.actual = actual
	a.actualSoftDeletePolicy = bucket.SoftDeletePolicy
	return true, nil
}

func (a *BucketAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating StorageBucket", "name", a.id)

	restBucket, err := protoToRESTBucket(a.desired)
	if err != nil {
		return err
	}

	restBucket.Name = a.id.Bucket

	if a.desiredObj.Spec.DefaultEventBasedHold != nil {
		restBucket.ForceSendFields = append(restBucket.ForceSendFields, "DefaultEventBasedHold")
	}

	if a.desiredObj.Spec.SoftDeletePolicy != nil {
		restBucket.SoftDeletePolicy = &gcpstorage.BucketSoftDeletePolicy{}
		if a.desiredObj.Spec.SoftDeletePolicy.RetentionDurationSeconds != nil {
			restBucket.SoftDeletePolicy.RetentionDurationSeconds = int64(*a.desiredObj.Spec.SoftDeletePolicy.RetentionDurationSeconds)
			restBucket.SoftDeletePolicy.ForceSendFields = []string{"RetentionDurationSeconds"}
		}
	}

	created, err := a.gcpClient.Buckets.Insert(a.id.Project, restBucket).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating StorageBucket %s: %w", a.id, err)
	}

	actual, err := restToProtoBucket(created)
	if err != nil {
		return err
	}

	a.actual = actual
	a.actualSoftDeletePolicy = created.SoftDeletePolicy
	return a.updateStatus(ctx, createOp, actual, created.SoftDeletePolicy)
}

func (a *BucketAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating StorageBucket", "name", a.id)

	diffs, updateMask, err := compareBucket(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if a.desiredObj.Spec.SoftDeletePolicy != nil && a.desiredObj.Spec.SoftDeletePolicy.RetentionDurationSeconds != nil {
		desiredSeconds := int64(*a.desiredObj.Spec.SoftDeletePolicy.RetentionDurationSeconds)
		var actualSeconds int64
		if a.actualSoftDeletePolicy != nil {
			actualSeconds = a.actualSoftDeletePolicy.RetentionDurationSeconds
		}
		if desiredSeconds != actualSeconds {
			diffs.AddField("softDeletePolicy.retentionDurationSeconds", actualSeconds, desiredSeconds)
		}
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected, skipping update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual, a.actualSoftDeletePolicy)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	log.V(2).Info("updating StorageBucket", "name", a.id, "updateMask", updateMask)

	restBucket, err := protoToRESTBucket(a.desired)
	if err != nil {
		return err
	}

	restBucket.Name = a.id.Bucket

	if a.desiredObj.Spec.DefaultEventBasedHold != nil {
		restBucket.ForceSendFields = append(restBucket.ForceSendFields, "DefaultEventBasedHold")
	}

	if a.desiredObj.Spec.SoftDeletePolicy != nil {
		restBucket.SoftDeletePolicy = &gcpstorage.BucketSoftDeletePolicy{}
		if a.desiredObj.Spec.SoftDeletePolicy.RetentionDurationSeconds != nil {
			restBucket.SoftDeletePolicy.RetentionDurationSeconds = int64(*a.desiredObj.Spec.SoftDeletePolicy.RetentionDurationSeconds)
			restBucket.SoftDeletePolicy.ForceSendFields = []string{"RetentionDurationSeconds"}
		}
	}

	updated, err := a.gcpClient.Buckets.Patch(a.id.Bucket, restBucket).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating StorageBucket %s: %w", a.id, err)
	}

	actual, err := restToProtoBucket(updated)
	if err != nil {
		return err
	}

	a.actual = actual
	a.actualSoftDeletePolicy = updated.SoftDeletePolicy
	return a.updateStatus(ctx, updateOp, actual, updated.SoftDeletePolicy)
}

func (a *BucketAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting StorageBucket", "name", a.id)

	err := a.gcpClient.Buckets.Delete(a.id.Bucket).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting StorageBucket %s: %w", a.id, err)
	}
	return true, nil
}

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

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.StorageBucketGVK)
	return u, nil
}

func (a *BucketAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Bucket, latestSoftDeletePolicy *gcpstorage.BucketSoftDeletePolicy) error {
	mapCtx := &direct.MapContext{}
	status := StorageBucketStatus_FromProto(mapCtx, latest, latestSoftDeletePolicy)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func StorageBucketStatus_FromProto(mapCtx *direct.MapContext, latest *pb.Bucket, latestSoftDeletePolicy *gcpstorage.BucketSoftDeletePolicy) *krm.StorageBucketStatus {
	if latest == nil {
		return nil
	}
	status := &krm.StorageBucketStatus{}
	status.SelfLink = direct.PtrTo(fmt.Sprintf("https://www.googleapis.com/storage/v1/b/%s", latest.Name))
	status.Url = direct.PtrTo(fmt.Sprintf("gs://%s", latest.Name))

	if latestSoftDeletePolicy != nil {
		status.ObservedState = &krm.StorageBucketObservedState{
			SoftDeletePolicy: &krm.StorageBucketSoftDeletePolicyObservedState{},
		}
		if latestSoftDeletePolicy.EffectiveTime != "" {
			status.ObservedState.SoftDeletePolicy.EffectiveTime = &latestSoftDeletePolicy.EffectiveTime
		} else {
			status.ObservedState.SoftDeletePolicy.EffectiveTime = direct.PtrTo("1970-01-01T00:00:00Z")
		}
		status.ObservedState.SoftDeletePolicy.RetentionDurationSeconds = &latestSoftDeletePolicy.RetentionDurationSeconds
	}

	return status
}

func compareBucket(ctx context.Context, actual, desired *pb.Bucket) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, StorageBucketSpec_FromProto, StorageBucketSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Bucket)

	populateDefaults := func(obj *pb.Bucket) {
		if obj.Location == "" {
			obj.Location = "US"
		}
		if obj.StorageClass == "" {
			obj.StorageClass = "STANDARD"
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
