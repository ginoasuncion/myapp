pipeline {
    agent any

    environment {
        GO_VERSION = '1.24.1'
        DOCKER_IMAGE = 'ginoasuncion/myapp:latest'
        EC2_IP = '3.79.245.226'
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

        stage('Deploy to EC2') {
            steps {
                sshagent (credentials: ['ec2-deploy-key']) {
                    sh '''
                    ssh -o StrictHostKeyChecking=no ec2-user@$EC2_IP <<EOF
                      sudo yum install -y docker
                      sudo systemctl start docker
                      sudo docker pull $DOCKER_IMAGE
                      sudo docker stop myapp || true
                      sudo docker rm myapp || true
                      sudo docker run -d --name myapp -p 4444:8080 $DOCKER_IMAGE
                    EOF
                    '''
                }
            }
        }
    }
}

