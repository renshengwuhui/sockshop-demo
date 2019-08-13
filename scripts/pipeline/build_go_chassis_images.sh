#!/bin/bash

set -e
set -x
BASEDIR=$PWD
WORKDIR=$BASEDIR/src/github.com/ServiceComb
CATALOGUE=$BASEDIR/cataloguedir
PAYMENT=$BASEDIR/paymentdir
USER=$BASEDIR/userdir
set +e
rm -rf $CATALOGUE $PAYMENT $USER $BASEDIR/src/
set -e
mkdir -p $WORKDIR
mkdir -p $CATALOGUE
mkdir -p $CATALOGUE/catalog
mkdir -p $PAYMENT
mkdir -p $PAYMENT/payment
mkdir -p $USER
mkdir -p $USER/user

cd $WORKDIR
set +e
rm -rf go-chassis
set -e
git clone https://github.com/ServiceComb/go-chassis.git
cd go-chassis
export CHECKOUT_VERSION=$go_chassis_version
if [ $CHECKOUT_VERSION == "latest" ]; then
    echo "using latest code"
else
    git checkout $CHECKOUT_VERSION
fi
gvt restore
echo "gvt restore success"

echo "catalogue started"

cd $BASEDIR
cp -r catalogue $WORKDIR/go-chassis/examples/

cd $WORKDIR/go-chassis/examples/catalogue/cmd/cataloguesvc/server/
export GOPATH=$BASEDIR
go get github.com/go-sql-driver/mysql
go get github.com/jmoiron/sqlx
CGO_ENABLED=0 GO_EXTLINK_ENABLED=0 go build --ldflags '-s -w -extldflags "-statis"' -a -o "catalogue"

cp -r conf catalogue $WORKDIR/go-chassis/examples/catalogue/images $CATALOGUE/catalog
cp -r $WORKDIR/go-chassis/examples/catalogue/start.sh $WORKDIR/go-chassis/examples/catalogue/Dockerfile $CATALOGUE/

cd $CATALOGUE/

IMAGE=catalogue
TAG=latest

docker build -t $IMAGE:$TAG .
docker save -o $IMAGE:$TAG.tar $IMAGE:$TAG

docker tag $IMAGE:$TAG registry.cn-north-1.huaweicloud.com/hwcse/sockshop-catalogue:$TAG
echo "build images success for catalogue"

echo "**********************"

echo "payment started"

cd $BASEDIR

cp -r payment $WORKDIR/go-chassis/examples/

cd $WORKDIR/go-chassis/examples/payment/cmd/paymentsvc/
export GOPATH=$BASEDIR
CGO_ENABLED=0 GO_EXTLINK_ENABLED=0 go build --ldflags '-s -w -extldflags "-statis"' -a -o "payment"

cp -r conf payment $PAYMENT/payment
cp -r $WORKDIR/go-chassis/examples/payment/start.sh $WORKDIR/go-chassis/examples/payment/Dockerfile $PAYMENT/

cd $PAYMENT/

IMAGE=payment
TAG=latest

docker build -t $IMAGE:$TAG .
docker save -o $IMAGE:$TAG.tar $IMAGE:$TAG

docker tag $IMAGE:$TAG registry.cn-north-1.huaweicloud.com/hwcse/sockshop-payment:$TAG

echo "build images success for payment"

echo "**********************"

echo "user started"

cd $BASEDIR

cp -r user $WORKDIR/go-chassis/examples/

cd $WORKDIR/go-chassis/examples/user/
export GOPATH=$BASEDIR
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
CGO_ENABLED=0 GO_EXTLINK_ENABLED=0 go build --ldflags '-s -w -extldflags "-statis"' -a -o "user"

cp -r conf user $USER/user
cp -r $WORKDIR/go-chassis/examples/user/start.sh $WORKDIR/go-chassis/examples/user/Dockerfile $USER/

cd $USER/

IMAGE=user
TAG=latest

docker build -t $IMAGE:$TAG .
docker save -o $IMAGE:$TAG.tar $IMAGE:$TAG

docker tag $IMAGE:$TAG registry.cn-north-1.huaweicloud.com/hwcse/sockshop-user:$TAG

echo "build image success for user"

