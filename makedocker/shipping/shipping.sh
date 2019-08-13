#!/bin/sh
cd /root/servicestage/shipping/
chmod -R 770 /root/servicestage/shipping
cp /etc/resolv.conf /tmp
sed -i s/"^.*search.*$"/"search"/g /tmp/resolv.conf
cat /tmp/resolv.conf > /etc/resolv.conf
java -jar shipping.jar

~                                                                                                                                                                        
~       
