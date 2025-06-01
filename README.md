# MyApp - Kubernetes Deployment with Jenkins

This project is a simple Go application deployed to a Kubernetes cluster using a Jenkins pipeline.

## Project Overview

- Language: Go
- Containerized with Docker
- CI/CD pipeline via Jenkins
- Deployed to Kubernetes using `kubectl`

## Pipeline Stages

1. **Test** – Runs unit tests
2. **Build** – Builds Go binary
3. **Docker Build** – Builds Docker image
4. **Docker Push** – Pushes image to Docker Hub
5. **Deploy to Kubernetes** – Applies deployment YAML via `kubectl`

## Prerequisites

- Jenkins configured with:
  - Go and Docker
  - DockerHub credentials (`dockerhub-creds`)
  - Kubernetes token credentials (`kubernetes-token`)
- Kubernetes cluster with access at `https://k8s:6443`

## Deployment Steps

1. **Docker Image**:  
   Pushed as `ginoasuncion/myapp:latest`

2. **Kubernetes YAML**:  
   Deployment defined in `myapp-deployment.yaml`

3. **Apply Deployment**:

```bash
kubectl apply -f myapp-deployment.yaml

