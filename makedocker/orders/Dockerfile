FROM openjdk:8-jre-alpine
ENV APP_ROOT=/root/servicestage/orders/
ENV LOG_ROOT=/var/log/orders/
RUN mkdir -p $APP_ROOT
RUN mkdir -p $LOG_ROOT

COPY ./orders.jar $APP_ROOT
COPY ./orders.sh $APP_ROOT


RUN cd $APP_ROOT && chmod -R 770 .
RUN chmod +x /root/servicestage/orders/orders.sh

ENTRYPOINT ["/root/servicestage/orders/orders.sh"]

