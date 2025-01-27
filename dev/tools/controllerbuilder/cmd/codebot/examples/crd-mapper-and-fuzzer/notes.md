1. In apis/filestore/ create generate.sh

```
#!/bin/bash
# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder


go run . generate-types \
    --service google.cloud.filestore.v1 \
    --api-version filestore.cnrm.cloud.google.com/v1alpha1 \
    --resource FilestoreInstance:Instance

go run . generate-mapper \
    --service google.cloud.filestore.v1 \
    --api-version filestore.cnrm.cloud.google.com/v1alpha1
```

1. Run:

```
chmod +x apis/filestore/generate.sh
apis/filestore/generate.sh
```


1. Run `dev/tasks/generate-crds`


1. Make git commits:

```
git add apis/filestore/generate.sh
git commit -m "filestore: add crd generation script"
```

```
git add apis/filestore/v1alpha1/
git add pkg/controller/direct/filestore/
git add config/crds/resources/
git commit -m "filestore: autogenerate types and CRDs"
```

1. Update Spec and Status CRDs:

TODO: Need to exclude the _existing_ CRD for this to work well
TODO: Merging is a pain.  Maybe split the Spec and ObservedState generation?


```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF
// +kcc:proto=google.cloud.filestore.v1.Instance
EOF
```


TODO: Often need to add projectRef / location.  Better input data?

```
import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

/* Immutable. The zone for this filestore instance. */
	Location string `json:"location"`

	/* The project that this resource belongs to. */
	ProjectRef *refs.ProjectRef `json:"projectRef"`
```

1. Iterate until `dev/tasks/generate-crds; go build ./pkg/controller/direct/filestore/...` compiles.

You may need to create some manual mappings.  TODO: Can we reduce these cases?

You may need to mark some fields as `// NOTYET`.

Some Refs may not be identified correctly, they might be `// NOTYET`

If you need to redefine a nested field, you can create a file like `v1alpha1/common_types.go` and simply paste the struct.  Then regenerate with `apis/filestore/generate.sh`.  The generator code should see the type with the same `// +kcc:proto=...` annotation and skip generation.

1. Make git commits:

```
git add apis/filestore/v1alpha1/instance_types.go
git add apis/filestore/v1alpha1/common_types.go # Maybe
git commit -m "filestore: update spec & status definitions"
```

```
git add apis/filestore/v1alpha1/
git add pkg/controller/direct/filestore/
git add config/crds/resources/
git commit -m "filestore: autogenerate types and CRDs"
```

1. Create instance_fuzzer.go

```
controllerbuilder prompt --src-dir ~/kcc/k8s-config-connector --proto-dir ~/kcc/k8s-config-connector/.build/third_party/googleapis/ <<EOF > pkg/controller/direct/filestore/instance_fuzzer.go
// +tool:crd-fuzzer
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Instance
// crd.type: FilestoreInstance
EOF
```

1. Add e.g. `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/filestore"` to imports in pkg/controller/direct/register/register.go

1. Run `dev/ci/presubmits/test-mappers-roundtrip` and iterate until we pass tests.  You will need to update the fields in the _fuzzer.go.  TODO: Does this get better once we have more examples of crd-fuzzer?

1. Commit the fuzzer

```
git add pkg/controller/direct/filestore/instance_fuzzer.go
git add pkg/controller/direct/register/register.go
git commit -m "FilestoreInstance: create fuzzer"
```

