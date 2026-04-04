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

package resourcemanager

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/orgpolicy/apiv2"

	orgpolicypb "cloud.google.com/go/orgpolicy/apiv2/orgpolicypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ResourceManagerPolicyGVK, NewPolicyModel)
}

func NewPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPolicy{config: *config}, nil
}

var _ directbase.Model = &modelPolicy{}

type modelPolicy struct {
	config config.ControllerConfig
}

func (m *modelPolicy) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Policy client: %w", err)
	}
	return gcpClient, err
}

func (m *modelPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ResourceManagerPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get orgpolicy GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type PolicyAdapter struct {
	id        *krm.PolicyIdentity
	gcpClient *gcp.Client
	desired   *krm.ResourceManagerPolicy
	actual    *orgpolicypb.Policy
}

var _ directbase.Adapter = &PolicyAdapter{}

// Find retrieves the GCP resource.
func (a *PolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Policy", "name", a.id)

	req := &orgpolicypb.GetPolicyRequest{Name: a.id.String()}
	policypb, err := a.gcpClient.GetPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Policy %q: %w", a.id, err)
	}

	a.actual = policypb
	return true, nil
}

// Create creates the resource in GCP
func (a *PolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Policy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := a.getDesiredProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &orgpolicypb.CreatePolicyRequest{
		Parent: a.id.Parent().String(),
		Policy: resource,
	}
	created, err := a.gcpClient.CreatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Policy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Policy", "name", a.id)

	status := &krm.ResourceManagerPolicyStatus{}
	status.ObservedState = ResourceManagerPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Etag = direct.LazyPtr(created.GetSpec().GetEtag())
	status.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, created.GetSpec().GetUpdateTime())

	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP
func (a *PolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Policy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := a.getDesiredProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if proto.Equal(a.actual.GetSpec(), desiredPb.GetSpec()) && proto.Equal(a.actual.GetDryRunSpec(), desiredPb.GetDryRunSpec()) {
		log.V(2).Info("Policy is already up to date", "name", a.id)
		status := &krm.ResourceManagerPolicyStatus{}
		status.ObservedState = ResourceManagerPolicyObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.LazyPtr(a.id.String())
		status.Etag = direct.LazyPtr(a.actual.GetSpec().GetEtag())
		status.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, a.actual.GetSpec().GetUpdateTime())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &orgpolicypb.UpdatePolicyRequest{
		Policy: desiredPb,
	}
	req.UpdateMask = &fieldmaskpb.FieldMask{
		Paths: []string{"spec", "dry_run_spec"},
	}

	updated, err := a.gcpClient.UpdatePolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Policy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Policy", "name", a.id)

	status := &krm.ResourceManagerPolicyStatus{}
	status.ObservedState = ResourceManagerPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Etag = direct.LazyPtr(updated.GetSpec().GetEtag())
	status.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, updated.GetSpec().GetUpdateTime())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *PolicyAdapter) getDesiredProto(mapCtx *direct.MapContext, spec *krm.ResourceManagerPolicySpec) *orgpolicypb.Policy {
	res := &orgpolicypb.Policy{}
	res.Name = a.id.String()

	// Handle V2 fields if present
	if len(spec.Rules) > 0 || spec.InheritFromParent != nil || spec.Reset != nil {
		res.Spec = &orgpolicypb.PolicySpec{
			Rules:             direct.Slice_ToProto(mapCtx, spec.Rules, PolicySpec_PolicyRule_ToProto),
			InheritFromParent: direct.ValueOf(spec.InheritFromParent),
			Reset_:            direct.ValueOf(spec.Reset),
		}
	} else {
		// Map legacy V1 fields to V2 rules
		if spec.BooleanPolicy != nil {
			res.Spec = &orgpolicypb.PolicySpec{
				Rules: []*orgpolicypb.PolicySpec_PolicyRule{
					{
						Kind: &orgpolicypb.PolicySpec_PolicyRule_Enforce{
							Enforce: spec.BooleanPolicy.Enforced,
						},
					},
				},
			}
		} else if spec.ListPolicy != nil {
			res.Spec = &orgpolicypb.PolicySpec{
				InheritFromParent: direct.ValueOf(spec.ListPolicy.InheritFromParent),
			}
			rule := &orgpolicypb.PolicySpec_PolicyRule{}
			if spec.ListPolicy.Allow != nil {
				if direct.ValueOf(spec.ListPolicy.Allow.All) {
					rule.Kind = &orgpolicypb.PolicySpec_PolicyRule_AllowAll{AllowAll: true}
				} else {
					rule.Kind = &orgpolicypb.PolicySpec_PolicyRule_Values{
						Values: &orgpolicypb.PolicySpec_PolicyRule_StringValues{
							AllowedValues: spec.ListPolicy.Allow.Values,
						},
					}
				}
			} else if spec.ListPolicy.Deny != nil {
				if direct.ValueOf(spec.ListPolicy.Deny.All) {
					rule.Kind = &orgpolicypb.PolicySpec_PolicyRule_DenyAll{DenyAll: true}
				} else {
					rule.Kind = &orgpolicypb.PolicySpec_PolicyRule_Values{
						Values: &orgpolicypb.PolicySpec_PolicyRule_StringValues{
							DeniedValues: spec.ListPolicy.Deny.Values,
						},
					}
				}
			}
			if rule.Kind != nil {
				res.Spec.Rules = append(res.Spec.Rules, rule)
			}
		} else if spec.RestorePolicy != nil {
			res.Spec = &orgpolicypb.PolicySpec{
				Reset_: spec.RestorePolicy.Default,
			}
		}
	}

	if spec.DryRunSpec != nil {
		res.DryRunSpec = PolicySpec_ToProto(mapCtx, spec.DryRunSpec)
	}

	return res
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	// Exporting back to the V2-style fields in ResourceManagerPolicy
	u := &unstructured.Unstructured{}

	obj := &krm.ResourceManagerPolicy{}
	mapCtx := &direct.MapContext{}

	obj.Spec.Rules = direct.Slice_FromProto(mapCtx, a.actual.GetSpec().GetRules(), PolicySpec_PolicyRule_FromProto)
	obj.Spec.InheritFromParent = direct.LazyPtr(a.actual.GetSpec().GetInheritFromParent())
	obj.Spec.Reset = direct.LazyPtr(a.actual.GetSpec().Reset_)
	obj.Spec.DryRunSpec = PolicySpec_FromProto(mapCtx, a.actual.GetDryRunSpec())
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	parentRef, constraint, err := krm.ParsePolicyExternal(a.actual.Name)
	if err != nil {
		return nil, err
	}
	obj.Spec.Constraint = constraint
	if parentRef.ProjectID != "" {
		obj.Spec.ProjectRef = &refs.ProjectRef{External: parentRef.ProjectID}
	} else if parentRef.FolderID != "" {
		obj.Spec.FolderRef = &refs.FolderRef{External: parentRef.FolderID}
	} else if parentRef.OrganizationID != "" {
		obj.Spec.OrganizationRef = &refs.OrganizationRef{External: parentRef.OrganizationID}
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ResourceManagerPolicyGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Policy", "name", a.id)

	req := &orgpolicypb.DeletePolicyRequest{Name: a.id.String()}
	err := a.gcpClient.DeletePolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Policy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Policy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Policy", "name", a.id)

	return true, nil
}
