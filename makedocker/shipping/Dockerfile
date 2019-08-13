FROM openjdk:8-jre-alpine
ENV APP_ROOT=/root/servicestage/shipping/
ENV LOG_ROOT=/var/log/shipping/
RUN mkdir -p $APP_ROOT
RUN mkdir -p $LOG_ROOT

COPY ./shipping.jar $APP_ROOT
COPY ./shipping.sh $APP_ROOT


RUN cd $APP_ROOT && chmod -R 770 .
RUN chmod +x /root/servicestage/shipping/shipping.sh

ENTRYPOINT ["/root/servicestage/shipping/shipping.sh"]

