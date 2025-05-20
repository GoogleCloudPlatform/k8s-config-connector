// Copyright 2022 Google LLC
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

package k8s

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/types"
)

// TODO: clean up old conditions used in handcrafted controllers
const (
	CNRMGroup                            = "cnrm.cloud.google.com"
	CNRMTestGroup                        = "test.cnrm.cloud.google.com"
	APIDomainSuffix                      = ".cnrm.cloud.google.com"
	SystemNamespace                      = "cnrm-system"
	ControllerManagerNamePrefix          = "cnrm-controller-manager"
	ControllerMaxConcurrentReconciles    = 20
	ReconcileDeadline                    = 59 * time.Minute
	TimeToLeaseExpiration                = 40 * time.Minute
	TimeToLeaseRenewal                   = 20 * time.Minute
	MeanReconcileReenqueuePeriod         = 10 * time.Minute
	JitterFactor                         = 2.0
	UpToDate                             = "UpToDate"
	UpToDateMessage                      = "The resource is up to date"
	Created                              = "Created"
	Creating                             = "Creating"
	CreatingMessage                      = "The resource is under creation"
	CreatedMessage                       = "Successfully created"
	CreateFailed                         = "CreateFailed"
	CreateFailedMessageTmpl              = "Create call failed: %v"
	Updating                             = "Updating"
	UpdatingMessage                      = "Update in progress"
	UpdateFailed                         = "UpdateFailed"
	Deleting                             = "Deleting"
	DeletingMessage                      = "Deletion in progress"
	Deleted                              = "Deleted"
	DeletedMessage                       = "Successfully deleted"
	DeleteFailed                         = "DeleteFailed"
	NoCondition                          = "NoCondition"
	DeleteFailedMessageTmpl              = "Delete call failed: %v"
	Unmanaged                            = "Unmanaged"
	UnmanagedMessageTmpl                 = "No controller is managing this resource. Check if a ConfigConnectorContext exists for resource's namespace, '%v'"
	ControllerFinalizerName              = "cnrm.cloud.google.com/finalizer"
	DeletionDefenderFinalizerName        = "cnrm.cloud.google.com/deletion-defender"
	DependencyNotReady                   = "DependencyNotReady"
	DependencyNotFound                   = "DependencyNotFound"
	DependencyInvalid                    = "DependencyInvalid"
	ManagementConflict                   = "ManagementConflict"
	PreActuationTransformFailed          = "PreActuationTransformFailed"
	PostActuationTransformFailed         = "PostActuationTransformFailed"
	DeletionPolicyDelete                 = "delete"
	DeletionPolicyAbandon                = "abandon"
	AnnotationPrefix                     = CNRMGroup
	NamespaceEnvVar                      = "NAMESPACE"
	ImmediateReconcileRequestsBufferSize = 10000
	MaxNumResourceWatcherRoutines        = 10000

	ControllerManagedFieldManager = "cnrm-controller-manager"
	UnmanagedDetectorFieldManager = "cnrm-unmanaged-detector"
	SupportsSSAManager            = "supports-ssa"

	// State into spec annotation values
	StateIntoSpecAnnotation = "cnrm.cloud.google.com/state-into-spec"
	StateMergeIntoSpec      = "merge"
	StateAbsentInSpec       = "absent"

	// Core kubernetes constants
	LastAppliedConfigurationAnnotation = "kubectl.kubernetes.io/last-applied-configuration"
	ManagedFieldsTypeFieldsV1          = "FieldsV1"

	ResourceIDFieldName = "resourceID"
	ResourceIDFieldPath = "spec." + ResourceIDFieldName

	// selfLink may not present in every KRM resource status.
	SelfLinkFieldName      = "selfLink"
	ObservedStateFieldName = "observedState"

	StabilityLevelStable = "stable"
	StabilityLevelAlpha  = "alpha"

	KCCAPIVersionV1Beta1  = "v1beta1"
	KCCAPIVersionV1Alpha1 = "v1alpha1"
)

var (
	DeletionPolicyAnnotation             = FormatAnnotation("deletion-policy")
	ReconcileIntervalInSecondsAnnotation = FormatAnnotation("reconcile-interval-in-seconds")

	// Annotations for Container objects
	ProjectIDAnnotation  = FormatAnnotation("project-id")
	FolderIDAnnotation   = FormatAnnotation("folder-id")
	OrgIDAnnotation      = FormatAnnotation("organization-id")
	ContainerAnnotations = []string{
		ProjectIDAnnotation,
		FolderIDAnnotation,
		OrgIDAnnotation,
	}

	// Internal Annotation to force reconciliation
	InternalForceReconcileAnnotation = CNRMTestGroup + "/reconcile-cookie"

	KCCComponentLabel    = FormatAnnotation("component")
	KCCSystemLabel       = FormatAnnotation("system")
	KCCVersionLabel      = FormatAnnotation("version")
	ScopedNamespaceLabel = FormatAnnotation("scoped-namespace")
	DCL2CRDLabel         = FormatAnnotation("dcl2crd")
	TF2CRDLabel          = FormatAnnotation("tf2crd")
	KCCStabilityLabel    = FormatAnnotation("stability-level")

	MutableButUnreadableFieldsAnnotation = FormatAnnotation("mutable-but-unreadable-fields")
	ObservedSecretVersionsAnnotation     = FormatAnnotation("observed-secret-versions")

	SupportsSSAAnnotation = FormatAnnotation("supports-ssa")

	BlueprintAttributionAnnotation = FormatAnnotation("blueprint")

	// TODO(kcc-eng): Adjust the timeout back down after b/237398742 is fixed.
	WebhookTimeoutSeconds = int32(10)

	ReservedStatusFieldNamesForFutureUse = []string{"generation"}

	NamespaceIDConfigMapNN = types.NamespacedName{
		Namespace: SystemNamespace,
		Name:      "namespace-id",
	}

	// IgnoredKindList contains special or deprecated Kinds that should be
	// ignored by the controllers.
	IgnoredKindList = map[string]bool{
		// ServiceMapping is a special resource type that does not make a call
		// to an underlying GCP API.
		// In addition, KCC no longer supports ServiceMapping CRD as of v1.50.0.
		"ServiceMapping": true,
		// KCC no longer supports GameServicesRealm CRD as of v1.101.0.
		"GameServicesRealm": true,
	}
)

func FormatAnnotation(annotationName string) string {
	return fmt.Sprintf("%v/%v", AnnotationPrefix, annotationName)
}
