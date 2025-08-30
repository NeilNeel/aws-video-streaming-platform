# AWS Video Platform

🎬 A progressively built video platform using Go and AWS - learning cloud architecture through hands-on development.

## 📊 Current Status (Updated January 2025)

**✅ What We Have Built:**

### Core Application
- ✅ Go HTTP server with Chi router v5.2.2
- ✅ RESTful API endpoints (`/health`, `/hello`, `/about`)
- ✅ JSON structured logging
- ✅ Comprehensive automated testing (`main_test.go`)

### Infrastructure (Terraform)
- ✅ Custom VPC with public subnet
- ✅ EC2 instance (t2.micro) with security groups
- ✅ IAM roles with S3 and SSM permissions
- ✅ S3 bucket for deployment artifacts
- ✅ SSH key pair management

### CI/CD Pipeline
- ✅ GitHub Actions workflow for build and test
- ✅ Automated Go application building
- ✅ Unit test execution on every push
- ✅ S3 artifact upload with commit SHA
- ✅ AWS credentials integration
- 🚧 EC2 deployment automation (in progress)

**🎯 Vision - What We're Building Toward:**
A mini-YouTube style platform with video upload, processing, and streaming capabilities.

## 🏗️ Current Architecture

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   GitHub    │───▶│GitHub Actions│───▶│     S3      │───▶│    EC2      │
│ Repository  │    │  (CI/CD)    │    │ Artifacts   │    │   Server    │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
       │                   │                   │                   │
       ▼                   ▼                   ▼                   ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│    Code     │    │Build & Test │    │   Binary    │    │   Running   │
│   Changes   │    │ Automation  │    │   Storage   │    │ Application │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
```

**Current Endpoints:**
- `GET /health` - Health check with JSON response
- `GET /hello` - Parameterized greeting
- `GET /about` - About page

## 🛠️ Technology Stack

### **Core Technologies**
- **Language**: Go 1.23.0
- **Web Framework**: Chi Router v5.2.2
- **Infrastructure**: Terraform (AWS Provider)
- **CI/CD**: GitHub Actions

### **AWS Services (Current)**
- **Compute**: EC2 (t2.micro) with IAM roles
- **Storage**: S3 for deployment artifacts
- **Management**: Systems Manager (SSM) for deployment
- **Networking**: Custom VPC, Public Subnet, Security Groups, Internet Gateway
- **Security**: IAM roles, SSH Key Pair, Restricted Security Group Rules

### **Planned AWS Services**
- **Storage**: S3 for video files (user uploads)
- **Processing**: Lambda with FFmpeg for transcoding
- **Database**: DynamoDB for video metadata
- **CDN**: CloudFront for video streaming
- **Monitoring**: CloudWatch for observability

### **Development Tools**
- **Version Control**: Git with GitHub
- **Testing**: Go testing framework
- **Documentation**: Markdown
- **Infrastructure as Code**: Terraform

## 📁 Project Structure

```
aws-video/
├── .github/                 # GitHub Actions workflows
│   └── workflows/
│       ├── build.yml       # CI/CD pipeline (build, test, deploy)
│       └── deploy.yml      # Deployment workflow (in progress)
├── docs/                   # Documentation
│   ├── setup-guide.md     # Complete deployment guide
│   └── api-reference.md   # API documentation
├── terraform/             # Infrastructure as Code
│   └── main.tf           # Complete AWS infrastructure
├── lambda/               # Future: Lambda functions
│   └── video-processor/  # Future: Video processing
├── screenshots/         # Documentation screenshots
├── main.go             # Go HTTP server
├── main_test.go        # Automated tests
├── go.mod             # Go dependencies
├── go.sum             # Go dependency checksums
└── README.md         # This file
```

## 🚀 Quick Start

### Prerequisites
- Go 1.23.0+
- Terraform
- AWS CLI configured
- GitHub account

### Local Development
```bash
# Clone and setup
git clone https://github.com/NeilNeel/aws-video-streaming-platform
cd aws-video

# Install dependencies
go mod tidy

# Run tests
go test -v

# Run locally
go run main.go
```

### Infrastructure Deployment
```bash
cd terraform
terraform init
terraform plan
terraform apply
```

## 🔄 CI/CD Pipeline

### Automated Workflow
```
Push to main → Build Go App → Run Tests → Upload to S3 → Deploy to EC2
```

### Current Capabilities
- ✅ **Continuous Integration**: Automated building and testing
- ✅ **Artifact Management**: S3 storage with commit SHA versioning
- ✅ **Quality Gates**: Tests must pass before deployment
- ✅ **Continuous Deployment**: EC2 deployment via SSM (WORKING!)

### GitHub Actions Workflows
- **build.yml**: Complete CI pipeline with S3 upload
- **deploy.yml**: EC2 deployment via Systems Manager (WORKING!)

## 🌐 Live Application

**Current deployment is LIVE and working!**
- **URL**: http://3.83.108.9:3000
- **Health Check**: http://3.83.108.9:3000/health
- **Test Endpoints**:
  ```bash
  curl http://3.83.108.9:3000/health
  curl http://3.83.108.9:3000/hello?name=Student
  curl http://3.83.108.9:3000/about
  curl http://3.83.108.9:3000/version
  ```

## 🧪 Testing

### Automated Testing
- **Unit Tests**: HTTP endpoint validation
- **Integration Tests**: Full request/response cycle testing
- **CI Integration**: Tests run on every push

### Test Coverage
- ✅ Health endpoint (status code, content-type, response body)
- ✅ Hello endpoint (with and without parameters)
- 🚧 Future: S3 upload endpoint testing

## 🔧 Current Development Status

### ✅ Completed Features
1. **Core Web Server**: Fully functional Go HTTP server
2. **Infrastructure**: Complete AWS setup with Terraform
3. **CI Pipeline**: Automated build, test, and artifact storage
4. **Quality Assurance**: Comprehensive testing framework

### ✅ Recently Completed
1. **Full CI/CD Pipeline**: Complete automated deployment to EC2
2. **SSM Integration**: Systems Manager successfully deploying applications
3. **Security Configuration**: Proper IAM roles and security groups

### 📋 Next Milestones
1. **Video Upload**: Add S3 file upload endpoints
2. **Video Processing**: Lambda-based transcoding with FFmpeg
3. **Video Streaming**: CloudFront CDN integration
4. **Multi-resolution**: 360p and 720p video processing

## 💰 Cost Estimation

- **EC2 t2.micro**: Free tier eligible
- **S3 Storage**: ~$0.023/GB/month
- **Lambda**: Free tier: 1M requests/month
- **Data Transfer**: First 1GB free/month
- **Systems Manager**: No additional cost

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes with tests
4. Ensure CI pipeline passes
5. Submit a pull request

## 📝 License

MIT License - see LICENSE file for details

## 🎯 Learning Progress

**✅ Currently Demonstrated:**
- Infrastructure as Code with Terraform
- AWS EC2, VPC, IAM, S3, and SSM integration
- Go HTTP server development with Chi router
- RESTful API design and implementation
- Automated testing with Go testing framework
- GitHub Actions CI/CD pipeline
- AWS SDK integration
- Professional development workflow

**🚧 Next Learning Goals:**
- Complete automated deployment pipeline
- S3 file upload integration
- Lambda serverless functions
- Video processing with FFmpeg
- DynamoDB for metadata storage
- CloudWatch monitoring and logging

---

**Built with ❤️ for learning AWS and modern DevOps practices**

**Last Updated**: January 29, 2025 - Full CI/CD Pipeline Working! 🎉
**Repository**: https://github.com/NeilNeel/aws-video-streaming-platform