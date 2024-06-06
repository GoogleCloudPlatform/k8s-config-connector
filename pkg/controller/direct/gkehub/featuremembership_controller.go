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

package gkehub

import (
	"context"
	"fmt"
	"reflect"

	featureapi "google.golang.org/api/gkehub/v1beta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

const ctrlName = "gkehubfeaturemembership-controller"

func init() {
	registry.RegisterModel(krm.GKEHubFeatureMembershipGVK, getGkeHubModel)
}

func getGkeHubModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gkeHubModel{config: config}, nil
}

type gkeHubModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &gkeHubModel{}

type gkeHubAdapter struct {
	membershipID string
	featureID    string
	projectID    string
	location     string

	desired *featureapi.MembershipFeatureSpec
	actual  *featureapi.Feature

	featureClient *featureapi.ProjectsLocationsFeaturesService
}

var _ directbase.Adapter = &gkeHubAdapter{}

// AdapterForObject implements the Model interface.
func (m *gkeHubModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}

	projectsLocationsFeaturesService, err := gcpClient.newProjectsLocationsFeaturesService(ctx)
	if err != nil {
		return nil, err
	}
	obj := &krm.GKEHubFeatureMembership{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	projectRef := &refs.ProjectRef{
		Name:      obj.Spec.ProjectRef.Name,
		Namespace: obj.Spec.ProjectRef.Namespace,
		External:  obj.Spec.ProjectRef.External,
	}
	project, err := refs.ResolveProject(ctx, reader, obj, projectRef)
	if err != nil {
		return nil, err
	}
	projectID := project.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	membership, err := resolveMembershipRef(ctx, reader, obj, projectID)
	if err != nil {
		return nil, err
	}
	feature, err := resolveFeatureRef(ctx, reader, obj, projectID)
	if err != nil {
		return nil, err
	}
	if err := resolveIAMReferences(ctx, reader, obj); err != nil {
		return nil, err
	}
	apiObj, err := featureMembershipSpecKRMtoMembershipFeatureSpecAPI(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return &gkeHubAdapter{
		membershipID:  membership.id,
		featureID:     feature.id,
		projectID:     projectID,
		location:      obj.Spec.Location,
		desired:       apiObj,
		featureClient: projectsLocationsFeaturesService,
	}, nil
}

func (m *gkeHubModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func resolveIAMReferences(ctx context.Context, reader client.Reader, obj *krm.GKEHubFeatureMembership) error {
	spec := obj.Spec
	if spec.Configmanagement != nil && spec.Configmanagement.ConfigSync != nil {
		if err := spec.Configmanagement.ConfigSync.MetricsGcpServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
			return err
		}
		if spec.Configmanagement.ConfigSync.Git != nil {
			if err := spec.Configmanagement.ConfigSync.Git.GcpServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
				return err
			}
		}
		if spec.Configmanagement.ConfigSync.Oci != nil {
			if err := spec.Configmanagement.ConfigSync.Oci.GcpServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *gkeHubAdapter) Find(ctx context.Context) (bool, error) {
	if a.membershipID == "" || a.featureID == "" {
		return false, nil
	}
	feature, err := a.featureClient.Get(a.featureID).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting feature %q: %w", a.featureID, err)
	}
	// Custom diff handling for hierarchyController(HNC), Hub API will always return {} for all false fields.
	// So here, we convert {} to all false.
	if feature.MembershipSpecs != nil {
		if mSpec, ok := feature.MembershipSpecs[a.membershipID]; ok {
			if mSpec.Configmanagement != nil || mSpec.Configmanagement.HierarchyController != nil {
				if reflect.DeepEqual(mSpec.Configmanagement.HierarchyController, featureapi.ConfigManagementHierarchyControllerConfig{}) {
					mSpec.Configmanagement.HierarchyController.EnableHierarchicalResourceQuota = false
					mSpec.Configmanagement.HierarchyController.EnablePodTreeLabels = false
					mSpec.Configmanagement.HierarchyController.Enabled = false
				}
			}
		}
	}
	a.actual = feature
	return true, nil
}

// Delete implements the Adapter interface.
func (a *gkeHubAdapter) Delete(ctx context.Context) (bool, error) {
	exist, err := a.Find(ctx)
	if err != nil {
		return false, fmt.Errorf("finding feature for %s:%w", a.featureID, err)
	}
	if !exist {
		// return (false, nil) if the object was not found but should be presumed deleted.
		return false, nil
	}
	// emptying the membershipspec is sufficient
	a.desired = &featureapi.MembershipFeatureSpec{}
	if _, err := a.patchMembershipSpec(ctx); err != nil {
		return false, fmt.Errorf("deleting membershipspec for %s: %w", a.membershipID, err)
	}
	return true, nil
}

func (a *gkeHubAdapter) patchMembershipSpec(ctx context.Context) ([]byte, error) {
	feature := a.actual
	// only change the feature configuration for the associated membership
	feature.MembershipSpecs[a.membershipID] = *a.desired
	op, err := a.featureClient.Patch(a.featureID, feature).UpdateMask("membershipSpecs").Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return op.Response, nil
}

func (a *gkeHubAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating gkehubfeaturemembership", "obj", u)

	_, err := a.patchMembershipSpec(ctx)
	if err != nil {
		return fmt.Errorf("failed to patch the MembershipSpec; %w", err)
	}
	log.V(2).Info("successfully created gkehubfeaturemembership")
	// no need to set the status from the api response for the  &krm.GKEHubFeatureMembershipStatus{} as the it only has generic status.
	return nil
}

func (a *gkeHubAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating object", "u", u)
	actual := a.actual.MembershipSpecs[a.membershipID]
	//  There are no output fields in the api Object, so we can compare the desired and the actaul directly.
	if !reflect.DeepEqual(a.desired.Configmanagement, actual.Configmanagement) || !reflect.DeepEqual(a.desired.Policycontroller, actual.Policycontroller) || !reflect.DeepEqual(a.desired.Mesh, actual.Mesh) {
		log.V(2).Info("diff detected, patching gkehubfeaturemembership")
		if _, err := a.patchMembershipSpec(ctx); err != nil {
			return fmt.Errorf("patching gkehubfeaturemembership failed: %w", err)
		}
		log.V(2).Info("successfully updated gkehubfeaturemembership")
	} else {
		log.V(2).Info("no diff, skipping updating gkehubfeaturemembership")
	}
	// no need to set the status from the api response for &krm.GKEHubFeatureMembershipStatus{} as the it only has generic status.
	return nil
}

func (a *gkeHubAdapter) Export(context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}
