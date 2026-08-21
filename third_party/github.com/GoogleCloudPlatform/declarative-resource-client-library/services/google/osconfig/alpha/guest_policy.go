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
package alpha

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type GuestPolicy struct {
	Name                *string                          `json:"name"`
	Description         *string                          `json:"description"`
	CreateTime          *string                          `json:"createTime"`
	UpdateTime          *string                          `json:"updateTime"`
	Assignment          *GuestPolicyAssignment           `json:"assignment"`
	Packages            []GuestPolicyPackages            `json:"packages"`
	PackageRepositories []GuestPolicyPackageRepositories `json:"packageRepositories"`
	Recipes             []GuestPolicyRecipes             `json:"recipes"`
	Etag                *string                          `json:"etag"`
	Project             *string                          `json:"project"`
}

func (r *GuestPolicy) String() string {
	return dcl.SprintResource(r)
}

// The enum GuestPolicyPackagesDesiredStateEnum.
type GuestPolicyPackagesDesiredStateEnum string

// GuestPolicyPackagesDesiredStateEnumRef returns a *GuestPolicyPackagesDesiredStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyPackagesDesiredStateEnumRef(s string) *GuestPolicyPackagesDesiredStateEnum {
	v := GuestPolicyPackagesDesiredStateEnum(s)
	return &v
}

func (v GuestPolicyPackagesDesiredStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"DESIRED_STATE_UNSPECIFIED", "INSTALLED", "REMOVED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyPackagesDesiredStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyPackagesManagerEnum.
type GuestPolicyPackagesManagerEnum string

// GuestPolicyPackagesManagerEnumRef returns a *GuestPolicyPackagesManagerEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyPackagesManagerEnumRef(s string) *GuestPolicyPackagesManagerEnum {
	v := GuestPolicyPackagesManagerEnum(s)
	return &v
}

func (v GuestPolicyPackagesManagerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MANAGER_UNSPECIFIED", "ANY", "APT", "YUM", "ZYPPER", "GOO"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyPackagesManagerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyPackageRepositoriesAptArchiveTypeEnum.
type GuestPolicyPackageRepositoriesAptArchiveTypeEnum string

// GuestPolicyPackageRepositoriesAptArchiveTypeEnumRef returns a *GuestPolicyPackageRepositoriesAptArchiveTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyPackageRepositoriesAptArchiveTypeEnumRef(s string) *GuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	v := GuestPolicyPackageRepositoriesAptArchiveTypeEnum(s)
	return &v
}

func (v GuestPolicyPackageRepositoriesAptArchiveTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ARCHIVE_TYPE_UNSPECIFIED", "DEB", "DEB_SRC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyPackageRepositoriesAptArchiveTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum.
type GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum string

// GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumRef returns a *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumRef(s string) *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	v := GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(s)
	return &v
}

func (v GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"TYPE_UNSPECIFIED", "VALIDATION", "DESIRED_STATE_CHECK", "DESIRED_STATE_ENFORCEMENT", "DESIRED_STATE_CHECK_POST_ENFORCEMENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum.
type GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum string

// GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumRef returns a *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumRef(s string) *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	v := GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(s)
	return &v
}

func (v GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INTERPRETER_UNSPECIFIED", "NONE", "SHELL", "POWERSHELL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum.
type GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum string

// GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumRef returns a *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumRef(s string) *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	v := GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(s)
	return &v
}

func (v GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"TYPE_UNSPECIFIED", "VALIDATION", "DESIRED_STATE_CHECK", "DESIRED_STATE_ENFORCEMENT", "DESIRED_STATE_CHECK_POST_ENFORCEMENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum.
type GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum string

// GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumRef returns a *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumRef(s string) *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	v := GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(s)
	return &v
}

func (v GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INTERPRETER_UNSPECIFIED", "NONE", "SHELL", "POWERSHELL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GuestPolicyRecipesDesiredStateEnum.
type GuestPolicyRecipesDesiredStateEnum string

// GuestPolicyRecipesDesiredStateEnumRef returns a *GuestPolicyRecipesDesiredStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func GuestPolicyRecipesDesiredStateEnumRef(s string) *GuestPolicyRecipesDesiredStateEnum {
	v := GuestPolicyRecipesDesiredStateEnum(s)
	return &v
}

func (v GuestPolicyRecipesDesiredStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"DESIRED_STATE_UNSPECIFIED", "INSTALLED", "REMOVED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GuestPolicyRecipesDesiredStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type GuestPolicyAssignment struct {
	empty                bool                               `json:"-"`
	GroupLabels          []GuestPolicyAssignmentGroupLabels `json:"groupLabels"`
	Zones                []string                           `json:"zones"`
	Instances            []string                           `json:"instances"`
	InstanceNamePrefixes []string                           `json:"instanceNamePrefixes"`
	OSTypes              []GuestPolicyAssignmentOSTypes     `json:"osTypes"`
}

