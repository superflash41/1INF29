#!/bin/bash

while IFS= read -r line; do
	words=($line)
	first="${words[0]}"
	unset 'words[0]'
	# words="${words[@]^}"  
	words+=("$first")
	new_ln=$(IFS=' '; echo "${words[@]^}")
	echo $new_ln >> newdata
done < data
mv newdata data