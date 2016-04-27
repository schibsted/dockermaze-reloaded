#!/bin/ash
# This sleep is required to
# wait for head to be ready
sleep 3
maze-client &
while true; do
  sleep 86400
done
