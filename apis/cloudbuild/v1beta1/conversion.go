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

package v1beta1

import (
	"fmt"

	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
)

func Convert_WorkerPool_KRM_To_API_v1(in *CloudBuildWorkerPool, out *cloudbuildpb.WorkerPool) error {
	if in == nil {
		return nil
	}
	// CloudBuildWorkerPool API has "Name" as output only field.
	// The "Name" is of the form "projects/<projectID>/locations/<location>/workerpools/<workerpoolID>"
	// out.Name = in.Name
	out.DisplayName = in.Spec.DisplayName

	// Custom
	outConfig := &cloudbuildpb.PrivatePoolV1Config{}
	if err := Convert_PrivatePoolV1Config_KRM_To_API_v1(in.Spec.PrivatePoolConfig, outConfig); err != nil {
		return err
	}
	out.Config = &cloudbuildpb.WorkerPool_PrivatePoolV1Config{
		PrivatePoolV1Config: outConfig,
	}
	return nil
}

func Convert_PrivatePoolV1Config_KRM_To_API_v1(in *PrivatePoolV1Config, out *cloudbuildpb.PrivatePoolV1Config) error {
	if in == nil {
		return nil
	}
	networkconfig := &cloudbuildpb.PrivatePoolV1Config_NetworkConfig{}
	if err := Convert_PrivatePoolV1Config_NetworkConfig_KRM_To_API_v1(in.NetworkConfig, networkconfig); err != nil {
		return err
	}
	out.NetworkConfig = networkconfig

	workerconfig := &cloudbuildpb.PrivatePoolV1Config_WorkerConfig{}
	if err := Convert_PrivatePoolV1Config_WorkerConfig_KRM_To_API_v1(in.WorkerConfig, workerconfig); err != nil {
		return err
	}
	out.WorkerConfig = workerconfig
	return nil
}

func Convert_PrivatePoolV1Config_NetworkConfig_KRM_To_API_v1(in *NetworkConfig, out *cloudbuildpb.PrivatePoolV1Config_NetworkConfig) error {
	if in == nil {
		return nil
	}
	obj := in.DeepCopy()
	out.PeeredNetworkIpRange = obj.PeeredNetworkIpRange

	// custom
	switch obj.EgressOption {
	case "EGRESS_OPTION_UNSPECIFIED":
		out.EgressOption = 0
	case "NO_PUBLIC_EGRESS":
		out.EgressOption = 1
	case "PUBLIC_EGRESS":
		out.EgressOption = 2
	default:
		return fmt.Errorf("unknown egressoption %s", obj.EgressOption)
	}

	if obj.PeeredNetworkRef.External != "" {
		out.PeeredNetwork = obj.PeeredNetworkRef.External
	}
	return nil
}

func Convert_PrivatePoolV1Config_WorkerConfig_KRM_To_API_v1(in *WorkerConfig, out *cloudbuildpb.PrivatePoolV1Config_WorkerConfig) error {
	if in == nil {
		return nil
	}
	obj := in.DeepCopy()
	out.DiskSizeGb = int64(obj.DiskSizeGb)
	out.MachineType = obj.MachineType
	return nil
}
