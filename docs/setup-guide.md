# Complete Setup Guide

This guide will walk you through deploying the AWS Video Processing Platform from scratch, with screenshots and detailed explanations.

## 📋 Table of Contents

1. [Prerequisites](#prerequisites)
2. [Phase 1: Basic Infrastructure](#phase-1-basic-infrastructure)
3. [Phase 2: Video Processing Pipeline](#phase-2-video-processing-pipeline)
4. [Phase 3: Production Enhancements](#phase-3-production-enhancements)
5. [Cleanup](#cleanup)

## Prerequisites

### Required Tools
- [ ] AWS CLI configured with appropriate permissions
- [ ] Terraform installed (version >= 1.0)
- [ ] Go installed (version >= 1.19)
- [ ] Git for version control

### AWS Permissions Required
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:*",
        "s3:*",
        "lambda:*",
        "iam:*",
        "dynamodb:*",
        "logs:*"
      ],
      "Resource": "*"
    }
  ]
}
```

## Phase 1: Basic Infrastructure

### Step 1: Clean Slate
If you have existing infrastructure, destroy it first:

```bash
terraform destroy
```

**Screenshot**: Take a screenshot of the destroy confirmation

### Step 2: Project Structure Setup

Create the enhanced project structure:

```bash
mkdir -p docs terraform lambda/video-processor
```

### Step 3: Enhanced Terraform Configuration

We'll rebuild with the video processing architecture in mind.

**Screenshot**: Show the new project structure in your file explorer

### Step 4: Deploy Base Infrastructure

```bash
cd terraform
terraform init
terraform plan
terraform apply
```

**Screenshots to capture**:
- [ ] `terraform plan` output
- [ ] `terraform apply` confirmation
- [ ] AWS Console showing created resources
- [ ] EC2 instance running
- [ ] S3 buckets created

### Step 5: Deploy Go Application

```bash
# SSH to EC2 instance
ssh -i go-server-key.pem ec2-user@<instance-ip>

# Install dependencies
sudo yum update -y
sudo yum install -y golang git

# Deploy application
# (detailed steps here)
```

**Screenshots to capture**:
- [ ] SSH connection successful
- [ ] Go installation
- [ ] Application running
- [ ] Health check response

## Phase 2: Video Processing Pipeline

### Step 6: S3 Integration

Add S3 upload functionality to the Go application.

**Screenshots to capture**:
- [ ] S3 buckets in AWS Console
- [ ] Upload endpoint testing
- [ ] Files appearing in S3

### Step 7: Lambda Function

Deploy the video processing Lambda function.

**Screenshots to capture**:
- [ ] Lambda function in AWS Console
- [ ] Test execution
- [ ] CloudWatch logs

### Step 8: End-to-End Testing

Test the complete video processing pipeline.

**Screenshots to capture**:
- [ ] Video upload
- [ ] Processing status
- [ ] Converted videos in S3
- [ ] Download different resolutions

## Phase 3: Production Enhancements

### Step 9: Database Integration

Add DynamoDB for video metadata.

### Step 10: Monitoring Setup

Configure CloudWatch dashboards and alarms.

### Step 11: Security Hardening

Implement proper IAM roles and security groups.

## Testing Checklist

- [ ] Health endpoint responds
- [ ] Video upload works
- [ ] Lambda processes videos
- [ ] Multiple resolutions available
- [ ] Download endpoints work
- [ ] Error handling works
- [ ] Monitoring captures metrics

## Cleanup

To avoid AWS charges:

```bash
terraform destroy
```

**Screenshot**: Confirmation that all resources are destroyed

## Troubleshooting

### Common Issues

1. **Terraform Permission Errors**
   - Solution: Check IAM permissions

2. **Lambda Timeout**
   - Solution: Increase timeout for large videos

3. **S3 Upload Fails**
   - Solution: Check CORS configuration

## Cost Monitoring

Keep track of AWS costs during development:
- Set up billing alerts
- Monitor S3 storage usage
- Track Lambda execution time

---

**Next**: [API Reference](api-reference.md)