import urllib.request
import re

url = "https://cloud.google.com/vertex-ai/docs/workbench/reference/rest/v1/projects.locations.instances/create"
try:
    with urllib.request.urlopen(url) as response:
        html = response.read().decode('utf-8')
        matches = re.findall(r'https://[^/]+\.googleapis\.com/[^"]+', html)
        for m in set(matches):
            if 'notebooks' in m or 'vertex' in m or 'instances' in m:
                print(m)
except Exception as e:
    print("Failed")
