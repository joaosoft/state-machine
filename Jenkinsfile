// this will start an executor on a Jenkins agent with the docker label
pipeline {
  agent any
  stages {
    stage('run tests') {
      steps {
        sh 'bash make utest'
        sh 'bash make itest'
      }
    }
    stage('build image') {
      steps {
        sh 'bash make build'
      }
    }
  }
}