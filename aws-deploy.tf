provider "aws" {
  region     = "us-east-1"
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

resource "aws_instance" "example" {
  ami           = "ami-0778521d914d23bc1"
  instance_type = "t2.medium"
  key_name      = var.key_pair_name

  vpc_security_group_ids = [aws_security_group.example.id]

  connection {
    type        = "ssh"
    user        = "ubuntu"
    host        = self.public_dns
    private_key = file("${var.key_pair_name}.pem")
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
    # sources = [for s in fileset("./mail-client", "**") : s if !can(regex("/node_modules/|/dist/|/.git/|/data/", "./${s}"))]
    source      = "./mail-client"
    destination = "/tmp/mail-client"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt-get update",
      "sudo mkdir -m 0755 -p /etc/apt/keyrings",
      "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg",
      "echo deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null",
      "sudo apt-get update",
      "sudo chmod a+r /etc/apt/keyrings/docker.gpg",
      "sudo apt-get update",
      "sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y",
      "sudo service docker start",
      "sudo mkdir /usr/local/shared/data",
      "echo 'ZS_USER=${var.zs_user}' > /tmp/.env",
      "echo 'ZS_PASSWORD=${var.zs_pass}' >> /tmp/.env",
      "sudo docker compose --env-file /tmp/.env -f /tmp/docker-compose.yml up -d"
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
    from_port   = 22
    to_port     = 22
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

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
