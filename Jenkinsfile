pipeline {
    agent any
    environment{
        SSH_USER = 'devxonic'
        SSH_HOST = '192.168.100.14'
    }
    stages {
        stage("Start") {
            steps {
                echo "Go-Lang Pipeline Execution Started."
            }
        }
        stage("Git Pull") {
            steps {
                sshagent(['ssh']) {
                    echo "Pulling latest code from Git repository..."
                    sh '''
                        ssh -o StrictHostKeyChecking=no $SSH_USER@$SSH_HOST \
                        "cd /home/devxonic/Projects/go-lang;
                        git pull origin main || echo 'Failed to pull latest code. Ensure the repository is properly set up.'"
                    '''
                }
            }
        }
        stage("SSH") {
            steps {
                    sshagent(['ssh']){
                    echo "Connecting to machine..."
                    sh '''
                        ssh -o StrictHostKeyChecking=no $SSH_USER@$SSH_HOST \
                        "export SUDO_ASKPASS=/tmp/mypass.sh;
                        sudo -A apt update;
                        sudo -A apt install golang-go -y;

                        cd /home/devxonic/Projects/go-lang;
                        ls -la;

                        npm init -y

                       # npm run build;
                        
                        go version;

                        go build main.go;
                        
                        ls -la"
                    '''
                }
            }
        }
        stage("Check Service Status") {
            steps {
                sshagent(credentials: ['ssh']) {
                    sh '''
                    ssh -o StrictHostKeyChecking=no $SSH_USER@$SSH_HOST \
                    "export SUDO_ASKPASS=/tmp/mypass.sh;
                    sudo -A systemctl status goweb.service;
                    
                    sudo -A systemctl restart goweb.service;
                    
                    sudo -A systemctl status goweb.service"                    
                    '''
                }
            }
        }
        stage("End") {
            steps {
                script {
                    if (currentBuild.result == null || currentBuild.result == 'SUCCESS') {
                        echo "Pipeline completed successfully. 🎉"
                    } else {
                        echo "Pipeline encountered errors. Please check the logs. ❌"
                    }
                }
            }
        }
    }
}
