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

	"k8s.io/apimachinery/pkg/types"
)

const (
	ConfigConnectorComponentName         = "configconnector"
	WorkloadIdentityAnnotation           = "iam.gke.io/gcp-service-account"
	ServiceAccountNamePrefix             = "cnrm-controller-manager-"
	ControllerManagerPodForClusterMode   = "cnrm-controller-manager-0"
	OperatorFinalizer                    = "configconnector.cnrm.cloud.google.com/finalizer"
	ConfigConnectorContextNamespaceLabel = "configconnectorcontext.cnrm.cloud.google.com/namespace"
	KCCFinalizer                         = "cnrm.cloud.google.com/finalizer"
	KCCSystemLabelSelectorRaw            = "cnrm.cloud.google.com/system"
	KCCSystemComponentLabel              = "cnrm.cloud.google.com/component"
	KCCControllerManagerComponent        = "cnrm-controller-manager"
	KCCUnmanagedDetectorComponent        = "cnrm-unmanaged-detector"
	CNRMDomain                           = "cnrm.cloud.google.com"
	CNRMSystemNamespace                  = "cnrm-system"
	NamespacedComponentLabel             = "cnrm.cloud.google.com/scoped-namespace"
	OperatorSystemNamespace              = "configconnector-operator-system"
	VersionAnnotation                    = "cnrm.cloud.google.com/version"
	OperatorVersionAnnotation            = "cnrm.cloud.google.com/operator-version"
	ProjectIDAnnotation                  = "cnrm.cloud.google.com/project-id"
	StableChannel                        = "stable"
	ConfigConnectorAllowedName           = "configconnector.core.cnrm.cloud.google.com"
	ConfigConnectorContextAllowedName    = "configconnectorcontext.core.cnrm.cloud.google.com"
	UpToDate                             = "UpToDate"
	UpToDateMessage                      = "ConfigConnector is up to date"
	UpdateFailed                         = "UpdateFailed"
	ControllerManagerService             = "cnrm-manager"
	NamespacedManagerServicePrefix       = "cnrm-manager-"
	NamespacedManagerServiceTmpl         = "cnrm-manager-${NAMESPACE?}"
	ServiceAccountProjectPolicy          = "SERVICE_ACCOUNT_PROJECT"
	ResourceProjectPolicy                = "RESOURCE_PROJECT"
	BillingProjectPolicy                 = "BILLING_PROJECT"
	UserProjectOverrideFlag              = "--user-project-override"
	BillingProjectFlag                   = "--billing-project"
	CNRMManagerContainerName             = "manager"
)

var (
	KCCControllerPodLabelSelectorRaw = fmt.Sprintf("%v=%v", KCCSystemComponentLabel, KCCControllerManagerComponent)

	// IgnoredCRDList contains CRDs that should be ignored by the operator.
	IgnoredCRDList = map[string]bool{
		// KCC no longer supports the ServiceMapping CRD as of v1.50.0, but customer
		// clusters may still contain a copy of the CRD and its CRs.
		// Ignore this CRD and its CRs to avoid blocking deletions of ConfigConnectorContext
		// objects on their existence (b/195157239).
		"servicemappings.core.cnrm.cloud.google.com": true,
		// GameServicesRealm CRD is removed from the install bundle as of
		// v1.101.0.
		"gameservicesrealms.gameservices.cnrm.cloud.google.com": true,
	}

	OperatorNamespaceIDConfigMapNN = types.NamespacedName{
		Namespace: OperatorSystemNamespace,
		Name:      "namespace-id",
	}
)
