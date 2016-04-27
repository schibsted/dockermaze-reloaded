#!/bin/bash

sleep 1
/home/radio/bin/radio_comm.rb
while true ; do
 # if ! pidof -X radio_comm.rb &> /dev/null ; then
 #   /home/radio/bin/radio_comm.rb
 # fi
 sleep 100
done