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

import unittest
from unittest.mock import MagicMock, AsyncMock, patch
import kubernetes
from src.server import MCPForGKEServer

class TestMCPForGKEServer(unittest.IsolatedAsyncioTestCase):

    async def asyncSetUp(self):
        self.mock_kubeconfig = "fake_kubeconfig"
        with patch('kubernetes.config.load_kube_config'), \
             patch('kubernetes.client.CoreV1Api'), \
             patch('kubernetes.client.AppsV1Api'), \
             patch('kubernetes.client.CustomObjectsApi'), \
             patch('kubernetes.client.ApiextensionsV1Api'):
            self.server = MCPForGKEServer(self.mock_kubeconfig)
            self.server.core_v1 = AsyncMock()
            self.server.apps_v1 = AsyncMock()
            self.server.custom_objects_api = AsyncMock()
            self.server.apiextensions_v1 = MagicMock()

    async def test_describe_custom_resource(self):
        # Mock CRD and Custom Object
        mock_crd = MagicMock()
        mock_crd.spec.names.kind = "MyCustomResource"
        mock_crd.spec.group = "stable.example.com"
        mock_version = MagicMock()
        mock_version.name = "v1"
        mock_crd.spec.versions = [mock_version]
        mock_crd.spec.names.plural = "mycustomresources"

        self.server.apiextensions_v1.list_custom_resource_definition.return_value = MagicMock(items=[mock_crd])

        mock_custom_object = {
            "apiVersion": "stable.example.com/v1",
            "kind": "MyCustomResource",
            "metadata": {"name": "test-cr", "namespace": "default"},
            "spec": {"key": "value"},
            "status": {
                "conditions": [
                    {"type": "Ready", "status": "True", "message": "Resource is ready", "reason": "Healthy"}
                ]
            }
        }
        self.server.custom_objects_api.get_namespaced_custom_object.return_value = AsyncMock(return_value=mock_custom_object)

        resource_kind = "MyCustomResource"
        resource_name = "test-cr"
        namespace = "default"

        result = await self.server.describe_resource(resource_kind, resource_name, namespace)

        self.server.apiextensions_v1.list_custom_resource_definition.assert_called_once()
        self.server.custom_objects_api.get_namespaced_custom_object.assert_called_once_with(
            group="stable.example.com",
            version="v1",
            name="test-cr",
            namespace="default",
            plural="mycustomresources"
        )
        self.assertIn("kind: MyCustomResource", result)
        self.assertIn("Resource Ready Status: True", result)
        self.assertIn("Message: Resource is ready", result)
        self.assertIn("Reason: Healthy", result)

    async def test_describe_custom_resource_not_found(self):
        self.server.apiextensions_v1.list_custom_resource_definition.return_value = MagicMock(items=[])

        resource_kind = "NonExistentCustomResource"
        resource_name = "test-cr"
        namespace = "default"

        result = await self.server.describe_resource(resource_kind, resource_name, namespace)
        self.assertIn(f"Could not find CRD for resource kind: {resource_kind}", result)

if __name__ == '__main__':
    unittest.main()
