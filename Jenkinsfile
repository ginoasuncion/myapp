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
                withCredentials([sshUserPrivateKey(
                    credentialsId: 'ec2-deploy-key',
                    keyFileVariable: 'KEY_FILE',
                    usernameVariable: 'SSH_USER'
                )]) {
                    sh '''
                        chmod 600 $KEY_FILE

                        # Copy binary and service file
                        scp -o StrictHostKeyChecking=no -i $KEY_FILE main $SSH_USER@$EC2_IP:/home/ec2-user/
                        scp -o StrictHostKeyChecking=no -i $KEY_FILE main.service $SSH_USER@$EC2_IP:/tmp/

                        # Install and start service
                        ssh -o StrictHostKeyChecking=no -i $KEY_FILE $SSH_USER@$EC2_IP '
                            sudo mv /tmp/main.service /etc/systemd/system/main.service &&
                            sudo chmod +x /home/ec2-user/main &&
                            sudo systemctl daemon-reexec &&
                            sudo systemctl daemon-reload &&
                            sudo systemctl restart main.service
                        '
                    '''
                }
            }
        }
    }
}
