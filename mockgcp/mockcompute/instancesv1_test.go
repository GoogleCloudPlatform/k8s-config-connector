// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockcompute

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// denyUnknownProjectStore reports unknown projects as PermissionDenied, as the real project store does for
// public image projects such as cos-cloud.
type denyUnknownProjectStore struct {
	*mockProjectStore
}

func (s *denyUnknownProjectStore) GetProjectByID(projectID string) (*projects.ProjectData, error) {
	if p, ok := s.projectsByID[projectID]; ok {
		return p, nil
	}
	return nil, status.Errorf(codes.PermissionDenied, "project %q: permission denied", projectID)
}

func (s *denyUnknownProjectStore) GetProjectByIDOrNumber(idOrNumber string) (*projects.ProjectData, error) {
	if p, ok := s.projectsByNumber[idOrNumber]; ok {
		return p, nil
	}
	return s.GetProjectByID(idOrNumber)
}

type testComputeServices struct {
	ctx       context.Context
	instances *InstancesV1
	disks     *DisksV1
	addresses *RegionalAddressesV1
	images    *ImagesV1
}

func newTestComputeServices() *testComputeServices {
	testProject := &projects.ProjectData{ID: "test-project", Number: 123456789012}
	otherProject := &projects.ProjectData{ID: "other-project", Number: 987654321098}
	projectStore := &denyUnknownProjectStore{mockProjectStore: &mockProjectStore{
		projectsByID: map[string]*projects.ProjectData{
			"test-project":  testProject,
			"other-project": otherProject,
		},
		projectsByNumber: map[string]*projects.ProjectData{
			"123456789012": testProject,
			"987654321098": otherProject,
		},
	}}
	env := &common.MockEnvironment{
		Projects: projectStore,
	}
	svc := New(env, storage.NewInMemoryStorage()).(*MockService)
	return &testComputeServices{
		ctx:       context.Background(),
		instances: &InstancesV1{MockService: svc},
		disks:     &DisksV1{MockService: svc},
		addresses: &RegionalAddressesV1{MockService: svc},
		images:    &ImagesV1{MockService: svc},
	}
}

func (s *testComputeServices) insertAddress(t *testing.T, project, region, name, ip string) {
	t.Helper()
	if _, err := s.addresses.Insert(s.ctx, &pb.InsertAddressRequest{
		Project: project,
		Region:  region,
		AddressResource: &pb.Address{
			Name:        PtrTo(name),
			Address:     PtrTo(ip),
			AddressType: PtrTo("INTERNAL"),
		},
	}); err != nil {
		t.Fatalf("Insert address %s/%s: %v", project, name, err)
	}
}

func (s *testComputeServices) insertDisk(t *testing.T, disk *pb.Disk) {
	t.Helper()
	if _, err := s.disks.Insert(s.ctx, &pb.InsertDiskRequest{
		Project:      "test-project",
		Zone:         "us-central1-a",
		DiskResource: disk,
	}); err != nil {
		t.Fatalf("Insert disk %s: %v", disk.GetName(), err)
	}
}

func (s *testComputeServices) insertInstance(project string, instance *pb.Instance) error {
	_, err := s.instances.Insert(s.ctx, &pb.InsertInstanceRequest{
		Project:          project,
		Zone:             "us-central1-a",
		InstanceResource: instance,
	})
	return err
}

func (s *testComputeServices) getInstance(t *testing.T, project, name string) *pb.Instance {
	t.Helper()
	instance, err := s.instances.Get(s.ctx, &pb.GetInstanceRequest{
		Project:  project,
		Zone:     "us-central1-a",
		Instance: name,
	})
	if err != nil {
		t.Fatalf("Get instance %s/%s: %v", project, name, err)
	}
	return instance
}

func bootDiskWithImage(sourceImage string) *pb.AttachedDisk {
	return &pb.AttachedDisk{
		Boot:       PtrTo(true),
		AutoDelete: PtrTo(true),
		InitializeParams: &pb.AttachedDiskInitializeParams{
			SourceImage: PtrTo(sourceImage),
		},
	}
}