type jsonGuestPolicyAssignment GuestPolicyAssignment

func (r *GuestPolicyAssignment) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyAssignment
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyAssignment
	} else {

		r.GroupLabels = res.GroupLabels

		r.Zones = res.Zones

		r.Instances = res.Instances

		r.InstanceNamePrefixes = res.InstanceNamePrefixes

		r.OSTypes = res.OSTypes

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyAssignment is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyAssignment *GuestPolicyAssignment = &GuestPolicyAssignment{empty: true}

func (r *GuestPolicyAssignment) Empty() bool {
	return r.empty
}

func (r *GuestPolicyAssignment) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyAssignment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyAssignmentGroupLabels struct {
	empty  bool              `json:"-"`
	Labels map[string]string `json:"labels"`
}

type jsonGuestPolicyAssignmentGroupLabels GuestPolicyAssignmentGroupLabels

func (r *GuestPolicyAssignmentGroupLabels) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyAssignmentGroupLabels
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyAssignmentGroupLabels
	} else {

		r.Labels = res.Labels

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyAssignmentGroupLabels is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyAssignmentGroupLabels *GuestPolicyAssignmentGroupLabels = &GuestPolicyAssignmentGroupLabels{empty: true}

func (r *GuestPolicyAssignmentGroupLabels) Empty() bool {
	return r.empty
}

func (r *GuestPolicyAssignmentGroupLabels) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyAssignmentGroupLabels) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyAssignmentOSTypes struct {
	empty          bool    `json:"-"`
	OSShortName    *string `json:"osShortName"`
	OSVersion      *string `json:"osVersion"`
	OSArchitecture *string `json:"osArchitecture"`
}

type jsonGuestPolicyAssignmentOSTypes GuestPolicyAssignmentOSTypes

func (r *GuestPolicyAssignmentOSTypes) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyAssignmentOSTypes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyAssignmentOSTypes
	} else {

		r.OSShortName = res.OSShortName

		r.OSVersion = res.OSVersion

		r.OSArchitecture = res.OSArchitecture

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyAssignmentOSTypes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyAssignmentOSTypes *GuestPolicyAssignmentOSTypes = &GuestPolicyAssignmentOSTypes{empty: true}

func (r *GuestPolicyAssignmentOSTypes) Empty() bool {
	return r.empty
}

func (r *GuestPolicyAssignmentOSTypes) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyAssignmentOSTypes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackages struct {
	empty        bool                                 `json:"-"`
	Name         *string                              `json:"name"`
	DesiredState *GuestPolicyPackagesDesiredStateEnum `json:"desiredState"`
	Manager      *GuestPolicyPackagesManagerEnum      `json:"manager"`
}

type jsonGuestPolicyPackages GuestPolicyPackages

func (r *GuestPolicyPackages) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyPackages
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyPackages
	} else {

		r.Name = res.Name

		r.DesiredState = res.DesiredState

		r.Manager = res.Manager

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyPackages is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyPackages *GuestPolicyPackages = &GuestPolicyPackages{empty: true}

func (r *GuestPolicyPackages) Empty() bool {
	return r.empty
}

func (r *GuestPolicyPackages) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackages) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositories struct {
	empty  bool                                  `json:"-"`
	Apt    *GuestPolicyPackageRepositoriesApt    `json:"apt"`
	Yum    *GuestPolicyPackageRepositoriesYum    `json:"yum"`
	Zypper *GuestPolicyPackageRepositoriesZypper `json:"zypper"`
	Goo    *GuestPolicyPackageRepositoriesGoo    `json:"goo"`
}

type jsonGuestPolicyPackageRepositories GuestPolicyPackageRepositories

func (r *GuestPolicyPackageRepositories) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyPackageRepositories
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyPackageRepositories
	} else {

		r.Apt = res.Apt

		r.Yum = res.Yum

		r.Zypper = res.Zypper

		r.Goo = res.Goo

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositories is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositories *GuestPolicyPackageRepositories = &GuestPolicyPackageRepositories{empty: true}

func (r *GuestPolicyPackageRepositories) Empty() bool {
	return r.empty
}

func (r *GuestPolicyPackageRepositories) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositories) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositoriesApt struct {
	empty        bool                                              `json:"-"`
	ArchiveType  *GuestPolicyPackageRepositoriesAptArchiveTypeEnum `json:"archiveType"`
	Uri          *string                                           `json:"uri"`
	Distribution *string                                           `json:"distribution"`
	Components   []string                                          `json:"components"`
	GpgKey       *string                                           `json:"gpgKey"`
}

