## Step 1. Apply the CR to create the composition of AKS and its attachment
```
kubectl apply -f 03-attached-1.yaml
```


## Step 2. Log in to the cluster and apply the attached cluster install manifest
```
AKS_NAME=$(cat 03-attached-1.yaml |yq .metadata.name)
ATTACHED_REGION=$(cat 03-attached-1.yaml |yq .spec.gcpRegion)
ATTACHED_PLATFORM_VERSION=$(cat 03-attached-1.yaml |yq .spec.attachedPlatformVersion)

gcloud container attached clusters generate-install-manifest \
  $AKS_NAME \
  --location=${ATTACHED_REGION} \
  --platform-version ${ATTACHED_PLATFORM_VERSION} \
  --output-file=/tmp/install-agent-${AKS_NAME}.yaml

az aks get-credentials --name ${AKS_NAME}-aks \
  --resource-group ${AKS_NAME}-rg 

kubectl apply -f /tmp/install-agent-${AKS_NAME}.yaml
```


## Commands to check progress:
```
kubectl get AttachedAKS -n team-aks

kubectl get ResourceGroup.resources.azure.com \
  -n team-aks

kubectl get managedcluster.containerservice.azure.com \
  -n team-aks

kubectl get cm -n team-aks

kubectl get containerattachedcluster \
  -n team-aks
```
