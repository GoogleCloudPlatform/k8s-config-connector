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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/identitytoolkit/alpha/identitytoolkit_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit/alpha"
)

// ConfigServer implements the gRPC interface for Config.
type ConfigServer struct{}

// ProtoToConfigSignInEmailHashConfigAlgorithmEnum converts a ConfigSignInEmailHashConfigAlgorithmEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum(e alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum) *alpha.ConfigSignInEmailHashConfigAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.ConfigSignInEmailHashConfigAlgorithmEnum(n[len("IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSignInHashConfigAlgorithmEnum converts a ConfigSignInHashConfigAlgorithmEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum(e alphapb.IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum) *alpha.ConfigSignInHashConfigAlgorithmEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum_name[int32(e)]; ok {
		e := alpha.ConfigSignInHashConfigAlgorithmEnum(n[len("IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailMethodEnum converts a ConfigNotificationSendEmailMethodEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum(e alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum) *alpha.ConfigNotificationSendEmailMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum_name[int32(e)]; ok {
		e := alpha.ConfigNotificationSendEmailMethodEnum(n[len("IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailSmtpSecurityModeEnum converts a ConfigNotificationSendEmailSmtpSecurityModeEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum(e alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum) *alpha.ConfigNotificationSendEmailSmtpSecurityModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum_name[int32(e)]; ok {
		e := alpha.ConfigNotificationSendEmailSmtpSecurityModeEnum(n[len("IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum converts a ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(e alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum) *alpha.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := alpha.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(n[len("IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum converts a ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(e alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum) *alpha.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := alpha.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(n[len("IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum converts a ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(e alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum) *alpha.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := alpha.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(n[len("IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailDnsInfoCustomDomainStateEnum converts a ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(e alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum) *alpha.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum_name[int32(e)]; ok {
		e := alpha.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(n[len("IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(e alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum) *alpha.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum_name[int32(e)]; ok {
		e := alpha.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(n[len("IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSubtypeEnum converts a ConfigSubtypeEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSubtypeEnum(e alphapb.IdentitytoolkitAlphaConfigSubtypeEnum) *alpha.ConfigSubtypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigSubtypeEnum_name[int32(e)]; ok {
		e := alpha.ConfigSubtypeEnum(n[len("IdentitytoolkitAlphaConfigSubtypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigMfaStateEnum converts a ConfigMfaStateEnum enum from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigMfaStateEnum(e alphapb.IdentitytoolkitAlphaConfigMfaStateEnum) *alpha.ConfigMfaStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IdentitytoolkitAlphaConfigMfaStateEnum_name[int32(e)]; ok {
		e := alpha.ConfigMfaStateEnum(n[len("IdentitytoolkitAlphaConfigMfaStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToConfigSignIn converts a ConfigSignIn object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSignIn(p *alphapb.IdentitytoolkitAlphaConfigSignIn) *alpha.ConfigSignIn {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigSignIn{
		Email:                ProtoToIdentitytoolkitAlphaConfigSignInEmail(p.GetEmail()),
		PhoneNumber:          ProtoToIdentitytoolkitAlphaConfigSignInPhoneNumber(p.GetPhoneNumber()),
		Anonymous:            ProtoToIdentitytoolkitAlphaConfigSignInAnonymous(p.GetAnonymous()),
		AllowDuplicateEmails: dcl.Bool(p.GetAllowDuplicateEmails()),
		HashConfig:           ProtoToIdentitytoolkitAlphaConfigSignInHashConfig(p.GetHashConfig()),
	}
	return obj
}

// ProtoToConfigSignInEmail converts a ConfigSignInEmail object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSignInEmail(p *alphapb.IdentitytoolkitAlphaConfigSignInEmail) *alpha.ConfigSignInEmail {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigSignInEmail{
		Enabled:          dcl.Bool(p.GetEnabled()),
		PasswordRequired: dcl.Bool(p.GetPasswordRequired()),
		HashConfig:       ProtoToIdentitytoolkitAlphaConfigSignInEmailHashConfig(p.GetHashConfig()),
	}
	return obj
}

// ProtoToConfigSignInEmailHashConfig converts a ConfigSignInEmailHashConfig object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSignInEmailHashConfig(p *alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfig) *alpha.ConfigSignInEmailHashConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigSignInEmailHashConfig{
		Algorithm:     ProtoToIdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum(p.GetAlgorithm()),
		SignerKey:     dcl.StringOrNil(p.GetSignerKey()),
		SaltSeparator: dcl.StringOrNil(p.GetSaltSeparator()),
		Rounds:        dcl.Int64OrNil(p.GetRounds()),
		MemoryCost:    dcl.Int64OrNil(p.GetMemoryCost()),
	}
	return obj
}

// ProtoToConfigSignInPhoneNumber converts a ConfigSignInPhoneNumber object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSignInPhoneNumber(p *alphapb.IdentitytoolkitAlphaConfigSignInPhoneNumber) *alpha.ConfigSignInPhoneNumber {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigSignInPhoneNumber{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigSignInAnonymous converts a ConfigSignInAnonymous object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSignInAnonymous(p *alphapb.IdentitytoolkitAlphaConfigSignInAnonymous) *alpha.ConfigSignInAnonymous {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigSignInAnonymous{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigSignInHashConfig converts a ConfigSignInHashConfig object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigSignInHashConfig(p *alphapb.IdentitytoolkitAlphaConfigSignInHashConfig) *alpha.ConfigSignInHashConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigSignInHashConfig{
		Algorithm:     ProtoToIdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum(p.GetAlgorithm()),
		SignerKey:     dcl.StringOrNil(p.GetSignerKey()),
		SaltSeparator: dcl.StringOrNil(p.GetSaltSeparator()),
		Rounds:        dcl.Int64OrNil(p.GetRounds()),
		MemoryCost:    dcl.Int64OrNil(p.GetMemoryCost()),
	}
	return obj
}

// ProtoToConfigNotification converts a ConfigNotification object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotification(p *alphapb.IdentitytoolkitAlphaConfigNotification) *alpha.ConfigNotification {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotification{
		SendEmail:     ProtoToIdentitytoolkitAlphaConfigNotificationSendEmail(p.GetSendEmail()),
		SendSms:       ProtoToIdentitytoolkitAlphaConfigNotificationSendSms(p.GetSendSms()),
		DefaultLocale: dcl.StringOrNil(p.GetDefaultLocale()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmail converts a ConfigNotificationSendEmail object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmail(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmail) *alpha.ConfigNotificationSendEmail {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendEmail{
		Method:                             ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum(p.GetMethod()),
		Smtp:                               ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailSmtp(p.GetSmtp()),
		ResetPasswordTemplate:              ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplate(p.GetResetPasswordTemplate()),
		VerifyEmailTemplate:                ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplate(p.GetVerifyEmailTemplate()),
		ChangeEmailTemplate:                ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplate(p.GetChangeEmailTemplate()),
		CallbackUri:                        dcl.StringOrNil(p.GetCallbackUri()),
		DnsInfo:                            ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailDnsInfo(p.GetDnsInfo()),
		RevertSecondFactorAdditionTemplate: ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(p.GetRevertSecondFactorAdditionTemplate()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailSmtp converts a ConfigNotificationSendEmailSmtp object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailSmtp(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtp) *alpha.ConfigNotificationSendEmailSmtp {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendEmailSmtp{
		SenderEmail:  dcl.StringOrNil(p.GetSenderEmail()),
		Host:         dcl.StringOrNil(p.GetHost()),
		Port:         dcl.Int64OrNil(p.GetPort()),
		Username:     dcl.StringOrNil(p.GetUsername()),
		Password:     dcl.StringOrNil(p.GetPassword()),
		SecurityMode: ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum(p.GetSecurityMode()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailResetPasswordTemplate converts a ConfigNotificationSendEmailResetPasswordTemplate object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplate(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplate) *alpha.ConfigNotificationSendEmailResetPasswordTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendEmailResetPasswordTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailVerifyEmailTemplate converts a ConfigNotificationSendEmailVerifyEmailTemplate object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplate(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplate) *alpha.ConfigNotificationSendEmailVerifyEmailTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendEmailVerifyEmailTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailChangeEmailTemplate converts a ConfigNotificationSendEmailChangeEmailTemplate object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplate(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplate) *alpha.ConfigNotificationSendEmailChangeEmailTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendEmailChangeEmailTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailDnsInfo converts a ConfigNotificationSendEmailDnsInfo object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailDnsInfo(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfo) *alpha.ConfigNotificationSendEmailDnsInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendEmailDnsInfo{
		CustomDomain:                  dcl.StringOrNil(p.GetCustomDomain()),
		UseCustomDomain:               dcl.Bool(p.GetUseCustomDomain()),
		PendingCustomDomain:           dcl.StringOrNil(p.GetPendingCustomDomain()),
		CustomDomainState:             ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(p.GetCustomDomainState()),
		DomainVerificationRequestTime: dcl.StringOrNil(p.GetDomainVerificationRequestTime()),
	}
	return obj
}

// ProtoToConfigNotificationSendEmailRevertSecondFactorAdditionTemplate converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) *alpha.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{
		SenderLocalPart:   dcl.StringOrNil(p.GetSenderLocalPart()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		SenderDisplayName: dcl.StringOrNil(p.GetSenderDisplayName()),
		Body:              dcl.StringOrNil(p.GetBody()),
		BodyFormat:        ProtoToIdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(p.GetBodyFormat()),
		ReplyTo:           dcl.StringOrNil(p.GetReplyTo()),
		Customized:        dcl.Bool(p.GetCustomized()),
	}
	return obj
}

// ProtoToConfigNotificationSendSms converts a ConfigNotificationSendSms object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendSms(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendSms) *alpha.ConfigNotificationSendSms {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendSms{
		UseDeviceLocale: dcl.Bool(p.GetUseDeviceLocale()),
		SmsTemplate:     ProtoToIdentitytoolkitAlphaConfigNotificationSendSmsSmsTemplate(p.GetSmsTemplate()),
	}
	return obj
}

// ProtoToConfigNotificationSendSmsSmsTemplate converts a ConfigNotificationSendSmsSmsTemplate object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigNotificationSendSmsSmsTemplate(p *alphapb.IdentitytoolkitAlphaConfigNotificationSendSmsSmsTemplate) *alpha.ConfigNotificationSendSmsSmsTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigNotificationSendSmsSmsTemplate{
		Content: dcl.StringOrNil(p.GetContent()),
	}
	return obj
}

// ProtoToConfigQuota converts a ConfigQuota object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigQuota(p *alphapb.IdentitytoolkitAlphaConfigQuota) *alpha.ConfigQuota {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigQuota{
		SignUpQuotaConfig: ProtoToIdentitytoolkitAlphaConfigQuotaSignUpQuotaConfig(p.GetSignUpQuotaConfig()),
	}
	return obj
}

// ProtoToConfigQuotaSignUpQuotaConfig converts a ConfigQuotaSignUpQuotaConfig object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigQuotaSignUpQuotaConfig(p *alphapb.IdentitytoolkitAlphaConfigQuotaSignUpQuotaConfig) *alpha.ConfigQuotaSignUpQuotaConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigQuotaSignUpQuotaConfig{
		Quota:         dcl.Int64OrNil(p.GetQuota()),
		StartTime:     dcl.StringOrNil(p.GetStartTime()),
		QuotaDuration: dcl.StringOrNil(p.GetQuotaDuration()),
	}
	return obj
}

// ProtoToConfigMonitoring converts a ConfigMonitoring object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigMonitoring(p *alphapb.IdentitytoolkitAlphaConfigMonitoring) *alpha.ConfigMonitoring {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigMonitoring{
		RequestLogging: ProtoToIdentitytoolkitAlphaConfigMonitoringRequestLogging(p.GetRequestLogging()),
	}
	return obj
}

// ProtoToConfigMonitoringRequestLogging converts a ConfigMonitoringRequestLogging object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigMonitoringRequestLogging(p *alphapb.IdentitytoolkitAlphaConfigMonitoringRequestLogging) *alpha.ConfigMonitoringRequestLogging {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigMonitoringRequestLogging{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToConfigMultiTenant converts a ConfigMultiTenant object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigMultiTenant(p *alphapb.IdentitytoolkitAlphaConfigMultiTenant) *alpha.ConfigMultiTenant {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigMultiTenant{
		AllowTenants:          dcl.Bool(p.GetAllowTenants()),
		DefaultTenantLocation: dcl.StringOrNil(p.GetDefaultTenantLocation()),
	}
	return obj
}

// ProtoToConfigClient converts a ConfigClient object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigClient(p *alphapb.IdentitytoolkitAlphaConfigClient) *alpha.ConfigClient {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigClient{
		ApiKey:            dcl.StringOrNil(p.GetApiKey()),
		Permissions:       ProtoToIdentitytoolkitAlphaConfigClientPermissions(p.GetPermissions()),
		FirebaseSubdomain: dcl.StringOrNil(p.GetFirebaseSubdomain()),
	}
	return obj
}

// ProtoToConfigClientPermissions converts a ConfigClientPermissions object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigClientPermissions(p *alphapb.IdentitytoolkitAlphaConfigClientPermissions) *alpha.ConfigClientPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigClientPermissions{
		DisabledUserSignup:   dcl.Bool(p.GetDisabledUserSignup()),
		DisabledUserDeletion: dcl.Bool(p.GetDisabledUserDeletion()),
	}
	return obj
}

// ProtoToConfigMfa converts a ConfigMfa object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigMfa(p *alphapb.IdentitytoolkitAlphaConfigMfa) *alpha.ConfigMfa {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigMfa{
		State: ProtoToIdentitytoolkitAlphaConfigMfaStateEnum(p.GetState()),
	}
	return obj
}

// ProtoToConfigBlockingFunctions converts a ConfigBlockingFunctions object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigBlockingFunctions(p *alphapb.IdentitytoolkitAlphaConfigBlockingFunctions) *alpha.ConfigBlockingFunctions {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigBlockingFunctions{}
	return obj
}

// ProtoToConfigBlockingFunctionsTriggers converts a ConfigBlockingFunctionsTriggers object from its proto representation.
func ProtoToIdentitytoolkitAlphaConfigBlockingFunctionsTriggers(p *alphapb.IdentitytoolkitAlphaConfigBlockingFunctionsTriggers) *alpha.ConfigBlockingFunctionsTriggers {
	if p == nil {
		return nil
	}
	obj := &alpha.ConfigBlockingFunctionsTriggers{
		FunctionUri: dcl.StringOrNil(p.GetFunctionUri()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToConfig converts a Config resource from its proto representation.
func ProtoToConfig(p *alphapb.IdentitytoolkitAlphaConfig) *alpha.Config {
	obj := &alpha.Config{
		SignIn:            ProtoToIdentitytoolkitAlphaConfigSignIn(p.GetSignIn()),
		Notification:      ProtoToIdentitytoolkitAlphaConfigNotification(p.GetNotification()),
		Quota:             ProtoToIdentitytoolkitAlphaConfigQuota(p.GetQuota()),
		Monitoring:        ProtoToIdentitytoolkitAlphaConfigMonitoring(p.GetMonitoring()),
		MultiTenant:       ProtoToIdentitytoolkitAlphaConfigMultiTenant(p.GetMultiTenant()),
		Subtype:           ProtoToIdentitytoolkitAlphaConfigSubtypeEnum(p.GetSubtype()),
		Client:            ProtoToIdentitytoolkitAlphaConfigClient(p.GetClient()),
		Mfa:               ProtoToIdentitytoolkitAlphaConfigMfa(p.GetMfa()),
		BlockingFunctions: ProtoToIdentitytoolkitAlphaConfigBlockingFunctions(p.GetBlockingFunctions()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetAuthorizedDomains() {
		obj.AuthorizedDomains = append(obj.AuthorizedDomains, r)
	}
	return obj
}

// ConfigSignInEmailHashConfigAlgorithmEnumToProto converts a ConfigSignInEmailHashConfigAlgorithmEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnumToProto(e *alpha.ConfigSignInEmailHashConfigAlgorithmEnum) alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum_value["ConfigSignInEmailHashConfigAlgorithmEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnum(0)
}

// ConfigSignInHashConfigAlgorithmEnumToProto converts a ConfigSignInHashConfigAlgorithmEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnumToProto(e *alpha.ConfigSignInHashConfigAlgorithmEnum) alphapb.IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum_value["ConfigSignInHashConfigAlgorithmEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnum(0)
}

// ConfigNotificationSendEmailMethodEnumToProto converts a ConfigNotificationSendEmailMethodEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnumToProto(e *alpha.ConfigNotificationSendEmailMethodEnum) alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum_value["ConfigNotificationSendEmailMethodEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnum(0)
}

// ConfigNotificationSendEmailSmtpSecurityModeEnumToProto converts a ConfigNotificationSendEmailSmtpSecurityModeEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnumToProto(e *alpha.ConfigNotificationSendEmailSmtpSecurityModeEnum) alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum_value["ConfigNotificationSendEmailSmtpSecurityModeEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnum(0)
}

// ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto(e *alpha.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum) alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum_value["ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto(e *alpha.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum) alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum_value["ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto(e *alpha.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum) alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum_value["ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(0)
}

// ConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto converts a ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto(e *alpha.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum) alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum_value["ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(0)
}

// ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto(e *alpha.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum) alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum_value["ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(0)
}

// ConfigSubtypeEnumToProto converts a ConfigSubtypeEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigSubtypeEnumToProto(e *alpha.ConfigSubtypeEnum) alphapb.IdentitytoolkitAlphaConfigSubtypeEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigSubtypeEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigSubtypeEnum_value["ConfigSubtypeEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigSubtypeEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigSubtypeEnum(0)
}

