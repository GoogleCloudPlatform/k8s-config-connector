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
import shutil
import re
import subprocess

GO_MODULE = "github.com/GoogleCloudPlatform/k8s-config-connector"

def to_abs_path(path: str, base_dir: str = None) -> str:
    """Converts a relative path to an absolute path based on the provided base directory."""
    if os.path.isabs(path) or base_dir is None:
        return path
    return os.path.join(base_dir, path)

def validate_api_path(api_path: str, base_dir: str = None) -> str:
    if base_dir is not None:
        api_path = to_abs_path(api_path, base_dir)
    if not api_path:
        raise ValueError(f"The API path '{api_path}' does not exist in the repository.")
    if not os.path.isfile(api_path) or not api_path.endswith('_types.go'):
        raise ValueError(f"The API path '{api_path}' must point to a Go file ending with '_types.go'.")
 
    return api_path

def validate_controller_path(controller_path: str, base_dir: str = None) -> str:
    if base_dir is not None:
        controller_path = to_abs_path(controller_path, base_dir)
    if not controller_path:
        raise ValueError(f"The API path '{controller_path}' does not exist in the repository.")
    if not os.path.isfile(controller_path) or not controller_path.endswith('_controller.go'):
        raise ValueError(f"The API path '{controller_path}' must point to a Go file ending with '_controller.go'.")
    return controller_path

def validate_fixture_path(fixture_path: str, base_dir: str = None) -> str:
    if base_dir is not None:
        fixture_path = to_abs_path(fixture_path, base_dir)
    if not fixture_path:
        raise ValueError(f"The fixture test path '{fixture_path}' does not exist in the repository.")
    if not os.path.isdir(fixture_path):
        raise ValueError(f"The fixture test path '{fixture_path}' must point to a directory")
 
    return fixture_path

def get_version_from_path(path: str) -> str:
    """
    Extracts the version from a file path.
    Assumes the version is in the format v[a-z0-9]+ (e.g., v1alpha1, v1beta1).
    """
    path = validate_api_path(path)
    match = re.search(r'v[0-9]+((alpha|beta)[0-9]+)?', path)
    if match:
        return match.group(0)
    raise ValueError(f"Could not find version in path: {path}")

def get_kind_from_path(path: str) -> str:
    """
    Extracts the kind from a file path.
    Assumes the kind is in the format xxx_types.go.
    """
    path = validate_api_path(path)
    with open(path, 'r') as f:
        content = f.read()
    match = re.search(r'var\s+([A-Z][a-zA-Z0-9]*)GVK\s*=\s*GroupVersion\.WithKind\("([A-Z][a-zA-Z0-9]*)"\)', content)
    if match and match.group(1) == match.group(2):
        return match.group(1)
    raise ValueError(f"Could not find kind in path: {path}")

def get_types_from_api_dir(api_dir_path: str, base_dir: str) -> set:
    """Parses all .go files in a directory and returns a set of defined type names."""
    types = set()
    abs_path = to_abs_path(api_dir_path, base_dir)
    if not os.path.isdir(abs_path):
        return types
    for filename in os.listdir(abs_path):
        if filename.endswith('.go'):
            filepath = os.path.join(abs_path, filename)
            with open(filepath, 'r') as f:
                content = f.read()
            # Regex to find 'type MyType struct'
            found_types = re.findall(r'type\s+([A-Z][a-zA-Z0-9_]+)\s+struct', content)
            types.update(found_types)
    return types


def get_controller_info(controller_path: str) -> (str, str):
    controller_path = validate_controller_path(controller_path)
    with open(controller_path, 'r') as f:
        content = f.read()
    version_match = re.search(r'// crd.version: (v[a-z0-9]+)', content)
    kind_match = re.search(r'// crd.type: (.*)', content)
    
    version = version_match.group(1).strip() if version_match else ""
    kind = kind_match.group(1).strip() if kind_match else ""
    return version, kind

