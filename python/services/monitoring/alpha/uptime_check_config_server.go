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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/monitoring/alpha/monitoring_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/monitoring/alpha"
)

// UptimeCheckConfigServer implements the gRPC interface for UptimeCheckConfig.
type UptimeCheckConfigServer struct{}

// ProtoToUptimeCheckConfigResourceGroupResourceTypeEnum converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(e alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum) *alpha.UptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum_name[int32(e)]; ok {
		e := alpha.UptimeCheckConfigResourceGroupResourceTypeEnum(n[len("MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckRequestMethodEnum converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(e alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum) *alpha.UptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum_name[int32(e)]; ok {
		e := alpha.UptimeCheckConfigHttpCheckRequestMethodEnum(n[len("MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigHttpCheckContentTypeEnum converts a UptimeCheckConfigHttpCheckContentTypeEnum enum from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(e alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum) *alpha.UptimeCheckConfigHttpCheckContentTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum_name[int32(e)]; ok {
		e := alpha.UptimeCheckConfigHttpCheckContentTypeEnum(n[len("MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigContentMatchersMatcherEnum converts a UptimeCheckConfigContentMatchersMatcherEnum enum from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(e alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum) *alpha.UptimeCheckConfigContentMatchersMatcherEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum_name[int32(e)]; ok {
		e := alpha.UptimeCheckConfigContentMatchersMatcherEnum(n[len("MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum"):])
		return &e
	}
	return nil
}

// ProtoToUptimeCheckConfigMonitoredResource converts a UptimeCheckConfigMonitoredResource object from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigMonitoredResource(p *alphapb.MonitoringAlphaUptimeCheckConfigMonitoredResource) *alpha.UptimeCheckConfigMonitoredResource {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigMonitoredResource{
		Type: dcl.StringOrNil(p.GetType()),
	}
	return obj
}

// ProtoToUptimeCheckConfigResourceGroup converts a UptimeCheckConfigResourceGroup object from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigResourceGroup(p *alphapb.MonitoringAlphaUptimeCheckConfigResourceGroup) *alpha.UptimeCheckConfigResourceGroup {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigResourceGroup{
		GroupId:      dcl.StringOrNil(p.GetGroupId()),
		ResourceType: ProtoToMonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheck converts a UptimeCheckConfigHttpCheck object from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigHttpCheck(p *alphapb.MonitoringAlphaUptimeCheckConfigHttpCheck) *alpha.UptimeCheckConfigHttpCheck {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigHttpCheck{
		RequestMethod: ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(p.GetRequestMethod()),
		UseSsl:        dcl.Bool(p.GetUseSsl()),
		Path:          dcl.StringOrNil(p.GetPath()),
		Port:          dcl.Int64OrNil(p.GetPort()),
		AuthInfo:      ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo(p.GetAuthInfo()),
		MaskHeaders:   dcl.Bool(p.GetMaskHeaders()),
		ContentType:   ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(p.GetContentType()),
		ValidateSsl:   dcl.Bool(p.GetValidateSsl()),
		Body:          dcl.StringOrNil(p.GetBody()),
	}
	return obj
}

// ProtoToUptimeCheckConfigHttpCheckAuthInfo converts a UptimeCheckConfigHttpCheckAuthInfo object from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo(p *alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo) *alpha.UptimeCheckConfigHttpCheckAuthInfo {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigHttpCheckAuthInfo{
		Username: dcl.StringOrNil(p.GetUsername()),
		Password: dcl.StringOrNil(p.GetPassword()),
	}
	return obj
}

// ProtoToUptimeCheckConfigTcpCheck converts a UptimeCheckConfigTcpCheck object from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigTcpCheck(p *alphapb.MonitoringAlphaUptimeCheckConfigTcpCheck) *alpha.UptimeCheckConfigTcpCheck {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigTcpCheck{
		Port: dcl.Int64OrNil(p.GetPort()),
	}
	return obj
}

// ProtoToUptimeCheckConfigContentMatchers converts a UptimeCheckConfigContentMatchers object from its proto representation.
func ProtoToMonitoringAlphaUptimeCheckConfigContentMatchers(p *alphapb.MonitoringAlphaUptimeCheckConfigContentMatchers) *alpha.UptimeCheckConfigContentMatchers {
	if p == nil {
		return nil
	}
	obj := &alpha.UptimeCheckConfigContentMatchers{
		Content: dcl.StringOrNil(p.GetContent()),
		Matcher: ProtoToMonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(p.GetMatcher()),
	}
	return obj
}

// ProtoToUptimeCheckConfig converts a UptimeCheckConfig resource from its proto representation.
func ProtoToUptimeCheckConfig(p *alphapb.MonitoringAlphaUptimeCheckConfig) *alpha.UptimeCheckConfig {
	obj := &alpha.UptimeCheckConfig{
		Name:              dcl.StringOrNil(p.GetName()),
		DisplayName:       dcl.StringOrNil(p.GetDisplayName()),
		MonitoredResource: ProtoToMonitoringAlphaUptimeCheckConfigMonitoredResource(p.GetMonitoredResource()),
		ResourceGroup:     ProtoToMonitoringAlphaUptimeCheckConfigResourceGroup(p.GetResourceGroup()),
		HttpCheck:         ProtoToMonitoringAlphaUptimeCheckConfigHttpCheck(p.GetHttpCheck()),
		TcpCheck:          ProtoToMonitoringAlphaUptimeCheckConfigTcpCheck(p.GetTcpCheck()),
		Period:            dcl.StringOrNil(p.GetPeriod()),
		Timeout:           dcl.StringOrNil(p.GetTimeout()),
		Project:           dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetContentMatchers() {
		obj.ContentMatchers = append(obj.ContentMatchers, *ProtoToMonitoringAlphaUptimeCheckConfigContentMatchers(r))
	}
	for _, r := range p.GetSelectedRegions() {
		obj.SelectedRegions = append(obj.SelectedRegions, r)
	}
	return obj
}

// UptimeCheckConfigResourceGroupResourceTypeEnumToProto converts a UptimeCheckConfigResourceGroupResourceTypeEnum enum to its proto representation.
func MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnumToProto(e *alpha.UptimeCheckConfigResourceGroupResourceTypeEnum) alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum_value["UptimeCheckConfigResourceGroupResourceTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(v)
	}
	return alphapb.MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnum(0)
}

// UptimeCheckConfigHttpCheckRequestMethodEnumToProto converts a UptimeCheckConfigHttpCheckRequestMethodEnum enum to its proto representation.
func MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnumToProto(e *alpha.UptimeCheckConfigHttpCheckRequestMethodEnum) alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum {
	if e == nil {
		return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum_value["UptimeCheckConfigHttpCheckRequestMethodEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(v)
	}
	return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnum(0)
}

// UptimeCheckConfigHttpCheckContentTypeEnumToProto converts a UptimeCheckConfigHttpCheckContentTypeEnum enum to its proto representation.
func MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnumToProto(e *alpha.UptimeCheckConfigHttpCheckContentTypeEnum) alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum {
	if e == nil {
		return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum_value["UptimeCheckConfigHttpCheckContentTypeEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(v)
	}
	return alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnum(0)
}

// UptimeCheckConfigContentMatchersMatcherEnumToProto converts a UptimeCheckConfigContentMatchersMatcherEnum enum to its proto representation.
func MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnumToProto(e *alpha.UptimeCheckConfigContentMatchersMatcherEnum) alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum {
	if e == nil {
		return alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(0)
	}
	if v, ok := alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum_value["UptimeCheckConfigContentMatchersMatcherEnum"+string(*e)]; ok {
		return alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(v)
	}
	return alphapb.MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnum(0)
}

// UptimeCheckConfigMonitoredResourceToProto converts a UptimeCheckConfigMonitoredResource object to its proto representation.
func MonitoringAlphaUptimeCheckConfigMonitoredResourceToProto(o *alpha.UptimeCheckConfigMonitoredResource) *alphapb.MonitoringAlphaUptimeCheckConfigMonitoredResource {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigMonitoredResource{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	mFilterLabels := make(map[string]string, len(o.FilterLabels))
	for k, r := range o.FilterLabels {
		mFilterLabels[k] = r
	}
	p.SetFilterLabels(mFilterLabels)
	return p
}

// UptimeCheckConfigResourceGroupToProto converts a UptimeCheckConfigResourceGroup object to its proto representation.
func MonitoringAlphaUptimeCheckConfigResourceGroupToProto(o *alpha.UptimeCheckConfigResourceGroup) *alphapb.MonitoringAlphaUptimeCheckConfigResourceGroup {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigResourceGroup{}
	p.SetGroupId(dcl.ValueOrEmptyString(o.GroupId))
	p.SetResourceType(MonitoringAlphaUptimeCheckConfigResourceGroupResourceTypeEnumToProto(o.ResourceType))
	return p
}

// UptimeCheckConfigHttpCheckToProto converts a UptimeCheckConfigHttpCheck object to its proto representation.
func MonitoringAlphaUptimeCheckConfigHttpCheckToProto(o *alpha.UptimeCheckConfigHttpCheck) *alphapb.MonitoringAlphaUptimeCheckConfigHttpCheck {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigHttpCheck{}
	p.SetRequestMethod(MonitoringAlphaUptimeCheckConfigHttpCheckRequestMethodEnumToProto(o.RequestMethod))
	p.SetUseSsl(dcl.ValueOrEmptyBool(o.UseSsl))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	p.SetAuthInfo(MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfoToProto(o.AuthInfo))
	p.SetMaskHeaders(dcl.ValueOrEmptyBool(o.MaskHeaders))
	p.SetContentType(MonitoringAlphaUptimeCheckConfigHttpCheckContentTypeEnumToProto(o.ContentType))
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
func MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfoToProto(o *alpha.UptimeCheckConfigHttpCheckAuthInfo) *alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigHttpCheckAuthInfo{}
	p.SetUsername(dcl.ValueOrEmptyString(o.Username))
	p.SetPassword(dcl.ValueOrEmptyString(o.Password))
	return p
}

// UptimeCheckConfigTcpCheckToProto converts a UptimeCheckConfigTcpCheck object to its proto representation.
func MonitoringAlphaUptimeCheckConfigTcpCheckToProto(o *alpha.UptimeCheckConfigTcpCheck) *alphapb.MonitoringAlphaUptimeCheckConfigTcpCheck {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigTcpCheck{}
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	return p
}

// UptimeCheckConfigContentMatchersToProto converts a UptimeCheckConfigContentMatchers object to its proto representation.
func MonitoringAlphaUptimeCheckConfigContentMatchersToProto(o *alpha.UptimeCheckConfigContentMatchers) *alphapb.MonitoringAlphaUptimeCheckConfigContentMatchers {
	if o == nil {
		return nil
	}
	p := &alphapb.MonitoringAlphaUptimeCheckConfigContentMatchers{}
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetMatcher(MonitoringAlphaUptimeCheckConfigContentMatchersMatcherEnumToProto(o.Matcher))
	return p
}

// UptimeCheckConfigToProto converts a UptimeCheckConfig resource to its proto representation.
func UptimeCheckConfigToProto(resource *alpha.UptimeCheckConfig) *alphapb.MonitoringAlphaUptimeCheckConfig {
	p := &alphapb.MonitoringAlphaUptimeCheckConfig{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetMonitoredResource(MonitoringAlphaUptimeCheckConfigMonitoredResourceToProto(resource.MonitoredResource))
	p.SetResourceGroup(MonitoringAlphaUptimeCheckConfigResourceGroupToProto(resource.ResourceGroup))
	p.SetHttpCheck(MonitoringAlphaUptimeCheckConfigHttpCheckToProto(resource.HttpCheck))
	p.SetTcpCheck(MonitoringAlphaUptimeCheckConfigTcpCheckToProto(resource.TcpCheck))
	p.SetPeriod(dcl.ValueOrEmptyString(resource.Period))
	p.SetTimeout(dcl.ValueOrEmptyString(resource.Timeout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sContentMatchers := make([]*alphapb.MonitoringAlphaUptimeCheckConfigContentMatchers, len(resource.ContentMatchers))
	for i, r := range resource.ContentMatchers {
		sContentMatchers[i] = MonitoringAlphaUptimeCheckConfigContentMatchersToProto(&r)
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
func (s *UptimeCheckConfigServer) applyUptimeCheckConfig(ctx context.Context, c *alpha.Client, request *alphapb.ApplyMonitoringAlphaUptimeCheckConfigRequest) (*alphapb.MonitoringAlphaUptimeCheckConfig, error) {
	p := ProtoToUptimeCheckConfig(request.GetResource())
	res, err := c.ApplyUptimeCheckConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UptimeCheckConfigToProto(res)
	return r, nil
}

// applyMonitoringAlphaUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Apply() method.
func (s *UptimeCheckConfigServer) ApplyMonitoringAlphaUptimeCheckConfig(ctx context.Context, request *alphapb.ApplyMonitoringAlphaUptimeCheckConfigRequest) (*alphapb.MonitoringAlphaUptimeCheckConfig, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyUptimeCheckConfig(ctx, cl, request)
}

// DeleteUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfig Delete() method.
func (s *UptimeCheckConfigServer) DeleteMonitoringAlphaUptimeCheckConfig(ctx context.Context, request *alphapb.DeleteMonitoringAlphaUptimeCheckConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigUptimeCheckConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUptimeCheckConfig(ctx, ProtoToUptimeCheckConfig(request.GetResource()))

}

// ListMonitoringAlphaUptimeCheckConfig handles the gRPC request by passing it to the underlying UptimeCheckConfigList() method.
func (s *UptimeCheckConfigServer) ListMonitoringAlphaUptimeCheckConfig(ctx context.Context, request *alphapb.ListMonitoringAlphaUptimeCheckConfigRequest) (*alphapb.ListMonitoringAlphaUptimeCheckConfigResponse, error) {
	cl, err := createConfigUptimeCheckConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUptimeCheckConfig(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.MonitoringAlphaUptimeCheckConfig
	for _, r := range resources.Items {
		rp := UptimeCheckConfigToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListMonitoringAlphaUptimeCheckConfigResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigUptimeCheckConfig(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
