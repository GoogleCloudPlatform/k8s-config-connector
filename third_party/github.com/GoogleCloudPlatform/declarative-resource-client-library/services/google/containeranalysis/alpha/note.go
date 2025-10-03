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

type Note struct {
	Name             *string            `json:"name"`
	ShortDescription *string            `json:"shortDescription"`
	LongDescription  *string            `json:"longDescription"`
	RelatedUrl       []NoteRelatedUrl   `json:"relatedUrl"`
	ExpirationTime   *string            `json:"expirationTime"`
	CreateTime       *string            `json:"createTime"`
	UpdateTime       *string            `json:"updateTime"`
	RelatedNoteNames []string           `json:"relatedNoteNames"`
	Vulnerability    *NoteVulnerability `json:"vulnerability"`
	Build            *NoteBuild         `json:"build"`
	Image            *NoteImage         `json:"image"`
	Package          *NotePackage       `json:"package"`
	Discovery        *NoteDiscovery     `json:"discovery"`
	Deployment       *NoteDeployment    `json:"deployment"`
	Attestation      *NoteAttestation   `json:"attestation"`
	Project          *string            `json:"project"`
}

func (r *Note) String() string {
	return dcl.SprintResource(r)
}

// The enum NoteVulnerabilitySeverityEnum.
type NoteVulnerabilitySeverityEnum string

// NoteVulnerabilitySeverityEnumRef returns a *NoteVulnerabilitySeverityEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilitySeverityEnumRef(s string) *NoteVulnerabilitySeverityEnum {
	v := NoteVulnerabilitySeverityEnum(s)
	return &v
}

func (v NoteVulnerabilitySeverityEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SEVERITY_UNSPECIFIED", "MINIMAL", "LOW", "MEDIUM", "HIGH", "CRITICAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilitySeverityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityDetailsAffectedVersionStartKindEnum.
type NoteVulnerabilityDetailsAffectedVersionStartKindEnum string

// NoteVulnerabilityDetailsAffectedVersionStartKindEnumRef returns a *NoteVulnerabilityDetailsAffectedVersionStartKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityDetailsAffectedVersionStartKindEnumRef(s string) *NoteVulnerabilityDetailsAffectedVersionStartKindEnum {
	v := NoteVulnerabilityDetailsAffectedVersionStartKindEnum(s)
	return &v
}

func (v NoteVulnerabilityDetailsAffectedVersionStartKindEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityDetailsAffectedVersionStartKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityDetailsAffectedVersionEndKindEnum.
type NoteVulnerabilityDetailsAffectedVersionEndKindEnum string

// NoteVulnerabilityDetailsAffectedVersionEndKindEnumRef returns a *NoteVulnerabilityDetailsAffectedVersionEndKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityDetailsAffectedVersionEndKindEnumRef(s string) *NoteVulnerabilityDetailsAffectedVersionEndKindEnum {
	v := NoteVulnerabilityDetailsAffectedVersionEndKindEnum(s)
	return &v
}

func (v NoteVulnerabilityDetailsAffectedVersionEndKindEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityDetailsAffectedVersionEndKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityDetailsFixedVersionKindEnum.
type NoteVulnerabilityDetailsFixedVersionKindEnum string

// NoteVulnerabilityDetailsFixedVersionKindEnumRef returns a *NoteVulnerabilityDetailsFixedVersionKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityDetailsFixedVersionKindEnumRef(s string) *NoteVulnerabilityDetailsFixedVersionKindEnum {
	v := NoteVulnerabilityDetailsFixedVersionKindEnum(s)
	return &v
}