type jsonGuestPolicyPackageRepositoriesApt GuestPolicyPackageRepositoriesApt

func (r *GuestPolicyPackageRepositoriesApt) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyPackageRepositoriesApt
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyPackageRepositoriesApt
	} else {

		r.ArchiveType = res.ArchiveType

		r.Uri = res.Uri

		r.Distribution = res.Distribution

		r.Components = res.Components

		r.GpgKey = res.GpgKey

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositoriesApt is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositoriesApt *GuestPolicyPackageRepositoriesApt = &GuestPolicyPackageRepositoriesApt{empty: true}

func (r *GuestPolicyPackageRepositoriesApt) Empty() bool {
	return r.empty
}

func (r *GuestPolicyPackageRepositoriesApt) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositoriesApt) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositoriesYum struct {
	empty       bool     `json:"-"`
	Id          *string  `json:"id"`
	DisplayName *string  `json:"displayName"`
	BaseUrl     *string  `json:"baseUrl"`
	GpgKeys     []string `json:"gpgKeys"`
}

type jsonGuestPolicyPackageRepositoriesYum GuestPolicyPackageRepositoriesYum

func (r *GuestPolicyPackageRepositoriesYum) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyPackageRepositoriesYum
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyPackageRepositoriesYum
	} else {

		r.Id = res.Id

		r.DisplayName = res.DisplayName

		r.BaseUrl = res.BaseUrl

		r.GpgKeys = res.GpgKeys

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositoriesYum is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositoriesYum *GuestPolicyPackageRepositoriesYum = &GuestPolicyPackageRepositoriesYum{empty: true}

func (r *GuestPolicyPackageRepositoriesYum) Empty() bool {
	return r.empty
}

func (r *GuestPolicyPackageRepositoriesYum) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositoriesYum) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositoriesZypper struct {
	empty       bool     `json:"-"`
	Id          *string  `json:"id"`
	DisplayName *string  `json:"displayName"`
	BaseUrl     *string  `json:"baseUrl"`
	GpgKeys     []string `json:"gpgKeys"`
}

type jsonGuestPolicyPackageRepositoriesZypper GuestPolicyPackageRepositoriesZypper

func (r *GuestPolicyPackageRepositoriesZypper) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyPackageRepositoriesZypper
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyPackageRepositoriesZypper
	} else {

		r.Id = res.Id

		r.DisplayName = res.DisplayName

		r.BaseUrl = res.BaseUrl

		r.GpgKeys = res.GpgKeys

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositoriesZypper is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositoriesZypper *GuestPolicyPackageRepositoriesZypper = &GuestPolicyPackageRepositoriesZypper{empty: true}

func (r *GuestPolicyPackageRepositoriesZypper) Empty() bool {
	return r.empty
}

func (r *GuestPolicyPackageRepositoriesZypper) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositoriesZypper) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyPackageRepositoriesGoo struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Url   *string `json:"url"`
}

type jsonGuestPolicyPackageRepositoriesGoo GuestPolicyPackageRepositoriesGoo

func (r *GuestPolicyPackageRepositoriesGoo) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyPackageRepositoriesGoo
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyPackageRepositoriesGoo
	} else {

		r.Name = res.Name

		r.Url = res.Url

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyPackageRepositoriesGoo is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyPackageRepositoriesGoo *GuestPolicyPackageRepositoriesGoo = &GuestPolicyPackageRepositoriesGoo{empty: true}

func (r *GuestPolicyPackageRepositoriesGoo) Empty() bool {
	return r.empty
}

func (r *GuestPolicyPackageRepositoriesGoo) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyPackageRepositoriesGoo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipes struct {
	empty        bool                                `json:"-"`
	Name         *string                             `json:"name"`
	Version      *string                             `json:"version"`
	Artifacts    []GuestPolicyRecipesArtifacts       `json:"artifacts"`
	InstallSteps []GuestPolicyRecipesInstallSteps    `json:"installSteps"`
	UpdateSteps  []GuestPolicyRecipesUpdateSteps     `json:"updateSteps"`
	DesiredState *GuestPolicyRecipesDesiredStateEnum `json:"desiredState"`
}

type jsonGuestPolicyRecipes GuestPolicyRecipes

func (r *GuestPolicyRecipes) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipes
	} else {

		r.Name = res.Name

		r.Version = res.Version

		r.Artifacts = res.Artifacts

		r.InstallSteps = res.InstallSteps

		r.UpdateSteps = res.UpdateSteps

		r.DesiredState = res.DesiredState

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipes *GuestPolicyRecipes = &GuestPolicyRecipes{empty: true}

func (r *GuestPolicyRecipes) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipes) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesArtifacts struct {
	empty         bool                               `json:"-"`
	Id            *string                            `json:"id"`
	Remote        *GuestPolicyRecipesArtifactsRemote `json:"remote"`
	Gcs           *GuestPolicyRecipesArtifactsGcs    `json:"gcs"`
	AllowInsecure *bool                              `json:"allowInsecure"`
}

