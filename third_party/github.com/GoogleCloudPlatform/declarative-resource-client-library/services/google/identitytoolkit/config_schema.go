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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLConfigSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "IdentityToolkit/Config",
			Description: "The IdentityToolkit Config resource",
			StructName:  "Config",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Config",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "config",
						Required:    true,
						Description: "A full instance of a Config",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Config",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "config",
						Required:    true,
						Description: "A full instance of a Config",
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Config": &dcl.Component{
					Title:           "Config",
					ID:              "projects/{{project}}/config",
					UsesStateHint:   true,
					ParentContainer: "project",
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
						},
						Properties: map[string]*dcl.Property{
							"authorizedDomains": &dcl.Property{
								Type:          "array",
								GoName:        "AuthorizedDomains",
								Description:   "List of domains authorized for OAuth redirects",
								ServerDefault: true,
								SendEmpty:     true,
								ListType:      "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"blockingFunctions": &dcl.Property{
								Type:          "object",
								GoName:        "BlockingFunctions",
								GoType:        "ConfigBlockingFunctions",
								Description:   "Configuration related to blocking functions.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"triggers": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type:   "object",
											GoType: "ConfigBlockingFunctionsTriggers",
											Properties: map[string]*dcl.Property{
												"functionUri": &dcl.Property{
													Type:        "string",
													GoName:      "FunctionUri",
													Description: "HTTP URI trigger for the Cloud Function.",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Cloudfunctions/Function",
															Field:    "httpsTrigger.url",
														},
													},
												},
												"updateTime": &dcl.Property{
													Type:        "string",
													Format:      "date-time",
													GoName:      "UpdateTime",
													ReadOnly:    true,
													Description: "When the trigger was changed.",
												},
											},
										},
										GoName:      "Triggers",
										Description: "Map of Trigger to event type. Key should be one of the supported event types: \"beforeCreate\", \"beforeSignIn\"",
									},
								},
							},
							"client": &dcl.Property{
								Type:          "object",
								GoName:        "Client",
								GoType:        "ConfigClient",
								Description:   "Options related to how clients making requests on behalf of a project should be configured.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"apiKey": &dcl.Property{
										Type:        "string",
										GoName:      "ApiKey",
										ReadOnly:    true,
										Description: "Output only. API key that can be used when making requests for this project.",
										Sensitive:   true,
									},
									"firebaseSubdomain": &dcl.Property{
										Type:        "string",
										GoName:      "FirebaseSubdomain",
										ReadOnly:    true,
										Description: "Output only. Firebase subdomain.",
									},
									"permissions": &dcl.Property{
										Type:        "object",
										GoName:      "Permissions",
										GoType:      "ConfigClientPermissions",
										Description: "Configuration related to restricting a user's ability to affect their account.",
										Properties: map[string]*dcl.Property{
											"disabledUserDeletion": &dcl.Property{
												Type:        "boolean",
												GoName:      "DisabledUserDeletion",
												Description: "When true, end users cannot delete their account on the associated project through any of our API methods",
											},
											"disabledUserSignup": &dcl.Property{
												Type:        "boolean",
												GoName:      "DisabledUserSignup",
												Description: "When true, end users cannot sign up for a new account on the associated project through any of our API methods",
											},
										},
									},
								},
							},
							"mfa": &dcl.Property{
								Type:          "object",
								GoName:        "Mfa",
								GoType:        "ConfigMfa",
								Description:   "Configuration for this project's multi-factor authentication, including whether it is active and what factors can be used for the second factor",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "ConfigMfaStateEnum",
										Description: "Whether MultiFactor Authentication has been enabled for this project. Possible values: STATE_UNSPECIFIED, DISABLED, ENABLED, MANDATORY",
										Enum: []string{
											"STATE_UNSPECIFIED",
											"DISABLED",
											"ENABLED",
											"MANDATORY",
										},
									},
								},
							},
							"monitoring": &dcl.Property{
								Type:          "object",
								GoName:        "Monitoring",
								GoType:        "ConfigMonitoring",
								Description:   "Configuration related to monitoring project activity.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"requestLogging": &dcl.Property{
										Type:        "object",
										GoName:      "RequestLogging",
										GoType:      "ConfigMonitoringRequestLogging",
										Description: "Configuration for logging requests made to this project to Stackdriver Logging",
										Properties: map[string]*dcl.Property{
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Whether logging is enabled for this project or not.",
											},
										},
									},
								},
							},
							"multiTenant": &dcl.Property{
								Type:          "object",
								GoName:        "MultiTenant",
								GoType:        "ConfigMultiTenant",
								Description:   "Configuration related to multi-tenant functionality.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"allowTenants": &dcl.Property{
										Type:        "boolean",
										GoName:      "AllowTenants",
										Description: "Whether this project can have tenants or not.",
									},
									"defaultTenantLocation": &dcl.Property{
										Type:        "string",
										GoName:      "DefaultTenantLocation",
										Description: "The default cloud parent org or folder that the tenant project should be created under. The parent resource name should be in the format of \"<type>/<number>\", such as \"folders/123\" or \"organizations/456\". If the value is not set, the tenant will be created under the same organization or folder as the agent project.",
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Cloudresourcemanager/Folder",
												Field:    "name",
											},
											&dcl.PropertyResourceReference{
												Resource: "Cloudresourcemanager/Organization",
												Field:    "name",
											},
										},
									},
								},
							},
							"notification": &dcl.Property{
								Type:          "object",
								GoName:        "Notification",
								GoType:        "ConfigNotification",
								Description:   "Configuration related to sending notifications to users.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"defaultLocale": &dcl.Property{
										Type:        "string",
										GoName:      "DefaultLocale",
										Description: "Default locale used for email and SMS in IETF BCP 47 format.",
									},
									"sendEmail": &dcl.Property{
										Type:        "object",
										GoName:      "SendEmail",
										GoType:      "ConfigNotificationSendEmail",
										Description: "Options for email sending.",
										Properties: map[string]*dcl.Property{
											"callbackUri": &dcl.Property{
												Type:        "string",
												GoName:      "CallbackUri",
												Description: "action url in email template.",
											},
											"changeEmailTemplate": &dcl.Property{
												Type:        "object",
												GoName:      "ChangeEmailTemplate",
												GoType:      "ConfigNotificationSendEmailChangeEmailTemplate",
												Description: "Email template for change email",
												Properties: map[string]*dcl.Property{
													"body": &dcl.Property{
														Type:        "string",
														GoName:      "Body",
														Description: "Email body",
														Immutable:   true,
													},
													"bodyFormat": &dcl.Property{
														Type:        "string",
														GoName:      "BodyFormat",
														GoType:      "ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum",
														Description: "Email body format Possible values: BODY_FORMAT_UNSPECIFIED, PLAIN_TEXT, HTML",
														Enum: []string{
															"BODY_FORMAT_UNSPECIFIED",
															"PLAIN_TEXT",
															"HTML",
														},
													},
													"customized": &dcl.Property{
														Type:        "boolean",
														GoName:      "Customized",
														ReadOnly:    true,
														Description: "Output only. Whether the body or subject of the email is customized.",
														Immutable:   true,
													},
													"replyTo": &dcl.Property{
														Type:        "string",
														GoName:      "ReplyTo",
														Description: "Reply-to address",
													},
													"senderDisplayName": &dcl.Property{
														Type:        "string",
														GoName:      "SenderDisplayName",
														Description: "Sender display name",
													},
													"senderLocalPart": &dcl.Property{
														Type:        "string",
														GoName:      "SenderLocalPart",
														Description: "Local part of From address",
													},
													"subject": &dcl.Property{
														Type:        "string",
														GoName:      "Subject",
														Description: "Subject of the email",
													},
												},
											},
											"dnsInfo": &dcl.Property{
												Type:        "object",
												GoName:      "DnsInfo",
												GoType:      "ConfigNotificationSendEmailDnsInfo",
												Description: "Information of custom domain DNS verification.",
												Properties: map[string]*dcl.Property{
													"customDomain": &dcl.Property{
														Type:        "string",
														GoName:      "CustomDomain",
														ReadOnly:    true,
														Description: "Output only. The applied verified custom domain.",
														Immutable:   true,
													},
													"customDomainState": &dcl.Property{
														Type:        "string",
														GoName:      "CustomDomainState",
														GoType:      "ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum",
														ReadOnly:    true,
														Description: "Output only. The current verification state of the custom domain. The custom domain will only be used once the domain verification is successful. Possible values: VERIFICATION_STATE_UNSPECIFIED, NOT_STARTED, IN_PROGRESS, FAILED, SUCCEEDED",
														Immutable:   true,
														Enum: []string{
															"VERIFICATION_STATE_UNSPECIFIED",
															"NOT_STARTED",
															"IN_PROGRESS",
															"FAILED",
															"SUCCEEDED",
														},
													},
													"domainVerificationRequestTime": &dcl.Property{
														Type:        "string",
														Format:      "date-time",
														GoName:      "DomainVerificationRequestTime",
														ReadOnly:    true,
														Description: "Output only. The timestamp of initial request for the current domain verification.",
														Immutable:   true,
													},
													"pendingCustomDomain": &dcl.Property{
														Type:        "string",
														GoName:      "PendingCustomDomain",
														ReadOnly:    true,
														Description: "Output only. The custom domain that's to be verified.",
														Immutable:   true,
													},
													"useCustomDomain": &dcl.Property{
														Type:        "boolean",
														GoName:      "UseCustomDomain",
														Description: "Whether to use custom domain.",
													},
												},
											},
											"method": &dcl.Property{
												Type:        "string",
												GoName:      "Method",
												GoType:      "ConfigNotificationSendEmailMethodEnum",
												Description: "The method used for sending an email. Possible values: METHOD_UNSPECIFIED, DEFAULT, CUSTOM_SMTP",
												Enum: []string{
													"METHOD_UNSPECIFIED",
													"DEFAULT",
													"CUSTOM_SMTP",
												},
											},
											"resetPasswordTemplate": &dcl.Property{
												Type:        "object",
												GoName:      "ResetPasswordTemplate",
												GoType:      "ConfigNotificationSendEmailResetPasswordTemplate",
												Description: "Email template for reset password",
												Properties: map[string]*dcl.Property{
													"body": &dcl.Property{
														Type:        "string",
														GoName:      "Body",
														Description: "Email body",
													},
													"bodyFormat": &dcl.Property{
														Type:        "string",
														GoName:      "BodyFormat",
														GoType:      "ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum",
														Description: "Email body format Possible values: BODY_FORMAT_UNSPECIFIED, PLAIN_TEXT, HTML",
														Enum: []string{
															"BODY_FORMAT_UNSPECIFIED",
															"PLAIN_TEXT",
															"HTML",
														},
													},
													"customized": &dcl.Property{
														Type:        "boolean",
														GoName:      "Customized",
														ReadOnly:    true,
														Description: "Output only. Whether the body or subject of the email is customized.",
														Immutable:   true,
													},
													"replyTo": &dcl.Property{
														Type:        "string",
														GoName:      "ReplyTo",
														Description: "Reply-to address",
													},
													"senderDisplayName": &dcl.Property{
														Type:        "string",
														GoName:      "SenderDisplayName",
														Description: "Sender display name",
													},
													"senderLocalPart": &dcl.Property{
														Type:        "string",
														GoName:      "SenderLocalPart",
														Description: "Local part of From address",
													},
													"subject": &dcl.Property{
														Type:        "string",
														GoName:      "Subject",
														Description: "Subject of the email",
													},
												},
											},
											"revertSecondFactorAdditionTemplate": &dcl.Property{
												Type:        "object",
												GoName:      "RevertSecondFactorAdditionTemplate",
												GoType:      "ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate",
												Description: "Email template for reverting second factor addition emails",
												Properties: map[string]*dcl.Property{
													"body": &dcl.Property{
														Type:        "string",
														GoName:      "Body",
														Description: "Email body",
														Immutable:   true,
													},
													"bodyFormat": &dcl.Property{
														Type:        "string",
														GoName:      "BodyFormat",
														GoType:      "ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum",
														Description: "Email body format Possible values: BODY_FORMAT_UNSPECIFIED, PLAIN_TEXT, HTML",
														Enum: []string{
															"BODY_FORMAT_UNSPECIFIED",
															"PLAIN_TEXT",
															"HTML",
														},
													},
													"customized": &dcl.Property{
														Type:        "boolean",
														GoName:      "Customized",
														ReadOnly:    true,
														Description: "Output only. Whether the body or subject of the email is customized.",
														Immutable:   true,
													},
													"replyTo": &dcl.Property{
														Type:        "string",
														GoName:      "ReplyTo",
														Description: "Reply-to address",
													},
													"senderDisplayName": &dcl.Property{
														Type:        "string",
														GoName:      "SenderDisplayName",
														Description: "Sender display name",
													},
													"senderLocalPart": &dcl.Property{
														Type:        "string",
														GoName:      "SenderLocalPart",
														Description: "Local part of From address",
													},
													"subject": &dcl.Property{
														Type:        "string",
														GoName:      "Subject",
														Description: "Subject of the email",
													},
												},
											},
											"smtp": &dcl.Property{
												Type:        "object",
												GoName:      "Smtp",
												GoType:      "ConfigNotificationSendEmailSmtp",
												Description: "Use a custom SMTP relay",
												Properties: map[string]*dcl.Property{
													"host": &dcl.Property{
														Type:        "string",
														GoName:      "Host",
														Description: "SMTP relay host",
													},
													"password": &dcl.Property{
														Type:        "string",
														GoName:      "Password",
														Description: "SMTP relay password",
														Sensitive:   true,
														Unreadable:  true,
													},
													"port": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "Port",
														Description: "SMTP relay port",
													},
													"securityMode": &dcl.Property{
														Type:        "string",
														GoName:      "SecurityMode",
														GoType:      "ConfigNotificationSendEmailSmtpSecurityModeEnum",
														Description: "SMTP security mode. Possible values: SECURITY_MODE_UNSPECIFIED, SSL, START_TLS",
														Enum: []string{
															"SECURITY_MODE_UNSPECIFIED",
															"SSL",
															"START_TLS",
														},
													},
													"senderEmail": &dcl.Property{
														Type:        "string",
														GoName:      "SenderEmail",
														Description: "Sender email for the SMTP relay",
													},
													"username": &dcl.Property{
														Type:        "string",
														GoName:      "Username",
														Description: "SMTP relay username",
													},
												},
											},
											"verifyEmailTemplate": &dcl.Property{
												Type:        "object",
												GoName:      "VerifyEmailTemplate",
												GoType:      "ConfigNotificationSendEmailVerifyEmailTemplate",
												Description: "Email template for verify email",
												Properties: map[string]*dcl.Property{
													"body": &dcl.Property{
														Type:        "string",
														GoName:      "Body",
														Description: "Email body",
														Immutable:   true,
													},
													"bodyFormat": &dcl.Property{
														Type:        "string",
														GoName:      "BodyFormat",
														GoType:      "ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum",
														Description: "Email body format Possible values: BODY_FORMAT_UNSPECIFIED, PLAIN_TEXT, HTML",
														Enum: []string{
															"BODY_FORMAT_UNSPECIFIED",
															"PLAIN_TEXT",
															"HTML",
														},
													},
													"customized": &dcl.Property{
														Type:        "boolean",
														GoName:      "Customized",
														ReadOnly:    true,
														Description: "Output only. Whether the body or subject of the email is customized.",
														Immutable:   true,
													},
													"replyTo": &dcl.Property{
														Type:        "string",
														GoName:      "ReplyTo",
														Description: "Reply-to address",
													},
													"senderDisplayName": &dcl.Property{
														Type:        "string",
														GoName:      "SenderDisplayName",
														Description: "Sender display name",
													},
													"senderLocalPart": &dcl.Property{
														Type:        "string",
														GoName:      "SenderLocalPart",
														Description: "Local part of From address",
													},
													"subject": &dcl.Property{
														Type:        "string",
														GoName:      "Subject",
														Description: "Subject of the email",
													},
												},
											},
										},
									},
									"sendSms": &dcl.Property{
										Type:        "object",
										GoName:      "SendSms",
										GoType:      "ConfigNotificationSendSms",
										Description: "Options for SMS sending.",
										Properties: map[string]*dcl.Property{
											"smsTemplate": &dcl.Property{
												Type:        "object",
												GoName:      "SmsTemplate",
												GoType:      "ConfigNotificationSendSmsSmsTemplate",
												ReadOnly:    true,
												Description: "Output only. The template to use when sending an SMS.",
												Immutable:   true,
												Properties: map[string]*dcl.Property{
													"content": &dcl.Property{
														Type:        "string",
														GoName:      "Content",
														ReadOnly:    true,
														Description: "Output only. The SMS's content. Can contain the following placeholders which will be replaced with the appropriate values: %APP_NAME% - For Android or iOS apps, the app's display name. For web apps, the domain hosting the application. %LOGIN_CODE% - The OOB code being sent in the SMS.",
														Immutable:   true,
													},
												},
											},
											"useDeviceLocale": &dcl.Property{
												Type:        "boolean",
												GoName:      "UseDeviceLocale",
												Description: "Whether to use the accept_language header for SMS.",
											},
										},
									},
								},
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project of the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"quota": &dcl.Property{
								Type:          "object",
								GoName:        "Quota",
								GoType:        "ConfigQuota",
								Description:   "Configuration related to quotas.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"signUpQuotaConfig": &dcl.Property{
										Type:        "object",
										GoName:      "SignUpQuotaConfig",
										GoType:      "ConfigQuotaSignUpQuotaConfig",
										Description: "Quota for the Signup endpoint, if overwritten. Signup quota is measured in sign ups per project per hour per IP.",
										Properties: map[string]*dcl.Property{
											"quota": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Quota",
												Description: "Corresponds to the 'refill_token_count' field in QuotaServer config",
											},
											"quotaDuration": &dcl.Property{
												Type:        "string",
												GoName:      "QuotaDuration",
												Description: "How long this quota will be active for",
											},
											"startTime": &dcl.Property{
												Type:        "string",
												Format:      "date-time",
												GoName:      "StartTime",
												Description: "When this quota will take affect",
											},
										},
									},
								},
							},
							"signIn": &dcl.Property{
								Type:          "object",
								GoName:        "SignIn",
								GoType:        "ConfigSignIn",
								Description:   "Configuration related to local sign in methods.",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"allowDuplicateEmails": &dcl.Property{
										Type:        "boolean",
										GoName:      "AllowDuplicateEmails",
										Description: "Whether to allow more than one account to have the same email.",
									},
									"anonymous": &dcl.Property{
										Type:        "object",
										GoName:      "Anonymous",
										GoType:      "ConfigSignInAnonymous",
										Description: "Configuration options related to authenticating an anonymous user.",
										Properties: map[string]*dcl.Property{
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Whether anonymous user auth is enabled for the project or not.",
											},
										},
									},
									"email": &dcl.Property{
										Type:        "object",
										GoName:      "Email",
										GoType:      "ConfigSignInEmail",
										Description: "Configuration options related to authenticating a user by their email address.",
										Properties: map[string]*dcl.Property{
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Whether email auth is enabled for the project or not.",
											},
											"hashConfig": &dcl.Property{
												Type:        "object",
												GoName:      "HashConfig",
												GoType:      "ConfigSignInEmailHashConfig",
												ReadOnly:    true,
												Description: "Output only. Hash config information.",
												Properties: map[string]*dcl.Property{
													"algorithm": &dcl.Property{
														Type:        "string",
														GoName:      "Algorithm",
														GoType:      "ConfigSignInEmailHashConfigAlgorithmEnum",
														ReadOnly:    true,
														Description: "Output only. Different password hash algorithms used in Identity Toolkit. Possible values: HASH_ALGORITHM_UNSPECIFIED, HMAC_SHA256, HMAC_SHA1, HMAC_MD5, SCRYPT, PBKDF_SHA1, MD5, HMAC_SHA512, SHA1, BCRYPT, PBKDF2_SHA256, SHA256, SHA512, STANDARD_SCRYPT",
														Enum: []string{
															"HASH_ALGORITHM_UNSPECIFIED",
															"HMAC_SHA256",
															"HMAC_SHA1",
															"HMAC_MD5",
															"SCRYPT",
															"PBKDF_SHA1",
															"MD5",
															"HMAC_SHA512",
															"SHA1",
															"BCRYPT",
															"PBKDF2_SHA256",
															"SHA256",
															"SHA512",
															"STANDARD_SCRYPT",
														},
													},
													"memoryCost": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "MemoryCost",
														ReadOnly:    true,
														Description: "Output only. Memory cost for hash calculation. Used by scrypt and other similar password derivation algorithms. See https://tools.ietf.org/html/rfc7914 for explanation of field.",
													},
													"rounds": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "Rounds",
														ReadOnly:    true,
														Description: "Output only. How many rounds for hash calculation. Used by scrypt and other similar password derivation algorithms.",
													},
													"saltSeparator": &dcl.Property{
														Type:        "string",
														GoName:      "SaltSeparator",
														ReadOnly:    true,
														Description: "Output only. Non-printable character to be inserted between the salt and plain text password in base64.",
													},
													"signerKey": &dcl.Property{
														Type:        "string",
														GoName:      "SignerKey",
														ReadOnly:    true,
														Description: "Output only. Signer key in base64.",
														Sensitive:   true,
													},
												},
											},
											"passwordRequired": &dcl.Property{
												Type:        "boolean",
												GoName:      "PasswordRequired",
												Description: "Whether a password is required for email auth or not. If true, both an email and password must be provided to sign in. If false, a user may sign in via either email/password or email link.",
											},
										},
									},
									"hashConfig": &dcl.Property{
										Type:        "object",
										GoName:      "HashConfig",
										GoType:      "ConfigSignInHashConfig",
										ReadOnly:    true,
										Description: "Output only. Hash config information.",
										Properties: map[string]*dcl.Property{
											"algorithm": &dcl.Property{
												Type:        "string",
												GoName:      "Algorithm",
												GoType:      "ConfigSignInHashConfigAlgorithmEnum",
												ReadOnly:    true,
												Description: "Output only. Different password hash algorithms used in Identity Toolkit. Possible values: HASH_ALGORITHM_UNSPECIFIED, HMAC_SHA256, HMAC_SHA1, HMAC_MD5, SCRYPT, PBKDF_SHA1, MD5, HMAC_SHA512, SHA1, BCRYPT, PBKDF2_SHA256, SHA256, SHA512, STANDARD_SCRYPT",
												Enum: []string{
													"HASH_ALGORITHM_UNSPECIFIED",
													"HMAC_SHA256",
													"HMAC_SHA1",
													"HMAC_MD5",
													"SCRYPT",
													"PBKDF_SHA1",
													"MD5",
													"HMAC_SHA512",
													"SHA1",
													"BCRYPT",
													"PBKDF2_SHA256",
													"SHA256",
													"SHA512",
													"STANDARD_SCRYPT",
												},
											},
											"memoryCost": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "MemoryCost",
												ReadOnly:    true,
												Description: "Output only. Memory cost for hash calculation. Used by scrypt and other similar password derivation algorithms. See https://tools.ietf.org/html/rfc7914 for explanation of field.",
											},
											"rounds": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Rounds",
												ReadOnly:    true,
												Description: "Output only. How many rounds for hash calculation. Used by scrypt and other similar password derivation algorithms.",
											},
											"saltSeparator": &dcl.Property{
												Type:        "string",
												GoName:      "SaltSeparator",
												ReadOnly:    true,
												Description: "Output only. Non-printable character to be inserted between the salt and plain text password in base64.",
											},
											"signerKey": &dcl.Property{
												Type:        "string",
												GoName:      "SignerKey",
												ReadOnly:    true,
												Description: "Output only. Signer key in base64.",
												Sensitive:   true,
											},
										},
									},
									"phoneNumber": &dcl.Property{
										Type:        "object",
										GoName:      "PhoneNumber",
										GoType:      "ConfigSignInPhoneNumber",
										Description: "Configuration options related to authenticated a user by their phone number.",
										Properties: map[string]*dcl.Property{
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Whether phone number auth is enabled for the project or not.",
											},
											"testPhoneNumbers": &dcl.Property{
												Type: "object",
												AdditionalProperties: &dcl.Property{
													Type: "string",
												},
												GoName:      "TestPhoneNumbers",
												Description: "A map of that can be used for phone auth testing.",
											},
										},
									},
								},
							},
							"subtype": &dcl.Property{
								Type:        "string",
								GoName:      "Subtype",
								GoType:      "ConfigSubtypeEnum",
								ReadOnly:    true,
								Description: "Output only. The subtype of this config. Possible values: SUBTYPE_UNSPECIFIED, IDENTITY_PLATFORM, FIREBASE_AUTH",
								Immutable:   true,
								Enum: []string{
									"SUBTYPE_UNSPECIFIED",
									"IDENTITY_PLATFORM",
									"FIREBASE_AUTH",
								},
							},
						},
					},
				},
			},
		},
	}
}
