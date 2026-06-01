import re

with open("apis/refs/v1beta1/networksecurityrefs.go", "r") as f:
    content = f.read()

content = re.sub(
    r'<<<<<<< HEAD\n(.*?)\n=======\n(.*?)\n>>>>>>> [^\n]+\n',
    r'\1\n\2\n',
    content,
    flags=re.DOTALL
)

with open("apis/refs/v1beta1/networksecurityrefs.go", "w") as f:
    f.write(content)
