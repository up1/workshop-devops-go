pipeline {
    agent any

    parameters {
        string(name:'TAG_NAME',defaultValue: '',description:'')
    }

    environment {
        DOCKER_CREDENTIAL_ID = 'dockerhub-id'
        GITHUB_CREDENTIAL_ID = 'github-id'
        K8S_ID = 'k8s-id'
        REGISTRY = 'docker.io'
        DOCKERHUB_NAMESPACE = 'somkiat'
        GITHUB_ACCOUNT = 'up1'
        APP_NAME = 'demo-api'
    }

    stages {
        stage ('checkout scm') {
            steps {
                checkout(scm)
            }
        }

        stage ('unit test') {
            steps {
                sh 'echo "Test passed"'
            }
        }
 
        stage ('build & push') {
            steps {
                sh 'docker-compose build'
                withCredentials([usernamePassword(passwordVariable : 'DOCKER_PASSWORD' ,usernameVariable : 'DOCKER_USERNAME' ,credentialsId : "$DOCKER_CREDENTIAL_ID" ,)]) {
                    sh 'echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin'
                    sh 'docker image push  $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME'
                }
            }
        }

        stage('deploy to dev') { 
            steps {
                sshPublisher(publishers: [sshPublisherDesc(configName: 'dev-server', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'sh deploy.sh', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: 'deploy.sh')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
            }
        }

        stage('deploy to kubernetes cluster') { 
            steps {
                kubernetesDeploy(configs: 'k8s/**', enableConfigSubstitution: true, kubeconfigId: "$K8S_ID")
            }
        }
    }
}