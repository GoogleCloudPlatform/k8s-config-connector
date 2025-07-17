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

import os
import re
from mcp.server.fastmcp import FastMCP
from mcp.types import Tool, ToolAnnotations
from mcp.server.fastmcp.prompts import Prompt
from typing import Callable, List, Any
import kubernetes
import json
from google.genai import types as gtypes
from google import genai
import yaml
import subprocess
from pkg import fs
from pkg import openapi2jsonschema
from pkg import promotion

class MCPForGKEServer(FastMCP):
    def __init__(self, kubeconfig: str):
        super().__init__()
        
        self.criteria = self.load_criteria()
        api_criteria = self.criteria.get("api", {})
        
        # Use prompt instead dedicated code since it requires a LLM call.
        self.add_tool(
            name="create_resource",
            description=api_criteria.get("description", "Provide a prompt to give LLM instructions to create a CustomResource object for the given CustomResourceDefinition. It should be preferred as the LLM prompt instruction to generates a complete and valid YAML manifest for a Kubernetes Custom Resource. It's the best first step for creating any new KCC resource."),
            fn=self.create_resource,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="validate_resource",
            description="Validates a Kubernetes resource manifest against its schema. Use this to check for errors before applying a configuration to a cluster, especially when you don't have access to a live cluster for server-side validation.",
            fn=self.validate_resource,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="update_custom_resource",
            description="Applies a partial update to an existing Kubernetes Custom Resource in a live cluster. Use this to modify specific fields of a resource, such as changing a label, annotation, or spec field, without needing the full resource manifest.",
            fn=self.update_custom_resource,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="describe_resource",
            description="Fetches the complete, live state of a specific Kubernetes resource from the cluster. This is useful for inspecting the current status, annotations, and spec of a resource to diagnose issues or confirm its configuration.",
            fn=self.describe_resource,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        # This is a trick since GEMINI CLI does not support MCP prompts yet.
        # We make a special tool to introduce the prompt usage into the GEMINI plan. 
        self.add_tool(
            name="list_task_scenarios",
            description="Provides a list of high-level task scenarios or workflows. Use this to discover multi-step actions you can perform, such as creating a resource and then verifying its status.",
            fn=self.list_prompts,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="promote_api",
            description="Promotes a KCC API to a new version. This involves updating the API version and running validation.",
            fn=self.promote_api,
            annotations=ToolAnnotations(readOnlyHint=False, destructiveHint=True),
        )
        
        self.add_tool(
            name="promote_controller",
            description="Promotes a KCC controller to a new version. This involves updating the controller's API imports and running validation.",
            fn=self.promote_controller,
            annotations=ToolAnnotations(readOnlyHint=False, destructiveHint=True),
        )

        self.add_tool(
            name="promote_tests",
            description="Promotes KCC test fixtures to a new version. This involves updating the test fixtures and running validation.",
            fn=self.promote_tests,
            annotations=ToolAnnotations(readOnlyHint=False, destructiveHint=True),
        )

        self.add_tool(
            name="scenario_promote",
            description="Gives a prompt to help Gemini promote a KCC resource to a target version. The prompt explains that the promotion happens in 3 steps, promote_api, promote_controller and promote_tests. the gemini should use the corresponding mcp tools, and if the step fails, gemini should fix it based on the response error and the expected output. If the step succeeds, move on to the next step.",
            fn=self.scenario_promote,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_prompt(Prompt(
            name="scenario_create_resource",
            description="Guides the LLM to generate a complete and valid YAML manifest for a new Kubernetes Custom Resource. Use this when a user wants to create a resource and has provided a high-level description of their needs. The prompt will instruct the LLM to use the `create_resource` tool to get detailed instructions on how to generate the YAML, and the `validate_resource` tool to validate the generated YAML.",
            fn=self.scenario_custom_create_resource
        ))
        
        self.add_prompt(Prompt(
            name="scenario_resource_status",
            description="Initiates a workflow to check the readiness and health of an existing Kubernetes Custom Resource in a cluster. This is the best tool to use when a user wants to know the status of a resource or troubleshoot why it is not ready. The prompt will guide the LLM to use the `describe_resource` tool to get the resource's status and then provide suggestions for fixing it, potentially using the `update_custom_resource` tool.",
            fn=self.scenario_resource_status
        ))
    
        # Load Kubernetes configuration
        # A cluster context is preferred and enables more tasks. If not given, we will do our best (e.g. kubeconform vs kubectl --dry-run=server) 
        try:
            kubernetes.config.load_kube_config(kubeconfig)
            self.core_v1 = kubernetes.client.CoreV1Api()
            self.apps_v1 = kubernetes.client.AppsV1Api()
            self.custom_objects_api = kubernetes.client.CustomObjectsApi()
            self.apiextensions_v1 = kubernetes.client.ApiextensionsV1Api()
        except Exception as e:
            print(f"Error loading Kubernetes configuration: {e}")
            self.core_v1 = None
            self.apps_v1 = None

    def load_criteria(self):
        criteria = {}
        criteria_dir = "criteria"
        if not os.path.exists(criteria_dir):
            # Fallback to an absolute path
            # This is useful when the script is not run from the root of the project
            # TODO: make this more robust
            criteria_dir = os.path.join(os.getcwd(), "criteria")
        
        if os.path.exists(criteria_dir):
            for filename in os.listdir(criteria_dir):
                if filename.endswith(".json"):
                    topic = os.path.splitext(filename)[0]
                    filepath = os.path.join(criteria_dir, filename)
                    with open(filepath, "r") as f:
                        criteria[topic] = json.load(f)
        return criteria

    async def scenario_custom_create_resource(self, resource_kind: str, name: str, namespace: str, descriptive_requirements: dict) -> str:
        """Constructs a detailed prompt that guides the LLM to generate a valid Kubernetes Custom Resource manifest.

        This function translates a user's high-level request into a precise, multi-step plan for the LLM.
        The generated prompt instructs the LLM to use the `create_resource` tool to generate the initial YAML
        and then use the `validate_resource` tool to ensure the output is correct before presenting it to the user.
        
        Args:
            resource_kind: The `kind` of the Kubernetes resource to create (e.g., "StorageBucket", "SQLInstance").
            name: The `metadata.name` for the resource.
            namespace: The `metadata.namespace` where the resource will be created.
            descriptive_requirements: A dictionary mapping resource field paths to desired values or natural language descriptions.
                For example, `{"spec.versioning.enabled": True, "spec.lifecycleRule": "30-day deletion policy"}`.
        """
        # Prepare the instruction for Gemini
        instruction = f"""Create a Kubernetes YAML CustomResource (CR) of kind '{resource_kind}' with the name '{name}' in the namespace '{namespace}'. 
        The resource should fulfill the following requirements: {descriptive_requirements}. 
        
        Remember to include all necessary fields and configurations based on the kind of resource.
        You can find the CustomResourceDefinition (CRD) for this resource kind in the https://github.com/GoogleCloudPlatform/k8s-config-connector/tree/<master|version>/config/crds/resources, 
        the CRD for kind is in the form of apiextensions.k8s.io_v1_customresourcedefinition_<kind_in_plural_in_lower_case>.<group_in_lower_case>.cnrm.cloud.google.com.yaml.
        
        Hints:
        * Use create_resource tool to get the detailed prompt instructions on how to create the object. For example, `create_resource(crd_path="/path/to/your/crd.yaml", resource_configs=dict_from_descriptive_requirements)`.
        * If `kubectl` is not available, or there is no GKE cluster as `kubectl context`. Use validate_resource tool to validate the generated CR object. For example, `validate_resource(cr=cr_object, crd_path="/path/to/your/crd.yaml")`
        * Use update_custom_resource tool to update the CR with specific requirements. For example, `update_custom_resource(resource_kind="{resource_kind}", resource_name="{name}", namespace="{namespace}", requirements=dict_from_descriptive_requirements)`.
        * Use describe_resource tool to get detailed information about the CR after it is created. For example, `describe_resource(resource_kind="{resource_kind}", resource_name="{name}", namespace="{namespace}")`
        """
        return instruction

    async def scenario_resource_status(self, resource_kind: str, name: str, namespace: str) -> str:
        """Generates a prompt that outlines a workflow for checking and diagnosing a Kubernetes resource's status.

        This function creates instructions for the LLM to first use the `describe_resource` tool to fetch the
        live state of the resource from the cluster. Based on the output, the LLM is then guided to analyze
        the `status.conditions` field, report the resource's health, and suggest fixes for any issues.
        
        Args:
            resource_kind: The `kind` of the resource to check (e.g., "IAMPolicyMember", "PubSubSubscription").
            name: The `metadata.name` of the resource.
            namespace: The `metadata.namespace` where the resource is located.
        """
        # Prepare the instruction for Gemini
        instruction = f"""Check if the Kubernetes CustomResource (CR) of kind '{resource_kind}' with the name '{name}' in the namespace '{namespace}' is ready and healthy. 
        If it is not ready, provide detailed information on what is missing or what needs to be fixed. You can assume the CR is already created and exists in the cluster, and
        the CR status inherits the Kubernetes core v1, which contains the `status.conditions` field that indicates the readiness of the resource, `status.ready` field that indicates if the resource is ready, and 
        `status.reason` and `status.message` fields that provide more information about the resource status.
        
        Hints:
        * Use describe_resource tool to get detailed information about the CR. For example, `describe_resource(resource_kind="{resource_kind}", resource_name="{name}", namespace="{namespace}")`
        * If the CR is not ready, provide suggestions on how to fix it.
        * Use update_custom_resource tool to update the CR with specific requirements. For example, `update_custom_resource(resource_kind="{resource_kind}", resource_name="{name}", namespace="{namespace}", requirements=dict_from_descriptive_requirements)`.
        to fix the CR.
        * If the CR is ready, return a message indicating that the CR is ready and healthy 
        """
        return instruction

    async def scenario_promote(self, kind: str, targetVersion: str, service: str, apiPath: str, controllerPath: str, testFixturePath: str) -> str:
        """Constructs a detailed prompt that guides the LLM to promote a KCC resource to a new version.

        This function translates a user's high-level request into a precise, multi-step plan for the LLM.
        The generated prompt instructs the LLM to use the `promote_api`, `promote_controller`, and `promote_tests` tools sequentially.
        
        Args:
            kind: The kind of the resource to promote. For example: `APIQuotaAdjusterSettings`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
            service: The service name for the resource. For example: `cloudquota`.
            apiPath: The path to the API definition file. For example: `apis/cloudquota/v1alpha1/quotaadjustersettings_types.go`.
            controllerPath: The path to the controller file. For example: `pkg/controller/direct/cloudquota/quotaadjustersettings_controller.go`.
            testFixturePath: The path to the test fixture. For example: `pkg/test/resourcefixture/testdata/basic/cloudquota/v1alpha1/apiquotaadjustersettings`.
        """
        # Prepare the instruction for Gemini
        instruction = f"""Promote the KCC resource of kind '{kind}' to version '{targetVersion}'. This is a three-step process.

1.  **Promote the API**: Use the `promote_api` tool. For example: `promote_api(apiPath='{apiPath}', targetVersion='{targetVersion}')`. If the tool returns an error, analyze the error message and the expected output to fix the problem and retry. Once successful, proceed to the next step.

2.  **Promote the Controller**: Use the `promote_controller` tool. For example: `promote_controller(controllerPath='{controllerPath}', apiPath='{apiPath}', targetVersion='{targetVersion}')`. If the tool returns an error, analyze the error message and the expected output to fix the problem and retry. Once successful, proceed to the next step.

3.  **Promote the Tests**: Use the `promote_tests` tool. For example: `promote_tests(testFixturePath='{testFixturePath}', kind='{kind}', targetVersion='{targetVersion}')`. If the tool returns an error, analyze the error message and the expected output to fix the problem and retry.

If all three steps are successful, the promotion is complete.
"""
        return instruction

    async def describe_resource(self, resource_kind: str, resource_name: str, namespace: str) -> str:
        """Get detailed information about a specific Kubernetes CustomResource object.

        Args:
            resource_kind: The kind of the resource, e.g. "SQLInstance", "SpannerDatabase", "Pod", "Deployment", etc.
            resource_name: The name of the resource.
            namespace: The namespace where the resource is located.
        """
        # use "kubectl get crd | grep <resource_kind in lowe case> | awk '{print $1}'" to get the CRD name
        # use "kubectl get <resource_kind>.<crd_name> -n <namespace> <resource_name> -o yaml" to get the resource
        # Analyze the resource condition to see if it is ready and healthy.
        if not self.core_v1 or not self.apps_v1:
            return "Kubernetes API client not initialized. Please check the configuration."

        try:
            # Handle common Kubernetes resources
            if resource_kind.lower() == "pod" or resource_kind.lower() == "pods":
                api_response = self.core_v1.read_namespaced_pod(name=resource_name, namespace=namespace)
            elif resource_kind.lower() == "deployment" or resource_kind.lower() == "deployments":
                api_response = self.apps_v1.read_namespaced_deployment(name=resource_name, namespace=namespace)
            elif resource_kind.lower() == "service" or resource_kind.lower() == "services":
                api_response = self.core_v1.read_namespaced_service(name=resource_name, namespace=namespace)
            else:
                # Assume it's a Custom Resource
                # 1. Get CRD details to find group and version
                crd_list = self.apiextensions_v1.list_custom_resource_definition()
                crd_found = None
                for crd in crd_list.items:
                    if crd.spec.names.kind.lower() == resource_kind.lower():
                        crd_found = crd
                        break
                if not crd_found:
                    return f"Could not find CRD for resource kind: {resource_kind}"

                group = crd_found.spec.group
                version = crd_found.spec.versions[0].name # Assuming the first version is the one to use
                plural = crd_found.spec.names.plural

                # 2. Get Custom Resource details
                api_response = self.custom_objects_api.get_namespaced_custom_object(
                    group=group,
                    version=version,
                    name=resource_name,
                    namespace=namespace,
                    plural=plural
                )

            # Analyze the resource condition to see if it is ready and healthy.
            status_info = ""
            if "status" in api_response and "conditions" in api_response["status"]:
                for condition in api_response["status"]["conditions"]:
                    if condition.get("type") == "Ready":
                        status_info += f"\nResource Ready Status: {condition.get('status')}"
                        if condition.get("message"):
                            status_info += f"\nMessage: {condition.get('message')}"
                        if condition.get("reason"):
                            status_info += f"\nReason: {condition.get('reason')}"
                        break
            
            return yaml.dump(api_response, default_flow_style=False, indent=2) + status_info

        except kubernetes.client.ApiException as e:
            return f"Error describing resource: {e.body}"
        except subprocess.CalledProcessError as e:
            return f"Error executing kubectl command: {e.stderr}"
        except Exception as e:
            return f"An unexpected error occurred: {e}"
        
    async def create_resource(self, crd_content: str = None, crd_path: str = None, custom_configs: dict = None) -> str:
        """Generates a detailed prompt instructing an LLM to create a YAML manifest for a Custom Resource.

        This function takes a Custom Resource Definition (CRD) and optional user configurations and
        produces a precise prompt. The prompt directs the LLM to generate the YAML manifest.

        Either `crd_content` or `crd_path` must be provided.

        Args:
            crd_content: A string containing the full YAML content of the Custom Resource Definition.
            crd_path: The filesystem path to the CRD YAML file. Used if `crd_content` is not provided.
            custom_configs: A dictionary of configurations to apply to the resource. Keys are dot-separated
                paths (e.g., "spec.replicas") and values are the desired settings. If not provided, the
                LLM will generate a manifest with placeholders based on the CRD schema.
            
        Returns:
            A detailed string prompt for the LLM to generate a YAML manifest.
        """
        if not crd_content and not crd_path:
            raise ValueError("Please provide either crd_content or crd_path.")

        if not crd_content:
            # read CRD from path
            crd_content = fs.read_yaml_file(crd_path)
            if not crd_content:
                raise ValueError(f"Error reading CRD from path: {crd_path}. Please check the file path and format.")
        
        api_criteria = self.criteria.get("api", {})
            
        # ask Gemini to generate
        instruction = """You are a Kubernetes expert. Your task is to create a YAML file for a Custom Resource based on the provided Custom Resource Definition (CRD).

Here is the CRD:
{crd}

Here are the custom configurations to apply:
{custom_configs}

Please follow these instructions carefully:
{instructions}

Quality Criteria:
{quality_criteria}

Return only YAML: Your final output should be only the generated YAML content, with no extra text or explanations.
""".format(
    crd=crd_content,
    custom_configs=custom_configs,
    instructions="\n".join([f"- {item}" for item in api_criteria.get("instructions", [])]),
    examples="\n".join([f"- {item['scenario']}: {item['request']}" for item in api_criteria.get("examples", [])]),
    quality_criteria="\n".join([f"- {item}" for item in api_criteria.get("quality_criteria", [])]),
)
        return instruction
    
    async def update_custom_resource(self, resource_kind: str, resource_name: str, namespace: str, requirements: dict) -> str:
        """Update a Kubernetes CustomResource object with specific requirements.

        Args:
            resource_kind: The kind of the resource to update, e.g., "Deployment", "Service", etc.
            resource_name: The name of the resource.
            namespace: The namespace where the resource is located.
            requirements: A dictionary containing the requirements to update the resource.
            It could include fields with specific value or user requirements.
            For example, {"spec.replicas": 3, "spec.template.spec.containers[0].image": "nginx:latest"}.
            
        Returns:
            A string indicating the result of the update operation.
        """
        if not requirements:
            return "No requirements provided for updating the CustomResource."
        
        # Convert requirements to JSON patch format
        patch_data = json.dumps(requirements)
        patch_command = f"kubectl patch {resource_kind.lower()} {resource_name} -n {namespace} --type='merge' -p='{patch_data}'"
        
        try:
            result = subprocess.run(patch_command, shell=True, capture_output=True, text=True, check=True)
            return f"Successfully updated {resource_kind} '{resource_name}' in namespace '{namespace}'.\n{result.stdout}"
        except subprocess.CalledProcessError as e:
            return f"Error updating resource: {e.stderr}"
        
    async def validate_resource(self, cr: str, crd_path: str, timeout: int = 30) -> str:
        """Validate a CustomResource object if kubectl --dry-run=server is not available.

        Args:
            cr: The CustomResource object to validate.
            crd_path: The path to the CustomResourceDefinition.
            timeout: The timeout for the validation command in seconds.
        """
        if not cr:
            return "Error: No CustomResource provided for validation."
        if not crd_path:
            return "Error: No CustomResourceDefinition path provided for validation."
        
        try:
            # read CRD from path
            files = openapi2jsonschema.run(crd_path)
            for file in files:
                abspath = os.path.join(os.getcwd(), file)
                if not os.path.exists(abspath):
                    raise ValueError(f"Error: Generated schema file not found at '{abspath}'")
                
                return self.run_kubeconform_with_yaml_content(abspath, cr, timeout)
        except Exception as e:
            return f"Error during schema generation or validation: {e}"
    
    def run_kubeconform_with_yaml_content(self, schema_location: str, yaml_content: str, timeout: int = 30) -> str:
        """
        Runs the kubeconform command, passing the YAML content via stdin.

        Args:
            schema_location (str): The path to the schema JSON file.
            yaml_content (str): The content of the YAML file as a string.
            timeout (int): The timeout for the command in seconds.

        Returns:
            subprocess.CompletedProcess: The result of the subprocess execution.
        """
        command = ["kubeconform", "-summary", "-output", "json", "-schema-location", schema_location, '-']
        try:
            result = subprocess.run(
                command,
                input=yaml_content.encode('utf-8'),
                capture_output=True,
                check=False,
                timeout=timeout
            )
            if result.returncode != 0:
                try:
                    # Try to parse the output as JSON for detailed error info
                    error_data = json.loads(result.stdout)
                    return json.dumps(error_data, indent=2)
                except json.JSONDecodeError:
                    # Fallback to stderr if JSON parsing fails
                    return f"kubeconform failed with return code {result.returncode}:\n{result.stderr.decode()}"
            
            if result.stderr:
                # Even with a zero return code, there might be warnings in stderr
                return f"kubeconform validation passed with warnings:\n{result.stderr.decode()}"

            return result.stdout.decode()

        except FileNotFoundError:
            return "Error: 'kubeconform' command not found. Please ensure it is installed and in your system's PATH."
        except subprocess.TimeoutExpired:
            return f"Error: kubeconform command timed out after {timeout} seconds."
        except Exception as e:
            return f"An unexpected error occurred while running kubeconform: {e}"

    async def promote_api(self, apiPath: str, targetVersion: str) -> dict:
        """Promotes a KCC API to a new version.

        This function takes the path to an API definition file and a target version.
        It copies the API file to a new directory corresponding to the target version,
        updates the version information within the file, and then runs validation
        to ensure the promoted API is well-formed.

        Args:
            apiPath: The path to the API definition file. For example: `apis/cloudquota/v1alpha1/quotaadjustersettings_types.go`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
        """
        result = promotion.promote_api_file(apiPath, targetVersion)
        if "error" in result:
            return result
        
        validation_result = promotion.validate_promotion(apiPath, targetVersion)
        if "error" in validation_result:
            return validation_result
            
        return result

    async def promote_controller(self, controllerPath: str, apiPath: str, targetVersion: str) -> dict:
        """Promotes a KCC controller to a new version.

        This function updates the controller file to use the new API version.
        It modifies the import paths in the controller to point to the new API version
        and then runs a validation step to ensure the controller code still compiles.

        Args:
            controllerPath: The path to the controller file. For example: `pkg/controller/direct/cloudquota/quotaadjustersettings_controller.go`.
            apiPath: The path to the API definition file. For example: `apis/cloudquota/v1alpha1/quotaadjustersettings_types.go`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
        """
        go_module = "github.com/GoogleCloudPlatform/k8s-config-connector"
        result = promotion.promote_controller_file(controllerPath, apiPath, targetVersion, go_module)
        if "error" in result:
            return result

        validation_result = promotion.validate_controller_compilation(controllerPath)
        if "error" in validation_result:
            return validation_result

        return result

    async def promote_tests(self, testFixturePath: str, kind: str, targetVersion: str) -> dict:
        """Promotes KCC test fixtures to a new version.

        This function copies the test fixture to a new directory corresponding to the
        target version, updates the `apiVersion` in the test YAML files, and then
        runs validation to ensure the tests are still valid.

        Args:
            testFixturePath: The path to the test fixture. For example: `pkg/test/resourcefixture/testdata/basic/cloudquota/v1alpha1/apiquotaadjustersettings`.
            kind: The kind of the resource. For example: `APIQuotaAdjusterSettings`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
        """
        result = promotion.promote_test_fixture(testFixturePath, targetVersion)
        if "error" in result:
            return result

        validation_result = promotion.validate_test_fixture(kind)
        if "error" in validation_result:
            return validation_result

        return result
       
