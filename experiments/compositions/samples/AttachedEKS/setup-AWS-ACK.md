# Set up AWS Controllers for Kubernetes (ACK)

Note: ACK support is only available for GCP managed [Config
Connector](https://cloud.google.com/config-connector/docs/overview).

Here is an overview of the ACK setup process:

1. Create a GCP service account.
1. Create an AWS role.
   1. Grant this role the permissions to manage the AWS resouces.
   1. Allow the GCP service account to impersonate the AWS role.
1. Create ACK controller(s) with the GCP service account.

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
export ACK_NAMESPACE=kontrollers-ack-system # Don’t change
export ACK_KSA_NAME=ack-controller # Don’t change
gcloud iam service-accounts add-iam-policy-binding ${GSA_EMAIL} \
 --role roles/iam.workloadIdentityUser \
 --member "serviceAccount:${WORKLOAD_IDENTITY_POOL}[${ACK_NAMESPACE}/${ACK_KSA_NAME}]" \
 --condition None
```

## Create an AWS role and allow the GCP service account to impersonate it

```
# In this example we use the AdministratorAccess policy to manage AWS resouces.
# The user can use a different policy to manage the resouces.
AWS_POLICY=arn:aws:iam::aws:policy/AdministratorAccess

# Get the subject of the GSA
GSA_SUB=$(gcloud iam service-accounts describe ${GSA_EMAIL} \
  --format "value(oauth2ClientId)")
export AWS_ROLE_NAME=${USER}-ack-role

# Create an AWS role and policy
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
 --assume-role-policy-document file:///tmp/trust-policy.json

aws iam attach-role-policy \
 --role-name ${AWS_ROLE_NAME} \
 --policy-arn ${AWS_POLICY}

AWS_ROLE_ARN=$(aws iam get-role --role-name ${AWS_ROLE_NAME} | jq -r .Role.Arn)
```


## Create ACK controller(s) with the GCP service account

We only support a subset of the ACK controllers. The user can create one or more
of the following ACK controllers:
- ec2-controller
- eks-controller
- iam-controller

Use the following commands to create the ACK controller(s):

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
