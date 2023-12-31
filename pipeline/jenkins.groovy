pipeline {
agent any
    parameters {

        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        choice(name: 'ARCH', choices: ['amd64', 'arm64'], description: 'Pick ARCH')
    }

    environment {
        REPO = 'https://github.com/TaRNaX/kbot'
        BRANCH = 'main'
    }

    stages {

        stage('Example') {
            steps {
                echo "Build for platform ${params.OS}"
                echo "Build for arch: ${params.ARCH}"
            }
        }
        
        stage('clone') {
            steps {
                echo 'CLONE REPOSITORY'
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }
        
        stage('test') {
            steps {
                echo 'TEST EXECUTION STARTED'
                sh 'make test'
            }
        }
        
        stage('build') {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'make build TARGETARCH=${ARCH} TARGETOS=${OS}'
            }
        }
        
        stage('image') {
            steps {
                script{
                    echo 'BUILD EXECUTION STARTED'
                    sh 'make image TARGETARCH=${ARCH} TARGETOS=${OS}' 
                }
            }
        }  
        
        stage('push') {
            steps {
                script{
                    docker.withRegistry('', 'dockerhub') {
                        sh 'make push TARGETARCH=${ARCH} TARGETOS=${OS}' 
                    }
                }
            }
        }
    }
}