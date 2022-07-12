#!/bin/bash

alacritty -o window.startup_mode=Fullscreen -e ./run.sh &

gol=$!

i3lock -c 00000000 -n

kill $gol
