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

package controllerprofile

import (
	"fmt"
	"testing"
)

func TestRenderSecureSourceManagerInstance(t *testing.T) {
	profile := &ControllerProfile{
		KCCService:    "securesourcemanager",
		KCCVersion:    "v1beta1",
		Kind:          "SecureSourceManagerInstance",
		ProtoResource: "Instance",
		Identity: IdentityTrait{
			Strategy:         IdentityClientField,
			TypeName:         "krm.SecureSourceManagerInstanceIdentity",
			UseGetIdentity:   true,
			HasFromExternal:  true,
			ParentStringExpr: "ParentString()",
			IDExpr:           "Instance",
			HasProjectRef:    true,
			HasLocation:      true,
			HasResourceID:    true,
		},
		Lifecycle: LifecycleTrait{
			HasGet:            true,
			GetMethodName:     "GetInstance",
			GetRequestType:    "GetInstanceRequest",
			GetNameField:      "Name",
			CanCreate:         true,
			CreateMethodName:  "CreateInstance",
			CreateRequestType: "CreateInstanceRequest",
			CreateParentField: "Parent",
			CreateIDField:     "InstanceId",
			CreateBodyField:   "Instance",
			UpdateStyle:       UpdateStyleUnsupported,
			DeleteStyle:       DeleteStyleStandard,
			DeleteMethodName:  "DeleteInstance",
			DeleteRequestType: "DeleteInstanceRequest",
			DeleteNameField:   "Name",
		},
		Execution: ExecutionTrait{
			Create: OpExecution{IsLRO: true, IsEmpty: false},
			Delete: OpExecution{IsLRO: true, IsEmpty: true},
		},
		Transport: TransportTrait{
			GCPClientImportPath:  "cloud.google.com/go/securesourcemanager/apiv1",
			GCPClientType:        "*gcp.Client",
			GCPClientConstructor: "gcp.NewRESTClient",
			PBImportPath:         "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb",
			PBPackageAlias:       "pb",
			HasReferences:        true,
			HasObservedState:     true,
			LabelDiffing:         LabelsNone,
			HasLabels:           true,
		},
	}

	code, err := RenderController(profile)
	if err != nil {
		t.Fatalf("RenderController failed: %v", err)
	}
	fmt.Println("=== GENERATED CODE ===")
	fmt.Println(string(code))
}
