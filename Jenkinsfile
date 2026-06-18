pipeline {
    agent any

    environment {

        REGISTRY = "nadzallafinalcloudacr.azurecr.io"

        AUTH_IMAGE = "${REGISTRY}/auth-service:latest"
        ORDER_IMAGE = "${REGISTRY}/order-service:latest"
        PAYMENT_IMAGE = "${REGISTRY}/payment-service:latest"
        DELIVERY_IMAGE = "${REGISTRY}/delivery-service:latest"
        SHIPMENT_IMAGE = "${REGISTRY}/shipment-service:latest"
        PICKUP_IMAGE = "${REGISTRY}/pickup-service:latest"
        WAREHOUSE_IMAGE = "${REGISTRY}/warehouse-service:latest"
        TRACKING_IMAGE = "${REGISTRY}/tracking-service:latest"
        NOTIFICATION_IMAGE = "${REGISTRY}/notification-service:latest"
        GATEWAY_IMAGE = "${REGISTRY}/api-gateway:latest"

        TEST_NETWORK = "test-net"
    }

    stages {
        // CHECKOUT
        stage('Checkout Repo') {
            steps {
                deleteDir()
                git branch: 'main', url: 'https://github.com/nadzallad/Final_Cloud.git'
            }
        }
        
        // UNIT TEST
        stage('Unit Test') {
            steps {

                dir('backend/auth-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test ./...'
                    }
                }

                dir('backend/payment-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test -v -run TestValidatePayment ./...'
                    }
                }

                dir('backend/order-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test -short ./...'
                    }
                }

                dir('backend/delivery-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test ./...'
                    }
                }

                dir('backend/shipment-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test ./...'
                    }
                }

                dir('backend/pickup-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test ./...'
                    }
                }

                dir('backend/warehouse-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test ./...'
                    }
                }

                dir('backend/tracking-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test -short ./...'
                    }
                }

                dir('backend/notification-service') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh '''
                        go test -short ./... \
                        -run "TestValidateNotification"
                        '''
                    }
                }

                dir('backend/api_gateway') {
                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go test ./...'
                    }
                }
            }
        }
        
        // LINT / VET
        stage('Lint / Vet') {
            steps {

                dir('backend/auth-service') {
                    sh 'go vet ./...'
                }

                dir('backend/payment-service') {
                    sh 'go vet ./...'
                }

                dir('backend/order-service') {
                    sh 'go vet ./...'
                }

                dir('backend/delivery-service') {
                    sh 'go vet ./...'
                }

                dir('backend/shipment-service') {
                    sh 'go vet ./...'
                }

                dir('backend/pickup-service') {
                    sh 'go vet ./...'
                }

                dir('backend/warehouse-service') {
                    sh 'go vet ./...'
                }

                dir('backend/tracking-service') {
                    sh 'go vet ./...'
                }

                dir('backend/notification-service') {
                    sh 'go vet ./...'
                }

                dir('backend/api_gateway') {
                    sh 'go vet ./...'
                }
            }
        }

        // BUILD IMAGE
        stage('Build Image') {
            steps {
                sh '''
                docker build -t $AUTH_IMAGE ./backend/auth-service
                docker build -t $PAYMENT_IMAGE ./backend/payment-service
                docker build -t $ORDER_IMAGE ./backend/order-service
                docker build -t $DELIVERY_IMAGE ./backend/delivery-service
                docker build -t $SHIPMENT_IMAGE ./backend/shipment-service
                docker build -t $PICKUP_IMAGE ./backend/pickup-service
                docker build -t $WAREHOUSE_IMAGE ./backend/warehouse-service
                docker build -t $TRACKING_IMAGE ./backend/tracking-service
                docker build -t $NOTIFICATION_IMAGE ./backend/notification-service
                docker build -t $GATEWAY_IMAGE ./backend/api_gateway
                '''
            }
        }


        // PUSH IMAGE
        stage('Push Image') {
                steps {
                    withCredentials([
                        usernamePassword(
                            credentialsId: 'azure-acr',
                            usernameVariable: 'USERNAME',
                            passwordVariable: 'PASSWORD'
                        )
                    ]) {

                        sh '''
                        echo "$PASSWORD" | docker login nadzallafinalcloudacr.azurecr.io \
                        -u "$USERNAME" \
                        --password-stdin

                        docker push $AUTH_IMAGE
                        docker push $ORDER_IMAGE
                        docker push $PAYMENT_IMAGE
                        docker push $DELIVERY_IMAGE
                        docker push $SHIPMENT_IMAGE
                        docker push $PICKUP_IMAGE
                        docker push $WAREHOUSE_IMAGE
                        docker push $TRACKING_IMAGE
                        docker push $NOTIFICATION_IMAGE
                        docker push $GATEWAY_IMAGE
                        '''
                    }
                }
            }

        stage('Azure Login') {
            steps {
                withCredentials([
                    string(credentialsId: 'AZURE_CLIENT_ID', variable: 'AZURE_CLIENT_ID'),
                    string(credentialsId: 'AZURE_CLIENT_SECRET', variable: 'AZURE_CLIENT_SECRET'),
                    string(credentialsId: 'AZURE_TENANT_ID', variable: 'AZURE_TENANT_ID')
                ]) {
                    sh '''
                    az login --service-principal \
                    -u $AZURE_CLIENT_ID \
                    -p $AZURE_CLIENT_SECRET \
                    --tenant $AZURE_TENANT_ID
                    '''
                }
            }
        }
        // DEPLOY
        stage('Deploy AKS') {
            steps {
                sh '''
                az aks get-credentials \
                --resource-group FinalCloudRG \
                --name finalcloud-aks \
                --overwrite-existing

                kubectl apply -f k8s/

                kubectl rollout status deployment/auth-service
                kubectl rollout status deployment/order-service
                kubectl rollout status deployment/payment-service
                kubectl rollout status deployment/delivery-service
                kubectl rollout status deployment/shipment-service
                kubectl rollout status deployment/pickup-service
                kubectl rollout status deployment/warehouse-service
                kubectl rollout status deployment/tracking-service
                kubectl rollout status deployment/notification-service
                kubectl rollout status deployment/api-gateway
                '''
            }
        }

        stage('Verify') {
            steps {
                sh '''
                kubectl get pods
                kubectl get svc
                '''
            }
        }

    }
}