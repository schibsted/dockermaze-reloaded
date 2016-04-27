FROM alpine:3.2
RUN apk update 
RUN apk add go ca-certificates
RUN rm -rf /var/cache/apk/*
RUN update-ca-certificates

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

COPY weapon.go /opt/weapon.go
COPY utils.go /opt/utils.go
COPY reload.sh /opt/reload.sh

RUN go build -o /opt/weapon /opt/*.go

ENTRYPOINT /opt/reload.sh
