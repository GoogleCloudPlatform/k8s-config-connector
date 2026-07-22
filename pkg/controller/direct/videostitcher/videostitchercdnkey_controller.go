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

package videostitcher

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/video/stitcher/apiv1"
	pb "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/videostitcher/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VideoStitcherCDNKeyGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.VideoStitcherClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewVideoStitcherClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building VideoStitcher client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VideoStitcherCDNKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.VideoStitcherCDNKeyIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := VideoStitcherCDNKeySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		model:     m,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.VideoStitcherCDNKeyIdentity
	gcpClient *gcp.VideoStitcherClient
	desired   *pb.CdnKey
	actual    *pb.CdnKey
	model     *model
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding VideoStitcherCDNKey", "id", a.id)

	req := &pb.GetCdnKeyRequest{
		Name: a.id.String(),
	}
	cdnKey, err := a.gcpClient.GetCdnKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VideoStitcherCDNKey %s: %w", a.id.String(), err)
	}

	a.actual = cdnKey
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VideoStitcherCDNKey", "id", a.id)

	req := &pb.CreateCdnKeyRequest{
		Parent:   a.id.ParentString(),
		CdnKeyId: a.id.CdnKey,
		CdnKey:   a.desired,
	}
	op, err := a.gcpClient.CreateCdnKey(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VideoStitcherCDNKey %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for VideoStitcherCDNKey %s creation: %w", a.id.String(), err)
	}

	// Fetch fully-populated resource after creation
	refetched, err := a.gcpClient.GetCdnKey(ctx, &pb.GetCdnKeyRequest{Name: a.id.String()})
	if err != nil {
		refetched = created
	}
	a.actual = refetched

	return a.updateStatus(ctx, createOp, refetched)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VideoStitcherCDNKey", "id", a.id)

	diffs, updateMask, err := a.compareCDNKey(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.CloneOf(a.desired)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdateCdnKeyRequest{
			CdnKey:     desiredCopy,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdateCdnKey(ctx, req)
		if err != nil {
			return fmt.Errorf("updating VideoStitcherCDNKey %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting VideoStitcherCDNKey %s update: %w", a.id.String(), err)
		}

		// Fetch fully-populated resource after update
		refetched, err := a.gcpClient.GetCdnKey(ctx, &pb.GetCdnKeyRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("fetching VideoStitcherCDNKey %s after update: %w", a.id.String(), err)
		}
		latest = refetched
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CdnKey) error {
	mapCtx := &direct.MapContext{}
	status := krm.VideoStitcherCDNKeyStatus{}
	status.ObservedState = VideoStitcherCDNKeyObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, &status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VideoStitcherCDNKey{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VideoStitcherCDNKeySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.CdnKey)
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.CdnKey)
	u.SetGroupVersionKind(krm.VideoStitcherCDNKeyGVK)

	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VideoStitcherCDNKey", "id", a.id)

	req := &pb.DeleteCdnKeyRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCdnKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent VideoStitcherCDNKey, assuming it was already deleted", "id", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting VideoStitcherCDNKey %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete VideoStitcherCDNKey %s: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *Adapter) compareCDNKey(ctx context.Context, actual, desired *pb.CdnKey) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VideoStitcherCDNKeySpec_FromProto, VideoStitcherCDNKeySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	// Since private keys and token keys are write-only fields and not returned by GCP,
	// we preserve the desired secrets in maskedActual to avoid false drifts.
	if desired.GetGoogleCdnKey() != nil && maskedActual.GetGoogleCdnKey() != nil {
		maskedActual.GetGoogleCdnKey().PrivateKey = desired.GetGoogleCdnKey().PrivateKey
	}
	if desired.GetMediaCdnKey() != nil && maskedActual.GetMediaCdnKey() != nil {
		maskedActual.GetMediaCdnKey().PrivateKey = desired.GetMediaCdnKey().PrivateKey
	}
	if desired.GetAkamaiCdnKey() != nil && maskedActual.GetAkamaiCdnKey() != nil {
		maskedActual.GetAkamaiCdnKey().TokenKey = desired.GetAkamaiCdnKey().TokenKey
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
