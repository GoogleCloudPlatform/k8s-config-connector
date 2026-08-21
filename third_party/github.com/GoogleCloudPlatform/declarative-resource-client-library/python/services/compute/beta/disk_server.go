// Copyright 2021 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for Disk.
type DiskServer struct{}

// ProtoToDiskGuestOSFeatureTypeEnum converts a DiskGuestOSFeatureTypeEnum enum from its proto representation.
func ProtoToComputeBetaDiskGuestOSFeatureTypeEnum(e betapb.ComputeBetaDiskGuestOSFeatureTypeEnum) *beta.DiskGuestOSFeatureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskGuestOSFeatureTypeEnum_name[int32(e)]; ok {
		e := beta.DiskGuestOSFeatureTypeEnum(n[len("ComputeBetaDiskGuestOSFeatureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOSFeatureTypeAltEnum converts a DiskGuestOSFeatureTypeAltEnum enum from its proto representation.
func ProtoToComputeBetaDiskGuestOSFeatureTypeAltEnum(e betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum) *beta.DiskGuestOSFeatureTypeAltEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum_name[int32(e)]; ok {
		e := beta.DiskGuestOSFeatureTypeAltEnum(n[len("ComputeBetaDiskGuestOSFeatureTypeAltEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskStatusEnum converts a DiskStatusEnum enum from its proto representation.
func ProtoToComputeBetaDiskStatusEnum(e betapb.ComputeBetaDiskStatusEnum) *beta.DiskStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskStatusEnum_name[int32(e)]; ok {
		e := beta.DiskStatusEnum(n[len("ComputeBetaDiskStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOSFeaturesTypeEnum converts a DiskGuestOSFeaturesTypeEnum enum from its proto representation.
func ProtoToComputeBetaDiskGuestOSFeaturesTypeEnum(e betapb.ComputeBetaDiskGuestOSFeaturesTypeEnum) *beta.DiskGuestOSFeaturesTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskGuestOSFeaturesTypeEnum_name[int32(e)]; ok {
		e := beta.DiskGuestOSFeaturesTypeEnum(n[len("ComputeBetaDiskGuestOSFeaturesTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOSFeaturesTypeAltsEnum converts a DiskGuestOSFeaturesTypeAltsEnum enum from its proto representation.
func ProtoToComputeBetaDiskGuestOSFeaturesTypeAltsEnum(e betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum) *beta.DiskGuestOSFeaturesTypeAltsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum_name[int32(e)]; ok {
		e := beta.DiskGuestOSFeaturesTypeAltsEnum(n[len("ComputeBetaDiskGuestOSFeaturesTypeAltsEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOSFeature converts a DiskGuestOSFeature resource from its proto representation.
func ProtoToComputeBetaDiskGuestOSFeature(p *betapb.ComputeBetaDiskGuestOSFeature) *beta.DiskGuestOSFeature {
	if p == nil {
		return nil
	}
	obj := &beta.DiskGuestOSFeature{
		Type: ProtoToComputeBetaDiskGuestOSFeatureTypeEnum(p.GetType()),
	}
	for _, r := range p.GetTypeAlt() {
		obj.TypeAlt = append(obj.TypeAlt, *ProtoToComputeBetaDiskGuestOSFeatureTypeAltEnum(r))
	}
	return obj
}

// ProtoToDiskEncryptionKey converts a DiskEncryptionKey resource from its proto representation.
func ProtoToComputeBetaDiskEncryptionKey(p *betapb.ComputeBetaDiskEncryptionKey) *beta.DiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.DiskEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToDiskGuestOSFeatures converts a DiskGuestOSFeatures resource from its proto representation.
func ProtoToComputeBetaDiskGuestOSFeatures(p *betapb.ComputeBetaDiskGuestOSFeatures) *beta.DiskGuestOSFeatures {
	if p == nil {
		return nil
	}
	obj := &beta.DiskGuestOSFeatures{
		Type: ProtoToComputeBetaDiskGuestOSFeaturesTypeEnum(p.GetType()),
	}
	for _, r := range p.GetTypeAlts() {
		obj.TypeAlts = append(obj.TypeAlts, *ProtoToComputeBetaDiskGuestOSFeaturesTypeAltsEnum(r))
	}
	return obj
}

// ProtoToDisk converts a Disk resource from its proto representation.
func ProtoToDisk(p *betapb.ComputeBetaDisk) *beta.Disk {
	obj := &beta.Disk{
		SelfLink:                    dcl.StringOrNil(p.SelfLink),
		Description:                 dcl.StringOrNil(p.Description),
		DiskEncryptionKey:           ProtoToComputeBetaDiskEncryptionKey(p.GetDiskEncryptionKey()),
		LabelFingerprint:            dcl.StringOrNil(p.LabelFingerprint),
		Name:                        dcl.StringOrNil(p.Name),
		Region:                      dcl.StringOrNil(p.Region),
		SizeGb:                      dcl.Int64OrNil(p.SizeGb),
		SourceImage:                 dcl.StringOrNil(p.SourceImage),
		SourceImageEncryptionKey:    ProtoToComputeBetaDiskEncryptionKey(p.GetSourceImageEncryptionKey()),
		SourceImageId:               dcl.StringOrNil(p.SourceImageId),
		SourceSnapshot:              dcl.StringOrNil(p.SourceSnapshot),
		SourceSnapshotEncryptionKey: ProtoToComputeBetaDiskEncryptionKey(p.GetSourceSnapshotEncryptionKey()),
		SourceSnapshotId:            dcl.StringOrNil(p.SourceSnapshotId),
		Type:                        dcl.StringOrNil(p.Type),
		Zone:                        dcl.StringOrNil(p.Zone),
		Project:                     dcl.StringOrNil(p.Project),
		Id:                          dcl.Int64OrNil(p.Id),
		Status:                      ProtoToComputeBetaDiskStatusEnum(p.GetStatus()),
		Options:                     dcl.StringOrNil(p.Options),
		LastAttachTimestamp:         dcl.StringOrNil(p.LastAttachTimestamp),
		LastDetachTimestamp:         dcl.StringOrNil(p.LastDetachTimestamp),
		PhysicalBlockSizeBytes:      dcl.Int64OrNil(p.PhysicalBlockSizeBytes),
		SourceDisk:                  dcl.StringOrNil(p.SourceDisk),
		SourceDiskId:                dcl.StringOrNil(p.SourceDiskId),
		Location:                    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetGuestOsFeature() {
		obj.GuestOSFeature = append(obj.GuestOSFeature, *ProtoToComputeBetaDiskGuestOSFeature(r))
	}
	for _, r := range p.GetLicense() {
		obj.License = append(obj.License, r)
	}
	for _, r := range p.GetReplicaZones() {
		obj.ReplicaZones = append(obj.ReplicaZones, r)
	}
	for _, r := range p.GetResourcePolicy() {
		obj.ResourcePolicy = append(obj.ResourcePolicy, r)
	}
	for _, r := range p.GetLicenses() {
		obj.Licenses = append(obj.Licenses, r)
	}
	for _, r := range p.GetGuestOsFeatures() {
		obj.GuestOSFeatures = append(obj.GuestOSFeatures, *ProtoToComputeBetaDiskGuestOSFeatures(r))
	}
	for _, r := range p.GetUsers() {
		obj.Users = append(obj.Users, r)
	}
	for _, r := range p.GetLicenseCodes() {
		obj.LicenseCodes = append(obj.LicenseCodes, r)
	}
	for _, r := range p.GetResourcePolicies() {
		obj.ResourcePolicies = append(obj.ResourcePolicies, r)
	}
	return obj
}

// DiskGuestOSFeatureTypeEnumToProto converts a DiskGuestOSFeatureTypeEnum enum to its proto representation.
func ComputeBetaDiskGuestOSFeatureTypeEnumToProto(e *beta.DiskGuestOSFeatureTypeEnum) betapb.ComputeBetaDiskGuestOSFeatureTypeEnum {
	if e == nil {
		return betapb.ComputeBetaDiskGuestOSFeatureTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskGuestOSFeatureTypeEnum_value["DiskGuestOSFeatureTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskGuestOSFeatureTypeEnum(v)
	}
	return betapb.ComputeBetaDiskGuestOSFeatureTypeEnum(0)
}

// DiskGuestOSFeatureTypeAltEnumToProto converts a DiskGuestOSFeatureTypeAltEnum enum to its proto representation.
func ComputeBetaDiskGuestOSFeatureTypeAltEnumToProto(e *beta.DiskGuestOSFeatureTypeAltEnum) betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum {
	if e == nil {
		return betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum_value["DiskGuestOSFeatureTypeAltEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum(v)
	}
	return betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum(0)
}

// DiskStatusEnumToProto converts a DiskStatusEnum enum to its proto representation.
func ComputeBetaDiskStatusEnumToProto(e *beta.DiskStatusEnum) betapb.ComputeBetaDiskStatusEnum {
	if e == nil {
		return betapb.ComputeBetaDiskStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskStatusEnum_value["DiskStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskStatusEnum(v)
	}
	return betapb.ComputeBetaDiskStatusEnum(0)
}

// DiskGuestOSFeaturesTypeEnumToProto converts a DiskGuestOSFeaturesTypeEnum enum to its proto representation.
func ComputeBetaDiskGuestOSFeaturesTypeEnumToProto(e *beta.DiskGuestOSFeaturesTypeEnum) betapb.ComputeBetaDiskGuestOSFeaturesTypeEnum {
	if e == nil {
		return betapb.ComputeBetaDiskGuestOSFeaturesTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskGuestOSFeaturesTypeEnum_value["DiskGuestOSFeaturesTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskGuestOSFeaturesTypeEnum(v)
	}
	return betapb.ComputeBetaDiskGuestOSFeaturesTypeEnum(0)
}

// DiskGuestOSFeaturesTypeAltsEnumToProto converts a DiskGuestOSFeaturesTypeAltsEnum enum to its proto representation.
func ComputeBetaDiskGuestOSFeaturesTypeAltsEnumToProto(e *beta.DiskGuestOSFeaturesTypeAltsEnum) betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum {
	if e == nil {
		return betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum(0)
	}
	if v, ok := betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum_value["DiskGuestOSFeaturesTypeAltsEnum"+string(*e)]; ok {
		return betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum(v)
	}
	return betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum(0)
}

// DiskGuestOSFeatureToProto converts a DiskGuestOSFeature resource to its proto representation.
func ComputeBetaDiskGuestOSFeatureToProto(o *beta.DiskGuestOSFeature) *betapb.ComputeBetaDiskGuestOSFeature {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaDiskGuestOSFeature{
		Type: ComputeBetaDiskGuestOSFeatureTypeEnumToProto(o.Type),
	}
	for _, r := range o.TypeAlt {
		p.TypeAlt = append(p.TypeAlt, betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum(betapb.ComputeBetaDiskGuestOSFeatureTypeAltEnum_value[string(r)]))
	}
	return p
}

// DiskEncryptionKeyToProto converts a DiskEncryptionKey resource to its proto representation.
func ComputeBetaDiskEncryptionKeyToProto(o *beta.DiskEncryptionKey) *betapb.ComputeBetaDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaDiskEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// DiskGuestOSFeaturesToProto converts a DiskGuestOSFeatures resource to its proto representation.
func ComputeBetaDiskGuestOSFeaturesToProto(o *beta.DiskGuestOSFeatures) *betapb.ComputeBetaDiskGuestOSFeatures {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaDiskGuestOSFeatures{
		Type: ComputeBetaDiskGuestOSFeaturesTypeEnumToProto(o.Type),
	}
	for _, r := range o.TypeAlts {
		p.TypeAlts = append(p.TypeAlts, betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum(betapb.ComputeBetaDiskGuestOSFeaturesTypeAltsEnum_value[string(r)]))
	}
	return p
}

// DiskToProto converts a Disk resource to its proto representation.
func DiskToProto(resource *beta.Disk) *betapb.ComputeBetaDisk {
	p := &betapb.ComputeBetaDisk{
		SelfLink:                    dcl.ValueOrEmptyString(resource.SelfLink),
		Description:                 dcl.ValueOrEmptyString(resource.Description),
		DiskEncryptionKey:           ComputeBetaDiskEncryptionKeyToProto(resource.DiskEncryptionKey),
		LabelFingerprint:            dcl.ValueOrEmptyString(resource.LabelFingerprint),
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		Region:                      dcl.ValueOrEmptyString(resource.Region),
		SizeGb:                      dcl.ValueOrEmptyInt64(resource.SizeGb),
		SourceImage:                 dcl.ValueOrEmptyString(resource.SourceImage),
		SourceImageEncryptionKey:    ComputeBetaDiskEncryptionKeyToProto(resource.SourceImageEncryptionKey),
		SourceImageId:               dcl.ValueOrEmptyString(resource.SourceImageId),
		SourceSnapshot:              dcl.ValueOrEmptyString(resource.SourceSnapshot),
		SourceSnapshotEncryptionKey: ComputeBetaDiskEncryptionKeyToProto(resource.SourceSnapshotEncryptionKey),
		SourceSnapshotId:            dcl.ValueOrEmptyString(resource.SourceSnapshotId),
		Type:                        dcl.ValueOrEmptyString(resource.Type),
		Zone:                        dcl.ValueOrEmptyString(resource.Zone),
		Project:                     dcl.ValueOrEmptyString(resource.Project),
		Id:                          dcl.ValueOrEmptyInt64(resource.Id),
		Status:                      ComputeBetaDiskStatusEnumToProto(resource.Status),
		Options:                     dcl.ValueOrEmptyString(resource.Options),
		LastAttachTimestamp:         dcl.ValueOrEmptyString(resource.LastAttachTimestamp),
		LastDetachTimestamp:         dcl.ValueOrEmptyString(resource.LastDetachTimestamp),
		PhysicalBlockSizeBytes:      dcl.ValueOrEmptyInt64(resource.PhysicalBlockSizeBytes),
		SourceDisk:                  dcl.ValueOrEmptyString(resource.SourceDisk),
		SourceDiskId:                dcl.ValueOrEmptyString(resource.SourceDiskId),
		Location:                    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.GuestOSFeature {
		p.GuestOsFeature = append(p.GuestOsFeature, ComputeBetaDiskGuestOSFeatureToProto(&r))
	}
	for _, r := range resource.License {
		p.License = append(p.License, r)
	}
	for _, r := range resource.ReplicaZones {
		p.ReplicaZones = append(p.ReplicaZones, r)
	}
	for _, r := range resource.ResourcePolicy {
		p.ResourcePolicy = append(p.ResourcePolicy, r)
	}
	for _, r := range resource.Licenses {
		p.Licenses = append(p.Licenses, r)
	}
	for _, r := range resource.GuestOSFeatures {
		p.GuestOsFeatures = append(p.GuestOsFeatures, ComputeBetaDiskGuestOSFeaturesToProto(&r))
	}
	for _, r := range resource.Users {
		p.Users = append(p.Users, r)
	}
	for _, r := range resource.LicenseCodes {
		p.LicenseCodes = append(p.LicenseCodes, r)
	}
	for _, r := range resource.ResourcePolicies {
		p.ResourcePolicies = append(p.ResourcePolicies, r)
	}

	return p
}

// ApplyDisk handles the gRPC request by passing it to the underlying Disk Apply() method.
func (s *DiskServer) applyDisk(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaDiskRequest) (*betapb.ComputeBetaDisk, error) {
	p := ProtoToDisk(request.GetResource())
	res, err := c.ApplyDisk(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DiskToProto(res)
	return r, nil
}

// ApplyDisk handles the gRPC request by passing it to the underlying Disk Apply() method.
func (s *DiskServer) ApplyComputeBetaDisk(ctx context.Context, request *betapb.ApplyComputeBetaDiskRequest) (*betapb.ComputeBetaDisk, error) {
	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDisk(ctx, cl, request)
}

// DeleteDisk handles the gRPC request by passing it to the underlying Disk Delete() method.
func (s *DiskServer) DeleteComputeBetaDisk(ctx context.Context, request *betapb.DeleteComputeBetaDiskRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDisk(ctx, ProtoToDisk(request.GetResource()))

}

// ListComputeBetaDisk handles the gRPC request by passing it to the underlying DiskList() method.
func (s *DiskServer) ListComputeBetaDisk(ctx context.Context, request *betapb.ListComputeBetaDiskRequest) (*betapb.ListComputeBetaDiskResponse, error) {
	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDisk(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaDisk
	for _, r := range resources.Items {
		rp := DiskToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaDiskResponse{Items: protos}, nil
}

func createConfigDisk(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
