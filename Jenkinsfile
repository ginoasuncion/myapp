pipeline {
    agent any

    environment {
        GO_VERSION = '1.24.1'
        DOCKER_IMAGE = 'ginoasuncion/myapp:latest'
        EC2_USER = 'ec2-user'
        EC2_HOST = '3.79.245.226'
        DEPLOY_KEY_ID = 'ec2-deploy-key'
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
                // Build a static Linux-compatible binary
                sh 'GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go'
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
                withCredentials([sshUserPrivateKey(
                    credentialsId: "${DEPLOY_KEY_ID}",
                    keyFileVariable: 'SSH_KEY',
                    usernameVariable: 'SSH_USER'
                )]) {
                    sh '''
                        scp -i $SSH_KEY -o StrictHostKeyChecking=no main $SSH_USER@$EC2_HOST:/home/ec2-user/main
                        ssh -i $SSH_KEY -o StrictHostKeyChecking=no $SSH_USER@$EC2_HOST 'sudo systemctl restart main.service && sudo systemctl status main.service'
                    '''
                }
            }
        }
    }
}
