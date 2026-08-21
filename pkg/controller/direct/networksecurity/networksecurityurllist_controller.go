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

package networksecurity

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/networksecurity/apiv1"
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityURLListGVK, NewURLListModel)
}

func NewURLListModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &urlListModel{config: *config}, nil
}

var _ directbase.Model = &urlListModel{}

type urlListModel struct {
	config config.ControllerConfig
}

func (m *urlListModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networksecurity client: %w", err)
	}

	return gcpClient, nil
}

func (m *urlListModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityURLList{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}
	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityURLListSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &urlList_Adapter{
		gcpClient: gcpClient,
		id:        id.(*krm.NetworkSecurityURLListIdentity),
		desired:   desired,
	}, nil
}

func (m *urlListModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type urlList_Adapter struct {
	gcpClient *gcp.Client
	id        *krm.NetworkSecurityURLListIdentity
	desired   *pb.UrlList
	actual    *pb.UrlList
}

var _ directbase.Adapter = &urlList_Adapter{}

func (a *urlList_Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting networksecurity url list", "name", a.id)

	req := &pb.GetUrlListRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetUrlList(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity url list %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *urlList_Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating networksecurity url list", "name", a.id)

	parent := a.id.ParentString()
	req := &pb.CreateUrlListRequest{
		Parent:    parent,
		UrlListId: a.id.UrlList,
		UrlList:   a.desired,
	}
	op, err := a.gcpClient.CreateUrlList(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networksecurity url list %s: %w", a.id.String(), err)
	}

	actual, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("networksecurity url list %s waiting for creation: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created networksecurity url list", "name", a.id.String())

	return a.updateStatus(ctx, createOp, actual)
}

func (a *urlList_Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating networksecurity url list", "name", a.id)

	diffs, updateMask, err := compareURLList(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	log.V(2).Info("fields need update", "name", a.id, "updateMask", updateMask)

	req := &pb.UpdateUrlListRequest{
		UpdateMask: updateMask,
		UrlList:    a.desired,
	}
	req.UrlList.Name = a.id.String()

	op, err := a.gcpClient.UpdateUrlList(ctx, req)
	if err != nil {
		return fmt.Errorf("updating networksecurity url list %s: %w", a.id, err)
	}
	latest, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("networksecurity url list %s waiting update: %w", a.id, err)
	}

	log.V(2).Info("successfully updated networksecurity url list", "name", a.id)

	return a.updateStatus(ctx, updateOp, latest)
}

func compareURLList(ctx context.Context, actual, desired *pb.UrlList) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityURLListSpec_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := NetworkSecurityURLListSpec_v1alpha1_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *urlList_Adapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.UrlList) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityURLListStatus{}
	status.ObservedState = NetworkSecurityURLListObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *urlList_Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *urlList_Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting networksecurity url list", "name", a.id)

	req := &pb.DeleteUrlListRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteUrlList(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting networksecurity url list %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("networksecurity url list %s waiting for deletion: %w", a.id.String(), err)
	}

	return true, nil
}
