podTemplate(label: 'golang', containers: [
        containerTemplate(name: 'golang', image: 'golang:1.9.2', ttyEnabled: true, command: 'cat')
    ]) {
    node('golang') {
        stage('Build') {
            sh("echo 'Build phase'")
            sh("pwd")
            sh("ls -al")
        }
        stage('Test') {
            container('golang') {
                sh("echo 'Build phase'")
                sh("pwd")
                sh("ls -al")
                sh("go test ./...")
            }
        }
        stage('Deploy') {
            sh("echo 'Build phase'")
            sh("pwd")
            sh("ls -al")
        }
    }
}
