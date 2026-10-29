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
// proto.service: google.appengine.v1.Versions
// proto.message: google.appengine.v1.Version
// crd.type: AppEngineStandardAppVersion
// crd.version: v1alpha1

package appengine

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	api "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
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
	registry.RegisterModel(krm.AppEngineStandardAppVersionGVK, NewAppEngineStandardAppVersionModel)
}

func NewAppEngineStandardAppVersionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &standardAppVersionModel{config: *config}, nil
}

type standardAppVersionModel struct {
	config config.ControllerConfig
}

var _ directbase.Model = &standardAppVersionModel{}

func (m *standardAppVersionModel) client(ctx context.Context) (*api.VersionsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewVersionsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building appengine versions client: %w", err)
	}

	return gcpClient, nil
}

func (m *standardAppVersionModel) servicesClient(ctx context.Context) (*api.ServicesClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewServicesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building appengine services client: %w", err)
	}

	return gcpClient, nil
}

func (m *standardAppVersionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AppEngineStandardAppVersion{}
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
	versionId := id.(*krm.AppEngineStandardAppVersionIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	servicesClient, err := m.servicesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &standardAppVersionAdapter{
		gcpClient:      gcpClient,
		servicesClient: servicesClient,
		id:             versionId,
		desired:        obj,
	}, nil
}

func (m *standardAppVersionModel) AdapterForURL(ctx context.Context, rawURL string) (directbase.Adapter, error) {
	parsedURL, err := url.Parse(rawURL)
	if err == nil && parsedURL.Path != "" {
		rawURL = strings.TrimPrefix(parsedURL.Path, "/v1/")
		rawURL = strings.TrimPrefix(rawURL, "/")
	}
	id := &krm.AppEngineStandardAppVersionIdentity{}
	if err := id.FromExternal(rawURL); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	servicesClient, err := m.servicesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &standardAppVersionAdapter{
		gcpClient:      gcpClient,
		servicesClient: servicesClient,
		id:             id,
	}, nil
}

type standardAppVersionAdapter struct {
	id             *krm.AppEngineStandardAppVersionIdentity
	gcpClient      *api.VersionsClient
	servicesClient *api.ServicesClient
	desired        *krm.AppEngineStandardAppVersion
	actual         *pb.Version
}

var _ directbase.Adapter = &standardAppVersionAdapter{}

func (a *standardAppVersionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AppEngineStandardAppVersion", "name", a.id.String())

	req := &pb.GetVersionRequest{
		Name: a.id.String(),
		View: pb.VersionView_FULL,
	}
	actual, err := a.gcpClient.GetVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AppEngineStandardAppVersion %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *standardAppVersionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AppEngineStandardAppVersion", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredPb := AppEngineStandardAppVersionSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Id = a.id.Version

	req := &pb.CreateVersionRequest{
		Parent:  a.id.ParentString(),
		Version: desiredPb,
	}

	op, err := a.gcpClient.CreateVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AppEngineStandardAppVersion %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AppEngineStandardAppVersion creation %s: %w", a.id.String(), err)
	}

	getReq := &pb.GetVersionRequest{
		Name: a.id.String(),
		View: pb.VersionView_FULL,
	}
	latest, err := a.gcpClient.GetVersion(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting AppEngineStandardAppVersion after creation %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created AppEngineStandardAppVersion", "name", a.id.String())
	return a.updateStatus(ctx, createOp, latest)
}

func normalizeStandardAppVersionForDiff(ctx context.Context, p *pb.Version) error {
	// Entrypoint is not returned by the App Engine Admin API GET endpoint (it is write-only upon creation).
	// Clear entrypoint to prevent spurious diffs against the unreadable field.
	p.Entrypoint = nil

	// Deployment file sha1_sum is computed and returned by the server, but omitted in KRM spec.
	if p.Deployment != nil {
		for _, f := range p.Deployment.Files {
			if f != nil {
				f.Sha1Sum = ""
			}
		}
	}
	return nil
}

func (a *standardAppVersionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AppEngineStandardAppVersion", "name", a.id.String())

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
		AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
		normalizeStandardAppVersionForDiff,
		true, // fineGrained
	)
	if err != nil {
		return err
	}

	if diffRes.Empty() {
		log.V(2).Info("no diff detected for AppEngineStandardAppVersion", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffRes.Diff.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffRes.Diff)

	req := &pb.UpdateVersionRequest{
		Name:       a.id.String(),
		Version:    diffRes.MergedDesired,
		UpdateMask: diffRes.FieldMask,
	}

	op, err := a.gcpClient.UpdateVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AppEngineStandardAppVersion %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AppEngineStandardAppVersion update %s: %w", a.id.String(), err)
	}

	getReq := &pb.GetVersionRequest{
		Name: a.id.String(),
		View: pb.VersionView_FULL,
	}
	latest, err := a.gcpClient.GetVersion(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting AppEngineStandardAppVersion after update %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated AppEngineStandardAppVersion", "name", a.id.String())
	return a.updateStatus(ctx, updateOp, latest)
}

func (a *standardAppVersionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AppEngineStandardAppVersion", "name", a.id.String())

	if direct.ValueOf(a.desired.Spec.NoopOnDestroy) {
		log.V(2).Info("skipping deletion due to noopOnDestroy", "name", a.id.String())
		return true, nil
	}

	if direct.ValueOf(a.desired.Spec.DeleteServiceOnDestroy) {
		req := &pb.DeleteServiceRequest{
			Name: a.id.ParentString(),
		}
		op, err := a.servicesClient.DeleteService(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("deleting AppEngineService %s: %w", a.id.ParentString(), err)
		}
		if err := op.Wait(ctx); err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			return false, fmt.Errorf("waiting for AppEngineService deletion %s: %w", a.id.ParentString(), err)
		}
		log.V(2).Info("successfully deleted AppEngineService", "name", a.id.ParentString())
		return true, nil
	}

	req := &pb.DeleteVersionRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AppEngineStandardAppVersion %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for AppEngineStandardAppVersion deletion %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted AppEngineStandardAppVersion", "name", a.id.String())
	return true, nil
}

func (a *standardAppVersionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AppEngineStandardAppVersion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AppEngineStandardAppVersionSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// TODO: spec.entrypoint is +required in the KRM schema, but it is write-only upon creation
	// and unreadable from GCP GET responses, so exported YAML will require manual entry before re-applying.

	obj.Spec.Project = &a.id.App
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Version)
	obj.Spec.ServiceRef = &krm.AppEngineServiceRef{
		External: "apps/" + a.id.App + "/services/" + a.id.Service,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Version)
	u.SetGroupVersionKind(krm.AppEngineStandardAppVersionGVK)

	return u, nil
}

func (a *standardAppVersionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Version) error {
	status := &krm.AppEngineStandardAppVersionStatus{}
	status.Name = direct.LazyPtr(latest.GetName())
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