func (v NoteVulnerabilityDetailsFixedVersionKindEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityDetailsFixedVersionKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3AttackVectorEnum.
type NoteVulnerabilityCvssV3AttackVectorEnum string

// NoteVulnerabilityCvssV3AttackVectorEnumRef returns a *NoteVulnerabilityCvssV3AttackVectorEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3AttackVectorEnumRef(s string) *NoteVulnerabilityCvssV3AttackVectorEnum {
	v := NoteVulnerabilityCvssV3AttackVectorEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3AttackVectorEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ATTACK_VECTOR_UNSPECIFIED", "ATTACK_VECTOR_NETWORK", "ATTACK_VECTOR_ADJACENT", "ATTACK_VECTOR_LOCAL", "ATTACK_VECTOR_PHYSICAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3AttackVectorEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3AttackComplexityEnum.
type NoteVulnerabilityCvssV3AttackComplexityEnum string

// NoteVulnerabilityCvssV3AttackComplexityEnumRef returns a *NoteVulnerabilityCvssV3AttackComplexityEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3AttackComplexityEnumRef(s string) *NoteVulnerabilityCvssV3AttackComplexityEnum {
	v := NoteVulnerabilityCvssV3AttackComplexityEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3AttackComplexityEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ATTACK_COMPLEXITY_UNSPECIFIED", "ATTACK_COMPLEXITY_LOW", "ATTACK_COMPLEXITY_HIGH"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3AttackComplexityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3PrivilegesRequiredEnum.
type NoteVulnerabilityCvssV3PrivilegesRequiredEnum string

// NoteVulnerabilityCvssV3PrivilegesRequiredEnumRef returns a *NoteVulnerabilityCvssV3PrivilegesRequiredEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3PrivilegesRequiredEnumRef(s string) *NoteVulnerabilityCvssV3PrivilegesRequiredEnum {
	v := NoteVulnerabilityCvssV3PrivilegesRequiredEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3PrivilegesRequiredEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"PRIVILEGES_REQUIRED_UNSPECIFIED", "PRIVILEGES_REQUIRED_NONE", "PRIVILEGES_REQUIRED_LOW", "PRIVILEGES_REQUIRED_HIGH"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3PrivilegesRequiredEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3UserInteractionEnum.
type NoteVulnerabilityCvssV3UserInteractionEnum string

// NoteVulnerabilityCvssV3UserInteractionEnumRef returns a *NoteVulnerabilityCvssV3UserInteractionEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3UserInteractionEnumRef(s string) *NoteVulnerabilityCvssV3UserInteractionEnum {
	v := NoteVulnerabilityCvssV3UserInteractionEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3UserInteractionEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"USER_INTERACTION_UNSPECIFIED", "USER_INTERACTION_NONE", "USER_INTERACTION_REQUIRED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3UserInteractionEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3ScopeEnum.
type NoteVulnerabilityCvssV3ScopeEnum string

// NoteVulnerabilityCvssV3ScopeEnumRef returns a *NoteVulnerabilityCvssV3ScopeEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3ScopeEnumRef(s string) *NoteVulnerabilityCvssV3ScopeEnum {
	v := NoteVulnerabilityCvssV3ScopeEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3ScopeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SCOPE_UNSPECIFIED", "SCOPE_UNCHANGED", "SCOPE_CHANGED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3ScopeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3ConfidentialityImpactEnum.
type NoteVulnerabilityCvssV3ConfidentialityImpactEnum string

// NoteVulnerabilityCvssV3ConfidentialityImpactEnumRef returns a *NoteVulnerabilityCvssV3ConfidentialityImpactEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3ConfidentialityImpactEnumRef(s string) *NoteVulnerabilityCvssV3ConfidentialityImpactEnum {
	v := NoteVulnerabilityCvssV3ConfidentialityImpactEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3ConfidentialityImpactEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"IMPACT_UNSPECIFIED", "IMPACT_HIGH", "IMPACT_LOW", "IMPACT_NONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3ConfidentialityImpactEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3IntegrityImpactEnum.
type NoteVulnerabilityCvssV3IntegrityImpactEnum string

// NoteVulnerabilityCvssV3IntegrityImpactEnumRef returns a *NoteVulnerabilityCvssV3IntegrityImpactEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3IntegrityImpactEnumRef(s string) *NoteVulnerabilityCvssV3IntegrityImpactEnum {
	v := NoteVulnerabilityCvssV3IntegrityImpactEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3IntegrityImpactEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"IMPACT_UNSPECIFIED", "IMPACT_HIGH", "IMPACT_LOW", "IMPACT_NONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3IntegrityImpactEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteVulnerabilityCvssV3AvailabilityImpactEnum.
type NoteVulnerabilityCvssV3AvailabilityImpactEnum string

// NoteVulnerabilityCvssV3AvailabilityImpactEnumRef returns a *NoteVulnerabilityCvssV3AvailabilityImpactEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteVulnerabilityCvssV3AvailabilityImpactEnumRef(s string) *NoteVulnerabilityCvssV3AvailabilityImpactEnum {
	v := NoteVulnerabilityCvssV3AvailabilityImpactEnum(s)
	return &v
}

func (v NoteVulnerabilityCvssV3AvailabilityImpactEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"IMPACT_UNSPECIFIED", "IMPACT_HIGH", "IMPACT_LOW", "IMPACT_NONE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteVulnerabilityCvssV3AvailabilityImpactEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteBuildSignatureKeyTypeEnum.
type NoteBuildSignatureKeyTypeEnum string

// NoteBuildSignatureKeyTypeEnumRef returns a *NoteBuildSignatureKeyTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteBuildSignatureKeyTypeEnumRef(s string) *NoteBuildSignatureKeyTypeEnum {
	v := NoteBuildSignatureKeyTypeEnum(s)
	return &v
}

func (v NoteBuildSignatureKeyTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"KEY_TYPE_UNSPECIFIED", "PGP_ASCII_ARMORED", "PKIX_PEM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteBuildSignatureKeyTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NotePackageDistributionArchitectureEnum.
type NotePackageDistributionArchitectureEnum string

// NotePackageDistributionArchitectureEnumRef returns a *NotePackageDistributionArchitectureEnum with the value of string s
// If the empty string is provided, nil is returned.
func NotePackageDistributionArchitectureEnumRef(s string) *NotePackageDistributionArchitectureEnum {
	v := NotePackageDistributionArchitectureEnum(s)
	return &v
}

func (v NotePackageDistributionArchitectureEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ARCHITECTURE_UNSPECIFIED", "X86", "X64"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NotePackageDistributionArchitectureEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NotePackageDistributionLatestVersionKindEnum.
type NotePackageDistributionLatestVersionKindEnum string

// NotePackageDistributionLatestVersionKindEnumRef returns a *NotePackageDistributionLatestVersionKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NotePackageDistributionLatestVersionKindEnumRef(s string) *NotePackageDistributionLatestVersionKindEnum {
	v := NotePackageDistributionLatestVersionKindEnum(s)
	return &v
}

func (v NotePackageDistributionLatestVersionKindEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"VERSION_KIND_UNSPECIFIED", "NORMAL", "MINIMUM", "MAXIMUM"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NotePackageDistributionLatestVersionKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum NoteDiscoveryAnalysisKindEnum.
type NoteDiscoveryAnalysisKindEnum string

// NoteDiscoveryAnalysisKindEnumRef returns a *NoteDiscoveryAnalysisKindEnum with the value of string s
// If the empty string is provided, nil is returned.
func NoteDiscoveryAnalysisKindEnumRef(s string) *NoteDiscoveryAnalysisKindEnum {
	v := NoteDiscoveryAnalysisKindEnum(s)
	return &v
}

func (v NoteDiscoveryAnalysisKindEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"NOTE_KIND_UNSPECIFIED", "VULNERABILITY", "BUILD", "IMAGE", "PACKAGE", "DEPLOYMENT", "DISCOVERY", "ATTESTATION", "UPGRADE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "NoteDiscoveryAnalysisKindEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type NoteRelatedUrl struct {
	empty bool    `json:"-"`
	Url   *string `json:"url"`
	Label *string `json:"label"`
}

type jsonNoteRelatedUrl NoteRelatedUrl

func (r *NoteRelatedUrl) UnmarshalJSON(data []byte) error {
	var res jsonNoteRelatedUrl
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteRelatedUrl
	} else {

		r.Url = res.Url

		r.Label = res.Label

	}
	return nil
}

// This object is used to assert a desired state where this NoteRelatedUrl is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteRelatedUrl *NoteRelatedUrl = &NoteRelatedUrl{empty: true}

func (r *NoteRelatedUrl) Empty() bool {
	return r.empty
}

func (r *NoteRelatedUrl) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteRelatedUrl) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerability struct {
	empty            bool                              `json:"-"`
	CvssScore        *float64                          `json:"cvssScore"`
	Severity         *NoteVulnerabilitySeverityEnum    `json:"severity"`
	Details          []NoteVulnerabilityDetails        `json:"details"`
	CvssV3           *NoteVulnerabilityCvssV3          `json:"cvssV3"`
	WindowsDetails   []NoteVulnerabilityWindowsDetails `json:"windowsDetails"`
	SourceUpdateTime *string                           `json:"sourceUpdateTime"`
}

type jsonNoteVulnerability NoteVulnerability

func (r *NoteVulnerability) UnmarshalJSON(data []byte) error {
	var res jsonNoteVulnerability
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteVulnerability
	} else {

		r.CvssScore = res.CvssScore

		r.Severity = res.Severity

		r.Details = res.Details

		r.CvssV3 = res.CvssV3

		r.WindowsDetails = res.WindowsDetails

		r.SourceUpdateTime = res.SourceUpdateTime

	}
	return nil
}

// This object is used to assert a desired state where this NoteVulnerability is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteVulnerability *NoteVulnerability = &NoteVulnerability{empty: true}

func (r *NoteVulnerability) Empty() bool {
	return r.empty
}

