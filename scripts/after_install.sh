#!/bin/bash

cd $HOME
echo "export GOPATH=$HOME/golang" >> $HOME/.bashrc
echo "export PATH=$PATH:$GOPATH/bin" >> $HOME/.bashrc
source $HOME/.bashrc
cd $HOME/golang/src/github.com/dearing/havoc_server
go build
go install
