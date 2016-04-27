#!/bin/ash
sleep 1
/opt/weapon
while true ; do
  if ! pidof weapon &> /dev/null ; then
    if ! stat /opt/weapon &> /dev/null; then
      go build -o /opt/weapon /opt/*.go
      /opt/weapon
    fi
  fi
  sleep 1
done
