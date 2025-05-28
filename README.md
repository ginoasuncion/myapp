# myapp

A simple Go web app that returns a JSON response. Built with Docker and deployed via Docker Hub.

## 🛠 Tech

- Go 1.22
- Docker (multi-stage build)
- Docker Hub: [ginoasuncion/myapp](https://hub.docker.com/r/ginoasuncion/myapp)

## 🚀 Build & Run (Locally)

```bash
# Build Docker image
docker build -t ginoasuncion/myapp .

# Run locally
docker run --rm -p 4444:4444 ginoasuncion/myapp

# Test
curl http://localhost:4444
```

## ☁️ Deploy to Docker VM

```bash
docker pull ginoasuncion/myapp
docker run -d --rm -p 4444:4444 ginoasuncion/myapp
curl http://<docker-vm-ip>:4444
```

## 📄 Response

```json
{
  "greeting": "Hello",
  "subject": "World",
  "host": "localhost:4444"
}
```

## 📁 Structure

```
main.go          # Go server
main_test.go     # Test
Dockerfile       # Multi-stage build
README.md        # This file
```

