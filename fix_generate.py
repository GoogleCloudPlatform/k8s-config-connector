import re

with open("apis/networksecurity/v1alpha1/generate.sh", "r") as f:
    content = f.read()

content = re.sub(
    r'<<<<<<< HEAD\n  --resource NetworkSecurityMirroringEndpointGroup:MirroringEndpointGroup \\\n=======\n  --resource NetworkSecurityMirroringEndpointGroupAssociation:MirroringEndpointGroupAssociation \\\n>>>>>>> [^\n]+\n',
    r'  --resource NetworkSecurityMirroringEndpointGroup:MirroringEndpointGroup \\\n  --resource NetworkSecurityMirroringEndpointGroupAssociation:MirroringEndpointGroupAssociation \\\n',
    content
)

with open("apis/networksecurity/v1alpha1/generate.sh", "w") as f:
    f.write(content)
