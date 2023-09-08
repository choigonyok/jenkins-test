package sum

func Sum(a, b int) int {
	c := a + b
	return c
}

func Multi(a, b int) int {
	c := a * b
	return c
}

// go plugin 설치
// Tool에서 golang 버전 적용
//sudo snap install openjdk

// pipeline {
//     agent any
//     tools {
//         go 'golang'
//     }

//     stages {
//         stage('Checkout') {
//             steps {
//                 git branch: 'main',
//                     credentialsId: 'test',
//                     url: 'https://github.com/choigonyok/jenkins-test.git'
//             }
//         }

//         stage('Test') {
//             steps {
//                 script {
//                     echo 'Running test'
//                     COVERAGE = sh (script: "go test -cover ./sum | tr -d '%' | awk -F ' ' '{print \$5}'", returnStdout: true).trim()
//                     if (COVERAGE >= "80") {
//                         echo "PASSED COVERAGE TEST: ${COVERAGE}%"
//                     } else {
//                         echo "YOUR TEST CODE COVERAGE IS ${COVERAGE}%"
//                         error('IT HAS TO LARGE THAN 80%, SO FAILED')
//                     }
//                 }
//             }
//         }
//     }
// }

// sudo apt update
// sudo apt install openjdk-17-jre
// java -version
// openjdk version "17.0.7" 2023-04-18
// OpenJDK Runtime Environment (build 17.0.7+7-Debian-1deb11u1)
// OpenJDK 64-Bit Server VM (build 17.0.7+7-Debian-1deb11u1, mixed mode, sharing)

// curl -fsSL https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key | sudo tee \
//   /usr/share/keyrings/jenkins-keyring.asc > /dev/null
// echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
//   https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
//   /etc/apt/sources.list.d/jenkins.list > /dev/null
// sudo apt-get update
// sudo apt-get install jenkins
