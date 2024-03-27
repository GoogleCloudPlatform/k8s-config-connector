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
package identitytoolkit

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/identitytoolkit"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Config struct{}

func ConfigToUnstructured(r *dclService.Config) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "identitytoolkit",
			Version: "ga",
			Type:    "Config",
		},
		Object: make(map[string]interface{}),
	}
	var rAuthorizedDomains []interface{}
	for _, rAuthorizedDomainsVal := range r.AuthorizedDomains {
		rAuthorizedDomains = append(rAuthorizedDomains, rAuthorizedDomainsVal)
	}
	u.Object["authorizedDomains"] = rAuthorizedDomains
	if r.BlockingFunctions != nil && r.BlockingFunctions != dclService.EmptyConfigBlockingFunctions {
		rBlockingFunctions := make(map[string]interface{})
		if r.BlockingFunctions.Triggers != nil {
			rBlockingFunctionsTriggers := make(map[string]interface{})
			for k, v := range r.BlockingFunctions.Triggers {
				rBlockingFunctionsTriggersMap := make(map[string]interface{})
				if v.FunctionUri != nil {
					rBlockingFunctionsTriggersMap["functionUri"] = *v.FunctionUri
				}
				if v.UpdateTime != nil {
					rBlockingFunctionsTriggersMap["updateTime"] = *v.UpdateTime
				}
				rBlockingFunctionsTriggers[k] = rBlockingFunctionsTriggersMap
			}
			rBlockingFunctions["triggers"] = rBlockingFunctionsTriggers
		}
		u.Object["blockingFunctions"] = rBlockingFunctions
	}
	if r.Client != nil && r.Client != dclService.EmptyConfigClient {
		rClient := make(map[string]interface{})
		if r.Client.ApiKey != nil {
			rClient["apiKey"] = *r.Client.ApiKey
		}
		if r.Client.FirebaseSubdomain != nil {
			rClient["firebaseSubdomain"] = *r.Client.FirebaseSubdomain
		}
		if r.Client.Permissions != nil && r.Client.Permissions != dclService.EmptyConfigClientPermissions {
			rClientPermissions := make(map[string]interface{})
			if r.Client.Permissions.DisabledUserDeletion != nil {
				rClientPermissions["disabledUserDeletion"] = *r.Client.Permissions.DisabledUserDeletion
			}
			if r.Client.Permissions.DisabledUserSignup != nil {
				rClientPermissions["disabledUserSignup"] = *r.Client.Permissions.DisabledUserSignup
			}
			rClient["permissions"] = rClientPermissions
		}
		u.Object["client"] = rClient
	}
	if r.Mfa != nil && r.Mfa != dclService.EmptyConfigMfa {
		rMfa := make(map[string]interface{})
		if r.Mfa.State != nil {
			rMfa["state"] = string(*r.Mfa.State)
		}
		u.Object["mfa"] = rMfa
	}
	if r.Monitoring != nil && r.Monitoring != dclService.EmptyConfigMonitoring {
		rMonitoring := make(map[string]interface{})
		if r.Monitoring.RequestLogging != nil && r.Monitoring.RequestLogging != dclService.EmptyConfigMonitoringRequestLogging {
			rMonitoringRequestLogging := make(map[string]interface{})
			if r.Monitoring.RequestLogging.Enabled != nil {
				rMonitoringRequestLogging["enabled"] = *r.Monitoring.RequestLogging.Enabled
			}
			rMonitoring["requestLogging"] = rMonitoringRequestLogging
		}
		u.Object["monitoring"] = rMonitoring
	}
	if r.MultiTenant != nil && r.MultiTenant != dclService.EmptyConfigMultiTenant {
		rMultiTenant := make(map[string]interface{})
		if r.MultiTenant.AllowTenants != nil {
			rMultiTenant["allowTenants"] = *r.MultiTenant.AllowTenants
		}
		if r.MultiTenant.DefaultTenantLocation != nil {
			rMultiTenant["defaultTenantLocation"] = *r.MultiTenant.DefaultTenantLocation
		}
		u.Object["multiTenant"] = rMultiTenant
	}
	if r.Notification != nil && r.Notification != dclService.EmptyConfigNotification {
		rNotification := make(map[string]interface{})
		if r.Notification.DefaultLocale != nil {
			rNotification["defaultLocale"] = *r.Notification.DefaultLocale
		}
		if r.Notification.SendEmail != nil && r.Notification.SendEmail != dclService.EmptyConfigNotificationSendEmail {
			rNotificationSendEmail := make(map[string]interface{})
			if r.Notification.SendEmail.CallbackUri != nil {
				rNotificationSendEmail["callbackUri"] = *r.Notification.SendEmail.CallbackUri
			}
			if r.Notification.SendEmail.ChangeEmailTemplate != nil && r.Notification.SendEmail.ChangeEmailTemplate != dclService.EmptyConfigNotificationSendEmailChangeEmailTemplate {
				rNotificationSendEmailChangeEmailTemplate := make(map[string]interface{})
				if r.Notification.SendEmail.ChangeEmailTemplate.Body != nil {
					rNotificationSendEmailChangeEmailTemplate["body"] = *r.Notification.SendEmail.ChangeEmailTemplate.Body
				}
				if r.Notification.SendEmail.ChangeEmailTemplate.BodyFormat != nil {
					rNotificationSendEmailChangeEmailTemplate["bodyFormat"] = string(*r.Notification.SendEmail.ChangeEmailTemplate.BodyFormat)
				}
				if r.Notification.SendEmail.ChangeEmailTemplate.Customized != nil {
					rNotificationSendEmailChangeEmailTemplate["customized"] = *r.Notification.SendEmail.ChangeEmailTemplate.Customized
				}
				if r.Notification.SendEmail.ChangeEmailTemplate.ReplyTo != nil {
					rNotificationSendEmailChangeEmailTemplate["replyTo"] = *r.Notification.SendEmail.ChangeEmailTemplate.ReplyTo
				}
				if r.Notification.SendEmail.ChangeEmailTemplate.SenderDisplayName != nil {
					rNotificationSendEmailChangeEmailTemplate["senderDisplayName"] = *r.Notification.SendEmail.ChangeEmailTemplate.SenderDisplayName
				}
				if r.Notification.SendEmail.ChangeEmailTemplate.SenderLocalPart != nil {
					rNotificationSendEmailChangeEmailTemplate["senderLocalPart"] = *r.Notification.SendEmail.ChangeEmailTemplate.SenderLocalPart
				}
				if r.Notification.SendEmail.ChangeEmailTemplate.Subject != nil {
					rNotificationSendEmailChangeEmailTemplate["subject"] = *r.Notification.SendEmail.ChangeEmailTemplate.Subject
				}
				rNotificationSendEmail["changeEmailTemplate"] = rNotificationSendEmailChangeEmailTemplate
			}
			if r.Notification.SendEmail.DnsInfo != nil && r.Notification.SendEmail.DnsInfo != dclService.EmptyConfigNotificationSendEmailDnsInfo {
				rNotificationSendEmailDnsInfo := make(map[string]interface{})
				if r.Notification.SendEmail.DnsInfo.CustomDomain != nil {
					rNotificationSendEmailDnsInfo["customDomain"] = *r.Notification.SendEmail.DnsInfo.CustomDomain
				}
				if r.Notification.SendEmail.DnsInfo.CustomDomainState != nil {
					rNotificationSendEmailDnsInfo["customDomainState"] = string(*r.Notification.SendEmail.DnsInfo.CustomDomainState)
				}
				if r.Notification.SendEmail.DnsInfo.DomainVerificationRequestTime != nil {
					rNotificationSendEmailDnsInfo["domainVerificationRequestTime"] = *r.Notification.SendEmail.DnsInfo.DomainVerificationRequestTime
				}
				if r.Notification.SendEmail.DnsInfo.PendingCustomDomain != nil {
					rNotificationSendEmailDnsInfo["pendingCustomDomain"] = *r.Notification.SendEmail.DnsInfo.PendingCustomDomain
				}
				if r.Notification.SendEmail.DnsInfo.UseCustomDomain != nil {
					rNotificationSendEmailDnsInfo["useCustomDomain"] = *r.Notification.SendEmail.DnsInfo.UseCustomDomain
				}
				rNotificationSendEmail["dnsInfo"] = rNotificationSendEmailDnsInfo
			}
			if r.Notification.SendEmail.Method != nil {
				rNotificationSendEmail["method"] = string(*r.Notification.SendEmail.Method)
			}
			if r.Notification.SendEmail.ResetPasswordTemplate != nil && r.Notification.SendEmail.ResetPasswordTemplate != dclService.EmptyConfigNotificationSendEmailResetPasswordTemplate {
				rNotificationSendEmailResetPasswordTemplate := make(map[string]interface{})
				if r.Notification.SendEmail.ResetPasswordTemplate.Body != nil {
					rNotificationSendEmailResetPasswordTemplate["body"] = *r.Notification.SendEmail.ResetPasswordTemplate.Body
				}
				if r.Notification.SendEmail.ResetPasswordTemplate.BodyFormat != nil {
					rNotificationSendEmailResetPasswordTemplate["bodyFormat"] = string(*r.Notification.SendEmail.ResetPasswordTemplate.BodyFormat)
				}
				if r.Notification.SendEmail.ResetPasswordTemplate.Customized != nil {
					rNotificationSendEmailResetPasswordTemplate["customized"] = *r.Notification.SendEmail.ResetPasswordTemplate.Customized
				}
				if r.Notification.SendEmail.ResetPasswordTemplate.ReplyTo != nil {
					rNotificationSendEmailResetPasswordTemplate["replyTo"] = *r.Notification.SendEmail.ResetPasswordTemplate.ReplyTo
				}
				if r.Notification.SendEmail.ResetPasswordTemplate.SenderDisplayName != nil {
					rNotificationSendEmailResetPasswordTemplate["senderDisplayName"] = *r.Notification.SendEmail.ResetPasswordTemplate.SenderDisplayName
				}
				if r.Notification.SendEmail.ResetPasswordTemplate.SenderLocalPart != nil {
					rNotificationSendEmailResetPasswordTemplate["senderLocalPart"] = *r.Notification.SendEmail.ResetPasswordTemplate.SenderLocalPart
				}
				if r.Notification.SendEmail.ResetPasswordTemplate.Subject != nil {
					rNotificationSendEmailResetPasswordTemplate["subject"] = *r.Notification.SendEmail.ResetPasswordTemplate.Subject
				}
				rNotificationSendEmail["resetPasswordTemplate"] = rNotificationSendEmailResetPasswordTemplate
			}
			if r.Notification.SendEmail.RevertSecondFactorAdditionTemplate != nil && r.Notification.SendEmail.RevertSecondFactorAdditionTemplate != dclService.EmptyConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
				rNotificationSendEmailRevertSecondFactorAdditionTemplate := make(map[string]interface{})
				if r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Body != nil {
					rNotificationSendEmailRevertSecondFactorAdditionTemplate["body"] = *r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Body
				}
				if r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.BodyFormat != nil {
					rNotificationSendEmailRevertSecondFactorAdditionTemplate["bodyFormat"] = string(*r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.BodyFormat)
				}
				if r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Customized != nil {
					rNotificationSendEmailRevertSecondFactorAdditionTemplate["customized"] = *r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Customized
				}
				if r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.ReplyTo != nil {
					rNotificationSendEmailRevertSecondFactorAdditionTemplate["replyTo"] = *r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.ReplyTo
				}
				if r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.SenderDisplayName != nil {
					rNotificationSendEmailRevertSecondFactorAdditionTemplate["senderDisplayName"] = *r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.SenderDisplayName
				}
				if r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.SenderLocalPart != nil {
					rNotificationSendEmailRevertSecondFactorAdditionTemplate["senderLocalPart"] = *r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.SenderLocalPart
				}
				if r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Subject != nil {
					rNotificationSendEmailRevertSecondFactorAdditionTemplate["subject"] = *r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Subject
				}
				rNotificationSendEmail["revertSecondFactorAdditionTemplate"] = rNotificationSendEmailRevertSecondFactorAdditionTemplate
			}
			if r.Notification.SendEmail.Smtp != nil && r.Notification.SendEmail.Smtp != dclService.EmptyConfigNotificationSendEmailSmtp {
				rNotificationSendEmailSmtp := make(map[string]interface{})
				if r.Notification.SendEmail.Smtp.Host != nil {
					rNotificationSendEmailSmtp["host"] = *r.Notification.SendEmail.Smtp.Host
				}
				if r.Notification.SendEmail.Smtp.Password != nil {
					rNotificationSendEmailSmtp["password"] = *r.Notification.SendEmail.Smtp.Password
				}
				if r.Notification.SendEmail.Smtp.Port != nil {
					rNotificationSendEmailSmtp["port"] = *r.Notification.SendEmail.Smtp.Port
				}
				if r.Notification.SendEmail.Smtp.SecurityMode != nil {
					rNotificationSendEmailSmtp["securityMode"] = string(*r.Notification.SendEmail.Smtp.SecurityMode)
				}
				if r.Notification.SendEmail.Smtp.SenderEmail != nil {
					rNotificationSendEmailSmtp["senderEmail"] = *r.Notification.SendEmail.Smtp.SenderEmail
				}
				if r.Notification.SendEmail.Smtp.Username != nil {
					rNotificationSendEmailSmtp["username"] = *r.Notification.SendEmail.Smtp.Username
				}
				rNotificationSendEmail["smtp"] = rNotificationSendEmailSmtp
			}
			if r.Notification.SendEmail.VerifyEmailTemplate != nil && r.Notification.SendEmail.VerifyEmailTemplate != dclService.EmptyConfigNotificationSendEmailVerifyEmailTemplate {
				rNotificationSendEmailVerifyEmailTemplate := make(map[string]interface{})
				if r.Notification.SendEmail.VerifyEmailTemplate.Body != nil {
					rNotificationSendEmailVerifyEmailTemplate["body"] = *r.Notification.SendEmail.VerifyEmailTemplate.Body
				}
				if r.Notification.SendEmail.VerifyEmailTemplate.BodyFormat != nil {
					rNotificationSendEmailVerifyEmailTemplate["bodyFormat"] = string(*r.Notification.SendEmail.VerifyEmailTemplate.BodyFormat)
				}
				if r.Notification.SendEmail.VerifyEmailTemplate.Customized != nil {
					rNotificationSendEmailVerifyEmailTemplate["customized"] = *r.Notification.SendEmail.VerifyEmailTemplate.Customized
				}
				if r.Notification.SendEmail.VerifyEmailTemplate.ReplyTo != nil {
					rNotificationSendEmailVerifyEmailTemplate["replyTo"] = *r.Notification.SendEmail.VerifyEmailTemplate.ReplyTo
				}
				if r.Notification.SendEmail.VerifyEmailTemplate.SenderDisplayName != nil {
					rNotificationSendEmailVerifyEmailTemplate["senderDisplayName"] = *r.Notification.SendEmail.VerifyEmailTemplate.SenderDisplayName
				}
				if r.Notification.SendEmail.VerifyEmailTemplate.SenderLocalPart != nil {
					rNotificationSendEmailVerifyEmailTemplate["senderLocalPart"] = *r.Notification.SendEmail.VerifyEmailTemplate.SenderLocalPart
				}
				if r.Notification.SendEmail.VerifyEmailTemplate.Subject != nil {
					rNotificationSendEmailVerifyEmailTemplate["subject"] = *r.Notification.SendEmail.VerifyEmailTemplate.Subject
				}
				rNotificationSendEmail["verifyEmailTemplate"] = rNotificationSendEmailVerifyEmailTemplate
			}
			rNotification["sendEmail"] = rNotificationSendEmail
		}
		if r.Notification.SendSms != nil && r.Notification.SendSms != dclService.EmptyConfigNotificationSendSms {
			rNotificationSendSms := make(map[string]interface{})
			if r.Notification.SendSms.SmsTemplate != nil && r.Notification.SendSms.SmsTemplate != dclService.EmptyConfigNotificationSendSmsSmsTemplate {
				rNotificationSendSmsSmsTemplate := make(map[string]interface{})
				if r.Notification.SendSms.SmsTemplate.Content != nil {
					rNotificationSendSmsSmsTemplate["content"] = *r.Notification.SendSms.SmsTemplate.Content
				}
				rNotificationSendSms["smsTemplate"] = rNotificationSendSmsSmsTemplate
			}
			if r.Notification.SendSms.UseDeviceLocale != nil {
				rNotificationSendSms["useDeviceLocale"] = *r.Notification.SendSms.UseDeviceLocale
			}
			rNotification["sendSms"] = rNotificationSendSms
		}
		u.Object["notification"] = rNotification
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Quota != nil && r.Quota != dclService.EmptyConfigQuota {
		rQuota := make(map[string]interface{})
		if r.Quota.SignUpQuotaConfig != nil && r.Quota.SignUpQuotaConfig != dclService.EmptyConfigQuotaSignUpQuotaConfig {
			rQuotaSignUpQuotaConfig := make(map[string]interface{})
			if r.Quota.SignUpQuotaConfig.Quota != nil {
				rQuotaSignUpQuotaConfig["quota"] = *r.Quota.SignUpQuotaConfig.Quota
			}
			if r.Quota.SignUpQuotaConfig.QuotaDuration != nil {
				rQuotaSignUpQuotaConfig["quotaDuration"] = *r.Quota.SignUpQuotaConfig.QuotaDuration
			}
			if r.Quota.SignUpQuotaConfig.StartTime != nil {
				rQuotaSignUpQuotaConfig["startTime"] = *r.Quota.SignUpQuotaConfig.StartTime
			}
			rQuota["signUpQuotaConfig"] = rQuotaSignUpQuotaConfig
		}
		u.Object["quota"] = rQuota
	}
	if r.SignIn != nil && r.SignIn != dclService.EmptyConfigSignIn {
		rSignIn := make(map[string]interface{})
		if r.SignIn.AllowDuplicateEmails != nil {
			rSignIn["allowDuplicateEmails"] = *r.SignIn.AllowDuplicateEmails
		}
		if r.SignIn.Anonymous != nil && r.SignIn.Anonymous != dclService.EmptyConfigSignInAnonymous {
			rSignInAnonymous := make(map[string]interface{})
			if r.SignIn.Anonymous.Enabled != nil {
				rSignInAnonymous["enabled"] = *r.SignIn.Anonymous.Enabled
			}
			rSignIn["anonymous"] = rSignInAnonymous
		}
		if r.SignIn.Email != nil && r.SignIn.Email != dclService.EmptyConfigSignInEmail {
			rSignInEmail := make(map[string]interface{})
			if r.SignIn.Email.Enabled != nil {
				rSignInEmail["enabled"] = *r.SignIn.Email.Enabled
			}
			if r.SignIn.Email.HashConfig != nil && r.SignIn.Email.HashConfig != dclService.EmptyConfigSignInEmailHashConfig {
				rSignInEmailHashConfig := make(map[string]interface{})
				if r.SignIn.Email.HashConfig.Algorithm != nil {
					rSignInEmailHashConfig["algorithm"] = string(*r.SignIn.Email.HashConfig.Algorithm)
				}
				if r.SignIn.Email.HashConfig.MemoryCost != nil {
					rSignInEmailHashConfig["memoryCost"] = *r.SignIn.Email.HashConfig.MemoryCost
				}
				if r.SignIn.Email.HashConfig.Rounds != nil {
					rSignInEmailHashConfig["rounds"] = *r.SignIn.Email.HashConfig.Rounds
				}
				if r.SignIn.Email.HashConfig.SaltSeparator != nil {
					rSignInEmailHashConfig["saltSeparator"] = *r.SignIn.Email.HashConfig.SaltSeparator
				}
				if r.SignIn.Email.HashConfig.SignerKey != nil {
					rSignInEmailHashConfig["signerKey"] = *r.SignIn.Email.HashConfig.SignerKey
				}
				rSignInEmail["hashConfig"] = rSignInEmailHashConfig
			}
			if r.SignIn.Email.PasswordRequired != nil {
				rSignInEmail["passwordRequired"] = *r.SignIn.Email.PasswordRequired
			}
			rSignIn["email"] = rSignInEmail
		}
		if r.SignIn.HashConfig != nil && r.SignIn.HashConfig != dclService.EmptyConfigSignInHashConfig {
			rSignInHashConfig := make(map[string]interface{})
			if r.SignIn.HashConfig.Algorithm != nil {
				rSignInHashConfig["algorithm"] = string(*r.SignIn.HashConfig.Algorithm)
			}
			if r.SignIn.HashConfig.MemoryCost != nil {
				rSignInHashConfig["memoryCost"] = *r.SignIn.HashConfig.MemoryCost
			}
			if r.SignIn.HashConfig.Rounds != nil {
				rSignInHashConfig["rounds"] = *r.SignIn.HashConfig.Rounds
			}
			if r.SignIn.HashConfig.SaltSeparator != nil {
				rSignInHashConfig["saltSeparator"] = *r.SignIn.HashConfig.SaltSeparator
			}
			if r.SignIn.HashConfig.SignerKey != nil {
				rSignInHashConfig["signerKey"] = *r.SignIn.HashConfig.SignerKey
			}
			rSignIn["hashConfig"] = rSignInHashConfig
		}
		if r.SignIn.PhoneNumber != nil && r.SignIn.PhoneNumber != dclService.EmptyConfigSignInPhoneNumber {
			rSignInPhoneNumber := make(map[string]interface{})
			if r.SignIn.PhoneNumber.Enabled != nil {
				rSignInPhoneNumber["enabled"] = *r.SignIn.PhoneNumber.Enabled
			}
			if r.SignIn.PhoneNumber.TestPhoneNumbers != nil {
				rSignInPhoneNumberTestPhoneNumbers := make(map[string]interface{})
				for k, v := range r.SignIn.PhoneNumber.TestPhoneNumbers {
					rSignInPhoneNumberTestPhoneNumbers[k] = v
				}
				rSignInPhoneNumber["testPhoneNumbers"] = rSignInPhoneNumberTestPhoneNumbers
			}
			rSignIn["phoneNumber"] = rSignInPhoneNumber
		}
		u.Object["signIn"] = rSignIn
	}
	if r.Subtype != nil {
		u.Object["subtype"] = string(*r.Subtype)
	}
	return u
}

