terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
    }
    tls = {
      source  = "hashicorp/tls"
    }
    local = {
      source  = "hashicorp/local"
    }
  }
}


provider "aws" {
    region = "us-east-1"
}

resource "aws_vpc" "main" {
    cidr_block = "10.0.0.0/16"
    enable_dns_support = true
    enable_dns_hostnames = true

    tags = {
        Name = "GoServerVPC"    
    }
}


resource "aws_subnet" "public" {
  vpc_id = aws_vpc.main.id
  cidr_block = "10.0.0.0/24"
  availability_zone = "us-east-1a"
  map_public_ip_on_launch = true
    
    tags = {
        Name = "GoServerPublicSubnet"
    }
}

resource "aws_internet_gateway" "igw" {
    vpc_id = aws_vpc.main.id

    tags = {
        Name = "GoServerInternetGateway"
    }
}

resource "aws_route_table" "table" {
    vpc_id = aws_vpc.main.id
    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.igw.id
    }

    tags = {
        Name = "GoServerRouteTable"
    }
}

resource "aws_route_table_association" "public_route" {
        subnet_id = aws_subnet.public.id
        route_table_id = aws_route_table.table.id
}

resource "aws_security_group" "web_sg" {
    name = "web_sg"
    description = "Allows SSH and HTTP access"

    tags = {
        Name = "GoServerSecurityGroup"
    }

    vpc_id = aws_vpc.main.id

    ingress {
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["98.109.109.51/32"]
    }

    ingress {
        from_port = 3000
        to_port = 3000
        protocol = "tcp"
        cidr_blocks = ["98.109.109.51/32"]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}

resource "tls_private_key" "go_server_private_key" {
    algorithm = "RSA"
    rsa_bits = 2048
}

resource "aws_key_pair" "go_server_key"{
    key_name = "go-server-key"
    public_key = tls_private_key.go_server_private_key.public_key_openssh
}

resource "local_file" "go_server_private_key_file" {
    content = tls_private_key.go_server_private_key.private_key_pem
    filename = "go-server-key.pem"
}

resource "aws_instance" "go_server" {
    ami = "ami-0de716d6197524dd9"
    instance_type = "t2.micro"
    subnet_id = aws_subnet.public.id
    vpc_security_group_ids = [aws_security_group.web_sg.id]
    key_name = aws_key_pair.go_server_key.key_name

    tags = {
        Name = "SimpleGoWebServer"
    }
}

output "instance_public_ip" {
  description = "Public IP address of the EC2 instance"
  value       = aws_instance.go_server.public_ip
}
