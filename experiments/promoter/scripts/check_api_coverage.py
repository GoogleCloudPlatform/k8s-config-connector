# This script verifies if a resource kind has full API coverage by running the API checks test suite.
# Full API coverage is a requirement for a resource to be promoted to Beta.

import json
import subprocess
import os
import glob
import yaml

def get_kind_to_plural_map(project_root):
    """Parses all CRDs to build a map from a resource Kind to its plural form."""
    kind_to_plural_map = {}
    crd_files = glob.glob(os.path.join(project_root, 'crds', '*.yaml'))
    for crd_file in crd_files:
        with open(crd_file, 'r') as f:
            try:
                crd = yaml.safe_load(f)
                if crd and 'spec' in crd and 'names' in crd['spec']:
                    names = crd['spec']['names']
                    if 'kind' in names and 'plural' in names:
                        kind = names['kind']
                        plural = names['plural']
                        kind_to_plural_map[kind] = plural
            except yaml.YAMLError as e:
                print(f"Warning: Could not parse CRD file {crd_file}: {e}")
    return kind_to_plural_map

def main():
    """Verifies if an alpha kind meets all beta level API requirements."""
    # Get the absolute path to the project root
    try:
        project_root = subprocess.check_output(
            ["git", "rev-parse", "--show-toplevel"], text=True
        ).strip()
    except subprocess.CalledProcessError:
        print("Error: This script must be run within a git repository.")
        exit(1)

    # Run a sanity check to update any necessary test data
    print("Running sanity check to update test data...")
    command = "WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/..."
    result = subprocess.run(command, shell=True, capture_output=True, text=True, cwd=project_root)
    if result.returncode != 0:
        print("Error during sanity check:")
        print(f"  stdout: {result.stdout}")
        print(f"  stderr: {result.stderr}")
        # Even if it fails, we continue, as it might be due to a specific test failing, 
        # and we want to check all candidates.

    candidates_file = os.path.join(project_root, 'experiments', 'promoter', 'results', 'candidates.json')

    with open(candidates_file, 'r') as f:
        candidates = json.load(f)

    kind_to_plural_map = get_kind_to_plural_map(project_root)
    updated_candidates = []
    exceptions_path = os.path.join(project_root, 'tests', 'apichecks', 'testdata', 'exceptions')
    exception_files = glob.glob(os.path.join(exceptions_path, '*'))

    for candidate in candidates:
        kind = candidate.get('kind')
        service = candidate.get('service')
        if not kind or not service:
            print(f"Skipping candidate without kind or service: {candidate}")
            candidate['apiValidation'] = False
            updated_candidates.append(candidate)
            continue

        print(f"Verifying API for kind: {kind}")

        # Check for exceptions
        kind_plural = kind_to_plural_map.get(kind)
        if not kind_plural:
            print(f"  FAIL: Could not determine plural for kind '{kind}'.")
            candidate['apiValidation'] = False
            updated_candidates.append(candidate)
            continue

        search_string = f"crd={kind_plural.lower()}.{service}.cnrm.cloud.google.com"
        
        found_exception = False
        for file_path in exception_files:
            if not os.path.isfile(file_path):
                continue
            with open(file_path, 'r') as f:
                if search_string in f.read():
                    print(f"  FAIL: Found exception for {kind} in {file_path}")
                    candidate['apiValidation'] = False
                    found_exception = True
                    break
        
        if found_exception:
            updated_candidates.append(candidate)
            continue

        # Run the specific test
        command = f"TARGET_KIND={kind} go test ./tests/apichecks/... -run TestCRDFieldPresenceInUnstructured"
        result = subprocess.run(command, shell=True, capture_output=True, text=True, cwd=project_root)

        if result.returncode == 0:
            print(f"  PASS: {kind}")
            candidate['apiValidation'] = True
        else:
            print(f"  FAIL: {kind}")
            print(f"    stdout: {result.stdout}")
            print(f"    stderr: {result.stderr}")
            candidate['apiValidation'] = False
        
        updated_candidates.append(candidate)

    with open(candidates_file, 'w') as f:
        json.dump(updated_candidates, f, indent=2)

    print("\nUpdated candidates.json with apiValidation results.")

if __name__ == "__main__":
    main()