func (r *NoteVulnerability) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerability) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityDetails struct {
	empty                bool                                          `json:"-"`
	SeverityName         *string                                       `json:"severityName"`
	Description          *string                                       `json:"description"`
	PackageType          *string                                       `json:"packageType"`
	AffectedCpeUri       *string                                       `json:"affectedCpeUri"`
	AffectedPackage      *string                                       `json:"affectedPackage"`
	AffectedVersionStart *NoteVulnerabilityDetailsAffectedVersionStart `json:"affectedVersionStart"`
	AffectedVersionEnd   *NoteVulnerabilityDetailsAffectedVersionEnd   `json:"affectedVersionEnd"`
	FixedCpeUri          *string                                       `json:"fixedCpeUri"`
	FixedPackage         *string                                       `json:"fixedPackage"`
	FixedVersion         *NoteVulnerabilityDetailsFixedVersion         `json:"fixedVersion"`
	IsObsolete           *bool                                         `json:"isObsolete"`
	SourceUpdateTime     *string                                       `json:"sourceUpdateTime"`
}

type jsonNoteVulnerabilityDetails NoteVulnerabilityDetails

func (r *NoteVulnerabilityDetails) UnmarshalJSON(data []byte) error {
	var res jsonNoteVulnerabilityDetails
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteVulnerabilityDetails
	} else {

		r.SeverityName = res.SeverityName

		r.Description = res.Description

		r.PackageType = res.PackageType

		r.AffectedCpeUri = res.AffectedCpeUri

		r.AffectedPackage = res.AffectedPackage

		r.AffectedVersionStart = res.AffectedVersionStart

		r.AffectedVersionEnd = res.AffectedVersionEnd

		r.FixedCpeUri = res.FixedCpeUri

		r.FixedPackage = res.FixedPackage

		r.FixedVersion = res.FixedVersion

		r.IsObsolete = res.IsObsolete

		r.SourceUpdateTime = res.SourceUpdateTime

	}
	return nil
}

// This object is used to assert a desired state where this NoteVulnerabilityDetails is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityDetails *NoteVulnerabilityDetails = &NoteVulnerabilityDetails{empty: true}

func (r *NoteVulnerabilityDetails) Empty() bool {
	return r.empty
}

func (r *NoteVulnerabilityDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityDetailsAffectedVersionStart struct {
	empty    bool                                                  `json:"-"`
	Epoch    *int64                                                `json:"epoch"`
	Name     *string                                               `json:"name"`
	Revision *string                                               `json:"revision"`
	Kind     *NoteVulnerabilityDetailsAffectedVersionStartKindEnum `json:"kind"`
	FullName *string                                               `json:"fullName"`
}

type jsonNoteVulnerabilityDetailsAffectedVersionStart NoteVulnerabilityDetailsAffectedVersionStart

func (r *NoteVulnerabilityDetailsAffectedVersionStart) UnmarshalJSON(data []byte) error {
	var res jsonNoteVulnerabilityDetailsAffectedVersionStart
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteVulnerabilityDetailsAffectedVersionStart
	} else {

		r.Epoch = res.Epoch

		r.Name = res.Name

		r.Revision = res.Revision

		r.Kind = res.Kind

		r.FullName = res.FullName

	}
	return nil
}

// This object is used to assert a desired state where this NoteVulnerabilityDetailsAffectedVersionStart is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityDetailsAffectedVersionStart *NoteVulnerabilityDetailsAffectedVersionStart = &NoteVulnerabilityDetailsAffectedVersionStart{empty: true}

func (r *NoteVulnerabilityDetailsAffectedVersionStart) Empty() bool {
	return r.empty
}

func (r *NoteVulnerabilityDetailsAffectedVersionStart) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityDetailsAffectedVersionStart) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityDetailsAffectedVersionEnd struct {
	empty    bool                                                `json:"-"`
	Epoch    *int64                                              `json:"epoch"`
	Name     *string                                             `json:"name"`
	Revision *string                                             `json:"revision"`
	Kind     *NoteVulnerabilityDetailsAffectedVersionEndKindEnum `json:"kind"`
	FullName *string                                             `json:"fullName"`
}

type jsonNoteVulnerabilityDetailsAffectedVersionEnd NoteVulnerabilityDetailsAffectedVersionEnd

func (r *NoteVulnerabilityDetailsAffectedVersionEnd) UnmarshalJSON(data []byte) error {
	var res jsonNoteVulnerabilityDetailsAffectedVersionEnd
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteVulnerabilityDetailsAffectedVersionEnd
	} else {

		r.Epoch = res.Epoch

		r.Name = res.Name

		r.Revision = res.Revision

		r.Kind = res.Kind

		r.FullName = res.FullName

	}
	return nil
}

// This object is used to assert a desired state where this NoteVulnerabilityDetailsAffectedVersionEnd is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityDetailsAffectedVersionEnd *NoteVulnerabilityDetailsAffectedVersionEnd = &NoteVulnerabilityDetailsAffectedVersionEnd{empty: true}

func (r *NoteVulnerabilityDetailsAffectedVersionEnd) Empty() bool {
	return r.empty
}

