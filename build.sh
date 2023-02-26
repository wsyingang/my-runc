#!/bin/bash
go build -o main .

if [ ! -d output ];then
  mkdir output
fi
mv main ./output