type jsonGuestPolicyRecipesArtifacts GuestPolicyRecipesArtifacts

func (r *GuestPolicyRecipesArtifacts) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesArtifacts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesArtifacts
	} else {

		r.Id = res.Id

		r.Remote = res.Remote

		r.Gcs = res.Gcs

		r.AllowInsecure = res.AllowInsecure

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesArtifacts is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesArtifacts *GuestPolicyRecipesArtifacts = &GuestPolicyRecipesArtifacts{empty: true}

func (r *GuestPolicyRecipesArtifacts) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesArtifacts) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesArtifacts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesArtifactsRemote struct {
	empty    bool    `json:"-"`
	Uri      *string `json:"uri"`
	Checksum *string `json:"checksum"`
}

type jsonGuestPolicyRecipesArtifactsRemote GuestPolicyRecipesArtifactsRemote

func (r *GuestPolicyRecipesArtifactsRemote) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesArtifactsRemote
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesArtifactsRemote
	} else {

		r.Uri = res.Uri

		r.Checksum = res.Checksum

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesArtifactsRemote is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesArtifactsRemote *GuestPolicyRecipesArtifactsRemote = &GuestPolicyRecipesArtifactsRemote{empty: true}

func (r *GuestPolicyRecipesArtifactsRemote) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesArtifactsRemote) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesArtifactsRemote) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesArtifactsGcs struct {
	empty      bool    `json:"-"`
	Bucket     *string `json:"bucket"`
	Object     *string `json:"object"`
	Generation *int64  `json:"generation"`
}

type jsonGuestPolicyRecipesArtifactsGcs GuestPolicyRecipesArtifactsGcs

func (r *GuestPolicyRecipesArtifactsGcs) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesArtifactsGcs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesArtifactsGcs
	} else {

		r.Bucket = res.Bucket

		r.Object = res.Object

		r.Generation = res.Generation

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesArtifactsGcs is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesArtifactsGcs *GuestPolicyRecipesArtifactsGcs = &GuestPolicyRecipesArtifactsGcs{empty: true}

func (r *GuestPolicyRecipesArtifactsGcs) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesArtifactsGcs) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesArtifactsGcs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallSteps struct {
	empty             bool                                             `json:"-"`
	FileCopy          *GuestPolicyRecipesInstallStepsFileCopy          `json:"fileCopy"`
	ArchiveExtraction *GuestPolicyRecipesInstallStepsArchiveExtraction `json:"archiveExtraction"`
	MsiInstallation   *GuestPolicyRecipesInstallStepsMsiInstallation   `json:"msiInstallation"`
	DpkgInstallation  *GuestPolicyRecipesInstallStepsDpkgInstallation  `json:"dpkgInstallation"`
	RpmInstallation   *GuestPolicyRecipesInstallStepsRpmInstallation   `json:"rpmInstallation"`
	FileExec          *GuestPolicyRecipesInstallStepsFileExec          `json:"fileExec"`
	ScriptRun         *GuestPolicyRecipesInstallStepsScriptRun         `json:"scriptRun"`
}

type jsonGuestPolicyRecipesInstallSteps GuestPolicyRecipesInstallSteps

