variable "aws_access_key" {
  type        = string
  description = "AWS Access Key"
  sensitive   = true
}

variable "aws_secret_key" {
  type        = string
  description = "AWS Secret Key"
  sensitive   = true
}

variable "key_pair_name" {
  type        = string
  description = "AWS Key pair"
  sensitive   = true
  default     = "personalmac"
}

variable "zs_user" {
  type        = string
  description = "Zinc search user"
  sensitive   = true
}

variable "zs_pass" {
  type        = string
  description = "Zinc search password"
  sensitive   = true
}
