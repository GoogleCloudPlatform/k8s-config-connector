// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package beta

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type FeatureMembership struct {
	Mesh               *FeatureMembershipMesh             `json:"mesh"`
	Configmanagement   *FeatureMembershipConfigmanagement `json:"configmanagement"`
	Policycontroller   *FeatureMembershipPolicycontroller `json:"policycontroller"`
	Project            *string                            `json:"project"`
	Location           *string                            `json:"location"`
	Feature            *string                            `json:"feature"`
	Membership         *string                            `json:"membership"`
	MembershipLocation *string                            `json:"membershipLocation"`
}

func (r *FeatureMembership) String() string {
	return dcl.SprintResource(r)
}

// The enum FeatureMembershipMeshManagementEnum.
type FeatureMembershipMeshManagementEnum string

// FeatureMembershipMeshManagementEnumRef returns a *FeatureMembershipMeshManagementEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureMembershipMeshManagementEnumRef(s string) *FeatureMembershipMeshManagementEnum {
	v := FeatureMembershipMeshManagementEnum(s)
	return &v
}

func (v FeatureMembershipMeshManagementEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MANAGEMENT_UNSPECIFIED", "MANAGEMENT_AUTOMATIC", "MANAGEMENT_MANUAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureMembershipMeshManagementEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FeatureMembershipMeshControlPlaneEnum.
type FeatureMembershipMeshControlPlaneEnum string

// FeatureMembershipMeshControlPlaneEnumRef returns a *FeatureMembershipMeshControlPlaneEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureMembershipMeshControlPlaneEnumRef(s string) *FeatureMembershipMeshControlPlaneEnum {
	v := FeatureMembershipMeshControlPlaneEnum(s)
	return &v
}

func (v FeatureMembershipMeshControlPlaneEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"CONTROL_PLANE_MANAGEMENT_UNSPECIFIED", "AUTOMATIC", "MANUAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureMembershipMeshControlPlaneEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum.
type FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum string

// FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumRef returns a *FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumRef(s string) *FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	v := FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(s)
	return &v
}

func (v FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MONITORING_BACKEND_UNSPECIFIED", "PROMETHEUS", "CLOUD_MONITORING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum.
type FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum string

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumRef returns a *FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumRef(s string) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum {
	v := FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(s)
	return &v
}

func (v FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INSTALL_SPEC_UNSPECIFIED", "INSTALL_SPEC_NOT_INSTALLED", "INSTALL_SPEC_ENABLED", "INSTALL_SPEC_SUSPENDED", "INSTALL_SPEC_DETACHED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum.
type FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum string

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumRef returns a *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumRef(s string) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum {
	v := FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(s)
	return &v
}

func (v FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MONITORING_BACKEND_UNSPECIFIED", "PROMETHEUS", "CLOUD_MONITORING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum.
type FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum string

// FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumRef returns a *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum with the value of string s
// If the empty string is provided, nil is returned.
func FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumRef(s string) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum {
	v := FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(s)
	return &v
}

func (v FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INSTALLATION_UNSPECIFIED", "NOT_INSTALLED", "ALL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type FeatureMembershipMesh struct {
	empty        bool                                   `json:"-"`
	Management   *FeatureMembershipMeshManagementEnum   `json:"management"`
	ControlPlane *FeatureMembershipMeshControlPlaneEnum `json:"controlPlane"`
}

type jsonFeatureMembershipMesh FeatureMembershipMesh

func (r *FeatureMembershipMesh) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipMesh
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipMesh
	} else {

		r.Management = res.Management

		r.ControlPlane = res.ControlPlane

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipMesh is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipMesh *FeatureMembershipMesh = &FeatureMembershipMesh{empty: true}

func (r *FeatureMembershipMesh) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipMesh) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipMesh) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipConfigmanagement struct {
	empty               bool                                                  `json:"-"`
	ConfigSync          *FeatureMembershipConfigmanagementConfigSync          `json:"configSync"`
	PolicyController    *FeatureMembershipConfigmanagementPolicyController    `json:"policyController"`
	Binauthz            *FeatureMembershipConfigmanagementBinauthz            `json:"binauthz"`
	HierarchyController *FeatureMembershipConfigmanagementHierarchyController `json:"hierarchyController"`
	Version             *string                                               `json:"version"`
}

type jsonFeatureMembershipConfigmanagement FeatureMembershipConfigmanagement

func (r *FeatureMembershipConfigmanagement) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipConfigmanagement
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipConfigmanagement
	} else {

		r.ConfigSync = res.ConfigSync

		r.PolicyController = res.PolicyController

		r.Binauthz = res.Binauthz

		r.HierarchyController = res.HierarchyController

		r.Version = res.Version

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipConfigmanagement is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipConfigmanagement *FeatureMembershipConfigmanagement = &FeatureMembershipConfigmanagement{empty: true}

func (r *FeatureMembershipConfigmanagement) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipConfigmanagement) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipConfigmanagement) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipConfigmanagementConfigSync struct {
	empty                         bool                                            `json:"-"`
	Git                           *FeatureMembershipConfigmanagementConfigSyncGit `json:"git"`
	SourceFormat                  *string                                         `json:"sourceFormat"`
	PreventDrift                  *bool                                           `json:"preventDrift"`
	MetricsGcpServiceAccountEmail *string                                         `json:"metricsGcpServiceAccountEmail"`
	Oci                           *FeatureMembershipConfigmanagementConfigSyncOci `json:"oci"`
}

