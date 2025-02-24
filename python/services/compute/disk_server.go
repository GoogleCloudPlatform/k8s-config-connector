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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Disk.
type DiskServer struct{}

// ProtoToDiskGuestOSFeatureTypeEnum converts a DiskGuestOSFeatureTypeEnum enum from its proto representation.
func ProtoToComputeDiskGuestOSFeatureTypeEnum(e computepb.ComputeDiskGuestOSFeatureTypeEnum) *compute.DiskGuestOSFeatureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeDiskGuestOSFeatureTypeEnum_name[int32(e)]; ok {
		e := compute.DiskGuestOSFeatureTypeEnum(n[len("ComputeDiskGuestOSFeatureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOSFeatureTypeAltEnum converts a DiskGuestOSFeatureTypeAltEnum enum from its proto representation.
func ProtoToComputeDiskGuestOSFeatureTypeAltEnum(e computepb.ComputeDiskGuestOSFeatureTypeAltEnum) *compute.DiskGuestOSFeatureTypeAltEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeDiskGuestOSFeatureTypeAltEnum_name[int32(e)]; ok {
		e := compute.DiskGuestOSFeatureTypeAltEnum(n[len("ComputeDiskGuestOSFeatureTypeAltEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskStatusEnum converts a DiskStatusEnum enum from its proto representation.
func ProtoToComputeDiskStatusEnum(e computepb.ComputeDiskStatusEnum) *compute.DiskStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeDiskStatusEnum_name[int32(e)]; ok {
		e := compute.DiskStatusEnum(n[len("ComputeDiskStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOSFeaturesTypeEnum converts a DiskGuestOSFeaturesTypeEnum enum from its proto representation.
func ProtoToComputeDiskGuestOSFeaturesTypeEnum(e computepb.ComputeDiskGuestOSFeaturesTypeEnum) *compute.DiskGuestOSFeaturesTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeDiskGuestOSFeaturesTypeEnum_name[int32(e)]; ok {
		e := compute.DiskGuestOSFeaturesTypeEnum(n[len("ComputeDiskGuestOSFeaturesTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOSFeaturesTypeAltsEnum converts a DiskGuestOSFeaturesTypeAltsEnum enum from its proto representation.
func ProtoToComputeDiskGuestOSFeaturesTypeAltsEnum(e computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum) *compute.DiskGuestOSFeaturesTypeAltsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum_name[int32(e)]; ok {
		e := compute.DiskGuestOSFeaturesTypeAltsEnum(n[len("ComputeDiskGuestOSFeaturesTypeAltsEnum"):])
		return &e
	}
	return nil
}

// ProtoToDiskGuestOSFeature converts a DiskGuestOSFeature resource from its proto representation.
func ProtoToComputeDiskGuestOSFeature(p *computepb.ComputeDiskGuestOSFeature) *compute.DiskGuestOSFeature {
	if p == nil {
		return nil
	}
	obj := &compute.DiskGuestOSFeature{
		Type: ProtoToComputeDiskGuestOSFeatureTypeEnum(p.GetType()),
	}
	for _, r := range p.GetTypeAlt() {
		obj.TypeAlt = append(obj.TypeAlt, *ProtoToComputeDiskGuestOSFeatureTypeAltEnum(r))
	}
	return obj
}

// ProtoToDiskEncryptionKey converts a DiskEncryptionKey resource from its proto representation.
func ProtoToComputeDiskEncryptionKey(p *computepb.ComputeDiskEncryptionKey) *compute.DiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.DiskEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToDiskGuestOSFeatures converts a DiskGuestOSFeatures resource from its proto representation.
func ProtoToComputeDiskGuestOSFeatures(p *computepb.ComputeDiskGuestOSFeatures) *compute.DiskGuestOSFeatures {
	if p == nil {
		return nil
	}
	obj := &compute.DiskGuestOSFeatures{
		Type: ProtoToComputeDiskGuestOSFeaturesTypeEnum(p.GetType()),
	}
	for _, r := range p.GetTypeAlts() {
		obj.TypeAlts = append(obj.TypeAlts, *ProtoToComputeDiskGuestOSFeaturesTypeAltsEnum(r))
	}
	return obj
}

// ProtoToDisk converts a Disk resource from its proto representation.
func ProtoToDisk(p *computepb.ComputeDisk) *compute.Disk {
	obj := &compute.Disk{
		SelfLink:                    dcl.StringOrNil(p.SelfLink),
		Description:                 dcl.StringOrNil(p.Description),
		DiskEncryptionKey:           ProtoToComputeDiskEncryptionKey(p.GetDiskEncryptionKey()),
		LabelFingerprint:            dcl.StringOrNil(p.LabelFingerprint),
		Name:                        dcl.StringOrNil(p.Name),
		Region:                      dcl.StringOrNil(p.Region),
		SizeGb:                      dcl.Int64OrNil(p.SizeGb),
		SourceImage:                 dcl.StringOrNil(p.SourceImage),
		SourceImageEncryptionKey:    ProtoToComputeDiskEncryptionKey(p.GetSourceImageEncryptionKey()),
		SourceImageId:               dcl.StringOrNil(p.SourceImageId),
		SourceSnapshot:              dcl.StringOrNil(p.SourceSnapshot),
		SourceSnapshotEncryptionKey: ProtoToComputeDiskEncryptionKey(p.GetSourceSnapshotEncryptionKey()),
		SourceSnapshotId:            dcl.StringOrNil(p.SourceSnapshotId),
		Type:                        dcl.StringOrNil(p.Type),
		Zone:                        dcl.StringOrNil(p.Zone),
		Project:                     dcl.StringOrNil(p.Project),
		Id:                          dcl.Int64OrNil(p.Id),
		Status:                      ProtoToComputeDiskStatusEnum(p.GetStatus()),
		Options:                     dcl.StringOrNil(p.Options),
		LastAttachTimestamp:         dcl.StringOrNil(p.LastAttachTimestamp),
		LastDetachTimestamp:         dcl.StringOrNil(p.LastDetachTimestamp),
		PhysicalBlockSizeBytes:      dcl.Int64OrNil(p.PhysicalBlockSizeBytes),
		SourceDisk:                  dcl.StringOrNil(p.SourceDisk),
		SourceDiskId:                dcl.StringOrNil(p.SourceDiskId),
		Location:                    dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetGuestOsFeature() {
		obj.GuestOSFeature = append(obj.GuestOSFeature, *ProtoToComputeDiskGuestOSFeature(r))
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
		obj.GuestOSFeatures = append(obj.GuestOSFeatures, *ProtoToComputeDiskGuestOSFeatures(r))
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
func ComputeDiskGuestOSFeatureTypeEnumToProto(e *compute.DiskGuestOSFeatureTypeEnum) computepb.ComputeDiskGuestOSFeatureTypeEnum {
	if e == nil {
		return computepb.ComputeDiskGuestOSFeatureTypeEnum(0)
	}
	if v, ok := computepb.ComputeDiskGuestOSFeatureTypeEnum_value["DiskGuestOSFeatureTypeEnum"+string(*e)]; ok {
		return computepb.ComputeDiskGuestOSFeatureTypeEnum(v)
	}
	return computepb.ComputeDiskGuestOSFeatureTypeEnum(0)
}

// DiskGuestOSFeatureTypeAltEnumToProto converts a DiskGuestOSFeatureTypeAltEnum enum to its proto representation.
func ComputeDiskGuestOSFeatureTypeAltEnumToProto(e *compute.DiskGuestOSFeatureTypeAltEnum) computepb.ComputeDiskGuestOSFeatureTypeAltEnum {
	if e == nil {
		return computepb.ComputeDiskGuestOSFeatureTypeAltEnum(0)
	}
	if v, ok := computepb.ComputeDiskGuestOSFeatureTypeAltEnum_value["DiskGuestOSFeatureTypeAltEnum"+string(*e)]; ok {
		return computepb.ComputeDiskGuestOSFeatureTypeAltEnum(v)
	}
	return computepb.ComputeDiskGuestOSFeatureTypeAltEnum(0)
}

// DiskStatusEnumToProto converts a DiskStatusEnum enum to its proto representation.
func ComputeDiskStatusEnumToProto(e *compute.DiskStatusEnum) computepb.ComputeDiskStatusEnum {
	if e == nil {
		return computepb.ComputeDiskStatusEnum(0)
	}
	if v, ok := computepb.ComputeDiskStatusEnum_value["DiskStatusEnum"+string(*e)]; ok {
		return computepb.ComputeDiskStatusEnum(v)
	}
	return computepb.ComputeDiskStatusEnum(0)
}

// DiskGuestOSFeaturesTypeEnumToProto converts a DiskGuestOSFeaturesTypeEnum enum to its proto representation.
func ComputeDiskGuestOSFeaturesTypeEnumToProto(e *compute.DiskGuestOSFeaturesTypeEnum) computepb.ComputeDiskGuestOSFeaturesTypeEnum {
	if e == nil {
		return computepb.ComputeDiskGuestOSFeaturesTypeEnum(0)
	}
	if v, ok := computepb.ComputeDiskGuestOSFeaturesTypeEnum_value["DiskGuestOSFeaturesTypeEnum"+string(*e)]; ok {
		return computepb.ComputeDiskGuestOSFeaturesTypeEnum(v)
	}
	return computepb.ComputeDiskGuestOSFeaturesTypeEnum(0)
}

// DiskGuestOSFeaturesTypeAltsEnumToProto converts a DiskGuestOSFeaturesTypeAltsEnum enum to its proto representation.
func ComputeDiskGuestOSFeaturesTypeAltsEnumToProto(e *compute.DiskGuestOSFeaturesTypeAltsEnum) computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum {
	if e == nil {
		return computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum(0)
	}
	if v, ok := computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum_value["DiskGuestOSFeaturesTypeAltsEnum"+string(*e)]; ok {
		return computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum(v)
	}
	return computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum(0)
}

// DiskGuestOSFeatureToProto converts a DiskGuestOSFeature resource to its proto representation.
func ComputeDiskGuestOSFeatureToProto(o *compute.DiskGuestOSFeature) *computepb.ComputeDiskGuestOSFeature {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeDiskGuestOSFeature{
		Type: ComputeDiskGuestOSFeatureTypeEnumToProto(o.Type),
	}
	for _, r := range o.TypeAlt {
		p.TypeAlt = append(p.TypeAlt, computepb.ComputeDiskGuestOSFeatureTypeAltEnum(computepb.ComputeDiskGuestOSFeatureTypeAltEnum_value[string(r)]))
	}
	return p
}

// DiskEncryptionKeyToProto converts a DiskEncryptionKey resource to its proto representation.
func ComputeDiskEncryptionKeyToProto(o *compute.DiskEncryptionKey) *computepb.ComputeDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeDiskEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// DiskGuestOSFeaturesToProto converts a DiskGuestOSFeatures resource to its proto representation.
func ComputeDiskGuestOSFeaturesToProto(o *compute.DiskGuestOSFeatures) *computepb.ComputeDiskGuestOSFeatures {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeDiskGuestOSFeatures{
		Type: ComputeDiskGuestOSFeaturesTypeEnumToProto(o.Type),
	}
	for _, r := range o.TypeAlts {
		p.TypeAlts = append(p.TypeAlts, computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum(computepb.ComputeDiskGuestOSFeaturesTypeAltsEnum_value[string(r)]))
	}
	return p
}

// DiskToProto converts a Disk resource to its proto representation.
func DiskToProto(resource *compute.Disk) *computepb.ComputeDisk {
	p := &computepb.ComputeDisk{
		SelfLink:                    dcl.ValueOrEmptyString(resource.SelfLink),
		Description:                 dcl.ValueOrEmptyString(resource.Description),
		DiskEncryptionKey:           ComputeDiskEncryptionKeyToProto(resource.DiskEncryptionKey),
		LabelFingerprint:            dcl.ValueOrEmptyString(resource.LabelFingerprint),
		Name:                        dcl.ValueOrEmptyString(resource.Name),
		Region:                      dcl.ValueOrEmptyString(resource.Region),
		SizeGb:                      dcl.ValueOrEmptyInt64(resource.SizeGb),
		SourceImage:                 dcl.ValueOrEmptyString(resource.SourceImage),
		SourceImageEncryptionKey:    ComputeDiskEncryptionKeyToProto(resource.SourceImageEncryptionKey),
		SourceImageId:               dcl.ValueOrEmptyString(resource.SourceImageId),
		SourceSnapshot:              dcl.ValueOrEmptyString(resource.SourceSnapshot),
		SourceSnapshotEncryptionKey: ComputeDiskEncryptionKeyToProto(resource.SourceSnapshotEncryptionKey),
		SourceSnapshotId:            dcl.ValueOrEmptyString(resource.SourceSnapshotId),
		Type:                        dcl.ValueOrEmptyString(resource.Type),
		Zone:                        dcl.ValueOrEmptyString(resource.Zone),
		Project:                     dcl.ValueOrEmptyString(resource.Project),
		Id:                          dcl.ValueOrEmptyInt64(resource.Id),
		Status:                      ComputeDiskStatusEnumToProto(resource.Status),
		Options:                     dcl.ValueOrEmptyString(resource.Options),
		LastAttachTimestamp:         dcl.ValueOrEmptyString(resource.LastAttachTimestamp),
		LastDetachTimestamp:         dcl.ValueOrEmptyString(resource.LastDetachTimestamp),
		PhysicalBlockSizeBytes:      dcl.ValueOrEmptyInt64(resource.PhysicalBlockSizeBytes),
		SourceDisk:                  dcl.ValueOrEmptyString(resource.SourceDisk),
		SourceDiskId:                dcl.ValueOrEmptyString(resource.SourceDiskId),
		Location:                    dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.GuestOSFeature {
		p.GuestOsFeature = append(p.GuestOsFeature, ComputeDiskGuestOSFeatureToProto(&r))
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
		p.GuestOsFeatures = append(p.GuestOsFeatures, ComputeDiskGuestOSFeaturesToProto(&r))
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
func (s *DiskServer) applyDisk(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeDiskRequest) (*computepb.ComputeDisk, error) {
	p := ProtoToDisk(request.GetResource())
	res, err := c.ApplyDisk(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DiskToProto(res)
	return r, nil
}

// ApplyDisk handles the gRPC request by passing it to the underlying Disk Apply() method.
func (s *DiskServer) ApplyComputeDisk(ctx context.Context, request *computepb.ApplyComputeDiskRequest) (*computepb.ComputeDisk, error) {
	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDisk(ctx, cl, request)
}

// DeleteDisk handles the gRPC request by passing it to the underlying Disk Delete() method.
func (s *DiskServer) DeleteComputeDisk(ctx context.Context, request *computepb.DeleteComputeDiskRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDisk(ctx, ProtoToDisk(request.GetResource()))

}

// ListComputeDisk handles the gRPC request by passing it to the underlying DiskList() method.
func (s *DiskServer) ListComputeDisk(ctx context.Context, request *computepb.ListComputeDiskRequest) (*computepb.ListComputeDiskResponse, error) {
	cl, err := createConfigDisk(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDisk(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeDisk
	for _, r := range resources.Items {
		rp := DiskToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeDiskResponse{Items: protos}, nil
}

func createConfigDisk(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
