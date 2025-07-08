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

class MCPForGKEServer(FastMCP):
    def __init__(self, kubeconfig: str):
        super().__init__()
        
        # Use prompt instead dedicated code since it requires a LLM call.
        self.add_tool(
            name="create_resource",
            description="Create a Kubernetes CustomResource object based on the provided CustomResourceDefinition (CRD) file.",
            fn=self.create_resource,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="validate_resource",
            description="Validate a Kubernetes CustomResource object against its CustomResourceDefinition (CRD) using kubeconform." +
            "If tool helps validating CR if a GKE cluster or kubectl is not available. Otherwise, the kubectl server side validation will be preferred.",
            fn=self.validate_resource,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="update_custom_resource",
            description="""Update a Kubernetes CustomResource object with specific requirements.
            "This tool expects the config change to be a nested dictionary representing a JSON merge patch.
            "For example, {"spec": {"replicas": 3}, "metadata": {"labels": {"app": "my-app"}}}.
            "The tool will update the CustomResource object in the cluster with the provided requirements. 
            "If the update fails, it will return an error message with details.
            
            One common use is to change the projectID from "default" to the $(gcloud config get core/project) in "cnrm.cloud.google.com/project-id" annotation. 
            This is a common error that the "default" namespace is used unindentedly as GCP project ID in annotations.
            """,
            fn=self.update_custom_resource,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_tool(
            name="describe_resource",
            description="Get detailed information about a specific Kubernetes resource.",
            fn=self.describe_resource,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        # This is a trick since GEMINI CLI does not support MCP prompts yet.
        # We make a special tool to introduce the prompt usage into the GEMINI plan. 
        self.add_tool(
            name="list_scenarios",
            description="Give a list of available prompts for Kubernetes related tasks." +
            "Each prompt is a specific scenario that is designed to help LLM understand the tasks."
            "Using prompt should be preferred before using tools, because the prompt can help determining which tools to use and the order of using them.",
            fn=self.list_prompts,
            annotations=ToolAnnotations(readOnlyHint=True),
        )
        
        self.add_prompt(Prompt(
            name="scenario_custom_create_resource",
            description="A prompt to create a Kubernetes CustomResource object based on user requirements.",
            fn=self.scenario_custom_create_resource
        ))
        
        self.add_prompt(Prompt(
            name="scenario_resource_status",
            description="A prompt to describe a Kubernetes CustomResource object in the GKE cluster. It specifically checks if the CR is in a healthy status.",
            fn=self.scenario_resource_status
        ))
    
        # Load Kubernetes configuration
        # A cluster context is preferred and enables more tasks. If not given, we will do our best (e.g. kubeconform vs kubectl --dry-run=server) 
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

    async def scenario_custom_create_resource(self, resource_kind: str, name: str, namespace: str, descriptive_requirements: dict) -> str:
        """ Generate a prompt for Gemini to create a Kubernetes CustomResource (CR) based on user requirements.
        
        Args:
            resource_kind: The kind of the resource to generate, e.g., "Deployment", "Service", etc.
            name: The name of the resource.
            namespace: The namespace where the resource will be created.
            descriptive_requirements: A dictionary containing user descriptive requirements for the resource.
            It could include fields with specific value or user requirements.
            For example, {"replica": 3, "image": "use latest ubuntu"}. 
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
        """ Generate a prompt for Gemini to check and (if possible) make a Kubernetes CustomResource (CR) ready and healthy.
        
        Args:
            resource_kind: The kind of the resource to check, e.g., "Deployment", "Service", etc.
            name: The name of the resource.
            namespace: The namespace where the resource is located.
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
        
    async def create_resource(self, crd_content: str|None, crd_path: str|None, custom_configs: dict) -> str:
        """Provide a prompt to give LLM instructions to create a CustomResource object for the given CustomResourceDefinition.

        Args:
            crd_content: The content of the CustomResourceDefinition. If not given, use crd_path  
            crd_path: Is the path to the CustomResourceDefinition. Required if crd_content is not given. 
            custom_configs: A dictionary containing the configurations to fill in the CustomResource object.
            It could include fields with specific value or user requirements.
            For example, {"spec.replicas": 3, "spec.template.spec.containers[0].image": "nginx:latest"}.
            If not provided, the tool will try to fill in as many fields as possible based on the CRD definition.
            
        Returns:
            The prompt message.
        """
        if not crd_content:
            # read CRD from path
            crd_content = fs.read_yaml_file(crd_path)
            if not crd_content:
                raise ValueError(f"Error reading CRD from path: {crd_path}. Please check the file path and format.")
            
        # ask Gemini to generate
        instruction = r"""Create a Kubernetes YAML CustomResource using the given CustomResourceDefinition (CRD) and fill in the content. Remember that (described in JSON path), 
    - If given, fill in the content as {custom_configs}.
    - the `.metadata.name` should be a DNS subdomain name as defined in RFC 1123.
    - try to understand the meaning of each field from the CRD definition.
    - try to configure as many fields as possible. 
    - the CRD is {crd}
    
    Remember you should only return the YAML formatted CR. 
    # """.format(crd=crd_content,custom_configs=custom_configs)
    #     print(instruction)
    #     messages = [gtypes.Content(
    #             role='user',
    #             parts=[gtypes.Part.from_text(text=instruction)]
    #     )]
        
        
    #     llm_response = self.google_client.models.generate_content(
    #         model="gemini-2.0-flash-001",
    #         contents=messages,
    #     )
        
    #     # Gemini
    #     if hasattr(llm_response, 'candidates'):
    #         pretty_yaml = ""
    #         for candidate in llm_response.candidates:
    #             if candidate.content.parts[0] != None:
    #                 cr_result = candidate.content.parts[0].text
    #                 pretty_yaml = yaml.dump(cr_result, default_flow_style=False, indent=2, sort_keys=False)
    #         return pretty_yaml
    #     if hasattr(llm_response, 'content'):
    #         if llm_response.content.parts[0] != None:
    #             cr_result = llm_response.content.parts[0].text
    #             pretty_yaml = yaml.dump(cr_result, default_flow_style=False, indent=2, sort_keys=False)
    #             return pretty_yaml
    #     raise ValueError("Error generating CustomResource object. Please check the input and try again.")
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
        # Use kubectl patch command to update the resource
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
        
    async def validate_resource(self, cr: str, crd_path: str) -> str:
        """Validate a CustomResource object if kubectl --dry-run=server is not available.

        Args:
            cr: The CustomResource object to validate.
            crd_path: The path to the CustomResourceDefinition.
        """
        if not cr:
            return "Error: No CustomResource provided for validation."
        if not crd_path:
            return "Error: No CustomResourceDefinition path provided for validation."
        # read CRD from path
        files = openapi2jsonschema.run(crd_path)
        for file in files:
            abspath = os.path.join(os.getcwd(), file)
            if not os.path.exists(abspath):
                raise(f"Error: File not found at '{abspath}'")
            
            return self.run_kubeconform_with_yaml_content(abspath, cr)
    
    def run_kubeconform_with_yaml_content(self, schema_location: str, yaml_content: str) -> str:
        """
        Runs the kubeconform command, passing the YAML content via stdin.

        Args:
            schema_location (str): The path to the schema JSON file.
            yaml_content (str): The content of the YAML file as a string.

        Returns:
            subprocess.CompletedProcess: The result of the subprocess execution.
        """
        # TODO let gemini decide how to call kubeconform
        command = ["kubeconform", "-summary", "-output", "json", "-schema-location", schema_location, '-']
        try:
            # Use subprocess.run to execute the command.
            # input: The string to be passed to the stdin of the child process.
            # text=True: Decodes stdin, stdout, and stderr using default encoding (usually UTF-8).
            # capture_output=True: Captures stdout and stderr.
            result = subprocess.run(
                command,
                input=yaml_content.encode('utf-8'),  # Encode the YAML content to bytes
                capture_output=True,
                check=False  # Raise a CalledProcessError if the command returns a non-zero exit code
            )
            print("kubeconform output (stdout):\n", result)
            if result.stderr:
                raise("kubeconform errors (stderr):\n", result.stderr)
            return result.stdout
        except subprocess.CalledProcessError as e:
            raise ValueError(f"Error running kubeconform: {e}\n")
        except FileNotFoundError:
            raise FileNotFoundError("Error: 'kubeconform' command not found. Make sure it's installed and in your PATH.\n")
        except Exception as e:
            raise Exception(f"An unexpected error occurred: {e}\n")
    