type jsonFeatureMembershipConfigmanagementConfigSync FeatureMembershipConfigmanagementConfigSync

func (r *FeatureMembershipConfigmanagementConfigSync) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipConfigmanagementConfigSync
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipConfigmanagementConfigSync
	} else {

		r.Git = res.Git

		r.SourceFormat = res.SourceFormat

		r.PreventDrift = res.PreventDrift

		r.MetricsGcpServiceAccountEmail = res.MetricsGcpServiceAccountEmail

		r.Oci = res.Oci

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipConfigmanagementConfigSync is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipConfigmanagementConfigSync *FeatureMembershipConfigmanagementConfigSync = &FeatureMembershipConfigmanagementConfigSync{empty: true}

func (r *FeatureMembershipConfigmanagementConfigSync) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipConfigmanagementConfigSync) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipConfigmanagementConfigSync) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipConfigmanagementConfigSyncGit struct {
	empty                  bool    `json:"-"`
	SyncRepo               *string `json:"syncRepo"`
	SyncBranch             *string `json:"syncBranch"`
	PolicyDir              *string `json:"policyDir"`
	SyncWaitSecs           *string `json:"syncWaitSecs"`
	SyncRev                *string `json:"syncRev"`
	SecretType             *string `json:"secretType"`
	HttpsProxy             *string `json:"httpsProxy"`
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail"`
}

type jsonFeatureMembershipConfigmanagementConfigSyncGit FeatureMembershipConfigmanagementConfigSyncGit

func (r *FeatureMembershipConfigmanagementConfigSyncGit) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipConfigmanagementConfigSyncGit
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipConfigmanagementConfigSyncGit
	} else {

		r.SyncRepo = res.SyncRepo

		r.SyncBranch = res.SyncBranch

		r.PolicyDir = res.PolicyDir

		r.SyncWaitSecs = res.SyncWaitSecs

		r.SyncRev = res.SyncRev

		r.SecretType = res.SecretType

		r.HttpsProxy = res.HttpsProxy

		r.GcpServiceAccountEmail = res.GcpServiceAccountEmail

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipConfigmanagementConfigSyncGit is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipConfigmanagementConfigSyncGit *FeatureMembershipConfigmanagementConfigSyncGit = &FeatureMembershipConfigmanagementConfigSyncGit{empty: true}

func (r *FeatureMembershipConfigmanagementConfigSyncGit) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipConfigmanagementConfigSyncGit) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipConfigmanagementConfigSyncGit) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipConfigmanagementConfigSyncOci struct {
	empty                  bool    `json:"-"`
	SyncRepo               *string `json:"syncRepo"`
	PolicyDir              *string `json:"policyDir"`
	SyncWaitSecs           *string `json:"syncWaitSecs"`
	SecretType             *string `json:"secretType"`
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail"`
}