func (r *GuestPolicyRecipesInstallSteps) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesInstallSteps
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesInstallSteps
	} else {

		r.FileCopy = res.FileCopy

		r.ArchiveExtraction = res.ArchiveExtraction

		r.MsiInstallation = res.MsiInstallation

		r.DpkgInstallation = res.DpkgInstallation

		r.RpmInstallation = res.RpmInstallation

		r.FileExec = res.FileExec

		r.ScriptRun = res.ScriptRun

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallSteps is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallSteps *GuestPolicyRecipesInstallSteps = &GuestPolicyRecipesInstallSteps{empty: true}

func (r *GuestPolicyRecipesInstallSteps) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesInstallSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsFileCopy struct {
	empty       bool    `json:"-"`
	ArtifactId  *string `json:"artifactId"`
	Destination *string `json:"destination"`
	Overwrite   *bool   `json:"overwrite"`
	Permissions *string `json:"permissions"`
}

type jsonGuestPolicyRecipesInstallStepsFileCopy GuestPolicyRecipesInstallStepsFileCopy

func (r *GuestPolicyRecipesInstallStepsFileCopy) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesInstallStepsFileCopy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesInstallStepsFileCopy
	} else {

		r.ArtifactId = res.ArtifactId

		r.Destination = res.Destination

		r.Overwrite = res.Overwrite

		r.Permissions = res.Permissions

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsFileCopy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsFileCopy *GuestPolicyRecipesInstallStepsFileCopy = &GuestPolicyRecipesInstallStepsFileCopy{empty: true}

func (r *GuestPolicyRecipesInstallStepsFileCopy) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesInstallStepsFileCopy) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsFileCopy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsArchiveExtraction struct {
	empty       bool                                                     `json:"-"`
	ArtifactId  *string                                                  `json:"artifactId"`
	Destination *string                                                  `json:"destination"`
	Type        *GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum `json:"type"`
}

type jsonGuestPolicyRecipesInstallStepsArchiveExtraction GuestPolicyRecipesInstallStepsArchiveExtraction

func (r *GuestPolicyRecipesInstallStepsArchiveExtraction) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesInstallStepsArchiveExtraction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesInstallStepsArchiveExtraction
	} else {

		r.ArtifactId = res.ArtifactId

		r.Destination = res.Destination

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsArchiveExtraction is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsArchiveExtraction *GuestPolicyRecipesInstallStepsArchiveExtraction = &GuestPolicyRecipesInstallStepsArchiveExtraction{empty: true}

func (r *GuestPolicyRecipesInstallStepsArchiveExtraction) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesInstallStepsArchiveExtraction) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsArchiveExtraction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsMsiInstallation struct {
	empty            bool     `json:"-"`
	ArtifactId       *string  `json:"artifactId"`
	Flags            []string `json:"flags"`
	AllowedExitCodes []int64  `json:"allowedExitCodes"`
}

type jsonGuestPolicyRecipesInstallStepsMsiInstallation GuestPolicyRecipesInstallStepsMsiInstallation

func (r *GuestPolicyRecipesInstallStepsMsiInstallation) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesInstallStepsMsiInstallation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesInstallStepsMsiInstallation
	} else {

		r.ArtifactId = res.ArtifactId

		r.Flags = res.Flags

		r.AllowedExitCodes = res.AllowedExitCodes

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsMsiInstallation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsMsiInstallation *GuestPolicyRecipesInstallStepsMsiInstallation = &GuestPolicyRecipesInstallStepsMsiInstallation{empty: true}

func (r *GuestPolicyRecipesInstallStepsMsiInstallation) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesInstallStepsMsiInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsMsiInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsDpkgInstallation struct {
	empty      bool    `json:"-"`
	ArtifactId *string `json:"artifactId"`
}

type jsonGuestPolicyRecipesInstallStepsDpkgInstallation GuestPolicyRecipesInstallStepsDpkgInstallation

func (r *GuestPolicyRecipesInstallStepsDpkgInstallation) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesInstallStepsDpkgInstallation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesInstallStepsDpkgInstallation
	} else {

		r.ArtifactId = res.ArtifactId

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsDpkgInstallation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsDpkgInstallation *GuestPolicyRecipesInstallStepsDpkgInstallation = &GuestPolicyRecipesInstallStepsDpkgInstallation{empty: true}

func (r *GuestPolicyRecipesInstallStepsDpkgInstallation) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesInstallStepsDpkgInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsDpkgInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsRpmInstallation struct {
	empty      bool    `json:"-"`
	ArtifactId *string `json:"artifactId"`
}

type jsonGuestPolicyRecipesInstallStepsRpmInstallation GuestPolicyRecipesInstallStepsRpmInstallation

func (r *GuestPolicyRecipesInstallStepsRpmInstallation) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesInstallStepsRpmInstallation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesInstallStepsRpmInstallation
	} else {

		r.ArtifactId = res.ArtifactId

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsRpmInstallation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsRpmInstallation *GuestPolicyRecipesInstallStepsRpmInstallation = &GuestPolicyRecipesInstallStepsRpmInstallation{empty: true}

func (r *GuestPolicyRecipesInstallStepsRpmInstallation) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesInstallStepsRpmInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsRpmInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsFileExec struct {
	empty            bool     `json:"-"`
	ArtifactId       *string  `json:"artifactId"`
	LocalPath        *string  `json:"localPath"`
	Args             []string `json:"args"`
	AllowedExitCodes []int64  `json:"allowedExitCodes"`
}

type jsonGuestPolicyRecipesInstallStepsFileExec GuestPolicyRecipesInstallStepsFileExec

func (r *GuestPolicyRecipesInstallStepsFileExec) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesInstallStepsFileExec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesInstallStepsFileExec
	} else {

		r.ArtifactId = res.ArtifactId

		r.LocalPath = res.LocalPath

		r.Args = res.Args

		r.AllowedExitCodes = res.AllowedExitCodes

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsFileExec is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsFileExec *GuestPolicyRecipesInstallStepsFileExec = &GuestPolicyRecipesInstallStepsFileExec{empty: true}

func (r *GuestPolicyRecipesInstallStepsFileExec) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesInstallStepsFileExec) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsFileExec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesInstallStepsScriptRun struct {
	empty            bool                                                    `json:"-"`
	Script           *string                                                 `json:"script"`
	AllowedExitCodes []int64                                                 `json:"allowedExitCodes"`
	Interpreter      *GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum `json:"interpreter"`
}

