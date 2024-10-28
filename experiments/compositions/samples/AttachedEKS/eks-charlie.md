## Apply the composition
kubectl apply -f 01-composition.yaml


## Create a GCP service account for this team and grant KCC permissions

This step follows the following documentation
https://cloud.google.com/config-connector/docs/how-to/install-namespaced

export NAMESPACE=team-eks
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


## Create a namespace for Alice's team
cat 02-context.yaml | envsubst | \
  kubectl apply  -f -