func TestInstancesV1_Insert_ResolvesNetworkIP(t *testing.T) {
	tests := []struct {
		name      string
		networkIP string
	}{
		{
			name:      "www URL",
			networkIP: "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/addresses/my-address",
		},
		{
			name:      "compute URL",
			networkIP: "https://compute.googleapis.com/compute/v1/projects/test-project/regions/us-central1/addresses/my-address",
		},
		{
			name:      "projects path",
			networkIP: "projects/test-project/regions/us-central1/addresses/my-address",
		},
		{
			name:      "regions path",
			networkIP: "regions/us-central1/addresses/my-address",
		},
		{
			name:      "plain IP",
			networkIP: "10.128.0.42",
		},
		{
			name:      "project number URL",
			networkIP: "https://www.googleapis.com/compute/v1/projects/123456789012/regions/us-central1/addresses/my-address",
		},
		{
			name:      "project number path",
			networkIP: "projects/123456789012/regions/us-central1/addresses/my-address",
		},
		{
			name:      "beta URL",
			networkIP: "https://www.googleapis.com/compute/beta/projects/test-project/regions/us-central1/addresses/my-address",
		},
		{
			name:      "alpha URL",
			networkIP: "https://compute.googleapis.com/compute/alpha/projects/test-project/regions/us-central1/addresses/my-address",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			s.insertAddress(t, "test-project", "us-central1", "my-address", "10.128.0.42")

			if err := s.insertInstance("test-project", &pb.Instance{
				Name:              PtrTo("inst"),
				MachineType:       PtrTo("n1-standard-1"),
				NetworkInterfaces: []*pb.NetworkInterface{{NetworkIP: PtrTo(tc.networkIP)}},
			}); err != nil {
				t.Fatalf("Insert() unexpected error: %v", err)
			}

			got := s.getInstance(t, "test-project", "inst").GetNetworkInterfaces()[0].GetNetworkIP()
			if got != "10.128.0.42" {
				t.Errorf("networkIP = %q, want %q", got, "10.128.0.42")
			}
		})
	}
}

func TestInstancesV1_Insert_NetworkIPErrors(t *testing.T) {
	tests := []struct {
		name      string
		networkIP string
		wantCode  codes.Code
		wantMsg   string
	}{
		{
			name:      "missing address",
			networkIP: "projects/test-project/regions/us-central1/addresses/missing",
			wantCode:  codes.NotFound,
		},
		{
			name:      "address reserved by another project",
			networkIP: "projects/other-project/regions/us-central1/addresses/shared",
			wantCode:  codes.InvalidArgument,
			wantMsg:   "reserved by another project",
		},
		{
			name:      "address in another region",
			networkIP: "projects/test-project/regions/us-east1/addresses/east",
			wantCode:  codes.InvalidArgument,
			wantMsg:   "Invalid value for field 'resource.networkInterfaces[0].networkIP'",
		},
		{
			name:      "address without an IP",
			networkIP: "projects/test-project/regions/us-central1/addresses/no-ip",
			wantCode:  codes.InvalidArgument,
			wantMsg:   "Invalid value for field 'resource.networkInterfaces[0].networkIP'",
		},
		{
			name:      "malformed",
			networkIP: "not-an-ip",
			wantCode:  codes.InvalidArgument,
			wantMsg:   "Invalid value for field 'resource.networkInterfaces[0].networkIP'",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			s.insertAddress(t, "other-project", "us-central1", "shared", "10.128.0.50")
			s.insertAddress(t, "test-project", "us-east1", "east", "10.142.0.7")
			s.insertAddress(t, "test-project", "us-central1", "no-ip", "")

			err := s.insertInstance("test-project", &pb.Instance{
				Name:              PtrTo("inst"),
				MachineType:       PtrTo("n1-standard-1"),
				NetworkInterfaces: []*pb.NetworkInterface{{NetworkIP: PtrTo(tc.networkIP)}},
			})
			if status.Code(err) != tc.wantCode {
				t.Fatalf("Insert() code = %v (err %v), want %v", status.Code(err), err, tc.wantCode)
			}
			if tc.wantMsg != "" && !strings.Contains(status.Convert(err).Message(), tc.wantMsg) {
				t.Errorf("Insert() message = %q, want it to contain %q", status.Convert(err).Message(), tc.wantMsg)
			}

			// A failed insert must not create the instance.
			if _, err := s.instances.Get(s.ctx, &pb.GetInstanceRequest{Project: "test-project", Zone: "us-central1-a", Instance: "inst"}); status.Code(err) != codes.NotFound {
				t.Errorf("Get() after failed insert code = %v, want %v", status.Code(err), codes.NotFound)
			}
		})
	}
}

