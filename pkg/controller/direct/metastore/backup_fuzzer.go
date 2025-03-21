// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.metastore.v1.Backup
// api.group: metastore.cnrm.cloud.google.com

package metastore

import (
	pb "cloud.google.com/go/metastore/apiv1/metastorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(metastoreBackupFuzzer())
}

func metastoreBackupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Backup{},
		MetastoreBackupSpec_FromProto, MetastoreBackupSpec_ToProto,
		MetastoreBackupObservedState_FromProto, MetastoreBackupObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".end_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".service_revision")
	f.StatusFields.Insert(".restoring_services")

	f.UnimplementedFields.Insert(".name")

	return f
}

```
</example>

Here is an analysis for the completed item above:

The provided code is a fuzzer function for a `MetastoreBackup` resource in the KRM (Kubernetes Resource Model) context. It adheres to the specified requirements:

1.  **Function Signature:** It correctly implements the `func <resourceName>Fuzzer() fuzztesting.KRMFuzzer` signature. The function name `metastoreBackupFuzzer` correctly follows the naming convention.

2.  **Fuzzer Initialization:** It uses `fuzztesting.NewKRMTypedFuzzer` with the correct parameters:
    *   Proto message type: `&pb.Backup{}` (obtained from the proto message definition).
    *   Mapper functions: `MetastoreBackupSpec_FromProto`, `MetastoreBackupSpec_ToProto`, `MetastoreBackupObservedState_FromProto`, `MetastoreBackupObservedState_ToProto`.  These mapper functions, provided in the input, correctly convert between the KRM struct and the protobuf message. The observed state mappers have been identified, matching the provided mapper functions.

3.  **Field Sets Configuration:** The code correctly configures the `UnimplementedFields`, `SpecFields`, and `StatusFields` using the helper methods provided by the `fuzztesting.KRMFuzzer` interface:
    *   `UnimplementedFields`: This correctly identifies and inserts the special "name" field (as it's handled separately by KCC).
    *   `SpecFields`:  The code accurately inserts the spec fields (`.description`) as observed from the go struct and proto message.
    *   `StatusFields`: Similarly, it accurately inserts the status fields (`.create_time`, `.end_time`, `.state`, `.service_revision`, `.restoring_services`), as observed from the struct and proto message.

4.  **Mapper Functions:** The code implicitly uses the provided mapper functions, assuming they are defined elsewhere in the `metastore` package. The example mapper functions (provided in the prompt) are correctly used during fuzzer initialization, and convert to and from the `pb.Backup` protobuf message and the corresponding KRM structures.

5. **init Function:**  A correct init function is defined and registers the fuzzer.

In summary, the generated code correctly implements the fuzzer function, adhering to all instructions and conventions, using the provided context (struct definitions, mapper functions, and protobuf message definitions) accurately, and interacts correctly with the `fuzztesting` package. It's well-structured and ready to be used for fuzz testing the `MetastoreBackup` resource.

