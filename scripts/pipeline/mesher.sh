#!/usr/bin/env bash
cd pwd

#frontend
mkdir build_image_frontend
cp -r ./front-end/* ./build_image_frontend/
cd build_image_frontend
docker build -t frontend:autobuild .
docker tag frontend:autobuild registry.cn-north-1.huaweicloud.com/hwcse/sockshop-frontend:latest
