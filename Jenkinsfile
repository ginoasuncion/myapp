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
        sh '''
          chmod 600 id_rsa
          ansible-playbook -i hosts.ini playbook.yml
        '''
      }
    }
  }
}

