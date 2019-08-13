FROM openjdk:8-jre-alpine
ENV APP_ROOT=/root/servicestage/carts/
ENV LOG_ROOT=/var/log/carts/
RUN mkdir -p $APP_ROOT
RUN mkdir -p $LOG_ROOT

COPY ./carts.jar $APP_ROOT
COPY ./carts.sh $APP_ROOT


RUN cd $APP_ROOT && chmod -R 770 .
RUN chmod +x /root/servicestage/carts/carts.sh

ENTRYPOINT ["/root/servicestage/carts/carts.sh"]

