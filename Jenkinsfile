pipeline {
    agent any

    tools {
        go "1.24.1"
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
            }
        }

        stage('Add SSH Fingerprint') {
            steps {
                withCredentials([
                    string(credentialsId: 'target-host', variable: 'DEPLOY_HOST')
                ]) {
                    sh 'mkdir -p ~/.ssh && ssh-keyscan -H $DEPLOY_HOST >> ~/.ssh/known_hosts'
                }
            }
        }

        stage('Check SSH Access') {
            steps {
                withCredentials([
                    sshUserPrivateKey(credentialsId: 'target-ssh-key', keyFileVariable: 'SSH_KEY', usernameVariable: 'DEPLOY_USER'),
                    string(credentialsId: 'target-host', variable: 'DEPLOY_HOST')
                ]) {
                    sh 'ssh -i $SSH_KEY $DEPLOY_USER@$DEPLOY_HOST "echo SSH OK"'
                }
            }
        }

        stage('Deploy Binary') {
            steps {
                withCredentials([
                    sshUserPrivateKey(credentialsId: 'target-ssh-key', keyFileVariable: 'SSH_KEY', usernameVariable: 'DEPLOY_USER'),
                    string(credentialsId: 'target-host', variable: 'DEPLOY_HOST')
                ]) {
                    sh 'scp -i $SSH_KEY main $DEPLOY_USER@$DEPLOY_HOST:/home/$DEPLOY_USER/main.tmp'
                    sh 'ssh -i $SSH_KEY $DEPLOY_USER@$DEPLOY_HOST "mv /home/$DEPLOY_USER/main.tmp /home/$DEPLOY_USER/main"'
                }
            }
        }

        stage('Deploy Service File') {
            steps {
                withCredentials([
                    sshUserPrivateKey(credentialsId: 'target-ssh-key', keyFileVariable: 'SSH_KEY', usernameVariable: 'DEPLOY_USER'),
                    string(credentialsId: 'target-host', variable: 'DEPLOY_HOST')
                ]) {
                    sh 'scp -i $SSH_KEY main.service $DEPLOY_USER@$DEPLOY_HOST:/home/$DEPLOY_USER/main.service'
                    sh """
                        ssh -i $SSH_KEY $DEPLOY_USER@$DEPLOY_HOST \\
                        'sudo mv /home/$DEPLOY_USER/main.service /etc/systemd/system/main.service && \\
                         sudo chmod 644 /etc/systemd/system/main.service && \\
                         sudo systemctl daemon-reload && \\
                         sudo systemctl enable --now main.service'
                    """
                }
            }
        }
    }
}
