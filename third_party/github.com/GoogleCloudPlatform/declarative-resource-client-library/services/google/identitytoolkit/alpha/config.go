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

type Config struct {
	SignIn            *ConfigSignIn            `json:"signIn"`
	Notification      *ConfigNotification      `json:"notification"`
	Quota             *ConfigQuota             `json:"quota"`
	Monitoring        *ConfigMonitoring        `json:"monitoring"`
	MultiTenant       *ConfigMultiTenant       `json:"multiTenant"`
	AuthorizedDomains []string                 `json:"authorizedDomains"`
	Subtype           *ConfigSubtypeEnum       `json:"subtype"`
	Client            *ConfigClient            `json:"client"`
	Mfa               *ConfigMfa               `json:"mfa"`
	BlockingFunctions *ConfigBlockingFunctions `json:"blockingFunctions"`
	Project           *string                  `json:"project"`
}

func (r *Config) String() string {
	return dcl.SprintResource(r)
}

// The enum ConfigSignInEmailHashConfigAlgorithmEnum.
type ConfigSignInEmailHashConfigAlgorithmEnum string

// ConfigSignInEmailHashConfigAlgorithmEnumRef returns a *ConfigSignInEmailHashConfigAlgorithmEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigSignInEmailHashConfigAlgorithmEnumRef(s string) *ConfigSignInEmailHashConfigAlgorithmEnum {
	v := ConfigSignInEmailHashConfigAlgorithmEnum(s)
	return &v
}

func (v ConfigSignInEmailHashConfigAlgorithmEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"HASH_ALGORITHM_UNSPECIFIED", "HMAC_SHA256", "HMAC_SHA1", "HMAC_MD5", "SCRYPT", "PBKDF_SHA1", "MD5", "HMAC_SHA512", "SHA1", "BCRYPT", "PBKDF2_SHA256", "SHA256", "SHA512", "STANDARD_SCRYPT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigSignInEmailHashConfigAlgorithmEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigSignInHashConfigAlgorithmEnum.
type ConfigSignInHashConfigAlgorithmEnum string

// ConfigSignInHashConfigAlgorithmEnumRef returns a *ConfigSignInHashConfigAlgorithmEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigSignInHashConfigAlgorithmEnumRef(s string) *ConfigSignInHashConfigAlgorithmEnum {
	v := ConfigSignInHashConfigAlgorithmEnum(s)
	return &v
}

func (v ConfigSignInHashConfigAlgorithmEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"HASH_ALGORITHM_UNSPECIFIED", "HMAC_SHA256", "HMAC_SHA1", "HMAC_MD5", "SCRYPT", "PBKDF_SHA1", "MD5", "HMAC_SHA512", "SHA1", "BCRYPT", "PBKDF2_SHA256", "SHA256", "SHA512", "STANDARD_SCRYPT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigSignInHashConfigAlgorithmEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigNotificationSendEmailMethodEnum.
type ConfigNotificationSendEmailMethodEnum string

// ConfigNotificationSendEmailMethodEnumRef returns a *ConfigNotificationSendEmailMethodEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigNotificationSendEmailMethodEnumRef(s string) *ConfigNotificationSendEmailMethodEnum {
	v := ConfigNotificationSendEmailMethodEnum(s)
	return &v
}

func (v ConfigNotificationSendEmailMethodEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"METHOD_UNSPECIFIED", "DEFAULT", "CUSTOM_SMTP"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigNotificationSendEmailMethodEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigNotificationSendEmailSmtpSecurityModeEnum.
type ConfigNotificationSendEmailSmtpSecurityModeEnum string

// ConfigNotificationSendEmailSmtpSecurityModeEnumRef returns a *ConfigNotificationSendEmailSmtpSecurityModeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigNotificationSendEmailSmtpSecurityModeEnumRef(s string) *ConfigNotificationSendEmailSmtpSecurityModeEnum {
	v := ConfigNotificationSendEmailSmtpSecurityModeEnum(s)
	return &v
}

func (v ConfigNotificationSendEmailSmtpSecurityModeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SECURITY_MODE_UNSPECIFIED", "SSL", "START_TLS"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigNotificationSendEmailSmtpSecurityModeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum.
type ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum string

// ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumRef returns a *ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumRef(s string) *ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	v := ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(s)
	return &v
}

func (v ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"BODY_FORMAT_UNSPECIFIED", "PLAIN_TEXT", "HTML"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum.
type ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum string

// ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumRef returns a *ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumRef(s string) *ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	v := ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(s)
	return &v
}

