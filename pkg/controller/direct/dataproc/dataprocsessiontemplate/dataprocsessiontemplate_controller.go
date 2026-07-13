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
// proto.service: google.cloud.dataproc.v1.SessionTemplateController
// proto.message: google.cloud.dataproc.v1.SessionTemplate
// crd.type: DataprocSessionTemplate
// crd.version: v1alpha1

package dataprocsessiontemplate

import (
	"context"
	"fmt"

	dataproc "cloud.google.com/go/dataproc/v2/apiv1"
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	dataprocgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/dataproc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DataprocSessionTemplateGVK, NewDataprocSessionTemplateModel)
}

func NewDataprocSessionTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sessionTemplateModel{config: *config}, nil
}

var _ directbase.Model = &sessionTemplateModel{}

type sessionTemplateModel struct {
	config config.ControllerConfig
}

func (m *sessionTemplateModel) client(ctx context.Context, projectID string) (*dataproc.SessionTemplateControllerClient, error) {
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

	gcpClient, err := dataproc.NewSessionTemplateControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc sessiontemplate client: %w", err)
	}

	return gcpClient, err
}

func (m *sessionTemplateModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataprocSessionTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	templateID, ok := id.(*krm.DataprocSessionTemplateIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", id)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx, templateID.Project)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := dataprocgcp.DataprocSessionTemplateSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = templateID.String()
	if desired.Labels == nil {
		desired.Labels = make(map[string]string)
	}
	for k, v := range obj.GetLabels() {
		desired.Labels[k] = v
	}
	desired.Labels["managed-by-cnrm"] = "true"

	return &sessionTemplateAdapter{
		gcpClient: gcpClient,
		id:        templateID,
		desired:   desired,
		obj:       obj,
	}, nil
}

func (m *sessionTemplateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type sessionTemplateAdapter struct {
	gcpClient *dataproc.SessionTemplateControllerClient
	id        *krm.DataprocSessionTemplateIdentity
	desired   *pb.SessionTemplate
	obj       *krm.DataprocSessionTemplate
	actual    *pb.SessionTemplate
}

var _ directbase.Adapter = &sessionTemplateAdapter{}

func (a *sessionTemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataproc sessiontemplate", "name", a.id)

	req := &pb.GetSessionTemplateRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSessionTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataproc sessiontemplate %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *sessionTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataproc sessiontemplate", "name", a.id)

	req := &pb.CreateSessionTemplateRequest{
		Parent:          a.id.ParentString(),
		SessionTemplate: a.desired,
	}
	latest, err := a.gcpClient.CreateSessionTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataproc sessiontemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created dataproc sessiontemplate in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, latest)
}

func (a *sessionTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataproc sessiontemplate", "name", a.id)

	var maskedActual *pb.SessionTemplate
	{
		mapCtx := &direct.MapContext{}
		spec := dataprocgcp.DataprocSessionTemplateSpec_v1alpha1_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		maskedActual = dataprocgcp.DataprocSessionTemplateSpec_v1alpha1_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
	}
	maskedActual.Name = a.desired.Name
	maskedActual.Labels = a.actual.Labels

	clonedDesired := proto.Clone(a.desired).(*pb.SessionTemplate)

	diffs, _, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	u := updateOp.GetUnstructured()
	diffs.Object = u
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateSessionTemplateRequest{
		SessionTemplate: a.desired,
	}
	latest, err := a.gcpClient.UpdateSessionTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("updating dataproc sessiontemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated dataproc sessiontemplate", "name", a.id)

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *sessionTemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataproc sessiontemplate", "name", a.id)

	req := &pb.DeleteSessionTemplateRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSessionTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting dataproc sessiontemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataproc sessiontemplate", "name", a.id)

	return true, nil
}

func (a *sessionTemplateAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SessionTemplate) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DataprocSessionTemplateStatus{}
	status.ObservedState = dataprocgcp.DataprocSessionTemplateObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export implements the Adapter interface.
func (a *sessionTemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataprocSessionTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(dataprocgcp.DataprocSessionTemplateSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Template)
	u.SetGroupVersionKind(krm.DataprocSessionTemplateGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}