func TestInstancesV1_Insert_CreatesBootDiskFromImageFamily(t *testing.T) {
	tests := []struct {
		name            string
		sourceImage     string
		wantImageSuffix string
	}{
		{
			name:            "image family path",
			sourceImage:     "projects/cos-cloud/global/images/family/cos-81-lts",
			wantImageSuffix: "cos-cloud/global/images/cos-81-12871-1196-0",
		},
		{
			name:            "shorthand",
			sourceImage:     "cos-cloud/cos-81-lts",
			wantImageSuffix: "cos-cloud/global/images/cos-81-12871-1196-0",
		},
		{
			name:            "current COS LTS family",
			sourceImage:     "projects/cos-cloud/global/images/family/cos-121-lts",
			wantImageSuffix: "cos-cloud/global/images/cos-121-18867-90-42",
		},
		{
			name:            "debian family",
			sourceImage:     "debian-cloud/debian-11",
			wantImageSuffix: "debian-cloud/global/images/debian-11-bullseye-v20231010",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			if err := s.insertInstance("test-project", &pb.Instance{
				Name:        PtrTo("inst"),
				MachineType: PtrTo("n1-standard-1"),
				Disks:       []*pb.AttachedDisk{bootDiskWithImage(tc.sourceImage)},
			}); err != nil {
				t.Fatalf("Insert() unexpected error: %v", err)
			}

			instance := s.getInstance(t, "test-project", "inst")
			attached := instance.GetDisks()[0]
			wantSource := "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/disks/inst"
			if attached.GetSource() != wantSource {
				t.Errorf("disk source = %q, want %q", attached.GetSource(), wantSource)
			}
			if attached.InitializeParams != nil {
				t.Errorf("disk initializeParams = %v, want nil", attached.InitializeParams)
			}

			disk, err := s.disks.Get(s.ctx, &pb.GetDiskRequest{Project: "test-project", Zone: "us-central1-a", Disk: "inst"})
			if err != nil {
				t.Fatalf("Get disk: %v", err)
			}
			if !strings.HasSuffix(disk.GetSourceImage(), tc.wantImageSuffix) {
				t.Errorf("disk sourceImage = %q, want suffix %q", disk.GetSourceImage(), tc.wantImageSuffix)
			}
			wantUser := "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/instances/inst"
			found := false
			for _, user := range disk.GetUsers() {
				if user == wantUser {
					found = true
				}
			}
			if !found {
				t.Errorf("disk users = %v, want to contain %q", disk.GetUsers(), wantUser)
			}
		})
	}
}

func TestInstancesV1_Insert_RollsBackBootDiskOnFailure(t *testing.T) {
	s := newTestComputeServices()
	if err := s.insertInstance("test-project", &pb.Instance{
		Name:        PtrTo("inst"),
		MachineType: PtrTo("n1-standard-1"),
	}); err != nil {
		t.Fatalf("Insert() unexpected error: %v", err)
	}

	// The instance name is taken, so the insert fails after the boot disk has been built.
	err := s.insertInstance("test-project", &pb.Instance{
		Name:        PtrTo("inst"),
		MachineType: PtrTo("n1-standard-1"),
		Disks:       []*pb.AttachedDisk{bootDiskWithImage("cos-cloud/cos-81-lts")},
	})
	if status.Code(err) != codes.AlreadyExists {
		t.Fatalf("Insert() code = %v (err %v), want %v", status.Code(err), err, codes.AlreadyExists)
	}

	_, err = s.disks.Get(s.ctx, &pb.GetDiskRequest{Project: "test-project", Zone: "us-central1-a", Disk: "inst"})
	if status.Code(err) != codes.NotFound {
		t.Errorf("Get disk after failed insert code = %v, want %v", status.Code(err), codes.NotFound)
	}
}

