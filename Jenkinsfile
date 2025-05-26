pipeline {
    agent any

    tools {
        go "1.24.1"
    }

    environment {
        DEPLOY_USER = 'laborant'
        DEPLOY_HOST = '172.16.0.3'
        SSH_KEY = '/var/lib/jenkins/.ssh/id_rsa' // ✅ use full path for Jenkins
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

        stage('Deploy Binary') {
            steps {
                sh 'scp -i $SSH_KEY main $DEPLOY_USER@$DEPLOY_HOST:~/main'
            }
        }

        stage('Deploy Service File') {
            steps {
                sh 'scp -i $SSH_KEY main.service $DEPLOY_USER@$DEPLOY_HOST:~/main.service'
                sh """
                ssh -i $SSH_KEY $DEPLOY_USER@$DEPLOY_HOST <<EOF
                    sudo mv ~/main.service /etc/systemd/system/main.service
                    sudo chmod 644 /etc/systemd/system/main.service
                    sudo systemctl daemon-reload
                    sudo systemctl enable main.service
                    sudo systemctl restart main.service
                EOF
                """
            }
        }
    }
}
