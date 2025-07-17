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

def promote_controller_file(controller_path: str, target_version: str) -> str:
    """
    Promotes a controller file to a target version.

    This involves:
    1.  Reading the content of the controller file.
    2.  Replacing all occurrences of "v1alpha1" with the target version.
    3.  Writing the modified content back to the controller file.
    4.  Returning the path of the modified file.
    """
    with open(controller_path, 'r') as f:
        content = f.read()

    new_content = content.replace("v1alpha1", target_version)

    with open(controller_path, 'w') as f:
        f.write(new_content)
        
    return controller_path

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
