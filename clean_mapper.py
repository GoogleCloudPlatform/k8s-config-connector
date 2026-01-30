# Copyright 2026 Google LLC
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

import re

with open('pkg/controller/direct/orgpolicy/mapper.generated.go', 'r') as f:
    content = f.read()

# Remove v1alpha1 import
content = re.sub(r'\t.*"github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"\n', '', content)

parts = content.split('func ')
new_parts = [parts[0]] # Header and imports
seen_funcs = set()

for part in parts[1:]:
    # Extract function name
    # func Name(...) ...
    # The part starts with Name(...)
    func_name = part.split('(')[0]

    if 'krmorgpolicyv1alpha1' in part:
        continue

    if func_name in seen_funcs:
        continue

    seen_funcs.add(func_name)
    new_parts.append(part)

new_content = 'func '.join(new_parts)

with open('pkg/controller/direct/orgpolicy/mapper.generated.go', 'w') as f:
    f.write(new_content)
