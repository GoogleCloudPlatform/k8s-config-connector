import json
import subprocess
import yaml
import glob

# Get candidates
files = glob.glob("config/crds/resources/*.yaml")
candidates = []
for f in files:
    with open(f) as file:
        try:
            doc = yaml.safe_load(file)
            if doc.get("metadata", {}).get("labels", {}).get("cnrm.cloud.google.com/dcl2crd") == "true":
                group = doc.get("spec", {}).get("group", "")
                kind = doc.get("spec", {}).get("names", {}).get("kind", "")
                versions = doc.get("spec", {}).get("versions", [])
                for v in versions:
                    if "beta" in v.get("name", ""):
                        group_short = group.split(".")[0]
                        types_pattern = f"apis/{group_short}/v1beta1/*_types.go"
                        if not glob.glob(types_pattern):
                            candidates.append((group_short, kind))
        except Exception:
            pass

candidates = sorted(list(set(candidates)))

# Get open issues
out = subprocess.check_output(["gh", "issue", "list", "--search", "Create generate.sh and types.go files for in:title state:open", "--json", "number,title,labels", "--limit", "100"])
open_issues = json.loads(out.decode('utf-8'))

issue_count = len(open_issues)
print(f"Total open issues: {issue_count}")

# We need to check both open and closed issues for the candidates
# So for each candidate, check if an issue exists
for group, kind in candidates:
    query = f"Create generate.sh and types.go files for {group} {kind} in:title"
    out = subprocess.check_output(["gh", "issue", "list", "--search", query, "--state", "all", "--json", "number,labels"])
    issues = json.loads(out.decode('utf-8'))
    
    if issues:
        issue = issues[0]
        number = issue["number"]
        existing_labels = [l["name"] for l in issue["labels"]]
        
        required_labels = ["overseer", "area/direct", "priority/medium"]
        missing_labels = [l for l in required_labels if l not in existing_labels]
        
        if missing_labels:
            print(f"Adding labels {missing_labels} to issue {number} for {group} {kind}")
            subprocess.check_call(["gh", "issue", "edit", str(number), "--add-label", ",".join(missing_labels)])
        else:
            print(f"Issue {number} for {group} {kind} already has required labels.")
    else:
        # Issue doesn't exist
        if issue_count >= 10:
            print(f"Skipping {group} {kind}: already 10 or more pending issues.")
        else:
            print(f"We should create an issue for {group} {kind}")
            # The prompt says "In a single run create at most one issue to avoid overwhelming the team."
            # So we only create one issue if issue_count < 10.
            # But wait, issue_count is already 13.
