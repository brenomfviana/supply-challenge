#!/bin/bash
cd "phase$1" &&
go build &&
echo "Executando..." &&
start=`date +%s%N`/1000000 &&
./"phase$1" $2 $3 &&
end=`date +%s%N`/1000000 &&
runtime=$(((end-start))) &&
echo "Time: $runtime milliseconds" &&
rm -f "phase$1" || rm -f "phase$1.exe"
