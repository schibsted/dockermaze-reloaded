FROM debian

RUN apt-get update && apt-get install -y --no-install-recommends \
	ruby \
  ruby-sinatra \
  ruby-rest-client \
  ruby-json \
  thin \
  golang \
	&& apt-get clean

RUN groupadd -r head && useradd -r -g head head

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

COPY files/ /opt/
ADD config/ /opt/config/
RUN chmod 755 /opt/server.rb /opt/start.bash
RUN go build -o /opt/head /opt/head.go

EXPOSE 80
EXPOSE 7777

ENTRYPOINT ["/opt/start.bash"]
