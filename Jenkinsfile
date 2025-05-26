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

        stage('Deploy') {
            steps {
                sshagent(['deploy-key']) {
                    sh 'scp -o StrictHostKeyChecking=no main laborant@target:~'
                }
            }
        }
    }
}
