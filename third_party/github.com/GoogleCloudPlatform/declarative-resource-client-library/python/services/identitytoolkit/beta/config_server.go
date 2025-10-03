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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/beta/identitytoolkit_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/beta"
)

// ConfigServer implements the gRPC interface for Config.
type ConfigServer struct{}

// ProtoToConfigSignInEmailHashConfigAlgorithmEnum converts a ConfigSignInEmailHashConfigAlgorithmEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum(e betapb.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum) *beta.ConfigSignInEmailHashConfigAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum_name[int32(e)]; ok {
		e := beta.ConfigSignInEmailHashConfigAlgorithmEnum(n[len("IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSignInHashConfigAlgorithmEnum converts a ConfigSignInHashConfigAlgorithmEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum(e betapb.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum) *beta.ConfigSignInHashConfigAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum_name[int32(e)]; ok {
		e := beta.ConfigSignInHashConfigAlgorithmEnum(n[len("IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailMethodEnum converts a ConfigNotificationSendEmailMethodEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailMethodEnum(e betapb.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum) *beta.ConfigNotificationSendEmailMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum_name[int32(e)]; ok {
		e := beta.ConfigNotificationSendEmailMethodEnum(n[len("IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailSmtpSecurityModeEnum converts a ConfigNotificationSendEmailSmtpSecurityModeEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum(e betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum) *beta.ConfigNotificationSendEmailSmtpSecurityModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum_name[int32(e)]; ok {
		e := beta.ConfigNotificationSendEmailSmtpSecurityModeEnum(n[len("IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum converts a ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(e betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum) *beta.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := beta.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(n[len("IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum converts a ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(e betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum) *beta.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := beta.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(n[len("IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum converts a ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(e betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum) *beta.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := beta.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(n[len("IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailDnsInfoCustomDomainStateEnum converts a ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(e betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum) *beta.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum_name[int32(e)]; ok {
		e := beta.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(n[len("IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(e betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum) *beta.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := beta.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(n[len("IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSubtypeEnum converts a ConfigSubtypeEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSubtypeEnum(e betapb.IdentitytoolkitBetaConfigSubtypeEnum) *beta.ConfigSubtypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigSubtypeEnum_name[int32(e)]; ok {
		e := beta.ConfigSubtypeEnum(n[len("IdentitytoolkitBetaConfigSubtypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigMfaStateEnum converts a ConfigMfaStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitBetaConfigMfaStateEnum(e betapb.IdentitytoolkitBetaConfigMfaStateEnum) *beta.ConfigMfaStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IdentitytoolkitBetaConfigMfaStateEnum_name[int32(e)]; ok {
		e := beta.ConfigMfaStateEnum(n[len("IdentitytoolkitBetaConfigMfaStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSignIn converts a ConfigSignIn object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSignIn(p *betapb.IdentitytoolkitBetaConfigSignIn) *beta.ConfigSignIn {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigSignIn{
		Email:                ProtoToIdentitytoolkitBetaConfigSignInEmail(p.GetEmail()),
		PhoneNumber:          ProtoToIdentitytoolkitBetaConfigSignInPhoneNumber(p.GetPhoneNumber()),
		Anonymous:            ProtoToIdentitytoolkitBetaConfigSignInAnonymous(p.GetAnonymous()),
		AllowDuplicateEmails: dcl.Bool(p.GetAllowDuplicateEmails()),
		HashConfig:           ProtoToIdentitytoolkitBetaConfigSignInHashConfig(p.GetHashConfig()),
	}
	return obj
}

// ProtoToConfigSignInEmail converts a ConfigSignInEmail object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSignInEmail(p *betapb.IdentitytoolkitBetaConfigSignInEmail) *beta.ConfigSignInEmail {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigSignInEmail{
		Enabled:          dcl.Bool(p.GetEnabled()),
		PasswordRequired: dcl.Bool(p.GetPasswordRequired()),
		HashConfig:       ProtoToIdentitytoolkitBetaConfigSignInEmailHashConfig(p.GetHashConfig()),
	}
	return obj
}

// ProtoToConfigSignInEmailHashConfig converts a ConfigSignInEmailHashConfig object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSignInEmailHashConfig(p *betapb.IdentitytoolkitBetaConfigSignInEmailHashConfig) *beta.ConfigSignInEmailHashConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigSignInEmailHashConfig{
		Algorithm:     ProtoToIdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum(p.GetAlgorithm()),
		SignerKey:     dcl.StringOrNil(p.GetSignerKey()),
		SaltSeparator: dcl.StringOrNil(p.GetSaltSeparator()),
		Rounds:        dcl.Int64OrNil(p.GetRounds()),
		MemoryCost:    dcl.Int64OrNil(p.GetMemoryCost()),
	}
	return obj
}

// ProtoToConfigSignInPhoneNumber converts a ConfigSignInPhoneNumber object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSignInPhoneNumber(p *betapb.IdentitytoolkitBetaConfigSignInPhoneNumber) *beta.ConfigSignInPhoneNumber {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigSignInPhoneNumber{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigSignInAnonymous converts a ConfigSignInAnonymous object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSignInAnonymous(p *betapb.IdentitytoolkitBetaConfigSignInAnonymous) *beta.ConfigSignInAnonymous {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigSignInAnonymous{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigSignInHashConfig converts a ConfigSignInHashConfig object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigSignInHashConfig(p *betapb.IdentitytoolkitBetaConfigSignInHashConfig) *beta.ConfigSignInHashConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigSignInHashConfig{
		Algorithm:     ProtoToIdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum(p.GetAlgorithm()),
		SignerKey:     dcl.StringOrNil(p.GetSignerKey()),
		SaltSeparator: dcl.StringOrNil(p.GetSaltSeparator()),
		Rounds:        dcl.Int64OrNil(p.GetRounds()),
		MemoryCost:    dcl.Int64OrNil(p.GetMemoryCost()),
	}
	return obj
}

// ProtoToConfigNotification converts a ConfigNotification object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotification(p *betapb.IdentitytoolkitBetaConfigNotification) *beta.ConfigNotification {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotification{
		SendEmail:     ProtoToIdentitytoolkitBetaConfigNotificationSendEmail(p.GetSendEmail()),
		SendSms:       ProtoToIdentitytoolkitBetaConfigNotificationSendSms(p.GetSendSms()),
		DefaultLocale: dcl.StringOrNil(p.GetDefaultLocale()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmail converts a ConfigNotificationSendEmail object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmail(p *betapb.IdentitytoolkitBetaConfigNotificationSendEmail) *beta.ConfigNotificationSendEmail {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendEmail{
		Method:                             ProtoToIdentitytoolkitBetaConfigNotificationSendEmailMethodEnum(p.GetMethod()),
		Smtp:                               ProtoToIdentitytoolkitBetaConfigNotificationSendEmailSmtp(p.GetSmtp()),
		ResetPasswordTemplate:              ProtoToIdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplate(p.GetResetPasswordTemplate()),
		VerifyEmailTemplate:                ProtoToIdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplate(p.GetVerifyEmailTemplate()),
		ChangeEmailTemplate:                ProtoToIdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplate(p.GetChangeEmailTemplate()),
		CallbackUri:                        dcl.StringOrNil(p.GetCallbackUri()),
		DnsInfo:                            ProtoToIdentitytoolkitBetaConfigNotificationSendEmailDnsInfo(p.GetDnsInfo()),
		RevertSecondFactorAdditionTemplate: ProtoToIdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(p.GetRevertSecondFactorAdditionTemplate()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailSmtp converts a ConfigNotificationSendEmailSmtp object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailSmtp(p *betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtp) *beta.ConfigNotificationSendEmailSmtp {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendEmailSmtp{
		SenderEmail:  dcl.StringOrNil(p.GetSenderEmail()),
		Host:         dcl.StringOrNil(p.GetHost()),
		Port:         dcl.Int64OrNil(p.GetPort()),
		Username:     dcl.StringOrNil(p.GetUsername()),
		Password:     dcl.StringOrNil(p.GetPassword()),
		SecurityMode: ProtoToIdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum(p.GetSecurityMode()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailResetPasswordTemplate converts a ConfigNotificationSendEmailResetPasswordTemplate object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplate(p *betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplate) *beta.ConfigNotificationSendEmailResetPasswordTemplate {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendEmailResetPasswordTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailVerifyEmailTemplate converts a ConfigNotificationSendEmailVerifyEmailTemplate object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplate(p *betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplate) *beta.ConfigNotificationSendEmailVerifyEmailTemplate {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendEmailVerifyEmailTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailChangeEmailTemplate converts a ConfigNotificationSendEmailChangeEmailTemplate object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplate(p *betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplate) *beta.ConfigNotificationSendEmailChangeEmailTemplate {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendEmailChangeEmailTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailDnsInfo converts a ConfigNotificationSendEmailDnsInfo object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailDnsInfo(p *betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfo) *beta.ConfigNotificationSendEmailDnsInfo {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendEmailDnsInfo{
		CustomDomain:                  dcl.StringOrNil(p.GetCustomDomain()),
		UseCustomDomain:               dcl.Bool(p.GetUseCustomDomain()),
		PendingCustomDomain:           dcl.StringOrNil(p.GetPendingCustomDomain()),
		CustomDomainState:             ProtoToIdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(p.GetCustomDomainState()),
		DomainVerificationRequestTime: dcl.StringOrNil(p.GetDomainVerificationRequestTime()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailRevertSecondFactorAdditionTemplate converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(p *betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) *beta.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendSms converts a ConfigNotificationSendSms object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendSms(p *betapb.IdentitytoolkitBetaConfigNotificationSendSms) *beta.ConfigNotificationSendSms {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendSms{
		UseDeviceLocale: dcl.Bool(p.GetUseDeviceLocale()),
		SmsTemplate:     ProtoToIdentitytoolkitBetaConfigNotificationSendSmsSmsTemplate(p.GetSmsTemplate()),
	}
	return obj
}

// ProtoToConfigNotificationSendSmsSmsTemplate converts a ConfigNotificationSendSmsSmsTemplate object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigNotificationSendSmsSmsTemplate(p *betapb.IdentitytoolkitBetaConfigNotificationSendSmsSmsTemplate) *beta.ConfigNotificationSendSmsSmsTemplate {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigNotificationSendSmsSmsTemplate{
		Content: dcl.StringOrNil(p.GetContent()),
	}
	return obj
}

// ProtoToConfigQuota converts a ConfigQuota object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigQuota(p *betapb.IdentitytoolkitBetaConfigQuota) *beta.ConfigQuota {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigQuota{
		SignUpQuotaConfig: ProtoToIdentitytoolkitBetaConfigQuotaSignUpQuotaConfig(p.GetSignUpQuotaConfig()),
	}
	return obj
}

// ProtoToConfigQuotaSignUpQuotaConfig converts a ConfigQuotaSignUpQuotaConfig object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigQuotaSignUpQuotaConfig(p *betapb.IdentitytoolkitBetaConfigQuotaSignUpQuotaConfig) *beta.ConfigQuotaSignUpQuotaConfig {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigQuotaSignUpQuotaConfig{
		Quota:         dcl.Int64OrNil(p.GetQuota()),
		StartTime:     dcl.StringOrNil(p.GetStartTime()),
		QuotaDuration: dcl.StringOrNil(p.GetQuotaDuration()),
	}
	return obj
}

// ProtoToConfigMonitoring converts a ConfigMonitoring object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigMonitoring(p *betapb.IdentitytoolkitBetaConfigMonitoring) *beta.ConfigMonitoring {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigMonitoring{
		RequestLogging: ProtoToIdentitytoolkitBetaConfigMonitoringRequestLogging(p.GetRequestLogging()),
	}
	return obj
}

// ProtoToConfigMonitoringRequestLogging converts a ConfigMonitoringRequestLogging object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigMonitoringRequestLogging(p *betapb.IdentitytoolkitBetaConfigMonitoringRequestLogging) *beta.ConfigMonitoringRequestLogging {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigMonitoringRequestLogging{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigMultiTenant converts a ConfigMultiTenant object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigMultiTenant(p *betapb.IdentitytoolkitBetaConfigMultiTenant) *beta.ConfigMultiTenant {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigMultiTenant{
		AllowTenants:          dcl.Bool(p.GetAllowTenants()),
		DefaultTenantLocation: dcl.StringOrNil(p.GetDefaultTenantLocation()),
	}
	return obj
}

// ProtoToConfigClient converts a ConfigClient object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigClient(p *betapb.IdentitytoolkitBetaConfigClient) *beta.ConfigClient {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigClient{
		ApiKey:            dcl.StringOrNil(p.GetApiKey()),
		Permissions:       ProtoToIdentitytoolkitBetaConfigClientPermissions(p.GetPermissions()),
		FirebaseSubdomain: dcl.StringOrNil(p.GetFirebaseSubdomain()),
	}
	return obj
}

// ProtoToConfigClientPermissions converts a ConfigClientPermissions object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigClientPermissions(p *betapb.IdentitytoolkitBetaConfigClientPermissions) *beta.ConfigClientPermissions {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigClientPermissions{
		DisabledUserSignup:   dcl.Bool(p.GetDisabledUserSignup()),
		DisabledUserDeletion: dcl.Bool(p.GetDisabledUserDeletion()),
	}
	return obj
}

// ProtoToConfigMfa converts a ConfigMfa object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigMfa(p *betapb.IdentitytoolkitBetaConfigMfa) *beta.ConfigMfa {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigMfa{
		State: ProtoToIdentitytoolkitBetaConfigMfaStateEnum(p.GetState()),
	}
	return obj
}

// ProtoToConfigBlockingFunctions converts a ConfigBlockingFunctions object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigBlockingFunctions(p *betapb.IdentitytoolkitBetaConfigBlockingFunctions) *beta.ConfigBlockingFunctions {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigBlockingFunctions{}
	return obj
}

// ProtoToConfigBlockingFunctionsTriggers converts a ConfigBlockingFunctionsTriggers object from its proto representation.
func ProtoToIdentitytoolkitBetaConfigBlockingFunctionsTriggers(p *betapb.IdentitytoolkitBetaConfigBlockingFunctionsTriggers) *beta.ConfigBlockingFunctionsTriggers {
	if p == nil {
		return nil
	}
	obj := &beta.ConfigBlockingFunctionsTriggers{
		FunctionUri: dcl.StringOrNil(p.GetFunctionUri()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToConfig converts a Config resource from its proto representation.
func ProtoToConfig(p *betapb.IdentitytoolkitBetaConfig) *beta.Config {
	obj := &beta.Config{
		SignIn:            ProtoToIdentitytoolkitBetaConfigSignIn(p.GetSignIn()),
		Notification:      ProtoToIdentitytoolkitBetaConfigNotification(p.GetNotification()),
		Quota:             ProtoToIdentitytoolkitBetaConfigQuota(p.GetQuota()),
		Monitoring:        ProtoToIdentitytoolkitBetaConfigMonitoring(p.GetMonitoring()),
		MultiTenant:       ProtoToIdentitytoolkitBetaConfigMultiTenant(p.GetMultiTenant()),
		Subtype:           ProtoToIdentitytoolkitBetaConfigSubtypeEnum(p.GetSubtype()),
		Client:            ProtoToIdentitytoolkitBetaConfigClient(p.GetClient()),
		Mfa:               ProtoToIdentitytoolkitBetaConfigMfa(p.GetMfa()),
		BlockingFunctions: ProtoToIdentitytoolkitBetaConfigBlockingFunctions(p.GetBlockingFunctions()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetAuthorizedDomains() {
		obj.AuthorizedDomains = append(obj.AuthorizedDomains, r)
	}
	return obj
}

// ConfigSignInEmailHashConfigAlgorithmEnumToProto converts a ConfigSignInEmailHashConfigAlgorithmEnum enum to its proto representation.
func IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnumToProto(e *beta.ConfigSignInEmailHashConfigAlgorithmEnum) betapb.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum_value["ConfigSignInEmailHashConfigAlgorithmEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnum(0)
}

// ConfigSignInHashConfigAlgorithmEnumToProto converts a ConfigSignInHashConfigAlgorithmEnum enum to its proto representation.
func IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnumToProto(e *beta.ConfigSignInHashConfigAlgorithmEnum) betapb.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum_value["ConfigSignInHashConfigAlgorithmEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnum(0)
}

// ConfigNotificationSendEmailMethodEnumToProto converts a ConfigNotificationSendEmailMethodEnum enum to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailMethodEnumToProto(e *beta.ConfigNotificationSendEmailMethodEnum) betapb.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum_value["ConfigNotificationSendEmailMethodEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigNotificationSendEmailMethodEnum(0)
}

// ConfigNotificationSendEmailSmtpSecurityModeEnumToProto converts a ConfigNotificationSendEmailSmtpSecurityModeEnum enum to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnumToProto(e *beta.ConfigNotificationSendEmailSmtpSecurityModeEnum) betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum_value["ConfigNotificationSendEmailSmtpSecurityModeEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnum(0)
}

// ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto(e *beta.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum) betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum_value["ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto(e *beta.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum) betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum_value["ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto(e *beta.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum) betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum_value["ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto converts a ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum enum to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto(e *beta.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum) betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum_value["ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(0)
}

// ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto(e *beta.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum) betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum_value["ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(0)
}

// ConfigSubtypeEnumToProto converts a ConfigSubtypeEnum enum to its proto representation.
func IdentitytoolkitBetaConfigSubtypeEnumToProto(e *beta.ConfigSubtypeEnum) betapb.IdentitytoolkitBetaConfigSubtypeEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigSubtypeEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigSubtypeEnum_value["ConfigSubtypeEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigSubtypeEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigSubtypeEnum(0)
}

// ConfigMfaStateEnumToProto converts a ConfigMfaStateEnum enum to its proto representation.
func IdentitytoolkitBetaConfigMfaStateEnumToProto(e *beta.ConfigMfaStateEnum) betapb.IdentitytoolkitBetaConfigMfaStateEnum {
	if e == nil {
		return betapb.IdentitytoolkitBetaConfigMfaStateEnum(0)
	}
	if v, ok := betapb.IdentitytoolkitBetaConfigMfaStateEnum_value["ConfigMfaStateEnum"+string(*e)]; ok {
		return betapb.IdentitytoolkitBetaConfigMfaStateEnum(v)
	}
	return betapb.IdentitytoolkitBetaConfigMfaStateEnum(0)
}

// ConfigSignInToProto converts a ConfigSignIn object to its proto representation.
func IdentitytoolkitBetaConfigSignInToProto(o *beta.ConfigSignIn) *betapb.IdentitytoolkitBetaConfigSignIn {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigSignIn{}
	p.SetEmail(IdentitytoolkitBetaConfigSignInEmailToProto(o.Email))
	p.SetPhoneNumber(IdentitytoolkitBetaConfigSignInPhoneNumberToProto(o.PhoneNumber))
	p.SetAnonymous(IdentitytoolkitBetaConfigSignInAnonymousToProto(o.Anonymous))
	p.SetAllowDuplicateEmails(dcl.ValueOrEmptyBool(o.AllowDuplicateEmails))
	p.SetHashConfig(IdentitytoolkitBetaConfigSignInHashConfigToProto(o.HashConfig))
	return p
}

// ConfigSignInEmailToProto converts a ConfigSignInEmail object to its proto representation.
func IdentitytoolkitBetaConfigSignInEmailToProto(o *beta.ConfigSignInEmail) *betapb.IdentitytoolkitBetaConfigSignInEmail {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigSignInEmail{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetPasswordRequired(dcl.ValueOrEmptyBool(o.PasswordRequired))
	p.SetHashConfig(IdentitytoolkitBetaConfigSignInEmailHashConfigToProto(o.HashConfig))
	return p
}

// ConfigSignInEmailHashConfigToProto converts a ConfigSignInEmailHashConfig object to its proto representation.
func IdentitytoolkitBetaConfigSignInEmailHashConfigToProto(o *beta.ConfigSignInEmailHashConfig) *betapb.IdentitytoolkitBetaConfigSignInEmailHashConfig {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigSignInEmailHashConfig{}
	p.SetAlgorithm(IdentitytoolkitBetaConfigSignInEmailHashConfigAlgorithmEnumToProto(o.Algorithm))
	p.SetSignerKey(dcl.ValueOrEmptyString(o.SignerKey))
	p.SetSaltSeparator(dcl.ValueOrEmptyString(o.SaltSeparator))
	p.SetRounds(dcl.ValueOrEmptyInt64(o.Rounds))
	p.SetMemoryCost(dcl.ValueOrEmptyInt64(o.MemoryCost))
	return p
}

// ConfigSignInPhoneNumberToProto converts a ConfigSignInPhoneNumber object to its proto representation.
func IdentitytoolkitBetaConfigSignInPhoneNumberToProto(o *beta.ConfigSignInPhoneNumber) *betapb.IdentitytoolkitBetaConfigSignInPhoneNumber {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigSignInPhoneNumber{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	mTestPhoneNumbers := make(map[string]string, len(o.TestPhoneNumbers))
	for k, r := range o.TestPhoneNumbers {
		mTestPhoneNumbers[k] = r
	}
	p.SetTestPhoneNumbers(mTestPhoneNumbers)
	return p
}

// ConfigSignInAnonymousToProto converts a ConfigSignInAnonymous object to its proto representation.
func IdentitytoolkitBetaConfigSignInAnonymousToProto(o *beta.ConfigSignInAnonymous) *betapb.IdentitytoolkitBetaConfigSignInAnonymous {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigSignInAnonymous{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ConfigSignInHashConfigToProto converts a ConfigSignInHashConfig object to its proto representation.
func IdentitytoolkitBetaConfigSignInHashConfigToProto(o *beta.ConfigSignInHashConfig) *betapb.IdentitytoolkitBetaConfigSignInHashConfig {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigSignInHashConfig{}
	p.SetAlgorithm(IdentitytoolkitBetaConfigSignInHashConfigAlgorithmEnumToProto(o.Algorithm))
	p.SetSignerKey(dcl.ValueOrEmptyString(o.SignerKey))
	p.SetSaltSeparator(dcl.ValueOrEmptyString(o.SaltSeparator))
	p.SetRounds(dcl.ValueOrEmptyInt64(o.Rounds))
	p.SetMemoryCost(dcl.ValueOrEmptyInt64(o.MemoryCost))
	return p
}

// ConfigNotificationToProto converts a ConfigNotification object to its proto representation.
func IdentitytoolkitBetaConfigNotificationToProto(o *beta.ConfigNotification) *betapb.IdentitytoolkitBetaConfigNotification {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotification{}
	p.SetSendEmail(IdentitytoolkitBetaConfigNotificationSendEmailToProto(o.SendEmail))
	p.SetSendSms(IdentitytoolkitBetaConfigNotificationSendSmsToProto(o.SendSms))
	p.SetDefaultLocale(dcl.ValueOrEmptyString(o.DefaultLocale))
	return p
}

// ConfigNotificationSendEmailToProto converts a ConfigNotificationSendEmail object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailToProto(o *beta.ConfigNotificationSendEmail) *betapb.IdentitytoolkitBetaConfigNotificationSendEmail {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendEmail{}
	p.SetMethod(IdentitytoolkitBetaConfigNotificationSendEmailMethodEnumToProto(o.Method))
	p.SetSmtp(IdentitytoolkitBetaConfigNotificationSendEmailSmtpToProto(o.Smtp))
	p.SetResetPasswordTemplate(IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateToProto(o.ResetPasswordTemplate))
	p.SetVerifyEmailTemplate(IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateToProto(o.VerifyEmailTemplate))
	p.SetChangeEmailTemplate(IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateToProto(o.ChangeEmailTemplate))
	p.SetCallbackUri(dcl.ValueOrEmptyString(o.CallbackUri))
	p.SetDnsInfo(IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoToProto(o.DnsInfo))
	p.SetRevertSecondFactorAdditionTemplate(IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto(o.RevertSecondFactorAdditionTemplate))
	return p
}

// ConfigNotificationSendEmailSmtpToProto converts a ConfigNotificationSendEmailSmtp object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailSmtpToProto(o *beta.ConfigNotificationSendEmailSmtp) *betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtp {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendEmailSmtp{}
	p.SetSenderEmail(dcl.ValueOrEmptyString(o.SenderEmail))
	p.SetHost(dcl.ValueOrEmptyString(o.Host))
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	p.SetPassword(dcl.ValueOrEmptyString(o.Password))
	p.SetSecurityMode(IdentitytoolkitBetaConfigNotificationSendEmailSmtpSecurityModeEnumToProto(o.SecurityMode))
	return p
}

// ConfigNotificationSendEmailResetPasswordTemplateToProto converts a ConfigNotificationSendEmailResetPasswordTemplate object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateToProto(o *beta.ConfigNotificationSendEmailResetPasswordTemplate) *betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplate {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitBetaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailVerifyEmailTemplateToProto converts a ConfigNotificationSendEmailVerifyEmailTemplate object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateToProto(o *beta.ConfigNotificationSendEmailVerifyEmailTemplate) *betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplate {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitBetaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailChangeEmailTemplateToProto converts a ConfigNotificationSendEmailChangeEmailTemplate object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateToProto(o *beta.ConfigNotificationSendEmailChangeEmailTemplate) *betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplate {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitBetaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailDnsInfoToProto converts a ConfigNotificationSendEmailDnsInfo object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoToProto(o *beta.ConfigNotificationSendEmailDnsInfo) *betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfo {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendEmailDnsInfo{}
	p.SetCustomDomain(dcl.ValueOrEmptyString(o.CustomDomain))
	p.SetUseCustomDomain(dcl.ValueOrEmptyBool(o.UseCustomDomain))
	p.SetPendingCustomDomain(dcl.ValueOrEmptyString(o.PendingCustomDomain))
	p.SetCustomDomainState(IdentitytoolkitBetaConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto(o.CustomDomainState))
	p.SetDomainVerificationRequestTime(dcl.ValueOrEmptyString(o.DomainVerificationRequestTime))
	return p
}

// ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto(o *beta.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) *betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitBetaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendSmsToProto converts a ConfigNotificationSendSms object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendSmsToProto(o *beta.ConfigNotificationSendSms) *betapb.IdentitytoolkitBetaConfigNotificationSendSms {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendSms{}
	p.SetUseDeviceLocale(dcl.ValueOrEmptyBool(o.UseDeviceLocale))
	p.SetSmsTemplate(IdentitytoolkitBetaConfigNotificationSendSmsSmsTemplateToProto(o.SmsTemplate))
	return p
}

// ConfigNotificationSendSmsSmsTemplateToProto converts a ConfigNotificationSendSmsSmsTemplate object to its proto representation.
func IdentitytoolkitBetaConfigNotificationSendSmsSmsTemplateToProto(o *beta.ConfigNotificationSendSmsSmsTemplate) *betapb.IdentitytoolkitBetaConfigNotificationSendSmsSmsTemplate {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigNotificationSendSmsSmsTemplate{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	return p
}

// ConfigQuotaToProto converts a ConfigQuota object to its proto representation.
func IdentitytoolkitBetaConfigQuotaToProto(o *beta.ConfigQuota) *betapb.IdentitytoolkitBetaConfigQuota {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigQuota{}
	p.SetSignUpQuotaConfig(IdentitytoolkitBetaConfigQuotaSignUpQuotaConfigToProto(o.SignUpQuotaConfig))
	return p
}

// ConfigQuotaSignUpQuotaConfigToProto converts a ConfigQuotaSignUpQuotaConfig object to its proto representation.
func IdentitytoolkitBetaConfigQuotaSignUpQuotaConfigToProto(o *beta.ConfigQuotaSignUpQuotaConfig) *betapb.IdentitytoolkitBetaConfigQuotaSignUpQuotaConfig {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigQuotaSignUpQuotaConfig{}
	p.SetQuota(dcl.ValueOrEmptyInt64(o.Quota))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetQuotaDuration(dcl.ValueOrEmptyString(o.QuotaDuration))
	return p
}

// ConfigMonitoringToProto converts a ConfigMonitoring object to its proto representation.
func IdentitytoolkitBetaConfigMonitoringToProto(o *beta.ConfigMonitoring) *betapb.IdentitytoolkitBetaConfigMonitoring {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigMonitoring{}
	p.SetRequestLogging(IdentitytoolkitBetaConfigMonitoringRequestLoggingToProto(o.RequestLogging))
	return p
}

// ConfigMonitoringRequestLoggingToProto converts a ConfigMonitoringRequestLogging object to its proto representation.
func IdentitytoolkitBetaConfigMonitoringRequestLoggingToProto(o *beta.ConfigMonitoringRequestLogging) *betapb.IdentitytoolkitBetaConfigMonitoringRequestLogging {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigMonitoringRequestLogging{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ConfigMultiTenantToProto converts a ConfigMultiTenant object to its proto representation.
func IdentitytoolkitBetaConfigMultiTenantToProto(o *beta.ConfigMultiTenant) *betapb.IdentitytoolkitBetaConfigMultiTenant {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigMultiTenant{}
	p.SetAllowTenants(dcl.ValueOrEmptyBool(o.AllowTenants))
	p.SetDefaultTenantLocation(dcl.ValueOrEmptyString(o.DefaultTenantLocation))
	return p
}

// ConfigClientToProto converts a ConfigClient object to its proto representation.
func IdentitytoolkitBetaConfigClientToProto(o *beta.ConfigClient) *betapb.IdentitytoolkitBetaConfigClient {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigClient{}
	p.SetApiKey(dcl.ValueOrEmptyString(o.ApiKey))
	p.SetPermissions(IdentitytoolkitBetaConfigClientPermissionsToProto(o.Permissions))
	p.SetFirebaseSubdomain(dcl.ValueOrEmptyString(o.FirebaseSubdomain))
	return p
}

// ConfigClientPermissionsToProto converts a ConfigClientPermissions object to its proto representation.
func IdentitytoolkitBetaConfigClientPermissionsToProto(o *beta.ConfigClientPermissions) *betapb.IdentitytoolkitBetaConfigClientPermissions {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigClientPermissions{}
	p.SetDisabledUserSignup(dcl.ValueOrEmptyBool(o.DisabledUserSignup))
	p.SetDisabledUserDeletion(dcl.ValueOrEmptyBool(o.DisabledUserDeletion))
	return p
}

// ConfigMfaToProto converts a ConfigMfa object to its proto representation.
func IdentitytoolkitBetaConfigMfaToProto(o *beta.ConfigMfa) *betapb.IdentitytoolkitBetaConfigMfa {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigMfa{}
	p.SetState(IdentitytoolkitBetaConfigMfaStateEnumToProto(o.State))
	return p
}

// ConfigBlockingFunctionsToProto converts a ConfigBlockingFunctions object to its proto representation.
func IdentitytoolkitBetaConfigBlockingFunctionsToProto(o *beta.ConfigBlockingFunctions) *betapb.IdentitytoolkitBetaConfigBlockingFunctions {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigBlockingFunctions{}
	mTriggers := make(map[string]*betapb.IdentitytoolkitBetaConfigBlockingFunctionsTriggers, len(o.Triggers))
	for k, r := range o.Triggers {
		mTriggers[k] = IdentitytoolkitBetaConfigBlockingFunctionsTriggersToProto(&r)
	}
	p.SetTriggers(mTriggers)
	return p
}

// ConfigBlockingFunctionsTriggersToProto converts a ConfigBlockingFunctionsTriggers object to its proto representation.
func IdentitytoolkitBetaConfigBlockingFunctionsTriggersToProto(o *beta.ConfigBlockingFunctionsTriggers) *betapb.IdentitytoolkitBetaConfigBlockingFunctionsTriggers {
	if o == nil {
		return nil
	}
	p := &betapb.IdentitytoolkitBetaConfigBlockingFunctionsTriggers{}
	p.SetFunctionUri(dcl.ValueOrEmptyString(o.FunctionUri))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// ConfigToProto converts a Config resource to its proto representation.
func ConfigToProto(resource *beta.Config) *betapb.IdentitytoolkitBetaConfig {
	p := &betapb.IdentitytoolkitBetaConfig{}
	p.SetSignIn(IdentitytoolkitBetaConfigSignInToProto(resource.SignIn))
	p.SetNotification(IdentitytoolkitBetaConfigNotificationToProto(resource.Notification))
	p.SetQuota(IdentitytoolkitBetaConfigQuotaToProto(resource.Quota))
	p.SetMonitoring(IdentitytoolkitBetaConfigMonitoringToProto(resource.Monitoring))
	p.SetMultiTenant(IdentitytoolkitBetaConfigMultiTenantToProto(resource.MultiTenant))
	p.SetSubtype(IdentitytoolkitBetaConfigSubtypeEnumToProto(resource.Subtype))
	p.SetClient(IdentitytoolkitBetaConfigClientToProto(resource.Client))
	p.SetMfa(IdentitytoolkitBetaConfigMfaToProto(resource.Mfa))
	p.SetBlockingFunctions(IdentitytoolkitBetaConfigBlockingFunctionsToProto(resource.BlockingFunctions))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sAuthorizedDomains := make([]string, len(resource.AuthorizedDomains))
	for i, r := range resource.AuthorizedDomains {
		sAuthorizedDomains[i] = r
	}
	p.SetAuthorizedDomains(sAuthorizedDomains)

	return p
}

// applyConfig handles the gRPC request by passing it to the underlying Config Apply() method.
func (s *ConfigServer) applyConfig(ctx context.Context, c *beta.Client, request *betapb.ApplyIdentitytoolkitBetaConfigRequest) (*betapb.IdentitytoolkitBetaConfig, error) {
	p := ProtoToConfig(request.GetResource())
	res, err := c.ApplyConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConfigToProto(res)
	return r, nil
}

// applyIdentitytoolkitBetaConfig handles the gRPC request by passing it to the underlying Config Apply() method.
func (s *ConfigServer) ApplyIdentitytoolkitBetaConfig(ctx context.Context, request *betapb.ApplyIdentitytoolkitBetaConfigRequest) (*betapb.IdentitytoolkitBetaConfig, error) {
	cl, err := createConfigConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyConfig(ctx, cl, request)
}

// DeleteConfig handles the gRPC request by passing it to the underlying Config Delete() method.
func (s *ConfigServer) DeleteIdentitytoolkitBetaConfig(ctx context.Context, request *betapb.DeleteIdentitytoolkitBetaConfigRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Config")

}

// ListIdentitytoolkitBetaConfig is a no-op method because Config has no list method.
func (s *ConfigServer) ListIdentitytoolkitBetaConfig(_ context.Context, _ *betapb.ListIdentitytoolkitBetaConfigRequest) (*betapb.ListIdentitytoolkitBetaConfigResponse, error) {
	return nil, nil
}

func createConfigConfig(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
