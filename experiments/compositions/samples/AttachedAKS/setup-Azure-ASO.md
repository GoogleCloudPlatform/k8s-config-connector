# Set up Azure Service Operator v2 (ASO)

Note: ASO support is only availabe for GCP managed [Config
Connector](https://cloud.google.com/config-connector/docs/overview).

Here is an overview of the ASO setup process:

1. Create a GCP service account.
1. Create an Azure managed identity.
   1. Grant this managed identity the permissions to manage the Azure resources.
   1. Allow the GCP service account to impersonate the Azure managed identity.
1. Create the ASO controller with the GCP service account.

## Create a GCP service account

```
PROJECT_ID=$(gcloud config get-value project)
PROJECT_NUMBER=$(gcloud projects describe $PROJECT_ID --format="value(projectNumber)")


# Create the service account
gcloud iam service-accounts create $USER-allotrope \
 --description="Allotrope proof of concept" \
 --display-name="Allotrope POC"

export GSA_EMAIL=$USER-allotrope@${PROJECT_ID}.iam.gserviceaccount.com

# Get the workload identity pool for the gke/kcc cluster
WORKLOAD_IDENTITY_POOL="${PROJECT_ID}.svc.id.goog"

# Grant workload identity bindings permissions
export ASO_NAMESPACE=kontrollers-azureserviceoperator-system # Don’t change
export ASO_KSA=azureserviceoperator-default # Don’t change
gcloud iam service-accounts add-iam-policy-binding ${GSA_EMAIL} \
 --role roles/iam.workloadIdentityUser \
 --member "serviceAccount:${WORKLOAD_IDENTITY_POOL}[${ASO_NAMESPACE}/${ASO_KSA}]" \
 --condition None
```

## Create an Azure managed identity and allow the GCP service account to
impersonate it

```
# Please change the parameters in this section
export MI_RESOURCE_GROUP="${USER}-kcc-demo"
export MI_NAME="${USER}-aso-mi"
AZURE_SUBSCRIPTION_ID=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
AZURE_TENANT_ID=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
AZURE_REGION=eastus

# Create a resource group
az group create -l ${AZURE_REGION} -n ${MI_RESOURCE_GROUP}

# Create an MI
az identity create --name ${MI_NAME} \
  --resource-group ${MI_RESOURCE_GROUP} \
  --location ${AZURE_REGION}

# Get the MI ID
AZURE_CLIENT_ID=$(az identity show \
  --name ${MI_NAME} --resource-group ${MI_RESOURCE_GROUP} \
  --query "clientId" -otsv)
MI_PRINCIPAL_ID=$(az identity show \
  --name ${MI_NAME} --resource-group ${MI_RESOURCE_GROUP} \
  --query "principalId" -otsv)

# Assign the permissions to this MI.
# The user can use other permissions to manage their resources.
az role assignment create \
  --assignee ${MI_PRINCIPAL_ID} \
  --role contributor \
  --scope /subscriptions/${AZURE_SUBSCRIPTION_ID}

# Allow the GCP service account to impersonate this MI
GSA_SUB=$(gcloud iam service-accounts describe ${GSA_EMAIL}  \
 --format "value(oauth2ClientId)")
az identity federated-credential create \
  --name gsa-azure-federated-credential \
  --identity-name ${MI_NAME} \
  --resource-group ${MI_RESOURCE_GROUP} \
  --issuer https://accounts.google.com \
  --subject ${GSA_SUB} \
  --audience api://AzureADTokenExchange
```

## Create the ASO controller

```
cat <<EOF > /tmp/aso.yaml
apiVersion: kontrollers.cnrm.cloud.google.com/v1alpha1
kind: ASOKontroller
metadata:
 name: asokontroller.kontrollers.cnrm.cloud.google.com
spec:
 defaultAzureSubscriptionID: "${AZURE_SUBSCRIPTION_ID}"
 defaultAzureTenantID: "${AZURE_TENANT_ID}"
 defaultAzureClientID: "${AZURE_CLIENT_ID}"
 crdPatterns:
 - "resources.azure.com/resourcegroup" # Use any pattern that fits your need.
 - "containerservice.azure.com/*"
 googleServiceAccount: "${GSA_EMAIL}"
EOF

kubectl apply -f /tmp/aso.yaml
```
