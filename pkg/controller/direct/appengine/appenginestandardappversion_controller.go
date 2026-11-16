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
	"k8s.io/apimachinery/pkg/util/sets"
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

	return &standardAppVersionAdapter{
		gcpClient: gcpClient,
		id:        versionId,
		desired:   obj,
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

	return &standardAppVersionAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type standardAppVersionAdapter struct {
	id        *krm.AppEngineStandardAppVersionIdentity
	gcpClient *api.VersionsClient
	desired   *krm.AppEngineStandardAppVersion
	actual    *pb.Version
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

	req, err := a.newCreateVersionRequest()
	if err != nil {
		return err
	}

	op, err := a.gcpClient.CreateVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AppEngineStandardAppVersion %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AppEngineStandardAppVersion creation %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created AppEngineStandardAppVersion", "name", a.id.String())
	return a.refreshAndUpdateStatus(ctx, createOp, created)
}

func (a *standardAppVersionAdapter) newCreateVersionRequest() (*pb.CreateVersionRequest, error) {
	mapCtx := &direct.MapContext{}
	desiredPb := AppEngineStandardAppVersionSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredPb.Id = a.id.Version
	return &pb.CreateVersionRequest{
		Parent:  a.id.ParentString(),
		Version: desiredPb,
	}, nil
}

func (a *standardAppVersionAdapter) refreshAndUpdateStatus(ctx context.Context, op directbase.Operation, fromLRO *pb.Version) error {
	log := klog.FromContext(ctx)
	latest := fromLRO
	getReq := &pb.GetVersionRequest{
		Name: a.id.String(),
		View: pb.VersionView_FULL,
	}
	if v, err := a.gcpClient.GetVersion(ctx, getReq); err != nil {
		log.V(2).Info("getting AppEngineStandardAppVersion failed; falling back to op.Wait result", "name", a.id.String(), "err", err)
	} else {
		latest = v
	}
	return a.updateStatus(ctx, op, latest)
}

func (a *standardAppVersionAdapter) diffNormalizer() func(context.Context, *pb.Version) error {
	if len(a.desired.Spec.Handlers) == 0 {
		return normalizeStandardAppVersionForDiff
	}
	return func(ctx context.Context, p *pb.Version) error {
		if err := normalizeStandardAppVersionForDiff(ctx, p); err != nil {
			return err
		}
		p.Handlers = normalizeHandlersForDiff(p.Handlers)
		return nil
	}
}

func normalizeStandardAppVersionForDiff(ctx context.Context, p *pb.Version) error {
	// Entrypoint is not returned by the App Engine Admin API GET endpoint (it is write-only upon creation).
	// Clear entrypoint to prevent spurious diffs against the unreadable field.
	p.Entrypoint = nil

	// Deployment.zip is write-only (never returned by GET), and deployment.files[*].sha1_sum
	// is computed by the server and omitted in the KRM spec.
	if p.Deployment != nil {
		p.Deployment.Zip = nil
		for _, f := range p.Deployment.Files {
			if f != nil {
				f.Sha1Sum = ""
			}
		}
	}
	return nil
}

func normalizeHandlersForDiff(handlers []*pb.UrlMap) []*pb.UrlMap {
	var out []*pb.UrlMap
	for _, h := range handlers {
		if h == nil {
			continue
		}
		if h.Login == pb.LoginRequirement_LOGIN_UNSPECIFIED {
			h.Login = pb.LoginRequirement_LOGIN_OPTIONAL
		}
		if h.SecurityLevel == pb.SecurityLevel_SECURE_UNSPECIFIED {
			h.SecurityLevel = pb.SecurityLevel_SECURE_OPTIONAL
		}
		if h.AuthFailAction == pb.AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED {
			h.AuthFailAction = pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT
		}
		if sf := h.GetStaticFiles(); sf != nil && sf.Expiration != nil && sf.Expiration.Seconds == 0 && sf.Expiration.Nanos == 0 {
			sf.Expiration = nil
		}
		out = append(out, h)
	}
	for len(out) > 0 && isRuntimeDefaultHandler(out[len(out)-1]) {
		out = out[:len(out)-1]
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func isRuntimeDefaultHandler(h *pb.UrlMap) bool {
	if h == nil {
		return false
	}
	return h.GetUrlRegex() == ".*" &&
		h.GetScript() != nil &&
		h.GetScript().GetScriptPath() == "auto" &&
		h.GetLogin() == pb.LoginRequirement_LOGIN_OPTIONAL &&
		h.GetSecurityLevel() == pb.SecurityLevel_SECURE_OPTIONAL &&
		h.GetAuthFailAction() == pb.AuthFailAction_AUTH_FAIL_ACTION_REDIRECT &&
		h.GetRedirectHttpResponseCode() == pb.UrlMap_REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED
}

type standardScalingType string

const (
	standardScalingUnspecified standardScalingType = ""
	standardScalingAutomatic   standardScalingType = "automatic"
	standardScalingBasic       standardScalingType = "basic"
	standardScalingManual      standardScalingType = "manual"
)

func (a *standardAppVersionAdapter) scalingTypeChange() (actual, desired standardScalingType, changed bool) {
	actual = scalingTypeOfVersion(a.actual)
	desired = scalingTypeOfSpec(&a.desired.Spec)
	changed = desired != standardScalingUnspecified && desired != actual
	return actual, desired, changed
}

// scalingTypeOfSpec returns the scaling type configured on the unmerged spec using
// the same precedence as AppEngineStandardAppVersionSpec_v1alpha1_ToProto
// (manual > basic > automatic).
func scalingTypeOfSpec(spec *krm.AppEngineStandardAppVersionSpec) standardScalingType {
	if spec == nil {
		return standardScalingUnspecified
	}
	if spec.ManualScaling != nil {
		return standardScalingManual
	}
	if spec.BasicScaling != nil {
		return standardScalingBasic
	}
	if spec.AutomaticScaling != nil {
		return standardScalingAutomatic
	}
	return standardScalingUnspecified
}

func scalingTypeOfVersion(v *pb.Version) standardScalingType {
	if v != nil {
		if v.GetManualScaling() != nil {
			return standardScalingManual
		}
		if v.GetBasicScaling() != nil {
			return standardScalingBasic
		}
	}
	return standardScalingAutomatic
}

// standardPatchableUpdatePaths are the update_mask paths App Engine's UpdateVersion accepts
// for a Standard version that the KRM spec can express (see the rpc UpdateVersion comment in
// google/appengine/v1/appengine.proto). App Engine rejects the whole request with
// INVALID_ARGUMENT if any other path is present.
var standardPatchableUpdatePaths = sets.New(
	"instance_class",
	"automatic_scaling.min_idle_instances",
	"automatic_scaling.max_idle_instances",
	"automatic_scaling.standard_scheduler_settings.min_instances",
	"automatic_scaling.standard_scheduler_settings.max_instances",
	"automatic_scaling.standard_scheduler_settings.target_cpu_utilization",
	"automatic_scaling.standard_scheduler_settings.target_throughput_utilization",
	"manual_scaling.instances",
)

type standardUpdateStrategy int

const (
	standardUpdateNone standardUpdateStrategy = iota
	standardUpdatePatch
	standardUpdateRedeploy
)

func chooseStandardUpdateStrategy(changedPaths []string, scalingTypeChanged bool) (standardUpdateStrategy, []string) {
	if scalingTypeChanged {
		return standardUpdateRedeploy, []string{"scaling type"}
	}
	if len(changedPaths) == 0 {
		return standardUpdateNone, nil
	}
	var reasons []string
	for _, p := range changedPaths {
		if !standardPatchableUpdatePaths.Has(p) {
			reasons = append(reasons, p)
		}
	}
	if len(reasons) > 0 {
		return standardUpdateRedeploy, reasons
	}
	return standardUpdatePatch, nil
}

// Update reconciles an existing version. App Engine's UpdateVersion accepts only a few
// update_mask paths, so any other change redeploys the version: CreateVersion with the
// same version id, which App Engine treats as an upsert.
func (a *standardAppVersionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AppEngineStandardAppVersion", "name", a.id.String())

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AppEngineStandardAppVersionSpec_v1alpha1_FromProto,
		AppEngineStandardAppVersionSpec_v1alpha1_ToProto,
		a.diffNormalizer(),
		true, // fineGrained
	)
	if err != nil {
		return err
	}

	actualScaling, desiredScaling, scalingChanged := a.scalingTypeChange()
	strategy, redeployReasons := chooseStandardUpdateStrategy(diffRes.FieldMask.GetPaths(), scalingChanged)
	if strategy == standardUpdateNone {
		log.V(2).Info("no diff detected for AppEngineStandardAppVersion", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	if scalingChanged {
		diffRes.Diff.AddField("scaling", string(actualScaling), string(desiredScaling))
	}
	diffRes.Diff.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffRes.Diff)

	if strategy == standardUpdatePatch {
		return a.patch(ctx, updateOp, diffRes)
	}
	log.V(2).Info("redeploying AppEngineStandardAppVersion with CreateVersion", "name", a.id.String(), "reasons", redeployReasons)
	return a.redeploy(ctx, updateOp)
}

func (a *standardAppVersionAdapter) patch(ctx context.Context, updateOp *directbase.UpdateOperation, diffRes *common.SpecDiffResult[*pb.Version]) error {
	log := klog.FromContext(ctx)
	req := &pb.UpdateVersionRequest{
		Name:       a.id.String(),
		Version:    diffRes.MergedDesired,
		UpdateMask: diffRes.FieldMask,
	}

	op, err := a.gcpClient.UpdateVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AppEngineStandardAppVersion %s: %w", a.id.String(), err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AppEngineStandardAppVersion update %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated AppEngineStandardAppVersion", "name", a.id.String())
	return a.refreshAndUpdateStatus(ctx, updateOp, updated)
}

func (a *standardAppVersionAdapter) redeploy(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	req, err := a.newCreateVersionRequest()
	if err != nil {
		return err
	}

	op, err := a.gcpClient.CreateVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("redeploying AppEngineStandardAppVersion %s: %w", a.id.String(), err)
	}

	deployed, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AppEngineStandardAppVersion redeploy %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully redeployed AppEngineStandardAppVersion", "name", a.id.String())
	return a.refreshAndUpdateStatus(ctx, updateOp, deployed)
}

func (a *standardAppVersionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AppEngineStandardAppVersion", "name", a.id.String())

	if direct.ValueOf(a.desired.Spec.NoopOnDestroy) {
		log.V(2).Info("skipping deletion due to noopOnDestroy", "name", a.id.String())
		return true, nil
	}

	if direct.ValueOf(a.desired.Spec.DeleteServiceOnDestroy) {
		// TODO: Deleting the parent service from a child version resource requires further architectural design.
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
