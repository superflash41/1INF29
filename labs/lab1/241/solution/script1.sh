#!/bin/bash
# superflash41
num=$1
last=$2
	
while [ $num -le $last ]; do 
	if [[ $num == *$3* ]]; then
		echo $num
	fi	
	num=$((num+1))
done
