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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/billingbudgets/beta/billingbudgets_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets/beta"
)

// BudgetServer implements the gRPC interface for Budget.
type BudgetServer struct{}

// ProtoToBudgetBudgetFilterCreditTypesTreatmentEnum converts a BudgetBudgetFilterCreditTypesTreatmentEnum enum from its proto representation.
func ProtoToBillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum(e betapb.BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum) *beta.BudgetBudgetFilterCreditTypesTreatmentEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum_name[int32(e)]; ok {
		e := beta.BudgetBudgetFilterCreditTypesTreatmentEnum(n[len("BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetBudgetFilterCalendarPeriodEnum converts a BudgetBudgetFilterCalendarPeriodEnum enum from its proto representation.
func ProtoToBillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum(e betapb.BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum) *beta.BudgetBudgetFilterCalendarPeriodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum_name[int32(e)]; ok {
		e := beta.BudgetBudgetFilterCalendarPeriodEnum(n[len("BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetThresholdRulesSpendBasisEnum converts a BudgetThresholdRulesSpendBasisEnum enum from its proto representation.
func ProtoToBillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum(e betapb.BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum) *beta.BudgetThresholdRulesSpendBasisEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum_name[int32(e)]; ok {
		e := beta.BudgetThresholdRulesSpendBasisEnum(n[len("BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetBudgetFilter converts a BudgetBudgetFilter object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetBudgetFilter(p *betapb.BillingbudgetsBetaBudgetBudgetFilter) *beta.BudgetBudgetFilter {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetBudgetFilter{
		CreditTypesTreatment: ProtoToBillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum(p.GetCreditTypesTreatment()),
		CalendarPeriod:       ProtoToBillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum(p.GetCalendarPeriod()),
		CustomPeriod:         ProtoToBillingbudgetsBetaBudgetBudgetFilterCustomPeriod(p.GetCustomPeriod()),
	}
	for _, r := range p.GetProjects() {
		obj.Projects = append(obj.Projects, r)
	}
	for _, r := range p.GetCreditTypes() {
		obj.CreditTypes = append(obj.CreditTypes, r)
	}
	for _, r := range p.GetServices() {
		obj.Services = append(obj.Services, r)
	}
	for _, r := range p.GetSubaccounts() {
		obj.Subaccounts = append(obj.Subaccounts, r)
	}
	return obj
}

// ProtoToBudgetBudgetFilterLabels converts a BudgetBudgetFilterLabels object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetBudgetFilterLabels(p *betapb.BillingbudgetsBetaBudgetBudgetFilterLabels) *beta.BudgetBudgetFilterLabels {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetBudgetFilterLabels{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriod converts a BudgetBudgetFilterCustomPeriod object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetBudgetFilterCustomPeriod(p *betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriod) *beta.BudgetBudgetFilterCustomPeriod {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetBudgetFilterCustomPeriod{
		StartDate: ProtoToBillingbudgetsBetaBudgetBudgetFilterCustomPeriodStartDate(p.GetStartDate()),
		EndDate:   ProtoToBillingbudgetsBetaBudgetBudgetFilterCustomPeriodEndDate(p.GetEndDate()),
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriodStartDate converts a BudgetBudgetFilterCustomPeriodStartDate object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetBudgetFilterCustomPeriodStartDate(p *betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriodStartDate) *beta.BudgetBudgetFilterCustomPeriodStartDate {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetBudgetFilterCustomPeriodStartDate{
		Year:  dcl.Int64OrNil(p.GetYear()),
		Month: dcl.Int64OrNil(p.GetMonth()),
		Day:   dcl.Int64OrNil(p.GetDay()),
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriodEndDate converts a BudgetBudgetFilterCustomPeriodEndDate object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetBudgetFilterCustomPeriodEndDate(p *betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriodEndDate) *beta.BudgetBudgetFilterCustomPeriodEndDate {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetBudgetFilterCustomPeriodEndDate{
		Year:  dcl.Int64OrNil(p.GetYear()),
		Month: dcl.Int64OrNil(p.GetMonth()),
		Day:   dcl.Int64OrNil(p.GetDay()),
	}
	return obj
}

// ProtoToBudgetAmount converts a BudgetAmount object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetAmount(p *betapb.BillingbudgetsBetaBudgetAmount) *beta.BudgetAmount {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetAmount{
		SpecifiedAmount:  ProtoToBillingbudgetsBetaBudgetAmountSpecifiedAmount(p.GetSpecifiedAmount()),
		LastPeriodAmount: ProtoToBillingbudgetsBetaBudgetAmountLastPeriodAmount(p.GetLastPeriodAmount()),
	}
	return obj
}

// ProtoToBudgetAmountSpecifiedAmount converts a BudgetAmountSpecifiedAmount object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetAmountSpecifiedAmount(p *betapb.BillingbudgetsBetaBudgetAmountSpecifiedAmount) *beta.BudgetAmountSpecifiedAmount {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetAmountSpecifiedAmount{
		CurrencyCode: dcl.StringOrNil(p.GetCurrencyCode()),
		Units:        dcl.Int64OrNil(p.GetUnits()),
		Nanos:        dcl.Int64OrNil(p.GetNanos()),
	}
	return obj
}

// ProtoToBudgetAmountLastPeriodAmount converts a BudgetAmountLastPeriodAmount object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetAmountLastPeriodAmount(p *betapb.BillingbudgetsBetaBudgetAmountLastPeriodAmount) *beta.BudgetAmountLastPeriodAmount {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetAmountLastPeriodAmount{}
	return obj
}

// ProtoToBudgetThresholdRules converts a BudgetThresholdRules object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetThresholdRules(p *betapb.BillingbudgetsBetaBudgetThresholdRules) *beta.BudgetThresholdRules {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetThresholdRules{
		ThresholdPercent: dcl.Float64OrNil(p.GetThresholdPercent()),
		SpendBasis:       ProtoToBillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum(p.GetSpendBasis()),
	}
	return obj
}

// ProtoToBudgetAllUpdatesRule converts a BudgetAllUpdatesRule object from its proto representation.
func ProtoToBillingbudgetsBetaBudgetAllUpdatesRule(p *betapb.BillingbudgetsBetaBudgetAllUpdatesRule) *beta.BudgetAllUpdatesRule {
	if p == nil {
		return nil
	}
	obj := &beta.BudgetAllUpdatesRule{
		PubsubTopic:                 dcl.StringOrNil(p.GetPubsubTopic()),
		SchemaVersion:               dcl.StringOrNil(p.GetSchemaVersion()),
		DisableDefaultIamRecipients: dcl.Bool(p.GetDisableDefaultIamRecipients()),
	}
	for _, r := range p.GetMonitoringNotificationChannels() {
		obj.MonitoringNotificationChannels = append(obj.MonitoringNotificationChannels, r)
	}
	return obj
}

// ProtoToBudget converts a Budget resource from its proto representation.
func ProtoToBudget(p *betapb.BillingbudgetsBetaBudget) *beta.Budget {
	obj := &beta.Budget{
		Name:           dcl.StringOrNil(p.GetName()),
		DisplayName:    dcl.StringOrNil(p.GetDisplayName()),
		BudgetFilter:   ProtoToBillingbudgetsBetaBudgetBudgetFilter(p.GetBudgetFilter()),
		Amount:         ProtoToBillingbudgetsBetaBudgetAmount(p.GetAmount()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		AllUpdatesRule: ProtoToBillingbudgetsBetaBudgetAllUpdatesRule(p.GetAllUpdatesRule()),
		BillingAccount: dcl.StringOrNil(p.GetBillingAccount()),
	}
	for _, r := range p.GetThresholdRules() {
		obj.ThresholdRules = append(obj.ThresholdRules, *ProtoToBillingbudgetsBetaBudgetThresholdRules(r))
	}
	return obj
}

// BudgetBudgetFilterCreditTypesTreatmentEnumToProto converts a BudgetBudgetFilterCreditTypesTreatmentEnum enum to its proto representation.
func BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnumToProto(e *beta.BudgetBudgetFilterCreditTypesTreatmentEnum) betapb.BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum {
	if e == nil {
		return betapb.BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum(0)
	}
	if v, ok := betapb.BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum_value["BudgetBudgetFilterCreditTypesTreatmentEnum"+string(*e)]; ok {
		return betapb.BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum(v)
	}
	return betapb.BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnum(0)
}

// BudgetBudgetFilterCalendarPeriodEnumToProto converts a BudgetBudgetFilterCalendarPeriodEnum enum to its proto representation.
func BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnumToProto(e *beta.BudgetBudgetFilterCalendarPeriodEnum) betapb.BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum {
	if e == nil {
		return betapb.BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum(0)
	}
	if v, ok := betapb.BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum_value["BudgetBudgetFilterCalendarPeriodEnum"+string(*e)]; ok {
		return betapb.BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum(v)
	}
	return betapb.BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnum(0)
}

// BudgetThresholdRulesSpendBasisEnumToProto converts a BudgetThresholdRulesSpendBasisEnum enum to its proto representation.
func BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnumToProto(e *beta.BudgetThresholdRulesSpendBasisEnum) betapb.BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum {
	if e == nil {
		return betapb.BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum(0)
	}
	if v, ok := betapb.BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum_value["BudgetThresholdRulesSpendBasisEnum"+string(*e)]; ok {
		return betapb.BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum(v)
	}
	return betapb.BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnum(0)
}

// BudgetBudgetFilterToProto converts a BudgetBudgetFilter object to its proto representation.
func BillingbudgetsBetaBudgetBudgetFilterToProto(o *beta.BudgetBudgetFilter) *betapb.BillingbudgetsBetaBudgetBudgetFilter {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetBudgetFilter{}
	p.SetCreditTypesTreatment(BillingbudgetsBetaBudgetBudgetFilterCreditTypesTreatmentEnumToProto(o.CreditTypesTreatment))
	p.SetCalendarPeriod(BillingbudgetsBetaBudgetBudgetFilterCalendarPeriodEnumToProto(o.CalendarPeriod))
	p.SetCustomPeriod(BillingbudgetsBetaBudgetBudgetFilterCustomPeriodToProto(o.CustomPeriod))
	sProjects := make([]string, len(o.Projects))
	for i, r := range o.Projects {
		sProjects[i] = r
	}
	p.SetProjects(sProjects)
	sCreditTypes := make([]string, len(o.CreditTypes))
	for i, r := range o.CreditTypes {
		sCreditTypes[i] = r
	}
	p.SetCreditTypes(sCreditTypes)
	sServices := make([]string, len(o.Services))
	for i, r := range o.Services {
		sServices[i] = r
	}
	p.SetServices(sServices)
	sSubaccounts := make([]string, len(o.Subaccounts))
	for i, r := range o.Subaccounts {
		sSubaccounts[i] = r
	}
	p.SetSubaccounts(sSubaccounts)
	mLabels := make(map[string]*betapb.BillingbudgetsBetaBudgetBudgetFilterLabels, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = BillingbudgetsBetaBudgetBudgetFilterLabelsToProto(&r)
	}
	p.SetLabels(mLabels)
	return p
}

// BudgetBudgetFilterLabelsToProto converts a BudgetBudgetFilterLabels object to its proto representation.
func BillingbudgetsBetaBudgetBudgetFilterLabelsToProto(o *beta.BudgetBudgetFilterLabels) *betapb.BillingbudgetsBetaBudgetBudgetFilterLabels {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetBudgetFilterLabels{}
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// BudgetBudgetFilterCustomPeriodToProto converts a BudgetBudgetFilterCustomPeriod object to its proto representation.
func BillingbudgetsBetaBudgetBudgetFilterCustomPeriodToProto(o *beta.BudgetBudgetFilterCustomPeriod) *betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriod {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriod{}
	p.SetStartDate(BillingbudgetsBetaBudgetBudgetFilterCustomPeriodStartDateToProto(o.StartDate))
	p.SetEndDate(BillingbudgetsBetaBudgetBudgetFilterCustomPeriodEndDateToProto(o.EndDate))
	return p
}

// BudgetBudgetFilterCustomPeriodStartDateToProto converts a BudgetBudgetFilterCustomPeriodStartDate object to its proto representation.
func BillingbudgetsBetaBudgetBudgetFilterCustomPeriodStartDateToProto(o *beta.BudgetBudgetFilterCustomPeriodStartDate) *betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriodStartDate {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriodStartDate{}
	p.SetYear(dcl.ValueOrEmptyInt64(o.Year))
	p.SetMonth(dcl.ValueOrEmptyInt64(o.Month))
	p.SetDay(dcl.ValueOrEmptyInt64(o.Day))
	return p
}

// BudgetBudgetFilterCustomPeriodEndDateToProto converts a BudgetBudgetFilterCustomPeriodEndDate object to its proto representation.
func BillingbudgetsBetaBudgetBudgetFilterCustomPeriodEndDateToProto(o *beta.BudgetBudgetFilterCustomPeriodEndDate) *betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriodEndDate {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetBudgetFilterCustomPeriodEndDate{}
	p.SetYear(dcl.ValueOrEmptyInt64(o.Year))
	p.SetMonth(dcl.ValueOrEmptyInt64(o.Month))
	p.SetDay(dcl.ValueOrEmptyInt64(o.Day))
	return p
}

// BudgetAmountToProto converts a BudgetAmount object to its proto representation.
func BillingbudgetsBetaBudgetAmountToProto(o *beta.BudgetAmount) *betapb.BillingbudgetsBetaBudgetAmount {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetAmount{}
	p.SetSpecifiedAmount(BillingbudgetsBetaBudgetAmountSpecifiedAmountToProto(o.SpecifiedAmount))
	p.SetLastPeriodAmount(BillingbudgetsBetaBudgetAmountLastPeriodAmountToProto(o.LastPeriodAmount))
	return p
}

// BudgetAmountSpecifiedAmountToProto converts a BudgetAmountSpecifiedAmount object to its proto representation.
func BillingbudgetsBetaBudgetAmountSpecifiedAmountToProto(o *beta.BudgetAmountSpecifiedAmount) *betapb.BillingbudgetsBetaBudgetAmountSpecifiedAmount {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetAmountSpecifiedAmount{}
	p.SetCurrencyCode(dcl.ValueOrEmptyString(o.CurrencyCode))
	p.SetUnits(dcl.ValueOrEmptyInt64(o.Units))
	p.SetNanos(dcl.ValueOrEmptyInt64(o.Nanos))
	return p
}

// BudgetAmountLastPeriodAmountToProto converts a BudgetAmountLastPeriodAmount object to its proto representation.
func BillingbudgetsBetaBudgetAmountLastPeriodAmountToProto(o *beta.BudgetAmountLastPeriodAmount) *betapb.BillingbudgetsBetaBudgetAmountLastPeriodAmount {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetAmountLastPeriodAmount{}
	return p
}

// BudgetThresholdRulesToProto converts a BudgetThresholdRules object to its proto representation.
func BillingbudgetsBetaBudgetThresholdRulesToProto(o *beta.BudgetThresholdRules) *betapb.BillingbudgetsBetaBudgetThresholdRules {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetThresholdRules{}
	p.SetThresholdPercent(dcl.ValueOrEmptyDouble(o.ThresholdPercent))
	p.SetSpendBasis(BillingbudgetsBetaBudgetThresholdRulesSpendBasisEnumToProto(o.SpendBasis))
	return p
}

// BudgetAllUpdatesRuleToProto converts a BudgetAllUpdatesRule object to its proto representation.
func BillingbudgetsBetaBudgetAllUpdatesRuleToProto(o *beta.BudgetAllUpdatesRule) *betapb.BillingbudgetsBetaBudgetAllUpdatesRule {
	if o == nil {
		return nil
	}
	p := &betapb.BillingbudgetsBetaBudgetAllUpdatesRule{}
	p.SetPubsubTopic(dcl.ValueOrEmptyString(o.PubsubTopic))
	p.SetSchemaVersion(dcl.ValueOrEmptyString(o.SchemaVersion))
	p.SetDisableDefaultIamRecipients(dcl.ValueOrEmptyBool(o.DisableDefaultIamRecipients))
	sMonitoringNotificationChannels := make([]string, len(o.MonitoringNotificationChannels))
	for i, r := range o.MonitoringNotificationChannels {
		sMonitoringNotificationChannels[i] = r
	}
	p.SetMonitoringNotificationChannels(sMonitoringNotificationChannels)
	return p
}

// BudgetToProto converts a Budget resource to its proto representation.
func BudgetToProto(resource *beta.Budget) *betapb.BillingbudgetsBetaBudget {
	p := &betapb.BillingbudgetsBetaBudget{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetBudgetFilter(BillingbudgetsBetaBudgetBudgetFilterToProto(resource.BudgetFilter))
	p.SetAmount(BillingbudgetsBetaBudgetAmountToProto(resource.Amount))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetAllUpdatesRule(BillingbudgetsBetaBudgetAllUpdatesRuleToProto(resource.AllUpdatesRule))
	p.SetBillingAccount(dcl.ValueOrEmptyString(resource.BillingAccount))
	sThresholdRules := make([]*betapb.BillingbudgetsBetaBudgetThresholdRules, len(resource.ThresholdRules))
	for i, r := range resource.ThresholdRules {
		sThresholdRules[i] = BillingbudgetsBetaBudgetThresholdRulesToProto(&r)
	}
	p.SetThresholdRules(sThresholdRules)

	return p
}

// applyBudget handles the gRPC request by passing it to the underlying Budget Apply() method.
func (s *BudgetServer) applyBudget(ctx context.Context, c *beta.Client, request *betapb.ApplyBillingbudgetsBetaBudgetRequest) (*betapb.BillingbudgetsBetaBudget, error) {
	p := ProtoToBudget(request.GetResource())
	res, err := c.ApplyBudget(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BudgetToProto(res)
	return r, nil
}

// applyBillingbudgetsBetaBudget handles the gRPC request by passing it to the underlying Budget Apply() method.
func (s *BudgetServer) ApplyBillingbudgetsBetaBudget(ctx context.Context, request *betapb.ApplyBillingbudgetsBetaBudgetRequest) (*betapb.BillingbudgetsBetaBudget, error) {
	cl, err := createConfigBudget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyBudget(ctx, cl, request)
}

// DeleteBudget handles the gRPC request by passing it to the underlying Budget Delete() method.
func (s *BudgetServer) DeleteBillingbudgetsBetaBudget(ctx context.Context, request *betapb.DeleteBillingbudgetsBetaBudgetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBudget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBudget(ctx, ProtoToBudget(request.GetResource()))

}

// ListBillingbudgetsBetaBudget handles the gRPC request by passing it to the underlying BudgetList() method.
func (s *BudgetServer) ListBillingbudgetsBetaBudget(ctx context.Context, request *betapb.ListBillingbudgetsBetaBudgetRequest) (*betapb.ListBillingbudgetsBetaBudgetResponse, error) {
	cl, err := createConfigBudget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBudget(ctx, request.GetBillingAccount())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.BillingbudgetsBetaBudget
	for _, r := range resources.Items {
		rp := BudgetToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListBillingbudgetsBetaBudgetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigBudget(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
