## Apply the compostion
```
kubectl apply -f 01-composition.yaml
```

## Create a GCP service account for this team and grant KCC permissions

This step follows the following documentation:
https://cloud.google.com/config-connector/docs/how-to/install-namespaced

```
export NAMESPACE=team-aks
export TEAM_GCP_SA_NAME="${NAMESPACE}"
export PROJECT_ID=$(gcloud config get-value project)
export TEAM_GSA_EMAIL="${TEAM_GCP_SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"
gcloud iam service-accounts create ${TEAM_GCP_SA_NAME} --project ${PROJECT_ID}
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member="serviceAccount:${TEAM_GSA_EMAIL}" \
    --role="roles/owner"
gcloud iam service-accounts add-iam-policy-binding \
    ${TEAM_GSA_EMAIL} \
    --member="serviceAccount:${PROJECT_ID}.svc.id.goog[cnrm-system/cnrm-controller-manager-${NAMESPACE}]" \
    --role="roles/iam.workloadIdentityUser" \
    --project ${PROJECT_ID}
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member="serviceAccount:${TEAM_GSA_EMAIL}" \
    --role="roles/monitoring.metricWriter"
WORKLOAD_IDENTITY_POOL="${PROJECT_ID}.svc.id.goog"
export ASO_NAMESPACE=azureserviceoperator-system # Don’t change
export ASO_KSA=azureserviceoperator-default # Don’t change
gcloud iam service-accounts add-iam-policy-binding ${TEAM_GSA_EMAIL} \
 --role roles/iam.workloadIdentityUser \
 --member "serviceAccount:${WORKLOAD_IDENTITY_POOL}[${ASO_NAMESPACE}/${ASO_KSA}]" \
 --condition None
```

## Grant ASO permission for the GCP service account used by this team

This step follows the [following documentation](setup-Azure-ASO.md#create-a-gcp-service-account)
```
export MI_RESOURCE_GROUP="$USER"-kcc-demo-${NAMESPACE}
export MI_NAME="$USER-aso-mi"
export AZURE_SUBSCRIPTION_ID=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
export AZURE_TENANT_ID=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
AZURE_REGION=eastus
```

## Create a resource group
```
az group create -l ${AZURE_REGION} -n ${MI_RESOURCE_GROUP}
```

## Create an MI
```
az identity create --name ${MI_NAME} \
  --resource-group ${MI_RESOURCE_GROUP} \
  --location ${AZURE_REGION}
```

## Get the MI ID
```
export AZURE_CLIENT_ID=$(az identity show \
  --name ${MI_NAME} --resource-group ${MI_RESOURCE_GROUP} \
  --query "clientId" -otsv)
```

## Assign permissions to this MI

Note: The user can use other permissions to manage their resources.
```
az role assignment create \
  --assignee $AZURE_CLIENT_ID \
  --role contributor \
  --scope /subscriptions/$AZURE_SUBSCRIPTION_ID
```

## Set up impersonation

This step allows the GCP service created
[here](setup-Azure-ASO.md#create-a-gcp-service-account) to impersonate the Azure
managed identity.

In this example, we use the default GCP service account.

```
DEFAULT_GSA_EMAIL=$(kubectl get asokontroller \
  asokontroller.kontrollers.cnrm.cloud.google.com \
  -ojson | jq -r .spec.googleServiceAccount)
DEFAULT_GSA_SUB=$(gcloud iam service-accounts describe ${TEAM_GSA_EMAIL} \
 --format "value(oauth2ClientId)")

az identity federated-credential create \
  --name gsa-azure-federated-credential \
  --identity-name ${MI_NAME} \
  --resource-group ${MI_RESOURCE_GROUP} \
  --issuer https://accounts.google.com \
  --subject ${DEFAULT_GSA_SUB} \
  --audience api://AzureADTokenExchange
```

## Create a namespace for Alice's team
```
cat 02-context.yaml | envsubst | \
  kubectl apply  -f -
```