func (v ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"BODY_FORMAT_UNSPECIFIED", "PLAIN_TEXT", "HTML"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum.
type ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum string

// ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumRef returns a *ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumRef(s string) *ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	v := ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(s)
	return &v
}

func (v ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"BODY_FORMAT_UNSPECIFIED", "PLAIN_TEXT", "HTML"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum.
type ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum string

// ConfigNotificationSendEmailDnsInfoCustomDomainStateEnumRef returns a *ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigNotificationSendEmailDnsInfoCustomDomainStateEnumRef(s string) *ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	v := ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(s)
	return &v
}

func (v ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"VERIFICATION_STATE_UNSPECIFIED", "NOT_STARTED", "IN_PROGRESS", "FAILED", "SUCCEEDED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum.
type ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum string

// ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumRef returns a *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumRef(s string) *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	v := ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(s)
	return &v
}

func (v ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"BODY_FORMAT_UNSPECIFIED", "PLAIN_TEXT", "HTML"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigSubtypeEnum.
type ConfigSubtypeEnum string

// ConfigSubtypeEnumRef returns a *ConfigSubtypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigSubtypeEnumRef(s string) *ConfigSubtypeEnum {
	v := ConfigSubtypeEnum(s)
	return &v
}

func (v ConfigSubtypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SUBTYPE_UNSPECIFIED", "IDENTITY_PLATFORM", "FIREBASE_AUTH"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigSubtypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ConfigMfaStateEnum.
type ConfigMfaStateEnum string

// ConfigMfaStateEnumRef returns a *ConfigMfaStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ConfigMfaStateEnumRef(s string) *ConfigMfaStateEnum {
	v := ConfigMfaStateEnum(s)
	return &v
}

func (v ConfigMfaStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "DISABLED", "ENABLED", "MANDATORY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ConfigMfaStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ConfigSignIn struct {
	empty                bool                     `json:"-"`
	Email                *ConfigSignInEmail       `json:"email"`
	PhoneNumber          *ConfigSignInPhoneNumber `json:"phoneNumber"`
	Anonymous            *ConfigSignInAnonymous   `json:"anonymous"`
	AllowDuplicateEmails *bool                    `json:"allowDuplicateEmails"`
	HashConfig           *ConfigSignInHashConfig  `json:"hashConfig"`
}

type jsonConfigSignIn ConfigSignIn

func (r *ConfigSignIn) UnmarshalJSON(data []byte) error {
	var res jsonConfigSignIn
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigSignIn
	} else {

		r.Email = res.Email

		r.PhoneNumber = res.PhoneNumber

		r.Anonymous = res.Anonymous

		r.AllowDuplicateEmails = res.AllowDuplicateEmails

		r.HashConfig = res.HashConfig

	}
	return nil
}

// This object is used to assert a desired state where this ConfigSignIn is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigSignIn *ConfigSignIn = &ConfigSignIn{empty: true}

func (r *ConfigSignIn) Empty() bool {
	return r.empty
}

func (r *ConfigSignIn) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigSignIn) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigSignInEmail struct {
	empty            bool                         `json:"-"`
	Enabled          *bool                        `json:"enabled"`
	PasswordRequired *bool                        `json:"passwordRequired"`
	HashConfig       *ConfigSignInEmailHashConfig `json:"hashConfig"`
}

type jsonConfigSignInEmail ConfigSignInEmail

func (r *ConfigSignInEmail) UnmarshalJSON(data []byte) error {
	var res jsonConfigSignInEmail
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigSignInEmail
	} else {

		r.Enabled = res.Enabled

		r.PasswordRequired = res.PasswordRequired

		r.HashConfig = res.HashConfig

	}
	return nil
}

// This object is used to assert a desired state where this ConfigSignInEmail is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigSignInEmail *ConfigSignInEmail = &ConfigSignInEmail{empty: true}

func (r *ConfigSignInEmail) Empty() bool {
	return r.empty
}

func (r *ConfigSignInEmail) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigSignInEmail) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigSignInEmailHashConfig struct {
	empty         bool                                      `json:"-"`
	Algorithm     *ConfigSignInEmailHashConfigAlgorithmEnum `json:"algorithm"`
	SignerKey     *string                                   `json:"signerKey"`
	SaltSeparator *string                                   `json:"saltSeparator"`
	Rounds        *int64                                    `json:"rounds"`
	MemoryCost    *int64                                    `json:"memoryCost"`
}

