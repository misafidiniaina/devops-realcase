resource "aws_security_group" "load_balancer" {
  name        = "${var.name}-load-balancer"
  description = "Public HTTP and HTTPS access for the application load balancer."
  vpc_id      = var.vpc_id

  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "HTTPS"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "All outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = merge(var.tags, {
    Name = "${var.name}-load-balancer"
  })
}

resource "aws_security_group" "application" {
  name        = "${var.name}-application"
  description = "Private access to the application container port."
  vpc_id      = var.vpc_id

  ingress {
    description     = "Application traffic from the load balancer"
    from_port       = var.app_port
    to_port         = var.app_port
    protocol        = "tcp"
    security_groups = [aws_security_group.load_balancer.id]
  }

  egress {
    description = "All outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = merge(var.tags, {
    Name = "${var.name}-application"
  })
}

resource "aws_security_group" "database" {
  name        = "${var.name}-database"
  description = "PostgreSQL access from the application security group only."
  vpc_id      = var.vpc_id

  ingress {
    description     = "PostgreSQL from the application"
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.application.id]
  }

  egress {
    description = "All outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = merge(var.tags, {
    Name = "${var.name}-database"
  })
}
