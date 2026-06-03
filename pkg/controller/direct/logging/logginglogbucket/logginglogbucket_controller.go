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
// proto.service: google.logging.v2.ConfigServiceV2
// proto.message: google.logging.v2.LogBucket
// crd.type: LoggingLogBucket
// crd.version: v1beta1

package logginglogbucket

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/logging/apiv2"
	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.LoggingLogBucketGVK, NewLoggingLogBucketModel)
}

func NewLoggingLogBucketModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLoggingLogBucket{config: *config}, nil
}

var _ directbase.Model = &modelLoggingLogBucket{}

type modelLoggingLogBucket struct {
	config config.ControllerConfig
}

func (m *modelLoggingLogBucket) client(ctx context.Context) (*gcp.ConfigClient, error) {
	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("building authenticated HTTP client: %w", err)
	}

	gcpClient, err := gcp.NewConfigRESTClient(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("building Logging Config client: %w", err)
	}
	return gcpClient, err
}

func (m *modelLoggingLogBucket) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LoggingLogBucket{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &LoggingLogBucketAdapter{
		id:        id.(*krm.LogBucketIdentity),
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelLoggingLogBucket) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LoggingLogBucketAdapter struct {
	id        *krm.LogBucketIdentity
	gcpClient *gcp.ConfigClient
	desired   *krm.LoggingLogBucket
	actual    *loggingpb.LogBucket
	reader    client.Reader
}

var _ directbase.Adapter = &LoggingLogBucketAdapter{}

func (a *LoggingLogBucketAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLogBucket", "name", a.id)

	req := &loggingpb.GetBucketRequest{Name: a.id.String()}
	bucket, err := a.gcpClient.GetBucket(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLogBucket %q: %w", a.id, err)
	}

	a.actual = bucket
	return true, nil
}

func (a *LoggingLogBucketAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LoggingLogBucket", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	if err := ResolveLoggingLogBucketRefs(ctx, a.reader, desired); err != nil {
		return err
	}
	resource := LoggingLogBucketSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := parentOfLogBucket(a.id)
	bucketID := a.id.ID()

	req := &loggingpb.CreateBucketRequest{
		Parent:   parent,
		BucketId: bucketID,
		Bucket:   resource,
	}
	created, err := a.gcpClient.CreateBucket(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LoggingLogBucket %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created LoggingLogBucket", "name", a.id)

	status := LoggingLogBucketStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLogBucketAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LoggingLogBucket", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	if err := ResolveLoggingLogBucketRefs(ctx, a.reader, desired); err != nil {
		return err
	}
	resource := LoggingLogBucketSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := &fieldmaskpb.FieldMask{
		Paths: []string{
			"description",
			"retention_days",
			"locked",
			"analytics_enabled",
			"restricted_fields",
			"index_configs",
			"cmek_settings",
		},
	}

	req := &loggingpb.UpdateBucketRequest{
		Name:       a.id.String(),
		Bucket:     resource,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateBucket(ctx, req)
	if err != nil {
		return fmt.Errorf("updating LoggingLogBucket %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated LoggingLogBucket", "name", a.id)

	status := LoggingLogBucketStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLogBucketAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingLogBucketGVK)

	u.Object = uObj
	return u, nil
}

func (a *LoggingLogBucketAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LoggingLogBucket", "name", a.id)

	req := &loggingpb.DeleteBucketRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteBucket(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LoggingLogBucket, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting LoggingLogBucket %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LoggingLogBucket", "name", a.id)
	return true, nil
}

func parentOfLogBucket(i *krm.LogBucketIdentity) string {
	if i.Project != "" {
		return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
	}
	if i.Folder != "" {
		return fmt.Sprintf("folders/%s/locations/%s", i.Folder, i.Location)
	}
	if i.Organization != "" {
		return fmt.Sprintf("organizations/%s/locations/%s", i.Organization, i.Location)
	}
	if i.BillingAccount != "" {
		return fmt.Sprintf("billingAccounts/%s/locations/%s", i.BillingAccount, i.Location)
	}
	if i.AccessPolicy != "" {
		return fmt.Sprintf("accessPolicies/%s/locations/%s", i.AccessPolicy, i.Location)
	}
	return ""
}
