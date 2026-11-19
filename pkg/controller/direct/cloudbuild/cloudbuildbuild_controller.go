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

// +generated:controller
// krm.group: cloudbuild.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CloudBuildBuild

package cloudbuild

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/cloudbuild/apiv1/v2"
	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudBuildBuildGVK, NewBuildModel)
}

func NewBuildModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBuild{config: *config}, nil
}

var _ directbase.Model = &modelBuild{}

type modelBuild struct {
	config config.ControllerConfig
}

func (m *modelBuild) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudBuildBuild{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CloudBuildBuildIdentity)

	// Get cloudbuild GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &BuildAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelBuild) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//cloudbuild.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CloudBuildBuildIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &BuildAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelBuild) client(ctx context.Context, location string) (*gcp.Client, error) {
	config := m.config
	if strings.Contains(config.UserAgent, "kcc/controller-manager") {
		// Cloud Build API returns 500 / rpc error Internal when User-Agent contains duplicate "kcc/" tokens.
		config.UserAgent = strings.Replace(config.UserAgent, "kcc/controller-manager", "controller-manager", 1)
	}
	var opts []option.ClientOption
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-cloudbuild.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Build client: %w", err)
	}
	return gcpClient, nil
}

type BuildAdapter struct {
	id        *krm.CloudBuildBuildIdentity
	gcpClient *gcp.Client
	desired   *krm.CloudBuildBuild
	actual    *cloudbuildpb.Build
}

var _ directbase.Adapter = &BuildAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BuildAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Build", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &cloudbuildpb.GetBuildRequest{
		Name: a.id.String(),
	}
	buildpb, err := a.gcpClient.GetBuild(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Build %q: %w", a.id, err)
	}

	a.actual = buildpb
	return true, nil
}

func cleanBuildStatus(status *krm.CloudBuildBuildStatus) {
	if status != nil && status.ObservedState != nil && status.ObservedState.Results != nil {
		var cleaned [][]byte
		for _, out := range status.ObservedState.Results.BuildStepOutputs {
			if len(out) > 0 {
				cleaned = append(cleaned, out)
			}
		}
		status.ObservedState.Results.BuildStepOutputs = cleaned
	}
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BuildAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Build", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudBuildBuildSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cloudbuildpb.CreateBuildRequest{
		Parent: a.id.ParentString(),
		Build:  resource,
	}
	op, err := a.gcpClient.CreateBuild(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Build %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Build %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Build", "name", a.id)

	if created.Name != "" {
		parsedID := &krm.CloudBuildBuildIdentity{}
		if err := parsedID.FromExternal(created.Name); err != nil {
			return err
		}
		a.id.Build = parsedID.Build
	} else if created.Id != "" {
		a.id.Build = created.Id
	}

	// Fetch fully-populated resource after creation
	latest, err := a.gcpClient.GetBuild(ctx, &cloudbuildpb.GetBuildRequest{Name: a.id.String()})
	if err != nil {
		latest = created
	}
	a.actual = latest

	status := &krm.CloudBuildBuildStatus{}
	status.ObservedState = CloudBuildBuildObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	cleanBuildStatus(status)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BuildAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Build", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CloudBuildBuildSpec_v1alpha1_FromProto,
		CloudBuildBuildSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating CloudBuildBuild is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.CloudBuildBuildStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudBuildBuildObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	cleanBuildStatus(status)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BuildAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudBuildBuild{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudBuildBuildSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refsv1beta1.ProjectRef{External: a.id.Project},
		Location:   a.id.Location,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CloudBuildBuildGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BuildAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deletion not supported for Build, removing finalizer without deleting GCP resource", "name", a.id)
	return true, nil
}