func TestInstancesV1_Delete_AutoDeletesBootDisk(t *testing.T) {
	tests := []struct {
		name       string
		autoDelete bool
	}{
		{name: "auto delete", autoDelete: true},
		{name: "keep disk", autoDelete: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			boot := bootDiskWithImage("cos-cloud/cos-81-lts")
			boot.AutoDelete = PtrTo(tc.autoDelete)
			if err := s.insertInstance("test-project", &pb.Instance{
				Name:        PtrTo("inst"),
				MachineType: PtrTo("n1-standard-1"),
				Disks:       []*pb.AttachedDisk{boot},
			}); err != nil {
				t.Fatalf("Insert() unexpected error: %v", err)
			}

			if _, err := s.instances.Delete(s.ctx, &pb.DeleteInstanceRequest{Project: "test-project", Zone: "us-central1-a", Instance: "inst"}); err != nil {
				t.Fatalf("Delete() unexpected error: %v", err)
			}

			disk, err := s.disks.Get(s.ctx, &pb.GetDiskRequest{Project: "test-project", Zone: "us-central1-a", Disk: "inst"})
			if tc.autoDelete {
				if status.Code(err) != codes.NotFound {
					t.Errorf("Get disk after delete code = %v, want %v", status.Code(err), codes.NotFound)
				}
				return
			}
			if err != nil {
				t.Fatalf("Get disk after delete: %v", err)
			}
			if len(disk.GetUsers()) != 0 {
				t.Errorf("disk users = %v, want none", disk.GetUsers())
			}
		})
	}
}

const (
	testInstanceSelfLink = "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/instances/inst"
	testDiskBase         = "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/disks/"
)

func TestInstancesV1_Insert_BootDiskNameCollision(t *testing.T) {
	s := newTestComputeServices()
	s.insertDisk(t, &pb.Disk{Name: PtrTo("inst")})

	err := s.insertInstance("test-project", &pb.Instance{
		Name:        PtrTo("inst"),
		MachineType: PtrTo("n1-standard-1"),
		Disks:       []*pb.AttachedDisk{bootDiskWithImage("cos-cloud/cos-81-lts")},
	})
	if status.Code(err) != codes.AlreadyExists {
		t.Fatalf("Insert() code = %v (err %v), want %v", status.Code(err), err, codes.AlreadyExists)
	}

	_, err = s.instances.Get(s.ctx, &pb.GetInstanceRequest{Project: "test-project", Zone: "us-central1-a", Instance: "inst"})
	if status.Code(err) != codes.NotFound {
		t.Errorf("Get instance after failed insert code = %v, want %v", status.Code(err), codes.NotFound)
	}
}

