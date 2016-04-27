FROM alpine:3.2
RUN apk update 
RUN apk add git g++ make
WORKDIR /tmp/
RUN git clone https://github.com/official-stockfish/Stockfish.git
RUN make -C /tmp/Stockfish/src build ARCH=x86-64
RUN make -C /tmp/Stockfish/src install
RUN rm -rf /tmp/Stockfish
RUN apk del git make
RUN rm -rf /var/cache/apk/*
COPY stockfish.sh /opt/stockfish.sh
EXPOSE 9999
ENTRYPOINT /opt/stockfish.sh