type jsonGuestPolicyRecipesInstallStepsScriptRun GuestPolicyRecipesInstallStepsScriptRun

func (r *GuestPolicyRecipesInstallStepsScriptRun) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesInstallStepsScriptRun
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesInstallStepsScriptRun
	} else {

		r.Script = res.Script

		r.AllowedExitCodes = res.AllowedExitCodes

		r.Interpreter = res.Interpreter

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesInstallStepsScriptRun is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesInstallStepsScriptRun *GuestPolicyRecipesInstallStepsScriptRun = &GuestPolicyRecipesInstallStepsScriptRun{empty: true}

func (r *GuestPolicyRecipesInstallStepsScriptRun) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesInstallStepsScriptRun) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesInstallStepsScriptRun) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateSteps struct {
	empty             bool                                            `json:"-"`
	FileCopy          *GuestPolicyRecipesUpdateStepsFileCopy          `json:"fileCopy"`
	ArchiveExtraction *GuestPolicyRecipesUpdateStepsArchiveExtraction `json:"archiveExtraction"`
	MsiInstallation   *GuestPolicyRecipesUpdateStepsMsiInstallation   `json:"msiInstallation"`
	DpkgInstallation  *GuestPolicyRecipesUpdateStepsDpkgInstallation  `json:"dpkgInstallation"`
	RpmInstallation   *GuestPolicyRecipesUpdateStepsRpmInstallation   `json:"rpmInstallation"`
	FileExec          *GuestPolicyRecipesUpdateStepsFileExec          `json:"fileExec"`
	ScriptRun         *GuestPolicyRecipesUpdateStepsScriptRun         `json:"scriptRun"`
}

type jsonGuestPolicyRecipesUpdateSteps GuestPolicyRecipesUpdateSteps

func (r *GuestPolicyRecipesUpdateSteps) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesUpdateSteps
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesUpdateSteps
	} else {

		r.FileCopy = res.FileCopy

		r.ArchiveExtraction = res.ArchiveExtraction

		r.MsiInstallation = res.MsiInstallation

		r.DpkgInstallation = res.DpkgInstallation

		r.RpmInstallation = res.RpmInstallation

		r.FileExec = res.FileExec

		r.ScriptRun = res.ScriptRun

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateSteps is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateSteps *GuestPolicyRecipesUpdateSteps = &GuestPolicyRecipesUpdateSteps{empty: true}

func (r *GuestPolicyRecipesUpdateSteps) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesUpdateSteps) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateSteps) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsFileCopy struct {
	empty       bool    `json:"-"`
	ArtifactId  *string `json:"artifactId"`
	Destination *string `json:"destination"`
	Overwrite   *bool   `json:"overwrite"`
	Permissions *string `json:"permissions"`
}

type jsonGuestPolicyRecipesUpdateStepsFileCopy GuestPolicyRecipesUpdateStepsFileCopy

func (r *GuestPolicyRecipesUpdateStepsFileCopy) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesUpdateStepsFileCopy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesUpdateStepsFileCopy
	} else {

		r.ArtifactId = res.ArtifactId

		r.Destination = res.Destination

		r.Overwrite = res.Overwrite

		r.Permissions = res.Permissions

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsFileCopy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsFileCopy *GuestPolicyRecipesUpdateStepsFileCopy = &GuestPolicyRecipesUpdateStepsFileCopy{empty: true}

func (r *GuestPolicyRecipesUpdateStepsFileCopy) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesUpdateStepsFileCopy) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsFileCopy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsArchiveExtraction struct {
	empty       bool                                                    `json:"-"`
	ArtifactId  *string                                                 `json:"artifactId"`
	Destination *string                                                 `json:"destination"`
	Type        *GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum `json:"type"`
}

