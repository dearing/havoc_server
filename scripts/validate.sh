#!/bin/bash

result=$(curl -s http://localhost:8080/wtf/)

if [[ "$result" =~ "404" ]]; then
    exit 0
else
    exit 1
fi
