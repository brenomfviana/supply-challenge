#!/bin/bash

# echo "Executando..." &&
# go run "Phase $1/supply-challenge.go"
cd "Phase_$1" &&
go build &&
echo "Executando..." &&
start=`date +%s%N`/1000000 &&
./"Phase_$1" $2 $3 &&
end=`date +%s%N`/1000000 &&
runtime=$(((end-start))) &&
echo "Time: $runtime milliseconds" &&
rm -f "Phase_$1" || rm -f "Phase_$1.exe"