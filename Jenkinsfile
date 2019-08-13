pipeline {
    agent any

    stages {
        stage('Build Mesher-Microservice Image') {
            steps {
                echo 'Building Mesher-Microservices Images'
                echo params.something
                sh 'bash +x ../../jenkinsScript/dockerlogin.sh'
                sh 'bash -x scripts/pipeline/mesher.sh'
            }
        }
        stage('Build Go-Microservice Image') {
             steps {
                  echo 'Building '
                  sh 'bash +x ../../jenkinsScript/dockerlogin.sh'
                  sh 'bash scripts/pipeline/build_go_chassis_images.sh'
             }
        }
        stage('Build Java-Microservice Image') {
            steps {
                echo 'Pushing Java-Microservice Images to SWR'
                sh 'bash +x ../../jenkinsScript/dockerlogin.sh'
                sh 'bash -x scripts/pipeline/java-chassis.sh'
            }
        }
        stage('Push Images to SWR') {
             steps {
                  echo 'Pushing Images to SWR'
                  sh 'bash +x ../../jenkinsScript/dockerlogin.sh'
                  sh 'bash -x scripts/pipeline/push_image_swr.sh'
             }
        }
    }
}
