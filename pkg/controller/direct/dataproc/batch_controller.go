// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:controller
// proto.service: google.cloud.dataproc.v1.BatchController
// proto.message: google.cloud.dataproc.v1.Batch
// crd.type: DataprocBatch
// crd.version: v1alpha1

package dataproc

import (
	"context"
	"fmt"
	"strings"

	dataproc "cloud.google.com/go/dataproc/apiv1"
	dataprocpb "cloud.google.com/go/dataproc/apiv1/dataprocpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/golang/protobuf/proto"
)

func init() {
	registry.RegisterModel(krm.DataprocBatchGVK, NewBatchModel)
}

func NewBatchModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &batchModel{config: *config}, nil
}

var _ directbase.Model = &batchModel{}

type batchModel struct {
	config config.ControllerConfig
}

func (m *batchModel) Client(ctx context.Context, projectID string) (*dataproc.BatchControllerClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := dataproc.NewBatchControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc batch client: %w", err)
	}

	return gcpClient, err
}

func (m *batchModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DataprocBatch{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBatchIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if obj.Spec.RuntimeConfig != nil {
		if obj.Spec.RuntimeConfig.Version != nil {
			if err := obj.Spec.RuntimeConfig.Version.ResolveValue(ctx, reader, obj.Spec.RuntimeConfig, obj, "version"); err != nil {
				return nil, fmt.Errorf("normalizing field \"version\": %w", err)
			}
		}
	}

	gcpClient, err := m.Client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &batchAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *batchModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//dataproc.googleapis.com/") {
		id, err := krm.ParseBatchExternal(url)
		if err != nil {
			log.V(2).Error(err, "url did not match DataprocBatch format", "url", url)
		} else {
			gcpClient, err := m.Client(ctx, id.Parent().ProjectID)
			if err != nil {
				return nil, err
			}
			return &batchAdapter{
				gcpClient: gcpClient,
				id:        id,
			}, nil
		}
	}
	return nil, nil
}

type batchAdapter struct {
	gcpClient *dataproc.BatchControllerClient
	id        *krm.BatchIdentity
	desired   *krm.DataprocBatch
	actual    *dataprocpb.Batch
}

var _ directbase.Adapter = &batchAdapter{}

func (a *batchAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting dataproc batch", "name", a.id)

	req := &dataprocpb.GetBatchRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetBatch(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataproc batch %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *batchAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating dataproc batch", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataprocBatchSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dataprocpb.CreateBatchRequest{
		Parent:  a.id.Parent().String(),
		Batch:   resource,
		BatchId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateBatch(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataproc batch %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("dataproc batch %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created dataproc batch in gcp", "name", a.id)

	status := &krm.DataprocBatchStatus{}
	status.ObservedState = DataprocBatchObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// DataprocBatch does not support update.
func (a *batchAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating dataproc batch", "name", a.id)

	desiredpb := DataprocBatchSpec_ToProto(&direct.MapContext{}, &a.desired.Spec)
	paths, err := common.CompareProtoMessage(desiredpb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) != 0 {
		log.V(2).Info("This resource does not support update", "name", a.id.String())
		return nil
	}

	status := &krm.DataprocBatchStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *batchAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting dataproc batch", "name", a.id)

	req := &dataprocpb.DeleteBatchRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteBatch(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting dataproc batch %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted dataproc batch", "name", a.id)

	return true, nil
}

// Export implements the Adapter interface.
func (a *batchAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataprocBatch{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataprocBatchSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	if a.actual.EnvironmentConfig != nil {
		obj.Spec.RuntimeConfig.Version = &krm.RuntimeConfigVersion{}
		obj.Spec.RuntimeConfig.Version.Value = proto.String(a.actual.EnvironmentConfig.ExecutionConfig.KmsKey)
	}

	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = direct.LazyPtr(a.id.Parent().Location)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.DataprocBatchGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

```
</out>


