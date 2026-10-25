// Copyright 2025 Google LLC
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
// krm.group: developerconnect.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DeveloperConnectGitRepositoryLink

package developerconnect

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/developerconnect/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/developerconnect/apiv1"
	developerconnectpb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DeveloperConnectGitRepositoryLinkGVK, NewGitRepositoryLinkModel)
}

func NewGitRepositoryLinkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGitRepositoryLink{config: *config}, nil
}

var _ directbase.Model = &modelGitRepositoryLink{}

type modelGitRepositoryLink struct {
	config config.ControllerConfig
}

func (m *modelGitRepositoryLink) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DeveloperConnectGitRepositoryLink{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DeveloperConnectGitRepositoryLinkIdentity)

	// Get developerconnect GCP client
	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &GitRepositoryLinkAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelGitRepositoryLink) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}
func (m *modelGitRepositoryLink) client(ctx context.Context, location string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("https://developerconnect.%s.rep.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building GitRepositoryLink client: %w", err)
	}
	return gcpClient, err
}

type GitRepositoryLinkAdapter struct {
	id        *krm.DeveloperConnectGitRepositoryLinkIdentity
	gcpClient *gcp.Client
	desired   *krm.DeveloperConnectGitRepositoryLink
	actual    *developerconnectpb.GitRepositoryLink
}

var _ directbase.Adapter = &GitRepositoryLinkAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GitRepositoryLinkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting GitRepositoryLink", "name", a.id)

	req := &developerconnectpb.GetGitRepositoryLinkRequest{
		Name: a.id.String(),
	}
	gitrepositorylinkpb, err := a.gcpClient.GetGitRepositoryLink(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GitRepositoryLink %q: %w", a.id, err)
	}

	a.actual = gitrepositorylinkpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GitRepositoryLinkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GitRepositoryLink", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DeveloperConnectGitRepositoryLinkSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &developerconnectpb.CreateGitRepositoryLinkRequest{
		Parent:              a.id.ParentString(),
		GitRepositoryLinkId: a.id.GitRepositoryLink,
		GitRepositoryLink:   resource,
	}
	op, err := a.gcpClient.CreateGitRepositoryLink(ctx, req)
	if err != nil {
		return fmt.Errorf("creating GitRepositoryLink %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("GitRepositoryLink %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created GitRepositoryLink", "name", a.id)

	status := &krm.DeveloperConnectGitRepositoryLinkStatus{}
	status.ObservedState = DeveloperConnectGitRepositoryLinkObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GitRepositoryLinkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GitRepositoryLink", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DeveloperConnectGitRepositoryLinkSpec_FromProto,
		DeveloperConnectGitRepositoryLinkSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating DeveloperConnectGitRepositoryLink is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.DeveloperConnectGitRepositoryLinkStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DeveloperConnectGitRepositoryLinkObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *GitRepositoryLinkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DeveloperConnectGitRepositoryLink{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DeveloperConnectGitRepositoryLinkSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DeveloperConnectGitRepositoryLinkGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GitRepositoryLinkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting GitRepositoryLink", "name", a.id)

	req := &developerconnectpb.DeleteGitRepositoryLinkRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteGitRepositoryLink(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent GitRepositoryLink, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting GitRepositoryLink %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted GitRepositoryLink", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete GitRepositoryLink %s: %w", a.id, err)
	}
	return true, nil
}
