with open("pkg/controller/direct/compute/utils.go", "r") as f:
    content = f.read()

# Replace any double } } at the end, etc.
lines = content.split('\n')
if lines[-18].strip() == '}':
    pass # ok

# let's just use replace to be safe.
