#!/bin/bash

pid="$(sudo lsof -i :5432 | awk '{print $2}' | head -2 | tail -1)"

while [[ $pid -gt 0 ]]
do
    sudo kill -9 $pid
    pid="$(sudo lsof -i :5432 | awk '{print $2}' | head -2 | tail -1)"
done

pid="$(sudo lsof -i :8080 | awk '{print $2}' | head -2 | tail -1)"
while [[ $pid -gt 0 ]]
do
    echo "$(expr length "$pid")"
    sudo kill -9 $pid
    pid="$(sudo lsof -i :8080 | awk '{print $2}' | head -2 | tail -1)"
done