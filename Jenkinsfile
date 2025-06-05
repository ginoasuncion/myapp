pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        sh '[ -f main ] && chmod +x main'
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