type jsonGuestPolicyRecipesUpdateStepsArchiveExtraction GuestPolicyRecipesUpdateStepsArchiveExtraction

func (r *GuestPolicyRecipesUpdateStepsArchiveExtraction) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesUpdateStepsArchiveExtraction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesUpdateStepsArchiveExtraction
	} else {

		r.ArtifactId = res.ArtifactId

		r.Destination = res.Destination

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsArchiveExtraction is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsArchiveExtraction *GuestPolicyRecipesUpdateStepsArchiveExtraction = &GuestPolicyRecipesUpdateStepsArchiveExtraction{empty: true}

func (r *GuestPolicyRecipesUpdateStepsArchiveExtraction) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesUpdateStepsArchiveExtraction) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsArchiveExtraction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsMsiInstallation struct {
	empty            bool     `json:"-"`
	ArtifactId       *string  `json:"artifactId"`
	Flags            []string `json:"flags"`
	AllowedExitCodes []int64  `json:"allowedExitCodes"`
}

type jsonGuestPolicyRecipesUpdateStepsMsiInstallation GuestPolicyRecipesUpdateStepsMsiInstallation

func (r *GuestPolicyRecipesUpdateStepsMsiInstallation) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesUpdateStepsMsiInstallation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesUpdateStepsMsiInstallation
	} else {

		r.ArtifactId = res.ArtifactId

		r.Flags = res.Flags

		r.AllowedExitCodes = res.AllowedExitCodes

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsMsiInstallation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsMsiInstallation *GuestPolicyRecipesUpdateStepsMsiInstallation = &GuestPolicyRecipesUpdateStepsMsiInstallation{empty: true}

func (r *GuestPolicyRecipesUpdateStepsMsiInstallation) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesUpdateStepsMsiInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsMsiInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsDpkgInstallation struct {
	empty      bool    `json:"-"`
	ArtifactId *string `json:"artifactId"`
}

type jsonGuestPolicyRecipesUpdateStepsDpkgInstallation GuestPolicyRecipesUpdateStepsDpkgInstallation

func (r *GuestPolicyRecipesUpdateStepsDpkgInstallation) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesUpdateStepsDpkgInstallation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesUpdateStepsDpkgInstallation
	} else {

		r.ArtifactId = res.ArtifactId

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsDpkgInstallation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsDpkgInstallation *GuestPolicyRecipesUpdateStepsDpkgInstallation = &GuestPolicyRecipesUpdateStepsDpkgInstallation{empty: true}

func (r *GuestPolicyRecipesUpdateStepsDpkgInstallation) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesUpdateStepsDpkgInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsDpkgInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsRpmInstallation struct {
	empty      bool    `json:"-"`
	ArtifactId *string `json:"artifactId"`
}

type jsonGuestPolicyRecipesUpdateStepsRpmInstallation GuestPolicyRecipesUpdateStepsRpmInstallation

func (r *GuestPolicyRecipesUpdateStepsRpmInstallation) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesUpdateStepsRpmInstallation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesUpdateStepsRpmInstallation
	} else {

		r.ArtifactId = res.ArtifactId

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsRpmInstallation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsRpmInstallation *GuestPolicyRecipesUpdateStepsRpmInstallation = &GuestPolicyRecipesUpdateStepsRpmInstallation{empty: true}

func (r *GuestPolicyRecipesUpdateStepsRpmInstallation) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesUpdateStepsRpmInstallation) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsRpmInstallation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsFileExec struct {
	empty            bool     `json:"-"`
	ArtifactId       *string  `json:"artifactId"`
	LocalPath        *string  `json:"localPath"`
	Args             []string `json:"args"`
	AllowedExitCodes []int64  `json:"allowedExitCodes"`
}

type jsonGuestPolicyRecipesUpdateStepsFileExec GuestPolicyRecipesUpdateStepsFileExec

func (r *GuestPolicyRecipesUpdateStepsFileExec) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesUpdateStepsFileExec
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesUpdateStepsFileExec
	} else {

		r.ArtifactId = res.ArtifactId

		r.LocalPath = res.LocalPath

		r.Args = res.Args

		r.AllowedExitCodes = res.AllowedExitCodes

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsFileExec is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsFileExec *GuestPolicyRecipesUpdateStepsFileExec = &GuestPolicyRecipesUpdateStepsFileExec{empty: true}

func (r *GuestPolicyRecipesUpdateStepsFileExec) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesUpdateStepsFileExec) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsFileExec) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GuestPolicyRecipesUpdateStepsScriptRun struct {
	empty            bool                                                   `json:"-"`
	Script           *string                                                `json:"script"`
	AllowedExitCodes []int64                                                `json:"allowedExitCodes"`
	Interpreter      *GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum `json:"interpreter"`
}

