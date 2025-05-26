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
                sh 'ls -l main'
            }
        }

        stage('Add SSH Fingerprint') {
            steps {
                sh 'mkdir -p ~/.ssh && ssh-keyscan -H target >> ~/.ssh/known_hosts'
            }
        }

        stage('Check SSH Access') {
            steps {
                sh 'ssh -i /var/lib/jenkins/.ssh/id_rsa laborant@target "echo SSH OK"'
            }
        }

        stage('Deploy') {
            steps {
                sh 'scp -i /var/lib/jenkins/.ssh/id_rsa main laborant@target:~'
            }
        }
    }
}
