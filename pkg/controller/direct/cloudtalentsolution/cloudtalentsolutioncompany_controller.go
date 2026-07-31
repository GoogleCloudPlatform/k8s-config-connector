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

package cloudtalentsolution

import (
	"context"
	"fmt"
	"strings"

	talent "cloud.google.com/go/talent/apiv4"
	pb "cloud.google.com/go/talent/apiv4/talentpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudtalentsolution/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CloudTalentSolutionCompanyGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*talent.CompanyClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := talent.NewCompanyRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building cloudtalentsolution client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudTalentSolutionCompany{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.CloudTalentSolutionCompanyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	if id.Tenant == "default" {
		var opts []option.ClientOption
		opts, err = m.config.RESTClientOptions()
		if err != nil {
			return nil, err
		}
		tenantClient, err := talent.NewTenantRESTClient(ctx, opts...)
		if err != nil {
			return nil, fmt.Errorf("building tenant client: %w", err)
		}
		defer tenantClient.Close()

		it := tenantClient.ListTenants(ctx, &pb.ListTenantsRequest{
			Parent: fmt.Sprintf("projects/%s", id.Project),
		})
		tenant, err := it.Next()
		if err != nil {
			return nil, fmt.Errorf("listing tenants to find default tenant: %w", err)
		}
		parts := strings.Split(tenant.GetName(), "/")
		if len(parts) < 4 {
			return nil, fmt.Errorf("unexpected tenant name format: %q", tenant.GetName())
		}
		id.Tenant = parts[3]
	}

	mapCtx := &direct.MapContext{}
	desired := CloudTalentSolutionCompanySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type adapter struct {
	gcpClient *talent.CompanyClient
	id        *krm.CloudTalentSolutionCompanyIdentity
	desired   *pb.Company
	actual    *pb.Company
	reader    client.Reader
}

var _ directbase.Adapter = &adapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting cloudtalentsolution company", "name", a.id)

	req := &pb.GetCompanyRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCompany(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting cloudtalentsolution company %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating cloudtalentsolution company", "name", a.id)

	req := &pb.CreateCompanyRequest{
		Parent:  a.id.ParentString(),
		Company: a.desired,
	}
	created, err := a.gcpClient.CreateCompany(ctx, req)
	if err != nil {
		return fmt.Errorf("creating cloudtalentsolution company %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created cloudtalentsolution company", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating cloudtalentsolution company", "name", a.id)

	maskedActual, err := mappers.OnlySpecFields(a.actual, CloudTalentSolutionCompanySpec_FromProto, CloudTalentSolutionCompanySpec_ToProto)
	if err != nil {
		return err
	}

	clonedDesired := proto.Clone(a.desired).(*pb.Company)
	maskedActual.Name = clonedDesired.Name

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateCompanyRequest{
		Company:    a.desired,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateCompany(ctx, req)
	if err != nil {
		return fmt.Errorf("updating cloudtalentsolution company %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated cloudtalentsolution company", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Company) error {
	mapCtx := &direct.MapContext{}
	observedState := CloudTalentSolutionCompanyObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.CloudTalentSolutionCompanyStatus{}
	status.ObservedState = observedState
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudTalentSolutionCompany{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudTalentSolutionCompanySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Company)
	u.SetGroupVersionKind(krm.CloudTalentSolutionCompanyGVK)
	u.Object = uObj
	return u, nil
}

// Delete deletes the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting cloudtalentsolution company", "name", a.id)

	req := &pb.DeleteCompanyRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteCompany(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting cloudtalentsolution company %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted cloudtalentsolution company", "name", a.id)
	return true, nil
}