// ConfigMfaStateEnumToProto converts a ConfigMfaStateEnum enum to its proto representation.
func IdentitytoolkitAlphaConfigMfaStateEnumToProto(e *alpha.ConfigMfaStateEnum) alphapb.IdentitytoolkitAlphaConfigMfaStateEnum {
	if e == nil {
		return alphapb.IdentitytoolkitAlphaConfigMfaStateEnum(0)
	}
	if v, ok := alphapb.IdentitytoolkitAlphaConfigMfaStateEnum_value["ConfigMfaStateEnum"+string(*e)]; ok {
		return alphapb.IdentitytoolkitAlphaConfigMfaStateEnum(v)
	}
	return alphapb.IdentitytoolkitAlphaConfigMfaStateEnum(0)
}

// ConfigSignInToProto converts a ConfigSignIn object to its proto representation.
func IdentitytoolkitAlphaConfigSignInToProto(o *alpha.ConfigSignIn) *alphapb.IdentitytoolkitAlphaConfigSignIn {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigSignIn{}
	p.SetEmail(IdentitytoolkitAlphaConfigSignInEmailToProto(o.Email))
	p.SetPhoneNumber(IdentitytoolkitAlphaConfigSignInPhoneNumberToProto(o.PhoneNumber))
	p.SetAnonymous(IdentitytoolkitAlphaConfigSignInAnonymousToProto(o.Anonymous))
	p.SetAllowDuplicateEmails(dcl.ValueOrEmptyBool(o.AllowDuplicateEmails))
	p.SetHashConfig(IdentitytoolkitAlphaConfigSignInHashConfigToProto(o.HashConfig))
	return p
}

