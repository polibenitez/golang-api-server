variable "region" {
  type        = string
  default     = "eu-west-1"
}

variable "profile" {
  type        = string
  default     = "poliBenitez"
}

# VPC variables

variable "vpc_cidr" {
  type        = string
  default     = "10.0.0.0/16"
}