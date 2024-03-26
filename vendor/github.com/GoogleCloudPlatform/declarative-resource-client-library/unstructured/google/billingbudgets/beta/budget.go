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
package billingbudgets

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Budget struct{}

func BudgetToUnstructured(r *dclService.Budget) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "billingbudgets",
			Version: "beta",
			Type:    "Budget",
		},
		Object: make(map[string]interface{}),
	}
	if r.AllUpdatesRule != nil && r.AllUpdatesRule != dclService.EmptyBudgetAllUpdatesRule {
		rAllUpdatesRule := make(map[string]interface{})
		if r.AllUpdatesRule.DisableDefaultIamRecipients != nil {
			rAllUpdatesRule["disableDefaultIamRecipients"] = *r.AllUpdatesRule.DisableDefaultIamRecipients
		}
		var rAllUpdatesRuleMonitoringNotificationChannels []interface{}
		for _, rAllUpdatesRuleMonitoringNotificationChannelsVal := range r.AllUpdatesRule.MonitoringNotificationChannels {
			rAllUpdatesRuleMonitoringNotificationChannels = append(rAllUpdatesRuleMonitoringNotificationChannels, rAllUpdatesRuleMonitoringNotificationChannelsVal)
		}
		rAllUpdatesRule["monitoringNotificationChannels"] = rAllUpdatesRuleMonitoringNotificationChannels
		if r.AllUpdatesRule.PubsubTopic != nil {
			rAllUpdatesRule["pubsubTopic"] = *r.AllUpdatesRule.PubsubTopic
		}
		if r.AllUpdatesRule.SchemaVersion != nil {
			rAllUpdatesRule["schemaVersion"] = *r.AllUpdatesRule.SchemaVersion
		}
		u.Object["allUpdatesRule"] = rAllUpdatesRule
	}
	if r.Amount != nil && r.Amount != dclService.EmptyBudgetAmount {
		rAmount := make(map[string]interface{})
		if r.Amount.LastPeriodAmount != nil && r.Amount.LastPeriodAmount != dclService.EmptyBudgetAmountLastPeriodAmount {
			rAmountLastPeriodAmount := make(map[string]interface{})
			rAmount["lastPeriodAmount"] = rAmountLastPeriodAmount
		}
		if r.Amount.SpecifiedAmount != nil && r.Amount.SpecifiedAmount != dclService.EmptyBudgetAmountSpecifiedAmount {
			rAmountSpecifiedAmount := make(map[string]interface{})
			if r.Amount.SpecifiedAmount.CurrencyCode != nil {
				rAmountSpecifiedAmount["currencyCode"] = *r.Amount.SpecifiedAmount.CurrencyCode
			}
			if r.Amount.SpecifiedAmount.Nanos != nil {
				rAmountSpecifiedAmount["nanos"] = *r.Amount.SpecifiedAmount.Nanos
			}
			if r.Amount.SpecifiedAmount.Units != nil {
				rAmountSpecifiedAmount["units"] = *r.Amount.SpecifiedAmount.Units
			}
			rAmount["specifiedAmount"] = rAmountSpecifiedAmount
		}
		u.Object["amount"] = rAmount
	}
	if r.BillingAccount != nil {
		u.Object["billingAccount"] = *r.BillingAccount
	}
	if r.BudgetFilter != nil && r.BudgetFilter != dclService.EmptyBudgetBudgetFilter {
		rBudgetFilter := make(map[string]interface{})
		if r.BudgetFilter.CalendarPeriod != nil {
			rBudgetFilter["calendarPeriod"] = string(*r.BudgetFilter.CalendarPeriod)
		}
		var rBudgetFilterCreditTypes []interface{}
		for _, rBudgetFilterCreditTypesVal := range r.BudgetFilter.CreditTypes {
			rBudgetFilterCreditTypes = append(rBudgetFilterCreditTypes, rBudgetFilterCreditTypesVal)
		}
		rBudgetFilter["creditTypes"] = rBudgetFilterCreditTypes
		if r.BudgetFilter.CreditTypesTreatment != nil {
			rBudgetFilter["creditTypesTreatment"] = string(*r.BudgetFilter.CreditTypesTreatment)
		}
		if r.BudgetFilter.CustomPeriod != nil && r.BudgetFilter.CustomPeriod != dclService.EmptyBudgetBudgetFilterCustomPeriod {
			rBudgetFilterCustomPeriod := make(map[string]interface{})
			if r.BudgetFilter.CustomPeriod.EndDate != nil && r.BudgetFilter.CustomPeriod.EndDate != dclService.EmptyBudgetBudgetFilterCustomPeriodEndDate {
				rBudgetFilterCustomPeriodEndDate := make(map[string]interface{})
				if r.BudgetFilter.CustomPeriod.EndDate.Day != nil {
					rBudgetFilterCustomPeriodEndDate["day"] = *r.BudgetFilter.CustomPeriod.EndDate.Day
				}
				if r.BudgetFilter.CustomPeriod.EndDate.Month != nil {
					rBudgetFilterCustomPeriodEndDate["month"] = *r.BudgetFilter.CustomPeriod.EndDate.Month
				}
				if r.BudgetFilter.CustomPeriod.EndDate.Year != nil {
					rBudgetFilterCustomPeriodEndDate["year"] = *r.BudgetFilter.CustomPeriod.EndDate.Year
				}
				rBudgetFilterCustomPeriod["endDate"] = rBudgetFilterCustomPeriodEndDate
			}
			if r.BudgetFilter.CustomPeriod.StartDate != nil && r.BudgetFilter.CustomPeriod.StartDate != dclService.EmptyBudgetBudgetFilterCustomPeriodStartDate {
				rBudgetFilterCustomPeriodStartDate := make(map[string]interface{})
				if r.BudgetFilter.CustomPeriod.StartDate.Day != nil {
					rBudgetFilterCustomPeriodStartDate["day"] = *r.BudgetFilter.CustomPeriod.StartDate.Day
				}
				if r.BudgetFilter.CustomPeriod.StartDate.Month != nil {
					rBudgetFilterCustomPeriodStartDate["month"] = *r.BudgetFilter.CustomPeriod.StartDate.Month
				}
				if r.BudgetFilter.CustomPeriod.StartDate.Year != nil {
					rBudgetFilterCustomPeriodStartDate["year"] = *r.BudgetFilter.CustomPeriod.StartDate.Year
				}
				rBudgetFilterCustomPeriod["startDate"] = rBudgetFilterCustomPeriodStartDate
			}
			rBudgetFilter["customPeriod"] = rBudgetFilterCustomPeriod
		}
		if r.BudgetFilter.Labels != nil {
			rBudgetFilterLabels := make(map[string]interface{})
			for k, v := range r.BudgetFilter.Labels {
				rBudgetFilterLabelsMap := make(map[string]interface{})
				var vValues []interface{}
				for _, vValuesVal := range v.Values {
					vValues = append(vValues, vValuesVal)
				}
				rBudgetFilterLabelsMap["values"] = vValues
				rBudgetFilterLabels[k] = rBudgetFilterLabelsMap
			}
			rBudgetFilter["labels"] = rBudgetFilterLabels
		}
		var rBudgetFilterProjects []interface{}
		for _, rBudgetFilterProjectsVal := range r.BudgetFilter.Projects {
			rBudgetFilterProjects = append(rBudgetFilterProjects, rBudgetFilterProjectsVal)
		}
		rBudgetFilter["projects"] = rBudgetFilterProjects
		var rBudgetFilterServices []interface{}
		for _, rBudgetFilterServicesVal := range r.BudgetFilter.Services {
			rBudgetFilterServices = append(rBudgetFilterServices, rBudgetFilterServicesVal)
		}
		rBudgetFilter["services"] = rBudgetFilterServices
		var rBudgetFilterSubaccounts []interface{}
		for _, rBudgetFilterSubaccountsVal := range r.BudgetFilter.Subaccounts {
			rBudgetFilterSubaccounts = append(rBudgetFilterSubaccounts, rBudgetFilterSubaccountsVal)
		}
		rBudgetFilter["subaccounts"] = rBudgetFilterSubaccounts
		u.Object["budgetFilter"] = rBudgetFilter
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	var rThresholdRules []interface{}
	for _, rThresholdRulesVal := range r.ThresholdRules {
		rThresholdRulesObject := make(map[string]interface{})
		if rThresholdRulesVal.SpendBasis != nil {
			rThresholdRulesObject["spendBasis"] = string(*rThresholdRulesVal.SpendBasis)
		}
		if rThresholdRulesVal.ThresholdPercent != nil {
			rThresholdRulesObject["thresholdPercent"] = *rThresholdRulesVal.ThresholdPercent
		}
		rThresholdRules = append(rThresholdRules, rThresholdRulesObject)
	}
	u.Object["thresholdRules"] = rThresholdRules
	return u
}

