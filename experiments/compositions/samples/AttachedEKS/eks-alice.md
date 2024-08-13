## Step 1. Apply CR to create the composition of AKS and its attachment
kubectl apply -f 03-attached-1.yaml


## Step 2. Log in to the cluster and apply the attached cluster install manifest
EKS_NAME=$(cat 03-attached-1.yaml |yq .metadata.name)
ATTACHED_REGION=$(cat 03-attached-1.yaml |yq .spec.gcpRegion)
ATTACHED_PLATFORM_VERSION=$(cat 03-attached-1.yaml |yq .spec.attachedPlatformVersion)

gcloud container attached clusters generate-install-manifest \
  $EKS_NAME \
  --location=${ATTACHED_REGION} \
  --platform-version ${ATTACHED_PLATFORM_VERSION} \
  --output-file=/tmp/install-agent-${EKS_NAME}.yaml

aws eks update-kubeconfig --name $EKS_NAME-cluster

kubectl apply -f /tmp/install-agent-${EKS_NAME}.yaml


## Commands to check progress:

kubectl get AttachedEKS -n team-eks

kubectl get vpc.ec2.services.k8s.aws \
  -n team-eks

kubectl get InternetGateway.ec2.services.k8s.aws \
  -n team-eks

kubectl get RouteTable.ec2.services.k8s.aws \
  -n team-eks

kubectl get subnet.ec2.services.k8s.aws \
  -n team-eks

kubectl get ElasticIPAddress.ec2.services.k8s.aws \
  -n team-eks

kubectl get NATGateway.ec2.services.k8s.aws \
  -n team-eks

kubectl get role.iam.services.k8s.aws \
  -n team-eks

kubectl get cluster.eks.services.k8s.aws \
  -n team-eks

kubectl get Nodegroup.eks.services.k8s.aws \
  -n team-eks

kubectl get AccessEntry.eks.services.k8s.aws \
  -n team-eks

kubectl get FieldExport.services.k8s.aws \
  -n team-eks

kubectl get cm \
  -n team-eks

kubectl get containerattachedcluster \
  -n team-eks
