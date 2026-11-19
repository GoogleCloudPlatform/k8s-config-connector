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
// krm.group: ftp.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: FTPServer

package ftp

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ftp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/ftp/apiv1"
	ftppb "cloud.google.com/go/ftp/apiv1/ftppb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.FTPServerGVK, NewServerModel)
}

func NewServerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelServer{config: *config}, nil
}

var _ directbase.Model = &modelServer{}

type modelServer struct {
	config config.ControllerConfig
}

func (m *modelServer) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.FTPServer{}
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
	id := i.(*krm.FTPServerIdentity)

	// Get ftp GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServerAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelServer) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//ftp.googleapis.com/") {
		return nil, nil
	}
	id := &krm.FTPServerIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServerAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelServer) client(ctx context.Context) (*gcp.CloudFtpClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudFtpRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Server client: %w", err)
	}
	return gcpClient, err
}

type ServerAdapter struct {
	id        *krm.FTPServerIdentity
	gcpClient *gcp.CloudFtpClient
	desired   *krm.FTPServer
	actual    *ftppb.Server
}

var _ directbase.Adapter = &ServerAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ServerAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Server", "name", a.id)

	req := &ftppb.GetServerRequest{
		Name: a.id.String(),
		View: ftppb.ServerView_SERVER_VIEW_FULL,
	}
	serverpb, err := a.gcpClient.GetServer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Server %q: %w", a.id, err)
	}

	a.actual = serverpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Server", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FTPServerSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = label.GCPLabels(a.desired)

	req := &ftppb.CreateServerRequest{
		Parent:   a.id.ParentString(),
		ServerId: a.id.Server,
		Server:   resource,
	}
	op, err := a.gcpClient.CreateServer(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Server %s: %w", a.id, err)
	}
	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("Server %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Server", "name", a.id)

	fullServer, err := a.gcpClient.GetServer(ctx, &ftppb.GetServerRequest{
		Name: a.id.String(),
		View: ftppb.ServerView_SERVER_VIEW_FULL,
	})
	if err != nil {
		return fmt.Errorf("getting Server %s after creation: %w", a.id, err)
	}

	status := &krm.FTPServerStatus{}
	status.ObservedState = FTPServerObservedState_FromProto(mapCtx, fullServer)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ServerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Server", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpecAndLabels(
		ctx,
		updateOp.GetUnstructured(),
		&a.desired.Spec,
		a.actual,
		"labels",
		FTPServerSpec_FromProto,
		FTPServerSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &ftppb.UpdateServerRequest{
			UpdateMask: updateMask,
			Server:     resource,
		}
		op, err := a.gcpClient.UpdateServer(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Server %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Server %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Server", "name", a.id)

		fullServer, err := a.gcpClient.GetServer(ctx, &ftppb.GetServerRequest{
			Name: a.id.String(),
			View: ftppb.ServerView_SERVER_VIEW_FULL,
		})
		if err != nil {
			return fmt.Errorf("getting Server %s after update: %w", a.id, err)
		}
		updated = fullServer
	}

	status := &krm.FTPServerStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = FTPServerObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ServerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.FTPServer{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FTPServerSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.FTPServerGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ServerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Server", "name", a.id)

	req := &ftppb.DeleteServerRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteServer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Server, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Server %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Server", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Server %s: %w", a.id, err)
	}
	return true, nil
}
