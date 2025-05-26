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

        stage('Check SSH Access') {
            steps {
                sh 'ssh -i /var/lib/jenkins/.ssh/id_rsa -o StrictHostKeyChecking=no laborant@target "echo SSH connection successful"'
            }
        }

        stage('Deploy') {
            steps {
                sh 'scp -i /var/lib/jenkins/.ssh/id_rsa -o StrictHostKeyChecking=no main laborant@target:~'
            }
        }
    }
}
