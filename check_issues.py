import os
import yaml
import glob
import subprocess
import json

# 1. Find candidates
resources_dir = "config/crds/resources"
candidates = []
for filename in os.listdir(resources_dir):
    if not filename.endswith(".yaml"):
        continue
    filepath = os.path.join(resources_dir, filename)
    with open(filepath, "r") as f:
        try:
            docs = list(yaml.safe_load_all(f))
        except:
            continue
        for doc in docs:
            if not doc:
                continue
            metadata = doc.get("metadata", {})
            labels = metadata.get("labels", {})
            if labels.get("cnrm.cloud.google.com/dcl2crd") == "true":
                spec = doc.get("spec", {})
                group = spec.get("group", "")
                kind = spec.get("names", {}).get("kind", "")
                
                short_group = group.split(".")[0]
                versions = spec.get("versions", [])
                has_beta = any("beta" in v.get("name", "") for v in versions)
                
                if has_beta:
                    types_pattern = f"apis/{short_group}/v1beta1/*_types.go"
                    if not glob.glob(types_pattern):
                        candidates.append((short_group, kind, group))

# 2. Get all existing issues related to the task
output = subprocess.check_output([
    "gh", "issue", "list", 
    "--search", "Create generate.sh and types.go files for in:title", 
    "--state", "all", 
    "--limit", "200", 
    "--json", "number,title,labels"
])
issues = json.loads(output.decode('utf-8'))

# 3. Get open (pending) issues count
pending_output = subprocess.check_output([
    "gh", "issue", "list", 
    "--search", "Create generate.sh and types.go files for in:title", 
    "--state", "open", 
    "--limit", "200", 
    "--json", "number"
])
pending_issues = json.loads(pending_output.decode('utf-8'))
pending_count = len(pending_issues)

# 4. Iterate over candidates
for group, kind, full_group in candidates:
    issue_title = f"Create generate.sh and types.go files for {group} {kind}"
    
    # Check if issue already exists (case insensitive)
    existing_issue = None
    for issue in issues:
        if issue["title"].lower() == issue_title.lower():
            existing_issue = issue
            break
            
    if existing_issue:
        # Check labels
        current_labels = [l["name"] for l in existing_issue["labels"]]
        required_labels = ["overseer", "area/direct", "priority/medium"]
        missing_labels = [l for l in required_labels if l not in current_labels]
        
        if missing_labels:
            print(f"Injecting missing labels {missing_labels} into issue #{existing_issue['number']}")
            subprocess.check_call([
                "gh", "issue", "edit", str(existing_issue["number"]), 
                "--add-label", ",".join(missing_labels)
            ])
        continue # Move to the next candidate
        
    # If no existing issue for this candidate
    if pending_count >= 10:
        print(f"There are already {pending_count} pending issues for this task (more than 10). Skipping creating new ones until some of the existing issues are resolved.")
        break
    else:
        # It shouldn't reach here given the current state of pending issues, but just in case
        print(f"Would create issue for {group} {kind}")
        break
