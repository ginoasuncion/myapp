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
        stage('Build') {
            steps {
                sh "go build -o main main.go"
            }
        }
        stage('Docker Login') {
            steps {
                withCredentials([string(credentialsId: 'dockerhub-pass', variable: 'DOCKER_PASS')]) {
                    sh '''
                        echo "$DOCKER_PASS" | docker login -u ginoasuncion --password-stdin
                    '''
                }
            }
        }
        stage('Docker Build') {
            steps {
                sh "docker build -t ginoasuncion/myapp:latest ."
            }
        }
        stage('Docker Push') {
            steps {
                sh "docker push ginoasuncion/myapp:latest"
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                withKubeConfig([credentialsId: 'kubernetes-token', serverUrl: 'https://k8s:6443']) {
                    sh "kubectl apply -f myapp-deployment.yaml"
                }
            }
        }
    }
}