func TestInstancesV1_Insert_BootDiskParams(t *testing.T) {
	tests := []struct {
		name     string
		diskName string
		diskType string
		sizeGb   int64
		wantName string
		wantType string
		wantSize int64
	}{
		{
			name:     "custom name, size and bare disk type",
			diskName: "custom-boot",
			diskType: "pd-ssd",
			sizeGb:   42,
			wantName: "custom-boot",
			wantType: "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/diskTypes/pd-ssd",
			wantSize: 42,
		},
		{
			name:     "zones path disk type",
			diskName: "custom-boot",
			diskType: "zones/us-central1-a/diskTypes/pd-balanced",
			sizeGb:   42,
			wantName: "custom-boot",
			wantType: "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/diskTypes/pd-balanced",
			wantSize: 42,
		},
		{
			name:     "projects path disk type",
			diskName: "custom-boot",
			diskType: "projects/test-project/zones/us-central1-a/diskTypes/hyperdisk",
			sizeGb:   42,
			wantName: "custom-boot",
			wantType: "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/diskTypes/hyperdisk",
			wantSize: 42,
		},
		{
			name:     "full URL disk type",
			diskName: "custom-boot",
			diskType: "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/diskTypes/pd-extreme",
			sizeGb:   42,
			wantName: "custom-boot",
			wantType: "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/diskTypes/pd-extreme",
			wantSize: 42,
		},
		{
			name:     "defaults to pd-standard and the image size",
			wantName: "inst",
			wantType: "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/diskTypes/pd-standard",
			wantSize: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			boot := bootDiskWithImage("cos-cloud/cos-81-lts")
			boot.InitializeParams.DiskName = PtrTo(tc.diskName)
			boot.InitializeParams.DiskType = PtrTo(tc.diskType)
			if tc.sizeGb != 0 {
				boot.InitializeParams.DiskSizeGb = PtrTo(tc.sizeGb)
			}
			boot.InitializeParams.Labels = map[string]string{"env": "test"}
			boot.DiskEncryptionKey = &pb.CustomerEncryptionKey{KmsKeyName: PtrTo("projects/test-project/locations/us-central1/keyRings/r/cryptoKeys/k")}

			if err := s.insertInstance("test-project", &pb.Instance{
				Name:        PtrTo("inst"),
				MachineType: PtrTo("n1-standard-1"),
				Disks:       []*pb.AttachedDisk{boot},
			}); err != nil {
				t.Fatalf("Insert() unexpected error: %v", err)
			}

			attached := s.getInstance(t, "test-project", "inst").GetDisks()[0]
			if want := testDiskBase + tc.wantName; attached.GetSource() != want {
				t.Errorf("disk source = %q, want %q", attached.GetSource(), want)
			}

			disk, err := s.disks.Get(s.ctx, &pb.GetDiskRequest{Project: "test-project", Zone: "us-central1-a", Disk: tc.wantName})
			if err != nil {
				t.Fatalf("Get disk %s: %v", tc.wantName, err)
			}
			if disk.GetType() != tc.wantType {
				t.Errorf("disk type = %q, want %q", disk.GetType(), tc.wantType)
			}
			if disk.GetSizeGb() != tc.wantSize {
				t.Errorf("disk sizeGb = %d, want %d", disk.GetSizeGb(), tc.wantSize)
			}
			if disk.GetLabels()["env"] != "test" {
				t.Errorf("disk labels = %v, want env=test", disk.GetLabels())
			}
			if got := disk.GetDiskEncryptionKey().GetKmsKeyName(); got != "projects/test-project/locations/us-central1/keyRings/r/cryptoKeys/k" {
				t.Errorf("disk kmsKeyName = %q, want the instance's boot disk key", got)
			}
		})
	}
}

func TestInstancesV1_Insert_AttachedDiskRewrite(t *testing.T) {
	s := newTestComputeServices()
	s.insertDisk(t, &pb.Disk{Name: PtrTo("existing")})

	// Index 0 is an existing disk that is not created by this insert; index 1 is a boot disk with
	// user-set attached-disk fields, which must be preserved.
	boot := bootDiskWithImage("cos-cloud/cos-81-lts")
	boot.InitializeParams.DiskName = PtrTo("my-boot")
	boot.InitializeParams.DiskSizeGb = PtrTo(int64(42))
	boot.DeviceName = PtrTo("user-device")
	boot.Interface = PtrTo("NVME")
	boot.DiskSizeGb = PtrTo(int64(99))
	if err := s.insertInstance("test-project", &pb.Instance{
		Name:        PtrTo("inst"),
		MachineType: PtrTo("n1-standard-1"),
		Disks: []*pb.AttachedDisk{
			{Source: PtrTo(testDiskBase + "existing")},
			boot,
		},
	}); err != nil {
		t.Fatalf("Insert() unexpected error: %v", err)
	}

	disks := s.getInstance(t, "test-project", "inst").GetDisks()
	if disks[0].GetSource() != testDiskBase+"existing" || disks[0].GetIndex() != 0 {
		t.Errorf("existing disk = source %q index %d, want untouched", disks[0].GetSource(), disks[0].GetIndex())
	}
	if disks[0].GetKind() != "" {
		t.Errorf("existing disk kind = %q, want it not rewritten", disks[0].GetKind())
	}

	got := disks[1]
	checks := map[string][2]string{
		"kind":       {got.GetKind(), "compute#attachedDisk"},
		"type":       {got.GetType(), "PERSISTENT"},
		"mode":       {got.GetMode(), "READ_WRITE"},
		"deviceName": {got.GetDeviceName(), "user-device"},
		"interface":  {got.GetInterface(), "NVME"},
		"source":     {got.GetSource(), testDiskBase + "my-boot"},
	}
	for field, c := range checks {
		if c[0] != c[1] {
			t.Errorf("boot disk %s = %q, want %q", field, c[0], c[1])
		}
	}
	if got.GetIndex() != 1 {
		t.Errorf("boot disk index = %d, want 1", got.GetIndex())
	}
	if got.GetDiskSizeGb() != 99 {
		t.Errorf("boot disk diskSizeGb = %d, want the user-set 99", got.GetDiskSizeGb())
	}
	if got.InitializeParams != nil {
		t.Errorf("boot disk initializeParams = %v, want nil", got.InitializeParams)
	}
}

