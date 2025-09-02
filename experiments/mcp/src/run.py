# Copyright 2025 Google LLC
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

import argparse
import asyncio
import os
import subprocess
from server import MCPForGKEServer


def main():
    KUBE_CONFIG_DEFAULT_LOCATION = os.environ.get('KUBECONFIG', '~/.kube/config')
    MCPWorkDir = os.environ.get('MCPWorkDir', subprocess.check_output(['git', 'rev-parse', '--show-toplevel'], text=True).strip())

    print(f"Using MCPWorkDir: {MCPWorkDir}")
    server = MCPForGKEServer(kubeconfig=KUBE_CONFIG_DEFAULT_LOCATION, absDir=MCPWorkDir)
    server.run(transport='stdio')


if __name__ == "__main__":
    main()