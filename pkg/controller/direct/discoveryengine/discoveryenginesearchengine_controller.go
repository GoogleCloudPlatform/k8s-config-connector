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
// proto.service: google.cloud.discoveryengine.v1.SiteSearchEngineService
// proto.message: google.cloud.discoveryengine.v1.SiteSearchEngine
// crd.type: DiscoveryEngineSearchEngine
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineSearchEngineGVK, NewSearchEngineModel, registry.CannotBeDeleted())
}

func NewSearchEngineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &searchEngineModel{config: *config}, nil
}

var _ directbase.Model = &searchEngineModel{}

type searchEngineModel struct {
	config config.ControllerConfig
}

func (m *searchEngineModel) client(ctx context.Context, projectID string) (*gcp.SiteSearchEngineClient, error) {
	var opts []option.ClientOption

	config := m.config

	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewSiteSearchEngineRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine site search client: %w", err)
	}

	return gcpClient, err
}

func (m *searchEngineModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineSearchEngine{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DiscoveryEngineSearchEngineIdentity)

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineSearchEngineSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &searchEngineAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *searchEngineModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//discoveryengine.googleapis.com/")
		id := &krm.DiscoveryEngineSearchEngineIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineSearchEngine format", "url", url)
			return nil, nil
		}
		gcpClient, err := m.client(ctx, id.Project)
		if err != nil {
			return nil, err
		}
		return &searchEngineAdapter{
			gcpClient: gcpClient,
			id:        id,
		}, nil
	}
	return nil, nil
}

type searchEngineAdapter struct {
	gcpClient *gcp.SiteSearchEngineClient
	id        *krm.DiscoveryEngineSearchEngineIdentity
	desired   *pb.SiteSearchEngine
	actual    *pb.SiteSearchEngine
}

var _ directbase.Adapter = &searchEngineAdapter{}

func (a *searchEngineAdapter) fullyQualifiedName() string {
	return a.id.String()
}

func (a *searchEngineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.fullyQualifiedName()
	log.V(2).Info("getting discoveryengine search engine", "name", fqn)

	req := &pb.GetSiteSearchEngineRequest{Name: fqn}
	actual, err := a.gcpClient.GetSiteSearchEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine search engine %q from gcp: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *searchEngineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating/getting discoveryengine search engine", "name", a.id)

	fqn := a.fullyQualifiedName()
	req := &pb.GetSiteSearchEngineRequest{Name: fqn}
	actual, err := a.gcpClient.GetSiteSearchEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("getting discoveryengine search engine %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully retrieved discoveryengine search engine in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, actual)
}

func (a *searchEngineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine search engine", "name", a.id)

	diffs, _, err := compareSearchEngine(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func compareSearchEngine(ctx context.Context, actual, desired *pb.SiteSearchEngine) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineSearchEngineSpec_v1alpha1_FromProto, DiscoveryEngineSearchEngineSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore name if needed

	clonedDesired := proto.CloneOf(desired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *searchEngineAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SiteSearchEngine) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DiscoveryEngineSearchEngineStatus{}
	status.ObservedState = DiscoveryEngineSearchEngineObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *searchEngineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineSearchEngine{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineSearchEngineSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{
		External: a.id.ParentString(),
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Datastore)
	u.SetGroupVersionKind(krm.DiscoveryEngineSearchEngineGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *searchEngineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine search engine (no-op)", "name", a.id)
	return true, nil
}
