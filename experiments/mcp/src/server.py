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
    def __init__(self, kubeconfig: str, absDir: str):
        super().__init__()
        
        self.absDir = absDir
        
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
            description="""Promotes a KCC API to a new version. This is a multi-step process that this tool automates. If this tool fails, you can attempt to perform the steps manually. The process involves:
1. Copying the entire source API package directory (e.g., `apis/myservice/v1alpha1`) to a new target directory (e.g., `apis/myservice/v1beta1`).
2. In the new target directory, for each `.go` file, it replaces all occurrences of the source version string (e.g., `v1alpha1`) with the target version string (e.g., `v1beta1`).
3. For the main `_types.go` file in the new target directory, it ensures two specific annotations are present on the primary Kind's struct definition:
    - It adds `// +kubebuilder:storageversion` to mark the new version as the storage version for the CRD.
    - It adds or appends to `// +kubebuilder:metadata:labels` the label `"internal.cloud.google.com/additional-versions={source_version}"` (e.g., `"internal.cloud.google.com/additional-versions=v1alpha1"`) to maintain a link to the previous version.
4. Finally, it runs a validation step to ensure the promotion was successful, which is equivalent to running `make generate-crds`.

This tool is not good for the case if there are other v1alpha1 resources in the APIs that depends on the target resource, or vice versa. In that case, you should use the `promote_api_prompt` tool to get a prompt to help Gemini promote the API in a more flexible way.
""",
            fn=self.promote_api,
            annotations=ToolAnnotations(readOnlyHint=False, destructiveHint=True),
        )
        
        self.add_tool(
            name="promote_api_prompt",
            description="Generates a detailed prompt to guide the user or LLM in promoting a KCC API, especially when dealing with dependencies between different API versions. Use this when `promote_api` fails due to compilation errors related to cross-version dependencies.",
            fn=self.promote_api_prompt,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="promote_api_validate",
            description="Validates the promotion of a KCC API to a new version. This tool can be run repeatedly without running promote_api. It runs the promote_validate_promotion and gives Gemini some instructions to run additional validation check itself.",
            fn=self.promote_api_validate,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="promote_controller",
            description="Promotes a KCC controller to a new version. This involves updating the controller's API imports and running validation. If the validation fails due to compilation errors related to package versions, you may need to manually adjust the imports. For example, if a type is still in `v1alpha1`, you may need to add `krmalpha 'github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1'` and change references from `krm.MyType` to `krmalpha.MyType`.",
            fn=self.promote_controller,
            annotations=ToolAnnotations(readOnlyHint=False, destructiveHint=True),
        )

        self.add_tool(
            name="promote_controller_validate",
            description="Validates the promotion of a KCC controller to a new version. This tool can be run repeatedly without running promote_controller.",
            fn=self.promote_controller_validate,
            annotations=ToolAnnotations(readOnlyHint=True),
        )

        self.add_tool(
            name="promote_controller_prompt",
            description="Generates a detailed prompt to guide the user or LLM in promoting a KCC controller, especially when dealing with dependencies between different API versions.",
            fn=self.promote_controller_prompt,
            annotations=ToolAnnotations(readOnlyHint=True),
        )

        self.add_tool(
            name="promote_tests",
            description="Promotes KCC test fixtures to a new version. This involves updating the test fixtures and running validation.",
            fn=self.promote_tests,
            annotations=ToolAnnotations(readOnlyHint=False, destructiveHint=True),
        )

        self.add_tool(
            name="scenario_promote",
            description="Gives a prompt to help Gemini promote a KCC resource to a target version. This tool reads from `experiments/promoter/results/candidates.json` to find the promotion candidate. The candidate's `apiCoverage` must be true. The prompt explains that the promotion happens in 3 steps, promote_api, promote_controller and promote_tests. the gemini should use the corresponding mcp tools, and if the step fails, gemini should fix it based on the response error and the expected output. If the step succeeds, move on to the next step. If the promote_api step fails with compilation errors, use the `promote_api_prompt` to get instructions on how to fix the dependencies.",
            fn=self.scenario_promote,
            annotations=ToolAnnotations(readOnlyHint=True),
        )

        self.add_tool(
            name="add_reference_doc",
            description="This function provides instructions to add reference documentation for a KCC resource. It accepts a KCC Kind name and the service it belongs to.",
            fn=self.add_reference_doc,
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
        try:
            criteria_dir = os.path.join(self.absDir, 'experiments/mcp/criteria')
        except (subprocess.CalledProcessError, FileNotFoundError):
            # Fallback if git is not available or not in a git repo
            script_dir = os.path.dirname(os.path.abspath(__file__))
            criteria_dir = os.path.join(script_dir, "..", "criteria")

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

    async def scenario_promote(self, kind: str, targetVersion: str) -> str:
        """Constructs a detailed prompt that guides the LLM to promote a KCC resource to a new version.

        This function finds the promotion candidate from the JSON content in `experiments/promoter/results/candidates.json`.
        If the candidate's `apiCoverage` is true, it constructs a multi-step plan for the LLM.
        The generated prompt instructs the LLM to use the `promote_api`, `promote_controller`, and `promote_tests` tools sequentially.
        If the `promote_api` step fails with compilation errors, it instructs the LLM to use `promote_api_prompt` to get instructions on how to fix the dependencies.
        
        Args:
            kind: The kind of the resource to promote. For example: `APIQuotaAdjusterSettings`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
        """
        candidates_json_path = os.path.join(self.absDir, 'experiments/promoter/results/candidates.json')
        with open(candidates_json_path, 'r') as f:
            candidates = json.load(f)
        
        candidate = next((c for c in candidates if c.get('kind') == kind), None)
        
        if not candidate:
            return f"Error: Could not find a promotion candidate for kind '{kind}'."
        
        if not candidate.get('apiCoverage'):
            return f"Error: The promotion candidate for kind '{kind}' is not ready to be promoted because 'apiCoverage' is not true."

        service = candidate.get('service')
        apiPath = candidate.get('apiPath')
        controllerPath = candidate.get('controllerPath')
        testFixturePath = candidate.get('testFixturePath')

        if not all([service, apiPath, controllerPath, testFixturePath]):
            return f"Error: The promotion candidate for kind '{kind}' is missing one or more required fields (service, apiPath, controllerPath, testFixturePath)."

        # Prepare the instruction for Gemini
        instruction = f"""Promote the KCC resource of kind '{kind}' to version '{targetVersion}'. This is a three-step process.

1.  **Promote the API**: Use the `promote_api` tool. For example: `promote_api(apiPath='{apiPath}', targetVersion='{targetVersion}')`. If the tool returns an error, analyze the error message. If the error is a compilation error, use the `promote_api_prompt` tool to get instructions on how to fix the dependencies. For example: `promote_api_prompt(apiPath='{apiPath}', targetVersion='{targetVersion}')`. After fixing the dependencies, retry the `promote_api` tool. Once successful, proceed to the next step.

2.  **Promote the Controller**: Use the `promote_controller` tool. For example: `promote_controller(controllerPath='{controllerPath}', apiPath='{apiPath}', targetVersion='{targetVersion}')`. If the tool returns an error, analyze the error message. If the error is a compilation error, use the `promote_controller_prompt` tool to get instructions on how to fix the dependencies. For example: `promote_controller_prompt(controllerPath='{controllerPath}', apiPath='{apiPath}', targetVersion='{targetVersion}')`. After fixing the dependencies, use `promote_controller_validate` to validate the controller promotion. Once successful, proceed to the next step.

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
                abspath = os.path.join(self.absDir, file)
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
        result = promotion.promote_api_file(apiPath, targetVersion, self.absDir)
        if "error" in result:
            return result
        
        validation_result = promotion.validate_promotion(apiPath, targetVersion, self.absDir)
        if "error" in validation_result:
            return validation_result
            
        return result

    async def promote_api_prompt(self, apiPath: str, targetVersion: str) -> str:
        """
        Generates a detailed prompt to guide the user or LLM in promoting a KCC API,
        especially when dealing with dependencies between different API versions.

        Args:
            apiPath: The path to the API definition file. For example: `apis/cloudquota/v1alpha1/quotaadjustersettings_types.go`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
        """
        abs_api_path = promotion.validate_api_path(apiPath, self.absDir)
        source_version = abs_api_path.split('/')[-2]
        instruction = f"""
You are about to promote the API at `{apiPath}` from `{source_version}` to `{targetVersion}`.

Promoting an API can be complex if the resource has dependencies on other resources, or if other resources depend on it. This is common in Kubernetes environments where resources reference each other.

**Instructions:**

1.  **Initial Promotion Attempt:**
    Start by running the `promote_api` tool:
    `promote_api(apiPath='{apiPath}', targetVersion='{targetVersion}')`

2.  **Analyze Compilation Errors:**
    If the `promote_api` tool fails during the validation step, it's likely due to compilation errors caused by version dependencies. Carefully examine the error messages. They will point to the files and types that need to be updated.

3.  **Fix Dependencies:**
    You will need to manually edit the Go files to resolve these dependencies. There are two common scenarios:

    **Scenario A: The promoted resource depends on an older version of another resource.**
    *   **Example:** Promoting `BackupDRBackupPlanAssociation` to `v1beta1`, which depends on `BackupPlanRef` from `v1alpha1`.
    *   **Symptom:** The compiler will complain that it cannot find the type for a field in the new `v1beta1` `_types.go` file.
    *   **Solution:**
        1.  Add a new import with an alias to the `v1alpha1` package in the `v1beta1` `_types.go` file.
            ```go
            import (
                backupdrv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
            )
            ```
        2.  Update the field to use the type from the aliased import.
            ```go
            // from:
            BackupPlanRef *BackupPlanRef `json:"backupPlanRef,omitempty"`
            // to:
            BackupPlanRef *backupdrv1alpha1.BackupPlanRef `json:"backupPlanRef,omitempty"`
            ```

    **Scenario B: An older version of a resource depends on the newly promoted resource.**
    *   **Example:** `ApigeeEnvironmentGroup` (`v1alpha1`) depends on `ApigeeOrganization`, which is being promoted to `v1beta1`.
    *   **Symptom:** The compiler will complain about a missing type in a `v1alpha1` file that references the type you are promoting.
    *   **Solution:**
        1.  Go to the `v1alpha1` `_types.go` file for the dependent resource (e.g., `environmentgroup_types.go`).
        2.  Add a new import with an alias to the new `v1beta1` package.
            ```go
            import (
                apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
            )
            ```
        3.  Update the field in the `v1alpha1` file to use the type from the new `v1beta1` aliased import.
            ```go
            // from:
            OrganizationRef *ApigeeOrganizationRef `json:"organizationRef"`
            // to:
            OrganizationRef *apigeev1beta1.ApigeeOrganizationRef `json:"organizationRef"`
            ```

4.  **Re-run Validation:**
    After you have applied the necessary code changes, you need to validate them. The easiest way is to re-run the original `promote_api` command. Since the files are already copied, it will likely just run the validation. If you want to only run validation, you might need a dedicated validation function if available.

By following these steps, you can handle complex API promotions with cross-version dependencies.
"""
        return instruction
        
    # async def promote_api_validate(self, apiPath: str, targetVersion: str) -> dict:
    async def promote_api_validate(self, apiPath: str, targetVersion: str) -> dict:
        """ Validates the promotion of a KCC API to a new version.

        Args:
            apiPath: The path to the API definition file. For example: `apis/cloudquota/v1alpha1/quotaadjustersettings_types.go`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
        """
        validation_result = promotion.validate_promotion(apiPath, targetVersion, self.absDir)
        if "error" in validation_result:
            return validation_result
        
        # TODO: Add more validation checks here.
        # For example, check if the new API version is correctly referenced in other parts of the codebase.
        # For example, check if the new API version is correctly referenced in the CRD.
        
        return {"message": "API promotion validation successful", "apiPath": apiPath, "targetVersion": targetVersion}


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
        abs_api_path = promotion.validate_api_path(apiPath, self.absDir)
        # Mixed versions case
        promotion.split_controller_imports(abs_api_path, controllerPath, targetVersion, self.absDir)
        result = {"new_controller_path": controllerPath}
        if "error" in result:
            return result

        return await self.promote_controller_validate(controllerPath, targetVersion)

    async def promote_controller_validate(self, controllerPath: str, targetVersion: str) -> dict:
        """Validates the promotion of a KCC controller to a new version.

        Args:
            controllerPath: The path to the controller file. For example: `pkg/controller/direct/cloudquota/quotaadjustersettings_controller.go`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
        """
        validation_result = promotion.validate_controller_compilation(controllerPath, self.absDir, targetVersion)
        if "error" in validation_result:
            return validation_result

        return {"message": "Controller promotion validation successful", "controllerPath": controllerPath}

    async def promote_controller_prompt(self, controllerPath: str, apiPath: str, targetVersion: str) -> str:
        """
        Generates a detailed prompt to guide the user or LLM in promoting a KCC controller,
        especially when dealing with dependencies between different API versions.

        Args:
            controllerPath: The path to the controller file. For example: `pkg/controller/direct/cloudquota/quotaadjustersettings_controller.go`.
            apiPath: The path to the API definition file. For example: `apis/cloudquota/v1alpha1/quotaadjustersettings_types.go`.
            targetVersion: The target version to promote to. For example: `v1beta1`.
        """
        abs_api_path = promotion.validate_api_path(apiPath, self.absDir)
        source_version = promotion.get_version_from_path(abs_api_path)
        instruction = f"""
You are about to promote the controller at `{controllerPath}` to support the API version `{targetVersion}`.

Promoting a controller is complex when the controller's directory contains files that need to reference types from both the new API version (`{targetVersion}`) and the old API version (`{source_version}`).

**Instructions:**

1.  **Initial Promotion Attempt:**
    Start by running the `promote_controller` tool:
    `promote_controller(controllerPath='{controllerPath}', apiPath='{apiPath}', targetVersion='{targetVersion}')`

2.  **Analyze Compilation Errors:**
    If the `promote_controller` tool fails, it's likely due to compilation errors. The compiler will complain about undefined types because the import aliases and type usages are incorrect.

3.  **Fix Import Aliases and Type Usages:**
    You need to manually edit the Go files in the controller directory (`{os.path.dirname(controllerPath)}`) to resolve these errors. The goal is to have two imports for the service's API:
    
    *   The **new version** (`{targetVersion}`) should be aliased to `krm`.
    *   The **old version** (`{source_version}`) should be aliased to a version-specific name, like `krm{source_version}`.

    **Example:**
    If you are promoting `ApigeeOrganization`, the imports in a controller file might need to look like this:

    ```go
    import (
        krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
        krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
    )
    ```

    Then, you must update all type usages in that file to use the correct alias:
    *   For types that were promoted (e.g., `ApigeeOrganization`), use the `krm` alias: `krm.ApigeeOrganization`.
    *   For types that remain in the old version (e.g., `ApigeeEnvironment`), use the version-specific alias: `krmv1alpha1.ApigeeEnvironment`.

4.  **Re-run Validation:**
    After you have manually corrected the import aliases and type usages in all necessary `.go` files in the directory, run the validation tool repeatedly until all compilation errors are resolved:
    `promote_controller_validate(controllerPath='{controllerPath}')`

By following these steps, you can correctly refactor the controller to handle mixed API versions.
"""
        return instruction

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
        result = promotion.promote_test_fixture(testFixturePath, targetVersion, self.absDir)
        if "error" in result:
            return result

        validation_result = promotion.validate_test_fixture(kind, self.absDir)
        if "error" in validation_result:
            return validation_result

        return result

    async def add_reference_doc(self, kind: str, service: str) -> str:
        """Provides instructions to add reference documentation for a KCC resource.

        Args:
            kind: The KCC Kind name (e.g., "StorageBucket").
            service: The service the kind belongs to (e.g., "storage").
        """
        if not kind or not service:
            return "Error: Both 'kind' and 'service' parameters are required."

        return f"""To add reference documentation for the kind '{kind}' under the service '{service}', follow these steps:

Step 1: Add samples for '{kind}' under '{service}' by following the instructions in '5.4 Add samples' from the document at 'docs/develop-resources/deep-dives/5-releases.md'.

Step 2: Once step 1 is complete, follow the instructions in sections 5.2, 5.3, and 5.5 of the same document ('docs/develop-resources/deep-dives/5-releases.md') to add the Google Docs for '{kind}'.

Step 3: After completing the previous steps, run 'make resource-docs' to generate the required code.
"""