type jsonFeatureMembershipConfigmanagementConfigSyncOci FeatureMembershipConfigmanagementConfigSyncOci

func (r *FeatureMembershipConfigmanagementConfigSyncOci) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipConfigmanagementConfigSyncOci
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipConfigmanagementConfigSyncOci
	} else {

		r.SyncRepo = res.SyncRepo

		r.PolicyDir = res.PolicyDir

		r.SyncWaitSecs = res.SyncWaitSecs

		r.SecretType = res.SecretType

		r.GcpServiceAccountEmail = res.GcpServiceAccountEmail

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipConfigmanagementConfigSyncOci is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipConfigmanagementConfigSyncOci *FeatureMembershipConfigmanagementConfigSyncOci = &FeatureMembershipConfigmanagementConfigSyncOci{empty: true}

func (r *FeatureMembershipConfigmanagementConfigSyncOci) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipConfigmanagementConfigSyncOci) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipConfigmanagementConfigSyncOci) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipConfigmanagementPolicyController struct {
	empty                    bool                                                         `json:"-"`
	Enabled                  *bool                                                        `json:"enabled"`
	ExemptableNamespaces     []string                                                     `json:"exemptableNamespaces"`
	ReferentialRulesEnabled  *bool                                                        `json:"referentialRulesEnabled"`
	LogDeniesEnabled         *bool                                                        `json:"logDeniesEnabled"`
	MutationEnabled          *bool                                                        `json:"mutationEnabled"`
	Monitoring               *FeatureMembershipConfigmanagementPolicyControllerMonitoring `json:"monitoring"`
	TemplateLibraryInstalled *bool                                                        `json:"templateLibraryInstalled"`
	AuditIntervalSeconds     *string                                                      `json:"auditIntervalSeconds"`
}

type jsonFeatureMembershipConfigmanagementPolicyController FeatureMembershipConfigmanagementPolicyController

func (r *FeatureMembershipConfigmanagementPolicyController) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipConfigmanagementPolicyController
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipConfigmanagementPolicyController
	} else {

		r.Enabled = res.Enabled

		r.ExemptableNamespaces = res.ExemptableNamespaces

		r.ReferentialRulesEnabled = res.ReferentialRulesEnabled

		r.LogDeniesEnabled = res.LogDeniesEnabled

		r.MutationEnabled = res.MutationEnabled

		r.Monitoring = res.Monitoring

		r.TemplateLibraryInstalled = res.TemplateLibraryInstalled

		r.AuditIntervalSeconds = res.AuditIntervalSeconds

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipConfigmanagementPolicyController is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipConfigmanagementPolicyController *FeatureMembershipConfigmanagementPolicyController = &FeatureMembershipConfigmanagementPolicyController{empty: true}

func (r *FeatureMembershipConfigmanagementPolicyController) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipConfigmanagementPolicyController) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipConfigmanagementPolicyController) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipConfigmanagementPolicyControllerMonitoring struct {
	empty    bool                                                                      `json:"-"`
	Backends []FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum `json:"backends"`
}

type jsonFeatureMembershipConfigmanagementPolicyControllerMonitoring FeatureMembershipConfigmanagementPolicyControllerMonitoring

func (r *FeatureMembershipConfigmanagementPolicyControllerMonitoring) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipConfigmanagementPolicyControllerMonitoring
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipConfigmanagementPolicyControllerMonitoring
	} else {

		r.Backends = res.Backends

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipConfigmanagementPolicyControllerMonitoring is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipConfigmanagementPolicyControllerMonitoring *FeatureMembershipConfigmanagementPolicyControllerMonitoring = &FeatureMembershipConfigmanagementPolicyControllerMonitoring{empty: true}

func (r *FeatureMembershipConfigmanagementPolicyControllerMonitoring) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipConfigmanagementPolicyControllerMonitoring) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipConfigmanagementPolicyControllerMonitoring) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipConfigmanagementBinauthz struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