// ConfigSignInEmailToProto converts a ConfigSignInEmail object to its proto representation.
func IdentitytoolkitAlphaConfigSignInEmailToProto(o *alpha.ConfigSignInEmail) *alphapb.IdentitytoolkitAlphaConfigSignInEmail {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigSignInEmail{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetPasswordRequired(dcl.ValueOrEmptyBool(o.PasswordRequired))
	p.SetHashConfig(IdentitytoolkitAlphaConfigSignInEmailHashConfigToProto(o.HashConfig))
	return p
}

// ConfigSignInEmailHashConfigToProto converts a ConfigSignInEmailHashConfig object to its proto representation.
func IdentitytoolkitAlphaConfigSignInEmailHashConfigToProto(o *alpha.ConfigSignInEmailHashConfig) *alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigSignInEmailHashConfig{}
	p.SetAlgorithm(IdentitytoolkitAlphaConfigSignInEmailHashConfigAlgorithmEnumToProto(o.Algorithm))
	p.SetSignerKey(dcl.ValueOrEmptyString(o.SignerKey))
	p.SetSaltSeparator(dcl.ValueOrEmptyString(o.SaltSeparator))
	p.SetRounds(dcl.ValueOrEmptyInt64(o.Rounds))
	p.SetMemoryCost(dcl.ValueOrEmptyInt64(o.MemoryCost))
	return p
}