def split_controller_imports(api_dir: str, controller_dir: str, target_version: str, base_dir: str):
    api_dir = validate_api_path(api_dir, base_dir)
    # kind = get_kind_from_path(api_dir)  # Ensure the API path is valid and has a kind
    source_version = get_version_from_path(api_dir)
    service = api_dir.split('/')[-3]
    promoted_kind = get_kind_from_path(api_dir)
    
    abs_controller_path = validate_controller_path(controller_dir, base_dir)
    abs_controller_dir = os.path.dirname(abs_controller_path)
    
    # Get all types from the source API directory
    source_api_dir = f'apis/{service}/{source_version}'
    all_source_types = get_types_from_api_dir(source_api_dir, base_dir)

    # All types related to the promoted kind are considered promoted
    promoted_types = {t for t in all_source_types if t.startswith(promoted_kind)}
    remaining_source_types = all_source_types - promoted_types

    # Refactor each controller file in the directory
    all_go_files = [f for f in os.listdir(abs_controller_dir) if f.endswith('.go')]
    for filename in all_go_files:
        filepath = os.path.join(abs_controller_dir, filename)
        with open(filepath, 'r') as f:
            content = f.read()

        original_import_path = f'"{GO_MODULE}/apis/{service}/{source_version}"'
        original_aliased_import = f'krm {original_import_path}'

        if original_aliased_import not in content:
            continue

        # Find all types used with the `krm` alias in this specific file
        used_krm_types = set(re.findall(r'\bkrm\.([A-Z][a-zA-Z0-9_]+)\b', content))

        # Check which categories of types are used
        uses_promoted = bool(used_krm_types.intersection(promoted_types))
        uses_remaining = bool(used_krm_types.intersection(remaining_source_types))

        new_content = content
        
        # 1. Replace usages of remaining types with a new alias
        if uses_remaining:
            new_alias_for_source = f'krm{source_version}'
            for type_name in remaining_source_types:
                # Important: only replace if it was actually a krm. type
                if type_name in used_krm_types:
                    new_content = re.sub(r'\bkrm\.%s\b' % type_name, f'{new_alias_for_source}.{type_name}', new_content)

        # 2. Construct the new import block
        new_import_block = []
        if uses_remaining:
            new_alias_for_source = f'krm{source_version}'
            new_import_block.append(f'\t{new_alias_for_source} {original_import_path}')
        
        if uses_promoted:
            target_import_path = f'"{GO_MODULE}/apis/{service}/{target_version}"'
            new_import_block.append(f'\tkrm {target_import_path}')

        # 3. Replace the original import line with the new block
        if new_import_block:
            # To handle cases where the import is in an import() block
            # We replace the line, preserving surrounding whitespace if possible
            new_content = new_content.replace(original_aliased_import, '\n'.join(new_import_block).strip())
        else:
            # If no types are used, remove the import line
            new_content = new_content.replace(original_aliased_import, '')

        with open(filepath, 'w') as f:
            f.write(new_content)

def promote_api_file(api_path: str, target_version: str, base_dir: str) -> dict:
    """
    Promotes an entire API package directory to a target version.
    """
    abs_api_path = validate_api_path(api_path, base_dir)
    try:
        source_version = get_version_from_path(abs_api_path)
        source_dir = os.path.dirname(abs_api_path)

        target_dir = source_dir.replace(source_version, target_version)
        os.makedirs(target_dir, exist_ok=True)

        resource_basename = os.path.basename(api_path).replace('_types.go', '')

        relevant_files = [
            'doc.go',
            'groupversion_info.go',
            'types.generated.go',
            f'{resource_basename}_types.go',
            f'{resource_basename}_identity.go',
            f'{resource_basename}_reference.go',
        ]

        for filename in relevant_files:
            source_filepath = os.path.join(source_dir, filename)
            target_filepath = os.path.join(target_dir, filename)

            if os.path.isfile(source_filepath) and filename.endswith('.go'):
                with open(source_filepath, 'r') as f:
                    content = f.read()

                new_content = content.replace(f"package {source_version}", f"package {target_version}")
                new_content = new_content.replace(source_version, target_version)

                with open(target_filepath, 'w') as f:
                    f.write(new_content)

                if filename.endswith('_types.go'):
                    with open(target_filepath, 'r') as f:
                        lines = f.readlines()

                    # Handle storageversion
                    storage_version_found = any("kubebuilder:storageversion" in line for line in lines)
                    if not storage_version_found:
                        for i, line in enumerate(lines):
                            if "// +k8s:openapi-gen=true" in line:
                                lines.insert(i, '// +kubebuilder:storageversion\n')
                                break

                    with open(target_filepath, 'w') as f:
                        f.writelines(lines)

            elif os.path.isfile(source_filepath):
                shutil.copy(source_filepath, target_filepath)

        return {"new_api_path": api_path.replace(source_version, target_version)}
    except Exception as e:
        return {"error": "ApiPromotionError", "message": str(e)}

