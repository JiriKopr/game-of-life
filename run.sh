#!/bin/bash

# Alacritty or tcell has some problem with running the script directly from command line
sleep 0.1

dirname=$(dirname $0)

$dirname/gol
