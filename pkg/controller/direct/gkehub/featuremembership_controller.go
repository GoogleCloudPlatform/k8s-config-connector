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
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"

	featureapi "google.golang.org/api/gkehub/v1beta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

const (
	ctrlName        = "gkehubfeaturemembership-controller"
	timeoutDuration = 20 * time.Minute
	baseDelay       = 5 * time.Second
)

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

	desired *krm.GKEHubFeatureMembership
	actual  *featureapi.Feature

	hubClient *gkeHubClient
	reader    client.Reader
}

var _ directbase.Adapter = &gkeHubAdapter{}

// AdapterForObject implements the Model interface.
func (m *gkeHubModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}
	hubClient, err := gcpClient.newGkeHubClient(ctx)
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
	project, err := refs.ResolveProject(ctx, reader, u.GetNamespace(), projectRef)
	if err != nil {
		return nil, err
	}
	projectID := project.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	if err := resolveIAMReferences(ctx, reader, obj); err != nil {
		return nil, err
	}
	return &gkeHubAdapter{
		projectID: projectID,
		location:  obj.Spec.Location,
		desired:   obj,
		hubClient: hubClient,
		reader:    reader,
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
	if &a.desired.Spec.FeatureRef == nil || &a.desired.Spec.MembershipRef == nil {
		return false, nil
	}
	// Resolve featureRef
	featureExternal, err := resolveFeatureRef(ctx, a.reader, a.desired, a.projectID)
	if err != nil {
		return false, err
	}
	a.featureID = featureExternal.id
	feature, err := a.hubClient.featureClient.Get(a.featureID).Context(ctx).ReturnPartialSuccess(true).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting feature %q: %w", a.featureID, err)
	}
	a.actual = feature

	// Build the membershipID(in format projects/*/locations/*/memberships/*) of the target membership.
	// We don't need to resolve the GKEHubMembership resource. Instead, we'll build a membershipID that matches the
	// membershipIDs in the GKEHubFeature's MembershipSpecs list. This allows us to manage the target membership
	// using the GKEHubFeatureMembership configuration, without needing to resolve the resource.
	membershipRef := a.desired.Spec.MembershipRef
	if membershipRef.External != "" {
		// use the provided external value
		a.membershipID = membershipRef.External
	} else {
		// build membershipID by its project, location and name
		if membershipLocation := a.desired.Spec.MembershipLocation; membershipLocation != nil {
			a.membershipID = fmt.Sprintf("projects/%s/locations/%s/memberships/%s", a.projectID, common.ValueOf(membershipLocation), membershipRef.Name)
		} else {
			a.membershipID = fmt.Sprintf("projects/%s/locations/global/memberships/%s", a.projectID, membershipRef.Name)
		}
	}

	// Find target membershipID from feature's MembershipSpecs
	canonicalizedMID, found, err := matchWithCanonicalMemebrshipID(a.membershipID, feature)
	if err != nil {
		return false, nil
	}
	if canonicalizedMID != "" {
		a.membershipID = canonicalizedMID
	}
	return found, nil
}

// Delete implements the Adapter interface.
func (a *gkeHubAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// emptying the membershipspec is sufficient
	a.desired = &krm.GKEHubFeatureMembership{}
	if _, err := a.patchMembershipSpec(ctx); err != nil {
		return false, fmt.Errorf("deleting membershipspec for %s: %w", a.membershipID, err)
	}
	return true, nil
}

func (a *gkeHubAdapter) patchMembershipSpec(ctx context.Context) ([]byte, error) {
	var feature *featureapi.Feature
	if a.actual != nil {
		feature = a.actual
	} else {
		// if the feature does not exist, create a new one.
		feature = &featureapi.Feature{
			Name: fmt.Sprintf("projects/%s/locations/%s/features/%s", a.projectID, a.location, a.featureID),
		}
	}
	mSpecs := make(map[string]featureapi.MembershipFeatureSpec)
	// only change the feature configuration for the associated membership
	desiredApiObj, err := featureMembershipSpecKRMtoMembershipFeatureSpecAPI(&a.desired.Spec)
	if err != nil {
		return nil, err
	}
	mSpecs[a.membershipID] = *desiredApiObj
	// MembershipSpecs is a map of membership spec. Here we only patch one membership.
	// GKE Hub server doesn't patch other memberships if they are not present in the membershipSpecs map.
	feature.MembershipSpecs = mSpecs
	op, err := a.hubClient.featureClient.Patch(a.featureID, feature).UpdateMask("membershipSpecs").Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return nil, fmt.Errorf("failed to wait for the operation: %w", err)
	}
	return op.Response, nil
}

func (a *gkeHubAdapter) waitForOp(ctx context.Context, op *featureapi.Operation) error {
	retryPeriod := baseDelay
	timeoutAt := time.Now().Add(timeoutDuration)
	for {
		current, err := a.hubClient.operationClient.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation status of %q failed: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error)
			} else {
				return nil
			}
		}
		if time.Now().After(timeoutAt) {
			return fmt.Errorf("operation timed out waiting for LRO after %s", timeoutDuration.String())
		}
		time.Sleep(retryPeriod)
		retryPeriod = retryPeriod * 2
	}
}

func (a *gkeHubAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating gkehubfeaturemembership", "obj", u)

	// todo: shall we resolve GKEHubMembership here to ensure its existence?
	// Can we add membershipSpec to GKEHubFeature, if the membership does not exist?

	_, err := a.patchMembershipSpec(ctx)
	if err != nil {
		return fmt.Errorf("failed to patch the MembershipSpec; %w", err)
	}
	log.V(2).Info("successfully created gkehubfeaturemembership")
	// no need to set the status from the api response for the  &krm.GKEHubFeatureMembershipStatus{} as the it only has generic status.
	return nil
}

func (a *gkeHubAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating object", "u", u)

	// todo: shall we resolve GKEHubMembership here to ensure its existence?
	// Can we update membershipSpec in GKEHubFeature's membershipSpecs field, if the membership does not exist?

	actual := a.actual.MembershipSpecs[a.membershipID]
	//  There are no output fields in the api Object, so we can compare the desired and the actaul directly.
	if len(diffFeatureMembership(&a.desired.Spec, &actual)) != 0 {
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
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEHubFeatureMembership{}
	mapCtx := &direct.MapContext{}
	m := a.actual.MembershipSpecs[a.membershipID]
	obj.Spec = direct.ValueOf(GKEHubFeatureMembershipSpec_FromProto(mapCtx, &m))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = krm.FeatureProjectRef{Name: a.projectID}
	obj.Spec.Location = a.location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// mID is in the format of "projects/{ProjectID}/locations/*/memberships/{membershipId}".
// matchWithCanonicalMemebrshipID matches the keys in the feature.membershipspec map, which is in the format of "projects/{ProjectNumber}/locations/*/memberships/{membershipId}".
func matchWithCanonicalMemebrshipID(mID string, feature *featureapi.Feature) (string, bool, error) {
	if feature.MembershipSpecs == nil {
		return "", false, nil
	}
	tokens := strings.Split(mID, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "memberships" {
		return "", false, fmt.Errorf("format of membership ID=%q was not known (use projects/*/locations/*/memberships/{membershipId}) ", mID)
	}
	suffix := strings.Join(tokens[2:], "/")
	for k := range feature.MembershipSpecs {
		if strings.HasSuffix(k, suffix) {
			return k, true, nil
		}
	}
	return "", false, nil
}
