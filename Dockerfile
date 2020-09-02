FROM golang:1.14
RUN mkdir /opt/loadtest /opt/loadtest/results_csv /opt/loadtest/results_html
WORKDIR /opt/loadtest
COPY . /opt/loadtest
