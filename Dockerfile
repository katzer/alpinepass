FROM golang:alpine
MAINTAINER Stefan Härtel "haertel@appplant.de"

ENV APP_HOME /code
RUN mkdir $APP_HOME
WORKDIR $APP_HOME