func TestInstancesV1_Insert_AttachedDiskDefaults(t *testing.T) {
	s := newTestComputeServices()
	boot := bootDiskWithImage("cos-cloud/cos-81-lts")
	boot.InitializeParams.DiskSizeGb = PtrTo(int64(42))
	if err := s.insertInstance("test-project", &pb.Instance{
		Name:        PtrTo("inst"),
		MachineType: PtrTo("n1-standard-1"),
		Disks:       []*pb.AttachedDisk{boot},
	}); err != nil {
		t.Fatalf("Insert() unexpected error: %v", err)
	}

	got := s.getInstance(t, "test-project", "inst").GetDisks()[0]
	checks := map[string][2]string{
		"kind":       {got.GetKind(), "compute#attachedDisk"},
		"type":       {got.GetType(), "PERSISTENT"},
		"mode":       {got.GetMode(), "READ_WRITE"},
		"deviceName": {got.GetDeviceName(), "persistent-disk-0"},
		"interface":  {got.GetInterface(), "SCSI"},
	}
	for field, c := range checks {
		if c[0] != c[1] {
			t.Errorf("boot disk %s = %q, want %q", field, c[0], c[1])
		}
	}
	if got.GetIndex() != 0 {
		t.Errorf("boot disk index = %d, want 0", got.GetIndex())
	}
	if got.GetDiskSizeGb() != 42 {
		t.Errorf("boot disk diskSizeGb = %d, want 42 from initializeParams", got.GetDiskSizeGb())
	}
}

func TestInstancesV1_Insert_LeavesDisksWithoutInitializeParamsAlone(t *testing.T) {
	tests := []struct {
		name string
		disk func() *pb.AttachedDisk
	}{
		{
			name: "scratch disk",
			disk: func() *pb.AttachedDisk {
				d := bootDiskWithImage("cos-cloud/cos-81-lts")
				d.Type = PtrTo("SCRATCH")
				return d
			},
		},
		{
			name: "boot disk that already has a source",
			disk: func() *pb.AttachedDisk {
				d := bootDiskWithImage("cos-cloud/cos-81-lts")
				d.Source = PtrTo(testDiskBase + "existing")
				return d
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			s.insertDisk(t, &pb.Disk{Name: PtrTo("existing")})
			disk := tc.disk()
			if err := s.insertInstance("test-project", &pb.Instance{
				Name:        PtrTo("inst"),
				MachineType: PtrTo("n1-standard-1"),
				Disks:       []*pb.AttachedDisk{disk},
			}); err != nil {
				t.Fatalf("Insert() unexpected error: %v", err)
			}

			_, err := s.disks.Get(s.ctx, &pb.GetDiskRequest{Project: "test-project", Zone: "us-central1-a", Disk: "inst"})
			if status.Code(err) != codes.NotFound {
				t.Errorf("Get disk inst code = %v, want %v: no disk should be created", status.Code(err), codes.NotFound)
			}
			got := s.getInstance(t, "test-project", "inst").GetDisks()[0]
			if got.InitializeParams == nil {
				t.Errorf("initializeParams = nil, want left untouched")
			}
			if got.GetSource() != disk.GetSource() {
				t.Errorf("source = %q, want %q", got.GetSource(), disk.GetSource())
			}
		})
	}
}