type jsonConfigSignInEmailHashConfig ConfigSignInEmailHashConfig

func (r *ConfigSignInEmailHashConfig) UnmarshalJSON(data []byte) error {
	var res jsonConfigSignInEmailHashConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigSignInEmailHashConfig
	} else {

		r.Algorithm = res.Algorithm

		r.SignerKey = res.SignerKey

		r.SaltSeparator = res.SaltSeparator

		r.Rounds = res.Rounds

		r.MemoryCost = res.MemoryCost

	}
	return nil
}

// This object is used to assert a desired state where this ConfigSignInEmailHashConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigSignInEmailHashConfig *ConfigSignInEmailHashConfig = &ConfigSignInEmailHashConfig{empty: true}

func (r *ConfigSignInEmailHashConfig) Empty() bool {
	return r.empty
}

func (r *ConfigSignInEmailHashConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigSignInEmailHashConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigSignInPhoneNumber struct {
	empty            bool              `json:"-"`
	Enabled          *bool             `json:"enabled"`
	TestPhoneNumbers map[string]string `json:"testPhoneNumbers"`
}

type jsonConfigSignInPhoneNumber ConfigSignInPhoneNumber

func (r *ConfigSignInPhoneNumber) UnmarshalJSON(data []byte) error {
	var res jsonConfigSignInPhoneNumber
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigSignInPhoneNumber
	} else {

		r.Enabled = res.Enabled

		r.TestPhoneNumbers = res.TestPhoneNumbers

	}
	return nil
}

// This object is used to assert a desired state where this ConfigSignInPhoneNumber is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigSignInPhoneNumber *ConfigSignInPhoneNumber = &ConfigSignInPhoneNumber{empty: true}

func (r *ConfigSignInPhoneNumber) Empty() bool {
	return r.empty
}

func (r *ConfigSignInPhoneNumber) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigSignInPhoneNumber) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigSignInAnonymous struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

type jsonConfigSignInAnonymous ConfigSignInAnonymous

func (r *ConfigSignInAnonymous) UnmarshalJSON(data []byte) error {
	var res jsonConfigSignInAnonymous
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigSignInAnonymous
	} else {

		r.Enabled = res.Enabled

	}
	return nil
}

// This object is used to assert a desired state where this ConfigSignInAnonymous is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigSignInAnonymous *ConfigSignInAnonymous = &ConfigSignInAnonymous{empty: true}

func (r *ConfigSignInAnonymous) Empty() bool {
	return r.empty
}

func (r *ConfigSignInAnonymous) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigSignInAnonymous) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigSignInHashConfig struct {
	empty         bool                                 `json:"-"`
	Algorithm     *ConfigSignInHashConfigAlgorithmEnum `json:"algorithm"`
	SignerKey     *string                              `json:"signerKey"`
	SaltSeparator *string                              `json:"saltSeparator"`
	Rounds        *int64                               `json:"rounds"`
	MemoryCost    *int64                               `json:"memoryCost"`
}

type jsonConfigSignInHashConfig ConfigSignInHashConfig

func (r *ConfigSignInHashConfig) UnmarshalJSON(data []byte) error {
	var res jsonConfigSignInHashConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigSignInHashConfig
	} else {

		r.Algorithm = res.Algorithm

		r.SignerKey = res.SignerKey

		r.SaltSeparator = res.SaltSeparator

		r.Rounds = res.Rounds

		r.MemoryCost = res.MemoryCost

	}
	return nil
}

// This object is used to assert a desired state where this ConfigSignInHashConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigSignInHashConfig *ConfigSignInHashConfig = &ConfigSignInHashConfig{empty: true}

func (r *ConfigSignInHashConfig) Empty() bool {
	return r.empty
}

