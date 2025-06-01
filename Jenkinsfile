pipeline {
    agent any

    environment {
        GO_VERSION = '1.24.1'
        DOCKER_IMAGE = 'ginoasuncion/myapp:latest'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Tool Install') {
            steps {
                script {
                    def go = tool name: "${GO_VERSION}", type: 'go'
                    env.GOROOT = go
                    env.PATH = "${go}/bin:${env.PATH}"
                }
            }
        }

        stage('Unit Test') {
            steps {
                sh 'go test -v ./...'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o main main.go'
            }
        }

        stage('Docker Login') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                }
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE .'
            }
        }

        stage('Docker Push') {
            steps {
                sh 'docker push $DOCKER_IMAGE'
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                withKubeConfig([credentialsId: 'kubernetes-token', serverUrl: 'https://k8s:6443']) {
                    sh 'kubectl apply -f myapp-deployment.yaml'
                }
            }
        }
    }
}

