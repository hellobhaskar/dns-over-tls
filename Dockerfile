FROM ubuntu:latest
LABEL maintainer "hellobhaskar@yahoo.co.in"
RUN  apt-get update -y &&  apt-get install -y golang ca-certificates git
RUN mkdir /app
RUN mkdir /app/gopkgs
ENV GOPATH /app/gopkgs
RUN go get github.com/miekg/dns
COPY srv.go /app/srv.go
RUN cd /app &&  go build srv.go
RUN apt-get remove -y golang git && apt-get autoremove -y
EXPOSE 53/tcp
EXPOSE 53/udp
CMD  ["/app/srv"]