func (r *NoteVulnerabilityDetailsAffectedVersionEnd) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityDetailsAffectedVersionEnd) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityDetailsFixedVersion struct {
	empty    bool                                          `json:"-"`
	Epoch    *int64                                        `json:"epoch"`
	Name     *string                                       `json:"name"`
	Revision *string                                       `json:"revision"`
	Kind     *NoteVulnerabilityDetailsFixedVersionKindEnum `json:"kind"`
	FullName *string                                       `json:"fullName"`
}

type jsonNoteVulnerabilityDetailsFixedVersion NoteVulnerabilityDetailsFixedVersion

func (r *NoteVulnerabilityDetailsFixedVersion) UnmarshalJSON(data []byte) error {
	var res jsonNoteVulnerabilityDetailsFixedVersion
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteVulnerabilityDetailsFixedVersion
	} else {

		r.Epoch = res.Epoch

		r.Name = res.Name

		r.Revision = res.Revision

		r.Kind = res.Kind

		r.FullName = res.FullName

	}
	return nil
}

// This object is used to assert a desired state where this NoteVulnerabilityDetailsFixedVersion is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityDetailsFixedVersion *NoteVulnerabilityDetailsFixedVersion = &NoteVulnerabilityDetailsFixedVersion{empty: true}

func (r *NoteVulnerabilityDetailsFixedVersion) Empty() bool {
	return r.empty
}

func (r *NoteVulnerabilityDetailsFixedVersion) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityDetailsFixedVersion) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityCvssV3 struct {
	empty                 bool                                              `json:"-"`
	BaseScore             *float64                                          `json:"baseScore"`
	ExploitabilityScore   *float64                                          `json:"exploitabilityScore"`
	ImpactScore           *float64                                          `json:"impactScore"`
	AttackVector          *NoteVulnerabilityCvssV3AttackVectorEnum          `json:"attackVector"`
	AttackComplexity      *NoteVulnerabilityCvssV3AttackComplexityEnum      `json:"attackComplexity"`
	PrivilegesRequired    *NoteVulnerabilityCvssV3PrivilegesRequiredEnum    `json:"privilegesRequired"`
	UserInteraction       *NoteVulnerabilityCvssV3UserInteractionEnum       `json:"userInteraction"`
	Scope                 *NoteVulnerabilityCvssV3ScopeEnum                 `json:"scope"`
	ConfidentialityImpact *NoteVulnerabilityCvssV3ConfidentialityImpactEnum `json:"confidentialityImpact"`
	IntegrityImpact       *NoteVulnerabilityCvssV3IntegrityImpactEnum       `json:"integrityImpact"`
	AvailabilityImpact    *NoteVulnerabilityCvssV3AvailabilityImpactEnum    `json:"availabilityImpact"`
}

type jsonNoteVulnerabilityCvssV3 NoteVulnerabilityCvssV3

func (r *NoteVulnerabilityCvssV3) UnmarshalJSON(data []byte) error {
	var res jsonNoteVulnerabilityCvssV3
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteVulnerabilityCvssV3
	} else {

		r.BaseScore = res.BaseScore

		r.ExploitabilityScore = res.ExploitabilityScore

		r.ImpactScore = res.ImpactScore

		r.AttackVector = res.AttackVector

		r.AttackComplexity = res.AttackComplexity

		r.PrivilegesRequired = res.PrivilegesRequired

		r.UserInteraction = res.UserInteraction

		r.Scope = res.Scope

		r.ConfidentialityImpact = res.ConfidentialityImpact

		r.IntegrityImpact = res.IntegrityImpact

		r.AvailabilityImpact = res.AvailabilityImpact

	}
	return nil
}

// This object is used to assert a desired state where this NoteVulnerabilityCvssV3 is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityCvssV3 *NoteVulnerabilityCvssV3 = &NoteVulnerabilityCvssV3{empty: true}

func (r *NoteVulnerabilityCvssV3) Empty() bool {
	return r.empty
}

func (r *NoteVulnerabilityCvssV3) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityCvssV3) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityWindowsDetails struct {
	empty       bool                                       `json:"-"`
	CpeUri      *string                                    `json:"cpeUri"`
	Name        *string                                    `json:"name"`
	Description *string                                    `json:"description"`
	FixingKbs   []NoteVulnerabilityWindowsDetailsFixingKbs `json:"fixingKbs"`
}

type jsonNoteVulnerabilityWindowsDetails NoteVulnerabilityWindowsDetails

func (r *NoteVulnerabilityWindowsDetails) UnmarshalJSON(data []byte) error {
	var res jsonNoteVulnerabilityWindowsDetails
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteVulnerabilityWindowsDetails
	} else {

		r.CpeUri = res.CpeUri

		r.Name = res.Name

		r.Description = res.Description

		r.FixingKbs = res.FixingKbs

	}
	return nil
}

