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
        // Permission to execute
        sh "chmod +x -R ${env.WORKSPACE}/../${env.JOB_NAME}@script"

        // Call SH
        sh "${env.WORKSPACE}/../${env.JOB_NAME}@script/make itest"
      }
    }

    stage('build image') {
      steps {
        sh 'make build'
      }
    }
  }
}