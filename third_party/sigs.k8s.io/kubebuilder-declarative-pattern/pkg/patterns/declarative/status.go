/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package declarative

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

// Status provides health and readiness information for a given DeclarativeObject
type Status interface {
	Reconciled
	Preflight
	VersionCheck

	// BuildStatus computes the new status for the object after a reconcile operation,
	// and writes it into statusInfo.Subject.
	// This function is should broadly map any reconciliation errors into the status,
	// and if no errors were reported should check that the applied objects are healthy and ready.
	// If this function returns a nil error, changes to the `status` of the object
	// will then be written back to the kube-apiserver.
	// If this function returns an error, the reconciliation function will return
	// an error without updating the status to the apiserver, the interpretation is that
	// we were unable to compute a new status.
	BuildStatus(ctx context.Context, statusInfo *StatusInfo) error
}

// LiveObjectReader exposes the state of objects on the cluster after the apply operation.
// Currently this is done by querying the cluster, but we will move this to record the state of objects after applying them.
type LiveObjectReader func(ctx context.Context, gvk schema.GroupVersionKind, nn types.NamespacedName) (*unstructured.Unstructured, error)

type Reconciled interface {
	// Reconciled is triggered when Reconciliation has occured.
	// The caller is encouraged to determine and surface the health of the reconcilation
	// on the DeclarativeObject.
	//
	// Deprecated: Prefer the BuildStatus method
	Reconciled(ctx context.Context, subject DeclarativeObject, manifest *manifest.Objects, err error) error
}

// BuildStatus computes the new status after a reconcile operation.
type BuildStatus interface {
	// BuildStatus computes the new status after a reconcile operation.
	BuildStatus(ctx context.Context, statusInfo *StatusInfo) error
}

type Preflight interface {
	// Preflight validates if the current state of the world is ready for reconciling.
	// Returning a non-nil error on this object will prevent Reconcile from running.
	// The caller is encouraged to surface the error status on the DeclarativeObject.
	Preflight(context.Context, DeclarativeObject) error
}

type VersionCheck interface {
	// VersionCheck checks if the version of the operator is greater than or equal to the
	// version requested by objects in the manifest, if it isn't it updates the status and
	// events and stops reconciling
	VersionCheck(context.Context, DeclarativeObject, *manifest.Objects) (bool, error)
}

// StatusBuilder provides a pluggable implementation of Status
type StatusBuilder struct {
	// ReconciledImpl is called after reconciliation
	//
	// Deprecated: Prefer the BuildStatus method
	ReconciledImpl Reconciled

	PreflightImpl    Preflight
	VersionCheckImpl VersionCheck

	// BuildStatus computes the status after a reconcile operation
	BuildStatusImpl BuildStatus
}

func (s *StatusBuilder) Reconciled(ctx context.Context, src DeclarativeObject, objs *manifest.Objects, err error) error {
	if s.ReconciledImpl != nil {
		return s.ReconciledImpl.Reconciled(ctx, src, objs, err)
	}
	return nil
}

func (s *StatusBuilder) Preflight(ctx context.Context, src DeclarativeObject) error {
	if s.PreflightImpl != nil {
		return s.PreflightImpl.Preflight(ctx, src)
	}
	return nil
}

func (s *StatusBuilder) VersionCheck(ctx context.Context, src DeclarativeObject, objs *manifest.Objects) (bool, error) {
	if s.VersionCheckImpl != nil {
		return s.VersionCheckImpl.VersionCheck(ctx, src, objs)
	}
	return true, nil
}

func (s *StatusBuilder) BuildStatus(ctx context.Context, statusInfo *StatusInfo) error {
	if s.BuildStatusImpl != nil {
		return s.BuildStatusImpl.BuildStatus(ctx, statusInfo)
	}
	return nil
}

var _ Status = &StatusBuilder{}
