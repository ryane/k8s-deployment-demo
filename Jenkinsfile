pipeline {
    agent any

    podTemplate(label: 'golang', containers: [
            containerTemplate(name: 'golang', image: 'golang:1.9.2', ttyEnabled: true, command: 'cat')
        ]) {
        node('golang') {

            stages {
                stage('Build') {
                    steps {
                        echo 'Build phase'
                    }
                }
                stage('Test') {
                    container('golang') {
                        steps {
                            echo 'Test phase'
                            sh 'go test ./...'
                        }
                    }
                }
                stage('Deploy') {
                    steps {
                        echo 'Deploy phase'
                    }
                }
            }
        }
    }
}
