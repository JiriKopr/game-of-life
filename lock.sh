#!/bin/bash

dirname=$(dirname $0)

alacritty -o window.startup_mode=Fullscreen -e "$dirname/run.sh" &

gol=$!

i3lock -c 00000000 -n

kill $gol
