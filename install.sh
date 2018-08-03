#!/bin/sh

EXE_PATH=`pwd`

echo "$EXE_PATH", $EXE_PATH

#sed "s:{{PATH}}:${EXE_PATH}:g" conf/novelsearch.tpl.conf > conf/novelsearch.conf
#cp conf/novelsearch.conf /etc/supervisor/

sed "s:{{PATH}}:${EXE_PATH}:g" start.tpl.sh > start.sh
chmod +x start.sh

echo $EXE_PATH/start.sh >> /etc/rc.local

