#!/bin/bash

i=1
for file in ./content/*; do
  curl -X POST http://localhost:8080/images/upload \
    -F "title=TestTitle$i" \
    -F "tags=tag$i,tag$((i+1))" \
    -F "payload=@${file};type=application/octet-stream"
  ((i++))
done