// ConfigSignInPhoneNumberToProto converts a ConfigSignInPhoneNumber object to its proto representation.
func IdentitytoolkitAlphaConfigSignInPhoneNumberToProto(o *alpha.ConfigSignInPhoneNumber) *alphapb.IdentitytoolkitAlphaConfigSignInPhoneNumber {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigSignInPhoneNumber{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	mTestPhoneNumbers := make(map[string]string, len(o.TestPhoneNumbers))
	for k, r := range o.TestPhoneNumbers {
		mTestPhoneNumbers[k] = r
	}
	p.SetTestPhoneNumbers(mTestPhoneNumbers)
	return p
}

// ConfigSignInAnonymousToProto converts a ConfigSignInAnonymous object to its proto representation.
func IdentitytoolkitAlphaConfigSignInAnonymousToProto(o *alpha.ConfigSignInAnonymous) *alphapb.IdentitytoolkitAlphaConfigSignInAnonymous {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigSignInAnonymous{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ConfigSignInHashConfigToProto converts a ConfigSignInHashConfig object to its proto representation.
func IdentitytoolkitAlphaConfigSignInHashConfigToProto(o *alpha.ConfigSignInHashConfig) *alphapb.IdentitytoolkitAlphaConfigSignInHashConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigSignInHashConfig{}
	p.SetAlgorithm(IdentitytoolkitAlphaConfigSignInHashConfigAlgorithmEnumToProto(o.Algorithm))
	p.SetSignerKey(dcl.ValueOrEmptyString(o.SignerKey))
	p.SetSaltSeparator(dcl.ValueOrEmptyString(o.SaltSeparator))
	p.SetRounds(dcl.ValueOrEmptyInt64(o.Rounds))
	p.SetMemoryCost(dcl.ValueOrEmptyInt64(o.MemoryCost))
	return p
}

// ConfigNotificationToProto converts a ConfigNotification object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationToProto(o *alpha.ConfigNotification) *alphapb.IdentitytoolkitAlphaConfigNotification {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotification{}
	p.SetSendEmail(IdentitytoolkitAlphaConfigNotificationSendEmailToProto(o.SendEmail))
	p.SetSendSms(IdentitytoolkitAlphaConfigNotificationSendSmsToProto(o.SendSms))
	p.SetDefaultLocale(dcl.ValueOrEmptyString(o.DefaultLocale))
	return p
}

// ConfigNotificationSendEmailToProto converts a ConfigNotificationSendEmail object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailToProto(o *alpha.ConfigNotificationSendEmail) *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmail {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendEmail{}
	p.SetMethod(IdentitytoolkitAlphaConfigNotificationSendEmailMethodEnumToProto(o.Method))
	p.SetSmtp(IdentitytoolkitAlphaConfigNotificationSendEmailSmtpToProto(o.Smtp))
	p.SetResetPasswordTemplate(IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateToProto(o.ResetPasswordTemplate))
	p.SetVerifyEmailTemplate(IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateToProto(o.VerifyEmailTemplate))
	p.SetChangeEmailTemplate(IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateToProto(o.ChangeEmailTemplate))
	p.SetCallbackUri(dcl.ValueOrEmptyString(o.CallbackUri))
	p.SetDnsInfo(IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoToProto(o.DnsInfo))
	p.SetRevertSecondFactorAdditionTemplate(IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto(o.RevertSecondFactorAdditionTemplate))
	return p
}

