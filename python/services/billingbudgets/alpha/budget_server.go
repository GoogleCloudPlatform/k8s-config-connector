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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/billingbudgets/alpha/billingbudgets_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/billingbudgets/alpha"
)

// BudgetServer implements the gRPC interface for Budget.
type BudgetServer struct{}

// ProtoToBudgetBudgetFilterCreditTypesTreatmentEnum converts a BudgetBudgetFilterCreditTypesTreatmentEnum enum from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum(e alphapb.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum) *alpha.BudgetBudgetFilterCreditTypesTreatmentEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum_name[int32(e)]; ok {
		e := alpha.BudgetBudgetFilterCreditTypesTreatmentEnum(n[len("BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetBudgetFilterCalendarPeriodEnum converts a BudgetBudgetFilterCalendarPeriodEnum enum from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum(e alphapb.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum) *alpha.BudgetBudgetFilterCalendarPeriodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum_name[int32(e)]; ok {
		e := alpha.BudgetBudgetFilterCalendarPeriodEnum(n[len("BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetThresholdRulesSpendBasisEnum converts a BudgetThresholdRulesSpendBasisEnum enum from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum(e alphapb.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum) *alpha.BudgetThresholdRulesSpendBasisEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum_name[int32(e)]; ok {
		e := alpha.BudgetThresholdRulesSpendBasisEnum(n[len("BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum"):])
		return &e
	}
	return nil
}

// ProtoToBudgetBudgetFilter converts a BudgetBudgetFilter object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetBudgetFilter(p *alphapb.BillingbudgetsAlphaBudgetBudgetFilter) *alpha.BudgetBudgetFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetBudgetFilter{
		CreditTypesTreatment: ProtoToBillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum(p.GetCreditTypesTreatment()),
		CalendarPeriod:       ProtoToBillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum(p.GetCalendarPeriod()),
		CustomPeriod:         ProtoToBillingbudgetsAlphaBudgetBudgetFilterCustomPeriod(p.GetCustomPeriod()),
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
func ProtoToBillingbudgetsAlphaBudgetBudgetFilterLabels(p *alphapb.BillingbudgetsAlphaBudgetBudgetFilterLabels) *alpha.BudgetBudgetFilterLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetBudgetFilterLabels{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriod converts a BudgetBudgetFilterCustomPeriod object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetBudgetFilterCustomPeriod(p *alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriod) *alpha.BudgetBudgetFilterCustomPeriod {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetBudgetFilterCustomPeriod{
		StartDate: ProtoToBillingbudgetsAlphaBudgetBudgetFilterCustomPeriodStartDate(p.GetStartDate()),
		EndDate:   ProtoToBillingbudgetsAlphaBudgetBudgetFilterCustomPeriodEndDate(p.GetEndDate()),
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriodStartDate converts a BudgetBudgetFilterCustomPeriodStartDate object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetBudgetFilterCustomPeriodStartDate(p *alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodStartDate) *alpha.BudgetBudgetFilterCustomPeriodStartDate {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetBudgetFilterCustomPeriodStartDate{
		Year:  dcl.Int64OrNil(p.GetYear()),
		Month: dcl.Int64OrNil(p.GetMonth()),
		Day:   dcl.Int64OrNil(p.GetDay()),
	}
	return obj
}

// ProtoToBudgetBudgetFilterCustomPeriodEndDate converts a BudgetBudgetFilterCustomPeriodEndDate object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetBudgetFilterCustomPeriodEndDate(p *alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodEndDate) *alpha.BudgetBudgetFilterCustomPeriodEndDate {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetBudgetFilterCustomPeriodEndDate{
		Year:  dcl.Int64OrNil(p.GetYear()),
		Month: dcl.Int64OrNil(p.GetMonth()),
		Day:   dcl.Int64OrNil(p.GetDay()),
	}
	return obj
}

// ProtoToBudgetAmount converts a BudgetAmount object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetAmount(p *alphapb.BillingbudgetsAlphaBudgetAmount) *alpha.BudgetAmount {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetAmount{
		SpecifiedAmount:  ProtoToBillingbudgetsAlphaBudgetAmountSpecifiedAmount(p.GetSpecifiedAmount()),
		LastPeriodAmount: ProtoToBillingbudgetsAlphaBudgetAmountLastPeriodAmount(p.GetLastPeriodAmount()),
	}
	return obj
}

// ProtoToBudgetAmountSpecifiedAmount converts a BudgetAmountSpecifiedAmount object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetAmountSpecifiedAmount(p *alphapb.BillingbudgetsAlphaBudgetAmountSpecifiedAmount) *alpha.BudgetAmountSpecifiedAmount {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetAmountSpecifiedAmount{
		CurrencyCode: dcl.StringOrNil(p.GetCurrencyCode()),
		Units:        dcl.Int64OrNil(p.GetUnits()),
		Nanos:        dcl.Int64OrNil(p.GetNanos()),
	}
	return obj
}

// ProtoToBudgetAmountLastPeriodAmount converts a BudgetAmountLastPeriodAmount object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetAmountLastPeriodAmount(p *alphapb.BillingbudgetsAlphaBudgetAmountLastPeriodAmount) *alpha.BudgetAmountLastPeriodAmount {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetAmountLastPeriodAmount{}
	return obj
}

