FROM golang:1.14
RUN mkdir /opt/loadtest
WORKDIR /opt/loadtest
COPY . /opt/loadtest
