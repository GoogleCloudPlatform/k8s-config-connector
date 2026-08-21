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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	monitoringpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/monitoring_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring"
)

// UptimeCheckConfigServer implements the gRPC interface for UptimeCheckConfig.
type UptimeCheckConfigServer struct{}

// ProtoToUptimeCheckConfigResourceGroupResourceTypeEnum converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum from its proto representation.
func ProtoToMonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(e monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum) *monitoring.UptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum_name[int32(e)]; ok {
		e := monitoring.UptimeCheckConfigResourceGroupResourceTypeEnum(n[len("MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckRequestMethodEnum converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum from its proto representation.
func ProtoToMonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(e monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum) *monitoring.UptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum_name[int32(e)]; ok {
		e := monitoring.UptimeCheckConfigHttpCheckRequestMethodEnum(n[len("MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckContentTypeEnum converts a UptimeCheckConfigHttpCheckContentTypeEnum enum from its proto representation.
func ProtoToMonitoringUptimeCheckConfigHttpCheckContentTypeEnum(e monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum) *monitoring.UptimeCheckConfigHttpCheckContentTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum_name[int32(e)]; ok {
		e := monitoring.UptimeCheckConfigHttpCheckContentTypeEnum(n[len("MonitoringUptimeCheckConfigHttpCheckContentTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigContentMatchersMatcherEnum converts a UptimeCheckConfigContentMatchersMatcherEnum enum from its proto representation.
func ProtoToMonitoringUptimeCheckConfigContentMatchersMatcherEnum(e monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum) *monitoring.UptimeCheckConfigContentMatchersMatcherEnum {
	if e == 0 {
		return nil
	}
	if n, ok := monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum_name[int32(e)]; ok {
		e := monitoring.UptimeCheckConfigContentMatchersMatcherEnum(n[len("MonitoringUptimeCheckConfigContentMatchersMatcherEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigMonitoredResource converts a UptimeCheckConfigMonitoredResource object from its proto representation.
func ProtoToMonitoringUptimeCheckConfigMonitoredResource(p *monitoringpb.MonitoringUptimeCheckConfigMonitoredResource) *monitoring.UptimeCheckConfigMonitoredResource {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigMonitoredResource{
		Type: dcl.StringOrNil(p.GetType()),
	}
	return obj
}

// ProtoToUptimeCheckConfigResourceGroup converts a UptimeCheckConfigResourceGroup object from its proto representation.
func ProtoToMonitoringUptimeCheckConfigResourceGroup(p *monitoringpb.MonitoringUptimeCheckConfigResourceGroup) *monitoring.UptimeCheckConfigResourceGroup {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigResourceGroup{
		GroupId:      dcl.StringOrNil(p.GetGroupId()),
		ResourceType: ProtoToMonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheck converts a UptimeCheckConfigHttpCheck object from its proto representation.
func ProtoToMonitoringUptimeCheckConfigHttpCheck(p *monitoringpb.MonitoringUptimeCheckConfigHttpCheck) *monitoring.UptimeCheckConfigHttpCheck {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigHttpCheck{
		RequestMethod: ProtoToMonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(p.GetRequestMethod()),
		UseSsl:        dcl.Bool(p.GetUseSsl()),
		Path:          dcl.StringOrNil(p.GetPath()),
		Port:          dcl.Int64OrNil(p.GetPort()),
		AuthInfo:      ProtoToMonitoringUptimeCheckConfigHttpCheckAuthInfo(p.GetAuthInfo()),
		MaskHeaders:   dcl.Bool(p.GetMaskHeaders()),
		ContentType:   ProtoToMonitoringUptimeCheckConfigHttpCheckContentTypeEnum(p.GetContentType()),
		ValidateSsl:   dcl.Bool(p.GetValidateSsl()),
		Body:          dcl.StringOrNil(p.GetBody()),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheckAuthInfo converts a UptimeCheckConfigHttpCheckAuthInfo object from its proto representation.
func ProtoToMonitoringUptimeCheckConfigHttpCheckAuthInfo(p *monitoringpb.MonitoringUptimeCheckConfigHttpCheckAuthInfo) *monitoring.UptimeCheckConfigHttpCheckAuthInfo {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigHttpCheckAuthInfo{
		Username: dcl.StringOrNil(p.GetUsername()),
		Password: dcl.StringOrNil(p.GetPassword()),
	}
	return obj
}

// ProtoToUptimeCheckConfigTcpCheck converts a UptimeCheckConfigTcpCheck object from its proto representation.
func ProtoToMonitoringUptimeCheckConfigTcpCheck(p *monitoringpb.MonitoringUptimeCheckConfigTcpCheck) *monitoring.UptimeCheckConfigTcpCheck {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigTcpCheck{
		Port: dcl.Int64OrNil(p.GetPort()),
	}
	return obj
}

// ProtoToUptimeCheckConfigContentMatchers converts a UptimeCheckConfigContentMatchers object from its proto representation.
func ProtoToMonitoringUptimeCheckConfigContentMatchers(p *monitoringpb.MonitoringUptimeCheckConfigContentMatchers) *monitoring.UptimeCheckConfigContentMatchers {
	if p == nil {
		return nil
	}
	obj := &monitoring.UptimeCheckConfigContentMatchers{
		Content: dcl.StringOrNil(p.GetContent()),
		Matcher: ProtoToMonitoringUptimeCheckConfigContentMatchersMatcherEnum(p.GetMatcher()),
	}
	return obj
}

// ProtoToUptimeCheckConfig converts a UptimeCheckConfig resource from its proto representation.
func ProtoToUptimeCheckConfig(p *monitoringpb.MonitoringUptimeCheckConfig) *monitoring.UptimeCheckConfig {
	obj := &monitoring.UptimeCheckConfig{
		Name:              dcl.StringOrNil(p.GetName()),
		DisplayName:       dcl.StringOrNil(p.GetDisplayName()),
		MonitoredResource: ProtoToMonitoringUptimeCheckConfigMonitoredResource(p.GetMonitoredResource()),
		ResourceGroup:     ProtoToMonitoringUptimeCheckConfigResourceGroup(p.GetResourceGroup()),
		HttpCheck:         ProtoToMonitoringUptimeCheckConfigHttpCheck(p.GetHttpCheck()),
		TcpCheck:          ProtoToMonitoringUptimeCheckConfigTcpCheck(p.GetTcpCheck()),
		Period:            dcl.StringOrNil(p.GetPeriod()),
		Timeout:           dcl.StringOrNil(p.GetTimeout()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetContentMatchers() {
		obj.ContentMatchers = append(obj.ContentMatchers, *ProtoToMonitoringUptimeCheckConfigContentMatchers(r))
	}
	for _, r := range p.GetSelectedRegions() {
		obj.SelectedRegions = append(obj.SelectedRegions, r)
	}
	return obj
}

// UptimeCheckConfigResourceGroupResourceTypeEnumToProto converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum to its proto representation.
func MonitoringUptimeCheckConfigResourceGroupResourceTypeEnumToProto(e *monitoring.UptimeCheckConfigResourceGroupResourceTypeEnum) monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum_value["UptimeCheckConfigResourceGroupResourceTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(v)
	}
	return monitoringpb.MonitoringUptimeCheckConfigResourceGroupResourceTypeEnum(0)
}

// UptimeCheckConfigHttpCheckRequestMethodEnumToProto converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum to its proto representation.
func MonitoringUptimeCheckConfigHttpCheckRequestMethodEnumToProto(e *monitoring.UptimeCheckConfigHttpCheckRequestMethodEnum) monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == nil {
		return monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(0)
	}
	if v, ok := monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum_value["UptimeCheckConfigHttpCheckRequestMethodEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(v)
	}
	return monitoringpb.MonitoringUptimeCheckConfigHttpCheckRequestMethodEnum(0)
}

// UptimeCheckConfigHttpCheckContentTypeEnumToProto converts a UptimeCheckConfigHttpCheckContentTypeEnum enum to its proto representation.
func MonitoringUptimeCheckConfigHttpCheckContentTypeEnumToProto(e *monitoring.UptimeCheckConfigHttpCheckContentTypeEnum) monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum {
	if e == nil {
		return monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum(0)
	}
	if v, ok := monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum_value["UptimeCheckConfigHttpCheckContentTypeEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum(v)
	}
	return monitoringpb.MonitoringUptimeCheckConfigHttpCheckContentTypeEnum(0)
}

// UptimeCheckConfigContentMatchersMatcherEnumToProto converts a UptimeCheckConfigContentMatchersMatcherEnum enum to its proto representation.
func MonitoringUptimeCheckConfigContentMatchersMatcherEnumToProto(e *monitoring.UptimeCheckConfigContentMatchersMatcherEnum) monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum {
	if e == nil {
		return monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum(0)
	}
	if v, ok := monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum_value["UptimeCheckConfigContentMatchersMatcherEnum"+string(*e)]; ok {
		return monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum(v)
	}
	return monitoringpb.MonitoringUptimeCheckConfigContentMatchersMatcherEnum(0)
}

// UptimeCheckConfigMonitoredResourceToProto converts a UptimeCheckConfigMonitoredResource object to its proto representation.
func MonitoringUptimeCheckConfigMonitoredResourceToProto(o *monitoring.UptimeCheckConfigMonitoredResource) *monitoringpb.MonitoringUptimeCheckConfigMonitoredResource {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigMonitoredResource{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	mFilterLabels := make(map[string]string, len(o.FilterLabels))
	for k, r := range o.FilterLabels {
		mFilterLabels[k] = r
	}
	p.SetFilterLabels(mFilterLabels)
	return p
}

// UptimeCheckConfigResourceGroupToProto converts a UptimeCheckConfigResourceGroup object to its proto representation.
func MonitoringUptimeCheckConfigResourceGroupToProto(o *monitoring.UptimeCheckConfigResourceGroup) *monitoringpb.MonitoringUptimeCheckConfigResourceGroup {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigResourceGroup{}
	p.SetGroupId(dcl.ValueOrEmptyString(o.GroupId))
	p.SetResourceType(MonitoringUptimeCheckConfigResourceGroupResourceTypeEnumToProto(o.ResourceType))
	return p
}

// UptimeCheckConfigHttpCheckToProto converts a UptimeCheckConfigHttpCheck object to its proto representation.
func MonitoringUptimeCheckConfigHttpCheckToProto(o *monitoring.UptimeCheckConfigHttpCheck) *monitoringpb.MonitoringUptimeCheckConfigHttpCheck {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigHttpCheck{}
	p.SetRequestMethod(MonitoringUptimeCheckConfigHttpCheckRequestMethodEnumToProto(o.RequestMethod))
	p.SetUseSsl(dcl.ValueOrEmptyBool(o.UseSsl))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	p.SetAuthInfo(MonitoringUptimeCheckConfigHttpCheckAuthInfoToProto(o.AuthInfo))
	p.SetMaskHeaders(dcl.ValueOrEmptyBool(o.MaskHeaders))
	p.SetContentType(MonitoringUptimeCheckConfigHttpCheckContentTypeEnumToProto(o.ContentType))
	p.SetValidateSsl(dcl.ValueOrEmptyBool(o.ValidateSsl))
	p.SetBody(dcl.ValueOrEmptyString(o.Body))
	mHeaders := make(map[string]string, len(o.Headers))
	for k, r := range o.Headers {
		mHeaders[k] = r
	}
	p.SetHeaders(mHeaders)
	return p
}

// UptimeCheckConfigHttpCheckAuthInfoToProto converts a UptimeCheckConfigHttpCheckAuthInfo object to its proto representation.
func MonitoringUptimeCheckConfigHttpCheckAuthInfoToProto(o *monitoring.UptimeCheckConfigHttpCheckAuthInfo) *monitoringpb.MonitoringUptimeCheckConfigHttpCheckAuthInfo {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigHttpCheckAuthInfo{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	p.SetPassword(dcl.ValueOrEmptyString(o.Password))
	return p
}

// UptimeCheckConfigTcpCheckToProto converts a UptimeCheckConfigTcpCheck object to its proto representation.
func MonitoringUptimeCheckConfigTcpCheckToProto(o *monitoring.UptimeCheckConfigTcpCheck) *monitoringpb.MonitoringUptimeCheckConfigTcpCheck {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigTcpCheck{}
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	return p
}

// UptimeCheckConfigContentMatchersToProto converts a UptimeCheckConfigContentMatchers object to its proto representation.
func MonitoringUptimeCheckConfigContentMatchersToProto(o *monitoring.UptimeCheckConfigContentMatchers) *monitoringpb.MonitoringUptimeCheckConfigContentMatchers {
	if o == nil {
		return nil
	}
	p := &monitoringpb.MonitoringUptimeCheckConfigContentMatchers{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetMatcher(MonitoringUptimeCheckConfigContentMatchersMatcherEnumToProto(o.Matcher))
	return p
}

// UptimeCheckConfigToProto converts a UptimeCheckConfig resource to its proto representation.
func UptimeCheckConfigToProto(resource *monitoring.UptimeCheckConfig) *monitoringpb.MonitoringUptimeCheckConfig {
	p := &monitoringpb.MonitoringUptimeCheckConfig{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetMonitoredResource(MonitoringUptimeCheckConfigMonitoredResourceToProto(resource.MonitoredResource))
	p.SetResourceGroup(MonitoringUptimeCheckConfigResourceGroupToProto(resource.ResourceGroup))
	p.SetHttpCheck(MonitoringUptimeCheckConfigHttpCheckToProto(resource.HttpCheck))
	p.SetTcpCheck(MonitoringUptimeCheckConfigTcpCheckToProto(resource.TcpCheck))
	p.SetPeriod(dcl.ValueOrEmptyString(resource.Period))
	p.SetTimeout(dcl.ValueOrEmptyString(resource.Timeout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sContentMatchers := make([]*monitoringpb.MonitoringUptimeCheckConfigContentMatchers, len(resource.ContentMatchers))
	for i, r := range resource.ContentMatchers {
		sContentMatchers[i] = MonitoringUptimeCheckConfigContentMatchersToProto(&r)
	}
	p.SetContentMatchers(sContentMatchers)
	sSelectedRegions := make([]string, len(resource.SelectedRegions))
	for i, r := range resource.SelectedRegions {
		sSelectedRegions[i] = r
	}
	p.SetSelectedRegions(sSelectedRegions)

	return p
}

// applyUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) applyUptimeCheckConfig(ctx context.Context, c *monitoring.Client, request *monitoringpb.ApplyMonitoringUptimeCheckConfigRequest) (*monitoringpb.MonitoringUptimeCheckConfig, error) {
	p := ProtoToUptimeCheckConfig(request.GetResource())
	res, err := c.ApplyUptimeCheckConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UptimeCheckConfigToProto(res)
	return r, nil
}

// applyMonitoringUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) ApplyMonitoringUptimeCheckConfig(ctx context.Context, request *monitoringpb.ApplyMonitoringUptimeCheckConfigRequest) (*monitoringpb.MonitoringUptimeCheckConfig, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyUptimeCheckConfig(ctx, cl, request)
}

// DeleteUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Delete() method.
func (s *UptimeCheckConfigServer) DeleteMonitoringUptimeCheckConfig(ctx context.Context, request *monitoringpb.DeleteMonitoringUptimeCheckConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigUptimeCheckConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUptimeCheckConfig(ctx, ProtoToUptimeCheckConfig(request.GetResource()))

}

// ListMonitoringUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfigList() method.
func (s *UptimeCheckConfigServer) ListMonitoringUptimeCheckConfig(ctx context.Context, request *monitoringpb.ListMonitoringUptimeCheckConfigRequest) (*monitoringpb.ListMonitoringUptimeCheckConfigResponse, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUptimeCheckConfig(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*monitoringpb.MonitoringUptimeCheckConfig
	for _, r := range resources.Items {
		rp := UptimeCheckConfigToProto(r)
		protos = append(protos, rp)
	}
	p := &monitoringpb.ListMonitoringUptimeCheckConfigResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigUptimeCheckConfig(ctx context.Context, service_account_file string) (*monitoring.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return monitoring.NewClient(conf), nil
}
