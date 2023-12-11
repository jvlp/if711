#!/bin/bash

for i in {1..100}; do
  echo "Run $i"
  echo "Run $i" >> e2_results.txt;
  ./e2 >> e2_results.txt;
done