type jsonFeatureMembershipConfigmanagementBinauthz FeatureMembershipConfigmanagementBinauthz

func (r *FeatureMembershipConfigmanagementBinauthz) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipConfigmanagementBinauthz
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipConfigmanagementBinauthz
	} else {

		r.Enabled = res.Enabled

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipConfigmanagementBinauthz is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipConfigmanagementBinauthz *FeatureMembershipConfigmanagementBinauthz = &FeatureMembershipConfigmanagementBinauthz{empty: true}

func (r *FeatureMembershipConfigmanagementBinauthz) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipConfigmanagementBinauthz) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipConfigmanagementBinauthz) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipConfigmanagementHierarchyController struct {
	empty                           bool  `json:"-"`
	Enabled                         *bool `json:"enabled"`
	EnablePodTreeLabels             *bool `json:"enablePodTreeLabels"`
	EnableHierarchicalResourceQuota *bool `json:"enableHierarchicalResourceQuota"`
}

type jsonFeatureMembershipConfigmanagementHierarchyController FeatureMembershipConfigmanagementHierarchyController

func (r *FeatureMembershipConfigmanagementHierarchyController) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipConfigmanagementHierarchyController
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipConfigmanagementHierarchyController
	} else {

		r.Enabled = res.Enabled

		r.EnablePodTreeLabels = res.EnablePodTreeLabels

		r.EnableHierarchicalResourceQuota = res.EnableHierarchicalResourceQuota

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipConfigmanagementHierarchyController is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipConfigmanagementHierarchyController *FeatureMembershipConfigmanagementHierarchyController = &FeatureMembershipConfigmanagementHierarchyController{empty: true}

func (r *FeatureMembershipConfigmanagementHierarchyController) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipConfigmanagementHierarchyController) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipConfigmanagementHierarchyController) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipPolicycontroller struct {
	empty                     bool                                                        `json:"-"`
	Version                   *string                                                     `json:"version"`
	PolicyControllerHubConfig *FeatureMembershipPolicycontrollerPolicyControllerHubConfig `json:"policyControllerHubConfig"`
}

type jsonFeatureMembershipPolicycontroller FeatureMembershipPolicycontroller

func (r *FeatureMembershipPolicycontroller) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipPolicycontroller
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipPolicycontroller
	} else {

		r.Version = res.Version

		r.PolicyControllerHubConfig = res.PolicyControllerHubConfig

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipPolicycontroller is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipPolicycontroller *FeatureMembershipPolicycontroller = &FeatureMembershipPolicycontroller{empty: true}

func (r *FeatureMembershipPolicycontroller) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipPolicycontroller) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipPolicycontroller) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipPolicycontrollerPolicyControllerHubConfig struct {
	empty                    bool                                                                       `json:"-"`
	InstallSpec              *FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum `json:"installSpec"`
	ExemptableNamespaces     []string                                                                   `json:"exemptableNamespaces"`
	ReferentialRulesEnabled  *bool                                                                      `json:"referentialRulesEnabled"`
	LogDeniesEnabled         *bool                                                                      `json:"logDeniesEnabled"`
	MutationEnabled          *bool                                                                      `json:"mutationEnabled"`
	Monitoring               *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring      `json:"monitoring"`
	AuditIntervalSeconds     *int64                                                                     `json:"auditIntervalSeconds"`
	ConstraintViolationLimit *int64                                                                     `json:"constraintViolationLimit"`
	PolicyContent            *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent   `json:"policyContent"`
}