func TestInstancesV1_Insert_StoredCustomImage(t *testing.T) {
	tests := []struct {
		name        string
		sourceImage string
	}{
		{name: "relative image", sourceImage: "global/images/custom-img"},
		{name: "project image path", sourceImage: "projects/test-project/global/images/custom-img"},
		{name: "relative family", sourceImage: "global/images/family/custom-fam"},
		{name: "project family path", sourceImage: "projects/test-project/global/images/family/custom-fam"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			if _, err := s.images.Insert(s.ctx, &pb.InsertImageRequest{
				Project: "test-project",
				ImageResource: &pb.Image{
					Name:   PtrTo("custom-img"),
					Family: PtrTo("custom-fam"),
				},
			}); err != nil {
				t.Fatalf("Insert image: %v", err)
			}
			image, err := s.images.Get(s.ctx, &pb.GetImageRequest{Project: "test-project", Image: "custom-img"})
			if err != nil {
				t.Fatalf("Get image: %v", err)
			}

			if err := s.insertInstance("test-project", &pb.Instance{
				Name:        PtrTo("inst"),
				MachineType: PtrTo("n1-standard-1"),
				Disks:       []*pb.AttachedDisk{bootDiskWithImage(tc.sourceImage)},
			}); err != nil {
				t.Fatalf("Insert() unexpected error: %v", err)
			}

			disk, err := s.disks.Get(s.ctx, &pb.GetDiskRequest{Project: "test-project", Zone: "us-central1-a", Disk: "inst"})
			if err != nil {
				t.Fatalf("Get disk: %v", err)
			}
			if want := image.GetSelfLink(); disk.GetSourceImage() != want {
				t.Errorf("disk sourceImage = %q, want %q", disk.GetSourceImage(), want)
			}
			if want := fmt.Sprintf("%d", image.GetId()); disk.GetSourceImageId() != want {
				t.Errorf("disk sourceImageId = %q, want %q", disk.GetSourceImageId(), want)
			}
		})
	}
}

func TestInstancesV1_Delete_ReleasesAttachedDisks(t *testing.T) {
	tests := []struct {
		name       string
		autoDelete bool
	}{
		{name: "not auto-delete: instance removed from users", autoDelete: false},
		{name: "auto-delete: disk removed", autoDelete: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			const otherInstance = "https://www.googleapis.com/compute/v1/projects/test-project/zones/us-central1-a/instances/other"
			s.insertDisk(t, &pb.Disk{Name: PtrTo("extra"), Users: []string{testInstanceSelfLink, otherInstance}})

			if err := s.insertInstance("test-project", &pb.Instance{
				Name:        PtrTo("inst"),
				MachineType: PtrTo("n1-standard-1"),
				Disks: []*pb.AttachedDisk{{
					Source:     PtrTo(testDiskBase + "extra"),
					AutoDelete: PtrTo(tc.autoDelete),
				}},
			}); err != nil {
				t.Fatalf("Insert() unexpected error: %v", err)
			}

			if _, err := s.instances.Delete(s.ctx, &pb.DeleteInstanceRequest{Project: "test-project", Zone: "us-central1-a", Instance: "inst"}); err != nil {
				t.Fatalf("Delete() unexpected error: %v", err)
			}

			disk, err := s.disks.Get(s.ctx, &pb.GetDiskRequest{Project: "test-project", Zone: "us-central1-a", Disk: "extra"})
			if tc.autoDelete {
				if status.Code(err) != codes.NotFound {
					t.Errorf("Get disk after delete code = %v, want %v", status.Code(err), codes.NotFound)
				}
				return
			}
			if err != nil {
				t.Fatalf("Get disk after delete: %v", err)
			}
			if len(disk.GetUsers()) != 1 || disk.GetUsers()[0] != otherInstance {
				t.Errorf("disk users = %v, want only %q", disk.GetUsers(), otherInstance)
			}
		})
	}
}
