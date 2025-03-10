package mockcompute

import (
	"context"
	"fmt"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsPreload = &MockService{}

func (s *MockService) Preload(ctx context.Context) error {
	if err := s.preloadImages(ctx); err != nil {
		return fmt.Errorf("preloading images: %w", err)
	}
	return nil
}

func (s *MockService) preloadImages(ctx context.Context) error {
	for _, image := range images {
		name, err := s.parseImageSelfLink(image.GetSelfLink())
		if err != nil {
			return fmt.Errorf("invalid image self link %q: %w", image.GetSelfLink(), err)
		}

		fqn := name.String()
		if err := s.storage.Create(ctx, fqn, image); err != nil {
			return fmt.Errorf("preloading image %q: %v", fqn, err)
		}
	}
	return nil
}

var images = []*pb.Image{
	{
		Architecture:              PtrTo("X86_64"),
		CreationTimestamp:         PtrTo("2025-02-13T14:26:09.387-08:00"),
		Description:               PtrTo("Debian, Debian GNU/Linux, 12 (bookworm), amd64 built on 20250212"),
		EnableConfidentialCompute: PtrTo(false),
		Family:                    PtrTo("debian-12"),
		GuestOsFeatures: []*pb.GuestOsFeature{
			{Type: PtrTo("UEFI_COMPATIBLE")},
			{Type: PtrTo("VIRTIO_SCSI_MULTIQUEUE")},
			{Type: PtrTo("GVNIC")},
			{Type: PtrTo("SEV_CAPABLE")},
			{Type: PtrTo("SEV_LIVE_MIGRATABLE_V2")},
		},
		Id:               PtrTo[uint64](3433243310150635375),
		Kind:             PtrTo("compute#image"),
		LabelFingerprint: PtrTo("iNBmVNCFF9w="),
		Labels: map[string]string{
			"public-image": "true",
		},
		LicenseCodes: []int64{
			2147286739765738111,
		},
		Licenses: []string{
			"https://www.googleapis.com/compute/v1/projects/debian-cloud/global/licenses/debian-12-bookworm",
		},
		Name:     PtrTo("debian-12-bookworm-v20250212"),
		SelfLink: PtrTo("https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-12-bookworm-v20250212"),
		StorageLocations: []string{
			"eu",
			"asia",
			"us",
		},
		SourceType: PtrTo("RAW"),
		RawDisk: &pb.RawDisk{
			Source:        PtrTo(""),
			ContainerType: PtrTo("TAR"),
		},
		ArchiveSizeBytes: PtrTo(int64(2056498176)),
		DiskSizeGb:       PtrTo(int64(10)),
		Status:           PtrTo("READY"),
	},
}
