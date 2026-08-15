/*
Copyright 2022 The Kubernetes Authors.

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

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

// Hook is the base interface implemented by a hook
type Hook interface {
}

// ApplyOperation contains the details of an Apply operation
type ApplyOperation struct {
	// Subject is the object we are reconciling
	Subject DeclarativeObject

	// Objects is the set of objects we are applying
	Objects *manifest.Objects

	// ApplierOptions is the set of options passed to the applier
	ApplierOptions *applier.ApplierOptions
}

// AfterApply is implemented by hooks that want to be called after every apply operation
type AfterApply interface {
	AfterApply(ctx context.Context, op *ApplyOperation) error
}

// BeforeApply is implemented by hooks that want to be called before every apply operation
type BeforeApply interface {
	BeforeApply(ctx context.Context, op *ApplyOperation) error
}

// UpdateStatusOperation contains the details of an Apply operation
type UpdateStatusOperation struct {
	// Subject is the object we are reconciling
	Subject DeclarativeObject
}

// AfterUpdateStatus is implemented by hooks that want to be called after the update-status phase
type AfterUpdateStatus interface {
	AfterUpdateStatus(ctx context.Context, op *UpdateStatusOperation) error
}

// BeforeUpdateStatus is implemented by hooks that want to be called before the update-status phase
type BeforeUpdateStatus interface {
	BeforeUpdateStatus(ctx context.Context, op *UpdateStatusOperation) error
}