// ProtoToBudgetThresholdRules converts a BudgetThresholdRules object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetThresholdRules(p *alphapb.BillingbudgetsAlphaBudgetThresholdRules) *alpha.BudgetThresholdRules {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetThresholdRules{
		ThresholdPercent: dcl.Float64OrNil(p.GetThresholdPercent()),
		SpendBasis:       ProtoToBillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum(p.GetSpendBasis()),
	}
	return obj
}

// ProtoToBudgetAllUpdatesRule converts a BudgetAllUpdatesRule object from its proto representation.
func ProtoToBillingbudgetsAlphaBudgetAllUpdatesRule(p *alphapb.BillingbudgetsAlphaBudgetAllUpdatesRule) *alpha.BudgetAllUpdatesRule {
	if p == nil {
		return nil
	}
	obj := &alpha.BudgetAllUpdatesRule{
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
func ProtoToBudget(p *alphapb.BillingbudgetsAlphaBudget) *alpha.Budget {
	obj := &alpha.Budget{
		Name:           dcl.StringOrNil(p.GetName()),
		DisplayName:    dcl.StringOrNil(p.GetDisplayName()),
		BudgetFilter:   ProtoToBillingbudgetsAlphaBudgetBudgetFilter(p.GetBudgetFilter()),
		Amount:         ProtoToBillingbudgetsAlphaBudgetAmount(p.GetAmount()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		AllUpdatesRule: ProtoToBillingbudgetsAlphaBudgetAllUpdatesRule(p.GetAllUpdatesRule()),
		BillingAccount: dcl.StringOrNil(p.GetBillingAccount()),
	}
	for _, r := range p.GetThresholdRules() {
		obj.ThresholdRules = append(obj.ThresholdRules, *ProtoToBillingbudgetsAlphaBudgetThresholdRules(r))
	}
	return obj
}

// BudgetBudgetFilterCreditTypesTreatmentEnumToProto converts a BudgetBudgetFilterCreditTypesTreatmentEnum enum to its proto representation.
func BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnumToProto(e *alpha.BudgetBudgetFilterCreditTypesTreatmentEnum) alphapb.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum {
	if e == nil {
		return alphapb.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum(0)
	}
	if v, ok := alphapb.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum_value["BudgetBudgetFilterCreditTypesTreatmentEnum"+string(*e)]; ok {
		return alphapb.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum(v)
	}
	return alphapb.BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnum(0)
}

// BudgetBudgetFilterCalendarPeriodEnumToProto converts a BudgetBudgetFilterCalendarPeriodEnum enum to its proto representation.
func BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnumToProto(e *alpha.BudgetBudgetFilterCalendarPeriodEnum) alphapb.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum {
	if e == nil {
		return alphapb.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum(0)
	}
	if v, ok := alphapb.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum_value["BudgetBudgetFilterCalendarPeriodEnum"+string(*e)]; ok {
		return alphapb.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum(v)
	}
	return alphapb.BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnum(0)
}

// BudgetThresholdRulesSpendBasisEnumToProto converts a BudgetThresholdRulesSpendBasisEnum enum to its proto representation.
func BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnumToProto(e *alpha.BudgetThresholdRulesSpendBasisEnum) alphapb.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum {
	if e == nil {
		return alphapb.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum(0)
	}
	if v, ok := alphapb.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum_value["BudgetThresholdRulesSpendBasisEnum"+string(*e)]; ok {
		return alphapb.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum(v)
	}
	return alphapb.BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnum(0)
}

// BudgetBudgetFilterToProto converts a BudgetBudgetFilter object to its proto representation.
func BillingbudgetsAlphaBudgetBudgetFilterToProto(o *alpha.BudgetBudgetFilter) *alphapb.BillingbudgetsAlphaBudgetBudgetFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetBudgetFilter{}
	p.SetCreditTypesTreatment(BillingbudgetsAlphaBudgetBudgetFilterCreditTypesTreatmentEnumToProto(o.CreditTypesTreatment))
	p.SetCalendarPeriod(BillingbudgetsAlphaBudgetBudgetFilterCalendarPeriodEnumToProto(o.CalendarPeriod))
	p.SetCustomPeriod(BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodToProto(o.CustomPeriod))
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
	mLabels := make(map[string]*alphapb.BillingbudgetsAlphaBudgetBudgetFilterLabels, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = BillingbudgetsAlphaBudgetBudgetFilterLabelsToProto(&r)
	}
	p.SetLabels(mLabels)
	return p
}

// BudgetBudgetFilterLabelsToProto converts a BudgetBudgetFilterLabels object to its proto representation.
func BillingbudgetsAlphaBudgetBudgetFilterLabelsToProto(o *alpha.BudgetBudgetFilterLabels) *alphapb.BillingbudgetsAlphaBudgetBudgetFilterLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetBudgetFilterLabels{}
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// BudgetBudgetFilterCustomPeriodToProto converts a BudgetBudgetFilterCustomPeriod object to its proto representation.
func BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodToProto(o *alpha.BudgetBudgetFilterCustomPeriod) *alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriod {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriod{}
	p.SetStartDate(BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodStartDateToProto(o.StartDate))
	p.SetEndDate(BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodEndDateToProto(o.EndDate))
	return p
}

// BudgetBudgetFilterCustomPeriodStartDateToProto converts a BudgetBudgetFilterCustomPeriodStartDate object to its proto representation.
func BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodStartDateToProto(o *alpha.BudgetBudgetFilterCustomPeriodStartDate) *alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodStartDate {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodStartDate{}
	p.SetYear(dcl.ValueOrEmptyInt64(o.Year))
	p.SetMonth(dcl.ValueOrEmptyInt64(o.Month))
	p.SetDay(dcl.ValueOrEmptyInt64(o.Day))
	return p
}

// BudgetBudgetFilterCustomPeriodEndDateToProto converts a BudgetBudgetFilterCustomPeriodEndDate object to its proto representation.
func BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodEndDateToProto(o *alpha.BudgetBudgetFilterCustomPeriodEndDate) *alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodEndDate {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetBudgetFilterCustomPeriodEndDate{}
	p.SetYear(dcl.ValueOrEmptyInt64(o.Year))
	p.SetMonth(dcl.ValueOrEmptyInt64(o.Month))
	p.SetDay(dcl.ValueOrEmptyInt64(o.Day))
	return p
}

// BudgetAmountToProto converts a BudgetAmount object to its proto representation.
func BillingbudgetsAlphaBudgetAmountToProto(o *alpha.BudgetAmount) *alphapb.BillingbudgetsAlphaBudgetAmount {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetAmount{}
	p.SetSpecifiedAmount(BillingbudgetsAlphaBudgetAmountSpecifiedAmountToProto(o.SpecifiedAmount))
	p.SetLastPeriodAmount(BillingbudgetsAlphaBudgetAmountLastPeriodAmountToProto(o.LastPeriodAmount))
	return p
}

// BudgetAmountSpecifiedAmountToProto converts a BudgetAmountSpecifiedAmount object to its proto representation.
func BillingbudgetsAlphaBudgetAmountSpecifiedAmountToProto(o *alpha.BudgetAmountSpecifiedAmount) *alphapb.BillingbudgetsAlphaBudgetAmountSpecifiedAmount {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetAmountSpecifiedAmount{}
	p.SetCurrencyCode(dcl.ValueOrEmptyString(o.CurrencyCode))
	p.SetUnits(dcl.ValueOrEmptyInt64(o.Units))
	p.SetNanos(dcl.ValueOrEmptyInt64(o.Nanos))
	return p
}

// BudgetAmountLastPeriodAmountToProto converts a BudgetAmountLastPeriodAmount object to its proto representation.
func BillingbudgetsAlphaBudgetAmountLastPeriodAmountToProto(o *alpha.BudgetAmountLastPeriodAmount) *alphapb.BillingbudgetsAlphaBudgetAmountLastPeriodAmount {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetAmountLastPeriodAmount{}
	return p
}

// BudgetThresholdRulesToProto converts a BudgetThresholdRules object to its proto representation.
func BillingbudgetsAlphaBudgetThresholdRulesToProto(o *alpha.BudgetThresholdRules) *alphapb.BillingbudgetsAlphaBudgetThresholdRules {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetThresholdRules{}
	p.SetThresholdPercent(dcl.ValueOrEmptyDouble(o.ThresholdPercent))
	p.SetSpendBasis(BillingbudgetsAlphaBudgetThresholdRulesSpendBasisEnumToProto(o.SpendBasis))
	return p
}

// BudgetAllUpdatesRuleToProto converts a BudgetAllUpdatesRule object to its proto representation.
func BillingbudgetsAlphaBudgetAllUpdatesRuleToProto(o *alpha.BudgetAllUpdatesRule) *alphapb.BillingbudgetsAlphaBudgetAllUpdatesRule {
	if o == nil {
		return nil
	}
	p := &alphapb.BillingbudgetsAlphaBudgetAllUpdatesRule{}
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
func BudgetToProto(resource *alpha.Budget) *alphapb.BillingbudgetsAlphaBudget {
	p := &alphapb.BillingbudgetsAlphaBudget{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetBudgetFilter(BillingbudgetsAlphaBudgetBudgetFilterToProto(resource.BudgetFilter))
	p.SetAmount(BillingbudgetsAlphaBudgetAmountToProto(resource.Amount))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetAllUpdatesRule(BillingbudgetsAlphaBudgetAllUpdatesRuleToProto(resource.AllUpdatesRule))
	p.SetBillingAccount(dcl.ValueOrEmptyString(resource.BillingAccount))
	sThresholdRules := make([]*alphapb.BillingbudgetsAlphaBudgetThresholdRules, len(resource.ThresholdRules))
	for i, r := range resource.ThresholdRules {
		sThresholdRules[i] = BillingbudgetsAlphaBudgetThresholdRulesToProto(&r)
	}
	p.SetThresholdRules(sThresholdRules)

	return p
}

// applyBudget handles the gRPC request by passing it to the underlying Budget Apply() method.
func (s *BudgetServer) applyBudget(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBillingbudgetsAlphaBudgetRequest) (*alphapb.BillingbudgetsAlphaBudget, error) {
	p := ProtoToBudget(request.GetResource())
	res, err := c.ApplyBudget(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BudgetToProto(res)
	return r, nil
}

// applyBillingbudgetsAlphaBudget handles the gRPC request by passing it to the underlying Budget Apply() method.
func (s *BudgetServer) ApplyBillingbudgetsAlphaBudget(ctx context.Context, request *alphapb.ApplyBillingbudgetsAlphaBudgetRequest) (*alphapb.BillingbudgetsAlphaBudget, error) {
	cl, err := createConfigBudget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyBudget(ctx, cl, request)
}

// DeleteBudget handles the gRPC request by passing it to the underlying Budget Delete() method.
func (s *BudgetServer) DeleteBillingbudgetsAlphaBudget(ctx context.Context, request *alphapb.DeleteBillingbudgetsAlphaBudgetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBudget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBudget(ctx, ProtoToBudget(request.GetResource()))

}

// ListBillingbudgetsAlphaBudget handles the gRPC request by passing it to the underlying BudgetList() method.
func (s *BudgetServer) ListBillingbudgetsAlphaBudget(ctx context.Context, request *alphapb.ListBillingbudgetsAlphaBudgetRequest) (*alphapb.ListBillingbudgetsAlphaBudgetResponse, error) {
	cl, err := createConfigBudget(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBudget(ctx, request.GetBillingAccount())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.BillingbudgetsAlphaBudget
	for _, r := range resources.Items {
		rp := BudgetToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListBillingbudgetsAlphaBudgetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigBudget(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
