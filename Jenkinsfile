pipeline {
    agent any

    tools {
        go "1.24.1"
    }

    environment {
        DEPLOY_USER = 'laborant'
        DEPLOY_HOST = '172.16.0.3'
        SSH_KEY = '/home/laborant/.ssh/id_rsa'
    }

    stages {
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o main main.go'
                sh 'ls -l main'
            }
        }

        stage('Add SSH Fingerprint') {
            steps {
                sh 'mkdir -p ~/.ssh && ssh-keyscan -H $DEPLOY_HOST >> ~/.ssh/known_hosts'
            }
        }

        stage('Check SSH Access') {
            steps {
                sh 'ssh -i $SSH_KEY $DEPLOY_USER@$DEPLOY_HOST "echo SSH OK"'
            }
        }

        stage('Deploy') {
            steps {
                sh 'scp -i $SSH_KEY main $DEPLOY_USER@$DEPLOY_HOST:~'
            }
        }
    }
}