def promote_controller_file(controller_path: str, api_path: str, target_version: str, go_module: str, base_dir: str) -> dict:
    """
    Updates the API import paths in all Go files within the controller's directory.
    """
    abs_api_path = validate_api_path(api_path, base_dir)
    try:
        abs_controller_path = validate_controller_path(controller_path, base_dir)
        source_version = get_version_from_path(abs_api_path)
        controller_dir = os.path.dirname(abs_controller_path)
        old_import_path = f'{go_module}/{os.path.dirname(api_path)}'
        new_import_path = old_import_path.replace(source_version, target_version)
        
        for filename in os.listdir(controller_dir):
            if filename.endswith('.go'):
                filepath = os.path.join(controller_dir, filename)
                with open(filepath, 'r') as f:
                    content = f.read()
                
                new_content = content.replace(old_import_path, new_import_path)
                
                with open(filepath, 'w') as f:
                    f.write(new_content)
                    
        return {"new_controller_path": controller_path}
    except Exception as e:
        return {"error": "ControllerPromotionError", "message": str(e)}

def promote_test_fixture(test_fixture_path: str, target_version: str, base_dir: str) -> dict:
    """
    Promotes a test fixture to a target version.
    """
    try:
        abs_test_fixture_path = validate_fixture_path(test_fixture_path, base_dir)
        source_version = get_version_from_path(abs_test_fixture_path)
        new_test_fixture_path_rel = test_fixture_path.replace(source_version, target_version)
        new_test_fixture_path_abs = to_abs_path(new_test_fixture_path_rel, base_dir)
        shutil.copytree(abs_test_fixture_path, new_test_fixture_path_abs)

        for root, _, files in os.walk(new_test_fixture_path_abs):
            for file in files:
                file_path = os.path.join(root, file)
                if not file_path.endswith('create.yaml') or not file_path.endswith('update.yaml'):
                    continue
                with open(file_path, 'r') as f:
                    content = f.read()
                
                new_content = content.replace(source_version, target_version)

                with open(file_path, 'w') as f:
                    f.write(new_content)

        return {"new_test_fixture_path": new_test_fixture_path_rel}
    except Exception as e:
        return {"error": "TestFixturePromotionError", "message": str(e)}

def validate_promotion(api_path: str, target_version: str, base_dir: str) -> dict:
    """
    Validates the promotion by running the CRD generation script.
    """
    abs_api_path = validate_api_path(api_path, base_dir)
    try:
        # Determine source and target directories.
        source_version = get_version_from_path(abs_api_path)
        source_dir = os.path.dirname(abs_api_path)
        target_dir = source_dir.replace(source_version, target_version)

        # Define the paths for the generated files to be deleted.
        source_generated_file = os.path.join(source_dir, 'zz_generated.deepcopy.go')
        target_generated_file = os.path.join(target_dir, 'zz_generated.deepcopy.go')

        # Delete the generated files if they exist.
        if os.path.exists(source_generated_file):
            os.remove(source_generated_file)
        if os.path.exists(target_generated_file):
            os.remove(target_generated_file)

        # Run the generation script and check for errors.
        script_path = to_abs_path('dev/tasks/generate-crds', base_dir)
        result = subprocess.run([script_path], capture_output=True, text=True, cwd=base_dir)
        if result.returncode != 0:
            return {"error": "ValidationFailed", "message": result.stderr}

        return {"success": True, "output": result.stdout}
    except Exception as e:
        return {"error": "ValidationException", "message": str(e)}

def validate_controller_compilation(controller_path: str, base_dir: str, target_version: str) -> dict:
    """
    Validates the controller by attempting to compile it.
    """
    abs_controller_path = validate_controller_path(controller_path, base_dir)
    service = os.path.basename(os.path.dirname(abs_controller_path))

    if not target_version:
        return {"error": "VersionNotFound", "message": f"Could not determine target version from {abs_controller_path}"}

    with open(abs_controller_path, 'r') as f:
        content = f.read()

    expected_import = f'"{GO_MODULE}/apis/{service}/{target_version}"'
    if expected_import not in content:
        return {"error": "ImportNotFound", "message": f"Expected import '{expected_import}' not found in {abs_controller_path}"}

    try:
        controller_dir = os.path.dirname(abs_controller_path)
        result = subprocess.run(['go', 'build', './...'], cwd=controller_dir, capture_output=True, text=True)
        if result.returncode != 0:
            return {"error": "CompilationFailed", "message": result.stderr}
        return {"success": True, "output": result.stdout}
    except Exception as e:
        return {"error": "CompilationException", "message": str(e)}

def validate_test_fixture(kind: str, base_dir: str) -> dict:
    """
    Validates the test fixture by running the hack/compare-mock script.
    """
    try:
        script_path = to_abs_path('hack/compare-mock', base_dir)
        arg = f'fixtures/{kind}'
        result = subprocess.run([script_path, arg], capture_output=True, text=True, cwd=base_dir)
        if result.returncode != 0:
            return {"error": "TestFixtureValidationFailed", "message": result.stderr}
        return {"success": True, "output": result.stdout}
    except Exception as e:
        return {"error": "TestFixtureValidationException", "message": str(e)}
