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

lines = []
with open('pkg/controller/direct/networkmanagement/connectivitytest_fuzzer.go', 'r') as f:
    for line in f:
        lines.append(line)

# Let's just rewrite the end of the file properly
import re
unique_lines = []
seen = set()
for line in lines:
    if "f.Unimplemented_NotYetTriaged" in line:
        if line not in seen:
            seen.add(line)
            unique_lines.append(line)
    else:
        unique_lines.append(line)

with open('pkg/controller/direct/networkmanagement/connectivitytest_fuzzer.go', 'w') as f:
    f.writelines(unique_lines)
