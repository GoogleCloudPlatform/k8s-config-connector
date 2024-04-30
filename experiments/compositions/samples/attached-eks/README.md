# Create a EKS cluster and attach to the GCP

1. Follow the [setup-Azure-ASO.md](setup-Azure-ASO.md) to setup the ASO.
2. Act as a platform admin, run the [eks-charlie.sh](eks-charlie.sh) for setup the composition and a user team.
3. Act as a user, run the [eks-alice.sh](eks-alice.sh) to create the attached EKS cluster.