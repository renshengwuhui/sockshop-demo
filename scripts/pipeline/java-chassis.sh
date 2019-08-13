#!/usr/bin/env bash
search='<servicecomb.version>0.2.0'
search1='<huaweicloud.version>2.1.11'
replace1='<huaweicloud.version>2.2.31'

export JAVA_CHASSIS_VERSION=$java_chassis_version
replace='<servicecomb.version>'$JAVA_CHASSIS_VERSION
BASEDIR=$PWD

for file in `find -maxdepth 1 -name 'pom.xml'`; do
  grep "$search" $file &> /dev/null
  if [ $? -ne 0 ]; then
    echo "Search string not found in $file!"
  else
    sed -i "s/$search/$replace/" $file
    sed -i "s/$search1/$replace1/" $file
  fi  
done

set -e

mvn clean install -Phuaweicloud -DskipTests
cd carts
chmod +x .

#carts
set +e
rm -rf build_image_carts
set -e
mkdir build_image_carts
cp target/carts.jar ./build_image_carts
cp ../makedocker/carts/Dockerfile build_image_carts
cp ../makedocker/carts/carts.sh build_image_carts
cd build_image_carts
docker build -t carts:autobuild .
docker tag carts:autobuild registry.cn-north-1.huaweicloud.com/hwcse/sockshop-carts:latest
cd ../..

cd orders
#orders
set +e
rm -rf build_image_orders
set -e
mkdir build_image_orders
cp ../orders/target/orders.jar build_image_orders
cp ../makedocker/orders/Dockerfile build_image_orders
cp ../makedocker/orders/orders.sh build_image_orders
cd build_image_orders
docker build -t orders:autobuild .
docker tag orders:autobuild registry.cn-north-1.huaweicloud.com/hwcse/sockshop-orders:latest
cd ../..

cd shipping
#shipping
set +e
rm -rf build_image_shipping
set -e
mkdir build_image_shipping
cp ../shipping/target/shipping.jar build_image_shipping
cp ../makedocker/shipping/Dockerfile build_image_shipping
cp ../makedocker/shipping/shipping.sh build_image_shipping
cd build_image_shipping
docker build -t shipping:autobuild .
docker tag shipping:autobuild registry.cn-north-1.huaweicloud.com/hwcse/sockshop-shipping:latest
cd ..
