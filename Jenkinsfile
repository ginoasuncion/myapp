pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        sh 'go build -o main main.go'
      }
    }

    stage('Deploy with Ansible') {
      steps {
        withCredentials([sshUserPrivateKey(credentialsId: 'ansible-key', keyFileVariable: 'SSH_KEY')]) {
          sh '''
            ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i hosts.ini --private-key $SSH_KEY playbook.yml
          '''
        }
      }
    }
  }
}