// This object is used to assert a desired state where this NoteVulnerabilityWindowsDetails is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityWindowsDetails *NoteVulnerabilityWindowsDetails = &NoteVulnerabilityWindowsDetails{empty: true}

func (r *NoteVulnerabilityWindowsDetails) Empty() bool {
	return r.empty
}

func (r *NoteVulnerabilityWindowsDetails) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityWindowsDetails) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteVulnerabilityWindowsDetailsFixingKbs struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
	Url   *string `json:"url"`
}

type jsonNoteVulnerabilityWindowsDetailsFixingKbs NoteVulnerabilityWindowsDetailsFixingKbs

func (r *NoteVulnerabilityWindowsDetailsFixingKbs) UnmarshalJSON(data []byte) error {
	var res jsonNoteVulnerabilityWindowsDetailsFixingKbs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteVulnerabilityWindowsDetailsFixingKbs
	} else {

		r.Name = res.Name

		r.Url = res.Url

	}
	return nil
}

// This object is used to assert a desired state where this NoteVulnerabilityWindowsDetailsFixingKbs is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteVulnerabilityWindowsDetailsFixingKbs *NoteVulnerabilityWindowsDetailsFixingKbs = &NoteVulnerabilityWindowsDetailsFixingKbs{empty: true}

func (r *NoteVulnerabilityWindowsDetailsFixingKbs) Empty() bool {
	return r.empty
}

func (r *NoteVulnerabilityWindowsDetailsFixingKbs) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteVulnerabilityWindowsDetailsFixingKbs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteBuild struct {
	empty          bool                `json:"-"`
	BuilderVersion *string             `json:"builderVersion"`
	Signature      *NoteBuildSignature `json:"signature"`
}

type jsonNoteBuild NoteBuild

func (r *NoteBuild) UnmarshalJSON(data []byte) error {
	var res jsonNoteBuild
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteBuild
	} else {

		r.BuilderVersion = res.BuilderVersion

		r.Signature = res.Signature

	}
	return nil
}

// This object is used to assert a desired state where this NoteBuild is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteBuild *NoteBuild = &NoteBuild{empty: true}

func (r *NoteBuild) Empty() bool {
	return r.empty
}

func (r *NoteBuild) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteBuild) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteBuildSignature struct {
	empty     bool                           `json:"-"`
	PublicKey *string                        `json:"publicKey"`
	Signature *string                        `json:"signature"`
	KeyId     *string                        `json:"keyId"`
	KeyType   *NoteBuildSignatureKeyTypeEnum `json:"keyType"`
}

type jsonNoteBuildSignature NoteBuildSignature

func (r *NoteBuildSignature) UnmarshalJSON(data []byte) error {
	var res jsonNoteBuildSignature
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteBuildSignature
	} else {

		r.PublicKey = res.PublicKey

		r.Signature = res.Signature

		r.KeyId = res.KeyId

		r.KeyType = res.KeyType

	}
	return nil
}

// This object is used to assert a desired state where this NoteBuildSignature is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteBuildSignature *NoteBuildSignature = &NoteBuildSignature{empty: true}

func (r *NoteBuildSignature) Empty() bool {
	return r.empty
}

func (r *NoteBuildSignature) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteBuildSignature) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteImage struct {
	empty       bool                  `json:"-"`
	ResourceUrl *string               `json:"resourceUrl"`
	Fingerprint *NoteImageFingerprint `json:"fingerprint"`
}

type jsonNoteImage NoteImage

func (r *NoteImage) UnmarshalJSON(data []byte) error {
	var res jsonNoteImage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteImage
	} else {

		r.ResourceUrl = res.ResourceUrl

		r.Fingerprint = res.Fingerprint

	}
	return nil
}

// This object is used to assert a desired state where this NoteImage is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteImage *NoteImage = &NoteImage{empty: true}

func (r *NoteImage) Empty() bool {
	return r.empty
}

func (r *NoteImage) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteImage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteImageFingerprint struct {
	empty  bool     `json:"-"`
	V1Name *string  `json:"v1Name"`
	V2Blob []string `json:"v2Blob"`
	V2Name *string  `json:"v2Name"`
}

type jsonNoteImageFingerprint NoteImageFingerprint

func (r *NoteImageFingerprint) UnmarshalJSON(data []byte) error {
	var res jsonNoteImageFingerprint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteImageFingerprint
	} else {

		r.V1Name = res.V1Name

		r.V2Blob = res.V2Blob

		r.V2Name = res.V2Name

	}
	return nil
}

// This object is used to assert a desired state where this NoteImageFingerprint is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteImageFingerprint *NoteImageFingerprint = &NoteImageFingerprint{empty: true}

