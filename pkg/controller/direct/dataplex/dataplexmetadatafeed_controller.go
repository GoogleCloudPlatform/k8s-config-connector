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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.MetadataFeed
// crd.type: DataplexMetadataFeed
// crd.version: v1alpha1

package dataplex

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/dataplex/apiv1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataplexMetadataFeedGVK, NewMetadataFeedModel)
}

func NewMetadataFeedModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &metadataFeedModel{config: config}, nil
}

var _ directbase.Model = &metadataFeedModel{}

type metadataFeedModel struct {
	config *config.ControllerConfig
}

func (m *metadataFeedModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexMetadataFeed{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	copied := obj.DeepCopy()
	mapCtx := &direct.MapContext{}
	desired := DataplexMetadataFeedSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	if err := normalizeMetadataFeed(ctx, m.config.ProjectMapper, desired); err != nil {
		return nil, fmt.Errorf("normalizing desired metadata feed: %w", err)
	}

	idI, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idI.(*krm.MetadataFeedIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type %T", idI)
	}

	adapter := &metadataFeedAdapter{
		id:            id,
		desired:       desired,
		reader:        reader,
		projectMapper: m.config.ProjectMapper,
	}

	// Get GCP client
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	catalogClient, err := gcpClient.catalogClient(ctx)
	if err != nil {
		return nil, err
	}
	adapter.gcpClient = catalogClient

	return adapter, nil
}

func (m *metadataFeedModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type metadataFeedAdapter struct {
	gcpClient     *gcp.CatalogClient
	id            *krm.MetadataFeedIdentity
	desired       *pb.MetadataFeed
	actual        *pb.MetadataFeed
	reader        client.Reader
	projectMapper *projects.ProjectMapper
}

var _ directbase.Adapter = &metadataFeedAdapter{}

func normalizeMetadataFeed(ctx context.Context, projectMapper *projects.ProjectMapper, feed *pb.MetadataFeed) error {
	if feed == nil || projectMapper == nil {
		return nil
	}
	if feed.Scope != nil {
		for i, p := range feed.Scope.Projects {
			normalized, err := projectMapper.ReplaceProjectNumberWithIDInLink(ctx, p)
			if err != nil {
				return err
			}
			feed.Scope.Projects[i] = normalized
		}
		for i, eg := range feed.Scope.EntryGroups {
			normalized, err := projectMapper.ReplaceProjectNumberWithIDInLink(ctx, eg)
			if err != nil {
				return err
			}
			feed.Scope.EntryGroups[i] = normalized
		}
	}
	if feed.Filters != nil {
		for i, et := range feed.Filters.EntryTypes {
			normalized, err := projectMapper.ReplaceProjectNumberWithIDInLink(ctx, et)
			if err != nil {
				return err
			}
			feed.Filters.EntryTypes[i] = normalized
		}
		for i, at := range feed.Filters.AspectTypes {
			normalized, err := projectMapper.ReplaceProjectNumberWithIDInLink(ctx, at)
			if err != nil {
				return err
			}
			feed.Filters.AspectTypes[i] = normalized
		}
	}
	if feed.GetPubsubTopic() != "" {
		normalized, err := projectMapper.ReplaceProjectNumberWithIDInLink(ctx, feed.GetPubsubTopic())
		if err != nil {
			return err
		}
		feed.Endpoint = &pb.MetadataFeed_PubsubTopic{PubsubTopic: normalized}
	}
	return nil
}

func (a *metadataFeedAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataplex metadatafeed", "name", a.id)

	req := &pb.GetMetadataFeedRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetMetadataFeed(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataplex metadatafeed %q from gcp: %w", a.id.String(), err)
	}

	if err := normalizeMetadataFeed(ctx, a.projectMapper, actual); err != nil {
		return false, fmt.Errorf("normalizing actual metadata feed: %w", err)
	}

	a.actual = actual
	return true, nil
}

func (a *metadataFeedAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataplex metadatafeed", "name", a.id)

	req := &pb.CreateMetadataFeedRequest{
		Parent:         a.id.ParentString(),
		MetadataFeedId: a.id.MetadataFeed,
		MetadataFeed:   a.desired,
	}

	op, err := a.gcpClient.CreateMetadataFeed(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataplex metadatafeed %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting create dataplex metadatafeed %s failed: %w", a.id, err)
	}

	log.V(2).Info("successfully created dataplex metadatafeed in gcp", "name", a.id)

	// Fetch fully-populated resource after creation
	actual, err := a.gcpClient.GetMetadataFeed(ctx, &pb.GetMetadataFeedRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching dataplex metadatafeed %s after create: %w", a.id, err)
	}
	if err := normalizeMetadataFeed(ctx, a.projectMapper, actual); err != nil {
		return fmt.Errorf("normalizing actual metadata feed: %w", err)
	}

	return a.updateStatus(ctx, createOp, actual)
}

func (a *metadataFeedAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataplex metadatafeed", "name", a.id)

	mapCtx := &direct.MapContext{}

	// Mask actual to only contain spec fields for correct diffing
	maskedActualSpec := DataplexMetadataFeedSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := DataplexMetadataFeedSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.MetadataFeed)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if diffs == nil || !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateMetadataFeedRequest{
		MetadataFeed: clonedDesired,
		UpdateMask:   updateMask,
	}
	req.MetadataFeed.Name = a.id.String()

	op, err := a.gcpClient.UpdateMetadataFeed(ctx, req)
	if err != nil {
		return fmt.Errorf("updating dataplex metadatafeed %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of dataplex metadatafeed %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated dataplex metadatafeed", "name", a.id)

	// Fetch fully-populated resource after update
	actual, err := a.gcpClient.GetMetadataFeed(ctx, &pb.GetMetadataFeedRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching dataplex metadatafeed %s after update: %w", a.id, err)
	}
	if err := normalizeMetadataFeed(ctx, a.projectMapper, actual); err != nil {
		return fmt.Errorf("normalizing actual metadata feed: %w", err)
	}

	return a.updateStatus(ctx, updateOp, actual)
}

func (a *metadataFeedAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MetadataFeed) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DataplexMetadataFeedStatus{}
	status.ObservedState = DataplexMetadataFeedObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *metadataFeedAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataplexMetadataFeed{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexMetadataFeedSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.MetadataFeed)
	u.SetGroupVersionKind(krm.DataplexMetadataFeedGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *metadataFeedAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataplex metadatafeed", "name", a.id)

	req := &pb.DeleteMetadataFeedRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteMetadataFeed(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent dataplex metadatafeed, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting dataplex metadatafeed %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of dataplex metadatafeed %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataplex metadatafeed", "name", a.id)

	return true, nil
}
