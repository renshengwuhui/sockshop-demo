#config example
#TARGET_VERSION=0.0.1                                                    ---------huawei cloud images repo save version.
ORIGIN_VERSION=1.2.3-SNAPSHOT                                            #---------images version been made "mvn -Psudo docker".
#TENANT_NAME=xxxxxxxxxxx                                                 ---------huawei cloud tenant name.
#REPO_ADDRESS=registry.cn-north-1.hwclouds.com                           #---------huawei cloud images repo address.
REPO_ADDRESS=swr.cn-north-1.myhuaweicloud.com                                          #---------huawei cloud images repo address.
#USER_NAME=xxxxx                                                         ---------user name: login huawei cloud images repo.
#PW=xxxxxxx                                                              ---------paasword: login huawei cloud images repo.
#CUSTOMER_REPO_NAME=acmeair-customer                                     ---------customer repo name ,created by huawei cloud. 
#BOOKING_REPO_NAME=acmeair-booking                                       ---------booking repo name ,created by huawei cloud. 
#WEBSITE_REPO_NAME=acmeair-website                                       ---------website repo name ,created by huawei cloud. 


export targetversion=1.1.0
#java sudo docker build
cd /opt/sockshop-demo/
#mvn clean  install   -Dmaven.test.skip=true  -settings=/home/root1/preethi/dcap/sockshop1/tank/acmeair_settings.xml
mvn install:install-file -Dfile=libs/java-sdk-core-2.0.1.jar -DgroupId=java-sdk-core -DartifactId=java-sdk-core -Dversion=2.0.1 -Dpackaging=jar
mvn clean install -Dmaven.test.skip=true
cd /opt/sockshop-demo/makedocker

cp /opt/sockshop-demo/makedocker/carts/Dockerfile  /opt/sockshop-demo/carts/target/
cp /opt/sockshop-demo/makedocker/carts/carts.sh  /opt/sockshop-demo/carts/target/
cp /opt/sockshop-demo/makedocker/orders/Dockerfile  /opt/sockshop-demo/orders/target/
cp /opt/sockshop-demo/makedocker/orders/orders.sh /opt/sockshop-demo/orders/target/
cp /opt/sockshop-demo/makedocker/shipping/Dockerfile  /opt/sockshop-demo/shipping/target/
cp /opt/sockshop-demo/makedocker/shipping/shipping.sh /opt/sockshop-demo/shipping/target/


cd /opt/sockshop-demo/carts/target/
sudo docker build -t sockshop-carts-service:$ORIGIN_VERSION  . 
cd /opt/sockshop-demo/orders/target/
sudo docker build -t sockshop-orders-service:$ORIGIN_VERSION  . 
cd /opt/sockshop-demo/shipping/target/
sudo docker build -t sockshop-shipping-service:$ORIGIN_VERSION  . 
#cd /home/root1/preethi/dcap/sockshop1/sockshop-demo/queue-master/target/
#sudo docker build -t sockshop-queuemaster-service:$ORIGIN_VERSION  . 

#front-end sudo docker build
#cd /home/root1/preethi/dcap/sockshop1/sockshop-demo/front-end/
#sudo docker build -t sockshop-frontend-service:$ORIGIN_VERSION .

#Go microservice build
#cd /home/root1/preethi/dcap/sockshop1/sockshop-demo/payment/
#sudo docker build -t sockshop-payment-service:$ORIGIN_VERSION .
#cd /home/root1/preethi/dcap/sockshop1/sockshop-demo/user/
#sudo docker build -t sockshop-user-service:$ORIGIN_VERSION .
#cd /home/root1/preethi/dcap/sockshop1/sockshop-demo/catalogue/
#sudo docker build -t sockshop-catalogue-service:$ORIGIN_VERSION .

#sudo docker tag/push
#sudo docker tag sockshop-frontend-service:$ORIGIN_VERSION  ${REPO_ADDRESS}/hwcse/sockshop-frontend:$targetversion
#sudo docker tag sockshop-payment-service:$ORIGIN_VERSION  ${REPO_ADDRESS}/hwcse/sockshop-payment:$targetversion
#sudo docker tag sockshop-user-service:$ORIGIN_VERSION  ${REPO_ADDRESS}/hwcse/sockshop-user:$targetversion
#sudo docker tag sockshop-catalogue-service:$ORIGIN_VERSION  ${REPO_ADDRESS}/hwcse/sockshop-cat:$targetversion
sudo docker tag sockshop-carts-service:$ORIGIN_VERSION  ${REPO_ADDRESS}/hwcse/sockshop-carts:$targetversion
sudo docker tag sockshop-orders-service:$ORIGIN_VERSION  ${REPO_ADDRESS}/hwcse/sockshop-orders:$targetversion
sudo docker tag sockshop-shipping-service:$ORIGIN_VERSION  ${REPO_ADDRESS}/hwcse/sockshop-shipping:$targetversion
#sudo docker tag sockshop-queuemaster-service:$ORIGIN_VERSION  ${REPO_ADDRESS}/hwcse/sockshop-queuemaster:$targetversion

#sudo docker login -u cn-north-1@89WO1KDCRPKDMSGK4KQH -p 21071575be7dbfbc2cfc876141b422d8212509f50ab44346b880e72126565691 ${REPO_ADDRESS}
#sudo docker login -u cn-north-1@CEOCLCHHQOZ602DRFQ5L -p 882d640dce0eb45cf833e7aad7f10aa8e5e22fe32ee2cc6fc7b2fd421f37f792 ${REPO_ADDRESS}
#sudo docker login -u cn-north-1@EDV0PE1TVESJ0Z9SFMV0 -p 68f44f4122fd4cbe5ba4e30909a56854d7836adc7465a60316fd2b3b9e60bbe3 ${REPO_ADDRESS}
sudo docker login -u cn-north-1@mRc8X3Uei9crPijr3tTP -p decec0ffb10c789126280e81ff638fc6f83fbbb405b336b96aaa5a3ac841bb9a swr.cn-north-1.myhuaweicloud.com

#sudo docker push ${REPO_ADDRESS}/hwcse/sockshop-frontend:$targetversion
#sudo docker push ${REPO_ADDRESS}/hwcse/sockshop-payment:$targetversion
#sudo docker push ${REPO_ADDRESS}/hwcse/sockshop-user:$targetversion
#sudo docker push ${REPO_ADDRESS}/hwcse/sockshop-cat:$targetversion
sudo docker push ${REPO_ADDRESS}/hwcse/sockshop-carts:$targetversion
sudo docker push ${REPO_ADDRESS}/hwcse/sockshop-orders:$targetversion
sudo docker push ${REPO_ADDRESS}/hwcse/sockshop-shipping:$targetversion
#sudo docker push ${REPO_ADDRESS}/hwcse/sockshop-queuemaster:$targetversion



