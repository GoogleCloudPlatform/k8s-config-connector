// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:controller
// proto.service: google.cloud.dataproc.v1.SessionController
// proto.message: google.cloud.dataproc.v1.Session
// crd.type: DataprocSession
// crd.version: v1alpha1

package session

import (
	"context"
	"fmt"
	"strings"

	dataproc "cloud.google.com/go/dataproc/v2/apiv1"
	dataprocpb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DataprocSessionGVK, NewSessionModel)
}

func NewSessionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sessionModel{config: *config}, nil
}

var _ directbase.Model = &sessionModel{}

type sessionModel struct {
	config config.ControllerConfig
}

func (m *sessionModel) Client(ctx context.Context, projectID string) (*dataproc.SessionControllerClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := dataproc.NewSessionControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc session client: %w", err)
	}

	return gcpClient, err
}

func getIdentity(ctx context.Context, reader client.Reader, obj *krm.DataprocSession) (*krm.DataprocSessionIdentity, error) {
	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := identity.(*krm.DataprocSessionIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", identity)
	}
	return id, nil
}

func (m *sessionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataprocSession{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := getIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.Client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &sessionAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *sessionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type sessionAdapter struct {
	gcpClient *dataproc.SessionControllerClient
	id        *krm.DataprocSessionIdentity
	desired   *krm.DataprocSession
	actual    *dataprocpb.Session
}

var _ directbase.Adapter = &sessionAdapter{}

func (a *sessionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataproc session", "name", a.id.String())

	req := &dataprocpb.GetSessionRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSession(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataproc session %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *sessionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataproc session", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataprocSessionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &dataprocpb.CreateSessionRequest{
		Parent:    parent,
		Session:   resource,
		SessionId: a.id.Session,
	}
	op, err := a.gcpClient.CreateSession(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataproc session %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("dataproc session %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created dataproc session in gcp", "name", a.id.String())

	status := &krm.DataprocSessionStatus{}
	status.ObservedState = DataprocSessionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// DataprocSession does not support update.
func (a *sessionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataproc session", "name", a.id.String())

	desiredpb := DataprocSessionSpec_ToProto(&direct.MapContext{}, &a.desired.Spec)
	paths, err := common.CompareProtoMessage(desiredpb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) != 0 {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)
		log.V(2).Info("This resource does not support update", "name", a.id.String())
		return nil
	}

	status := &krm.DataprocSessionStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *sessionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataproc session", "name", a.id.String())

	req := &dataprocpb.DeleteSessionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteSession(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting dataproc session %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		if !strings.Contains(err.Error(), "mismatched message type") {
			return false, fmt.Errorf("waiting delete dataproc session %s: %w", a.id.String(), err)
		}
	}
	log.V(2).Info("successfully deleted dataproc session", "name", a.id.String())

	return true, nil
}

// Export implements the Adapter interface.
func (a *sessionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataprocSession{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataprocSessionSpec_FromProto(mapCtx, a.actual))
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
	u.SetName(a.id.Session)
	u.SetGroupVersionKind(krm.DataprocSessionGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}
