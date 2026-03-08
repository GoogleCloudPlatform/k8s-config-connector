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

package resourceoverrides

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func GetContainerNodePoolResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ContainerNodePool",
	}
	ro.Overrides = append(ro.Overrides, addConfidentialInstanceTypeEnumForNodePool())
	return ro
}

func addConfidentialInstanceTypeEnumForNodePool() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		spec := schema.Properties["spec"]

		// confidentialNodes in nodeConfig
		if nc, ok := spec.Properties["nodeConfig"]; ok {
			if cn, ok := nc.Properties["confidentialNodes"]; ok {
				if cit, ok := cn.Properties["confidentialInstanceType"]; ok {
					cit.Enum = []apiextensions.JSON{
						{Raw: []byte(`"SEV"`)},
						{Raw: []byte(`"SEV_SNP"`)},
					}
					cn.Properties["confidentialInstanceType"] = cit
					nc.Properties["confidentialNodes"] = cn
					spec.Properties["nodeConfig"] = nc
				}
			}
		}
		schema.Properties["spec"] = spec
		return nil
	}
	return o
}
