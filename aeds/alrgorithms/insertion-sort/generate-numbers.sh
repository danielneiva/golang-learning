#!/bin/bash
rm $1
touch $1
shuf -i 0-2147483647 -n $2 >> $1