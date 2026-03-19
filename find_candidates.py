import os
import yaml
import glob

crd_dir = "config/crds/resources"
apis_dir = "apis"

for root, _, files in os.walk(crd_dir):
    for f in files:
        if not f.endswith(".yaml"): continue
        path = os.path.join(root, f)
        with open(path, "r") as file:
            try:
                docs = list(yaml.safe_load_all(file))
                for doc in docs:
                    if not doc or "metadata" not in doc: continue
                    labels = doc["metadata"].get("labels", {})
                    if labels.get("cnrm.cloud.google.com/dcl2crd") != "true": continue
                    
                    spec = doc.get("spec", {})
                    group = spec.get("group", "")
                    short_group = group.split(".")[0]
                    kind = spec.get("names", {}).get("kind", "")
                    
                    versions = spec.get("versions", [])
                    beta_versions = [v["name"] for v in versions if "beta" in v.get("name", "")]
                    
                    if not beta_versions: continue
                    
                    for version in beta_versions:
                        # check if types.go exists
                        types_pattern = os.path.join(apis_dir, short_group, version, "*_types.go")
                        if not glob.glob(types_pattern):
                            print(f"{short_group} {kind} {version}")
                            break
            except Exception as e:
                pass