func UnstructuredToBudget(u *unstructured.Resource) (*dclService.Budget, error) {
	r := &dclService.Budget{}
	if _, ok := u.Object["allUpdatesRule"]; ok {
		if rAllUpdatesRule, ok := u.Object["allUpdatesRule"].(map[string]interface{}); ok {
			r.AllUpdatesRule = &dclService.BudgetAllUpdatesRule{}
			if _, ok := rAllUpdatesRule["disableDefaultIamRecipients"]; ok {
				if b, ok := rAllUpdatesRule["disableDefaultIamRecipients"].(bool); ok {
					r.AllUpdatesRule.DisableDefaultIamRecipients = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.AllUpdatesRule.DisableDefaultIamRecipients: expected bool")
				}
			}
			if _, ok := rAllUpdatesRule["monitoringNotificationChannels"]; ok {
				if s, ok := rAllUpdatesRule["monitoringNotificationChannels"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.AllUpdatesRule.MonitoringNotificationChannels = append(r.AllUpdatesRule.MonitoringNotificationChannels, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.AllUpdatesRule.MonitoringNotificationChannels: expected []interface{}")
				}
			}
			if _, ok := rAllUpdatesRule["pubsubTopic"]; ok {
				if s, ok := rAllUpdatesRule["pubsubTopic"].(string); ok {
					r.AllUpdatesRule.PubsubTopic = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AllUpdatesRule.PubsubTopic: expected string")
				}
			}
			if _, ok := rAllUpdatesRule["schemaVersion"]; ok {
				if s, ok := rAllUpdatesRule["schemaVersion"].(string); ok {
					r.AllUpdatesRule.SchemaVersion = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.AllUpdatesRule.SchemaVersion: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.AllUpdatesRule: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["amount"]; ok {
		if rAmount, ok := u.Object["amount"].(map[string]interface{}); ok {
			r.Amount = &dclService.BudgetAmount{}
			if _, ok := rAmount["lastPeriodAmount"]; ok {
				if _, ok := rAmount["lastPeriodAmount"].(map[string]interface{}); ok {
					r.Amount.LastPeriodAmount = &dclService.BudgetAmountLastPeriodAmount{}
				} else {
					return nil, fmt.Errorf("r.Amount.LastPeriodAmount: expected map[string]interface{}")
				}
			}
			if _, ok := rAmount["specifiedAmount"]; ok {
				if rAmountSpecifiedAmount, ok := rAmount["specifiedAmount"].(map[string]interface{}); ok {
					r.Amount.SpecifiedAmount = &dclService.BudgetAmountSpecifiedAmount{}
					if _, ok := rAmountSpecifiedAmount["currencyCode"]; ok {
						if s, ok := rAmountSpecifiedAmount["currencyCode"].(string); ok {
							r.Amount.SpecifiedAmount.CurrencyCode = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Amount.SpecifiedAmount.CurrencyCode: expected string")
						}
					}
					if _, ok := rAmountSpecifiedAmount["nanos"]; ok {
						if i, ok := rAmountSpecifiedAmount["nanos"].(int64); ok {
							r.Amount.SpecifiedAmount.Nanos = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Amount.SpecifiedAmount.Nanos: expected int64")
						}
					}
					if _, ok := rAmountSpecifiedAmount["units"]; ok {
						if i, ok := rAmountSpecifiedAmount["units"].(int64); ok {
							r.Amount.SpecifiedAmount.Units = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Amount.SpecifiedAmount.Units: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Amount.SpecifiedAmount: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Amount: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["billingAccount"]; ok {
		if s, ok := u.Object["billingAccount"].(string); ok {
			r.BillingAccount = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.BillingAccount: expected string")
		}
	}
	if _, ok := u.Object["budgetFilter"]; ok {
		if rBudgetFilter, ok := u.Object["budgetFilter"].(map[string]interface{}); ok {
			r.BudgetFilter = &dclService.BudgetBudgetFilter{}
			if _, ok := rBudgetFilter["calendarPeriod"]; ok {
				if s, ok := rBudgetFilter["calendarPeriod"].(string); ok {
					r.BudgetFilter.CalendarPeriod = dclService.BudgetBudgetFilterCalendarPeriodEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.BudgetFilter.CalendarPeriod: expected string")
				}
			}
			if _, ok := rBudgetFilter["creditTypes"]; ok {
				if s, ok := rBudgetFilter["creditTypes"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.BudgetFilter.CreditTypes = append(r.BudgetFilter.CreditTypes, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.BudgetFilter.CreditTypes: expected []interface{}")
				}
			}
			if _, ok := rBudgetFilter["creditTypesTreatment"]; ok {
				if s, ok := rBudgetFilter["creditTypesTreatment"].(string); ok {
					r.BudgetFilter.CreditTypesTreatment = dclService.BudgetBudgetFilterCreditTypesTreatmentEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.BudgetFilter.CreditTypesTreatment: expected string")
				}
			}
			if _, ok := rBudgetFilter["customPeriod"]; ok {
				if rBudgetFilterCustomPeriod, ok := rBudgetFilter["customPeriod"].(map[string]interface{}); ok {
					r.BudgetFilter.CustomPeriod = &dclService.BudgetBudgetFilterCustomPeriod{}
					if _, ok := rBudgetFilterCustomPeriod["endDate"]; ok {
						if rBudgetFilterCustomPeriodEndDate, ok := rBudgetFilterCustomPeriod["endDate"].(map[string]interface{}); ok {
							r.BudgetFilter.CustomPeriod.EndDate = &dclService.BudgetBudgetFilterCustomPeriodEndDate{}
							if _, ok := rBudgetFilterCustomPeriodEndDate["day"]; ok {
								if i, ok := rBudgetFilterCustomPeriodEndDate["day"].(int64); ok {
									r.BudgetFilter.CustomPeriod.EndDate.Day = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod.EndDate.Day: expected int64")
								}
							}
							if _, ok := rBudgetFilterCustomPeriodEndDate["month"]; ok {
								if i, ok := rBudgetFilterCustomPeriodEndDate["month"].(int64); ok {
									r.BudgetFilter.CustomPeriod.EndDate.Month = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod.EndDate.Month: expected int64")
								}
							}
							if _, ok := rBudgetFilterCustomPeriodEndDate["year"]; ok {
								if i, ok := rBudgetFilterCustomPeriodEndDate["year"].(int64); ok {
									r.BudgetFilter.CustomPeriod.EndDate.Year = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod.EndDate.Year: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod.EndDate: expected map[string]interface{}")
						}
					}
					if _, ok := rBudgetFilterCustomPeriod["startDate"]; ok {
						if rBudgetFilterCustomPeriodStartDate, ok := rBudgetFilterCustomPeriod["startDate"].(map[string]interface{}); ok {
							r.BudgetFilter.CustomPeriod.StartDate = &dclService.BudgetBudgetFilterCustomPeriodStartDate{}
							if _, ok := rBudgetFilterCustomPeriodStartDate["day"]; ok {
								if i, ok := rBudgetFilterCustomPeriodStartDate["day"].(int64); ok {
									r.BudgetFilter.CustomPeriod.StartDate.Day = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod.StartDate.Day: expected int64")
								}
							}
							if _, ok := rBudgetFilterCustomPeriodStartDate["month"]; ok {
								if i, ok := rBudgetFilterCustomPeriodStartDate["month"].(int64); ok {
									r.BudgetFilter.CustomPeriod.StartDate.Month = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod.StartDate.Month: expected int64")
								}
							}
							if _, ok := rBudgetFilterCustomPeriodStartDate["year"]; ok {
								if i, ok := rBudgetFilterCustomPeriodStartDate["year"].(int64); ok {
									r.BudgetFilter.CustomPeriod.StartDate.Year = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod.StartDate.Year: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod.StartDate: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.BudgetFilter.CustomPeriod: expected map[string]interface{}")
				}
			}
			if _, ok := rBudgetFilter["labels"]; ok {
				if rBudgetFilterLabels, ok := rBudgetFilter["labels"].(map[string]interface{}); ok {
					m := make(map[string]dclService.BudgetBudgetFilterLabels)
					for k, v := range rBudgetFilterLabels {
						if objval, ok := v.(map[string]interface{}); ok {
							var rBudgetFilterLabelsObj dclService.BudgetBudgetFilterLabels
							if _, ok := objval["values"]; ok {
								if s, ok := objval["values"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rBudgetFilterLabelsObj.Values = append(rBudgetFilterLabelsObj.Values, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rBudgetFilterLabelsObj.Values: expected []interface{}")
								}
							}
							m[k] = rBudgetFilterLabelsObj
						} else {
							return nil, fmt.Errorf("r.BudgetFilter.Labels: expected map[string]interface{}")
						}
					}
					r.BudgetFilter.Labels = m
				} else {
					return nil, fmt.Errorf("r.BudgetFilter.Labels: expected map[string]interface{}")
				}
			}
			if _, ok := rBudgetFilter["projects"]; ok {
				if s, ok := rBudgetFilter["projects"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.BudgetFilter.Projects = append(r.BudgetFilter.Projects, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.BudgetFilter.Projects: expected []interface{}")
				}
			}
			if _, ok := rBudgetFilter["services"]; ok {
				if s, ok := rBudgetFilter["services"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.BudgetFilter.Services = append(r.BudgetFilter.Services, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.BudgetFilter.Services: expected []interface{}")
				}
			}
			if _, ok := rBudgetFilter["subaccounts"]; ok {
				if s, ok := rBudgetFilter["subaccounts"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.BudgetFilter.Subaccounts = append(r.BudgetFilter.Subaccounts, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.BudgetFilter.Subaccounts: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.BudgetFilter: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["thresholdRules"]; ok {
		if s, ok := u.Object["thresholdRules"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rThresholdRules dclService.BudgetThresholdRules
					if _, ok := objval["spendBasis"]; ok {
						if s, ok := objval["spendBasis"].(string); ok {
							rThresholdRules.SpendBasis = dclService.BudgetThresholdRulesSpendBasisEnumRef(s)
						} else {
							return nil, fmt.Errorf("rThresholdRules.SpendBasis: expected string")
						}
					}
					if _, ok := objval["thresholdPercent"]; ok {
						if f, ok := objval["thresholdPercent"].(float64); ok {
							rThresholdRules.ThresholdPercent = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("rThresholdRules.ThresholdPercent: expected float64")
						}
					}
					r.ThresholdRules = append(r.ThresholdRules, rThresholdRules)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ThresholdRules: expected []interface{}")
		}
	}
	return r, nil
}

func GetBudget(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBudget(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetBudget(ctx, r)
	if err != nil {
		return nil, err
	}
	return BudgetToUnstructured(r), nil
}

func ListBudget(ctx context.Context, config *dcl.Config, billingAccount string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListBudget(ctx, billingAccount)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, BudgetToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyBudget(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBudget(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToBudget(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyBudget(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return BudgetToUnstructured(r), nil
}

func BudgetHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBudget(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToBudget(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyBudget(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteBudget(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToBudget(u)
	if err != nil {
		return err
	}
	return c.DeleteBudget(ctx, r)
}

func BudgetID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToBudget(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Budget) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"billingbudgets",
		"Budget",
		"beta",
	}
}

func (r *Budget) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Budget) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Budget) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Budget) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Budget) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Budget) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Budget) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetBudget(ctx, config, resource)
}

func (r *Budget) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyBudget(ctx, config, resource, opts...)
}

func (r *Budget) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return BudgetHasDiff(ctx, config, resource, opts...)
}

func (r *Budget) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteBudget(ctx, config, resource)
}

func (r *Budget) ID(resource *unstructured.Resource) (string, error) {
	return BudgetID(resource)
}

func init() {
	unstructured.Register(&Budget{})
}
