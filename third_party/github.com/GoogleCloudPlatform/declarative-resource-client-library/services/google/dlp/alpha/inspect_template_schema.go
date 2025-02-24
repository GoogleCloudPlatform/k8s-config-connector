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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLInspectTemplateSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Dlp/InspectTemplate",
			Description: "The Dlp InspectTemplate resource",
			StructName:  "InspectTemplate",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a InspectTemplate",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "inspectTemplate",
						Required:    true,
						Description: "A full instance of a InspectTemplate",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a InspectTemplate",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "inspectTemplate",
						Required:    true,
						Description: "A full instance of a InspectTemplate",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a InspectTemplate",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "inspectTemplate",
						Required:    true,
						Description: "A full instance of a InspectTemplate",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all InspectTemplate",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many InspectTemplate",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"InspectTemplate": &dcl.Component{
					Title: "InspectTemplate",
					ID:    "{{parent}}/inspectTemplates/{{name}}",
					Locations: []string{
						"region",
					},
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"parent",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The creation timestamp of an inspectTemplate.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Short description (max 256 chars).",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Display name (max 256 chars).",
							},
							"inspectConfig": &dcl.Property{
								Type:        "object",
								GoName:      "InspectConfig",
								GoType:      "InspectTemplateInspectConfig",
								Description: "The core content of the template. Configuration of the scanning process.",
								Properties: map[string]*dcl.Property{
									"contentOptions": &dcl.Property{
										Type:        "array",
										GoName:      "ContentOptions",
										Description: "List of options defining data content to scan. If empty, text, images, and other content will be included.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "InspectTemplateInspectConfigContentOptionsEnum",
											Enum: []string{
												"CONTENT_UNSPECIFIED",
												"CONTENT_TEXT",
												"CONTENT_IMAGE",
											},
										},
									},
									"customInfoTypes": &dcl.Property{
										Type:        "array",
										GoName:      "CustomInfoTypes",
										Description: "CustomInfoTypes provided by the user. See https://cloud.google.com/dlp/docs/creating-custom-infotypes to learn more.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "InspectTemplateInspectConfigCustomInfoTypes",
											Properties: map[string]*dcl.Property{
												"dictionary": &dcl.Property{
													Type:        "object",
													GoName:      "Dictionary",
													GoType:      "InspectTemplateInspectConfigCustomInfoTypesDictionary",
													Description: "A list of phrases to detect as a CustomInfoType.",
													Conflicts: []string{
														"regex",
														"surrogateType",
														"storedType",
													},
													Properties: map[string]*dcl.Property{
														"cloudStoragePath": &dcl.Property{
															Type:        "object",
															GoName:      "CloudStoragePath",
															GoType:      "InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath",
															Description: "Newline-delimited file of words in Cloud Storage. Only a single file is accepted.",
															Conflicts: []string{
																"wordList",
															},
															Properties: map[string]*dcl.Property{
																"path": &dcl.Property{
																	Type:        "string",
																	GoName:      "Path",
																	Description: "A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt",
																},
															},
														},
														"wordList": &dcl.Property{
															Type:        "object",
															GoName:      "WordList",
															GoType:      "InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList",
															Description: "List of words or phrases to search for.",
															Conflicts: []string{
																"cloudStoragePath",
															},
															Properties: map[string]*dcl.Property{
																"words": &dcl.Property{
																	Type:        "array",
																	GoName:      "Words",
																	Description: "Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits. [required]",
																	SendEmpty:   true,
																	ListType:    "list",
																	Items: &dcl.Property{
																		Type:   "string",
																		GoType: "string",
																	},
																},
															},
														},
													},
												},
												"exclusionType": &dcl.Property{
													Type:        "string",
													GoName:      "ExclusionType",
													GoType:      "InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum",
													Description: "If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned. It still can be used for rules matching. Possible values: EXCLUSION_TYPE_UNSPECIFIED, EXCLUSION_TYPE_EXCLUDE",
													Enum: []string{
														"EXCLUSION_TYPE_UNSPECIFIED",
														"EXCLUSION_TYPE_EXCLUDE",
													},
												},
												"infoType": &dcl.Property{
													Type:        "object",
													GoName:      "InfoType",
													GoType:      "InspectTemplateInspectConfigCustomInfoTypesInfoType",
													Description: "CustomInfoType can either be a new infoType, or an extension of built-in infoType, when the name matches one of existing infoTypes and that infoType is specified in `InspectContent.info_types` field. Specifying the latter adds findings to the one detected by the system. If built-in info type is not specified in `InspectContent.info_types` list then the name is treated as a custom info type.",
													Properties: map[string]*dcl.Property{
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
														},
													},
												},
												"likelihood": &dcl.Property{
													Type:        "string",
													GoName:      "Likelihood",
													GoType:      "InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum",
													Description: "Likelihood to return for this CustomInfoType. This base value can be altered by a detection rule if the finding meets the criteria specified by the rule. Defaults to `VERY_LIKELY` if not specified. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY",
													Enum: []string{
														"LIKELIHOOD_UNSPECIFIED",
														"VERY_UNLIKELY",
														"UNLIKELY",
														"POSSIBLE",
														"LIKELY",
														"VERY_LIKELY",
													},
												},
												"regex": &dcl.Property{
													Type:        "object",
													GoName:      "Regex",
													GoType:      "InspectTemplateInspectConfigCustomInfoTypesRegex",
													Description: "Regular expression based CustomInfoType.",
													Conflicts: []string{
														"dictionary",
														"surrogateType",
														"storedType",
													},
													Properties: map[string]*dcl.Property{
														"groupIndexes": &dcl.Property{
															Type:        "array",
															GoName:      "GroupIndexes",
															Description: "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "integer",
																Format: "int64",
																GoType: "int64",
															},
														},
														"pattern": &dcl.Property{
															Type:        "string",
															GoName:      "Pattern",
															Description: "Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
														},
													},
												},
												"storedType": &dcl.Property{
													Type:        "object",
													GoName:      "StoredType",
													GoType:      "InspectTemplateInspectConfigCustomInfoTypesStoredType",
													Description: "Load an existing `StoredInfoType` resource for use in `InspectDataSource`. Not currently supported in `InspectContent`.",
													Conflicts: []string{
														"dictionary",
														"regex",
														"surrogateType",
													},
													Properties: map[string]*dcl.Property{
														"createTime": &dcl.Property{
															Type:        "string",
															Format:      "date-time",
															GoName:      "CreateTime",
															ReadOnly:    true,
															Description: "Timestamp indicating when the version of the `StoredInfoType` used for inspection was created. Output-only field, populated by the system.",
														},
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "Resource name of the requested `StoredInfoType`, for example `organizations/433245324/storedInfoTypes/432452342` or `projects/project-id/storedInfoTypes/432452342`.",
															ResourceReferences: []*dcl.PropertyResourceReference{
																&dcl.PropertyResourceReference{
																	Resource: "Dlp/StoredInfoType",
																	Field:    "name",
																	Parent:   true,
																},
															},
														},
													},
												},
												"surrogateType": &dcl.Property{
													Type:        "object",
													GoName:      "SurrogateType",
													GoType:      "InspectTemplateInspectConfigCustomInfoTypesSurrogateType",
													Description: "Message for detecting output from deidentification transformations that support reversing.",
													Conflicts: []string{
														"dictionary",
														"regex",
														"storedType",
													},
													Properties: map[string]*dcl.Property{},
												},
											},
										},
									},
									"excludeInfoTypes": &dcl.Property{
										Type:        "boolean",
										GoName:      "ExcludeInfoTypes",
										Description: "When true, excludes type information of the findings.",
									},
									"includeQuote": &dcl.Property{
										Type:        "boolean",
										GoName:      "IncludeQuote",
										Description: "When true, a contextual quote from the data that triggered a finding is included in the response; see Finding.quote.",
									},
									"infoTypes": &dcl.Property{
										Type:        "array",
										GoName:      "InfoTypes",
										Description: "Restricts what info_types to look for. The values must correspond to InfoType values returned by ListInfoTypes or listed at https://cloud.google.com/dlp/docs/infotypes-reference. When no InfoTypes or CustomInfoTypes are specified in a request, the system may automatically choose what detectors to run. By default this may be all types, but may change over time as detectors are updated. If you need precise control and predictability as to what detectors are run you should specify specific InfoTypes listed in the reference, otherwise a default list will be used, which may change over time.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "InspectTemplateInspectConfigInfoTypes",
											Properties: map[string]*dcl.Property{
												"name": &dcl.Property{
													Type:        "string",
													GoName:      "Name",
													Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
												},
											},
										},
									},
									"limits": &dcl.Property{
										Type:        "object",
										GoName:      "Limits",
										GoType:      "InspectTemplateInspectConfigLimits",
										Description: "Configuration to control the number of findings returned.",
										Properties: map[string]*dcl.Property{
											"maxFindingsPerInfoType": &dcl.Property{
												Type:        "array",
												GoName:      "MaxFindingsPerInfoType",
												Description: "Configuration of findings limit given for specified infoTypes.",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType",
													Properties: map[string]*dcl.Property{
														"infoType": &dcl.Property{
															Type:        "object",
															GoName:      "InfoType",
															GoType:      "InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType",
															Description: "Type of information the findings limit applies to. Only one limit per info_type should be provided. If InfoTypeLimit does not have an info_type, the DLP API applies the limit against all info_types that are found but not specified in another InfoTypeLimit.",
															Properties: map[string]*dcl.Property{
																"name": &dcl.Property{
																	Type:        "string",
																	GoName:      "Name",
																	Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
																},
															},
														},
														"maxFindings": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "MaxFindings",
															Description: "Max findings limit for the given infoType.",
														},
													},
												},
											},
											"maxFindingsPerItem": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "MaxFindingsPerItem",
												Description: "Max number of findings that will be returned for each item scanned. When set within `InspectJobConfig`, the maximum returned is 2000 regardless if this is set higher. When set within `InspectContentRequest`, this field is ignored.",
											},
											"maxFindingsPerRequest": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "MaxFindingsPerRequest",
												Description: "Max number of findings that will be returned per request/job. When set within `InspectContentRequest`, the maximum returned is 2000 regardless if this is set higher.",
											},
										},
									},
									"minLikelihood": &dcl.Property{
										Type:        "string",
										GoName:      "MinLikelihood",
										GoType:      "InspectTemplateInspectConfigMinLikelihoodEnum",
										Description: "Only returns findings equal or above this threshold. The default is POSSIBLE. See https://cloud.google.com/dlp/docs/likelihood to learn more. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY",
										Enum: []string{
											"LIKELIHOOD_UNSPECIFIED",
											"VERY_UNLIKELY",
											"UNLIKELY",
											"POSSIBLE",
											"LIKELY",
											"VERY_LIKELY",
										},
									},
									"ruleSet": &dcl.Property{
										Type:        "array",
										GoName:      "RuleSet",
										Description: "Set of rules to apply to the findings for this InspectConfig. Exclusion rules, contained in the set are executed in the end, other rules are executed in the order they are specified for each info type.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "InspectTemplateInspectConfigRuleSet",
											Properties: map[string]*dcl.Property{
												"infoTypes": &dcl.Property{
													Type:        "array",
													GoName:      "InfoTypes",
													Description: "List of infoTypes this rule set is applied to.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "InspectTemplateInspectConfigRuleSetInfoTypes",
														Properties: map[string]*dcl.Property{
															"name": &dcl.Property{
																Type:        "string",
																GoName:      "Name",
																Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
															},
														},
													},
												},
												"rules": &dcl.Property{
													Type:        "array",
													GoName:      "Rules",
													Description: "Set of rules to be applied to infoTypes. The rules are applied in order.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "object",
														GoType: "InspectTemplateInspectConfigRuleSetRules",
														Properties: map[string]*dcl.Property{
															"exclusionRule": &dcl.Property{
																Type:        "object",
																GoName:      "ExclusionRule",
																GoType:      "InspectTemplateInspectConfigRuleSetRulesExclusionRule",
																Description: "Exclusion rule.",
																Conflicts: []string{
																	"hotwordRule",
																},
																Properties: map[string]*dcl.Property{
																	"dictionary": &dcl.Property{
																		Type:        "object",
																		GoName:      "Dictionary",
																		GoType:      "InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary",
																		Description: "Dictionary which defines the rule.",
																		Conflicts: []string{
																			"regex",
																			"excludeInfoTypes",
																		},
																		Properties: map[string]*dcl.Property{
																			"cloudStoragePath": &dcl.Property{
																				Type:        "object",
																				GoName:      "CloudStoragePath",
																				GoType:      "InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath",
																				Description: "Newline-delimited file of words in Cloud Storage. Only a single file is accepted.",
																				Conflicts: []string{
																					"wordList",
																				},
																				Properties: map[string]*dcl.Property{
																					"path": &dcl.Property{
																						Type:        "string",
																						GoName:      "Path",
																						Description: "A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt",
																					},
																				},
																			},
																			"wordList": &dcl.Property{
																				Type:        "object",
																				GoName:      "WordList",
																				GoType:      "InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList",
																				Description: "List of words or phrases to search for.",
																				Conflicts: []string{
																					"cloudStoragePath",
																				},
																				Properties: map[string]*dcl.Property{
																					"words": &dcl.Property{
																						Type:        "array",
																						GoName:      "Words",
																						Description: "Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits. [required]",
																						SendEmpty:   true,
																						ListType:    "list",
																						Items: &dcl.Property{
																							Type:   "string",
																							GoType: "string",
																						},
																					},
																				},
																			},
																		},
																	},
																	"excludeInfoTypes": &dcl.Property{
																		Type:        "object",
																		GoName:      "ExcludeInfoTypes",
																		GoType:      "InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes",
																		Description: "Set of infoTypes for which findings would affect this rule.",
																		Conflicts: []string{
																			"dictionary",
																			"regex",
																		},
																		Properties: map[string]*dcl.Property{
																			"infoTypes": &dcl.Property{
																				Type:        "array",
																				GoName:      "InfoTypes",
																				Description: "InfoType list in ExclusionRule rule drops a finding when it overlaps or contained within with a finding of an infoType from this list. For example, for `InspectionRuleSet.info_types` containing \"PHONE_NUMBER\"` and `exclusion_rule` containing `exclude_info_types.info_types` with \"EMAIL_ADDRESS\" the phone number findings are dropped if they overlap with EMAIL_ADDRESS finding. That leads to \"555-222-2222@example.org\" to generate only a single finding, namely email address.",
																				SendEmpty:   true,
																				ListType:    "list",
																				Items: &dcl.Property{
																					Type:   "object",
																					GoType: "InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes",
																					Properties: map[string]*dcl.Property{
																						"name": &dcl.Property{
																							Type:        "string",
																							GoName:      "Name",
																							Description: "Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type. When sending Cloud DLP results to Data Catalog, infoType names should conform to the pattern `[A-Za-z0-9$-_]{1,64}`.",
																						},
																					},
																				},
																			},
																		},
																	},
																	"matchingType": &dcl.Property{
																		Type:        "string",
																		GoName:      "MatchingType",
																		GoType:      "InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum",
																		Description: "How the rule is applied, see MatchingType documentation for details. Possible values: MATCHING_TYPE_UNSPECIFIED, MATCHING_TYPE_FULL_MATCH, MATCHING_TYPE_PARTIAL_MATCH, MATCHING_TYPE_INVERSE_MATCH",
																		Enum: []string{
																			"MATCHING_TYPE_UNSPECIFIED",
																			"MATCHING_TYPE_FULL_MATCH",
																			"MATCHING_TYPE_PARTIAL_MATCH",
																			"MATCHING_TYPE_INVERSE_MATCH",
																		},
																	},
																	"regex": &dcl.Property{
																		Type:        "object",
																		GoName:      "Regex",
																		GoType:      "InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex",
																		Description: "Regular expression which defines the rule.",
																		Conflicts: []string{
																			"dictionary",
																			"excludeInfoTypes",
																		},
																		Properties: map[string]*dcl.Property{
																			"groupIndexes": &dcl.Property{
																				Type:        "array",
																				GoName:      "GroupIndexes",
																				Description: "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
																				SendEmpty:   true,
																				ListType:    "list",
																				Items: &dcl.Property{
																					Type:   "integer",
																					Format: "int64",
																					GoType: "int64",
																				},
																			},
																			"pattern": &dcl.Property{
																				Type:        "string",
																				GoName:      "Pattern",
																				Description: "Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
																			},
																		},
																	},
																},
															},
															"hotwordRule": &dcl.Property{
																Type:   "object",
																GoName: "HotwordRule",
																GoType: "InspectTemplateInspectConfigRuleSetRulesHotwordRule",
																Conflicts: []string{
																	"exclusionRule",
																},
																Properties: map[string]*dcl.Property{
																	"hotwordRegex": &dcl.Property{
																		Type:        "object",
																		GoName:      "HotwordRegex",
																		GoType:      "InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex",
																		Description: "Regular expression pattern defining what qualifies as a hotword.",
																		Properties: map[string]*dcl.Property{
																			"groupIndexes": &dcl.Property{
																				Type:        "array",
																				GoName:      "GroupIndexes",
																				Description: "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
																				SendEmpty:   true,
																				ListType:    "list",
																				Items: &dcl.Property{
																					Type:   "integer",
																					Format: "int64",
																					GoType: "int64",
																				},
																			},
																			"pattern": &dcl.Property{
																				Type:        "string",
																				GoName:      "Pattern",
																				Description: "Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
																			},
																		},
																	},
																	"likelihoodAdjustment": &dcl.Property{
																		Type:        "object",
																		GoName:      "LikelihoodAdjustment",
																		GoType:      "InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment",
																		Description: "Likelihood adjustment to apply to all matching findings.",
																		Properties: map[string]*dcl.Property{
																			"fixedLikelihood": &dcl.Property{
																				Type:        "string",
																				GoName:      "FixedLikelihood",
																				GoType:      "InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum",
																				Description: "Set the likelihood of a finding to a fixed value. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY",
																				Conflicts: []string{
																					"relativeLikelihood",
																				},
																				Enum: []string{
																					"LIKELIHOOD_UNSPECIFIED",
																					"VERY_UNLIKELY",
																					"UNLIKELY",
																					"POSSIBLE",
																					"LIKELY",
																					"VERY_LIKELY",
																				},
																			},
																			"relativeLikelihood": &dcl.Property{
																				Type:        "integer",
																				Format:      "int64",
																				GoName:      "RelativeLikelihood",
																				Description: "Increase or decrease the likelihood by the specified number of levels. For example, if a finding would be `POSSIBLE` without the detection rule and `relative_likelihood` is 1, then it is upgraded to `LIKELY`, while a value of -1 would downgrade it to `UNLIKELY`. Likelihood may never drop below `VERY_UNLIKELY` or exceed `VERY_LIKELY`, so applying an adjustment of 1 followed by an adjustment of -1 when base likelihood is `VERY_LIKELY` will result in a final likelihood of `LIKELY`.",
																				Conflicts: []string{
																					"fixedLikelihood",
																				},
																			},
																		},
																	},
																	"proximity": &dcl.Property{
																		Type:        "object",
																		GoName:      "Proximity",
																		GoType:      "InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity",
																		Description: "Proximity of the finding within which the entire hotword must reside. The total length of the window cannot exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be used to match substrings of the finding itself. For example, the certainty of a phone number regex \"(d{3}) d{3}-d{4}\" could be adjusted upwards if the area code is known to be the local area code of a company office using the hotword regex \"(xxx)\", where \"xxx\" is the area code in question.",
																		Properties: map[string]*dcl.Property{
																			"windowAfter": &dcl.Property{
																				Type:        "integer",
																				Format:      "int64",
																				GoName:      "WindowAfter",
																				Description: "Number of characters after the finding to consider.",
																			},
																			"windowBefore": &dcl.Property{
																				Type:        "integer",
																				Format:      "int64",
																				GoName:      "WindowBefore",
																				Description: "Number of characters before the finding to consider.",
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location of the resource",
								Immutable:   true,
							},
							"locationId": &dcl.Property{
								Type:        "string",
								GoName:      "LocationId",
								ReadOnly:    true,
								Description: "Output only. The geographic location where this resource is stored.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "The template name. The template will have one of the following formats: `projects/PROJECT_ID/inspectTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/inspectTemplates/TEMPLATE_ID`;",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"parent": &dcl.Property{
								Type:                "string",
								GoName:              "Parent",
								Description:         "The parent of the resource",
								Immutable:           true,
								ForwardSlashAllowed: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Organization",
										Field:    "name",
										Parent:   true,
									},
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The last update timestamp of an inspectTemplate.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
