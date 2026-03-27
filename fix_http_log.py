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

import os
files = [
    "pkg/test/resourcefixture/testdata/basic/pubsub/v1beta1/pubsubsubscription/pubsubsubscription-cloudstorage/_http.log",
    "pkg/test/resourcefixture/testdata/basic/run/v1beta1/runjob/runjob-gcsvolume/_http.log"
]

to_replace = """{
    "kind": "storage#objects",
    "prefixes": [
      "testfolder",
      "testmanagedfolder"
    ]
  }"""
to_replace_2 = """{
    "kind": "storage#objects",
    "prefixes": [
      "testfolder",
      "testmanagedfolder"
    ]
}"""

replacement = """{
    "kind": "storage#objects"
  }"""
replacement_2 = """{
    "kind": "storage#objects"
}"""

for file_path in files:
    with open(file_path, "r") as f:
        content = f.read()
    
    import re
    # We can just use a regex to replace "prefixes": [\n ... \n]
    content = re.sub(r',\s*"prefixes": \[\s*"testfolder",\s*"testmanagedfolder"\s*\]', '', content)
    
    with open(file_path, "w") as f:
        f.write(content)

