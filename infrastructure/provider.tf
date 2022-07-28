terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "2.19.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.22.0"
    }
  }
}

provider "docker" {}

provider "aws" {
  profile    = var.aws_profile
  region = var.region
}