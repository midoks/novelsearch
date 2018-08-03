#!/bin/sh

EXE_PATH=`pwd`

echo "$EXE_PATH", $EXE_PATH
sed "s:{{PATH}}:${EXE_PATH}:g" conf/novelsearch.tpl.conf > conf/novelsearch.conf

cp conf/novelsearch.conf /etc/supervisor/