func (r *NoteImageFingerprint) Empty() bool {
	return r.empty
}

func (r *NoteImageFingerprint) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteImageFingerprint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NotePackage struct {
	empty        bool                      `json:"-"`
	Name         *string                   `json:"name"`
	Distribution []NotePackageDistribution `json:"distribution"`
}

type jsonNotePackage NotePackage

func (r *NotePackage) UnmarshalJSON(data []byte) error {
	var res jsonNotePackage
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNotePackage
	} else {

		r.Name = res.Name

		r.Distribution = res.Distribution

	}
	return nil
}

// This object is used to assert a desired state where this NotePackage is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNotePackage *NotePackage = &NotePackage{empty: true}

func (r *NotePackage) Empty() bool {
	return r.empty
}

func (r *NotePackage) String() string {
	return dcl.SprintResource(r)
}

func (r *NotePackage) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NotePackageDistribution struct {
	empty         bool                                     `json:"-"`
	CpeUri        *string                                  `json:"cpeUri"`
	Architecture  *NotePackageDistributionArchitectureEnum `json:"architecture"`
	LatestVersion *NotePackageDistributionLatestVersion    `json:"latestVersion"`
	Maintainer    *string                                  `json:"maintainer"`
	Url           *string                                  `json:"url"`
	Description   *string                                  `json:"description"`
}

type jsonNotePackageDistribution NotePackageDistribution

func (r *NotePackageDistribution) UnmarshalJSON(data []byte) error {
	var res jsonNotePackageDistribution
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNotePackageDistribution
	} else {

		r.CpeUri = res.CpeUri

		r.Architecture = res.Architecture

		r.LatestVersion = res.LatestVersion

		r.Maintainer = res.Maintainer

		r.Url = res.Url

		r.Description = res.Description

	}
	return nil
}

// This object is used to assert a desired state where this NotePackageDistribution is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNotePackageDistribution *NotePackageDistribution = &NotePackageDistribution{empty: true}

func (r *NotePackageDistribution) Empty() bool {
	return r.empty
}

func (r *NotePackageDistribution) String() string {
	return dcl.SprintResource(r)
}

func (r *NotePackageDistribution) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NotePackageDistributionLatestVersion struct {
	empty    bool                                          `json:"-"`
	Epoch    *int64                                        `json:"epoch"`
	Name     *string                                       `json:"name"`
	Revision *string                                       `json:"revision"`
	Kind     *NotePackageDistributionLatestVersionKindEnum `json:"kind"`
	FullName *string                                       `json:"fullName"`
}

type jsonNotePackageDistributionLatestVersion NotePackageDistributionLatestVersion

func (r *NotePackageDistributionLatestVersion) UnmarshalJSON(data []byte) error {
	var res jsonNotePackageDistributionLatestVersion
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNotePackageDistributionLatestVersion
	} else {

		r.Epoch = res.Epoch

		r.Name = res.Name

		r.Revision = res.Revision

		r.Kind = res.Kind

		r.FullName = res.FullName

	}
	return nil
}

// This object is used to assert a desired state where this NotePackageDistributionLatestVersion is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNotePackageDistributionLatestVersion *NotePackageDistributionLatestVersion = &NotePackageDistributionLatestVersion{empty: true}

func (r *NotePackageDistributionLatestVersion) Empty() bool {
	return r.empty
}

func (r *NotePackageDistributionLatestVersion) String() string {
	return dcl.SprintResource(r)
}

func (r *NotePackageDistributionLatestVersion) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteDiscovery struct {
	empty        bool                           `json:"-"`
	AnalysisKind *NoteDiscoveryAnalysisKindEnum `json:"analysisKind"`
}

type jsonNoteDiscovery NoteDiscovery

func (r *NoteDiscovery) UnmarshalJSON(data []byte) error {
	var res jsonNoteDiscovery
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteDiscovery
	} else {

		r.AnalysisKind = res.AnalysisKind

	}
	return nil
}

// This object is used to assert a desired state where this NoteDiscovery is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteDiscovery *NoteDiscovery = &NoteDiscovery{empty: true}

func (r *NoteDiscovery) Empty() bool {
	return r.empty
}

func (r *NoteDiscovery) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteDiscovery) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteDeployment struct {
	empty       bool     `json:"-"`
	ResourceUri []string `json:"resourceUri"`
}

type jsonNoteDeployment NoteDeployment

func (r *NoteDeployment) UnmarshalJSON(data []byte) error {
	var res jsonNoteDeployment
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteDeployment
	} else {

		r.ResourceUri = res.ResourceUri

	}
	return nil
}

// This object is used to assert a desired state where this NoteDeployment is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteDeployment *NoteDeployment = &NoteDeployment{empty: true}

func (r *NoteDeployment) Empty() bool {
	return r.empty
}

