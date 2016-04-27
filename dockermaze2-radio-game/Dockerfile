FROM debian

RUN apt-get update && apt-get install -y --no-install-recommends \
	gcc-multilib \
	libc6-dev-i386 \
	ruby \
  ruby-rest-client \
	&& apt-get clean

RUN groupadd -r radio && useradd -r -g radio radio

COPY files/ /home/radio/bin/

RUN chmod 755 /home/radio/bin/radio_comm.rb /home/radio/bin/start.bash /home/radio/bin/compile.sh
RUN chmod 644 /home/radio/bin/create_mac.s

RUN /home/radio/bin/compile.sh
RUN chown -R radio:radio /home/radio/bin/*
RUN chmod +x /usr/local/bin/create_mac

ENTRYPOINT ["/home/radio/bin/start.bash"]

USER radio