FROM alpine:3.2
RUN apk update 
RUN apk add bash sed 
RUN rm -rf /var/cache/apk/*
COPY arms.sh /opt/arms.sh
COPY engage.sh /opt/engage.sh
ENTRYPOINT /opt/engage.sh
