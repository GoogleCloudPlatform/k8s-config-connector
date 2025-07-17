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

def get_version_from_path(path: str) -> str:
    """
    Extracts the version from a file path.
    Assumes the version is in the format v[a-z0-9]+ (e.g., v1alpha1, v1beta1).
    """
    match = re.search(r'v[a-z0-9]+', path)
    if match:
        return match.group(0)
    raise ValueError(f"Could not find version in path: {path}")

def promote_api_file(api_path: str, target_version: str) -> str:
    """
    Promotes an entire API package directory to a target version.

    This involves:
    1.  Determining the source and target directory paths.
    2.  Creating the target directory.
    3.  Iterating through all files in the source directory.
    4.  For each .go file, updating the package declaration and replacing
        occurrences of the source version with the target version.
    5.  Copying non-.go files as is.
    6.  Returning the path of the new, promoted API file corresponding to the
        original api_path.
    """
    # Determine the source version and directory from the input path.
    source_version = get_version_from_path(api_path)
    source_dir = os.path.dirname(api_path)
    
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
                
                # Replace other occurrences of the source version. This is a broad
                # replacement, but it's necessary to update import paths and other
                # versioned references.
                new_content = new_content.replace(source_version, target_version)
                
                with open(target_filepath, 'w') as f:
                    f.write(new_content)
            # Otherwise, just copy the file.
            else:
                shutil.copy(source_filepath, target_filepath)

    # Return the path to the new API file.
    return api_path.replace(source_version, target_version)

def promote_controller_file(controller_path: str, api_path: str, target_version: str, go_module: str) -> str:
    """
    Updates the API import paths in all Go files within the controller's directory.

    This involves:
    1.  Determining the source and target API import paths.
    2.  Iterating through all .go files in the controller's directory.
    3.  Replacing the old API import path with the new one.
    4.  Returning the path of the controller directory.
    """
    source_version = get_version_from_path(api_path)
    
    controller_dir = os.path.dirname(controller_path)
    for filename in os.listdir(controller_dir):
        if filename.endswith('.go'):
            filepath = os.path.join(controller_dir, filename)
            with open(filepath, 'r') as f:
                content = f.read()
            
            new_content = content.replace(f'/{source_version}', f'/{target_version}')
            
            with open(filepath, 'w') as f:
                f.write(new_content)
                
    return controller_path

def validate_controller_compilation(controller_path: str) -> tuple[bool, str]:
    """
    Validates the controller by attempting to compile it.

    This involves:
    1.  Determining the controller's directory.
    2.  Running `go build` in that directory.
    3.  Returning whether the build was successful and any output.
    """
    controller_dir = os.path.dirname(controller_path)
    result = subprocess.run(['go', 'build', './...'], cwd=controller_dir, capture_output=True, text=True)
    if result.returncode != 0:
        return False, result.stderr
    return True, result.stdout

def promote_test_fixture(test_fixture_path: str, target_version: str) -> str:
    """
    Promotes a test fixture to a target version.

    This involves:
    1.  Copying the entire source test fixture directory to a new target directory.
    2.  Iterating through all files in the new target directory.
    3.  Replacing all occurrences of the source version with the target version in the content of each file.
    4.  Returning the path of the new test fixture directory.
    """
    source_version = get_version_from_path(test_fixture_path)
    new_test_fixture_path = test_fixture_path.replace(source_version, target_version)
    shutil.copytree(test_fixture_path, new_test_fixture_path)

    for root, _, files in os.walk(new_test_fixture_path):
        for file in files:
            file_path = os.path.join(root, file)
            with open(file_path, 'r') as f:
                content = f.read()
            
            new_content = content.replace(source_version, target_version)

            with open(file_path, 'w') as f:
                f.write(new_content)

    return new_test_fixture_path

def validate_promotion(api_path: str, target_version: str) -> tuple[bool, str]:
    """
    Validates the promotion by running the CRD generation script.

    This involves:
    1.  Determining the source and target directories.
    2.  Deleting the zz_generated.deepcopy.go files from both directories.
    3.  Running the dev/tasks/generate-crds script.
    4.  Returning whether the script was successful and any output.
    """
    # Determine source and target directories.
    source_version = get_version_from_path(api_path)
    source_dir = os.path.dirname(api_path)
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
    result = subprocess.run(['./dev/tasks/generate-crds'], capture_output=True, text=True)
    if result.returncode != 0:
        return False, result.stderr

    return True, result.stdout
