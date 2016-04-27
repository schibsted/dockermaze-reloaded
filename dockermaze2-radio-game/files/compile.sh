#!/bin/bash

exec gcc -m32 /home/radio/bin/create_mac.s -masm=intel -fno-stack-protector -O0 -o /usr/local/bin/create_mac