func (r *ConfigSignInHashConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigSignInHashConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotification struct {
	empty         bool                         `json:"-"`
	SendEmail     *ConfigNotificationSendEmail `json:"sendEmail"`
	SendSms       *ConfigNotificationSendSms   `json:"sendSms"`
	DefaultLocale *string                      `json:"defaultLocale"`
}

type jsonConfigNotification ConfigNotification

func (r *ConfigNotification) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotification
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotification
	} else {

		r.SendEmail = res.SendEmail

		r.SendSms = res.SendSms

		r.DefaultLocale = res.DefaultLocale

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotification is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotification *ConfigNotification = &ConfigNotification{empty: true}

func (r *ConfigNotification) Empty() bool {
	return r.empty
}

func (r *ConfigNotification) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotification) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendEmail struct {
	empty                              bool                                                           `json:"-"`
	Method                             *ConfigNotificationSendEmailMethodEnum                         `json:"method"`
	Smtp                               *ConfigNotificationSendEmailSmtp                               `json:"smtp"`
	ResetPasswordTemplate              *ConfigNotificationSendEmailResetPasswordTemplate              `json:"resetPasswordTemplate"`
	VerifyEmailTemplate                *ConfigNotificationSendEmailVerifyEmailTemplate                `json:"verifyEmailTemplate"`
	ChangeEmailTemplate                *ConfigNotificationSendEmailChangeEmailTemplate                `json:"changeEmailTemplate"`
	CallbackUri                        *string                                                        `json:"callbackUri"`
	DnsInfo                            *ConfigNotificationSendEmailDnsInfo                            `json:"dnsInfo"`
	RevertSecondFactorAdditionTemplate *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate `json:"revertSecondFactorAdditionTemplate"`
}

type jsonConfigNotificationSendEmail ConfigNotificationSendEmail

func (r *ConfigNotificationSendEmail) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendEmail
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendEmail
	} else {

		r.Method = res.Method

		r.Smtp = res.Smtp

		r.ResetPasswordTemplate = res.ResetPasswordTemplate

		r.VerifyEmailTemplate = res.VerifyEmailTemplate

		r.ChangeEmailTemplate = res.ChangeEmailTemplate

		r.CallbackUri = res.CallbackUri

		r.DnsInfo = res.DnsInfo

		r.RevertSecondFactorAdditionTemplate = res.RevertSecondFactorAdditionTemplate

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendEmail is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendEmail *ConfigNotificationSendEmail = &ConfigNotificationSendEmail{empty: true}

func (r *ConfigNotificationSendEmail) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendEmail) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendEmail) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendEmailSmtp struct {
	empty        bool                                             `json:"-"`
	SenderEmail  *string                                          `json:"senderEmail"`
	Host         *string                                          `json:"host"`
	Port         *int64                                           `json:"port"`
	Username     *string                                          `json:"username"`
	Password     *string                                          `json:"password"`
	SecurityMode *ConfigNotificationSendEmailSmtpSecurityModeEnum `json:"securityMode"`
}

type jsonConfigNotificationSendEmailSmtp ConfigNotificationSendEmailSmtp

func (r *ConfigNotificationSendEmailSmtp) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendEmailSmtp
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendEmailSmtp
	} else {

		r.SenderEmail = res.SenderEmail

		r.Host = res.Host

		r.Port = res.Port

		r.Username = res.Username

		r.Password = res.Password

		r.SecurityMode = res.SecurityMode

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendEmailSmtp is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendEmailSmtp *ConfigNotificationSendEmailSmtp = &ConfigNotificationSendEmailSmtp{empty: true}

func (r *ConfigNotificationSendEmailSmtp) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendEmailSmtp) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendEmailSmtp) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendEmailResetPasswordTemplate struct {
	empty             bool                                                            `json:"-"`
	SenderLocalPart   *string                                                         `json:"senderLocalPart"`
	Subject           *string                                                         `json:"subject"`
	SenderDisplayName *string                                                         `json:"senderDisplayName"`
	Body              *string                                                         `json:"body"`
	BodyFormat        *ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum `json:"bodyFormat"`
	ReplyTo           *string                                                         `json:"replyTo"`
	Customized        *bool                                                           `json:"customized"`
}

type jsonConfigNotificationSendEmailResetPasswordTemplate ConfigNotificationSendEmailResetPasswordTemplate

