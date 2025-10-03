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
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type CryptoKey struct {
	Name                     *string                   `json:"name"`
	Primary                  *CryptoKeyPrimary         `json:"primary"`
	Purpose                  *CryptoKeyPurposeEnum     `json:"purpose"`
	CreateTime               *string                   `json:"createTime"`
	NextRotationTime         *string                   `json:"nextRotationTime"`
	RotationPeriod           *string                   `json:"rotationPeriod"`
	VersionTemplate          *CryptoKeyVersionTemplate `json:"versionTemplate"`
	Labels                   map[string]string         `json:"labels"`
	ImportOnly               *bool                     `json:"importOnly"`
	DestroyScheduledDuration *string                   `json:"destroyScheduledDuration"`
	Project                  *string                   `json:"project"`
	Location                 *string                   `json:"location"`
	KeyRing                  *string                   `json:"keyRing"`
}

func (r *CryptoKey) String() string {
	return dcl.SprintResource(r)
}

// The enum CryptoKeyPrimaryStateEnum.
type CryptoKeyPrimaryStateEnum string

// CryptoKeyPrimaryStateEnumRef returns a *CryptoKeyPrimaryStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func CryptoKeyPrimaryStateEnumRef(s string) *CryptoKeyPrimaryStateEnum {
	v := CryptoKeyPrimaryStateEnum(s)
	return &v
}

func (v CryptoKeyPrimaryStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"CRYPTO_KEY_VERSION_STATE_UNSPECIFIED", "PENDING_GENERATION", "ENABLED", "DISABLED", "DESTROYED", "DESTROY_SCHEDULED", "PENDING_IMPORT", "IMPORT_FAILED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CryptoKeyPrimaryStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CryptoKeyPrimaryProtectionLevelEnum.
type CryptoKeyPrimaryProtectionLevelEnum string

// CryptoKeyPrimaryProtectionLevelEnumRef returns a *CryptoKeyPrimaryProtectionLevelEnum with the value of string s
// If the empty string is provided, nil is returned.
func CryptoKeyPrimaryProtectionLevelEnumRef(s string) *CryptoKeyPrimaryProtectionLevelEnum {
	v := CryptoKeyPrimaryProtectionLevelEnum(s)
	return &v
}

func (v CryptoKeyPrimaryProtectionLevelEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"PROTECTION_LEVEL_UNSPECIFIED", "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CryptoKeyPrimaryProtectionLevelEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CryptoKeyPrimaryAlgorithmEnum.
type CryptoKeyPrimaryAlgorithmEnum string

// CryptoKeyPrimaryAlgorithmEnumRef returns a *CryptoKeyPrimaryAlgorithmEnum with the value of string s
// If the empty string is provided, nil is returned.
func CryptoKeyPrimaryAlgorithmEnumRef(s string) *CryptoKeyPrimaryAlgorithmEnum {
	v := CryptoKeyPrimaryAlgorithmEnum(s)
	return &v
}

func (v CryptoKeyPrimaryAlgorithmEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED", "GOOGLE_SYMMETRIC_ENCRYPTION", "RSA_SIGN_PSS_2048_SHA256", "RSA_SIGN_PSS_3072_SHA256", "RSA_SIGN_PSS_4096_SHA256", "RSA_SIGN_PSS_4096_SHA512", "RSA_SIGN_PKCS1_2048_SHA256", "RSA_SIGN_PKCS1_3072_SHA256", "RSA_SIGN_PKCS1_4096_SHA256", "RSA_SIGN_PKCS1_4096_SHA512", "RSA_DECRYPT_OAEP_2048_SHA256", "RSA_DECRYPT_OAEP_3072_SHA256", "RSA_DECRYPT_OAEP_4096_SHA256", "RSA_DECRYPT_OAEP_4096_SHA512", "EC_SIGN_P256_SHA256", "EC_SIGN_P384_SHA384", "EC_SIGN_SECP256K1_SHA256", "HMAC_SHA256", "EXTERNAL_SYMMETRIC_ENCRYPTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CryptoKeyPrimaryAlgorithmEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CryptoKeyPrimaryAttestationFormatEnum.
type CryptoKeyPrimaryAttestationFormatEnum string

// CryptoKeyPrimaryAttestationFormatEnumRef returns a *CryptoKeyPrimaryAttestationFormatEnum with the value of string s
// If the empty string is provided, nil is returned.
func CryptoKeyPrimaryAttestationFormatEnumRef(s string) *CryptoKeyPrimaryAttestationFormatEnum {
	v := CryptoKeyPrimaryAttestationFormatEnum(s)
	return &v
}

func (v CryptoKeyPrimaryAttestationFormatEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"ATTESTATION_FORMAT_UNSPECIFIED", "CAVIUM_V1_COMPRESSED", "CAVIUM_V2_COMPRESSED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CryptoKeyPrimaryAttestationFormatEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CryptoKeyPurposeEnum.
type CryptoKeyPurposeEnum string

// CryptoKeyPurposeEnumRef returns a *CryptoKeyPurposeEnum with the value of string s
// If the empty string is provided, nil is returned.
func CryptoKeyPurposeEnumRef(s string) *CryptoKeyPurposeEnum {
	v := CryptoKeyPurposeEnum(s)
	return &v
}

func (v CryptoKeyPurposeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"CRYPTO_KEY_PURPOSE_UNSPECIFIED", "ENCRYPT_DECRYPT", "ASYMMETRIC_SIGN", "ASYMMETRIC_DECRYPT", "MAC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CryptoKeyPurposeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CryptoKeyVersionTemplateProtectionLevelEnum.
type CryptoKeyVersionTemplateProtectionLevelEnum string

// CryptoKeyVersionTemplateProtectionLevelEnumRef returns a *CryptoKeyVersionTemplateProtectionLevelEnum with the value of string s
// If the empty string is provided, nil is returned.
func CryptoKeyVersionTemplateProtectionLevelEnumRef(s string) *CryptoKeyVersionTemplateProtectionLevelEnum {
	v := CryptoKeyVersionTemplateProtectionLevelEnum(s)
	return &v
}

func (v CryptoKeyVersionTemplateProtectionLevelEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"PROTECTION_LEVEL_UNSPECIFIED", "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CryptoKeyVersionTemplateProtectionLevelEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum CryptoKeyVersionTemplateAlgorithmEnum.
type CryptoKeyVersionTemplateAlgorithmEnum string

// CryptoKeyVersionTemplateAlgorithmEnumRef returns a *CryptoKeyVersionTemplateAlgorithmEnum with the value of string s
// If the empty string is provided, nil is returned.
func CryptoKeyVersionTemplateAlgorithmEnumRef(s string) *CryptoKeyVersionTemplateAlgorithmEnum {
	v := CryptoKeyVersionTemplateAlgorithmEnum(s)
	return &v
}

func (v CryptoKeyVersionTemplateAlgorithmEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED", "GOOGLE_SYMMETRIC_ENCRYPTION", "RSA_SIGN_PSS_2048_SHA256", "RSA_SIGN_PSS_3072_SHA256", "RSA_SIGN_PSS_4096_SHA256", "RSA_SIGN_PSS_4096_SHA512", "RSA_SIGN_PKCS1_2048_SHA256", "RSA_SIGN_PKCS1_3072_SHA256", "RSA_SIGN_PKCS1_4096_SHA256", "RSA_SIGN_PKCS1_4096_SHA512", "RSA_DECRYPT_OAEP_2048_SHA256", "RSA_DECRYPT_OAEP_3072_SHA256", "RSA_DECRYPT_OAEP_4096_SHA256", "RSA_DECRYPT_OAEP_4096_SHA512", "EC_SIGN_P256_SHA256", "EC_SIGN_P384_SHA384", "EC_SIGN_SECP256K1_SHA256", "HMAC_SHA256", "EXTERNAL_SYMMETRIC_ENCRYPTION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "CryptoKeyVersionTemplateAlgorithmEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type CryptoKeyPrimary struct {
	empty                          bool                                            `json:"-"`
	Name                           *string                                         `json:"name"`
	State                          *CryptoKeyPrimaryStateEnum                      `json:"state"`
	ProtectionLevel                *CryptoKeyPrimaryProtectionLevelEnum            `json:"protectionLevel"`
	Algorithm                      *CryptoKeyPrimaryAlgorithmEnum                  `json:"algorithm"`
	Attestation                    *CryptoKeyPrimaryAttestation                    `json:"attestation"`
	CreateTime                     *string                                         `json:"createTime"`
	GenerateTime                   *string                                         `json:"generateTime"`
	DestroyTime                    *string                                         `json:"destroyTime"`
	DestroyEventTime               *string                                         `json:"destroyEventTime"`
	ImportJob                      *string                                         `json:"importJob"`
	ImportTime                     *string                                         `json:"importTime"`
	ImportFailureReason            *string                                         `json:"importFailureReason"`
	ExternalProtectionLevelOptions *CryptoKeyPrimaryExternalProtectionLevelOptions `json:"externalProtectionLevelOptions"`
	ReimportEligible               *bool                                           `json:"reimportEligible"`
}

type jsonCryptoKeyPrimary CryptoKeyPrimary

func (r *CryptoKeyPrimary) UnmarshalJSON(data []byte) error {
	var res jsonCryptoKeyPrimary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCryptoKeyPrimary
	} else {

		r.Name = res.Name

		r.State = res.State

		r.ProtectionLevel = res.ProtectionLevel

		r.Algorithm = res.Algorithm

		r.Attestation = res.Attestation

		r.CreateTime = res.CreateTime

		r.GenerateTime = res.GenerateTime

		r.DestroyTime = res.DestroyTime

		r.DestroyEventTime = res.DestroyEventTime

		r.ImportJob = res.ImportJob

		r.ImportTime = res.ImportTime

		r.ImportFailureReason = res.ImportFailureReason

		r.ExternalProtectionLevelOptions = res.ExternalProtectionLevelOptions

		r.ReimportEligible = res.ReimportEligible

	}
	return nil
}

// This object is used to assert a desired state where this CryptoKeyPrimary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyCryptoKeyPrimary *CryptoKeyPrimary = &CryptoKeyPrimary{empty: true}

func (r *CryptoKeyPrimary) Empty() bool {
	return r.empty
}

func (r *CryptoKeyPrimary) String() string {
	return dcl.SprintResource(r)
}

func (r *CryptoKeyPrimary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CryptoKeyPrimaryAttestation struct {
	empty      bool                                   `json:"-"`
	Format     *CryptoKeyPrimaryAttestationFormatEnum `json:"format"`
	Content    *string                                `json:"content"`
	CertChains *CryptoKeyPrimaryAttestationCertChains `json:"certChains"`
}

type jsonCryptoKeyPrimaryAttestation CryptoKeyPrimaryAttestation

func (r *CryptoKeyPrimaryAttestation) UnmarshalJSON(data []byte) error {
	var res jsonCryptoKeyPrimaryAttestation
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCryptoKeyPrimaryAttestation
	} else {

		r.Format = res.Format

		r.Content = res.Content

		r.CertChains = res.CertChains

	}
	return nil
}

// This object is used to assert a desired state where this CryptoKeyPrimaryAttestation is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyCryptoKeyPrimaryAttestation *CryptoKeyPrimaryAttestation = &CryptoKeyPrimaryAttestation{empty: true}

func (r *CryptoKeyPrimaryAttestation) Empty() bool {
	return r.empty
}

func (r *CryptoKeyPrimaryAttestation) String() string {
	return dcl.SprintResource(r)
}

func (r *CryptoKeyPrimaryAttestation) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CryptoKeyPrimaryAttestationCertChains struct {
	empty                bool     `json:"-"`
	CaviumCerts          []string `json:"caviumCerts"`
	GoogleCardCerts      []string `json:"googleCardCerts"`
	GooglePartitionCerts []string `json:"googlePartitionCerts"`
}

type jsonCryptoKeyPrimaryAttestationCertChains CryptoKeyPrimaryAttestationCertChains

func (r *CryptoKeyPrimaryAttestationCertChains) UnmarshalJSON(data []byte) error {
	var res jsonCryptoKeyPrimaryAttestationCertChains
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCryptoKeyPrimaryAttestationCertChains
	} else {

		r.CaviumCerts = res.CaviumCerts

		r.GoogleCardCerts = res.GoogleCardCerts

		r.GooglePartitionCerts = res.GooglePartitionCerts

	}
	return nil
}

// This object is used to assert a desired state where this CryptoKeyPrimaryAttestationCertChains is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyCryptoKeyPrimaryAttestationCertChains *CryptoKeyPrimaryAttestationCertChains = &CryptoKeyPrimaryAttestationCertChains{empty: true}

func (r *CryptoKeyPrimaryAttestationCertChains) Empty() bool {
	return r.empty
}

func (r *CryptoKeyPrimaryAttestationCertChains) String() string {
	return dcl.SprintResource(r)
}

func (r *CryptoKeyPrimaryAttestationCertChains) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CryptoKeyPrimaryExternalProtectionLevelOptions struct {
	empty          bool    `json:"-"`
	ExternalKeyUri *string `json:"externalKeyUri"`
}

type jsonCryptoKeyPrimaryExternalProtectionLevelOptions CryptoKeyPrimaryExternalProtectionLevelOptions

func (r *CryptoKeyPrimaryExternalProtectionLevelOptions) UnmarshalJSON(data []byte) error {
	var res jsonCryptoKeyPrimaryExternalProtectionLevelOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCryptoKeyPrimaryExternalProtectionLevelOptions
	} else {

		r.ExternalKeyUri = res.ExternalKeyUri

	}
	return nil
}

// This object is used to assert a desired state where this CryptoKeyPrimaryExternalProtectionLevelOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyCryptoKeyPrimaryExternalProtectionLevelOptions *CryptoKeyPrimaryExternalProtectionLevelOptions = &CryptoKeyPrimaryExternalProtectionLevelOptions{empty: true}

func (r *CryptoKeyPrimaryExternalProtectionLevelOptions) Empty() bool {
	return r.empty
}

func (r *CryptoKeyPrimaryExternalProtectionLevelOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *CryptoKeyPrimaryExternalProtectionLevelOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type CryptoKeyVersionTemplate struct {
	empty           bool                                         `json:"-"`
	ProtectionLevel *CryptoKeyVersionTemplateProtectionLevelEnum `json:"protectionLevel"`
	Algorithm       *CryptoKeyVersionTemplateAlgorithmEnum       `json:"algorithm"`
}

type jsonCryptoKeyVersionTemplate CryptoKeyVersionTemplate

func (r *CryptoKeyVersionTemplate) UnmarshalJSON(data []byte) error {
	var res jsonCryptoKeyVersionTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyCryptoKeyVersionTemplate
	} else {

		r.ProtectionLevel = res.ProtectionLevel

		r.Algorithm = res.Algorithm

	}
	return nil
}

// This object is used to assert a desired state where this CryptoKeyVersionTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyCryptoKeyVersionTemplate *CryptoKeyVersionTemplate = &CryptoKeyVersionTemplate{empty: true}

func (r *CryptoKeyVersionTemplate) Empty() bool {
	return r.empty
}

func (r *CryptoKeyVersionTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *CryptoKeyVersionTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *CryptoKey) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloudkms",
		Type:    "CryptoKey",
		Version: "alpha",
	}
}

func (r *CryptoKey) ID() (string, error) {
	if err := extractCryptoKeyFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                       dcl.ValueOrEmptyString(nr.Name),
		"primary":                    dcl.ValueOrEmptyString(nr.Primary),
		"purpose":                    dcl.ValueOrEmptyString(nr.Purpose),
		"create_time":                dcl.ValueOrEmptyString(nr.CreateTime),
		"next_rotation_time":         dcl.ValueOrEmptyString(nr.NextRotationTime),
		"rotation_period":            dcl.ValueOrEmptyString(nr.RotationPeriod),
		"version_template":           dcl.ValueOrEmptyString(nr.VersionTemplate),
		"labels":                     dcl.ValueOrEmptyString(nr.Labels),
		"import_only":                dcl.ValueOrEmptyString(nr.ImportOnly),
		"destroy_scheduled_duration": dcl.ValueOrEmptyString(nr.DestroyScheduledDuration),
		"project":                    dcl.ValueOrEmptyString(nr.Project),
		"location":                   dcl.ValueOrEmptyString(nr.Location),
		"key_ring":                   dcl.ValueOrEmptyString(nr.KeyRing),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/keyRings/{{key_ring}}/cryptoKeys/{{name}}", params), nil
}

const CryptoKeyMaxPage = -1

type CryptoKeyList struct {
	Items []*CryptoKey

	nextToken string

	pageSize int32

	resource *CryptoKey
}

func (l *CryptoKeyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *CryptoKeyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listCryptoKey(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListCryptoKey(ctx context.Context, project, location, keyRing string) (*CryptoKeyList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListCryptoKeyWithMaxResults(ctx, project, location, keyRing, CryptoKeyMaxPage)

}

func (c *Client) ListCryptoKeyWithMaxResults(ctx context.Context, project, location, keyRing string, pageSize int32) (*CryptoKeyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &CryptoKey{
		Project:  &project,
		Location: &location,
		KeyRing:  &keyRing,
	}
	items, token, err := c.listCryptoKey(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetCryptoKey(ctx context.Context, r *CryptoKey) (*CryptoKey, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractCryptoKeyFields(r)

	b, err := c.getCryptoKeyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalCryptoKey(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.KeyRing = r.KeyRing
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeCryptoKeyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractCryptoKeyFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyCryptoKey(ctx context.Context, rawDesired *CryptoKey, opts ...dcl.ApplyOption) (*CryptoKey, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *CryptoKey
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyCryptoKeyHelper(c, ctx, rawDesired, opts...)
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

func applyCryptoKeyHelper(c *Client, ctx context.Context, rawDesired *CryptoKey, opts ...dcl.ApplyOption) (*CryptoKey, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyCryptoKey...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractCryptoKeyFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.cryptoKeyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToCryptoKeyDiffs(c.Config, fieldDiffs, opts)
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
	var ops []cryptoKeyApiOperation
	if create {
		ops = append(ops, &createCryptoKeyOperation{})
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
	return applyCryptoKeyDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyCryptoKeyDiff(c *Client, ctx context.Context, desired *CryptoKey, rawDesired *CryptoKey, ops []cryptoKeyApiOperation, opts ...dcl.ApplyOption) (*CryptoKey, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetCryptoKey(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createCryptoKeyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapCryptoKey(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeCryptoKeyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeCryptoKeyNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeCryptoKeyDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractCryptoKeyFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractCryptoKeyFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffCryptoKey(c, newDesired, newState)
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

func (r *CryptoKey) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
