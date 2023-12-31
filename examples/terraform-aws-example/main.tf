terraform {
  required_version = ">= 0.12.26"
}
provider "aws" {
  region = "us-east-2"
}

resource "aws_instance" "example" {
  ami = "ami-0d5d9d301c853a04a"
  instance_type = "t2.micro"
  vpc_security_group_ids = [aws_security_group.instance.id]
  user_data = <<EOF
#!/bin/bash
echo "Hello, FnitFnit!" > index.html
nohup busybox httpd -f -p 8080 &
EOF
}

resource "aws_security_group" "instance" {
  ingress {
    from_port = 8080
    to_port = 8080
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "public_ip" {
  value = aws_instance.example.public_ip
}