func (r *ConfigNotificationSendEmailResetPasswordTemplate) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendEmailResetPasswordTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendEmailResetPasswordTemplate
	} else {

		r.SenderLocalPart = res.SenderLocalPart

		r.Subject = res.Subject

		r.SenderDisplayName = res.SenderDisplayName

		r.Body = res.Body

		r.BodyFormat = res.BodyFormat

		r.ReplyTo = res.ReplyTo

		r.Customized = res.Customized

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendEmailResetPasswordTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendEmailResetPasswordTemplate *ConfigNotificationSendEmailResetPasswordTemplate = &ConfigNotificationSendEmailResetPasswordTemplate{empty: true}

func (r *ConfigNotificationSendEmailResetPasswordTemplate) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendEmailResetPasswordTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendEmailResetPasswordTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendEmailVerifyEmailTemplate struct {
	empty             bool                                                          `json:"-"`
	SenderLocalPart   *string                                                       `json:"senderLocalPart"`
	Subject           *string                                                       `json:"subject"`
	SenderDisplayName *string                                                       `json:"senderDisplayName"`
	Body              *string                                                       `json:"body"`
	BodyFormat        *ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum `json:"bodyFormat"`
	ReplyTo           *string                                                       `json:"replyTo"`
	Customized        *bool                                                         `json:"customized"`
}

type jsonConfigNotificationSendEmailVerifyEmailTemplate ConfigNotificationSendEmailVerifyEmailTemplate

func (r *ConfigNotificationSendEmailVerifyEmailTemplate) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendEmailVerifyEmailTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendEmailVerifyEmailTemplate
	} else {

		r.SenderLocalPart = res.SenderLocalPart

		r.Subject = res.Subject

		r.SenderDisplayName = res.SenderDisplayName

		r.Body = res.Body

		r.BodyFormat = res.BodyFormat

		r.ReplyTo = res.ReplyTo

		r.Customized = res.Customized

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendEmailVerifyEmailTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendEmailVerifyEmailTemplate *ConfigNotificationSendEmailVerifyEmailTemplate = &ConfigNotificationSendEmailVerifyEmailTemplate{empty: true}

func (r *ConfigNotificationSendEmailVerifyEmailTemplate) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendEmailVerifyEmailTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendEmailVerifyEmailTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendEmailChangeEmailTemplate struct {
	empty             bool                                                          `json:"-"`
	SenderLocalPart   *string                                                       `json:"senderLocalPart"`
	Subject           *string                                                       `json:"subject"`
	SenderDisplayName *string                                                       `json:"senderDisplayName"`
	Body              *string                                                       `json:"body"`
	BodyFormat        *ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum `json:"bodyFormat"`
	ReplyTo           *string                                                       `json:"replyTo"`
	Customized        *bool                                                         `json:"customized"`
}

type jsonConfigNotificationSendEmailChangeEmailTemplate ConfigNotificationSendEmailChangeEmailTemplate

func (r *ConfigNotificationSendEmailChangeEmailTemplate) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendEmailChangeEmailTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendEmailChangeEmailTemplate
	} else {

		r.SenderLocalPart = res.SenderLocalPart

		r.Subject = res.Subject

		r.SenderDisplayName = res.SenderDisplayName

		r.Body = res.Body

		r.BodyFormat = res.BodyFormat

		r.ReplyTo = res.ReplyTo

		r.Customized = res.Customized

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendEmailChangeEmailTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendEmailChangeEmailTemplate *ConfigNotificationSendEmailChangeEmailTemplate = &ConfigNotificationSendEmailChangeEmailTemplate{empty: true}

func (r *ConfigNotificationSendEmailChangeEmailTemplate) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendEmailChangeEmailTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendEmailChangeEmailTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendEmailDnsInfo struct {
	empty                         bool                                                     `json:"-"`
	CustomDomain                  *string                                                  `json:"customDomain"`
	UseCustomDomain               *bool                                                    `json:"useCustomDomain"`
	PendingCustomDomain           *string                                                  `json:"pendingCustomDomain"`
	CustomDomainState             *ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum `json:"customDomainState"`
	DomainVerificationRequestTime *string                                                  `json:"domainVerificationRequestTime"`
}

type jsonConfigNotificationSendEmailDnsInfo ConfigNotificationSendEmailDnsInfo

func (r *ConfigNotificationSendEmailDnsInfo) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendEmailDnsInfo
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendEmailDnsInfo
	} else {

		r.CustomDomain = res.CustomDomain

		r.UseCustomDomain = res.UseCustomDomain

		r.PendingCustomDomain = res.PendingCustomDomain

		r.CustomDomainState = res.CustomDomainState

		r.DomainVerificationRequestTime = res.DomainVerificationRequestTime

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendEmailDnsInfo is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendEmailDnsInfo *ConfigNotificationSendEmailDnsInfo = &ConfigNotificationSendEmailDnsInfo{empty: true}

func (r *ConfigNotificationSendEmailDnsInfo) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendEmailDnsInfo) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendEmailDnsInfo) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate struct {
	empty             bool                                                                         `json:"-"`
	SenderLocalPart   *string                                                                      `json:"senderLocalPart"`
	Subject           *string                                                                      `json:"subject"`
	SenderDisplayName *string                                                                      `json:"senderDisplayName"`
	Body              *string                                                                      `json:"body"`
	BodyFormat        *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum `json:"bodyFormat"`
	ReplyTo           *string                                                                      `json:"replyTo"`
	Customized        *bool                                                                        `json:"customized"`
}

