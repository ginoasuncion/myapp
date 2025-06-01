pipeline {
    agent any

    tools {
        go "1.24.1"
    }

    stages {
        stage('Unit Test') {
            steps {
                sh "go test -v ./..."
            }
        }

        stage('Build Binary') {
            steps {
                sh "go build -o main main.go"
            }
        }

        stage('Docker Build') {
            steps {
                sh "docker build -t ttl.sh/myapp:1h ."
            }
        }

        stage('Docker Push') {
            steps {
                sh "docker push ttl.sh/myapp:1h"
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                withKubeConfig([credentialsId: 'kubernetes-token', serverUrl: 'https://k8s:6443']) {
                    sh "kubectl apply -f k8s/myapp-deployment.yaml"
                }
            }
        }
    }
}

