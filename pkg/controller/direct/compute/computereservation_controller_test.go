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

package compute

import (
	"context"
	"sort"
	"testing"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
)

func TestCompareComputeReservation(t *testing.T) {
	tests := []struct {
		name          string
		desired       *krm.ComputeReservationSpec
		actual        *computepb.Reservation
		wantDiff      bool
		wantPaths     []string
		wantUpdateErr string
	}{
		{
			name: "brownfield reservation with omitted optional immutable fields",
			desired: &krm.ComputeReservationSpec{
				Zone: direct.LazyPtr("us-central1-a"),
				SpecificReservation: &krm.ReservationSpecificReservation{
					Count: direct.LazyPtr(int64(2)),
					InstanceProperties: &krm.ReservationInstanceProperties{
						MachineType: direct.LazyPtr("n1-highmem-8"),
					},
				},
			},
			actual: &computepb.Reservation{
				Name:                        proto.String("res-1"),
				Description:                 proto.String("pre-existing reservation"),
				Zone:                        proto.String("https://www.googleapis.com/compute/v1/projects/proj-1/zones/us-central1-a"),
				SpecificReservationRequired: proto.Bool(true),
				SpecificReservation: &computepb.AllocationSpecificSKUReservation{
					Count:        proto.Int64(2),
					AssuredCount: proto.Int64(2),
					InUseCount:   proto.Int64(1),
					InstanceProperties: &computepb.AllocationSpecificSKUAllocationReservedInstanceProperties{
						MachineType:    proto.String("n1-highmem-8"),
						MinCpuPlatform: proto.String("Intel Skylake"),
						GuestAccelerators: []*computepb.AcceleratorConfig{
							{
								AcceleratorCount: proto.Int32(1),
								AcceleratorType:  proto.String("nvidia-tesla-v100"),
							},
						},
						LocalSsds: []*computepb.AllocationSpecificSKUAllocationAllocatedInstancePropertiesReservedDisk{
							{
								DiskSizeGb: proto.Int64(375),
								Interface:  proto.String("NVME"),
							},
						},
					},
				},
				ShareSettings: &computepb.ShareSettings{
					ShareType: proto.String("LOCAL"),
				},
			},
			wantDiff:  false,
			wantPaths: nil,
		},
		{
			name: "brownfield reservation resize when optional immutable fields are omitted",
			desired: &krm.ComputeReservationSpec{
				Zone: direct.LazyPtr("us-central1-a"),
				SpecificReservation: &krm.ReservationSpecificReservation{
					Count: direct.LazyPtr(int64(4)),
					InstanceProperties: &krm.ReservationInstanceProperties{
						MachineType: direct.LazyPtr("e2-medium"),
					},
				},
			},
			actual: &computepb.Reservation{
				Name:                        proto.String("res-1"),
				Description:                 proto.String("created via gcloud"),
				Zone:                        proto.String("https://www.googleapis.com/compute/v1/projects/proj-1/zones/us-central1-a"),
				SpecificReservationRequired: proto.Bool(false),
				SpecificReservation: &computepb.AllocationSpecificSKUReservation{
					Count:      proto.Int64(1),
					InUseCount: proto.Int64(1),
					InstanceProperties: &computepb.AllocationSpecificSKUAllocationReservedInstanceProperties{
						MachineType:    proto.String("e2-medium"),
						MinCpuPlatform: proto.String("Intel Broadwell"),
					},
				},
				ShareSettings: &computepb.ShareSettings{
					ShareType: proto.String("LOCAL"),
				},
			},
			wantDiff:  true,
			wantPaths: []string{"specific_reservation"},
		},
		{
			name: "brownfield shared reservation with omitted shareSettings",
			desired: &krm.ComputeReservationSpec{
				Zone: direct.LazyPtr("us-central1-a"),
				SpecificReservation: &krm.ReservationSpecificReservation{
					Count: direct.LazyPtr(int64(1)),
					InstanceProperties: &krm.ReservationInstanceProperties{
						MachineType: direct.LazyPtr("n1-standard-1"),
					},
				},
			},
			actual: &computepb.Reservation{
				Name: proto.String("res-shared"),
				Zone: proto.String("https://www.googleapis.com/compute/v1/projects/proj-1/zones/us-central1-a"),
				SpecificReservation: &computepb.AllocationSpecificSKUReservation{
					Count: proto.Int64(1),
					InstanceProperties: &computepb.AllocationSpecificSKUAllocationReservedInstanceProperties{
						MachineType: proto.String("n1-standard-1"),
					},
				},
				ShareSettings: &computepb.ShareSettings{
					ShareType: proto.String("SPECIFIC_PROJECTS"),
					ProjectMap: map[string]*computepb.ShareSettingsProjectConfig{
						"111111": {ProjectId: proto.String("111111")},
					},
				},
			},
			wantDiff:  false,
			wantPaths: nil,
		},
		{
			name: "shared reservation project removal via projectMap",
			desired: &krm.ComputeReservationSpec{
				Zone: direct.LazyPtr("us-central1-a"),
				SpecificReservation: &krm.ReservationSpecificReservation{
					Count: direct.LazyPtr(int64(1)),
					InstanceProperties: &krm.ReservationInstanceProperties{
						MachineType: direct.LazyPtr("n1-standard-1"),
					},
				},
				ShareSettings: &krm.ShareSettings{
					ShareType: direct.LazyPtr("SPECIFIC_PROJECTS"),
					ProjectMap: []krm.ShareSettingsProjectMap{
						{
							KeyRef: &refs.ExtendedProjectRef{External: "111111"},
							Value: &krm.ShareSettingsProjectConfig{
								ProjectIDRef: &refs.ProjectRef{External: "111111"},
							},
						},
					},
				},
			},
			actual: &computepb.Reservation{
				Name: proto.String("res-shared"),
				Zone: proto.String("https://www.googleapis.com/compute/v1/projects/proj-1/zones/us-central1-a"),
				SpecificReservation: &computepb.AllocationSpecificSKUReservation{
					Count: proto.Int64(1),
					InstanceProperties: &computepb.AllocationSpecificSKUAllocationReservedInstanceProperties{
						MachineType: proto.String("n1-standard-1"),
					},
				},
				ShareSettings: &computepb.ShareSettings{
					ShareType: proto.String("SPECIFIC_PROJECTS"),
					ProjectMap: map[string]*computepb.ShareSettingsProjectConfig{
						"111111": {ProjectId: proto.String("111111")},
						"222222": {ProjectId: proto.String("222222")},
					},
				},
			},
			wantDiff:  true,
			wantPaths: []string{"share_settings"},
		},
		{
			name: "explicit change to immutable description fails validation",
			desired: &krm.ComputeReservationSpec{
				Description: direct.LazyPtr("updated description"),
				Zone:        direct.LazyPtr("us-central1-a"),
				SpecificReservation: &krm.ReservationSpecificReservation{
					Count: direct.LazyPtr(int64(1)),
					InstanceProperties: &krm.ReservationInstanceProperties{
						MachineType: direct.LazyPtr("e2-medium"),
					},
				},
			},
			actual: &computepb.Reservation{
				Name:        proto.String("res-1"),
				Description: proto.String("original description"),
				Zone:        proto.String("https://www.googleapis.com/compute/v1/projects/proj-1/zones/us-central1-a"),
				SpecificReservation: &computepb.AllocationSpecificSKUReservation{
					Count: proto.Int64(1),
					InstanceProperties: &computepb.AllocationSpecificSKUAllocationReservedInstanceProperties{
						MachineType: proto.String("e2-medium"),
					},
				},
			},
			wantDiff:      true,
			wantPaths:     []string{"description"},
			wantUpdateErr: `field "description" is immutable`,
		},
		{
			name: "explicit change to immutable minCpuPlatform fails validation",
			desired: &krm.ComputeReservationSpec{
				Zone: direct.LazyPtr("us-central1-a"),
				SpecificReservation: &krm.ReservationSpecificReservation{
					Count: direct.LazyPtr(int64(1)),
					InstanceProperties: &krm.ReservationInstanceProperties{
						MachineType:    direct.LazyPtr("e2-medium"),
						MinCpuPlatform: direct.LazyPtr("Intel Ice Lake"),
					},
				},
			},
			actual: &computepb.Reservation{
				Name: proto.String("res-1"),
				Zone: proto.String("https://www.googleapis.com/compute/v1/projects/proj-1/zones/us-central1-a"),
				SpecificReservation: &computepb.AllocationSpecificSKUReservation{
					Count: proto.Int64(1),
					InstanceProperties: &computepb.AllocationSpecificSKUAllocationReservedInstanceProperties{
						MachineType:    proto.String("e2-medium"),
						MinCpuPlatform: proto.String("Intel Broadwell"),
					},
				},
			},
			wantDiff:      true,
			wantPaths:     []string{"specific_reservation"},
			wantUpdateErr: `field "specific_reservation.instance_properties.min_cpu_platform" is immutable`,
		},
		{
			name: "explicit change to immutable shareType fails validation",
			desired: &krm.ComputeReservationSpec{
				Zone: direct.LazyPtr("us-central1-a"),
				SpecificReservation: &krm.ReservationSpecificReservation{
					Count: direct.LazyPtr(int64(1)),
					InstanceProperties: &krm.ReservationInstanceProperties{
						MachineType: direct.LazyPtr("e2-medium"),
					},
				},
				ShareSettings: &krm.ShareSettings{
					ShareType: direct.LazyPtr("SPECIFIC_PROJECTS"),
				},
			},
			actual: &computepb.Reservation{
				Name: proto.String("res-1"),
				Zone: proto.String("https://www.googleapis.com/compute/v1/projects/proj-1/zones/us-central1-a"),
				SpecificReservation: &computepb.AllocationSpecificSKUReservation{
					Count: proto.Int64(1),
					InstanceProperties: &computepb.AllocationSpecificSKUAllocationReservedInstanceProperties{
						MachineType: proto.String("e2-medium"),
					},
				},
				ShareSettings: &computepb.ShareSettings{
					ShareType: proto.String("LOCAL"),
				},
			},
			wantDiff:      true,
			wantPaths:     []string{"share_settings"},
			wantUpdateErr: `field "share_settings.share_type" is immutable`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			diff, updateMask, err := compareComputeReservation(ctx, tc.desired, tc.actual)
			if tc.wantUpdateErr != "" {
				if err == nil || err.Error() != tc.wantUpdateErr {
					t.Fatalf("compareComputeReservation() error = %v, want %q", err, tc.wantUpdateErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("compareComputeReservation() unexpected error: %v", err)
			}

			if diff.HasDiff() != tc.wantDiff {
				t.Errorf("compareComputeReservation() HasDiff() = %v, want %v", diff.HasDiff(), tc.wantDiff)
			}

			gotPaths := updateMask.GetPaths()
			sort.Strings(gotPaths)
			sort.Strings(tc.wantPaths)
			if d := cmp.Diff(tc.wantPaths, gotPaths); d != "" {
				t.Errorf("compareComputeReservation() paths diff (-want +got):\n%s", d)
			}
		})
	}
}
