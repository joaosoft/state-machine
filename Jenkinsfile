// this will start an executor on a Jenkins agent with the docker label
pipeline {
  agent any
  stages {

    stage('Build') {
      steps {
        deleteDir()
        checkout scm
      }
    }

    stage('run tests') {
      steps {
        make utest
        make itest
      }
    }

    stage('build image') {
      steps {
        make build
      }
    }
  }
}