// ConfigNotificationSendEmailSmtpToProto converts a ConfigNotificationSendEmailSmtp object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailSmtpToProto(o *alpha.ConfigNotificationSendEmailSmtp) *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtp {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailSmtp{}
	p.SetSenderEmail(dcl.ValueOrEmptyString(o.SenderEmail))
	p.SetHost(dcl.ValueOrEmptyString(o.Host))
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	p.SetPassword(dcl.ValueOrEmptyString(o.Password))
	p.SetSecurityMode(IdentitytoolkitAlphaConfigNotificationSendEmailSmtpSecurityModeEnumToProto(o.SecurityMode))
	return p
}

// ConfigNotificationSendEmailResetPasswordTemplateToProto converts a ConfigNotificationSendEmailResetPasswordTemplate object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateToProto(o *alpha.ConfigNotificationSendEmailResetPasswordTemplate) *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitAlphaConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailVerifyEmailTemplateToProto converts a ConfigNotificationSendEmailVerifyEmailTemplate object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateToProto(o *alpha.ConfigNotificationSendEmailVerifyEmailTemplate) *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitAlphaConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailChangeEmailTemplateToProto converts a ConfigNotificationSendEmailChangeEmailTemplate object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateToProto(o *alpha.ConfigNotificationSendEmailChangeEmailTemplate) *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitAlphaConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendEmailDnsInfoToProto converts a ConfigNotificationSendEmailDnsInfo object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoToProto(o *alpha.ConfigNotificationSendEmailDnsInfo) *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfo{}
	p.SetCustomDomain(dcl.ValueOrEmptyString(o.CustomDomain))
	p.SetUseCustomDomain(dcl.ValueOrEmptyBool(o.UseCustomDomain))
	p.SetPendingCustomDomain(dcl.ValueOrEmptyString(o.PendingCustomDomain))
	p.SetCustomDomainState(IdentitytoolkitAlphaConfigNotificationSendEmailDnsInfoCustomDomainStateEnumToProto(o.CustomDomainState))
	p.SetDomainVerificationRequestTime(dcl.ValueOrEmptyString(o.DomainVerificationRequestTime))
	return p
}

// ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto converts a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateToProto(o *alpha.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) *alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	p.SetSenderLocalPart(dcl.ValueOrEmptyString(o.SenderLocalPart))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetSenderDisplayName(dcl.ValueOrEmptyString(o.SenderDisplayName))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	p.SetBodyFormat(IdentitytoolkitAlphaConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumToProto(o.BodyFormat))
	p.SetReplyTo(dcl.ValueOrEmptyString(o.ReplyTo))
	p.SetCustomized(dcl.ValueOrEmptyBool(o.Customized))
	return p
}

// ConfigNotificationSendSmsToProto converts a ConfigNotificationSendSms object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendSmsToProto(o *alpha.ConfigNotificationSendSms) *alphapb.IdentitytoolkitAlphaConfigNotificationSendSms {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendSms{}
	p.SetUseDeviceLocale(dcl.ValueOrEmptyBool(o.UseDeviceLocale))
	p.SetSmsTemplate(IdentitytoolkitAlphaConfigNotificationSendSmsSmsTemplateToProto(o.SmsTemplate))
	return p
}

// ConfigNotificationSendSmsSmsTemplateToProto converts a ConfigNotificationSendSmsSmsTemplate object to its proto representation.
func IdentitytoolkitAlphaConfigNotificationSendSmsSmsTemplateToProto(o *alpha.ConfigNotificationSendSmsSmsTemplate) *alphapb.IdentitytoolkitAlphaConfigNotificationSendSmsSmsTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigNotificationSendSmsSmsTemplate{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	return p
}

