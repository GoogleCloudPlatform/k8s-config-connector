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

package appengine

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	api "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.AppEngineFlexibleAppVersionGVK, newFlexibleAppVersionModel)
}

func newFlexibleAppVersionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &flexibleAppVersionModel{config: config}, nil
}

type flexibleAppVersionModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &flexibleAppVersionModel{}

func (m *flexibleAppVersionModel) client(ctx context.Context) (*api.VersionsClient, *api.ServicesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, nil, err
	}

	versionsClient, err := api.NewVersionsRESTClient(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("building appengine versions client: %w", err)
	}

	servicesClient, err := api.NewServicesRESTClient(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("building appengine services client: %w", err)
	}

	return versionsClient, servicesClient, nil
}

func (m *flexibleAppVersionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AppEngineFlexibleAppVersion{}
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
	versionID := id.(*krm.AppEngineFlexibleAppVersionIdentity)

	versionsClient, servicesClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &flexibleAppVersionAdapter{
		versionsClient: versionsClient,
		servicesClient: servicesClient,
		id:             versionID,
		desired:        obj,
	}, nil
}

func (m *flexibleAppVersionModel) AdapterForURL(ctx context.Context, rawURL string) (directbase.Adapter, error) {
	parsedURL, err := url.Parse(rawURL)
	if err == nil && parsedURL.Path != "" {
		rawURL = strings.TrimPrefix(parsedURL.Path, "/v1/")
		rawURL = strings.TrimPrefix(rawURL, "/")
	}
	id := &krm.AppEngineFlexibleAppVersionIdentity{}
	if err := id.FromExternal(rawURL); err != nil {
		return nil, nil
	}

	versionsClient, servicesClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &flexibleAppVersionAdapter{
		versionsClient: versionsClient,
		servicesClient: servicesClient,
		id:             id,
	}, nil
}

type flexibleAppVersionAdapter struct {
	id             *krm.AppEngineFlexibleAppVersionIdentity
	versionsClient *api.VersionsClient
	servicesClient *api.ServicesClient
	desired        *krm.AppEngineFlexibleAppVersion
	actual         *pb.Version
}

var _ directbase.Adapter = &flexibleAppVersionAdapter{}

func (a *flexibleAppVersionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AppEngineFlexibleAppVersion", "name", a.id.String())

	req := &pb.GetVersionRequest{
		Name: a.id.String(),
		View: pb.VersionView_FULL,
	}
	actual, err := a.versionsClient.GetVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AppEngineFlexibleAppVersion %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *flexibleAppVersionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AppEngineFlexibleAppVersion", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredPb := AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Id = a.id.Version
	desiredPb.Env = "flex"

	req := &pb.CreateVersionRequest{
		Parent:  a.id.ParentString(),
		Version: desiredPb,
	}

	opLRO, err := a.versionsClient.CreateVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AppEngineFlexibleAppVersion %s: %w", a.id.String(), err)
	}

	created, err := opLRO.Wait(ctx)
	if err != nil {
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return fmt.Errorf("waiting for AppEngineFlexibleAppVersion %s creation: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully created AppEngineFlexibleAppVersion", "name", a.id.String())

	// Fetch fully-populated version
	getReq := &pb.GetVersionRequest{
		Name: a.id.String(),
		View: pb.VersionView_FULL,
	}
	refetched, err := a.versionsClient.GetVersion(ctx, getReq)
	if err != nil {
		refetched = created
	}

	return a.updateStatus(ctx, createOp, refetched)
}

func (a *flexibleAppVersionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AppEngineFlexibleAppVersion", "name", a.id.String())

	normalize := func(ctx context.Context, p *pb.Version) error {
		if p == nil || p.Deployment == nil || p.Deployment.Container == nil {
			return nil
		}
		img := p.Deployment.Container.Image
		// GCP resolves container image tags (e.g., image:tag) to digests (image@sha256:...).
		// Normalize image by trimming any tag or digest suffix to compare the base repository path.
		if idx := strings.IndexAny(img, ":@"); idx != -1 {
			p.Deployment.Container.Image = img[:idx]
		}
		return nil
	}

	diffs, _, err := common.CompareBrownfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AppEngineFlexibleAppVersionSpec_v1alpha1_FromProto,
		AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for AppEngineFlexibleAppVersion", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	mapCtx := &direct.MapContext{}
	desiredPb := AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Id = a.id.Version
	desiredPb.Env = "flex"

	// Flexible app version update in App Engine Admin API uses CreateVersion (POST) with existing version ID.
	req := &pb.CreateVersionRequest{
		Parent:  a.id.ParentString(),
		Version: desiredPb,
	}

	opLRO, err := a.versionsClient.CreateVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AppEngineFlexibleAppVersion %s: %w", a.id.String(), err)
	}

	updated, err := opLRO.Wait(ctx)
	if err != nil {
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return fmt.Errorf("waiting for AppEngineFlexibleAppVersion %s update: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully updated AppEngineFlexibleAppVersion", "name", a.id.String())

	// Fetch fully-populated version
	getReq := &pb.GetVersionRequest{
		Name: a.id.String(),
		View: pb.VersionView_FULL,
	}
	refetched, err := a.versionsClient.GetVersion(ctx, getReq)
	if err != nil {
		refetched = updated
	}

	return a.updateStatus(ctx, updateOp, refetched)
}

func (a *flexibleAppVersionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AppEngineFlexibleAppVersion", "name", a.id.String())

	if a.desired != nil && direct.ValueOf(a.desired.Spec.NoopOnDestroy) {
		log.V(2).Info("noopOnDestroy set to true, skipping deletion of AppEngineFlexibleAppVersion", "name", a.id.String())
		return true, nil
	}

	if a.desired != nil && direct.ValueOf(a.desired.Spec.DeleteServiceOnDestroy) {
		serviceName := a.id.ParentString()
		log.V(2).Info("deleteServiceOnDestroy set to true, deleting service", "name", serviceName)
		req := &pb.DeleteServiceRequest{
			Name: serviceName,
		}
		opLRO, err := a.servicesClient.DeleteService(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("deleting AppEngine service %s: %w", serviceName, err)
		}
		if err := opLRO.Wait(ctx); err != nil {
			if err.Error() != "unsupported result type <nil>: <nil>" {
				return false, fmt.Errorf("waiting for deletion of AppEngine service %s: %w", serviceName, err)
			}
		}
		log.V(2).Info("successfully deleted AppEngine service", "name", serviceName)
		return true, nil
	}

	req := &pb.DeleteVersionRequest{
		Name: a.id.String(),
	}
	opLRO, err := a.versionsClient.DeleteVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AppEngineFlexibleAppVersion %s: %w", a.id.String(), err)
	}

	if err := opLRO.Wait(ctx); err != nil {
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting for deletion of AppEngineFlexibleAppVersion %s: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully deleted AppEngineFlexibleAppVersion", "name", a.id.String())

	return true, nil
}

func (a *flexibleAppVersionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AppEngineFlexibleAppVersion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AppEngineFlexibleAppVersionSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Project = &a.id.App
	obj.Spec.ServiceRef = &krm.AppEngineServiceRef{
		External: "apps/" + a.id.App + "/services/" + a.id.Service,
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Version)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Version)
	u.SetGroupVersionKind(krm.AppEngineFlexibleAppVersionGVK)

	return u, nil
}

func (a *flexibleAppVersionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Version) error {
	mapCtx := &direct.MapContext{}
	status := AppEngineFlexibleAppVersionStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
