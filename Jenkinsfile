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
                        "export SUDO_ASKPASS=/tmp/mypass.sh;
                        sudo -A apt update;
                        sudo -A apt install golang-go -y;

                        cd /home/devxonic/Projects/go-lang;
                        ls -la;

                        # Initialize npm (if needed)
                        if [ ! package.json ]; then
                            npm init -y
                            jq '.scripts.build="echo No build script defined"' package.json > temp.json && mv temp.json package.json
                        fi

                        # Run npm build (if applicable)
                        npm run build || echo "Skipping npm build as no meaningful script is defined;

                        # Build the Go application
                        if [ -f main.go ]; then
                            go build -o goweb main.go
                        else
                            echo "main.go not found!"
                            exit 1
                        fi
                        
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
