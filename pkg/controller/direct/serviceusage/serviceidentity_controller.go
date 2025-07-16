// Copyright 2024 Google LLC
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
// proto.service: google.api.serviceusage.v1beta1.ServiceUsage
// proto.message: google.api.serviceusage.v1beta1.ServiceIdentity
// crd.type: ServiceIdentity
// crd.version: v1beta1

package serviceusage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/serviceusage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"

	gcp "google.golang.org/api/serviceusage/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// How long to wait for an LRO to complete
const operationPollTimeout = 5 * time.Minute

// GoogleApisServiceAgent was historically used by multiple services, but for more granular permission control we moved to P4SAs instead.
// MIGs were tricky to move though, and so instead of switching to a P4SA, they adopted the GoogleApisServiceAgent as their P4SA
// We still want to support this though, so we recognize this "magic value"
const GoogleApisServiceAgent = "cloudservices.gserviceaccount.com"

func init() {
	reconcileGate := &SecretReconcileGate{}
	registry.RegisterModelWithReconcileGate(krm.ServiceIdentityGVK, NewServiceIdentityModel, reconcileGate)
}

type SecretReconcileGate struct {
	optIn kccpredicate.OptInToDirectReconciliation
}

var _ kccpredicate.ReconcileGate = &SecretReconcileGate{}

func (r *SecretReconcileGate) ShouldReconcile(o *unstructured.Unstructured) bool {
	if r.optIn.ShouldReconcile(o) {
		return true
	}
	obj := &krm.ServiceIdentity{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &obj); err != nil {
		return false
	}
	serviceName := direct.ValueOf(obj.Spec.ResourceID)
	if serviceName == "" {
		serviceName = obj.GetName()
	}
	return serviceName == GoogleApisServiceAgent
}

func NewServiceIdentityModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &serviceIdentityModel{config: *config}, nil
}

var _ directbase.Model = &serviceIdentityModel{}

type serviceIdentityModel struct {
	config config.ControllerConfig
}

func (m *serviceIdentityModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ServiceIdentity{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get ServiceUsage GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	gcpBeta, err := gcpClient.newV1Beta1Client(ctx)
	if err != nil {
		return nil, err
	}
	return &serviceIdentityAdapter{
		gcpBeta:       gcpBeta,
		id:            id.(*krm.ServiceIdentityIdentity),
		desired:       obj,
		projectMapper: m.config.ProjectMapper,
	}, nil
}

func (m *serviceIdentityModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// This resource type does not have a direct GETtable URL in the traditional sense.
	return nil, nil
}

type serviceIdentityAdapter struct {
	projectMapper *projects.ProjectMapper
	gcpBeta       *gcp.APIService

	id      *krm.ServiceIdentityIdentity
	desired *krm.ServiceIdentity
	actual  *gcp.ServiceIdentity // This will be populated from status or after generation
}

var _ directbase.Adapter = &serviceIdentityAdapter{}

// Find determines if the service identity has already been generated and recorded.
func (a *serviceIdentityAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	// If status fields are populated, we consider the identity "found"
	// as its details are known to KCC.
	email := direct.ValueOf(a.desired.Status.Email)
	uniqueID := ""   //direct.ValueOf(a.desired.Status.UniqueID)
	if email != "" { // && uniqueID != "" {
		log.V(2).Info("service identity found in KRM status", "email", email)
		a.actual = &gcp.ServiceIdentity{
			Email:    email,
			UniqueId: uniqueID,
		}
		return true, nil
	}

	log.V(2).Info("service identity not yet recorded in KRM status, needs generation")
	return false, nil
}

// Create generates the service identity using the GenerateServiceIdentity RPC.
func (a *serviceIdentityAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)

	fqn := a.id.String()
	log.V(2).Info("generating service identity", "fqn", fqn)

	status := &krm.ServiceIdentityStatus{}
	if a.id.Service == GoogleApisServiceAgent {
		// Special case: the MIG "P4SA" (which is also the "Google APIs Service Agent")
		// Format is always <projectNumber>@cloudservices.gserviceaccount.com

		projectNumber, err := a.projectMapper.LookupProjectNumber(ctx, a.id.ParentID.ProjectID)
		if err != nil {
			return fmt.Errorf("looking up project number for %q: %w", a.id.ParentID.ProjectID, err)
		}
		email := fmt.Sprintf("%d@cloudservices.gserviceaccount.com", projectNumber)
		status.Email = &email
	} else {
		//   - parent: Name of the consumer and service to generate an identity for. The
		//     `GenerateServiceIdentity` methods currently support projects, folders,
		//     organizations. Example parents would be:
		//     `projects/123/services/example.googleapis.com`
		//     `folders/123/services/example.googleapis.com`
		//     `organizations/123/services/example.googleapis.com`.

		op, err := a.gcpBeta.Services.GenerateServiceIdentity(fqn).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("generating service identity for %q: %w", fqn, err)
		}

		var serviceIdentity *gcp.GoogleApiServiceusageV1beta1ServiceIdentity
		start := time.Now()
		for {
			if op.Done {
				identity := &gcp.GoogleApiServiceusageV1beta1ServiceIdentity{}
				if err := json.Unmarshal(op.Response, identity); err != nil {
					return fmt.Errorf("error parsing response: %w", err)
				}
				serviceIdentity = identity
				break
			}
			if time.Since(start) > operationPollTimeout {
				return fmt.Errorf("timed out waiting for service identity generation for %q", fqn)
			}
			time.Sleep(2 * time.Second)
			if err := ctx.Err(); err != nil {
				return err
			}
			op, err = a.gcpBeta.Operations.Get(op.Name).Context(ctx).Do()
			if err != nil {
				return fmt.Errorf("waiting for service identity generation for %q: %w", fqn, err)
			}
		}

		log.V(2).Info("successfully generated service identity", "fqn", fqn, "identity.email", serviceIdentity.Email, "identity.uniqueId", serviceIdentity.UniqueId)

		// It really doesn't seem worthwhile to use the mapper here

		// observedState := krm.ServiceIdentityObservedState{
		// 	Email:    direct.ValueOf(serviceIdentity.Email),
		// 	UniqueID: direct.ValueOf(serviceIdentity.UniqueId),
		// }

		// status.ObservedState = &observedState

		status.Email = direct.LazyPtr(serviceIdentity.Email)
		// status.UniqueID = direct.LazyPtr(serviceIdentity.UniqueId)
	}

	// status.ExternalRef = direct.LazyPtr(parent)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update for ServiceIdentity is a no-op as it should only be called when we already have a service identity.
func (a *serviceIdentityAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)

	fqn := a.id.String()
	log.V(2).Info("update is a no-op for ServiceIdentity", "fqn", fqn)

	return nil
}

// Export is not meaningful for ServiceIdentity as it has no configurable spec from the GCP resource.
func (a *serviceIdentityAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("export not implemented")
}

// Delete is a no-op for ServiceIdentity as they are managed by GCP with the service.
func (a *serviceIdentityAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("delete is a no-op for ServiceIdentity", "parent", fqn)
	return true, nil // Indicate successful (no-op) deletion
}
