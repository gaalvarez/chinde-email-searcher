# CHINDE MAIL SEARCHER

Chinde is a project created to use the [ZincSearch](https://docs.zinc.dev/) search engine, with a set of Enron Corp emails. Three projects have been created to allow filtering and searching for emails.

- mail-client: frontend client application -> more details [aqui](./mail-client/README.md)
- mail-backend: backend server application -> more details [aqui](./mail-backend/README.md)
- mail-indexer: application to index the content -> more details [aqui](./mail-indexer/README.md)

## Instalation

### Clone this repository

```sh
git clone https://github.com/gaalvarez/chinde-email-searcher
```

### Create Access Keys on AWS

- Input IAM in search input, and select IAM service
- Click on Manage access keys button
- Scroll to find Acces keys section
- Click on Create access key button
- Follow the instructions

### Create Key Pair (for ssh connection) on AWS

- Input key pairs in search input, and select key pairs feature
- Click on Create key pair button
- Input a name (remember this name)
- Select the options RSA and .pem options
- Click on Create key pair
- Put the generated file on root folder of code repository

### Install Terraform

[Follow this instructions](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)

### Export the values for terraform variables

The variables at available [here](./variables.tf).

And define user and password for Zinc Search API and UI.

```sh
export TF_VAR_aws_access_key="your_access_key"
export TF_VAR_aws_secret_key="your_secret_key"
export TF_VAR_aws_key_pair_name="your_key_pair_name"
export TF_VAR_zs_user="zincsearch user"
export TF_VAR_zs_pass="zincsearch password"
```

### Deploy using terraform

```sh
terraform init
terraform plan
terraform apply
```

### Identify publi IP

Open your AWS console, open the EC2 service and open the instance, copy the Public IPv4 DNS

### Access to app

Open your browser and paste <http://[your> public IPV4 DNS]

## Mail Indexation

Please follow [this instructions](./mail-indexer/README.md)
