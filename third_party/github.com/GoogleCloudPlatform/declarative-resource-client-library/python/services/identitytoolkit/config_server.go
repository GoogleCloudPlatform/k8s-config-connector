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
package server

import (
	"context"
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	identitytoolkitpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/identitytoolkit_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit"
)

// ConfigServer implements the gRPC interface for Config.
type ConfigServer struct{}

// ProtoToConfigSignInEmailHashConfigAlgorithmEnum converts a ConfigSignInEmailHashConfigAlgorithmEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum(e identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum) *identitytoolkit.ConfigSignInEmailHashConfigAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigSignInEmailHashConfigAlgorithmEnum(n[len("IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSignInHashConfigAlgorithmEnum converts a ConfigSignInHashConfigAlgorithmEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigSignInHashConfigAlgorithmEnum(e identitytoolkitpb.IdentitytoolkitConfigSignInHashConfigAlgorithmEnum) *identitytoolkit.ConfigSignInHashConfigAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigSignInHashConfigAlgorithmEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigSignInHashConfigAlgorithmEnum(n[len("IdentitytoolkitConfigSignInHashConfigAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailMethodEnum converts a ConfigNotificationSendEmailMethodEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailMethodEnum(e identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailMethodEnum) *identitytoolkit.ConfigNotificationSendEmailMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailMethodEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigNotificationSendEmailMethodEnum(n[len("IdentitytoolkitConfigNotificationSendEmailMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailSmtpSecurityModeEnum converts a ConfigNotificationSendEmailSmtpSecurityModeEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum(e identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum) *identitytoolkit.ConfigNotificationSendEmailSmtpSecurityModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigNotificationSendEmailSmtpSecurityModeEnum(n[len("IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum converts a ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(e identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum) *identitytoolkit.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(n[len("IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum converts a ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(e identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum) *identitytoolkit.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(n[len("IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum converts a ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(e identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum) *identitytoolkit.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(n[len("IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailDnsInfoCustomDomainStateEnum converts a ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(e identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum) *identitytoolkit.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(n[len("IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(e identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum) *identitytoolkit.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(n[len("IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSubtypeEnum converts a ConfigSubtypeEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigSubtypeEnum(e identitytoolkitpb.IdentitytoolkitConfigSubtypeEnum) *identitytoolkit.ConfigSubtypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigSubtypeEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigSubtypeEnum(n[len("IdentitytoolkitConfigSubtypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigMfaStateEnum converts a ConfigMfaStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitConfigMfaStateEnum(e identitytoolkitpb.IdentitytoolkitConfigMfaStateEnum) *identitytoolkit.ConfigMfaStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := identitytoolkitpb.IdentitytoolkitConfigMfaStateEnum_name[int32(e)]; ok {
		e := identitytoolkit.ConfigMfaStateEnum(n[len("IdentitytoolkitConfigMfaStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSignIn converts a ConfigSignIn object from its proto representation.
func ProtoToIdentitytoolkitConfigSignIn(p *identitytoolkitpb.IdentitytoolkitConfigSignIn) *identitytoolkit.ConfigSignIn {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigSignIn{
		Email:                ProtoToIdentitytoolkitConfigSignInEmail(p.GetEmail()),
		PhoneNumber:          ProtoToIdentitytoolkitConfigSignInPhoneNumber(p.GetPhoneNumber()),
		Anonymous:            ProtoToIdentitytoolkitConfigSignInAnonymous(p.GetAnonymous()),
		AllowDuplicateEmails: dcl.Bool(p.GetAllowDuplicateEmails()),
		HashConfig:           ProtoToIdentitytoolkitConfigSignInHashConfig(p.GetHashConfig()),
	}
	return obj
}

// ProtoToConfigSignInEmail converts a ConfigSignInEmail object from its proto representation.
func ProtoToIdentitytoolkitConfigSignInEmail(p *identitytoolkitpb.IdentitytoolkitConfigSignInEmail) *identitytoolkit.ConfigSignInEmail {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigSignInEmail{
		Enabled:          dcl.Bool(p.GetEnabled()),
		PasswordRequired: dcl.Bool(p.GetPasswordRequired()),
		HashConfig:       ProtoToIdentitytoolkitConfigSignInEmailHashConfig(p.GetHashConfig()),
	}
	return obj
}

// ProtoToConfigSignInEmailHashConfig converts a ConfigSignInEmailHashConfig object from its proto representation.
func ProtoToIdentitytoolkitConfigSignInEmailHashConfig(p *identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfig) *identitytoolkit.ConfigSignInEmailHashConfig {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigSignInEmailHashConfig{
		Algorithm:     ProtoToIdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum(p.GetAlgorithm()),
		SignerKey:     dcl.StringOrNil(p.GetSignerKey()),
		SaltSeparator: dcl.StringOrNil(p.GetSaltSeparator()),
		Rounds:        dcl.Int64OrNil(p.GetRounds()),
		MemoryCost:    dcl.Int64OrNil(p.GetMemoryCost()),
	}
	return obj
}

// ProtoToConfigSignInPhoneNumber converts a ConfigSignInPhoneNumber object from its proto representation.
func ProtoToIdentitytoolkitConfigSignInPhoneNumber(p *identitytoolkitpb.IdentitytoolkitConfigSignInPhoneNumber) *identitytoolkit.ConfigSignInPhoneNumber {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigSignInPhoneNumber{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigSignInAnonymous converts a ConfigSignInAnonymous object from its proto representation.
func ProtoToIdentitytoolkitConfigSignInAnonymous(p *identitytoolkitpb.IdentitytoolkitConfigSignInAnonymous) *identitytoolkit.ConfigSignInAnonymous {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigSignInAnonymous{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigSignInHashConfig converts a ConfigSignInHashConfig object from its proto representation.
func ProtoToIdentitytoolkitConfigSignInHashConfig(p *identitytoolkitpb.IdentitytoolkitConfigSignInHashConfig) *identitytoolkit.ConfigSignInHashConfig {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigSignInHashConfig{
		Algorithm:     ProtoToIdentitytoolkitConfigSignInHashConfigAlgorithmEnum(p.GetAlgorithm()),
		SignerKey:     dcl.StringOrNil(p.GetSignerKey()),
		SaltSeparator: dcl.StringOrNil(p.GetSaltSeparator()),
		Rounds:        dcl.Int64OrNil(p.GetRounds()),
		MemoryCost:    dcl.Int64OrNil(p.GetMemoryCost()),
	}
	return obj
}

// ProtoToConfigNotification converts a ConfigNotification object from its proto representation.
func ProtoToIdentitytoolkitConfigNotification(p *identitytoolkitpb.IdentitytoolkitConfigNotification) *identitytoolkit.ConfigNotification {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotification{
		SendEmail:     ProtoToIdentitytoolkitConfigNotificationSendEmail(p.GetSendEmail()),
		SendSms:       ProtoToIdentitytoolkitConfigNotificationSendSms(p.GetSendSms()),
		DefaultLocale: dcl.StringOrNil(p.GetDefaultLocale()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmail converts a ConfigNotificationSendEmail object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmail(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmail) *identitytoolkit.ConfigNotificationSendEmail {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendEmail{
		Method:                             ProtoToIdentitytoolkitConfigNotificationSendEmailMethodEnum(p.GetMethod()),
		Smtp:                               ProtoToIdentitytoolkitConfigNotificationSendEmailSmtp(p.GetSmtp()),
		ResetPasswordTemplate:              ProtoToIdentitytoolkitConfigNotificationSendEmailResetPasswordTemplate(p.GetResetPasswordTemplate()),
		VerifyEmailTemplate:                ProtoToIdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplate(p.GetVerifyEmailTemplate()),
		ChangeEmailTemplate:                ProtoToIdentitytoolkitConfigNotificationSendEmailChangeEmailTemplate(p.GetChangeEmailTemplate()),
		CallbackUri:                        dcl.StringOrNil(p.GetCallbackUri()),
		DnsInfo:                            ProtoToIdentitytoolkitConfigNotificationSendEmailDnsInfo(p.GetDnsInfo()),
		RevertSecondFactorAdditionTemplate: ProtoToIdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(p.GetRevertSecondFactorAdditionTemplate()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailSmtp converts a ConfigNotificationSendEmailSmtp object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailSmtp(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtp) *identitytoolkit.ConfigNotificationSendEmailSmtp {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendEmailSmtp{
		SenderEmail:  dcl.StringOrNil(p.GetSenderEmail()),
		Host:         dcl.StringOrNil(p.GetHost()),
		Port:         dcl.Int64OrNil(p.GetPort()),
		Username:     dcl.StringOrNil(p.GetUsername()),
		Password:     dcl.StringOrNil(p.GetPassword()),
		SecurityMode: ProtoToIdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum(p.GetSecurityMode()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailResetPasswordTemplate converts a ConfigNotificationSendEmailResetPasswordTemplate object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailResetPasswordTemplate(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplate) *identitytoolkit.ConfigNotificationSendEmailResetPasswordTemplate {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendEmailResetPasswordTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailVerifyEmailTemplate converts a ConfigNotificationSendEmailVerifyEmailTemplate object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplate(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplate) *identitytoolkit.ConfigNotificationSendEmailVerifyEmailTemplate {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendEmailVerifyEmailTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailChangeEmailTemplate converts a ConfigNotificationSendEmailChangeEmailTemplate object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailChangeEmailTemplate(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplate) *identitytoolkit.ConfigNotificationSendEmailChangeEmailTemplate {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendEmailChangeEmailTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailDnsInfo converts a ConfigNotificationSendEmailDnsInfo object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailDnsInfo(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfo) *identitytoolkit.ConfigNotificationSendEmailDnsInfo {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendEmailDnsInfo{
		CustomDomain:                  dcl.StringOrNil(p.GetCustomDomain()),
		UseCustomDomain:               dcl.Bool(p.GetUseCustomDomain()),
		PendingCustomDomain:           dcl.StringOrNil(p.GetPendingCustomDomain()),
		CustomDomainState:             ProtoToIdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(p.GetCustomDomainState()),
		DomainVerificationRequestTime: dcl.StringOrNil(p.GetDomainVerificationRequestTime()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailRevertSecondFactorAdditionTemplate converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) *identitytoolkit.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendSms converts a ConfigNotificationSendSms object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendSms(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendSms) *identitytoolkit.ConfigNotificationSendSms {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendSms{
		UseDeviceLocale: dcl.Bool(p.GetUseDeviceLocale()),
		SmsTemplate:     ProtoToIdentitytoolkitConfigNotificationSendSmsSmsTemplate(p.GetSmsTemplate()),
	}
	return obj
}

// ProtoToConfigNotificationSendSmsSmsTemplate converts a ConfigNotificationSendSmsSmsTemplate object from its proto representation.
func ProtoToIdentitytoolkitConfigNotificationSendSmsSmsTemplate(p *identitytoolkitpb.IdentitytoolkitConfigNotificationSendSmsSmsTemplate) *identitytoolkit.ConfigNotificationSendSmsSmsTemplate {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigNotificationSendSmsSmsTemplate{
		Content: dcl.StringOrNil(p.GetContent()),
	}
	return obj
}

// ProtoToConfigQuota converts a ConfigQuota object from its proto representation.
func ProtoToIdentitytoolkitConfigQuota(p *identitytoolkitpb.IdentitytoolkitConfigQuota) *identitytoolkit.ConfigQuota {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigQuota{
		SignUpQuotaConfig: ProtoToIdentitytoolkitConfigQuotaSignUpQuotaConfig(p.GetSignUpQuotaConfig()),
	}
	return obj
}

// ProtoToConfigQuotaSignUpQuotaConfig converts a ConfigQuotaSignUpQuotaConfig object from its proto representation.
func ProtoToIdentitytoolkitConfigQuotaSignUpQuotaConfig(p *identitytoolkitpb.IdentitytoolkitConfigQuotaSignUpQuotaConfig) *identitytoolkit.ConfigQuotaSignUpQuotaConfig {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigQuotaSignUpQuotaConfig{
		Quota:         dcl.Int64OrNil(p.GetQuota()),
		StartTime:     dcl.StringOrNil(p.GetStartTime()),
		QuotaDuration: dcl.StringOrNil(p.GetQuotaDuration()),
	}
	return obj
}

// ProtoToConfigMonitoring converts a ConfigMonitoring object from its proto representation.
func ProtoToIdentitytoolkitConfigMonitoring(p *identitytoolkitpb.IdentitytoolkitConfigMonitoring) *identitytoolkit.ConfigMonitoring {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigMonitoring{
		RequestLogging: ProtoToIdentitytoolkitConfigMonitoringRequestLogging(p.GetRequestLogging()),
	}
	return obj
}

// ProtoToConfigMonitoringRequestLogging converts a ConfigMonitoringRequestLogging object from its proto representation.
func ProtoToIdentitytoolkitConfigMonitoringRequestLogging(p *identitytoolkitpb.IdentitytoolkitConfigMonitoringRequestLogging) *identitytoolkit.ConfigMonitoringRequestLogging {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigMonitoringRequestLogging{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigMultiTenant converts a ConfigMultiTenant object from its proto representation.
func ProtoToIdentitytoolkitConfigMultiTenant(p *identitytoolkitpb.IdentitytoolkitConfigMultiTenant) *identitytoolkit.ConfigMultiTenant {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigMultiTenant{
		AllowTenants:          dcl.Bool(p.GetAllowTenants()),
		DefaultTenantLocation: dcl.StringOrNil(p.GetDefaultTenantLocation()),
	}
	return obj
}

// ProtoToConfigClient converts a ConfigClient object from its proto representation.
func ProtoToIdentitytoolkitConfigClient(p *identitytoolkitpb.IdentitytoolkitConfigClient) *identitytoolkit.ConfigClient {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigClient{
		ApiKey:            dcl.StringOrNil(p.GetApiKey()),
		Permissions:       ProtoToIdentitytoolkitConfigClientPermissions(p.GetPermissions()),
		FirebaseSubdomain: dcl.StringOrNil(p.GetFirebaseSubdomain()),
	}
	return obj
}

// ProtoToConfigClientPermissions converts a ConfigClientPermissions object from its proto representation.
func ProtoToIdentitytoolkitConfigClientPermissions(p *identitytoolkitpb.IdentitytoolkitConfigClientPermissions) *identitytoolkit.ConfigClientPermissions {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigClientPermissions{
		DisabledUserSignup:   dcl.Bool(p.GetDisabledUserSignup()),
		DisabledUserDeletion: dcl.Bool(p.GetDisabledUserDeletion()),
	}
	return obj
}

// ProtoToConfigMfa converts a ConfigMfa object from its proto representation.
func ProtoToIdentitytoolkitConfigMfa(p *identitytoolkitpb.IdentitytoolkitConfigMfa) *identitytoolkit.ConfigMfa {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigMfa{
		State: ProtoToIdentitytoolkitConfigMfaStateEnum(p.GetState()),
	}
	return obj
}

// ProtoToConfigBlockingFunctions converts a ConfigBlockingFunctions object from its proto representation.
func ProtoToIdentitytoolkitConfigBlockingFunctions(p *identitytoolkitpb.IdentitytoolkitConfigBlockingFunctions) *identitytoolkit.ConfigBlockingFunctions {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigBlockingFunctions{}
	return obj
}

// ProtoToConfigBlockingFunctionsTriggers converts a ConfigBlockingFunctionsTriggers object from its proto representation.
func ProtoToIdentitytoolkitConfigBlockingFunctionsTriggers(p *identitytoolkitpb.IdentitytoolkitConfigBlockingFunctionsTriggers) *identitytoolkit.ConfigBlockingFunctionsTriggers {
	if p == nil {
		return nil
	}
	obj := &identitytoolkit.ConfigBlockingFunctionsTriggers{
		FunctionUri: dcl.StringOrNil(p.GetFunctionUri()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToConfig converts a Config resource from its proto representation.
func ProtoToConfig(p *identitytoolkitpb.IdentitytoolkitConfig) *identitytoolkit.Config {
	obj := &identitytoolkit.Config{
		SignIn:            ProtoToIdentitytoolkitConfigSignIn(p.GetSignIn()),
		Notification:      ProtoToIdentitytoolkitConfigNotification(p.GetNotification()),
		Quota:             ProtoToIdentitytoolkitConfigQuota(p.GetQuota()),
		Monitoring:        ProtoToIdentitytoolkitConfigMonitoring(p.GetMonitoring()),
		MultiTenant:       ProtoToIdentitytoolkitConfigMultiTenant(p.GetMultiTenant()),
		Subtype:           ProtoToIdentitytoolkitConfigSubtypeEnum(p.GetSubtype()),
		Client:            ProtoToIdentitytoolkitConfigClient(p.GetClient()),
		Mfa:               ProtoToIdentitytoolkitConfigMfa(p.GetMfa()),
		BlockingFunctions: ProtoToIdentitytoolkitConfigBlockingFunctions(p.GetBlockingFunctions()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetAuthorizedDomains() {
		obj.AuthorizedDomains = append(obj.AuthorizedDomains, r)
	}
	return obj
}

// ConfigSignInEmailHashConfigAlgorithmEnumToProto converts a ConfigSignInEmailHashConfigAlgorithmEnum enum to its proto representation.
func IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnumToProto(e *identitytoolkit.ConfigSignInEmailHashConfigAlgorithmEnum) identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum_value["ConfigSignInEmailHashConfigAlgorithmEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnum(0)
}

// ConfigSignInHashConfigAlgorithmEnumToProto converts a ConfigSignInHashConfigAlgorithmEnum enum to its proto representation.
func IdentitytoolkitConfigSignInHashConfigAlgorithmEnumToProto(e *identitytoolkit.ConfigSignInHashConfigAlgorithmEnum) identitytoolkitpb.IdentitytoolkitConfigSignInHashConfigAlgorithmEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigSignInHashConfigAlgorithmEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigSignInHashConfigAlgorithmEnum_value["ConfigSignInHashConfigAlgorithmEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigSignInHashConfigAlgorithmEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigSignInHashConfigAlgorithmEnum(0)
}

// ConfigNotificationSendEmailMethodEnumToProto converts a ConfigNotificationSendEmailMethodEnum enum to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailMethodEnumToProto(e *identitytoolkit.ConfigNotificationSendEmailMethodEnum) identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailMethodEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailMethodEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailMethodEnum_value["ConfigNotificationSendEmailMethodEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailMethodEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailMethodEnum(0)
}

// ConfigNotificationSendEmailSmtpSecurityModeEnumToProto converts a ConfigNotificationSendEmailSmtpSecurityModeEnum enum to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnumToProto(e *identitytoolkit.ConfigNotificationSendEmailSmtpSecurityModeEnum) identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum_value["ConfigNotificationSendEmailSmtpSecurityModeEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnum(0)
}

// ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto(e *identitytoolkit.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum) identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum_value["ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto(e *identitytoolkit.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum) identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum_value["ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto(e *identitytoolkit.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum) identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum_value["ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto converts a ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum enum to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto(e *identitytoolkit.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum) identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum_value["ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(0)
}

// ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto(e *identitytoolkit.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum) identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum_value["ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(0)
}

// ConfigSubtypeEnumToProto converts a ConfigSubtypeEnum enum to its proto representation.
func IdentitytoolkitConfigSubtypeEnumToProto(e *identitytoolkit.ConfigSubtypeEnum) identitytoolkitpb.IdentitytoolkitConfigSubtypeEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigSubtypeEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigSubtypeEnum_value["ConfigSubtypeEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigSubtypeEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigSubtypeEnum(0)
}

// ConfigMfaStateEnumToProto converts a ConfigMfaStateEnum enum to its proto representation.
func IdentitytoolkitConfigMfaStateEnumToProto(e *identitytoolkit.ConfigMfaStateEnum) identitytoolkitpb.IdentitytoolkitConfigMfaStateEnum {
	if e == nil {
		return identitytoolkitpb.IdentitytoolkitConfigMfaStateEnum(0)
	}
	if v, ok := identitytoolkitpb.IdentitytoolkitConfigMfaStateEnum_value["ConfigMfaStateEnum"+string(*e)]; ok {
		return identitytoolkitpb.IdentitytoolkitConfigMfaStateEnum(v)
	}
	return identitytoolkitpb.IdentitytoolkitConfigMfaStateEnum(0)
}

// ConfigSignInToProto converts a ConfigSignIn object to its proto representation.
func IdentitytoolkitConfigSignInToProto(o *identitytoolkit.ConfigSignIn) *identitytoolkitpb.IdentitytoolkitConfigSignIn {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigSignIn{}
	p.SetEmail(IdentitytoolkitConfigSignInEmailToProto(o.Email))
	p.SetPhoneNumber(IdentitytoolkitConfigSignInPhoneNumberToProto(o.PhoneNumber))
	p.SetAnonymous(IdentitytoolkitConfigSignInAnonymousToProto(o.Anonymous))
	p.SetAllowDuplicateEmails(dcl.ValueOrEmptyBool(o.AllowDuplicateEmails))
	p.SetHashConfig(IdentitytoolkitConfigSignInHashConfigToProto(o.HashConfig))
	return p
}

// ConfigSignInEmailToProto converts a ConfigSignInEmail object to its proto representation.
func IdentitytoolkitConfigSignInEmailToProto(o *identitytoolkit.ConfigSignInEmail) *identitytoolkitpb.IdentitytoolkitConfigSignInEmail {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigSignInEmail{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetPasswordRequired(dcl.ValueOrEmptyBool(o.PasswordRequired))
	p.SetHashConfig(IdentitytoolkitConfigSignInEmailHashConfigToProto(o.HashConfig))
	return p
}

// ConfigSignInEmailHashConfigToProto converts a ConfigSignInEmailHashConfig object to its proto representation.
func IdentitytoolkitConfigSignInEmailHashConfigToProto(o *identitytoolkit.ConfigSignInEmailHashConfig) *identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfig {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigSignInEmailHashConfig{}
	p.SetAlgorithm(IdentitytoolkitConfigSignInEmailHashConfigAlgorithmEnumToProto(o.Algorithm))
	p.SetSignerKey(dcl.ValueOrEmptyString(o.SignerKey))
	p.SetSaltSeparator(dcl.ValueOrEmptyString(o.SaltSeparator))
	p.SetRounds(dcl.ValueOrEmptyInt64(o.Rounds))
	p.SetMemoryCost(dcl.ValueOrEmptyInt64(o.MemoryCost))
	return p
}

// ConfigSignInPhoneNumberToProto converts a ConfigSignInPhoneNumber object to its proto representation.
func IdentitytoolkitConfigSignInPhoneNumberToProto(o *identitytoolkit.ConfigSignInPhoneNumber) *identitytoolkitpb.IdentitytoolkitConfigSignInPhoneNumber {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigSignInPhoneNumber{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	mTestPhoneNumbers := make(map[string]string, len(o.TestPhoneNumbers))
	for k, r := range o.TestPhoneNumbers {
		mTestPhoneNumbers[k] = r
	}
	p.SetTestPhoneNumbers(mTestPhoneNumbers)
	return p
}

// ConfigSignInAnonymousToProto converts a ConfigSignInAnonymous object to its proto representation.
func IdentitytoolkitConfigSignInAnonymousToProto(o *identitytoolkit.ConfigSignInAnonymous) *identitytoolkitpb.IdentitytoolkitConfigSignInAnonymous {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigSignInAnonymous{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ConfigSignInHashConfigToProto converts a ConfigSignInHashConfig object to its proto representation.
func IdentitytoolkitConfigSignInHashConfigToProto(o *identitytoolkit.ConfigSignInHashConfig) *identitytoolkitpb.IdentitytoolkitConfigSignInHashConfig {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigSignInHashConfig{}
	p.SetAlgorithm(IdentitytoolkitConfigSignInHashConfigAlgorithmEnumToProto(o.Algorithm))
	p.SetSignerKey(dcl.ValueOrEmptyString(o.SignerKey))
	p.SetSaltSeparator(dcl.ValueOrEmptyString(o.SaltSeparator))
	p.SetRounds(dcl.ValueOrEmptyInt64(o.Rounds))
	p.SetMemoryCost(dcl.ValueOrEmptyInt64(o.MemoryCost))
	return p
}

// ConfigNotificationToProto converts a ConfigNotification object to its proto representation.
func IdentitytoolkitConfigNotificationToProto(o *identitytoolkit.ConfigNotification) *identitytoolkitpb.IdentitytoolkitConfigNotification {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotification{}
	p.SetSendEmail(IdentitytoolkitConfigNotificationSendEmailToProto(o.SendEmail))
	p.SetSendSms(IdentitytoolkitConfigNotificationSendSmsToProto(o.SendSms))
	p.SetDefaultLocale(dcl.ValueOrEmptyString(o.DefaultLocale))
	return p
}

// ConfigNotificationSendEmailToProto converts a ConfigNotificationSendEmail object to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailToProto(o *identitytoolkit.ConfigNotificationSendEmail) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmail {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmail{}
	p.SetMethod(IdentitytoolkitConfigNotificationSendEmailMethodEnumToProto(o.Method))
	p.SetSmtp(IdentitytoolkitConfigNotificationSendEmailSmtpToProto(o.Smtp))
	p.SetResetPasswordTemplate(IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateToProto(o.ResetPasswordTemplate))
	p.SetVerifyEmailTemplate(IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateToProto(o.VerifyEmailTemplate))
	p.SetChangeEmailTemplate(IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateToProto(o.ChangeEmailTemplate))
	p.SetCallbackUri(dcl.ValueOrEmptyString(o.CallbackUri))
	p.SetDnsInfo(IdentitytoolkitConfigNotificationSendEmailDnsInfoToProto(o.DnsInfo))
	p.SetRevertSecondFactorAdditionTemplate(IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto(o.RevertSecondFactorAdditionTemplate))
	return p
}

// ConfigNotificationSendEmailSmtpToProto converts a ConfigNotificationSendEmailSmtp object to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailSmtpToProto(o *identitytoolkit.ConfigNotificationSendEmailSmtp) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtp {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailSmtp{}
	p.SetSenderEmail(dcl.ValueOrEmptyString(o.SenderEmail))
	p.SetHost(dcl.ValueOrEmptyString(o.Host))
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	p.SetPassword(dcl.ValueOrEmptyString(o.Password))
	p.SetSecurityMode(IdentitytoolkitConfigNotificationSendEmailSmtpSecurityModeEnumToProto(o.SecurityMode))
	return p
}

// ConfigNotificationSendEmailResetPasswordTemplateToProto converts a ConfigNotificationSendEmailResetPasswordTemplate object to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateToProto(o *identitytoolkit.ConfigNotificationSendEmailResetPasswordTemplate) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplate {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailVerifyEmailTemplateToProto converts a ConfigNotificationSendEmailVerifyEmailTemplate object to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateToProto(o *identitytoolkit.ConfigNotificationSendEmailVerifyEmailTemplate) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplate {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailChangeEmailTemplateToProto converts a ConfigNotificationSendEmailChangeEmailTemplate object to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateToProto(o *identitytoolkit.ConfigNotificationSendEmailChangeEmailTemplate) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplate {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailDnsInfoToProto converts a ConfigNotificationSendEmailDnsInfo object to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailDnsInfoToProto(o *identitytoolkit.ConfigNotificationSendEmailDnsInfo) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfo {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailDnsInfo{}
	p.SetCustomDomain(dcl.ValueOrEmptyString(o.CustomDomain))
	p.SetUseCustomDomain(dcl.ValueOrEmptyBool(o.UseCustomDomain))
	p.SetPendingCustomDomain(dcl.ValueOrEmptyString(o.PendingCustomDomain))
	p.SetCustomDomainState(IdentitytoolkitConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto(o.CustomDomainState))
	p.SetDomainVerificationRequestTime(dcl.ValueOrEmptyString(o.DomainVerificationRequestTime))
	return p
}

// ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate object to its proto representation.
func IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto(o *identitytoolkit.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendSmsToProto converts a ConfigNotificationSendSms object to its proto representation.
func IdentitytoolkitConfigNotificationSendSmsToProto(o *identitytoolkit.ConfigNotificationSendSms) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendSms {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendSms{}
	p.SetUseDeviceLocale(dcl.ValueOrEmptyBool(o.UseDeviceLocale))
	p.SetSmsTemplate(IdentitytoolkitConfigNotificationSendSmsSmsTemplateToProto(o.SmsTemplate))
	return p
}

// ConfigNotificationSendSmsSmsTemplateToProto converts a ConfigNotificationSendSmsSmsTemplate object to its proto representation.
func IdentitytoolkitConfigNotificationSendSmsSmsTemplateToProto(o *identitytoolkit.ConfigNotificationSendSmsSmsTemplate) *identitytoolkitpb.IdentitytoolkitConfigNotificationSendSmsSmsTemplate {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigNotificationSendSmsSmsTemplate{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	return p
}

// ConfigQuotaToProto converts a ConfigQuota object to its proto representation.
func IdentitytoolkitConfigQuotaToProto(o *identitytoolkit.ConfigQuota) *identitytoolkitpb.IdentitytoolkitConfigQuota {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigQuota{}
	p.SetSignUpQuotaConfig(IdentitytoolkitConfigQuotaSignUpQuotaConfigToProto(o.SignUpQuotaConfig))
	return p
}

// ConfigQuotaSignUpQuotaConfigToProto converts a ConfigQuotaSignUpQuotaConfig object to its proto representation.
func IdentitytoolkitConfigQuotaSignUpQuotaConfigToProto(o *identitytoolkit.ConfigQuotaSignUpQuotaConfig) *identitytoolkitpb.IdentitytoolkitConfigQuotaSignUpQuotaConfig {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigQuotaSignUpQuotaConfig{}
	p.SetQuota(dcl.ValueOrEmptyInt64(o.Quota))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetQuotaDuration(dcl.ValueOrEmptyString(o.QuotaDuration))
	return p
}

// ConfigMonitoringToProto converts a ConfigMonitoring object to its proto representation.
func IdentitytoolkitConfigMonitoringToProto(o *identitytoolkit.ConfigMonitoring) *identitytoolkitpb.IdentitytoolkitConfigMonitoring {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigMonitoring{}
	p.SetRequestLogging(IdentitytoolkitConfigMonitoringRequestLoggingToProto(o.RequestLogging))
	return p
}

// ConfigMonitoringRequestLoggingToProto converts a ConfigMonitoringRequestLogging object to its proto representation.
func IdentitytoolkitConfigMonitoringRequestLoggingToProto(o *identitytoolkit.ConfigMonitoringRequestLogging) *identitytoolkitpb.IdentitytoolkitConfigMonitoringRequestLogging {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigMonitoringRequestLogging{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ConfigMultiTenantToProto converts a ConfigMultiTenant object to its proto representation.
func IdentitytoolkitConfigMultiTenantToProto(o *identitytoolkit.ConfigMultiTenant) *identitytoolkitpb.IdentitytoolkitConfigMultiTenant {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigMultiTenant{}
	p.SetAllowTenants(dcl.ValueOrEmptyBool(o.AllowTenants))
	p.SetDefaultTenantLocation(dcl.ValueOrEmptyString(o.DefaultTenantLocation))
	return p
}

// ConfigClientToProto converts a ConfigClient object to its proto representation.
func IdentitytoolkitConfigClientToProto(o *identitytoolkit.ConfigClient) *identitytoolkitpb.IdentitytoolkitConfigClient {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigClient{}
	p.SetApiKey(dcl.ValueOrEmptyString(o.ApiKey))
	p.SetPermissions(IdentitytoolkitConfigClientPermissionsToProto(o.Permissions))
	p.SetFirebaseSubdomain(dcl.ValueOrEmptyString(o.FirebaseSubdomain))
	return p
}

// ConfigClientPermissionsToProto converts a ConfigClientPermissions object to its proto representation.
func IdentitytoolkitConfigClientPermissionsToProto(o *identitytoolkit.ConfigClientPermissions) *identitytoolkitpb.IdentitytoolkitConfigClientPermissions {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigClientPermissions{}
	p.SetDisabledUserSignup(dcl.ValueOrEmptyBool(o.DisabledUserSignup))
	p.SetDisabledUserDeletion(dcl.ValueOrEmptyBool(o.DisabledUserDeletion))
	return p
}

// ConfigMfaToProto converts a ConfigMfa object to its proto representation.
func IdentitytoolkitConfigMfaToProto(o *identitytoolkit.ConfigMfa) *identitytoolkitpb.IdentitytoolkitConfigMfa {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigMfa{}
	p.SetState(IdentitytoolkitConfigMfaStateEnumToProto(o.State))
	return p
}

// ConfigBlockingFunctionsToProto converts a ConfigBlockingFunctions object to its proto representation.
func IdentitytoolkitConfigBlockingFunctionsToProto(o *identitytoolkit.ConfigBlockingFunctions) *identitytoolkitpb.IdentitytoolkitConfigBlockingFunctions {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigBlockingFunctions{}
	mTriggers := make(map[string]*identitytoolkitpb.IdentitytoolkitConfigBlockingFunctionsTriggers, len(o.Triggers))
	for k, r := range o.Triggers {
		mTriggers[k] = IdentitytoolkitConfigBlockingFunctionsTriggersToProto(&r)
	}
	p.SetTriggers(mTriggers)
	return p
}

// ConfigBlockingFunctionsTriggersToProto converts a ConfigBlockingFunctionsTriggers object to its proto representation.
func IdentitytoolkitConfigBlockingFunctionsTriggersToProto(o *identitytoolkit.ConfigBlockingFunctionsTriggers) *identitytoolkitpb.IdentitytoolkitConfigBlockingFunctionsTriggers {
	if o == nil {
		return nil
	}
	p := &identitytoolkitpb.IdentitytoolkitConfigBlockingFunctionsTriggers{}
	p.SetFunctionUri(dcl.ValueOrEmptyString(o.FunctionUri))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// ConfigToProto converts a Config resource to its proto representation.
func ConfigToProto(resource *identitytoolkit.Config) *identitytoolkitpb.IdentitytoolkitConfig {
	p := &identitytoolkitpb.IdentitytoolkitConfig{}
	p.SetSignIn(IdentitytoolkitConfigSignInToProto(resource.SignIn))
	p.SetNotification(IdentitytoolkitConfigNotificationToProto(resource.Notification))
	p.SetQuota(IdentitytoolkitConfigQuotaToProto(resource.Quota))
	p.SetMonitoring(IdentitytoolkitConfigMonitoringToProto(resource.Monitoring))
	p.SetMultiTenant(IdentitytoolkitConfigMultiTenantToProto(resource.MultiTenant))
	p.SetSubtype(IdentitytoolkitConfigSubtypeEnumToProto(resource.Subtype))
	p.SetClient(IdentitytoolkitConfigClientToProto(resource.Client))
	p.SetMfa(IdentitytoolkitConfigMfaToProto(resource.Mfa))
	p.SetBlockingFunctions(IdentitytoolkitConfigBlockingFunctionsToProto(resource.BlockingFunctions))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sAuthorizedDomains := make([]string, len(resource.AuthorizedDomains))
	for i, r := range resource.AuthorizedDomains {
		sAuthorizedDomains[i] = r
	}
	p.SetAuthorizedDomains(sAuthorizedDomains)

	return p
}

// applyConfig handles the gRPC request by passing it to the underlying Config Apply() method.
func (s *ConfigServer) applyConfig(ctx context.Context, c *identitytoolkit.Client, request *identitytoolkitpb.ApplyIdentitytoolkitConfigRequest) (*identitytoolkitpb.IdentitytoolkitConfig, error) {
	p := ProtoToConfig(request.GetResource())
	res, err := c.ApplyConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConfigToProto(res)
	return r, nil
}

// applyIdentitytoolkitConfig handles the gRPC request by passing it to the underlying Config Apply() method.
func (s *ConfigServer) ApplyIdentitytoolkitConfig(ctx context.Context, request *identitytoolkitpb.ApplyIdentitytoolkitConfigRequest) (*identitytoolkitpb.IdentitytoolkitConfig, error) {
	cl, err := createConfigConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyConfig(ctx, cl, request)
}

// DeleteConfig handles the gRPC request by passing it to the underlying Config Delete() method.
func (s *ConfigServer) DeleteIdentitytoolkitConfig(ctx context.Context, request *identitytoolkitpb.DeleteIdentitytoolkitConfigRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Config")

}

// ListIdentitytoolkitConfig is a no-op method because Config has no list method.
func (s *ConfigServer) ListIdentitytoolkitConfig(_ context.Context, _ *identitytoolkitpb.ListIdentitytoolkitConfigRequest) (*identitytoolkitpb.ListIdentitytoolkitConfigResponse, error) {
	return nil, nil
}

func createConfigConfig(ctx context.Context, service_account_file string) (*identitytoolkit.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return identitytoolkit.NewClient(conf), nil
}
