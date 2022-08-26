#!/bin/bash

# ls -v — natural sort of (version) numbers within text
files=`ls -v | grep -i "^bandit*"`
last_file=`echo "$files" | tail -1`

if [ -f "$1" ]; then
	last_file=$1
else
	echo "$1 does not exist."
	exit 1
fi


echo "ssh to $last_file@bandit.labs.overthewire.org"
sshpass -p `cat "$last_file"` ssh $last_file@bandit.labs.overthewire.org -p 2220
 
