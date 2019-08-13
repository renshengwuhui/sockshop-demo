#!/bin/sh
cd /root/servicestage/orders/
chmod -R 770 /root/servicestage/orders
cp /etc/resolv.conf /tmp
sed -i s/"^.*search.*$"/"search"/g /tmp/resolv.conf
cat /tmp/resolv.conf > /etc/resolv.conf
java -jar orders.jar