type jsonConfigNotificationSendEmailRevertSecondFactorAdditionTemplate ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate

func (r *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendEmailRevertSecondFactorAdditionTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendEmailRevertSecondFactorAdditionTemplate
	} else {

		r.SenderLocalPart = res.SenderLocalPart

		r.Subject = res.Subject

		r.SenderDisplayName = res.SenderDisplayName

		r.Body = res.Body

		r.BodyFormat = res.BodyFormat

		r.ReplyTo = res.ReplyTo

		r.Customized = res.Customized

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendEmailRevertSecondFactorAdditionTemplate *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate = &ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{empty: true}

func (r *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendSms struct {
	empty           bool                                  `json:"-"`
	UseDeviceLocale *bool                                 `json:"useDeviceLocale"`
	SmsTemplate     *ConfigNotificationSendSmsSmsTemplate `json:"smsTemplate"`
}

type jsonConfigNotificationSendSms ConfigNotificationSendSms

func (r *ConfigNotificationSendSms) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendSms
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendSms
	} else {

		r.UseDeviceLocale = res.UseDeviceLocale

		r.SmsTemplate = res.SmsTemplate

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendSms is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendSms *ConfigNotificationSendSms = &ConfigNotificationSendSms{empty: true}

func (r *ConfigNotificationSendSms) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendSms) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendSms) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigNotificationSendSmsSmsTemplate struct {
	empty   bool    `json:"-"`
	Content *string `json:"content"`
}

type jsonConfigNotificationSendSmsSmsTemplate ConfigNotificationSendSmsSmsTemplate

func (r *ConfigNotificationSendSmsSmsTemplate) UnmarshalJSON(data []byte) error {
	var res jsonConfigNotificationSendSmsSmsTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigNotificationSendSmsSmsTemplate
	} else {

		r.Content = res.Content

	}
	return nil
}

// This object is used to assert a desired state where this ConfigNotificationSendSmsSmsTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigNotificationSendSmsSmsTemplate *ConfigNotificationSendSmsSmsTemplate = &ConfigNotificationSendSmsSmsTemplate{empty: true}

func (r *ConfigNotificationSendSmsSmsTemplate) Empty() bool {
	return r.empty
}

func (r *ConfigNotificationSendSmsSmsTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigNotificationSendSmsSmsTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigQuota struct {
	empty             bool                          `json:"-"`
	SignUpQuotaConfig *ConfigQuotaSignUpQuotaConfig `json:"signUpQuotaConfig"`
}

type jsonConfigQuota ConfigQuota

func (r *ConfigQuota) UnmarshalJSON(data []byte) error {
	var res jsonConfigQuota
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigQuota
	} else {

		r.SignUpQuotaConfig = res.SignUpQuotaConfig

	}
	return nil
}

// This object is used to assert a desired state where this ConfigQuota is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigQuota *ConfigQuota = &ConfigQuota{empty: true}

func (r *ConfigQuota) Empty() bool {
	return r.empty
}

func (r *ConfigQuota) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigQuota) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigQuotaSignUpQuotaConfig struct {
	empty         bool    `json:"-"`
	Quota         *int64  `json:"quota"`
	StartTime     *string `json:"startTime"`
	QuotaDuration *string `json:"quotaDuration"`
}

type jsonConfigQuotaSignUpQuotaConfig ConfigQuotaSignUpQuotaConfig

func (r *ConfigQuotaSignUpQuotaConfig) UnmarshalJSON(data []byte) error {
	var res jsonConfigQuotaSignUpQuotaConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigQuotaSignUpQuotaConfig
	} else {

		r.Quota = res.Quota

		r.StartTime = res.StartTime

		r.QuotaDuration = res.QuotaDuration

	}
	return nil
}

// This object is used to assert a desired state where this ConfigQuotaSignUpQuotaConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigQuotaSignUpQuotaConfig *ConfigQuotaSignUpQuotaConfig = &ConfigQuotaSignUpQuotaConfig{empty: true}

func (r *ConfigQuotaSignUpQuotaConfig) Empty() bool {
	return r.empty
}

