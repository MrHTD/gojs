pipeline {
    agent any
    environment{
        SSH_USER = 'devxonic'
        SSH_HOST = '192.168.100.14'
    }
    stages {
        stage("SSH") {
            steps {
                    sshagent(['ssh']){
                    echo "Connecting to machine..."
                    sh '''
                        ssh -o StrictHostKeyChecking=no $SSH_USER@$SSH_HOST \
                        "apt install golang-go -y;
                        cd /home/devxonic/Projects/go-lang;
                        ls -la;

                        npm init -y;
                        
                        npm run build;
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
                    sudo -A systemctl start goweb.service;
                    sudo -A systemctl reload goweb.service;
                    sudo -A systemctl status goweb.service"                    
                    '''
                }
            }
        }
    }
}
