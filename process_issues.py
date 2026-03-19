import json
import subprocess
import glob
import yaml

candidates = []
for filepath in glob.glob("config/crds/resources/*.yaml"):
    try:
        with open(filepath, "r") as f:
            docs = yaml.safe_load_all(f)
            for doc in docs:
                if not doc: continue
                labels = doc.get("metadata", {}).get("labels", {})
                if labels.get("cnrm.cloud.google.com/dcl2crd") == "true":
                    group = doc.get("spec", {}).get("group", "").split(".")[0]
                    kind = doc.get("spec", {}).get("names", {}).get("kind", "")
                    versions = [v.get("name") for v in doc.get("spec", {}).get("versions", [])]
                    
                    if "v1beta1" in versions:
                        types_file = f"apis/{group}/v1beta1/{kind.lower()}_types.go"
                        if not glob.glob(f"apis/{group}/v1beta1/*_types.go"):
                            candidates.append((group, kind))
    except Exception as e:
        pass

# Ensure deterministic order
candidates = sorted(list(set(candidates)))

with open("existing_issues.json") as f:
    issues = json.load(f)

# Count pending
pending_issues = sum(1 for i in issues if i.get("state", "OPEN").upper() == "OPEN")

required_labels = {"overseer", "area/direct", "priority/medium"}

issue_created = False
printed_pending_message = False

for group, kind in candidates:
    target_title = f"Create generate.sh and types.go files for {group} {kind}".lower()
    
    # Check if issue exists
    existing_issue = None
    for issue in issues:
        if issue["title"].lower() == target_title:
            existing_issue = issue
            break
            
    if existing_issue:
        # Check labels
        issue_labels = {l["name"] for l in existing_issue.get("labels", [])}
        missing_labels = required_labels - issue_labels
        if missing_labels:
            print(f"Injecting missing labels {missing_labels} for issue #{existing_issue['number']}")
            labels_args = ",".join(missing_labels)
            subprocess.run(["gh", "issue", "edit", str(existing_issue["number"]), "--add-label", labels_args])
    else:
        # Issue doesn't exist
        if pending_issues >= 10:
            if not printed_pending_message:
                print(f"There are already {pending_issues} pending issues and skip creating new ones until some of the existing issues are resolved.")
                printed_pending_message = True
        elif not issue_created:
            print(f"CREATE:{group}:{kind}")
            issue_created = True
            break
