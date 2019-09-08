#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/ginApi/
git pull https://github.com/feiyuWeb/ginApi.git
cd webserver/
./ginApi &