func UnstructuredToConfig(u *unstructured.Resource) (*dclService.Config, error) {
	r := &dclService.Config{}
	if _, ok := u.Object["authorizedDomains"]; ok {
		if s, ok := u.Object["authorizedDomains"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.AuthorizedDomains = append(r.AuthorizedDomains, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.AuthorizedDomains: expected []interface{}")
		}
	}
	if _, ok := u.Object["blockingFunctions"]; ok {
		if rBlockingFunctions, ok := u.Object["blockingFunctions"].(map[string]interface{}); ok {
			r.BlockingFunctions = &dclService.ConfigBlockingFunctions{}
			if _, ok := rBlockingFunctions["triggers"]; ok {
				if rBlockingFunctionsTriggers, ok := rBlockingFunctions["triggers"].(map[string]interface{}); ok {
					m := make(map[string]dclService.ConfigBlockingFunctionsTriggers)
					for k, v := range rBlockingFunctionsTriggers {
						if objval, ok := v.(map[string]interface{}); ok {
							var rBlockingFunctionsTriggersObj dclService.ConfigBlockingFunctionsTriggers
							if _, ok := objval["functionUri"]; ok {
								if s, ok := objval["functionUri"].(string); ok {
									rBlockingFunctionsTriggersObj.FunctionUri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rBlockingFunctionsTriggersObj.FunctionUri: expected string")
								}
							}
							if _, ok := objval["updateTime"]; ok {
								if s, ok := objval["updateTime"].(string); ok {
									rBlockingFunctionsTriggersObj.UpdateTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rBlockingFunctionsTriggersObj.UpdateTime: expected string")
								}
							}
							m[k] = rBlockingFunctionsTriggersObj
						} else {
							return nil, fmt.Errorf("r.BlockingFunctions.Triggers: expected map[string]interface{}")
						}
					}
					r.BlockingFunctions.Triggers = m
				} else {
					return nil, fmt.Errorf("r.BlockingFunctions.Triggers: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.BlockingFunctions: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["client"]; ok {
		if rClient, ok := u.Object["client"].(map[string]interface{}); ok {
			r.Client = &dclService.ConfigClient{}
			if _, ok := rClient["apiKey"]; ok {
				if s, ok := rClient["apiKey"].(string); ok {
					r.Client.ApiKey = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Client.ApiKey: expected string")
				}
			}
			if _, ok := rClient["firebaseSubdomain"]; ok {
				if s, ok := rClient["firebaseSubdomain"].(string); ok {
					r.Client.FirebaseSubdomain = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Client.FirebaseSubdomain: expected string")
				}
			}
			if _, ok := rClient["permissions"]; ok {
				if rClientPermissions, ok := rClient["permissions"].(map[string]interface{}); ok {
					r.Client.Permissions = &dclService.ConfigClientPermissions{}
					if _, ok := rClientPermissions["disabledUserDeletion"]; ok {
						if b, ok := rClientPermissions["disabledUserDeletion"].(bool); ok {
							r.Client.Permissions.DisabledUserDeletion = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Client.Permissions.DisabledUserDeletion: expected bool")
						}
					}
					if _, ok := rClientPermissions["disabledUserSignup"]; ok {
						if b, ok := rClientPermissions["disabledUserSignup"].(bool); ok {
							r.Client.Permissions.DisabledUserSignup = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Client.Permissions.DisabledUserSignup: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Client.Permissions: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Client: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["mfa"]; ok {
		if rMfa, ok := u.Object["mfa"].(map[string]interface{}); ok {
			r.Mfa = &dclService.ConfigMfa{}
			if _, ok := rMfa["state"]; ok {
				if s, ok := rMfa["state"].(string); ok {
					r.Mfa.State = dclService.ConfigMfaStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Mfa.State: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Mfa: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["monitoring"]; ok {
		if rMonitoring, ok := u.Object["monitoring"].(map[string]interface{}); ok {
			r.Monitoring = &dclService.ConfigMonitoring{}
			if _, ok := rMonitoring["requestLogging"]; ok {
				if rMonitoringRequestLogging, ok := rMonitoring["requestLogging"].(map[string]interface{}); ok {
					r.Monitoring.RequestLogging = &dclService.ConfigMonitoringRequestLogging{}
					if _, ok := rMonitoringRequestLogging["enabled"]; ok {
						if b, ok := rMonitoringRequestLogging["enabled"].(bool); ok {
							r.Monitoring.RequestLogging.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Monitoring.RequestLogging.Enabled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Monitoring.RequestLogging: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Monitoring: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["multiTenant"]; ok {
		if rMultiTenant, ok := u.Object["multiTenant"].(map[string]interface{}); ok {
			r.MultiTenant = &dclService.ConfigMultiTenant{}
			if _, ok := rMultiTenant["allowTenants"]; ok {
				if b, ok := rMultiTenant["allowTenants"].(bool); ok {
					r.MultiTenant.AllowTenants = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.MultiTenant.AllowTenants: expected bool")
				}
			}
			if _, ok := rMultiTenant["defaultTenantLocation"]; ok {
				if s, ok := rMultiTenant["defaultTenantLocation"].(string); ok {
					r.MultiTenant.DefaultTenantLocation = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.MultiTenant.DefaultTenantLocation: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MultiTenant: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["notification"]; ok {
		if rNotification, ok := u.Object["notification"].(map[string]interface{}); ok {
			r.Notification = &dclService.ConfigNotification{}
			if _, ok := rNotification["defaultLocale"]; ok {
				if s, ok := rNotification["defaultLocale"].(string); ok {
					r.Notification.DefaultLocale = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Notification.DefaultLocale: expected string")
				}
			}
			if _, ok := rNotification["sendEmail"]; ok {
				if rNotificationSendEmail, ok := rNotification["sendEmail"].(map[string]interface{}); ok {
					r.Notification.SendEmail = &dclService.ConfigNotificationSendEmail{}
					if _, ok := rNotificationSendEmail["callbackUri"]; ok {
						if s, ok := rNotificationSendEmail["callbackUri"].(string); ok {
							r.Notification.SendEmail.CallbackUri = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Notification.SendEmail.CallbackUri: expected string")
						}
					}
					if _, ok := rNotificationSendEmail["changeEmailTemplate"]; ok {
						if rNotificationSendEmailChangeEmailTemplate, ok := rNotificationSendEmail["changeEmailTemplate"].(map[string]interface{}); ok {
							r.Notification.SendEmail.ChangeEmailTemplate = &dclService.ConfigNotificationSendEmailChangeEmailTemplate{}
							if _, ok := rNotificationSendEmailChangeEmailTemplate["body"]; ok {
								if s, ok := rNotificationSendEmailChangeEmailTemplate["body"].(string); ok {
									r.Notification.SendEmail.ChangeEmailTemplate.Body = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ChangeEmailTemplate.Body: expected string")
								}
							}
							if _, ok := rNotificationSendEmailChangeEmailTemplate["bodyFormat"]; ok {
								if s, ok := rNotificationSendEmailChangeEmailTemplate["bodyFormat"].(string); ok {
									r.Notification.SendEmail.ChangeEmailTemplate.BodyFormat = dclService.ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ChangeEmailTemplate.BodyFormat: expected string")
								}
							}
							if _, ok := rNotificationSendEmailChangeEmailTemplate["customized"]; ok {
								if b, ok := rNotificationSendEmailChangeEmailTemplate["customized"].(bool); ok {
									r.Notification.SendEmail.ChangeEmailTemplate.Customized = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ChangeEmailTemplate.Customized: expected bool")
								}
							}
							if _, ok := rNotificationSendEmailChangeEmailTemplate["replyTo"]; ok {
								if s, ok := rNotificationSendEmailChangeEmailTemplate["replyTo"].(string); ok {
									r.Notification.SendEmail.ChangeEmailTemplate.ReplyTo = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ChangeEmailTemplate.ReplyTo: expected string")
								}
							}
							if _, ok := rNotificationSendEmailChangeEmailTemplate["senderDisplayName"]; ok {
								if s, ok := rNotificationSendEmailChangeEmailTemplate["senderDisplayName"].(string); ok {
									r.Notification.SendEmail.ChangeEmailTemplate.SenderDisplayName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ChangeEmailTemplate.SenderDisplayName: expected string")
								}
							}
							if _, ok := rNotificationSendEmailChangeEmailTemplate["senderLocalPart"]; ok {
								if s, ok := rNotificationSendEmailChangeEmailTemplate["senderLocalPart"].(string); ok {
									r.Notification.SendEmail.ChangeEmailTemplate.SenderLocalPart = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ChangeEmailTemplate.SenderLocalPart: expected string")
								}
							}
							if _, ok := rNotificationSendEmailChangeEmailTemplate["subject"]; ok {
								if s, ok := rNotificationSendEmailChangeEmailTemplate["subject"].(string); ok {
									r.Notification.SendEmail.ChangeEmailTemplate.Subject = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ChangeEmailTemplate.Subject: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Notification.SendEmail.ChangeEmailTemplate: expected map[string]interface{}")
						}
					}
					if _, ok := rNotificationSendEmail["dnsInfo"]; ok {
						if rNotificationSendEmailDnsInfo, ok := rNotificationSendEmail["dnsInfo"].(map[string]interface{}); ok {
							r.Notification.SendEmail.DnsInfo = &dclService.ConfigNotificationSendEmailDnsInfo{}
							if _, ok := rNotificationSendEmailDnsInfo["customDomain"]; ok {
								if s, ok := rNotificationSendEmailDnsInfo["customDomain"].(string); ok {
									r.Notification.SendEmail.DnsInfo.CustomDomain = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.DnsInfo.CustomDomain: expected string")
								}
							}
							if _, ok := rNotificationSendEmailDnsInfo["customDomainState"]; ok {
								if s, ok := rNotificationSendEmailDnsInfo["customDomainState"].(string); ok {
									r.Notification.SendEmail.DnsInfo.CustomDomainState = dclService.ConfigNotificationSendEmailDnsInfoCustomDomainStateEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.DnsInfo.CustomDomainState: expected string")
								}
							}
							if _, ok := rNotificationSendEmailDnsInfo["domainVerificationRequestTime"]; ok {
								if s, ok := rNotificationSendEmailDnsInfo["domainVerificationRequestTime"].(string); ok {
									r.Notification.SendEmail.DnsInfo.DomainVerificationRequestTime = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.DnsInfo.DomainVerificationRequestTime: expected string")
								}
							}
							if _, ok := rNotificationSendEmailDnsInfo["pendingCustomDomain"]; ok {
								if s, ok := rNotificationSendEmailDnsInfo["pendingCustomDomain"].(string); ok {
									r.Notification.SendEmail.DnsInfo.PendingCustomDomain = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.DnsInfo.PendingCustomDomain: expected string")
								}
							}
							if _, ok := rNotificationSendEmailDnsInfo["useCustomDomain"]; ok {
								if b, ok := rNotificationSendEmailDnsInfo["useCustomDomain"].(bool); ok {
									r.Notification.SendEmail.DnsInfo.UseCustomDomain = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.DnsInfo.UseCustomDomain: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Notification.SendEmail.DnsInfo: expected map[string]interface{}")
						}
					}
					if _, ok := rNotificationSendEmail["method"]; ok {
						if s, ok := rNotificationSendEmail["method"].(string); ok {
							r.Notification.SendEmail.Method = dclService.ConfigNotificationSendEmailMethodEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Notification.SendEmail.Method: expected string")
						}
					}
					if _, ok := rNotificationSendEmail["resetPasswordTemplate"]; ok {
						if rNotificationSendEmailResetPasswordTemplate, ok := rNotificationSendEmail["resetPasswordTemplate"].(map[string]interface{}); ok {
							r.Notification.SendEmail.ResetPasswordTemplate = &dclService.ConfigNotificationSendEmailResetPasswordTemplate{}
							if _, ok := rNotificationSendEmailResetPasswordTemplate["body"]; ok {
								if s, ok := rNotificationSendEmailResetPasswordTemplate["body"].(string); ok {
									r.Notification.SendEmail.ResetPasswordTemplate.Body = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ResetPasswordTemplate.Body: expected string")
								}
							}
							if _, ok := rNotificationSendEmailResetPasswordTemplate["bodyFormat"]; ok {
								if s, ok := rNotificationSendEmailResetPasswordTemplate["bodyFormat"].(string); ok {
									r.Notification.SendEmail.ResetPasswordTemplate.BodyFormat = dclService.ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ResetPasswordTemplate.BodyFormat: expected string")
								}
							}
							if _, ok := rNotificationSendEmailResetPasswordTemplate["customized"]; ok {
								if b, ok := rNotificationSendEmailResetPasswordTemplate["customized"].(bool); ok {
									r.Notification.SendEmail.ResetPasswordTemplate.Customized = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ResetPasswordTemplate.Customized: expected bool")
								}
							}
							if _, ok := rNotificationSendEmailResetPasswordTemplate["replyTo"]; ok {
								if s, ok := rNotificationSendEmailResetPasswordTemplate["replyTo"].(string); ok {
									r.Notification.SendEmail.ResetPasswordTemplate.ReplyTo = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ResetPasswordTemplate.ReplyTo: expected string")
								}
							}
							if _, ok := rNotificationSendEmailResetPasswordTemplate["senderDisplayName"]; ok {
								if s, ok := rNotificationSendEmailResetPasswordTemplate["senderDisplayName"].(string); ok {
									r.Notification.SendEmail.ResetPasswordTemplate.SenderDisplayName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ResetPasswordTemplate.SenderDisplayName: expected string")
								}
							}
							if _, ok := rNotificationSendEmailResetPasswordTemplate["senderLocalPart"]; ok {
								if s, ok := rNotificationSendEmailResetPasswordTemplate["senderLocalPart"].(string); ok {
									r.Notification.SendEmail.ResetPasswordTemplate.SenderLocalPart = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ResetPasswordTemplate.SenderLocalPart: expected string")
								}
							}
							if _, ok := rNotificationSendEmailResetPasswordTemplate["subject"]; ok {
								if s, ok := rNotificationSendEmailResetPasswordTemplate["subject"].(string); ok {
									r.Notification.SendEmail.ResetPasswordTemplate.Subject = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.ResetPasswordTemplate.Subject: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Notification.SendEmail.ResetPasswordTemplate: expected map[string]interface{}")
						}
					}
					if _, ok := rNotificationSendEmail["revertSecondFactorAdditionTemplate"]; ok {
						if rNotificationSendEmailRevertSecondFactorAdditionTemplate, ok := rNotificationSendEmail["revertSecondFactorAdditionTemplate"].(map[string]interface{}); ok {
							r.Notification.SendEmail.RevertSecondFactorAdditionTemplate = &dclService.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
							if _, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["body"]; ok {
								if s, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["body"].(string); ok {
									r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Body = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Body: expected string")
								}
							}
							if _, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["bodyFormat"]; ok {
								if s, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["bodyFormat"].(string); ok {
									r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.BodyFormat = dclService.ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.BodyFormat: expected string")
								}
							}
							if _, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["customized"]; ok {
								if b, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["customized"].(bool); ok {
									r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Customized = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Customized: expected bool")
								}
							}
							if _, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["replyTo"]; ok {
								if s, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["replyTo"].(string); ok {
									r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.ReplyTo = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.ReplyTo: expected string")
								}
							}
							if _, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["senderDisplayName"]; ok {
								if s, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["senderDisplayName"].(string); ok {
									r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.SenderDisplayName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.SenderDisplayName: expected string")
								}
							}
							if _, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["senderLocalPart"]; ok {
								if s, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["senderLocalPart"].(string); ok {
									r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.SenderLocalPart = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.SenderLocalPart: expected string")
								}
							}
							if _, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["subject"]; ok {
								if s, ok := rNotificationSendEmailRevertSecondFactorAdditionTemplate["subject"].(string); ok {
									r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Subject = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.RevertSecondFactorAdditionTemplate.Subject: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Notification.SendEmail.RevertSecondFactorAdditionTemplate: expected map[string]interface{}")
						}
					}
					if _, ok := rNotificationSendEmail["smtp"]; ok {
						if rNotificationSendEmailSmtp, ok := rNotificationSendEmail["smtp"].(map[string]interface{}); ok {
							r.Notification.SendEmail.Smtp = &dclService.ConfigNotificationSendEmailSmtp{}
							if _, ok := rNotificationSendEmailSmtp["host"]; ok {
								if s, ok := rNotificationSendEmailSmtp["host"].(string); ok {
									r.Notification.SendEmail.Smtp.Host = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.Smtp.Host: expected string")
								}
							}
							if _, ok := rNotificationSendEmailSmtp["password"]; ok {
								if s, ok := rNotificationSendEmailSmtp["password"].(string); ok {
									r.Notification.SendEmail.Smtp.Password = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.Smtp.Password: expected string")
								}
							}
							if _, ok := rNotificationSendEmailSmtp["port"]; ok {
								if i, ok := rNotificationSendEmailSmtp["port"].(int64); ok {
									r.Notification.SendEmail.Smtp.Port = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.Smtp.Port: expected int64")
								}
							}
							if _, ok := rNotificationSendEmailSmtp["securityMode"]; ok {
								if s, ok := rNotificationSendEmailSmtp["securityMode"].(string); ok {
									r.Notification.SendEmail.Smtp.SecurityMode = dclService.ConfigNotificationSendEmailSmtpSecurityModeEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.Smtp.SecurityMode: expected string")
								}
							}
							if _, ok := rNotificationSendEmailSmtp["senderEmail"]; ok {
								if s, ok := rNotificationSendEmailSmtp["senderEmail"].(string); ok {
									r.Notification.SendEmail.Smtp.SenderEmail = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.Smtp.SenderEmail: expected string")
								}
							}
							if _, ok := rNotificationSendEmailSmtp["username"]; ok {
								if s, ok := rNotificationSendEmailSmtp["username"].(string); ok {
									r.Notification.SendEmail.Smtp.Username = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.Smtp.Username: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Notification.SendEmail.Smtp: expected map[string]interface{}")
						}
					}
					if _, ok := rNotificationSendEmail["verifyEmailTemplate"]; ok {
						if rNotificationSendEmailVerifyEmailTemplate, ok := rNotificationSendEmail["verifyEmailTemplate"].(map[string]interface{}); ok {
							r.Notification.SendEmail.VerifyEmailTemplate = &dclService.ConfigNotificationSendEmailVerifyEmailTemplate{}
							if _, ok := rNotificationSendEmailVerifyEmailTemplate["body"]; ok {
								if s, ok := rNotificationSendEmailVerifyEmailTemplate["body"].(string); ok {
									r.Notification.SendEmail.VerifyEmailTemplate.Body = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.VerifyEmailTemplate.Body: expected string")
								}
							}
							if _, ok := rNotificationSendEmailVerifyEmailTemplate["bodyFormat"]; ok {
								if s, ok := rNotificationSendEmailVerifyEmailTemplate["bodyFormat"].(string); ok {
									r.Notification.SendEmail.VerifyEmailTemplate.BodyFormat = dclService.ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.VerifyEmailTemplate.BodyFormat: expected string")
								}
							}
							if _, ok := rNotificationSendEmailVerifyEmailTemplate["customized"]; ok {
								if b, ok := rNotificationSendEmailVerifyEmailTemplate["customized"].(bool); ok {
									r.Notification.SendEmail.VerifyEmailTemplate.Customized = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.VerifyEmailTemplate.Customized: expected bool")
								}
							}
							if _, ok := rNotificationSendEmailVerifyEmailTemplate["replyTo"]; ok {
								if s, ok := rNotificationSendEmailVerifyEmailTemplate["replyTo"].(string); ok {
									r.Notification.SendEmail.VerifyEmailTemplate.ReplyTo = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.VerifyEmailTemplate.ReplyTo: expected string")
								}
							}
							if _, ok := rNotificationSendEmailVerifyEmailTemplate["senderDisplayName"]; ok {
								if s, ok := rNotificationSendEmailVerifyEmailTemplate["senderDisplayName"].(string); ok {
									r.Notification.SendEmail.VerifyEmailTemplate.SenderDisplayName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.VerifyEmailTemplate.SenderDisplayName: expected string")
								}
							}
							if _, ok := rNotificationSendEmailVerifyEmailTemplate["senderLocalPart"]; ok {
								if s, ok := rNotificationSendEmailVerifyEmailTemplate["senderLocalPart"].(string); ok {
									r.Notification.SendEmail.VerifyEmailTemplate.SenderLocalPart = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.VerifyEmailTemplate.SenderLocalPart: expected string")
								}
							}
							if _, ok := rNotificationSendEmailVerifyEmailTemplate["subject"]; ok {
								if s, ok := rNotificationSendEmailVerifyEmailTemplate["subject"].(string); ok {
									r.Notification.SendEmail.VerifyEmailTemplate.Subject = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendEmail.VerifyEmailTemplate.Subject: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Notification.SendEmail.VerifyEmailTemplate: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Notification.SendEmail: expected map[string]interface{}")
				}
			}
			if _, ok := rNotification["sendSms"]; ok {
				if rNotificationSendSms, ok := rNotification["sendSms"].(map[string]interface{}); ok {
					r.Notification.SendSms = &dclService.ConfigNotificationSendSms{}
					if _, ok := rNotificationSendSms["smsTemplate"]; ok {
						if rNotificationSendSmsSmsTemplate, ok := rNotificationSendSms["smsTemplate"].(map[string]interface{}); ok {
							r.Notification.SendSms.SmsTemplate = &dclService.ConfigNotificationSendSmsSmsTemplate{}
							if _, ok := rNotificationSendSmsSmsTemplate["content"]; ok {
								if s, ok := rNotificationSendSmsSmsTemplate["content"].(string); ok {
									r.Notification.SendSms.SmsTemplate.Content = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Notification.SendSms.SmsTemplate.Content: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Notification.SendSms.SmsTemplate: expected map[string]interface{}")
						}
					}
					if _, ok := rNotificationSendSms["useDeviceLocale"]; ok {
						if b, ok := rNotificationSendSms["useDeviceLocale"].(bool); ok {
							r.Notification.SendSms.UseDeviceLocale = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Notification.SendSms.UseDeviceLocale: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Notification.SendSms: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Notification: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["quota"]; ok {
		if rQuota, ok := u.Object["quota"].(map[string]interface{}); ok {
			r.Quota = &dclService.ConfigQuota{}
			if _, ok := rQuota["signUpQuotaConfig"]; ok {
				if rQuotaSignUpQuotaConfig, ok := rQuota["signUpQuotaConfig"].(map[string]interface{}); ok {
					r.Quota.SignUpQuotaConfig = &dclService.ConfigQuotaSignUpQuotaConfig{}
					if _, ok := rQuotaSignUpQuotaConfig["quota"]; ok {
						if i, ok := rQuotaSignUpQuotaConfig["quota"].(int64); ok {
							r.Quota.SignUpQuotaConfig.Quota = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Quota.SignUpQuotaConfig.Quota: expected int64")
						}
					}
					if _, ok := rQuotaSignUpQuotaConfig["quotaDuration"]; ok {
						if s, ok := rQuotaSignUpQuotaConfig["quotaDuration"].(string); ok {
							r.Quota.SignUpQuotaConfig.QuotaDuration = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Quota.SignUpQuotaConfig.QuotaDuration: expected string")
						}
					}
					if _, ok := rQuotaSignUpQuotaConfig["startTime"]; ok {
						if s, ok := rQuotaSignUpQuotaConfig["startTime"].(string); ok {
							r.Quota.SignUpQuotaConfig.StartTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Quota.SignUpQuotaConfig.StartTime: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Quota.SignUpQuotaConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Quota: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["signIn"]; ok {
		if rSignIn, ok := u.Object["signIn"].(map[string]interface{}); ok {
			r.SignIn = &dclService.ConfigSignIn{}
			if _, ok := rSignIn["allowDuplicateEmails"]; ok {
				if b, ok := rSignIn["allowDuplicateEmails"].(bool); ok {
					r.SignIn.AllowDuplicateEmails = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.SignIn.AllowDuplicateEmails: expected bool")
				}
			}
			if _, ok := rSignIn["anonymous"]; ok {
				if rSignInAnonymous, ok := rSignIn["anonymous"].(map[string]interface{}); ok {
					r.SignIn.Anonymous = &dclService.ConfigSignInAnonymous{}
					if _, ok := rSignInAnonymous["enabled"]; ok {
						if b, ok := rSignInAnonymous["enabled"].(bool); ok {
							r.SignIn.Anonymous.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.SignIn.Anonymous.Enabled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.SignIn.Anonymous: expected map[string]interface{}")
				}
			}
			if _, ok := rSignIn["email"]; ok {
				if rSignInEmail, ok := rSignIn["email"].(map[string]interface{}); ok {
					r.SignIn.Email = &dclService.ConfigSignInEmail{}
					if _, ok := rSignInEmail["enabled"]; ok {
						if b, ok := rSignInEmail["enabled"].(bool); ok {
							r.SignIn.Email.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.SignIn.Email.Enabled: expected bool")
						}
					}
					if _, ok := rSignInEmail["hashConfig"]; ok {
						if rSignInEmailHashConfig, ok := rSignInEmail["hashConfig"].(map[string]interface{}); ok {
							r.SignIn.Email.HashConfig = &dclService.ConfigSignInEmailHashConfig{}
							if _, ok := rSignInEmailHashConfig["algorithm"]; ok {
								if s, ok := rSignInEmailHashConfig["algorithm"].(string); ok {
									r.SignIn.Email.HashConfig.Algorithm = dclService.ConfigSignInEmailHashConfigAlgorithmEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.SignIn.Email.HashConfig.Algorithm: expected string")
								}
							}
							if _, ok := rSignInEmailHashConfig["memoryCost"]; ok {
								if i, ok := rSignInEmailHashConfig["memoryCost"].(int64); ok {
									r.SignIn.Email.HashConfig.MemoryCost = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.SignIn.Email.HashConfig.MemoryCost: expected int64")
								}
							}
							if _, ok := rSignInEmailHashConfig["rounds"]; ok {
								if i, ok := rSignInEmailHashConfig["rounds"].(int64); ok {
									r.SignIn.Email.HashConfig.Rounds = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.SignIn.Email.HashConfig.Rounds: expected int64")
								}
							}
							if _, ok := rSignInEmailHashConfig["saltSeparator"]; ok {
								if s, ok := rSignInEmailHashConfig["saltSeparator"].(string); ok {
									r.SignIn.Email.HashConfig.SaltSeparator = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.SignIn.Email.HashConfig.SaltSeparator: expected string")
								}
							}
							if _, ok := rSignInEmailHashConfig["signerKey"]; ok {
								if s, ok := rSignInEmailHashConfig["signerKey"].(string); ok {
									r.SignIn.Email.HashConfig.SignerKey = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.SignIn.Email.HashConfig.SignerKey: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.SignIn.Email.HashConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rSignInEmail["passwordRequired"]; ok {
						if b, ok := rSignInEmail["passwordRequired"].(bool); ok {
							r.SignIn.Email.PasswordRequired = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.SignIn.Email.PasswordRequired: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.SignIn.Email: expected map[string]interface{}")
				}
			}
			if _, ok := rSignIn["hashConfig"]; ok {
				if rSignInHashConfig, ok := rSignIn["hashConfig"].(map[string]interface{}); ok {
					r.SignIn.HashConfig = &dclService.ConfigSignInHashConfig{}
					if _, ok := rSignInHashConfig["algorithm"]; ok {
						if s, ok := rSignInHashConfig["algorithm"].(string); ok {
							r.SignIn.HashConfig.Algorithm = dclService.ConfigSignInHashConfigAlgorithmEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.SignIn.HashConfig.Algorithm: expected string")
						}
					}
					if _, ok := rSignInHashConfig["memoryCost"]; ok {
						if i, ok := rSignInHashConfig["memoryCost"].(int64); ok {
							r.SignIn.HashConfig.MemoryCost = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.SignIn.HashConfig.MemoryCost: expected int64")
						}
					}
					if _, ok := rSignInHashConfig["rounds"]; ok {
						if i, ok := rSignInHashConfig["rounds"].(int64); ok {
							r.SignIn.HashConfig.Rounds = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.SignIn.HashConfig.Rounds: expected int64")
						}
					}
					if _, ok := rSignInHashConfig["saltSeparator"]; ok {
						if s, ok := rSignInHashConfig["saltSeparator"].(string); ok {
							r.SignIn.HashConfig.SaltSeparator = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SignIn.HashConfig.SaltSeparator: expected string")
						}
					}
					if _, ok := rSignInHashConfig["signerKey"]; ok {
						if s, ok := rSignInHashConfig["signerKey"].(string); ok {
							r.SignIn.HashConfig.SignerKey = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.SignIn.HashConfig.SignerKey: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.SignIn.HashConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rSignIn["phoneNumber"]; ok {
				if rSignInPhoneNumber, ok := rSignIn["phoneNumber"].(map[string]interface{}); ok {
					r.SignIn.PhoneNumber = &dclService.ConfigSignInPhoneNumber{}
					if _, ok := rSignInPhoneNumber["enabled"]; ok {
						if b, ok := rSignInPhoneNumber["enabled"].(bool); ok {
							r.SignIn.PhoneNumber.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.SignIn.PhoneNumber.Enabled: expected bool")
						}
					}
					if _, ok := rSignInPhoneNumber["testPhoneNumbers"]; ok {
						if rSignInPhoneNumberTestPhoneNumbers, ok := rSignInPhoneNumber["testPhoneNumbers"].(map[string]interface{}); ok {
							m := make(map[string]string)
							for k, v := range rSignInPhoneNumberTestPhoneNumbers {
								if s, ok := v.(string); ok {
									m[k] = s
								}
							}
							r.SignIn.PhoneNumber.TestPhoneNumbers = m
						} else {
							return nil, fmt.Errorf("r.SignIn.PhoneNumber.TestPhoneNumbers: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.SignIn.PhoneNumber: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SignIn: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["subtype"]; ok {
		if s, ok := u.Object["subtype"].(string); ok {
			r.Subtype = dclService.ConfigSubtypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Subtype: expected string")
		}
	}
	return r, nil
}

func GetConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConfig(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	return ConfigToUnstructured(r), nil
}

func ApplyConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConfig(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToConfig(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyConfig(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ConfigToUnstructured(r), nil
}

func ConfigHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConfig(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToConfig(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyConfig(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteConfig(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func ConfigID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToConfig(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Config) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"identitytoolkit",
		"Config",
		"ga",
	}
}

func (r *Config) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Config) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Config) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Config) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Config) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Config) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Config) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetConfig(ctx, config, resource)
}

func (r *Config) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyConfig(ctx, config, resource, opts...)
}

func (r *Config) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ConfigHasDiff(ctx, config, resource, opts...)
}

func (r *Config) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteConfig(ctx, config, resource)
}

func (r *Config) ID(resource *unstructured.Resource) (string, error) {
	return ConfigID(resource)
}

func init() {
	unstructured.Register(&Config{})
}
