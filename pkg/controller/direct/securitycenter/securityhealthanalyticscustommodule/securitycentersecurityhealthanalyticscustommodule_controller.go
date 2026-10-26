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

package securityhealthanalyticscustommodule

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/securitycenter/apiv1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/securitycenter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SecuritycenterSecurityHealthAnalyticsCustomModuleGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SecurityCenter REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecuritycenterSecurityHealthAnalyticsCustomModule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.SecuritycenterSecurityHealthAnalyticsCustomModuleIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.SecuritycenterSecurityHealthAnalyticsCustomModuleIdentity
	gcpClient *gcp.Client
	desired   *krm.SecuritycenterSecurityHealthAnalyticsCustomModule
	actual    *pb.SecurityHealthAnalyticsCustomModule
}

var _ directbase.Adapter = &Adapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the request.
func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	// Rule 2: Find() Pre-Check Guard for server-generated ID
	if a.id.CustomModule == "" {
		return false, nil
	}

	log.V(2).Info("getting SecuritycenterSecurityHealthAnalyticsCustomModule", "name", a.id.String())
	req := &pb.GetSecurityHealthAnalyticsCustomModuleRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecuritycenterSecurityHealthAnalyticsCustomModule %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and updates status.
func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SecuritycenterSecurityHealthAnalyticsCustomModule", "parent", a.id.ParentString())
	mapCtx := &direct.MapContext{}

	desiredPb := securitycenter.SecuritycenterSecurityHealthAnalyticsCustomModuleSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Rule 3: Create() Request Payload - do not set Name, let GCP generate the identifier.
	desiredPb.Name = ""

	req := &pb.CreateSecurityHealthAnalyticsCustomModuleRequest{
		Parent:                              a.id.ParentString(),
		SecurityHealthAnalyticsCustomModule: desiredPb,
	}

	created, err := a.gcpClient.CreateSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SecuritycenterSecurityHealthAnalyticsCustomModule: %w", err)
	}
	log.V(2).Info("successfully created SecuritycenterSecurityHealthAnalyticsCustomModule", "name", created.GetName())

	status := &krm.SecuritycenterSecurityHealthAnalyticsCustomModuleStatus{}
	status.ObservedState = securitycenter.SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.GetName())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and updates status.
func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SecuritycenterSecurityHealthAnalyticsCustomModule", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredPb := securitycenter.SecuritycenterSecurityHealthAnalyticsCustomModuleSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(desiredPb).(*pb.SecurityHealthAnalyticsCustomModule)
	actualSpec := securitycenter.SecuritycenterSecurityHealthAnalyticsCustomModuleSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := securitycenter.SecuritycenterSecurityHealthAnalyticsCustomModuleSpec_ToProto(mapCtx, actualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	updated := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
	} else {
		log.V(2).Info("fields need update", "name", a.id.String(), "updateMask", updateMask)

		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredPb.Name = a.id.String()
		req := &pb.UpdateSecurityHealthAnalyticsCustomModuleRequest{
			UpdateMask:                          updateMask,
			SecurityHealthAnalyticsCustomModule: desiredPb,
		}
		updated, err = a.gcpClient.UpdateSecurityHealthAnalyticsCustomModule(ctx, req)
		if err != nil {
			return fmt.Errorf("updating SecuritycenterSecurityHealthAnalyticsCustomModule %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated SecuritycenterSecurityHealthAnalyticsCustomModule", "name", a.id.String())
	}

	status := &krm.SecuritycenterSecurityHealthAnalyticsCustomModuleStatus{}
	status.ObservedState = securitycenter.SecuritycenterSecurityHealthAnalyticsCustomModuleObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.GetName())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecuritycenterSecurityHealthAnalyticsCustomModule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(securitycenter.SecuritycenterSecurityHealthAnalyticsCustomModuleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.Organization}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.CustomModule)
	u.SetGroupVersionKind(krm.SecuritycenterSecurityHealthAnalyticsCustomModuleGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SecuritycenterSecurityHealthAnalyticsCustomModule", "name", a.id.String())

	req := &pb.DeleteSecurityHealthAnalyticsCustomModuleRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSecurityHealthAnalyticsCustomModule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent SecuritycenterSecurityHealthAnalyticsCustomModule", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting SecuritycenterSecurityHealthAnalyticsCustomModule %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted SecuritycenterSecurityHealthAnalyticsCustomModule", "name", a.id.String())
	return true, nil
}