func (r *NoteDeployment) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteDeployment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteAttestation struct {
	empty bool                 `json:"-"`
	Hint  *NoteAttestationHint `json:"hint"`
}

type jsonNoteAttestation NoteAttestation

func (r *NoteAttestation) UnmarshalJSON(data []byte) error {
	var res jsonNoteAttestation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteAttestation
	} else {

		r.Hint = res.Hint

	}
	return nil
}

// This object is used to assert a desired state where this NoteAttestation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteAttestation *NoteAttestation = &NoteAttestation{empty: true}

func (r *NoteAttestation) Empty() bool {
	return r.empty
}

func (r *NoteAttestation) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteAttestation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type NoteAttestationHint struct {
	empty             bool    `json:"-"`
	HumanReadableName *string `json:"humanReadableName"`
}

type jsonNoteAttestationHint NoteAttestationHint

func (r *NoteAttestationHint) UnmarshalJSON(data []byte) error {
	var res jsonNoteAttestationHint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyNoteAttestationHint
	} else {

		r.HumanReadableName = res.HumanReadableName

	}
	return nil
}

// This object is used to assert a desired state where this NoteAttestationHint is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyNoteAttestationHint *NoteAttestationHint = &NoteAttestationHint{empty: true}

func (r *NoteAttestationHint) Empty() bool {
	return r.empty
}

func (r *NoteAttestationHint) String() string {
	return dcl.SprintResource(r)
}

func (r *NoteAttestationHint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Note) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "container_analysis",
		Type:    "Note",
		Version: "alpha",
	}
}

func (r *Note) ID() (string, error) {
	if err := extractNoteFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":               dcl.ValueOrEmptyString(nr.Name),
		"short_description":  dcl.ValueOrEmptyString(nr.ShortDescription),
		"long_description":   dcl.ValueOrEmptyString(nr.LongDescription),
		"related_url":        dcl.ValueOrEmptyString(nr.RelatedUrl),
		"expiration_time":    dcl.ValueOrEmptyString(nr.ExpirationTime),
		"create_time":        dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":        dcl.ValueOrEmptyString(nr.UpdateTime),
		"related_note_names": dcl.ValueOrEmptyString(nr.RelatedNoteNames),
		"vulnerability":      dcl.ValueOrEmptyString(nr.Vulnerability),
		"build":              dcl.ValueOrEmptyString(nr.Build),
		"image":              dcl.ValueOrEmptyString(nr.Image),
		"package":            dcl.ValueOrEmptyString(nr.Package),
		"discovery":          dcl.ValueOrEmptyString(nr.Discovery),
		"deployment":         dcl.ValueOrEmptyString(nr.Deployment),
		"attestation":        dcl.ValueOrEmptyString(nr.Attestation),
		"project":            dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/notes/{{name}}", params), nil
}

const NoteMaxPage = -1

type NoteList struct {
	Items []*Note

	nextToken string

	pageSize int32

	resource *Note
}

func (l *NoteList) HasNext() bool {
	return l.nextToken != ""
}

func (l *NoteList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listNote(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListNote(ctx context.Context, project string) (*NoteList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListNoteWithMaxResults(ctx, project, NoteMaxPage)

}

func (c *Client) ListNoteWithMaxResults(ctx context.Context, project string, pageSize int32) (*NoteList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Note{
		Project: &project,
	}
	items, token, err := c.listNote(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &NoteList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetNote(ctx context.Context, r *Note) (*Note, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractNoteFields(r)

	b, err := c.getNoteRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalNote(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeNoteNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractNoteFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteNote(ctx context.Context, r *Note) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Note resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Note...")
	deleteOp := deleteNoteOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllNote deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllNote(ctx context.Context, project string, filter func(*Note) bool) error {
	listObj, err := c.ListNote(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllNote(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllNote(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyNote(ctx context.Context, rawDesired *Note, opts ...dcl.ApplyOption) (*Note, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Note
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyNoteHelper(c, ctx, rawDesired, opts...)
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

func applyNoteHelper(c *Client, ctx context.Context, rawDesired *Note, opts ...dcl.ApplyOption) (*Note, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyNote...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractNoteFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.noteDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToNoteDiffs(c.Config, fieldDiffs, opts)
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
	var ops []noteApiOperation
	if create {
		ops = append(ops, &createNoteOperation{})
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
	return applyNoteDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyNoteDiff(c *Client, ctx context.Context, desired *Note, rawDesired *Note, ops []noteApiOperation, opts ...dcl.ApplyOption) (*Note, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetNote(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createNoteOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapNote(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeNoteNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeNoteNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeNoteDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractNoteFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractNoteFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffNote(c, newDesired, newState)
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
