#!/bin/bash

for i in {1..59}; do
    echo $i
    ./flume-client -category kafka
done
