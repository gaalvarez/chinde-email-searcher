provider "aws" {
  region = "us-east-1"
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

resource "aws_instance" "example" {
  ami           = "ami-0778521d914d23bc1"
  instance_type = "t2.micro"

  vpc_security_group_ids = [aws_security_group.example.id]

  connection {
    type        = "ssh"
    user        = "ec2-user"
    host        = "44.211.187.149"
    private_key = file("~/.ssh/id_rsa")
  }

  provisioner "file" {
    source      = "./docker-compose.yml"
    destination = "/tmp/docker-compose.yml"
  }

  provisioner "file" {
    source      = "./mail-backend"
    destination = "/tmp/mail-backend"
  }

  provisioner "file" {
    source      = "./mail-client"
    destination = "/tmp/mail-client"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum install -y docker",
      "sudo service docker start",
      "sudo docker-compose -f /tmp/docker-compose.yml up -d"
    ]
  }

}

resource "aws_security_group" "example" {
  name        = "example"
  description = "Allow HTTP and SSH traffic"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 4080
    to_port     = 4080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}