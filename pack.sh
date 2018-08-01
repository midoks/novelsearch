#!/bin/sh

bee pack -exs=.go:.DS_Store:.tmp:vendor:logs:pack.sh -be GOOS=darwin
mv novelsearch.tar.gz ../novelsearch.macosx.tar.gz

bee pack -exs=.go:.DS_Store:.tmp:vendor:logs:pack.sh -be GOOS=linux
mv novelsearch.tar.gz ../novelsearch.linux.tar.gz

bee pack -exs=.go:.DS_Store:.tmp:vendor:logs:pack.sh -be GOOS=windows
mv novelsearch.tar.gz ../novelsearch.win64.tar.gz