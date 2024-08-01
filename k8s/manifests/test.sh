#!/bin/bash

# Define the cluster name and region
CLUSTER_NAME="golang-demo-eks-RhdefV3D"
REGION="eu-west-2"

# Get security group IDs associated with the EKS cluster
SG_IDS=$(aws eks describe-cluster --name $CLUSTER_NAME --region $REGION --query "cluster.resourcesVpcConfig.securityGroupIds" --output text)

# Loop through each security group and update rules
for SG_ID in $SG_IDS; do
  echo "Updating security group: $SG_ID"

  # Allow all inbound traffic
  aws ec2 authorize-security-group-ingress --group-id $SG_ID --protocol all --port all --cidr 0.0.0.0/0 || echo "Ingress rule already exists or failed to add."

  # Allow all outbound traffic
  aws ec2 authorize-security-group-egress --group-id $SG_ID --protocol all --port all --cidr 0.0.0.0/0 || echo "Egress rule already exists or failed to add."
done