// ConfigQuotaToProto converts a ConfigQuota object to its proto representation.
func IdentitytoolkitAlphaConfigQuotaToProto(o *alpha.ConfigQuota) *alphapb.IdentitytoolkitAlphaConfigQuota {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigQuota{}
	p.SetSignUpQuotaConfig(IdentitytoolkitAlphaConfigQuotaSignUpQuotaConfigToProto(o.SignUpQuotaConfig))
	return p
}

// ConfigQuotaSignUpQuotaConfigToProto converts a ConfigQuotaSignUpQuotaConfig object to its proto representation.
func IdentitytoolkitAlphaConfigQuotaSignUpQuotaConfigToProto(o *alpha.ConfigQuotaSignUpQuotaConfig) *alphapb.IdentitytoolkitAlphaConfigQuotaSignUpQuotaConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigQuotaSignUpQuotaConfig{}
	p.SetQuota(dcl.ValueOrEmptyInt64(o.Quota))
	p.SetStartTime(dcl.ValueOrEmptyString(o.StartTime))
	p.SetQuotaDuration(dcl.ValueOrEmptyString(o.QuotaDuration))
	return p
}

// ConfigMonitoringToProto converts a ConfigMonitoring object to its proto representation.
func IdentitytoolkitAlphaConfigMonitoringToProto(o *alpha.ConfigMonitoring) *alphapb.IdentitytoolkitAlphaConfigMonitoring {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigMonitoring{}
	p.SetRequestLogging(IdentitytoolkitAlphaConfigMonitoringRequestLoggingToProto(o.RequestLogging))
	return p
}

