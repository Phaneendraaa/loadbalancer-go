﻿Custom Load Balancer in Golang for AWS EC2 with Health Monitoring and Auto-Scaling
Built a Golang-based custom load balancing solution to distribute traffic across AWS EC2 instances hosting a web application. The system monitors instance health, automatically replacing unhealthy instances by terminating and launching new ones—without relying on AWS-managed services like ELB or Auto Scaling, allowing for high customization.

Key Features:
Developed traffic distribution and health check mechanisms entirely in Golang.

Implemented auto-replacement of unhealthy instances using Golang and AWS CLI.

Leveraged Golang’s concurrency model (goroutines) to efficiently handle parallel health checks.

Ensured high availability by automatically managing instance lifecycle without external services.

Technologies Used:
Golang

AWS EC2

AWS CLI

Skills:
Amazon Web Services (AWS)

Go (Programming Language)