func (r *ConfigQuotaSignUpQuotaConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigQuotaSignUpQuotaConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigMonitoring struct {
	empty          bool                            `json:"-"`
	RequestLogging *ConfigMonitoringRequestLogging `json:"requestLogging"`
}

type jsonConfigMonitoring ConfigMonitoring

func (r *ConfigMonitoring) UnmarshalJSON(data []byte) error {
	var res jsonConfigMonitoring
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigMonitoring
	} else {

		r.RequestLogging = res.RequestLogging

	}
	return nil
}

// This object is used to assert a desired state where this ConfigMonitoring is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigMonitoring *ConfigMonitoring = &ConfigMonitoring{empty: true}

func (r *ConfigMonitoring) Empty() bool {
	return r.empty
}

func (r *ConfigMonitoring) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigMonitoring) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigMonitoringRequestLogging struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

type jsonConfigMonitoringRequestLogging ConfigMonitoringRequestLogging

func (r *ConfigMonitoringRequestLogging) UnmarshalJSON(data []byte) error {
	var res jsonConfigMonitoringRequestLogging
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigMonitoringRequestLogging
	} else {

		r.Enabled = res.Enabled

	}
	return nil
}

// This object is used to assert a desired state where this ConfigMonitoringRequestLogging is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigMonitoringRequestLogging *ConfigMonitoringRequestLogging = &ConfigMonitoringRequestLogging{empty: true}

func (r *ConfigMonitoringRequestLogging) Empty() bool {
	return r.empty
}

func (r *ConfigMonitoringRequestLogging) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigMonitoringRequestLogging) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigMultiTenant struct {
	empty                 bool    `json:"-"`
	AllowTenants          *bool   `json:"allowTenants"`
	DefaultTenantLocation *string `json:"defaultTenantLocation"`
}

type jsonConfigMultiTenant ConfigMultiTenant

func (r *ConfigMultiTenant) UnmarshalJSON(data []byte) error {
	var res jsonConfigMultiTenant
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigMultiTenant
	} else {

		r.AllowTenants = res.AllowTenants

		r.DefaultTenantLocation = res.DefaultTenantLocation

	}
	return nil
}

// This object is used to assert a desired state where this ConfigMultiTenant is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigMultiTenant *ConfigMultiTenant = &ConfigMultiTenant{empty: true}

func (r *ConfigMultiTenant) Empty() bool {
	return r.empty
}

func (r *ConfigMultiTenant) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigMultiTenant) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigClient struct {
	empty             bool                     `json:"-"`
	ApiKey            *string                  `json:"apiKey"`
	Permissions       *ConfigClientPermissions `json:"permissions"`
	FirebaseSubdomain *string                  `json:"firebaseSubdomain"`
}

type jsonConfigClient ConfigClient

func (r *ConfigClient) UnmarshalJSON(data []byte) error {
	var res jsonConfigClient
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigClient
	} else {

		r.ApiKey = res.ApiKey

		r.Permissions = res.Permissions

		r.FirebaseSubdomain = res.FirebaseSubdomain

	}
	return nil
}

// This object is used to assert a desired state where this ConfigClient is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigClient *ConfigClient = &ConfigClient{empty: true}

func (r *ConfigClient) Empty() bool {
	return r.empty
}

func (r *ConfigClient) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigClient) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigClientPermissions struct {
	empty                bool  `json:"-"`
	DisabledUserSignup   *bool `json:"disabledUserSignup"`
	DisabledUserDeletion *bool `json:"disabledUserDeletion"`
}

type jsonConfigClientPermissions ConfigClientPermissions

func (r *ConfigClientPermissions) UnmarshalJSON(data []byte) error {
	var res jsonConfigClientPermissions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigClientPermissions
	} else {

		r.DisabledUserSignup = res.DisabledUserSignup

		r.DisabledUserDeletion = res.DisabledUserDeletion

	}
	return nil
}

// This object is used to assert a desired state where this ConfigClientPermissions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigClientPermissions *ConfigClientPermissions = &ConfigClientPermissions{empty: true}

func (r *ConfigClientPermissions) Empty() bool {
	return r.empty
}

func (r *ConfigClientPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigClientPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigMfa struct {
	empty bool                `json:"-"`
	State *ConfigMfaStateEnum `json:"state"`
}

type jsonConfigMfa ConfigMfa

func (r *ConfigMfa) UnmarshalJSON(data []byte) error {
	var res jsonConfigMfa
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigMfa
	} else {

		r.State = res.State

	}
	return nil
}

// This object is used to assert a desired state where this ConfigMfa is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigMfa *ConfigMfa = &ConfigMfa{empty: true}

func (r *ConfigMfa) Empty() bool {
	return r.empty
}

