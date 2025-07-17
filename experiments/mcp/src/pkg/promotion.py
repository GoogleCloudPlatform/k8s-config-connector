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

def get_git_root():
    """Returns the root directory of the git repository."""
    try:
        return subprocess.check_output(['git', 'rev-parse', '--show-toplevel'], text=True).strip()
    except subprocess.CalledProcessError as e:
        raise RuntimeError("Could not find git repository root.") from e

GIT_ROOT = get_git_root()

def to_abs_path(path: str) -> str:
    """Converts a relative path to an absolute path based on the git root."""
    if os.path.isabs(path):
        return path
    return os.path.join(GIT_ROOT, path)

def get_version_from_path(path: str) -> str:
    """
    Extracts the version from a file path.
    Assumes the version is in the format v[a-z0-9]+ (e.g., v1alpha1, v1beta1).
    """
    match = re.search(r'v[a-z0-9]+', path)
    if match:
        return match.group(0)
    raise ValueError(f"Could not find version in path: {path}")

def get_kind_from_path(path: str) -> str:
    """
    Extracts the kind from a file path.
    Assumes the kind is in the format xxx_types.go.
    """
    base_name = os.path.basename(path)
    kind_name = base_name.replace('_types.go', '')
    # Convert snake_case to CamelCase
    return ''.join(word.title() for word in kind_name.split('_'))

def promote_api_file(api_path: str, target_version: str) -> dict:
    """
    Promotes an entire API package directory to a target version.
    """
    try:
        abs_api_path = to_abs_path(api_path)
        # Determine the source version and directory from the input path.
        source_version = get_version_from_path(abs_api_path)
        source_dir = os.path.dirname(abs_api_path)
        
        # Determine the target directory and create it.
        target_dir = source_dir.replace(source_version, target_version)
        os.makedirs(target_dir, exist_ok=True)

        # Iterate through all files in the source directory.
        for filename in os.listdir(source_dir):
            source_filepath = os.path.join(source_dir, filename)
            target_filepath = os.path.join(target_dir, filename)

            if os.path.isfile(source_filepath):
                # If the file is a Go file, process its content.
                if filename.endswith('.go'):
                    with open(source_filepath, 'r') as f:
                        content = f.read()
                    
                    # Replace the package declaration.
                    new_content = content.replace(f"package {source_version}", f"package {target_version}")
                    
                    # Replace other occurrences of the source version.
                    new_content = new_content.replace(source_version, target_version)
                    
                    with open(target_filepath, 'w') as f:
                        f.write(new_content)

                    # Add storage version annotation
                    if filename.endswith('_types.go'):
                        kind = get_kind_from_path(filename)
                        with open(target_filepath, 'r') as f:
                            lines = f.readlines()
                        
                        new_lines = []
                        for line in lines:
                            new_lines.append(line)
                            if f"type {kind} struct" in line:
                                new_lines.insert(len(new_lines) - 1, '// +kubebuilder:storageversion\n')
                        
                        with open(target_filepath, 'w') as f:
                            f.writelines(new_lines)

                # Otherwise, just copy the file.
                else:
                    shutil.copy(source_filepath, target_filepath)

        # Return the path to the new API file.
        return {"new_api_path": api_path.replace(source_version, target_version)}
    except Exception as e:
        return {"error": "ApiPromotionError", "message": str(e)}

def promote_controller_file(controller_path: str, api_path: str, target_version: str, go_module: str) -> dict:
    """
    Updates the API import paths in all Go files within the controller's directory.
    """
    try:
        abs_api_path = to_abs_path(api_path)
        abs_controller_path = to_abs_path(controller_path)
        source_version = get_version_from_path(abs_api_path)
        
        controller_dir = os.path.dirname(abs_controller_path)
        for filename in os.listdir(controller_dir):
            if filename.endswith('.go'):
                filepath = os.path.join(controller_dir, filename)
                with open(filepath, 'r') as f:
                    content = f.read()
                
                new_content = content.replace(f'/{source_version}', f'/{target_version}')
                
                with open(filepath, 'w') as f:
                    f.write(new_content)
                    
        return {"new_controller_path": controller_path}
    except Exception as e:
        return {"error": "ControllerPromotionError", "message": str(e)}

def promote_test_fixture(test_fixture_path: str, target_version: str) -> dict:
    """
    Promotes a test fixture to a target version.
    """
    try:
        abs_test_fixture_path = to_abs_path(test_fixture_path)
        source_version = get_version_from_path(abs_test_fixture_path)
        new_test_fixture_path_rel = test_fixture_path.replace(source_version, target_version)
        new_test_fixture_path_abs = to_abs_path(new_test_fixture_path_rel)
        shutil.copytree(abs_test_fixture_path, new_test_fixture_path_abs)

        for root, _, files in os.walk(new_test_fixture_path_abs):
            for file in files:
                file_path = os.path.join(root, file)
                with open(file_path, 'r') as f:
                    content = f.read()
                
                new_content = content.replace(source_version, target_version)

                with open(file_path, 'w') as f:
                    f.write(new_content)

        return {"new_test_fixture_path": new_test_fixture_path_rel}
    except Exception as e:
        return {"error": "TestFixturePromotionError", "message": str(e)}

def validate_promotion(api_path: str, target_version: str) -> dict:
    """
    Validates the promotion by running the CRD generation script.
    """
    try:
        abs_api_path = to_abs_path(api_path)
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
        script_path = to_abs_path('dev/tasks/generate-crds')
        result = subprocess.run([script_path], capture_output=True, text=True, cwd=GIT_ROOT)
        if result.returncode != 0:
            return {"error": "ValidationFailed", "message": result.stderr}

        return {"success": True, "output": result.stdout}
    except Exception as e:
        return {"error": "ValidationException", "message": str(e)}

def validate_controller_compilation(controller_path: str) -> dict:
    """
    Validates the controller by attempting to compile it.
    """
    try:
        abs_controller_path = to_abs_path(controller_path)
        controller_dir = os.path.dirname(abs_controller_path)
        result = subprocess.run(['go', 'build', './...'], cwd=controller_dir, capture_output=True, text=True)
        if result.returncode != 0:
            return {"error": "CompilationFailed", "message": result.stderr}
        return {"success": True, "output": result.stdout}
    except Exception as e:
        return {"error": "CompilationException", "message": str(e)}

def validate_test_fixture(kind: str) -> dict:
    """
    Validates the test fixture by running the hack/compare-mock script.
    """
    try:
        script_path = to_abs_path('hack/compare-mock')
        arg = f'fixtures/{kind}'
        result = subprocess.run([script_path, arg], capture_output=True, text=True, cwd=GIT_ROOT)
        if result.returncode != 0:
            return {"error": "TestFixtureValidationFailed", "message": result.stderr}
        return {"success": True, "output": result.stdout}
    except Exception as e:
        return {"error": "TestFixtureValidationException", "message": str(e)}