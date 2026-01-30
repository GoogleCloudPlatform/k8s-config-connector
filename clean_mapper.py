import re

with open('pkg/controller/direct/orgpolicy/mapper.generated.go', 'r') as f:
    content = f.read()

# Remove v1alpha1 import
content = re.sub(r'\t.*"github.com/GoogleCloudPlatform/k8s-config-connector/apis/orgpolicy/v1alpha1"\n', '', content)

parts = content.split('func ')
new_parts = [parts[0]] # Header and imports
seen_funcs = set()

for part in parts[1:]:
    # Extract function name
    # func Name(...) ...
    # The part starts with Name(...)
    func_name = part.split('(')[0]

    if 'krmorgpolicyv1alpha1' in part:
        continue

    if func_name in seen_funcs:
        continue

    seen_funcs.add(func_name)
    new_parts.append(part)

new_content = 'func '.join(new_parts)

with open('pkg/controller/direct/orgpolicy/mapper.generated.go', 'w') as f:
    f.write(new_content)
