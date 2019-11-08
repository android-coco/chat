#!/bin/sh
 rm -rf ./release
 mkdir  release
 # mac
 # make
 # linux
 make linux
 chmod +x ./bin/chat_server
 cp -r config ./release/
 rm -r ./release/config/config.demo.yaml
 rm -rf ./release/config/config.go
 rm -rf ./bin/mnt
 cp -r bin ./release/
 cp -r ./static ./release/
 cp -r ./view ./release/