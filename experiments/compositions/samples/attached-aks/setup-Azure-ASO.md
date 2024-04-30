# Setup Azure Service Operator v2 (ASO)

Note: The ASO support is only availabe for GCP managed [Config Connector](https://cloud.google.com/config-connector/docs/overview).

Here is an overview of the ASO setup process. 

1. Create a GCP service account (or use the default service account).
1. Create an Azure managed identity.
   1. Grant this managed identity the permissions to manage the Azure resouces.
   1. Allow the GCP service account to inpersonate the Azure managed identity.
1. Create ASO controller with the GCP service account.

## Create a GCP service account

### [Optional] Use the default GCP service account
User can optionally skip this step if they want to use the default service account in GCP managed [Config Connector](https://cloud.google.com/config-connector/docs/overview). And the default service account can be obtained by 

```
PROJECT_ID=$(gcloud config get-value project)
PROJECT_NUMBER=$(gcloud projects describe $PROJECT_ID --format="value(projectNumber)")

GSA_EMAIL="service-${PROJECT_NUMBER}@gcp-sa-yakima.iam.gserviceaccount.com"
```

### Create a GCP service account

```
PROJECT_ID=$(gcloud config get-value project)
PROJECT_NUMBER=$(gcloud projects describe $PROJECT_ID --format="value(projectNumber)")


# Create service account
gcloud iam service-accounts create $USER-allotrope \
 --description="Allotrope Proof of concept" \
 --display-name="Allotrope POC"

export GSA_EMAIL=$USER-allotrope@${PROJECT_ID}.iam.gserviceaccount.com

# Get the workload identity pool for the gke/kcc cluster
WORKLOAD_IDENTITY_POOL="${PROJECT_ID}.svc.id.goog"

# grant workload identity bindings permissions
export ACK_NAMESPACE=ack-system    # Don’t change
export ACK_KSA_NAME=ack-controller # Don’t change
gcloud iam service-accounts add-iam-policy-binding ${GSA_EMAIL} \
 --role roles/iam.workloadIdentityUser \
 --member "serviceAccount:${WORKLOAD_IDENTITY_POOL}[${ASO_NAMESPACE}/${ASO_KSA}]" \
 --condition None
gcloud iam service-accounts add-iam-policy-binding ${GSA_EMAIL} \
 --role roles/iam.workloadIdentityUser \
 --member "serviceAccount:${WORKLOAD_IDENTITY_POOL}[${ACK_NAMESPACE}/${ACK_KSA_NAME}]" \
 --condition None
```

## Create an Azure managed identity

```
# Please change the parameters in this session
export MI_RESOURCE_GROUP="$USER"
export MI_NAME="$USER-aso-mi"
AZURE_SUBSCRIPTION_ID=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
AZURE_TENANT_ID=XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
AZURE_REGION=eastus

# Create a resource group
az group create -l ${AZURE_REGION} -n ${MI_RESOURCE_GROUP}

# Create a MI
az identity create --name ${MI_NAME} --resource-group ${MI_RESOURCE_GROUP} --location ${AZURE_REGION}

# Get the MI ID
AZURE_CLIENT_ID=$(az identity show --ids /subscriptions/${AZURE_SUBSCRIPTION_ID}/resourcegroups/${MI_RESOURCE_GROUP}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/${MI_NAME} --query "clientId" -otsv)
MI_PRINCIPAL_ID=$(az identity show --ids /subscriptions/${AZURE_SUBSCRIPTION_ID}/resourcegroups/${MI_RESOURCE_GROUP}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/${MI_NAME} --query "principalId" -otsv)

# Assign the permissions to this MI.
# User can use other permissions to manage their resources.
az role assignment create --assignee $MI_PRINCIPAL_ID --role contributor --scope /subscriptions/$AZURE_SUBSCRIPTION_ID

# Allow GCP service account to inpersonate this MI
az identity federated-credential create \
  --name gsa-azure-federated-credential \
  --identity-name ${MI_NAME} \
  --resource-group ${MI_RESOURCE_GROUP} \
  --issuer https://accounts.google.com \
  --subject ${GSA_UNIQUE_ID} \
  --audience api://AzureADTokenExchange
```

## Create ASO controller with the GCP service account.

```
cat <<EOF > /tmp/aso.yaml
apiVersion: kontrollers.cnrm.cloud.google.com/v1alpha1
kind: ASOKontroller
metadata:
 name: asokontroller.kontrollers.cnrm.cloud.google.com
spec:
 defaultAzureSubscriptionID: "$AZURE_SUBSCRIPTION_ID"
 defaultAzureTenentID: "$AZURE_TENANT_ID"
 defaultAzureClientID: "$AZURE_CLIENT_ID"
 crdPatterns:
 - "resources.azure.com/resourcegroup" # Use any pattern that fits your need.
 - "containerservice.azure.com/*"
 googleServiceAccount: "${GSA_EMAIL}"
EOF

kubectl apply -f /tmp/aso.yaml
```