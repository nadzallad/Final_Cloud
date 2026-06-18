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
    }

    stages {

        stage('Checkout') {
            steps {
                deleteDir()
                git branch: 'main',
                    url: 'https://github.com/nadzallad/Final_Cloud.git'
            }
        }

        stage('Build Images') {
            steps {
                sh '''
                docker build -t $AUTH_IMAGE ./backend/auth-service
                docker build -t $ORDER_IMAGE ./backend/order-service
                docker build -t $PAYMENT_IMAGE ./backend/payment-service
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

        stage('Push Images') {
            steps {
                withCredentials([
                    usernamePassword(
                        credentialsId: 'azure-acr',
                        usernameVariable: 'ACR_USER',
                        passwordVariable: 'ACR_PASS'
                    )
                ]) {
                    sh '''
                    echo "$ACR_PASS" | docker login $REGISTRY \
                    -u "$ACR_USER" \
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
                    string(credentialsId: 'AZURE_CLIENT_ID', variable: 'CLIENT_ID'),
                    string(credentialsId: 'AZURE_CLIENT_SECRET', variable: 'CLIENT_SECRET'),
                    string(credentialsId: 'AZURE_TENANT_ID', variable: 'TENANT_ID')
                ]) {
                    sh '''
                    az login --service-principal \
                    -u $CLIENT_ID \
                    -p $CLIENT_SECRET \
                    --tenant $TENANT_ID
                    '''
                }
            }
        }

        stage('Deploy AKS') {
            steps {
                sh '''
                az aks get-credentials \
                --resource-group FinalCloudRG \
                --name finalcloud-aks \
                --overwrite-existing

                kubectl apply -f k8s/
                kubectl get pods
                kubectl get svc
                '''
            }
        }
    }

    post {
        success {
            echo 'Deployment to AKS successful'
        }

        failure {
            echo 'Pipeline failed'
        }
    }
}
