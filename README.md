# AWS Video Platform

🎬 A progressively built video platform using Go and AWS - learning cloud architecture through hands-on development.

## 📊 Current Status

**✅ What We Have Built:**
- Basic Go HTTP server with Chi router
- AWS EC2 instance with Terraform
- Custom VPC with security groups
- Health check endpoint (`/health`)
- Hello world endpoints (`/hello`, `/about`)
- JSON structured logging

**🎯 Vision - What We're Building Toward:**
A mini-YouTube style platform with video upload, processing, and streaming capabilities.

## 🏗️ Current Architecture

```
┌─────────────┐    ┌─────────────┐
│   Client    │───▶│  Go Server  │
│ (Browser)   │    │   (EC2)     │
└─────────────┘    └─────────────┘
```

**Current Endpoints:**
- `GET /health` - Health check
- `GET /hello` - Hello world
- `GET /about` - About page

## 🛠️ Technology Stack

### **Core Technologies**
- **Language**: Go 1.23.0
- **Web Framework**: Chi Router v5.2.2
- **Infrastructure**: Terraform (AWS Provider)

### **AWS Services (Current)**
- **Compute**: EC2 (t2.micro)
- **Networking**: Custom VPC, Public Subnet, Security Groups, Internet Gateway
- **Security**: SSH Key Pair, Restricted Security Group Rules

### **Planned AWS Services**
- **Storage**: S3 for video files
- **Processing**: Lambda with FFmpeg
- **Database**: DynamoDB for metadata
- **CDN**: CloudFront for streaming
- **Monitoring**: CloudWatch

### **Development Tools**
- **SSH Client**: PuTTY (Windows)
- **Version Control**: Git
- **Documentation**: Markdown

## 📁 Project Structure

```
aws-video/
├── docs/                    # Documentation
│   ├── setup-guide.md      # Complete deployment guide
│   └── api-reference.md    # API documentation
├── terraform/              # Infrastructure as Code
│   └── main.tf            # Terraform configuration
├── lambda/                 # Future: Lambda functions
│   └── video-processor/   # Future: Video processing
├── screenshots/           # Documentation screenshots
├── main.go               # Go HTTP server
├── go.mod               # Go dependencies
├── go.sum               # Go dependency checksums
└── README.md           # This file
```

## 🚀 Quick Start

1. **Clone the repository**
   ```bash
   git clone <your-repo>
   cd simple-http-server
   ```

2. **Follow the setup guide**
   - See [docs/setup-guide.md](docs/setup-guide.md) for detailed instructions

## 📖 Documentation

- [Setup Guide](docs/setup-guide.md) - Complete deployment walkthrough
- [API Reference](docs/api-reference.md) - Endpoint documentation
- [Architecture Guide](docs/architecture.md) - System design details

## 🔧 Development

### Local Development
```bash
go mod tidy
go run main.go
```

### Testing
```bash
curl http://localhost:3000/health
```

## 🌐 Deployment

This project uses Infrastructure as Code with Terraform. See the setup guide for complete deployment instructions.

## 💰 Cost Estimation

- **EC2 t2.micro**: Free tier eligible
- **S3 Storage**: ~$0.023/GB/month
- **Lambda**: Free tier: 1M requests/month
- **Data Transfer**: First 1GB free/month

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add documentation
5. Submit a pull request

## 📝 License

MIT License - see LICENSE file for details

## 🎯 Learning Progress

**✅ Currently Demonstrated:**
- Infrastructure as Code with Terraform
- AWS EC2 and VPC networking
- Go HTTP server development
- RESTful API endpoints
- JSON structured logging
- SSH key management

**🚧 Next Learning Goals:**
- S3 file storage integration
- Lambda serverless functions
- Video processing with FFmpeg
- DynamoDB for metadata
- CloudWatch monitoring

---

**Built with ❤️ for learning AWS**