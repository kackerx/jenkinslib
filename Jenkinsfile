#!Groovy
@Library('jenkinslib@master') _

def tools = new org.devops.tools()

String workspace = '/opt/jenkins/workspace'

// Pipeline
pipeline {
    agent {
        node {
            label 'build'
            customWorkspace "${workspace}"
        }
    }

    options {
        timestamps()
        skipDefaultCheckout()
        disableConcurrentBuilds()
        timeout(time: 1, unit: 'HOURS')
    }

    stages {
        stage('GetCode') {
            steps {
                timeout(time: 5, unit: 'MINUTES') {
                    script {
                        println('获取代码')
                        tools.PringMsg('get 代码')

                        // input
                        input id: 'Id10', message: '继续吗', ok: '继续吗', parameters: [choice(choices: ['a', 'b'], description: '呵呵', name: 'param1')], submitter: 'kacker,admin'

                    }
                }
            }
        }
        stage('Build') {
            // 是否执行该stage
            when {
                expression {
                    return env.name == 'kacker'
                }
            }
            steps {
                timeout(time: 20, unit: 'MINUTES') {
                    script {
                        println('应用打包')

                        // tool
                        mvnHome = tool "m2"
                        println(mvnHome)

                        sh "${mvnHome}/bin/mvn --version"
                    }
                }
            }
        }
        stage('CodeScan') {
            steps {
                timeout(time: 30, unit: 'MINUTES') {
                    script {
                        println('代码扫描')
                    }
                }
            }
        }
    }

    post {
        always {
            script {
                println('always ouput')
            }
        }

        success {
            script {
                currentBuild.description = '构建成功'
            }
        }

        failure {
            script {
                currentBuild.description = '构建失败'
            }
        }

        aborted {
            script {
                currentBuild.description = '构建终止'
            }
        }
    }
}