// ConfigMonitoringRequestLoggingToProto converts a ConfigMonitoringRequestLogging object to its proto representation.
func IdentitytoolkitAlphaConfigMonitoringRequestLoggingToProto(o *alpha.ConfigMonitoringRequestLogging) *alphapb.IdentitytoolkitAlphaConfigMonitoringRequestLogging {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigMonitoringRequestLogging{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// ConfigMultiTenantToProto converts a ConfigMultiTenant object to its proto representation.
func IdentitytoolkitAlphaConfigMultiTenantToProto(o *alpha.ConfigMultiTenant) *alphapb.IdentitytoolkitAlphaConfigMultiTenant {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigMultiTenant{}
	p.SetAllowTenants(dcl.ValueOrEmptyBool(o.AllowTenants))
	p.SetDefaultTenantLocation(dcl.ValueOrEmptyString(o.DefaultTenantLocation))
	return p
}

// ConfigClientToProto converts a ConfigClient object to its proto representation.
func IdentitytoolkitAlphaConfigClientToProto(o *alpha.ConfigClient) *alphapb.IdentitytoolkitAlphaConfigClient {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigClient{}
	p.SetApiKey(dcl.ValueOrEmptyString(o.ApiKey))
	p.SetPermissions(IdentitytoolkitAlphaConfigClientPermissionsToProto(o.Permissions))
	p.SetFirebaseSubdomain(dcl.ValueOrEmptyString(o.FirebaseSubdomain))
	return p
}

// ConfigClientPermissionsToProto converts a ConfigClientPermissions object to its proto representation.
func IdentitytoolkitAlphaConfigClientPermissionsToProto(o *alpha.ConfigClientPermissions) *alphapb.IdentitytoolkitAlphaConfigClientPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigClientPermissions{}
	p.SetDisabledUserSignup(dcl.ValueOrEmptyBool(o.DisabledUserSignup))
	p.SetDisabledUserDeletion(dcl.ValueOrEmptyBool(o.DisabledUserDeletion))
	return p
}

// ConfigMfaToProto converts a ConfigMfa object to its proto representation.
func IdentitytoolkitAlphaConfigMfaToProto(o *alpha.ConfigMfa) *alphapb.IdentitytoolkitAlphaConfigMfa {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigMfa{}
	p.SetState(IdentitytoolkitAlphaConfigMfaStateEnumToProto(o.State))
	return p
}

// ConfigBlockingFunctionsToProto converts a ConfigBlockingFunctions object to its proto representation.
func IdentitytoolkitAlphaConfigBlockingFunctionsToProto(o *alpha.ConfigBlockingFunctions) *alphapb.IdentitytoolkitAlphaConfigBlockingFunctions {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigBlockingFunctions{}
	mTriggers := make(map[string]*alphapb.IdentitytoolkitAlphaConfigBlockingFunctionsTriggers, len(o.Triggers))
	for k, r := range o.Triggers {
		mTriggers[k] = IdentitytoolkitAlphaConfigBlockingFunctionsTriggersToProto(&r)
	}
	p.SetTriggers(mTriggers)
	return p
}

// ConfigBlockingFunctionsTriggersToProto converts a ConfigBlockingFunctionsTriggers object to its proto representation.
func IdentitytoolkitAlphaConfigBlockingFunctionsTriggersToProto(o *alpha.ConfigBlockingFunctionsTriggers) *alphapb.IdentitytoolkitAlphaConfigBlockingFunctionsTriggers {
	if o == nil {
		return nil
	}
	p := &alphapb.IdentitytoolkitAlphaConfigBlockingFunctionsTriggers{}
	p.SetFunctionUri(dcl.ValueOrEmptyString(o.FunctionUri))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// ConfigToProto converts a Config resource to its proto representation.
func ConfigToProto(resource *alpha.Config) *alphapb.IdentitytoolkitAlphaConfig {
	p := &alphapb.IdentitytoolkitAlphaConfig{}
	p.SetSignIn(IdentitytoolkitAlphaConfigSignInToProto(resource.SignIn))
	p.SetNotification(IdentitytoolkitAlphaConfigNotificationToProto(resource.Notification))
	p.SetQuota(IdentitytoolkitAlphaConfigQuotaToProto(resource.Quota))
	p.SetMonitoring(IdentitytoolkitAlphaConfigMonitoringToProto(resource.Monitoring))
	p.SetMultiTenant(IdentitytoolkitAlphaConfigMultiTenantToProto(resource.MultiTenant))
	p.SetSubtype(IdentitytoolkitAlphaConfigSubtypeEnumToProto(resource.Subtype))
	p.SetClient(IdentitytoolkitAlphaConfigClientToProto(resource.Client))
	p.SetMfa(IdentitytoolkitAlphaConfigMfaToProto(resource.Mfa))
	p.SetBlockingFunctions(IdentitytoolkitAlphaConfigBlockingFunctionsToProto(resource.BlockingFunctions))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sAuthorizedDomains := make([]string, len(resource.AuthorizedDomains))
	for i, r := range resource.AuthorizedDomains {
		sAuthorizedDomains[i] = r
	}
	p.SetAuthorizedDomains(sAuthorizedDomains)

	return p
}

// applyConfig handles the gRPC request by passing it to the underlying Config Apply() method.
func (s *ConfigServer) applyConfig(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIdentitytoolkitAlphaConfigRequest) (*alphapb.IdentitytoolkitAlphaConfig, error) {
	p := ProtoToConfig(request.GetResource())
	res, err := c.ApplyConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConfigToProto(res)
	return r, nil
}

// applyIdentitytoolkitAlphaConfig handles the gRPC request by passing it to the underlying Config Apply() method.
func (s *ConfigServer) ApplyIdentitytoolkitAlphaConfig(ctx context.Context, request *alphapb.ApplyIdentitytoolkitAlphaConfigRequest) (*alphapb.IdentitytoolkitAlphaConfig, error) {
	cl, err := createConfigConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyConfig(ctx, cl, request)
}

// DeleteConfig handles the gRPC request by passing it to the underlying Config Delete() method.
func (s *ConfigServer) DeleteIdentitytoolkitAlphaConfig(ctx context.Context, request *alphapb.DeleteIdentitytoolkitAlphaConfigRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for Config")

}

// ListIdentitytoolkitAlphaConfig is a no-op method because Config has no list method.
func (s *ConfigServer) ListIdentitytoolkitAlphaConfig(_ context.Context, _ *alphapb.ListIdentitytoolkitAlphaConfigRequest) (*alphapb.ListIdentitytoolkitAlphaConfigResponse, error) {
	return nil, nil
}

func createConfigConfig(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
