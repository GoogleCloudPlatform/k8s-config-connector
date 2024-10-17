// Copyright 2024 Google LLC
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

package sql

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func ApplySQLInstanceGCPDefaults(in *krm.SQLInstance, out *api.DatabaseInstance) {
	if in.Spec.Settings.ActivationPolicy == nil {
		out.Settings.ActivationPolicy = "ALWAYS"
	}
	if in.Spec.Settings.AvailabilityType == nil {
		out.Settings.AvailabilityType = "ZONAL"
	}
	if in.Spec.Settings.DiskType == nil {
		out.Settings.DataDiskType = "PD_SSD"
	}
	if in.Spec.Settings.Edition == nil {
		out.Settings.Edition = "ENTERPRISE"
	}
	if in.Spec.Settings.PricingPlan == nil {
		out.Settings.PricingPlan = "PER_USE"
	}
	if in.Spec.Settings.DiskAutoresize == nil {
		out.Settings.StorageAutoResize = direct.PtrTo(true)
	}
}
