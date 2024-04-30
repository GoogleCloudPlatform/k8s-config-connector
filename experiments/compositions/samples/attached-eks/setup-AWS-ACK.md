# Setup AWS Controllers for Kubernetes (ACK) 

Note: The ACK support is only availabe for GCP managed [Config Connector](https://cloud.google.com/config-connector/docs/overview).

Here is an overview of the ACK setup process. 

1. Create a GCP service account (or use the default service account).
1. Create a AWS role.
   1. Grant this role the permissions to manage the AWS resouces.
   1. Allow the GCP service account to inpersonate the AWS role.
1. Create ACK controller(s) with the GCP service account.

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

## Create AWS role and allow GCP service account to inpersonate it

```
# Here we use AdministratorAccess policy to manage AWS resouces.
# User can use other policy to manage the resouces.
AWS_POLICY=arn:aws:iam::aws:policy/AdministratorAccess 

# Get the subject of the GSA
GSA_SUB=$(gcloud iam service-accounts describe ${GSA_EMAIL}  \
 --format "value(oauth2ClientId)")
export AWS_ROLE_NAME=$USER-ack-role

# Create AWS role and policy
cat > /tmp/trust-policy.json << EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Effect": "Allow",
     "Principal": {
       "Federated": "accounts.google.com"
     },
     "Action": "sts:AssumeRoleWithWebIdentity",
     "Condition": {
       "StringEquals": {
         "accounts.google.com:oaud": "sts.amazonaws.com",
         "accounts.google.com:aud": "${GSA_SUB}",
         "accounts.google.com:sub": "${GSA_SUB}"
       }
     }
   }
 ]
}
EOF
aws iam create-role \
 --role-name ${AWS_ROLE_NAME} \
 --assume-role-policy-document file:#/tmp/trust-policy.json

aws iam attach-role-policy \
 --role-name ${AWS_ROLE_NAME} \
 --policy-arn ${AWS_POLICY}

AWS_ROLE_ARN=$(aws iam get-role --role-name ${AWS_ROLE_NAME} | jq -r .Role.Arn)
```


## Create ACK controller(s) with the GCP service account.

We only support part of the ACK controllers. 
User can create one or more ACK controllers of
- ec2-controller
- eks-controller
- iam-controller

Please use this commands to create the ACK controllers.

```
AWS_REGION=us-west-2 # or other regions

cat <<EOF > /tmp/ack-config.yaml
apiVersion: kontrollers.cnrm.cloud.google.com/v1alpha1
kind: ACKKontroller
metadata:
 name: ackkontroller.kontrollers.cnrm.cloud.google.com
spec:
 defaultAWSRegion: "${AWS_REGION}"
 defaultRoleARN: "${AWS_ROLE_ARN}"
 controllers:
 - "ec2-controller" # Choose one or more controllers.
 - "eks-controller" # Choose one or more controllers.
 - "iam-controller" # Choose one or more controllers.
 googleServiceAccount: "${GSA_EMAIL}"
EOF

kubectl apply -f /tmp/ack-config.yaml
```