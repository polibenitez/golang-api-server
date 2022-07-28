resource "aws_ecs_cluster" "cluster" {
  name = "go-api-ecs-cluster"

  capacity_providers = ["FARGATE_SPOT", "FARGATE"]

  default_capacity_provider_strategy {
    capacity_provider = "FARGATE_SPOT"
  }

  setting {
    name  = "containerInsights"
    value = "disabled"
  }
}

module "ecs-fargate" {
  source  = "umotif-public/ecs-fargate/aws"
  version = "~> 6.5.1"

  name_prefix        = "go-api-ecs"
  vpc_id             = aws_vpc.ecs_vpc.id
  private_subnet_ids = [aws_subnet.private_subnet_1.id, aws_subnet.private_subnet_2.id]

  cluster_id = aws_ecs_cluster.cluster.id

  task_container_image   = "go-api"
  task_definition_cpu    = 256
  task_definition_memory = 512

  task_container_port             = 80
  task_container_assign_public_ip = true

  load_balanced = false

  target_groups = [
    {
      target_group_name = "tg-fargate-go-api"
      container_port    = 9999
    }
  ]

  health_check = {
    port = "9999"
    path = "/"
  }

  tags = {
    Environment = "test"
    Project     = "go-api"
  }
}