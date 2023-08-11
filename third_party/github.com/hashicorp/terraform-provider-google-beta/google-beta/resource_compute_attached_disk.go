// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	tpgcompute "github.com/hashicorp/terraform-provider-google-beta/google-beta/services/compute"
	compute "google.golang.org/api/compute/v0.beta"
)

func findDiskByName(disks []*compute.AttachedDisk, id string) *compute.AttachedDisk {
	return tpgcompute.FindDiskByName(disks, id)
}
