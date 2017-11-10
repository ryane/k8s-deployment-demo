pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Build phase'
            }
        }
        stage('Test') {
            steps {
                echo 'Test phase'
                sh 'go test ./...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploy phase'
            }
        }
    }
}
