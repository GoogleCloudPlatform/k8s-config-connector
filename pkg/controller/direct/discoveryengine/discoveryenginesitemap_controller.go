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
// proto.message: google.cloud.discoveryengine.v1.Sitemap
// crd.type: DiscoveryEngineSitemap
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
	registry.RegisterModel(krm.DiscoveryEngineSitemapGVK, NewSitemapModel)
}

func NewSitemapModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sitemapModel{config: *config}, nil
}

var _ directbase.Model = &sitemapModel{}

type sitemapModel struct {
	config config.ControllerConfig
}

func (m *sitemapModel) client(ctx context.Context, projectID string) (*gcp.SiteSearchEngineClient, error) {
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
		return nil, fmt.Errorf("building discoveryengine site search engine client: %w", err)
	}

	return gcpClient, err
}

func (m *sitemapModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineSitemap{}
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
	id := identity.(*krm.DiscoveryEngineSitemapIdentity)

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineSitemapSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &sitemapAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *sitemapModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//discoveryengine.googleapis.com/")
		id := &krm.DiscoveryEngineSitemapIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineSitemap format", "url", url)
			return nil, nil
		}
		gcpClient, err := m.client(ctx, id.Project)
		if err != nil {
			return nil, err
		}
		return &sitemapAdapter{
			gcpClient: gcpClient,
			id:        id,
		}, nil
	}
	return nil, nil
}

type sitemapAdapter struct {
	gcpClient *gcp.SiteSearchEngineClient
	id        *krm.DiscoveryEngineSitemapIdentity
	desired   *pb.Sitemap
	actual    *pb.Sitemap
}

var _ directbase.Adapter = &sitemapAdapter{}

func (a *sitemapAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine sitemap", "name", a.id)

	if a.id.Sitemap == "" {
		return false, nil
	}

	req := &pb.FetchSitemapsRequest{
		Parent: a.id.ParentString(),
	}
	resp, err := a.gcpClient.FetchSitemaps(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("fetching discoveryengine sitemaps for parent %q: %w", a.id.ParentString(), err)
	}

	for _, meta := range resp.GetSitemapsMetadata() {
		s := meta.GetSitemap()
		if s == nil {
			continue
		}
		id := &krm.DiscoveryEngineSitemapIdentity{}
		if err := id.FromExternal(s.GetName()); err == nil {
			if id.Sitemap == a.id.Sitemap {
				a.actual = s
				return true, nil
			}
		}
	}

	return false, nil
}

func (a *sitemapAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine sitemap", "parent", a.id.ParentString())

	desired := proto.Clone(a.desired).(*pb.Sitemap)
	desired.Name = ""

	req := &pb.CreateSitemapRequest{
		Parent:  a.id.ParentString(),
		Sitemap: desired,
	}
	op, err := a.gcpClient.CreateSitemap(ctx, req)
	if err != nil && strings.Contains(err.Error(), "Only Advanced Site Search data stores are permitted") {
		log.V(2).Info("enabling advanced site search for data store", "parent", a.id.ParentString())
		enableOp, enableErr := a.gcpClient.EnableAdvancedSiteSearch(ctx, &pb.EnableAdvancedSiteSearchRequest{
			SiteSearchEngine: a.id.ParentString(),
		})
		if enableErr != nil {
			return fmt.Errorf("enabling advanced site search on %s: %w", a.id.ParentString(), enableErr)
		}
		if _, enableErr := enableOp.Wait(ctx); enableErr != nil {
			return fmt.Errorf("waiting for advanced site search enablement on %s: %w", a.id.ParentString(), enableErr)
		}
		op, err = a.gcpClient.CreateSitemap(ctx, req)
	}
	if err != nil {
		return fmt.Errorf("creating discoveryengine sitemap %s: %w", a.id.ParentString(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for discoveryengine sitemap creation: %w", err)
	}
	log.V(2).Info("successfully created discoveryengine sitemap in gcp", "name", created.GetName())

	createdID := &krm.DiscoveryEngineSitemapIdentity{}
	if err := createdID.FromExternal(created.GetName()); err != nil {
		return fmt.Errorf("parsing created sitemap name %q: %w", created.GetName(), err)
	}
	a.id.Sitemap = createdID.Sitemap

	return a.updateStatus(ctx, createOp, created)
}

func (a *sitemapAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine sitemap", "name", a.id)

	desired := proto.Clone(a.desired).(*pb.Sitemap)
	desired.Name = a.id.String()

	diffs, _, err := a.compare(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("DiscoveryEngineSitemap is immutable and cannot be updated")
}

func (a *sitemapAdapter) compare(ctx context.Context, actual, desired *pb.Sitemap) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineSitemapSpec_v1alpha1_FromProto, DiscoveryEngineSitemapSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Sitemap)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	return diffs, updateMask, nil
}

func (a *sitemapAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Sitemap) error {
	status := &krm.DiscoveryEngineSitemapStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineSitemapObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *sitemapAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineSitemap{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineSitemapSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{
		External: fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", a.id.Project, a.id.Location, a.id.Collection, a.id.DataStore),
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Sitemap)
	u.SetGroupVersionKind(krm.DiscoveryEngineSitemapGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *sitemapAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine sitemap", "name", a.id)

	if a.id.Sitemap == "" {
		return true, nil
	}

	found, err := a.Find(ctx)
	if err != nil {
		return false, err
	}
	if !found || a.actual == nil {
		return true, nil
	}

	req := &pb.DeleteSitemapRequest{
		Name: a.actual.GetName(),
	}
	op, err := a.gcpClient.DeleteSitemap(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting discoveryengine sitemap %s: %w", a.actual.GetName(), err)
	}

	if err := op.Wait(ctx); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for discoveryengine sitemap deletion %s: %w", a.actual.GetName(), err)
	}

	log.V(2).Info("successfully deleted discoveryengine sitemap", "name", a.id)
	return true, nil
}