type jsonGuestPolicyRecipesUpdateStepsScriptRun GuestPolicyRecipesUpdateStepsScriptRun

func (r *GuestPolicyRecipesUpdateStepsScriptRun) UnmarshalJSON(data []byte) error {
	var res jsonGuestPolicyRecipesUpdateStepsScriptRun
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGuestPolicyRecipesUpdateStepsScriptRun
	} else {

		r.Script = res.Script

		r.AllowedExitCodes = res.AllowedExitCodes

		r.Interpreter = res.Interpreter

	}
	return nil
}

// This object is used to assert a desired state where this GuestPolicyRecipesUpdateStepsScriptRun is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGuestPolicyRecipesUpdateStepsScriptRun *GuestPolicyRecipesUpdateStepsScriptRun = &GuestPolicyRecipesUpdateStepsScriptRun{empty: true}

func (r *GuestPolicyRecipesUpdateStepsScriptRun) Empty() bool {
	return r.empty
}

func (r *GuestPolicyRecipesUpdateStepsScriptRun) String() string {
	return dcl.SprintResource(r)
}

func (r *GuestPolicyRecipesUpdateStepsScriptRun) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *GuestPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "os_config",
		Type:    "GuestPolicy",
		Version: "alpha",
	}
}

func (r *GuestPolicy) ID() (string, error) {
	if err := extractGuestPolicyFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                 dcl.ValueOrEmptyString(nr.Name),
		"description":          dcl.ValueOrEmptyString(nr.Description),
		"create_time":          dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":          dcl.ValueOrEmptyString(nr.UpdateTime),
		"assignment":           dcl.ValueOrEmptyString(nr.Assignment),
		"packages":             dcl.ValueOrEmptyString(nr.Packages),
		"package_repositories": dcl.ValueOrEmptyString(nr.PackageRepositories),
		"recipes":              dcl.ValueOrEmptyString(nr.Recipes),
		"etag":                 dcl.ValueOrEmptyString(nr.Etag),
		"project":              dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/guestPolicies/{{name}}", params), nil
}

const GuestPolicyMaxPage = -1

type GuestPolicyList struct {
	Items []*GuestPolicy

	nextToken string

	pageSize int32

	resource *GuestPolicy
}

func (l *GuestPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *GuestPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listGuestPolicy(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListGuestPolicy(ctx context.Context, project string) (*GuestPolicyList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListGuestPolicyWithMaxResults(ctx, project, GuestPolicyMaxPage)

}

func (c *Client) ListGuestPolicyWithMaxResults(ctx context.Context, project string, pageSize int32) (*GuestPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &GuestPolicy{
		Project: &project,
	}
	items, token, err := c.listGuestPolicy(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &GuestPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetGuestPolicy(ctx context.Context, r *GuestPolicy) (*GuestPolicy, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractGuestPolicyFields(r)

	b, err := c.getGuestPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalGuestPolicy(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeGuestPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractGuestPolicyFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteGuestPolicy(ctx context.Context, r *GuestPolicy) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("GuestPolicy resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting GuestPolicy...")
	deleteOp := deleteGuestPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllGuestPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllGuestPolicy(ctx context.Context, project string, filter func(*GuestPolicy) bool) error {
	listObj, err := c.ListGuestPolicy(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllGuestPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllGuestPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyGuestPolicy(ctx context.Context, rawDesired *GuestPolicy, opts ...dcl.ApplyOption) (*GuestPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *GuestPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyGuestPolicyHelper(c, ctx, rawDesired, opts...)
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

func applyGuestPolicyHelper(c *Client, ctx context.Context, rawDesired *GuestPolicy, opts ...dcl.ApplyOption) (*GuestPolicy, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyGuestPolicy...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractGuestPolicyFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.guestPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToGuestPolicyDiffs(c.Config, fieldDiffs, opts)
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
	var ops []guestPolicyApiOperation
	if create {
		ops = append(ops, &createGuestPolicyOperation{})
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
	return applyGuestPolicyDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyGuestPolicyDiff(c *Client, ctx context.Context, desired *GuestPolicy, rawDesired *GuestPolicy, ops []guestPolicyApiOperation, opts ...dcl.ApplyOption) (*GuestPolicy, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetGuestPolicy(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createGuestPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapGuestPolicy(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeGuestPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeGuestPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeGuestPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractGuestPolicyFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractGuestPolicyFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffGuestPolicy(c, newDesired, newState)
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
