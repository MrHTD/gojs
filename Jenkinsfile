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
                        "apt install golang-go -y
                        cd /home/devxonic/Projects/go-lang; \
                        ls -la; \

                        npm init; \
                        
                        npm run build; \
                        go version; \

                        go build main.go; \
                        
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
                    "sudo -n systemctl status goweb.service; \
                    sudo -n systemctl start goweb.service; \
                    sudo -n systemctl reload goweb.service; \
                    sudo -n systemctl status goweb.service"                    
                    '''
                }
            }
        }
    }
}
