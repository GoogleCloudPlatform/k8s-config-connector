import re
with open("apis/networksecurity/v1alpha1/generate.sh", "r") as f:
    content = f.read()

content = re.sub(r'<<<<<<< HEAD\ncd \$\{REPO_ROOT\}\n=======\n(.*?)\n>>>>>>> [^\n]+\n', r'\1\n', content, flags=re.DOTALL)

with open("apis/networksecurity/v1alpha1/generate.sh", "w") as f:
    f.write(content)