type jsonFeatureMembershipPolicycontrollerPolicyControllerHubConfig FeatureMembershipPolicycontrollerPolicyControllerHubConfig

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfig) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipPolicycontrollerPolicyControllerHubConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfig
	} else {

		r.InstallSpec = res.InstallSpec

		r.ExemptableNamespaces = res.ExemptableNamespaces

		r.ReferentialRulesEnabled = res.ReferentialRulesEnabled

		r.LogDeniesEnabled = res.LogDeniesEnabled

		r.MutationEnabled = res.MutationEnabled

		r.Monitoring = res.Monitoring

		r.AuditIntervalSeconds = res.AuditIntervalSeconds

		r.ConstraintViolationLimit = res.ConstraintViolationLimit

		r.PolicyContent = res.PolicyContent

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipPolicycontrollerPolicyControllerHubConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfig *FeatureMembershipPolicycontrollerPolicyControllerHubConfig = &FeatureMembershipPolicycontrollerPolicyControllerHubConfig{empty: true}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfig) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring struct {
	empty    bool                                                                               `json:"-"`
	Backends []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum `json:"backends"`
}

type jsonFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	} else {

		r.Backends = res.Backends

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{empty: true}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent struct {
	empty           bool                                                                                    `json:"-"`
	TemplateLibrary *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary `json:"templateLibrary"`
}

type jsonFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
	} else {

		r.TemplateLibrary = res.TemplateLibrary

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{empty: true}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary struct {
	empty        bool                                                                                                    `json:"-"`
	Installation *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum `json:"installation"`
}

type jsonFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) UnmarshalJSON(data []byte) error {
	var res jsonFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary
	} else {

		r.Installation = res.Installation

	}
	return nil
}

// This object is used to assert a desired state where this FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{empty: true}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) Empty() bool {
	return r.empty
}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) String() string {
	return dcl.SprintResource(r)
}

func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *FeatureMembership) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "gke_hub",
		Type:    "FeatureMembership",
		Version: "beta",
	}
}

func (r *FeatureMembership) ID() (string, error) {
	if err := extractFeatureMembershipFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"mesh":                dcl.ValueOrEmptyString(nr.Mesh),
		"configmanagement":    dcl.ValueOrEmptyString(nr.Configmanagement),
		"policycontroller":    dcl.ValueOrEmptyString(nr.Policycontroller),
		"project":             dcl.ValueOrEmptyString(nr.Project),
		"location":            dcl.ValueOrEmptyString(nr.Location),
		"feature":             dcl.ValueOrEmptyString(nr.Feature),
		"membership":          dcl.ValueOrEmptyString(nr.Membership),
		"membership_location": dcl.ValueOrEmptyString(nr.MembershipLocation),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/features/{{feature}}/memberships/{{membership}}", params), nil
}

const FeatureMembershipMaxPage = -1

type FeatureMembershipList struct {
	Items []*FeatureMembership

	nextToken string

	resource *FeatureMembership
}

func (c *Client) DeleteFeatureMembership(ctx context.Context, r *FeatureMembership) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("FeatureMembership resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting FeatureMembership...")
	deleteOp := deleteFeatureMembershipOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllFeatureMembership deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllFeatureMembership(ctx context.Context, project, location, feature string, filter func(*FeatureMembership) bool) error {
	listObj, err := c.ListFeatureMembership(ctx, project, location, feature)
	if err != nil {
		return err
	}

	err = c.deleteAllFeatureMembership(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllFeatureMembership(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyFeatureMembership(ctx context.Context, rawDesired *FeatureMembership, opts ...dcl.ApplyOption) (*FeatureMembership, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *FeatureMembership
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyFeatureMembershipHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyFeatureMembershipHelper(c *Client, ctx context.Context, rawDesired *FeatureMembership, opts ...dcl.ApplyOption) (*FeatureMembership, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyFeatureMembership...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractFeatureMembershipFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.featureMembershipDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToFeatureMembershipDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []featureMembershipApiOperation
	if create {
		ops = append(ops, &createFeatureMembershipOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyFeatureMembershipDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyFeatureMembershipDiff(c *Client, ctx context.Context, desired *FeatureMembership, rawDesired *FeatureMembership, ops []featureMembershipApiOperation, opts ...dcl.ApplyOption) (*FeatureMembership, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetFeatureMembership(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createFeatureMembershipOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapFeatureMembership(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeFeatureMembershipNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeFeatureMembershipNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeFeatureMembershipDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractFeatureMembershipFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractFeatureMembershipFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffFeatureMembership(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