func (r *ConfigMfa) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigMfa) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigBlockingFunctions struct {
	empty    bool                                       `json:"-"`
	Triggers map[string]ConfigBlockingFunctionsTriggers `json:"triggers"`
}

type jsonConfigBlockingFunctions ConfigBlockingFunctions

func (r *ConfigBlockingFunctions) UnmarshalJSON(data []byte) error {
	var res jsonConfigBlockingFunctions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigBlockingFunctions
	} else {

		r.Triggers = res.Triggers

	}
	return nil
}

// This object is used to assert a desired state where this ConfigBlockingFunctions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigBlockingFunctions *ConfigBlockingFunctions = &ConfigBlockingFunctions{empty: true}

func (r *ConfigBlockingFunctions) Empty() bool {
	return r.empty
}

func (r *ConfigBlockingFunctions) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigBlockingFunctions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ConfigBlockingFunctionsTriggers struct {
	empty       bool    `json:"-"`
	FunctionUri *string `json:"functionUri"`
	UpdateTime  *string `json:"updateTime"`
}

type jsonConfigBlockingFunctionsTriggers ConfigBlockingFunctionsTriggers

func (r *ConfigBlockingFunctionsTriggers) UnmarshalJSON(data []byte) error {
	var res jsonConfigBlockingFunctionsTriggers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyConfigBlockingFunctionsTriggers
	} else {

		r.FunctionUri = res.FunctionUri

		r.UpdateTime = res.UpdateTime

	}
	return nil
}

// This object is used to assert a desired state where this ConfigBlockingFunctionsTriggers is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyConfigBlockingFunctionsTriggers *ConfigBlockingFunctionsTriggers = &ConfigBlockingFunctionsTriggers{empty: true}

func (r *ConfigBlockingFunctionsTriggers) Empty() bool {
	return r.empty
}

func (r *ConfigBlockingFunctionsTriggers) String() string {
	return dcl.SprintResource(r)
}

func (r *ConfigBlockingFunctionsTriggers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Config) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "identity_toolkit",
		Type:    "Config",
		Version: "alpha",
	}
}

func (r *Config) ID() (string, error) {
	if err := extractConfigFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"sign_in":            dcl.ValueOrEmptyString(nr.SignIn),
		"notification":       dcl.ValueOrEmptyString(nr.Notification),
		"quota":              dcl.ValueOrEmptyString(nr.Quota),
		"monitoring":         dcl.ValueOrEmptyString(nr.Monitoring),
		"multi_tenant":       dcl.ValueOrEmptyString(nr.MultiTenant),
		"authorized_domains": dcl.ValueOrEmptyString(nr.AuthorizedDomains),
		"subtype":            dcl.ValueOrEmptyString(nr.Subtype),
		"client":             dcl.ValueOrEmptyString(nr.Client),
		"mfa":                dcl.ValueOrEmptyString(nr.Mfa),
		"blocking_functions": dcl.ValueOrEmptyString(nr.BlockingFunctions),
		"project":            dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.Nprintf("projects/{{project}}/config", params), nil
}

const ConfigMaxPage = -1

type ConfigList struct {
	Items []*Config

	nextToken string

	resource *Config
}

func (c *Client) GetConfig(ctx context.Context, r *Config) (*Config, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractConfigFields(r)

	b, err := c.getConfigRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalConfig(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeConfigNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractConfigFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) ApplyConfig(ctx context.Context, rawDesired *Config, opts ...dcl.ApplyOption) (*Config, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Config
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyConfigHelper(c, ctx, rawDesired, opts...)
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

func applyConfigHelper(c *Client, ctx context.Context, rawDesired *Config, opts ...dcl.ApplyOption) (*Config, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyConfig...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractConfigFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.configDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToConfigDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		return nil, dcl.ApplyInfeasibleError{Message: "No initial state found for singleton resource."}
	} else {
		for _, d := range diffs {
			if d.UpdateOp == nil {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) no update method found for field", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}
	var ops []configApiOperation
	for _, d := range diffs {
		ops = append(ops, d.UpdateOp)
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
	return applyConfigDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyConfigDiff(c *Client, ctx context.Context, desired *Config, rawDesired *Config, ops []configApiOperation, opts ...dcl.ApplyOption) (*Config, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetConfig(ctx, desired)
	if err != nil {
		return nil, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeConfigNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeConfigDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractConfigFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractConfigFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffConfig(c, newDesired, newState)
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
