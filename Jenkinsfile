pipeline {
    agent any
    environment{
        SSH_USER = 'devxonic'
        SSH_HOST = '192.168.100.14'
    }
    stages {
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

                        npm run build;

                        # Build the Go application
                        if [ main.go ]; then
                            go build -o goweb main.go
                        else
                            echo "main.go not found!"
                            exit 1
                        fi
                        
                        go version;

                        # go build main.go;
                        
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
                    sudo -A systemctl status goweb.service || true;
                    sudo -A systemctl start goweb.service;
                    sudo -A systemctl reload goweb.service;
                    sudo -A systemctl status goweb.service"                    
                    '''
                }
            }
        }
    }
}
