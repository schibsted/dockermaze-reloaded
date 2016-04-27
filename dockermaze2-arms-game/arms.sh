#!/bin/bash

# DOCKERBOT ARMS VERIFICATION MODULE
#
# This module verifies that the arms can solve 
# basic problems in coordination with the head
# and the tools present in the robot.
#
# The verification is successful when all
# mate-in-one puzzles are successfully
# solved with any of the winning moves.

set -e
echo "Connecting to the head..."
exec 7<>/dev/tcp/head/7777
echo "Retrieving puzzles..."
while read puzzle <&7; do
  if [[ "$puzzle" == *"EOF"* ]]; then
    echo "End of transfer received."
    echo "EOF" >&7
    head -1 <&7
    break
  fi
  echo "Position: $puzzle"
  echo "Solving for mate in one..."
  echo "position fen $puzzle" >&9
  echo "go depth 1" >&9
  while true; do
    if read line <&9; then
      if [[ "$line" == *"bestmove"* ]]; then
        echo "Sending best move to the head..."
        echo "$line" | sed -n 's/.*bestmove \(.*\)/\1/p' >&7
        break
      fi
    fi
  done
done
