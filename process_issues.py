import json
import os
import subprocess
import glob
import yaml

with open('all_issues.json') as f:
    all_issues = json.load(f)

# Count pending (open) issues
open_issues = [issue for issue in all_issues if issue.get("state", "OPEN") == "OPEN"]
pending_count = len(open_issues)

# Candidate resources
candidate_files = []
for root, _, files in os.walk("config/crds/resources"):
    for file in files:
        if file.endswith(".yaml"):
            filepath = os.path.join(root, file)
            with open(filepath) as f:
                content = f.read()
                if "cnrm.cloud.google.com/dcl2crd: \"true\"" in content:
                    candidate_files.append(filepath)

candidates = []
for filepath in candidate_files:
    with open(filepath) as f:
        docs = yaml.safe_load_all(f)
        for doc in docs:
            if not doc: continue
            if doc.get("kind") == "CustomResourceDefinition":
                group = doc["spec"]["group"]
                kind = doc["spec"]["names"]["kind"]
                plural = doc["spec"]["names"]["plural"]
                versions = doc["spec"]["versions"]
                
                is_beta = any(v["name"] == "v1beta1" for v in versions)
                if not is_beta:
                    continue
                
                group_prefix = group.split(".")[0]
                
                types_exists = False
                dir_path = f"apis/{group_prefix}/v1beta1/"
                if os.path.isdir(dir_path):
                    for fname in os.listdir(dir_path):
                        if fname.lower().replace("_", "") == f"{kind.lower()}types.go" or fname.endswith("_types.go") and kind.lower() in fname.lower():
                            types_exists = True
                            break
                        if fname == f"{kind.lower()}_types.go":
                            types_exists = True
                            
                if not types_exists:
                    candidates.append((group, kind, group_prefix, plural))

required_labels = {"overseer", "area/direct", "priority/medium"}

# To limit issue creation
issue_created = False

for group, kind, group_prefix, plural in candidates:
    # Build possible titles
    title1 = f"Create generate.sh and types.go files for {group} {kind}"
    title2 = f"Create generate.sh and types.go files for {group_prefix} {kind}"
    title3 = f"Create generate.sh and types.go files for {group_prefix.capitalize()} {kind}"
    title4 = f"Create generate.sh and types.go files for {group_prefix.lower()} {kind}"
    
    existing_issue = None
    for issue in all_issues:
        t = issue["title"].lower()
        if t == title1.lower() or t == title2.lower() or t == title3.lower() or t == title4.lower():
            existing_issue = issue
            break
            
    if existing_issue:
        existing_labels = {lbl["name"] for lbl in existing_issue.get("labels", [])}
        missing_labels = required_labels - existing_labels
        if missing_labels:
            print(f"Injecting missing labels {missing_labels} for issue #{existing_issue['number']} ({existing_issue['title']})")
            subprocess.run(["gh", "issue", "edit", str(existing_issue["number"]), "--add-label", ",".join(missing_labels)])
        continue
    else:
        # No existing issue found
        if not issue_created:
            if pending_count > 10:
                print(f"There are already {pending_count} pending issues. Skipping creation of new issue for {group} {kind}.")
                # Don't create, but don't exit either so we can continue fixing labels for other candidates
            else:
                print(f"Would create issue for {group} {kind} ...")
                # Creating issue
                body = f"""As part of moving resources from terraform controllers to direct controllers (Epic #5954), we need to create the Go types for `{kind}`.

Currently, `{kind}` is managed by the Terraform controller (marked with `tf2crd=true`). The goal is to create the Go types in `apis/{group_prefix}/v1beta1/` so that we can eventually migrate the controller implementation to the "direct" approach.

### Instructions

1.  **Create a generate.sh**:
    Create `apis/{group_prefix}/v1beta1/generate.sh` which includes `{kind}`.
    It likely maps to something like `google.cloud.{group_prefix}.v1`.
    Example:
    ```bash
    go run . generate-types \\
      --service google.cloud.{group_prefix}.v1 \\
      --api-version {group}/v1beta1 \\
      --resource {kind}:{kind} \\
      --include-skipped-output

    go run . generate-mapper \\
      --service google.cloud.{group_prefix}.v1 \\
      --api-version {group}/v1beta1 \\
      --include-skipped-output
    ```

2.  Set the write permission on the new `apis/{group_prefix}/v1beta1/generate.sh` file. You should do this by running both `chmod +x apis/{group_prefix}/v1beta1/generate.sh` and `git add --chmod=+x apis/{group_prefix}/v1beta1/generate.sh`.

3.  **Generate Scaffolding**:
    Run `apis/{group_prefix}/v1beta1/generate.sh`. This should create `apis/{group_prefix}/v1beta1/{kind.lower()}_types.go`.

4.  **Iterate on Types**:
    Compare the generated CRD with the existing one using `dev/tasks/diff-crds`.
    Modify `apis/{group_prefix}/v1beta1/{kind.lower()}_types.go` until the CRD matches the existing one at `config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_{plural}.{group}.yaml`.

    **Acceptance Criteria:**
    - Running `dev/tasks/diff-crds` should not show differences (or minimal acceptable ones like descriptions).
    - Ensure that running the check_crd_equivalence MCP on the CRD should return EQUIVALENT.
    - Changes to the schema (fields added/removed) are NOT acceptable.

5.  **Copyright Headers**:
    Ensure that new files have the correct copyright header:
    ```go
    // Copyright 2026 Google LLC
    ```
    Please do not change the copyright on existing files.

6.  **Labels**:
    Ensure the controller-runtime annotations match the existing CRD labels, including:
    ```go
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
    // +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
    ```
    The goal is to maintain these annotations, not add an annotation if it is missing.

7.  **Status**:
    `status.observedGeneration` should be an `*int64`.

8. **Generate Mappers**:
   - Running `dev/tasks/generate-types-and-mappers` will generate the mapper code once the `apis/{group_prefix}/v1beta1/{kind.lower()}_types.go` file is generating an equivalent CRD.
   - Run `make all-binary` to ensure the generated mapper code compiles. Please fix any issue discovered by this compilation.

This issue is part of Epic #5954.
"""
                title = f"Create generate.sh and types.go files for {group} {kind}"
                with open('body.txt', 'w') as bf:
                    bf.write(body)
                
                # I use the group as full Group
                print(f"Creating issue: {title}")
                subprocess.run(["gh", "issue", "create", "--title", title, "--body-file", "body.txt", "--label", "overseer,area/direct,priority/medium", "--milestone", "none"]) # milestones can't easily be set by name unless configured, ignoring
                issue_created = True

